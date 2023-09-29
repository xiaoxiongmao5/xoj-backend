/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:31:11
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-29 22:18:53
 * @FilePath: /xoj-backend/model/entity/Question.go
 * @Description: 实体类，表字段
 */
package entity

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
)

func init() {
	mylog.Log.Info("init begin: entity.Question")

	orm.RegisterModel(new(Question))

	mylog.Log.Info("init end  : entity.Question")
}

// 题目
type Question struct {
	ID          int64     `json:"id" orm:"column(id);auto;description(id)"`
	Title       string    `json:"title" orm:"column(title);size(512);null;description(标题)"`
	Content     string    `json:"content" orm:"column(content);type(text);null;description(内容)"`
	Tags        string    `json:"tags" orm:"column(tags);size(1024);null;description(标签列表-json 数组)"`
	Answer      string    `json:"answer" orm:"column(answer);type(text);null;description(题目答案)"`
	Submitnum   int32     `json:"submitNum" orm:"column(submitNum);default(0);description(题目提交数)"`
	Acceptednum int32     `json:"acceptedNum" orm:"column(acceptedNum);default(0);description(题目通过数)"`
	Judgecase   string    `json:"judgeCase" orm:"column(judgeCase);type(text);null;description(判题用例-json 数组)"`
	Judgeconfig string    `json:"judgeConfig" orm:"column(judgeConfig);type(text);null;description(判题配置-json 对象)"`
	Thumbnum    int32     `json:"thumbNum" orm:"column(thumbNum);default(0);description(点赞数)"`
	Favournum   int32     `json:"favourNum" orm:"column(favourNum);default(0);description(收藏数)"`
	Userid      int64     `json:"userId" orm:"column(userId);index;description(创建用户 id)"`
	CreateTime  time.Time `json:"createTime" orm:"column(createTime);auto_now_add;type(datetime);description(创建时间)"`
	UpdateTime  time.Time `json:"updateTime" orm:"column(updateTime);auto_now;type(datetime);description(更新时间)"`
	IsDelete    int32     `json:"isDelete" orm:"column(isDelete);default(0);description(是否删除)"`
}

func (this *Question) TableEngine() string {
	return "INNODB"
}

func (this *Question) TableName() string {
	return "question"
}
