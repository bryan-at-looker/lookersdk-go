# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**ContentFavoriteId** | Pointer to **int64** | Content Favorite Id | [optional] [readonly] 
**ContentMetadataId** | Pointer to **int64** | Id of content metadata | [optional] [readonly] 
**Description** | Pointer to **string** | Description | [optional] 
**Hidden** | **bool** | Is Hidden | [optional] 
**Id** | **string** | Unique Id | [optional] [readonly] 
**Model** | [**LookModel**](LookModel.md) |  | [optional] 
**QueryTimezone** | Pointer to **string** | Timezone in which the Dashboard will run by default. | [optional] 
**Readonly** | **bool** | Is Read-only | [optional] [readonly] 
**RefreshInterval** | Pointer to **string** | Refresh Interval, as a time duration phrase like \&quot;2 hours 30 minutes\&quot;. A number with no time units will be interpreted as whole seconds. | [optional] 
**RefreshIntervalToI** | Pointer to **int64** | Refresh Interval in milliseconds | [optional] [readonly] 
**Space** | [**SpaceBase**](SpaceBase.md) |  | [optional] 
**Folder** | [**FolderBase**](FolderBase.md) |  | [optional] 
**Title** | Pointer to **string** | Dashboard Title | [optional] 
**UserId** | Pointer to **int64** | Id of User | [optional] [readonly] 
**BackgroundColor** | Pointer to **string** | Background color | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Time that the Dashboard was created. | [optional] [readonly] 
**CrossfilterEnabled** | **bool** | Enables crossfiltering in dashboards - only available in dashboards-next (beta) | [optional] 
**DashboardElements** | Pointer to [**[]DashboardElement**](DashboardElement.md) | Elements | [optional] [readonly] 
**DashboardFilters** | Pointer to [**[]DashboardFilter**](DashboardFilter.md) | Filters | [optional] [readonly] 
**DashboardLayouts** | Pointer to [**[]DashboardLayout**](DashboardLayout.md) | Layouts | [optional] [readonly] 
**Deleted** | **bool** | Whether or not a dashboard is &#39;soft&#39; deleted. | [optional] 
**DeletedAt** | Pointer to [**time.Time**](time.Time.md) | Time that the Dashboard was &#39;soft&#39; deleted. | [optional] [readonly] 
**DeleterId** | Pointer to **int64** | Id of User that &#39;soft&#39; deleted the dashboard. | [optional] [readonly] 
**EditUri** | Pointer to **string** | Relative path of URI of LookML file to edit the dashboard (LookML dashboard only). | [optional] [readonly] 
**FavoriteCount** | Pointer to **int64** | Number of times favorited | [optional] [readonly] 
**LastAccessedAt** | Pointer to [**time.Time**](time.Time.md) | Time the dashboard was last accessed | [optional] [readonly] 
**LastViewedAt** | Pointer to [**time.Time**](time.Time.md) | Time last viewed in the Looker web UI | [optional] [readonly] 
**LoadConfiguration** | Pointer to **string** | configuration option that governs how dashboard loading will happen. | [optional] 
**LookmlLinkId** | Pointer to **string** | Links this dashboard to a particular LookML dashboard such that calling a **sync** operation on that LookML dashboard will update this dashboard to match. | [optional] 
**ShowFiltersBar** | **bool** | Show filters bar.  **Security Note:** This property only affects the *cosmetic* appearance of the dashboard, not a user&#39;s ability to access data. Hiding the filters bar does **NOT** prevent users from changing filters by other means. For information on how to set up secure data access control policies, see [Control User Access to Data](https://looker.com/docs/r/api/control-access) | [optional] 
**ShowTitle** | **bool** | Show title | [optional] 
**Slug** | Pointer to **string** | Content Metadata Slug | [optional] 
**SpaceId** | Pointer to **string** | Id of Space | [optional] 
**FolderId** | Pointer to **string** | Id of folder | [optional] 
**TextTileTextColor** | Pointer to **string** | Color of text on text tiles | [optional] 
**TileBackgroundColor** | Pointer to **string** | Tile background color | [optional] 
**TileTextColor** | Pointer to **string** | Tile text color | [optional] 
**TitleColor** | Pointer to **string** | Title color | [optional] 
**ViewCount** | Pointer to **int64** | Number of times viewed in the Looker web UI | [optional] [readonly] 
**Appearance** | [**DashboardAppearance**](DashboardAppearance.md) |  | [optional] 
**PreferredViewer** | Pointer to **string** | The preferred route for viewing this dashboard (ie: dashboards or dashboards-next) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


