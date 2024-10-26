package handler

import (
	"fmt"
	"net/http"
	"regexp"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := Templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "<h1  style='color: brown;'>Not Found 404</h1>")
		return
	}
}

func extractFormData(w http.ResponseWriter, r *http.Request) string {
	err := r.ParseForm()
	if err != nil {
		renderTemplate(w, "errorPage.html", http.StatusInternalServerError)
		return "500"
	}

	text := r.FormValue("text")

	if len(text)>1000{
		renderTemplate(w, "errorPage.html", http.StatusBadRequest)
		return "400"
	}

	if text == "" {
		Pagedata.FormError = "You need to provide a text"
	}
	banner := r.FormValue("banner")
	if !IsBanner(banner) {
		renderTemplate(w, "errorPage.html", http.StatusBadRequest)
		return "500"
	}

	if textReg := regexp.MustCompile(`^\r\n+`); textReg.MatchString(text) {
		Pagedata.Text = "\r\n" + text
	} else {
		Pagedata.Text = text
	}
	Pagedata.Banner = banner

	return "200"
}

func IsBanner(banner string) bool {
	return banner == "standard" || banner == "shadow" || banner == "thinkertoy"
}


