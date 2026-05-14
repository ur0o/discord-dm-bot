# discord-dm-bot

DiscordのBotから特定ユーザーにDMを送信するためのGoライブラリです。

このReadmeはClaude Codeによって書かれています。

## 特徴

- シンプルなAPI設計
- 自動リトライ機能（Rate Limitやサーバーエラーのハンドリング）
- タイムアウト・リトライ間隔のカスタマイズ可能
- 軽量な依存関係

## インストール

```bash
go get github.com/ur0o/discord-dm-bot
```

## 前提条件

このライブラリを使用するには以下が必要です：

1. **Botトークン**: [Discord Developer Portal](https://discord.com/developers/applications)でBotを作成し、トークンを取得
2. **ユーザーID**: DMを送信したい対象ユーザーのDiscord ID

### ユーザーIDの取得方法

1. Discordの設定 → 詳細設定 → 開発者モードを有効化
2. 対象ユーザーを右クリック → IDをコピー

## 基本的な使い方

```go
package main

import (
    "fmt"
    "log"

    ddm "github.com/ur0o/discord-dm-bot"
)

func main() {
    botToken := "YOUR_BOT_TOKEN"
    userID := "TARGET_USER_ID"

    // クライアントの初期化
    client, err := ddm.New(botToken, userID)
    if err != nil {
        log.Fatal(err)
    }

    // メッセージの送信
    res, err := client.Message.Post("Hello from Bot!")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Message sent:", string(res))
}
```

## オプション設定

リトライ回数、リトライ間隔、タイムアウトをカスタマイズできます。

```go
import (
    "time"
    ddm "github.com/ur0o/discord-dm-bot"
)

client, _ := ddm.New(botToken, userID)

// オプション付きでメッセージ送信
res, err := client.Message.Post(
    "Hello with custom options!",
    ddm.MaxRetry(3),                          // 最大3回リトライ
    ddm.RetryInterval(500 * time.Millisecond), // リトライ間隔500ms
    ddm.Timeout(15 * time.Second),             // タイムアウト15秒
)
```

### 利用可能なオプション

| オプション | 型 | デフォルト値 | 説明 |
|-----------|-----|-------------|------|
| `MaxRetry` | `uint` | `0` | 最大リトライ回数 |
| `RetryInterval` | `time.Duration` | `100ms` | リトライ間隔 |
| `Timeout` | `time.Duration` | `10s` | リクエストタイムアウト |

### 自動リトライ対象

以下のHTTPステータスコードで自動的にリトライされます：
- `429` Too Many Requests (Rate Limit)
- `500` Internal Server Error

## API リファレンス

### `ddm.New(botToken, userID string) (*Client, error)`

新しいクライアントを作成します。初期化時に対象ユーザーとのDMチャンネルIDを自動取得します。

**パラメータ:**
- `botToken`: Discord BotのトークンOAuth2
- `userID`: 送信先ユーザーのDiscord ID

**戻り値:**
- `*Client`: 初期化されたクライアント
- `error`: エラーが発生した場合

### `client.Message.Post(message string, options ...request.Option) ([]byte, error)`

指定したユーザーにDMを送信します。

**パラメータ:**
- `message`: 送信するメッセージ内容
- `options`: リクエストオプション（可変長引数）

**戻り値:**
- `[]byte`: Discord APIからのレスポンス（JSON形式）
- `error`: エラーが発生した場合

### オプション関数

以下のオプション関数が`ddm`パッケージから直接利用できます：

#### `ddm.MaxRetry(maxRetry uint) request.Option`
最大リトライ回数を設定します（デフォルト: 0）

#### `ddm.RetryInterval(retryInterval time.Duration) request.Option`
リトライ間隔を設定します（デフォルト: 100ms）

#### `ddm.Timeout(timeout time.Duration) request.Option`
リクエストタイムアウトを設定します（デフォルト: 10秒）

## ライセンス

MIT License
