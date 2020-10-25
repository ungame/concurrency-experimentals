package utils

import (
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func LoadFromFile(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return b
}

func GetRandom() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func BoolToInt(v bool) int8 {
	if v {
		return 1
	}
	return 0
}

func IntToBool(v int8) bool {
	return v > 0
}
