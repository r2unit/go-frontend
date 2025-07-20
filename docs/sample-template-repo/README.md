# Sample Frontend Template Repository

This is a sample template repository for use with the frontend-framework. It demonstrates the structure and files needed to create a frontend application using the framework.

## Structure

This repository contains:

- `config/routes.json`: Configuration file for routes
- `templates/`: Common templates (header, footer)
- `pages/`: Page-specific templates (home, about, 404)
- `assets/`: Directory for static files (CSS, JS, images)

## How to Use

1. Clone this repository and rename it to match your project (e.g., `frontend-your-project`)
2. Modify the templates and pages to match your design
3. Update the routes.json file to define your routes
4. Add your static assets (CSS, JS, images) to the assets directory
5. Push to GitHub

The GitHub Workflow in the frontend-framework repository will automatically build a Docker container for your frontend when you push to the main or develop branch.

## Templates

The templates in this repository are simple examples. In a real project, you would:

1. Use proper Go template syntax (the comments in these files show where template code would go)
2. Create more complex templates with real content
3. Add more pages as needed

## Testing Locally

To test your frontend locally before pushing to GitHub, follow the instructions in the [Template Repository Structure](../template-repository-structure.md) documentation.

## Learn More

For more information about how to use the frontend-framework, see the [main documentation](../../README.md).