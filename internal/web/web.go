package web

import (
	"L0/internal/entity"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ModelGetter func(string) *entity.Model

type Web struct {
	Files    string
	GetModel ModelGetter
}

func (w *Web) getHandler() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{id}", func(writer http.ResponseWriter, request *http.Request) {
		model := w.GetModel(chi.URLParam(request, "id"))
		if model == nil {
			writer.WriteHeader(404)
		} else {
			b, err := json.Marshal(*model)
			if err != nil {
				writer.WriteHeader(404)
			} else {
				writer.Write(b)
			}
		}

	})
	r.Handle("/", http.FileServer(http.Dir(w.Files)))
	return r
}

func (w *Web) Run(address string) {
	http.ListenAndServe(address, w.getHandler())
}
