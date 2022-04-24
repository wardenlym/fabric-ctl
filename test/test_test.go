package test

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"

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
