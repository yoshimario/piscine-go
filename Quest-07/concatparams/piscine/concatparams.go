package piscine

import "fmt"

func ConcatParams(args []string) string {
    var result string
    for i, arg := range args {
        if i == 0 {
            result = arg
        } else {
            result = fmt.Sprintf("%s\n%s", result, arg)
        }
    }
    return result
}
