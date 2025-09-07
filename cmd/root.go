package cmd

import (
	"cli_notes/internal/service"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const (
	idFlag    = "id"
	titleFlag = "title"
	tagFlag   = "tag"
	bodyFlag  = "body"
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

		fmt.Printf("note added successfully with id: %d\n", id)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all notes",
	Long:  "list all notes from db to console in order: ID, Title, Body, Tag",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := service.List()
		if err != nil {
			log.Fatal(err)
		}
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete note/notes",
	Long:  "delete note/notes by key (id, title, tag)",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		id, err := cmd.Flags().GetString("id")
		if err != nil {
			log.Fatal(err)
		}
		title, err := cmd.Flags().GetString("title")
		if err != nil {
			log.Fatal(err)
		}
		tag, err := cmd.Flags().GetString("tag")
		if err != nil {
			log.Fatal(err)
		}
		var res int64
		switch {
		case len(id) > 0:
			res, err = service.Delete(id, "id")
		case len(title) > 0:
			res, err = service.Delete(title, "title")
		case len(tag) > 0:
			res, err = service.Delete(tag, "tag")
		default:
			fmt.Println("no key provided")
		}
		if err != nil {
			log.Fatal(err)
		}
		if res == 0 {
			fmt.Println("no one note was found with this arguments")
		}
		fmt.Printf("notes was successfully deleted: %d\n", res)

	},
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit note by name",
	Long:  "edit note name, body or tag by name",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		title, err := cmd.Flags().GetString("t")
		if err != nil {
			log.Fatal(err)
		}

		id, err := cmd.Flags().GetString("i")
		if err != nil {
			log.Fatal(err)
		}

		var flag, arg string

		switch {
		case len(title) > 0:
			flag = "title"
			arg = title
		case len(id) > 0:
			flag = "id"
			arg = id
		default:
			log.Fatal("no searching flag given")
		}
		// getting value flag that will be changed
		var valueToEdit bool
		valueToEdit, err = cmd.Flags().GetBool(bodyFlag)
		if err != nil {
			log.Fatal(err)
		}

		if valueToEdit == true {
			_, err = service.Edit(flag, arg, bodyFlag)
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		valueToEdit, err = cmd.Flags().GetBool(titleFlag)
		if err != nil {
			log.Fatal(err)
		}

		if valueToEdit == true {
			_, err = service.Edit(flag, arg, titleFlag)
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		valueToEdit, err = cmd.Flags().GetBool(tagFlag)
		if err != nil {
			log.Fatal(err)
		}

		if valueToEdit == true {
			_, err = service.Edit(flag, arg, tagFlag)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		log.Fatal("no flag to edit given")

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
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(editCmd)

	deleteCmd.Flags().String(idFlag, "", "delete note by id")
	deleteCmd.Flags().String(titleFlag, "", "delete note by title")
	deleteCmd.Flags().String(tagFlag, "", "delete note by tag")

	editCmd.Flags().Bool(bodyFlag, false, "edit note's body")
	editCmd.Flags().Bool(titleFlag, false, "edit note's title")
	editCmd.Flags().Bool(tagFlag, false, "edit note's tag")

	editCmd.Flags().String("i", "", "search note by id")
	editCmd.Flags().String("t", "", "search note by title")

}
