package main

import (
	"fmt"
	"os"

	"github.com/wangyoucao577/go-project-layout/stringutil"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(stringutil.Reverse(os.Args[1]))
	} else {
		fmt.Println() // print empty string
	}
}
