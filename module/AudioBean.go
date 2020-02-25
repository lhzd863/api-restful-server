package module

type AudioBean struct {
	Catalog    string `json:"catalog"`
	Name       string `json:"name"`
	Info       string `json:"info"`
	Url        string `json:"url"`
	Context    string `json:"context"`
	Amount     string `json:"amount"`
	Img        string `json:"img"`
	Udt        string `json:"udt"`
	Status     string `json:"status"`
	Announcer  string `json:"announcer"`
	SubmitUser string `json:"submituser"`
        Author     string `json:"author"`
        Format     string `json:"format"`
}
