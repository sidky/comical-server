package handlers

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5String(id string) string {
	hasher := md5.New()
	hasher.Write([]byte(id))
	return hex.EncodeToString(hasher.Sum(nil))
}
