package gocale

import (
	"encoding/json"
	"fmt"

	"github.com/huuff/go-defaults"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type LocaleEncoding int
const (
  JsonLocaleEncoding LocaleEncoding = iota
)

type Localizer struct {
  bundle *i18n.Bundle
}

type LocalizerConfig struct {
  defaultLocale *language.Tag
  encoding *LocaleEncoding 
  enabledLocales []string
  path string
}

func NewLocalizer(config LocalizerConfig) Localizer {
  defaultLocale := defaults.DefaultPtr(config.defaultLocale, &language.English)
  enabledLocales := defaults.DefaultPtr[[]string](&config.enabledLocales, &[]string { defaultLocale.String() } )

  bundle := i18n.NewBundle(*defaultLocale)
  bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

  path := defaults.DefaultString(config.path, "/translations")
  for _, lang := range *enabledLocales {
    bundle.MustLoadMessageFile(fmt.Sprintf("%s/%s.json", path, lang))
  }

  return Localizer { bundle }
}

func (l Localizer) Translate(id, lang string) (string, error) {
  localizer := i18n.NewLocalizer(l.bundle, lang)

  cfg := &i18n.LocalizeConfig {
    DefaultMessage: &i18n.Message {
      ID: id,
      Other: id,
      One: id,
    },
  }

  str, err := localizer.Localize(cfg)
  if err != nil {
    return id, err
  }

  return str, nil
}

