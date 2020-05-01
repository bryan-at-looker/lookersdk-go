# Go API client for looker

### Authorization  

The Looker API uses Looker **API3** credentials for authorization and access control. Looker admins can create API3 credentials on Looker's **Admin/Users** page. Pass API3 credentials to the **_/login** endpoint to obtain a temporary access_token. Include that access_token in the Authorization header of Looker API requests. For details, see [Looker API Authorization](https://looker.com/docs/r/api/authorization)  

For Go, we have taken care of the login by including a LoginContext function.

## Environment variables

This SDK requires four environment variables to be set, you can see examples in [.env](./env.example) or below:

```shell
LOOKERSDK_BASE_URL=https://customer.looker.com:19999
LOOKERSDK_API_VERSION=3.1
LOOKERSDK_CLIENT_ID=abc123
LOOKERSDK_CLIENT_SECRET=zyx098
```

LOOKERSDK_CLIENT_ID and LOOKERSDK_CLIENT_SECRET can be given to you by your Looker admin.

## Use

Install the following dependencies:

```shell
go get github.com/bryan-at-looker/lookersdkgo
```

Put the package under your project folder and add the following in import:

```golang
import "lookersdkgo"
```

create a client and login

```
client := lookersdkgo.NewAPIClient()
ctx := lookersdkgo.LoginContext(context.TODO(), client)
```

Test an API Call

```golang
me, _, err := client.UserApi.Me(ctx, nil)
if err != nil {
    log.Fatal(err)
}
log.Println(*me.DisplayName)
```