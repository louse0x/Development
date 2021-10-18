package main

/*


 */

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
	var dirs = [3]string{"raw", "log", "result"}
	for _, value := range dirs {
		if _, err := os.Stat(filepath.Join(wd, value)); err != nil {
			// Dir IsNotExist
			if os.IsNotExist(err) {
				os.MkdirAll(value, os.ModePerm)
			}
		}
	}

}

func main() {
	// Cli Start
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	// Passive Collection @jugo/passsive.go

	jugo.Passive()

}
