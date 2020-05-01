# BackupConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Type** | Pointer to **string** | Type of backup: looker-s3 or custom-s3 | [optional] 
**CustomS3Bucket** | Pointer to **string** | Name of bucket for custom-s3 backups | [optional] 
**CustomS3BucketRegion** | Pointer to **string** | Name of region where the bucket is located | [optional] 
**CustomS3Key** | Pointer to **string** | (Write-Only) AWS S3 key used for custom-s3 backups | [optional] 
**CustomS3Secret** | Pointer to **string** | (Write-Only) AWS S3 secret used for custom-s3 backups | [optional] 
**Url** | Pointer to **string** | Link to get this item | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


