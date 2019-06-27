package lib

import (
    "fmt"
    "os"
    "os/exec"
)

var clearCommandText string

func ClearConsole() {
    c := exec.Command(clearCommandText)
    c.Stdout = os.Stdout
    c.Run()
}

func SetClearCommand() error {
    var err error = nil

    err = exec.Command("cli").Run()
    if err == nil {
        clearCommandText = "cli"
        return nil
    }

    err = exec.Command("clear").Run()
    if err == nil {
        clearCommandText = "clear"
        return nil
    }

    return fmt.Errorf("Unsupported OS")
}
