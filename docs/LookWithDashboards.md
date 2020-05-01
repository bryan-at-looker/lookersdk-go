# LookWithDashboards

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**ContentMetadataId** | Pointer to **int64** | Id of content metadata | [optional] [readonly] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**Title** | Pointer to **string** | Look Title | [optional] 
**ContentFavoriteId** | Pointer to **int64** | Content Favorite Id | [optional] [readonly] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Time that the Look was created. | [optional] [readonly] 
**Deleted** | **bool** | Whether or not a look is &#39;soft&#39; deleted. | [optional] 
**DeletedAt** | Pointer to [**time.Time**](time.Time.md) | Time that the Look was deleted. | [optional] [readonly] 
**DeleterId** | Pointer to **int64** | Id of User that deleted the look. | [optional] [readonly] 
**Description** | Pointer to **string** | Description | [optional] 
**EmbedUrl** | Pointer to **string** | Embed Url | [optional] [readonly] 
**ExcelFileUrl** | Pointer to **string** | Excel File Url | [optional] [readonly] 
**FavoriteCount** | Pointer to **int64** | Number of times favorited | [optional] [readonly] 
**GoogleSpreadsheetFormula** | Pointer to **string** | Google Spreadsheet Formula | [optional] [readonly] 
**ImageEmbedUrl** | Pointer to **string** | Image Embed Url | [optional] [readonly] 
**IsRunOnLoad** | **bool** | auto-run query when Look viewed | [optional] 
**LastAccessedAt** | Pointer to [**time.Time**](time.Time.md) | Time that the Look was last accessed by any user | [optional] [readonly] 
**LastUpdaterId** | Pointer to **int64** | Id of User that last updated the look. | [optional] [readonly] 
**LastViewedAt** | Pointer to [**time.Time**](time.Time.md) | Time last viewed in the Looker web UI | [optional] [readonly] 
**Model** | [**LookModel**](LookModel.md) |  | [optional] 
**Public** | Pointer to **bool** | Is Public | [optional] 
**PublicSlug** | Pointer to **string** | Public Slug | [optional] [readonly] 
**PublicUrl** | Pointer to **string** | Public Url | [optional] [readonly] 
**QueryId** | Pointer to **int64** | Query Id | [optional] 
**ShortUrl** | Pointer to **string** | Short Url | [optional] [readonly] 
**Space** | [**SpaceBase**](SpaceBase.md) |  | [optional] 
**Folder** | [**FolderBase**](FolderBase.md) |  | [optional] 
**SpaceId** | Pointer to **string** | Space Id | [optional] 
**FolderId** | Pointer to **string** | Folder Id | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | Time that the Look was updated. | [optional] [readonly] 
**UserId** | Pointer to **int64** | User Id | [optional] 
**ViewCount** | Pointer to **int64** | Number of times viewed in the Looker web UI | [optional] [readonly] 
**User** | [**UserIdOnly**](UserIdOnly.md) |  | [optional] 
**Dashboards** | Pointer to [**[]DashboardBase**](DashboardBase.md) | Dashboards | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


