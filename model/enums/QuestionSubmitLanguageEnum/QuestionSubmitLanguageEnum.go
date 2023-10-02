/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-29 21:42:35
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-02 13:20:00
 * @FilePath: /xoj-backend/model/enums/QuestionSubmitLanguageEnum.go
 */

package questionsubmitlanguageenum

type QuestionSubmitLanguageEnum string

func (this QuestionSubmitLanguageEnum) GetValue() string {
	return string(this)
}

func (this QuestionSubmitLanguageEnum) GetText() string {
	return QuestionSubmitLanguageEnumName[this]
}

const (
	JAVA      QuestionSubmitLanguageEnum = "java"
	CPLUSPLUS QuestionSubmitLanguageEnum = "cpp"
	GOLANG    QuestionSubmitLanguageEnum = "go"
)

var QuestionSubmitLanguageEnumName = map[QuestionSubmitLanguageEnum]string{
	JAVA:      "java",
	CPLUSPLUS: "cpp",
	GOLANG:    "go",
}

// 根据 value 获取枚举
func GetEnumByValue(value string) {

}
