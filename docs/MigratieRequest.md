# MigratieRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NieuweAansluitnummer** | **string** |  | 
**Straat** | Pointer to **string** |  | [optional] 
**Huisnummer** | Pointer to **string** |  | [optional] 
**Atsp** | **int32** |  | 
**Memo** | Pointer to **string** |  | [optional] 

## Methods

### NewMigratieRequest

`func NewMigratieRequest(nieuweAansluitnummer string, atsp int32, ) *MigratieRequest`

NewMigratieRequest instantiates a new MigratieRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigratieRequestWithDefaults

`func NewMigratieRequestWithDefaults() *MigratieRequest`

NewMigratieRequestWithDefaults instantiates a new MigratieRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNieuweAansluitnummer

`func (o *MigratieRequest) GetNieuweAansluitnummer() string`

GetNieuweAansluitnummer returns the NieuweAansluitnummer field if non-nil, zero value otherwise.

### GetNieuweAansluitnummerOk

`func (o *MigratieRequest) GetNieuweAansluitnummerOk() (*string, bool)`

GetNieuweAansluitnummerOk returns a tuple with the NieuweAansluitnummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNieuweAansluitnummer

`func (o *MigratieRequest) SetNieuweAansluitnummer(v string)`

SetNieuweAansluitnummer sets NieuweAansluitnummer field to given value.


### GetStraat

`func (o *MigratieRequest) GetStraat() string`

GetStraat returns the Straat field if non-nil, zero value otherwise.

### GetStraatOk

`func (o *MigratieRequest) GetStraatOk() (*string, bool)`

GetStraatOk returns a tuple with the Straat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStraat

`func (o *MigratieRequest) SetStraat(v string)`

SetStraat sets Straat field to given value.

### HasStraat

`func (o *MigratieRequest) HasStraat() bool`

HasStraat returns a boolean if a field has been set.

### GetHuisnummer

`func (o *MigratieRequest) GetHuisnummer() string`

GetHuisnummer returns the Huisnummer field if non-nil, zero value otherwise.

### GetHuisnummerOk

`func (o *MigratieRequest) GetHuisnummerOk() (*string, bool)`

GetHuisnummerOk returns a tuple with the Huisnummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuisnummer

`func (o *MigratieRequest) SetHuisnummer(v string)`

SetHuisnummer sets Huisnummer field to given value.

### HasHuisnummer

`func (o *MigratieRequest) HasHuisnummer() bool`

HasHuisnummer returns a boolean if a field has been set.

### GetAtsp

`func (o *MigratieRequest) GetAtsp() int32`

GetAtsp returns the Atsp field if non-nil, zero value otherwise.

### GetAtspOk

`func (o *MigratieRequest) GetAtspOk() (*int32, bool)`

GetAtspOk returns a tuple with the Atsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtsp

`func (o *MigratieRequest) SetAtsp(v int32)`

SetAtsp sets Atsp field to given value.


### GetMemo

`func (o *MigratieRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *MigratieRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *MigratieRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *MigratieRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


