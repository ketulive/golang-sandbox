/*
コマンドを実行するスクリプト
*/
package main

import (
    "fmt"
    "log"
    "os/exec"
)

func main() {
    out, error := exec.Command("date").Output()
    if error != nil {
        log.Printf("エラー: %s", error)
        return
    }
    fmt.Printf("現在時刻: %s", out)
}