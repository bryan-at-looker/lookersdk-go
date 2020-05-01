# LookmlModelExplore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Fully qualified name model plus explore name | [optional] [readonly] 
**Name** | Pointer to **string** | Explore name | [optional] [readonly] 
**Description** | **string** | Description | [optional] [readonly] 
**Label** | Pointer to **string** | Label | [optional] [readonly] 
**Scopes** | Pointer to **[]string** | Scopes | [optional] [readonly] 
**CanTotal** | **bool** | Can Total | [optional] [readonly] 
**CanSave** | **bool** | Can Save | [optional] [readonly] 
**CanExplain** | **bool** | Can Explain | [optional] [readonly] 
**CanPivotInDb** | **bool** | Can pivot in the DB | [optional] [readonly] 
**CanSubtotal** | **bool** | Can use subtotals | [optional] [readonly] 
**HasTimezoneSupport** | **bool** | Has timezone support | [optional] [readonly] 
**SupportsCostEstimate** | **bool** | Cost estimates supported | [optional] [readonly] 
**ConnectionName** | Pointer to **string** | Connection name | [optional] [readonly] 
**NullSortTreatment** | Pointer to **string** | How nulls are sorted, possible values are \&quot;low\&quot;, \&quot;high\&quot;, \&quot;first\&quot; and \&quot;last\&quot; | [optional] [readonly] 
**Files** | Pointer to **[]string** | List of model source files | [optional] [readonly] 
**SourceFile** | Pointer to **string** | Primary source_file file | [optional] [readonly] 
**ProjectName** | Pointer to **string** | Name of project | [optional] [readonly] 
**ModelName** | Pointer to **string** | Name of model | [optional] [readonly] 
**ViewName** | Pointer to **string** | Name of view | [optional] [readonly] 
**Hidden** | **bool** | Is hidden | [optional] [readonly] 
**SqlTableName** | Pointer to **string** | A sql_table_name expression that defines what sql table the view/explore maps onto. Example: \&quot;prod_orders2 AS orders\&quot; in a view named orders. | [optional] [readonly] 
**AccessFilterFields** | Pointer to **[]string** | (DEPRECATED) Array of access filter field names | [optional] [readonly] 
**AccessFilters** | Pointer to [**[]LookmlModelExploreAccessFilter**](LookmlModelExploreAccessFilter.md) | Access filters | [optional] [readonly] 
**Aliases** | Pointer to [**[]LookmlModelExploreAlias**](LookmlModelExploreAlias.md) | Aliases | [optional] [readonly] 
**AlwaysFilter** | Pointer to [**[]LookmlModelExploreAlwaysFilter**](LookmlModelExploreAlwaysFilter.md) | Always filter | [optional] [readonly] 
**ConditionallyFilter** | Pointer to [**[]LookmlModelExploreConditionallyFilter**](LookmlModelExploreConditionallyFilter.md) | Conditionally filter | [optional] [readonly] 
**IndexFields** | Pointer to **[]string** | Array of index fields | [optional] [readonly] 
**Sets** | Pointer to [**[]LookmlModelExploreSet**](LookmlModelExploreSet.md) | Sets | [optional] [readonly] 
**Tags** | Pointer to **[]string** | An array of arbitrary string tags provided in the model for this explore. | [optional] [readonly] 
**Errors** | Pointer to [**[]LookmlModelExploreError**](LookmlModelExploreError.md) | Errors | [optional] [readonly] 
**Fields** | [**LookmlModelExploreFieldset**](LookmlModelExploreFieldset.md) |  | [optional] 
**Joins** | Pointer to [**[]LookmlModelExploreJoins**](LookmlModelExploreJoins.md) | Views joined into this explore | [optional] [readonly] 
**GroupLabel** | Pointer to **string** | Label used to group explores in the navigation menus | [optional] [readonly] 
**SupportedMeasureTypes** | [**[]LookmlModelExploreSupportedMeasureType**](LookmlModelExploreSupportedMeasureType.md) | An array of items describing which custom measure types are supported for creating a custom measure &#39;baed_on&#39; each possible dimension type. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


