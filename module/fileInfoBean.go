package module

type FileInfoBean struct {
	Name  string `json:"name"`
	Size  string `json:"size"`
	Cdt   string `json:"cdt"`
	IsDir string `json:"isdir"`
	Url   string `json:"url"`
}
