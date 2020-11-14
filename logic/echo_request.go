package logic

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request : on echo
func Request() (string, error) {
	url := "http://localhost:9000"
	api := "/echo"

	req, _ := http.NewRequest("GET", url+api, nil)

	client := new(http.Client)
	resp, e := client.Do(req)
	if e != nil {
		return "", e
	}

	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	str := string(byteArray)
	fmt.Println(str)

	return str, nil
}
