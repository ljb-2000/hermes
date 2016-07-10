package main

import (
	"html/template"
	"io"
	"time"
)

type RulesPage struct {
	supervisor *AgentSupervisor
}

type RuleViewModel struct {
	Name           string
	TimeSinceEvent string
}

func NewRuleViewModel(agent RegisteredAgent) RuleViewModel {
	vm := RuleViewModel{}
	vm.Name = agent.name
	vm.TimeSinceEvent = "-"
	if agent.lastEventTime.Ok {
		vm.TimeSinceEvent = PrettyDuration(time.Since(agent.lastEventTime.Time))
	}
	return vm
}

func NewRulesPage(supervisor *AgentSupervisor) RulesPage {
	return RulesPage{supervisor}
}

func (p RulesPage) Rules() []RuleViewModel {
	agents := p.supervisor.Agents()
	vms := make([]RuleViewModel, len(agents))
	for i, agent := range agents {
		agent.mu.Lock()
		vms[i] = NewRuleViewModel(agent)
		agent.mu.Unlock()
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
