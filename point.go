package libdht

// Distance represents a scalar value that measures the separation between two
// Points. While typically numeric, it can be tailored to represent other types
// of distance metrics.
type Distance[D any] interface {
	// Compare compares the numeric value of the key with another key of the
	// same type. It returns -1 if the key is numerically less than other, +
	// 1 if it is greater and 0 if both keys are equal.
	Compare(other D) int
}

// Point symbolizes a specific location within the DHT's geography.
// Typically, Peer and Content identifiers map to a Point, delineating the
// connections between peers and determining how content is allocated among
// them.
type Point[P any, D Distance[D]] interface {
	// Distance calculates the spatial separation between the current Point and
	// the provided Point P.
	Distance(P) D
	// Equal checks whether the current Point and the provided Point P occupy
	// the same spatial position.
	Equal(P) bool
}
