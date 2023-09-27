/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:27:02
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-27 21:25:14
 * @FilePath: /xoj-backend/service/question.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package questionservice

import (
	"database/sql"

	"context"

	beecontext "github.com/beego/beego/v2/server/web/context"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/dbsq"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/vo"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mydb"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	userservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/userService"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

func ValidQuestion(ctx *beecontext.Context, question *entity.Question, add bool) {
	title := question.Title
	content := question.Content
	tags := question.Tags
	answer := question.Answer
	judgeCase := question.Judgecase
	judgeConfig := question.Judgeconfig
	// 创建时，参数不能为空
	if add {
		if utils.IsAnyBlank(title, content, tags) {
			myresq.Abort(ctx, myresq.PARAMS_ERROR, "")
			return
		}
	}
	// 有参数则校验
	if len(title) > 80 {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "标题过长")
		return
	}
	if len(content) > 8192 {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "内容过长")
		return
	}
	if len(answer) > 8192 {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "答案过长")
		return
	}
	if len(judgeCase) > 8192 {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "判题用例过长")
		return
	}
	if len(judgeConfig) > 8192 {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "判题配置过长")
		return
	}
}

func GetQuestionVO(original *entity.Question) vo.QuestionVO {
	questionVO := vo.Obj2Vo(original)
	// 关联查询用户信息
	loginUser := &dbsq.User{}
	if original.Userid > 0 {
		loginUser, _ = userservice.GetUserInfoById(original.Userid)
	}
	userVO := userservice.GetUserVO(loginUser)
	questionVO.UserVO = userVO
	return questionVO
}

func Save(dbparams *dbsq.AddQuestionParams) (sql.Result, error) {
	conn, err := mydb.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.AddQuestion(ctx, dbparams)
}

func GetById(id int64) (*dbsq.Question, error) {
	conn, err := mydb.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.GetQuestionById(ctx, id)
}

func UpdateById(dbparams *dbsq.UpdateQuestionParams) error {
	conn, err := mydb.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.UpdateQuestion(ctx, dbparams)
}

func RemoveById(id int64) error {
	conn, err := mydb.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.DeleteQuestion(ctx, id)
}
