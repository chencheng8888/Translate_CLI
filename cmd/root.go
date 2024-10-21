package cmd

import (
	"fmt"
	"os"
	useapi "translate_cli/UseAPI"

	"github.com/spf13/cobra"
)
var tranPattern string 

var rootCmd = &cobra.Command{
	Use: "tran",
	Short: "tran is a tool which help you translate a word or a sentence(English) to Chinese",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			from string
			to string
			txt = args[0]
		)
		switch tranPattern {
		case "":
			from = "en"
			to = "zh-CHS"
		case "ec":
			from = "en"
			to = "zh-CHS"
		case "ce":
			from = "zh-CHS"
			to = "en"
		default:
			fmt.Println("the pattern is Illegal")	
			return		
		}
		if txt == ""{
			fmt.Println("the txt should not be null")
			return
		}
		res,err := useapi.Translate(txt,from,to)
		if err!=nil {
			fmt.Printf("The translation failed, maybe because of the network, or maybe because of your configuration error,reason:%v",err)
			return
		}
		for _,v := range res{
			fmt.Println(v)
		}
	},
}
func init(){
	rootCmd.Flags().StringVarP(&tranPattern,"pattern","p","ec","ec is \"English\" to \"Chinese\",ce is \"Chinese\" to \"English\"")
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
}