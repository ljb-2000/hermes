package main

import "testing"

type testAgent struct{}

func (t testAgent) State() interface{} {
	return 0
}

func TestAgentsCanBeRegistered(t *testing.T) {
	supervisor := NewAgentSupervisor()
	agent := testAgent{}
	supervisor.Register(agent)
}

func TestRegisteredAgentsCanBeAccessed(t *testing.T) {
	supervisor := NewAgentSupervisor()
	agent := testAgent{}
	supervisor.Register(agent)

	registered := supervisor.Agents()

	if len(registered) != 1 {
		t.Errorf("wrong number of registed agents: %v\n", len(registered))
		return
	}

	if registered[0] != agent {
		t.Error("wrong agent registered")
	}
}
