package i18n

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed locales/*.json
var localeFS embed.FS

var (
	catalogs map[string]map[string]string
	current  = "en"
	fallback = "en"
)

func init() {
	catalogs = make(map[string]map[string]string)

	entries, err := localeFS.ReadDir("locales")
	if err != nil {
		return
	}

	for _, e := range entries {
		data, err := localeFS.ReadFile("locales/" + e.Name())
		if err != nil {
			continue
		}
		var m map[string]string
		if err := json.Unmarshal(data, &m); err != nil {
			continue
		}
		lang := e.Name()[:len(e.Name())-len(".json")]
		catalogs[lang] = m
	}
}

func Available() []string {
	langs := make([]string, 0, len(catalogs))
	if _, ok := catalogs[fallback]; ok {
		langs = append(langs, fallback)
	}
	for lang := range catalogs {
		if lang != fallback {
			langs = append(langs, lang)
		}
	}
	return langs
}

func SetLanguage(lang string) {
	if _, ok := catalogs[lang]; ok {
		current = lang
		return
	}
	current = fallback
}

func T(key string, args ...any) string {
	if msg, ok := catalogs[current][key]; ok {
		return format(msg, args...)
	}
	if msg, ok := catalogs[fallback][key]; ok {
		return format(msg, args...)
	}
	return key
}

func format(msg string, args ...any) string {
	if len(args) == 0 {
		return msg
	}
	return fmt.Sprintf(msg, args...)
}
