package libdht

// NodeID is a generic node identifier and not equal to a Kademlia key. Some
// implementations use NodeID's as preimages for Kademlia keys. Kademlia keys
// are used for calculating distances between nodes while NodeID's are the
// original logical identifier of a node.
//
// The NodeID interface only defines a method that returns the Kademlia key
// for the given NodeID. E.g., the operation to go from a NodeID to a Kademlia key
// can be as simple as hashing the NodeID.
//
// Implementations may choose to equate NodeID's and Kademlia keys.
type NodeID[D Distance[D], P Point[P, D]] interface {
	// Key returns the Kademlia key of the given NodeID. E.g., NodeID's can be
	// preimages to Kademlia keys, in which case, Key() could return the SHA256
	// of NodeID.
	Key() P
}

type Request[D Distance[D], P Point[P, D], N NodeID[D, P]] interface {
	Target() P
}

type Response[D Distance[D], P Point[P, D], N NodeID[D, P]] interface {
	CloserNodes() []N
}
