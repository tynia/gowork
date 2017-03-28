package libutil

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"mae_proj/MAE/common/logging"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var (
	ERR_CODE_SUCCESS      = 0
	ERR_CODE_OK           = 0
	ERR_CODE_PARA_ERROR   = 1001
	ERR_CODE_DB_ERROR     = 1002
	ERR_CODE_SYSTEM_ERROR = 1003
	ERR_CODE_AUTH_ERROR   = 1004

	ERR_STR_SUCCESS      = "ok"
	ERR_STR_OK           = "success"
	ERR_STR_PARA_ERROR   = "para error"
	ERR_STR_DB_ERROR     = "db error"
	ERR_STR_SYSTEM_ERROR = "system error"
	ERR_STR_AUTH_ERROR   = "auth error"

	// error code for project server
	ERR_CODE_PROJ_BEGIN = 10000

	// error code for image server
	ERR_CODE_IMAGE_BEGIN = 20000

	// error code for application server
	ERR_CODE_APPLICATION_BEGIN = 30000

	ERR_CODE_PROJ_PARA_NULL          = 10001
	ERR_CODE_PROJ_UPDATE_PARA_ERROR  = 10002
	ERR_CODE_PROJ_APPS_ARE_RUNNING   = 10003
	ERR_CODE_PROJ_GIT_PATH_INVALID   = 10004
	ERR_CODE_PROJ_PARSE_DATE_ERROR   = 10005
	ERR_CODE_PROJ_NOT_EXIST_ERROR    = 10006
	ERR_CODE_PROJ_NAME_ALREADY_EXIST = 10007
	ERR_CODE_PROJ_NO_RECORD_UPDATED  = 10008
	ERR_CODE_PROJ_LIST_ERROR         = 10009
	ERR_CODE_TAG_LIST_ERROR          = 10010

	ERR_STR_PROJ_PARA_NULL          = "parameter %s is null or empty"
	ERR_STR_PROJ_UPDATE_PARA_ERROR  = "Project update parameters error"
	ERR_STR_PROJ_APPS_ARE_RUNNING   = "one or more applications of the deleting project are running"
	ERR_STR_PROJ_GIT_PATH_INVALID   = "invalid git path %s"
	ERR_STR_PROJ_PARSE_DATE_ERROR   = "parse date error"
	ERR_STR_PROJ_NOT_EXIST_ERROR    = "%s not exist"
	ERR_STR_PROJ_NAME_ALREADY_EXIST = "project name already exist"
	ERR_STR_PROJ_NO_RECORD_UPDATED  = "no record updated"
	ERR_STR_PROJ_LIST_ERROR         = "get Project list error"
	ERR_STR_TAG_LIST_ERROR          = "get tag list error"

	ERR_CODE_IMAGE_QUERY_PROJ_ERROR            = 20001
	ERR_CODE_USER_USER_DEFINE_DOCKERFILE_ERROR = 20002
	ERR_CODE_USER_PROJ_INFO_FIELD_NULL_ERROR   = 20003
	ERR_CODE_PROJ_GIT_PATH_ERROR               = 20004
	ERR_CODE_IMAGE_IMG_NOT_EXIST               = 20005
	ERR_CODE_DOCKERFILE_TEMPLATE_URL_ERROR     = 20006
	ERR_CODE_IMAGE_QUERY_TEMPL_ERROR           = 20007
	ERR_CODE_PROJ_TEMPL_INFO_FIELD_NULL_ERROR  = 20008
	ERR_CODE_IMAGE_EXIST       			       = 20009

	ERR_STR_IMAGE_QUERY_PROJ_ERROR            = "user %s doesn't have project %s record"
	ERR_STR_USER_USER_DEFINE_DOCKERFILE_ERROR = "not support user define dockerfile now"
	ERR_STR_USER_PROJ_INFO_FIELD_NULL_ERROR   = "user project info field null error, git_path=%s, dockerfile_path=%s"
	ERR_STR_PROJ_GIT_PATH_ERROR               = "git path not match ssh format"
	ERR_STR_IMAGE_IMG_NOT_EXIST               = "image not exist"
	ERR_STR_IMAGE_HAS_FIRST_IMGINFO_ERROR     = "user %s has created first image with projectid=%s"
	ERR_STR_DOCKERFILE_TEMPLATE_URL_ERROR     = "dockerfile template url error: %s"
	ERR_STR_IMAGE_QUERY_TEMPL_ERROR           = "fail to query template info of project %s"
	ERR_STR_PROJ_TEMPL_INFO_FIELD_NULL_ERROR  = "project template info field null error, portList=%s"
	ERR_STR_IMAGE_EXIST						  = "image already exist"

	ERR_CODE_APPLICATION_INVALID_CMD = 30001
	ERR_STR_APPLICATION_INVALID_CMD  = "invalid cmd"

	ERR_CODE_APPLICATION_INVALID_APPID = 30002
	ERR_STR_APPLICATION_INVALID_APPID  = "invalid appid"

	ERR_CODE_APPLICATION_CREATE_SVC_FAILED = 30003
	ERR_STR_APPLICATION_CREATE_SVC_FAILED  = "create svc failed"

	ERR_CODE_APPLICATION_CREATE_RC_FAILED = 30004
	ERR_STR_APPLICATION_CREATE_RC_FAILED  = "create rc failed"

	ERR_CODE_APPLICATION_UPDATE_STATUS_FAILED = 30005
	ERR_STR_APPLICATION_UPDATE_STATUS_FAILED  = "update status failed"

	ERR_CODE_APPLICATION_DELETE_FAILED = 30006
	ERR_STR_APPLICATION_DELETE_FAILED  = "delete svc or rc failed"

	ERR_CODE_APPLICATION_SCALE_FAILED = 30007
	ERR_STR_APPLICATION_SCALE_FAILED  = "scale  rc failed"

	ERR_CODE_APPLICATION_UPDATE_REPLICAS_FAILED = 30008
	ERR_STR_APPLICATION_UPDATE_REPLICAS_FAILED  = "update replicas failed"

	ERR_CODE_APPLICATION_GET_IMAGE_INFO_FAILED = 30009
	ERR_STR_APPLICATION_GET_IMAGE_INFO_FAILED  = "get image info failed"

	ERR_CODE_APPLICATION_DELETE_APP_FAILED = 30010
	ERR_STR_APPLICATION_DELETE_APP_FAILED  = "delete app failed"

	ERR_CODE_APPLICATION_IMAGEID_EXIST = 30011
	ERR_STR_APPLICATION_IMAGEID_EXIST  = "imageid already exist"

	ERR_CODE_APPLICATION_APPNAME_EXIST = 30012
	ERR_STR_APPLICATION_APPNAME_EXIST  = "app name already exist"

	ERR_CODE_APPLICATION_ROLLING_UPDATING = 30013
	ERR_STR_APPLICATION_ROLLING_UPDATING  = "app is rolling updating"

	ERR_CODE_APPLICATION_ROLLING_UPDATE_COMMIT = 30014
	ERR_STR_APPLICATION_ROLLING_UPDATE_COMMIT  = "rolling update commit"

	ERR_CODE_APPLICATION_INVALID_IMAGE_PORT = 30015
	ERR_STR_APPLICATION_INVALID_IMAGE_PORT  = "invalid image port"

	ERR_CODE_APPLICATION_CREATE_TEMPLATE_FAILED = 30016
	ERR_STR_APPLICATION_CREATE_TEMPLATE_FAILED  = "create template failed"

	ERR_CODE_APPLICATION_INVALID_IMAGE_STATUS = 30017
	ERR_STR_APPLICATION_INVALID_IMAGE_STATUS  = "invalid image port"

	ERR_CODE_APPLICATION_ALREADY_RUNNING = 30018
	ERR_STR_APPLICATION_ALREADY_RUNNING  = "app already running"

	//CMQ错误码
	ERR_CODE_AUTH_PARA_NULL = 1000001
	ERR_STR_AUTH_PARA_NULL  = "parameter %s is null or empty"

	ErrMap = make(map[int]string)
)

