package constants

import (
	"context"
	"os"
	"path"
	"runtime"

	"gopkg.in/yaml.v3"
)

type CountryAlpha2Code string

type CountryAlpha3Code string

type CountryNumericCode string

type Country struct {
	EngName string             `json:"eng_name" yaml:"eng_name"`
	Alpha2  CountryAlpha2Code  `json:"alpha2" yaml:"alpha2"`
	Alpha3  CountryAlpha3Code  `json:"alpha3" yaml:"alpha3"`
	Numeric CountryNumericCode `json:"numeric" yaml:"numeric"`
}

func LoadCountries(ctx context.Context) ([]*Country, error) {
	_, filename, _, _ := runtime.Caller(1)
	currentDirname := preparePath(path.Dir(filename), "constants")
	yamlFile, err := os.ReadFile(currentDirname + "countries.yaml")
	if err != nil {
		return nil, err
	}

	var out []*Country
	if err := yaml.Unmarshal(yamlFile, &out); err != nil {
		return nil, err
	}

	return out, nil
}
