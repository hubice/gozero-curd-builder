package utils

import "strings"

// DbType2Type 数据库转换成go
func DbType2Type(name string) string {
	list := map[string]string {
		"int": "int64",
		"timestamp": "int64",
		"date": "string",
		"varchar": "string",
		"char": "string",
		"text": "string",
	}
	for k, v := range list {
		if strings.Contains(name, k) {
			return v
		} else {
			return "string"
		}
	}
	return name
}

// Case2CamelLower 下划线转换成小写驼峰
func Case2CamelLower(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	name = strings.Replace(name, " ", "", -1)
	return strings.ToLower(name[0:1]) + name[1:]
}

// Case2CamelUpper 下划线转换成大写驼峰
func Case2CamelUpper(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// Case2Empty 下划线转换为空
func Case2Empty(name string) string {
	return strings.Replace(name, "_", "", -1)
}

// Case2Mid 下划线转中划线
func Case2Mid(name string) string {
	return strings.Replace(name, "_", "-", -1)
}
