package solisapi

import (
	"testing"
)

func TestSignAuthorizationPayload(t *testing.T) {
	authorizationPayload := "POST\nkxdxk7rbAsrzSIWgEwhH4w==\napplication/json\nFri, 26 Jul 2019 06:00:46 GMT\n/v1/api/userStationList"
	authorization := signAuthorizationPayload(authorizationPayload, "6680182547")

	expectedValue := "nBYQWeuzy3Y+gp67BN8zXTmvSDk="

	if authorization != expectedValue {
		t.Fatalf("authorization not calculated correctly")
	}
}
