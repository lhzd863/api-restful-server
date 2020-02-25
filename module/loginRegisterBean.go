package module

type LoginRegisterBean struct {
        UserId       string     `json:"userid"`
        UserName     string     `json:"username"`
        Password     string     `json:"password"`
        Alias        string     `json:"alias"`
        Mail         string     `json:"mail"`
}

