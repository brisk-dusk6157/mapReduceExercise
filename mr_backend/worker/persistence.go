package worker

import (
	"encoding/json"
	"github.com/brisk-dusk6157/mapReduceExercise/mr_client"
	"log"
	"os"
)

func readFileContent(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to open %s", filename)
	}
	return string(data)
}

func writeKeyValues(filename string, kvs []*mr_client.KeyValue) {
	kvsJson, err := json.Marshal(kvs)
	if err != nil {
		log.Fatalf("Error marshaling to json")
	}
	err = os.WriteFile(filename, kvsJson, 0o644)
	if err != nil {
		log.Fatalf("Failed to write %s", filename)
	}
}
