# DashboardBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**ContentFavoriteId** | Pointer to **int64** | Content Favorite Id | [optional] [readonly] 
**ContentMetadataId** | Pointer to **int64** | Id of content metadata | [optional] [readonly] 
**Description** | Pointer to **string** | Description | [optional] [readonly] 
**Hidden** | **bool** | Is Hidden | [optional] [readonly] 
**Id** | **string** | Unique Id | [optional] [readonly] 
**Model** | [**LookModel**](LookModel.md) |  | [optional] 
**QueryTimezone** | Pointer to **string** | Timezone in which the Dashboard will run by default. | [optional] [readonly] 
**Readonly** | **bool** | Is Read-only | [optional] [readonly] 
**RefreshInterval** | Pointer to **string** | Refresh Interval, as a time duration phrase like \&quot;2 hours 30 minutes\&quot;. A number with no time units will be interpreted as whole seconds. | [optional] [readonly] 
**RefreshIntervalToI** | Pointer to **int64** | Refresh Interval in milliseconds | [optional] [readonly] 
**Space** | [**SpaceBase**](SpaceBase.md) |  | [optional] 
**Folder** | [**FolderBase**](FolderBase.md) |  | [optional] 
**Title** | Pointer to **string** | Dashboard Title | [optional] [readonly] 
**UserId** | Pointer to **int64** | Id of User | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


