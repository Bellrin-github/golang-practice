package user

type User struct {
    Type     string `json:"type"`
    Settings Setting `json:"settings"`
}

type Setting struct {
    Type     string `json:"type"`
    IsFirst  int8 `json:"isFirst"`
    IsSecond int8 `json:"isSecond"`
}

