# Frontend Framework

A Go framework for rendering frontend pages with dynamic route configuration, designed for a multi-repository architecture.

## Overview

This framework allows you to create web applications with dynamically configured routes and handlers based on a JSON configuration file. It provides template rendering, static file serving, and middleware for handling 404 redirects.

The framework is designed to be used in a multi-repository architecture:
- **Framework Repository**: Contains the core framework code
- **Template Repositories**: Contain templates, pages, assets, and route configuration

## Features

- Dynamic route configuration via JSON
- Template rendering with common and page-specific templates
- Static file serving
- 404 redirect middleware
- Cookie policy support
- Multi-repository architecture with automated builds
- Docker containerization

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

## Multi-Repository Architecture

This framework is designed to be used in a multi-repository setup:

1. **Framework Repository** (this repository):
   - Contains the core framework code
   - Provides the Go package for rendering templates and handling routes
   - Includes the GitHub Workflow for automated builds

2. **Template Repositories**:
   - Named with the pattern `frontend-*` (e.g., `frontend-website`, `frontend-blog`)
   - Contain only templates, pages, assets, and route configuration
   - No Go code required
   - Automatically built into Docker containers using the framework

### How It Works

When you push to a template repository:
1. The GitHub Workflow detects if the repository name starts with `frontend-`
2. If it does, it checks out both the template repository and the framework repository
3. It combines them into a build directory
4. It builds a Docker container and pushes it to GitHub Container Registry

### Docker Images

The framework uses a specific naming and tagging scheme for Docker images:

#### Image Names

Docker images are named after the repository name (without the organization prefix). For example, a repository named `nerdpitch-cloud/frontend-website` will produce a Docker image named `ghcr.io/frontend-website`.

#### Tagging Strategy

The framework uses Semantic Versioning for Docker image tags:

1. **For the develop branch:**
   - Format: `develop-[SemVer]` (e.g., `develop-1.2.3`)
   - If no semantic versioning Git tags exist, fallback to `develop-[branch]-[sha]`

2. **For other branches (including main):**
   - Format: `[SemVer]` (e.g., `1.2.3`)
   - If no semantic versioning Git tags exist, fallback to `[branch]-[sha]`
   - The `latest` tag is also applied to images built from the main branch

To use semantic versioning, create Git tags following the pattern `v1.2.3` in your repository.

#### Container Metadata

Each Docker image includes metadata that links it to the source GitHub repository:

- `org.opencontainers.image.title`: The repository name
- `org.opencontainers.image.description`: Description of the image
- `org.opencontainers.image.source`: URL to the GitHub repository
- `org.opencontainers.image.authors`: r2unit@proton.me
- `org.opencontainers.image.vendor`: Nerdpitch Cloud

## Directory Structure

### Framework Repository

```
frontend-framework/
├── .github/
│   └── workflows/     # GitHub Workflows for automated builds
├── pkg/
│   └── framework/     # Framework code
├── docs/              # Documentation
├── go.mod             # Go module file
├── go.sum             # Go dependencies
├── main.go            # Main application
└── dockerfile         # Dockerfile for building containers
```

### Template Repository

```
frontend-your-project/
├── assets/            # Static files (CSS, JS, images)
├── config/            # Configuration files
│   └── routes.json    # Route configuration
├── pages/             # Page-specific templates
└── templates/         # Common templates
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
