package car_test

import (
	"context"

	blocks "github.com/ipfs/go-block-format"
	cid "github.com/ipfs/go-cid"
	car "github.com/ipld/go-car"
)

// tests deleted because they depended on merkledag

type countingReadStore struct {
	bs    car.ReadStore
	count int
}

func (rs *countingReadStore) Get(ctx context.Context, c cid.Cid) (blocks.Block, error) {
	rs.count++
	return rs.bs.Get(ctx, c)
}
