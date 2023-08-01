package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/swaggo/http-swagger"
	"mini-project-new/docs"
	"mini-project-new/internal/handlers"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	//r.Use(middleware.CORS)

	docs.SwaggerInfo.Title = "Event Management"
	docs.SwaggerInfo.Version = "v1"
	conf := httpSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.Get("/swagger/*", httpSwagger.Handler(conf))

	handlers.RouteHandler(r)

	fmt.Printf("Server listening on port %s\n", "8080")
	http.ListenAndServe(":8080", r)
}
