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
