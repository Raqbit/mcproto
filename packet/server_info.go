package packet

import (
	"encoding/json"
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
	"io"
)

//go:generate go run ../tools/genpacket/genpacket.go -packet=ServerInfoPacket -output=server_info_gen.go

const ServerInfoPacketID int32 = 0x00

// https://wiki.vg/Protocol#Response
type ServerInfoPacket struct {
	Response ServerInfo
}

func (*ServerInfoPacket) Info() PacketInfo {
	return PacketInfo{
		ID:              ServerInfoPacketID,
		Direction:       types.ClientBound,
		ConnectionState: types.ConnectionStateStatus,
	}
}

func (*ServerInfoPacket) String() string {
	return "ServerInfo"
}

type Version struct {
	Name     string `json:"name"`     // Version name
	Protocol int32  `json:"protocol"` // Version protocol number
}

// Server info player
type Player struct {
	Name string `json:"name"` // Player name
	ID   string `json:"id"`   // Player UUID
}

// Server info players
type Players struct {
	Max    int32    `json:"max"`    // Max amount of players allowed
	Online int32    `json:"online"` // Amount of players online
	Sample []Player `json:"sample"` // Sample of online players
}

// Server ping response
// https://wiki.vg/Server_List_Ping#Response
type ServerInfo struct {
	Version     Version           `json:"version"`           // Server version info
	Players     Players           `json:"players"`           // Server player info
	Description ServerDescription `json:"description"`       // Server description
	Favicon     string            `json:"favicon,omitempty"` // Server favicon
}

func (s *ServerInfo) Write(w io.Writer) error {
	bytes, err := json.Marshal(s)

	if err != nil {
		return err
	}

	str := enc.String(bytes)
	return str.Write(w)
}

func (s *ServerInfo) Read(r io.Reader) error {
	var str enc.String

	if err := str.Read(r); err != nil {
		return err
	}

	return json.Unmarshal([]byte(str), s)
}

// ServerDescription can be both a string (legacy) or TextComponent JSON structure
type ServerDescription types.TextComponent

// Lenient server description parser,
func (c *ServerDescription) UnmarshalJSON(data []byte) error {
	// The data starts with quotes which means it's a string, not an object
	if data[0] == '"' {
		var text string
		if err := json.Unmarshal(data, &text); err != nil {
			return err
		}

		c.Text = text
	} else {
		// Data to unmarshal is not a string, we can parse it as a regular text component
		var out types.TextComponent
		if err := json.Unmarshal(data, &out); err != nil {
			return err
		}
		*c = ServerDescription(out)
	}

	return nil
}