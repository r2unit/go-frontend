# Frontend Framework

A Go framework for rendering frontend pages with dynamic route configuration.

## Overview

This framework allows you to create web applications with dynamically configured routes and handlers based on a JSON configuration file. It provides template rendering, static file serving, and middleware for handling 404 redirects.

## Features

- Dynamic route configuration via JSON
- Template rendering with common and page-specific templates
- Static file serving
- 404 redirect middleware
- Cookie policy support

## Installation

```bash
go get github.com/nerdpitch-cloud/frontend
```

## Usage

### 1. Create a routes.json configuration file

```json
{
  "routes": [
    {
      "path": "/",
      "template": "home",
      "title": "Home Page",
      "data": {
        "showCookiePolicy": true
      }
    },
    {
      "path": "/about",
      "template": "about",
      "title": "About Us",
      "data": {
        "showCookiePolicy": true
      }
    }
  ],
  "static": {
    "path": "/static/",
    "dir": "assets"
  },
  "templates": {
    "common": "templates/*.html",
    "pages": "pages/*.html"
  }
}
```

### 2. Create your templates

Create your HTML templates in the templates/ and pages/ directories.

### 3. Use the framework in your application

```go
package main

import (
	"log"
	"net/http"

	"github.com/nerdpitch-cloud/frontend/pkg/framework"
)

func main() {
	// Create a new framework instance
	fw, err := framework.New("config/routes.json")
	if err != nil {
		log.Fatalf("Error creating framework: %v", err)
	}

	// Register handlers
	fw.RegisterHandlers()

	// Apply middleware
	handler := fw.NotFoundRedirectMiddleware(http.DefaultServeMux)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
```

## Directory Structure

The recommended directory structure for your project is:

```
your-project/
├── assets/           # Static files (CSS, JS, images)
├── config/           # Configuration files
│   └── routes.json   # Route configuration
├── pages/            # Page-specific templates
├── templates/        # Common templates
├── go.mod            # Go module file
├── go.sum            # Go dependencies
└── main.go           # Main application
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
