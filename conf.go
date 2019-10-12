package goconf

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	filepath string
	conflist map[string]map[string]string
}

func NewConfig(filepath string) (*Config, error) {
	c := new(Config)
	c.filepath = filepath
	conflist, err := c.ReadList()
	if err != nil {
		return c, err
	}

	c.conflist = conflist
	return c, nil
}

//获取一个配置项
func (c *Config) GetValue(section, name string) (string, bool) {
	value, has := c.conflist[section][name]

	if has {
		return value, true
	}
	return "", false
}

//获取一个 bool 值配置
func (c *Config) GetBool(section, name string) bool {
	value, has := c.GetValue(section, name)
	if has {
		if strings.ToLower(value) == "true" {
			return true
		}
	}

	return false
}

//获取一个 string 配置
func (c *Config) GetString(section, name string) (string, bool) {
	return c.GetValue(section, name)
}

//获取一个 int 配置
func (c *Config) GetInt(section, name string) (int, bool) {
	value, has := c.GetValue(section, name)
	if has {
		n, err := strconv.Atoi(value)
		if err == nil {
			return n, true
		}
	}

	return 0, false
}

//读取配置
func (c *Config) ReadList() (map[string]map[string]string, error) {

	sectionData := make(map[string]map[string]string)
	kvData := make(map[string]string)

	file, err := os.Open(c.filepath)
	if err != nil {
		return sectionData, err
	}
	defer file.Close()

	var section string
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)

		if err != nil {
			i := strings.IndexAny(line, "=")
			if i >= 0 {
				kvData[strings.TrimSpace(line[0:i])] = strings.TrimSpace(line[i+1:])
			}
			if len(kvData) > 0 {
				sectionData[section] = kvData
			}

			return sectionData, nil
		}

		switch {
		case len(line) == 0:
		case line[0] == '#':
		case line[0] == '[' && line[len(line)-1] == ']':
			str := strings.TrimSpace(line[1 : len(line)-1])
			if section == "" {
				section = str
			}

			if section != str && len(kvData) > 0 {
				sectionData[section] = kvData
				kvData = make(map[string]string)
				section = str
			}

		default:
			i := strings.IndexAny(line, "=")
			kvData[strings.TrimSpace(line[0:i])] = strings.TrimSpace(line[i+1:])
		}

	}

	return sectionData, nil
}
