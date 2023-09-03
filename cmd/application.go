package cmd

type Application struct {
	name     string
	commands map[string]Command
}

func NewApplication() *Application {
	app := &Application{name: "vimes", commands: make(map[string]Command)}

	return app
}

func (app *Application) AddCommand(cmd Command) {
	app.commands[cmd.Name()] = cmd
}

func (app *Application) Name() string {
	return app.name
}

func (app *Application) Run(args []string) int {
	for _, command := range app.commands {
		if len(args) > 0 && command.Name() == args[0] {
			return command.Execute()
		}
	}

	return app.commands["help"].Execute()
}
