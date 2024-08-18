package main

import (
	"fmt"
	logs "go-studies/package_logs"
)

func main() {
	log := "please replace '👎' with '👍'"
	replacedLog := logs.Replace(log, '👎', '👍')
	fmt.Println(replacedLog)
}
