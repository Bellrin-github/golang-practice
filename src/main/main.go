package main

import (
    "./lib"
    "./const/mainTask"
    "./gameMain"
    "./user"
    "fmt"
    "os"
)

var userSetting user.User
var task mainTask.MainTask
var gameMode int

func main() {
    err := lib.SetClearCommand()
    if err != nil {
        lib.ClearConsole()
        fmt.Println(err)
        os.Exit(0)
    }

    defer func() {
        lib.SystemMessage([]string{
            "プログラムを終了しました。",
        })
    }()

    userSetting = gameMain.Opening()

    task = mainTask.MAIN_MENU

    for ; task != mainTask.EXIT_GAME; {
        switch task {
            case mainTask.MAIN_MENU:
                task = gameMain.MainMenu()

            case mainTask.MODE_SELECT:
                gameMode, task = gameMain.ModeSelect()

            case mainTask.STAGE_SELECT:
                gameMode, task = gameMain.StageSelect()

            case mainTask.MANUAL:
                lib.SystemMessage([]string{
                    "未実装",
                })

            case mainTask.GAME_MAIN:
                lib.SystemMessage([]string{
                    "未実装",
                })

            default:
                lib.SystemMessage([]string{
                    "不正なタスク",
                })
                os.Exit(0)
        }
    }

    lib.SystemMessage([]string{
        "またお会いしましょう。",
    })
}
