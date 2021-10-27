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
	"cari/cmd/film"
	"fmt"
	"github.com/spf13/cobra"
)

// filmCmd represents the film command
var filmCmd = &cobra.Command{
	Use:   "film",
	Short: "cari film dari terminal",
	Long: `lelah cari film dari browser, coba cari dari terminal mungkin anda suka `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Mulai")
		title,_ := cmd.Flags().GetString("title")
		year,_ := cmd.Flags().GetString("year")

		if title != "" && year == "" {
			film.SearchFilm(title)
		} else if title == "" && year != "" {
			film.SearchByYear(year)
		}

	},
}

func init() {
	rootCmd.AddCommand(filmCmd)

	// berdasarkan judul
	filmCmd.Flags().StringP("title", "t", "", "example: -t after OR --title after")
	// berdasarkan tahun
	filmCmd.Flags().StringP("year", "y", "", "example: -y 2014 OR --year 2014")
}
