# QueryTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **string** | Unique Id | [optional] [readonly] 
**QueryId** | Pointer to **int64** | Id of query | [optional] 
**Query** | [**Query**](Query.md) |  | [optional] 
**GenerateLinks** | **bool** | whether or not to generate links in the query response. | [optional] 
**ForceProduction** | **bool** | Use production models to run query (even is user is in dev mode). | [optional] 
**PathPrefix** | Pointer to **string** | Prefix to use for drill links. | [optional] 
**Cache** | **bool** | Whether or not to use the cache | [optional] 
**ServerTableCalcs** | **bool** | Whether or not to run table calculations on the server | [optional] 
**CacheOnly** | **bool** | Retrieve any results from cache even if the results have expired. | [optional] 
**CacheKey** | Pointer to **string** | cache key used to cache query. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of query task. | [optional] 
**Source** | Pointer to **string** | Source of query task. | [optional] 
**Runtime** | Pointer to **float32** | Runtime of prior queries. | [optional] [readonly] 
**RebuildPdts** | **bool** | Rebuild PDTS used in query. | [optional] 
**ResultSource** | Pointer to **string** | Source of the results of the query. | [optional] [readonly] 
**LookId** | Pointer to **int64** | Id of look associated with query. | [optional] 
**DashboardId** | Pointer to **string** | Id of dashboard associated with query. | [optional] 
**ResultFormat** | Pointer to **string** | The data format of the query results. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


