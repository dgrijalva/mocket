package mocket

import (
	"testing"
	"net"
	"bytes"
)

func TestSetup(t *testing.T) {
	m := New()
	if _, ok := m.Server().(net.Conn); !ok {
		t.Errorf("mocket.Server() should be a net.Conn")
	}
	if _, ok := m.Client().(net.Conn); !ok {
		t.Errorf("mocket.Client() should be a net.Conn")
	}
}

func TestBasicRW(t *testing.T) {
	m := New()
	c, s := m.Client(), m.Server()
	
	data := []byte("ping")
	
	c.Write(data)
	r := make([]byte, 4)
	l, e := s.Read(r)
	if i := 4; l != i {
		t.Errorf("length of server read should be %v was %v", i, l)
	}
	if e != nil {
		t.Errorf("server read should not produce error, was %v", e)
	}
	if !bytes.Equal(data, r) {
		t.Errorf("expected %v was %v", data, r)
	}
}