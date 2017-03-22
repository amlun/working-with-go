/**
 * rest-api-fetch.go
 *
 * Example fetching data from a REST API endpoint parsing and displaying
 * This example using new listing in Reddit golang group. The key to
 * understanding JSON decoding is understanding the data structure you
 * receive the data. Once data structs are mapped, it becomes relatively
 * straight forward
 *
 */

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// create data structures to hold JSON data
// this needs to match the fields in the JSON feed
// see sample data fetch at bottom of this file

// create an individual entry data structure
// this does not need to hold every field, just the ones we want
type Entry struct {
	Cid      string
	Openid   string
	CN       string
	NickName string
	Avatar   string
}

// the feed is the full JSON data structure
// this sets up the array of Entry types (defined above)
type Feed struct {
	Data []Entry
}

func main() {
	// url of REST endpoint we are grabbing data from
	url := "http://nr-cmt-dev.ksmobile.net/comment/get?resid=test_in&aid=1&sg=3c377e5ff058f37f0dcf208f99d9c4c3&app=NR"

	// fetch url
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error fetching:", err)
	}
	// defer response close
	defer resp.Body.Close()

	// confirm we received an OK status
	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Error Status not OK:", resp.StatusCode)
	}

	// read the entire body of the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading body:", err)
	}

	// create an empty instance of Feed struct
	// this is what gets filled in when unmarshaling JSON
	var entries Feed
	if err := json.Unmarshal(body, &entries); err != nil {
		log.Fatalln("Error decoing JSON", err)
	}

	// loop through the children and create entry objects
	for _, ed := range entries.Data {
		entry := ed
		log.Println(">>>")
		log.Println("Comment_id  :", entry.Cid)
		log.Println("Open_id     :", entry.Openid)
		log.Println("Content     :", entry.CN)
		log.Println("NickName    :", entry.NickName)
		log.Println("Avatar      :", entry.Avatar)
	}
}
