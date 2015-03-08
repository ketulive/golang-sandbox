/*
マップ（Key-Value）の使い方
*/
package main

import (
    "fmt"
)

func main() {
    var month map[int]string = map[int]string{
        1: "January",
    }
    month[2] = "February"
    month[3] = "March"
    month[4] = "April"
    month[5] = "May"

    fmt.Println(month)
    fmt.Println(month[1])

    jar := month[1]
    fmt.Println(jar)

    // 指定したキーがこのマップに格納されているか
    _, ok := month[2]
    fmt.Println(ok)
    if ok {
        // データがあった場合
        fmt.Println("データがあります")
    }

    // for文で順番に処理する
    for key, value := range month {
        fmt.Printf("%d %s\n", key, value)
    }

    // マップからデータを消す
    delete(month, 1)
    fmt.Println(month)
}
