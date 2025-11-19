# 암호화폐 자동충전 API

Go로 구현한 간단한 자동충전 API 서버입니다.  
Echo 프레임워크를 사용하며, **SimpleSwap**을 통해 거래를 생성하고 상태를 조회할 수 있습니다.

> 주의: 이 프로젝트는 비공식 API를 사용하며, 학습·연구 목적에서만 이용해야 합니다.
> 실제 서비스·자동화·금융 거래 등 실사용은 절대 금지됩니다.
> 사용으로 인해 발생하는 모든 책임은 사용자에게 있습니다.
> 
>문의: [sora2931@proton.me](mailto:sora2931@proton.me)

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


