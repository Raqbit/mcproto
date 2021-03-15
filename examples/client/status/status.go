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
	"net"
	"os"
	"strconv"
)

const (
	ProtocolVersion = 578
)

var (
	ServerHost = flag.String("host", "", "server host")
	ServerPort = flag.Uint("port", 25565, "server port")
)

func main() {
	flag.Parse()

	if *ServerHost == "" {
		flag.Usage()
		os.Exit(2)
	}

	ctx := context.Background()

	address := net.JoinHostPort(*ServerHost, strconv.Itoa(int(*ServerPort)))

	conn, err := mcproto.Dial(address, types.ClientSide)

	if err != nil {
		log.Fatalf("tcp dial error: %s", err)
	}

	err = conn.WritePacket(ctx,
		&packet.HandshakePacket{
			ProtoVer:   ProtocolVersion,
			ServerAddr: enc.String(*ServerHost),
			ServerPort: enc.UnsignedShort(*ServerPort),
			NextState:  types.ConnectionStateStatus,
		})

	// Actually update our connection as well
	conn.SwitchState(types.ConnectionStateStatus)

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
