// Code generated by "genpacket -packet=ClientSettingsPacket -output=client_settings_gen.go"; DO NOT EDIT.

package packet

import "io"

func (c *ClientSettingsPacket) Write(w io.Writer) error {
	var err error
	if err = c.Lang.Write(w); err != nil {
		return err
	}
	if err = c.ViewDistance.Write(w); err != nil {
		return err
	}
	if err = c.ChatVisibility.Write(w); err != nil {
		return err
	}
	if err = c.EnableChatColors.Write(w); err != nil {
		return err
	}
	if err = c.DisplayedSkinParts.Write(w); err != nil {
		return err
	}
	if err = c.MainHand.Write(w); err != nil {
		return err
	}
	return nil
}
func (c *ClientSettingsPacket) Read(r io.Reader) error {
	var err error
	if err = c.Lang.Read(r); err != nil {
		return err
	}
	if err = c.ViewDistance.Read(r); err != nil {
		return err
	}
	if err = c.ChatVisibility.Read(r); err != nil {
		return err
	}
	if err = c.EnableChatColors.Read(r); err != nil {
		return err
	}
	if err = c.DisplayedSkinParts.Read(r); err != nil {
		return err
	}
	if err = c.MainHand.Read(r); err != nil {
		return err
	}
	return nil
}
