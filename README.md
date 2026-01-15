# sacloud/simple-notification-api-go
さくらのクラウド シンプル通知 Go言語向け APIライブラリ

マニュアル: https://manual.sakura.ad.jp/cloud/appliance/simplenotification/


## 概要
sacloud/simple-notification-api-goはさくらのクラウド シンプル通知 APIをGo言語から利用するためのAPIライブラリです。




## ogenによるコード生成

以下のコマンドを実行

```
$ go get -tool github.com/ogen-go/ogen/cmd/ogen@latest
$ go tool ogen -package v1 -target apis/v1 -clean -config ogen-config.yaml ./openapi/openapi.yaml
```


## License

`simple-notification-api-go` Copyright (C) 2026- The sacloud/simple-notification-api-go authors.
This project is published under [Apache 2.0 License](LICENSE).