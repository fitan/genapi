/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"log"

	"github.com/spf13/cobra"
)

// genApiCmd represents the genApi command
var genApiCmd = &cobra.Command{
	Use:   "gen-api",
	Short: "gen api template",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("src: %v, dest: %v", *src, *dest)
	},
}

func init() {
	rootCmd.AddCommand(genApiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genApiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genApiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
