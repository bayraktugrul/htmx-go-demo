# htmx-go-demo

This project is a demo application built using Go, Htmx, TailwindCSS, Go Templ, and Air.
It serves as a foundational template to demonstrate how these technologies
can be integrated to create dynamic and responsive web applications, providing a solid starting point
for developing applications using these technologies.

## Stack

| Technology |        Purpose        |               Doc                |
|:----------:|:---------------------:|:--------------------------------:|
|     Go     |        backend        |         https://go.dev/          |
|    Htmx    | dynamic interactivity |        https://htmx.org/         |
|  Tailwind  |     styling, css      |     https://tailwindcss.com      |
|   Templ    |  template generation  |   https://github.com/a-h/templ   |
|    Air     |      live reload      | https://github.com/air-verse/air |

## Install

### Air

```sh
go install github.com/air-verse/air@latest
```

### Tailwind

#### For Intel Mac (x64)

```sh
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-x64
chmod +x tailwindcss-macos-x64
mv tailwindcss-macos-x64 tailwindcss
```

#### For Apple Silicon (arm64)

```sh
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
chmod +x tailwindcss-macos-arm64
mv tailwindcss-macos-arm64 tailwindcss
```

## Run

#### Generate Templ
```sh
make templ-generate
```

#### Rebuild Tailwind
```sh
make tailwind-build
```

#### Run in Dev
```sh
make dev
```

#### Build
```sh
make build
```

#### Gif:
<img src="https://github.com/bayraktugrul/htmx-go-demo/blob/main/htmx.gif" width="500" />