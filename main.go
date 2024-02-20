package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

func Html() *HTMLElement  { return NewHTMLElement("html") }
func Head() *HTMLElement  { return NewHTMLElement("head") }
func Body() *HTMLElement  { return NewHTMLElement("body") }
func Title() *HTMLElement { return NewHTMLElement("title") }
func Div() *HTMLElement   { return NewHTMLElement("div") }
func Image() *HTMLElement { return NewHTMLElement("image").Closing() }

func timeTaken(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer timeTaken(time.Now(), "main")
	doc := Html().
		SetChildren(
			Head().
				SetChildren(
					Title().SetText("Example Page"),
				),
			Body().
				SetChildren(
					Div().
						SetID("content").
						SetAttributes(A{
							"hx-get":     "/get-me-some-sugar",
							"hx-trigger": "load",
						}).
						SetClasses(
							"container",
							"mx-auto",
							"px-4",
						).
						SetChildren(
							Div().
								SetClasses("htmx-indicator").
								SetText("Loading..."),
						),
					Image().
						SetClasses("w-96").
						SetAttributes(A{"src": "https://example.com/image.png"}),
				),
		)

	// render the document
	var sb strings.Builder
	doc.Render(&sb, 0, false)
	fmt.Println(sb.String())

	// Output json:
	jsonBytes, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		fmt.Println("Error serializing to JSON:", err)
		return
	}
	fmt.Println(string(jsonBytes))

	// Deserialize JSON back into the struct
	var root HTMLElement
	err = json.Unmarshal(jsonBytes, &root)
	if err != nil {
		fmt.Println("Error deserializing JSON:", err)
		return
	}
}
