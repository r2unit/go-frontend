package framework

import (
	"net/http"
)

// Framework represents the web framework
type Framework struct {
	Config       *Config
	Renderer     *TemplateRenderer
	AllowedPaths map[string]bool
}

// New creates a new Framework instance
func New(configPath string) (*Framework, error) {
	// Load configuration
	config, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	// Create template renderer
	renderer, err := NewTemplateRenderer(config.Templates)
	if err != nil {
		return nil, err
	}

	// Create framework
	framework := &Framework{
		Config:       config,
		Renderer:     renderer,
		AllowedPaths: make(map[string]bool),
	}

	return framework, nil
}

// RegisterHandlers registers all handlers defined in the configuration
func (f *Framework) RegisterHandlers() {
	// Register static file handler
	fs := http.FileServer(http.Dir(f.Config.Static.Dir))
	http.Handle(f.Config.Static.Path, http.StripPrefix(f.Config.Static.Path, fs))

	// Register route handlers
	for _, route := range f.Config.Routes {
		// Create a handler for this route
		handler := f.createHandler(route)

		// Register the handler
		http.HandleFunc(route.Path, handler)

		// Add to allowed paths
		f.AllowedPaths[route.Path] = true
	}
}

// createHandler creates a handler function for a route
func (f *Framework) createHandler(route Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create page data
		data := NewPageData(route.Title, route.Data)

		// Check for cookie policy if needed
		if data.ShowCookiePolicy {
			if _, err := r.Cookie("cookieAccepted"); err == nil {
				data.ShowCookiePolicy = false
			}
		}

		// Render the template
		f.Renderer.Render(w, route.Template, data)
	}
}

// NotFoundRedirectMiddleware creates middleware that redirects to 404 page for non-existent paths
func (f *Framework) NotFoundRedirectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow static files
		if r.URL.Path == "/" || r.URL.Path == "/404" || r.URL.Path == f.Config.Static.Path ||
			(len(r.URL.Path) > len(f.Config.Static.Path) && r.URL.Path[:len(f.Config.Static.Path)] == f.Config.Static.Path) {
			next.ServeHTTP(w, r)
			return
		}

		// Check if path is allowed
		if _, ok := f.AllowedPaths[r.URL.Path]; !ok {
			http.Redirect(w, r, "/404", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
