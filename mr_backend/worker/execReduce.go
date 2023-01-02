package worker

import (
	"fmt"
	"sort"
	"strings"
)

func (w *Worker) execReduce(part int, intermediaryFiles []string) string {

	keyBatches := make(map[string][]string)
	for _, file := range intermediaryFiles {
		for _, kv := range readKeyValues(file) {
			keyBatches[kv.Key] = append(keyBatches[kv.Key], kv.Value)
		}
	}

	keys := make([]string, len(keyBatches))
	i := 0
	for key := range keyBatches {
		keys[i] = key
		i++
	}
	sort.Slice(keys, func(i, j int) bool {
		return strings.Compare(keys[i], keys[j]) == -1
	})

	var outputLines []string
	for _, key := range keys {
		lines := w.impl.Reduce(key, keyBatches[key])
		outputLines = append(outputLines, lines...)
	}

	resultFile := fmt.Sprintf("result-%d.dat", part)
	writeLines(resultFile, outputLines)
	return resultFile
}
