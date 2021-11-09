package main
 
import (
	"flag"
	"fmt"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
	"os"
    "os/exec"
	"strconv"
)

type Fees struct {
    Fastfee    int    `json:"fastestFee"`
    Halfhourfee	int    `json:"halfHourFee"`
    Hourfee	int    `json:"hourFee"`
    MinimumFee	int    `json:"minimumFee"`
}

func main() {
	minfee := flag.Int("alertfee", 10, "maximum fee for which to trigger an alert")
	flag.Parse()
    response, err := http.Get("https://mempool.space/api/v1/fees/recommended")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var responseObject Fees
    json.Unmarshal(responseData, &responseObject)
	fmt.Printf("current fee %s, alert fee %s\n", strconv.Itoa(responseObject.Fastfee), strconv.Itoa(*minfee))
	var triggerfee = *minfee
	if responseObject.Fastfee < triggerfee {
	app := "ls"
    arg := "-ltrh"
	cmd := exec.Command(app, arg)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Print(err.Error())
        return
    }
	fmt.Print(string(stdout))
	} 
}
