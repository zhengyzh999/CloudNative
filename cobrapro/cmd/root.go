package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobrapro",
	Short: "short desc",
	Long:  `long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd start")
		// 打印flag
		fmt.Println("root cmd end")
	},
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	var b *bool
	rootCmd.PersistentFlags().Bool("viper", true, "")
	rootCmd.PersistentFlags().StringP("author", "a", "Your Name", "")
	rootCmd.PersistentFlags().BoolVar(b, "v", true, "")
	rootCmd.Flags()
}
