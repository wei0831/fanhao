package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wei0831/fanhao/utli"
)

var cmdRename = &cobra.Command{
	Use:   "rename toFind",
	Short: "rename",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utli.Rename(dir, to, args[0], sperator, numbering, exclude, wet)
	},
}

func init() {
	rootCmd.AddCommand(cmdRename)
	cmdRename.PersistentFlags().BoolVarP(&wet, "wet", "w", false, "wet run, commit changes")
	cmdRename.PersistentFlags().StringVarP(&dir, "dir", "d", ".", "search target dir")
	cmdRename.PersistentFlags().StringVarP(&to, "to", "t", "", "move to target dir")
	cmdRename.PersistentFlags().StringVarP(&exclude, "exclude", "e", "", "exclude the pattern")
	cmdRename.PersistentFlags().StringVarP(&sperator, "sperator", "s", "-", "fanhao -sperator- number")
	cmdRename.PersistentFlags().StringVarP(&numbering, "numbering", "n", `\d\d\d`, "fanhao sperator -number-")
}
