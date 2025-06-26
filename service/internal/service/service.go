package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"go-openoj/service/internal/define"
)

func ServiceSubmit(codedata define.Submit) (string, error) {
	// 这里之后要改成从数据库里读取
	stdin, err := os.ReadFile("../../static/testdata/input.txt")
	if err != nil {
		return "", fmt.Errorf("read stdin err: %s", err)
	}
	stdout, err := os.ReadFile("../../static/testdata/output.txt")
	if err != nil {
		return "", fmt.Errorf("read stdout err: %s", err)
	}
	// log.Print(string(stdin))
	url := "http://localhost:5050/submit"

	// 构造请求体结构体
	payload := map[string]string{
		"code":  codedata.Code,
		"stdin": string(stdin),
	}

	// 编码为 JSON
	data, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("json marshal error: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	log.Printf("body: %s", body)
	var res define.Judgeresp
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", err
	}
	log.Print(res.Output)
	log.Print(string(stdout))
	if res.Output == string(stdout) {
		return fmt.Sprintf("Result: %s", "Accepted"), nil
	}
	return fmt.Sprintf("Result: %s", "Wrong Answer"), nil
}
