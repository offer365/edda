package tools

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// 随机生成字符串指定个数的字符串
func RandString(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano() + rand.Int63())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RunTime(start int64) string {
	online := time.Now().Unix() - start
	d := online / 86400
	h := (online - d*86400) / 3600
	m := (online - d*86400 - h*3600) / 60
	s := online - d*86400 - h*3600 - m*60
	return fmt.Sprintf("在线 %02d天%02d小时%02d分钟%02d秒.", d, h, m, s)
}

func Md5Sum(obj interface{}) string {
	if obj == nil {
		return ""
	}
	code, err := json.Marshal(obj)
	if err != nil {
		return ""
	}

	h := md5.New()
	h = sha256.New()
	h.Write(code)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}