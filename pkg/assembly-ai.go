package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type IClient interface {
	UploadFile(string) (string, error)
	StartTranscription(string) (string, error)
	GetTranscribedText(string) (string, error)
}

type Client struct {
	ApiKey string
}

const TRANSCRIPT_URL = "https://api.assemblyai.com/v2/transcript"
const UPLOAD_URL = "https://api.assemblyai.com/v2/upload"

func NewClient(apiKey string) IClient {
	return &Client{ApiKey: apiKey}
}

func (c *Client) UploadFile(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", UPLOAD_URL, bytes.NewBuffer(data))
	req.Header.Set("authorization", c.ApiKey)
	res, err := client.Do(req)

	defer res.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	if err, isExist := result["error"]; isExist {
		return "", errors.New(err.(string))
	}
	return result["upload_url"].(string), nil
}

func (c *Client) StartTranscription(auditUrl string) (string, error) {
	values := map[string]string{"audio_url": auditUrl}
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", TRANSCRIPT_URL, bytes.NewBuffer(jsonData))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", c.ApiKey)
	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	if err, isExist := result["error"]; isExist {
		return "", errors.New(err.(string))
	}

	return result["id"].(string), nil
}

func (c *Client) GetTranscribedText(id string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", TRANSCRIPT_URL+"/"+id, nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", c.ApiKey)
	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	if err, isExist := result["error"]; isExist {
		return "", errors.New(err.(string))
	}

	if status, isExist := result["status"]; isExist && status == "processing" {
		return "", nil
	}
	return result["text"].(string), nil
}
