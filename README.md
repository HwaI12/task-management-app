# タスク管理アプリケーション

React・Go・MySQL・Dockerを利用したタスク管理アプリ

## 注意

初めてDockerを構築し、バックエンドをGo言語で書きました。
また、初めてデータベースにMySQLを使用し、データベース接続やReactとGoの連携を行いました。

ディレクトリ構成やコーディングなど改善点があれば教えていただけると嬉しいです！

## セットアップ手順

### 前提条件

- Docker
- Docker Compose

### プロジェクトのクローン

```sh
git clone https://github.com/HwaI12/task-management-app.git
cd task-management-app
```

### 環境変数の設定

`variables.env` ファイルを作成し、以下を追加：

```env
MYSQL_ROOT_PASSWORD=youre mysql root password
MYSQL_DATABASE=task_management
MYSQL_USER=task_user
MYSQL_PASSWORD=youre mysql password
```

### アプリケーションの起動

```sh
docker-compose up --build
```