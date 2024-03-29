package CO1Cache

import (
	"fmt"
	"os"
	"time"
)

func Verify(file string) bool {
	f, err := os.Stat(Filename(file))

	if os.IsNotExist(err) {
		return false
	}

	if err != nil {
		fmt.Printf("Unable to read cache file: %v\n", err)
		return false
	}

	return !f.IsDir() && time.Now().Sub(f.ModTime()) < time.Second * time.Duration(_CacheDuration())
}
