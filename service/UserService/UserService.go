package userservice

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/constant"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/user"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	userroleenum "github.com/xiaoxiongmao5/xoj/xoj-backend/model/enums/UserRoleEnum"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/vo"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mydb"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

// 用户注册
func UserRegister(ctx *context.Context, userAccount, userPassword, checkUserPassword string) (num int64) {
	// 校验
	if utils.IsAnyBlank(userAccount, userPassword) {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "参数为空")
		return
	}
	if len(userAccount) < 4 || len(userAccount) > 16 {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "用户账号长度不合格要求")
		return
	}
	if len(userPassword) < 6 || len(userPassword) > 16 {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "用户密码长度不合格要求")
		return
	}
	// 密码和校验密码相同
	if !utils.CheckSame[string]("密码和校验密码相同", userPassword, checkUserPassword) {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "两次输入的密码不一致")
		return
	}

	userObj := entity.User{}
	if err := mydb.O.QueryTable(new(entity.User)).Filter("userAccount", userAccount).One(&userObj); err == nil || err == orm.ErrMultiRows {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "账号重复")
		return
	}

	// 加密
	hashPassword, err := utils.HashPasswordByBcrypt(userPassword)
	if err != nil {
		myresq.Abort(ctx, myresq.OPERATION_ERROR, "注册失败，加密错误")
		return
	}

	// 插入数据
	userObj.UserAccount = userAccount
	userObj.UserPassword = hashPassword
	userObj.UserName = userAccount
	num, err = Save(&userObj)
	if err != nil {
		myresq.Abort(ctx, myresq.OPERATION_ERROR, "注册失败，数据库错误")
		return
	}

	return
}

// 用户登录
func UserLogin(ctx *context.Context, userAccount, userPassword string) (loginUserVO vo.LoginUserVO) {
	// 校验
	if utils.IsAnyBlank(userAccount, userPassword) {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "参数为空")
		return
	}
	if len(userAccount) < 4 || len(userAccount) > 16 {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "账号错误")
		return
	}
	if len(userPassword) < 6 || len(userPassword) > 16 {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "密码错误")
		return
	}

	// 查询用户是否存在
	userObj := entity.User{}
	if err := mydb.O.QueryTable(new(entity.User)).Filter("userAccount", userAccount).One(&userObj); err != nil {
		if err == orm.ErrMultiRows {
			mylog.Log.Errorf("user 表中存在 userAccount[%s],userPassword[%s] 的多条记录, qs.One err=[%v]", userAccount, userPassword, err.Error())
			myresq.Abort(ctx, myresq.OPERATION_ERROR, "用户不存在或密码错误")
			return
		}
		if err == orm.ErrNoRows {
			mylog.Log.Errorf("user login failed, userAccount[%s] cannot match userPassword[%s]", userAccount, userPassword, err.Error())
			myresq.Abort(ctx, myresq.PARAMS_ERROR, "用户不存在或密码错误")
			return
		}
	}

	// 校验密码是否正确
	if err := utils.CheckHashPasswordByBcrypt(userObj.UserPassword, userPassword); err != nil {
		myresq.Abort(ctx, myresq.OPERATION_ERROR, "登录失败")
		return
	}

	// 记录用户的登录态
	sess, _ := mydb.GlobalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	defer sess.SessionRelease(ctx.Request.Context(), ctx.ResponseWriter)
	sess.Set(ctx.Request.Context(), "USER_LOGIN_STATE", userObj)

	return GetLoginUserVO(&userObj)
}

// 用户登录（微信开放平台）
func UserLoginByMpOpen(wxOAuth2UserInfo interface{}) {

}

// 获取当前登录用户
func GetLoginUser(ctx *context.Context) *entity.User {
	// 先判断是否已登录
	// SessionStart 根据当前请求返回 session 对象
	sess, _ := mydb.GlobalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	defer sess.SessionRelease(ctx.Request.Context(), ctx.ResponseWriter)
	userObj := sess.Get(ctx.Request.Context(), "USER_LOGIN_STATE")
	if userObj == nil {
		myresq.Abort(ctx, myresq.NOT_LOGIN_ERROR, "")
		return nil
	}
	currentUser, ok := userObj.(entity.User)
	if !ok || currentUser.ID <= 0 {
		myresq.Abort(ctx, myresq.NOT_LOGIN_ERROR, "")
		return nil
	}
	// 从数据库查询（追求性能的话可以注释，直接走缓存）
	loginUser, err := GetById(currentUser.ID)
	if err != nil {
		myresq.Abort(ctx, myresq.NOT_LOGIN_ERROR, "")
		return nil
	}
	return loginUser
}

// 获取当前登录用户（允许未登录）
//
//	@param ctx
func GetLoginUserPermitNull(ctx *context.Context) {

}

// 是否为管理员
//
//	@param user
//	@return bool
func IsAdmin(user *entity.User) bool {
	return utils.CheckSame[string]("是否为管理员", user.UserRole, userroleenum.ADMIN.GetValue())
}

