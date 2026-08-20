package main

import (
	"fmt"
	"os"

	"gitlab.com/Orsacle/Dumpware/i18n"
	"gitlab.com/Orsacle/Dumpware/src/catalog"
	"gitlab.com/Orsacle/Dumpware/src/installer"
	"gitlab.com/Orsacle/Dumpware/src/system"
	"gitlab.com/Orsacle/Dumpware/src/wizard"
)

func main() {
	lang := wizard.SelectLanguage(i18n.Available())
	i18n.SetLanguage(lang)

	fmt.Println(i18n.T("welcome.title"))
	fmt.Println()

	detectedOS := system.DetectOS()
	selectedOS := wizard.SelectOS(detectedOS)

	// only relevant for stinky people
	var pm system.PackageManager
	if selectedOS == system.Linux {
		available := system.DetectPackageManagers()
		pm = wizard.SelectPackageManager(available)
	}

	selectedIDs := make(map[string]bool)
	var profileItems []catalog.Item
	var profileName string
	var manualItems []catalog.Item

	if profile := wizard.SelectProfile(catalog.Profiles()); profile != nil {
		profileItems = catalog.ItemsByIDs(profile.ItemIDs)
		for _, item := range profileItems {
			selectedIDs[item.ID] = true
		}
		profileName = i18n.T(profile.NameKey)

		installNow := wizard.ShowProfileOverview(profileName, profileItems)
		if !installNow {
			manualItems = collectManualItems(selectedIDs)
		}
	} else {
		manualItems = collectManualItems(selectedIDs)
	}

	selected := append(append([]catalog.Item{}, profileItems...), manualItems...)

	if len(selected) == 0 {
		fmt.Println(i18n.T("prompt.nothingSelected"))
		return
	}

	if !wizard.Confirm(selectedOS, pm, profileName, manualItems) {
		fmt.Println(i18n.T("prompt.aborted"))
		return
	}

	if err := installer.Install(selectedOS, pm, selected); err != nil {
		fmt.Fprintln(os.Stderr, i18n.T("error.install"), err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println(i18n.T("success.done"))
}

func collectManualItems(selectedIDs map[string]bool) []catalog.Item {
	var items []catalog.Item
	for _, category := range catalog.All() {
		chosen := wizard.SelectItems(category)
		for _, item := range chosen {
			if !selectedIDs[item.ID] {
				items = append(items, item)
				selectedIDs[item.ID] = true
			}
		}
	}
	return items
}