//初始化包内的变量，自动执行
func init() {
	ErrMap[ERR_CODE_OK] = ERR_STR_OK
	ErrMap[ERR_CODE_PARA_ERROR] = ERR_STR_PARA_ERROR
	ErrMap[ERR_CODE_DB_ERROR] = ERR_STR_DB_ERROR
	ErrMap[ERR_CODE_SYSTEM_ERROR] = ERR_STR_SYSTEM_ERROR
	ErrMap[ERR_CODE_AUTH_ERROR] = ERR_STR_AUTH_ERROR
	ErrMap[ERR_CODE_SUCCESS] = ERR_STR_SUCCESS

	ErrMap[ERR_CODE_PROJ_PARA_NULL] = ERR_STR_PROJ_PARA_NULL
	ErrMap[ERR_CODE_PROJ_UPDATE_PARA_ERROR] = ERR_STR_PROJ_UPDATE_PARA_ERROR
	ErrMap[ERR_CODE_PROJ_APPS_ARE_RUNNING] = ERR_STR_PROJ_APPS_ARE_RUNNING
	ErrMap[ERR_CODE_PROJ_GIT_PATH_INVALID] = ERR_STR_PROJ_GIT_PATH_INVALID
	ErrMap[ERR_CODE_PROJ_PARSE_DATE_ERROR] = ERR_STR_PROJ_PARSE_DATE_ERROR
	ErrMap[ERR_CODE_PROJ_NOT_EXIST_ERROR] = ERR_STR_PROJ_NOT_EXIST_ERROR
	ErrMap[ERR_CODE_PROJ_NAME_ALREADY_EXIST] = ERR_STR_PROJ_NAME_ALREADY_EXIST
	ErrMap[ERR_CODE_PROJ_NO_RECORD_UPDATED] = ERR_STR_PROJ_NO_RECORD_UPDATED
	ErrMap[ERR_CODE_PROJ_LIST_ERROR] = ERR_STR_PROJ_LIST_ERROR
	ErrMap[ERR_CODE_TAG_LIST_ERROR] = ERR_STR_TAG_LIST_ERROR

	ErrMap[ERR_CODE_IMAGE_QUERY_PROJ_ERROR] = ERR_STR_IMAGE_QUERY_PROJ_ERROR
	ErrMap[ERR_CODE_USER_USER_DEFINE_DOCKERFILE_ERROR] = ERR_STR_USER_USER_DEFINE_DOCKERFILE_ERROR
	ErrMap[ERR_CODE_USER_PROJ_INFO_FIELD_NULL_ERROR] = ERR_STR_USER_PROJ_INFO_FIELD_NULL_ERROR
	ErrMap[ERR_CODE_PROJ_GIT_PATH_ERROR] = ERR_STR_PROJ_GIT_PATH_ERROR
	ErrMap[ERR_CODE_IMAGE_IMG_NOT_EXIST] = ERR_STR_IMAGE_IMG_NOT_EXIST
	ErrMap[ERR_CODE_DOCKERFILE_TEMPLATE_URL_ERROR] = ERR_STR_DOCKERFILE_TEMPLATE_URL_ERROR
	ErrMap[ERR_CODE_IMAGE_QUERY_TEMPL_ERROR] = ERR_STR_IMAGE_QUERY_TEMPL_ERROR
	ErrMap[ERR_CODE_PROJ_TEMPL_INFO_FIELD_NULL_ERROR] = ERR_STR_PROJ_TEMPL_INFO_FIELD_NULL_ERROR
	ErrMap[ERR_CODE_IMAGE_EXIST] = ERR_STR_IMAGE_EXIST

	ErrMap[ERR_CODE_APPLICATION_INVALID_CMD] = ERR_STR_APPLICATION_INVALID_CMD
	ErrMap[ERR_CODE_APPLICATION_INVALID_APPID] = ERR_STR_APPLICATION_INVALID_APPID
	ErrMap[ERR_CODE_APPLICATION_CREATE_SVC_FAILED] = ERR_STR_APPLICATION_CREATE_SVC_FAILED
	ErrMap[ERR_CODE_APPLICATION_CREATE_RC_FAILED] = ERR_STR_APPLICATION_CREATE_RC_FAILED
	ErrMap[ERR_CODE_APPLICATION_UPDATE_STATUS_FAILED] = ERR_STR_APPLICATION_UPDATE_STATUS_FAILED
	ErrMap[ERR_CODE_APPLICATION_DELETE_FAILED] = ERR_STR_APPLICATION_DELETE_FAILED
	ErrMap[ERR_CODE_APPLICATION_SCALE_FAILED] = ERR_STR_APPLICATION_SCALE_FAILED
	ErrMap[ERR_CODE_APPLICATION_UPDATE_REPLICAS_FAILED] = ERR_STR_APPLICATION_UPDATE_REPLICAS_FAILED
	ErrMap[ERR_CODE_APPLICATION_GET_IMAGE_INFO_FAILED] = ERR_STR_APPLICATION_GET_IMAGE_INFO_FAILED
	ErrMap[ERR_CODE_APPLICATION_DELETE_APP_FAILED] = ERR_STR_APPLICATION_DELETE_APP_FAILED
	ErrMap[ERR_CODE_APPLICATION_IMAGEID_EXIST] = ERR_STR_APPLICATION_IMAGEID_EXIST
	ErrMap[ERR_CODE_APPLICATION_APPNAME_EXIST] = ERR_STR_APPLICATION_APPNAME_EXIST
	ErrMap[ERR_CODE_APPLICATION_ROLLING_UPDATING] = ERR_STR_APPLICATION_ROLLING_UPDATING
	ErrMap[ERR_CODE_APPLICATION_ROLLING_UPDATE_COMMIT] = ERR_STR_APPLICATION_ROLLING_UPDATE_COMMIT
	ErrMap[ERR_CODE_APPLICATION_INVALID_IMAGE_PORT] = ERR_STR_APPLICATION_INVALID_IMAGE_PORT
	ErrMap[ERR_CODE_APPLICATION_CREATE_TEMPLATE_FAILED] = ERR_STR_APPLICATION_CREATE_TEMPLATE_FAILED
	ErrMap[ERR_CODE_APPLICATION_INVALID_IMAGE_STATUS] = ERR_STR_APPLICATION_INVALID_IMAGE_STATUS
	ErrMap[ERR_CODE_APPLICATION_ALREADY_RUNNING] = ERR_STR_APPLICATION_ALREADY_RUNNING

	ErrMap[ERR_CODE_AUTH_PARA_NULL] = ERR_STR_AUTH_PARA_NULL
}

