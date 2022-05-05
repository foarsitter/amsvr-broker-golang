# ATSP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Naam** | **string** |  | 
**Code** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewATSP

`func NewATSP(id int32, naam string, ) *ATSP`

NewATSP instantiates a new ATSP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewATSPWithDefaults

`func NewATSPWithDefaults() *ATSP`

NewATSPWithDefaults instantiates a new ATSP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ATSP) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ATSP) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ATSP) SetId(v int32)`

SetId sets Id field to given value.


### GetNaam

`func (o *ATSP) GetNaam() string`

GetNaam returns the Naam field if non-nil, zero value otherwise.

### GetNaamOk

`func (o *ATSP) GetNaamOk() (*string, bool)`

GetNaamOk returns a tuple with the Naam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaam

`func (o *ATSP) SetNaam(v string)`

SetNaam sets Naam field to given value.


### GetCode

`func (o *ATSP) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ATSP) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ATSP) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ATSP) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ATSP) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ATSP) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


