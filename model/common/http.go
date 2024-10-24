package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/awoyai/gin-temp/global"
	"go.uber.org/zap"
)

var cli = &http.Client{}

func SendRequest(req *http.Request, resp any) error {
	t := reflect.TypeOf(resp)
	if t.Kind() != reflect.Ptr {
		return fmt.Errorf("resp")
	}
	response, err := cli.Do(req)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		global.LOG.Error("sendRequest failed", zap.Any("response", *response))
		return fmt.Errorf("sendRequest fail, status code: %d", response.StatusCode)
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	if err := json.Unmarshal(body, resp); err != nil {
		return err
	}
	return nil
}
