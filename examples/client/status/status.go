package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Raqbit/mcproto"
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/packet"
	"github.com/Raqbit/mcproto/types"
	"log"
	"os"
	"strconv"
)

const (
	ProtocolVersion = 578
)

var (
	ServerHost = flag.String("host", "", "server host")
	ServerPort = flag.String("port", "", "server port")
)

func main() {
	flag.Parse()

	if *ServerHost == "" {
		flag.Usage()
		os.Exit(2)
	}

	ctx := context.Background()

	conn, err := mcproto.Dial(*ServerHost, *ServerPort, types.ClientSide)

	if err != nil {
		log.Fatalf("mcproto dial error: %s", err)
	}

	addr := conn.RemoteAddress()
	port, err := strconv.Atoi(addr.Port)

	if err != nil {
		log.Fatal(err)
	}

	err = conn.WritePacket(ctx,
		&packet.HandshakePacket{
			ProtoVer:   ProtocolVersion,
			ServerAddr: enc.String(addr.Host),
			ServerPort: enc.UnsignedShort(port),
			NextState:  types.ConnectionStateStatus,
		})

	// Actually update our connection as well
	conn.SetState(types.ConnectionStateStatus)

	if err != nil {
		log.Fatalf("could not write handshake packet: %s", err)
	}

	err = conn.WritePacket(ctx, &packet.ServerQueryPacket{})

	if err != nil {
		log.Fatalf("could not write request packet: %s", err)
	}

	resp, err := conn.ReadPacket(ctx)

	if err != nil {
		log.Fatalf("could not read response packet: %s", err)
	}

	response, ok := resp.(*packet.ServerInfoPacket)

	if !ok {
		log.Fatalf("Server sent unexpected packet: %s", resp.String())
	}

	jsonRes, _ := json.Marshal(response.Response)

	fmt.Println(string(jsonRes))
}
