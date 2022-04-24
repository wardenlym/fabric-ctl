package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"

	"github.com/kr/pretty"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"

	_ "github.com/hyperledger/fabric-sdk-go/pkg/gateway"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	_ "github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	_ "github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	_ "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	_ "github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
)

func Test_depends_sdk(t *testing.T) {
	sdk, err := fabsdk.New(nil)

	if err != nil {
		t.Log(errors.Wrap(err, "fab sdk error"))
		return
	}
	defer sdk.Close()
}

func Test_depends_client_channel(t *testing.T) {
	channelclient, err := channel.New(nil)

	if err != nil {
		t.Log(errors.Wrap(err, "fab sdk error"))
		return
	}

	fmt.Printf("channelclient: %v\n", channelclient)
}

func Test_docker_client(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}

func Test_cobra(t *testing.T) {
	rootCmd := &cobra.Command{}

	rootCmd = rootCmd
	pretty.Printf("rootCmd: %# v\n", pretty.Formatter(rootCmd))
}
