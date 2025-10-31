package handlers

import "markets-api/internal/environment"

type ClientEnvironment struct {
	*environment.ClientEnvironment
}

type ServerEnvironment struct {
	*environment.ServerEnvironment
}

func NewClientEnvironment() *ClientEnvironment {
	return &ClientEnvironment{
		ClientEnvironment: environment.NewClientEnvironment(),
	}
}

func NewServerEnvironment() *ServerEnvironment {
	return &ServerEnvironment{
		ServerEnvironment: environment.NewServerEnvironment(),
	}
}
