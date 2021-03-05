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
	"fmt"

	"regexp"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate a input string vs regex expression",
	Long:  `valid a string input according a regex expression`,
	Run: func(cmd *cobra.Command, args []string) {
		str := args[0]
		regexDel := "(\\.bex-dominio-stg.svc:8080)"
		regexComp := "([A-Z]|[a-z]|-|)+(\\.bex-dominio-stg.svc:8080$)"
		fmt.Println(str)
		match, _ := regexp.MatchString(regexComp, str)

		if match {
			fmt.Println("hay coincidencia")
			strConcat, _ := cmd.Flags().GetString("cadena")
			if strConcat != "" {
				fmt.Println(strConcat)
			} else {
				strConcat = "http://localhost:8080/"
			}
			str := args[0]
			rDel, _ := regexp.Compile(regexDel)
			strFinal := strConcat + rDel.ReplaceAllString(str, "")
			fmt.Println(strFinal)
		} else {
			fmt.Println("no hay coincidencia")
		}

	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringP("cadena", "c", "http://localhost:8080/", `Sera la cadena a concatenar`)
	validateCmd.MarkFlagRequired("cadena")
}
