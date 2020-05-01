# LdapConfigTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **string** | Additional details for error cases | [optional] [readonly] 
**Issues** | Pointer to [**[]LdapConfigTestIssue**](LDAPConfigTestIssue.md) | Array of issues/considerations about the result | [optional] [readonly] 
**Message** | Pointer to **string** | Short human readable test about the result | [optional] [readonly] 
**Status** | Pointer to **string** | Test status code: always &#39;success&#39; or &#39;error&#39; | [optional] [readonly] 
**Trace** | Pointer to **string** | A more detailed trace of incremental results during auth tests | [optional] [readonly] 
**User** | [**LdapUser**](LDAPUser.md) |  | [optional] 
**Url** | Pointer to **string** | Link to ldap config | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


