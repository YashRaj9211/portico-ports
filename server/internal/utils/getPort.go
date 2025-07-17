package utils

import (
	"fmt"
	"sync"
)

var (
	currPort  = 30000
	portMutex = sync.RWMutex{}
)

func GetPort() string {
	portMutex.Lock()
	defer portMutex.Unlock()
	currPort++
	return fmt.Sprintf(":%d", currPort)
}