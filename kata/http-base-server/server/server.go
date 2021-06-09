package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type App struct {
	srv *http.Server
}

func NewApp(addr string) *App {
	a := &App{}
	a.srv = &http.Server{
		Addr:    addr,
		Handler: a.handlers(),
	}

	return a
}

func (a *App) Run() error {
	return a.srv.ListenAndServe()
}

func (a *App) handlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := struct {
			Msg string `json:"msg"`
		}{
			Msg: "Hello dude",
		}

		body, _ := json.Marshal(message)

		w.Header().Set("Content-Type", "application/json")

		_, err := w.Write(body)
		if err != nil {
			log.Panicf("write error: %s", err)
		}
	})

	return mux
}
