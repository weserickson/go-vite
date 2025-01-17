package gvite_plugins

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"

	"github.com/vitelabs/go-vite/cmd/nodemanager"
	"github.com/vitelabs/go-vite/cmd/utils"
)

var (
	pluginDataCommand = cli.Command{
		Action:   utils.MigrateFlags(pluginDataAction),
		Name:     "pluginData",
		Usage:    "pluginData",
		Category: "PLUGIN DATA COMMANDS",
		Flags:    configFlags,
		Description: `
recreate plugin data.
`,
	}
)

func pluginDataAction(ctx *cli.Context) error {
	// Create and start the node based on the CLI flags
	nodeManager, err := nodemanager.NewPluginDataNodeManager(ctx, nodemanager.FullNodeMaker{})
	if err != nil {
		log.Error(fmt.Sprintf("new Node error, %+v", err))
		return err
	}
	if err := nodeManager.Start(); err != nil {
		log.Error(err.Error())
		fmt.Println(err.Error())
		return err
	}

	os.Exit(0)
	return nil
}
