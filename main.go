package main

import (
	"fmt"
	"jugo/jugo"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func init() {
	// Check Dirs
	wd, _ := os.Getwd()
	var dirs = [2]string{"log", "result"}
	for _, value := range dirs {
		if _, err := os.Stat(filepath.Join(wd, value)); err != nil {
			// Dir IsNotExist
			if os.IsNotExist(err) {
				os.MkdirAll(value, os.ModePerm)
			}
		}
	}
	// Check config.toml
	if _, err := os.Stat(filepath.Join(wd, "config.toml")); err != nil {
		// config.toml IsNotExist
		if os.IsNotExist(err) {
			//os.Create("config.toml")
			file, ferr := os.OpenFile("config.toml", os.O_WRONLY|os.O_CREATE, 0666)
			if ferr != nil {
				log.Fatal(ferr)
			}
			defer file.Close()
			// config.toml Content
			str := "[proxy]\n\n\n[api]\n"
			file.WriteString(str)
		}
	}
}

func main() {
	// Cli Start
	cli_str := `
    ___  ___  ___  ________  ________     
   |\  \|\  \|\  \|\   ____\|\   __  \    
   \ \  \ \  \\\  \ \  \___|\ \  \|\  \   
 __ \ \  \ \  \\\  \ \  \  __\ \  \\\  \  
|\  \\_\  \ \  \\\  \ \  \|\  \ \  \\\  \ 
\ \________\ \_______\ \_______\ \_______\
 \|________|\|_______|\|_______|\|_______|                              
			   `

	app := &cli.App{
		Name:    "Jugo",
		Version: "0.1.0dev",
		Usage:   "An external information gathering toole.",
		Action: func(c *cli.Context) error {
			fmt.Printf("%s", cli_str)
			return nil
		},
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			// target
			Name:  "t",
			Value: "www.target.com",
			Usage: "A given domain name or ip address",
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:  "h",
		Usage: "Please visit: https://github.com/louse0x/Jugo/help.md",
	}

	app.Run(os.Args)

	// Passive Collection @jugo/passsive.go
	//fmt.Printf("%c[1;0;31m%s%c[0m %-30s\n", 0x1b, "[x]", 0x1b, "Beian information api interface returns an exception")
	jugo.Passive("www.baidu.com")
	// fmt.Println("\n\n")
	// fmt.Println(jugo.RandomHeader())
	fmt.Println(jugo.PassiveData)

}
