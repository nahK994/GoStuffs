package ocp

import "fmt"

type Mobile struct {
	brand      string
	memorySize int
}

type Matcher interface {
	isMatched(m Mobile) bool
}

type Brand struct {
	name string
}

func (b Brand) isMatched(m Mobile) bool {
	return m.brand == b.name
}

type Memory struct {
	size int
}

func (memory Memory) isMatched(m Mobile) bool {
	return m.memorySize == memory.size
}

type Filter struct{}

func (f Filter) Filter(mobiles []Mobile, matcher Matcher) {
	for _, j := range mobiles {
		fmt.Println(j.brand, "-> IsFiltered:", matcher.isMatched(j))
	}
}

func Example1() {
	filterBrand := Brand{name: "Samsung"}
	filterSize := Memory{size: 1024}

	m1 := Mobile{
		brand:      "Samsung",
		memorySize: 1024,
	}
	m2 := Mobile{
		brand:      "Nokia",
		memorySize: 2048,
	}

	var mobiles []Mobile
	mobiles = append(mobiles, m1)
	mobiles = append(mobiles, m2)

	f := Filter{}
	f.Filter(mobiles, filterBrand)
	f.Filter(mobiles, filterSize)
}
