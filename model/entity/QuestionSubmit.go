/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:31:11
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-29 00:57:54
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
	mylog.Log.Info("init begin: entity.QuestionSubmit")

	orm.RegisterModel(new(QuestionSubmit))

	mylog.Log.Info("init end  : entity.QuestionSubmit")
}

// 题目提交
type QuestionSubmit struct {
	ID         int64     `json:"id" orm:"column(id);auto;description(id)"`
	Language   string    `json:"language" orm:"column(language);size(128);description(编程语言)"`
	Code       string    `json:"code" orm:"column(code);type(text);description(用户代码)"`
	JudgeInfo  string    `json:"judgeInfo" orm:"column(judgeInfo);type(text);null;description(判题信息-json 对象)"`
	Status     int32     `json:"status" orm:"column(status);default(0);description(判题状态:0-待判题、1-判题中、2-成功、3-失败）)"`
	QuestionId int64     `json:"questionId" orm:"column(questionId);index;description(题目 id)"`
	Userid     int64     `json:"userId" orm:"column(userId);index;description(创建用户 id)"`
	CreateTime time.Time `json:"createTime" orm:"column(createTime);auto_now_add;type(datetime);description(创建时间)"`
	UpdateTime time.Time `json:"updateTime" orm:"column(updateTime);auto_now;type(datetime);description(更新时间)"`
	IsDelete   int32     `json:"isDelete" orm:"column(isDelete);default(0);description(是否删除)"`
}

func (this *QuestionSubmit) TableEngine() string {
	return "INNODB"
}

// 自定义表名
func (this *QuestionSubmit) TableName() string {
	return "question_submit"
}
