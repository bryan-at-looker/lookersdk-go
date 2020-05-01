# UserLoginLockout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Key** | Pointer to **string** | Hash of user&#39;s client id | [optional] [readonly] 
**AuthType** | Pointer to **string** | Authentication method for login failures | [optional] [readonly] 
**Ip** | Pointer to **string** | IP address of most recent failed attempt | [optional] [readonly] 
**UserId** | Pointer to **int64** | User ID | [optional] [readonly] 
**RemoteId** | Pointer to **string** | Remote ID of user if using LDAP | [optional] [readonly] 
**FullName** | Pointer to **string** | User&#39;s name | [optional] [readonly] 
**Email** | Pointer to **string** | Email address associated with the user&#39;s account | [optional] [readonly] 
**FailCount** | Pointer to **int64** | Number of failures that triggered the lockout | [optional] [readonly] 
**LockoutAt** | Pointer to [**time.Time**](time.Time.md) | Time when lockout was triggered | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


