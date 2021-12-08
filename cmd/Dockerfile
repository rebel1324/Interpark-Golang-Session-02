#==================================================
# BUILD IMAGE
#==================================================

# 1.17.4 GO환경의 UBUNTU/DEBIAN, BULLSEYE 버전 환경 이미지
FROM golang:1.17.4-bullseye AS build

# 1.11 go module형식 ON (go mod init)
#ENV GO111MODULE=on
# 빌드 대상 linux
#ENV GOOS=linux
# x64 CPU (64bit)
#ENV GOARCH=amd64

# 빌드 대상 폴더 세팅
RUN mkdir /app
WORKDIR /app

# go.mod, go.sum만 복사해서 이 상태 유지
COPY go.* .
# go.mod, go.sum이 바뀔때만 패키지 다운로드
RUN go mod download

# MACOS: CGO ISSUE
# UNFORTUNATELY, this will cause issues when we make something related with c interface such as librdkafka
ENV CGO_ENABLED=0

#소스코드 복사 및 빌드
COPY . .
RUN go build -o bin/main
RUN chmod +x bin/main

#==================================================
# RUNTIME IMAGE
#==================================================

# 아무것도 없는 빈 이미지
# 사이즈 0byte
FROM alpine:latest AS runtime

# 기존 빌드 이미지에서 빌드한 것들을 가져온다
COPY --from=build /app/bin /app/bin
WORKDIR /app/bin

# 8080포트 노출
EXPOSE 8080

# 어플리케이션 실행
CMD ["./main"]
