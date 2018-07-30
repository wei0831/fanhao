package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wei0831/fanhao/utli"
)

var cmdMove = &cobra.Command{
	Use:   "move toFind",
	Short: "move",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utli.Move(dir, to, args[0], exclude, wet)
	},
}

func init() {
	rootCmd.AddCommand(cmdMove)
	cmdMove.PersistentFlags().BoolVarP(&wet, "wet", "w", false, "wet run, commit changes")
	cmdMove.PersistentFlags().StringVarP(&dir, "dir", "d", ".", "search target dir")
	cmdMove.PersistentFlags().StringVarP(&to, "to", "t", "", "move to target dir")
	cmdMove.PersistentFlags().StringVarP(&exclude, "exclude", "e", "", "exclude the pattern")
}
