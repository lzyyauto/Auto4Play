package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
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
)

type configuration struct {
	// Env          Environment             `json:"env"`
	// ServerCfg    *ServerConfiguration    `json:"server"`
	// RedisCfg     *RedisConfiguration     `json:"redis"`
	MysqlCfg *MysqlConfiguration `json:"mysql"`
	// AlipayCfg    *AlipayConfiguration    `json:"alipay"`
	// WechatPayCfg *WechatPayConfiguration `json:"wechatpay"`
	// AccountCfg   *AccountConfiguration   `json:"account"`
}

type MysqlConfiguration struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	var configBytes []byte
	var err error
	configENV := os.Getenv("CONFIG")
	if configENV == "" {
		_, f, _, _ := runtime.Caller(0)
		cfgFile := fmt.Sprintf("%s/config.json", filepath.Dir(f))
		configBytes, err = ioutil.ReadFile(cfgFile)
		if err != nil {
			log.Fatalf("Load configuration failed: %s\n", err)
		}
	} else {
		configBytes = []byte(configENV)
	}
	var conf configuration
	if err = json.Unmarshal(configBytes, &conf); err != nil {
		log.Fatalf("Load configuration error: %s\n", err)
	}

	MysqlCfg = conf.MysqlCfg

	configLog, _ := json.MarshalIndent(conf, "", "  ")
	log.Printf("[INFO] Configuration load success: \n%s\n", configLog)
}
