package utils

import (
	"os"
	"fmt"
)

func DeleteFile(path string) {
	// delete file
	var err = os.Remove(path)
	if err != nil {
		return
	}
	fmt.Println("==> done deleting file")
}
