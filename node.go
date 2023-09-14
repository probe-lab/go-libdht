package libdht

// NodeID is a generic node identifier and not equal to a Point. Points are used
// for calculating distances between nodes while NodeID's are the original
// logical identifier of a node.
//
// The NodeID interface only defines a method that returns the Kademlia key
// for the given NodeID. E.g., the operation to go from a NodeID to a Point
// can be as simple as hashing the NodeID.
//
// Implementations may choose to equate NodeID's and Points.
type NodeID[D Distance[D], P Point[P, D]] interface {
	// Key returns the Point associated with the given NodeID.
	Key() P
}

// Request represent the request message format to perform a DHT lookup. All
// request message must contain the target location within the underlying
// geography.
type Request[D Distance[D], P Point[P, D], N NodeID[D, P]] interface {
	// Target specifies the Point of interest in the request.
	Target() P
}

// Response represents a reply to a DHT lookup request from a node in the
// network, typically providing a list of nodes that are closer to a target
// Point or satisfying some query criteria.
type Response[D Distance[D], P Point[P, D], N NodeID[D, P]] interface {
	// CloserNodes fetches a list of NodeIDs that are proximate to the target
	// Point or fulfill the original request's criteria.
	CloserNodes() []N
}
