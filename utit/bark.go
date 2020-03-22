package utit

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/lzyyauto/auto4play/config"
)

type BarkResponse struct {
	Code    int
	Data    string
	Message string
}

func SendUrl(title, url string) error {
	finalurl := "/" + title + "?url=" + url
	return send(finalurl)
}

func SendMessage(text string) error {
	finalurl := "/" + text
	return send(finalurl)
}

func send(url string) error {
	barkcfg := config.BarkCfg
	url = barkcfg.BarkUrl + barkcfg.BarkKey + url

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var re BarkResponse
	err = json.Unmarshal(body, &re)
	if err != nil {
		return err
	}
	if re.Code != 200 {
		return errors.New(re.Message)
	}

	return nil
}
