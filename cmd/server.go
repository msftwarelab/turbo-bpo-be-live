package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lonmarsDev/bpo-golang-grahpql/api"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/generated"
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/resolver"
	requestMiddleWare "github.com/lonmarsDev/bpo-golang-grahpql/internal/middleware"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/config"
	"github.com/rs/cors"
)

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

type Server struct{}

func (srv Server) start() {

	port := os.Getenv("PORT")
	if port == "" {
		port = config.AppConfig.GetString("port")
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	router := chi.NewRouter()
	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "HEAD", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"X-Requested-With", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		Debug:            false,
		AllowCredentials: true,
	})

	router.Use(middleware.Timeout(60 * time.Second)) // 1minute timeout
	router.Use(cors.Handler)
	router.Use(requestMiddleWare.Middleware())
	router.Use(requestMiddleWare.LoginLog())
	router.Handle("/", handler.Playground("TurboBPO GraphQL playground", "/graphql"))
	router.Handle("/graphql",
		handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}})),
	)
	//apiRoute(router)
	api.Initialize(router)
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "../resources")
	FileServer(router, "/static", http.Dir(filesDir))

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
	return
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
