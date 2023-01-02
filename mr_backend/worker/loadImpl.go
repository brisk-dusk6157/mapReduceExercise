package worker

import (
	"github.com/brisk-dusk6157/mapReduceExercise/mr_client"
	"hash/fnv"
	"log"
	"plugin"
)

type Impl struct {
	Map       func(string, string) []mr_client.KeyValue
	Reduce    func(string, []string) []string
	Partition func(string, int) int
}

func loadImpl(path string) Impl {
	p, err := plugin.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	xmapf, err := p.Lookup("Map")
	if err != nil {
		log.Fatalf("cannot find Map in %v", path)
	}
	mapf := xmapf.(func(string, string) []mr_client.KeyValue)
	xreducef, err := p.Lookup("Reduce")
	if err != nil {
		log.Fatalf("cannot find Reduce in %v", path)
	}
	reducef := xreducef.(func(string, []string) []string)
	return Impl{
		Map:       mapf,
		Reduce:    reducef,
		Partition: partition,
	}
}

func partition(key string, nParts int) int {
	return ihash(key) % nParts
}

func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}
