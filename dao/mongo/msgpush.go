package mongo

import (
	"github.com/felix-xqs/go_example/conf"
	"github.com/felix-xqs/golog"
	"github.com/xqs-go-lib/gosdk/mygin"
)

//生成本库的client实例，想要查询就getsessionAndCollection
//var MsgPushMongoClient = &mongo.client{}
const (
	DBNameTest         = "mg_test_0"
	CollectionNameUser = "user"
)

var TestDao *mygin.Client

func InitTestDao() {
	var err error
	TestDao, err = mygin.NewClient(conf.DBS[DBNameTest], DBNameTest, CollectionNameUser)
	if err != nil {
		golog.Error("", err)
	}
}
