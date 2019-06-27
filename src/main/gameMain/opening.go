package gameMain

import (
    "../lib"
    "../user"
    "encoding/json"
    "io/ioutil"
    "os"
)

func Opening() user.User {
    var isFirst bool = false
    var userSetting user.User

    userSettingRaw, err := ioutil.ReadFile("./save/userSetting.sav")
    if err != nil {
        isFirst = true
    }

    if isFirst {
        lib.SystemMessage([]string{
            "Go言語練習用プログラムです。",
            "メッセージの最後にに ▼ が表示されている場合、",
            "ENTERキーの押下で次へ進みます。",
        })

        lib.SystemMessage([]string{
            "ENTERキーでの「次へ」操作は理解頂けたようですね。",
        })

        lib.SystemMessage([]string{
            "プログラムはいつでも Ctrl + c コマンドで終了する事が出来ます。",
        })

        lib.SystemMessage([]string{
            "今回が初回起動のようですね。",
            "設定ファイルを作成します。",
        })

        userSetting.Type = "User"
        userSetting.Settings.Type = "Setting"
        userSetting.Settings.IsFirst = 1
        userSetting.Settings.IsSecond = 2
        outputJson, err := json.Marshal(&userSetting)
        if err != nil {
            lib.SystemMessage([]string{
                "Jsonファイルの変換に失敗しました。",
            })
            os.Exit(0)
        }

        ioutil.WriteFile("./save/userSetting.sav", outputJson, os.ModePerm)

        lib.SystemMessage([]string{
            "設定ファイルの作成に成功しました。",
            "ゲームを開始します。",
        })

        return userSetting

    }
    lib.SystemMessage([]string{
        "またお会いしましたね。",
        "早速ゲームを開始します。",
    })

    json.Unmarshal(userSettingRaw, &userSetting)
    userSettingRaw = nil

    return userSetting
}
