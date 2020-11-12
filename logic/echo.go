package logic

import (
	"net/http"
	"os"
	"strconv"
	"log"
)

// Metric : Abstraction for metric data
type Metric struct {
	metrics map[string]int
}

func (m Metric) get(w http.ResponseWriter, req *http.Request) {

}

func (m Metric) set(w http.ResponseWriter, req *http.Request) {

}

func (m Metric) inc(w http.ResponseWriter, req *http.Request) {
	
}

func (m Metric) echo(w http.ResponseWriter, req *http.Request) {
	
}

// Init : Starting an echo server
func Init() {
  metric := Metric{metrics: map[string]int{}}
  http.HandleFunc("/get", metric.get)
  http.HandleFunc("/set", metric.set)
	http.HandleFunc("/inc", metric.inc)
	http.HandleFunc("/echo", metric.echo)

  portnum := 9000
  if len(os.Args) > 1 {
    portnum, _ = strconv.Atoi(os.Args[1])
	}
	
  log.Printf("Going to listen on port %d\n", portnum)
  log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portnum), nil))
}