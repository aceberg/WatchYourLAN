# How to encrypt password with bcrypt?

It is not safe to store password unencrypted, so this app uses `bcrypt` encryption. There are several ways to encrypt your password.

## 1. Set password through web GUI
Then the app will encrypt it for you.

## 2. Encrypt password yourself
On Linux encryption can be done with `htpasswd` command:
```sh
htpasswd -nbBC 10 USER YourSecretPassword | sed 's/USER://'
```

## 3. Encrypt password online
There are online tools for `bcrypt` encryption.