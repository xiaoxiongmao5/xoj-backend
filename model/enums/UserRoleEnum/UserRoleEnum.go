/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 15:24:36
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-02 13:19:55
 * @FilePath: /xoj-backend/model/enums/UserRoleEnum.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package userroleenum

type UserRoleEnum string

func (this UserRoleEnum) GetValue() string {
	return string(this)
}

func (this UserRoleEnum) GetText() string {
	return UserRoleEnumName[this]
}

const (
	USER  UserRoleEnum = "user"
	ADMIN UserRoleEnum = "admin"
	BAN   UserRoleEnum = "ban"
)

var UserRoleEnumName = map[UserRoleEnum]string{
	USER:  "用户",
	ADMIN: "管理员",
	BAN:   "被封号",
}
