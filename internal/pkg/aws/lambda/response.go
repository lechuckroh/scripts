package lambda

type Function struct {
	FunctionName string `json:"FunctionName"`
}

type ListFunctionsResponse struct {
	Functions []Function `json:"Functions"`
	NextToken string     `json:"NextToken"`
}
