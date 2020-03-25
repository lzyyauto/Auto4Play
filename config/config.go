package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

type Environment string

const (
	EnvProd  Environment = "production"
	EnvDev   Environment = "development"
	EnvLocal Environment = "local"
)

var (
	// Env          Environment
	// ServerCfg    *ServerConfiguration
	// RedisCfg     *RedisConfiguration
	MysqlCfg *MysqlConfiguration
	// AlipayCfg    *AlipayConfiguration
	// WechatPayCfg *WechatPayConfiguration
	// AccountCfg   *AccountConfiguration
	BarkCfg *BarkConfiguration
)

type configuration struct {
	// Env          Environment             `json:"env"`
	// ServerCfg    *ServerConfiguration    `json:"server"`
	// RedisCfg     *RedisConfiguration     `json:"redis"`
	MysqlCfg *MysqlConfiguration `json:"mysql"`
	// AlipayCfg    *AlipayConfiguration    `json:"alipay"`
	// WechatPayCfg *WechatPayConfiguration `json:"wechatpay"`
	// AccountCfg   *AccountConfiguration   `json:"account"`
	BarkCfg *BarkConfiguration `json:"bark"`
}

type MysqlConfiguration struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type BarkConfiguration struct {
	BarkUrl string `json:"barkurl"`
	BarkKey string `json:"barkkey"`
}

func init() {
	var configBytes []byte
	var err error
	configPath := flag.String("config", "./config/config.json", "配置文件路径")
	configBytes, err = ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("Load configuration failed: %s\n", err)
	}

	var conf configuration
	if err = json.Unmarshal(configBytes, &conf); err != nil {
		log.Fatalf("Load configuration error: %s\n", err)
	}

	MysqlCfg = conf.MysqlCfg
	BarkCfg = conf.BarkCfg

	configLog, _ := json.MarshalIndent(conf, "", "  ")
	log.Printf("[INFO] Configuration load success: \n%s\n", configLog)
}
