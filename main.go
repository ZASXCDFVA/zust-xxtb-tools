package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		println("configuration not found")
		return
	}

	cfg, err := ParseConfig(os.Args[1])
	if err != nil {
		println(err.Error())
		return
	}

	telegramToken = cfg.Telegram.Token
	telegramChatId = cfg.Telegram.ChatId

	client := http.Client{
		Transport: &MaskTransport{},
	}

	success := 0
	failures := make([]string, 0, len(cfg.Users))

	for _, user := range cfg.Users {
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))

		statusBody := url.Values{}

		statusBody.Set("id", user.Id)
		statusBody.Set("user_type", "1")
		statusBody.Set("environment_type", "101")
		statusBody.Set("round", "")
		statusBody.Set("module_id", "63")

		response, err := client.Post(
			"https://v-xxtb.zust.edu.cn/api/Ncov2019/get_record",
			"application/x-www-form-urlencoded",
			strings.NewReader(statusBody.Encode()),
		)

		if err != nil {
			failures = append(failures, fmt.Sprintf("dial for %s: %s", user.Id, err.Error()))

			continue
		}

		status := &CheckInStatus{}

		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			failures = append(failures, fmt.Sprintf("parse status of %s: %s", user.Id, err.Error()))

			continue
		}

		if status.Code/100 != 2 {
			failures = append(failures, fmt.Sprintf("query status of %s: %s", user.Id, status.Message))

			continue
		}

		for _, data := range status.Data {
			time.Sleep(time.Second * time.Duration(rand.Intn(5)))

			if data.State == 1 {
				success++

				continue
			}

			values := url.Values{}

			values.Set("id", user.Id)
			values.Set("user_type", "1")
			values.Set("environment_type", "101")
			values.Set("r_id", strconv.Itoa(data.Id))
			values.Set("student_id", user.Id)
			values.Set("state", "1")
			values.Set("is_body_ok", "1")
			values.Set("is_gl", "0")
			values.Set("is_tl", "0")
			values.Set("is_jc", "0")
			values.Set("is_2_man", "0")
			values.Set("is_family", "0")
			values.Set("user_location", "1")
			values.Set("location_province", user.Province)
			values.Set("location_city", user.City)
			values.Set("location_country", user.Country)
			values.Set("morning_state", "1")
			values.Set("morning_temperature", "")
			values.Set("afternoon_state", "1")
			values.Set("afternoon_temperature", "")
			values.Set("is_jkm", "1")
			values.Set("round", "")
			values.Set("module_id", "63")

			response, err := client.Post(
				"https://v-xxtb.zust.edu.cn/api/Ncov2019/update_record_student",
				"application/x-www-form-urlencoded",
				strings.NewReader(values.Encode()),
			)

			if err != nil {
				failures = append(failures, fmt.Sprintf("dial for %s/%d: %s", user.Id, data.Id, err.Error()))

				continue
			}

			result := &CheckInResult{}

			if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
				failures = append(failures, fmt.Sprintf("parse result of %s/%d: %s", user.Id, data.Id, err.Error()))

				continue
			}

			if result.Code/100 != 2 {
				failures = append(failures, fmt.Sprintf("update record %s/%d: %s", user.Id, data.Id, result.Message))

				continue
			}

			success++
		}
	}

	output := bytes.NewBuffer([]byte{})

	output.WriteString(fmt.Sprintf("%d/%d successfully\n\n", success, len(cfg.Users)))

	if len(failures) > 0 {
		output.WriteString("failures:\n")

		for _, failure := range failures {
			output.WriteString(failure)
			output.WriteString("\n")
		}
	}

	log(output.String())
}
