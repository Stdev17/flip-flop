package logic

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

// Request : on echo
func Request() error {
	url := "http://localhost:9000"

  req, _ := http.NewRequest("GET", url, nil)

  client := new(http.Client)
	resp, e := client.Do(req)
	if e != nil {
		return e
	}

  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
	
	return nil
}