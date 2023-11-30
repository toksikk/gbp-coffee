package main

import (
	"github.com/bwmarrin/discordgo"
	plugin "github.com/toksikk/gbp-coffee/plugin"
)

// PluginName is the name of the plugin
var PluginName = plugin.PluginName

// PluginVersion is the version of the plugin, set by the compiler
var PluginVersion = ""

// PluginBuilddate is the builddate of the plugin, set by the compiler
var PluginBuilddate = ""

// Start the plugin
func Start(discord *discordgo.Session) {
	plugin.Start(discord)
}
