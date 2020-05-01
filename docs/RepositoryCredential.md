# RepositoryCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **string** | Unique Id | [optional] [readonly] 
**RootProjectId** | **string** | Root project Id | [optional] [readonly] 
**RemoteUrl** | **string** | Git remote repository url | [optional] [readonly] 
**GitUsername** | Pointer to **string** | Git username for HTTPS authentication. | [optional] 
**GitPassword** | Pointer to **string** | (Write-Only) Git password for HTTPS authentication. | [optional] 
**SshPublicKey** | Pointer to **string** | Public deploy key for SSH authentication. | [optional] 
**IsConfigured** | **bool** | Whether the credentials have been configured for the Git Repository. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


