package commands

import (
	"fmt"
	"ricer/internal/filesys"
	"ricer/internal/theme"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ricer",
	Short: "Uhhh idk, change some values to anothers \"via\" theme files...?",
	Args:  cobra.ExactArgs(1),
}

var listCmd = &cobra.Command{
	Use:   "themes",
	Short: "list of available themes",
	Run:   listThemes,
}

var listSubjectsCmd = &cobra.Command{
	Use:   "subjects",
	Short: "list of files that will be changed by ricer",
	Run:   listSubjectsOfChange,
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set theme by theme id",
	Run:   setTheme,
}

var showChangemapCmd = &cobra.Command{
	Use:   "changemap [id]",
	Short: "show changemap for theme (by theme id)",
	Run:   showChangemap,
	Args:  cobra.ExactArgs(1),
}

// Execute executes the root command.
func Execute() error {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(listSubjectsCmd)
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(showChangemapCmd)

	return rootCmd.Execute()
}

func listThemes(cmd *cobra.Command, args []string) {
	themes := theme.GetAll()
	for index, theme := range themes.Files {
		fmt.Printf("%v: %v\n", index+1, theme.Name)
	}
}

func setTheme(cmd *cobra.Command, args []string) {
	themeCol := theme.GetAll()
	subjTheme := themeCol.GetByName(args[0])
	theme.Submit(subjTheme)
}

func showChangemap(cmd *cobra.Command, args []string) {
	themeCol := theme.GetAll()
	subjTheme := themeCol.GetByName(args[0])
	changeMap := theme.CreateChangeMapForCurrent(subjTheme)
	for _, change := range changeMap {
		where := "all files"
		if len(change.FilePath) > 0 {
			where = change.FilePath
		}
		fmt.Printf("[%v] %v: %v => %v\n", where, change.Code, change.From, change.To)
	}
}

func listSubjectsOfChange(cmd *cobra.Command, args []string) {
	files := filesys.GetSubjectFiles()
	for _, file := range files {
		fmt.Printf("%v\n", file)
	}
}
