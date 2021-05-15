package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//SendTelegramMessage will send message to telegram group
func SendTelegramMessage(botcode, chatid, message string) (response *http.Response, err error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	telegram_bot_url := fmt.Sprintf("https://api.telegram.org/%s/sendMessage?chat_id=%s&parse_mode=HTML", botcode, chatid)
	values := map[string]string{"text": message}
	jsonStr, _ := json.Marshal(values)
	req, req_err := http.NewRequest("POST", telegram_bot_url, bytes.NewBuffer(jsonStr))
	req.Header.Set("content-type", "application/json")

	if req_err != nil {
		fmt.Printf("The HTTP New Request created with  error %s\n", req_err)
	}

	resp, resp_err := client.Do(req)
	return resp, resp_err
}
