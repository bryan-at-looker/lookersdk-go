# Go API client for looker

### Authorization  

The Looker API uses Looker **API3** credentials for authorization and access control. Looker admins can create API3 credentials on Looker's **Admin/Users** page. Pass API3 credentials to the **_/login** endpoint to obtain a temporary access_token. Include that access_token in the Authorization header of Looker API requests. For details, see [Looker API Authorization](https://looker.com/docs/r/api/authorization)  

For Go, we have taken care of the login by including a LoginContext function.

### Client SDKs  

The Looker API is a RESTful system that should be usable by any programming language capable of making HTTPS requests. Client SDKs for a variety of programming languages can be generated from the Looker API's Swagger JSON metadata to streamline use of the Looker API in your applications. A client SDK for Ruby is available as an example. For more information, see [Looker API Client SDKs](https://looker.com/docs/r/api/client_sdks)  

### Try It Out!  

The 'api-docs' page served by the Looker instance includes 'Try It Out!' buttons for each API method. After logging in with API3 credentials, you can use the \"Try It Out!\" buttons to call the API directly from the documentation page to interactively explore API features and responses.  Note! With great power comes great responsibility: The \"Try It Out!\" button makes API calls to your live Looker instance. Be especially careful with destructive API operations such as `delete_user` or similar. There is no \"undo\" for API operations.  ### Versioning  Future releases of Looker will expand this API release-by-release to securely expose more and more of the core power of Looker to API client applications. API endpoints marked as \"beta\" may receive breaking changes without warning (but we will try to avoid doing that). Stable (non-beta) API endpoints should not receive breaking changes in future releases. For more information, see [Looker API Versioning](https://looker.com/docs/r/api/versioning)  ### In This Release  The following are a few examples of noteworthy items that have changed between API 3.0 and API 3.1. For more comprehensive coverage of API changes, please see the release notes for your Looker release.   

## Installation

Install the following dependencies:

```shell
go get github.com/bryan-at-looker/lookersdkgo
go get github.com/joho/godotenv
```

Put the package under your project folder and add the following in import:

```shell
import "looker"
```

## Documentation for API Endpoints

