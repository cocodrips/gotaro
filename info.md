# goでslackbot作る

## 作業計画

1. slackの仕様大体調べる
1. 簡単な投稿テストする
1. goのWAFの仕様大体調べる（いまここ）
1. slackへの投稿（incoming）, slackからのPOSTをhandleする用（outgoing）に分かれて開発
1. GAEかHerokuにデプロイ

## slack API

|= api =|= detail =|
|Incoming WebHooks|slackがURL発行, URLへPOSTすると指定channel,botname,msgで投稿する|
|Outgoing WebHooks|slackへ指定のmsgが投稿されると指定したURLに投稿内容がPOSTされる, 指定URLにリクエストが来たらハンドルするwebapp作ったらよし|
|Slack API|メッセージ投稿, ファイルアップロード・チーム情報やチャンネル情報の取得, 絵文字のリストとれたりとか（いらん）|

- 細かいAPIはSlackAPI, 手軽にリアルタイム応答使おうとおもったら Incoming/Outgoing Webhook使うと簡単だよみたいに見えた（わからんけど）

## Incoming Webhooksの設定
https://bolders.slack.com/services/7897851120

## Outgoing Webhooksの設定
https://bolders.slack.com/services/7897499748

Slack API
https://api.slack.com/

## goのおすすめWAF

@cocodrips たのむ


