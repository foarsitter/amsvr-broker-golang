# AansluitingBestandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Naam** | Pointer to **string** |  | [optional] 
**Omschrijving** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**NullableAansluitingBestandTypeEnum**](AansluitingBestandTypeEnum.md) | overige &#x3D; Overige, pve &#x3D; PVE, upd &#x3D; UPD, nva &#x3D; NVA (nota van aanvulling), nvw &#x3D; NVW (nota van wijzigingen), gms_mutatierapport &#x3D; GMS mutatierapport | [optional] 
**File** | Pointer to **Nullable*os.File** |  | [optional] 

## Methods

### NewAansluitingBestandRequest

`func NewAansluitingBestandRequest() *AansluitingBestandRequest`

NewAansluitingBestandRequest instantiates a new AansluitingBestandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAansluitingBestandRequestWithDefaults

`func NewAansluitingBestandRequestWithDefaults() *AansluitingBestandRequest`

NewAansluitingBestandRequestWithDefaults instantiates a new AansluitingBestandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNaam

`func (o *AansluitingBestandRequest) GetNaam() string`

GetNaam returns the Naam field if non-nil, zero value otherwise.

### GetNaamOk

`func (o *AansluitingBestandRequest) GetNaamOk() (*string, bool)`

GetNaamOk returns a tuple with the Naam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaam

`func (o *AansluitingBestandRequest) SetNaam(v string)`

SetNaam sets Naam field to given value.

### HasNaam

`func (o *AansluitingBestandRequest) HasNaam() bool`

HasNaam returns a boolean if a field has been set.

### GetOmschrijving

`func (o *AansluitingBestandRequest) GetOmschrijving() string`

GetOmschrijving returns the Omschrijving field if non-nil, zero value otherwise.

### GetOmschrijvingOk

`func (o *AansluitingBestandRequest) GetOmschrijvingOk() (*string, bool)`

GetOmschrijvingOk returns a tuple with the Omschrijving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmschrijving

`func (o *AansluitingBestandRequest) SetOmschrijving(v string)`

SetOmschrijving sets Omschrijving field to given value.

### HasOmschrijving

`func (o *AansluitingBestandRequest) HasOmschrijving() bool`

HasOmschrijving returns a boolean if a field has been set.

### SetOmschrijvingNil

`func (o *AansluitingBestandRequest) SetOmschrijvingNil(b bool)`

 SetOmschrijvingNil sets the value for Omschrijving to be an explicit nil

### UnsetOmschrijving
`func (o *AansluitingBestandRequest) UnsetOmschrijving()`

UnsetOmschrijving ensures that no value is present for Omschrijving, not even an explicit nil
### GetType

`func (o *AansluitingBestandRequest) GetType() AansluitingBestandTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AansluitingBestandRequest) GetTypeOk() (*AansluitingBestandTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AansluitingBestandRequest) SetType(v AansluitingBestandTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *AansluitingBestandRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AansluitingBestandRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AansluitingBestandRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFile

`func (o *AansluitingBestandRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AansluitingBestandRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AansluitingBestandRequest) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *AansluitingBestandRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *AansluitingBestandRequest) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *AansluitingBestandRequest) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


