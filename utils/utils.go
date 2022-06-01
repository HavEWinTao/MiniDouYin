package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
)

// AES加密密钥
// 用于生成随机密码
// 必须是16位, 可以优化成动态配置

// GetReturnData 打包数据和Code,Msg
func GetReturnData(dt interface{}, msgString string) *gin.H {
	var flag bool
	if msgString == "SUCCESS" {
		flag = true
	} else {
		flag = false
	}
	result := gin.H{"isSuccess": flag, "Data": dt}
	return &result
}

// GetSHAEncode 数据加密
// SHA256加密
func GetSHAEncode(str string) string {
	w := sha256.New()
	//将str写入到w中
	io.WriteString(w, str)
	//w.Sum(nil)将w的hash转成[]byte格式
	bw := w.Sum(nil)
	//将 bw 转成字符串
	shyster2 := hex.EncodeToString(bw)
	return shyster2
}
