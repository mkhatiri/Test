package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) showPatient(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		files := []string{
			"./ui/html/patients.page.tmpl",
			"./ui/html/base.layout.tmpl",
			"./ui/html/footer.partial.tmpl",
		}

		ts, err := template.ParseFiles(files...)

		if err != nil {
			app.errorLog.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = ts.Execute(w, nil)

		if err != nil {
			app.errorLog.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	i, err := fmt.Fprintf(w, "Show patient number %d", id)
	app.infoLog.Println(i, err)
}

func (app *application) createPatient(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.Header().Add("name", "med")

		w.Header().Del("name")
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		// this is equivalent to
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("createPatient"))
}

func (app *application) downloadHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./ui/static/files/go.pdf")
}
