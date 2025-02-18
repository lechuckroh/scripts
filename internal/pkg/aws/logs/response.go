package logs

// DescribeLogStreamsResponse represents the response structure for log streams
type DescribeLogStreamsResponse struct {
	LogStreams []LogStream `json:"logStreams"`
}

// LogStream represents a single log stream in the response
type LogStream struct {
	LogStreamName      string `json:"logStreamName"`
	LastEventTimestamp uint64 `json:"lastEventTimestamp"`
}
