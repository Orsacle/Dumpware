package catalog

import "gitlab.com/Orsacle/Dumpware/src/system"

func All() []Category {
	return []Category{
		{
			NameKey: "category.languages",
			Items: []Item{
				{
					ID:             "python",
					Name:           "Python",
					DescriptionKey: "item.python.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "python"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Python.Python.3.12"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "python3"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "python3"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "python"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "python3"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "python3"}},
					},
				},
				{
					ID:             "nodejs",
					Name:           "Node.js",
					DescriptionKey: "item.nodejs.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "node"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "OpenJS.NodeJS"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "nodejs", "npm"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "nodejs", "npm"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "nodejs", "npm"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "nodejs", "npm"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "nodejs", "npm"}},
					},
				},
				{
					ID:             "php",
					Name:           "PHP",
					DescriptionKey: "item.php.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "php"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "PHP.PHP.8.4"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "php"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "php"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "php"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "php"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "php84"}},
					},
				},
				{
					ID:             "rust",
					Name:           "Rust",
					DescriptionKey: "item.rust.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "rust"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Rustlang.Rustup"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "rustc", "cargo"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "rust", "cargo"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "rust"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "rust", "cargo"}},
						// no Zypper entry: openSUSE's rust/cargo live in the
						// devel:languages:rust OBS project, not Tumbleweed's
						// default repos - see LINUX_EXCEPTIONS
					},
				},
				{
					ID:             "go",
					Name:           "Go",
					DescriptionKey: "item.go.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "go"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "GoLang.Go"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "golang-go"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "golang"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "go"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "go"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "go"}},
					},
				},
				{
					ID:             "java",
					Name:           "Java (OpenJDK)",
					DescriptionKey: "item.java.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "openjdk"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "EclipseAdoptium.Temurin.21.JDK"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "default-jdk"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "java-21-openjdk-devel"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "jdk-openjdk"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "java-21-openjdk-devel"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "openjdk21"}},
					},
				},
				{
					ID:             "ruby",
					Name:           "Ruby",
					DescriptionKey: "item.ruby.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "ruby"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "RubyInstallerTeam.Ruby.3.4"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "ruby-full"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "ruby"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "ruby"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "ruby"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "ruby", "ruby-dev"}},
					},
				},
			},
		},
		{
			NameKey: "category.ides",
			Items: []Item{
				{
					ID:             "vscode",
					Name:           "Visual Studio Code",
					DescriptionKey: "item.vscode.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "visual-studio-code"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Microsoft.VisualStudioCode"}},
					// no distro ships an official repo package; Microsoft only publishes
					// via its own apt/yum repo, which Dumpware doesn't add automatically
				},
				{
					ID:             "vim",
					Name:           "Vim",
					DescriptionKey: "item.vim.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "vim"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Vim.Vim"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "vim"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "vim"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "vim"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "vim"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "vim"}},
					},
				},
				{
					ID:             "neovim",
					Name:           "Neovim",
					DescriptionKey: "item.neovim.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "neovim"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Neovim.Neovim"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "neovim"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "neovim"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "neovim"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "neovim"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "neovim"}},
					},
				},
				{
					ID:             "sublimetext",
					Name:           "Sublime Text",
					DescriptionKey: "item.sublimetext.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "sublime-text"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "SublimeHQ.SublimeText.4"}},
					// Linux builds require Sublime's own repo on every distro
					// (Arch: AUR-only) - no default-repo package anywhere
				},
			},
		},
		{
			NameKey: "category.browsers",
			Items: []Item{
				{
					ID:             "firefox",
					Name:           "Firefox",
					DescriptionKey: "item.firefox.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "firefox"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "firefox"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "firefox"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "firefox"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "firefox"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "firefox"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "firefox"}},
					},
				},
				{
					ID:             "chromium",
					Name:           "Chromium",
					DescriptionKey: "item.chromium.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "chromium"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Hibbiki.Chromium"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "chromium"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "chromium"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "chromium"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "chromium"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "chromium"}},
					},
				},
				{
					ID:             "brave",
					Name:           "Brave",
					DescriptionKey: "item.brave.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "brave-browser"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Brave.Brave"}},
					// every Linux distro needs Brave's own repo (or AUR on Arch) -
					// no default-repo package to install without adding one first
				},
			},
		},
		{
			NameKey: "category.databases",
			Items: []Item{
				{
					ID:             "mysql",
					Name:           "MySQL",
					DescriptionKey: "item.mysql.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "mysql"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Oracle.MySQL"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "mysql-server"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "mysql-server"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "mysql"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "mysql"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "mariadb", "mariadb-client"}},
					},
				},
				{
					ID:             "postgresql",
					Name:           "PostgreSQL",
					DescriptionKey: "item.postgresql.description",
					// pinned to 17 on brew/winget/apk: brew's unversioned
					// "postgresql" formula currently aliases to the
					// soon-disabled postgresql@14, and Alpine has no
					// unversioned "postgresql" package at all - only
					// versioned ones (postgresql16, postgresql17, ...).
					MacOS:   InstallCommand{Command: "brew", Args: []string{"install", "postgresql@17"}},
					Windows: InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "PostgreSQL.PostgreSQL.17"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "postgresql"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "postgresql-server"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "postgresql"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "postgresql-server"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "postgresql17"}},
					},
				},
				{
					ID:             "redis",
					Name:           "Redis",
					DescriptionKey: "item.redis.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "redis"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Redis.Redis"}},
					// Fedora, Arch and openSUSE dropped the "redis" package
					// after Redis's 2024 license change and now ship
					// Valkey (a BSD-licensed, protocol-compatible fork)
					// instead. Debian/Ubuntu kept shipping redis-server
					// as-is, so Apt is the only one still named "redis".
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "redis-server"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "valkey"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "valkey"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "valkey"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "valkey"}},
					},
				},
				{
					ID:             "sqlite",
					Name:           "SQLite",
					DescriptionKey: "item.sqlite.description",
					MacOS:   InstallCommand{Command: "brew", Args: []string{"install", "sqlite"}},
					Windows: InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "SQLite.SQLite"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "sqlite3"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "sqlite"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "sqlite"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "sqlite3"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "sqlite"}},
					},
				},
			},
		},
		{
			NameKey: "category.devtools",
			Items: []Item{
				{
					ID:             "git",
					Name:           "Git",
					DescriptionKey: "item.git.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "git"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Git.Git"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "git"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "git"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "git"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "git"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "git"}},
					},
				},
				{
					ID:             "docker",
					Name:           "Docker",
					DescriptionKey: "item.docker.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "docker"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Docker.DockerDesktop"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "docker.io"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "moby-engine"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "docker"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "docker"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "docker"}},
					},
				},
				{
					ID:             "tmux",
					Name:           "tmux",
					DescriptionKey: "item.tmux.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "tmux"}},
					// win zero val - tmux doesn't run natively on Windows
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "tmux"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "tmux"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "tmux"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "tmux"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "tmux"}},
					},
				},
				{
					ID:             "lazygit",
					Name:           "lazygit",
					DescriptionKey: "item.lazygit.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "lazygit"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "JesseDuffield.lazygit"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "lazygit"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "lazygit"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "lazygit"}},
					},
				},
				{
					ID:             "gh",
					Name:           "GitHub CLI",
					DescriptionKey: "item.gh.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "gh"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "GitHub.cli"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "gh"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "gh"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "github-cli"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "gh"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "github-cli"}},
					},
				},
			},
		},
		{
			NameKey: "category.creative",
			Items: []Item{
				{
					ID:             "inkscape",
					Name:           "Inkscape",
					DescriptionKey: "item.inkscape.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "inkscape"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Inkscape.Inkscape"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "inkscape"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "inkscape"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "inkscape"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "inkscape"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "inkscape"}},
					},
				},
				{
					ID:             "krita",
					Name:           "Krita",
					DescriptionKey: "item.krita.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "krita"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "KDE.Krita"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "krita"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "krita"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "krita"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "krita"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "krita"}},
					},
				},
				{
					ID:             "blender",
					Name:           "Blender",
					DescriptionKey: "item.blender.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "blender"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "BlenderFoundation.Blender"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "blender"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "blender"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "blender"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "blender"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "blender"}},
					},
				},
				{
					ID:             "audacity",
					Name:           "Audacity",
					DescriptionKey: "item.audacity.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "audacity"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Audacity.Audacity"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "audacity"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "audacity"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "audacity"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "audacity"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "audacity"}},
					},
				},
				{
					ID:             "kdenlive",
					Name:           "Kdenlive",
					DescriptionKey: "item.kdenlive.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "kdenlive"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "KDE.Kdenlive"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "kdenlive"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "kdenlive"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "kdenlive"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "kdenlive"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "kdenlive"}},
					},
				},
			},
		},
		{
			NameKey: "category.apps",
			Items: []Item{
				{
					ID:             "firefox",
					Name:           "Firefox",
					DescriptionKey: "item.firefox.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "firefox"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Mozilla.Firefox"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "firefox"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "firefox"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "firefox"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "firefox"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "firefox"}},
					},
				},
				{
					ID:             "thunderbird",
					Name:           "Thunderbird",
					DescriptionKey: "item.thunderbird.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "thunderbird"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Mozilla.Thunderbird"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "thunderbird"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "thunderbird"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "thunderbird"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "MozillaThunderbird"}}, // openSUSE kept the old package name
						system.Apk:    {Command: "apk", Args: []string{"add", "thunderbird"}},
					},
				},
				{
					ID:             "vlc",
					Name:           "VLC",
					DescriptionKey: "item.vlc.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "vlc"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "VideoLAN.VLC"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "vlc"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "vlc"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "vlc"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "vlc"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "vlc"}},
					},
				},
				{
					ID:             "wireshark",
					Name:           "Wireshark",
					DescriptionKey: "item.wireshark.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "wireshark"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "WiresharkFoundation.Wireshark"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "wireshark"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "wireshark"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "wireshark-qt"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "wireshark"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "wireshark"}},
					},
				},
				{
					ID:             "steam",
					Name:           "Steam",
					DescriptionKey: "item.steam.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "steam"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Valve.Steam"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "steam"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "steam"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "steam"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "steam"}},
					},
				},
				{
					ID:             "bitwarden",
					Name:           "Bitwarden",
					DescriptionKey: "item.bitwarden.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "bitwarden"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Bitwarden.Bitwarden"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "bitwarden"}},
					},
				},
				{
					ID:             "keepassxc",
					Name:           "KeePassXC",
					DescriptionKey: "item.keepassxc.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "keepassxc"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "KeePassXCTeam.KeePassXC"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "keepassxc"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "keepassxc"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "keepassxc"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "keepassxc"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "keepassxc"}},
					},
				},
				{
					ID:             "qbittorrent",
					Name:           "qBittorrent",
					DescriptionKey: "item.qbittorrent.description",
					// NOTE: Homebrew has scheduled ts cask to be disabled
					MacOS:   InstallCommand{Command: "brew", Args: []string{"install", "--cask", "qbittorrent"}},
					Windows: InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "qBittorrent.qBittorrent"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "qbittorrent"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "qbittorrent"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "qbittorrent"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "qbittorrent"}},
					},
				},
				{
					ID:             "obs",
					Name:           "OBS Studio",
					DescriptionKey: "item.obs.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "obs"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "OBSProject.OBSStudio"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "obs-studio"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "obs-studio"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "obs-studio"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "obs-studio"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "obs-studio"}},
					},
				},
				{
					ID:             "p7zip",
					Name:           "7-Zip",
					DescriptionKey: "item.p7zip.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "sevenzip"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "7zip.7zip"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "p7zip-full"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "p7zip"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "p7zip"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "p7zip"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "p7zip"}},
					},
				},
				{
					ID:             "htop",
					Name:           "htop",
					DescriptionKey: "item.htop.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "htop"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "htop"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "htop"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "htop"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "htop"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "htop"}},
					},
				},
				{
					ID:             "gimp",
					Name:           "GIMP",
					DescriptionKey: "item.gimp.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "gimp"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "GIMP.GIMP"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "gimp"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "gimp"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "gimp"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "gimp"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "gimp"}},
					},
				},
				{
					ID:             "mullvad",
					Name:           "Mullvad VPN",
					DescriptionKey: "item.mullvad.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "mullvad-vpn"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "MullvadVPN.MullvadVPN"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "mullvad-vpn"}}, // no package in the standard deb or rpm repos
					},
				},
				{
					ID:             "ventoy",
					Name:           "Ventoy",
					DescriptionKey: "item.ventoy.description",
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Ventoy.Ventoy"}},
					// mac and linux zero val
				},
				{
					ID:             "clamav",
					Name:           "ClamAV",
					DescriptionKey: "item.clamav.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "clamav"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Cisco.ClamAV"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "clamav"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "clamav"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "clamav"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "clamav"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "clamav"}},
					},
				},
				{
					ID:             "fastfetch",
					Name:           "fastfetch",
					DescriptionKey: "item.fastfetch.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "fastfetch"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Fastfetch-cli.Fastfetch"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "fastfetch"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "fastfetch"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "fastfetch"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "fastfetch"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "fastfetch"}},
						system.Xbps:   {Command: "xbps-install", Args: []string{"-S", "fastfetch"}},
					},
				},
				{
					ID:             "cava",
					Name:           "cava",
					DescriptionKey: "item.cava.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "cava"}},
					// win zero val
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "cava"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "cava"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "cava"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "cava"}},
					},
				},
				{
					ID:             "ttyclock",
					Name:           "tty-clock",
					DescriptionKey: "item.ttyclock.description",
					// mac and win zero val
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt: {Command: "apt-get", Args: []string{"install", "-y", "tty-clock"}},
						system.Dnf: {Command: "dnf", Args: []string{"install", "-y", "tty-clock"}},
						// pacman req aur helper
					},
				},
				{
					ID:             "libreoffice",
					Name:           "LibreOffice",
					DescriptionKey: "item.libreoffice.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "libreoffice"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "TheDocumentFoundation.LibreOffice"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "libreoffice"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "libreoffice"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "libreoffice-fresh"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "libreoffice"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "libreoffice"}},
					},
				},
				{
					ID:             "nextcloud",
					Name:           "Nextcloud Client",
					DescriptionKey: "item.nextcloud.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "--cask", "nextcloud"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Nextcloud.NextcloudDesktop"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "nextcloud-client"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "nextcloud-client"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "nextcloud-client"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "nextcloud-client"}},
						// no trusted src for apk
					},
				},
				{
					ID:             "syncthing",
					Name:           "Syncthing",
					DescriptionKey: "item.syncthing.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "syncthing"}},
					Windows:        InstallCommand{Command: "winget", Args: []string{"install", "-e", "--id", "Syncthing.Syncthing"}},
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "syncthing"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "syncthing"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "syncthing"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "syncthing"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "syncthing"}},
					},
				},
				{
					ID:             "cmatrix",
					Name:           "cmatrix",
					DescriptionKey: "item.cmatrix.description",
					MacOS:          InstallCommand{Command: "brew", Args: []string{"install", "cmatrix"}},
					// win zero val
					Linux: map[system.PackageManager]InstallCommand{
						system.Apt:    {Command: "apt-get", Args: []string{"install", "-y", "cmatrix"}},
						system.Dnf:    {Command: "dnf", Args: []string{"install", "-y", "cmatrix"}},
						system.Pacman: {Command: "pacman", Args: []string{"-S", "--noconfirm", "cmatrix"}},
						system.Zypper: {Command: "zypper", Args: []string{"install", "-y", "cmatrix"}},
						system.Apk:    {Command: "apk", Args: []string{"add", "cmatrix"}},
					},
				},
			},
		},
	}
}
