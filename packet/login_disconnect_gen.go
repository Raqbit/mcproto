// Code generated by "genpacket -packet=LoginDisconnectPacket -output=login_disconnect_gen.go"; DO NOT EDIT.

package packet

import "io"

func (l *LoginDisconnectPacket) Write(w io.Writer) error {
	var err error
	if err = l.Reason.Write(w); err != nil {
		return err
	}
	return nil
}
func (l *LoginDisconnectPacket) Read(r io.Reader) error {
	var err error
	if err = l.Reason.Read(r); err != nil {
		return err
	}
	return nil
}
