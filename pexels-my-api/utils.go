package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Client) SearchPhotos(query string, perPage, page int) (*SearchResult, error) { // it return resutl or err
	//The fmt. Sprintf function in the GO programming language is a function used to return a formatted string
	url := fmt.Sprintf(PhotoApi+"/search?query=%s&per_page=%d&page=%d", query, perPage, page)
	resp, err := c.requestDoWithAuth("GET", url) // we make request here, and capture to resp
	defer resp.Body.Close()                      // we close this function,when the function is finish

	data, err := ioutil.ReadAll(resp.Body) // we use ioutil package to read the body

	// fmt.Printf("data from resp.Body: %v\n", data)

	if err != nil { // if error is not nil, show err
		return nil, err
	}
	var result SearchResult // SearchResult is a struct, so it is a type

	// we capture result
	// Unmarshal save data to result
	err = json.Unmarshal(data, &result) // https://pkg.go.dev/encoding/json#Unmarshal
	return &result, err                 // so we return SearchResult "object"
}

func (c *Client) requestDoWithAuth(method, url string) (*http.Response, error) { // response http.Response or error

	// we capture NewRequest to req
	// https://pkg.go.dev/net/http
	req, err := http.NewRequest(method, url, nil) //http that we imported, and NewRequest method from http

	if err != nil {
		return nil, err
	}

	// we setup Header
	req.Header.Add("Authorization", c.Token) // c give us acces to Token in struct, because c is the client

	// we send request
	// Once the http.Request is created and configured, you use the Do method of http.DefaultClient to send the request to the server
	resp, err := c.hc.Do(req) // hc is http.Client (from struct), we call Do method

	if err != nil {
		return resp, err
	}

	// we check remaining numbers of access to Api
	// strconv - string representation
	// example:
	// i, err := strconv.Atoi("-42")
	// s := strconv.Itoa(-42)

	times, err := strconv.Atoi(resp.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return resp, nil
	} else {
		c.RemainingTimes = int32(times) // we set remainigTimes in Struct field to int32
	}
	return resp, nil
}
