package main

import (
	"fmt"

	"go-openoj/internal/utils"
)

func main() {
	code := `
#include <stdio.h>
int main() {
    int a, b;
    scanf("%d%d", &a, &b);
    printf("%d\n", a + b);
    return 0;
}
`
	req := utils.JudgeRequest{
		Code:       code,
		Language:   "c",
		TimeLimit:  1000,
		InputFile:  "testdata/input.txt",
		OutputFile: "testdata/output.txt",
	}
	result := utils.RunJudge(req)
	fmt.Printf("Result: %s\nDetail: %s\n", result.Status, result.Message)
}
