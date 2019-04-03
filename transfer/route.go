package transfer

import (
	"encoding/json"
	"fmt"
	"http-transfer/proto/me"
	"net/http"
	"strings"

	"github.com/golang/glog"
)

type RouteHandler struct {
}

//http://localhost:8000/api/v1/selector?ad_network=xxx&data=urlencode(json)
func (b *RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	adNetwork := r.Form.Get("ad_network")
	if _, ok := serviceManager[adNetwork]; !ok {
		fmt.Fprintf(w, "%s not registered, please check!", adNetwork)
		return
	}

	adapter := serviceManager[adNetwork]

	data := r.Form.Get("data")
	//TODO
	fmt.Printf("[%s] me request => %s\n", strings.ToUpper(adNetwork), string(data))
	glog.V(5).Infof("[ME-%s] me requset: %s", strings.ToUpper(adNetwork), string(data))
	var req me.MeRequest
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		fmt.Fprintf(w, "wrong format data, please check!")
		return
	}
	beegoReq, err := adapter.MeRequest2AdapterRequest(&req)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	res, err := beegoReq.DoRequest()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	meRes, err := adapter.AdapterResponse2MeResponse(&req, res)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	bs, _ := json.Marshal(meRes)
	//TODO
	fmt.Printf("[%s] me response => %s\n", strings.ToUpper(adNetwork), string(bs))
	glog.V(5).Infof("[%s-ME] me response: %s", strings.ToUpper(adNetwork), string(bs))
	fmt.Fprintf(w, string(bs))
	return
}
