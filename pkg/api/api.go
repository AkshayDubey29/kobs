package api

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/kobsio/kobs/pkg/api/clusters"
	"github.com/kobsio/kobs/pkg/api/middleware/auth"
	"github.com/kobsio/kobs/pkg/api/middleware/httplog"
	"github.com/kobsio/kobs/pkg/api/middleware/metrics"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

var (
	log     = logrus.WithFields(logrus.Fields{"package": "api"})
	address string
)

// init is used to define all flags, which are needed for the api server. We have to define the address, where the api
// server is listen on.
func init() {
	defaultAddress := ":15220"
	if os.Getenv("KOBS_API_ADDRESS") != "" {
		defaultAddress = os.Getenv("KOBS_API_ADDRESS")
	}

	flag.StringVar(&address, "api.address", defaultAddress, "The address, where the API server is listen on.")
}

// Server implements the api server. The api server is used to serve the rest api for kobs.
type Server struct {
	server *http.Server
}

// Start starts serving the api server.
func (s *Server) Start() {
	log.Infof("API server listen on %s.", s.server.Addr)

	if err := s.server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.WithError(err).Error("API server died unexpected.")
		} else {
			log.Info("API server was stopped.")
		}
	}
}

// Stop terminates the api server gracefully.
func (s *Server) Stop() {
	log.Debugf("Start shutdown of the API server.")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		log.WithError(err).Error("Graceful shutdown of the API server failed.")
	}
}

// New return a new api server. It creates the underlying http server, with the defined address from the api.address
// flag. When the development flag is set we also set some cors option, so we do not have to care about cors for
// development.
// We exclude the health check from all middlewares, because the health check just returns 200. Therefore we do not need
// our defined middlewares like request id, metrics, auth or loggin. This also makes it easier to analyze the logs in a
// Kubernetes cluster where the health check is called every x seconds, because we generate less logs.
func New(loadedClusters *clusters.Clusters, pluginsRouter chi.Router, isDevelopment bool) (*Server, error) {
	router := chi.NewRouter()

	if isDevelopment {
		router.Use(cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		}))
	}

	router.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, nil)
	})

	router.Route("/api", func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.Recoverer)
		r.Use(middleware.URLFormat)
		r.Use(metrics.Metrics)
		r.Use(auth.Handler(loadedClusters))
		r.Use(httplog.NewStructuredLogger(log.Logger))
		r.Use(render.SetContentType(render.ContentTypeJSON))

		r.Get("/user", auth.UserHandler)
		r.Mount("/clusters", clusters.NewRouter(loadedClusters))
		r.Mount("/plugins", pluginsRouter)
	})

	return &Server{
		server: &http.Server{
			Addr:    address,
			Handler: router,
		},
	}, nil
}
