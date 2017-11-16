package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"time"
)

type Termbox struct {
	Width  int
	Height int
	Shit   []rune
}

func NewTermbox() *Termbox {
	return &Termbox{
		Shit: NewShit(),
	}
}

func (t *Termbox) SetSize() {
	t.Width, t.Height = termbox.Size()
}

func (t *Termbox) Display() {
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputAlt)
	t.SetSize()
	t.Draw()

	tick := time.NewTicker(1 * time.Second)
	stop := make(chan bool)

	go func(t *Termbox) {
		for {
			select {
			case <- tick.C:
				t.Shit = append(t.Shit, '💩')
				t.Draw()
			case <- stop:
				break
			}
		}
	}(t)

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC:
				tick.Stop()
				close(stop)
				break loop
			case termbox.KeyCtrlD:
				tick.Stop()
				close(stop)
				break loop
			}
		}

		t.Draw()
	}
}

func (t *Termbox) Draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	t.Print(0, 0, termbox.ColorDefault, termbox.ColorDefault, t.Shit)
	termbox.Flush()
}

func (t *Termbox) Print(x, y int, fg, bg termbox.Attribute, msg []rune) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		width := runewidth.RuneWidth(c)
		if x+width > t.Width {
			x = 0
			y++
		} else {
			x = x + width
		}
	}
}
