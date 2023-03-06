package output

import (
	"encoding/json"
	"fmt"
)

func AnyToString(input any) string {
    output, err := json.Marshal(input)
    if err != nil {
        fmt.Println("Oh no, something went wrong!")
        return ""
    }
    return string(output)
}
