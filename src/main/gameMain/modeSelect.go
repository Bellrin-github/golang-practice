package gameMain

import (
    "../lib"
    "../const/mainTask"
    "strconv"
)

func ModeSelect() (int, mainTask.MainTask) {
    inputText, _ := lib.SystemMessage([]string{
        "5  : 5x5 モード",
        "6  : 6x6 モード",
        "7  : 7x7 モード",
        "8  : 8x8 モード",
        "9  : 9x9 モード",
        "10 : 10x10 モード",
        "",
        "99 : メニュー画面に戻る",
        "",
        "いずれかの番号を入力してENTERキーを押してください。",
    })

    i, err := strconv.Atoi(inputText)
    if err != nil || i < 5 || (i > 10 && i != 99) {
        return i, mainTask.MODE_SELECT
    }

    if i == 99 {
        return i, mainTask.MAIN_MENU
    }

    return i, mainTask.STAGE_SELECT
}
