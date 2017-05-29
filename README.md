# Auth API
[![Build Status](https://travis-ci.org/charakoba-com/auth-api.svg?branch=master)](https://travis-ci.org/charakoba-com/auth-api)
---

Authentication API Server for charakoba.com

## Routings

| method | path       | description                             |
|:------:|:-----------|:----------------------------------------|
| ANY    | /          | health check                            |
| POST   | /user      | create user                             |
| DELETE | /user      | delete user                             |
| GET    | /user/list | get user list                           |
| POST   | /auth      | authenticate with username and password |
| GET    | /algorithm | get signing algorithm                   |
| GET    | /alg       | alias for /algorithm                    |
| POST   | /verify    | verify authorization token              |
| GET    | /key       | get public key for verify auth token    |
