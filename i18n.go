package main

import (
	"fmt"

	"github.com/Unknwon/i18n"
)

func InitI18n(langs []string) error {
	for _, lang := range langs {
		i18n.SetMessage(lang, fmt.Sprintf("./langs/locale_%s.ini", lang))
	}
	return i18n.ReloadLangs(langs...)
}
