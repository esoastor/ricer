package theme_test

import (
	// "log"
	"log"
	"ricer/internal/theme"
	"ricer/test/helpers"
	"slices"
	"testing"
)

func TestParseMetaStartGood(t *testing.T) {
	options := []string{
		"[ file  /sdf/asgd.asdg   ]",
		"[   file     /sdf/asgd.asdg   ]",
		"[file /sdf/asgd.asdg]",
	}

	for _, option := range options {
		tag, file := theme.ParseStartMeta(option)
		if tag != "file" {
			t.Fatalf("%v != %v", tag, "file")
		}
		if file != "/sdf/asgd.asdg" {
			t.Fatal("WRONG")
		}
	}
}

func TestParseMetaStartBad(t *testing.T) {
	options := []string{
		"[sfdgsdg]",
		"[asdfasdf sadfsadf asdfsadf]",
		"[",
	}

	for _, option := range options {
		tag, file := theme.ParseStartMeta(option)
		if tag != "" || file != "" {
			t.Fatal("must be empty")
		}
	}
}

func TestParseMetaEndGood(t *testing.T) {
	options := []string{
		"[endfile]",
		"[    end file     ]",
		"         [ end file   ]   ",
	}

	for _, option := range options {
		tag := theme.ParseEndMeta(option)
		if tag != "file" {
			t.Fatalf("invalid tag (%v)\nparsed \"%v\"", tag, option)
		}
	}
}

func TestParseMetaEndBad(t *testing.T) {
	options := []string{
		"[endfilfdsae]",
		"[    end file  sdafasfdhk    ]",
		"      safsf   [ end file   ]   ",
	}

	for _, option := range options {
		tag := theme.ParseEndMeta(option)
		if tag != "" {
			t.Fatalf("invalid tag (%v)\nparsed \"%v\"", tag, option)
		}
	}
}
func TestGetGoodTheme(t *testing.T) {
	keyPathMerges := []string{
		"globalVar1",
		"var1coolSoft/styles.css",
		"themecoolSoft/styles.css",
		"var1/some/folder/test.css",
		"var2/some/folder/test.css",
		"var3/some/folder/test.css",
		"glovalVar2",
	}
	expectedRows := 7

	themeFile := helpers.GetGoodTheme()

	testTheme, err := themeFile.FormTheme()

	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	if len(testTheme) != expectedRows {
		t.Fatalf("Wrong number: %v != %v", expectedRows, len(testTheme))
	}

	for _, row := range testTheme {
		exists := slices.Contains(keyPathMerges, row.Key+row.Meta.Path)
		if !exists {
			t.Fatalf("Row not found: \"%v\"", row)
		}
	}
}

func TestGetBadTheme(t *testing.T) {
	themeFile := helpers.GetBadTheme()
	_, err := themeFile.FormTheme()
	if err == nil {
		t.Fatalf("Should return error")
	}
}

