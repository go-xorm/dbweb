package main

import (
	"fmt"
	"strings"

	"github.com/Unknwon/i18n"
)

func InitI18n(langs []string) error {
	for _, lang := range langs {
		i18n.SetMessage(lang, fmt.Sprintf("./langs/locale_%s.ini", strings.ToLower(lang)))
	}
	return i18n.ReloadLangs(langs...)
}
