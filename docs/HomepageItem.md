# HomepageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**ContentCreatedBy** | Pointer to **string** | Name of user who created the content this item is based on | [optional] [readonly] 
**ContentFavoriteId** | Pointer to **int64** | Content favorite id associated with the item this content is based on | [optional] [readonly] 
**ContentMetadataId** | Pointer to **int64** | Content metadata id associated with the item this content is based on | [optional] [readonly] 
**ContentUpdatedAt** | Pointer to **string** | Last time the content that this item is based on was updated | [optional] [readonly] 
**CustomDescription** | Pointer to **string** | Custom description entered by the user, if present | [optional] 
**CustomImageDataBase64** | Pointer to **string** | (Write-Only) base64 encoded image data | [optional] 
**CustomImageUrl** | Pointer to **string** | Custom image_url entered by the user, if present | [optional] [readonly] 
**CustomTitle** | Pointer to **string** | Custom title entered by the user, if present | [optional] 
**CustomUrl** | Pointer to **string** | Custom url entered by the user, if present | [optional] 
**DashboardId** | Pointer to **int64** | Dashboard to base this item on | [optional] 
**Description** | Pointer to **string** | The actual description for display | [optional] [readonly] 
**FavoriteCount** | Pointer to **int64** | Number of times content has been favorited, if present | [optional] [readonly] 
**HomepageSectionId** | Pointer to **string** | Associated Homepage Section | [optional] 
**Id** | **string** | Unique Id | [optional] [readonly] 
**ImageUrl** | Pointer to **string** | The actual image_url for display | [optional] [readonly] 
**Location** | Pointer to **string** | The container folder name of the content | [optional] [readonly] 
**LookId** | Pointer to **int64** | Look to base this item on | [optional] 
**LookmlDashboardId** | Pointer to **string** | LookML Dashboard to base this item on | [optional] 
**Order** | Pointer to **int64** | An arbitrary integer representing the sort order within the section | [optional] 
**SectionFetchTime** | Pointer to **float32** | Number of seconds it took to fetch the section this item is in | [optional] [readonly] 
**Title** | Pointer to **string** | The actual title for display | [optional] [readonly] 
**Url** | Pointer to **string** | The actual url for display | [optional] [readonly] 
**UseCustomDescription** | **bool** | Whether the custom description should be used instead of the content description, if the item is associated with content | [optional] 
**UseCustomImage** | **bool** | Whether the custom image should be used instead of the content image, if the item is associated with content | [optional] 
**UseCustomTitle** | **bool** | Whether the custom title should be used instead of the content title, if the item is associated with content | [optional] 
**UseCustomUrl** | **bool** | Whether the custom url should be used instead of the content url, if the item is associated with content | [optional] 
**ViewCount** | Pointer to **int64** | Number of times content has been viewed, if present | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


