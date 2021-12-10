# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [activity.proto](#activity.proto)
    - [ActivityInfo](#activity.ActivityInfo)
    - [ListAllRequest](#activity.ListAllRequest)
    - [ListAllResponse](#activity.ListAllResponse)
    - [ListByTypeRequest](#activity.ListByTypeRequest)
    - [ListByTypeResponse](#activity.ListByTypeResponse)
  
    - [Activity](#activity.Activity)
  
- [Scalar Value Types](#scalar-value-types)



<a name="activity.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## activity.proto



<a name="activity.ActivityInfo"></a>

### ActivityInfo
活动信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| type | [int64](#int64) |  | 活动类型 |
| level | [int64](#int64) |  | 开放等级 |
| status | [int64](#int64) |  | 活动撞他 |
| startAt | [int64](#int64) |  | 开始时间（Unix时间戳） |
| endAt | [int64](#int64) |  | 结束时间（Unix时间戳） |






<a name="activity.ListAllRequest"></a>

### ListAllRequest
列出所有活动的请求信息






<a name="activity.ListAllResponse"></a>

### ListAllResponse
列出所有活动的相应信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| list | [ActivityInfo](#activity.ActivityInfo) | repeated | 活动列表 |






<a name="activity.ListByTypeRequest"></a>

### ListByTypeRequest
按类型列出活动的请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [int64](#int64) |  |  |






<a name="activity.ListByTypeResponse"></a>

### ListByTypeResponse
按类型列出活动的相应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| list | [ActivityInfo](#activity.ActivityInfo) | repeated | 活动列表 |





 

 

 


<a name="activity.Activity"></a>

### Activity


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListAll | [ListAllRequest](#activity.ListAllRequest) | [ListAllResponse](#activity.ListAllResponse) | 列出所有活动 |
| ListByType | [ListByTypeRequest](#activity.ListByTypeRequest) | [ListByTypeResponse](#activity.ListByTypeResponse) | 按类型列出活动（比如：底部） |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

