package main

import (
	"strconv"
	"time"

	termbox "github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetOutputMode(termbox.Output256)

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	termbox.Clear(0, 0)

	render()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
					return
				default:
					render()
					time.Sleep(time.Millisecond * 100)
				}
			}
		}
	}
}

func render() {
	termbox.Clear(0, 0)
	termbox.SetCell(10, 10, ' ', 10, 10)

	width := 10
	var colorNumber termbox.Attribute = 0

	for y := 0; y < 30 && colorNumber <= 256; y++ {
		for col := 0; col < 10 && colorNumber <= 256; col++ {

			zstr := strconv.Itoa(int(colorNumber))
			l := len(zstr)
			if l < (width) {
				for i := l; i < width; i++ {
					zstr += " "
				}
			}
			runes := []rune(zstr)
			for x2, letter := range runes {
				termbox.SetCell((col*width)+x2, y, letter, 0, colorNumber)
			}
			colorNumber++
		}
	}
	termbox.Flush()
}
