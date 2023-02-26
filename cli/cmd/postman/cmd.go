package postman

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var (
	outputDirectory string

	Cmd = &cobra.Command{
		Use:   "postman",
		Short: "Postman tools",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	toThunderClientCmd = &cobra.Command{
		Use:   "thunder",
		Short: "Convert a collection from Postman to Thunder Client",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			raw, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatal(err)
			}

			var collection PostmanCollection
			json.Unmarshal(raw, &collection)

			tCollection, tClient := collection.ToThunderClientCollection()

			if err := tCollection.WriteFile(outputDirectory); err != nil {
				log.Fatal(err)
			}

			if err := tClient.WriteFile(outputDirectory); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	// Commands
	Cmd.AddCommand(toThunderClientCmd)

	toThunderClientCmd.Flags().StringVarP(&outputDirectory, "output-directory", "d", outputDirectory, "output directory")
	toThunderClientCmd.MarkFlagRequired("output-directory")
}
