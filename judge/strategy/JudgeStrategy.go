/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-02 14:23:36
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-02 15:04:16
 * @FilePath: /xoj-backend/judge/strategy/JudgeStrategy.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package strategy

import "github.com/xiaoxiongmao5/xoj/xoj-backend/judge/codesandbox/model"

// 判题策略
type JudgeStrategy interface {
	DoJudge(judgeContext JudgeContext) model.JudgeInfo
}
