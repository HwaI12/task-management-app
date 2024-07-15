# 適切なベースイメージを使用
FROM golang:1.22

# Debianベースのイメージを使用する場合は、apt-getを使用してSupervisorをインストール
RUN apt-get update && apt-get install -y supervisor

# Supervisorの設定ディレクトリを作成
RUN mkdir -p /etc/supervisor/conf.d

# Goアプリケーションのコードをコピー
COPY . /app
WORKDIR /app

# Goアプリケーションをビルド
RUN go build -o main ./cmd/main.go

# Supervisorの設定ファイルをコピー
COPY supervisord.conf /etc/supervisor/supervisord.conf

# Supervisorを実行するコマンド
CMD ["supervisord", "-c", "/etc/supervisor/supervisord.conf"]
