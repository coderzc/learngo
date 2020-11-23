package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s1 := " aBc"
	s2 := "100a"
	s3 := s1 + s2
	fmt.Println(s3)
	fmt.Println(strings.HasPrefix(s3, "a"))               //判断前缀
	fmt.Println(strings.HasSuffix(s3, "0"))               //判断后缀
	fmt.Println(strings.Contains(s3, "9"))                //字符串包含关系
	fmt.Println(strings.Index(s3, "0"))                   //判断子字符串或字符在父字符串中出现的位置（索引）
	fmt.Println(strings.LastIndex(s3, "0"))               //最后出现位置的索引
	fmt.Println(strings.Replace(s3, "0", "1", -1))        //如果 n = -1 则替换所有字符串
	fmt.Println(strings.Count(s3, "0"))                   //出现的非重叠次数
	fmt.Println(strings.Repeat(s3, 2))                    //重复字符串
	fmt.Println(strings.ToLower(s3))                      //修改字符串大小写
	fmt.Println(strings.ToUpper(s3))                      //修改字符串大小写
	fmt.Println(strings.TrimSpace(s3))                    //修剪字符串 去掉开头和结尾空格
	fmt.Println(strings.Trim(strings.TrimSpace(s3), "a")) //修剪字符串 去掉开头和结尾字符串

	//截取含中文子串,因为go的是string是变长编码，所以要先转[]rune
	runes := []rune(s2)
	fmt.Println(string(runes[:len(runes)-1]))

	// 分割字符串
	split := strings.Split("1,2,3,4,5,6,", ",")
	for k, v := range split {
		fmt.Println(k, v)
	}

	// parseInt
	if i, err := parsInt("111"); err == nil {
		fmt.Println("parsInt:", i)
	}
}

func parsInt(orig string) (int, error) {
	an, err := strconv.Atoi(orig)
	if err != nil {
		//fmt.Println(err)
		return 0, fmt.Errorf("orig %s is not an integer - exiting with error\n", orig)
	}
	return an, nil
}
