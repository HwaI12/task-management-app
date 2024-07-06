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