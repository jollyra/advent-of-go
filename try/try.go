package main

import "fmt"
import "time"

type Bot struct{ HighDest, LowDest, Chips chan int }

func spawnBot() *Bot {
	return &Bot{
		Chips: make(chan int, 1),
	}
}

func (bot *Bot) Start(highDest chan int, lowDest chan int) {
	go func() {
		for {
			x := <-bot.Chips
			y := <-bot.Chips
			fmt.Println("sending", x, y)
			if x >= y {
				highDest <- x
				lowDest <- y
			} else {
				highDest <- y
				lowDest <- x
			}

		}
	}()
}

type Bin struct{ Value int }

func (bin *Bin) Start(ch chan int) {
	go func() {
		for {
			ch <- bin.Value
			fmt.Println("chip retrieved from bin", bin.Value)
		}
	}()
}

func main() {
	fmt.Println("go")

	out0 := make(chan int, 1)
	out1 := make(chan int, 1)

	bin0 := Bin{0}
	bin1 := Bin{1}

	bot0 := spawnBot()
	bot1 := spawnBot()

	bot0.Start(bot1.Chips, bot1.Chips)
	bot1.Start(out0, out1)
	time.Sleep(1 * time.Second)

	bin0.Start(bot0.Chips)
	bin1.Start(bot0.Chips)
	time.Sleep(1 * time.Second)

	x := <-out0
	y := <-out1

	fmt.Println("x, y =", x, y)
}
