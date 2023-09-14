package gocale

import (
  "testing"
  "golang.org/x/text/language"
)

func TestEnglishAndSpanish(t *testing.T) {
  localizer := NewLocalizer(LocalizerConfig {
    defaultLocale: &language.English,
    enabledLocales: []string { "es", "en" },
    path: "test_translations",
  }) 

  key := "test"
  englishTranslation, err := localizer.Get(key, "en")

  if err != nil {
    t.Errorf("Error translating %s for en: %v", key, err)
  }

  spanishTranslation, err := localizer.Get(key, "es")

  if err != nil {
    t.Errorf("Error translating %s for en: %v", key, err)
  }

  if englishTranslation != "test" {
    t.Errorf("(%s for en) Expected: %s, Got: %s", key, "test", englishTranslation)
  }

  if spanishTranslation != "prueba" {
    t.Errorf("(%s for es) Expected: %s, Got: %s", key, "prueba", spanishTranslation)
  }
}
