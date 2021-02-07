package app

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func UpdateNumer(ctx workflow.Context, mess string) error {
	fmt.Println("Entering UpdateNumer Workflow")
	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    500,
	}
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retrypolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	fmt.Println("Executing the GetInt")
	err := workflow.ExecuteActivity(ctx, GetInt).Get(ctx, nil)
	if err != nil {
		return err
	}
	fmt.Println("Executing the UpdateInt")
	err = workflow.ExecuteActivity(ctx, UpdateInt).Get(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}
