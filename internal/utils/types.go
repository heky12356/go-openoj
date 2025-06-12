package utils

type JudgeRequest struct {
	Code       string // 源代码内容
	Language   string // "c" 或 "cpp"
	TimeLimit  int    // ms
	InputFile  string
	OutputFile string
}

type JudgeResult struct {
	Status  string // "AC", "WA", "TLE", "RE", "CE"
	Message string
}
