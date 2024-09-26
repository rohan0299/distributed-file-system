package p2p

import(
	"net"
	"sync"
	"fmt"
)

// The Remote Node over a TCP Established Connection
type TCPPeer struct {
	// Underlying Connection of Peer
	comm net.Conn

	// Dial and Retrieve Connection -> outbound == true
	// Listen and Accept Connection -> outbound == false
	outbound bool

}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		comm: conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mu sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP Accept Error: %s\n", err)
		}

		go t.handleComm(conn)
	}
}

func (t *TCPTransport) handleComm(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	fmt.Printf("New Incoming Connection: %+v\n", peer)
}