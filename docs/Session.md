# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**IpAddress** | Pointer to **string** | IP address of user when this session was initiated | [optional] [readonly] 
**Browser** | Pointer to **string** | User&#39;s browser type | [optional] [readonly] 
**OperatingSystem** | Pointer to **string** | User&#39;s Operating System | [optional] [readonly] 
**City** | Pointer to **string** | City component of user location (derived from IP address) | [optional] [readonly] 
**State** | Pointer to **string** | State component of user location (derived from IP address) | [optional] [readonly] 
**Country** | Pointer to **string** | Country component of user location (derived from IP address) | [optional] [readonly] 
**CredentialsType** | Pointer to **string** | Type of credentials used for logging in this session | [optional] [readonly] 
**ExtendedAt** | Pointer to **string** | Time when this session was last extended by the user | [optional] [readonly] 
**ExtendedCount** | Pointer to **int64** | Number of times this session was extended | [optional] [readonly] 
**SudoUserId** | Pointer to **int64** | Actual user in the case when this session represents one user sudo&#39;ing as another | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | Time when this session was initiated | [optional] [readonly] 
**ExpiresAt** | Pointer to **string** | Time when this session will expire | [optional] [readonly] 
**Url** | Pointer to **string** | Link to get this item | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


