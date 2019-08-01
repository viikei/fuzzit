/*
Copyright © 2019 fuzzit.dev, inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/fuzzitdev/fuzzit/client"
	"log"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <resource_path>",
	Short: "get information about targets or jobs",
	Long: `Example: 
	./fuzzit get targets # retrieve all targets
	./fuzzit get targets/<target_id> # retrieve specific target
	./fuzzit get targets/<target_id>/jobs # retrieve all jobs for target
	./fuzzit get targets/<target_id>/jobs/<job_id> # retrieve specific job`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.LoadFuzzitFromCache()
		if err != nil {
			log.Fatal(err)
		}
		err = c.GetResource(args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}