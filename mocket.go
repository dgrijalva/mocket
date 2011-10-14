package mocket

import "net"
import "bytes"
import "io"
import "os"

type Mocket struct {
	server, client *bytes.Buffer
	closed bool
}

func New()*Mocket {
	m := &Mocket{server: &bytes.Buffer{}, client: &bytes.Buffer{}}
	return m
}

func (m *Mocket) Server()net.Conn {
	return &side{m, m.server, m.client}
}

func (m *Mocket) Client()net.Conn {
	return &side{m, m.client, m.server}
}

func (m *Mocket) Close() {
	
}

type side struct {
	m *Mocket
	r io.Reader
	w io.Writer
}

func (s *side) Read(p []byte) (n int, err os.Error) {
	return s.r.Read(p)
}

func (s *side) Write(p []byte) (n int, err os.Error) {
	return s.w.Write(p)
}

func (s *side) LocalAddr() net.Addr {
	return nil
}

func (s *side) RemoteAddr() net.Addr {
	return nil
}

func (s *side) SetTimeout(nsec int64) os.Error {
	return nil
}

func (s *side) SetReadTimeout(nsec int64) os.Error {
	return nil
}

func (s *side) SetWriteTimeout(nsec int64) os.Error {
	return nil
}

func (s *side) Close() os.Error {
	s.m.Close()
	return nil
}