# go-web

- assets

  - https://stripe.com/
  - https://www.gnu.org/software/make/manual/

    ```
    brew update && brew install make
    ```

    ```
    # 설치검증
    $ make

    # makefile 이 없어서 멈춘다는 메시지가 나온다면 성공!
    > make: *** No targets specified and no makefile found.  Stop.
    ```

- init

```sh
# go mod init <프로젝트이름>
go mod init goweb
docker-compose -f asset/go-web-infra/infra.yaml up -d
```


## 인트로
- 왜 고를 공부해야하나
  - PPH 대비 서버 사용량 엄청 적음


## 웹사이트 만들기
- 트레픽의 흐름
  - DNS 네임서버 >>  쿼리해서 IP 얻어옴
  - IP 로 서버 접속해서 HTML 같은 자원 가져옴
  - 


## go.mod 파일에 관하여

```sh
go mod init myapp
```


## 8월31일 28번 강의부터

- rkwl
  - 고모듈 시작한거
  - ??
  - 템플릿 페이지 레이아웃 만든거임


### 고랭컨벤션

```
/[모듈] 
    /cmd
      /web
    /pkg
      /handlers
      /render
```

- 함수의 첫글자 : 소문자면 >> 프라이빗
- 함수의 첫글자가 대문자 >> 퍼블릭 >> 비저블함