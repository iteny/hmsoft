package common

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
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
