package module

type MetaConf struct {
	Apiversion      string `yaml:"apiversion"`
	Name            string `yaml:"name"`
	Port            string `yaml:"port"`
	SendMailCmd     string `yaml:"sendmailcmd"`
	JwtKey          string `yaml:"jwtkey"`
	Mail            string `yaml:"mail"`
	ZhangXin        string `yaml:"zhangxin"`
	SendZhangXinCmd string `yaml:"sendzhangxincmd"`
	UploadPath      string `yaml:"uploadpath"`
	BboltDBPath     string `yaml:"bboltdbpath"`
	UploadRmtUrl    string `yaml:"uploadrmturl"`
        LeastsqScript   string `yaml:"leastsqscript"`
        LeastsqRegExp   string `yaml:"leastsqregexp"`
        FitScript       string `yaml:"fitscript"`
        FitGrepExp      string `yaml:"fitgrepexp"`
        CorrScript      string `yaml:"corrscript"`
        CorrGrepExp     string `yaml:"corrgrepexp"`
}
