package gameMain

import (
    "../lib"
    "../const/mainTask"
)

func MainMenu() mainTask.MainTask {
    for {
        inputText, _ := lib.SystemMessage([]string{
            "A : モード選択",
            "B : ルール説明",
            "C : ゲーム終了",
            "",
            "A, B, C いずれかを入力してENTERキーを押してください。",
        })

        switch inputText {
            case "A", "a":
                return mainTask.MODE_SELECT
            case "B", "b":
                return mainTask.MANUAL
            case "C", "c":
                return mainTask.EXIT_GAME
            default:
                // 処理無し
        }
    }
}
