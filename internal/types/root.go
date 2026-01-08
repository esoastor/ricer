package types

type Theme map[string]string

type Config struct {
	ThemesPath string `yaml:"themesPath"`
	SubjectPath string `yaml:"subjectsPath"`
	Exclude []string `yaml:"exclude"`
}

