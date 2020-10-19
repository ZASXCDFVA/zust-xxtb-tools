package main

type CheckInStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Id    int `json:"id"`
		State int `json:"state"`
	} `json:"data"`
}

type CheckInResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
