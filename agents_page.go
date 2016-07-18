package main

import (
	"html/template"
	"io"
	"time"
)

type AgentsPage struct {
	supervisor *AgentSupervisor
}

type AgentViewModel struct {
	Name           string
	TimeSinceEvent string
}

func NewAgentViewModel(agent RegisteredAgent) AgentViewModel {
	vm := AgentViewModel{}
	vm.Name = agent.name
	vm.TimeSinceEvent = "-"
	if agent.lastEventTime.Ok {
		vm.TimeSinceEvent = PrettyDuration(time.Since(agent.lastEventTime.Time))
	}
	return vm
}

func NewAgentsPage(supervisor *AgentSupervisor) AgentsPage {
	return AgentsPage{supervisor}
}

func (p AgentsPage) Agents() []AgentViewModel {
	agents := p.supervisor.Agents()
	vms := make([]AgentViewModel, len(agents))
	for i, agent := range agents {
		agent.mu.Lock()
		vms[i] = NewAgentViewModel(agent)
		agent.mu.Unlock()
	}
	return vms
}

func (p AgentsPage) Render(w io.Writer) {
	tmpl, err := template.ParseFiles("templates/agents.html")
	if err != nil {
		panic("Could not read agents template")
	}

	tmpl.Execute(w, p)
}
