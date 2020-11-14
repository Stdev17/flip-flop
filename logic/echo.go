package logic

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Metric : Abstraction for metric data
type Metric struct {
	metrics map[string]int
}

func (m Metric) echo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Echo from the server!")
}

// Init : Starting an echo server
func Init() {
	e := echo.New()

	metric := Metric{metrics: map[string]int{}}
	e.GET("/echo", metric.getEcho)

	portnum := 9000
	if len(os.Args) > 1 {
		portnum, _ = strconv.Atoi(os.Args[1])
	}

	log.Printf("Going to listen on port %d\n", portnum)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(portnum)))
}
