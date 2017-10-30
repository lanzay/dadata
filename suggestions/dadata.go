/*
Golang client library for DaData.ru (https://dadata.ru/).

While implemented only suggest (https://dadata.ru/api/suggest/).
*/
package suggestions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	//"log"
	//"errors"
)

const BASE_URL = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/suggest/"

type DaData struct {
	ApiKey     string
	httpClient *http.Client
}

type Query struct {
	Query string `json:"query"`
	Count int    `json:"count"`
}

/*
Create new client of DaData.

Api and secret keys see on profile page (https://dadata.ru/profile/).
*/
func NewDaData(apiKey string) *DaData {
	return NewDaDataCustomClient(apiKey, &http.Client{})
}

/*
Create new custom client of DaData. By example, this option should be used to Google AppEngine:
    ctx := appengine.NewContext(request)
    appEngineClient := urlfetch.Client(ctx)
    daData:= NewDaDataCustomClient(apiKey, secretKey, appEngineClient)
*/
func NewDaDataCustomClient(apiKey string, httpClient *http.Client) *DaData {
	return &DaData{
		ApiKey:     apiKey,
		httpClient: httpClient,
	}
}

func (daData *DaData) sendRequest(lastUrlPart string, query Query, result interface{}) error {
	buffer := &bytes.Buffer{}
	
	if encodeErr := json.NewEncoder(buffer).Encode(query); nil != encodeErr {
		fmt.Printf("encodeErr: %v", encodeErr)
		return encodeErr
	}

	//fmt.Println("[BODY]",buffer.String())
	
	request, requestErr := http.NewRequest("POST", BASE_URL+lastUrlPart, buffer)
	if nil != requestErr {
		fmt.Printf("requestErr: %v", requestErr)
		return requestErr
	}
	
	//request.Header.Set("User-Agent","Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; chromeframe/12.0.742.112)")
	request.Header.Add("Authorization", fmt.Sprintf("Token %s", daData.ApiKey))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	
	response, httpErr := daData.httpClient.Do(request)
	if nil != httpErr {
		fmt.Printf("httpErr: %v", httpErr)
		return httpErr
	}
	
	defer response.Body.Close()
	
	if http.StatusOK != response.StatusCode {
		return fmt.Errorf("Request error %v", response.Status)
	}
	
	if decodeErr := json.NewDecoder(response.Body).Decode(&result); nil != decodeErr {
		fmt.Printf("decodeErr: %v", decodeErr)
		return decodeErr
	}
	
	return nil
	
}
