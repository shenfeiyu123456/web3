package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

/*
*
函数用func 声明
函数可以有一个或者多个参数，需要有参数类型，用，分割
函数可以有一个或者多个参数进行返回，需要有返回值类型用，分割
函数的参数是可以选的，返回值也是可以选的
*/
/**
值传递，传递参数时，将参数复制一份传递到函数中，对参数进行调整后，不影响参数值，Go语言默认值传递
引用传递，传递参数时，将参数的地址传递到函数中，对参数进行调整后，影响参数值
*/
type ResultMsg struct {
	Code int
	Data interface{}
	Msg  string
}

func setData(ResultMsg *ResultMsg) {
	ResultMsg.Code = 1
	ResultMsg.Data = map[string]interface{}{
		"name":    "张三",
		"age":     18,
		"address": "郑州",
	}
	ResultMsg.Msg = "success"
}

func jsonPrint(ResultMsg *ResultMsg) {
	jsons, err := json.Marshal(ResultMsg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsons))
}

// MD5 方法
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// 格式刷时间
func getTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前时间字符串
func getCurrentTimeInt() int64 {
	return time.Now().Unix()
}

// 生成签名
func createSin(param map[string]interface{}) string {
	var key []string
	var str = ""
	for k, _ := range param {
		key = append(key, k)
	}
	fmt.Println("key", key)
	sort.Strings(key)
	fmt.Println("排序后的key", key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v\n", key[i], param[key[i]])
			fmt.Println("str", str)
		} else {
			str = str + fmt.Sprintf("&xl_%v=%v", key[i], param[key[i]])
			fmt.Println("str", str)
		}
	}
	//自定义密钥
	secreate := "123456"
	//自定义签名算法
	sign := MD5(MD5(str) + MD5(secreate))
	return sign
}

func main() {
	ResultMsg := ResultMsg{
		Code: 1,
		Data: map[string]interface{}{
			"姓名": "张三",
			"年纪": 24,
			"地址": "郑州",
		},
		Msg: "success",
	}
	jsonPrint(&ResultMsg)
	str := "12345"
	var str1 string
	str1 = MD5(str)
	//str1 := MD5(str)
	fmt.Println(str1)
	timestr := getTimeStr()
	fmt.Println(timestr)
	timemills := getCurrentTimeInt()
	fmt.Println(timemills)
	param := map[string]interface{}{
		"name": "tom",
		"pwd":  "123456",
		"age":  18,
	}
	fmt.Println(createSin(param))
}
