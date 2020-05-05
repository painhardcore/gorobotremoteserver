package gorobotremoteserver

// Userkeyword is used for registering in Service.
type Userkeyword interface {
	Name() string
	Run(args []string) Result
}

// Result is used for returning results from keyword
type Result struct {
	Status      string      `xml:"status"` // can use package const
	Return      interface{} `xml:"return"` // use struct instead of map
	Error       string      `xml:"error"`
	Traceback   string      `xml:"traceback"` // Logs
	Fatal       bool        `xml:"fatal"`
	Continuable bool        `xml:"continuable"`
}
