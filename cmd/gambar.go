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
	"cari/cmd/gambar"
	"cari/cmd/helper"
	"errors"
	"github.com/spf13/cobra"
	"runtime"
)

// gambarCmd represents the gambar command
var gambarCmd = &cobra.Command{
	Use:   "gambar",
	Short: "download gambar dari page pertama berdasarkan keyword",
	Run: func(cmd *cobra.Command, args []string) {
		dir,_ := cmd.Flags().GetString("dir")
		category,_ := cmd.Flags().GetString("cat")
		page,err := cmd.Flags().GetInt("page")
		helper.PrintIfError(err)
		thread,err := cmd.Flags().GetInt("thread")
		helper.PrintIfError(err)

		if len(dir) == 0 || len(category) == 0 {
			panic(errors.New("parameter wajib diisi"))
		}
		gambar.VisitWeb(thread,page,dir,category)
	},
}

func init() {
	rootCmd.AddCommand(gambarCmd)
	gambarCmd.Flags().StringP("dir", "d", "", "tempat gambar di download")
	gambarCmd.Flags().StringP("cat", "c", "", "category gambar yang akan didownload")
	gambarCmd.Flags().IntP("page", "p", 1,"download gambar dari page 1 - <page>")
	gambarCmd.Flags().IntP("thread", "t", runtime.NumCPU(),"jumlah worker")
}
