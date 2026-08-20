package system

import "os/exec"

type PackageManager int

const (
	UnknownPM PackageManager = iota
	Apt
	Dnf
	Pacman
	Zypper
	Apk
	Xbps
)

func (pm PackageManager) String() string {
	switch pm {
	case Apt:
		return "apt"
	case Dnf:
		return "dnf"
	case Pacman:
		return "pacman"
	case Zypper:
		return "zypper"
	case Apk:
		return "apk"
	case Xbps:
		return "xbps"
	default:
		return "unknown"
	}
}

func (pm PackageManager) binary() string {
	switch pm {
	case Apt:
		return "apt-get"
	case Dnf:
		return "dnf"
	case Pacman:
		return "pacman"
	case Zypper:
		return "zypper"
	case Apk:
		return "apk"
	case Xbps:
		return "xbps-install"
	default:
		return ""
	}
}

func allPackageManagers() []PackageManager {
	return []PackageManager{Apt, Dnf, Pacman, Zypper, Apk, Xbps}
}

func DetectPackageManagers() []PackageManager {
	var found []PackageManager
	for _, pm := range allPackageManagers() {
		if _, err := exec.LookPath(pm.binary()); err == nil {
			found = append(found, pm)
		}
	}
	return found
}
