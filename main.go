package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"steampipe-plugin-eol/eol"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: eol.Plugin})
}
