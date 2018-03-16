package main

import (
	"net/http"
	"testing"
	"net/http/httptest"
)

func Test_Get(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	HomePageHandle(res, req)

	if res.Code != 200 {
		t.Fatalf("Expected status to == 200, but got %d", res.Code)
	}
}

func Test_Post(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", nil)
	HomePageHandle(res, req)
	if res.Code == 200 {
		t.Fatalf("Expected status to == 200, but got %d", res.Code)
	}
}