// 用户注销
func UserLogout(ctx *context.Context) bool {
	sess, _ := mydb.GlobalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	defer sess.SessionRelease(ctx.Request.Context(), ctx.ResponseWriter)
	userObj := sess.Get(ctx.Request.Context(), "USER_LOGIN_STATE")
	if userObj == nil {
		myresq.Abort(ctx, myresq.NOT_LOGIN_ERROR, "")
		return true
	}
	currentUser, ok := userObj.(entity.User)
	if !ok || currentUser.ID <= 0 {
		myresq.Abort(ctx, myresq.NOT_LOGIN_ERROR, "未登录")
		return true
	}
	// 移除登录态
	err := sess.Delete(ctx.Request.Context(), "USER_LOGIN_STATE")
	if err != nil {
		myresq.Abort(ctx, myresq.OPERATION_ERROR, "用户注销失败")
		return false
	}
	return true
}

// 获取脱敏的已登录用户信息
func GetLoginUserVO(userObj *entity.User) vo.LoginUserVO {
	loginUserVO := vo.LoginUserVO{}
	utils.CopyStructFields(*userObj, &loginUserVO)
	return loginUserVO
}

// 获取脱敏的用户信息
func GetUserVO(userObj *entity.User) vo.UserVO {
	userVO := vo.UserVO{}
	utils.CopyStructFields(*userObj, &userVO)
	return userVO
}

// 获取脱敏的用户信息列表
//
//	@param list
//	@return respdata
func ListUserVO(list []*entity.User) (respdata []vo.UserVO) {
	if utils.IsEmpty(list) {
		return
	}
	for _, one := range list {
		respdata = append(respdata, GetUserVO(one))
	}
	return
}

// 获取查询条件
//
//	@param qs
//	@param queryRequest
//	@return orm.QuerySeter
func GetQuerySeter(qs orm.QuerySeter, queryRequest user.UserQueryRequest) orm.QuerySeter {
	id := queryRequest.ID
	unionId := queryRequest.UnionId
	mpOpenId := queryRequest.MpOpenId
	userName := queryRequest.UserName
	userProfile := queryRequest.UserProfile
	userRole := queryRequest.UserRole
	sortField := queryRequest.PageRequest.SortField
	sortOrder := queryRequest.PageRequest.SortOrder

	if id != 0 {
		qs = qs.Filter("id", id)
	}
	if utils.IsNotBlank(unionId) {
		qs = qs.Filter("unionId", unionId)
	}
	if utils.IsNotBlank(mpOpenId) {
		qs = qs.Filter("mpOpenId", mpOpenId)
	}
	if utils.IsNotBlank(userRole) {
		qs = qs.Filter("userRole", userRole)
	}
	if utils.IsNotBlank(userProfile) {
		qs = qs.Filter("userProfile__icontains", userProfile)
	}
	if utils.IsNotBlank(userName) {
		qs = qs.Filter("userName__icontains", userName)
	}

	if utils.IsNotBlank(sortField) {
		order := sortField
		if utils.CheckSame[string]("检查排序是否一样", sortOrder, constant.SORT_ORDER_DESC) {
			order = "-" + order
		}
		qs = qs.OrderBy(order)
	}
	qs = qs.Filter("isDelete", 0)
	return qs
}

func ListByIds(ids []int64) ([]*entity.User, error) {
	qs := mydb.O.QueryTable(new(entity.User))
	qs = qs.Filter("id__in", ids).Filter("isDelete", 0)
	var users []*entity.User
	_, err := qs.All(&users)
	if err != nil {
		mylog.Log.Errorf("User ListByIds qs.All error: %v", err.Error())
		return users, err
	}
	return users, nil
}

func GetById(id int64) (*entity.User, error) {
	var userObj entity.User
	err := mydb.O.QueryTable(new(entity.User)).Filter("id", id).Filter("isDelete", 0).One(&userObj)
	if err == orm.ErrMultiRows {
		mylog.Log.Errorf("User 表中存在 id=[%d] 的多条记录, qs.One err=[%v]", id, err.Error())
		return nil, err
	}
	if err == orm.ErrNoRows {
		mylog.Log.Errorf("User 表没有找到 id=[%d] 的记录, qs.One err=[%v]", id, err.Error())
		return nil, err
	}
	return &userObj, nil
}

func Save(userObj *entity.User) (int64, error) {
	num, err := mydb.O.Insert(userObj)
	if err != nil {
		return -1, err
	}
	return num, nil
}

func UpdateById(userObj *entity.User) error {
	num, err := mydb.O.Update(userObj)
	if err != nil {
		return err
	}
	if num == 0 {
		return errors.New("无更新影响条目")
	}
	return nil
}

func RemoveById(id int64) error {
	userObj, err := GetById(id)
	if err != nil {
		return nil
	}
	userObj.IsDelete = 1
	num, err := mydb.O.Update(userObj)
	if err != nil {
		return err
	}
	if num == 0 {
		return nil
	}
	return nil
}
