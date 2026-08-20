package system

import "runtime"

type OS int

const (
	Unknown OS = iota
	MacOS
	Linux
	Windows
)

func (o OS) String() string {
	switch o {
	case MacOS:
		return "macOS"
	case Linux:
		return "Linux"
	case Windows:
		return "Windows"
	default:
		return "Unknown"
	}
}

func AllOS() []OS {
	return []OS{MacOS, Linux, Windows}
}

func DetectOS() OS {
	switch runtime.GOOS {
	case "darwin":
		return MacOS
	case "linux":
		return Linux
	case "windows":
		return Windows
	default:
		return Unknown
	}
}
