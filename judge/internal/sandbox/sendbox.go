package sandbox

import (
	"fmt"
	"os/exec"
	"strconv"
)

type Sendbox struct {
	Boxid string
}

func InitSendbox() (Sendbox, error) {
	var sd Sendbox
	var st bool
	boxID := 0

	// 初始化沙箱
	err := exec.Command("isolate", "--box-id="+strconv.Itoa(boxID), "--init").Run()
	if err != nil {
		st = true
	}
	for st {
		boxID += 1
		err = exec.Command("isolate", "--box-id="+strconv.Itoa(boxID), "--init").Run()
		if err == nil {
			st = false
		}
		if boxID > 999 {
			return sd, fmt.Errorf("无法创建沙箱")
		}
	}
	sd.Boxid = strconv.Itoa(boxID)
	return sd, nil
}
