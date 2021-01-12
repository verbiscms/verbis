# Auth

Verbis uses secure token authentication for all routes that are not public. Fresh tokens are generated when after a
successful login and will be regenerated once the user logs out, or the session expires.

📢 **Storing the token:** Private endpoints should feature a header with the key `token` to authorise requests.

There are a collection of `Auth` endpoints described below that help with password resets as wells as authentication.

## Login

Logging in requires an email and password to be sent, if the user is authenticated a token will be sent back to **store
for future requests**.

Required fields:
- `email`
- `password`

👉 `POST` to `/login`

```json
{
	"email": "hello@verbiscms.com",
	"password": "mypassword"
}
```

**Example Response**

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully logged in & session started",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
	},
	"data": {
		"id": 1,
		"uuid": "7d14f9da-5412-11eb-ae93-0242ac130002",
		"first_name": "Verbis",
		"last_name": "CMS",
		"email": "hello@verbiscms.com",
		"facebook": null,
		"twitter": null,
		"linked_in": null,
		"instagram": null,
		"biography": null,
		"role": {
			"id": 0,
			"name": "",
			"description": ""
		},
		"profile_picture_id": null,
		"email_verified_at": null,
		"created_at": "2020-01-01T12:00:00Z",
		"updated_at": "2020-01-01T12:00:00Z",
		"token": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"token_last_used": "2020-01-01T12:00:00Z"
	}
}
```

## Logout

Logs the user out everywhere by obtaining the `token` key. As such, it doesn't require any JSON body.

Required fields:
- `token`

👉 `POST` to `/logout`

**Example Response**

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully logged out",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
	},
	"data": {}
}
```



## Send Password Reset

Sends a password reset email to the user if they have forgotten their password. This route creates a unique token within
the database for secure verification.

Required fields:
- `email`

```json
{
	"email": "hello@verbiscms.com"
}
```

👉 `POST` to `/password/email`

**Example Response**

```json
{
	"status": 200,
	"error": false,
	"message": "A fresh verification link has been sent to your email",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
	},
	"data": {}
}
```

## Verify Password Token

As Verbis is headless, meaning a completely separate admin interface is used to interact with the API, password token's
can be verified using this endpoint. The token is passed in the URL, and the endpoint returns a 404 if the token could
not be verified.

👉 `POST` to `password/verify/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`

**Example Response**

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully verified token",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
	},
	"data": {}
}
```

## Password Reset

Resets a users' password, after they have clicked the verification email link. `new_password` and `confirm_password`
need to be equal and at least 8 characters in length. The token passed from the email also needs to be valid in order to
reset the users password. This can be obtained from the last part of the URL.

Required fields:
- `new_password`
- `confirm_password`
- `token`

👉 `POST` to `/password/reset`

```json
{
	"new_password": "mypassword",
	"confirm_password": "mypassword",
	"token": "f60a267969416107c68c6133ff00d88b"
}
```

**Example Response**

```json
{
	"status": 200,
	"error": false,
	"message": "Successfully reset password",
	"meta": {
		"request_time": "2021-01-01 12:00:00.000000 +0000 UTC",
		"response_time": "2021-01-01 12:00:20.200000 +0000 UTC",
		"latency_time": "20.000ms"
	},
	"data": {}
}
```