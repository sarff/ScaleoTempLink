package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func goDotEnvVariable(key string) string {

	// load .env file
	errLoad := godotenv.Load(".env")

	if errLoad != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"pass"`
	ChatID   int64  `json:"chat_id"`
}

type Info struct {
	Info InfoVal `json:"info"`
}

type InfoVal struct {
	Link string `json:"one_time_login_link"`
}

func main() {
	f_email := flag.String("f_email", "", "")
	f_pass := flag.String("f_pass", "", "")
	f_chat := flag.Int64("f_chat", 0, "")
	flag.Parse()

	url := goDotEnvVariable("SCALEO_URL")
	method := "POST"
	client := &http.Client{}

	if *f_email == "" {
		jsonFile, err := os.Open("users.json")
		if err != nil {
			fmt.Println(err)
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		var users Users
		json.Unmarshal(byteValue, &users)

		for i := 0; i < len(users.Users); i++ {
			var rawText = fmt.Sprintf(`{
	    "email": "%s",
	    "password": "%s"
		}`, users.Users[i].Email, users.Users[i].Password)
			payload := strings.NewReader(rawText)
			req, err := http.NewRequest(method, url, payload)
			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Add("Content-Type", "application/json")

			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			var info Info
			json.Unmarshal(body, &info)
			link := info.Info.Link

			chatId := users.Users[i].ChatID
			run(chatId, link)

		}
	} else {
		var rawText = fmt.Sprintf(`{
	    "email": "%s",
	    "password": "%s"
		}`, *f_email, *f_pass)
		payload := strings.NewReader(rawText)
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var info Info
		json.Unmarshal(body, &info)
		link := info.Info.Link

		chatId := *f_chat
		run(chatId, link)
	}

}
