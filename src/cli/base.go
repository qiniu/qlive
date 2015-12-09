package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type LiveHub struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Hub       string `json:"hub"`
}

func (this *LiveHub) ToJson() (jsonStr string, err error) {
	jsonData, mErr := json.Marshal(this)
	if mErr != nil {
		err = errors.New(fmt.Sprintf("Marshal hub info failed, %s", mErr))
		return
	}
	jsonStr = string(jsonData)
	return
}

func (this *LiveHub) String() string {
	return fmt.Sprintf("AccessKey: %s \nSecretKey: %s \nHub: %s", this.AccessKey, this.SecretKey, this.Hub)
}

func (this *LiveHub) Set(accessKey string, secretKey string, hub string) (err error) {
	qLiveHubFolder := ".qlive"
	if _, sErr := os.Stat(qLiveHubFolder); sErr != nil {
		if mErr := os.MkdirAll(qLiveHubFolder, 0775); mErr != nil {
			err = errors.New(fmt.Sprintf("Mkdir `%s' failed, %s", qLiveHubFolder, mErr.Error()))
			return
		}
	}
	qLiveHubFile := filepath.Join(qLiveHubFolder, "hub.json")

	fp, openErr := os.Create(qLiveHubFile)
	if openErr != nil {
		err = errors.New(fmt.Sprintf("Open hub info file error, %s", openErr.Error()))
		return
	}
	defer fp.Close()

	this.AccessKey = accessKey
	this.SecretKey = secretKey
	this.Hub = hub

	jsonStr, mErr := this.ToJson()
	if mErr != nil {
		err = mErr
		return
	}

	_, wErr := fp.WriteString(jsonStr)
	if wErr != nil {
		err = errors.New(fmt.Sprintf("Write hub info failed, %s", wErr.Error()))
		return
	}

	return
}

func (this *LiveHub) Get() (err error) {
	qLiveHubFile := filepath.Join(".qlive", "hub.json")
	fp, openErr := os.Open(qLiveHubFile)
	if openErr != nil {
		err = errors.New(fmt.Sprintf("Get hub info error, %s", openErr.Error()))
		return
	}
	defer fp.Close()
	accountBytes, readErr := ioutil.ReadAll(fp)
	if readErr != nil {
		err = errors.New(fmt.Sprintf("Read hub info error, %s", readErr.Error()))
		return
	}
	if umError := json.Unmarshal(accountBytes, this); umError != nil {
		err = errors.New(fmt.Sprintf("Parse hub info error, %s", umError.Error()))
		return
	}

	return
}
