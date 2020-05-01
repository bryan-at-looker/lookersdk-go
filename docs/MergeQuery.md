# MergeQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**ColumnLimit** | Pointer to **string** | Column Limit | [optional] 
**DynamicFields** | Pointer to **string** | Dynamic Fields | [optional] 
**Id** | **string** | Unique Id | [optional] [readonly] 
**Pivots** | Pointer to **[]string** | Pivots | [optional] 
**ResultMakerId** | Pointer to **int64** | Unique to get results | [optional] [readonly] 
**Sorts** | Pointer to **[]string** | Sorts | [optional] 
**SourceQueries** | Pointer to [**[]MergeQuerySourceQuery**](MergeQuerySourceQuery.md) | Source Queries defining the results to be merged. | [optional] 
**Total** | **bool** | Total | [optional] 
**VisConfig** | Pointer to **map[string]string** | Visualization Config | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


