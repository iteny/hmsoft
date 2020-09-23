package common

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"strconv"
	"unsafe"
)

func (c *BaseFunc) Md5(s string) string {
	hash := md5.New()
	buf := []byte(s)
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
func (c *BaseFunc) Sha1(s string) string {
	hash := sha1.New()
	buf := []byte(s)
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//Sha1 Plus Md5
func (c *BaseFunc) Sha1PlusMd5(s string) string {
	return c.Sha1(c.Md5(s))
}

// 字符串转整形
func (c *BaseFunc) ConvertToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

// 整形转字符串
func (c *BaseFunc) ConvertToString(s int) string {
	str := strconv.Itoa(s)
	return str
}

// 字符串转到成数组byte
func (c *BaseFunc) Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s)) // 获取s的起始地址开始后的两个 uintptr 指针
	h := [3]uintptr{x[0], x[1], x[1]}      // 构造三个指针数组
	return *(*[]byte)(unsafe.Pointer(&h))
}
