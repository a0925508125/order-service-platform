package errcode

const Success = 200000

const (
	// 公用錯誤代碼
	CommonParamError        = 200001 + iota //參數錯誤
	CommonUnKnowError                       //未知錯誤
	CommonGRPCError                         //GRPC錯誤
	CommonConvertError                      //資料格式錯誤
	CommonTokenValidError                   //token驗證錯誤
	CommonDataNotFoundError                 //查無資料
	DefaultFaultError       = 400000 + iota
)
