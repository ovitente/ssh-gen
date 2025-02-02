package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

type sshConfigOpts struct {
	optName  string
	optValue string
}

type sshConfigBlock struct {
	hostAlias    string
	hostName     string
	userName     string
	portNumber   string
	identityFile string
	additionOpts sshConfigOpts
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func writeToConfig(
	hostAlias string,
	hostName string,
	userName string,
	portNumber string,
	identityFile string,
	additionOpts sshConfigOpts,
) {

	// Filling options structure with values
	//confOpt := sshConfigOpts{
	//	"IdentitiesOnly",
	//	"yes",
	//}

	// Sample printing structure values to the console
	fmt.Printf(
		"Host %s\n  Hostname %s\n  User %s\n  Port %s\n  IdentityFile %s\n %s %s\n\n",
		hostAlias,
		hostName,
		userName,
		portNumber,
		identityFile,
		additionOpts.optName,
		additionOpts.optValue,
	)

	// Creating file and filling it
	//	f, err := os.Create("config")
	//	if err != nil {
	//		fmt.Println(err)
	//		f.Close()
	//		return
	//	}

	//	err = f.Close()
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
}

func main() {

	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := "1pWf496p36IXX_MagZiObkJFwZmwLPDWNQq7OT0E2Ato"
	readRange := "Instances!B2:H"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
		return
	}

	sshuser := os.Args[1:]

	for _, v := range resp.Values {
		//fmt.Fprintln(f, v[3], v[5], v[6]) // Это надо перенести в функцию записи в файл
		hostAlias, ok := v[3].(string)
		if !ok {
			return
		}
		hostName, ok := v[5].(string)
		if !ok {
			return
		}
		portNumber, ok := v[6].(string)
		if !ok {
			return
		}
		writeToConfig(hostAlias, hostName, sshuser[0], portNumber, "~/.ssh/work", sshConfigOpts{optName: " IdentitiesOnly", optValue: "yes"})
	}

}
