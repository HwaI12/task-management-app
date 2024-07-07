- [2024/7/6](#202476)
  - [目的](#目的)
  - [進捗](#進捗)
  - [技術選定](#技術選定)
    - [フロントエンド](#フロントエンド)
    - [バックエンド](#バックエンド)
    - [データベース](#データベース)
    - [コンテナ管理](#コンテナ管理)
  - [詰まったところ](#詰まったところ)
    - [yarnエラー](#yarnエラー)
    - [MySQLの警告](#mysqlの警告)
      - [警告1](#警告1)
      - [警告2](#警告2)
      - [警告3](#警告3)
      - [警告4](#警告4)
- [2024/7/7](#202477)
  
# 2024/7/6

## 目的

- タスク管理アプリの上位互換
  - タスクを投稿
  - そのタスクを達成するまでに必要なプロセスを記入
  - 自分でも記入できるし、生成AIに記入させることも可能
  - そのタスクをどのくらい進めたか進捗を登録
  - 達成した際に個人で成果物を投稿することができる
  - 他のユーザにも見れるように成果物を投稿することもできる

## 進捗
- バックエンドとフロントエンドを繋げ、Docker構築した。
- 次はログイン機能を作成したい。

## 技術選定

### フロントエンド

- **ライブラリ**: React
- **言語**: TypeScript
- **パッケージマネージャ**: Yarn

### バックエンド

- **言語**: Go
- **フレームワーク**: Gorilla Mux

### データベース

- **種類**: MySQL

### コンテナ管理

- **ツール**: Docker, Docker Compose
- CI/CDでGitHub Actionsを利用したい
- ゆくゆくはAWS EC2にデプロイしたい

## 詰まったところ
### yarnエラー

- インストールしようとしたらエラーが発生

```bash
yarn add axios
yarn add v1.22.21
[1/5] 🔍  Validating package.json...
error frontend@0.1.0: The engine "node" is incompatible with this module. Expected version ">=14.15.0 <=18.x". Got "19.8.1"
error Found incompatible module.
info Visit https://yarnpkg.com/en/docs/cli/add for documentation about this command.
```

- 原因: Node.jsのバージョンが正しく設定されていなかった。
- 対策: VSCodeを再起動して解決。

### MySQLの警告
#### 警告1
```bash
The syntax '--skip-host-cache' is deprecated and will be removed in a future release. Please use SET GLOBAL host_cache_size=0 instead.
```
**対策**
- my.cnf ファイルに skip-host-cache の代わりに host_cache_size=0 を設定。

#### 警告2
```bash
Setting lower_case_table_names=2 because file system for /var/lib/mysql/ is case insensitive
```

**対策**
- この警告は情報としてのみ認識し、特に対応は不要。

#### 警告3
```bash
CA certificate ca.pem is self signed.
```

**対策**
- 自己署名証明書を開発環境で使用するのは問題ないが、本番環境ではCA署名の証明書を使用することを検討。

#### 警告4
```bash
Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.
```

**対策**
- my.cnf ファイルで --pid-file のパスを /var/run/mysqld/mysqld.pid に変更。

修正後の my.cnf ファイル:
```my.cnf
[mysqld]
character-set-server=utf8mb4
collation-server=utf8mb4_unicode_ci
skip-host-cache
host_cache_size=0
pid-file = /var/run/mysqld/mysqld.pid

[client]
default-character-set=utf8mb4
```

# 2024/7/7
