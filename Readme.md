# Creating a Web Server in Go

This tutorial is intended for people with some experience writing Go who want to learn more about HTTP servers in Go. Some Go concepts are explained in detail and some, like how functions and variables are declared and used, are skipped.

## build a web server:

- Handle incoming requests.
- Route requests to different handlers.
- Add logging middleware—a layer between the OS and your software application that can process data.

## learning

:heavy_check_mark: How to serve static assets.

:heavy_check_mark: How to create an HTTP handler.

:heavy_check_mark: How to route requests dynamically.

:heavy_check_mark: How to create middleware.

## Requirements

- Go (version 1.11 or higher). You can download the latest version from the [Go site](https://golang.org/dl/).
- Knowledge of Go syntax and basic concepts.
- Access to the terminal window on your computer. The commands assume you're using Bash or an equivalent.
- A code editor.

## Run the server

To start, download the starting code

```bash
git clone https://github.com/StyvenSoft/webserver-go.git
```
Go to project directory:

```bash
cd webserver-go
```
Use the auto setup script:

```bash
bash server-setup.sh
```
The web server is now running at `localhost:8080`

---

## Reference docs

[HTTP package documentation](https://golang.org/pkg/net/http/)

func [ListenAndServe](https://golang.org/pkg/net/http/#ListenAndServe)

func [FileServer](https://golang.org/pkg/net/http/#ListenAndServe)

type [ServeMux](https://golang.org/pkg/net/http/#ServeMux)

func [StripPrefix](https://golang.org/pkg/net/http/#StripPrefix)
