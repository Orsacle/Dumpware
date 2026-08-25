package catalog

import "github.com/Orsacle/Dumpware/src/system"

type InstallCommand struct {
	Command string
	Args    []string
}

type Item struct {
	ID             string
	Name           string
	DescriptionKey string

	MacOS   InstallCommand
	Windows InstallCommand
	Linux   map[system.PackageManager]InstallCommand
}

func (i Item) CommandFor(os system.OS, pm system.PackageManager) (cmd InstallCommand, ok bool) {
	switch os {
	case system.MacOS:
		return i.MacOS, i.MacOS.Command != ""
	case system.Windows:
		return i.Windows, i.Windows.Command != ""
	case system.Linux:
		cmd, ok = i.Linux[pm]
		return cmd, ok
	default:
		return InstallCommand{}, false
	}
}

type Category struct {
	NameKey string
	Items   []Item
}

type Profile struct {
	NameKey string
	ItemIDs []string
}
