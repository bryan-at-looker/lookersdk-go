# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**AvatarUrl** | Pointer to **string** | URL for the avatar image (may be generic) | [optional] [readonly] 
**AvatarUrlWithoutSizing** | Pointer to **string** | URL for the avatar image (may be generic), does not specify size | [optional] [readonly] 
**CredentialsApi3** | Pointer to [**[]CredentialsApi3**](CredentialsApi3.md) | API 3 credentials | [optional] [readonly] 
**CredentialsEmail** | [**CredentialsEmail**](CredentialsEmail.md) |  | [optional] 
**CredentialsEmbed** | Pointer to [**[]CredentialsEmbed**](CredentialsEmbed.md) | Embed credentials | [optional] [readonly] 
**CredentialsGoogle** | [**CredentialsGoogle**](CredentialsGoogle.md) |  | [optional] 
**CredentialsLdap** | [**CredentialsLdap**](CredentialsLDAP.md) |  | [optional] 
**CredentialsLookerOpenid** | [**CredentialsLookerOpenid**](CredentialsLookerOpenid.md) |  | [optional] 
**CredentialsOidc** | [**CredentialsOidc**](CredentialsOIDC.md) |  | [optional] 
**CredentialsSaml** | [**CredentialsSaml**](CredentialsSaml.md) |  | [optional] 
**CredentialsTotp** | [**CredentialsTotp**](CredentialsTotp.md) |  | [optional] 
**DisplayName** | Pointer to **string** | Full name for display (available only if both first_name and last_name are set) | [optional] [readonly] 
**Email** | Pointer to **string** | EMail address | [optional] [readonly] 
**EmbedGroupSpaceId** | Pointer to **int64** | (Embed only) ID of user&#39;s group space based on the external_group_id optionally specified during embed user login | [optional] [readonly] 
**FirstName** | Pointer to **string** | First name | [optional] 
**GroupIds** | Pointer to **[]int64** | Array of ids of the groups for this user | [optional] [readonly] 
**HomeSpaceId** | Pointer to **string** | ID string for user&#39;s home space | [optional] 
**HomeFolderId** | Pointer to **string** | ID string for user&#39;s home folder | [optional] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**IsDisabled** | **bool** | Account has been disabled | [optional] 
**LastName** | Pointer to **string** | Last name | [optional] 
**Locale** | Pointer to **string** | User&#39;s preferred locale. User locale takes precedence over Looker&#39;s system-wide default locale. Locale determines language of display strings and date and numeric formatting in API responses. Locale string must be a 2 letter language code or a combination of language code and region code: &#39;en&#39; or &#39;en-US&#39;, for example. | [optional] 
**LookerVersions** | Pointer to **[]string** | Array of strings representing the Looker versions that this user has used (this only goes back as far as &#39;3.54.0&#39;) | [optional] [readonly] 
**ModelsDirValidated** | Pointer to **bool** | User&#39;s dev workspace has been checked for presence of applicable production projects | [optional] 
**PersonalSpaceId** | Pointer to **int64** | ID of user&#39;s personal space | [optional] [readonly] 
**PersonalFolderId** | Pointer to **int64** | ID of user&#39;s personal folder | [optional] [readonly] 
**PresumedLookerEmployee** | **bool** | User is identified as an employee of Looker | [optional] [readonly] 
**RoleIds** | Pointer to **[]int64** | Array of ids of the roles for this user | [optional] [readonly] 
**Sessions** | Pointer to [**[]Session**](Session.md) | Active sessions | [optional] [readonly] 
**UiState** | Pointer to **map[string]string** | Per user dictionary of undocumented state information owned by the Looker UI. | [optional] 
**VerifiedLookerEmployee** | **bool** | User is identified as an employee of Looker who has been verified via Looker corporate authentication | [optional] [readonly] 
**RolesExternallyManaged** | **bool** | User&#39;s roles are managed by an external directory like SAML or LDAP and can not be changed directly. | [optional] [readonly] 
**Url** | Pointer to **string** | Link to get this item | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


