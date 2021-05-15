package cowin

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//getCowinApiRepsone function give json array data of API
func GetCowinApiRepsone(districtid string, date string) (resp_data []byte, err error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	url := fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?district_id=%s&date=%s", districtid, date)
	req, req_err := http.NewRequest("GET", url, nil)
	if req_err != nil {
		fmt.Printf("The HTTP New Request created with  error %s\n", req_err)
	}
	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)

	resp, resp_err := client.Do(req)
	if resp_err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", resp_err)
	}

	return ioutil.ReadAll(resp.Body)

}
