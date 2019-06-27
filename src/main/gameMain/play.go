package gameMain

import (
    "../const/mainTask"
    "../const/playTask"
    "../const/mark"
    "fmt"
    "../lib"
    "../question"
    "strconv"
    "strings"
)

var task playTask.PlayTask
var stageArray [][]int

func Play(gameMode int, gameStage int) mainTask.MainTask {
    task = playTask.INIT

    for task != playTask.GAME_OVER {
        switch task {
            case playTask.INIT:
                task = createStage(gameMode, gameStage)

            case playTask.ACTION_SELECT:
                task = actionSelect()

            case playTask.SEE_NORTH:
                task = seeNorth()

            case playTask.SEE_EAST:
                task = seeEast()

            case playTask.SEE_SOUTH:
                task = seeSouth()

            case playTask.SEE_WEST:
                task = seeWest()

            case playTask.ANSWER:
                
            case playTask.GAME_OVER:
                
        }
    }

    return mainTask.MAIN_MENU
}

func createStage(gameMode int, gameStage int) playTask.PlayTask {
    res, err := question.Load(gameMode, gameStage)
    if err != nil {
        fmt.Println(err)
        lib.InputWait("")
        return playTask.GAME_OVER
    }

    stageArray = res

    return playTask.ACTION_SELECT
}

func actionSelect() playTask.PlayTask {
    inputText, _ := lib.SystemMessage([]string{
        "行動を入力選択してください。",
        "",
        "1 : 北から見る",
        "2 : 東から見る",
        "3 : 南から見る",
        "4 : 西から見る",
        "",
        "5 : 回答する",
        "",
        "いずれかの番号を入力してENTERキーを押してください。",
    })

    userInputNumber, err := strconv.Atoi(inputText)
    if err != nil{
        return playTask.ACTION_SELECT
    }

    switch userInputNumber {
        case 1:
            return playTask.SEE_NORTH
        case 2:
            return playTask.SEE_EAST
        case 3:
            return playTask.SEE_SOUTH
        case 4:
            return playTask.SEE_WEST
        case 5:
            return playTask.ANSWER
    }

    return playTask.ACTION_SELECT
}

func seeNorth() playTask.PlayTask {
    strs := make([]int, len(stageArray))

    for _, y := range stageArray {
        j := 0
        for i:=len(y)-1; i>=0; i-- {
            if y[i] > 0 && strs[j] == 0 {
                strs[j] = y[i]
            }
            j++
        }
    }

    lib.SystemMessage([]string{
        getSeeStringFormat(strs),
        "",
        "あなたは北から空間を見た",
    })

    return playTask.ACTION_SELECT
}

func seeEast() playTask.PlayTask {
    strs := make([]int, len(stageArray))

    for yIndex := len(stageArray) - 1; yIndex >= 0; yIndex-- {
        j := 0
        for i:=len(stageArray[yIndex])-1; i>=0; i-- {
            if stageArray[yIndex][i] > 0 && strs[yIndex] == 0 {
                strs[yIndex] = stageArray[yIndex][i]
            }
            j++
        }
    }

    lib.SystemMessage([]string{
        getSeeStringFormat(strs),
        "",
        "あなたは東から空間を見た",
    })

    return playTask.ACTION_SELECT
}

func seeSouth() playTask.PlayTask {
    strs := make([]int, len(stageArray))

    for yIndex := len(stageArray) - 1; yIndex >= 0; yIndex--{
        for i, x := range stageArray[yIndex] {
            if x > 0 && strs[i] == 0 {
                strs[i] = x
            }
        }
    }

    lib.SystemMessage([]string{
        getSeeStringFormat(strs),
        "",
        "あなたは南から空間を見た",
    })

    return playTask.ACTION_SELECT
}

func seeWest() playTask.PlayTask {
    strs := make([]int, len(stageArray))

    for yIndex, y := range stageArray {
        for _, x := range y {
            if x > 0 && strs[yIndex] == 0 {
                strs[yIndex] = x
            }
        }
    }

    lib.SystemMessage([]string{
        getSeeStringFormat(strs),
        "",
        "あなたは西から空間を見た",
    })

    return playTask.ACTION_SELECT
}

func getSeeStringFormat(strs []int) string {
    marks := make([]string, len(strs))
    for key, value := range strs {
        marks[key] = string(mark.GetMark(value))
    }

    topLine := "-" + strings.Repeat("--", len(stageArray)-1) + "--"
    centerLine := "|" + strings.Join(marks, "|") + "|"
    bottomLine := "-" + strings.Repeat("--", len(stageArray)-1) + "--"

    return topLine + "\n" + centerLine + "\n" + bottomLine
}
