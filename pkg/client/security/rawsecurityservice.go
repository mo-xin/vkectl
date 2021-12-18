package security
import (
	"net/url"

	"github.com/volcengine/vkectl/pkg/client"

	"github.com/volcengine/vkectl/pkg/model/security/kitex_gen/paastob/productivity/common"
	"github.com/volcengine/vkectl/pkg/model/security/kitex_gen/paastob/vke/security"
)

// SecurityserviceRawApi is a base client
type SecurityserviceRawApi struct {
	Client *client.Client
}

// NewAPIClient 生成一个客户端
func NewAPIClient(ak, sk, host, service, region string) *SecurityserviceRawApi {
	c := client.NewBaseClient()
	c.ServiceInfo = client.NewServiceInfo()
	c.ServiceInfo.Host = host
	c.ServiceInfo.Credentials.AccessKeyID = ak
	c.ServiceInfo.Credentials.SecretAccessKey = sk
	c.ServiceInfo.Credentials.Service = service
	c.ServiceInfo.Credentials.Region = region
	c.SdkVersion = client.DefaultSdkVersion

	return &SecurityserviceRawApi{Client: c}
}
func (p *SecurityserviceRawApi) StartScan(body string, query url.Values) (r *common.EmptyResponse, statusCode int, err error) {
	action := "StartScan"
	r = &common.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *SecurityserviceRawApi) ListBenchmarks(body string, query url.Values) (r *security.ListBenchmarksResponse, statusCode int, err error) {
	action := "ListBenchmarks"
	r = &security.ListBenchmarksResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *SecurityserviceRawApi) GetCronScan(body string, query url.Values) (r *security.GetCronScanResponse, statusCode int, err error) {
	action := "GetCronScan"
	r = &security.GetCronScanResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *SecurityserviceRawApi) UpdateCronScan(body string, query url.Values) (r *common.EmptyResponse, statusCode int, err error) {
	action := "UpdateCronScan"
	r = &common.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *SecurityserviceRawApi) ListNodeReports(body string, query url.Values) (r *security.ListNodeReportsResponse, statusCode int, err error) {
	action := "ListNodeReports"
	r = &security.ListNodeReportsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *SecurityserviceRawApi) StartNodeScan(body string, query url.Values) (r *common.EmptyResponse, statusCode int, err error) {
	action := "StartNodeScan"
	r = &common.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *SecurityserviceRawApi) ListCheckItems(body string, query url.Values) (r *security.ListCheckItemsResponse, statusCode int, err error) {
	action := "ListCheckItems"
	r = &security.ListCheckItemsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *SecurityserviceRawApi) GetCheckItem(body string, query url.Values) (r *security.GetCheckItemResponse, statusCode int, err error) {
	action := "GetCheckItem"
	r = &security.GetCheckItemResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}
