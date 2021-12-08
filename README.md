# 인터파크 기술 공유 세션 02

![interpark_logo.png](data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAANIAAAA2CAYAAACyTMfWAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyJpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuMC1jMDYxIDY0LjE0MDk0OSwgMjAxMC8xMi8wNy0xMDo1NzowMSAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENTNS4xIFdpbmRvd3MiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6RkRFNThCQUIyM0NCMTFFNzg3Q0U5NTA3NEY5QzhFNkUiIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6RkRFNThCQUMyM0NCMTFFNzg3Q0U5NTA3NEY5QzhFNkUiPiA8eG1wTU06RGVyaXZlZEZyb20gc3RSZWY6aW5zdGFuY2VJRD0ieG1wLmlpZDpGREU1OEJBOTIzQ0IxMUU3ODdDRTk1MDc0RjlDOEU2RSIgc3RSZWY6ZG9jdW1lbnRJRD0ieG1wLmRpZDpGREU1OEJBQTIzQ0IxMUU3ODdDRTk1MDc0RjlDOEU2RSIvPiA8L3JkZjpEZXNjcmlwdGlvbj4gPC9yZGY6UkRGPiA8L3g6eG1wbWV0YT4gPD94cGFja2V0IGVuZD0iciI/PlYptT0AABfCSURBVHja7F0JmNzE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)

## 진행자

**넥스트커머스랩 데이터플랫폼팀** 이규연 연구원 

## 목적

Golang을 사용하여 쉽게 배포할 수 있는 서비스 제작

최대한 많은 항목을 진행이 가능했으면 바램입니다.

1. **바퀴를 두번 발명하지 않기**
   이미 어느정도 자기가 하려는 것을 이해하고 있다면, 여러 사람들이 만들어 놓은 오픈 소스 패키지를 이용한 서비스 제작하기
2. **도커 이미지를 최대한 경량화해서 배포** 
   서비스 이미지 유지보수 및 배포시 시간이 오래 걸리지 않도록 최대한 작은 이미지 제작

   그리고 최대한 컨테이너에서 이상한 일을 할 수 없도록 컨테이너 보안 처리 해보기 
3. **다른 언어에서 처리할 수 있는 것들을 golang을 이용하여 api 만들어보기**

    - [ ] 간단한 HTTP 서버

    * [ ] 간단한 JSON 파싱
        * [ ] JSON 검증
        * [ ] 비정형 데이터 순회
        * [ ] 데이터 맵핑 및 golang의 특성 리마인드
    * [ ] Template엔진 사용해보기
    * [ ] 데이터베이스 연동 
        * [ ] Mongodb
        * [ ] Mysql
        * [ ] Redis
    * [ ] 파일 업로드
        * [ ] 파일 후처리 (썸네일 이미지 만들기)
            * [ ] PngQuant를 이용한 PNG 압축
            * [ ] OpenCV를 이용한 이미지 후처리 및 이미지 리사이즈
        * [ ] 파일 생성 (압축 파일 생성 / 액셀 파일 만들기)
            * [ ] 속도 더 빠르게 해보기 (압축 파일 처리시 병렬 처리 가능 항목 확인 및 병렬 처리)
    * [ ] 헤더 연동
        * [ ] JWT
        * [ ] OAUTH
    * [ ] 웹 소켓 서버 제작 
    * [ ] 다른 컨테이너랑 연동
        * [ ] 웹훅형식을 이용한 푸시 알람 사용해보기 (gotify)
        * [ ] 웹소켓을 이용한 서비스 알람 만들어보기 (websocket)
        * [ ] JOB Management


## 기타

* Golang 기술 공유 세션 1은 문서가 완성되는 대로 올라갈 예정입니다.
* 이 프로젝트는 진행중인 세션을 기반으로 하기 때문에, 이 리포지토리는 세션이 진행될때마다 업데이트가 될 예정입니다
* 주 개발 환경은 다음과 같습니다. 혹시 동일한 세팅이 필요하신 분이 있으신 경우 아래와 같이 설정하시면 됩니다.
  * 세션 개발 환경
    * MacOS Catalina
    * VSCode + VIM
    * fish + zsh
  * 개인 개발 환경
    * Linux 
      * `WSL/VirtualBox` Debian Buster, SSH-Terminal
      * `Machine` Arch Linux (Manjaro)
        * `Desktop Environment` KDE Plasma
        * `IME` uim - 벼루
    * Jetbrain Goland + VIM
    * fish
    * 개인 dotfile은 **[dotfiles 리포지토리](https://github.com/rebel1324/dotfiles)**에서 확인 가능합니다.

## 질문

* 연구소 내에서 질문하셔도 됩니다.
  * 긴 질문은 사내 위키로 부탁드립니다.
* 연구소 이외의 분들은 Issue 또는 Discussion을 이용해 주세요

## Contact

프로젝트와 관련한 질문은 이메일로 받지 않습니다.

* 세션 위키 (Github)
* 세션 Discussion (Github)
* **인터파크 이메일**: kylee1337@interpark.com
* **개인 이메일**: rebel1324@gmail.com