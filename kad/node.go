package kad

import "github.com/libp2p/go-libdht"

// NodeID represents the Kademlia Node Identifier, only based on Key both as
// the Point and Distance definition.
type NodeID[K Key[K]] interface {
	libdht.NodeID[K, K]
}

// Request represents the Kademlia Request, only depending on Key, both for
// Point and Distance definitions, and NodeID.
type Request[K Key[K], N NodeID[K]] interface {
	libdht.Request[K, K, N]

	// EmptyResponse returns the format of the associated Response. It is
	// useful to parse the received response to the expected format.
	EmptyResponse() Response[K, N]
}

// Response represents the Kademlia Response, only depending on Keym both for
// Point and Distance definitions, and NodeID.
type Response[K Key[K], N NodeID[K]] interface {
	libdht.Response[K, K, N]
}
