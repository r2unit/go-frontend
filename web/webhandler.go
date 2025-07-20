package web

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var templates *template.Template

var allowedPaths = map[string]bool{}

func init() {
	templates = template.New("")
	var err error

	_, err = templates.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing common templates: %v", err)
	}

	_, err = templates.ParseGlob("pages/*.html")
	if err != nil {
		log.Fatalf("Error parsing page templates: %v", err)
	}
}

func RegisterHandlers() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", HomeHandler)
	allowedPaths["/"] = true

	http.HandleFunc("/contact", ContactHandler)
	allowedPaths["/contact"] = true

	http.HandleFunc("/example", ExampleHandler)
	allowedPaths["/example"] = true

	http.HandleFunc("/404", NotFoundHandler)
	allowedPaths["/404"] = true
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, tmpl, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(buf.Bytes())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	showCookiePolicy := false
	if _, err := r.Cookie("cookieAccepted"); err != nil {
		showCookiePolicy = true
	}

	data := struct {
		Title            string
		ShowCookiePolicy bool
	}{
		Title:            "Nerdpitch - Get yourself ready for the cloud for the rest of us nerds",
		ShowCookiePolicy: showCookiePolicy,
	}

	RenderTemplate(w, "home", data)
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	showCookiePolicy := false
	if _, err := r.Cookie("cookieAccepted"); err != nil {
		showCookiePolicy = true
	}

	data := struct {
		Title            string
		ShowCookiePolicy bool
	}{
		Title:            "Nerdpitch - Example Page ",
		ShowCookiePolicy: showCookiePolicy,
	}

	RenderTemplate(w, "example", data)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	showCookiePolicy := false
	if _, err := r.Cookie("cookieAccepted"); err != nil {
		showCookiePolicy = true
	}

	data := struct {
		Title            string
		ShowCookiePolicy bool
	}{
		Title:            "Contact Us",
		ShowCookiePolicy: showCookiePolicy,
	}

	RenderTemplate(w, "contact", data)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Page Not Found",
	}
	w.WriteHeader(http.StatusNotFound)
	RenderTemplate(w, "404", data)
}

func NotFoundRedirectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/static/") {
			next.ServeHTTP(w, r)
			return
		}

		if _, ok := allowedPaths[r.URL.Path]; !ok {
			http.Redirect(w, r, "/404", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
