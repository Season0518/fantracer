package driver

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Bot OneBot配置的链接信息
type Bot struct {
	WebSocket   string `yaml:"websocket"`
	Http        string `yaml:"http"`
	AccessToken string `yaml:"access_token"`
}

type MySQL struct {
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

// Greeting 欢迎新成员文案
// : 加载文案时使用CQ码或占位符，简化程序逻辑
type Greeting struct {
	Target   int64    `yaml:"target"`
	Text     string   `yaml:"text"`
	MediaURL []string `yaml:"media_url"`
}

// Mail 邮件服务日志上报
type Mail struct {
	Interval int    `yaml:"post_interval"` //上报的间隔，单位为小时
	Addr     string `yaml:"smtp_address"`
	Port     int    `yaml:"smtp_port"`
	Account  string `yaml:"sender_account"`
	Key      string `yaml:"sender_key"`
	To       string `yaml:"mail_to"`
}

type Config struct {
	Bot      Bot        `yaml:"Bot"`
	MySQL    MySQL      `yaml:"MySQL"`
	Greeting []Greeting `yaml:"Greeting"`
	Mail     Mail       `yaml:"Mail"`
}

var Base *Config

// InitCfg 加载配置文件 #39c5bb
// Todo: 需要校验配置文件，部分字段不能为空或格式错误。
func InitCfg(absPath ...string) error {
	Base = &Config{}

	if len(absPath) == 0 {
		absPath = append(absPath, "config.yaml")
	} else if len(absPath) > 1 {
		return os.ErrInvalid
	}
	data, err := os.ReadFile(absPath[0])
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, Base)
	if err != nil {
		return err
	}

	return nil
}
