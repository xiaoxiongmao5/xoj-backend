/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-02 13:25:20
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-02 14:56:23
 * @FilePath: /xoj-backend/judge/JudgeService.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package judge

import "github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"

// 判题服务
type JudgeService interface {
	DoJudge(questionsubmitId int64) *entity.QuestionSubmit
}
