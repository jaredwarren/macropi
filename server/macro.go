package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jaredwarren/macroPi/log"
)

type GenericPage map[string]any

func (h *HTMLServer) Home(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	// logger := log.GetLogger(ctx)
	homepageTpl := template.Must(template.New("base").ParseFiles(
		"templates/base.html",
		"templates/home.html",
	))

	cfg := h.Macro

	render(w, r, homepageTpl, &GenericPage{
		"title":  "home",
		"macros": cfg.Macros,
	})
}

func ListMacros(w http.ResponseWriter, r *http.Request) {
	fmt.Println("~~~~~")
	ctx := r.Context()
	logger := log.GetLogger(ctx)
	logger.Info("ListMacros!")
	files := []string{
		"templates/home.html",
		"templates/base.html",
	}
	homepageTpl := template.Must(template.ParseFiles(files...))

	render(w, r, homepageTpl, nil)
}

func GetEditMacroForm(w http.ResponseWriter, r *http.Request) {
	homepageTpl := template.Must(template.New("base").ParseFiles(
		"templates/edit.html",
	))

	// cfg := h.Macro

	render(w, r, homepageTpl, &GenericPage{
		"title": "home",
		// "macros": cfg.Macros,
	})
}

func ShowMacroForm(w http.ResponseWriter, r *http.Request) {

}

func UpdateMacro(w http.ResponseWriter, r *http.Request) {

}

func DeleteMacro(w http.ResponseWriter, r *http.Request) {

}

func RunMacro(w http.ResponseWriter, r *http.Request) {

}
