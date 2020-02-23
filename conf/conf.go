package conf

import (
	"fmt"
	"path/filepath"
)

type Conf struct {
	Database        string   `yaml:"database"`
	User            string   `yaml:"user"`
	Password        string   `yaml:"password"`
	RemoveUrl       string   `yaml:"removeUrl"`
	PhotoSavePath   string   `yaml:"photoSavePath"`
	EmailAddr       string   `yaml:"emailAddr"`
	ToEmailAddr     []string `yaml:"toEmailAddr"`
	Subject         string   `yaml:"subject"`
	RedisAddr       []string `yaml:"redisAddr"`
	RedisPassWord   string   `yaml:"redisPassword"`
	LoginStatusLong int      `yaml:"loginStatusLong"`
	CookietLong     int      `yaml:"cookietLong"`
	PemPath         string   `yaml:"pemPath"`
	SslPath         string   `yaml:"sslPath"`
}

const (
	C_ConfFilePath = "conf" + string(filepath.Separator) + "conf.yaml"
)

func (c *Conf) GetDataBaseUrl() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		c.User, c.Password, c.RemoveUrl, c.Database)
	return dsn
}