//按统一格式生成http回复
//{“code”:0,”msg”:”ok”,”data”:”” }
//参数说明：
//msg支持变参
//data支持任何对象，注意如果data是struct，struct内部只有公有成员能正常Marshal,若json的key首字母小写需使用tag来解决
type HttpResponseData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type LogRequestInfo struct {
	Method       string
	URL          string
	User_id      string
	X_Auth_Token string
	Form         interface{}
	Body         interface{}
}

func LogGetResponseString(req *http.Request, code int, data interface{}, msg ...interface{}) []byte {
	ret := GetResponseString(code, data, msg...)
	defer req.Body.Close()
	userid := ""
	userid_cookie, _err := req.Cookie("user_id")
	if _err != nil || userid_cookie == nil {
		//logging.Error("LogGetResponseString get userid failed,request=%+v", req)

	} else {
		userid = userid_cookie.Value
	}

	body := ""
	b_body, err := ioutil.ReadAll(req.Body)
	if err == nil {
		body = string(b_body)
	}

	logReq := LogRequestInfo{
		Method:       req.Method,
		URL:          req.RequestURI,
		User_id:      userid,
		X_Auth_Token: req.Header.Get("X-Auth-Token"),
		Form:         req.Form,
		Body:         body,
	}
	logging.Info("HANDLE_LOG:request=%+v,response=%s", logReq, string(ret))
	return ret
}

