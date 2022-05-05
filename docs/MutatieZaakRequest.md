# MutatieZaakRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aansluiting** | **int32** |  | 
**Type** | Pointer to [**MutatieZaakTypeEnum**](MutatieZaakTypeEnum.md) |  | [optional] 
**Reden** | Pointer to **NullableString** |  | [optional] 
**OmschrijvingVanWijziging** | **string** |  | 
**BouwdelenToevoegen** | Pointer to **bool** |  | [optional] 
**Sectoren** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewMutatieZaakRequest

`func NewMutatieZaakRequest(aansluiting int32, omschrijvingVanWijziging string, ) *MutatieZaakRequest`

NewMutatieZaakRequest instantiates a new MutatieZaakRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutatieZaakRequestWithDefaults

`func NewMutatieZaakRequestWithDefaults() *MutatieZaakRequest`

NewMutatieZaakRequestWithDefaults instantiates a new MutatieZaakRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAansluiting

`func (o *MutatieZaakRequest) GetAansluiting() int32`

GetAansluiting returns the Aansluiting field if non-nil, zero value otherwise.

### GetAansluitingOk

`func (o *MutatieZaakRequest) GetAansluitingOk() (*int32, bool)`

GetAansluitingOk returns a tuple with the Aansluiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAansluiting

`func (o *MutatieZaakRequest) SetAansluiting(v int32)`

SetAansluiting sets Aansluiting field to given value.


### GetType

`func (o *MutatieZaakRequest) GetType() MutatieZaakTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MutatieZaakRequest) GetTypeOk() (*MutatieZaakTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MutatieZaakRequest) SetType(v MutatieZaakTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *MutatieZaakRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReden

`func (o *MutatieZaakRequest) GetReden() string`

GetReden returns the Reden field if non-nil, zero value otherwise.

### GetRedenOk

`func (o *MutatieZaakRequest) GetRedenOk() (*string, bool)`

GetRedenOk returns a tuple with the Reden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReden

`func (o *MutatieZaakRequest) SetReden(v string)`

SetReden sets Reden field to given value.

### HasReden

`func (o *MutatieZaakRequest) HasReden() bool`

HasReden returns a boolean if a field has been set.

### SetRedenNil

`func (o *MutatieZaakRequest) SetRedenNil(b bool)`

 SetRedenNil sets the value for Reden to be an explicit nil

### UnsetReden
`func (o *MutatieZaakRequest) UnsetReden()`

UnsetReden ensures that no value is present for Reden, not even an explicit nil
### GetOmschrijvingVanWijziging

`func (o *MutatieZaakRequest) GetOmschrijvingVanWijziging() string`

GetOmschrijvingVanWijziging returns the OmschrijvingVanWijziging field if non-nil, zero value otherwise.

### GetOmschrijvingVanWijzigingOk

`func (o *MutatieZaakRequest) GetOmschrijvingVanWijzigingOk() (*string, bool)`

GetOmschrijvingVanWijzigingOk returns a tuple with the OmschrijvingVanWijziging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmschrijvingVanWijziging

`func (o *MutatieZaakRequest) SetOmschrijvingVanWijziging(v string)`

SetOmschrijvingVanWijziging sets OmschrijvingVanWijziging field to given value.


### GetBouwdelenToevoegen

`func (o *MutatieZaakRequest) GetBouwdelenToevoegen() bool`

GetBouwdelenToevoegen returns the BouwdelenToevoegen field if non-nil, zero value otherwise.

### GetBouwdelenToevoegenOk

`func (o *MutatieZaakRequest) GetBouwdelenToevoegenOk() (*bool, bool)`

GetBouwdelenToevoegenOk returns a tuple with the BouwdelenToevoegen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBouwdelenToevoegen

`func (o *MutatieZaakRequest) SetBouwdelenToevoegen(v bool)`

SetBouwdelenToevoegen sets BouwdelenToevoegen field to given value.

### HasBouwdelenToevoegen

`func (o *MutatieZaakRequest) HasBouwdelenToevoegen() bool`

HasBouwdelenToevoegen returns a boolean if a field has been set.

### GetSectoren

`func (o *MutatieZaakRequest) GetSectoren() []int32`

GetSectoren returns the Sectoren field if non-nil, zero value otherwise.

### GetSectorenOk

`func (o *MutatieZaakRequest) GetSectorenOk() (*[]int32, bool)`

GetSectorenOk returns a tuple with the Sectoren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectoren

`func (o *MutatieZaakRequest) SetSectoren(v []int32)`

SetSectoren sets Sectoren field to given value.

### HasSectoren

`func (o *MutatieZaakRequest) HasSectoren() bool`

HasSectoren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


