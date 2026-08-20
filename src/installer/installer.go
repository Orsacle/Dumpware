package installer

import (
	"fmt"
	"os"
	"os/exec"

	"gitlab.com/Orsacle/Dumpware/i18n"
	"gitlab.com/Orsacle/Dumpware/src/catalog"
	"gitlab.com/Orsacle/Dumpware/src/system"
)

func Install(selectedOS system.OS, pm system.PackageManager, items []catalog.Item) error {
	if selectedOS == system.Windows {
		if err := updateWingetSources(); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", i18n.T("installer.sourceUpdateFailed"), err)
		}
	}

	for _, item := range items {
		if selectedOS == system.Windows && item.ID == "php" {
			fmt.Printf("%s %s ...\n", i18n.T("installer.installing"), item.Name)
			if err := InstallPHPWindows(); err != nil {
				return fmt.Errorf("%s: %w", item.Name, err)
			}
			continue
		}

		if selectedOS == system.Linux && item.ID == "mullvad" && (pm == system.Apt || pm == system.Dnf) {
			fmt.Printf("%s %s ...\n", i18n.T("installer.installing"), item.Name)
			var err error
			switch pm {
			case system.Apt:
				err = InstallMullvadApt()
			case system.Dnf:
				err = InstallMullvadDnf()
			}
			if err != nil {
				return fmt.Errorf("%s: %w", item.Name, err)
			}
			continue
		}

		if selectedOS == system.Linux && item.ID == "ventoy" {
			fmt.Printf("%s %s ...\n", i18n.T("installer.installing"), item.Name)
			if err := InstallVentoyLinux(); err != nil {
				return fmt.Errorf("%s: %w", item.Name, err)
			}
			continue
		}

		cmd, ok := item.CommandFor(selectedOS, pm)
		if !ok {
			fmt.Printf(i18n.T("installer.noCommand")+"\n", item.Name)
			continue
		}

		fmt.Printf("%s %s ...\n", i18n.T("installer.installing"), item.Name)

		execCmd := exec.Command(cmd.Command, cmd.Args...)
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		execCmd.Stdin = os.Stdin

		if err := execCmd.Run(); err != nil {
			return fmt.Errorf("%s: %w", item.Name, err)
		}
	}
	return nil
}

// good measure
func updateWingetSources() error {
	fmt.Println(i18n.T("installer.updatingSources"))

	cmd := exec.Command("winget", "source", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
