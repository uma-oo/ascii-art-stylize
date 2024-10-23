package handler

import (
	"net/http"
	"regexp"
	"text/template"

	"asciiWeb/internal"
)

type Data struct {
	Text      string
	Banner    string
	FormError string
	AsciiArt  string
}

var (
	Pagedata  = Data{}
	Templates *template.Template
)

// handler for the path "/"
func HandleMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] != "" {
		renderTemplate(w, "errorPage.html", http.StatusNotFound)
		return
	}
	renderTemplate(w, "index.html", Pagedata)
	Pagedata = Data{}
}

// handler for the path "/ascii-art
func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderTemplate(w, "errorPage.html", http.StatusMethodNotAllowed)
		return
	}

	if extractFormData(w, r) != "200" {
		return
	}

	formatedText := regexp.MustCompile(`\r\n`).ReplaceAllString(Pagedata.Text, `\n`)
	asciiArt, err, st := internal.Ascii(formatedText, Pagedata.Banner)
	if st == "500" {
		w.WriteHeader(http.StatusInternalServerError)
		renderTemplate(w, "errorPage.html", http.StatusInternalServerError)
		return
	} else if st == "405" {
		w.WriteHeader(http.StatusBadRequest)
		Pagedata.FormError = err
		renderTemplate(w, "index.html", Pagedata)
		Pagedata.FormError = ""
		return
	} else if st == "404" {
		w.WriteHeader(http.StatusNotFound)
		renderTemplate(w, "errorPage.html", http.StatusNotFound)
		return

	}

	Pagedata.AsciiArt = asciiArt
	renderTemplate(w, "index.html", Pagedata)
	Pagedata = Data{}
}
