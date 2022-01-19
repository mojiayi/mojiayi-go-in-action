package param

/**
* 公共传入参数
 */
type BaseParam struct {
	/*
	* 请求来源IP地址
	 */
	clientIP string
	/*
	* 请求来源浏览器或客户端
	 */
	clientAgent string
	/*
	* 请求访问时间
	 */
	timestamp int64
	/*
	* 请求追踪id
	 */
	traceId string
}

func (p *BaseParam) GetClientIP() string {
	return p.clientIP
}

func (p *BaseParam) SetClientIP(clientIP string) *BaseParam {
	p.clientIP = clientIP
	return p
}

func (p *BaseParam) GetClientAgent() string {
	return p.clientAgent
}

func (p *BaseParam) SetClientAgent(clientAgent string) *BaseParam {
	p.clientAgent = clientAgent
	return p
}

func (p *BaseParam) GetTimestamp() int64 {
	return p.timestamp
}

func (p *BaseParam) SetTimestamp(timestamp int64) *BaseParam {
	p.timestamp = timestamp
	return p
}

func (p *BaseParam) GetTraceId() string {
	return p.traceId
}

func (p *BaseParam) SetTraceId(traceId string) *BaseParam {
	p.traceId = traceId
	return p
}
