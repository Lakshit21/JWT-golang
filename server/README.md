## Project Setup

```
server-backend/
│
├── .env
├── go.mod
├── go.sum
├── main.go
└── config/
    └── db.go
└── controllers/
    └── authController.go
└── models/
    └── userModel.go
└── routes/
    └── authRoutes.go
└── utils/
    └── jwtUtils.go
```

## (dot)env setup
```
touch .env
```
and setup the following
```
MONGO_URI=<your-mongodb-atlas-uri>
JWT_SECRET=<your-jwt-secret-key>
PORT=8080
```

## Supported Features
- user defined email and password (like workday)

## Test JWT-AUTH Using Postman
These are steps to follow
### /register route
HTTP METHOD: POST
BODY>RAW(JSON)

request body
```
{
    "first_name" : "Tom",
    "last_name": "Riddle",
    "Email": "youknowwho@gmail.com",
    "Password": "avadakadavra"
}
```

response body
```
{
    "InsertedID": <some-random-id>
}
```

### /login route
HTTP METHOD: POST
BODY>RAW(JSON)

request body
```
{
    "Email": "youknowwho@gmail.com",
    "Password": "expelliarmus"
}
```

response body
```
{
    "error": "Invalid credentials"
}
```

request body
```
{
    "Email": "youknowwho@gmail.com",
    "Password": "avadakadavra"
}
```

response body
```
{
    "token": "HEADER.PAYLOAD.SIGNATURE"
}
```

verify the token at [https://jwt.io/]

(PS. Do not forget to add current ip in mongodb atlas)

### /profile route

HTTP METHOD: GET
HEADER

request body
```
Key = Authorization
Value = Bearer <HEADER.PAYLOAD.SIGNATURE>
```
<HEADER.PAYLOAD.SIGNATURE> captured at the time of /login response body!

response body
```
{
    "created_at": "2024-09-29T06:02:19.782Z",
    "email": "youknowwho@gmail.com",
    "first_name": "Tom",
    "last_name": "Riddle",
    "updated_at": "2024-09-29T06:02:19.782Z"
}
```


## Todo
1. Add _id & email both while token creation
2. After creating user, should we add token into the database? If Yes/No why?
3. add otp 2 factor email otp authentication
4. logout route


## Personal Notes
To Add multiple claims
1. go to utils/jwtUtils.go
2. inside GenerateToken(args) change arguments
3. update the GenerateToken calls in controllers/authController.go
