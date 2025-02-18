package main

import (
	"fmt"
	"github.com/lechuckroh/scripts/internal/pkg/aws/lambda"
	"github.com/lechuckroh/scripts/internal/pkg/aws/logs"
	"github.com/lechuckroh/scripts/internal/pkg/timeutil"
	"sort"
)

func listLastExecutedTimes() {
	functionNames, err := lambda.ListFunctionNames()
	if err != nil {
		fmt.Printf("failed to fetch lambda function names: %v", err)
		return
	}

	sort.Strings(functionNames)

	for _, name := range functionNames {
		lastExecutdAt, err := logs.GetLastExecutdAt(name)
		if err != nil {
			fmt.Printf("null %s\n", name)
			continue
		}

		lastExecutedTime := timeutil.FromMillis(int64(lastExecutdAt))
		fmt.Printf("%v %s\n", lastExecutedTime, name)
	}
}

func main() {
	listLastExecutedTimes()
}
