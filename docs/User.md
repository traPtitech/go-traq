# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | ユーザーUUID | [optional] 
**Name** | **string** | traQ ID | [optional] 
**DisplayName** | **string** | 表示名 | [optional] 
**IconFileId** | **string** | アイコンファイルUUID | [optional] 
**Bot** | **bool** | BOTかどうか | [optional] 
**TwitterId** | **string** | ツイッターID | [optional] 
**LastOnline** | [**time.Time**](time.Time.md) | 最終オンライン日時(オンラインに１度もなってない場合はnull) | [optional] 
**IsOnline** | **bool** | 現在オンラインかどうか | [optional] 
**Suspended** | **bool** | アカウントが停止中かどうか | [optional] 
**AccountStatus** | **int32** | アカウントの状態 (0:停止,1:有効,2:一時停止) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


