# Infrastructure

DB接続など、外部パッケージを操作する場合にここに置きます。

外部ライブラリに依存した書き方（`router := gin.Default()`など）はInfrastructure内に閉じ込めます。

外部パッケージとインターフェースをimportできます。