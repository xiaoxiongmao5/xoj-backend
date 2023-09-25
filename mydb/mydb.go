package mydb

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
)

var DB *sql.DB

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
