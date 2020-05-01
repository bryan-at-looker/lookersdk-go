# OidcConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**AlternateEmailLoginAllowed** | **bool** | Allow alternate email-based login via &#39;/login/email&#39; for admins and for specified users with the &#39;login_special_email&#39; permission. This option is useful as a fallback during ldap setup, if ldap config problems occur later, or if you need to support some users who are not in your ldap directory. Looker email/password logins are always disabled for regular users when ldap is enabled. | [optional] 
**Audience** | Pointer to **string** | OpenID Provider Audience | [optional] 
**AuthRequiresRole** | **bool** | Users will not be allowed to login at all unless a role for them is found in OIDC if set to true | [optional] 
**AuthorizationEndpoint** | Pointer to **string** | OpenID Provider Authorization Url | [optional] 
**DefaultNewUserGroupIds** | Pointer to **[]int64** | (Write-Only) Array of ids of groups that will be applied to new users the first time they login via OIDC | [optional] 
**DefaultNewUserGroups** | Pointer to [**[]Group**](Group.md) | (Read-only) Groups that will be applied to new users the first time they login via OIDC | [optional] [readonly] 
**DefaultNewUserRoleIds** | Pointer to **[]int64** | (Write-Only) Array of ids of roles that will be applied to new users the first time they login via OIDC | [optional] 
**DefaultNewUserRoles** | Pointer to [**[]Role**](Role.md) | (Read-only) Roles that will be applied to new users the first time they login via OIDC | [optional] [readonly] 
**Enabled** | **bool** | Enable/Disable OIDC authentication for the server | [optional] 
**Groups** | Pointer to [**[]OidcGroupRead**](OIDCGroupRead.md) | (Read-only) Array of mappings between OIDC Groups and Looker Roles | [optional] [readonly] 
**GroupsAttribute** | Pointer to **string** | Name of user record attributes used to indicate groups. Used when &#39;groups_finder_type&#39; is set to &#39;grouped_attribute_values&#39; | [optional] 
**GroupsWithRoleIds** | Pointer to [**[]OidcGroupWrite**](OIDCGroupWrite.md) | (Read/Write) Array of mappings between OIDC Groups and arrays of Looker Role ids | [optional] 
**Identifier** | Pointer to **string** | Relying Party Identifier (provided by OpenID Provider) | [optional] 
**Issuer** | Pointer to **string** | OpenID Provider Issuer | [optional] 
**ModifiedAt** | Pointer to [**time.Time**](time.Time.md) | When this config was last modified | [optional] [readonly] 
**ModifiedBy** | Pointer to **int64** | User id of user who last modified this config | [optional] [readonly] 
**NewUserMigrationTypes** | Pointer to **string** | Merge first-time oidc login to existing user account by email addresses. When a user logs in for the first time via oidc this option will connect this user into their existing account by finding the account with a matching email address by testing the given types of credentials for existing users. Otherwise a new user account will be created for the user. This list (if provided) must be a comma separated list of string like &#39;email,ldap,google&#39; | [optional] 
**Scopes** | Pointer to **[]string** | Array of scopes to request. | [optional] 
**Secret** | Pointer to **string** | (Write-Only) Relying Party Secret (provided by OpenID Provider) | [optional] 
**SetRolesFromGroups** | **bool** | Set user roles in Looker based on groups from OIDC | [optional] 
**TestSlug** | Pointer to **string** | Slug to identify configurations that are created in order to run a OIDC config test | [optional] [readonly] 
**TokenEndpoint** | Pointer to **string** | OpenID Provider Token Url | [optional] 
**UserAttributeMapEmail** | Pointer to **string** | Name of user record attributes used to indicate email address field | [optional] 
**UserAttributeMapFirstName** | Pointer to **string** | Name of user record attributes used to indicate first name | [optional] 
**UserAttributeMapLastName** | Pointer to **string** | Name of user record attributes used to indicate last name | [optional] 
**UserAttributes** | Pointer to [**[]OidcUserAttributeRead**](OIDCUserAttributeRead.md) | (Read-only) Array of mappings between OIDC User Attributes and Looker User Attributes | [optional] [readonly] 
**UserAttributesWithIds** | Pointer to [**[]OidcUserAttributeWrite**](OIDCUserAttributeWrite.md) | (Read/Write) Array of mappings between OIDC User Attributes and arrays of Looker User Attribute ids | [optional] 
**UserinfoEndpoint** | Pointer to **string** | OpenID Provider User Information Url | [optional] 
**Url** | Pointer to **string** | Link to get this item | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


