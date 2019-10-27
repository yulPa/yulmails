package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gitlab.com/tortuemat/yulmails/plugins/sdk"
)

var (
	pluginName    = "spamassassin"
	version       = "0.1.0"
	customHeaders = []string{
		"X-Spam-Checker-Version",
		"X-Spam-Level",
		"X-Spam-Status",
	}
)

func newResult(s int, d string, exec int64) *sdk.Result {
	return &sdk.Result{
		Score:    s,
		Name:     pluginName,
		Version:  version,
		ExecTime: exec,
		Details:  d,
	}
}

func getScore(m *mail.Message) (int, error) {
	header := m.Header
	status := header.Get("X-spam-status")
	re := regexp.MustCompile(`score=([0-9]*\.?[0-9]*)`)
	match := re.FindStringSubmatch(status)
	s := match[1]
	if s == "" {
		return 5, nil
	}
	score, err := strconv.Atoi(s)
	if err != nil {
		return 5, nil
	}
	return score, nil
}

func getCustomHeaders(m *mail.Message) []map[string]string {
	headers := make([]map[string]string, 0, len(customHeaders))
	for _, h := range customHeaders {
		headers = append(headers, map[string]string{h: m.Header.Get(h)})
	}
	return headers
}

func testEmailAPI(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	res, exec, raw, err := spamassassinCheck(body)
	if err != nil {
		// TODO: status code + stack error + JSON
		fmt.Fprintf(w, err.Error())
	}
	s, err := getScore(res)
	if err != nil {
		// TODO: status code + stack error + JSON
		fmt.Fprintf(w, err.Error())
	}
	result := newResult(s, raw, exec)
	if len(customHeaders) > 0 {
		result.Headers = getCustomHeaders(res)
	}
	payload, err := json.Marshal(result)
	if err != nil {
		// TODO: status code + stack error + JSON
		fmt.Fprintf(w, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func spamassassinCheck(m []byte) (*mail.Message, int64, string, error) {
	start := time.Now()
	cmd := fmt.Sprintf("echo %s | spamassassin", m)
	delta := time.Since(start)
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return nil, 0, "", err
	}
	// since o is an email, we can parse it
	// with net/mail
	r := strings.NewReader(string(output))
	email, err := mail.ReadMessage(r)
	if err != nil {
		return nil, 0, "", err
	}
	return email, int64(delta), string(output), nil
}

func main() {
	http.HandleFunc("/check", testEmailAPI)
	log.Fatal(http.ListenAndServe(":12800", nil))
}
