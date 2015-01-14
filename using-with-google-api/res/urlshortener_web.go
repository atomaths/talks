package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"code.google.com/p/goauth2/oauth"
)

var (
	clientId     = flag.String("id", "", "Client ID")
	clientSecret = flag.String("secret", "", "Client Secret")
	cacheFile    = flag.String("cache", "url.cache.json", "Token cache file")
	code         = flag.String("code", "", "Authorization Code")
)

const usageMsg = `
Usage: yourl http://goo.gl/xxxxx     (to look up details)
       yourl http://example.com/long (to shorten)

To obtain a request token you must specify both -id and -secret.

To obtain Client ID and Secret, see the "OAuth 2 Credentials" section under
the "API Access" tab on this page: https://code.google.com/apis/console/

Once you have completed the OAuth flow, the credentials should be stored inside
the file specified by -cache and you may run without the -id and -secret flags.
`

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		fmt.Fprint(os.Stderr, usageMsg)
		os.Exit(2)

	}

	config := &oauth.Config{
		ClientId:     *clientId,
		ClientSecret: *clientSecret,
		Scope:        urlshortener.UrlshortenerScope,
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://accounts.google.com/o/oauth2/token",
		TokenCache:   oauth.CacheFile(*cacheFile),
	}

	// Set up a Transport using the config.
	transport := &oauth.Transport{Config: config}

	// Try to pull the token from the cache; if this fails, we need to get one.
	token, err := config.TokenCache.Token()
	if err != nil {
		if *clientId == "" || *clientSecret == "" {
			flag.Usage()
			fmt.Fprint(os.Stderr, usageMsg)
			os.Exit(2)
		}
		if *code == "" {
			// Get an authorization code from the data provider.
			// ("Please ask the user if I can access this resource.")
			url := config.AuthCodeURL("")
			fmt.Println("Visit this URL to get a code, then run again with -code=YOUR_CODE\n")
			fmt.Println(url)
			return
		}
		// Exchange the authorization code for an access token.
		// ("Here's the code you gave the user, now give me a token!")
		token, err = transport.Exchange(*code)
		if err != nil {
			log.Fatal("Exchange:", err)
		}
		// (The Exchange method will automatically cache the token.)
		fmt.Printf("Token is cached in %v\n", config.TokenCache)
	}

	// Make the actual request using the cached token to authenticate.
	// ("Here's the token, let me in!")
	transport.Token = token

	httpClient := transport.Client()

	svc, err := urlshortener.New(httpClient)
	if err != nil {
		panic(err)
	}

	urlstr := flag.Arg(0)

	// short -> long
	if strings.HasPrefix(urlstr, "http://goo.gl/") || strings.HasPrefix(urlstr, "https://goo.gl/") {
		url, err := svc.Url.Get(urlstr).Do()
		if err != nil {
			log.Fatalf("URL Get: %v", err)
		}
		fmt.Printf("Lookup of %s: %s\n", urlstr, url.LongUrl)
		return
	}

	// long -> short
	url, err := svc.Url.Insert(&urlshortener.Url{
		Kind:    "urlshortener#url", // Not really needed
		LongUrl: urlstr,
	}).Do()

	if err != nil {
		log.Fatalf("URL Insert: %v", err)
	}

	fmt.Printf("Shortened %s => %s\n", urlstr, url.Id)
}
