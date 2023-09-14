package libdht

type Distance[D any] interface {
	// Compare compares the numeric value of the key with another key of the same type.
	// It returns -1 if the key is numerically less than other, +1 if it is greater
	// and 0 if both keys are equal.
	Compare(other D) int
}

type Point[P any, D Distance[D]] interface {
	// Distance returns a Distance between two points
	Distance(P) D
	Equal(P) bool
}
