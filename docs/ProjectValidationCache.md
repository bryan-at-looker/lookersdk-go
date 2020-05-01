# ProjectValidationCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ProjectError**](ProjectError.md) | A list of project errors | [optional] [readonly] 
**ProjectDigest** | Pointer to **string** | A hash value computed from the project&#39;s current state | [optional] [readonly] 
**ModelsNotValidated** | Pointer to [**[]ModelsNotValidated**](ModelsNotValidated.md) | A list of models which were not fully validated | [optional] [readonly] 
**ComputationTime** | Pointer to **float32** | Duration of project validation in seconds | [optional] [readonly] 
**Stale** | **bool** | If true, the cached project validation results are no longer accurate because the project has changed since the cached results were calculated | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


