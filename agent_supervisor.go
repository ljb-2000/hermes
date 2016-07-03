package main

// AgentSupervisor contains many agents, and provides high-level access
// to aggregate actions on those agents.
type AgentSupervisor struct {
	agents []Agent
}

// Register adds an Agent to the set of supervised agents
func (s *AgentSupervisor) Register(a Agent) {
	s.agents = append(s.agents, a)
}

// Agents returns a slice of all supervised agents
func (s AgentSupervisor) Agents() []Agent {
	return s.agents
}

// Run runs all agents, and the supervisor itself.
func (s AgentSupervisor) Run() {
	for _, agent := range s.agents {
		agent.Run()
	}
}

// NewAgentSupervisor creates an agent supervisor
func NewAgentSupervisor() AgentSupervisor {
	return AgentSupervisor{}
}
