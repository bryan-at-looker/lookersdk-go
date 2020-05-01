# RunningQueries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**User** | [**UserPublic**](UserPublic.md) |  | [optional] 
**Query** | [**Query**](Query.md) |  | [optional] 
**SqlQuery** | [**SqlQuery**](SqlQuery.md) |  | [optional] 
**Look** | [**LookBasic**](LookBasic.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | Date/Time Query was initiated | [optional] [readonly] 
**CompletedAt** | Pointer to **string** | Date/Time Query was completed | [optional] [readonly] 
**QueryId** | Pointer to **string** | Query Id | [optional] [readonly] 
**Source** | Pointer to **string** | Source (look, dashboard, queryrunner, explore, etc.) | [optional] [readonly] 
**NodeId** | Pointer to **string** | Node Id | [optional] [readonly] 
**Slug** | Pointer to **string** | Slug | [optional] [readonly] 
**QueryTaskId** | Pointer to **string** | ID of a Query Task | [optional] [readonly] 
**CacheKey** | Pointer to **string** | Cache Key | [optional] [readonly] 
**ConnectionName** | Pointer to **string** | Connection | [optional] [readonly] 
**Dialect** | Pointer to **string** | Dialect | [optional] [readonly] 
**ConnectionId** | Pointer to **string** | Connection ID | [optional] [readonly] 
**Message** | Pointer to **string** | Additional Information(Error message or verbose status) | [optional] [readonly] 
**Status** | Pointer to **string** | Status description | [optional] [readonly] 
**Runtime** | Pointer to **float64** | Number of seconds elapsed running the Query | [optional] [readonly] 
**Sql** | Pointer to **string** | SQL text of the query as run | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


