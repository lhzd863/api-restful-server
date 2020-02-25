package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"

	"github.com/lhzd863/tools-server/db"
	"github.com/lhzd863/tools-server/ex"
	"github.com/lhzd863/tools-server/jwt"
	"github.com/lhzd863/tools-server/module"
	"github.com/lhzd863/tools-server/util"

	"github.com/satori/go.uuid"
)

// JobResource is the REST layer to the User domain
type ResponseResource struct {
	// normally one would use DAO (data access object)
	Data map[string]interface{}
}

// WebService creates a new service that can handle REST requests for User resources.
func (rrs ResponseResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/api").
		Consumes("*/*").
		Produces(restful.MIME_JSON, restful.MIME_JSON) // you can specify this per route as well

	tags := []string{"mtoken"}

	ws.Route(ws.GET("/health").To(rrs.HealthHandler).
		// docs
		Doc("Health").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/login").To(rrs.LoginHandler).
		// docs
		Doc("login info").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/home/myinfo").To(rrs.MyInfoHandler).
		// docs
		Doc("login info").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/register").To(rrs.RegisterHandler).
		// docs
		Doc("register info").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/tool/list").To(rrs.ToolListHandler).
		// docs
		Doc("Get tools list").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/tool/add").To(rrs.ToolAddHandler).
		// docs
		Doc("Post tools list").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/tool/update").To(rrs.ToolUpdateHandler).
		// docs
		Doc("Post tools list,update value same id").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/tool/delete").To(rrs.ToolRemoveHandler).
		// docs
		Doc("Delete tools list").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/app/upload").To(rrs.UploadHandler).
		// docs
		Doc("Upload file").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/app/upload/filelist").To(rrs.UploadFileListHandler).
		// docs
		Doc("my file list").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/app/upload/filedelete").To(rrs.UploadFileDeleteHandler).
		// docs
		Doc("my file delete").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/app/cal/leastsq").To(rrs.CalLeastsqHandler).
		// docs
		Doc("leastsq").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/app/cal/fitline").To(rrs.CalFitLineHandler).
		// docs
		Doc("fit line").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/app/cal/corr").To(rrs.CalCorrHandler).
		// docs
		Doc("corr").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/app/field/add").To(rrs.FieldAddHandler).
		// docs
		Doc("field add").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/app/field/remove").To(rrs.FieldRemoveHandler).
		// docs
		Doc("field remove").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/app/field/qry").To(rrs.FieldQryHandler).
		// docs
		Doc("field remove").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseResource{}). // on the response
		Returns(200, "OK", ResponseResource{}).
		Returns(404, "Not Found", nil))

        ws.Route(ws.POST("/buy/product/add").To(rrs.BuyAddProductHandler).
                // docs
                Doc("buy product").
                Metadata(restfulspec.KeyOpenAPITags, tags).
                Writes(ResponseResource{}). // on the response
                Returns(200, "OK", ResponseResource{}).
                Returns(404, "Not Found", nil))

        ws.Route(ws.POST("/buy/product/get").To(rrs.BuyGetProductHandler).
                // docs
                Doc("buy get product").
                Metadata(restfulspec.KeyOpenAPITags, tags).
                Writes(ResponseResource{}). // on the response
                Returns(200, "OK", ResponseResource{}).
                Returns(404, "Not Found", nil))
       
        ws.Route(ws.POST("/audio/catalog/list").To(rrs.AudioCatalogListHandler).
                // docs
                Doc("audio catalog list").
                Metadata(restfulspec.KeyOpenAPITags, tags).
                Writes(ResponseResource{}). // on the response
                Returns(200, "OK", ResponseResource{}).
                Returns(404, "Not Found", nil))

        ws.Route(ws.POST("/audio/add").To(rrs.AudioAddHandler).
                // docs
                Doc("audio add").
                Metadata(restfulspec.KeyOpenAPITags, tags).
                Writes(ResponseResource{}). // on the response
                Returns(200, "OK", ResponseResource{}).
                Returns(404, "Not Found", nil))
        
        ws.Route(ws.POST("/audio/remove").To(rrs.AudioRemoveHandler).
                // docs
                Doc("audio add").
                Metadata(restfulspec.KeyOpenAPITags, tags).
                Writes(ResponseResource{}). // on the response
                Returns(200, "OK", ResponseResource{}).
                Returns(404, "Not Found", nil))

        ws.Route(ws.POST("/audio/get").To(rrs.AudioGetHandler).
                // docs
                Doc("audio get").
                Metadata(restfulspec.KeyOpenAPITags, tags).
                Writes(ResponseResource{}). // on the response
                Returns(200, "OK", ResponseResource{}).
                Returns(404, "Not Found", nil))

        ws.Route(ws.POST("/audio/search").To(rrs.AudioSearchHandler).
                // docs
                Doc("audio search").
                Metadata(restfulspec.KeyOpenAPITags, tags).
                Writes(ResponseResource{}). // on the response
                Returns(200, "OK", ResponseResource{}).
                Returns(404, "Not Found", nil))

	return ws
}

