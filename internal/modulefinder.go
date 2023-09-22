package internal

type ModuleFinder struct {
}

type Module struct {
	name string
}

func NewModuleFinder() *PortFinder {
	return &PortFinder{}
}

func (list *ModuleFinder) FindModules() []Module {
	// Need to retrieve a list of all loaded kernel modules.
	return nil
}
