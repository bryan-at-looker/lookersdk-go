# LdapUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllEmails** | Pointer to **[]string** | Array of user&#39;s email addresses and aliases for use in migration | [optional] [readonly] 
**Attributes** | Pointer to **map[string]string** | Dictionary of user&#39;s attributes (name/value) | [optional] [readonly] 
**Email** | Pointer to **string** | Primary email address | [optional] [readonly] 
**FirstName** | Pointer to **string** | First name | [optional] [readonly] 
**Groups** | Pointer to **[]string** | Array of user&#39;s groups (group names only) | [optional] [readonly] 
**LastName** | Pointer to **string** | Last Name | [optional] [readonly] 
**LdapDn** | Pointer to **string** | LDAP&#39;s distinguished name for the user record | [optional] [readonly] 
**LdapId** | Pointer to **string** | LDAP&#39;s Unique ID for the user | [optional] [readonly] 
**Roles** | Pointer to **[]string** | Array of user&#39;s roles (role names only) | [optional] [readonly] 
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Url** | Pointer to **string** | Link to ldap config | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


