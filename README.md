# cram-api

Challenge Response Authentication Mechanism (CRAM) Example, Using Elliptic Curve Digital Signature Algorithm.

## [1] Registration

```shell
curl --location --request POST 'http://localhost:3000/v1/challenge/registration' \
--data-raw '{
    "userId": "1",
    "pk": "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE1Pp8aJj3w61AW78eKxd+wiHYERzJti4I2dngMVVMcZ759TLxLETJPlku5MFkNlGJNLu4GT1bt8lEvxi5h3A8nA=="
}'
```

Response

```json
{
    "code": 200,
    "message": "OK"
}
```

## [2] Request

```shell
curl --location --request POST 'http://localhost:3000/v1/challenge/request' \
--data-raw '{
    "userId": "1"
}'
```

Response

```json
{
  "code": 200,
  "data": {
    "challenge": "577cb5f5-34d2-43ea-8a41-c67cdbf80a4a"
  },
  "message": "OK"
}
```

## [3] Verify

```shell
curl --location --request POST 'http://localhost:3000/v1/challenge/verify' \
--data-raw '{
    "userId": "1",
    "challenge": "577cb5f5-34d2-43ea-8a41-c67cdbf80a4a",
    "signature": "MEQCIH6X8KrktBuTaV3StA0TYe4OuGoqoP4EM4hleuHVUlCQAiBHso4R8EUTIUIWfVgI4qpW1qn7Zkz9h0CmyJypgBRyBA==",
    "nonce": "f8a01ad7-abbb-43d6-bc86-d811214e7cf9"
}'
```

Response

```json
{
  "code": 200,
  "data": {
    "token": "5a2d29b9-346c-40db-aa7a-7422dc7e57ab"
  },
  "message": "OK"
}
```