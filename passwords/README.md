## Passwords

### Reset Password

We use PASETO under the hood: https://paseto.io/

Set your env values first: `PASETO_PRIVATE_KEY` & `PASETO_PUBLIC_KEY`

You could generate those using `ed25519`:

```go
publicKey, privateKey, _ := ed25519.GenerateKey(nil)

fmt.Println("public key:  ", hex.EncodeToString(publicKey))
fmt.Println("private key: ", hex.EncodeToString(privateKey))
```

Generate new reset token:
```go
passwords.NewResetToken(userID string, expiration time.Time)
```

Get UserID from reset token:
```go
passwords.GetUserIDFromToken(token string)
```
