package cmd

type Command interface {
	Name() string
	Execute() int
}
