/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-08 18:44:13
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-10 15:19:59
 */
package rpcapiservice

import (
	"context"

	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/rpc_api"
	questionsubmitservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/questionSubmitService"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

type QuestionSubmitServerImpl struct {
	rpc_api.UnimplementedQuestionSubmitServer
}

func (s *QuestionSubmitServerImpl) GetById(ctx context.Context, in *rpc_api.QuestionSubmitGetByIdReq) (*rpc_api.RpcQuestionSubmitObj, error) {
	mylog.Log.Infof("[RPC] GetById: QuestionSubmitId = %d", in.QuestionSubmitId)
	questionSubmitObj, err := questionsubmitservice.GetById(in.QuestionSubmitId)
	if err != nil {
		return nil, err
	}
	return Convert2QuestionSubmitUpdateByIdResp(questionSubmitObj), nil
}

func Convert2QuestionSubmitUpdateByIdResp(info *entity.QuestionSubmit) *rpc_api.RpcQuestionSubmitObj {
	var res rpc_api.RpcQuestionSubmitObj
	utils.CopyStructFields(*info, &res)

	res.CreateTime = ConvertTimeToTimestamp(info.CreateTime)
	res.UpdateTime = ConvertTimeToTimestamp(info.UpdateTime)
	return &res
}

func (s *QuestionSubmitServerImpl) UpdateById(ctx context.Context, in *rpc_api.RpcQuestionSubmitObj) (*rpc_api.CommonUpdateByIdResp, error) {
	mylog.Log.Infof("[RPC] UpdateById: QuestionSubmitId = %d, JudgeInfo=%v, Status=%v, QuestionId=%v, UserId=%v", in.Id, in.JudgeInfo, in.Status, in.QuestionId, in.UserId)

	questionSubmitObj := Convert2EntityQuestionSubmitObj(in)

	mylog.Log.Infof("Convert2EntityQuestionSubmitObj后: QuestionSubmitId = %d, JudgeInfo=%v, Status=%v, QuestionId=%v, UserId=%v", questionSubmitObj.Id, questionSubmitObj.JudgeInfo, questionSubmitObj.Status, questionSubmitObj.QuestionId, questionSubmitObj.UserId)

	err := questionsubmitservice.UpdateById(questionSubmitObj)
	if err != nil {
		return &rpc_api.CommonUpdateByIdResp{Result: false}, err
	}
	return &rpc_api.CommonUpdateByIdResp{Result: true}, nil
}

func Convert2EntityQuestionSubmitObj(info *rpc_api.RpcQuestionSubmitObj) *entity.QuestionSubmit {
	var res entity.QuestionSubmit
	utils.CopyStructFields(*info, &res)

	res.CreateTime = ConvertTimestampToTime(info.CreateTime)
	res.UpdateTime = ConvertTimestampToTime(info.UpdateTime)
	res.IsDelete = info.IsDelete
	return &res
}
