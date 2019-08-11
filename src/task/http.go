package task

import (
    "net/http"
    "fmt"
    "bytes"
    "io/ioutil"
)

func HTTPSend(data []byte, settings map[string]interface{}) (map[string][]byte, error) {
    reader := bytes.NewReader(data)
    request, err := http.NewRequest("POST", settings["url"].(string), reader)
    if err != nil {
        fmt.Println(err.Error())
        return nil, err
    }
    request.Header.Set("Content-Type", "application/json;charset=UTF-8")
    client := http.Client{}
    resp, err := client.Do(request)
    if err != nil {
        fmt.Println(err.Error())
        return nil, err
    }
    respBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err.Error())
        return nil, err
    }

    fmt.Println(string(respBytes))
    return nil, nil
}
