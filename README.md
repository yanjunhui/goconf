# goconf
配置文件读取


## 初始化

```go
conf, err := NewConfig("./config.ini")
```


## 获取一个值

```go
conf, has = ini.GetValue("conf_02", "B")
```

## 获取整型(int)类型值

```go
valueN, has := conf.GetInt("conf_01", "NumberName")
```

## 获取字符串类型值

```go
value, has := conf.GetString("conf_01", "StringName")
```


## 获取一个布尔类型

```go
ok := conf.GetString("conf_01", "BoolName")
```



## 完整示例
```go

	//加载配置文件
	conf, err := NewConfig("./config.conf")
	if err != nil {
		t.Error(err)
		return
	}

	//获取字符串配置
	value, has := conf.GetString("conf_01", "StringName")
	if has && value == "aaa" {
	} else {
		t.Errorf("获取字符配置项: %s, 获取值: %s", "StringName", value)
	}


	//获取整型数值配置
	valueN, has := conf.GetInt("conf_01", "NumberName")
	if has && valueN == 123  {
	} else {
		t.Errorf("获取整型配置项: %s, 获取值: %d", "NumberName", valueN)
	}


	//获取布尔类型配置
	valueB := conf.GetBool("conf_01", "BoolName")
	if valueB  {
	} else {
		t.Error("获取布尔值配置项失败")
	}


	//获取配置2
	value, has = conf.GetValue("conf_02", "B")
	if has && value == "bbb" {
	} else {
		t.Errorf("配置值: %s, 获取值: %s", "B", value)
	}


```
