package module

type BuyProductBean struct {
        ProductId            string `json:"productid"`
	ProductName          string `json:"productname"`
	ImgUrl1              string `json:"imgurl1"`
        ImgUrl2              string `json:"imgurl2"`
        ImgUrl3              string `json:"imgurl3"`
        ImgUrl4              string `json:"imgurl4"`
        Price                int64  `json:"price"`
        VideoUrl             string `json:"videourl"`
        Cts                  string `json:"cts"`
        Contributor          string `json:"contributor"`
}

