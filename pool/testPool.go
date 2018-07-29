package main

import (
	"Golang/important/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutine = 5
	poolRes      = 2
)

func main() {
	//var pool  sync.Pool
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)

	p, err := pool.New(CreateConnection, poolRes)
	if err != nil {
		log.Println(err)
		return
	}

	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {
			dbQuery(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("开始关闭资源池")
	p.Close()
}

func dbQuery(query int, pool *pool.Pool) {
	conn, err := pool.Acquiry()
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("第%d个查询，使用的是ID为%d的数据库连接", query, conn.(*dbConnection).ID)
}

type dbConnection struct {
	ID int32
}

func (db *dbConnection) Close() error {
	log.Println("关闭连接", db.ID)
	return nil
}

var idCounter int32

func CreateConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id}, nil
}
