package solisapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const jsonBody = "{\"pageNo\":1,\"pageSize\":10}"
const jsonBodyMD5 = "kxdxk7rbAsrzSIWgEwhH4w=="

func FetchProduction(config *SolisConfig) (*string, error) {
	if !IsValid(config) {
		return nil, fmt.Errorf("valid config must be provided")
	}

	// Get list of user's stations
	stationList, err := getStationList(config)
	if err != nil {
		return nil, err
	}

	// Select the first station, a normal small house would only have the one station
	station := stationList.Data.Page.Records[0]

	power := fmt.Sprintf("%f", station.Power)

	return &power, nil
}

func getStationList(config *SolisConfig) (*StationListResponse, error) {
	// Get list of user's stations
	req, err := createGetStationListRequest(config)
	if err != nil {
		return nil, err
	}

	httpClient := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get station list: [%d] %s", res.StatusCode, res.Status)
	}

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	stationList := StationListResponse{}
	err = json.Unmarshal(responseBody, &stationList)
	if err != nil {
		return nil, err
	}

	return &stationList, nil
}

func createGetStationListRequest(config *SolisConfig) (*http.Request, error) {
	method := "POST"
	url := "/v1/api/userStationList"
	uri := fmt.Sprintf("%s%s", config.SolisApiUrl, url)
	bodyReader := bytes.NewReader([]byte(jsonBody))

	req, err := http.NewRequest(
		method,
		uri,
		bodyReader,
	)

	if err != nil {
		return nil, err
	}

	contentType := "application/json"
	utcTime := time.Now().UTC().Format(time.RFC1123)
	authorizationPayload := strings.Join([]string{method, jsonBodyMD5, contentType, utcTime, url}, "\n")
	authorization := signAuthorizationPayload(authorizationPayload, config.SolisApiSecret)

	req.Header.Set("Authorization", fmt.Sprintf("API %s:%s", config.SolisApiId, authorization))
	req.Header.Set("Content-MD5", jsonBodyMD5)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Date", utcTime)

	return req, nil
}

func signAuthorizationPayload(payload string, key string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(payload))

	signature := hash.Sum(nil)

	return base64.StdEncoding.EncodeToString(signature)
}
