package main

import (
	"html/template"
	"io"
)

type RulesPage struct{}

func NewRulesPage() RulesPage {
	return RulesPage{}
}

func (p RulesPage) RuleNames() []string {
	return []string{
		"XKCD",
		"Simplycast",
	}
}

func (p RulesPage) Render(w io.Writer) {
	tmpl, err := template.ParseFiles("templates/rules.html")
	if err != nil {
		panic("Could not read rules template")
	}

	tmpl.Execute(w, p)
}
