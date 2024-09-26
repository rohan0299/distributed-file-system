package p2p

// Peer Represents A Node
type Peer interface {

}

// Transport Handles Communication B/W Nodes In A Network
// TCP, UDP, Web Sockets
type Transport interface {
	ListenAndAccept() error
}
