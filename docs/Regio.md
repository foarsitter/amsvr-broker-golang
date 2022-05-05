# Regio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Naam** | **string** |  | 
**Meldkamer** | **int32** |  | 
**Code** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRegio

`func NewRegio(id int32, naam string, meldkamer int32, ) *Regio`

NewRegio instantiates a new Regio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegioWithDefaults

`func NewRegioWithDefaults() *Regio`

NewRegioWithDefaults instantiates a new Regio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Regio) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Regio) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Regio) SetId(v int32)`

SetId sets Id field to given value.


### GetNaam

`func (o *Regio) GetNaam() string`

GetNaam returns the Naam field if non-nil, zero value otherwise.

### GetNaamOk

`func (o *Regio) GetNaamOk() (*string, bool)`

GetNaamOk returns a tuple with the Naam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaam

`func (o *Regio) SetNaam(v string)`

SetNaam sets Naam field to given value.


### GetMeldkamer

`func (o *Regio) GetMeldkamer() int32`

GetMeldkamer returns the Meldkamer field if non-nil, zero value otherwise.

### GetMeldkamerOk

`func (o *Regio) GetMeldkamerOk() (*int32, bool)`

GetMeldkamerOk returns a tuple with the Meldkamer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeldkamer

`func (o *Regio) SetMeldkamer(v int32)`

SetMeldkamer sets Meldkamer field to given value.


### GetCode

`func (o *Regio) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Regio) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Regio) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Regio) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Regio) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Regio) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


