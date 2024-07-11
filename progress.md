- [2024/7/6](#202476)
  - [環境構築](#環境構築)
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
  - [進捗](#進捗-1)
    - [認証機能](#認証機能)
      - [Backend](#backend)
      - [frontend](#frontend)
      - [GitHub](#github)
  - [詰まったところ](#詰まったところ-1)
    - [コンテナ間通信ができない場合の原因](#コンテナ間通信ができない場合の原因)
- [2024/7/8](#202478)
  - [進捗](#進捗-2)
    - [コード修正](#コード修正)
      - [backend](#backend-1)
    - [UI/UX修正](#uiux修正)
  - [画面設計](#画面設計)
    - [認証関連画面](#認証関連画面)
    - [個人関連画面](#個人関連画面)
    - [共有用画面](#共有用画面)
    - [主要機能](#主要機能)
- [2024/7/9](#202479)
  - [MySQLの中身を確認する方法](#mysqlの中身を確認する方法)
- [2024/7/10](#2024710)
  - [詰まったところ](#詰まったところ-2)
    - [React-Select](#react-select)
  
# 2024/7/6

## 環境構築
[Go + MySQL + React のDocker開発環境を作成する](https://qiita.com/makosinhori/items/c695774bef249a2014a6)
を元に環境構築を行なった

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
## 進捗
### 認証機能
#### Backend
- `register.go`: 新規登録用にデータベースと連携
  ```bash
  curl -X POST http://localhost:8000/register -d '{
                "username": "wa",
                "email": "wa@example.com",
                "password_hash": "wa"
            }' -H "Content-Type: application/json"
  ```
- `login.go`: ログイン用にデータベースと連携
  ```bash
  curl -v -X POST http://localhost:8000/login -d '{
                "username": "wa",
                "password": "wa"
            }' -H "Content-Type: application/json"
  ```
- `delete_user.go`: 退会用にデータベースと連携
  ```bash
  curl -X POST http://localhost:8000/delete -d '{
             "email": "wa@example.com"
         }' -H "Content-Type: application/json"
  ```
- `main.go`: ルーティング・データベース接続・ポート・CORS
#### frontend
**フロントエンドとバックエンドの連携**
- `Register.tsx`: 3000/registerから新規登録可能にした
- `Login.tsx`: 3000/loginからログイン可能にした
- `DeteteAccount.tsx`: 3000/deleteから退会できるようにした
  - しかしemailを使って退会しているため認証データとパスワードを使って退会できるようにしたい

**画面移動**
- 認証前は/Homeには移動できない
- 認証後は必ず/Registerと/loginには移動できない

#### GitHub
ISSUEを使ってタスクを管理するために、Templateを作成した

## 詰まったところ
### コンテナ間通信ができない場合の原因

```bash
backend-1   | 2024/07/07 02:57:12 Error saving user: dial tcp: lookup db on 127.0.0.11:53: no such host
```
**原因**
- MySQLのコンテナの起動が終わる前に、GoがDBコンテナ接続に行ってしまっていたため。
**詳細**
- dbのコンテナにdepends_onやlinks属性を書いていたとしても、goのコンテナはdbコンテナが完全に準備ができるのを待たずに起動してしまう(Dockerはそこまでの制御をできない)。らしい

**対応策**
Supervisor経由でコンテナを起動する。
Supervisorを使うと、DBコンテナに接続できないせいでGoのコンテナが落ちても、プロセス監視によってまた再起動できるため。

参考記事: [コンテナ間通信ができない場合の原因について](https://qiita.com/satofujino/items/cbfc0dcca36f48bd17d3)

# 2024/7/8
## 進捗
### コード修正
#### backend
- `main.go`を分割した
  - `db/database.go`: データベースに接続
  - `router/router.go`: ルーティング
### UI/UX修正
- `/delete`はいつかやる。今はまだやるつもりなし
- `/`, `/signin`, `/signup`, `/logout`のUIを修正

## 画面設計
### 認証関連画面
1. **ログイン画面**
   - ユーザ名とパスワードの入力フォーム
   - パスワードのリセットオプション

2. **新規登録画面**
   - ユーザ名、メールアドレス、パスワードの入力フォーム
   - 利用規約やプライバシーポリシーの同意チェックボックス

3. **ログアウト画面**
   - ログアウトの確認メッセージ

4. **退会画面**
   - アカウント削除の確認メッセージ
   - 削除手続きのフォーム（パスワードの再確認）

### 個人関連画面
1. **プロフィール画面**
   - ユーザ名、プロフィール写真の表示と編集
   - 投稿したタスクと制作物の一覧表示
   - いいねした他人の制作物の表示

2. **通知画面**
   - 新着通知の一覧表示
   - 通知のクリア機能

3. **タスク投稿画面**
   - タスクのタイトル、説明、期限、優先度の入力フォーム
   - ファイルの添付オプション

4. **制作物投稿画面**
   - 制作物のタイトル、説明、カテゴリ、画像のアップロードフォーム
   - タグの追加オプション

### 共有用画面
1. **ホーム画面**
   - 自分のプロフィールへのリンク
   - 他人の投稿した制作物の一覧表示
   - 制作物をクリックすると詳細ページへのリンク
   - ヘッダーにホームと通知画面へのリンク

2. **人のプロフィール画面**
   - ユーザ名、プロフィール写真の表示
   - そのユーザが投稿した制作物の一覧表示

### 主要機能
1. **ログイン管理**
   - ログインのセキュリティ管理
   - セッションの管理

2. **タスク管理**
   - タスクの作成、編集、削除機能
   - タスクのリスト表示とフィルタリング（期限、優先度など）

3. **タスクの手順生成**
   - ユーザが設定した目標に基づいたタスクの自動生成または手動設定

4. **成果物管理**
   - タスク達成時の成果物（画像やドキュメントなど）のアップロード機能
   - タスクとの連携機能

5. **進捗共有**
   - 他のユーザとのタスクの進捗共有
   - 成果物に対するいいねやコメント機能

# 2024/7/9
やること
1. **プロフィール画面**
   - ユーザ名、プロフィール写真の表示と編集
   - 投稿したタスクと制作物の一覧表示
   - いいねした他人の制作物の表示
2. **制作物投稿画面**
   - 制作物のタイトル、説明、カテゴリ、画像のアップロードフォーム
   - タグの追加オプション

- [x] 認証したユーザのID以外がURLに書かれた場合、共有用のUIになるようにしたい
- [x] /userid/addtaskでタスクの投稿ができてしまうが、
- [x] タスククリックしたらタスクの内容が表示されるページ作成

今後やること
- タスクをクリックした後のUI/UXの修正

## MySQLの中身を確認する方法

1. ターミナルで以下のコマンドを実行する:
    ```sh
    docker-compose exec db mysql -u root -p
    ```
2. パスワードを入力する。
3. データベースを選択する:
    ```sql
    USE task_management;
    ```
4. テーブルの一覧を表示する:
    ```sql
    SHOW TABLES;
    ```
5. テーブルの中身を表示する:
    ```sql
    SELECT * FROM users;
    ```

# 2024/7/10
## 詰まったところ
- フロントエンドでのログを追加したい
  ```typescript
  export const log = (message: string, level: 'debug' | 'info' | 'warn' | 'error' = 'info') => {
    console.log(`[${level.toUpperCase()}] ${message}`);
  };

  log('This is an error message', 'error');
  ```
  ```dockerfile
  # ベースイメージ
  FROM node:18-alpine
  
  # 作業ディレクトリを設定
  WORKDIR /app
  
  # 依存関係のインストール
  COPY package*.json ./
  COPY yarn.lock ./
  RUN yarn install

  # アプリケーションのソースをコピー
  COPY . .

  # ポートの公開
  EXPOSE 3000

  # アプリケーションの起動
  CMD ["yarn", "start"]
  ```
  ```yml
  services:
   frontend:
      build: ./frontend
      ports:
      - '3000:3000'
      tty: true
      stdin_open: true
      environment:
      - CHOKIDAR_USEPOLLING=true
      - REACT_APP_LOG_LEVEL=debug
      logging:
      driver: "json-file"
      options:
         max-size: "200k"
         max-file: "10"
  ```
  と追加し、`docker-compose logs frontendとログを確認したところ、表示されなかった
- バックエンドのディレクトリ構成を修正したいが、今の技術では難しいことがわかった。
  - データベースに保存はできるが、ハッシュ化したパスワードの取得などのログインやユーザの取得ができなかった
  - POSTはできるがGETができない

- [DockerでReact＋TypeScript環境を作ってみた～formatterを添えて～](https://logical-studio.com/develop/backend/20211217-docker-react-formatter/)を元に再構築したみたが、eslintのconfigファイルが`.eslintrc`だったものが、9.0.0以降はデフォルトが`eslint.config.mjs`になるそう。あまり文献がなかったため断念

### React-Select
```typescript
<Select
    styles={customStyles}
 />
```
このように使用できるが、customStylesでwidthを設定すると文字によって横幅が変わってしまう。

解決策
```typescript
<div style={{width:'100%'}}>
  <Select
      styles={customStyles}
  />
</div>
```
横幅が修正された！