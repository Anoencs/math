/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// sumCmd represents the sum command
var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		isMinus, err := cmd.Flags().GetBool("minus")
		if err != nil {
			log.Panic(err)
		}
		var result int
		if len(args) >= 1 {
			result, err = strconv.Atoi(args[0])
		}
		if err != nil {
			log.Panic(err)
		}
		if isMinus {
			for i := 1; i < len(args); i++ {
				num, _ := strconv.Atoi(args[i])
				result -= num
			}
		} else {
			for i := 1; i < len(args); i++ {
				num, _ := strconv.Atoi(args[i])
				result += num
			}
		}
		fmt.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(sumCmd)
}
