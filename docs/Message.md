# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | メッセージUUID | [optional] 
**UserId** | **string** | 投稿者UUID | [optional] 
**ParentChannelId** | **string** | 投稿先チャンネルUUID | [optional] 
**Content** | **string** | メッセージ本体 | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | メッセージ投稿日時 | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | メッセージ更新日時 | [optional] 
**Pin** | **bool** | ピン留めされているかどうか | [optional] 
**Reported** | **bool** | 自分が通報しているかどうか | [optional] 
**StampList** | [**[]MessageStamp**](MessageStamp.md) | メッセージスタンプ配列 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


