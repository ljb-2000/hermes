package main

import (
	"fmt"
	"sync"
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
	agent         Agent
	lastEventTime MaybeTime
	name          string
	mu            sync.Mutex
}

func (r *RegisteredAgent) Name() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.name
}

func (r *RegisteredAgent) LastEventTime() MaybeTime {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.lastEventTime
}

// Register adds an Agent to the set of supervised agents
func (s *AgentSupervisor) Register(theName string, theAgent Agent) {
	s.agents = append(s.agents, RegisteredAgent{
		agent:         theAgent,
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
	for i := range s.agents {
		go s.recordEvents(&s.agents[i])
		go s.agents[i].agent.Run()
	}
}

func (s *AgentSupervisor) recordEvents(theAgent *RegisteredAgent) {
	for _ = range theAgent.agent.Events() {
		fmt.Println("got event from agent", theAgent.name)
		theAgent.mu.Lock()
		theAgent.lastEventTime = MaybeTime{true, time.Now()}
		theAgent.mu.Unlock()
	}
}

// NewAgentSupervisor creates an agent supervisor
func NewAgentSupervisor() AgentSupervisor {
	return AgentSupervisor{}
}
