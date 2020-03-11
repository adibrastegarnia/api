# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/protocols/log/log.proto](#atomix/protocols/log/log.proto)
    - [CompactionSpec](#atomix.protocols.log.CompactionSpec)
    - [LogProtocol](#atomix.protocols.log.LogProtocol)
    - [StorageSpec](#atomix.protocols.log.StorageSpec)
  
    - [MemberGroupStrategy](#atomix.protocols.log.MemberGroupStrategy)
    - [StorageLevel](#atomix.protocols.log.StorageLevel)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/protocols/log/log.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/protocols/log/log.proto



<a name="atomix.protocols.log.CompactionSpec"></a>

### CompactionSpec
Partition group compaction configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dynamic | [bool](#bool) |  |  |
| free_disk_buffer | [double](#double) |  |  |






<a name="atomix.protocols.log.LogProtocol"></a>

### LogProtocol
Log protocol configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member_group_strategy | [MemberGroupStrategy](#atomix.protocols.log.MemberGroupStrategy) |  |  |
| storage | [StorageSpec](#atomix.protocols.log.StorageSpec) |  |  |
| compaction | [CompactionSpec](#atomix.protocols.log.CompactionSpec) |  |  |






<a name="atomix.protocols.log.StorageSpec"></a>

### StorageSpec
Partition group storage configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| max_entry_size | [uint32](#uint32) |  |  |
| segment_size | [uint32](#uint32) |  |  |
| level | [StorageLevel](#atomix.protocols.log.StorageLevel) |  |  |
| flush_on_commit | [bool](#bool) |  |  |





 


<a name="atomix.protocols.log.MemberGroupStrategy"></a>

### MemberGroupStrategy
Member group strategy

| Name | Number | Description |
| ---- | ------ | ----------- |
| HOST_AWARE | 0 |  |
| RACK_AWARE | 1 |  |
| ZONE_AWARE | 2 |  |



<a name="atomix.protocols.log.StorageLevel"></a>

### StorageLevel
Storage level

| Name | Number | Description |
| ---- | ------ | ----------- |
| DISK | 0 |  |
| MAPPED | 1 |  |


 

 

 



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

