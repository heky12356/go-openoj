package service

import (
	"fmt"
	"os"
	"os/exec"
)

func ServiceSubmit(code string) (string, error) {
	// 将用户代码保存到文件中
	err := os.WriteFile("submissions/test.c", []byte(code), 0o644)
	if err != nil {
		return "", fmt.Errorf("write file err: %s", err)
	}

	// 编译用户代码
	cmd := exec.Command("gcc", "submissions/test.c", "-o", "tmp/a.out")
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("编译错误： %s", err)
	}
	// 执行 isolate 进行判题
	err = exec.Command("isolate", "--cleanup").Run()
	if err != nil {
		return "", fmt.Errorf("--cleanup err: %s", err)
	}
	err = exec.Command("isolate", "--init").Run()
	if err != nil {
		return "", fmt.Errorf("--init err: %s", err)
	}

	// 将文件复制到沙箱中
	err = exec.Command("cp", "testcases/in1.txt", "/var/lib/isolate/0/box/in1.txt").Run()
	if err != nil {
		return "", fmt.Errorf("cp in1 err: %s", err)
	}
	err = exec.Command("cp", "tmp/a.out", "/var/lib/isolate/0/box/a.out").Run()
	if err != nil {
		return "", fmt.Errorf("cp a.out err: %s", err)
	}

	runCmd := exec.Command("isolate",
		"--run",
		"--stdin=in1.txt",
		"--stdout=user_out.txt",
		"--",
		"a.out",
	)
	err = runCmd.Run()
	if err != nil {
		return "", fmt.Errorf("run err: %s", err)
	}
	userOutput, _ := os.ReadFile("/var/lib/isolate/0/box/user_out.txt")
	expectedOutput, _ := os.ReadFile("testcases/out1.txt")

	if string(userOutput) == string(expectedOutput) {
		return "Accepted", nil
	} else {
		return "Wrong Answer", nil
	}
}
