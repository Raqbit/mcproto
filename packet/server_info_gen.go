// Code generated by "genpacket -packet=ServerInfoPacket -output=server_info_gen.go"; DO NOT EDIT.

package packet

import "io"

func (s *ServerInfoPacket) Write(w io.Writer) error {
	var err error
	if err = s.Response.Write(w); err != nil {
		return err
	}
	return nil
}
func (s *ServerInfoPacket) Read(r io.Reader) error {
	var err error
	if err = s.Response.Read(r); err != nil {
		return err
	}
	return nil
}
