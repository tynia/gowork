package hash

import (
	"io"
	"sync"
	"time"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
)

var (
	autoIncId = uint64(0)
	autoIncIdMutex = sync.Mutex{}
)


func GenHashString32(ts uint64, key string) string {
	var id string
	var prefix string
	var suffix string
	var tmpIncId uint64

	autoIncIdMutex.Lock()
	tmpIncId = autoIncId
	autoIncId += 1
	autoIncIdMutex.Unlock()

	if ts <= 0 {
		ts = uint64(time.Now().Unix())
	}

	prefix = strconv.FormatUint(ts, 10)
	hs := sha256.New()
	io.WriteString(hs, key)
	strHash := hs.Sum(nil)
	suffix = hex.EncodeToString(strHash[len(key):len(key)+8])
	id = prefix + suffix + strconv.FormatUint(100000 + tmpIncId, 10)

	return id
}

