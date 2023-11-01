package main

import (
	"fmt"
	"sync"
)

type Command interface {
	Execute()
}

type SingletonCommandManager struct {
	commands []Command
}

var instance *SingletonCommandManager
var once sync.Once

func GetCommandManagerInstance() *SingletonCommandManager {
	once.Do(func() {
		instance = &SingletonCommandManager{}
	})
	return instance
}

func (manager *SingletonCommandManager) AddCommand(command Command) {
	manager.commands = append(manager.commands, command)
}

func (manager *SingletonCommandManager) ExecuteCommands() {
	fmt.Println("Executing all registered commands:")
	for _, command := range manager.commands {
		command.Execute()
	}
}

type ConcreteCommand struct {
	message string
}

func (cc *ConcreteCommand) Execute() {
	fmt.Println(cc.message)
}

func main() {
	commandManager := GetCommandManagerInstance()

	commandManager.AddCommand(&ConcreteCommand{"Command 1"})
	commandManager.AddCommand(&ConcreteCommand{"Command 2"})
	commandManager.AddCommand(&ConcreteCommand{"Command 3"})

	commandManager.ExecuteCommands()
}
