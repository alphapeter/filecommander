package commands

import (
	"../cfg"
)

func Df() ([]string, error) {
	settings := cfg.GetSettings()
	rootNames := []string{}
	for i := 0; i < len(settings.Roots); i++ {
		rootNames = append(rootNames, settings.Roots[i].Name)
	}
	return rootNames, nil
}
