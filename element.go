package main

import (
	"fmt"
	"io"
	"strings"
)

type (
	A map[string]string

	HTMLElement struct {
		ID          string         `json:"id,omitempty"`
		TagName     string         `json:"tag_name"`
		Attributes  A              `json:"attributes,omitempty"`
		Classes     []string       `json:"classes,omitempty"`
		Text        string         `json:"text,omitempty"`
		Children    []*HTMLElement `json:"children,omitempty"`
		SelfClosing bool           `json:"self_closing,omitempty"`
	}
)

func NewHTMLElement(tagName string) *HTMLElement {
	return &HTMLElement{
		TagName:    tagName,
		Attributes: make(A),
	}
}

func (e *HTMLElement) AddChildren(children ...*HTMLElement) *HTMLElement {
	for _, child := range children {
		e.Children = append(e.Children, child)
	}
	return e
}

func (e *HTMLElement) AddAttributes(attributes A) *HTMLElement {
	for key, attribute := range attributes {
		e.Attributes[key] = attribute
	}
	return e
}

func (e *HTMLElement) AddClasses(classes ...string) *HTMLElement {
	for _, class := range classes {
		e.Classes = append(e.Classes, class)
	}
	return e
}

func (e *HTMLElement) SetID(id string) *HTMLElement {
	e.ID = id
	return e
}

func (e *HTMLElement) SetText(text string) *HTMLElement {
	e.Text = text
	return e
}

func (e *HTMLElement) Closing() *HTMLElement {
	e.SelfClosing = true
	return e
}

func (e *HTMLElement) renderAttributes() string {
	var builder strings.Builder
	if e.ID != "" {
		fmt.Fprintf(&builder, `id="%s" `, e.ID)
	}
	if len(e.Classes) > 0 {
		fmt.Fprintf(&builder, `class="%s" `, strings.Join(e.Classes, " "))
	}
	for key, value := range e.Attributes {
		fmt.Fprintf(&builder, `%s="%s" `, key, value)
	}
	attrStr := builder.String()
	return strings.TrimSpace(attrStr)
}

func (e *HTMLElement) Render(w io.Writer, depth int, minify bool) {
	delimeter, indent, textIndent := "\n", strings.Repeat("  ", depth), strings.Repeat("  ", depth+1)
	if minify {
		delimeter, indent, textIndent = "", "", ""
	}

	attrs := e.renderAttributes()
	if len(attrs) > 0 {
		attrs = " " + attrs
	}
	tagClosure := ">"
	if e.SelfClosing {
		tagClosure = "/>"
	}
	fmt.Fprintf(w, "%s<%s%s%s%s", indent, e.TagName, attrs, tagClosure, delimeter)

	if !e.SelfClosing && e.Text != "" {
		fmt.Fprintf(w, "%s%s%s", textIndent, e.Text, delimeter)
	}

	for _, child := range e.Children {
		child.Render(w, depth+1, minify)
	}

	if !e.SelfClosing {
		fmt.Fprintf(w, "%s</%s>%s", indent, e.TagName, delimeter)
	}
}
