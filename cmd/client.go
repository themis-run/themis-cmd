package cmd

import (
	"sync"

	"go.themis.run/themis-cmd/conf"
	"go.themis.run/themisclient"
)

var client *themisclient.Client
var one sync.Once

func GetThemisClient() *themisclient.Client {
	var err error
	one.Do(func() {
		serverName, address := conf.GetThemisServer()
		opt := []themisclient.Option{
			themisclient.WithServerName(serverName),
			themisclient.WithServerAddress(address),
		}

		config := themisclient.NewConfigration(opt...)

		client, err = themisclient.NewClient(config)
	})
	if err != nil {
		panic(err)
	}

	return client
}
