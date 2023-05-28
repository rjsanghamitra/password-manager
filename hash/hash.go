package hash

import (
	"crypto/md5"
	"encoding/hex"
	"hash/fnv"
	"strconv"
)

// this function generates hash values based on the fnv-1a algorithm
func Hash(s string) string {
	temp := fnv.New32a()
	temp.Write([]byte(s))
	return strconv.FormatUint(uint64(temp.Sum32()), 10)
}

// this function generates hash values using the MD5 hashing algorithm
func GetMD5Hash(s string) string {
	temp := md5.Sum([]byte(s))
	return hex.EncodeToString(temp[:])
}
