package kadtest

import (
	"crypto/sha256"

	"github.com/libp2p/go-libdht/kad"
	"github.com/plprobelab/go-kademlia/key"
)

// ID is a concrete implementation of the NodeID interface.
type ID[K kad.Key[K]] struct {
	key K
}

// interface assertion. Using the concrete key type of key.Key8 does not
// limit the validity of the assertion for other key types.
var _ kad.NodeID[key.Key8] = (*ID[key.Key8])(nil)

// NewID returns a new Kademlia identifier that implements the NodeID interface.
// Instead of deriving the Kademlia key from a NodeID, this method directly takes
// the Kademlia key.
func NewID[K kad.Key[K]](k K) *ID[K] {
	return &ID[K]{key: k}
}

// Key returns the Kademlia key that is used by, e.g., the routing table
// implementation to group nodes into buckets. The returned key was manually
// defined in the ID constructor NewID and not derived via, e.g., hashing
// a preimage.
func (i ID[K]) Key() K {
	return i.key
}

func (i ID[K]) Equal(other K) bool {
	return i.key.Compare(other) == 0
}

func (i ID[K]) String() string {
	return key.HexString(i.key)
}

type StringID string

var _ kad.NodeID[kad.Key256] = (*StringID)(nil)

func NewStringID(s string) *StringID {
	return (*StringID)(&s)
}

func (s StringID) Key() kad.Key256 {
	h := sha256.New()
	h.Write([]byte(s))
	return kad.NewKey256(h.Sum(nil))
}

func (s StringID) NodeID() kad.NodeID[kad.Key256] {
	return &s
}

func (s StringID) Equal(other string) bool {
	return string(s) == other
}

func (s StringID) String() string {
	return string(s)
}

var _ kad.NodeID[kad.Key256] = keyID[kad.Key256]{}
var _ kad.Response[kad.Key256, keyID[kad.Key256]] = resp[kad.Key256, keyID[kad.Key256]]{}
var _ kad.Request[kad.Key256, keyID[kad.Key256]] = req[kad.Key256, keyID[kad.Key256]]{}

type keyID[K kad.Key[K]] struct {
	key K
}

func (k keyID[K]) Key() K {
	return k.key
}

type resp[K kad.Key[K], N kad.NodeID[K]] struct {
	peers []N
}

func (r resp[K, N]) CloserNodes() []N {
	return r.peers
}

type req[K kad.Key[K], N kad.NodeID[K]] struct {
	targetId N
}

func (r req[K, N]) Target() K {
	return r.targetId.Key()
}

func (r req[K, N]) EmptyResponse() kad.Response[K, N] {
	return resp[K, N]{}
}
