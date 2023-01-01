package worker

import "fmt"

func (w *Worker) execMap(file string) map[int]string {
	fmt.Printf("mTask done file=%s\n", file)
	return nil
}
