package transfer

import (
	"github.com/CodeDing/http-transfer/httplib"
	"github.com/CodeDing/http-transfer/proto/me"
	"net/http"
)

type Service interface {
	MeRequest2AdapterRequest(*me.MeRequest) (*httplib.BeegoHTTPRequest, error)
	AdapterResponse2MeResponse(*me.MeRequest, *http.Response) (*me.MeResponse, error)
}

var serviceManager = make(map[string]Service)

func Register(name string, service Service) {
	if service == nil {
		panic("register service is nil, name: " + name)
	}
	if _, ok := serviceManager[name]; ok {
		panic("service already duplicate register: " + name)
	}
	serviceManager[name] = service
}
