# lib

於專案根目錄執行
```
rm -rf internal/lib
git submodule add https://gitlab.alpha.blockaction.tech/bab/lib.git internal/lib
git submodule update --init --recursive
```

定期更新子專案
```
git submodule update --init --recursive
cd internal/lib
git pull origin main
```


使用此lib請update各自專案下的Dockerfile
```
# 使用 Go 官方映像作為基礎
FROM golang:1.20 as base

# 設置工作目錄
WORKDIR /app
ENV TZ=Asia/Taipei
ENV GIN_MODE=debug
EXPOSE 8080


FROM golang:1.20 as build

ARG OAUTH_TOKEN
ARG CI_SERVER_HOST
ENV OAUTH_TOKEN=${OAUTH_TOKEN}
ENV CI_SERVER_HOST=${CI_SERVER_HOST}

# 複製源代碼
COPY . .

RUN git config --global user.name "CI Runner"
RUN git config --global user.email "cirunner@ci.runner"
RUN git config --global url."https://oauth2:${OAUTH_TOKEN}@${CI_SERVER_HOST}/".insteadOf "https://${CI_SERVER_HOST}/"
RUN git submodule update --init --recursive
RUN cd internal/lib && \
    git pull origin main
RUN ls internal/lib

# 複製 go.mod 和 go.sum 文件
COPY go.mod go.sum ./
# 下載依賴
ENV GOPATH=/usr/local/go/bin
RUN go mod download
# 構建應用程序
RUN go build -o app


FROM base AS final
COPY --from=build /go/app ./
COPY --from=build /go/config ./config
# 沒有swagger的不用加下面這行
COPY --from=build /go/docs ./docs

# 運行應用程序
CMD ["./app"]
```

## model使用
```
import (
	libmodel "main/internal/lib/model"
)

# 存入
data := libmodel.Account{
    ...
}
# 檢查not null 每個表有對應的檢查function
err := libmodel.ValidateAccount(data)
if err != nil {
    # 失败
    ...
} else {
    # 成功
    ...
}

```