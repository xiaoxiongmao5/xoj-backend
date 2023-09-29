/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:42:42
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-29 18:16:59
 * @FilePath: /xoj-backend/model/entity/User.go
 * @Description: 实体类，表字段
 */
package entity

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
)

func init() {
	mylog.Log.Info("init begin: entity.User")

	// 注册模型
	orm.RegisterModel(new(User))

	mylog.Log.Info("init end  : entity.User")
}

// 用户
type User struct {
	ID           int64     `json:"id" orm:"column(id);auto;description(id)"`
	UserAccount  string    `json:"userAccount" orm:"column(userAccount);unique;description(账号)"`
	UserPassword string    `json:"userPassword" orm:"column(userPassword);size(512);description(密码)"`
	UnionId      string    `json:"unionId" orm:"column(unionId);null;index;description(微信开放平台id)"`
	MpOpenId     string    `json:"mpOpenId" orm:"column(mpOpenId);null;description(公众号openId)"`
	UserName     string    `json:"userName" orm:"column(userName);null;description(用户昵称)"`
	UserAvatar   string    `json:"userAvatar" orm:"column(userAvatar);size(1024);null;description(用户头像)"`
	UserProfile  string    `json:"userProfile" orm:"column(userProfile);size(512);null;description(用户简介)"`
	Gender       int32     `json:"gender" orm:"column(gender);null;description(性别)"`
	UserRole     string    `json:"userRole" orm:"column(userRole);default(user);description(用户角色：user/admin/ban)"`
	CreateTime   time.Time `json:"createTime" orm:"column(createTime);auto_now_add;type(datetime);description(创建时间)"`
	UpdateTime   time.Time `json:"updateTime" orm:"column(updateTime);auto_now;type(datetime);description(更新时间)"`
	IsDelete     int32     `json:"isDelete" orm:"column(isDelete);default(0);description(是否删除)"`
}

// 设置引擎为 INNODB
func (this *User) TableEngine() string {
	return "INNODB"
}

// 自定义表名
func (this *User) TableName() string {
	return "user"
}
