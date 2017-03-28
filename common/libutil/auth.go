package libutil

import (
	"mae_proj/MAE/common/logging"
	"mae_proj/MAE/common/openstackauth"
	"net/http"
)

//从http.Request获取userid和tokenid
func getHttpAuthInfo(req *http.Request, userid, token *string) bool {
	userid_cookie, _err := req.Cookie("user_id")
	if _err != nil {
		logging.Error("getUserAuthInfo get userid failed,request=%v", req)
		return false
	}
	*userid = userid_cookie.Value
	*token = req.Header.Get("X-Auth-Token")
	if *token == "" {
		logging.Error("get X-Auth-Token header failed,request=%v", req)
		return false
	}
	return true
}

//验证用户，返回userid，是否成功
func AuthUserToken(req *http.Request) (string, bool) {
	userid := ""
	token := ""
	if !getHttpAuthInfo(req, &userid, &token) {
		logging.Error("AuthUserToken getHttpAuthInfo from request failed,request=%v", req)
		return "", false
	}

	_, err := openstackauth.AuthClient.GetUser(userid, token, false)
	if err != nil {
		logging.Error("AuthUserToken openstackauth GetUser failed,request=%v", req)
		return "", false
	}
	logging.Debug("AuthUserToken success userid=%s,token=%s", userid, token)
	return userid, true
}

//验证用户，返回userinfo，是否成功
func AuthUserTokenWithUserinfo(req *http.Request) (openstackauth.CommonUserInfo, bool) {
	userid := ""
	token := ""
	if !getHttpAuthInfo(req, &userid, &token) {
		logging.Error("AuthUserToken getUserAuthInfo failed,request=%v", req)
		return openstackauth.CommonUserInfo{}, false
	}

	logging.Debug("AuthUserToken userid=%s,token=%s", userid, token)

	userinfo, err := openstackauth.AuthClient.GetUser(userid, token, true)
	if err != nil {
		logging.Error("AuthUserToken openstackauth GetUser failed,request=%v", req)
		return openstackauth.CommonUserInfo{}, false
	}

	logging.Debug("AuthUserTokenWithUserinfo ok, userinfo=%+v", userinfo)
	return userinfo, true
}

func AuthUserTokenWithUserinfo2(req *http.Request) (*openstackauth.CommonUserInfo, bool) {
	return openstackauth.AuthToken(req)
}
