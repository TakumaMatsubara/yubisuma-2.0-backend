# ベースとなるイメージを指定します
FROM golang:1.20

# GOPROXYを指定して、Goの公式モジュールミラーを使用します（オプション）
ENV GOPROXY=https://proxy.golang.org,direct

# 作業ディレクトリを設定します
WORKDIR /app

# 依存関係のファイルをコピーします
COPY go.mod go.sum ./

# 依存関係をダウンロードします
RUN go mod download

# ソースコードをコピーします
COPY . .


# アプリケーションをビルドします
# RUN go build -o main .

# # アプリケーションを実行します
# CMD ["./main"]
CMD ["go", "run", "main.go"]