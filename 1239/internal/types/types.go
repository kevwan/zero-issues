// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginReply struct {
	Id uint32 `json:"id"`
}
