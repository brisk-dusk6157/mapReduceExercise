package worker

import (
	"encoding/json"
	"fmt"
	"github.com/brisk-dusk6157/mapReduceExercise/mr_client"
	"os"
	"strings"
)

// TODO: Use temp files and rename for atomic behavior

func readFileContent(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func writeKeyValues(filename string, kvs []mr_client.KeyValue) {
	kvsJson, err := json.Marshal(kvs)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filename, kvsJson, 0o644)
	if err != nil {
		panic(err)
	}
}

func readKeyValues(filename string) (kvs []mr_client.KeyValue) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(filename)
		panic(err)
	}
	err = json.Unmarshal(data, &kvs)
	if err != nil {
		panic(err)
	}
	return
}

func writeLines(filename string, lines []string) {
	content := strings.Join(lines, "\n")

	err := os.WriteFile(filename, []byte(content), 0o644)
	if err != nil {
		panic(err)
	}
}
