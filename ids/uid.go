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

// RandZoomUID 把唯一ID随机放大scale倍再加上base，结果仍保持唯一
// 无法反推原UID，用于生成一个对外的唯一ID
func RandZoomUID(i int64, scale int64, base int64) int64 {
	return base + i*scale + rand.Int63n(scale)
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
