package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func api_call(params map[string]string) ([]byte, error) {
	client := &http.Client{}
	url := ServerURL + "/api/get_account_detail_data_dict"

	api_request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	api_request.Header.Set("authtoken", Authtoken)

	q := api_request.URL.Query()
	for key, value := range params {
		if value != "" {
			q.Add(key, value)
		}
	}
	api_request.URL.RawQuery = q.Encode()

	resp, err := client.Do(api_request)
	if err != nil {
		return nil, fmt.Errorf("Error in API call: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func get_account(ctx context.Context, username string) (SecurdenDataSourceModel, int, string) {
	var account SecurdenDataSourceModel
	params := map[string]string{
		"username": username,
	}
	var Response struct {
		Password   string `json:"password"`
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
		Error      struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
	body, err := api_call(params)
	if err != nil {
		return account, 500, fmt.Sprintf("Error in API call: %v", err)
	}
	json.Unmarshal(body, &Response)
	if Response.StatusCode != 200 {
		if Response.Error.Message != "" {
			return account, Response.StatusCode, Response.Error.Message
		}
		return account, Response.StatusCode, Response.Message
	}
	account.Password = types.StringValue(Response.Password)
	return account, Response.StatusCode, Response.Message
}
