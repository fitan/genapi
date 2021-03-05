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
	"github.com/spf13/cobra"
	"log"
)

// genEntCmd represents the genEnt command
var genEntCmd = &cobra.Command{
	Use:   "gen-ent",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//pkg.Load(src, dest, nodes)
		log.Printf("src: %v, dest: %v, nodes: %v", *src, *dest, *nodes)
	},
}

var src *string
var dest *string
var nodes *[]string

func init() {
	rootCmd.AddCommand(genEntCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genEntCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genEntCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	nodes = genEntCmd.Flags().StringSliceP("node", "n", nil, "ent schema node")
	src = genEntCmd.Flags().StringP("src", "s", "./ent/schema", "ent schema src")
	dest = genEntCmd.Flags().StringP("dest", "d", "./gen", "generate dest")
}
