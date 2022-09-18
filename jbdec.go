package jbdec

import (
	"bytes"
	"fmt"
)

type Decoder struct {
	b   []byte
	pos int
}

type Token struct {
	Type  Type
	Pos   int
	value []byte
	err   error
}

type Type byte

const (
	BeginArray     = Type('[')
	BeginObject    = Type('{')
	EndArray       = Type(']')
	EndObject      = Type('}')
	NameSeparator  = Type(':')
	ValueSeparator = Type(',')

	Null  = Type('n')
	True  = Type('t')
	False = Type('f')

	String = Type('S')
	Number = Type('N')

	EOF   = Type(0x00)
	Error = Type(0xFF)

	escape = '\\'
)

func New(b []byte) *Decoder {
	return &Decoder{
		b:   b,
		pos: 0,
	}
}

var (
	byteNull  = []byte("null")
	byteTrue  = []byte("true")
	byteFalse = []byte("false")
)

func (d *Decoder) Next() Token {
	d.skipWS()

	if d.pos >= len(d.b) {
		return token(EOF, nil, 0, nil)
	}

	first := d.b[d.pos]

	//log.Printf("%d %x %s", d.pos, first, string(d.b[d.pos:]))

	switch first {
	case byte(BeginArray):
		d.pos++
		return token(BeginArray, nil, d.pos-1, nil)

	case byte(EndArray):
		d.pos++
		return token(EndArray, nil, d.pos-1, nil)

	case byte(BeginObject):
		d.pos++
		return token(BeginObject, nil, d.pos-1, nil)

	case byte(EndObject):
		d.pos++
		return token(EndObject, nil, d.pos-1, nil)

	case byte(NameSeparator):
		d.pos++
		return token(NameSeparator, nil, d.pos-1, nil)

	case byte(ValueSeparator):
		d.pos++
		return token(ValueSeparator, nil, d.pos-1, nil)

	case 'n':
		ok := d.progressIf(byteNull...)
		if ok {
			return token(Null, byteNull, d.pos-4, nil)
		}

	case 't':
		ok := d.progressIf(byteTrue...)
		if ok {
			return token(True, byteTrue, d.pos-4, nil)
		}

	case 'f':
		ok := d.progressIf(byteFalse...)
		if ok {
			return token(False, byteFalse, d.pos-5, nil)
		}

	default:
		//nop
	}

	if isNumComponent(first) {
		//log.Print("number")
		p := d.pos + 1
		l := len(d.b)
		for {
			if p >= l {
				break
			}
			if !isNumComponent(d.b[p]) {
				break
			}
			p++
		}
		s := d.pos
		d.pos = p
		return token(Number, d.b[s:p], s, nil)
	}

	if first != '"' {
		return token(Error, nil, d.pos, fmt.Errorf("invalid token %v", first))
	}

	escaping := false
	p := d.pos + 1
	l := len(d.b)
	b := d.b
	for {
		if p >= l {
			break
		}

		if b[p] == '\\' {
			escaping = true
			p++
			continue
		}

		if !escaping && b[p] == '"' {
			p++
			break
		}

		escaping = false

		p++
	}

	s := d.pos
	d.pos = p

	//log.Print("string", s, p, string(d.b[s:p]))

	return token(String, d.b[s:p], s, nil)
}

func (d *Decoder) skipWS() {
	b := d.b
	p := d.pos
	l := len(d.b)

	for {
		if p >= l {
			return
		}

		if !isWS(b[p]) {
			d.pos = p
			return
		}

		p++
	}
}

func (d *Decoder) progressIf(test ...byte) bool {
	b := d.b
	p := d.pos
	l := len(d.b)
	tl := len(test)

	if l < p+tl {
		return false
	}

	if !bytes.Equal(b[p:p+tl], test) {
		return false
	}
	d.pos += tl

	return true
}

func token(typ Type, value []byte, pos int, err error) Token {
	return Token{
		Type:  typ,
		Pos:   pos,
		value: value,
		err:   err,
	}
}

func isWS(c byte) bool {
	return c == 0x20 || c == 0x09 || c == 0x0A || c == 0x0D
}

func isNumComponent(c byte) bool {
	return ('0' <= c && c <= '9') || c == '-' || c == '+' || c == '.' || c == 'e' || c == 'E'
}

func (t Token) String() string {
	return string(t.value)
}

func (t Token) Bytes() []byte {
	return t.value
}

func (t Token) Error() error {
	return t.err
}
