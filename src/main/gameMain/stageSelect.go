package gameMain

import (
    "../lib"
    "../const/mainTask"
    "strconv"
)

func StageSelect() (int, mainTask.MainTask) {
    inputText, _ := lib.SystemMessage([]string{
        "1  2  3  4  5  6  7  8  9  10",
        "",
        "99 : モード選択画面に戻る",
        "",
        "いずれかの番号を入力してENTERキーを押してください。",
    })

    gameMode, err := strconv.Atoi(inputText)
    if err != nil || gameMode < 0 || (gameMode > 10 && gameMode != 99) {
        return gameMode, mainTask.STAGE_SELECT
    }

    if gameMode == 99 {
        return gameMode, mainTask.MODE_SELECT
    }

    return gameMode, mainTask.GAME_MAIN
}
