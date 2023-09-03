package cmd

import "fmt"

type HelpCommand struct {
	name string
	app  *Application
}

func NewHelpCommand(app *Application) *HelpCommand {
	return &HelpCommand{name: "help", app: app}
}

func (cmd *HelpCommand) Name() string {
	return cmd.name
}

func (cmd *HelpCommand) Execute() int {
	fmt.Printf("%s is a tool to detect unexpected changes in ports, files, and kernel modules\n", cmd.app.Name())
	fmt.Println()
	fmt.Printf("Usage %s <command> [arguments]\n", cmd.app.Name())
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println()

	for _, command := range cmd.app.commands {
		fmt.Printf("\t%s\n", command.Name())
	}

	fmt.Println()
	fmt.Printf("Use %s help <topic> for more information about that topic.\n", cmd.app.Name())
	fmt.Println()

	return 0
}
