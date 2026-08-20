// dear god this is ai generated!!

package installer

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	mullvadKeyringURL  = "https://repository.mullvad.net/deb/mullvad-keyring.asc"
	mullvadKeyringPath = "/usr/share/keyrings/mullvad-keyring.asc"
	mullvadAptListPath = "/etc/apt/sources.list.d/mullvad.list"
	mullvadDnfRepoURL  = "https://repository.mullvad.net/rpm/stable/mullvad.repo"
)

// InstallMullvadApt adds Mullvad's official APT repository (Debian/Ubuntu
// and derivatives) and installs the package from it.
func InstallMullvadApt() error {
	codename, err := debianCodename()
	if err != nil {
		return fmt.Errorf("could not determine distro codename: %w", err)
	}

	arch, err := dpkgArch()
	if err != nil {
		return fmt.Errorf("could not determine dpkg architecture: %w", err)
	}

	if err := downloadFile(mullvadKeyringURL, mullvadKeyringPath); err != nil {
		return fmt.Errorf("could not fetch Mullvad signing key: %w", err)
	}

	repoLine := fmt.Sprintf(
		"deb [signed-by=%s arch=%s] https://repository.mullvad.net/deb/stable %s main\n",
		mullvadKeyringPath, arch, codename,
	)
	if err := os.WriteFile(mullvadAptListPath, []byte(repoLine), 0o644); err != nil {
		return fmt.Errorf("could not write %s: %w", mullvadAptListPath, err)
	}

	if err := runCmd("apt-get", "update"); err != nil {
		return fmt.Errorf("apt-get update failed: %w", err)
	}
	return runCmd("apt-get", "install", "-y", "mullvad-vpn")
}

// InstallMullvadDnf adds Mullvad's official RPM repository (Fedora) and
// installs the package from it.
func InstallMullvadDnf() error {
	// dnf5 (Fedora 41+) uses "addrepo"; dnf4 (Fedora 40 and earlier) uses the
	// legacy "--add-repo" flag instead. Try the modern syntax first, fall
	// back to the old one if that fails.
	err := runCmd("dnf", "config-manager", "addrepo", "--from-repofile="+mullvadDnfRepoURL)
	if err != nil {
		if err2 := runCmd("dnf", "config-manager", "--add-repo", mullvadDnfRepoURL); err2 != nil {
			return fmt.Errorf("could not add Mullvad repo (dnf5 attempt: %v; dnf4 attempt: %w)", err, err2)
		}
	}
	return runCmd("dnf", "install", "-y", "mullvad-vpn")
}

// debianCodename reads VERSION_CODENAME from /etc/os-release. Distros that
// don't rebase their codename onto Debian's (e.g. some Kali/Mint variants)
// may not have a matching entry on Mullvad's repo server; that's a Mullvad
// limitation, not something Dumpware can work around generically.
func debianCodename() (string, error) {
	f, err := os.Open("/etc/os-release")
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "VERSION_CODENAME=") {
			return strings.Trim(strings.TrimPrefix(line, "VERSION_CODENAME="), `"`), nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("VERSION_CODENAME not found in /etc/os-release")
}

func dpkgArch() (string, error) {
	out, err := exec.Command("dpkg", "--print-architecture").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
