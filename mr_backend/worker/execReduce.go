package worker

import "fmt"

func (w *Worker) execReduce(part int, intermediaryFiles []string) string {
	fmt.Printf("rTask done part=%d\n", part)
	return ""
}
