package groupware

import (
	"encoding/json"
	"fmt"
	"github.com/mattermost/mattermost-server/mlog"
	"github.com/mattermost/mattermost-server/utils"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

const loginUrl = "https://api-gsimplexi.cafe24.com/groupware/api/new/GroupwareLogin.php"

type Response struct {
	Code string
	Data []ResponseData
}

type ResponseData struct {
	Userflag string
	Restflag string
	PwdChangeflag string
}

// 그룹웨어 로그인
func Login(empId string, passwd string) (bool, error) {
	resp, err := http.PostForm(loginUrl, url.Values{"userId": {"gsimplexi"}, "empId": {empId}, "passwd": {utils.HashMd5(passwd)}})
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	var response Response

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return false, err
	}

	mlog.Debug(fmt.Sprintf("Login result: %v", response))

	if response.Code != "0000" || response.Data[0].Userflag == "false" || response.Data[0].Restflag == "true" {
		err := errors.New("Login failed")
		return false, err
	}

	return true, nil
}
