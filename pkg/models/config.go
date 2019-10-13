package models

//Config Main Configuration File Structure
type Config struct {
	WebPort       string   `json:"webport"`
	WebAddress    string   `json:"webaddress"`
	DatabaseName  string   `json:"databasename"`
	APIURL        string   `json:"apiurl"`
	APIKEY	      string 	`json:"apikey"`
	AppName       string   `json:"appname"`
	LogFile       string   `json:"logfile"`
	KeyFile    string   `json:"keyfile"`
	CrtFile        string   `json:"crtfile"`
}

