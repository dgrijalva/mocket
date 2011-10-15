package mocket

import "net"
import "bytes"
import "io"
import "os"
import "sync"

type Mocket struct {
	server, client *buffer
	closed bool
}

func New()*Mocket {
	m := &Mocket{server: newBuffer(&bytes.Buffer{}), client: newBuffer(&bytes.Buffer{})}
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

type buffer struct {
	buf *bytes.Buffer
	cond *sync.Cond
}

func newBuffer(buf *bytes.Buffer)*buffer {
	return &buffer{
		buf: buf,
		cond: sync.NewCond(new(sync.Mutex)),
	}
}

func (b *buffer) Read(d []byte)(int, os.Error) {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()
	for b.buf.Len() < 1 {
	    b.cond.Wait()
	}
	return b.buf.Read(d)
}

func (b *buffer) Write(d []byte)(int, os.Error) {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()
	i, e := b.buf.Write(d)
	b.cond.Signal()
	return i, e
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