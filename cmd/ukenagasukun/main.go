package main

import (
	"fmt"
	"os"

	"github.com/t98s/ukenagasukun/internal/discordwrapper"
)

func main() {

	var (
		guildID       = os.Getenv("GUILD_ID")
		token1        = os.Getenv("TOKEN1")
		token2        = os.Getenv("TOKEN2")
		fromChannelID = os.Getenv("FROM_CHANNEL_ID")
		toChannelID   = os.Getenv("TO_CHANNEL_ID")
		err           error
	)

	if guildID == "" || token1 == "" || token2 == "" || fromChannelID == "" || toChannelID == "" {
		fmt.Println("Set environment variables correctly")
		return
	}

	voiceConnection1, discord1, err := discordwrapper.JoinAndGetVoiceConnection(token1, guildID, fromChannelID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer voiceConnection1.Close()
	defer discord1.Close()

	voiceConnection2, discord2, err := discordwrapper.JoinAndGetVoiceConnection(token2, guildID, toChannelID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer voiceConnection2.Close()
	defer discord2.Close()

	discordwrapper.Ukenagashi(voiceConnection1, voiceConnection2)
}
