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
