package main

import (
	"github.com/brisk-dusk6157/mapReduceExercise/mr_client"
	"strings"
)

func Map(filename string, content string) (out []mr_client.KeyValue) {
	for _, word := range strings.Fields(content) {
		out = append(out, mr_client.KeyValue{word, "1"})
	}
	return
}

func Reduce(key string, values []string) []string {
	return nil
}
