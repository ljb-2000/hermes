package main

import (
	"fmt"
	"time"
)

// AgentSupervisor contains many agents, and provides high-level access
// to aggregate actions on those agents.
type AgentSupervisor struct {
	agents []RegisteredAgent
}

// MaybeTime is an optional time value. No guaruntee is made about the content
// of the time if ok is false.
type MaybeTime struct {
	Ok   bool
	Time time.Time
}

// RegisteredAgent is an Agent, along with metadata related to the storage
// and operation of that agent, as required by AgentSupervisor.
type RegisteredAgent struct {
	Agent
	lastEventTime MaybeTime
	name          string
}

// Register adds an Agent to the set of supervised agents
func (s *AgentSupervisor) Register(theName string, theAgent Agent) {
	s.agents = append(s.agents, RegisteredAgent{
		Agent:         theAgent,
		lastEventTime: MaybeTime{},
		name:          theName,
	})
}

// Agents returns a slice of all supervised agents
func (s AgentSupervisor) Agents() []RegisteredAgent {
	return s.agents
}

// Run runs all agents, and the supervisor itself.
func (s AgentSupervisor) Run() {
	for _, agent := range s.agents {
		go s.recordEvents(&agent)
		go agent.Run()
	}
}

func (s *AgentSupervisor) recordEvents(theAgent *RegisteredAgent) {
	for _ = range theAgent.Events() {
		fmt.Println("got event from agent", theAgent.name)
		theAgent.lastEventTime = MaybeTime{true, time.Now()}
	}
}

// NewAgentSupervisor creates an agent supervisor
func NewAgentSupervisor() AgentSupervisor {
	return AgentSupervisor{}
}