func (rrs *ResponseResource) HealthHandler(request *restful.Request, response *restful.Response) {
	list := []module.RetBean{}
	list = append(list, module.RetBean{Code: "200", Status: "ok", Err: "", Data: "{}"})
	response.WriteEntity(list)
}

func (rrs *ResponseResource) LoginHandler(request *restful.Request, response *restful.Response) {

	loginbean := new(module.LoginBean)
	err := request.ReadEntity(&loginbean)
	if err != nil {
		log.Println(err)
		response.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		response.WriteErrorString(401, "401: Not Authorized Parameter Input Error")
		return
	}

	bt := db.NewBoltDB(conf.BboltDBPath, "user")
	defer bt.Close()

	jsonstr := bt.Get(loginbean.UserName)
	logregisterbean := new(module.LoginRegisterBean)
	json.Unmarshal([]byte(jsonstr.(string)), &logregisterbean)

	if logregisterbean.Password != loginbean.Password {
		response.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		response.WriteErrorString(401, "401: Not Authorized UserName Or Password Input Error")
		return
	}

	exp1, _ := strconv.ParseFloat(fmt.Sprintf("%v", time.Now().Unix()+3600*24*30), 64)
	claims := map[string]interface{}{
		"iss": loginbean.UserName,
		"exp": exp1,
	}
	key := []byte(conf.JwtKey)
	encoded, encodeErr := jwt.Encode(
		claims,
		key,
		"HS256",
	)
	if encodeErr != nil {
		fmt.Printf("Failed to encode: ", encodeErr)
		response.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		response.WriteErrorString(401, "401: Not Authorized AccessToken Generator Error")
		return
	}
	acctoken := string(encoded)

	retlst := []module.LoginSucBean{}
	retlst = append(retlst, module.LoginSucBean{AccessToken: acctoken, UserName: loginbean.UserName})
	response.WriteEntity(retlst)
}

func (rrs *ResponseResource) RegisterHandler(request *restful.Request, response *restful.Response) {

	loginregisterbean := new(module.LoginRegisterBean)
	err := request.ReadEntity(&loginregisterbean)
	if err != nil {
		log.Println(err)
		response.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		response.WriteErrorString(401, "401: Not Authorized Parameter Input Error")
		return
	}
	u1 := uuid.Must(uuid.NewV4())
	loginregisterbean.UserId = fmt.Sprint(u1)

	exp1, _ := strconv.ParseFloat(fmt.Sprintf("%v", time.Now().Unix()+3600*24*30), 64)
	claims := map[string]interface{}{
		"iss": loginregisterbean.UserName,
		"exp": exp1,
	}
	key := []byte(conf.JwtKey)
	encoded, encodeErr := jwt.Encode(
		claims,
		key,
		"HS256",
	)
	if encodeErr != nil {
		fmt.Printf("Failed to encode: ", encodeErr)
	}
	acctoken := string(encoded)

	jsonstr, _ := json.Marshal(loginregisterbean)
	bt := db.NewBoltDB(conf.BboltDBPath, "user")
	defer bt.Close()

	bt.Set(loginregisterbean.UserName, string(jsonstr))
	bt.Close()

	retlst := []module.LoginSucBean{}
	retlst = append(retlst, module.LoginSucBean{UserName: loginregisterbean.UserName, AccessToken: acctoken})
	response.WriteEntity(retlst)
}

func (rrs *ResponseResource) AlertHandler(request *restful.Request, response *restful.Response) {
	reqParams, err := url.ParseQuery(request.Request.URL.RawQuery)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(reqParams)

	notification := new(module.Notification)
	err = request.ReadEntity(&notification)
	if err != nil {
		log.Println(err)
		return
	}
	list := []module.RetBean{}
	list = append(list, module.RetBean{Code: "200", Status: "ok", Err: "", Data: "{}"})
	response.WriteEntity(list)
}

