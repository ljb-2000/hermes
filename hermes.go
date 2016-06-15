package main

func main() {
	supervisor := NewAgentSupervisor()
	RunWebServer(&supervisor)
}
