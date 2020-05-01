# LookmlModelExploreJoins

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of this join (and name of the view to join) | [optional] [readonly] 
**DependentFields** | Pointer to **[]string** | Fields referenced by the join | [optional] [readonly] 
**Fields** | **[]string** | Fields of the joined view to pull into this explore | [optional] [readonly] 
**ForeignKey** | **string** | Name of the dimension in this explore whose value is in the primary key of the joined view | [optional] [readonly] 
**From** | **string** | Name of view to join | [optional] [readonly] 
**OuterOnly** | Pointer to **bool** | Specifies whether all queries must use an outer join | [optional] [readonly] 
**Relationship** | Pointer to **string** | many_to_one, one_to_one, one_to_many, many_to_many | [optional] [readonly] 
**RequiredJoins** | **[]string** | Names of joins that must always be included in SQL queries | [optional] [readonly] 
**SqlForeignKey** | **string** | SQL expression that produces a foreign key | [optional] [readonly] 
**SqlOn** | **string** | SQL ON expression describing the join condition | [optional] [readonly] 
**SqlTableName** | **string** | SQL table name to join | [optional] [readonly] 
**Type** | **string** | The join type: left_outer, full_outer, inner, or cross | [optional] [readonly] 
**ViewLabel** | **string** | Label to display in UI selectors | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