func (rrs *ResponseResource) ToolListHandler(request *restful.Request, response *restful.Response) {

	bt := db.NewBoltDB(conf.BboltDBPath, "tools")
	defer bt.Close()

	toolslist := bt.Scan()
	retlst := make([]interface{}, 0)
	for _, v := range toolslist {
		for _, v1 := range v.(map[string]interface{}) {
			tools := new(module.Tools)
			err := json.Unmarshal([]byte(v1.(string)), &tools)
			if err != nil {
				log.Println(err)
				continue
			}
			retlst = append(retlst, tools)
		}
	}
	response.WriteEntity(retlst)

}

func (rrs *ResponseResource) MyFileListHandler(request *restful.Request, response *restful.Response) {

}

func (rrs *ResponseResource) MyInfoHandler(request *restful.Request, response *restful.Response) {
	loginregisterbean := new(module.LoginRegisterBean)
	err := request.ReadEntity(&loginregisterbean)
	if err != nil {
		log.Println(err)
		response.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		response.WriteErrorString(401, "401: Not Authorized Parameter Input Error")
		return
	}

	bt := db.NewBoltDB(conf.BboltDBPath, "user")
	defer bt.Close()

	jsonstr := bt.Get(loginregisterbean.UserName)
	lsb := new(module.LoginRegisterBean)
	json.Unmarshal([]byte(jsonstr.(string)), &lsb)

	retlst := []module.LoginRegisterBean{}
	retlst = append(retlst, module.LoginRegisterBean{UserName: loginregisterbean.UserName, UserId: lsb.UserId, Alias: lsb.Alias, Mail: lsb.Mail, Password: lsb.Password})
	response.WriteEntity(retlst)

}

func (rrs *ResponseResource) ToolAddHandler(request *restful.Request, response *restful.Response) {

	tools := new(module.Tools)
	err := request.ReadEntity(&tools)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(http.StatusNotFound, "Parse json error.")
		return
	}
	u1 := uuid.Must(uuid.NewV4())
	tools.ToolsId = fmt.Sprint(u1)
	timeStr := time.Now().Format("2006-01-02")
	tools.Cdt = timeStr

	jsonstr, _ := json.Marshal(tools)
	bt := db.NewBoltDB(conf.BboltDBPath, "tools")
	defer bt.Close()

	bt.Set(tools.ToolsId+","+tools.Version, string(jsonstr))
	bt.Close()
	//fmt.Println(string(jsonstr))
	list := []module.RetBean{}
	list = append(list, module.RetBean{Code: "200", Status: "ok", Err: "", Data: "{}"})
	response.WriteEntity(list)
}

func (rrs *ResponseResource) ToolUpdateHandler(request *restful.Request, response *restful.Response) {

	tools := new(module.Tools)
	err := request.ReadEntity(&tools)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(http.StatusNotFound, "Parse json error.")
		return
	}

	jsonstr, _ := json.Marshal(tools)
	bt := db.NewBoltDB(conf.BboltDBPath, "tools")
	defer bt.Close()

	bt.Set(tools.ToolsId+","+tools.Version, string(jsonstr))
	bt.Close()
	//fmt.Println(string(jsonstr))
	list := []module.RetBean{}
	list = append(list, module.RetBean{Code: "200", Status: "ok", Err: "", Data: "{}"})
	response.WriteEntity(list)
}

func (rrs *ResponseResource) ToolRemoveHandler(request *restful.Request, response *restful.Response) {

	tools := new(module.Tools)
	err := request.ReadEntity(&tools)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(http.StatusNotFound, "Parse json error.")
		return
	}

	bt := db.NewBoltDB(conf.BboltDBPath, "tools")
	defer bt.Close()

	bt.Remove(tools.ToolsId + "," + tools.Version)
	bt.Close()
	//fmt.Println(string(jsonstr))
}

func (rrs *ResponseResource) UploadHandler(request *restful.Request, response *restful.Response) {
	reqParams, err := url.ParseQuery(request.Request.URL.RawQuery)
	if err != nil {
		response.WriteErrorString(401, "parse url parameter err.")
		log.Println(err)
		return
	}

	username, err := util.JwtAccessTokenUserName(fmt.Sprint(reqParams["accesstoken"][0]), conf.JwtKey)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(401, "accesstoken parse username err.")
		return
	}
	os.Mkdir(conf.UploadPath+"/"+username, 0777)

	req := request.Request
	mr, err := req.MultipartReader()
	if err != nil {
		log.Println(err)
		response.WriteErrorString(500, "Open file error.")
		return
	}

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			response.WriteErrorString(500, "Open file error.")
			return
		}
		log.Println(p.Header)
		formName := p.FormName()
		fileName := p.FileName()
		if formName != "" && fileName == "" {
			formValue, _ := ioutil.ReadAll(p)
			log.Println(formValue)
		}
		if fileName != "" {
			fileData, _ := ioutil.ReadAll(p)
			err = ioutil.WriteFile(conf.UploadPath+"/"+username+"/"+fileName, fileData, 0666)
			if err != nil {
				log.Println(err)
				response.WriteErrorString(500, "Open file error.")
				return
			}
		}
	}
	list := []module.RetBean{}
	list = append(list, module.RetBean{Code: "200", Status: "ok", Err: "", Data: "{}"})
	response.WriteEntity(list)
}

