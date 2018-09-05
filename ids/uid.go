package ids

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

// NewUUID 创建一个UUID
func NewUUID() string {
	u, _ := uuid.NewV4()
	raw := u.String()
	return strings.Replace(raw, "-", "", -1)
}

// UniqueRandID 唯一随机ID
func UniqueRandID(i int64) *big.Int {
	m := big.NewInt(i)
	n := big.NewInt(zoomRange)
	m.Mul(m, n)

	rd := rand.Int63n(zoomRange)
	m.Add(m, big.NewInt(rd))
	return m
}

var unixNanoID int64
var unixNanoIDLock sync.Mutex

// UnixNanoID 纳秒时间戳 1e9(Hz)
func UnixNanoID(namespace string) (string, error) {
	unixNanoIDLock.Lock()
	defer unixNanoIDLock.Unlock()

	id := time.Now().UnixNano()
	if id < unixNanoID {
		return "", errors.New("Generate a pass timestamp")
	}

	// concurrence incr
	if id == unixNanoID {
		id++
	}

	unixNanoID = id
	return fmt.Sprintf("%s-%d", namespace, id), nil
}
