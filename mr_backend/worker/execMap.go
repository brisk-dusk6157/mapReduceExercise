package worker

import (
	"fmt"
	"github.com/brisk-dusk6157/mapReduceExercise/mr_client"
)

func (w *Worker) execMap(taskId int, file string) (intermediaryFiles map[int]string) {
	content := readFileContent(file)

	kvs := w.impl.Map(file, content)

	// TODO: is it possible to do map[int][]*mr_client.KeyValue?
	partedKvs := make(map[int][]mr_client.KeyValue)
	for _, kv := range kvs {
		part := w.impl.Hash(kv.Key) % w.nParts
		partedKvs[part] = append(partedKvs[part], kv)
	}

	intermediaryFiles = make(map[int]string)
	for part, kvs := range partedKvs {
		intermediaryFiles[part] = fmt.Sprintf("intermediary-%d-%d.json", part, taskId)
		writeKeyValues(intermediaryFiles[part], kvs)
	}
	return
}
