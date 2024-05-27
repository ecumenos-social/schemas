package constants

import (
	"context"
	_ "embed"

	"gopkg.in/yaml.v3"
)

type LanguageAlpha1Code string

type LanguageAlpha2Code string

type LanguageAlpha3Code string

type LanguageType string

type Language struct {
	Alpha1   LanguageAlpha1Code `json:"alpha1" yaml:"alpha1"`
	Alpha2   LanguageAlpha2Code `json:"alpha2" yaml:"alpha2"`
	Alpha3   LanguageAlpha3Code `json:"alpha3" yaml:"alpha3"`
	EngName  string             `json:"eng_name" yaml:"eng_name"`
	Type     LanguageType       `json:"type" yaml:"type"`
	MostUsed bool               `json:"most_used" yaml:"most_used"`
}

//go:embed languages.yaml
var languagesFile []byte

func LoadLanguages(ctx context.Context) ([]*Language, error) {
	var out []*Language
	if err := yaml.Unmarshal(languagesFile, &out); err != nil {
		return nil, err
	}

	return out, nil
}
