package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gophergala2016/Gobots/easyai"
)

var (
	serverAddress = flag.String("server_address", "", "address of API server")
	token         = flag.String("token", "", "secret token for bot")
	botName       = flag.String("bot_name", "aggro", "which bot to use")
)

const (
	exitFail  = 1
	exitUsage = 64
)

var bots = map[string]easyai.Factory{
	"aggro": func(gameID string) easyai.AI {
		return aggro{}
	},
}

func main() {
	flag.Parse()
	if *serverAddress == "" {
		fmt.Fprintln(os.Stderr, "samplebots: must specify -server_address flag")
		os.Exit(exitUsage)
	}
	if *token == "" {
		fmt.Fprintln(os.Stderr, "samplebots: must specify -token flag")
		os.Exit(exitUsage)
	}
	factory := bots[*botName]
	if factory == nil {
		fmt.Fprintln(os.Stderr, "samplebots: unknown bot name. Known bots:")
		for name := range bots {
			fmt.Fprintf(os.Stderr, "* %s\n", name)
		}
		os.Exit(exitUsage)
	}

	c, err := easyai.Dial(*serverAddress)
	if err != nil {
		fmt.Fprintln(os.Stderr, "samplebots: dial:", err)
		os.Exit(exitFail)
	}
	if err = c.RegisterAI(*token, factory); err != nil {
		fmt.Fprintln(os.Stderr, "samplebots: register:", err)
		os.Exit(exitFail)
	}
	fmt.Fprintln(os.Stderr, "Connected. Ctrl-C or send SIGINT to disconnect.")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT)
	<-sig
	signal.Stop(sig)
	fmt.Fprintln(os.Stderr, "Interrupted. Quitting...")
	if err := c.Close(); err != nil {
		fmt.Fprintln(os.Stderr, "samplebots: close:", err)
		os.Exit(exitFail)
	}
}
