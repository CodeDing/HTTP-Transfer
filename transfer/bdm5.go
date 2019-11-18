package transfer

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/CodeDing/http-transfer/config"
	"github.com/CodeDing/http-transfer/httplib"
	api "github.com/CodeDing/http-transfer/proto/bdm5"
	"github.com/CodeDing/http-transfer/proto/me"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

var BDM_VERSION = "v5.6.0"

var numbers string = "0123456789"

const BDM5 = "bdm5"

const (
	BDM_API_VERSION_MAJOR uint32 = 5
	BDM_API_VERSION_MINOR uint32 = 6
	BDM_API_VERSION_MICRO uint32 = 0
)

//默认1：1 手机（含 iTouch），2 平板，3 电视
var Bdm5DeviceTypeMap = map[int]api.Device_DeviceType{
	1: api.Device_PHONE,
	2: api.Device_TABLET,
	3: api.Device_TABLET,
}

//系统 android,ios
var Bdm5OSTypeMap = map[string]api.Device_OsType{
	"android": api.Device_ANDROID,
	"ios":     api.Device_IOS,
}

//默认1：0 未知, 1 WIFI，2 2G，3 3G，4 4G
var Bdm5ConnectionTypeMap = map[int]api.Network_ConnectionType{
	0: api.Network_CONNECTION_UNKNOWN,
	1: api.Network_WIFI,
	2: api.Network_CELL_2G,
	3: api.Network_CELL_3G,
	4: api.Network_CELL_4G,
}

//默认1：0 未知,1 移动，2 联通，3 电信
var Bdm5OperatorTypeMap = map[int]api.Network_OperatorType{
	0: api.Network_UNKNOWN_OPERATOR,
	1: api.Network_CHINA_MOBILE,
	2: api.Network_CHINA_UNICOM,
	3: api.Network_CHINA_TELECOM,
}

var bdmApi string

var bdm5 *Bdm5

func init() {
	bdmApi = config.String(ADAPTER, BDM5)
	if bdmApi == "" {
		fmt.Printf("[%s] request url empty\n", strings.ToUpper(BDM5))
		os.Exit(-1)
	}
	fmt.Printf("[%s] url => %s\n", strings.ToUpper(BDM5), bdmApi)
	Register(BDM5, bdm5)
}

//百度联盟v5.6.0版本
type Bdm5 struct {
}

