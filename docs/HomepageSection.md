# HomepageSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Time at which this section was created. | [optional] [readonly] 
**DeletedAt** | Pointer to [**time.Time**](time.Time.md) | Time at which this section was deleted. | [optional] 
**DetailUrl** | Pointer to **string** | A URL pointing to a page showing further information about the content in the section. | [optional] [readonly] 
**HomepageId** | Pointer to **int64** | Id reference to parent homepage | [optional] 
**HomepageItems** | Pointer to [**[]HomepageItem**](HomepageItem.md) | Items in the homepage section | [optional] [readonly] 
**Id** | **string** | Unique Id | [optional] [readonly] 
**IsHeader** | **bool** | Is this a header section (has no items) | [optional] [readonly] 
**ItemOrder** | Pointer to **[]int64** | ids of the homepage items in the order they should be displayed | [optional] 
**Title** | Pointer to **string** | Name of row | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | Time at which this section was last updated. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the content found in this section. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


