package types

type Config struct {
	ThemesPath   string   `yaml:"themesPath"`
	SubjectPath  string   `yaml:"subjectsPath"`
	Exclude      []string `yaml:"exclude"`
	AfterCommand []string `yaml:"afterCommand"`
}

type ChangeMap struct {
	From     string
	To       string
	Code     string
	FilePath string
}

type ThemeFile struct {
	Path string
	Name string
}

type ThemeRow struct {
	Key   string
	Value string
	Meta  ThemeRowMeta
}

func (row *ThemeRow) FormId() string {
	return row.Key + row.Meta.Path
}

type ThemeRowMeta struct {
	Path string
}

type MultilineStruct struct {
	IsMultilineValue bool
	MultilineKey     string
	MultilineValue   string
}
