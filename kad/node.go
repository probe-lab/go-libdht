package kad

import "github.com/libp2p/go-libdht"

type NodeID[K Key[K]] interface {
	libdht.NodeID[K, K]
}

type Request[K Key[K], N NodeID[K]] interface {
	libdht.Request[K, K, N]
	EmptyResponse() Response[K, N]
}

type Response[K Key[K], N NodeID[K]] interface {
	libdht.Response[K, K, N]
}

var _ NodeID[Key256] = keyID[Key256]{}
var _ Response[Key256, keyID[Key256]] = resp[Key256, keyID[Key256]]{}
var _ Request[Key256, keyID[Key256]] = req[Key256, keyID[Key256]]{}

type keyID[K Key[K]] struct {
	key K
}

func (k keyID[K]) Key() K {
	return k.key
}

type resp[K Key[K], N NodeID[K]] struct {
	peers []N
}

func (r resp[K, N]) CloserNodes() []N {
	return r.peers
}

type req[K Key[K], N NodeID[K]] struct {
	targetId N
}

func (r req[K, N]) Target() K {
	return r.targetId.Key()
}

func (r req[K, N]) EmptyResponse() Response[K, N] {
	return resp[K, N]{}
}