func (rrs *ResponseResource) UploadFileListHandler(request *restful.Request, response *restful.Response) {
	reqParams, err := url.ParseQuery(request.Request.URL.RawQuery)
	if err != nil {
		response.WriteErrorString(401, "parse url parameter err.")
		log.Println(err)
		return
	}

	username, err := util.JwtAccessTokenUserName(fmt.Sprint(reqParams["accesstoken"][0]), conf.JwtKey)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(401, "accesstoken parse username err.")
		return
	}
	retlst := []module.FileInfoBean{}
	files, _ := ioutil.ReadDir(conf.UploadPath + "/" + username)
	for _, f := range files {
		finfo, _ := os.Stat(conf.UploadPath + "/" + username + "/" + f.Name())
		fname := f.Name()
		//fctime:= time.Now().UTC().Format("2006-01-02 15:04:05")
		fisdir := "-"
		if finfo.IsDir() {
			fisdir = "d"
		}
		fsize := strconv.FormatInt(finfo.Size(), 10)
		fctime := finfo.ModTime().Format("2006-01-02 15:04:05")
		furl := conf.UploadRmtUrl + "/" + username + "/" + f.Name()

		retlst = append(retlst, module.FileInfoBean{Name: fname, Size: fsize, Cdt: fctime, IsDir: fisdir, Url: furl})
	}
	response.WriteEntity(retlst)
}

func (rrs *ResponseResource) UploadFileDeleteHandler(request *restful.Request, response *restful.Response) {
	reqParams, err := url.ParseQuery(request.Request.URL.RawQuery)
	if err != nil {
		response.WriteErrorString(401, "parse url parameter err.")
		log.Println(err)
		return
	}
	finfo := new(module.FileInfoBean)
	err = request.ReadEntity(&finfo)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(http.StatusNotFound, "Parse json error.")
		return
	}

	username, err := util.JwtAccessTokenUserName(fmt.Sprint(reqParams["accesstoken"][0]), conf.JwtKey)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(401, "accesstoken parse username err.")
		return
	}
	err = os.Remove(conf.UploadPath + "/" + username + "/" + finfo.Name)
	if err != nil {
		response.WriteErrorString(401, "delete file "+conf.UploadPath+"/"+username+"/"+finfo.Name+" err.")
	}
	list := []module.RetBean{}
	list = append(list, module.RetBean{Code: "200", Status: "ok", Err: "", Data: "{}"})
	response.WriteEntity(list)
}

func (rrs *ResponseResource) CalLeastsqHandler(request *restful.Request, response *restful.Response) {
	reqParams, err := url.ParseQuery(request.Request.URL.RawQuery)
	if err != nil {
		response.WriteErrorString(401, "parse url parameter err.")
		log.Println(err)
		return
	}

	leastsq := new(module.LeastsqBean)
	err = request.ReadEntity(&leastsq)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(http.StatusNotFound, "Parse json error.")
		return
	}

	username, err := util.JwtAccessTokenUserName(fmt.Sprint(reqParams["accesstoken"][0]), conf.JwtKey)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(401, "accesstoken parse username err.")
		return
	}
	os.Mkdir(conf.UploadPath+"/"+username, 0777)

	u1 := uuid.Must(uuid.NewV4())
	pngid := fmt.Sprint(u1)

	cmd := "python " + conf.LeastsqScript + " -d \"" + leastsq.CalVal + "\" -v \"" + leastsq.RealVal + "\" -j \"" + leastsq.Title + "\" -p \"" + conf.UploadPath + "/" + username + "/" + pngid + ".png\""
	log.Println(cmd)
	scret, err := ex.NewExec().Execmd(cmd)
	scret = strings.Replace(scret, "\n", "", -1)

	//reg := regexp.MustCompile(conf.LeastsqRegExp)
	//regret := reg.FindStringSubmatch(scret)[1]
	log.Println(scret)
	arr := strings.Split(scret, ",")
	if len(arr) != 7 {
		response.WriteErrorString(401, "leastsq script cal err.")
		return
	}

	leastsqlst := []module.LeastsqBean{}
	leastsqlst = append(leastsqlst, module.LeastsqBean{Aval: arr[0], Bval: arr[1], Rate: arr[2], Subv: arr[3], PredictVal: arr[4], RealSeq: arr[5], Url: conf.UploadRmtUrl + "/" + username + "/" + pngid + ".png"})
	response.WriteEntity(leastsqlst)
}

