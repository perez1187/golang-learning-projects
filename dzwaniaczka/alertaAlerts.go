package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	AlertaApi = "https://alerta.tenesys.pl/api/alerts?status=ack&severity=critical&limit=1"
)

type Server struct {
	Token string
	hc    http.Client
}

// function creating Client
func NewServer(token string) *Server { // return a point to Client
	c := http.Client{}
	return &Server{Token: token, hc: c}
}

type SearchAllerts struct {
	Status string  `json:"status"`
	Alerts []Alert `json:"alerts"`
}

type Alert struct {
	Environment string   `json:"environment"`
	Event       string   `json:"event"`
	Service     []string `json:"service"`
	Resource    string   `json:"resource"`
}

// type Serv struct {
// 	Testy string
// }

func (c *Server) SearchAllerts() (*SearchAllerts, error) {

	resp, err := c.requestDoWithAuth("GET", AlertaApi)

	defer resp.Body.Close()
	if err != nil {
		log.Println("ERROR: ALERTA: requestDoWithAuth with err: ", err)
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body) // ioutil package to read the body

	if err != nil {
		log.Println("ERROR: ALERTA: SearchAllerts() read body with err: ", err)
		return nil, err
	}
	var result SearchAllerts

	// capture result
	err = json.Unmarshal(data, &result)

	return &result, err
}

func (c *Server) requestDoWithAuth(method, url string) (*http.Response, error) { // response http.Response or error

	// log.Println("INFO: requestDoWithAuth START")
	req, err := http.NewRequest(method, url, nil) //http that we imported, and NewRequest method from http

	// fmt.Println("req, ", req)
	if err != nil {
		log.Println("ERROR ALERTA requestDoWithAuth with error:  ", err)
		return nil, err
	}

	// we setup Header with Token
	tokenAlerta := "Key " + c.Token
	req.Header.Add("Authorization", tokenAlerta) // c give us acces to Token in struct, because c is the client

	// we send request
	log.Println("INFO: sending a request to Alerta - start")
	fmt.Println("Sending req")
	resp, err := c.hc.Do(req) // hc is http.Client (from struct), we call Do method for send question to api

	if err != nil {
		log.Println("ERROR ALERTA resp, err := c.hc.Do(req) with error:  ", err)
		return resp, err
	}
	log.Println("INFO: sending a request to Alerta - success")

	return resp, nil
}
