package web

import (
	"L0/internal/entity"
	"encoding/json"
	"fmt"
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
		id := chi.URLParam(request, "id")
		fmt.Println("Поиск заказа:", id)
		model := w.GetModel(id)
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