func (rrs *ResponseResource) CalFitLineHandler(request *restful.Request, response *restful.Response) {
	reqParams, err := url.ParseQuery(request.Request.URL.RawQuery)
	if err != nil {
		response.WriteErrorString(401, "parse url parameter err.")
		log.Println(err)
		return
	}

	fitline := new(module.FitLineBean)
	err = request.ReadEntity(&fitline)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(http.StatusNotFound, "Parse json error.")
		return
	}

	username, err := util.JwtAccessTokenUserName(fmt.Sprint(reqParams["accesstoken"][0]), conf.JwtKey)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(401, "accesstoken parse username err.")
		return
	}
	os.Mkdir(conf.UploadPath+"/"+username, 0777)

	u1 := uuid.Must(uuid.NewV4())
	pngid := fmt.Sprint(u1)

	cmd := "python " + conf.FitScript + " -x \"" + fitline.XVal + "\" -y \"" + fitline.YVal + "\" -t \"" + fitline.Title + "\" -p \"" + conf.UploadPath + "/" + username + "/" + pngid + ".png\" -m \"" + fitline.OptTime + "\" -s \"" + fitline.Loss + "\""
	log.Println(cmd)
	scret, err := ex.NewExec().Execmd(cmd)
	scret = strings.Replace(scret, "\n", "", -1)

	reg := regexp.MustCompile(conf.FitGrepExp)
	regret := reg.FindStringSubmatch(scret)
	log.Println(regret[3])
	log.Println(len(regret))
	if len(regret) != 4 {
		response.WriteErrorString(401, "fit line script cal err.")
		return
	}

	fitLinelst := []module.FitLineBean{}
	fitLinelst = append(fitLinelst, module.FitLineBean{Aval: fmt.Sprint(regret[1]), Bval: fmt.Sprint(regret[2]), Loss: fmt.Sprint(regret[3]), Url: conf.UploadRmtUrl + "/" + username + "/" + pngid + ".png"})
	response.WriteEntity(fitLinelst)
}

func (rrs *ResponseResource) CalCorrHandler(request *restful.Request, response *restful.Response) {
	reqParams, err := url.ParseQuery(request.Request.URL.RawQuery)
	if err != nil {
		response.WriteErrorString(401, "parse url parameter err.")
		log.Println(err)
		return
	}

	corrbean := new(module.CorrBean)
	err = request.ReadEntity(&corrbean)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(http.StatusNotFound, "Parse json error.")
		return
	}

	username, err := util.JwtAccessTokenUserName(fmt.Sprint(reqParams["accesstoken"][0]), conf.JwtKey)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(401, "accesstoken parse username err.")
		return
	}
	os.Mkdir(conf.UploadPath+"/"+username, 0777)

	//u1 := uuid.Must(uuid.NewV4())
	//pngid := fmt.Sprint(u1)

	cmd := "python " + conf.CorrScript + " -x \"" + corrbean.XVal + "\" -y \"" + corrbean.YVal + "\" "
	log.Println(cmd)
	scret, err := ex.NewExec().Execmd(cmd)
	scret = strings.Replace(scret, "\n", "", -1)

	reg := regexp.MustCompile(conf.CorrGrepExp)
	regret := reg.FindStringSubmatch(scret)
	log.Println(len(regret))
	if len(regret) != 2 {
		response.WriteErrorString(401, "fit line script cal err.")
		return
	}

	corrlst := []module.CorrBean{}
	corrlst = append(corrlst, module.CorrBean{Corr: fmt.Sprint(regret[1])})
	response.WriteEntity(corrlst)
}

