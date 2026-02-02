package theme_test

import (
	"log"
	"ricer/internal/theme"
	"ricer/test/helpers"
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
	keyPathMap := map[string]string{
		"globalVar1": "",
		"theme":      "/nvim/theme.lua",
		"var1":       "/some/folder/test.css",
		"var2":       "/some/folder/test.css",
		"var3":       "/some/folder/test.css",
		"glovalVar2": "",
	}
	expectedRows := 6

	themeFile := helpers.GetGoodTheme()

	testTheme := theme.GetTheme(themeFile)

	if len(testTheme) != expectedRows {
		t.Fatalf("Wrong number: %v != %v", expectedRows, len(testTheme))
	}

	for _, row := range testTheme {
		expected := keyPathMap[row.Key]
		got := row.Meta.Path
		if got != expected {
			t.Fatalf("Wrong path: \"%v\" != \"%v\"", expected, got)
		}
	}
}

// todo TestGetBadTheme

