package clickhouse

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/kobsio/kobs/pkg/api/clusters"
	"github.com/kobsio/kobs/pkg/api/middleware/errresponse"
	"github.com/kobsio/kobs/pkg/api/plugins/plugin"
	"github.com/kobsio/kobs/plugins/clickhouse/pkg/instance"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

// Route is the route under which the plugin should be registered in our router for the rest api.
const Route = "/clickhouse"

var (
	log = logrus.WithFields(logrus.Fields{"package": "clickhouse"})
)

// Config is the structure of the configuration for the clickhouse plugin.
type Config []instance.Config

// Router implements the router for the resources plugin, which can be registered in the router for our rest api.
type Router struct {
	*chi.Mux
	clusters  *clusters.Clusters
	instances []*instance.Instance
}

func (router *Router) getInstance(name string) *instance.Instance {
	for _, i := range router.instances {
		if i.Name == name {
			return i
		}
	}

	return nil
}

func (router *Router) getFields(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	filter := r.URL.Query().Get("filter")
	fieldType := r.URL.Query().Get("fieldType")

	log.WithFields(logrus.Fields{"name": name, "filter": filter, "fieldType": fieldType}).Tracef("getFields")

	i := router.getInstance(name)
	if i == nil {
		errresponse.Render(w, r, nil, http.StatusBadRequest, "Could not find instance name")
		return
	}

	fields := i.GetFields(filter, fieldType)
	log.WithFields(logrus.Fields{"fields": len(fields)}).Tracef("getFields")
	render.JSON(w, r, fields)
}

// getLogs implements the special handling when the user selected the "logs" options for the "view" configuration. This
// options is intended to use together with the kobsio/fluent-bit-clickhouse Fluent Bit plugin and provides a custom
// query language to get the logs from ClickHouse.
func (router *Router) getLogs(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	query := r.URL.Query().Get("query")
	order := r.URL.Query().Get("order")
	orderBy := r.URL.Query().Get("orderBy")
	timeStart := r.URL.Query().Get("timeStart")
	timeEnd := r.URL.Query().Get("timeEnd")

	log.WithFields(logrus.Fields{"name": name, "query": query, "order": order, "orderBy": orderBy, "timeStart": timeStart, "timeEnd": timeEnd}).Tracef("getLogs")

	i := router.getInstance(name)
	if i == nil {
		errresponse.Render(w, r, nil, http.StatusBadRequest, "Could not find instance name")
		return
	}

	parsedTimeStart, err := strconv.ParseInt(timeStart, 10, 64)
	if err != nil {
		errresponse.Render(w, r, err, http.StatusBadRequest, "Could not parse start time")
		return
	}

	parsedTimeEnd, err := strconv.ParseInt(timeEnd, 10, 64)
	if err != nil {
		errresponse.Render(w, r, err, http.StatusBadRequest, "Could not parse end time")
		return
	}

	// Query for larger time ranges can took several minutes to be completed. To avoid that the connection is closed for
	// these long running requests by a load balancer which sits infront of kobs, we are writing a newline character
	// every 10 seconds. We shouldn't write sth. else, because this would make parsing the response in the React UI more
	// diffucult and with the newline character parsing works in the same ways as it was before.
	done := make(chan bool)

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				if f, ok := w.(http.Flusher); ok {
					// We do not set the processing status code, so that the queries always are returning a 200. This is
					// necessary because Go doesn't allow to set a new status code once the header was written.
					// See: https://github.com/golang/go/issues/36734
					// For that we also have to handle errors, when the status code is 200 in the React UI.
					// See plugins/clickhouse/src/components/page/Logs.tsx#L64
					// w.WriteHeader(http.StatusProcessing)
					w.Write([]byte("\n"))
					f.Flush()
				}
			}
		}
	}()

	defer func() {
		done <- true
	}()

	documents, fields, count, took, buckets, err := i.GetLogs(r.Context(), query, order, orderBy, 1000, parsedTimeStart, parsedTimeEnd)
	if err != nil {
		errresponse.Render(w, r, err, http.StatusBadRequest, "Could not get logs")
		return
	}

	data := struct {
		Documents []map[string]interface{} `json:"documents"`
		Fields    []string                 `json:"fields"`
		Count     int64                    `json:"count"`
		Took      int64                    `json:"took"`
		Buckets   []instance.Bucket        `json:"buckets"`
	}{
		documents,
		fields,
		count,
		took,
		buckets,
	}

	render.JSON(w, r, data)
}

// getAggregation returns the columns and rows for the user given aggregation request. The aggregation data must
// provided in the body of the request and is the run against the specified Clichouse instance.
func (router *Router) getAggregation(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	log.WithFields(logrus.Fields{"name": name}).Tracef("getAggregation")

	i := router.getInstance(name)
	if i == nil {
		errresponse.Render(w, r, nil, http.StatusBadRequest, "Could not find instance name")
		return
	}

	var aggregationData instance.Aggregation

	err := json.NewDecoder(r.Body).Decode(&aggregationData)
	if err != nil {
		errresponse.Render(w, r, err, http.StatusBadRequest, "Could not decode request body")
		return
	}

	done := make(chan bool)

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				if f, ok := w.(http.Flusher); ok {
					// w.WriteHeader(http.StatusProcessing)
					w.Write([]byte("\n"))
					f.Flush()
				}
			}
		}
	}()

	defer func() {
		done <- true
	}()

	rows, columns, err := i.GetAggregation(r.Context(), aggregationData)
	if err != nil {
		errresponse.Render(w, r, err, http.StatusBadRequest, "Error while running aggregation")
		return
	}

	data := struct {
		Rows    []map[string]interface{} `json:"rows"`
		Columns []string                 `json:"columns"`
	}{
		rows,
		columns,
	}

	render.JSON(w, r, data)
}

// Register returns a new router which can be used in the router for the kobs rest api.
func Register(clusters *clusters.Clusters, plugins *plugin.Plugins, config Config) (chi.Router, []*instance.Instance) {
	var instances []*instance.Instance

	for _, cfg := range config {
		instance, err := instance.New(cfg)
		if err != nil {
			log.WithError(err).WithFields(logrus.Fields{"name": cfg.Name}).Fatalf("Could not create ClickHouse instance")
		}

		instances = append(instances, instance)

		plugins.Append(plugin.Plugin{
			Name:        cfg.Name,
			DisplayName: cfg.DisplayName,
			Description: cfg.Description,
			Type:        "clickhouse",
		})
	}

	router := Router{
		chi.NewRouter(),
		clusters,
		instances,
	}

	router.Get("/fields/{name}", router.getFields)
	router.Get("/logs/{name}", router.getLogs)
	router.Post("/aggregation/{name}", router.getAggregation)

	return router, instances
}
