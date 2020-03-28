package commontest

import (
    "testing"
    "encoding/json"
)

func TestJson2Map(t *testing.T) {
    jsonStr := `
    {
        "name":"liangyongxing",
        "age":12
    }
    `
    var mapResult map[string]interface{}
    //使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
    if err := json.Unmarshal([]byte(jsonStr), &mapResult); err != nil {
        t.Fatal(err)
    }
    t.Log(mapResult)
}

