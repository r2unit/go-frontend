# Remote Repositories Workflow

This document explains the GitHub Workflow that builds frontend applications using templates, assets, pages, and configuration from remote repositories under the `/r2unit/` user.

## Overview

The `build-from-remote-repos.yml` workflow is designed to:

1. Build the Go framework code
2. Use templates, assets, pages, and configuration from multiple remote repositories
3. Build a Docker container for each remote repository
4. Push the containers to GitHub Container Registry

This allows you to maintain a single framework codebase while having multiple frontend applications with different templates, assets, pages, and configurations.

## How It Works

The workflow:

1. Runs automatically every 24 hours at midnight UTC, or can be manually triggered
2. Uses a matrix strategy to build from a hardcoded list of repositories under the `/r2unit/` user
3. For each repository:
   - Checks out the framework repository (the current repo)
   - Checks out the remote template repository (from `r2unit/{repo}`)
   - Sets up Go
   - Prepares a build directory by copying framework code and template files from the remote repository
   - Sets up Docker Buildx
   - Logs in to GitHub Container Registry
   - Extracts metadata for Docker
   - Builds and pushes a Docker image

## Configuration

### Hardcoded Repository List

The workflow uses a hardcoded list of repositories to check and build from. You can modify this list in the workflow file:

```yaml
strategy:
  matrix:
    # Hardcoded list of repositories under the /r2unit/ user to check and build from
    repo: [
      'frontend-main',
      'frontend-blog',
      'frontend-docs',
      'frontend-dashboard'
      # Add more repositories as needed
    ]
```

### Private Repositories

If the remote repositories are private, you'll need to use a Personal Access Token (PAT) with appropriate permissions. Uncomment and use the `token` parameter in the checkout step:

```yaml
- name: Checkout remote template repository
  uses: actions/checkout@v3
  with:
    repository: r2unit/${{ matrix.repo }}
    path: template-repo
    token: ${{ secrets.REPO_ACCESS_TOKEN }}
```

Make sure to add the `REPO_ACCESS_TOKEN` secret to your repository.

## Triggering the Workflow

### Automatic Triggering

The workflow is automatically triggered every 24 hours at midnight UTC via a scheduled cron job.

### Manual Triggering

You can also trigger the workflow manually from the GitHub Actions tab. When triggering manually, you can optionally specify a single repository to build:

1. Go to the Actions tab in your repository
2. Select "Build Frontend from Remote Repositories" from the workflows list
3. Click "Run workflow"
4. Optionally enter a repository name in the "Build only a specific repository" field
5. Click "Run workflow"

## Docker Images

The workflow builds Docker images with the following naming and tagging scheme:

### Image Names

Docker images are named after the repository name. For example, a repository named `frontend-blog` will produce a Docker image named `ghcr.io/frontend-blog`.

### Tagging Strategy

The workflow uses Semantic Versioning for Docker image tags:

1. **For the develop branch:**
   - Format: `develop-[SemVer]` (e.g., `develop-1.2.3`)
   - If no semantic versioning Git tags exist, fallback to `develop-[branch]-[sha]`

2. **For other branches (including main):**
   - Format: `[SemVer]` (e.g., `1.2.3`)
   - If no semantic versioning Git tags exist, fallback to `[branch]-[sha]`
   - The `latest` tag is also applied to images built from the main branch

To use semantic versioning, create Git tags following the pattern `v1.2.3` in your repository.

## Container Metadata

Each Docker image includes metadata that links it to the source GitHub repository:

- `org.opencontainers.image.title`: The repository name
- `org.opencontainers.image.description`: Description of the image
- `org.opencontainers.image.source`: URL to the GitHub repository
- `org.opencontainers.image.authors`: r2unit@proton.me
- `org.opencontainers.image.vendor`: r2unit

## Remote Repository Structure

Each remote repository should have the following structure:

```
frontend-your-project/
├── assets/            # Static files (CSS, JS, images)
├── config/            # Configuration files
│   └── routes.json    # Route configuration
├── pages/             # Page-specific templates
└── templates/         # Common templates
```

The workflow will copy these directories from the remote repository to the build directory.