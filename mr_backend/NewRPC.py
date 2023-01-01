import argparse
from pathlib import Path

coordinator_path_template = "coordinator/handle{basename}.go"
coordinator_file_template = """package coordinator

import "github.com/brisk-dusk6157/mapReduceExercise/mr_backend/schemas"

func (c *Coordinator) {basename}(args *schemas.{basename}Args, reply *schemas.{basename}Reply) error {{
	return nil
}}
"""

worker_path_template = "worker/call{basename}.go"
worker_file_template = """package worker

import (
	"log"
	"github.com/brisk-dusk6157/mapReduceExercise/mr_backend/schemas"
)

func (w *Worker) call{basename}() schemas.{basename}Reply {{
	args := schemas.{basename}Args{{}}
	reply := schemas.{basename}Reply{{}}
	err := w.call("Coordinator.{basename}", &args, &reply)
	if err != nil {{
		log.Fatal("{basename} call failed: ", err)
	}}
	return reply
}}
"""

struct_path_template = "schemas/{basename}.go"
struct_file_template = """package schemas

type {basename}Args struct {{
}}

type {basename}Reply struct {{
}}
"""

parser = argparse.ArgumentParser()
parser.add_argument("basename", nargs="+")

if __name__ == '__main__':
    args = parser.parse_args()

    for basename in args.basename:
        Path(struct_path_template.format(basename=basename)).write_text(
            struct_file_template.format(basename=basename))
        Path(coordinator_path_template.format(basename=basename)).write_text(
            coordinator_file_template.format(basename=basename))
        Path(worker_path_template.format(basename=basename)).write_text(
            worker_file_template.format(basename=basename))
