package pkg

import (
	"fmt"
	"testing"
)

func TestEncryption(t *testing.T) {
	encrypFile, err := Encryption("G:\\backupAgent\\domain\\pkg\\printlogo.go")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(encrypFile)
}
