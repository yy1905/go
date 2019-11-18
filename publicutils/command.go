package publicutils

import (
	"os/exec"
	"fmt"
)

func Command(cmd string)(cmdOut string)  {

	//cmd := "cat /proc/cpuinfo | egrep '^model name' | uniq | awk '{print substr($0, index($0,$4))}'"

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	res := string(out)
	return res

}
