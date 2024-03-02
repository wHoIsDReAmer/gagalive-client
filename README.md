# Gagalive client
gagalive client is unofficial client implemented golang

## Analysis

* 패킷 타입: 웹소켓
* 패킷 암호화 방식: 없음 (NoEncryption)
> ⚠ 보안에 굉장히 취약함
* 서버 주소: ws://rchat.gagalive.kr:8082/


**패킷 종류**
- **Send**:
    - 세션 등록: `Y[세션]`
    - 닉네임 등록: `L[닉네임]|@@@randomchat`
    - 메세지 보내기: `#[메세지]`
    - 스팸 방지 문자 해결: ``P``

- **Receive**:
    - 채팅 받기 (낯선 사람에게): `:[고유아이디|고유아이디]`
    - 채팅 받기 (시스템 메세지): `G[메세지]`
    - 채팅 시작 알림: `C[My Temp ID]|[Opponent Temp ID]`

## Usage example
```go
client := gagalive.NewGagaClient()

client.SetOnConnected(func() {
    client.Send("foo")
})

client.SetOnMessage(func(msg string) {
    // Solve the captcha
    if strings.Contains(msg, "방지 문자: ") {
        captcha := strings.Split(msg, "방지 문자: ")[1]
        solvedCaptcha := gagalive.NewSolveMacro().Solve(captcha)

        terminal.AddString(solvedCaptcha)
        client.Send(solvedCaptcha)
        return
    }
})

client.SetOnDisconnected(func() {
    fmt.Println("Bye anonymous")
    // when it disconnected do reconnect
    client.Connect()
})
```