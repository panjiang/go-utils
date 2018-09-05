package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// HTTP 方法
const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	PATCH  string = "PATCH"
	DELETE string = "DELETE"
)

// Debug 打印调试信息
var Debug = false

// BaseAuth 基本验证
type BaseAuth struct {
	Username string
	Password string
}

// Request 请求
type Request struct {
	Method   string
	Host     string
	Path     string
	BaseAuth *BaseAuth
	Query    map[string]string
}

// Response 返回
type Response struct {
	Status int
	Body   []byte
}

// DebugString 调试打印
func (r *Response) DebugString() string {
	body, _ := PrettyJSON(r.Body)
	return fmt.Sprintf("Status: %d, Body: %s", r.Status, body)
}

func debug(a ...interface{}) {
	if Debug {
		fmt.Fprintln(os.Stdout, a...)
	}
}

// PrettyJSON 格式化JSON
func PrettyJSON(body []byte) (string, error) {
	var obj map[string]interface{}
	err := json.Unmarshal(body, &obj)
	if err != nil {
		return "", err
	}

	b, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Do 执行请求
func Do(data *Request) (*Response, error) {
	u, err := url.Parse(data.Host)
	if err != nil {
		return nil, err
	}
	u.Path = data.Path

	for k, v := range data.Query {
		q := u.Query()
		q.Set(k, v)
		u.RawQuery = q.Encode()
	}

	us := u.String()
	debug(data.Method, us)
	req, err := http.NewRequest(data.Method, us, nil)
	if err != nil {
		return nil, err
	}

	if data.BaseAuth != nil {
		req.SetBasicAuth(data.BaseAuth.Username, data.BaseAuth.Password)
		debug("Authorization:", req.Header.Get("Authorization"))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := Response{
		Status: resp.StatusCode,
		Body:   body,
	}
	debug("Response:", res.Status, string(res.Body))
	return &res, nil
}
