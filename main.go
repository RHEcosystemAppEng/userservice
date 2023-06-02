package main

import "userservice-go/env"

func main() {
	env.LoadEnvVars("./")
	InitializeAndStartServer()
}
