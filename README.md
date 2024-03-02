# An example of Authentication with Go

This projects shows how to build a minimal full stack web app using Go, HTMX, Templ and Tailwindcss **with Authentication (OAuth)**.

Checkout the [Full Stack GO](https://github.com/sikozonpc/fullstack-go-htmx) project to learn more about the project structure.

## Structure

All the HTML are stored in `*.templ` files in the `/views` and `/components` directories.
The `/handlers` directory contains the Go handlers for the different routes that serve those Templ components.

## Installation

There are a few tools that you need to install to run the project.
So make sure you have the following tools installed on your machine.

- [Templ (for the UI layer)](https://templ.guide/quick-start/installation)
- [Tailwindcss CLI (CSS styling library)](https://tailwindcss.com/docs/installation)
- [Migrate (for DB migrations)](https://github.com/golang-migrate/migrate/tree/v4.17.0/cmd/migrate)
- [Air (for live reloading)](https://github.com/cosmtrek/air)

Adittionally, it's recommended to install a syntax highlighting and templ LSP integration:
[the official Templ documentation](https://templ.guide/quick-start/installation#editor-support).

## Running the project

Firstly make sure you have a MySQL database running on your machine or just swap for any storage you like under `/store`.

Don't forget to check the `.env.example` file and inject those environment variables into your environment, it's optimized for the cloud enviroment so it's recommended to inject them at runtime, for example using [direnv](https://direnv.net/).

> If you want to inject them manually into a `.env`, install go-dotenv and adjust the `config/config.go` file to read the `.env` file instead.

Then, for the best development experience, run the project in 3 different processes by running:

```bash
air # for the go server live reloading
make tailwind # for the tailwindcss live reloading
make templ # for the templ files live generation and reloading
```
