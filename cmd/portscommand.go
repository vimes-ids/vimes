package cmd

import (
	"fmt"
	vimes "vimes/vimes"
)

type PortsCommand struct {
	name       string
	portFinder *vimes.PortFinder
}

func NewPortsCommand(portList *vimes.PortFinder) *PortsCommand {
	return &PortsCommand{name: "ports", portFinder: portList}
}

func (cmd *PortsCommand) Name() string {
	return cmd.name
}

func (cmd *PortsCommand) Execute() int {
	ports, err := cmd.portFinder.FindPorts()

	if err != nil {
		panic(err)
	}

	fmt.Println("The following ports are currently listening:")
	fmt.Println()

	for _, port := range ports {
		fmt.Printf("%s:%d (%s)\n", port.IP, port.Port, port.Protocol)
	}

	return 0
}
