//go get -t -d github.com/tebeka/selenium
// mkdir vendor
//go mod vendor
//cd vendor
//go run init.go --alsologtostderr  --download_browsers --download_latest

//java -jar selenium-server.jar
//java -jar selenium-server-standalone-3.4.jar
package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
	from_loc := "DEL"
	to_loc := "BOM"
	date := "30/10/2021"
	url := "https://www.makemytrip.com/flight/search?tripType=O&itinerary=" + from_loc + "-" + to_loc + "-" + date + "&paxType=A-1_C-0_I-0&cabinClass=E&sTime=1597828876664&forwardFlowRequired=true"

	// Start a Selenium WebDriver server instance (if one is not already running).
	const (
		// These paths will be different on your system.
		// seleniumPath    = "./vendor/selenium-server"
		seleniumPath = "./chromedriver"
		port         = 4444
	)
	opts := []selenium.ServiceOption{
		// selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		// selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		// selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	// selenium.SetDebug(true)
	service, err := selenium.NewChromeDriverService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))

	err = wd.Get(url)
	time.Sleep(4 * time.Second)

	//Find Flight Light Div
	flightListDiv, err := wd.FindElements(selenium.ByCSSSelector, "div[class='fli-list  simpleow '")
	for _, fl := range flightListDiv {
		text, err := fl.Text()
		fliArr := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
		_ = fliArr
		if err != nil {
			panic(err)
		}
		if len(fliArr) > 7 {
			log.Println("cheapest " + fliArr[0] + " flight is from " + fliArr[2] + "(" + fliArr[1] + ") and to" + fliArr[6] + " to " + "(" + fliArr[5] + ") is " + fliArr[7])
			break
		}
	}
}
