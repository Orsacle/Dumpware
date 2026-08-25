package catalog

func Profiles() []Profile {
	return []Profile{
		{
			NameKey: "profile.security",
			ItemIDs: []string{"wireshark", "mullvad", "bitwarden", "clamav", "p7zip"},
		},
		{
			NameKey: "profile.gaming",
			ItemIDs: []string{"steam", "obs"},
		},
		{
			NameKey: "profile.coding",
			ItemIDs: []string{"python", "nodejs", "php", "mysql", "git", "docker", "vim"},
		},
		{
			NameKey: "profile.terminal",
			ItemIDs: []string{"htop", "fastfetch", "cava", "cmatrix"}, // tty-clock not included
		},
		{
			NameKey: "profile.oob",
			ItemIDs: []string{"firefox", "neovim", "vlc", "libreoffice"},
		},
	}
}

func ItemsByIDs(ids []string) []Item {
	lookup := make(map[string]Item)
	for _, cat := range All() {
		for _, item := range cat.Items {
			lookup[item.ID] = item
		}
	}

	var result []Item
	for _, id := range ids {
		if item, ok := lookup[id]; ok {
			result = append(result, item)
		}
	}
	return result
}
