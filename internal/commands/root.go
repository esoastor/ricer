package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"ricer/internal/filesys"
	"strconv"
)

var rootCmd = &cobra.Command{
	Use:   "ricer",
	Short: "Uhhh idk, change some values to anothers \"via\" theme files...?",
	Args:  cobra.ExactArgs(1),
}

var listCmd = &cobra.Command{
	Use:   "themes",
	Short: "list available themes",
	Run:   listThemes,
}

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "show current theme",
	Run:   showCurrentTheme,
}

var listSubjectsCmd = &cobra.Command{
	Use:   "subjects",
	Short: "how files that gonna change",
	Run:   listSubjectsOfChange,
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set theme by number",
	Run:   setTheme,
}

var showChangemapCmd = &cobra.Command{
	Use:   "changemap [id]",
	Short: "set theme by number",
	Run:   showChangemap,
	Args:  cobra.ExactArgs(1),
}

// Execute executes the root command.
func Execute() error {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(currentCmd)
	rootCmd.AddCommand(listSubjectsCmd)
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(showChangemapCmd)

	return rootCmd.Execute()
}

func listThemes(cmd *cobra.Command, args []string) {
	themes := filesys.GetThemes()
	for index, theme := range themes {
		fmt.Printf("%v: %v\n", index+1, theme.Name)
	}
}

func showCurrentTheme(cmd *cobra.Command, args []string) {
	cur := filesys.GetCurrentTheme()
	fmt.Printf("%v\nwhat?? its always current. Nonsense\n", cur)
}

func setTheme(cmd *cobra.Command, args []string) {
	themes := filesys.GetThemes()
	index, err := strconv.Atoi(args[0])
	if err != nil || len(themes) < index {
		fmt.Print("error")
		panic("aa")
	}
	theme := themes[index-1]

	filesys.SubmitTheme(theme)
}

func showChangemap(cmd *cobra.Command, args []string) {
	themes := filesys.GetThemes()
	index, err := strconv.Atoi(args[0])
	if err != nil || len(themes) < index {
		fmt.Print("error")
		panic("aa")
	}
	theme := themes[index-1]
	// todo: failed to get current theme
	changeMap := filesys.CreateChangeMapForCurrent(theme)
	for _, change := range changeMap {
		where := "all files"
		if len(change.File) > 0 {
			where = change.File
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
