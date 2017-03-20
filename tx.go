package ipldeth

import (
	"bytes"
	"fmt"

	cid "gx/ipfs/QmV5gPoRsjN1Gid3LMdNZTyfCtP2DsvqEbMAmz82RmmiGk/go-cid"
	node "gx/ipfs/QmYDscK7dmdo2GZ9aumS8s5auUUAH5mR1jvj5pYhWusfK7/go-ipld-node"
	mh "gx/ipfs/QmbZ6Cee2uHjG7hf19qLHppgKDRtaG4CVtMzdmK9VCVqLu/go-multihash"

	types "github.com/ethereum/go-ethereum/core/types"
)

type Tx struct {
	tx *types.Transaction
}

func (b *Tx) Cid() *cid.Cid {
	c, err := cid.Prefix{
		Codec:    cid.EthereumTx,
		Version:  1,
		MhType:   mh.KECCAK_256,
		MhLength: -1,
	}.Sum(b.RawData())
	if err != nil {
		panic(err)
	}
	return c
}

func (t *Tx) MarshalJSON() ([]byte, error) {
	return t.tx.MarshalJSON()
}

func (b *Tx) Copy() node.Node {
	panic("dont use this yet")
}

func (b *Tx) Links() []*node.Link {
	return nil
}

func (b *Tx) Loggable() map[string]interface{} {
	return map[string]interface{}{
		"type": "ethereum_block",
	}
}

func (b *Tx) RawData() []byte {
	buf := new(bytes.Buffer)
	if err := b.tx.EncodeRLP(buf); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func (b *Tx) Resolve(p []string) (interface{}, []string, error) {
	if len(p) == 0 {
		return b, nil, nil
	}

	switch p[0] {
	case "nonce":
		return b.tx.Nonce(), p[1:], nil
	case "gasPrice":
		return b.tx.GasPrice(), p[1:], nil
	case "gas":
		return b.tx.Gas(), p[1:], nil
	case "toAddress":
		return b.tx.To(), p[1:], nil
	case "value":
		return b.tx.Value(), p[1:], nil
	case "data":
		return b.tx.Data(), p[1:], nil
	default:
		return nil, nil, fmt.Errorf("no such link")
	}
}

func (b *Tx) ResolveLink(p []string) (*node.Link, []string, error) {
	obj, rest, err := b.Resolve(p)
	if err != nil {
		return nil, nil, err
	}
	lnk, ok := obj.(*node.Link)
	if !ok {
		return nil, nil, fmt.Errorf("was not a link")
	}

	return lnk, rest, nil
}

func (b *Tx) Size() (uint64, error) {
	panic("don't do size")
}

func (b *Tx) Stat() (*node.NodeStat, error) {
	return &node.NodeStat{}, nil
}

func (b *Tx) String() string {
	return fmt.Sprintf("<EthereumTx %s>", b.Cid())
}

func (b *Tx) Tree(p string, depth int) []string {
	return nil
}

func (b *Tx) BaseTx() *types.Transaction {
	return b.tx
}
