package enums

type ResCodeCommon int
type ResCodeOfManage int
type ResCodeOfUser int
type ResCodeOfInterface int

const (
	// ResCodeCommon
	Success            ResCodeCommon = iota // 0
	ParameterError                          // 1
	AuthenticationFail                      // 2
	Unauthorized
	NotAdminRole
	GenerateRandomKeyFailed
	GenerateTokenFailed
	HandleSQLError
)

const (
	// ResCodeOfManage
	UpdateIPRateLimitConfigFailed ResCodeOfManage = 2000
)

// const (
// 	// ResCodeOfInterface
// 	InterfaceNotExist ResCodeOfInterface = 4000
// 	ListInterfaceFailed
// 	CreateInterfaceFailed
// 	UpdateInterfaceFailed
// 	DeleteInterfaceFailed
// 	OnlineInterfaceFailed
// 	OfflineInterfaceFailed
// 	InvokeInterfaceFailed
// 	UpdateInvokeLeftCountFailed
// )