func (rrs *ResponseResource) FieldAddHandler(request *restful.Request, response *restful.Response) {
	reqParams, err := url.ParseQuery(request.Request.URL.RawQuery)
	if err != nil {
		response.WriteErrorString(401, "parse url parameter err.")
		log.Println(err)
		return
	}

	cn2en := new(module.Cn2EnFieldBean)
	err = request.ReadEntity(&cn2en)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(http.StatusNotFound, "Parse json error.")
		return
	}

	username, err := util.JwtAccessTokenUserName(fmt.Sprint(reqParams["accesstoken"][0]), conf.JwtKey)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(401, "accesstoken parse username err.")
		return
	}

	timeStr := time.Now().Format("2006-01-02 15:04:05")
	cn2en.Cts = timeStr
        cn2en.Contributor = username

	jsonstr, _ := json.Marshal(cn2en)
	bt := db.NewBoltDB(conf.BboltDBPath, "cn2en")
	defer bt.Close()

	err = bt.Set(cn2en.ENName+","+cn2en.CNName, string(jsonstr))
        if err != nil {
                log.Println(err)
                response.WriteErrorString(401, "data in db update error.")
                return
        }
        log.Println(string(jsonstr))
	list := []module.RetBean{}
	list = append(list, module.RetBean{Code: "200", Status: "ok", Err: "", Data: "{}"})
	response.WriteEntity(list)
}

func (rrs *ResponseResource) FieldRemoveHandler(request *restful.Request, response *restful.Response) {

	cn2en := new(module.Cn2EnFieldBean)
	err := request.ReadEntity(&cn2en)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(http.StatusNotFound, "Parse json error.")
		return
	}

	bt := db.NewBoltDB(conf.BboltDBPath, "cn2en")
	defer bt.Close()

	err=bt.Remove(cn2en.ENName + "," + cn2en.CNName)
        if err != nil {
                log.Println(err)
                response.WriteErrorString(401, "data in db update error.")
                return
        }
        list := []module.RetBean{}
        list = append(list, module.RetBean{Code: "200", Status: "ok", Err: "", Data: "{}"})
        response.WriteEntity(list)
}

func (rrs *ResponseResource) FieldQryHandler(request *restful.Request, response *restful.Response) {
	cn2enqry := new(module.Cn2EnFieldQryBean)
	err := request.ReadEntity(&cn2enqry)
	if err != nil {
		log.Println(err)
		response.WriteErrorString(http.StatusNotFound, "Parse json error.")
		return
	}
        log.Println(cn2enqry.Rege)
	bt := db.NewBoltDB(conf.BboltDBPath, "cn2en")
	defer bt.Close()

	strlist := bt.Scan()
	retlst := make([]interface{}, 0)
	for _, v := range strlist {
		for k1, v1 := range v.(map[string]interface{}) {
			reg := regexp.MustCompile(cn2enqry.Rege)
			sstr := reg.FindAllString(k1, -1)
			if len(sstr) == 0 {
				continue
			}
			cn2en := new(module.Cn2EnFieldBean)
			err := json.Unmarshal([]byte(v1.(string)), &cn2en)
			if err != nil {
				log.Println(err)
				continue
			}
                        cn2en.Key = cn2en.ENName+","+cn2en.CNName
			retlst = append(retlst, cn2en)
		}
	}
        //log.Println(retlst)
	response.WriteEntity(retlst)
}

func (rrs *ResponseResource) BuyAddProductHandler(request *restful.Request, response *restful.Response) {
        reqParams, err := url.ParseQuery(request.Request.URL.RawQuery)
        if err != nil {
                response.WriteErrorString(401, "parse url parameter err.")
                log.Println(err)
                return
        }
		
        username, err := util.JwtAccessTokenUserName(fmt.Sprint(reqParams["accesstoken"][0]), conf.JwtKey)
        if err != nil {
                log.Println(err)
                response.WriteErrorString(401, "accesstoken parse username err.")
                return
        }

        buyproduct := new(module.BuyProductBean)
        err = request.ReadEntity(&buyproduct)
        if err != nil {
                log.Println(err)
                response.WriteErrorString(http.StatusNotFound, "Parse json error.")
                return
        }

        u1 := uuid.Must(uuid.NewV4())
        productid := fmt.Sprint(u1)

        bt := db.NewBoltDB(conf.BboltDBPath, "product")
        defer bt.Close()
        
        timeStr := time.Now().Format("2006-01-02 15:04:05")
        buyproduct.Cts = timeStr
        buyproduct.Contributor = username
        buyproduct.ProductId = productid

        jsonstr, _ := json.Marshal(buyproduct)

        err = bt.Set(buyproduct.ProductId, string(jsonstr))
        if err != nil {
                log.Println(err)
                response.WriteErrorString(401, "data in db update error.")
                return
        }

        //log.Println(retlst)
        list := []module.RetBean{}
        list = append(list, module.RetBean{Code: "200", Status: "ok", Err: "", Data: "{}"})
        response.WriteEntity(list)
}

