package main

import (
	"fmt"

	"imp/addons/gitlab"
)

func main() {

}

func testGitlabGetProject() {
    fmt.Println(gitlab.GetProjectById(14355))
}

func testGitlabGetProjectHooks() {
    fmt.Println(gitlab.GetProjectHooksById(14355))
}

func testGitlabPutProjectHooks() {
    hook := gitlab.GetProjectHooksById(14355)[0]
    hook.Url = "https://ciao.test"
    hook.SslEnable = false
    gitlab.PutProjectHooksById(14355, hook)
    fmt.Println(gitlab.GetProjectHooksById(14355)[0])
}
