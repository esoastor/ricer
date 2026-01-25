package types

type Config struct {
	ThemesPath string `yaml:"themesPath"`
	SubjectPath string `yaml:"subjectsPath"`
	Exclude []string `yaml:"exclude"`
}

type ChangeMap struct {
	From string
	To string
}
