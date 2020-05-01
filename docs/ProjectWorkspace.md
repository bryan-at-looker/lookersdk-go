# ProjectWorkspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**ProjectId** | Pointer to **string** | The id of the project | [optional] [readonly] 
**WorkspaceId** | Pointer to **string** | The id of the local workspace containing the project files | [optional] [readonly] 
**GitStatus** | Pointer to **string** | The status of the local git directory | [optional] [readonly] 
**GitHead** | Pointer to **string** | Git head revision name | [optional] [readonly] 
**GitBranch** | [**GitBranch**](GitBranch.md) |  | [optional] 
**LookmlType** | Pointer to **string** | The lookml syntax used by all files in this project | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


