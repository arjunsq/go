package main

import (
	"fmt"
)

type driver interface {
	connect()
	close()
}

type mysql struct {
	path string
}

type postqr struct {
	path string
	key  int
}

func main() {
	m := &mysql{}
	dbconnect(m)
	p := &postqr{}
	dbconnect(p)
}

func (m *mysql) connect() {
	m.path = "localhost"
	fmt.Println("Mysql is connecting to ", m.path)
}

func (m *mysql) close() {
	fmt.Println("Mysql closed")
}

func (p *postqr) connect() {
	p.path = "localhost"
	p.key = 5
	fmt.Println("postqr is connecting to ", p.path, " key: ", p.key)
}

func (p *postqr) close() {
	fmt.Println("postqr closed")
}

func dbconnect(d driver) {
	//Check system status
	d.connect()
	// do something
	//everything ok
	d.close()
}
