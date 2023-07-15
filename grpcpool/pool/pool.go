package pool

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"log"
	"sync"
)

type ClientPool struct {
	pool sync.Pool
}

func GetPool(target string, opts ...grpc.DialOption) (*ClientPool, error) {
	return &ClientPool{
		pool: sync.Pool{
			New: func() any {
				conn, err := grpc.Dial(target, opts...)
				if err != nil {
					log.Fatal(err)
				}
				return conn
			},
		},
	}, nil
}
func (c *ClientPool) Get() *grpc.ClientConn {
	conn := c.pool.Get().(*grpc.ClientConn)
	if conn == nil || conn.GetState() == connectivity.TransientFailure || conn.GetState() == connectivity.Shutdown {
		if conn != nil {
			conn.Close()
		}
		conn = c.pool.Get().(*grpc.ClientConn)
	}
	return conn
}
func (c *ClientPool) Put(conn *grpc.ClientConn) {
	if conn == nil {
		return
	}
	if conn.GetState() == connectivity.TransientFailure || conn.GetState() == connectivity.Shutdown {
		conn.Close()
		return
	}
	c.pool.Put(conn)
}
