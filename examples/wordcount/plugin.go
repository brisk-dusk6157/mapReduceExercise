package main

import (
	"fmt"
	"github.com/brisk-dusk6157/mapReduceExercise/mr_client"
	"strings"
)

func Map(filename string, content string) (out []mr_client.KeyValue) {
	for _, word := range strings.Fields(content) {
		out = append(out, mr_client.KeyValue{word, "1"})
	}
	return
}

func Reduce(key string, values []string) (results []string) {
	results = append(results, fmt.Sprintf("%s %d", key, len(values)))
	return
}
