package cmd

import (
	"cli_notes/internal/service"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add note (exact args position: title, body, tags)",
	Long:  "add note with title, body and tags. Exact args position: title, body, tags. Note gets ID in db",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		body := args[1]
		tag := args[2]

		id, err := service.Add(title, body, tag)
		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(1)
		}

		fmt.Println(id)
	},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli_notes",
	Short: "A brief description of your service",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your service. For example:

Cobra is a CLI library for Go that empowers applications.
This service is a tool to generate the needed files
to quickly create a Cobra service.`,
	// Uncomment the following line if your bare service
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your service.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli_notes.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(addCmd)
}
