package main

import (
	"html/template"
	"io"
)

type RulesPage struct {
	supervisor *AgentSupervisor
}

type RuleViewModel struct {
	Name          string
	LastEventTime MaybeTime
}

func NewRulesPage(supervisor *AgentSupervisor) RulesPage {
	return RulesPage{supervisor}
}

func (p RulesPage) Rules() []RuleViewModel {
	agents := p.supervisor.Agents()
	vms := make([]RuleViewModel, len(agents))
	for i, agent := range agents {
		vms[i] = RuleViewModel{agent.name, agent.lastEventTime}
	}
	return vms
}

func (p RulesPage) Render(w io.Writer) {
	tmpl, err := template.ParseFiles("templates/rules.html")
	if err != nil {
		panic("Could not read rules template")
	}

	tmpl.Execute(w, p)
}
