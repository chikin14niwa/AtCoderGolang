# 大元
FROM golang:alpine

# アップデートとgitのインストール
RUN apk add --update && apk add git

# appディレクトリの作成
RUN mkdir /go/src/app

# ワーキングディレクトリの設定
WORKDIR /go/src/app

# ホストのファイルをコンテナの作業ディレクトリに移行
ADD ./src /go/src/app/