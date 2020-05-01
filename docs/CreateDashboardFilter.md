# CreateDashboardFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Id | [optional] [readonly] 
**DashboardId** | Pointer to **string** | Id of Dashboard | 
**Name** | Pointer to **string** | Name of filter | 
**Title** | Pointer to **string** | Title of filter | 
**Type** | Pointer to **string** | Type of filter: one of date, number, string, or field | 
**DefaultValue** | Pointer to **string** | Default value of filter | [optional] 
**Model** | Pointer to **string** | Model of filter (required if type &#x3D; field) | [optional] 
**Explore** | Pointer to **string** | Explore of filter (required if type &#x3D; field) | [optional] 
**Dimension** | Pointer to **string** | Dimension of filter (required if type &#x3D; field) | [optional] 
**Field** | Pointer to **map[string]string** | Field information | [optional] [readonly] 
**Row** | Pointer to **int64** | Display order of this filter relative to other filters | [optional] 
**ListensToFilters** | Pointer to **[]string** | Array of listeners for faceted filters | [optional] 
**AllowMultipleValues** | **bool** | Whether the filter allows multiple filter values | [optional] 
**Required** | **bool** | Whether the filter requires a value to run the dashboard | [optional] 
**UiConfig** | Pointer to **map[string]string** | The visual configuration for this filter. Used to set up how the UI for this filter should appear. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


