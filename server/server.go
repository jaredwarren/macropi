package server

import (
	"bytes"
	"context"
	"html/template"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/jaredwarren/macroPi/log"
	"github.com/spf13/viper"
)

// Config provides basic configuration
type Config struct {
	Addr         string `yaml:"addr"`
	HTTPS        bool   `yaml:"https"`
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	// Db           db.DBer
	Logger log.Logger
}

// HTMLServer represents the web service that serves up HTML
type HTMLServer struct {
	Logger log.Logger
	Config *Config

	server *http.Server
	wg     sync.WaitGroup
}

// Start launches the HTML Server
func (h *HTMLServer) Start() {
	logger := h.Logger

	// Setup Handlers
	r := mux.NewRouter()
	r.Use(NewLoggingMiddleware(logger))

	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// login-required methods
	// home
	// CRUD macro
	// CRUD profile(note: switch profile key is a "macro")
	r.HandleFunc("/", ListMacros).Methods(http.MethodGet)

	// r.HandleFunc("/keys", ListKeys).Methods(http.MethodGet)
	//
	// Macros
	//
	r.HandleFunc("/macros", ListMacros).Methods(http.MethodGet)

	macsub := r.PathPrefix("/macro/{macro_id}").Subrouter()

	// htmx
	macsub.HandleFunc("", GetMacroRow).Methods(http.MethodGet)
	macsub.HandleFunc("/edit", GetMacroEditRowForm).Methods(http.MethodGet)
	macsub.HandleFunc("", UpdateMacro).Methods(http.MethodPost, http.MethodPut)
	macsub.HandleFunc("", DeleteMacro).Methods(http.MethodDelete)

	// Run
	macsub.HandleFunc("/run", RunMacro).Methods(http.MethodGet)

	h.server = &http.Server{
		Addr:           h.Config.Addr,
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}

	// Start the listener
	h.wg.Add(1)
	go func() {
		logger.Info("Starting HTTP server", log.Any("host", h.server.Addr), log.Any("https", viper.GetBool("https")))
		if viper.GetBool("host.https") {
			h.server.ListenAndServeTLS("localhost.crt", "localhost.key")
		} else {
			h.server.ListenAndServe()
		}
		h.wg.Done()
	}()
}

func NewLoggingMiddleware(l log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = log.WithLogger(ctx, l)
			// s.logger.Debug(r.RequestURI, log.Any("r", r))
			// s.logger.Debug("[request] - " + r.RequestURI)
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// Stop turns off the HTML Server
func (h *HTMLServer) StopHTTPServer() error {
	// Create a context to attempt a graceful 5 second shutdown.
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	h.Logger.Info("Stopping HTTP service...")

	// Attempt the graceful shutdown by closing the listener
	// and completing all inflight requests
	if err := h.server.Shutdown(ctx); err != nil {
		// Looks like we timed out on the graceful shutdown. Force close.
		if err := h.server.Close(); err != nil {
			h.Logger.Error("error stopping HTML service", log.Error(err))
			return err
		}
	}

	// Wait for the listener to report that it is closed.
	h.wg.Wait()
	h.Logger.Info("HTTP service stopped")
	return nil
}

// Render a template, or server error.
func render(w http.ResponseWriter, r *http.Request, tpl *template.Template, data interface{}) {
	ctx := r.Context()
	logger := log.GetLogger(ctx)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, data); err != nil {
		logger.Error("template render error", log.Error(err), log.Any("data", data))
		return
	}
	_, err := w.Write(buf.Bytes())
	if err != nil {
		logger.Error("template write error", log.Error(err), log.Any("data", data))
		return
	}
}
