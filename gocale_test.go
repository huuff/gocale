package gocale

import (
	"embed"
	"testing"

	"golang.org/x/text/language"
)

func TestEnglishAndSpanish(t *testing.T) {
  localizer := NewLocalizer(LocalizerConfig {
    DefaultLocale: &language.English,
    EnabledLocales: []string { "es", "en" },
    Path: "test_translations",
  }) 

  key := "test"
  englishTranslation, err := localizer.Translate(key, "en")

  if err != nil {
    t.Errorf("Error translating %s for en: %v", key, err)
  }

  spanishTranslation, err := localizer.Translate(key, "es")

  if err != nil {
    t.Errorf("Error translating %s for en: %v", key, err)
  }

  if englishTranslation != "test1" {
    t.Errorf("(%s for en) Expected: %s, Got: %s", key, "test1", englishTranslation)
  }

  if spanishTranslation != "prueba" {
    t.Errorf("(%s for es) Expected: %s, Got: %s", key, "prueba", spanishTranslation)
  }
}

//go:embed test_translations/*
var EmbeddedTranslations embed.FS

func TestEmbedded(t *testing.T) {
  localizer := NewLocalizer(LocalizerConfig {
    DefaultLocale: &language.English,
    EnabledLocales: []string { "es", "en" },
    EmbeddedFS: &EmbeddedTranslations,
    // TODO: It's a bit absurd that I need to provide the host path if its embedded... can't I work around this? Embed at the first level? Detect the directory name?
    Path: "test_translations",
  })

  key := "test"
  englishTranslation, err := localizer.Translate(key, "en")

  if err != nil {
    t.Errorf("Error translating %s for en: %v", key, err)
  }

  spanishTranslation, err := localizer.Translate(key, "es")

  if err != nil {
    t.Errorf("Error translating %s for en: %v", key, err)
  }

  if englishTranslation != "test1" {
    t.Errorf("(%s for en) Expected: %s, Got: %s", key, "test1", englishTranslation)
  }

  if spanishTranslation != "prueba" {
    t.Errorf("(%s for es) Expected: %s, Got: %s", key, "prueba", spanishTranslation)
  }
}
