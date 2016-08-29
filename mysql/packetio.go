package mysql

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

const (
	defaultReaderSize = 8 * 1024
	defaultWriterSize = 8 * 1024
)

type PacketIO struct {
	rb *bufio.Reader
	wb io.Writer

	Sequence uint8
}

func NewPacketIO(conn net.Conn) *PacketIO {
	p := &PacketIO{
		rb:       bufio.NewReaderSize(conn, defaultReaderSize),
		wb:       conn,
		Sequence: 0,
	}

	return p
}

// from tidb
func (p *PacketIO) readOnePacket() ([]byte, error) {
	var header [4]byte

	if _, err := io.ReadFull(p.rb, header[:]); err != nil {
		return nil, ErrBadConn
	}

	//  length could be zero
	// 	length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
	// 	if length < 1 { /////////////  length could be zero
	// 		return nil, errInvalidPayloadLen.Gen("invalid payload length %d", length)
	// 	}

	sequence := uint8(header[3])
	if sequence != p.Sequence {
		return nil, fmt.Errorf("invalid sequence %d != %d", sequence, p.Sequence)
	}

	p.Sequence++

	length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)

	data := make([]byte, length)
	if _, err := io.ReadFull(p.rb, data); err != nil {
		return nil, ErrBadConn
	}
	return data, nil
}

func (p *PacketIO) ReadPacket() ([]byte, error) {
	data, err := p.readOnePacket()
	if err != nil {
		return nil, err
	}

	if len(data) < MaxPayloadLen {
		return data, nil
	}

	// handle muliti-packet
	for {
		buf, err := p.readOnePacket()
		if err != nil {
			return nil, err
		}

		data = append(data, buf...)

		if len(buf) < MaxPayloadLen {
			break
		}
	}

	return data, nil
}

//data already have header
func (p *PacketIO) WritePacket(data []byte) error {
	length := len(data) - 4

	for length >= MaxPayloadLen {

		data[0] = 0xff
		data[1] = 0xff
		data[2] = 0xff

		data[3] = p.Sequence

		if n, err := p.wb.Write(data[:4+MaxPayloadLen]); err != nil {
			return ErrBadConn
		} else if n != (4 + MaxPayloadLen) {
			return ErrBadConn
		} else {
			p.Sequence++
			length -= MaxPayloadLen
			data = data[MaxPayloadLen:]
		}
	}

	data[0] = byte(length)
	data[1] = byte(length >> 8)
	data[2] = byte(length >> 16)
	data[3] = p.Sequence

	if n, err := p.wb.Write(data); err != nil {
		return ErrBadConn
	} else if n != len(data) {
		return ErrBadConn
	} else {
		p.Sequence++
		return nil
	}
}
