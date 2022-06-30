package discordwrapper

import (
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

func JoinAndGetVoiceConnection(token, guildID, channelID string) (*discordgo.VoiceConnection, *discordgo.Session, error) {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, nil, err
	}

	err = discord.Open()
	if err != nil {
		return nil, nil, err
	}

	voiceConnection, err := discord.ChannelVoiceJoin(guildID, channelID, false, false)
	if err != nil {
		return nil, nil, err
	}

	return voiceConnection, discord, nil
}

func Ukenagashi(v1 *discordgo.VoiceConnection, v2 *discordgo.VoiceConnection) {

	recv := make(chan *discordgo.Packet, 2)
	go dgvoice.ReceivePCM(v1, recv)

	send := make(chan []int16, 2)
	go dgvoice.SendPCM(v2, send)

	v2.Speaking(true)
	defer v2.Speaking(false)

	for {
		p, ok := <-recv
		if !ok {
			return
		}
		send <- p.PCM
	}
}
