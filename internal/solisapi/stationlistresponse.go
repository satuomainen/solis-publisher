package solisapi

type StationListResponse struct {
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	Page StationListPage `json:"page"`
}

type StationListPage struct {
	Records []Station `json:"records"`
}

type Station struct {
	Power float32 `json:"power"`
}
