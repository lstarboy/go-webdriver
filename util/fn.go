package util

import (
	"encoding/base64"
	"math/rand"
	"os"
	"time"
)

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 97
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func RandInt(start, end int) int {
	if end <= start {
		return start
	}
	mid := end - start
	return start + rand.Int()%mid
}

func SaveImage(fileName string, base64Str string) {
	//解压
	dist, _ := base64.StdEncoding.DecodeString(base64Str)
	//写入新文件
	f, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer func() {
		_ = f.Close()
	}()

	_, _ = f.Write(dist)
}
