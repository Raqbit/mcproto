// Code generated by "genpacket -packet=HandshakePacket -output=handshake_gen.go"; DO NOT EDIT.

package packet

import "io"

func (h *HandshakePacket) Write(w io.Writer) error {
	var err error
	if err = h.ProtoVer.Write(w); err != nil {
		return err
	}
	if err = h.ServerAddr.Write(w); err != nil {
		return err
	}
	if err = h.ServerPort.Write(w); err != nil {
		return err
	}
	if err = h.NextState.Write(w); err != nil {
		return err
	}
	return nil
}
func (h *HandshakePacket) Read(r io.Reader) error {
	var err error
	if err = h.ProtoVer.Read(r); err != nil {
		return err
	}
	if err = h.ServerAddr.Read(r); err != nil {
		return err
	}
	if err = h.ServerPort.Read(r); err != nil {
		return err
	}
	if err = h.NextState.Read(r); err != nil {
		return err
	}
	return nil
}
