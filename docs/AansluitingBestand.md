# AansluitingBestand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Naam** | Pointer to **string** |  | [optional] 
**Omschrijving** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**NullableAansluitingBestandTypeEnum**](AansluitingBestandTypeEnum.md) | overige &#x3D; Overige, pve &#x3D; PVE, upd &#x3D; UPD, nva &#x3D; NVA (nota van aanvulling), nvw &#x3D; NVW (nota van wijzigingen), gms_mutatierapport &#x3D; GMS mutatierapport | [optional] 
**File** | Pointer to **NullableString** |  | [optional] 
**Created** | **time.Time** |  | [readonly] 
**Modified** | **time.Time** |  | [readonly] 

## Methods

### NewAansluitingBestand

`func NewAansluitingBestand(id int32, created time.Time, modified time.Time, ) *AansluitingBestand`

NewAansluitingBestand instantiates a new AansluitingBestand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAansluitingBestandWithDefaults

`func NewAansluitingBestandWithDefaults() *AansluitingBestand`

NewAansluitingBestandWithDefaults instantiates a new AansluitingBestand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AansluitingBestand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AansluitingBestand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AansluitingBestand) SetId(v int32)`

SetId sets Id field to given value.


### GetNaam

`func (o *AansluitingBestand) GetNaam() string`

GetNaam returns the Naam field if non-nil, zero value otherwise.

### GetNaamOk

`func (o *AansluitingBestand) GetNaamOk() (*string, bool)`

GetNaamOk returns a tuple with the Naam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaam

`func (o *AansluitingBestand) SetNaam(v string)`

SetNaam sets Naam field to given value.

### HasNaam

`func (o *AansluitingBestand) HasNaam() bool`

HasNaam returns a boolean if a field has been set.

### GetOmschrijving

`func (o *AansluitingBestand) GetOmschrijving() string`

GetOmschrijving returns the Omschrijving field if non-nil, zero value otherwise.

### GetOmschrijvingOk

`func (o *AansluitingBestand) GetOmschrijvingOk() (*string, bool)`

GetOmschrijvingOk returns a tuple with the Omschrijving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmschrijving

`func (o *AansluitingBestand) SetOmschrijving(v string)`

SetOmschrijving sets Omschrijving field to given value.

### HasOmschrijving

`func (o *AansluitingBestand) HasOmschrijving() bool`

HasOmschrijving returns a boolean if a field has been set.

### SetOmschrijvingNil

`func (o *AansluitingBestand) SetOmschrijvingNil(b bool)`

 SetOmschrijvingNil sets the value for Omschrijving to be an explicit nil

### UnsetOmschrijving
`func (o *AansluitingBestand) UnsetOmschrijving()`

UnsetOmschrijving ensures that no value is present for Omschrijving, not even an explicit nil
### GetType

`func (o *AansluitingBestand) GetType() AansluitingBestandTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AansluitingBestand) GetTypeOk() (*AansluitingBestandTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AansluitingBestand) SetType(v AansluitingBestandTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *AansluitingBestand) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AansluitingBestand) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AansluitingBestand) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFile

`func (o *AansluitingBestand) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AansluitingBestand) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AansluitingBestand) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *AansluitingBestand) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *AansluitingBestand) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *AansluitingBestand) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetCreated

`func (o *AansluitingBestand) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AansluitingBestand) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AansluitingBestand) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *AansluitingBestand) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AansluitingBestand) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AansluitingBestand) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


