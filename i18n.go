package main

import (
	"fmt"
	"strings"

	"github.com/Unknwon/i18n"
	options "github.com/go-xorm/dbweb/modules/options"
)

func InitI18n(langs []string) error {
	for _, lang := range langs {
		data, err := options.Locale(fmt.Sprintf("locale_%s.ini", strings.ToLower(lang)))
		if err != nil {
			return err
		}
		i18n.SetMessage(lang, data)
	}
	return i18n.ReloadLangs(langs...)
}
