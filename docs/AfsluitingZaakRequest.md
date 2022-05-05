# AfsluitingZaakRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aansluiting** | **int32** |  | 
**Reden** | **string** |  | 
**RedenVanAfwijzing** | Pointer to **NullableString** |  | [optional] 
**GeplandeDatum** | **string** | Voor wanneer staat de afsluiting gepland? | 

## Methods

### NewAfsluitingZaakRequest

`func NewAfsluitingZaakRequest(aansluiting int32, reden string, geplandeDatum string, ) *AfsluitingZaakRequest`

NewAfsluitingZaakRequest instantiates a new AfsluitingZaakRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfsluitingZaakRequestWithDefaults

`func NewAfsluitingZaakRequestWithDefaults() *AfsluitingZaakRequest`

NewAfsluitingZaakRequestWithDefaults instantiates a new AfsluitingZaakRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAansluiting

`func (o *AfsluitingZaakRequest) GetAansluiting() int32`

GetAansluiting returns the Aansluiting field if non-nil, zero value otherwise.

### GetAansluitingOk

`func (o *AfsluitingZaakRequest) GetAansluitingOk() (*int32, bool)`

GetAansluitingOk returns a tuple with the Aansluiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAansluiting

`func (o *AfsluitingZaakRequest) SetAansluiting(v int32)`

SetAansluiting sets Aansluiting field to given value.


### GetReden

`func (o *AfsluitingZaakRequest) GetReden() string`

GetReden returns the Reden field if non-nil, zero value otherwise.

### GetRedenOk

`func (o *AfsluitingZaakRequest) GetRedenOk() (*string, bool)`

GetRedenOk returns a tuple with the Reden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReden

`func (o *AfsluitingZaakRequest) SetReden(v string)`

SetReden sets Reden field to given value.


### GetRedenVanAfwijzing

`func (o *AfsluitingZaakRequest) GetRedenVanAfwijzing() string`

GetRedenVanAfwijzing returns the RedenVanAfwijzing field if non-nil, zero value otherwise.

### GetRedenVanAfwijzingOk

`func (o *AfsluitingZaakRequest) GetRedenVanAfwijzingOk() (*string, bool)`

GetRedenVanAfwijzingOk returns a tuple with the RedenVanAfwijzing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedenVanAfwijzing

`func (o *AfsluitingZaakRequest) SetRedenVanAfwijzing(v string)`

SetRedenVanAfwijzing sets RedenVanAfwijzing field to given value.

### HasRedenVanAfwijzing

`func (o *AfsluitingZaakRequest) HasRedenVanAfwijzing() bool`

HasRedenVanAfwijzing returns a boolean if a field has been set.

### SetRedenVanAfwijzingNil

`func (o *AfsluitingZaakRequest) SetRedenVanAfwijzingNil(b bool)`

 SetRedenVanAfwijzingNil sets the value for RedenVanAfwijzing to be an explicit nil

### UnsetRedenVanAfwijzing
`func (o *AfsluitingZaakRequest) UnsetRedenVanAfwijzing()`

UnsetRedenVanAfwijzing ensures that no value is present for RedenVanAfwijzing, not even an explicit nil
### GetGeplandeDatum

`func (o *AfsluitingZaakRequest) GetGeplandeDatum() string`

GetGeplandeDatum returns the GeplandeDatum field if non-nil, zero value otherwise.

### GetGeplandeDatumOk

`func (o *AfsluitingZaakRequest) GetGeplandeDatumOk() (*string, bool)`

GetGeplandeDatumOk returns a tuple with the GeplandeDatum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeplandeDatum

`func (o *AfsluitingZaakRequest) SetGeplandeDatum(v string)`

SetGeplandeDatum sets GeplandeDatum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


