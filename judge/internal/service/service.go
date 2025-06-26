package service

import (
	"fmt"
	"os"
	"os/exec"

	"go-judge/internal/sandbox"
)

func ServiceSubmit(code string, stdin string) (string, error) {
	// 初始化沙箱
	sd, err := sandbox.InitSendbox()
	if err != nil {
		return "", err
	}
	boxID := sd.Boxid
	boxPath := "/var/lib/isolate/" + boxID + "/box/"
	metaPath := "/var/lib/isolate/" + boxID + "/meta.txt"

	// 将用户代码保存到文件中
	err = os.WriteFile("submissions/test.c", []byte(code), 0o644)
	if err != nil {
		return "", fmt.Errorf("write file err: %s", err)
	}

	// 保存标准输入
	err = os.WriteFile("submissions/stdin.txt", []byte(stdin), 0o644)
	if err != nil {
		return "", fmt.Errorf("write file err: %s", err)
	}

	// 编译用户代码
	cmd := exec.Command("gcc", "submissions/test.c", "-o", "tmp/a.out")
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("编译错误： %s", err)
	}

	// 将文件复制到沙箱中
	err = exec.Command("cp", "submissions/stdin.txt", boxPath+"stdin.txt").Run()
	if err != nil {
		return "", fmt.Errorf("cp stdin err: %s", err)
	}
	err = exec.Command("cp", "tmp/a.out", boxPath+"a.out").Run()
	if err != nil {
		return "", fmt.Errorf("cp a.out err: %s", err)
	}

	runCmd := exec.Command("isolate",
		"--box-id="+boxID,
		"--run",
		// 下面两行都是在沙箱内获取in1，和输出user_out
		"--stdin=stdin.txt",
		"--stdout=user_out.txt",
		// 这个meta反而是在外文件夹生成
		"--meta="+metaPath,
		"--",
		"a.out",
	)
	err = runCmd.Run()
	if err != nil {
		return "", fmt.Errorf("run err: %s", err)
	}
	userOutput, _ := os.ReadFile(boxPath + "user_out.txt")
	return string(userOutput), nil
}
