package gocale

import (
	"encoding/json"
	"fmt"

	"github.com/huuff/go-defaults"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
  "embed"
)

type LocaleEncoding int
const (
  JsonLocaleEncoding LocaleEncoding = iota
)

type Localizer struct {
  bundle *i18n.Bundle
}

type LocalizerConfig struct {
  DefaultLocale *language.Tag
  Encoding *LocaleEncoding 
  EnabledLocales []string
  Path string
  EmbeddedFS *embed.FS
}

func NewLocalizer(config LocalizerConfig) Localizer {

  defaultLocale := defaults.DefaultPtr(config.DefaultLocale, &language.English)
  enabledLocales := defaults.DefaultPtr[[]string](&config.EnabledLocales, &[]string { defaultLocale.String() } )

  bundle := i18n.NewBundle(*defaultLocale)
  bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

  path := defaults.DefaultString(config.Path, "/translations")
  for _, lang := range *enabledLocales {
    if config.EmbeddedFS != nil {
      bundle.LoadMessageFileFS(config.EmbeddedFS, fmt.Sprintf("%s/%s.json", path, lang))
    } else {
      bundle.MustLoadMessageFile(fmt.Sprintf("%s/%s.json", path, lang))
    } 
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

func (l Localizer) TranslateAll(ids []string, lang string) (map[string]string, error) {
  results := make(map[string]string) 

  for _, id := range ids {
    translation, err := l.Translate(id, lang)

    if err != nil {
      return nil, err
    }

    results[id] = translation
  }

  return results, nil
}
