// Copyright © 2017 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package lorawan

// IsConfirmed returns wheter the message is a confirmed message
func (m *Message) IsConfirmed() bool {
	return m.MType == MType_CONFIRMED_UP || m.MType == MType_CONFIRMED_DOWN
}

func (m *Message) initMACPayload() *MACPayload {
	m.Major = Major_LORAWAN_R1
	macPayload := new(MACPayload)
	m.Payload = &Message_MacPayload{MacPayload: macPayload}
	return macPayload
}

// InitUplink initializes an unconfirmed uplink message
func (m *Message) InitUplink() *MACPayload {
	m.MType = MType_UNCONFIRMED_UP
	return m.initMACPayload()
}

// InitDownlink initializes an unconfirmed downlink message
func (m *Message) InitDownlink() *MACPayload {
	m.MType = MType_UNCONFIRMED_DOWN
	return m.initMACPayload()
}
