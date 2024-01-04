package utils

import (
	"encoding/base32"
	"encoding/hex"
	"strconv"
	"strings"
)

// Convert hex to []Bytes
// param hexx - An hex string
// return bytes
func Hex2Bt(hexx string) []byte {
	data, err := hex.DecodeString(hexx)
	if err != nil {
		panic(err)
	}
	return data
}

// Convert an []Bytes to hex
// param b - An []Bytes
// return string
func Bt2Hex(b []byte) string {
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b)
	return strings.ToUpper(string(dst))
}

func Address2Bt(a string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(a)
}

func Bt2Address(s []byte) string {
	return base32.StdEncoding.EncodeToString(s)
}

func Str2Uint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func Str2Uint8(s string) uint8 {
	v, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		panic(err)
	}
	return uint8(v)
}

func Str2Float64(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func Int64ToStr(s int64) string {
	return strconv.FormatInt(s, 10)
}

func StringToInt64(s string) *int64 {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil
	}
	return &i64
}
