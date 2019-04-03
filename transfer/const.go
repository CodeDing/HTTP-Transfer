package transfer

const (
	ADAPTER = "adapter"
)

const (
	ANDROID = "android"
	IOS     = "ios"
)

const (
	IMEI = "imei"
	IDFA = "idfa"
)

//必填，联网方式，默认1：0 未知, 1 WIFI，2 2G，3 3G，4 4G
const (
	WIFI = 1
	N2G  = 2
	N3G  = 3
	N4G  = 4
)

//必填，运营商，默认1：0 未知,1 移动，2 联通，3 电信
//China Union Communication Corporation ---CUCC联通
//China Mobile Communications Corporation---CMCC移动
//CHINATELECOM -- CTCC电信
//中国电信：China Telecom
//中国移动：China Mobile
//中国联通：China Unicom

const (
	CMCC = 1
	CUCC = 2
	CTCC = 3
)

const (
	IMAGE     = "image"
	TEXT_ICON = "text-icon"
	TEXT = "text"
	VIDEO     = "video"
)

const (
	SURFING_WEB = 1
	DOWNLOAD    = 0
)

//API http err code

const (
	UNMARSHAL_RESP_ERR = 999
	NO_AD              = 996
	ILLEGAL_PARAMS     = 5001
)

const (
	AD_IS_NIL  = "广告数组异常!"
	AD_SUCCESS = "success"
)