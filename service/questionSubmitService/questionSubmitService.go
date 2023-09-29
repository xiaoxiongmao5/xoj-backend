/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-29 09:20:16
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-29 20:37:03
 * @FilePath: /xoj-backend/service/QuestionSubmitService/QuestionSubmitService.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package questionsubmitservice

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/questionsubmit"
)

// 题目提交
func DoQuestionSubmit(ctx *context.Context) {
	// 校验编程语言是否合法

	// 判断实体是否存在，根据类别获取实体

	// 是否已提交题目

	// 每个用户串行提交题目

	// 设置初始状态

	// 执行判题服务
}

// 获取查询条件
func GetQuerySeter(qs orm.QuerySeter, queryRequest questionsubmit.QuestionSubmitQueryRequest) orm.QuerySeter {
	return qs
}

// 获取题目封装
func GetQuestionSubmitVO() {
	// 脱敏：仅本人和管理员能看见自己（提交 userId 和登录用户 id 不同）提交的代码

	// 处理脱敏
}

// 分页获取题目封装
func GetQuestionSubmitVOPage() {

}
