package main

import (
	"flag"
	"fmt"
	"github.com/AdityaAnandCodes/GoStats/api"
	"github.com/AdityaAnandCodes/GoStats/cli"
)


func main(){
	mode := flag.String("mode","cli","Mode:cli or api")
	flag.Parse()
	switch *mode{
		case "cli" :
			cli.LiveDashboard()
		case "api" :
			api.RunServer()
		default:
			fmt.Println("Invalid mode. Use --mode=cli or --mode=api")
	}  
}