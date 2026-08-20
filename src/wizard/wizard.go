package wizard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gitlab.com/Orsacle/Dumpware/i18n"
	"gitlab.com/Orsacle/Dumpware/src/catalog"
	"gitlab.com/Orsacle/Dumpware/src/system"
)

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func SelectLanguage(available []string) string {
	fmt.Println("Select language:")
	for i, lang := range available {
		fmt.Printf("  [%d] %s\n", i+1, lang)
	}
	fmt.Print("> ")
	input := readLine()

	if idx, err := strconv.Atoi(input); err == nil && idx >= 1 && idx <= len(available) {
		return available[idx-1]
	}
	for _, lang := range available {
		if strings.EqualFold(input, lang) {
			return lang
		}
	}
	return "en"
}

func SelectOS(detected system.OS) system.OS {
	options := system.AllOS()

	fmt.Println(i18n.T("prompt.selectOS"))
	for i, o := range options {
		marker := ""
		if o == detected {
			marker = i18n.T("wizard.detected")
		}
		fmt.Printf("  [%d] %s%s\n", i+1, o, marker)
	}
	fmt.Print("> ")
	input := readLine()

	if input == "" && detected != system.Unknown {
		return detected
	}
	if idx, err := strconv.Atoi(input); err == nil && idx >= 1 && idx <= len(options) {
		return options[idx-1]
	}
	if detected != system.Unknown {
		return detected
	}
	return options[0]
}

func SelectPackageManager(available []system.PackageManager) system.PackageManager {
	if len(available) == 0 {
		fmt.Println(i18n.T("wizard.noPackageManagerFound"))
		return system.UnknownPM
	}

	fmt.Println(i18n.T("prompt.selectPackageManager"))
	for i, pm := range available {
		fmt.Printf("  [%d] %s\n", i+1, pm)
	}
	fmt.Print("> ")
	input := readLine()

	if idx, err := strconv.Atoi(input); err == nil && idx >= 1 && idx <= len(available) {
		return available[idx-1]
	}
	return available[0]
}

func SelectProfile(profiles []catalog.Profile) *catalog.Profile {
	if len(profiles) == 0 {
		return nil
	}

	fmt.Println()
	fmt.Println(i18n.T("prompt.selectProfile"))
	for i, p := range profiles {
		fmt.Printf("  [%d] %s\n", i+1, i18n.T(p.NameKey))
	}
	fmt.Println(i18n.T("wizard.profileHint"))
	fmt.Print("> ")
	input := readLine()

	if input == "" {
		return nil
	}
	if idx, err := strconv.Atoi(input); err == nil && idx >= 1 && idx <= len(profiles) {
		return &profiles[idx-1]
	}
	return nil
}

func ShowProfileOverview(profileName string, items []catalog.Item) bool {
	fmt.Println()
	fmt.Printf("== %s ==\n", profileName)
	for _, item := range items {
		fmt.Printf("  - %s\n", item.Name)
	}
	fmt.Println()
	fmt.Println(i18n.T("wizard.profileOverviewHint"))
	fmt.Print("> ")
	input := readLine()

	return input == "1"
}

func SelectItems(category catalog.Category) []catalog.Item {
	fmt.Println()
	fmt.Printf("== %s ==\n", i18n.T(category.NameKey))
	for i, item := range category.Items {
		fmt.Printf("  [%d] %s - %s\n", i+1, item.Name, i18n.T(item.DescriptionKey))
	}
	fmt.Println(i18n.T("wizard.selectItemsHint"))
	fmt.Print("> ")
	input := readLine()

	if input == "" {
		return nil
	}

	var selected []catalog.Item
	for _, part := range strings.Split(input, ",") {
		idx, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil || idx < 1 || idx > len(category.Items) {
			continue
		}
		selected = append(selected, category.Items[idx-1])
	}
	return selected
}

func Confirm(selectedOS system.OS, pm system.PackageManager, profileName string, manualItems []catalog.Item) bool {
	fmt.Println()
	fmt.Println(i18n.T("wizard.summaryHeader"))
	fmt.Printf("  OS: %s\n", selectedOS)
	if selectedOS == system.Linux {
		fmt.Printf("  Package Manager: %s\n", pm)
	}
	if profileName != "" {
		fmt.Printf("  - %s\n", i18n.T("wizard.profilePackage", profileName))
	}
	for _, item := range manualItems {
		fmt.Printf("  - %s\n", item.Name)
	}
	fmt.Println()
	fmt.Printf("%s %s ", i18n.T("prompt.confirmInstall"), i18n.T("wizard.yesNoHint"))

	input := strings.ToLower(readLine())
	return strings.HasPrefix(input, "y") || strings.HasPrefix(input, "j")
}
