package main

import (
	"github.com/ctcpip/notifize"
	"github.com/go-co-op/gocron/v2"
	"time"
	"fmt"
	"os"
)

func main() {
	s, _ := gocron.NewScheduler()

	j, _ := s.NewJob(
		gocron.DurationJob(1*time.Second), // temporary
		gocron.NewTask(func() { notify() }),
	)

	fmt.Println(j.ID())

	s.Start()

	select {}
}

func notify() {
	notifize.Display("Drink Water", "It's time to drink water 🗣️🗣️🗣️", true, "/home/radi8/Code/Golang/Drink-Agua/icon.png")
	os.Exit(0)
}
