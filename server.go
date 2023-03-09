package stupidstuff

import (
	"log"
	"net"
)

type ServerOpts struct {
	ListenAddr string
}

type Server struct {
	ServerOpts
	cache Cacher
}

func NewServer(o ServerOpts, c Cacher) Server {
	return Server{
		ServerOpts: o,
		cache:      c,
	}
}

func (s Server) StartAndListen() error {
	b, err := net.Listen("tcp", s.ServerOpts.ListenAddr)
	if err != nil {
		return err
	}

	log.Printf("server listen on %s", s.ServerOpts.ListenAddr)

	for {
		conn, err := b.Accept()
		if err != nil {
			continue
		}
		go s.handleConn(conn)
	}
}

func (s Server) handleConn(c net.Conn) {
	defer c.Close()

	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf)
		if err != nil {
			log.Printf("[server][handleConn] error read %s", err)
			break
		}

		m := buf[:n]
		go s.handleCmd(c, m)
	}
}

func (s Server) handleCmd(c net.Conn, msg []byte) {
	m, err := parseMessage(msg)
	if err != nil {
		c.Write([]byte(err.Error() + "\n"))
		return
	}

	switch m.Cmd {
	case CMDSet:
		err := s.cache.Set(m.Key, m.Value)
		if err != nil {
			c.Write([]byte(err.Error() + "\n"))
			return
		}
		c.Write([]byte("ok\n"))
	case CMDGet:
		v, err := s.cache.Get(m.Key)
		if err != nil {
			c.Write([]byte(err.Error() + "\n"))
			return
		}
		wv := string(v) + "\n"
		c.Write([]byte(wv))
	}
}
