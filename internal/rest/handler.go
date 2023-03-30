package rest

import (
	"YYY/internal/adapter"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"net/http"
)

func Handler(a *adapter.Adapter) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			HandlerIndex(w)
		})
		r.Get("/groups", a.GetGroups)
		r.Route("/group", func(r chi.Router) {
			r.Post("/", a.PostGroup)
			r.Get("/{id}", a.GetGroupsByID)
		})
	})

	return router
}

func HandlerIndex(out io.Writer) {
	_, err := fmt.Fprintf(out, "hello world")
	if err != nil {
		return
	}
}
