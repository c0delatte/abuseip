package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"

	"github.com/jmoiron/jsonq"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	file := usr.HomeDir + "/.config/ipdb.txt"
	API_KEY, err := os.ReadFile(file)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Please put your API_KEY in", file)
		os.Exit(0)

	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		client := http.Client{}
		req, err := http.NewRequest("GET", "https://api.abuseipdb.com/api/v2/check?ipAddress="+scanner.Text(), nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)

		}

		req.Header = http.Header{
			"Key":          []string{strings.TrimRight(string(API_KEY), "\n")},
			"Content-Type": []string{"application/json"},
		}

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)

		data := map[string]interface{}{}
		dec := json.NewDecoder(strings.NewReader(string(body)))
		dec.Decode(&data)
		jq := jsonq.NewQuery(data)

		countryCode, err := jq.String("data", "countryCode")
		abuseScore, err := jq.Int("data", "abuseConfidenceScore")
		totalReports, err := jq.Int("data", "totalReports")
		fmt.Print("[", scanner.Text(), "] Code: ", countryCode, " | Score: ", abuseScore, " | Total Reports: ", totalReports, "\n")
	}
}
