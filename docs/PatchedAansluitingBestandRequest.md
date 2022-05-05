# PatchedAansluitingBestandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Naam** | Pointer to **string** |  | [optional] 
**Omschrijving** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**NullableAansluitingBestandTypeEnum**](AansluitingBestandTypeEnum.md) | overige &#x3D; Overige, pve &#x3D; PVE, upd &#x3D; UPD, nva &#x3D; NVA (nota van aanvulling), nvw &#x3D; NVW (nota van wijzigingen), gms_mutatierapport &#x3D; GMS mutatierapport | [optional] 
**File** | Pointer to **Nullable*os.File** |  | [optional] 

## Methods

### NewPatchedAansluitingBestandRequest

`func NewPatchedAansluitingBestandRequest() *PatchedAansluitingBestandRequest`

NewPatchedAansluitingBestandRequest instantiates a new PatchedAansluitingBestandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAansluitingBestandRequestWithDefaults

`func NewPatchedAansluitingBestandRequestWithDefaults() *PatchedAansluitingBestandRequest`

NewPatchedAansluitingBestandRequestWithDefaults instantiates a new PatchedAansluitingBestandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNaam

`func (o *PatchedAansluitingBestandRequest) GetNaam() string`

GetNaam returns the Naam field if non-nil, zero value otherwise.

### GetNaamOk

`func (o *PatchedAansluitingBestandRequest) GetNaamOk() (*string, bool)`

GetNaamOk returns a tuple with the Naam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaam

`func (o *PatchedAansluitingBestandRequest) SetNaam(v string)`

SetNaam sets Naam field to given value.

### HasNaam

`func (o *PatchedAansluitingBestandRequest) HasNaam() bool`

HasNaam returns a boolean if a field has been set.

### GetOmschrijving

`func (o *PatchedAansluitingBestandRequest) GetOmschrijving() string`

GetOmschrijving returns the Omschrijving field if non-nil, zero value otherwise.

### GetOmschrijvingOk

`func (o *PatchedAansluitingBestandRequest) GetOmschrijvingOk() (*string, bool)`

GetOmschrijvingOk returns a tuple with the Omschrijving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmschrijving

`func (o *PatchedAansluitingBestandRequest) SetOmschrijving(v string)`

SetOmschrijving sets Omschrijving field to given value.

### HasOmschrijving

`func (o *PatchedAansluitingBestandRequest) HasOmschrijving() bool`

HasOmschrijving returns a boolean if a field has been set.

### SetOmschrijvingNil

`func (o *PatchedAansluitingBestandRequest) SetOmschrijvingNil(b bool)`

 SetOmschrijvingNil sets the value for Omschrijving to be an explicit nil

### UnsetOmschrijving
`func (o *PatchedAansluitingBestandRequest) UnsetOmschrijving()`

UnsetOmschrijving ensures that no value is present for Omschrijving, not even an explicit nil
### GetType

`func (o *PatchedAansluitingBestandRequest) GetType() AansluitingBestandTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedAansluitingBestandRequest) GetTypeOk() (*AansluitingBestandTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedAansluitingBestandRequest) SetType(v AansluitingBestandTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedAansluitingBestandRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PatchedAansluitingBestandRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PatchedAansluitingBestandRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFile

`func (o *PatchedAansluitingBestandRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *PatchedAansluitingBestandRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *PatchedAansluitingBestandRequest) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *PatchedAansluitingBestandRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *PatchedAansluitingBestandRequest) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *PatchedAansluitingBestandRequest) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


