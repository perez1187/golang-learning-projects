package main

// import "net/http"

// type AlertFromAlerta struct {
// 	Environment string `json:"environment"`
// 	Event       string `json:"event"`
// 	Service     string `json:"service"`
// 	Resource    string `json:"resource"`
// }

// type Client struct {
// 	Token string
// 	hc    http.Client
// }

// type SearchAllerts struct {
// 	Status string  `json:"status"`
// 	Alerts []Alert `json:"alerts"`
// }

// type Alert struct {
// 	Environment string   `json:"environment"`
// 	Event       string   `json:"event"`
// 	Service     []string `json:"service"`
// 	Resource    string   `json:"resource"`
// }

// type Server struct {
// 	alertch     chan AlertFromAlerta
// 	TokenAlerta string
// 	hc          http.Client
// 	quitch      chan struct{} // zero memory allocation
// }

// const (
// 	AlertaApi = "https://alerta.tenesys.pl/api/alerts?status=ack&severity=critical&limit=1"
// )
