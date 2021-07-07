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

//import (
//	"github.com/fitan/genapi/pkg"
//	"github.com/spf13/cobra"
//)
//
//// genEntCmd represents the genEnt command
//var genEntCmd = &cobra.Command{
//	Use:   "ent",
//	Short: "gen ent restful",
//	Long:  ``,
//	Run: func(cmd *cobra.Command, args []string) {
//		//pkg.Load(src, dest, nodes)
//		if *genEntV2 == "2" {
//			pkg.LoadV2(*genEntSrc, "./gen/entt2")
//		} else {
//			pkg.Load(*genEntSrc, *genEntDest)
//		}
//
//	},
//}
//
//var genEntSrc *string
//var genEntDest *string
//var genEntV2 *string
//
//func init() {
//	rootCmd.AddCommand(genEntCmd)
//	genEntSrc = genEntCmd.Flags().StringP("src", "s", "./ent/schema", "ent schema src.")
//	genEntDest = genEntCmd.Flags().StringP("dest", "d", "./gen/entt", "generate dest.")
//	genEntV2 = genEntCmd.Flags().StringP("version", "v", "2", "version")
//
//	// Here you will define your flags and configuration settings.
//
//	// Cobra supports Persistent Flags which will work for this command
//	// and all subcommands, e.g.:
//	// genEntCmd.PersistentFlags().String("foo", "", "A help for foo")
//
//	// Cobra supports local flags which will only run when this command
//	// is called directly, e.g.:
//	// genEntCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
//
//}
