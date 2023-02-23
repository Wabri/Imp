package gitlab

import (
	"encoding/json"
	"strconv"

	"imp/utils/http"
)

var RequestHandler http.RequestHandler = http.RequestHandler{ 
    Url: base_url + api_prefix,
    TokenHeader: "PRIVATE-TOKEN",
    Token: token,
}

func GetProjectById(id int) Project {
    var project Project
    RequestHandler.Url = RequestHandler.Url + "/projects/" + strconv.Itoa(id)

    raw := http.GetRequest(RequestHandler)
    json.Unmarshal(raw, &project) 

    RequestHandler.Url = base_url + api_prefix
    return project
}

func GetProjectHooksById(id int) []Hook {
    var hooks []Hook
    RequestHandler.Url = RequestHandler.Url + "/projects/" + strconv.Itoa(id) + "/hooks"

    raw := http.GetRequest(RequestHandler)
    json.Unmarshal(raw, &hooks) 
    
    RequestHandler.Url = base_url + api_prefix
    return hooks
}

func PutProjectHooksById(id int, hook Hook) bool {
    RequestHandler.Url = RequestHandler.Url + "/projects/" + strconv.Itoa(id) + "/hooks/" + strconv.Itoa(hook.Id) 

    payload := make([][2]string, 2)
    payload[0][0] = "url"
    payload[0][1] = hook.Url
    payload[1][0] = "enable_ssl_verification"
    payload[1][1] = strconv.FormatBool(hook.SslEnable)
    valid := http.PutRequest(RequestHandler, payload)

    RequestHandler.Url = base_url + api_prefix
    return valid
}

func DeleteProjectHooksById(id int, hook_id int) bool {
    RequestHandler.Url = RequestHandler.Url + "/projects/" + strconv.Itoa(id) + "/hooks/" + strconv.Itoa(hook_id)

    result := http.DeleteRequest(RequestHandler)
    
    RequestHandler.Url = base_url + api_prefix
    return result
}

func SearchOnProjectById(id int, search string) []SearchResult{
    var searchResult []SearchResult
    RequestHandler.Url = RequestHandler.Url + "/projects/" + strconv.Itoa(id) + "/search?scope=blobs&search=" + search

    raw := http.GetRequest(RequestHandler)
    json.Unmarshal(raw, &searchResult) 

    RequestHandler.Url = base_url + api_prefix

    return searchResult
}
