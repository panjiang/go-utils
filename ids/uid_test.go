package ids

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewUUID(t *testing.T) {
	t.Log(NewUUID())
	t.Log(NewUUID())
	t.Log(NewUUID())
}

func TestUniqueRandID(t *testing.T) {
	t.Log(UniqueRandID(1))
	t.Log(UniqueRandID(2))
	t.Log(UniqueRandID(3))
	t.Log(UniqueRandID(100000))
}

func TestRandZoomUID(t *testing.T) {
	rand.Seed(time.Now().Unix())
	t.Log(RandZoomUID(0, 123, 10000))
	t.Log(RandZoomUID(1, 123, 10000))
	t.Log(RandZoomUID(2, 123, 10000))
	t.Log(RandZoomUID(3, 123, 10000))
	t.Log(RandZoomUID(300, 123, 10000))
	t.Log(RandZoomUID(30000, 123, 10000))
	t.Log(RandZoomUID(30001, 123, 10000))
	t.Log(RandZoomUID(300000, 123, 10000))
}
