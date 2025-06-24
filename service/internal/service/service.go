package service

import (
	"fmt"

	"go-openoj/service/internal/define"
	"go-openoj/service/internal/utils"
)

func ServiceSubmit(codedata define.Submit) (string, error) {
	// 保存源代码到数据库
	req := utils.JudgeRequest{
		Code:       codedata.Code,
		Language:   codedata.Language,
		TimeLimit:  1000,
		InputFile:  "static/testdata/input.txt",
		OutputFile: "static/testdata/output.txt",
	}
	result := utils.RunJudge(req)
	return fmt.Sprintf("Result: %s\nDetail: %s\n", result.Status, result.Message), nil
}
