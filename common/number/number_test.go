package number

import (
	"fmt"
	"testing"
	"time"
)

func TestNewNumber(t *testing.T) {
	var aa = ""
	for i := 0; i < 10000; i++ {
		kk, err := NextSerialNumber("kkk", 1, 0, 0, 2, aa)
		if err != nil {
			fmt.Println(err)

		}
		aa = kk
		fmt.Println(kk)
		time.Sleep(time.Second / 5)
	}

}
