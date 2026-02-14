package theme_test

import (
	"ricer/internal/theme"
	"ricer/internal/types"
	"ricer/test/helpers"
	"strings"
	"testing"
)

func TestCreateChangeMap(t *testing.T) {
	controlChangeMap := []types.ChangeMap{
		{From: "#fbaded", To: "#fbad00", Code: "globalVar1", FilePath: ""},
		{From: "#1f2fff", To: "#10aaff", Code: "var1", FilePath: "coolSoft/styles.css"},
		{From: `{
        "some-theme",
        local params = test test params
    }`, To: `{
        "new-theme",
        local params = test test params
    }`, Code: "theme", FilePath: "coolSoft/styles.css"},
		{From: "testVal val val", To: "testVal val vaaaaaaal", Code: "var1", FilePath: "/some/folder/test.css"},
	}

	const changedKeysNumberFromGoodToDiffThemes = 4
	from := helpers.GetGoodTheme()
	to := helpers.GetDiffTheme()

	changeMap := theme.CreateChangeMap(from, to)

	changesLen := len(changeMap)
	if changesLen != changedKeysNumberFromGoodToDiffThemes {
		t.Fatalf("wrong changes number: %v (expecting %v)", changesLen, changedKeysNumberFromGoodToDiffThemes)
	}

	for index, change := range changeMap {
		control := controlChangeMap[index]
		fromOk := clearState(change.From) == clearState(control.From)
		toOk := clearState(change.To) == clearState(control.To)
		codeOk := change.Code == control.Code
		fileOk := change.FilePath == control.FilePath
		if fromOk && toOk && codeOk && fileOk {
			continue
		}
		t.Fatalf("\n%v != %v\n || \n%v != %v", clearState(change.From), clearState(control.From), clearState(change.To), clearState(control.To))
	}
}

func clearState(state string) string {
	state = strings.ReplaceAll(state, "\n", "")
	state = strings.ReplaceAll(state, " ", "")
	return state
}

func TestReplaceByChangeMap(t *testing.T) {
	subjectsFiles := helpers.GetSubjects()
	expectedFiles := helpers.GetExpected()

	current := helpers.GetSubjectCurrentTheme()
	applied := helpers.GetSubjectThemeTheme()

	changeMap := theme.CreateChangeMap(current, applied)

	for key, value := range subjectsFiles {
		expected := expectedFiles[key].Content
		replaced := theme.ReplaceByChangeMap(changeMap, value.Content, value.Path)

		expectedFlat := strings.ReplaceAll(expected, " ", "")
		expectedFlat = strings.ReplaceAll(expectedFlat, "\n", "")

		replacedFlat := strings.ReplaceAll(replaced, " ", "")
		replacedFlat = strings.ReplaceAll(replacedFlat, "\n", "")

		if expectedFlat != replacedFlat {
			t.Fatalf("NOT EQUAL:\n%v\nAND\n%v", replaced, expected)
		} 	
	}
}
