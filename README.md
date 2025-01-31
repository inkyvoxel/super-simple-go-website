# SUPER SIMPLE GO WEBSITE

This is a super simple starter template for websites written in [Go](https://go.dev/).

It tries to follow best practices from the [Standard Go Project Layout repository](https://github.com/golang-standards/project-layout).

How to run the project:

1. Clone the repo
1. `cd super-simple-go-website`
1. `go run cmd/super-simple-go-website/main.go`
1. Go to `http://localhost:8080/`

## Project structure

- `/cmd/super-simple-go-website/main.go` is the main entry point for the application. Register your HTTP handlers here.
- `/internal` directory is for Go packages intended to be used only within this project. Utility functions, business logic, etc.
- `/web/app` directory is for your application specific code, such as HTTP handlers, middleware, routing, etc.
- `/web/static` directory is for static files such as CSS, JavaScript, images, etc.
- `/web/template` directory is for HTML templates used when rendering views.
