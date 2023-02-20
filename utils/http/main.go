package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func PutRequest(requestHandler RequestHandler, payload [][2]string) bool {
    form := url.Values{}
    for _, value := range payload {
        form.Add(value[0],value[1])
    }

    request, err := http.NewRequest(http.MethodPut, requestHandler.Url, strings.NewReader(form.Encode()))
    request.Header.Set(requestHandler.TokenHeader, requestHandler.Token)
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    if err != nil {
        fmt.Println("Oh no, something went wrong!")
        return false
    }

    response, err := http.DefaultClient.Do(request)
    if err != nil {
        fmt.Println("Oh no, something went wrong!")
        return false
    }
    defer response.Body.Close()

    return true
}

func GetRequest(requestHandler RequestHandler) []byte {
    request, err := http.NewRequest(http.MethodGet, requestHandler.Url, nil)
    request.Header.Set(requestHandler.TokenHeader, requestHandler.Token)
    if err != nil {
        fmt.Println("Oh no, something went wrong!")
        return nil
    }

    response, err := http.DefaultClient.Do(request)
    if err != nil {
        fmt.Println("Oh no, something went wrong!")
        return nil
    }

    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Oh no, something went wrong!")
        return nil
    }

    return body
}

func DeleteRequest(requestHandler RequestHandler) bool {
    request, err := http.NewRequest(http.MethodDelete, requestHandler.Url, nil)
    request.Header.Set(requestHandler.TokenHeader, requestHandler.Token)
    if err != nil {
        fmt.Println("Oh no, something went wrong!")
        return false
    }

    response, err := http.DefaultClient.Do(request)
    if err != nil {
        fmt.Println("Oh no, something went wrong!")
        return false
    }
    defer response.Body.Close()

    return true
}
