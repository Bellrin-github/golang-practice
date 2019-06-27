package question

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

func Load(gameMode int, stageNumber int) ([][]int, error) {
    file, err := os.Open(fmt.Sprintf("./question/csvs/%03d_%03d.csv", gameMode, stageNumber))
    if err != nil {
        fmt.Println(err)
        return nil, fmt.Errorf(fmt.Sprintf("csvファイルの読み込み失敗\n./question/csvs/%03d_%03d.csv", gameMode, stageNumber))
    }
    defer file.Close()

    reader := csv.NewReader(file)
    var line []string

    stage := make([][]int, gameMode)
    for i:=0; i<gameMode; i++{
        stage[i] = make([]int, gameMode)
    }

    for row := 0 ; ; row++ {
        line, err = reader.Read()
        if err != nil {
            break
        }

        for key, value := range line {
            n, err := strconv.Atoi(value)
            if err != nil{
                stage[row][key] = 0
                continue
            }

            stage[row][key] = n
        }
    }

    return stage, nil
}