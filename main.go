package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	sess, err := discordgo.New("Bot MTEzNzIwMjUwNzU1OTQ4OTU1Ng.GPwd-q.ZM_1CBJbDM93ZJgZDquwpJHfSbu_neQhgzNh6c")
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "boykisser" {
			s.ChannelMessageSend(m.ChannelID, "Ambatukammmm")
		}

		sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

		err = sess.Open()
		if err != nil {
			log.Fatal(err)
		}

		defer sess.Close()

		fmt.Println("Mantaf")

		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	})
}
