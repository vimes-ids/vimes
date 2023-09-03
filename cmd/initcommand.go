package cmd

type InitCommand struct {
	name string
}

func NewInitCommand() *InitCommand {
	return &InitCommand{name: "init"}
}

func (cmd *InitCommand) Name() string {
	return cmd.name
}

func (cmd *InitCommand) Execute() int {
	return 0
}
