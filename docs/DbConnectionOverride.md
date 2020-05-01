# DbConnectionOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | **string** | Context in which to override (&#x60;pdt&#x60; is the only allowed value) | [optional] 
**Host** | Pointer to **string** | Host name/address of server | [optional] 
**Port** | Pointer to **string** | Port number on server | [optional] 
**Username** | Pointer to **string** | Username for server authentication | [optional] 
**Password** | Pointer to **string** | (Write-Only) Password for server authentication | [optional] 
**HasPassword** | **bool** | Whether or not the password is overridden in this context | [optional] [readonly] 
**Certificate** | Pointer to **string** | (Write-Only) Base64 encoded Certificate body for server authentication (when appropriate for dialect). | [optional] 
**FileType** | Pointer to **string** | (Write-Only) Certificate keyfile type - .json or .p12 | [optional] 
**Database** | Pointer to **string** | Database name | [optional] 
**Schema** | Pointer to **string** | Scheme name | [optional] 
**JdbcAdditionalParams** | Pointer to **string** | Additional params to add to JDBC connection string | [optional] 
**AfterConnectStatements** | Pointer to **string** | SQL statements (semicolon separated) to issue after connecting to the database. Requires &#x60;custom_after_connect_statements&#x60; license feature | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


