package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaredwarren/macroPi/log"
	"github.com/jaredwarren/macroPi/macro"
)

const (
	EditRowErrorTemplate = "templates/macro/edit_row_error.html"
)

type GenericPage map[string]any

func ListMacros(w http.ResponseWriter, r *http.Request) {
	homepageTpl := template.Must(template.New("base").ParseFiles(
		"templates/base.html",
		"templates/home.html",
	))

	ms := macro.ListtMacros()

	render(w, r, homepageTpl, &GenericPage{
		"title":  "home",
		"macros": ms,
	})
}

// func ListMacros(w http.ResponseWriter, r *http.Request) {
// 	homepageTpl := template.Must(template.ParseFiles(
// 		"templates/home.html",
// 		"templates/base.html",
// 	))
// 	render(w, r, homepageTpl, nil)
// }

func GetMacroEditRowForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	macroID := vars["macro_id"]

	var m *macro.Macro
	if macroID == "new" {
		m = &macro.Macro{}
	} else {
		m = macro.GetMacro(macroID)
		if m == nil {
			renderError(EditRowErrorTemplate, fmt.Errorf("not found"), w, r)
			return
		}
	}

	homepageTpl := template.Must(template.New("base").ParseFiles(
		"templates/macro/edit_row.html",
	))

	render(w, r, homepageTpl, &GenericPage{
		"id":    macroID,
		"macro": m,
	})
}

func GetMacroRow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.GetLogger(ctx)

	vars := mux.Vars(r)
	macroID := vars["macro_id"]

	m := macro.GetMacro(macroID)
	if m == nil {
		logger.Error("GetMacro error", log.Any("id", macroID))
		renderError(EditRowErrorTemplate, fmt.Errorf("not found"), w, r)
		return
	}

	homepageTpl := template.Must(template.New("base").ParseFiles(
		"templates/macro/get_row.html",
	))

	render(w, r, homepageTpl, &GenericPage{
		"id":    macroID,
		"macro": m,
	})
}

func UpdateMacro(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.GetLogger(ctx)

	vars := mux.Vars(r)
	macroID := vars["macro_id"]
	isNew := (macroID == "new")

	err := r.ParseForm()
	if err != nil {
		logger.Error("ParseForm error", log.Error(err))
		renderError(EditRowErrorTemplate, err, w, r)
		return
	}

	label := r.PostForm.Get("label")
	command := r.PostForm.Get("command")
	om := &macro.Macro{
		Label:   label,
		Command: command,
	}
	i, err := macro.UpdateMacro(macroID, om)
	if err != nil {
		logger.Error("UpdateMacro error", log.Error(err))
		renderError(EditRowErrorTemplate, err, w, r)
		return
	}

	// update id in vars, specifically for "new"
	vars["macro_id"] = fmt.Sprintf("%d", i)
	mux.SetURLVars(r, vars)

	GetMacroRow(w, r)
	if isNew {
		homepageTpl := template.Must(template.New("base").ParseFiles(
			"templates/macro/edit_row_new.html",
		))

		render(w, r, homepageTpl, &GenericPage{
			"err": "",
		})
	}
}

func DeleteMacro(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.GetLogger(ctx)

	vars := mux.Vars(r)
	macroID := vars["macro_id"]

	err := r.ParseForm()
	if err != nil {
		renderError(EditRowErrorTemplate, err, w, r)
		return
	}

	label := r.PostForm.Get("label")
	command := r.PostForm.Get("command")

	logger.Error("Delete", log.Any("macroID", macroID), log.Any("command", command), log.Any("label", label))

	err = macro.DeleteMacro(macroID)
	if err != nil {
		renderError(EditRowErrorTemplate, err, w, r)
		return
	}
}

func renderError(tpl string, err error, w http.ResponseWriter, r *http.Request) {
	homepageTpl := template.Must(template.New("base").ParseFiles(
		tpl,
	))

	render(w, r, homepageTpl, &GenericPage{
		"err": err.Error(),
	})
}

func RunMacro(w http.ResponseWriter, r *http.Request) {

}
