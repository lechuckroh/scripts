package logs

import (
	"encoding/json"
	"fmt"
	"github.com/bitfield/script"
)

// GetLastExecutdAt 함수는 지정한 람다 함수의 마지막 실행 시각을 반환합니다.
func GetLastExecutdAt(functionName string) (uint64, error) {
	// 람다 함수 로그 그룹명
	logGroupName := fmt.Sprintf("/aws/lambda/%s", functionName)

	// 최근 로그 스트림 조회
	cmd := fmt.Sprintf("aws logs describe-log-streams --log-group-name %s --order-by LastEventTime --descending --limit 1 --output json", logGroupName)
	output, err := script.Exec(cmd).String()
	if err != nil {
		return 0, fmt.Errorf("failed to execute command: %v", err)
	}

	// 결과 파싱
	var logStreams DescribeLogStreamsResponse
	if err := json.Unmarshal([]byte(output), &logStreams); err != nil {
		return 0, fmt.Errorf("failed to parse output: %v", err)
	}

	// 실행 로그가 없는 경우 에러 반환
	if len(logStreams.LogStreams) == 0 {
		return 0, fmt.Errorf("no log streams found for function: %s", functionName)
	}

	// 최근 실행 결과 반환
	lastExecutionTime := logStreams.LogStreams[0].LastEventTimestamp
	return lastExecutionTime, nil
}
