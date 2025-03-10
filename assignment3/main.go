package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const URL = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

func GetData() (*string, error) {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Errorf("Error Request url: %s", URL)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Error Reading body %s", URL)
		return nil, err
	}
	strBody := string(body)
	return &strBody, nil
}

func FormatData(data string) []string {
	// replace data
	for _, o := range []string{".", ",", "\n"} {
		for strings.Contains(data, o) {
			data = strings.ReplaceAll(data, o, " ")
		}
	}
	data = strings.TrimSuffix(data, " ")
	for strings.Contains(data, "  ") {
		data = strings.ReplaceAll(data, "  ", " ")
	}
	data = strings.ToLower(data)
	result := strings.Split(data, " ")
	return result
}

func ConvertToMap(data []string) map[string]int {
	mapData := map[string]int{}
	for _, res := range data {
		if _, ok := mapData[res]; !ok {
			mapData[res] = 1
		} else {
			mapData[res] += 1
		}

	}
	return mapData
}

func BuildData() (map[string]int, error) {
	data, err := GetData()
	if err != nil {
		fmt.Errorf("Error get data from api: %v", err)
		return nil, err
	}
	result := FormatData(*data)

	mapData := ConvertToMap(result)
	return mapData, nil
}

func GetBeefSumary(w http.ResponseWriter, r *http.Request) {
	mapData, err := BuildData()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respData := struct {
		Beef map[string]int `json:"beef"`
	}{
		Beef: mapData,
	}

	jsonData, err := json.Marshal(respData)
	if err != nil {
		http.Error(w, "Error Marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/beef/summary", GetBeefSumary)
	http.ListenAndServe(":8080", nil)
}