func (rrs *ResponseResource) BuyGetProductHandler(request *restful.Request, response *restful.Response) {
        buyproduct := new(module.BuyProductBean)
        err := request.ReadEntity(&buyproduct)
        if err != nil {
                log.Println(err)
                response.WriteErrorString(http.StatusNotFound, "Parse json error.")
                return
        }
        bt := db.NewBoltDB(conf.BboltDBPath, "product")
        defer bt.Close()

        retlst := make([]interface{}, 0)
        //log.Println(retlst)
        bp := bt.Get(buyproduct.ProductId)
        if bp != nil {
            buyproduct1 := new(module.BuyProductBean)
            err := json.Unmarshal([]byte(bp.(string)), &buyproduct1)
            if err != nil {
               log.Println(err)
            }
            retlst = append(retlst, buyproduct1)
        }
        response.WriteEntity(retlst)
}

func (rrs *ResponseResource) AudioAddHandler(request *restful.Request, response *restful.Response) {
        reqParams, err := url.ParseQuery(request.Request.URL.RawQuery)
        if err != nil {
                response.WriteErrorString(401, "parse url parameter err.")
                log.Println(err)
                return
        }

        username, err := util.JwtAccessTokenUserName(fmt.Sprint(reqParams["accesstoken"][0]), conf.JwtKey)
        if err != nil {
                log.Println(err)
                response.WriteErrorString(401, "accesstoken parse username err.")
                return
        }

        audiobean := new(module.AudioBean)
        err = request.ReadEntity(&audiobean)
        if err != nil {
                log.Println(err)
                response.WriteErrorString(http.StatusNotFound, "Parse json error.")
                return
        }
        bt := db.NewBoltDB(conf.BboltDBPath, "audio_catalog")
        defer bt.Close()

        timeStr := time.Now().Format("2006-01-02 15:04:05")
        audiobean.Udt = timeStr
        audiobean.SubmitUser = username

        jsonstr, _ := json.Marshal(audiobean)
        err = bt.Set(audiobean.Catalog+"/"+audiobean.Name, string(jsonstr))
        if err != nil {
                log.Println(err)
                response.WriteErrorString(401, "data in db update error.")
                return
        }

        //log.Println(retlst)
        list := []module.RetBean{}
        list = append(list, module.RetBean{Code: "200", Status: "ok", Err: "", Data: "{}"})
        response.WriteEntity(list)       
}

func (rrs *ResponseResource) AudioRemoveHandler(request *restful.Request, response *restful.Response) {
        audiobean := new(module.AudioBean)
        err := request.ReadEntity(&audiobean)
        if err != nil {
                log.Println(err)
                response.WriteErrorString(http.StatusNotFound, "Parse json error.")
                return
        }
        bt := db.NewBoltDB(conf.BboltDBPath, "audio_catalog")
        defer bt.Close()

        err=bt.Remove(audiobean.Catalog+"/"+audiobean.Name)
        if err != nil {
                log.Println(err)
                response.WriteErrorString(401, "data in db update error.")
                return
        }
        list := []module.RetBean{}
        list = append(list, module.RetBean{Code: "200", Status: "ok", Err: "", Data: "{}"})
        response.WriteEntity(list)

}

func (rrs *ResponseResource) AudioGetHandler(request *restful.Request, response *restful.Response) {
        audiobean := new(module.AudioBean)
        err := request.ReadEntity(&audiobean)
        if err != nil {
                log.Println(err)
                response.WriteErrorString(http.StatusNotFound, "Parse json error.")
                return
        }
        bt := db.NewBoltDB(conf.BboltDBPath, "audio_catalog")
        defer bt.Close()
 
        retlst := make([]interface{}, 0)
        ab := bt.Get(audiobean.Catalog+"/"+audiobean.Name)
        if ab != nil {
            ab1 := new(module.AudioBean)
            err = json.Unmarshal([]byte(ab.(string)), &ab1)
            if err != nil {
               log.Println(err)
            }
            retlst = append(retlst, ab1)
        }
        response.WriteEntity(retlst)

}

