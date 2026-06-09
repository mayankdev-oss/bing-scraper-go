package main

import(
	"github.com/PuerkitoBio/goquery"
	"strings"
	"net/http"
	"time"
	"math/rand"
	"fmt"
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

func main() {


}