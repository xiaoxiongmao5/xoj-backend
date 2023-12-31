/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-11 14:25:30
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-15 22:38:49
 * @FilePath: /xoj-backend/mydb/mydb.go
 * @Description: mydb
 */
package mydb

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/client/orm/filter/bean"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/config"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
)

var DB *sql.DB
var O orm.Ormer

func init() {
	mylog.Log.Info("init begin: mydb")

	// 注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 注册数据库连接
	orm.RegisterDataBase("default", "mysql", config.AppConfig.Database.SavePath)

	// 显示注册默认值的Filter
	builder := bean.NewDefaultValueFilterChainBuilder(nil, true, true)
	orm.AddGlobalFilterChain(builder.FilterChain)

	// 创建一个 Orm 实例对象，用于执行数据库操作。 NewOrm 的同时会执行 orm.BootStrap (整个 app 只执行一次)，用以验证模型之间的定义并缓存。（大多数情况下，应该尽量复用Orm 实例，因为本身Orm实例被设计为无状态的，一个数据库对应一个Orm实例）（ps: 但是在使用事务的时候，我们会返回TxOrm的实例，它本身是有状态的，一个事务对应一个TxOrm实例。在使用TxOrm时候，任何衍生查询都是在该事务内。）
	O = orm.NewOrm()

	// 执行数据库表结构同步操作。告诉 ORM 在默认数据库连接上执行同步操作，第二个参数 false 表示不强制删除已存在的表，第三个参数 true 表示打印同步操作的日志。这通常在应用程序启动时执行，以确保数据库表结构与 ORM 模型定义一致。
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		mylog.Log.Error("数据库表结构同步操作失败[orm.RunSyncdb] err=", err.Error())
	}

	mylog.Log.Info("init end  : mydb")
}

func Close(dbCli *sql.DB) error {
	if dbCli != nil {
		return dbCli.Close()
	}
	return nil
}

// 创建数据库连接池
// 使用：
// // 初始化数据库连接池
//
//	if mydb.DB, err = mydb.ConnectionPool(config.AppConfig.Database.SavePath, config.AppConfig.Database.MaxOpenConns); err != nil {
//		panic(err)
//	}
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
