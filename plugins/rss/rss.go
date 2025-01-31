package rss

import (
	"net/http"
	"sync"

	"github.com/kobsio/kobs/pkg/api/clusters"
	"github.com/kobsio/kobs/pkg/api/plugins/plugin"
	"github.com/kobsio/kobs/plugins/rss/pkg/feed"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
)

// Route is the route under which the plugin should be registered in our router for the rest api.
const Route = "/rss"

var (
	log = logrus.WithFields(logrus.Fields{"package": "rss"})
)

// Config is the structure of the configuration for the rss plugin.
type Config struct{}

// Router implements the router for the resources plugin, which can be registered in the router for our rest api.
type Router struct {
	*chi.Mux
	clusters *clusters.Clusters
	config   Config
}

// getFeed returns a feed with the retrieved items from the given links.
func (router *Router) getFeed(w http.ResponseWriter, r *http.Request) {
	urls := r.URL.Query()["url"]
	sortBy := r.URL.Query().Get("sortBy")

	var feeds []*gofeed.Feed
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			fp := gofeed.NewParser()
			feed, err := fp.ParseURL(url)
			if err != nil {
				log.WithError(err).Error("Error while getting feed")
			}

			if feed != nil {
				feeds = append(feeds, feed)
			}

			wg.Done()
		}(url)
	}

	wg.Wait()

	items := feed.Transform(feeds, sortBy)

	log.WithFields(logrus.Fields{"links": len(urls), "sortBy": sortBy, "items": len(items)}).Tracef("getFeed")

	render.JSON(w, r, items)
}

// Register returns a new router which can be used in the router for the kobs rest api.
func Register(clusters *clusters.Clusters, plugins *plugin.Plugins, config Config) chi.Router {
	plugins.Append(plugin.Plugin{
		Name:        "rss",
		DisplayName: "RSS",
		Description: "Get the latest status updates of your third party services.",
		Type:        "rss",
	})

	router := Router{
		chi.NewRouter(),
		clusters,
		config,
	}

	router.Get("/feed", router.getFeed)

	return router
}