func (rrs *ResponseResource) AudioCatalogListHandler(request *restful.Request, response *restful.Response) {
        audiobean := new(module.AudioBean)
        err := request.ReadEntity(&audiobean)
        if err != nil {
                log.Println(err)
                response.WriteErrorString(http.StatusNotFound, "Parse json error.")
                return
        }
        bt := db.NewBoltDB(conf.BboltDBPath, "audio_catalog")
        defer bt.Close()

        strlist := bt.Scan()
        retlst := make([]interface{}, 0)
        for _, v := range strlist {
            for k1, v1 := range v.(map[string]interface{}) {
              arr := strings.Split(k1, "/")
              if arr[0] != audiobean.Catalog {
                 continue
              }
              ab_v := new(module.AudioBean)
              err = json.Unmarshal([]byte(v1.(string)), &ab_v)
              if err != nil {
                   log.Println(err)
                   continue
              }
              retlst = append(retlst, ab_v)
            }
        }
        //log.Println(retlst)
        response.WriteEntity(retlst)

}

func (rrs *ResponseResource) AudioSearchHandler(request *restful.Request, response *restful.Response) {
        audiobean := new(module.AudioBean)
        err := request.ReadEntity(&audiobean)
        if err != nil {
                log.Println(err)
                response.WriteErrorString(http.StatusNotFound, "Parse json error.")
                return
        }
        bt := db.NewBoltDB(conf.BboltDBPath, "audio_catalog")
        defer bt.Close()

        strlist := bt.Scan()
        retlst := make([]interface{}, 0)
        for _, v := range strlist {
            for k1, v1 := range v.(map[string]interface{}) {
              reg := regexp.MustCompile(audiobean.Catalog)
              sstr := reg.FindAllString(k1, -1)
              if len(sstr) == 0 {
                  continue
              }
              ab_v := new(module.AudioBean)
              err = json.Unmarshal([]byte(v1.(string)), &ab_v)
              if err != nil {
                   log.Println(err)
                   continue
              }
              retlst = append(retlst, ab_v)
            }
        }
        //log.Println(retlst)
        response.WriteEntity(retlst)

}


var (
	cfg  = flag.String("conf", "conf.yaml", "basic config")
	conf *module.MetaConf
)

func main() {
	flag.Parse()
	conf = new(module.MetaConf)
	yamlFile, err := ioutil.ReadFile(*cfg)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}
	err = yaml.UnmarshalStrict(yamlFile, conf)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}

	restful.Filter(globalOauth)

	// Optionally, you may need to enable CORS for the UI to work.
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		CookiesAllowed: false,
		Container:      restful.DefaultContainer}
	restful.DefaultContainer.Filter(cors.Filter)

	rrs := ResponseResource{Data: make(map[string]interface{})}
	restful.DefaultContainer.Add(rrs.WebService())

	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/home/k8s/tomcat/swagger-ui-3.22.0/dist"))))

	log.Printf("Get the API using http://localhost:" + conf.Port + "/apidocs.json")
	log.Printf("Open Swagger UI using http://localhost:" + conf.Port + "/apidocs/?url=http://localhost:" + conf.Port + "/apidocs.json")
	log.Fatal(http.ListenAndServe(":"+conf.Port, nil))
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "UserService",
			Description: "Resource for managing Users",
			Contact: &spec.ContactInfo{
				Name:  "john",
				Email: "john@doe.rp",
				URL:   "http://johndoe.org",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "users",
		Description: "Managing users"}}}
}

// Global Filter
func globalOauth(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	u, err := url.Parse(req.Request.URL.String())
	if err != nil {
		log.Println("parse url error")
		return
	}
	if u.Path == "/api/login" || u.Path == "/api/register" {
		chain.ProcessFilter(req, resp)
		return
	}
	reqParams, err := url.ParseQuery(req.Request.URL.RawQuery)
	if err != nil {
		log.Printf("Failed to decode: %v", err)
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}
	tokenstring := fmt.Sprint(reqParams["accesstoken"][0])
	var claimsDecoded map[string]interface{}
	decodeErr := jwt.Decode([]byte(tokenstring), &claimsDecoded, []byte(conf.JwtKey))
	if decodeErr != nil {
		log.Printf("Failed to decode: %s (%s)", decodeErr, tokenstring)
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}

	exp := claimsDecoded["exp"].(float64)
	exp1, _ := strconv.ParseFloat(fmt.Sprintf("%v", time.Now().Unix()+0), 64)

	if (exp - exp1) < 0 {
		log.Printf("Failed to decode: %v %v %v", exp, exp1, (exp - exp1))
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401: Not Authorized AccessToken Expired ,Please login")
		return
	}
	//fmt.Println((exp - exp1))
	chain.ProcessFilter(req, resp)
}
