package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)



var BingDomains = map[string]string{  //alll the domains inserted into a map type formate which have jsonish type and it's a key-value pair tho
	"com":"com", "uk": "co.uk",
}

type SearchResult struct{
	ResultRank int
	ResultURL string
	ResultTitle string
	ResultDesc string
}

var userAgents=[]string{  //a list of browswer it creates, u need to randomize the user ageent, can't use one while making user request
"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"                          
}

func randomUserAgent()(string) {
randomIndex := rand.IntN(len(userAgents))
return userAgents[randomIndex]
}

func buildBingURL(searchTerm string, country string, page int64, count int64,) (string) {
//modifiedArg1:= strings.ReplaceAll(searchTerm, "-", " ")
modifiedArg1:= strings.ReplaceAll(searchTerm, " ", "+")
result:= fmt.Sprintf("https://www.bing.%s/search?q=%s&first=%d&count=%d", country, modifiedArg1, page, count)
return result

}

func executeScrapeRequest(targetURL string) (*http.Response, error){
req, err:=http.NewRequest("GET", targetURL, nil)
if err!=nil{
	log.Fatalf("Error while creatjng the request:", err)
}
req.Header.Set("User-Agent", randomUserAgent)

// client :=http.Client{
// 	Timeout: 10* time.Second,
// }

// resp, err:=client.Do(req)
// if err!=nil{
// 	log.Fatalf("Error sending the request")
// }

resp, err:=http.DefaultClient.Do(req)
if err!=nil{
 log.Fatalf("Error while sending the request: ", err)
}
//defer resp.Body.Close()
return resp
}

func parseBingResults(resp *http.Response) ([]SearchResult, error) {
	
doc, err := goquery.NewDocumentFromReader(resp.Body)
if err != nil {
    return nil, err
}
var results []SearchResult
doc.Find("li.b_algo").Each(func(i int, s *goquery.Selection)){
title := s.Find("h2 a").Text()
desc := s.Find("p").Text()
link, exists := s.Find("h2 a").Attr("href")
item := SearchResult{
    ResultRank:  i + 1,  
    ResultTitle: title,
    ResultURL:   link,
    ResultDesc:  desc,
}
results append(result, item)
}
return results, nil
}
func main() {


}