package controller

import "strconv"
import "crypto/sha1"
import "encoding/hex"
import "github.com/boyxp/nova/router"
import "github.com/boyxp/nova/exception"
import "github.com/boyxp/nova/response"
import "github.com/boyxp/nova/time"

import "api/model"
import "api/common"

func init() {
   router.Register(&Url{})
}

type Url struct {
}

//生成短连接
func (U *Url) Gen(url string, _pwd string, _expire string) any {
	if len(url)<11 || url[0:4]!="http" {
		exception.Throw("url格式错误", 2001)
	}

	data := map[string]string{"url":url}

	if _pwd!="" {
		data["password"] = U.sha1(_pwd)
	}

	if _expire!="" {
		data["expire_at"] = time.Date("Y-m-d H:i:s", time.Strtotime(_expire))
	}

	uid  := model.Url.Insert(data)
	hash := common.Short{}.Hash(uid)

	model.Url.Where("id", strconv.Itoa(int(uid))).Update(map[string]string{"hash":hash})

	return hash
}

//转换为连接
func (U *Url) Open(hash string) {
	id := common.Short{}.Convert(hash)

	if id==0 {
		exception.Throw("连接格式错误", 2002)
	}

	info := model.Url.Where("id", strconv.Itoa(id)).Where("expire_at", ">", time.Date("Y-m-d H:i:s")).Find()
	if info==nil {
		exception.Throw("连接不存在或已过期", 2003)
	}

	(&response.Redirect{}).Render(info["url"])
}

func (U *Url) sha1(str string) string {
	m := sha1.New()
	m.Write([]byte(str))

	return hex.EncodeToString(m.Sum(nil))
}
