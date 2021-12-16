package main

import "fmt"

type vehicle interface {
	license() string
	name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) license() string {
	return c.licenceNo
}

func (c car) name() string {
	return c.brand
}

func main() {
	var v vehicle = car{licenceNo: "378dh4", brand: "Honda"}
	fmt.Println(v.license())
	fmt.Println(v.name())
}
