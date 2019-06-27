package main

import (
    "fmt"
)

const (
    ROOM_TASK_INIT = iota + 1
    ROOM_TASK_ACTION_INPUT_WAIT
)

var roomTask int = ROOM_TASK_INIT

func RoomMain() bool {
    switch roomTask {
        // 初期化
        case ROOM_TASK_INIT:
            init()
        // 行動入力町
        case ROOM_TASK_ACTION_INPUT_WAIT:
            
        default:
            return false
    }

    return true
}

func Init() {

}
