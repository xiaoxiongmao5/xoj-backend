package mydb

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
)

var DB *sql.DB
var O orm.Ormer

func init() {
	mylog.Log.Info("init begin: mydb")

	// 注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 注册数据库连接
	orm.RegisterDataBase("default", "mysql", "root:@/xoj?charset=utf8&parseTime=true")

	// 创建一个 Orm 实例对象，用于执行数据库操作。 NewOrm 的同时会执行 orm.BootStrap (整个 app 只执行一次)，用以验证模型之间的定义并缓存。（大多数情况下，应该尽量复用Orm 实例，因为本身Orm实例被设计为无状态的，一个数据库对应一个Orm实例）（ps: 但是在使用事务的时候，我们会返回TxOrm的实例，它本身是有状态的，一个事务对应一个TxOrm实例。在使用TxOrm时候，任何衍生查询都是在该事务内。）
	O = orm.NewOrm()

	mylog.Log.Info("init end  : mydb")
}

// 创建数据库连接池
func ConnectionPool(savePath string, maxOpenConns int) (*sql.DB, error) {
	db, err := sql.Open("mysql", savePath)
	if err != nil {
		mylog.Log.Error("数据库连接失败, err=", err)
		return nil, err
	}
	// 设置最大连接池大小
	db.SetMaxOpenConns(maxOpenConns)
	mylog.Log.Infof("数据库连接成功,savePath=[%s],maxOpenConns=[%d]", savePath, maxOpenConns)
	return db, nil
}

// 从连接池中获取连接
func GetConn() (*sql.Conn, error) {
	if DB == nil {
		return nil, errors.New("DB database connection is nil")
	}
	ctx := context.Background()
	conn, err := DB.Conn(ctx)
	if err != nil {
		mylog.Log.Error("从连接池中获取连接失败, err=", err)
		return nil, err
	}
	return conn, nil
}

// 获取数据库连接，最多重试 maxRetries 次
func GetConnWithRetry(maxRetries int) (*sql.Conn, error) {
	for retry := 0; retry < maxRetries; retry++ {
		conn, err := GetConn()
		if err == nil {
			return conn, nil
		}
		// 如果连接失败，等待一段时间后重试
		time.Sleep(1 * time.Second)
	}

	return nil, errors.New("failed to obtain database connection after retries")
}
