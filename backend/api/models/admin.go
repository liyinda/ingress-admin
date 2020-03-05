package models

import (
    "fmt"
    //orm "github.com/liyinda/viewdns/backend/api/database"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

// 用户登录 from JSON
type LoginJson struct {
        Username     string `form:"username" json:"username" xml:"username"  binding:"required"`
        Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

//管理员结构体
type Admin struct {
    Adminname string `yaml:"admin"`
    Password string `yaml:"password"`
}

//获取管理员密码
func (adm *Admin) GetPassword() (password string, err error) {
    //orm.GetEtcdkey(login_name)
    data, err := ioutil.ReadFile("config/admin.yaml")
    if err != nil {
        fmt.Print(err)
    }

    err = yaml.Unmarshal([]byte(data), &adm)
    if err != nil {
        fmt.Print(err)
    }
    //fmt.Printf("%v", adm.Password)
    return adm.Password, err

}
