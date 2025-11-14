# 암호화폐 자동충전 API

Go로 구현한 간단한 자동충전 API 서버입니다.  
Echo 프레임워크를 사용하며, **SimpleSwap**을 통해 거래를 생성하고 상태를 조회할 수 있습니다.

> 주의: 사용된 SimpleSwap 서비스의 API는 **공식 API가 아닌 비공식 API**입니다.  
> 문제가 발생하면 [sora2931@proton.me](mailto:sora2931@proton.me)로 연락해주세요.

## API 사용 예시

### 1. 새 거래 생성

**요청**
```bash
curl "http://localhost:8080/api/new?crypto=btc&amount=0.001"
```

**응답**

```json
{
  "status": "success",
  "pubId": "aaa"
}
```

---

### 2. 거래 조회

**요청**

```bash
curl "http://localhost:8080/api/check?id=aaa"
```

**응답**

```json
{
  "status": "success",
  "data": {
    "status": "completed",
    "address": "btckslankk12...",
    "txid": "aaaa..."
  }
}
```


