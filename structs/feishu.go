package structs

var Uuids []string

type Approve struct {
	Challenge string `json:"challenge"`
	Token     string `json:"token"`
	Type      string `json:"type"`
}

type Response struct {
	Challenge string `json:"challenge"`
}

// ApproveTrigger 审批事件订阅
type ApproveTrigger struct {
	UUID  string `json:"uuid"`
	Event Event  `json:"event"`
	Token string `json:"token"`
	Ts    string `json:"ts"`
	Type  string `json:"type"`
}
type Event struct {
	AppID        string `json:"app_id"`
	ApprovalCode string `json:"approval_code"`
	GenerateType string `json:"generate_type"`
	InstanceCode string `json:"instance_code"`
	OpenID       string `json:"open_id"`
	OperateTime  string `json:"operate_time"`
	Status       string `json:"status"`
	TaskID       string `json:"task_id"`
	TenantKey    string `json:"tenant_key"`
	Type         string `json:"type"`
	UserID       string `json:"user_id"`
}

// InstanceInfo 实例详情请求body
type InstanceInfo struct {
	InstanceCode string `json:"instance_code"`
	Locale       string `json:"locale"`
	OpenID       string `json:"open_id"`
	UserID       string `json:"user_id"`
}

// RequestToken 获取飞书token请求体
type RequestToken struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

// ResponseToken 获取飞书token响应体
type ResponseToken struct {
	code              int
	msg               string
	TenantAccessToken string `json:"tenant_access_token"`
	expire            int
}

// Instance 实例详情响应body
type Instance struct {
	Code int    `json:"code"`
	Data Data   `json:"data"`
	Msg  string `json:"msg"`
}
type TaskList struct {
	EndTime   int64  `json:"end_time"`
	ID        string `json:"id"`
	NodeID    string `json:"node_id"`
	NodeName  string `json:"node_name"`
	OpenID    string `json:"open_id"`
	StartTime int64  `json:"start_time"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	UserID    string `json:"user_id"`
}
type CcUserList struct {
	CcID   string `json:"cc_id"`
	OpenID string `json:"open_id"`
	UserID string `json:"user_id"`
}
type Timeline struct {
	CreateTime int64        `json:"create_time"`
	Ext        string       `json:"ext"`
	OpenID     string       `json:"open_id"`
	Type       string       `json:"type"`
	UserID     string       `json:"user_id,omitempty"`
	NodeKey    string       `json:"node_key,omitempty"`
	TaskID     string       `json:"task_id,omitempty"`
	CcUserList []CcUserList `json:"cc_user_list,omitempty"`
	OpenIDList []string     `json:"open_id_list,omitempty"`
	UserIDList []string     `json:"user_id_list,omitempty"`
}
type Data struct {
	ApprovalCode string        `json:"approval_code"`
	ApprovalName string        `json:"approval_name"`
	CommentList  []interface{} `json:"comment_list"`
	DepartmentID string        `json:"department_id"`
	EndTime      int64         `json:"end_time"`
	Form         string        `json:"form"`
	OpenID       string        `json:"open_id"`
	SerialNumber string        `json:"serial_number"`
	StartTime    int64         `json:"start_time"`
	Status       string        `json:"status"`
	TaskList     []TaskList    `json:"task_list"`
	Timeline     []Timeline    `json:"timeline"`
	UserID       string        `json:"user_id"`
	UUID         string        `json:"uuid"`
}

// Form 审批里自定义表单内容
type Form []struct {
	ID     string      `json:"id"`
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Ext    interface{} `json:"ext"`
	Value  string      `json:"value"`
	Option Option      `json:"option,omitempty"`
}
type Option struct {
	Key  string `json:"key"`
	Text string `json:"text"`
}
