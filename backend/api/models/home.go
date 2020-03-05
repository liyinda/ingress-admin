package models

import (
	"encoding/json"
	//"errors"
	//"fmt"
	//orm "github.com/liyinda/ingress-admin/backend/api/database"
)

type DomainName struct {
	Id         string `json:"id"`
	Domainname string `json:"domainname"`
	Status     string `json:"status"`
	Type       string `json:"type"`
	Ttl        string `json:"ttl"`
	Rdata      string `json:"rdata"`
}

type EtcdRdata struct {
	Host string `json:"host"`
	Ttl  int    `json:"ttl"`
}

//添加dns的a记录
/*func (a_records DNS_A) Dnsadd() (id int64, err error) {
    //etcdctl put

}
*/

//列出etcd中所有列信息
//etcdctl get /skydns --prefix --keys-only=true
//如果读取不到etcd信息返回异常?
func (domainnames *DomainName) DnsList() (domains []DomainName, err error) {

	s := `[{"id": "650000199402171139","domainname": "www.baidu.com","status": "deleted","type": "name","rdata": "10.10.10.10","ttl": "4316"},{"id": "650000199402171131","domainname": "sdfsdf","status": "deleted","type": "name","rdata": "2017-03-11 05:35:45","ttl": "4317"}]`
	//s, _ := orm.GetEtcdPrefix()

	err = json.Unmarshal([]byte(s), &domains)
	if err != nil {
		return
	}
	return

}

//func (etcdrdata *EtcdRdata) DnsAdd(key string, host string, ttl int) bool {
//	etcdrdata.Host = host
//	etcdrdata.Ttl = ttl
//
//	value, err := json.Marshal(etcdrdata)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println("----" + key)
//
//	//orm.PutEtcdKV(key, string(value))
//	return true
//}
//
//func (etcdrdata *EtcdRdata) DnsDel(key string) (res string, err error) {
//	//如果删除成功则值为1
//	if orm.DelEtcdKV(key) == 1 {
//		res = "deleted"
//		return res, nil
//	}
//	return res, errors.New("catnot deleted")
//}
