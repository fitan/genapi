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
	"github.com/fitan/genapi/pkg/gen_mgr"
	public2 "github.com/fitan/genapi/public"
	"github.com/spf13/cobra"
	"log"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		switch *genType {
		case "ent":
			if *genName == "" {
				for _, ent := range public2.GetGenConf().Gen.Ent {
					gen_mgr.LoadV2(ent.Src, ent.Dest)
				}
				return
			}

			ent := public2.GetConfKey().GetEnt(*genName)
			if ent == nil {
				log.Panicln("Unknown ent name" + *genName)
			}
			gen_mgr.LoadV2(ent.Src,ent.Dest)
		case "api":
			//b , _ := json.Marshal(public2.GetGenConf())
			//spew.Dump(string(b))
			if *genName == "" {
				for _, api := range public2.GetGenConf().Gen.API {
					gen_mgr.DepthGen(api.Src,api.Dest, gen_mgr.GenApi)
				}
				return
			}
			api := public2.GetConfKey().GetApi(*genName)
			if api == nil {
				log.Panicln("Unknown api name" + *genName)
			}

			gen_mgr.DepthGen(api.Src, api.Dest, gen_mgr.GenApi)
		case "ts":
			if *genName == "" {
				for _, ts := range public2.GetGenConf().Gen.Ts {
					gen_mgr.DepthGen(ts.Src, ts.Dest, gen_mgr.GenTs)
				}
				return
			}
			ts := public2.GetConfKey().GetTs(*genName)
			if ts == nil {
				log.Panicln("Unknown ts name" + *genName)
			}

			gen_mgr.DepthGen(ts.Src, ts.Dest, gen_mgr.GenTs)

		default:
		}

	},
}

var genType *string
var genName *string
func init() {
	rootCmd.AddCommand(genCmd)
	genType = genCmd.Flags().StringP("type", "t", "", "gen type: ent,api")
	genCmd.MarkFlagRequired("type")
	genName = genCmd.Flags().StringP("name", "n", "", "gen name")


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
