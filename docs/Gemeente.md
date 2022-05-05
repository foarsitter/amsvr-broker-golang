# Gemeente

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Naam** | **string** |  | 
**Regio** | **int32** |  | 

## Methods

### NewGemeente

`func NewGemeente(id int32, naam string, regio int32, ) *Gemeente`

NewGemeente instantiates a new Gemeente object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGemeenteWithDefaults

`func NewGemeenteWithDefaults() *Gemeente`

NewGemeenteWithDefaults instantiates a new Gemeente object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Gemeente) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Gemeente) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Gemeente) SetId(v int32)`

SetId sets Id field to given value.


### GetNaam

`func (o *Gemeente) GetNaam() string`

GetNaam returns the Naam field if non-nil, zero value otherwise.

### GetNaamOk

`func (o *Gemeente) GetNaamOk() (*string, bool)`

GetNaamOk returns a tuple with the Naam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaam

`func (o *Gemeente) SetNaam(v string)`

SetNaam sets Naam field to given value.


### GetRegio

`func (o *Gemeente) GetRegio() int32`

GetRegio returns the Regio field if non-nil, zero value otherwise.

### GetRegioOk

`func (o *Gemeente) GetRegioOk() (*int32, bool)`

GetRegioOk returns a tuple with the Regio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegio

`func (o *Gemeente) SetRegio(v int32)`

SetRegio sets Regio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


