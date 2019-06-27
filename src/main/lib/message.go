package lib

import (
    "bufio"
    "fmt"
    "os"
)

func GetUserImputString()  (string, error) {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()

    return scanner.Text(), nil
}

func GetUserImputStringln() (string, error) {
    text, err := GetUserImputString()
    fmt.Println("")

    return text, err
}

func SystemMessage(m []string) (string, error) {
    ClearConsole()

    var inputText string = ""
    var err error = nil

    cnt := len(m) - 1
    for key, value := range m {
        if key < cnt {
            fmt.Println(value)
            continue
        }

        inputText, err = InputWait(value)
    }

    return inputText, err
}

func InputWait(message string) (string, error) {
    fmt.Print(message)
    fmt.Println(" ▼")

    return GetUserImputString()
}
