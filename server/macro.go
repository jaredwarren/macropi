package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jaredwarren/macroPi/log"
	"github.com/jaredwarren/macroPi/macro"
)

const (
	EditRowErrorTemplate = "templates/macro/edit_row_error.html"
)

type GenericPage map[string]any

func (h *HTMLServer) ListMacros(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.GetLogger(ctx)
	homepageTpl := template.Must(template.New("base").ParseFiles(
		"templates/base.html",
		"templates/home.html",
	))

	ms, err := h.DB.ListtMacros()
	if err != nil {
		logger.Error("ListMacro error", log.Error(err))
		// What to return here!
	}

	render(w, r, homepageTpl, &GenericPage{
		"title":  "home",
		"macros": ms,
	})
}

func (h *HTMLServer) GetMacroEditRowForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.GetLogger(ctx)

	vars := mux.Vars(r)
	macroID := vars["macro_id"]
	logger = logger.With(log.Any("id", macroID))
	isNew := macroID == "new"

	var m *macro.Macro
	if isNew {
		m = &macro.Macro{}
	} else {
		m, err := h.DB.GetMacro(macroID)
		if err != nil {
			logger.Error("GetMacro error", log.Error(err))
			renderError(EditRowErrorTemplate, fmt.Errorf("not found"), w, r)
			return
		}
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
		"isNew": isNew,
	})
}

func (h *HTMLServer) GetMacroRow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.GetLogger(ctx)

	vars := mux.Vars(r)
	macroID := vars["macro_id"]
	logger = logger.With(log.Any("id", macroID))

	// show "add new" button if adding a new macro
	if macroID == "new" {
		homepageTpl := template.Must(template.New("base").ParseFiles(
			"templates/macro/edit_row_new.html",
		))

		render(w, r, homepageTpl, &GenericPage{
			"err": "",
		})
		return
	}

	m, err := h.DB.GetMacro(macroID)
	if err != nil {
		logger.Error("GetMacro error", log.Error(err))
		renderError(EditRowErrorTemplate, fmt.Errorf("not found"), w, r)
		return
	}
	if m == nil {
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

func (h *HTMLServer) UpdateMacro(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.GetLogger(ctx)

	vars := mux.Vars(r)
	macroID := vars["macro_id"]
	isNew := (macroID == "new")
	if isNew {
		macroID = uuid.New().String()
	}
	logger = logger.With(log.Any("id", macroID), log.Any("is_new", isNew))

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
	logger = logger.With(log.Any("macro", om))

	err = h.DB.UpdateMacro(macroID, om)
	if err != nil {
		logger.Error("UpdateMacro error", log.Error(err))
		renderError(EditRowErrorTemplate, err, w, r)
		return
	}

	// update id in vars, specifically for "new"
	vars["macro_id"] = macroID
	mux.SetURLVars(r, vars)

	// return new or updated macro
	h.GetMacroRow(w, r)
	// show "add new" button if adding a new macro
	if isNew {
		homepageTpl := template.Must(template.New("base").ParseFiles(
			"templates/macro/edit_row_new.html",
		))

		render(w, r, homepageTpl, &GenericPage{
			"err": "",
		})
	}
}

func (h *HTMLServer) DeleteMacro(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.GetLogger(ctx)

	vars := mux.Vars(r)
	macroID := vars["macro_id"]
	logger = logger.With(log.Any("id", macroID))

	err := r.ParseForm()
	if err != nil {
		renderError(EditRowErrorTemplate, err, w, r)
		return
	}

	label := r.PostForm.Get("label")
	command := r.PostForm.Get("command")
	logger = logger.With(log.Any("command", command), log.Any("label", label))

	logger.Error("Delete")

	err = h.DB.DeleteMacro(macroID)
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
