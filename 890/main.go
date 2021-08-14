package main

import (
	"net/http"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/rest"
)

type codedResponseWriter struct {
	w    http.ResponseWriter
	r    *http.Request
	code int
}

func (w *codedResponseWriter) Flush() {
	if flusher, ok := w.w.(http.Flusher); ok {
		flusher.Flush()
	}
}

func (w *codedResponseWriter) Header() http.Header {
	return w.w.Header()
}

func (w *codedResponseWriter) Write(bytes []byte) (int, error) {
	if w.code == http.StatusInternalServerError {
		w.w.Write([]byte("oops"))
	}
	return w.w.Write(bytes)
}

func (w *codedResponseWriter) WriteHeader(code int) {
	w.w.WriteHeader(code)
	w.code = code
	if w.code == http.StatusInternalServerError {
		w.w.Write([]byte("oops"))
	}
}

func main() {
	srv := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Name: "adhoc",
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Host: "127.0.0.1",
		Port: 8888,
	})
	srv.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/panic",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			panic("bad things happen")
		},
	})
	srv.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/error",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		},
	})
	srv.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if p := recover(); p != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("oops"))
				}
			}()

			cw := &codedResponseWriter{
				w: w,
				r: r,
			}
			next(cw, r)
		}
	})
	defer srv.Stop()
	srv.Start()
}
