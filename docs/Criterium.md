# Criterium

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Naam** | **int32** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCriterium

`func NewCriterium(id int32, naam int32, ) *Criterium`

NewCriterium instantiates a new Criterium object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriteriumWithDefaults

`func NewCriteriumWithDefaults() *Criterium`

NewCriteriumWithDefaults instantiates a new Criterium object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Criterium) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Criterium) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Criterium) SetId(v int32)`

SetId sets Id field to given value.


### GetNaam

`func (o *Criterium) GetNaam() int32`

GetNaam returns the Naam field if non-nil, zero value otherwise.

### GetNaamOk

`func (o *Criterium) GetNaamOk() (*int32, bool)`

GetNaamOk returns a tuple with the Naam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaam

`func (o *Criterium) SetNaam(v int32)`

SetNaam sets Naam field to given value.


### GetDescription

`func (o *Criterium) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Criterium) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Criterium) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Criterium) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Criterium) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Criterium) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


