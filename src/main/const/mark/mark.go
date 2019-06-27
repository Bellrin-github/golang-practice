package mark

type Mark string

const (
    SPACE Mark = " "
    SQUEAR Mark = "■"
    TRIANGLE Mark = "▲"
    CIRCLE Mark = "●"
    RHOMBUS Mark = "◆"
    TRIANGLE_R Mark = "▼"
    STAR Mark = "★"
)

func GetMark(n int) Mark {
    switch n {
        case 0:
            return SPACE
        case 1:
            return SQUEAR
        case 2:
            return TRIANGLE
        case 3:
            return CIRCLE
        case 4:
            return RHOMBUS
        case 5:
            return TRIANGLE_R
        case 6:
            return STAR
    }

    return SPACE
}