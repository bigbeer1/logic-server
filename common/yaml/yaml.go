package yaml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/*
*
读取配置文件
*/
func ReadConfigByPath(file string, config interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	// 把yaml形式的字符串解析成struct类型
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return err
	}
	//fmt.Printf("Config: %+v", config)
	return nil
}

/*
*
写入配置文件
*/
func WriteConfigByPath(file string, config interface{}) error {
	// 对象转换
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	// 写入文件
	err = ioutil.WriteFile(file, data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func ReadConfig(name, folder string, config interface{}) error {
	// 如果没有后缀则补上
	if IsEmpty(filepath.Ext(name)) {
		name = fmt.Sprintf("%s.yml", name)
	}
	// 配置文件路径
	var configPath string
	if filepath.IsAbs(folder) {
		configPath = filepath.Join(folder, name)
	} else {
		configPath = filepath.Join(GetCurrentDirectory(), folder, name)
	}

	return ReadConfigByPath(configPath, config)
}

func WriteConfig(name, folder string, config interface{}) error {
	// 如果没有后缀则补上
	if IsEmpty(filepath.Ext(name)) {
		name = fmt.Sprintf("%s.yml", name)
	}
	// 配置文件路径
	configPath := GetCurrentDirectory()
	configPath = filepath.Join(configPath, folder, name)

	return WriteConfigByPath(configPath, config)
}

/*
*
字符串是否为空
*/
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

/*
获取程序运行路径
*/
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	//return strings.Replace(dir, "\\", "/", -1)
	return dir
}
