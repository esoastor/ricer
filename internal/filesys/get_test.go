package filesys_test

import (
	// "log"
	"ricer/internal/filesys"
	"ricer/internal/types"
	"ricer/test/helpers"
	"strings"
	"testing"
)

func TestCreateChangeMap(t *testing.T) {
	controlChangeMap := []types.ChangeMap{
		{From: "#fbaded", To: "#fbad00", Code: "globalVar1", File: ""},
		{From: `{
        "some-theme",
        local params = test test params
    }`, To: `{
        "new-theme",
        local params = test test params
    }`, Code: "theme", File: "/nvim/theme.lua"},
		{From: "testVal val val", To: "testVal val vaaaaaaal", Code: "var1", File: "/some/folder/test.css"},
	}

	const changedKeysNumberFromGoodToDiffThemes = 3
	from := helpers.GetGoodTheme()
	to := helpers.GetDiffTheme()

	changeMap := filesys.CreateChangeMap(from, to)

	changesLen := len(changeMap)
	if changesLen != changedKeysNumberFromGoodToDiffThemes {
		t.Fatalf("wrong changes number: %v (expecting %v)", changesLen, changedKeysNumberFromGoodToDiffThemes)
	}

	for index, change := range changeMap {
		control := controlChangeMap[index]
		fromOk := clearState(change.From) == clearState(control.From)
		toOk := clearState(change.To) == clearState(control.To)
		codeOk := change.Code == control.Code
		fileOk := change.File == control.File
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
