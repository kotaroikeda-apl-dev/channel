# Go の Channel を使ったサンプル

## 概要

このプロジェクトは、Go 言語の `channel` を活用し、並行処理でデータをやり取りする方法を示します。
このサンプルでは、`channel` を使用してデータを送信し、受信する基本的な方法を学びます。

## `channel` とは？

`channel` は、Go 言語が提供するデータの受け渡し機構であり、複数の `goroutine` 間で安全にデータをやり取りするために使用されます。

このサンプルでは、`channel` を使って、並列処理を行う `goroutine` からデータを送信し、メイン関数でデータを受信する方法を解説します。

## **実行方法**

```sh
go run cmd/basic/main.go  #チャネルの基本実装
go run cmd/multi/main.go  #複数のgoroutineからデータを受信
go run cmd/limit/main.go  #ジョブの並列実行を制限
go run cmd/waitGroup/main.go  #sync.WaitGroup + channelでジョブの同期処理
go run cmd/error/main.go  #エラーチャネルを使用
```

## **学習ポイント**

1. **`channel` を使うことで `goroutine` 間でデータをやり取りできる**
2. **`chan` 型を定義し、`<-` でデータの送受信を行う**
3. **送信側 (`ch <- value`) がデータを送ると、受信側 (`<-ch`) が受け取るまでブロックされる**
4. **バッファなしチャネル（`make(chan string)`）では、送信と受信が同期的に行われる**
5. **`close(ch)` を使うことで、チャネルを明示的に閉じることができる**
6. **非同期で複数の `worker` を起動し、それぞれの完了メッセージを `channel` に送信**
7. **`defer wg.Done()` を使うことで、関数の終了時に確実に `wg.Done()` を実行し、ゴルーチンが終了するたびに `sync.WaitGroup` のカウントが減る**
8. **チャネルを close() しないと、受信側が range を抜けずにデッドロックが発生する可能性がある**
9. **エラー発生時に errors チャネルを使うことで、処理を中断せずにエラーメッセージを送信できる**

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
