package main

import (
	"flag"
	"log"
	"time"

	"github.com/Bots-Bots-Bots/pkg/commands"
	"github.com/Bots-Bots-Bots/pkg/tcpclient"
)

func main() {
	// command line flags
	publicId := flag.String("publicId", "", "room you want to join to")
	key := flag.String("key", "", "key")
	secret := flag.String("secret", "", "secret")

	flag.Parse()

	c := tcpclient.New("pds.dev.ifi.tv:2020")

	room := commands.NewRoom(*publicId)

	// join room
	if err := c.Write(room.Pass(*key, *secret), 0); err != nil {
		log.Fatal(err)
	}

	cmd, err := c.Read(time.Second)
	if err != nil {
		log.Println("read error:", err)
	} else {
		log.Println("Command:", cmd.Name)
		log.Println("Args:", cmd.Args)
	}

	if err := c.Write(room.Join(), 0); err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("Read")
		cmd, err := c.Read(0)
		if err != nil {
			log.Println("read error:", err)
			continue
		}
		log.Println("from chat:", cmd.Name, cmd.Args)
		if cmd.Name == "SAY" && cmd.Args["username"] != "loyaltyBot" {
			if err := c.Write(room.Say(cmd.Args["message"]), 0); err != nil {
				log.Println("write error:", err)
			}
		}
	}
}
