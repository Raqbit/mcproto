// Code generated by "genpacket -packet=PongPacket -output=pong_gen.go"; DO NOT EDIT.

package packet

import "io"

func (p *PongPacket) Write(w io.Writer) error {
	var err error
	if err = p.Payload.Write(w); err != nil {
		return err
	}
	return nil
}
func (p *PongPacket) Read(r io.Reader) error {
	var err error
	if err = p.Payload.Read(r); err != nil {
		return err
	}
	return nil
}