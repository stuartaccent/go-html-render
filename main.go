package main

import (
	"log"
	"net/http"
)

func Html() *HTMLElement   { return NewHTMLElement("html") }
func Head() *HTMLElement   { return NewHTMLElement("head") }
func Title() *HTMLElement  { return NewHTMLElement("title") }
func Script() *HTMLElement { return NewHTMLElement("script") }
func Body() *HTMLElement   { return NewHTMLElement("body") }
func Div() *HTMLElement    { return NewHTMLElement("div") }
func Image() *HTMLElement  { return NewHTMLElement("image").Closing() }

var (
	root = Html().
		SetChildren(
			Head().
				SetChildren(
					Title().SetText("Example Page"),
					Script().
						SetAttributes(A{
							"src":         "https://unpkg.com/htmx.org@1.9.10",
							"integrity":   "sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC",
							"crossorigin": "anonymous",
						}),
				),
			Body().
				SetClasses(
					"container",
					"mx-auto",
					"px-4",
				).
				SetChildren(
					Div().
						SetID("hello").
						SetAttributes(A{
							"hx-get":     "/hello",
							"hx-trigger": "load delay:0.3s",
						}).
						SetChildren(
							Div().
								SetClasses("htmx-indicator").
								SetText("Loading..."),
						),
					Image().
						SetAttributes(A{
							"src":   "https://picsum.photos/800/600",
							"width": "800",
						}),
				),
		)

	hello = Div().SetText("Hello, World!")
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		root.Render(w, 0, true)
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		hello.Render(w, 0, true)
	})
	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
