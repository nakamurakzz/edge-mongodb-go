# edge-mongodb-go
以下、2つの機能を持ちます
- 一時データをmongoDBに保存する
- mongoDBからデータを取得してTimeStreamに保存する

# 使い方
## 一時データをmongoDBに保存する
```bash
$ ./bin/main -type=register -temperture=20 -humidity=30
```

## mongoDBからデータを取得してTimeStreamに保存する
```bash
$ ./bin/main -type=publish
```

- TimeStreamに登録されていないデータの登録を行います