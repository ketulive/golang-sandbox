/*
ゴルーチンサンプル
*/
package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

func main() {
    wait := new(sync.WaitGroup)
    urls := []string {
        "https://www.google.co.jp/",
        "http://www.yahoo.co.jp/",
        "http://www.livedoor.com/",
        "https://twitter.com/",
    }
    for _, url := range urls {
        // WaitGroupに追加する（カウントを増やす）
        wait.Add(1)
        // 取得処理をゴルーチンで実行する
        go func(url string) {
            res, err := http.Get(url)
            if err != nil {
                log.Fatal(err)
            }
            defer res.Body.Close()
            fmt.Println(url, res.Status)
            // WaitGroupから削除する（カウントを減らす）
            wait.Done()
        }(url)
    }
    // カウントがゼロになるまで待ち合わせする
    wait.Wait()
}
