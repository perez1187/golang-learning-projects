package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	PhotoApi = "https://api.pexels.com/v1"
	VideoApi = "https://api.pexels.com/videos"
)

type Client struct {
	Token          string
	hc             http.Client
	RemainingTimes int32 // how many free downloading from pixels api
}

// function creating Client
func NewClient(token string) *Client { // return a point to Client
	c := http.Client{}
	return &Client{Token: token, hc: c}
}

type SearchResult struct {
	Page         int32   `json:"page"` //json value comes from api, and value like Page we use in our program
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_Results`
	NextPage     string  `json:"next_page`
	Photos       []Photo `json:"photos` // slice of photos
}

type Photo struct {
	Id     int32  `json:"id"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
	Url    string `json:"url"`
	// Photographer    string      `json:"photographer"`
	// PhotographerUrl string      `json:"photographer_url"`
	// Src             PhotoSource `json:"src"` // another struct
}

// type PhotoSource struct {
// 	Original  string `json:"original"`
// 	Large     string `json:"large"`
// 	Large2x   string `json:"large2x"`
// 	Medium    string `json:"medium"`
// 	Small     string `json:small"`
// 	Potrait   string `json:"portrait"`
// 	Square    string `json:"square"`
// 	Landscape string `json:"landscape"`
// 	Tiny      string `json:"tiny"`
// }

// func (c *Client) SearchPhotos(query string, perPage, page int) (*SearchResult, error) { // it return resutl or err
// 	//The fmt. Sprintf function in the GO programming language is a function used to return a formatted string
// 	url := fmt.Sprintf(PhotoApi+"/search?query=%s&per_page=%d&page=%d", query, perPage, page)
// 	resp, err := c.requestDoWithAuth("GET", url) // we make request here, and capture to resp
// 	defer resp.Body.Close()                      // we close this function,when the function is finish

// 	data, err := ioutil.ReadAll(resp.Body) // we use ioutil package to read the body

// 	// fmt.Printf("data from resp.Body: %v\n", data)

// 	if err != nil { // if error is not nil, show err
// 		return nil, err
// 	}
// 	var result SearchResult // SearchResult is a struct, so it is a type

// 	// we capture result
// 	// Unmarshal save data to result
// 	err = json.Unmarshal(data, &result) // https://pkg.go.dev/encoding/json#Unmarshal
// 	return &result, err                 // so we return SearchResult "object"
// }

// func (c *Client) requestDoWithAuth(method, url string) (*http.Response, error) { // response http.Response or error

// 	// we capture NewRequest to req
// 	// https://pkg.go.dev/net/http
// 	req, err := http.NewRequest(method, url, nil) //http that we imported, and NewRequest method from http

// 	if err != nil {
// 		return nil, err
// 	}

// 	// we setup Header
// 	req.Header.Add("Authorization", c.Token) // c give us acces to Token in struct, because c is the client

// 	// we send request
// 	// Once the http.Request is created and configured, you use the Do method of http.DefaultClient to send the request to the server
// 	resp, err := c.hc.Do(req) // hc is http.Client (from struct), we call Do method

// 	if err != nil {
// 		return resp, err
// 	}

// 	// we check remaining numbers of access to Api
// 	// strconv - string representation
// 	// example:
// 	// i, err := strconv.Atoi("-42")
// 	// s := strconv.Itoa(-42)

// 	times, err := strconv.Atoi(resp.Header.Get("X-Ratelimit-Remaining"))
// 	if err != nil {
// 		return resp, nil
// 	} else {
// 		c.RemainingTimes = int32(times) // we set remainigTimes in Struct field to int32
// 	}
// 	return resp, nil
// }

func main() {
	fmt.Print("hello world\n")
	// we manually set token
	os.Setenv("PexelsToken", "VKSbcM6xRE1punB1SWPXueRe6Osb6xAkYarp5G5Wy0QexaroT1drmQKL") // create account on pexels

	// we get token from env
	var TOKEN = os.Getenv("PexelsToken")

	fmt.Printf("Token: %v \n", TOKEN)

	//we create client to work with pexels api
	var c = NewClient(TOKEN)

	fmt.Printf("c: %v\n", c)
	fmt.Printf("c.Token: %v\n", c.Token)
	fmt.Printf("c.hc: %v\n", c.hc)
	fmt.Printf("c.RemainingTimes: %v\n", c.RemainingTimes)

	result, err := c.SearchPhotos("waves", 1, 1) //in brackects, query string, perPage, page int

	// debug Auth
	// query := "waves"
	// perPage := 15
	// page := 1
	// url := fmt.Sprintf(PhotoApi+"/search?query=%s&per_page=%d&page=%d", query, perPage, page)
	// resp, err := c.requestDoWithAuth("GET", url) // we make request here, and capture to resp

	//now we can handle the error
	if err != nil {
		fmt.Errorf("Search error %v", err)
	}

	fmt.Println(result)
	fmt.Printf("c.RemainingTimes new: %v\n", c.RemainingTimes)
	// if there are no results
	// if resul.Page == 0 {
	// 	fmt.Errorf("search results is wrong")
	// }

	// //if everythink is ok, we print result
	// fmt.Println(result)
}
