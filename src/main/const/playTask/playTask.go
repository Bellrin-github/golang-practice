package playTask

type PlayTask int

const (
    INIT PlayTask = iota + 1
    ACTION_SELECT
    SEE_NORTH
    SEE_EAST
    SEE_SOUTH
    SEE_WEST
    ANSWER
    GAME_OVER
)
