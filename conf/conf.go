package conf

import (
	"flag"
	"fmt"
	yconf "github.com/felix-xqs/conf"
	"github.com/felix-xqs/golog"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path"
	"path/filepath"
)

type Conf struct{
	App *struct {
		Name string  `mapstructure:"name"`
		Port int `mapstructure:"port"`
	} `mapstructure:"app"`
	Redis map[string]RedisConf `mapstructure:"redis"`
	Mongo map[string]MongoConf `mapstructure:"mongo"`

	Logger *golog.Config `mapstructure:"logger"`
	Services *struct{
		PayCenter string `mapstructure:"payCenter"`
		PushCenter string `mapstructure:"payCenter"`
	} `mapstructure:"services"`
}

var (
	Gin *gin.Engine
	C   Conf
	RDS YRedisMap
	DBS  YMongoMap
	Env string
)
func initConfig(env string){
	configDir := getConfDir()
	configName := fmt.Sprintf("%s.conf.yaml", env)
	configPath := path.Join(configDir, configName)
	Env = env
	C = Conf{}
	err := yconf.LoadConfig(configPath, &C)
	if err != nil {
		log.Fatal(err)
	}
}
func getConfDir() string {
	dir := "conf"
	for i := 0; i < 3; i++ {
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			break
		}
		dir = filepath.Join("..", dir)
	}

	return dir
}
func init() {
	var env string
	flag.StringVar(&env, "env", "dev", "set env")
	var test bool
	flag.BoolVar(&test, "test", false, "set test case flag")
	flag.Parse()

	initConfig(env)
	gin.SetMode(gin.ReleaseMode)

	golog.Init(C.Logger)

	DBS = GetMongoMap(C.Mongo)
	RDS = GetRedisMap(C.Redis)
}
