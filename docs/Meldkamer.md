# Meldkamer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Naam** | **string** |  | 
**Afkorting** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMeldkamer

`func NewMeldkamer(id int32, naam string, ) *Meldkamer`

NewMeldkamer instantiates a new Meldkamer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeldkamerWithDefaults

`func NewMeldkamerWithDefaults() *Meldkamer`

NewMeldkamerWithDefaults instantiates a new Meldkamer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Meldkamer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Meldkamer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Meldkamer) SetId(v int32)`

SetId sets Id field to given value.


### GetNaam

`func (o *Meldkamer) GetNaam() string`

GetNaam returns the Naam field if non-nil, zero value otherwise.

### GetNaamOk

`func (o *Meldkamer) GetNaamOk() (*string, bool)`

GetNaamOk returns a tuple with the Naam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaam

`func (o *Meldkamer) SetNaam(v string)`

SetNaam sets Naam field to given value.


### GetAfkorting

`func (o *Meldkamer) GetAfkorting() string`

GetAfkorting returns the Afkorting field if non-nil, zero value otherwise.

### GetAfkortingOk

`func (o *Meldkamer) GetAfkortingOk() (*string, bool)`

GetAfkortingOk returns a tuple with the Afkorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfkorting

`func (o *Meldkamer) SetAfkorting(v string)`

SetAfkorting sets Afkorting field to given value.

### HasAfkorting

`func (o *Meldkamer) HasAfkorting() bool`

HasAfkorting returns a boolean if a field has been set.

### SetAfkortingNil

`func (o *Meldkamer) SetAfkortingNil(b bool)`

 SetAfkortingNil sets the value for Afkorting to be an explicit nil

### UnsetAfkorting
`func (o *Meldkamer) UnsetAfkorting()`

UnsetAfkorting ensures that no value is present for Afkorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


