# About

gRPC を用いてリアルタイムチャットアプリを作ってみる(server 側)

# Setup

- サーバーを 8080 ポートで起動する

```
$ make run
```

- MySQL のデータを初期化する

```
$ make reset_migration
```

# Preview

![output-palette](https://user-images.githubusercontent.com/38310693/114688320-8be5c300-9d4f-11eb-9901-ac0efb7b5ebd.gif)


# 構成
## Domain層
- ドメインモデルの定義・ライフサイクルの表現
### domain/model/
- エンティティ
- 値オブジェクト
- ファクトリ
- リポジトリ

### domain/service/
- ドメインサービス

## Infrastructure層
- Domain層で定義したRepositoryを実現する

### infrastructure/repository/
- Domain層で定義したRepositoryの実装クラス

### infrastructure/mysql/
- mysqlとのコネクション確保
- DBデータ⇆Entityの変換を行う独自のFactory(DTO)

### infrastructure/redis/
- redisのコネクション確保


## Application層
- Domain情報を駆使して、ユースケースを進行させる

### application/${USECASE_NAME}/
- 入力値(InputData)
- 出力値(OutputData)
- ユースケースを実現するinteractor
- interactorのインターフェースであるinput_port


## Interfaces層
### interafces/controller/
- リクエストを整形しアプリケーションサービスに伝搬、出力のレスポンスも担う(controller/presenter)

----------------------------------

## proto/
- protoファイル定義・自動生成コード

## mock/
- gomockでモックした関数

## test/
- モックを利用したユニットテスト


