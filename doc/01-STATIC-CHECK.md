# Golang에서의 정적 분석

## 정적분석?

정적 분석은 코드를 체크하여 코드에 문제가 없는지 확인하는 과정.

정적 분석을 통해서 만든 코드가 쓸만한 코드인지 미리 확인하여 빌드하고나서 왜 안되는지 궁금해하는 경우를 꽤나 많이 줄여줌.

## 설치하기 

### GOLANG 1.17+ 의 경우...

```bash
# 패키지 바이너리 설치
go install honnef.co/go/tools/cmd/staticcheck@latest
# 실행이 안되면 PATH를 추가하도록 하자
export PATH=$PATH:$GOPATH/bin
```

## 사용방법 

```bash
# 바이너리 실행
staticcheck .
# 실행이 안되면 PATH를 추가하도록 하자
export PATH=$PATH:$GOPATH/bin
```
