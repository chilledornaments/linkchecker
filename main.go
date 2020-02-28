package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func checkLink(link string) (int, bool) {

	client := http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest("GET", link, nil)

	if err != nil {
		panic(fmt.Sprintf("Error creating HTTP client for %s", link))
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error checking link", link, ",", err.Error())
		return 500, false
	}

	resp.Body.Close()

	if resp.StatusCode == 404 {
		return 404, false
	}

	if resp.StatusCode == 401 {
		return 401, false
	}

	return resp.StatusCode, true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must pass only one argument. Arguments passed:", len(os.Args))
		os.Exit(1)
	}

	if !strings.HasPrefix(os.Args[1], "http") {
		fmt.Println("Link must start with http")
		os.Exit(1)
	}
	fmt.Println("Checking link", os.Args[1])

	v, r := checkLink(os.Args[1])

	if !r {
		fmt.Println(fmt.Sprintf("%d found for %s", v, os.Args[1]))
	} else {
		fmt.Println("Link ok for", os.Args[1])
	}

}
