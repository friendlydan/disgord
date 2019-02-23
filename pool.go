package disgord

import "sync"

type Pool interface {
	Put(x Reseter)
	Get() (x Reseter)
}

type pool struct {
	pool sync.Pool

	// New specifies a function to generate a
	// value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
	New func() Reseter
}

// Put resets the object before it is put back into the pool. We reset it here
// to quickly detect if there are other owners than the pool as it is inserted.
func (p *pool) Put(x Reseter) {
	x.Reset()
	p.pool.Put(x)
}

// Get selects an arbitrary item from the Pool, removes it from the
// Pool, and returns it to the caller.
// Get may choose to ignore the pool and treat it as empty.
// Callers should not assume any relation between values passed to Put and
// the values returned by Get.
//
// This assumes that p.New is always set.
func (p *pool) Get() (x Reseter) {
	var ok bool
	if x, ok = p.pool.Get().(Reseter); x == nil || !ok {
		x = p.New()
	}

	return
}

//////////////////////////////////////////////////////
//
// Client Pools
//
//////////////////////////////////////////////////////

func (c *client) ChannelPool() Pool {
	return c.channelPool
}
func (c *client) MessagePool() Pool {
	return c.userPool
}
func (c *client) UserPool() Pool {
	return c.userPool
}
