package main

import "fmt"

func main() {
	var siteMap map[string]string
	siteMap = make(map[string]string)

	siteMap["Google"] = "谷歌"
	siteMap["Runoob"] = "菜鸟教程"
	siteMap["Baidu"] = "百度"
	siteMap["Wiki"] = "维基百科"

	for site := range siteMap {
		fmt.Println(site, "首都是", siteMap[site])
	}

	name, ok := siteMap["Facebook"]

	if ok {
		fmt.Println("Facebook 的 站点是", name)
	} else {
		fmt.Println("Facebook 站点不存在")
	}
}