func (m *Bdm5) MeRequest2AdapterRequest(req *me.MeRequest) (*httplib.BeegoHTTPRequest, error) {
	if req.Selector == nil || req.Slot == nil || req.Device == nil {
		return nil, fmt.Errorf("[%s] request params is illegal!", strings.ToUpper(BDM5))
	}
	var bdm5Req = &api.MobadsRequest{}
	bdm5Req.RequestId = proto.String(fmt.Sprintf("%x", md5.Sum([]byte(req.Id))))
	apiVersion := &api.Version{}
	apiVersion.Major = proto.Uint32(BDM_API_VERSION_MAJOR)
	apiVersion.Minor = proto.Uint32(BDM_API_VERSION_MINOR)
	apiVersion.Micro = proto.Uint32(BDM_API_VERSION_MICRO)
	bdm5Req.ApiVersion = apiVersion

	//APP
	app := &api.App{}
	app.AppId = proto.String(req.Selector.AppId)
	app.AppPackage = proto.String(req.Selector.PackageName)

	appVersion, err := ParseVersion(req.Selector.PackageVersion)
	if err != nil {
		appVersion = &api.Version{}
		appVersion.Major = proto.Uint32(1)
	}
	app.AppVersion = appVersion
	bdm5Req.App = app

	device := &api.Device{}
	if deviceType, ok := Bdm5DeviceTypeMap[req.Device.DeviceType]; ok {
		device.DeviceType = &deviceType
	}
	if osType, ok := Bdm5OSTypeMap[req.Device.Os]; ok {
		device.OsType = &osType
	}
	osVersion, err := ParseVersion(req.Device.OsVersion)
	if err != nil {
		osVersion = &api.Version{}
		osVersion.Major = proto.Uint32(5)
	}
	device.OsVersion = osVersion

	device.Vendor = []byte(req.Device.Vendor)
	device.Model = []byte(req.Device.Model)

	device.ScreenSize = &api.Size{}
	device.ScreenSize.Height = proto.Uint32(uint32(req.Slot.Height))
	device.ScreenSize.Width = proto.Uint32(uint32(req.Slot.Width))

	device.Udid = &api.UdId{}
	device.Udid.Imei = proto.String(req.Device.Imei)
	device.Udid.Mac = proto.String(req.Device.Mac)
	device.Udid.AndroidId = proto.String(req.Device.AndroidId)
	device.Udid.Idfa = proto.String(req.Device.Idfa)
	device.Udid.AndroididMd5 = proto.String(fmt.Sprintf("%x", md5.Sum([]byte(req.Device.AndroidId))))
	device.Udid.IdfaMd5 = proto.String(req.Device.IdfaMd5)
	device.Udid.ImeiMd5 = proto.String(req.Device.ImeiMd5)
	bdm5Req.Device = device

	//network
	network := &api.Network{}
	network.Ipv4 = proto.String(req.Device.Ip)
	if networkType, ok := Bdm5ConnectionTypeMap[req.Device.Network]; ok {
		network.ConnectionType = &networkType
	}
	if orientationType, ok := Bdm5OperatorTypeMap[req.Device.Orientation]; ok {
		network.OperatorType = &orientationType
	}
	bdm5Req.Network = network

	slot := &api.AdSlot{}
	slot.AdslotId = proto.String(req.Selector.SlotId)
	slot.AdslotSize = &api.Size{}
	slot.AdslotSize.Width = proto.Uint32(uint32(req.Slot.Width))
	slot.AdslotSize.Height = proto.Uint32(uint32(req.Slot.Height))

	bdm5Req.Adslot = slot
	bdm5Req.IsDebug = proto.Bool(config.BoolDefault("bdm5", "isdebug", false))
	datas, err := proto.Marshal(bdm5Req)
	if err != nil {
		return nil, err
	}
	//TODO
	js, _ := json.Marshal(bdm5Req)
	glog.V(5).Infof("[%s] request(json) => %s\n", strings.ToUpper(BDM5), string(js))
	ret := httplib.Post(bdmApi)
	ret.Header("Content-Type", "application/octet-stream")
	ret.Header("Content-Length", fmt.Sprintf("%d", len(datas)))
	ret.Body(datas)
	return ret, nil
}