func GetResponseString(code int, data interface{}, msg ...interface{}) []byte {
	response := &HttpResponseData{
		Code: code,
		Data: data,
	}

	msgfmt, exist := ErrMap[code]
	if !exist { //找不到说明代码有错误，panic
		logging.Error("GetResponseString failed errorcode:%d not exist", code)
		panic("GetResponseString failed errorcode:" + strconv.Itoa(code) + " not exist")
	}
	var msginfo = msgfmt
	if len(msg) != 0 {

		msginfo = fmt.Sprintf(msgfmt, msg...)
	}
	response.Msg = msginfo

	ret, err := json.Marshal(response)
	if err != nil {
		logging.Error("GetResponseString failed,code=%d,data=%v,msg=%v", code, data, msg)
		panic("GetResponseString failed, " + "code=" + strconv.Itoa(code) + err.Error())
	}

	return ret
}

func ParseRequestToMap(req *http.Request) (map[string]interface{}, string, error) {
	//logging.Debug("req: %s", req)
	var bodybuffer [1000]byte
	jsonobjRequestMap := make(map[string]interface{})
	//var httppath string
	httppath := html.EscapeString(req.URL.Path)
	len, err := req.Body.Read(bodybuffer[0:1000])
	if len <= 0 {
		return jsonobjRequestMap, "", nil
	}
	if err != nil && err.Error() != "EOF" {
		logging.Error("req.Body.Read failed,path:%s,err:%s", httppath, err.Error())
		return jsonobjRequestMap, "", err
	}
	recvbody := string(bodybuffer[0:len])
	recvslice := bodybuffer[0:len]
	//logging.Debug("%s recv body %s,len=%d", httppath, recvbody, len)

	err = json.Unmarshal(recvslice, &jsonobjRequestMap)
	if err != nil {
		logging.Error("invalid body: %s", recvbody)
		return jsonobjRequestMap, recvbody, err
	}
	for key, value := range jsonobjRequestMap {
		if reflect.TypeOf(value).String() != "string" {
			logging.Error("invalid json value type,key:%s,value:%s", key, value)
			return make(map[string]interface{}), recvbody, nil
		}
	}
	req.Body = ioutil.NopCloser(strings.NewReader(recvbody))

	return jsonobjRequestMap, recvbody, nil

}

//写入信息
/*func LogGetResponseString(req *http.Request, stime time.Time, code int, data interface{}, msg ...interface{}) []byte {
	ret := GetResponseString(code, data, msg...)
	defer req.Body.Close()

	body := ""
	b_body, err := ioutil.ReadAll(req.Body)
	if err == nil {
		body = string(b_body)
	}

	logReq := LogRequestInfo{
		Method: req.Method,
		URL:    req.RequestURI,
		Form:   req.Form,
		Body:   body,
	}
	use_time := time.Now().Sub(stime)
	logging.Info("HANDLE_LOG:uri=%s,use_time=%v,request=%+v,response=%s", req.RequestURI, use_time, logReq, string(ret))
	return ret
}
*/
