package utils

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func RunJudge(req JudgeRequest) JudgeResult {
	// 创建临时目录
	tmpDir, err := ioutil.TempDir("", "judge")
	if err != nil {
		return JudgeResult{"SystemError", "Cannot create temp dir"}
	}
	defer os.RemoveAll(tmpDir)

	// 保存源代码
	codePath := filepath.Join(tmpDir, "main.c")
	if req.Language == "cpp" {
		codePath = filepath.Join(tmpDir, "main.cpp")
	}
	if err := os.WriteFile(codePath, []byte(req.Code), 0o644); err != nil {
		return JudgeResult{"SystemError", "Write code failed"}
	}

	// 编译
	exePath := filepath.Join(tmpDir, "main")
	compiler := "gcc"
	if req.Language == "cpp" {
		compiler = "g++"
	}
	cmd := exec.Command(compiler, codePath, "-o", exePath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return JudgeResult{"CE", string(out)}
	}

	// 执行程序
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(req.TimeLimit)*time.Millisecond)
	defer cancel()

	runCmd := exec.CommandContext(ctx, exePath)
	inputData, _ := os.ReadFile(req.InputFile)
	runCmd.Stdin = bytes.NewReader(inputData)
	outputBuf := &bytes.Buffer{}
	runCmd.Stdout = outputBuf

	err = runCmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return JudgeResult{"TLE", "Time limit exceeded"}
	}
	if err != nil {
		return JudgeResult{"RE", err.Error()}
	}

	// 对比输出
	expected, _ := os.ReadFile(req.OutputFile)
	if string(outputBuf.Bytes()) == string(expected) {
		return JudgeResult{"AC", "Accepted"}
	} else {
		return JudgeResult{"WA", "Wrong Answer"}
	}
}
