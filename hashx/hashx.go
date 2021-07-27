package hashx

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Md5(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256(data string ) string  {
	hash := sha256.New()
	//输入数据
	hash.Write([]byte(data))
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode
}