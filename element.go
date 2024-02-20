package main

import (
	"fmt"
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

func (e *HTMLElement) SetChildren(children ...*HTMLElement) *HTMLElement {
	e.Children = children
	return e
}

func (e *HTMLElement) SetAttributes(attributes A) *HTMLElement {
	e.Attributes = attributes
	return e
}

func (e *HTMLElement) SetClasses(classes ...string) *HTMLElement {
	e.Classes = classes
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
	var attrs []string
	if len(e.Classes) > 0 {
		attrs = append(attrs, fmt.Sprintf(`class="%s"`, strings.Join(e.Classes, " ")))
	}
	for key, value := range e.Attributes {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, key, value))
	}
	return strings.Join(attrs, " ")
}

func (e *HTMLElement) Render(sb *strings.Builder, depth int, minify bool) {
	newLine := "\n"
	indent := strings.Repeat("  ", depth)
	textIndent := strings.Repeat("  ", depth+1)
	if minify {
		newLine = ""
		indent = ""
		textIndent = ""
	}

	attrs := e.renderAttributes()
	if attrs != "" {
		attrs = " " + attrs
	}

	id := e.ID
	if id != "" {
		id = fmt.Sprintf(` id="%s"`, id)
	}

	tagFormat := "%s<%s%s%s%s>%s"

	if e.SelfClosing {
		sb.WriteString(fmt.Sprintf(tagFormat, indent, e.TagName, id, attrs, "/", newLine))
	} else {
		sb.WriteString(fmt.Sprintf(tagFormat, indent, e.TagName, id, attrs, "", newLine))
		if e.Text != "" {
			sb.WriteString(fmt.Sprintf("%s%s%s", textIndent, e.Text, newLine))
		}
		for _, child := range e.Children {
			child.Render(sb, depth+1, minify)
		}
		sb.WriteString(fmt.Sprintf("%s</%s>%s", indent, e.TagName, newLine))
	}
}
