package main

import (
	"github.com/urfave/cli"
	"github.com/nsf/termbox-go"
)

func HolyShit(c *cli.Context) error {
	t := NewTermbox()
	// pp.Println(t)
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	t.Display()

	return nil
}
