package conf

import (
"fmt"
"testing"
)

func Test(t *testing.T) {

	//加载配置文件
	ini, err := NewConfig("./config.ini")
	if err != nil {
		t.Error(err)
		return
	}

	//获取配置1
	value, has := ini.GetValue("conf_01", "A")
	if has && value == "aaa" {
	} else {
		t.Errorf("配置值: %s, 获取值: %s", "A", value)
	}

	//获取配置2
	value, has = ini.GetValue("conf_02", "B")
	if has && value == "bbb" {
	} else {
		t.Errorf("配置值: %s, 获取值: %s", "B", value)
	}


	fmt.Println(ini)

}