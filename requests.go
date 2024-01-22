package yandexlogistic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func post[T APIResponse](token, base, endpoint string, vals url.Values, payload []byte) (res T, err error) {
	u, err := joinURL(base, endpoint, vals)
	if err != nil {
		return res, err
	}
	fmt.Println("YANDEX BODY REQUEST", string(payload))
	statusCode, cnt, err := sendPostRequest(token, u, payload)
	if err != nil {
		return res, err
	}

	// Debug
	dst := &bytes.Buffer{}
	if err = json.Indent(dst, cnt, "", "  "); err != nil {
		log.Println(err)
	}
	fmt.Println("YANDEX BODY RESPONSE", dst.String())
	// End Debug

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	err = check(statusCode, res)
	return
}

func check(code int, r APIResponse) error {
	b := r.Base()
	// обрабатываем только те, что допустимы по документации YandexAPI https://yandex.ru/dev/logistics/api/about/intro.html
	// все остальное в default error
	switch code {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusTooManyRequests:
		return &APIError{code: Code(b.Code), message: b.Message}
	default:
		return &APIError{code: "unknown", message: "unknown error"}

	}
}

// sendGetRequest is used to send an HTTP GET request.
func sendGetRequest(url string) (int, []byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, []byte{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, []byte{}, err
	}

	return resp.StatusCode, data, nil
}

// sendPostRequest is used to send an HTTP POST request.
func sendPostRequest(token, url string, payload []byte) (int, []byte, error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return 0, []byte{}, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Accept-Language", "ru")

	var client = new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return 0, []byte{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, []byte{}, err
	}

	return res.StatusCode, data, nil
}

func joinURL(base, endpoint string, vals url.Values) (addr string, err error) {
	addr, err = url.JoinPath(base, endpoint)
	if err != nil {
		return
	}

	if vals != nil {
		if queries := vals.Encode(); queries != "" {
			addr = fmt.Sprintf("%s?%s", addr, queries)
		}
	}

	return
}
