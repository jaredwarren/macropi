package server

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/home.html",
		"templates/base.html",
	}
	homepageTpl := template.Must(template.ParseFiles(files...))

	render(w, r, homepageTpl, nil)
}

func ListMacros(http.ResponseWriter, *http.Request) {

}

func ShowMacroForm(http.ResponseWriter, *http.Request) {

}

func UpdateMacro(http.ResponseWriter, *http.Request) {

}

func DeleteMacro(http.ResponseWriter, *http.Request) {

}

func RunMacro(http.ResponseWriter, *http.Request) {

}
