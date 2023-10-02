/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-29 21:42:35
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-02 13:19:51
 * @FilePath: /xoj-backend/model/enums/QuestionSubmitStatusEnum.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package questionsubmitstatusenum

type QuestionSubmitStatusEnum int32

func (this QuestionSubmitStatusEnum) GetValue() int32 {
	return int32(this)
}

func (this QuestionSubmitStatusEnum) GetText() string {
	return QuestionSubmitStatusEnumName[this]
}

const (
	WAITING QuestionSubmitStatusEnum = 0
	RUNNING QuestionSubmitStatusEnum = 1
	SUCCEED QuestionSubmitStatusEnum = 2
	FAILED  QuestionSubmitStatusEnum = 3
)

var QuestionSubmitStatusEnumName = map[QuestionSubmitStatusEnum]string{
	WAITING: "等待中",
	RUNNING: "判题中",
	SUCCEED: "成功",
	FAILED:  "失败",
}
