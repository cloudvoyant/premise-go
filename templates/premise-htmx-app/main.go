package main

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:embed web/*
var assets embed.FS

// Private helpers ------------------------------------------------------------
func handler() http.Handler {
	mux := http.NewServeMux()
	page := template.Must(template.ParseFS(assets, "web/index.html"))
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		if err := page.Execute(w, nil); err != nil {
			log.Print(err)
		}
	})
	mux.HandleFunc("POST /greet", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, "<p>Hello, "+template.HTMLEscapeString(r.FormValue("name"))+"!</p>"); err != nil {
			log.Print(err)
		}
	})
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assets))))
	return mux
}
func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	server := &http.Server{Addr: addr, Handler: handler(), ReadHeaderTimeout: 5 * time.Second}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		<-ctx.Done()
		shutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = server.Shutdown(shutdown)
	}()
	log.Printf("listening on %s", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
