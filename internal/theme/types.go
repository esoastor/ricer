package theme

type ThemeFile struct {
	Path string
	Name string
}

func (tf ThemeFile) FormTheme() []ThemeRow {
	return GetTheme(tf)
}

type ThemeRow struct {
	Key string
	Value string
	Meta ThemeRowMeta
}

type ThemeRowMeta struct {
	Path string
}

