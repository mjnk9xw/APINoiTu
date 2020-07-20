package models

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Models struct {
}

func newModels() *Models {
	return &Models{}
}

func (m *Models) CheckTu(preSeq, seq string) (bool, error) {

	preSeq = strings.ToLower(preSeq)
	seq = strings.ToLower(seq)
	preSeq = strings.TrimSpace(preSeq)
	seq = strings.TrimSpace(seq)
	if preSeq == "" || seq == "" {
		return false, errors.New("Từ gửi lên trống")
	}
	if len(preSeq) > 15 || len(seq) > 15 {
		return false, errors.New("Từ quá dài > 14 kí tự")
	}

	preSeqSplit := strings.Split(preSeq, " ")
	if len(preSeqSplit) != 2 {
		return false, errors.New("Từ sai cấu trúc")
	}
	seqSplit := strings.Split(seq, " ")
	if len(seqSplit) != 2 {
		return false, errors.New("Từ sai cấu trúc")
	}

	if preSeqSplit[1] != seqSplit[0] {
		return false, errors.New("Từ không giống yêu cầu của trò chơi")
	}

	checkNghia := m.callAPIvtudien(seq)
	if !checkNghia {
		return false, errors.New("Từ không có nghĩa")
	}
	return true, errors.New("Từ ok nè")
}

func (m *Models) callAPIvtudien(seq string) bool {
	apiUrl := "https://vtudien.com/dstu"
	data := url.Values{}
	data.Set("tu", seq)
	data.Set("src", "viet")
	data.Set("tgt", "viet")
	data.Set("tudien", "dictionary")
	data.Set("key", "testkey")

	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String()

	client := &http.Client{}
	r, e := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	if e != nil {
		log.Println("Server error")
		return false
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return false
	}

	reqOK := m.splitString(string(body), seq)
	return reqOK
}
func (m *Models) splitString(allStr string, seq string) bool {
	listString := strings.Split(allStr, "||")
	for _, v := range listString {
		if v == seq {
			return true
		}
	}
	return false
}
