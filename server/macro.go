package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaredwarren/macroPi/log"
	"github.com/jaredwarren/macroPi/macro"
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
	vars := mux.Vars(r)
	macroID := vars["macro_id"]

	m := macro.GetMacro(macroID)
	if m == nil {
		// TODO: return "edit_error.html" stub
		m = &macro.Macro{
			Label:   "not found",
			Command: "error",
		}
	}

	homepageTpl := template.Must(template.New("base").ParseFiles(
		"templates/edit.html",
	))

	render(w, r, homepageTpl, &GenericPage{
		"title": "home",
		"id":    macroID,
		"macro": m,
	})
}

func ShowMacroForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	macroID := vars["macro_id"]

	m := macro.GetMacro(macroID)
	if m == nil {
		// TODO: return "edit_error.html" stub
		m = &macro.Macro{
			Label:   "not found",
			Command: "error",
		}
	}

	homepageTpl := template.Must(template.New("base").ParseFiles(
		"templates/get.html",
	))

	render(w, r, homepageTpl, &GenericPage{
		"title": "home",
		"id":    macroID,
		"macro": m,
	})
}

func UpdateMacro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	macroID := vars["macro_id"]

	om := macro.GetMacro(macroID)
	if om == nil {
		// TODO: return "edit_error.html" stub
		om = &macro.Macro{
			Label:   "not found",
			Command: "error",
		}
	}

	// TODO: get params
	// update macro

	homepageTpl := template.Must(template.New("base").ParseFiles(
		"templates/get.html",
	))

	render(w, r, homepageTpl, &GenericPage{
		"title": "home",
		"id":    macroID,
		"macro": om,
	})
}

func DeleteMacro(w http.ResponseWriter, r *http.Request) {

}

func RunMacro(w http.ResponseWriter, r *http.Request) {

}
