package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"strings"

	"github.com/fatih/color"
	"net/http"
	"regexp"
)

func hasSchema(uri string) (b bool) {
	if m, _ := regexp.MatchString("https?://", uri); !m {
		return false
	}
	return true
}

func main() {
	app := cli.NewApp()
	app.Name = "httr"
	app.Usage = "Display Http Response Headers"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) {
		red := color.New(color.FgRed, color.Bold).SprintFunc()
		green := color.New(color.FgGreen, color.Bold).SprintFunc()
		yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
		blue := color.New(color.FgBlue, color.Bold).SprintFunc()
		cyan := color.New(color.FgCyan).SprintFunc()

		if len(c.Args()) == 0 {
			fmt.Printf("%s: type \"httr -h\" or \"httr help\" for usage\n", red("Error"))
			return
		}

		request_uri := c.Args()[0]
		if !hasSchema(request_uri) {
			request_uri = "http://" + request_uri
		}

		// DO NOT follow the redirects
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}

		response, err := client.Head(request_uri)
		if err != nil {
			fmt.Printf("%s: Failed to establish a new connection to %s\n%s\n", red("Error"), request_uri, err)
			return
		}

		statusCode := green(response.StatusCode)
		if response.StatusCode != http.StatusOK {
			statusCode = yellow(response.Status)
		}

		fmt.Printf("Proto: %v\n", blue(response.Proto))
		fmt.Printf("Status: %v\n", statusCode)
		fmt.Printf("\n%s\n", blue("HTTP Response Headers"))

		for key, value := range response.Header {
			fmt.Printf("%s: %s\n", key, cyan(strings.Join(value, ", ")))
		}

	}
	app.Run(os.Args)
}
