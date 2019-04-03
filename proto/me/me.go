package me

type MeRequest struct {
	Id       string             `json:"id"`
	Slot     *MeRequestSlot     `json:"slot"`
	Selector *MeRequestSelector `json:"selector"`
	Device   *MeRequestDevice   `json:"device"`
	App      *MeRequestApp      `json:"app"`
	User     *MeRequestUser     `json:"user"`
	Region   *MeRequestRegion   `json:"region"`
	Ext      *MeRequestExt      `json:"ext"`
}

type MeRequestSlot struct {
	Id              string `json:"id"`
	CurrentSelector string `json:"current_selector"`
	SlotName        string `json:"slot_name"`
	MediaType       string `json:"media_type"`
	Https           int    `json:"https"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
}

type MeRequestSelector struct {
	SlotId          string `json:"slot_id"`
	AppId           string `json:"app_id"`
	SlotType        int    `json:"slot_type"`
	NativeType      int    `json:"native_type"`
	PackageName     string `json:"package_name"`
	PackageVersion  string `json:"package_version"`
	AppName         string `json:"app_name"`
	AppStoreUrl     string `json:"app_storeurl"`
	Template        int    `json:"template"`
	AdId            int    `json:"ad_id"`
	SupportDeeplink int    `json:"support_deeplink"`
	SupportClkxy    int    `json:"support_clickxy"`
	AdCount         int    `json:"ad_count"`
	Price           int    `json:"price"`
	PriceType       int    `json:"price_type"`
	Position        int    `json:"position"`
	AppKey          string `json:"app_key"`
}

type MeRequestDevice struct {
	Ip              string    `json:"ip"`
	Ua              string    `json:"ua"`
	Imei            string    `json:"imei"`
	ImeiMd5         string    `json:"imei_md5"`
	AndroidId       string    `json:"android_id"`
	Idfa            string    `json:"idfa"`
	IdfaMd5         string    `json:"idfa_md5"`
	PhysicalWidth   int       `json:"physical_width"`
	PhysicalHeight  int       `json:"physical_height"`
	PhysicalDensity string    `json:"physical_density"`
	Mac             string    `json:"mac"`
	Vendor          string    `json:"vendor"`
	Model           string    `json:"model"`
	Os              string    `json:"os"`
	OsVersion       string    `json:"os_version"`
	Network         int       `json:"network"`
	Operator        int       `json:"operator"`
	DeviceType      int       `json:"type"`
	Location        *Location `json:"location"`
	Orientation     int       `json:"orientation"`
}

type Location struct {
	Lat       string `json:"lat"`
	Lon       string `json:"lng"`
	Timestamp int64  `json:"timestamp"`
}

type MeRequestApp struct {
	Version     string `json:"version"`
	PackageName string `json:"package_name"`
}

type MeRequestUser struct {
	Id     string `json:"id"`
	Gender int    `json:"gender"`
	Age    string `json:"age"`
}

type MeRequestRegion struct {
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
}

type MeRequestExt struct {
	FilterCURL string `json:"filter_curl"`
	FilterURL  string `json:"filter_url"`
}

type MeResponse struct {
	Code    int                `json:"code"`
	Msg     string             `json:"msg"`
	Id      string             `json:"id"`
	Width   int                `json:"width"`
	Height  int                `json:"height"`
	AdId    int                `json:"ad_id"`
	SlotId  string             `json:"slot_id"`
	Content *MeResponseContent `json:"content"`
	Union   string             `json:"union"`
	Ext     *MeResponseExt     `json:"ext,omitempty"`
}

type MeResponseContent struct {
	ContentType   string                          `json:"type"`
	Action        int                             `json:"action"`
	Url           string                          `json:"url"`
	CUrl          string                          `json:"c_url"`
	Imp           []string                        `json:"imp"`
	Clk           []string                        `json:"clk"`
	Title         string                          `json:"title,omitempty"`
	Desc          string                          `json:"desc,omitempty"`
	ExtUrls       []string                        `json:"ext_urls,omitempty"`
	IconUrl       string                          `json:"icon_url,omitempty"`
	DpUrl         string                          `json:"dp_url,omitempty"`
	DpClk         []string                        `json:"dp_clk,omitempty"`
	DownStartUrls []string                        `json:"down_start_urls,omitempty"`
	DownUrls      []string                        `json:"down_urls,omitempty"`
	InstStartUrls []string                        `json:"inst_start_urls,omitempty"`
	InstUrls      []string                        `json:"inst_urls,omitempty"`
	App           *MeResponseContentApp           `json:"app,omitempty"`
	Duration      int                             `json:"duration"`
	VideoTracking *MeResponseContentVideoTracking `json:"video_tracking,omitempty"`
	HTML          string                          `json:"html,omitempty"`
}

type MeResponseContentApp struct {
	PackageName string `json:"package_name,omitempty"`
	IconUrl     string `json:"icon_url,omitempty"`
	Size        int    `json:"size,omitempty"`
	Name        string `json:"name,omitempty"`
}

type MeResponseContentVideoTracking struct {
	End   []string `json:"end"`
	Start []string `json:"start"`
	Close []string `json:"close"`
}

type MeResponseExt struct {
}
