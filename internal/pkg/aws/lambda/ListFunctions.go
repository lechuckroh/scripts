package lambda

import (
	"encoding/json"
	"fmt"
	"github.com/bitfield/script"
)

func ListFunctionNames() ([]string, error) {
	functions, err := ListFunctions()
	if err != nil {
		return nil, err
	}
	var functionNames []string
	for _, res := range functions {
		functionNames = append(functionNames, res.FunctionName)
	}
	return functionNames, nil
}

func ListFunctions() ([]Function, error) {
	var functions []Function
	nextToken := ""

	for {
		// 람다함수 목록 페이지네이션 조회 커맨드
		cmd := "aws lambda list-functions --output json"
		if nextToken != "" {
			cmd += fmt.Sprintf(" --starting-token %s", nextToken)
		}

		// 커맨드 실행
		output, err := script.Exec(cmd).String()
		if err != nil {
			return nil, fmt.Errorf("cmd: %s, error: %v", cmd, err)
		}

		// 실행결과 파싱
		var res ListFunctionsResponse
		if err := json.Unmarshal([]byte(output), &res); err != nil {
			return nil, fmt.Errorf("failed to parse output: %v", err)
		}

		// 결과 처리
		for _, f := range res.Functions {
			functions = append(functions, f)
		}

		// 다음 페이지가 없는 경우 종료
		if res.NextToken == "" {
			break
		}
		nextToken = res.NextToken
	}

	return functions, nil
}
