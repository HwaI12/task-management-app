# タスク管理アプリケーション

このプロジェクトは、ReactとGoを使用したタスク管理アプリケーションです。

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