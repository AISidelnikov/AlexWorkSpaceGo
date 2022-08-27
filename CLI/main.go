package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

var name = flag.String("name", "World", "A name to say hello to.")
var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish languige.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish languige.")
}

var opts struct {
	Name string `short:"n" long:"name" default:"World"
	description:"A name to say hello to."`
	Spanish bool `short:"s" long:"spanish"
	description:"Use Spanish Languge"`
}

func main() {
	/*flag.Parse()
	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}*/

	/*flags.Parse(&opts)
	if opts.Spanish == true {
		fmt.Printf("Hola %s\n", opts.Name)
	} else {
		fmt.Printf("Hello %s\n", opts.Name)
	}*/
	app := cli.NewApp()
	app.Usage = "Count up or down"
	app.Commands = []cli.Command{
		{
			Name: "up", ShortName: "u",
			Usage: "Count Up",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "stop, s",
					Usage: "Value to count up to",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("stop")
				if start <= 0 {
					fmt.Println("Stop cannot be negative.")
				}
				for i := 1; i <= start; i++ {
					fmt.Println(i)
				}
				return nil
			},
		},
		{
			Name: "down", ShortName: "d",
			Usage: "Count Down",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start, s",
					Usage: "Start counting down from",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("start")
				if start < 0 {
					fmt.Println("Start cannot be negative.")
				}
				for i := start; i >= 0; i-- {
					fmt.Println(i)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
