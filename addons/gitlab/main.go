package gitlab

import (
	"encoding/json"
	"reflect"
	"regexp"
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

func GetProjectHookById(idProject int, idHook int) Hook {
    var hook Hook
    RequestHandler.Url = RequestHandler.Url + "/projects/" + strconv.Itoa(idProject) + "/hooks/" + strconv.Itoa(idHook)

    raw := http.GetRequest(RequestHandler)
    json.Unmarshal(raw, &hook)
    
    RequestHandler.Url = base_url + api_prefix
    return hook
}

func PutProjectHooksById(id int, hookId int, hook Hook) bool {
    RequestHandler.Url = RequestHandler.Url + "/projects/" + strconv.Itoa(id) + "/hooks/" + strconv.Itoa(hookId) 

    values := reflect.ValueOf(hook)
    types := values.Type()
    re := regexp.MustCompile(`json:\"([^ ]+)\"`)
    payload := make([][2]string, values.NumField())
    for i := 0; i < values.NumField(); i++ {
	match := re.FindStringSubmatch(string(types.Field(i).Tag))[1]
	payload[i][0] = match
	switch values.Field(i).Type().String() {
	case "int":
	    payload[i][1] = strconv.FormatInt(values.Field(i).Int(), 10)
	case "bool":
	    payload[i][1] = strconv.FormatBool(values.Field(i).Bool())
	case "string":
	    payload[i][1] = values.Field(i).String()
	}
    }

    valid := http.PutRequest(RequestHandler, payload)

    RequestHandler.Url = base_url + api_prefix
    return valid
}

func ChangeSettingById(id int, key string, value string) bool {
    RequestHandler.Url = RequestHandler.Url + "/projects/" + strconv.Itoa(id)
    payload :=  [][2]string{{key,value}}

    result := http.PutRequest(RequestHandler, payload)

    RequestHandler.Url = base_url + api_prefix
    return result
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