func (m *Bdm5) AdapterResponse2MeResponse(meReq *me.MeRequest, response *http.Response) (*me.MeResponse, error) {
	bdmResp := &api.MobadsResponse{}
	meResp := &me.MeResponse{}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to read body!", strings.ToUpper(BDM5))
	}

	defer response.Body.Close()
	err = proto.Unmarshal(body, bdmResp)
	if err != nil {
		meResp.Code = 999 //result异常
		meResp.Msg = ""
		return meResp, err
	}
	//TODO
	js, _ := json.Marshal(bdmResp)
	glog.V(5).Infof("[%s] response(json) => %s", strings.ToUpper(BDM5), string(js))

	if bdmResp.GetErrorCode() != 0 {
		meResp.Code = int(bdmResp.GetErrorCode())
		meResp.Msg = ""
		return meResp, nil
	}
	meResp.Id = meReq.Id
	meResp.Width = meReq.Slot.Width
	meResp.Height = meReq.Slot.Height
	meResp.AdId = meReq.Selector.AdId
	meResp.SlotId = meReq.Slot.Id
	meResp.Union = BDM5

	//广告信息
	for _, v := range bdmResp.GetAds() {
		if len(v.GetMetaGroup()) > 0 && v.GetMetaGroup()[0] != nil && meResp.Content == nil {
			meta := v.GetMetaGroup()[0]
			content := &me.MeResponseContent{}
			app := &me.MeResponseContentApp{}
			app.PackageName = meta.GetAppPackage()
			app.Size = int(meta.GetAppSize())
			if len(meta.GetIconSrc()) > 0 {
				app.IconUrl = meta.GetIconSrc()[0]
			}
			app.Name = meta.GetApkName()
			content.App = app

			switch meta.GetCreativeType() {
			case api.MaterialMeta_TEXT:
				content.ContentType = TEXT
			case api.MaterialMeta_IMAGE:
				content.ContentType = IMAGE
			case api.MaterialMeta_TEXT_ICON:
				content.ContentType = TEXT_ICON
			case api.MaterialMeta_VIDEO:
				content.ContentType = VIDEO
			case api.MaterialMeta_NO_TYPE:
				//百度返回no_type时，有时候他带有所有字段，可以任意指定类型
				content.ContentType = TEXT_ICON
				//但有时no_type里面并没有title和description，指定为text就无用了
				if len(meta.GetTitle()) == 0 && len(meta.GetDescription()) == 0 {
					content.ContentType = IMAGE
				}
			default:
				content.ContentType = TEXT
			}
			//0 下载 1 打开网页
			switch meta.GetInteractionType() {
			case api.MaterialMeta_DOWNLOAD:
				content.Action = 0
			case api.MaterialMeta_DEEPLINK:
				content.Action = 0
				content.DpClk = []string{string(meta.GetDeeplinkUrl())}
				content.InstUrls = []string{(string(meta.GetFallbackUrl()))}

			default:
				content.Action = 1
			}
			content.Url = meta.GetClickUrl()
			content.CUrl = meta.GetClickUrl()
			content.Imp = meta.GetWinNoticeUrl()
			content.Clk = []string{meta.GetClickUrl()}
			content.Title = string(meta.GetTitle())
			if len(meta.GetDescription()) > 0 {
				content.Desc = string(meta.GetDescription()[0])
			} else {
				content.Desc = content.Title
			}

			for idx, v := range meta.GetImageSrc() {
				if idx == 0 {
					content.Url = v
				} else {
					if content.ExtUrls == nil {
						content.ExtUrls = []string{}
					}
					content.ExtUrls = append(content.ExtUrls, v)
				}
			}
			content.IconUrl = app.IconUrl
			meResp.Content = content
		}
	}

	return meResp, nil
}

func ParseVersion(s string) (*api.Version, error) {
	ver := &api.Version{}
	ver.Major = proto.Uint32(1)
	if len(s) == 0 {
		return ver, fmt.Errorf("Version string empty")
	}

	// Split into major.minor.(patch+pr+meta)
	parts := strings.SplitN(s, ".", 3)
	if len(parts) < 2 {
		return ver, fmt.Errorf("No Major.Minor.Patch elements found")
	}

	// Major
	if !containsOnly(parts[0], numbers) {
		return ver, fmt.Errorf("Invalid character(s) found in major number %q", parts[0])
	}
	if hasLeadingZeroes(parts[0]) {
		return ver, fmt.Errorf("Major number must not contain leading zeroes %q", parts[0])
	}
	major, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		return ver, err
	}

	// Minor
	if !containsOnly(parts[1], numbers) {
		return ver, fmt.Errorf("Invalid character(s) found in minor number %q", parts[1])
	}
	if hasLeadingZeroes(parts[1]) {
		return ver, fmt.Errorf("Minor number must not contain leading zeroes %q", parts[1])
	}
	minor, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return ver, err
	}
	if len(parts) > 2 {

		// Micro
		if !containsOnly(parts[2], numbers) {
			return ver, fmt.Errorf("Invalid character(s) found in minor number %q", parts[1])
		}
		if hasLeadingZeroes(parts[2]) {
			return ver, fmt.Errorf("Minor number must not contain leading zeroes %q", parts[1])
		}
		micro, err := strconv.ParseUint(parts[2], 10, 32)
		if err != nil {
			return ver, err
		}
		ver.Micro = proto.Uint32(uint32(micro))
	}
	ver.Major = proto.Uint32(uint32(major))
	ver.Minor = proto.Uint32(uint32(minor))
	return ver, nil
}

func containsOnly(s string, set string) bool {
	return strings.IndexFunc(s, func(r rune) bool {
		return !strings.ContainsRune(set, r)
	}) == -1
}

func hasLeadingZeroes(s string) bool {
	return len(s) > 1 && s[0] == '0'
}
