package conf

import (
	"testing"
)

func Test(t *testing.T) {

	//加载配置文件
	ini, err := NewConfig("./config.ini")
	if err != nil {
		t.Error(err)
		return
	}

	//获取字符串配置
	value, has := ini.GetString("conf_01", "StringName")
	if has && value == "aaa" {
	} else {
		t.Errorf("获取字符配置项: %s, 获取值: %s", "StringName", value)
	}


	//获取整型数值配置
	valueN, has := ini.GetInt("conf_01", "NumberName")
	if has && valueN == 123  {
	} else {
		t.Errorf("获取整型配置项: %s, 获取值: %d", "NumberName", valueN)
	}


	//获取布尔类型配置
	valueB := ini.GetBool("conf_01", "BoolName")
	if valueB  {
	} else {
		t.Error("获取布尔值配置项失败")
	}


	//获取配置2
	value, has = ini.GetValue("conf_02", "B")
	if has && value == "bbb" {
	} else {
		t.Errorf("配置值: %s, 获取值: %s", "B", value)
	}

}