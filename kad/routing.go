package kad

import "github.com/libp2p/go-libdht"

// RoutingTable is the Kademlia Routing Table depending on Kademlia Key both
// for Point and Distance definition.
type RoutingTable[K Key[K], N NodeID[K]] interface {
	libdht.RoutingTable[K, K, N]
}
