package util

import (
	"testing"
)

func TestCheckPassword(t *testing.T) {
	inputPassword := "82790085228cf8a1e3bac41f45271e5f"
	userPassword := "c8318217c958f2e9897babd86838ebad3f2de81858a46e10cc31f28e82be9f72"
	passwordStr := "DZ8#*>"
	password, err := CheckPassword(inputPassword, userPassword, passwordStr)
	t.Log(password, err)
	//	str := "123456"
	//	//方法一
	//	data := []byte(str)
	//	has := md5.Sum(data)
	//	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	//	fmt.Println(md5str1)
	//	//方法二
	//	w := md5.New()
	//	io.WriteString(w, str)
	//	//将str写入到w中
	//	md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	//
	//	fmt.Println(md5str2)

}
