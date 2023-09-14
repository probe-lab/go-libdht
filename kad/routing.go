package kad

import "github.com/libp2p/go-libdht"

type RoutingTable[K Key[K], N NodeID[K]] interface {
	libdht.RoutingTable[K, K, N]
}
