package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/motapratik/covid19-vaccine-tracker/cowin"
	"github.com/motapratik/covid19-vaccine-tracker/telegram"
)

var CenterStatusMap = make(map[int]int)

func main() {

	if len(os.Args) < 4 {
		fmt.Println("Arguments are missing..!!")
		os.Exit(1)
	}
	botcode := os.Args[1]
	chatid := os.Args[2]
	districtcode := os.Args[3]
	agecheck, _ := strconv.Atoi(os.Args[4])

	// Continues Checking Start
	dt := time.Now()
	// DD-MM-YYYY
	Today := fmt.Sprintf("%d-%02d-%02d", dt.Day(), dt.Month(), dt.Year())
	data, err := cowin.GetCowinApiRepsone(districtcode, Today)
	if err == nil {
		var center Centerlist
		err := json.Unmarshal(data, &center)
		if err == nil {
			SendVaccineUpdates(center, botcode, chatid, agecheck)
		}
	}

	//Take 1 sec sleep
	time.Sleep(1 * time.Second)
	// Continues Checking End
}

func SendVaccineUpdates(center Centerlist, botcode string, chatid string, agecheck int) {
	var telegramMsg string
	for i := 0; i < len(center.CentersArray); i++ {
		center_name := center.CentersArray[i].Name
		address := center.CentersArray[i].Address
		pincode := strconv.Itoa(center.CentersArray[i].PinCode)
		fee_type := center.CentersArray[i].FeeType
		datecount := -1
		telegramMsg = ""
		otherdatedata := ""
		for j := 0; j < len(center.CentersArray[i].Sessions); j++ {
			// If Age group match and Vaccine available then only
			if (center.CentersArray[i].Sessions[j].MinAgeLimit == agecheck) && (center.CentersArray[i].Sessions[j].AvailableCapacity > 0) {
				datecount++
				date := center.CentersArray[i].Sessions[j].Date
				vaccine_type := center.CentersArray[i].Sessions[j].Vaccine
				available := strconv.Itoa(center.CentersArray[i].Sessions[j].AvailableCapacity)
				// Center's first date detail
				if datecount == 0 {
					age_group := strconv.Itoa(center.CentersArray[i].Sessions[j].MinAgeLimit) + "+"
					// Telegram Post message doese not support option for NEW Line. So this is hack to add new line is post message
					telegramMsg = fmt.Sprintf("<b>Center Name: </b>%s<pre>                                            </pre><b>Address: </b>%s<pre>                                            </pre><b>Pincode: </b>%s<pre>                        </pre><b>Age Group: </b>%s<pre>                        </pre><b>Date: </b>%s<pre>                        </pre><b>Vaccine: </b>%s<pre>                                                </pre><b>Available: </b>%s<pre>                        </pre><b>Fee Type: </b>%s",
						center_name, address, pincode, age_group, date, vaccine_type, available, fee_type)
				} else {
					otherdatedata = otherdatedata + date + "|" + vaccine_type + "|" + available + "; "
				}
			}
		} // Session Loop End here

		// Send Telegram Message
		if telegramMsg != "" {
			// Check if for Other then Today Vaccine available then add in meesage
			if otherdatedata != "" {
				// telegram post message new line formating hack
				telegramMsg = telegramMsg + "<pre>                        </pre><b>OtherDates: </b>" + otherdatedata
			}
			telegramMsg = telegramMsg + "<pre>                                            </pre>" + "<a>https://www.cowin.gov.in</a>"
			telegram_resp, telegram_err := telegram.SendTelegramMessage(botcode, chatid, telegramMsg)
			if telegram_err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", telegram_err)
			}
			if telegram_resp.StatusCode != 200 {
				fmt.Printf("Bad Response: %d\n", telegram_resp.StatusCode)
			}
		}
	} // Centers Loop End here
}
