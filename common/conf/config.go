package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
)

func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local",
		d.UserName,
		d.Password,
		d.Host,
		d.DBName,
		d.Charset,
		d.ParseTime,
	)
}

type JWTConfig struct {
	Secret  string
	Expires time.Duration
}

func InitConfig() {
	vp := viper.New()
	workDirectory, err := os.Getwd()
	if err != nil {
		klog.Fatal(err)
	}
	vp.AddConfigPath(workDirectory + "/conf")
	for workDirectory != "/" {
		vp.AddConfigPath(workDirectory + "/conf")
		workDirectory = filepath.Dir(workDirectory)
	}
	vp.SetConfigName("conf")
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil {
		klog.Fatal(err)
	}
	vp.UnmarshalKey("Server", &Server)
	vp.UnmarshalKey("Database", &Database)
	vp.UnmarshalKey("JWT", &JWT)
	JWT.Expires *= time.Hour
}
