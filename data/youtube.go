package data

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Kind  string  `json:"kind"`
	Items []Items `json:"items"`
}

type Items struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}

func GetSubscribers() (Items, error) {
	responseData, err := http.Get("https://www.googleapis.com/youtube/v3/channels?key=AIzaSyCZ2FEhLG4_YQeJbOecGRhNAAaXALghqJE&part=snippet%2CcontentDetails%2Cstatistics&id=UCwFl9Y49sWChrddQTD9QhRA")
	if err != nil {
		return Items{}, err
	}
	bytesData, err := ioutil.ReadAll(responseData.Body)
	if err != nil {
		return Items{}, err
	}
	var response Response
	json.Unmarshal(bytesData, &response)

	return response.Items[0], nil

}
