package main

import (
	"fmt"
	"github.com/brisk-dusk6157/mapReduceExercise/mr_client"
	"math/rand"
	"strings"
	"time"
)

func Map(filename string, content string) (out []mr_client.KeyValue) {
	if rand.Intn(100) < 85 {
		time.Sleep(5 * time.Second)
	}
	for _, word := range strings.Fields(content) {
		out = append(out, mr_client.KeyValue{word, "1"})
	}
	return
}

func Reduce(key string, values []string) (results []string) {
	if rand.Intn(100) < 85 {
		time.Sleep(5 * time.Second)
	}
	results = append(results, fmt.Sprintf("%s %d", key, len(values)))
	return
}
