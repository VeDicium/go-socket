package sermo

type Request struct {
	Method  string                 `json:"method"`
	URL     string                 `json:"url"`
	Query   map[string]interface{} `json:"query"`
	Body    map[string]interface{} `json:"body"`
	Headers map[string]interface{} `json:"headers"`

	RequestID string `json:"requestId"`
}

type Response struct {
	Type string      `json:"type"`
	URL  string      `json:"url"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`

	RequestID string `json:"requestId"`
	Client    Client `json:"-"`
}

func (r Response) Send(res Response) (int, error) {
	res.Type = r.Type
	res.URL = r.URL
	res.RequestID = r.RequestID
	return r.Client.Write(res)
}

type Route struct {
	Method        string
	Version       string
	URL           string
	RouteFunction func(req Request, res Response) (int, error)
}
type Routes []Route

func (r *Routes) Get(version string, url string, routeFunction func(req Request, res Response) (int, error)) {
	r.RegisterRoute("get", version, url, routeFunction)
}

func (r *Routes) Post(version string, url string, routeFunction func(req Request, res Response) (int, error)) {
	r.RegisterRoute("post", version, url, routeFunction)
}

func (r *Routes) Put(version string, url string, routeFunction func(req Request, res Response) (int, error)) {
	r.RegisterRoute("put", version, url, routeFunction)
}

func (r *Routes) Patch(version string, url string, routeFunction func(req Request, res Response) (int, error)) {
	r.RegisterRoute("patch", version, url, routeFunction)
}

func (r *Routes) Delete(version string, url string, routeFunction func(req Request, res Response) (int, error)) {
	r.RegisterRoute("delete", version, url, routeFunction)
}

func (r *Routes) RegisterRoute(method string, version string, url string, routeFunction func(req Request, res Response) (int, error)) {
	route := Route{
		Method:        method,
		Version:       version,
		URL:           "/" + version + url,
		RouteFunction: routeFunction,
	}

	*r = append(*r, route)
}
