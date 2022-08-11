# バックエンド環境

## ＜ローカル環境構築＞

### 【インストールなど】

- [こちら](https://qiita.com/waka424/items/4cb8ff710b0eda940489)を参照して MongoDB をインストールしてください。また、[こちら](https://dev.classmethod.jp/articles/introducing-mongodb-compass/)の GUI ツールを入れると便利と思います。
- Golang のバージョンは 1.16 を使用しているので、[こちら](https://qiita.com/walkers/items/761b2a5e58849176a633)を参考にしてバージョン指定でインストール。
- VSCode をインストールしておいてください。[こちら](https://future-architect.github.io/articles/20201117/)を参考にデバッガツールをインストール。

### 【DB のセットアップ】

- 「coda_db」の名前で DB を作成してください。
- [こちら](https://dev.classmethod.jp/articles/db-migrate-with-golang-migrate/)を参考に golang-migrate をインストール。
- API のフォルダで`migrate -database "mongodb://localhost:27017/coda_db" -path "./migrate/" up 4`の形式で migrate コマンドを実行してください。

### 【API の準備】

- API のフォルダで`go mod vendor`を実行してください。

### 【API のローカル起動】

- VSCode より`main.go`を選択して`Start Debbuging`を実行してください。
