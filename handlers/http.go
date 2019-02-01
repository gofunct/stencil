package net

import (
	"fmt"
	"github.com/gofunct/gofs/scripter"
	"github.com/gofunct/stencil"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"net/http"
	"net/http/pprof"
)

type MuxFunc func(*http.ServeMux)

type RouterFunc func(route *mux.Router)

func InitializeRouter() RouterFunc {
	return func(handler *mux.Router) {
		handler.HandleFunc("/debug/pprof", http.HandlerFunc(pprof.Index))
		handler.HandleFunc("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		handler.HandleFunc("/debug/trace", http.HandlerFunc(pprof.Trace))
		handler.HandleFunc("/debug/profile", http.HandlerFunc(pprof.Profile))
		handler.HandleFunc("/debug/symbol", http.HandlerFunc(pprof.Symbol))
		handler.Handle("/metrics", promhttp.Handler())
	}
}

func InitializeMux() MuxFunc {
	return func(handler *http.ServeMux) {
		handler.HandleFunc("/debug/pprof", http.HandlerFunc(pprof.Index))
		handler.HandleFunc("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		handler.HandleFunc("/debug/trace", http.HandlerFunc(pprof.Trace))
		handler.HandleFunc("/debug/profile", http.HandlerFunc(pprof.Profile))
		handler.HandleFunc("/debug/symbol", http.HandlerFunc(pprof.Symbol))
		handler.Handle("/metrics", promhttp.Handler())
	}
}

func Get(pattern string, url string) MuxFunc {
	return func(handler *http.ServeMux) {
		handler.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			client := http.Client{
				Transport:     nil,
				CheckRedirect: nil,
				Jar:           nil,
				Timeout:       0,
			}
			resp, err := client.Get(url)
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
		})
	}
}

func Post(pattern, url, content string, reader io.Reader) MuxFunc {
	return func(handler *http.ServeMux) {
		handler.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			client := http.Client{
				Transport:     nil,
				CheckRedirect: nil,
				Jar:           nil,
				Timeout:       0,
			}
			resp, err := client.Post(url, "", reader)
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
		})
	}
}

func Load(pattern, url string) MuxFunc {
	return func(handler *http.ServeMux) {
		handler.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			stencil.GetAndWrite(url, w)
		})
	}
}

func Script(pattern string, scripter *scripter.Scripter) MuxFunc {
	return func(handler *http.ServeMux) {
		handler.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			if scripter.GetBits() == nil {
				scripter.Run()
			}
			scripter.WriteTo(w)
		})
	}
}
