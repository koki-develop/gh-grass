package cmd

import (
	"fmt"
	"os"

	"github.com/cli/go-gh"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "grass",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("hi world, this is the gh-grass extension!")
		client, err := gh.RESTClient(nil)
		if err != nil {
			return err
		}
		response := struct{ Login string }{}
		err = client.Get("user", &response)
		if err != nil {
			return err
		}
		fmt.Printf("running as %s\n", response.Login)

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
