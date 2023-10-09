/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-08 18:44:13
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-08 19:26:55
 */
package rpcapiservice

import (
	"context"

	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/rpc_api"
	questionservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/questionService"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

type QuestionServerImpl struct {
	rpc_api.UnimplementedQuestionServer
}

func (s *QuestionServerImpl) GetById(ctx context.Context, in *rpc_api.QuestionGetByIdReq) (*rpc_api.QuestionGetByIdResp, error) {
	mylog.Log.Infof("[RPC] GetById: QuestionId = %d", in.QuestionId)
	quesionObj, err := questionservice.GetById(in.QuestionId)
	if err != nil {
		return nil, err
	}
	return Convert2QuestionGetByIdResp(quesionObj), nil
}

func Convert2QuestionGetByIdResp(info *entity.Question) *rpc_api.QuestionGetByIdResp {
	var res rpc_api.QuestionGetByIdResp
	utils.CopyStructFields(*info, &res)

	res.CreateTime = ConvertTimeToTimestamp(info.CreateTime)
	res.UpdateTime = ConvertTimeToTimestamp(info.UpdateTime)
	return &res
}
