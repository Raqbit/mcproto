// Code generated by "genpacket -packet=LoginStartPacket -output=login_start_gen.go"; DO NOT EDIT.

package packet

import "io"

func (l *LoginStartPacket) Write(w io.Writer) error {
	var err error
	if err = l.Name.Write(w); err != nil {
		return err
	}
	return nil
}
func (l *LoginStartPacket) Read(r io.Reader) error {
	var err error
	if err = l.Name.Read(r); err != nil {
		return err
	}
	return nil
}
