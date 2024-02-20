# Go HTML Render

This is a simple HTML renderer Kotlin style for Go. The beauty of this approach is its type safety.
JSON can be ripped out of the struct and used to rebuild it again. HTML becomes the product of the code, not the code itself.

Inspired by [ThePrimeTime](https://youtu.be/zJNkIJCQohU?si=feVfbkkfbRe_Afry).

## Usage

Declare blocks of HTML using the provided functions. 
```go
func Html() *HTMLElement  { return NewHTMLElement("html") }
func Head() *HTMLElement  { return NewHTMLElement("head") }
func Body() *HTMLElement  { return NewHTMLElement("body") }
func Title() *HTMLElement { return NewHTMLElement("title") }
func Div() *HTMLElement   { return NewHTMLElement("div") }
func Image() *HTMLElement { return NewHTMLElement("image").Closing() }
```

Set the attributes and children of the elements using the provided functions.
```go
doc := Html().
    AddChildren(
        Head().
            AddChildren(
                Title().SetText("Example Page"),
            ),
        Body().
            AddChildren(
                Div().
                    SetID("content").
                    AddAttributes(A{
                        "hx-get":     "/get-me-some-sugar",
                        "hx-trigger": "load",
                    }).
                    AddClasses(
                        "container",
                        "mx-auto",
                        "px-4",
                    ).
                    AddChildren(
                        Div().
                            AddClasses("htmx-indicator").
                            SetText("Loading..."),
                    ),
                Image().
                    AddClasses("w-96").
                    AddAttributes(A{"src": "https://example.com/image.png"}),
            ),
    )

// render the document
var sb strings.Builder
doc.Render(&sb, 0, false)
fmt.Println(sb.String())
```

The above code will produce the following HTML:
```html
<html>
  <head>
    <title>
      Example Page
    </title>
  </head>
  <body>
    <div id="content" class="container mx-auto px-4" hx-get="/get-me-some-sugar" hx-trigger="load">
      <div class="htmx-indicator">
        Loading...
      </div>
    </div>
    <image class="w-96" src="https://example.com/image.png"/>
  </body>
</html>
```
