package mainTask

type MainTask int

const (
    MAIN_MENU MainTask = iota + 1
    MODE_SELECT
    MANUAL
    EXIT_GAME
    STAGE_SELECT
    PLAY
)
