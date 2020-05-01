# RenderTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | Date/Time render task was created | [optional] [readonly] 
**DashboardFilters** | Pointer to **string** | Filter values to apply to the dashboard queries, in URL query format | [optional] [readonly] 
**DashboardId** | Pointer to **int64** | Id of dashboard to render | [optional] [readonly] 
**DashboardStyle** | Pointer to **string** | Dashboard layout style: single_column or tiled | [optional] [readonly] 
**FinalizedAt** | Pointer to **string** | Date/Time render task was completed | [optional] [readonly] 
**Height** | Pointer to **int64** | Output height in pixels. Flowed layouts may ignore this value. | [optional] [readonly] 
**Id** | **string** | Id of this render task | [optional] [readonly] 
**LookId** | Pointer to **int64** | Id of look to render | [optional] [readonly] 
**LookmlDashboardId** | Pointer to **string** | Id of lookml dashboard to render | [optional] [readonly] 
**QueryId** | Pointer to **int64** | Id of query to render | [optional] [readonly] 
**QueryRuntime** | Pointer to **float64** | Number of seconds elapsed running queries | [optional] [readonly] 
**RenderRuntime** | Pointer to **float64** | Number of seconds elapsed rendering data | [optional] [readonly] 
**ResultFormat** | Pointer to **string** | Output format: pdf, png, or jpg | [optional] [readonly] 
**Runtime** | Pointer to **float64** | Total seconds elapsed for render task | [optional] [readonly] 
**Status** | Pointer to **string** | Render task status: enqueued_for_query, querying, enqueued_for_render, rendering, success, failure | [optional] [readonly] 
**StatusDetail** | Pointer to **string** | Additional information about the current status | [optional] [readonly] 
**UserId** | Pointer to **int64** | The user account permissions in which the render task will execute | [optional] [readonly] 
**Width** | Pointer to **int64** | Output width in pixels | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


