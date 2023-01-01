package worker

import (
	"fmt"
)

func (w *Worker) execReduce(part int, intermediaryFiles []string) string {

	keyBatches := make(map[string][]string)
	for _, file := range intermediaryFiles {
		for _, kv := range readKeyValues(file) {
			keyBatches[kv.Key] = append(keyBatches[kv.Key], kv.Value)
		}
	}

	var outputLines []string
	for key, batch := range keyBatches {
		lines := w.impl.Reduce(key, batch)
		outputLines = append(outputLines, lines...)
	}

	resultFile := fmt.Sprintf("result-%d.dat", part)
	writeLines(resultFile, outputLines)
	return resultFile
}
