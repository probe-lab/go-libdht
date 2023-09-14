package libdht

// RoutingTable is the interface all DHT Routing Tables types support.
type RoutingTable[D Distance[D], P Point[P, D], N NodeID[D, P]] interface {
	// AddNode tries to add a peer to the routing table. It returns true if
	// the node was added and false if it wasn't added, e.g., because it
	// was already part of the routing table.
	//
	// Because NodeID[K]'s are often preimages to Kademlia keys K
	// there's no way to derive a NodeID[K] from just K. Therefore, to be
	// able to return NodeID[K]'s from the `NearestNodes` method, this
	// `AddNode` method signature takes a NodeID[K] instead of only K.
	//
	// Nodes added to the routing table are grouped into buckets based on their
	// XOR distance to the local node's identifier. The details of the XOR
	// arithmetics are defined on K.
	AddNode(N) bool

	// RemoveKey tries to remove a node identified by its Kademlia key from the
	// routing table.
	//
	// It returns true if the key existed in the routing table and was removed.
	// It returns false if the key didn't exist in the routing table and
	// therefore, was not removed.
	RemoveNode(N) bool

	// NearestNodes returns the given number of closest nodes to a given
	// Kademlia key that are currently present in the routing table.
	// The returned list of nodes will be ordered from closest to furthest and
	// contain at maximum the given number of entries, but also possibly less
	// if the number exceeds the number of nodes in the routing table.
	NearestNodes(P, int) []N

	// GetNode returns the node identified by the supplied Kademlia key or a zero
	// value if the node is not present in the routing table. The boolean second
	// return value indicates whether the node was found in the table.
	GetNode(P) (N, bool)
}
