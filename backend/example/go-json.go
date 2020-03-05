package main

import (
    "encoding/json"
    "fmt"
)

//把结构体都改小写
type User struct {
    UserName string `json:"user_name"` //json的tag标记
    Nickname string `json:"nickname"`
    Age      int
    Birthday string
    Sex      string
    Email    string
    Phone    string
}

func testStruct() {
    user1 := &User{
        UserName: "超哥",
        Nickname: "大头哥",
        Age:      18,
        Birthday: "2008/8/8",
        Sex:      "男",
        Email:    "mahuateng@qq.com",
        Phone:    "110",
    }

    //开始json序列化
    data, err := json.Marshal(user1)
    if err != nil {
        fmt.Printf("json.marshal failed,err:", err)
        return
    }
    fmt.Printf("%s\n", string(data))
}

func testInt() {
    var a = 18
    //开始json序列化
    data, err := json.Marshal(a)
    if err != nil {
        fmt.Printf("json.marshal failed,err:", err)
        return
    }
    fmt.Printf("%s\n", string(data))

}

func testMap() {
    var m map[string]interface{}     //声明map
    m = make(map[string]interface{}) //必须初始化map分配内存
    m["username"] = "user1"
    m["age"] = 18
    m["sex"] = "man"
    fmt.Println(m)
    data, err := json.Marshal(m)
    if err != nil {
        fmt.Printf("json.marshal failed,err:", err)
        return
    }
    fmt.Printf("%s\n", string(data))

}

func testSlice() {
    //定义一个slice，元素是map
    var m map[string]interface{}
    var s []map[string]interface{}
    m = make(map[string]interface{})
    m["username"] = "user1"
    m["age"] = 18
    m["sex"] = "man"
    s = append(s, m)
    m = make(map[string]interface{})
    m["username"]="user2"
    m["age"]=188
    m["sex"]="male"
    s=append(s,m)
    data, err := json.Marshal(s)
    if err != nil {
        fmt.Printf("json.marshal failed,err:", err)
        return
    }
    fmt.Printf("%s\n", string(data))

}
func main() {
    testStruct() //结构体的序列化
    testInt()//序列化数值
    testMap()//序列化map
    testSlice()//序列化切片
}
