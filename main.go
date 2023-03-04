package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jaredwarren/macroPi/config"
	"github.com/jaredwarren/macroPi/key"
	"github.com/jaredwarren/macroPi/log"
	"github.com/jaredwarren/macroPi/server"
	"github.com/jaredwarren/macroPi/usb"
)

func main() {
	logger := log.NewStdLogger(log.Debug)

	err := config.InitConfig()
	if err != nil {
		logger.Fatal("Init", log.Error(err))
	}
	cfg := config.Get()

	// Init Server
	htmlServer := server.HTMLServer{
		Logger: logger,
		Config: cfg.Host,
		Macro:  cfg.Macro,
	}
	htmlServer.Start()
	defer htmlServer.StopHTTPServer()

	// Wait for Shutdown
	WaitForInterupt()

	logger.Info("main :shutting down")

	return

	w, err := usb.NewHID()
	if err != nil {
		fmt.Printf("[Warning] - %s\n", err)
		// ignore and continue with "test" writer
	}
	defer w.Close()

	// TODO: setup server

	fmt.Println("starting...")
	time.Sleep(3 * time.Second)
	fmt.Println("go!")

	i := key.StringToInput("Jared")
	key.RunInput(w, i)
	return

	key.PressAndReleaseKey(w, key.MuteKey, key.NoMod)
	time.Sleep(2 * time.Second)

	key.PressAndReleaseKey(w, key.VoluemUpKey, key.NoMod)
	time.Sleep(2 * time.Second)

	key.PressAndReleaseKey(w, key.AKey, key.RShift)
	time.Sleep(2 * time.Second)

	key.PressAndReleaseKey(w, key.AKey, key.NoMod)
	time.Sleep(2 * time.Second)

	key.PressAndReleaseKey(w, key.AKey, key.LShift|key.RCtrl)
}

func WaitForInterupt() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}
