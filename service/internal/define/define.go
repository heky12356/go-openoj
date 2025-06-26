package define

type Submit struct {
	Language string `json:"language"` // 语言
	Code     string `json:"code"`     // 源代码
}

type Judgeresp struct {
	Output string `json:"output"`
}
