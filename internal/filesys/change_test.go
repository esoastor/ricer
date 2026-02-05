package filesys_test

import (
	// "log"
	"log"
	"ricer/internal/filesys"
	"ricer/test/helpers"
	"strings"
	"testing"
)

func TestReplaceByChangeMap(t *testing.T) {
	subjectsFiles := helpers.GetSubjects()
	expectedFiles := helpers.GetExpected()

	current := helpers.GetSubjectCurrentTheme()
	applied := helpers.GetSubjectThemeTheme()

	changeMap := filesys.CreateChangeMap(current, applied)

	for key, value := range subjectsFiles {
		expected := expectedFiles[key].Content
		replaced := filesys.ReplaceByChangeMap(changeMap, value.Content, value.Path)

		expectedFlat := strings.ReplaceAll(expected, " ", "")
		expectedFlat = strings.ReplaceAll(expectedFlat, "\n", "")

		replacedFlat := strings.ReplaceAll(replaced, " ", "")
		replacedFlat = strings.ReplaceAll(replacedFlat, "\n", "")

		if expectedFlat != replacedFlat {
			t.Fatalf("NOT EQUAL:\n%v\nAND\n%v", replaced, expected)
		} else {
			log.Print("OK\n")
		}
		// log.Printf("exp: %v\nget: %v\n", expectedFlat, replacedFlat)
	}
	// log.Printf("subj: %v\n", changeMap)
}
