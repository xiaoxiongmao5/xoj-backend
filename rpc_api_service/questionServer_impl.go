/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-08 18:44:13
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-10 15:30:00
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

func (s *QuestionServerImpl) GetById(ctx context.Context, in *rpc_api.QuestionGetByIdReq) (*rpc_api.RpcQuestionObj, error) {
	mylog.Log.Infof("[RPC] GetById: QuestionId = %d", in.QuestionId)
	quesionObj, err := questionservice.GetById(in.QuestionId)
	if err != nil {
		return nil, err
	}
	return Convert2QuestionGetByIdResp(quesionObj), nil
}

func Convert2QuestionGetByIdResp(info *entity.Question) *rpc_api.RpcQuestionObj {
	var res rpc_api.RpcQuestionObj
	utils.CopyStructFields(*info, &res)

	res.CreateTime = ConvertTimeToTimestamp(info.CreateTime)
	res.UpdateTime = ConvertTimeToTimestamp(info.UpdateTime)
	return &res
}

func (s *QuestionServerImpl) Add1AcceptedNum(ctx context.Context, in *rpc_api.RpcQuestionObj) (*rpc_api.CommonUpdateByIdResp, error) {
	mylog.Log.Infof("[RPC] Add1AcceptedNum: QuestionId = %d, Title=%v, AcceptedNum=%v", in.Id, in.Title, in.AcceptedNum)

	var questionObj entity.Question
	utils.CopyStructFields(*in, &questionObj)

	// 更新题目通过数+1
	questionObj.AcceptedNum++

	err := questionservice.UpdateById(&questionObj, "acceptedNum")
	if err != nil {
		return &rpc_api.CommonUpdateByIdResp{Result: false}, err
	}
	return &rpc_api.CommonUpdateByIdResp{Result: true}, nil
}
