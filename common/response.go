package common

import (
	corev1 "k8s.io/api/core/v1"
	"time"
)

type Deploy struct {
	Name string `json:"name"`
}

type Ns struct {
	Name     string                `json:"name"`
	Status   corev1.NamespacePhase `json:"status"`
	CreateAt string                `json:"create_at"`
}

type Node struct {
	Name     string                 `json:"name"`
	Ip       string                 `json:"ip"`
	Status   corev1.ConditionStatus `json:"status"`
	CreateAt string                 `json:"create_at"`
}

type Pod struct {
	Name      string `json:"name"`
	Namespase string `json:"namespase"`
	Ready     bool   `json:"ready"`
	Status    corev1.ConditionStatus
	NodeIp    string `json:"node_ip"`
}

type Service struct {
	Name string `json:"name"`
}

type Spark struct {
	Name        string    `json:"name"`
	Uid         uint      `json:"u_id"`
	Sid         uint      `json:"s_id"`
	PodList     []Pod     `json:"pod_list"`
	DeployList  []Deploy  `json:"deploy_list"`
	ServiceList []Service `json:"service_list"`
}

type UserInfo struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Role      uint      `json:"role"`
	Avatar    string    `json:"avatar"`
}

type Response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type ResponseOfValidator struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  interface{} `json:"status_msg,omitempty"`
}

func ValidatorResponse(err error) ResponseOfValidator {
	return ResponseOfValidator{
		-1,
		translate(err),
	}
}

var OK = Response{StatusCode: 0, StatusMsg: "success"}
var NoRole = Response{StatusCode: -1, StatusMsg: "权限不够"}
var NoToken = Response{StatusCode: -1, StatusMsg: "No Token"}
var TokenExpired = Response{StatusCode: -1, StatusMsg: "token过期"}
var NoUid = Response{StatusCode: -1, StatusMsg: "Uid获取失败"}

type LoginResponse struct {
	Response
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

type UserListResponse struct {
	Response
	Page     int        `json:"page"`
	UserList []UserInfo `json:"user_list"`
}

type NodeListResponse struct {
	Response
	Length   int    `json:"length"`
	NodeList []Node `json:"node_list"`
}
type NsListResponse struct {
	Response
	Length int  `json:"length"`
	NsList []Ns `json:"ns_list"`
}
type PodListResponse struct {
	Response
	Length  int   `json:"length"`
	PodList []Pod `json:"pod_list"`
}

// DeployListResponse pod控制器返回结果
type DeployListResponse struct {
	Response
	Length     int      `json:"length"`
	DeployList []Deploy `json:"deploy_list"`
}

// ServiceListResponse 服务返回结果
type ServiceListResponse struct {
	Response
	Length      int       `json:"length"`
	ServiceList []Service `json:"service_list"`
}

type SparkListResponse struct {
	Response
	Length    int     `json:"length"`
	SparkList []Spark `json:"spark_list"`
}