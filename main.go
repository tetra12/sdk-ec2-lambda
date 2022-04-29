package main

import (
	"context"

	aj "github.com/choria-io/asyncjobs"
)

func LabdaEC2(_ context.Context, log aj.Logger, task *aj.Task) (interface{}, error) {
	log.Infof("Labda EC2: Handling task %s", task.ID)
	return "success", nil
}
