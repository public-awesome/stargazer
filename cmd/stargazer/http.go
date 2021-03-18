package main

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
)

func startHTTP(ctx context.Context) {
	m := &http.ServeMux{}
	m.HandleFunc("/stargazer/healthz", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})
	server := &http.Server{Addr: ":1670", Handler: m}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Err(err).Msg("error running http server")
		}
	}()
	<-ctx.Done()
}
