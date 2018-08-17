package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m       sync.Mutex
	res     chan io.Closer
	factory func() (io.Closer, error)
	closed  bool
}

var ErrPoolClosed = errors.New("the pool is closed!")

func New(fn func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("It's a invalid size")
	}
	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquiry() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		log.Println("Acquiry : 共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		return p.factory()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	close(p.res) //关闭通道，不让写入了..但是注意,还可以读

	for r := range p.res {
		r.Close()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}
	select {
	case p.res <- r:
		log.Println("资源释放到资源池了")
	default:
		log.Println("资源池满了,请释放")
		r.Close()
	}
}
