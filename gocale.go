package gocale

import (
  "encoding/json"
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
}

// TODO: Finish this
func NewLocalizer(config LocalizerConfig) Localizer {
  bundle := i18n.NewBundle(*config.defaultLocale)
  bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

  //for _, lang := range languages {
    
  //}

  return Localizer { bundle }
}
