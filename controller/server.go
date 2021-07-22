package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

const (
	serverPort int    = 8080
	awsRegion  string = "us-east-1"
)

type Server struct {
	Router chi.Router
}

func NewServer() *Server {
	s := Server{}
	r := chi.NewRouter()
	s.Router = r
	return &s
}

func (s *Server) Initialize() {
	s.initializeRoutes()
	fmt.Println("Backend successfully initialized and listening")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", serverPort), s.Router))
}

func (s *Server) initializeRoutes() {
	s.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Access-Control-Allow-Headers", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
		Debug:            false,
	}))
	s.Router.Use(render.SetContentType(render.ContentTypeJSON))
	// Create backend endpoints for the frontend to consume
	s.Router.Route("/api/buckets", func(r chi.Router) {
		r.Get("/", s.BucketsRouter)
	})
	s.Router.Route("/api/objects", func(r chi.Router) {
		r.Route("/bucket/{bucketName}", func(r chi.Router) {
			r.Get("/", s.ObjectsRouter)
			r.Post("/", s.BucketsRouter)
			r.Delete("/", s.BucketsRouter)
			r.Route("/key/{objKey}", func(r chi.Router) {
				r.Delete("/", s.ObjectRouter)
				r.Get("/", s.ObjectRouter)
				r.Post("/", s.ObjectRouter)
			})
		})
	})
}