All URIs are relative to *https://customer.looker.com:443/api/3.1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApiAuthApi* | [**Login**](docs/ApiAuthApi.md#login) | **Post** /login | Login
*ApiAuthApi* | [**LoginUser**](docs/ApiAuthApi.md#loginuser) | **Post** /login/{user_id} | Login user
*ApiAuthApi* | [**Logout**](docs/ApiAuthApi.md#logout) | **Delete** /logout | Logout
*AuthApi* | [**AllUserLoginLockouts**](docs/AuthApi.md#alluserloginlockouts) | **Get** /user_login_lockouts | Get All User Login Lockouts
*AuthApi* | [**CreateOidcTestConfig**](docs/AuthApi.md#createoidctestconfig) | **Post** /oidc_test_configs | Create OIDC Test Configuration
*AuthApi* | [**CreateSamlTestConfig**](docs/AuthApi.md#createsamltestconfig) | **Post** /saml_test_configs | Create SAML Test Configuration
*AuthApi* | [**CreateSsoEmbedUrl**](docs/AuthApi.md#createssoembedurl) | **Post** /embed/sso_url | Create SSO Embed Url
*AuthApi* | [**DeleteOidcTestConfig**](docs/AuthApi.md#deleteoidctestconfig) | **Delete** /oidc_test_configs/{test_slug} | Delete OIDC Test Configuration
*AuthApi* | [**DeleteSamlTestConfig**](docs/AuthApi.md#deletesamltestconfig) | **Delete** /saml_test_configs/{test_slug} | Delete SAML Test Configuration
*AuthApi* | [**DeleteUserLoginLockout**](docs/AuthApi.md#deleteuserloginlockout) | **Delete** /user_login_lockout/{key} | Delete User Login Lockout
*AuthApi* | [**FetchAndParseSamlIdpMetadata**](docs/AuthApi.md#fetchandparsesamlidpmetadata) | **Post** /fetch_and_parse_saml_idp_metadata | Parse SAML IdP Url
*AuthApi* | [**ForcePasswordResetAtNextLoginForAllUsers**](docs/AuthApi.md#forcepasswordresetatnextloginforallusers) | **Put** /password_config/force_password_reset_at_next_login_for_all_users | Force password reset
*AuthApi* | [**LdapConfig**](docs/AuthApi.md#ldapconfig) | **Get** /ldap_config | Get LDAP Configuration
*AuthApi* | [**OidcConfig**](docs/AuthApi.md#oidcconfig) | **Get** /oidc_config | Get OIDC Configuration
*AuthApi* | [**OidcTestConfig**](docs/AuthApi.md#oidctestconfig) | **Get** /oidc_test_configs/{test_slug} | Get OIDC Test Configuration
*AuthApi* | [**ParseSamlIdpMetadata**](docs/AuthApi.md#parsesamlidpmetadata) | **Post** /parse_saml_idp_metadata | Parse SAML IdP XML
*AuthApi* | [**PasswordConfig**](docs/AuthApi.md#passwordconfig) | **Get** /password_config | Get Password Config
*AuthApi* | [**SamlConfig**](docs/AuthApi.md#samlconfig) | **Get** /saml_config | Get SAML Configuration
*AuthApi* | [**SamlTestConfig**](docs/AuthApi.md#samltestconfig) | **Get** /saml_test_configs/{test_slug} | Get SAML Test Configuration
*AuthApi* | [**SearchUserLoginLockouts**](docs/AuthApi.md#searchuserloginlockouts) | **Get** /user_login_lockouts/search | Search User Login Lockouts
*AuthApi* | [**SessionConfig**](docs/AuthApi.md#sessionconfig) | **Get** /session_config | Get Session Config
*AuthApi* | [**TestLdapConfigAuth**](docs/AuthApi.md#testldapconfigauth) | **Put** /ldap_config/test_auth | Test LDAP Auth
*AuthApi* | [**TestLdapConfigConnection**](docs/AuthApi.md#testldapconfigconnection) | **Put** /ldap_config/test_connection | Test LDAP Connection
*AuthApi* | [**TestLdapConfigUserAuth**](docs/AuthApi.md#testldapconfiguserauth) | **Put** /ldap_config/test_user_auth | Test LDAP User Auth
*AuthApi* | [**TestLdapConfigUserInfo**](docs/AuthApi.md#testldapconfiguserinfo) | **Put** /ldap_config/test_user_info | Test LDAP User Info
*AuthApi* | [**UpdateLdapConfig**](docs/AuthApi.md#updateldapconfig) | **Patch** /ldap_config | Update LDAP Configuration
*AuthApi* | [**UpdateOidcConfig**](docs/AuthApi.md#updateoidcconfig) | **Patch** /oidc_config | Update OIDC Configuration
*AuthApi* | [**UpdatePasswordConfig**](docs/AuthApi.md#updatepasswordconfig) | **Patch** /password_config | Update Password Config
*AuthApi* | [**UpdateSamlConfig**](docs/AuthApi.md#updatesamlconfig) | **Patch** /saml_config | Update SAML Configuration
*AuthApi* | [**UpdateSessionConfig**](docs/AuthApi.md#updatesessionconfig) | **Patch** /session_config | Update Session Config
*ColorCollectionApi* | [**AllColorCollections**](docs/ColorCollectionApi.md#allcolorcollections) | **Get** /color_collections | Get all Color Collections
*ColorCollectionApi* | [**ColorCollection**](docs/ColorCollectionApi.md#colorcollection) | **Get** /color_collections/{collection_id} | Get Color Collection by ID
*ColorCollectionApi* | [**ColorCollectionsCustom**](docs/ColorCollectionApi.md#colorcollectionscustom) | **Get** /color_collections/custom | Get all Custom Color Collections
*ColorCollectionApi* | [**ColorCollectionsStandard**](docs/ColorCollectionApi.md#colorcollectionsstandard) | **Get** /color_collections/standard | Get all Standard Color Collections
*ColorCollectionApi* | [**CreateColorCollection**](docs/ColorCollectionApi.md#createcolorcollection) | **Post** /color_collections | Create ColorCollection
*ColorCollectionApi* | [**DefaultColorCollection**](docs/ColorCollectionApi.md#defaultcolorcollection) | **Get** /color_collections/default | Get Default Color Collection
*ColorCollectionApi* | [**DeleteColorCollection**](docs/ColorCollectionApi.md#deletecolorcollection) | **Delete** /color_collections/{collection_id} | Delete ColorCollection
*ColorCollectionApi* | [**SetDefaultColorCollection**](docs/ColorCollectionApi.md#setdefaultcolorcollection) | **Put** /color_collections/default | Set Default Color Collection
*ColorCollectionApi* | [**UpdateColorCollection**](docs/ColorCollectionApi.md#updatecolorcollection) | **Patch** /color_collections/{collection_id} | Update Custom Color collection
*ConfigApi* | [**AllLegacyFeatures**](docs/ConfigApi.md#alllegacyfeatures) | **Get** /legacy_features | Get All Legacy Features
*ConfigApi* | [**AllLocales**](docs/ConfigApi.md#alllocales) | **Get** /locales | Get All Locales
*ConfigApi* | [**AllTimezones**](docs/ConfigApi.md#alltimezones) | **Get** /timezones | Get All Timezones
*ConfigApi* | [**BackupConfiguration**](docs/ConfigApi.md#backupconfiguration) | **Get** /backup_configuration | Get Backup Configuration
*ConfigApi* | [**CreateDigestEmailSend**](docs/ConfigApi.md#createdigestemailsend) | **Post** /digest_email_send | Deliver digest email contents
*ConfigApi* | [**CustomWelcomeEmail**](docs/ConfigApi.md#customwelcomeemail) | **Get** /custom_welcome_email | Get Custom Welcome Email
*ConfigApi* | [**DigestEmailsEnabled**](docs/ConfigApi.md#digestemailsenabled) | **Get** /digest_emails_enabled | Get Digest_emails
*ConfigApi* | [**InternalHelpResources**](docs/ConfigApi.md#internalhelpresources) | **Get** /internal_help_resources_enabled | Get Internal Help Resources
*ConfigApi* | [**InternalHelpResourcesContent**](docs/ConfigApi.md#internalhelpresourcescontent) | **Get** /internal_help_resources_content | Get Internal Help Resources Content
*ConfigApi* | [**LegacyFeature**](docs/ConfigApi.md#legacyfeature) | **Get** /legacy_features/{legacy_feature_id} | Get Legacy Feature
*ConfigApi* | [**UpdateBackupConfiguration**](docs/ConfigApi.md#updatebackupconfiguration) | **Patch** /backup_configuration | Update Backup Configuration
*ConfigApi* | [**UpdateCustomWelcomeEmail**](docs/ConfigApi.md#updatecustomwelcomeemail) | **Patch** /custom_welcome_email | Update Custom Welcome Email Content
*ConfigApi* | [**UpdateCustomWelcomeEmailTest**](docs/ConfigApi.md#updatecustomwelcomeemailtest) | **Put** /custom_welcome_email_test | Send a test welcome email to the currently logged in user with the supplied content 
*ConfigApi* | [**UpdateDigestEmailsEnabled**](docs/ConfigApi.md#updatedigestemailsenabled) | **Patch** /digest_emails_enabled | Update Digest_emails
*ConfigApi* | [**UpdateInternalHelpResources**](docs/ConfigApi.md#updateinternalhelpresources) | **Patch** /internal_help_resources | Update internal help resources configuration
*ConfigApi* | [**UpdateInternalHelpResourcesContent**](docs/ConfigApi.md#updateinternalhelpresourcescontent) | **Patch** /internal_help_resources_content | Update internal help resources content
*ConfigApi* | [**UpdateLegacyFeature**](docs/ConfigApi.md#updatelegacyfeature) | **Patch** /legacy_features/{legacy_feature_id} | Update Legacy Feature
*ConfigApi* | [**UpdateWhitelabelConfiguration**](docs/ConfigApi.md#updatewhitelabelconfiguration) | **Put** /whitelabel_configuration | Update Whitelabel configuration
*ConfigApi* | [**Versions**](docs/ConfigApi.md#versions) | **Get** /versions | Get ApiVersion
*ConfigApi* | [**WhitelabelConfiguration**](docs/ConfigApi.md#whitelabelconfiguration) | **Get** /whitelabel_configuration | Get Whitelabel configuration
*ConnectionApi* | [**AllConnections**](docs/ConnectionApi.md#allconnections) | **Get** /connections | Get All Connections
*ConnectionApi* | [**AllDialectInfos**](docs/ConnectionApi.md#alldialectinfos) | **Get** /dialect_info | Get All Dialect Infos
*ConnectionApi* | [**Connection**](docs/ConnectionApi.md#connection) | **Get** /connections/{connection_name} | Get Connection
*ConnectionApi* | [**CreateConnection**](docs/ConnectionApi.md#createconnection) | **Post** /connections | Create Connection
*ConnectionApi* | [**DeleteConnection**](docs/ConnectionApi.md#deleteconnection) | **Delete** /connections/{connection_name} | Delete Connection
*ConnectionApi* | [**DeleteConnectionOverride**](docs/ConnectionApi.md#deleteconnectionoverride) | **Delete** /connections/{connection_name}/connection_override/{override_context} | Delete Connection Override
*ConnectionApi* | [**TestConnection**](docs/ConnectionApi.md#testconnection) | **Put** /connections/{connection_name}/test | Test Connection
*ConnectionApi* | [**TestConnectionConfig**](docs/ConnectionApi.md#testconnectionconfig) | **Put** /connections/test | Test Connection Configuration
*ConnectionApi* | [**UpdateConnection**](docs/ConnectionApi.md#updateconnection) | **Patch** /connections/{connection_name} | Update Connection
*ContentApi* | [**AllContentMetadataAccesses**](docs/ContentApi.md#allcontentmetadataaccesses) | **Get** /content_metadata_access | Get All Content Metadata Accesses
*ContentApi* | [**AllContentMetadatas**](docs/ContentApi.md#allcontentmetadatas) | **Get** /content_metadata | Get All Content Metadatas
*ContentApi* | [**ContentFavorite**](docs/ContentApi.md#contentfavorite) | **Get** /content_favorite/{content_favorite_id} | Get Favorite Content
*ContentApi* | [**ContentMetadata**](docs/ContentApi.md#contentmetadata) | **Get** /content_metadata/{content_metadata_id} | Get Content Metadata
*ContentApi* | [**ContentValidation**](docs/ContentApi.md#contentvalidation) | **Get** /content_validation | Validate Content
*ContentApi* | [**CreateContentFavorite**](docs/ContentApi.md#createcontentfavorite) | **Post** /content_favorite | Create Favorite Content
*ContentApi* | [**CreateContentMetadataAccess**](docs/ContentApi.md#createcontentmetadataaccess) | **Post** /content_metadata_access | Create Content Metadata Access
*ContentApi* | [**DeleteContentFavorite**](docs/ContentApi.md#deletecontentfavorite) | **Delete** /content_favorite/{content_favorite_id} | Delete Favorite Content
*ContentApi* | [**DeleteContentMetadataAccess**](docs/ContentApi.md#deletecontentmetadataaccess) | **Delete** /content_metadata_access/{content_metadata_access_id} | Delete Content Metadata Access
*ContentApi* | [**SearchContentFavorites**](docs/ContentApi.md#searchcontentfavorites) | **Get** /content_favorite/search | Search Favorite Contents
*ContentApi* | [**SearchContentViews**](docs/ContentApi.md#searchcontentviews) | **Get** /content_view/search | Search Content Views
*ContentApi* | [**UpdateContentMetadata**](docs/ContentApi.md#updatecontentmetadata) | **Patch** /content_metadata/{content_metadata_id} | Update Content Metadata
*ContentApi* | [**UpdateContentMetadataAccess**](docs/ContentApi.md#updatecontentmetadataaccess) | **Put** /content_metadata_access/{content_metadata_access_id} | Update Content Metadata Access
*ContentApi* | [**VectorThumbnail**](docs/ContentApi.md#vectorthumbnail) | **Get** /vector_thumbnail/{type}/{resource_id} | Get Vector Thumbnail
*DashboardApi* | [**AllDashboards**](docs/DashboardApi.md#alldashboards) | **Get** /dashboards | Get All Dashboards
*DashboardApi* | [**CreateDashboard**](docs/DashboardApi.md#createdashboard) | **Post** /dashboards | Create Dashboard
*DashboardApi* | [**CreateDashboardElement**](docs/DashboardApi.md#createdashboardelement) | **Post** /dashboard_elements | Create DashboardElement
*DashboardApi* | [**CreateDashboardFilter**](docs/DashboardApi.md#createdashboardfilter) | **Post** /dashboard_filters | Create Dashboard Filter
*DashboardApi* | [**CreateDashboardLayout**](docs/DashboardApi.md#createdashboardlayout) | **Post** /dashboard_layouts | Create DashboardLayout
*DashboardApi* | [**Dashboard**](docs/DashboardApi.md#dashboard) | **Get** /dashboards/{dashboard_id} | Get Dashboard
*DashboardApi* | [**DashboardDashboardElements**](docs/DashboardApi.md#dashboarddashboardelements) | **Get** /dashboards/{dashboard_id}/dashboard_elements | Get All DashboardElements
*DashboardApi* | [**DashboardDashboardFilters**](docs/DashboardApi.md#dashboarddashboardfilters) | **Get** /dashboards/{dashboard_id}/dashboard_filters | Get All Dashboard Filters
*DashboardApi* | [**DashboardDashboardLayouts**](docs/DashboardApi.md#dashboarddashboardlayouts) | **Get** /dashboards/{dashboard_id}/dashboard_layouts | Get All DashboardLayouts
*DashboardApi* | [**DashboardElement**](docs/DashboardApi.md#dashboardelement) | **Get** /dashboard_elements/{dashboard_element_id} | Get DashboardElement
*DashboardApi* | [**DashboardFilter**](docs/DashboardApi.md#dashboardfilter) | **Get** /dashboard_filters/{dashboard_filter_id} | Get Dashboard Filter
*DashboardApi* | [**DashboardLayout**](docs/DashboardApi.md#dashboardlayout) | **Get** /dashboard_layouts/{dashboard_layout_id} | Get DashboardLayout
*DashboardApi* | [**DashboardLayoutComponent**](docs/DashboardApi.md#dashboardlayoutcomponent) | **Get** /dashboard_layout_components/{dashboard_layout_component_id} | Get DashboardLayoutComponent
*DashboardApi* | [**DashboardLayoutDashboardLayoutComponents**](docs/DashboardApi.md#dashboardlayoutdashboardlayoutcomponents) | **Get** /dashboard_layouts/{dashboard_layout_id}/dashboard_layout_components | Get All DashboardLayoutComponents
*DashboardApi* | [**DashboardLookml**](docs/DashboardApi.md#dashboardlookml) | **Get** /dashboards/lookml/{dashboard_id} | Get lookml of a UDD
*DashboardApi* | [**DeleteDashboard**](docs/DashboardApi.md#deletedashboard) | **Delete** /dashboards/{dashboard_id} | Delete Dashboard
*DashboardApi* | [**DeleteDashboardElement**](docs/DashboardApi.md#deletedashboardelement) | **Delete** /dashboard_elements/{dashboard_element_id} | Delete DashboardElement
*DashboardApi* | [**DeleteDashboardFilter**](docs/DashboardApi.md#deletedashboardfilter) | **Delete** /dashboard_filters/{dashboard_filter_id} | Delete Dashboard Filter
*DashboardApi* | [**DeleteDashboardLayout**](docs/DashboardApi.md#deletedashboardlayout) | **Delete** /dashboard_layouts/{dashboard_layout_id} | Delete DashboardLayout
*DashboardApi* | [**ImportLookmlDashboard**](docs/DashboardApi.md#importlookmldashboard) | **Post** /dashboards/{lookml_dashboard_id}/import/{space_id} | Import LookML Dashboard
*DashboardApi* | [**SearchDashboardElements**](docs/DashboardApi.md#searchdashboardelements) | **Get** /dashboard_elements/search | Search Dashboard Elements
*DashboardApi* | [**SearchDashboards**](docs/DashboardApi.md#searchdashboards) | **Get** /dashboards/search | Search Dashboards
*DashboardApi* | [**SyncLookmlDashboard**](docs/DashboardApi.md#synclookmldashboard) | **Patch** /dashboards/{lookml_dashboard_id}/sync | Sync LookML Dashboard
*DashboardApi* | [**UpdateDashboard**](docs/DashboardApi.md#updatedashboard) | **Patch** /dashboards/{dashboard_id} | Update Dashboard
*DashboardApi* | [**UpdateDashboardElement**](docs/DashboardApi.md#updatedashboardelement) | **Patch** /dashboard_elements/{dashboard_element_id} | Update DashboardElement
*DashboardApi* | [**UpdateDashboardFilter**](docs/DashboardApi.md#updatedashboardfilter) | **Patch** /dashboard_filters/{dashboard_filter_id} | Update Dashboard Filter
*DashboardApi* | [**UpdateDashboardLayout**](docs/DashboardApi.md#updatedashboardlayout) | **Patch** /dashboard_layouts/{dashboard_layout_id} | Update DashboardLayout
*DashboardApi* | [**UpdateDashboardLayoutComponent**](docs/DashboardApi.md#updatedashboardlayoutcomponent) | **Patch** /dashboard_layout_components/{dashboard_layout_component_id} | Update DashboardLayoutComponent
*DataActionApi* | [**FetchRemoteDataActionForm**](docs/DataActionApi.md#fetchremotedataactionform) | **Post** /data_actions/form | Fetch Remote Data Action Form
*DataActionApi* | [**PerformDataAction**](docs/DataActionApi.md#performdataaction) | **Post** /data_actions | Send a Data Action
*DatagroupApi* | [**AllDatagroups**](docs/DatagroupApi.md#alldatagroups) | **Get** /datagroups | Get All Datagroups
*DatagroupApi* | [**Datagroup**](docs/DatagroupApi.md#datagroup) | **Get** /datagroups/{datagroup_id} | Get Datagroup
*DatagroupApi* | [**UpdateDatagroup**](docs/DatagroupApi.md#updatedatagroup) | **Patch** /datagroups/{datagroup_id} | Update Datagroup
*FolderApi* | [**AllFolders**](docs/FolderApi.md#allfolders) | **Get** /folders | Get All Folders
*FolderApi* | [**CreateFolder**](docs/FolderApi.md#createfolder) | **Post** /folders | Create Folder
*FolderApi* | [**DeleteFolder**](docs/FolderApi.md#deletefolder) | **Delete** /folders/{folder_id} | Delete Folder
*FolderApi* | [**Folder**](docs/FolderApi.md#folder) | **Get** /folders/{folder_id} | Get Folder
*FolderApi* | [**FolderAncestors**](docs/FolderApi.md#folderancestors) | **Get** /folders/{folder_id}/ancestors | Get Folder Ancestors
*FolderApi* | [**FolderChildren**](docs/FolderApi.md#folderchildren) | **Get** /folders/{folder_id}/children | Get Folder Children
*FolderApi* | [**FolderChildrenSearch**](docs/FolderApi.md#folderchildrensearch) | **Get** /folders/{folder_id}/children/search | Search Folder Children
*FolderApi* | [**FolderDashboards**](docs/FolderApi.md#folderdashboards) | **Get** /folders/{folder_id}/dashboards | Get Folder Dashboards
*FolderApi* | [**FolderLooks**](docs/FolderApi.md#folderlooks) | **Get** /folders/{folder_id}/looks | Get Folder Looks
*FolderApi* | [**FolderParent**](docs/FolderApi.md#folderparent) | **Get** /folders/{folder_id}/parent | Get Folder Parent
*FolderApi* | [**SearchFolders**](docs/FolderApi.md#searchfolders) | **Get** /folders/search | Search Folders
*FolderApi* | [**UpdateFolder**](docs/FolderApi.md#updatefolder) | **Patch** /folders/{folder_id} | Update Folder
*GroupApi* | [**AddGroupGroup**](docs/GroupApi.md#addgroupgroup) | **Post** /groups/{group_id}/groups | Add a Group to Group
*GroupApi* | [**AddGroupUser**](docs/GroupApi.md#addgroupuser) | **Post** /groups/{group_id}/users | Add a User to Group
*GroupApi* | [**AllGroupGroups**](docs/GroupApi.md#allgroupgroups) | **Get** /groups/{group_id}/groups | Get All Groups in Group
*GroupApi* | [**AllGroupUsers**](docs/GroupApi.md#allgroupusers) | **Get** /groups/{group_id}/users | Get All Users in Group
*GroupApi* | [**AllGroups**](docs/GroupApi.md#allgroups) | **Get** /groups | Get All Groups
*GroupApi* | [**CreateGroup**](docs/GroupApi.md#creategroup) | **Post** /groups | Create Group
*GroupApi* | [**DeleteGroup**](docs/GroupApi.md#deletegroup) | **Delete** /groups/{group_id} | Delete Group
*GroupApi* | [**DeleteGroupFromGroup**](docs/GroupApi.md#deletegroupfromgroup) | **Delete** /groups/{group_id}/groups/{deleting_group_id} | Deletes a Group from Group
*GroupApi* | [**DeleteGroupUser**](docs/GroupApi.md#deletegroupuser) | **Delete** /groups/{group_id}/users/{user_id} | Remove a User from Group
*GroupApi* | [**DeleteUserAttributeGroupValue**](docs/GroupApi.md#deleteuserattributegroupvalue) | **Delete** /groups/{group_id}/attribute_values/{user_attribute_id} | Delete User Attribute Group Value
*GroupApi* | [**Group**](docs/GroupApi.md#group) | **Get** /groups/{group_id} | Get Group
*GroupApi* | [**SearchGroups**](docs/GroupApi.md#searchgroups) | **Get** /groups/search | Search Groups
*GroupApi* | [**UpdateGroup**](docs/GroupApi.md#updategroup) | **Patch** /groups/{group_id} | Update Group
*GroupApi* | [**UpdateUserAttributeGroupValue**](docs/GroupApi.md#updateuserattributegroupvalue) | **Patch** /groups/{group_id}/attribute_values/{user_attribute_id} | Set User Attribute Group Value
*HomepageApi* | [**AllHomepageItems**](docs/HomepageApi.md#allhomepageitems) | **Get** /homepage_items | Get All Homepage Items
*HomepageApi* | [**AllHomepageSections**](docs/HomepageApi.md#allhomepagesections) | **Get** /homepage_sections | Get All Homepage sections
*HomepageApi* | [**AllHomepages**](docs/HomepageApi.md#allhomepages) | **Get** /homepages | Get All Homepages
*HomepageApi* | [**CreateHomepage**](docs/HomepageApi.md#createhomepage) | **Post** /homepages | Create Homepage
*HomepageApi* | [**CreateHomepageItem**](docs/HomepageApi.md#createhomepageitem) | **Post** /homepage_items | Create Homepage Item
*HomepageApi* | [**CreateHomepageSection**](docs/HomepageApi.md#createhomepagesection) | **Post** /homepage_sections | Create Homepage section
*HomepageApi* | [**DeleteHomepage**](docs/HomepageApi.md#deletehomepage) | **Delete** /homepages/{homepage_id} | Delete Homepage
*HomepageApi* | [**DeleteHomepageItem**](docs/HomepageApi.md#deletehomepageitem) | **Delete** /homepage_items/{homepage_item_id} | Delete Homepage Item
*HomepageApi* | [**DeleteHomepageSection**](docs/HomepageApi.md#deletehomepagesection) | **Delete** /homepage_sections/{homepage_section_id} | Delete Homepage section
*HomepageApi* | [**Homepage**](docs/HomepageApi.md#homepage) | **Get** /homepages/{homepage_id} | Get Homepage
*HomepageApi* | [**HomepageItem**](docs/HomepageApi.md#homepageitem) | **Get** /homepage_items/{homepage_item_id} | Get Homepage Item
*HomepageApi* | [**HomepageSection**](docs/HomepageApi.md#homepagesection) | **Get** /homepage_sections/{homepage_section_id} | Get Homepage section
*HomepageApi* | [**SearchHomepages**](docs/HomepageApi.md#searchhomepages) | **Get** /homepages/search | Search Homepages
*HomepageApi* | [**UpdateHomepage**](docs/HomepageApi.md#updatehomepage) | **Patch** /homepages/{homepage_id} | Update Homepage
*HomepageApi* | [**UpdateHomepageItem**](docs/HomepageApi.md#updatehomepageitem) | **Patch** /homepage_items/{homepage_item_id} | Update Homepage Item
*HomepageApi* | [**UpdateHomepageSection**](docs/HomepageApi.md#updatehomepagesection) | **Patch** /homepage_sections/{homepage_section_id} | Update Homepage section
*IntegrationApi* | [**AcceptIntegrationHubLegalAgreement**](docs/IntegrationApi.md#acceptintegrationhublegalagreement) | **Post** /integration_hubs/{integration_hub_id}/accept_legal_agreement | Accept Integration Hub Legal Agreement
*IntegrationApi* | [**AllIntegrationHubs**](docs/IntegrationApi.md#allintegrationhubs) | **Get** /integration_hubs | Get All Integration Hubs
*IntegrationApi* | [**AllIntegrations**](docs/IntegrationApi.md#allintegrations) | **Get** /integrations | Get All Integrations
*IntegrationApi* | [**CreateIntegrationHub**](docs/IntegrationApi.md#createintegrationhub) | **Post** /integration_hubs | Create Integration Hub
*IntegrationApi* | [**DeleteIntegrationHub**](docs/IntegrationApi.md#deleteintegrationhub) | **Delete** /integration_hubs/{integration_hub_id} | Delete Integration Hub
*IntegrationApi* | [**FetchIntegrationForm**](docs/IntegrationApi.md#fetchintegrationform) | **Post** /integrations/{integration_id}/form | Fetch Remote Integration Form
*IntegrationApi* | [**Integration**](docs/IntegrationApi.md#integration) | **Get** /integrations/{integration_id} | Get Integration
*IntegrationApi* | [**IntegrationHub**](docs/IntegrationApi.md#integrationhub) | **Get** /integration_hubs/{integration_hub_id} | Get Integration Hub
*IntegrationApi* | [**TestIntegration**](docs/IntegrationApi.md#testintegration) | **Post** /integrations/{integration_id}/test | Test integration
*IntegrationApi* | [**UpdateIntegration**](docs/IntegrationApi.md#updateintegration) | **Patch** /integrations/{integration_id} | Update Integration
*IntegrationApi* | [**UpdateIntegrationHub**](docs/IntegrationApi.md#updateintegrationhub) | **Patch** /integration_hubs/{integration_hub_id} | Update Integration Hub
*LookApi* | [**AllLooks**](docs/LookApi.md#alllooks) | **Get** /looks | Get All Looks
*LookApi* | [**CreateLook**](docs/LookApi.md#createlook) | **Post** /looks | Create Look
*LookApi* | [**DeleteLook**](docs/LookApi.md#deletelook) | **Delete** /looks/{look_id} | Delete Look
*LookApi* | [**Look**](docs/LookApi.md#look) | **Get** /looks/{look_id} | Get Look
*LookApi* | [**RunLook**](docs/LookApi.md#runlook) | **Get** /looks/{look_id}/run/{result_format} | Run Look
*LookApi* | [**SearchLooks**](docs/LookApi.md#searchlooks) | **Get** /looks/search | Search Looks
*LookApi* | [**UpdateLook**](docs/LookApi.md#updatelook) | **Patch** /looks/{look_id} | Update Look
*LookmlModelApi* | [**AllLookmlModels**](docs/LookmlModelApi.md#alllookmlmodels) | **Get** /lookml_models | Get All LookML Models
*LookmlModelApi* | [**CreateLookmlModel**](docs/LookmlModelApi.md#createlookmlmodel) | **Post** /lookml_models | Create LookML Model
*LookmlModelApi* | [**DeleteLookmlModel**](docs/LookmlModelApi.md#deletelookmlmodel) | **Delete** /lookml_models/{lookml_model_name} | Delete LookML Model
*LookmlModelApi* | [**LookmlModel**](docs/LookmlModelApi.md#lookmlmodel) | **Get** /lookml_models/{lookml_model_name} | Get LookML Model
*LookmlModelApi* | [**LookmlModelExplore**](docs/LookmlModelApi.md#lookmlmodelexplore) | **Get** /lookml_models/{lookml_model_name}/explores/{explore_name} | Get LookML Model Explore
*LookmlModelApi* | [**UpdateLookmlModel**](docs/LookmlModelApi.md#updatelookmlmodel) | **Patch** /lookml_models/{lookml_model_name} | Update LookML Model
*ProjectApi* | [**AllGitBranches**](docs/ProjectApi.md#allgitbranches) | **Get** /projects/{project_id}/git_branches | Get All Git Branches
*ProjectApi* | [**AllGitConnectionTests**](docs/ProjectApi.md#allgitconnectiontests) | **Get** /projects/{project_id}/git_connection_tests | Get All Git Connection Tests
*ProjectApi* | [**AllLookmlTests**](docs/ProjectApi.md#alllookmltests) | **Get** /projects/{project_id}/lookml_tests | Get All LookML Tests
*ProjectApi* | [**AllProjectFiles**](docs/ProjectApi.md#allprojectfiles) | **Get** /projects/{project_id}/files | Get All Project Files
*ProjectApi* | [**AllProjects**](docs/ProjectApi.md#allprojects) | **Get** /projects | Get All Projects
*ProjectApi* | [**CreateGitBranch**](docs/ProjectApi.md#creategitbranch) | **Post** /projects/{project_id}/git_branch | Checkout New Git Branch
*ProjectApi* | [**CreateGitDeployKey**](docs/ProjectApi.md#creategitdeploykey) | **Post** /projects/{project_id}/git/deploy_key | Create Deploy Key
*ProjectApi* | [**CreateProject**](docs/ProjectApi.md#createproject) | **Post** /projects | Create Project
*ProjectApi* | [**DeleteGitBranch**](docs/ProjectApi.md#deletegitbranch) | **Delete** /projects/{project_id}/git_branch/{branch_name} | Delete a Git Branch
*ProjectApi* | [**DeleteRepositoryCredential**](docs/ProjectApi.md#deleterepositorycredential) | **Delete** /projects/{root_project_id}/credential/{credential_id} | Delete Repository Credential
*ProjectApi* | [**DeployToProduction**](docs/ProjectApi.md#deploytoproduction) | **Post** /projects/{project_id}/deploy_to_production | Deploy To Production
*ProjectApi* | [**FindGitBranch**](docs/ProjectApi.md#findgitbranch) | **Get** /projects/{project_id}/git_branch/{branch_name} | Find a Git Branch
*ProjectApi* | [**GetAllRepositoryCredentials**](docs/ProjectApi.md#getallrepositorycredentials) | **Get** /projects/{root_project_id}/credentials | Get All Repository Credentials
*ProjectApi* | [**GitBranch**](docs/ProjectApi.md#gitbranch) | **Get** /projects/{project_id}/git_branch | Get Active Git Branch
*ProjectApi* | [**GitDeployKey**](docs/ProjectApi.md#gitdeploykey) | **Get** /projects/{project_id}/git/deploy_key | Git Deploy Key
*ProjectApi* | [**Manifest**](docs/ProjectApi.md#manifest) | **Get** /projects/{project_id}/manifest | Get Manifest
*ProjectApi* | [**Project**](docs/ProjectApi.md#project) | **Get** /projects/{project_id} | Get Project
*ProjectApi* | [**ProjectFile**](docs/ProjectApi.md#projectfile) | **Get** /projects/{project_id}/files/file | Get Project File
*ProjectApi* | [**ProjectValidationResults**](docs/ProjectApi.md#projectvalidationresults) | **Get** /projects/{project_id}/validate | Cached Project Validation Results
*ProjectApi* | [**ProjectWorkspace**](docs/ProjectApi.md#projectworkspace) | **Get** /projects/{project_id}/current_workspace | Get Project Workspace
*ProjectApi* | [**ResetProjectToProduction**](docs/ProjectApi.md#resetprojecttoproduction) | **Post** /projects/{project_id}/reset_to_production | Reset To Production
*ProjectApi* | [**ResetProjectToRemote**](docs/ProjectApi.md#resetprojecttoremote) | **Post** /projects/{project_id}/reset_to_remote | Reset To Remote
*ProjectApi* | [**RunGitConnectionTest**](docs/ProjectApi.md#rungitconnectiontest) | **Get** /projects/{project_id}/git_connection_tests/{test_id} | Run Git Connection Test
*ProjectApi* | [**RunLookmlTest**](docs/ProjectApi.md#runlookmltest) | **Get** /projects/{project_id}/lookml_tests/run | Run LookML Test
*ProjectApi* | [**UpdateGitBranch**](docs/ProjectApi.md#updategitbranch) | **Put** /projects/{project_id}/git_branch | Update Project Git Branch
*ProjectApi* | [**UpdateProject**](docs/ProjectApi.md#updateproject) | **Patch** /projects/{project_id} | Update Project
*ProjectApi* | [**UpdateRepositoryCredential**](docs/ProjectApi.md#updaterepositorycredential) | **Put** /projects/{root_project_id}/credential/{credential_id} | Create Repository Credential
*ProjectApi* | [**ValidateProject**](docs/ProjectApi.md#validateproject) | **Post** /projects/{project_id}/validate | Validate Project
*QueryApi* | [**AllRunningQueries**](docs/QueryApi.md#allrunningqueries) | **Get** /running_queries | Get All Running Queries
*QueryApi* | [**CreateMergeQuery**](docs/QueryApi.md#createmergequery) | **Post** /merge_queries | Create Merge Query
*QueryApi* | [**CreateQuery**](docs/QueryApi.md#createquery) | **Post** /queries | Create Query
*QueryApi* | [**CreateQueryTask**](docs/QueryApi.md#createquerytask) | **Post** /query_tasks | Run Query Async
*QueryApi* | [**CreateSqlQuery**](docs/QueryApi.md#createsqlquery) | **Post** /sql_queries | Create SQL Runner Query
*QueryApi* | [**KillQuery**](docs/QueryApi.md#killquery) | **Delete** /running_queries/{query_task_id} | Kill Running Query
*QueryApi* | [**MergeQuery**](docs/QueryApi.md#mergequery) | **Get** /merge_queries/{merge_query_id} | Get Merge Query
*QueryApi* | [**Query**](docs/QueryApi.md#query) | **Get** /queries/{query_id} | Get Query
*QueryApi* | [**QueryForSlug**](docs/QueryApi.md#queryforslug) | **Get** /queries/slug/{slug} | Get Query for Slug
*QueryApi* | [**QueryTask**](docs/QueryApi.md#querytask) | **Get** /query_tasks/{query_task_id} | Get Async Query Info
*QueryApi* | [**QueryTaskMultiResults**](docs/QueryApi.md#querytaskmultiresults) | **Get** /query_tasks/multi_results | Get Multiple Async Query Results
*QueryApi* | [**QueryTaskResults**](docs/QueryApi.md#querytaskresults) | **Get** /query_tasks/{query_task_id}/results | Get Async Query Results
*QueryApi* | [**RunInlineQuery**](docs/QueryApi.md#runinlinequery) | **Post** /queries/run/{result_format} | Run Inline Query
*QueryApi* | [**RunQuery**](docs/QueryApi.md#runquery) | **Get** /queries/{query_id}/run/{result_format} | Run Query
*QueryApi* | [**RunSqlQuery**](docs/QueryApi.md#runsqlquery) | **Post** /sql_queries/{slug}/run/{result_format} | Run SQL Runner Query
*QueryApi* | [**RunUrlEncodedQuery**](docs/QueryApi.md#runurlencodedquery) | **Get** /queries/models/{model_name}/views/{view_name}/run/{result_format} | Run Url Encoded Query
*QueryApi* | [**SqlQuery**](docs/QueryApi.md#sqlquery) | **Get** /sql_queries/{slug} | Get SQL Runner Query
*RenderTaskApi* | [**CreateDashboardRenderTask**](docs/RenderTaskApi.md#createdashboardrendertask) | **Post** /render_tasks/dashboards/{dashboard_id}/{result_format} | Create Dashboard Render Task
*RenderTaskApi* | [**CreateLookRenderTask**](docs/RenderTaskApi.md#createlookrendertask) | **Post** /render_tasks/looks/{look_id}/{result_format} | Create Look Render Task
*RenderTaskApi* | [**CreateLookmlDashboardRenderTask**](docs/RenderTaskApi.md#createlookmldashboardrendertask) | **Post** /render_tasks/lookml_dashboards/{dashboard_id}/{result_format} | Create Lookml Dashboard Render Task
*RenderTaskApi* | [**CreateQueryRenderTask**](docs/RenderTaskApi.md#createqueryrendertask) | **Post** /render_tasks/queries/{query_id}/{result_format} | Create Query Render Task
*RenderTaskApi* | [**RenderTask**](docs/RenderTaskApi.md#rendertask) | **Get** /render_tasks/{render_task_id} | Get Render Task
*RenderTaskApi* | [**RenderTaskResults**](docs/RenderTaskApi.md#rendertaskresults) | **Get** /render_tasks/{render_task_id}/results | Render Task Results
*RoleApi* | [**AllModelSets**](docs/RoleApi.md#allmodelsets) | **Get** /model_sets | Get All Model Sets
*RoleApi* | [**AllPermissionSets**](docs/RoleApi.md#allpermissionsets) | **Get** /permission_sets | Get All Permission Sets
*RoleApi* | [**AllPermissions**](docs/RoleApi.md#allpermissions) | **Get** /permissions | Get All Permissions
*RoleApi* | [**AllRoles**](docs/RoleApi.md#allroles) | **Get** /roles | Get All Roles
*RoleApi* | [**CreateModelSet**](docs/RoleApi.md#createmodelset) | **Post** /model_sets | Create Model Set
*RoleApi* | [**CreatePermissionSet**](docs/RoleApi.md#createpermissionset) | **Post** /permission_sets | Create Permission Set
*RoleApi* | [**CreateRole**](docs/RoleApi.md#createrole) | **Post** /roles | Create Role
*RoleApi* | [**DeleteModelSet**](docs/RoleApi.md#deletemodelset) | **Delete** /model_sets/{model_set_id} | Delete Model Set
*RoleApi* | [**DeletePermissionSet**](docs/RoleApi.md#deletepermissionset) | **Delete** /permission_sets/{permission_set_id} | Delete Permission Set
*RoleApi* | [**DeleteRole**](docs/RoleApi.md#deleterole) | **Delete** /roles/{role_id} | Delete Role
*RoleApi* | [**ModelSet**](docs/RoleApi.md#modelset) | **Get** /model_sets/{model_set_id} | Get Model Set
*RoleApi* | [**PermissionSet**](docs/RoleApi.md#permissionset) | **Get** /permission_sets/{permission_set_id} | Get Permission Set
*RoleApi* | [**Role**](docs/RoleApi.md#role) | **Get** /roles/{role_id} | Get Role
*RoleApi* | [**RoleGroups**](docs/RoleApi.md#rolegroups) | **Get** /roles/{role_id}/groups | Get Role Groups
*RoleApi* | [**RoleUsers**](docs/RoleApi.md#roleusers) | **Get** /roles/{role_id}/users | Get Role Users
*RoleApi* | [**SearchModelSets**](docs/RoleApi.md#searchmodelsets) | **Get** /model_sets/search | Search Model Sets
*RoleApi* | [**SearchPermissionSets**](docs/RoleApi.md#searchpermissionsets) | **Get** /permission_sets/search | Search Permission Sets
*RoleApi* | [**SearchRoles**](docs/RoleApi.md#searchroles) | **Get** /roles/search | Search Roles
*RoleApi* | [**SetRoleGroups**](docs/RoleApi.md#setrolegroups) | **Put** /roles/{role_id}/groups | Update Role Groups
*RoleApi* | [**SetRoleUsers**](docs/RoleApi.md#setroleusers) | **Put** /roles/{role_id}/users | Update Role Users
*RoleApi* | [**UpdateModelSet**](docs/RoleApi.md#updatemodelset) | **Patch** /model_sets/{model_set_id} | Update Model Set
*RoleApi* | [**UpdatePermissionSet**](docs/RoleApi.md#updatepermissionset) | **Patch** /permission_sets/{permission_set_id} | Update Permission Set
*RoleApi* | [**UpdateRole**](docs/RoleApi.md#updaterole) | **Patch** /roles/{role_id} | Update Role
*ScheduledPlanApi* | [**AllScheduledPlans**](docs/ScheduledPlanApi.md#allscheduledplans) | **Get** /scheduled_plans | Get All Scheduled Plans
*ScheduledPlanApi* | [**CreateScheduledPlan**](docs/ScheduledPlanApi.md#createscheduledplan) | **Post** /scheduled_plans | Create Scheduled Plan
*ScheduledPlanApi* | [**DeleteScheduledPlan**](docs/ScheduledPlanApi.md#deletescheduledplan) | **Delete** /scheduled_plans/{scheduled_plan_id} | Delete Scheduled Plan
*ScheduledPlanApi* | [**ScheduledPlan**](docs/ScheduledPlanApi.md#scheduledplan) | **Get** /scheduled_plans/{scheduled_plan_id} | Get Scheduled Plan
*ScheduledPlanApi* | [**ScheduledPlanRunOnce**](docs/ScheduledPlanApi.md#scheduledplanrunonce) | **Post** /scheduled_plans/run_once | Run Scheduled Plan Once
*ScheduledPlanApi* | [**ScheduledPlanRunOnceById**](docs/ScheduledPlanApi.md#scheduledplanrunoncebyid) | **Post** /scheduled_plans/{scheduled_plan_id}/run_once | Run Scheduled Plan Once by Id
*ScheduledPlanApi* | [**ScheduledPlansForDashboard**](docs/ScheduledPlanApi.md#scheduledplansfordashboard) | **Get** /scheduled_plans/dashboard/{dashboard_id} | Scheduled Plans for Dashboard
*ScheduledPlanApi* | [**ScheduledPlansForLook**](docs/ScheduledPlanApi.md#scheduledplansforlook) | **Get** /scheduled_plans/look/{look_id} | Scheduled Plans for Look
*ScheduledPlanApi* | [**ScheduledPlansForLookmlDashboard**](docs/ScheduledPlanApi.md#scheduledplansforlookmldashboard) | **Get** /scheduled_plans/lookml_dashboard/{lookml_dashboard_id} | Scheduled Plans for LookML Dashboard
*ScheduledPlanApi* | [**ScheduledPlansForSpace**](docs/ScheduledPlanApi.md#scheduledplansforspace) | **Get** /scheduled_plans/space/{space_id} | Scheduled Plans for Space
*ScheduledPlanApi* | [**UpdateScheduledPlan**](docs/ScheduledPlanApi.md#updatescheduledplan) | **Patch** /scheduled_plans/{scheduled_plan_id} | Update Scheduled Plan
*SessionApi* | [**Session**](docs/SessionApi.md#session) | **Get** /session | Get Session
*SessionApi* | [**UpdateSession**](docs/SessionApi.md#updatesession) | **Patch** /session | Update Session
*SpaceApi* | [**AllSpaces**](docs/SpaceApi.md#allspaces) | **Get** /spaces | Get All Spaces
*SpaceApi* | [**CreateSpace**](docs/SpaceApi.md#createspace) | **Post** /spaces | Create Space
*SpaceApi* | [**DeleteSpace**](docs/SpaceApi.md#deletespace) | **Delete** /spaces/{space_id} | Delete Space
*SpaceApi* | [**SearchSpaces**](docs/SpaceApi.md#searchspaces) | **Get** /spaces/search | Search Spaces
*SpaceApi* | [**Space**](docs/SpaceApi.md#space) | **Get** /spaces/{space_id} | Get Space
*SpaceApi* | [**SpaceAncestors**](docs/SpaceApi.md#spaceancestors) | **Get** /spaces/{space_id}/ancestors | Get Space Ancestors
*SpaceApi* | [**SpaceChildren**](docs/SpaceApi.md#spacechildren) | **Get** /spaces/{space_id}/children | Get Space Children
*SpaceApi* | [**SpaceChildrenSearch**](docs/SpaceApi.md#spacechildrensearch) | **Get** /spaces/{space_id}/children/search | Search Space Children
*SpaceApi* | [**SpaceDashboards**](docs/SpaceApi.md#spacedashboards) | **Get** /spaces/{space_id}/dashboards | Get Space Dashboards
*SpaceApi* | [**SpaceLooks**](docs/SpaceApi.md#spacelooks) | **Get** /spaces/{space_id}/looks | Get Space Looks
*SpaceApi* | [**SpaceParent**](docs/SpaceApi.md#spaceparent) | **Get** /spaces/{space_id}/parent | Get Space Parent
*SpaceApi* | [**UpdateSpace**](docs/SpaceApi.md#updatespace) | **Patch** /spaces/{space_id} | Update Space
*ThemeApi* | [**ActiveThemes**](docs/ThemeApi.md#activethemes) | **Get** /themes/active | Get Active Themes
*ThemeApi* | [**AllThemes**](docs/ThemeApi.md#allthemes) | **Get** /themes | Get All Themes
*ThemeApi* | [**CreateTheme**](docs/ThemeApi.md#createtheme) | **Post** /themes | Create Theme
*ThemeApi* | [**DefaultTheme**](docs/ThemeApi.md#defaulttheme) | **Get** /themes/default | Get Default Theme
*ThemeApi* | [**DeleteTheme**](docs/ThemeApi.md#deletetheme) | **Delete** /themes/{theme_id} | Delete Theme
*ThemeApi* | [**SearchThemes**](docs/ThemeApi.md#searchthemes) | **Get** /themes/search | Search Themes
*ThemeApi* | [**SetDefaultTheme**](docs/ThemeApi.md#setdefaulttheme) | **Put** /themes/default | Set Default Theme
*ThemeApi* | [**Theme**](docs/ThemeApi.md#theme) | **Get** /themes/{theme_id} | Get Theme
*ThemeApi* | [**ThemeOrDefault**](docs/ThemeApi.md#themeordefault) | **Get** /themes/theme_or_default | Get Theme or Default
*ThemeApi* | [**UpdateTheme**](docs/ThemeApi.md#updatetheme) | **Patch** /themes/{theme_id} | Update Theme
*ThemeApi* | [**ValidateTheme**](docs/ThemeApi.md#validatetheme) | **Post** /themes/validate | Validate Theme
*UserApi* | [**AllUserCredentialsApi3s**](docs/UserApi.md#allusercredentialsapi3s) | **Get** /users/{user_id}/credentials_api3 | Get All API 3 Credentials
*UserApi* | [**AllUserCredentialsEmbeds**](docs/UserApi.md#allusercredentialsembeds) | **Get** /users/{user_id}/credentials_embed | Get All Embedding Credentials
*UserApi* | [**AllUserSessions**](docs/UserApi.md#allusersessions) | **Get** /users/{user_id}/sessions | Get All Web Login Sessions
*UserApi* | [**AllUsers**](docs/UserApi.md#allusers) | **Get** /users | Get All Users
*UserApi* | [**CreateUser**](docs/UserApi.md#createuser) | **Post** /users | Create User
*UserApi* | [**CreateUserCredentialsApi3**](docs/UserApi.md#createusercredentialsapi3) | **Post** /users/{user_id}/credentials_api3 | Create API 3 Credential
*UserApi* | [**CreateUserCredentialsEmail**](docs/UserApi.md#createusercredentialsemail) | **Post** /users/{user_id}/credentials_email | Create Email/Password Credential
*UserApi* | [**CreateUserCredentialsEmailPasswordReset**](docs/UserApi.md#createusercredentialsemailpasswordreset) | **Post** /users/{user_id}/credentials_email/password_reset | Create Password Reset Token
*UserApi* | [**CreateUserCredentialsTotp**](docs/UserApi.md#createusercredentialstotp) | **Post** /users/{user_id}/credentials_totp | Create Two-Factor Credential
*UserApi* | [**DeleteUser**](docs/UserApi.md#deleteuser) | **Delete** /users/{user_id} | Delete User
*UserApi* | [**DeleteUserAttributeUserValue**](docs/UserApi.md#deleteuserattributeuservalue) | **Delete** /users/{user_id}/attribute_values/{user_attribute_id} | Delete User Attribute User Value
*UserApi* | [**DeleteUserCredentialsApi3**](docs/UserApi.md#deleteusercredentialsapi3) | **Delete** /users/{user_id}/credentials_api3/{credentials_api3_id} | Delete API 3 Credential
*UserApi* | [**DeleteUserCredentialsEmail**](docs/UserApi.md#deleteusercredentialsemail) | **Delete** /users/{user_id}/credentials_email | Delete Email/Password Credential
*UserApi* | [**DeleteUserCredentialsEmbed**](docs/UserApi.md#deleteusercredentialsembed) | **Delete** /users/{user_id}/credentials_embed/{credentials_embed_id} | Delete Embedding Credential
*UserApi* | [**DeleteUserCredentialsGoogle**](docs/UserApi.md#deleteusercredentialsgoogle) | **Delete** /users/{user_id}/credentials_google | Delete Google Auth Credential
*UserApi* | [**DeleteUserCredentialsLdap**](docs/UserApi.md#deleteusercredentialsldap) | **Delete** /users/{user_id}/credentials_ldap | Delete LDAP Credential
*UserApi* | [**DeleteUserCredentialsLookerOpenid**](docs/UserApi.md#deleteusercredentialslookeropenid) | **Delete** /users/{user_id}/credentials_looker_openid | Delete Looker OpenId Credential
*UserApi* | [**DeleteUserCredentialsOidc**](docs/UserApi.md#deleteusercredentialsoidc) | **Delete** /users/{user_id}/credentials_oidc | Delete OIDC Auth Credential
*UserApi* | [**DeleteUserCredentialsSaml**](docs/UserApi.md#deleteusercredentialssaml) | **Delete** /users/{user_id}/credentials_saml | Delete Saml Auth Credential
*UserApi* | [**DeleteUserCredentialsTotp**](docs/UserApi.md#deleteusercredentialstotp) | **Delete** /users/{user_id}/credentials_totp | Delete Two-Factor Credential
*UserApi* | [**DeleteUserSession**](docs/UserApi.md#deleteusersession) | **Delete** /users/{user_id}/sessions/{session_id} | Delete Web Login Session
*UserApi* | [**Me**](docs/UserApi.md#me) | **Get** /user | Get Current User
*UserApi* | [**SearchUsers**](docs/UserApi.md#searchusers) | **Get** /users/search | Search Users
*UserApi* | [**SearchUsersNames**](docs/UserApi.md#searchusersnames) | **Get** /users/search/names/{pattern} | Search User Names
*UserApi* | [**SetUserAttributeUserValue**](docs/UserApi.md#setuserattributeuservalue) | **Patch** /users/{user_id}/attribute_values/{user_attribute_id} | Set User Attribute User Value
*UserApi* | [**SetUserRoles**](docs/UserApi.md#setuserroles) | **Put** /users/{user_id}/roles | Set User Roles
*UserApi* | [**UpdateUser**](docs/UserApi.md#updateuser) | **Patch** /users/{user_id} | Update User
*UserApi* | [**UpdateUserCredentialsEmail**](docs/UserApi.md#updateusercredentialsemail) | **Patch** /users/{user_id}/credentials_email | Update Email/Password Credential
*UserApi* | [**User**](docs/UserApi.md#user) | **Get** /users/{user_id} | Get User by Id
*UserApi* | [**UserAttributeUserValues**](docs/UserApi.md#userattributeuservalues) | **Get** /users/{user_id}/attribute_values | Get User Attribute Values
*UserApi* | [**UserCredentialsApi3**](docs/UserApi.md#usercredentialsapi3) | **Get** /users/{user_id}/credentials_api3/{credentials_api3_id} | Get API 3 Credential
*UserApi* | [**UserCredentialsEmail**](docs/UserApi.md#usercredentialsemail) | **Get** /users/{user_id}/credentials_email | Get Email/Password Credential
*UserApi* | [**UserCredentialsEmbed**](docs/UserApi.md#usercredentialsembed) | **Get** /users/{user_id}/credentials_embed/{credentials_embed_id} | Get Embedding Credential
*UserApi* | [**UserCredentialsGoogle**](docs/UserApi.md#usercredentialsgoogle) | **Get** /users/{user_id}/credentials_google | Get Google Auth Credential
*UserApi* | [**UserCredentialsLdap**](docs/UserApi.md#usercredentialsldap) | **Get** /users/{user_id}/credentials_ldap | Get LDAP Credential
*UserApi* | [**UserCredentialsLookerOpenid**](docs/UserApi.md#usercredentialslookeropenid) | **Get** /users/{user_id}/credentials_looker_openid | Get Looker OpenId Credential
*UserApi* | [**UserCredentialsOidc**](docs/UserApi.md#usercredentialsoidc) | **Get** /users/{user_id}/credentials_oidc | Get OIDC Auth Credential
*UserApi* | [**UserCredentialsSaml**](docs/UserApi.md#usercredentialssaml) | **Get** /users/{user_id}/credentials_saml | Get Saml Auth Credential
*UserApi* | [**UserCredentialsTotp**](docs/UserApi.md#usercredentialstotp) | **Get** /users/{user_id}/credentials_totp | Get Two-Factor Credential
*UserApi* | [**UserForCredential**](docs/UserApi.md#userforcredential) | **Get** /users/credential/{credential_type}/{credential_id} | Get User by Credential Id
*UserApi* | [**UserRoles**](docs/UserApi.md#userroles) | **Get** /users/{user_id}/roles | Get User Roles
*UserApi* | [**UserSession**](docs/UserApi.md#usersession) | **Get** /users/{user_id}/sessions/{session_id} | Get Web Login Session
*UserAttributeApi* | [**AllUserAttributeGroupValues**](docs/UserAttributeApi.md#alluserattributegroupvalues) | **Get** /user_attributes/{user_attribute_id}/group_values | Get User Attribute Group Values
*UserAttributeApi* | [**AllUserAttributes**](docs/UserAttributeApi.md#alluserattributes) | **Get** /user_attributes | Get All User Attributes
*UserAttributeApi* | [**CreateUserAttribute**](docs/UserAttributeApi.md#createuserattribute) | **Post** /user_attributes | Create User Attribute
*UserAttributeApi* | [**DeleteUserAttribute**](docs/UserAttributeApi.md#deleteuserattribute) | **Delete** /user_attributes/{user_attribute_id} | Delete User Attribute
*UserAttributeApi* | [**SetUserAttributeGroupValues**](docs/UserAttributeApi.md#setuserattributegroupvalues) | **Post** /user_attributes/{user_attribute_id}/group_values | Set User Attribute Group Values
*UserAttributeApi* | [**UpdateUserAttribute**](docs/UserAttributeApi.md#updateuserattribute) | **Patch** /user_attributes/{user_attribute_id} | Update User Attribute
*UserAttributeApi* | [**UserAttribute**](docs/UserAttributeApi.md#userattribute) | **Get** /user_attributes/{user_attribute_id} | Get User Attribute
*WorkspaceApi* | [**AllWorkspaces**](docs/WorkspaceApi.md#allworkspaces) | **Get** /workspaces | Get All Workspaces
*WorkspaceApi* | [**Workspace**](docs/WorkspaceApi.md#workspace) | **Get** /workspaces/{workspace_id} | Get Workspace


## Documentation For Models

 - [AccessToken](docs/AccessToken.md)
 - [ApiSession](docs/ApiSession.md)
 - [ApiVersion](docs/ApiVersion.md)
 - [ApiVersionElement](docs/ApiVersionElement.md)
 - [BackupConfiguration](docs/BackupConfiguration.md)
 - [ColorCollection](docs/ColorCollection.md)
 - [ColorStop](docs/ColorStop.md)
 - [ContentFavorite](docs/ContentFavorite.md)
 - [ContentMeta](docs/ContentMeta.md)
 - [ContentMetaGroupUser](docs/ContentMetaGroupUser.md)
 - [ContentValidation](docs/ContentValidation.md)
 - [ContentValidationDashboard](docs/ContentValidationDashboard.md)
 - [ContentValidationDashboardElement](docs/ContentValidationDashboardElement.md)
 - [ContentValidationDashboardFilter](docs/ContentValidationDashboardFilter.md)
 - [ContentValidationError](docs/ContentValidationError.md)
 - [ContentValidationFolder](docs/ContentValidationFolder.md)
 - [ContentValidationLook](docs/ContentValidationLook.md)
 - [ContentValidationSpace](docs/ContentValidationSpace.md)
 - [ContentValidatorError](docs/ContentValidatorError.md)
 - [ContentView](docs/ContentView.md)
 - [ContinuousPalette](docs/ContinuousPalette.md)
 - [CreateDashboardFilter](docs/CreateDashboardFilter.md)
 - [CreateDashboardRenderTask](docs/CreateDashboardRenderTask.md)
 - [CreateFolder](docs/CreateFolder.md)
 - [CreateQueryTask](docs/CreateQueryTask.md)
 - [CreateSpace](docs/CreateSpace.md)
 - [CredentialsApi3](docs/CredentialsApi3.md)
 - [CredentialsEmail](docs/CredentialsEmail.md)
 - [CredentialsEmbed](docs/CredentialsEmbed.md)
 - [CredentialsGoogle](docs/CredentialsGoogle.md)
 - [CredentialsLdap](docs/CredentialsLdap.md)
 - [CredentialsLookerOpenid](docs/CredentialsLookerOpenid.md)
 - [CredentialsOidc](docs/CredentialsOidc.md)
 - [CredentialsSaml](docs/CredentialsSaml.md)
 - [CredentialsTotp](docs/CredentialsTotp.md)
 - [CustomWelcomeEmail](docs/CustomWelcomeEmail.md)
 - [Dashboard](docs/Dashboard.md)
 - [DashboardAppearance](docs/DashboardAppearance.md)
 - [DashboardBase](docs/DashboardBase.md)
 - [DashboardElement](docs/DashboardElement.md)
 - [DashboardFilter](docs/DashboardFilter.md)
 - [DashboardLayout](docs/DashboardLayout.md)
 - [DashboardLayoutComponent](docs/DashboardLayoutComponent.md)
 - [DashboardLookml](docs/DashboardLookml.md)
 - [DataActionForm](docs/DataActionForm.md)
 - [DataActionFormField](docs/DataActionFormField.md)
 - [DataActionFormSelectOption](docs/DataActionFormSelectOption.md)
 - [DataActionRequest](docs/DataActionRequest.md)
 - [DataActionResponse](docs/DataActionResponse.md)
 - [DataActionUserState](docs/DataActionUserState.md)
 - [Datagroup](docs/Datagroup.md)
 - [DbConnection](docs/DbConnection.md)
 - [DbConnectionBase](docs/DbConnectionBase.md)
 - [DbConnectionOverride](docs/DbConnectionOverride.md)
 - [DbConnectionTestResult](docs/DbConnectionTestResult.md)
 - [DelegateOauthTest](docs/DelegateOauthTest.md)
 - [Dialect](docs/Dialect.md)
 - [DialectInfo](docs/DialectInfo.md)
 - [DialectInfoOptions](docs/DialectInfoOptions.md)
 - [DigestEmailSend](docs/DigestEmailSend.md)
 - [DigestEmails](docs/DigestEmails.md)
 - [DiscretePalette](docs/DiscretePalette.md)
 - [EmbedSsoUrl](docs/EmbedSsoUrl.md)
 - [EmbedSsoUrlParams](docs/EmbedSsoUrlParams.md)
 - [Error](docs/Error.md)
 - [Folder](docs/Folder.md)
 - [FolderBase](docs/FolderBase.md)
 - [GitBranch](docs/GitBranch.md)
 - [GitConnectionTest](docs/GitConnectionTest.md)
 - [GitConnectionTestResult](docs/GitConnectionTestResult.md)
 - [GitStatus](docs/GitStatus.md)
 - [Group](docs/Group.md)
 - [GroupIdForGroupInclusion](docs/GroupIdForGroupInclusion.md)
 - [GroupIdForGroupUserInclusion](docs/GroupIdForGroupUserInclusion.md)
 - [Homepage](docs/Homepage.md)
 - [HomepageItem](docs/HomepageItem.md)
 - [HomepageSection](docs/HomepageSection.md)
 - [ImportedProject](docs/ImportedProject.md)
 - [Integration](docs/Integration.md)
 - [IntegrationHub](docs/IntegrationHub.md)
 - [IntegrationParam](docs/IntegrationParam.md)
 - [IntegrationRequiredField](docs/IntegrationRequiredField.md)
 - [IntegrationTestResult](docs/IntegrationTestResult.md)
 - [InternalHelpResources](docs/InternalHelpResources.md)
 - [InternalHelpResourcesContent](docs/InternalHelpResourcesContent.md)
 - [LdapConfig](docs/LdapConfig.md)
 - [LdapConfigTestIssue](docs/LdapConfigTestIssue.md)
 - [LdapConfigTestResult](docs/LdapConfigTestResult.md)
 - [LdapGroupRead](docs/LdapGroupRead.md)
 - [LdapGroupWrite](docs/LdapGroupWrite.md)
 - [LdapUser](docs/LdapUser.md)
 - [LdapUserAttributeRead](docs/LdapUserAttributeRead.md)
 - [LdapUserAttributeWrite](docs/LdapUserAttributeWrite.md)
 - [LegacyFeature](docs/LegacyFeature.md)
 - [Locale](docs/Locale.md)
 - [LocalizationSettings](docs/LocalizationSettings.md)
 - [Look](docs/Look.md)
 - [LookBasic](docs/LookBasic.md)
 - [LookModel](docs/LookModel.md)
 - [LookWithDashboards](docs/LookWithDashboards.md)
 - [LookWithQuery](docs/LookWithQuery.md)
 - [LookmlModel](docs/LookmlModel.md)
 - [LookmlModelExplore](docs/LookmlModelExplore.md)
 - [LookmlModelExploreAccessFilter](docs/LookmlModelExploreAccessFilter.md)
 - [LookmlModelExploreAlias](docs/LookmlModelExploreAlias.md)
 - [LookmlModelExploreAlwaysFilter](docs/LookmlModelExploreAlwaysFilter.md)
 - [LookmlModelExploreConditionallyFilter](docs/LookmlModelExploreConditionallyFilter.md)
 - [LookmlModelExploreError](docs/LookmlModelExploreError.md)
 - [LookmlModelExploreField](docs/LookmlModelExploreField.md)
 - [LookmlModelExploreFieldEnumeration](docs/LookmlModelExploreFieldEnumeration.md)
 - [LookmlModelExploreFieldMapLayer](docs/LookmlModelExploreFieldMapLayer.md)
 - [LookmlModelExploreFieldSqlCase](docs/LookmlModelExploreFieldSqlCase.md)
 - [LookmlModelExploreFieldTimeInterval](docs/LookmlModelExploreFieldTimeInterval.md)
 - [LookmlModelExploreFieldset](docs/LookmlModelExploreFieldset.md)
 - [LookmlModelExploreJoins](docs/LookmlModelExploreJoins.md)
 - [LookmlModelExploreSet](docs/LookmlModelExploreSet.md)
 - [LookmlModelExploreSupportedMeasureType](docs/LookmlModelExploreSupportedMeasureType.md)
 - [LookmlModelNavExplore](docs/LookmlModelNavExplore.md)
 - [LookmlTest](docs/LookmlTest.md)
 - [LookmlTestResult](docs/LookmlTestResult.md)
 - [Manifest](docs/Manifest.md)
 - [MergeFields](docs/MergeFields.md)
 - [MergeQuery](docs/MergeQuery.md)
 - [MergeQuerySourceQuery](docs/MergeQuerySourceQuery.md)
 - [ModelSet](docs/ModelSet.md)
 - [ModelsNotValidated](docs/ModelsNotValidated.md)
 - [OidcConfig](docs/OidcConfig.md)
 - [OidcGroupRead](docs/OidcGroupRead.md)
 - [OidcGroupWrite](docs/OidcGroupWrite.md)
 - [OidcUserAttributeRead](docs/OidcUserAttributeRead.md)
 - [OidcUserAttributeWrite](docs/OidcUserAttributeWrite.md)
 - [PasswordConfig](docs/PasswordConfig.md)
 - [Permission](docs/Permission.md)
 - [PermissionSet](docs/PermissionSet.md)
 - [Project](docs/Project.md)
 - [ProjectError](docs/ProjectError.md)
 - [ProjectFile](docs/ProjectFile.md)
 - [ProjectValidation](docs/ProjectValidation.md)
 - [ProjectValidationCache](docs/ProjectValidationCache.md)
 - [ProjectWorkspace](docs/ProjectWorkspace.md)
 - [Query](docs/Query.md)
 - [QueryTask](docs/QueryTask.md)
 - [RenderTask](docs/RenderTask.md)
 - [RepositoryCredential](docs/RepositoryCredential.md)
 - [ResultMakerFilterables](docs/ResultMakerFilterables.md)
 - [ResultMakerFilterablesListen](docs/ResultMakerFilterablesListen.md)
 - [ResultMakerWithIdVisConfigAndDynamicFields](docs/ResultMakerWithIdVisConfigAndDynamicFields.md)
 - [Role](docs/Role.md)
 - [RunningQueries](docs/RunningQueries.md)
 - [SamlConfig](docs/SamlConfig.md)
 - [SamlGroupRead](docs/SamlGroupRead.md)
 - [SamlGroupWrite](docs/SamlGroupWrite.md)
 - [SamlMetadataParseResult](docs/SamlMetadataParseResult.md)
 - [SamlUserAttributeRead](docs/SamlUserAttributeRead.md)
 - [SamlUserAttributeWrite](docs/SamlUserAttributeWrite.md)
 - [ScheduledPlan](docs/ScheduledPlan.md)
 - [ScheduledPlanDestination](docs/ScheduledPlanDestination.md)
 - [Session](docs/Session.md)
 - [SessionConfig](docs/SessionConfig.md)
 - [Snippet](docs/Snippet.md)
 - [Space](docs/Space.md)
 - [SpaceBase](docs/SpaceBase.md)
 - [SqlQuery](docs/SqlQuery.md)
 - [SqlQueryCreate](docs/SqlQueryCreate.md)
 - [Theme](docs/Theme.md)
 - [ThemeSettings](docs/ThemeSettings.md)
 - [Timezone](docs/Timezone.md)
 - [UpdateFolder](docs/UpdateFolder.md)
 - [UpdateSpace](docs/UpdateSpace.md)
 - [User](docs/User.md)
 - [UserAttribute](docs/UserAttribute.md)
 - [UserAttributeGroupValue](docs/UserAttributeGroupValue.md)
 - [UserAttributeWithValue](docs/UserAttributeWithValue.md)
 - [UserIdOnly](docs/UserIdOnly.md)
 - [UserLoginLockout](docs/UserLoginLockout.md)
 - [UserPublic](docs/UserPublic.md)
 - [ValidationError](docs/ValidationError.md)
 - [ValidationErrorDetail](docs/ValidationErrorDetail.md)
 - [WelcomeEmailTest](docs/WelcomeEmailTest.md)
 - [WhitelabelConfiguration](docs/WhitelabelConfiguration.md)
 - [Workspace](docs/Workspace.md)
 - [WriteScheduledPlan](docs/WriteScheduledPlan.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Author

support@looker.com

