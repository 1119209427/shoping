package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppPort string `json:"app_port"`
	AppHost  string `json:"app_host"`
	AppHttps bool   `json:"app_https"`
	DataBase DataBaseConfig `json:"data_base"`
	Redis RedisConfig `json:"redis"`
	Jwt JwtCfg `json:"jwt"`
	Email EmailConfig `json:"email"`
}
type JwtCfg struct {
	SigningKey string `json:"signing_key"`

}
type EmailConfig struct {
	ServiceEmail string `json:"service_email"`
	ServicePwd   string `json:"service_pwd"`
	SmtpPort     string `json:"smtp_port"`
	SmtpHost     string `json:"smtp_host"`
}
type DataBaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`

}
type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

var cfg *Config
//获取解析文件
func GetCfg () *Config{
	return cfg
}
func init(){
	err:=ParamConfig("./config/json")
	if err!=nil{
		panic(err)
	}
}
func ParamConfig(path string)error{
	file,err:=os.Open(path)
	if err!=nil{
		return err

	}
	defer file.Close()
	reader:=bufio.NewReader(file)
	decoder:=json.NewDecoder(reader)
	err=decoder.Decode(&cfg)
	if err!=nil{
		return err
	}
	return nil

}