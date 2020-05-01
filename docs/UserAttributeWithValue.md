# UserAttributeWithValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Name** | Pointer to **string** | Name of user attribute | [optional] [readonly] 
**Label** | Pointer to **string** | Human-friendly label for user attribute | [optional] [readonly] 
**Rank** | Pointer to **int64** | Precedence for setting value on user (lowest wins) | [optional] [readonly] 
**Value** | Pointer to **string** | Value of attribute for user | [optional] 
**UserId** | Pointer to **int64** | Id of User | [optional] [readonly] 
**UserCanEdit** | **bool** | Can the user set this value | [optional] [readonly] 
**ValueIsHidden** | **bool** | If true, the \&quot;value\&quot; field will be null, because the attribute settings block access to this value | [optional] [readonly] 
**UserAttributeId** | Pointer to **int64** | Id of User Attribute | [optional] [readonly] 
**Source** | Pointer to **string** | How user got this value for this attribute | [optional] [readonly] 
**HiddenValueDomainWhitelist** | Pointer to **string** | If this user attribute is hidden, whitelist of destinations to which it may be sent. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


