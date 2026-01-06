package errcode

var CodeMapMessage = map[int32]string{
	Success: "操作成功",

	//CommonParamError:        "參數錯誤",
	//CommonUnKnowError:       "不知名錯誤",
	//CommonConvertError:      "資料格式錯誤",
	//CommonGRPCError:         "GRPC 錯誤",
	//CommonTokenValidError:   "token 驗證錯誤",
	//CommonDataNotFoundError: "查無資料",
	//CommonUploadError:       "上传档案错误",
	//CommonUploadDataError:   "上传档案内容错误",
	//
	////MemberInfoError:  "玩家資訊錯誤",
	//MemberIsGaming:   "玩家在遊戲中",
	//MemberIsBetting:  "玩家在下注中",
	//MemberTokenError: "玩家token錯誤",
	//
	//DefaultFaultError:     "預設錯誤",
	//BetOrderSyncFail:      "注單同步失敗",
	//CommonGameStatusError: "游戏关闭失败，请洽IT人员",
	//
	////驗證問題
	//AuthGroupNotFoundError:  "使用者无群组",
	//InsufficientPermissions: "使用者权限不足",
}
