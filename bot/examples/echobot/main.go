package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Bots-Bots-Bots/pkg/bot"
	"github.com/Bots-Bots-Bots/pkg/commands"
)

func main() {
	// command line flags
	publicId := flag.String("publicId", "", "room you want to join to")
	key := flag.String("key", "", "key")
	secret := flag.String("secret", "", "secret")
	host := flag.String("secret", "pds.dev.ifi.tv:2020", "Bot secret")

	flag.Parse()

	// Create a bot
	b, err := bot.New(*host, *key, *secret, *publicId)
	if err != nil {
		log.Fatal(err)
	}

	// handle LEAVE on server shutdown
	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			cmd, err := b.Read()
			if err != nil {
				log.Println("read error:", err)
				continue
			}
			switch cmd.Name {
			case commands.LSay:
				if cmd.Get("bot") == "false" {
					if err := b.Say(cmd.Get("message")); err != nil {
						log.Println("write error:", err)
					}
				}
			}
		}
	}()

	<-done
	fmt.Println("Leaving chat room")
	b.Leave()
}