## Tokens

### Access Token

We use PASETO under the hood: https://paseto.io/

Set your env values first: `PASETO_SYMMETRIC_KEY` (32 bytes key)

You could generate those using `tokens.GenerateSymmetricKey()` function:

```go
fmt.Println("generated symmetric key:  ", tokens.GenerateSymmetricKey())
```

Generate new reset token:
```go
passwords.NewAccessToken(payload string, expiration time.Time)
```

Get Payload from reset token:
```go
passwords.GetPayloadFromToken(token string)
```

See tests in order to set and get a given JSON payload.
