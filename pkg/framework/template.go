package framework

import (
	"bytes"
	"html/template"
	"net/http"
)

// TemplateRenderer handles template rendering
type TemplateRenderer struct {
	templates *template.Template
}

// NewTemplateRenderer creates a new TemplateRenderer
func NewTemplateRenderer(config TemplateConfig) (*TemplateRenderer, error) {
	templates := template.New("")
	var err error

	// Parse common templates
	_, err = templates.ParseGlob(config.Common)
	if err != nil {
		return nil, err
	}

	// Parse page templates
	_, err = templates.ParseGlob(config.Pages)
	if err != nil {
		return nil, err
	}

	return &TemplateRenderer{
		templates: templates,
	}, nil
}

// Render renders a template with the given data
func (tr *TemplateRenderer) Render(w http.ResponseWriter, tmpl string, data interface{}) error {
	var buf bytes.Buffer
	if err := tr.templates.ExecuteTemplate(&buf, tmpl, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Write(buf.Bytes())
	return nil
}

// PageData represents the data passed to a template
type PageData struct {
	Title            string
	ShowCookiePolicy bool
	CustomData       map[string]interface{}
}

// NewPageData creates a new PageData with the given title and custom data
func NewPageData(title string, customData map[string]interface{}) PageData {
	showCookiePolicy := false
	if val, ok := customData["showCookiePolicy"]; ok {
		if boolVal, ok := val.(bool); ok {
			showCookiePolicy = boolVal
		}
	}

	return PageData{
		Title:            title,
		ShowCookiePolicy: showCookiePolicy,
		CustomData:       customData,
	}
}
