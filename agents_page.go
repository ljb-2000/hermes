package main

import (
	"html/template"
	"io"
	"time"
)

type AgentsPage struct {
	lister AgentLister
}

type AgentViewModel struct {
	listing AgentListing
}

func (vm AgentViewModel) Name() string {
	return vm.listing.Name
}

// TimeSinceEvent returns a string formated representation of the time that
// has passed since an agent yielded an event.
func (vm AgentViewModel) TimeSinceEvent() string {
	if vm.listing.LastEventTime.Ok {
		return PrettyDuration(time.Since(vm.listing.LastEventTime.Time))
	}
	return "-"
}

func NewAgentsPage(lister AgentLister) AgentsPage {
	return AgentsPage{lister}
}

// Agents is the list of agents on tha page.
func (p AgentsPage) Agents() []AgentViewModel {
	listings := p.lister.ListAgents()
	vms := make([]AgentViewModel, len(listings))
	for i, listing := range listings {
		vms[i] = AgentViewModel{listing}
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
