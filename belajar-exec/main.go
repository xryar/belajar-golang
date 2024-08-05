package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var output1, _ = exec.Command("ls").Output()
	fmt.Printf("-> ls\n%s\n", string(output1))

	var output2, _ = exec.Command("/C").Output()
	fmt.Printf("-> pwd\n%s\n", string(output2))

	var output3, _ = exec.Command("git", "config", "user.name").Output()
	fmt.Printf("-> git config user.name\n%s\n", string(output3))

}

// direkomendasikan seperti ini
// func recomendationUse() {
// 	if runtime.GOOS == "windows" {
// 		output, err = exec.Command("cmd", "/C", "git config user.name").Output()
// 	} else {
// 		output, err = exec.Command("bash", "-c", "git config user.name").Output()
// 	}
// }
