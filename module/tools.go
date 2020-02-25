package module

type Tools struct {
    ToolsId           string            `json:"toolsid"`
    ToolsName         string            `json:"toolsname"`
    Version           string            `json:"version"`
    Img               string            `json:"img"`
    Author            string            `json:"author"`
    Url               string            `json:"url"`
    AccessCnt         string            `json:"accesscnt"`
    Des               string            `json:"des"`
    Cdt               string            `json:"cdt"`
    Tag               string            `json:"tag"`
    UrlIcon           string            `json:"urlicon"`
    Stars             int64             `json:"stars"`
}
