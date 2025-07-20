# Template Repository Structure

This document describes how to set up a new frontend repository that uses the frontend-framework.

## Repository Naming

To use the automated build workflow, your repository name should start with `frontend-`. For example:
- `frontend-website`
- `frontend-blog`
- `frontend-dashboard`

## Directory Structure

Your template repository should have the following structure:

```
frontend-your-project/
├── assets/           # Static files (CSS, JS, images)
├── config/           # Configuration files
│   └── routes.json   # Route configuration
├── pages/            # Page-specific templates
├── templates/        # Common templates
└── README.md         # Documentation for your frontend
```

## Required Files

### config/routes.json

This file defines the routes for your frontend application. Example:

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

### templates/

This directory contains common templates used across multiple pages. At minimum, you should have:

- `header.html`: The header template
- `footer.html`: The footer template
- Any other common components

### pages/

This directory contains page-specific templates. Each page should have its own template file. For example:

- `home.html`: The home page template
- `about.html`: The about page template
- `404.html`: The 404 page template

## How It Works

When you push to the `main` or `develop` branch of your repository, the GitHub Workflow will:

1. Check if your repository name starts with `frontend-`
2. If it does, it will:
   - Check out your repository
   - Check out the frontend-framework repository
   - Combine them into a build directory
   - Build a Docker container
   - Push the container to GitHub Container Registry

The resulting container will include:
- The framework code from the frontend-framework repository
- Your templates, pages, assets, and routes.json

## Testing Locally

To test your frontend locally before pushing to GitHub, you can:

1. Clone both repositories:
   ```bash
   git clone https://github.com/nerdpitch-cloud/frontend-framework.git
   git clone https://github.com/your-org/frontend-your-project.git
   ```

2. Create a build directory:
   ```bash
   mkdir -p build
   ```

3. Copy the framework code:
   ```bash
   cp -r frontend-framework/pkg build/
   cp frontend-framework/main.go build/
   cp frontend-framework/go.mod build/
   cp frontend-framework/go.sum build/
   ```

4. Copy your template files:
   ```bash
   mkdir -p build/templates build/pages build/assets build/config
   cp -r frontend-your-project/templates/* build/templates/
   cp -r frontend-your-project/pages/* build/pages/
   cp -r frontend-your-project/assets/* build/assets/
   cp frontend-your-project/config/routes.json build/config/
   ```

5. Build and run the container:
   ```bash
   cd build
   docker build -t frontend-your-project -f ../frontend-framework/dockerfile .
   docker run -p 8080:8080 frontend-your-project
   ```

6. Open http://localhost:8080 in your browser to see your frontend.