package executor

import (
	"fmt"
	_ "os/exec"
)

func Execute(files string, outputPrefix string, inPlace string) error {
	//exec.Command("short", "-k", "-f", filesSlice...)

	fmt.Println(files)
	fmt.Println(outputPrefix)
	fmt.Println(inPlace)

	return nil
}
