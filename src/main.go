package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "Wyst",
		Version: "alpha_0.0.1",
		Authors: []*cli.Author{
			{Name: "Orus"},
		},
		Usage: "The wyst compiler",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Specifies the output binary file name",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return cli.Exit("Error: <file> argument is required.", 1)
			}
			file := c.Args().Get(0)
			if FileExists(file) {
				root := NewRoot(file)
				WriteRoot(root)
				cmd := exec.Command("go", "build", "-o", "../out", "./")
				cmd.Dir = "./build"

				output, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Println("Error:", err)
				}
				fmt.Println(string(output))
				// os.RemoveAll(OUT_FOLDER)
			} else {
				return cli.Exit(fmt.Sprintf("Error: file '%s' not found.", file), 1)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
