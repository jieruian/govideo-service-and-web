package main

import "log"

type ConnLimiter struct {
	courrentConn int
	bucket       chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		courrentConn: cc,
		bucket:       make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	log.Printf("长度%d,总长度%d", len(cl.bucket), cl.courrentConn)
	if len(cl.bucket) >= cl.courrentConn {
		log.Panicln("Reached the rate limitaion")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("新的链接来了%d", c)
}
