package gbpcoffee

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

var messages = []string{"moin", "hi", "morgen", "morgn", "guten morgen", "servus", "servas", "dere", "oida", "porst", "prost", "grias di", "gude", "spinotwachtldroha", "scheipi", "heisl", "gschissana", "christkindl"}

// PluginName is the name of the plugin
var PluginName = "coffee"

// PluginVersion is the version of the plugin, set by the compiler
var PluginVersion = ""

// PluginBuilddate is the builddate of the plugin, set by the compiler
var PluginBuilddate = ""

// Start the plugin
func Start(discord *discordgo.Session) {
	discord.AddHandler(onMessageCreate)
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, v := range messages {
		if v == strings.ToLower(m.Content) {
			var err error

			if m.Author.ID == "263959699764805642" || m.Author.ID == "217697101818232832" {
				err = s.MessageReactionAdd(m.ChannelID, m.ID, "🍵")
				if err != nil {
					logrus.Info(err)
				}
			} else {
				err = s.MessageReactionAdd(m.ChannelID, m.ID, "☕")
				if err != nil {
					logrus.Info(err)
				}
			}

			// faces
			if m.Author.ID == "269898849714307073" {
				err = s.MessageReactionAdd(m.ChannelID, m.ID, ":sidus:576309032789475328")
				if err != nil {
					logrus.Info(err)
				}
			}
			if m.Author.ID == "125230846629249024" {
				err = s.MessageReactionAdd(m.ChannelID, m.ID, ":sikk:355329009824825355")
				if err != nil {
					logrus.Info(err)
				}
			}
		}
	}
}
