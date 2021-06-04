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
	"github.com/fitan/genapi/pkg"
	"github.com/fitan/genapi/pkg/gen_apiV2"
	"github.com/marcinwyszynski/directory_tree"
	"github.com/spf13/cobra"
	"path"
)

func DepthGen(tree *directory_tree.Node, Dir string)  {
	context := gen_apiV2.NewApiContext()
	context.Load(tree.FullPath)
	context.Parse()
	for _, file := range context.Files {
		if len(file.Funcs) != 0 {
			pkg.GenApiV2(context.Files, Dir)
			break
		}
	}

	for _,node := range tree.Children {
		if node.Info.IsDir {
			DepthGen(node,path.Join(Dir, node.Info.Name))
		}
	}
}

// router2Cmd represents the router2 command
var router2Cmd = &cobra.Command{
	Use:   "router2",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tree, err := directory_tree.NewTree("./logic")
		if err != nil {
			panic(err)
		}

		DepthGen(tree, "./gen2/handler")



		//for fileName, f := range context.Files {
		//	fmt.Println("fileName:  ",fileName)
		//	fmt.Printf("%# v", pretty.Formatter(f.Funcs))
		//	spew.Dump(f.Funcs)
		//}
	},
}

func init() {
	rootCmd.AddCommand(router2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// router2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// router2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
