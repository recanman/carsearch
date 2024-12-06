// made by recanman
package notify

import "github.com/bwmarrin/discordgo"

type DiscordNotifier struct {
	Session   *discordgo.Session
	ChannelID string
}

func NewDiscordNotifier(token string, channelId string) (*DiscordNotifier, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return &DiscordNotifier{Session: session, ChannelID: channelId}, nil
}

func (d *DiscordNotifier) Notify(embed *discordgo.MessageEmbed) error {
	_, err := d.Session.ChannelMessageSendEmbed(d.ChannelID, embed)
	return err
}

func (d *DiscordNotifier) NotifyError(err error) error {
	_, err = d.Session.ChannelMessageSend(d.ChannelID, err.Error())
	return err
}
