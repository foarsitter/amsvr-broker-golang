# PaginatedWaarschuwingsadresList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]Waarschuwingsadres**](Waarschuwingsadres.md) |  | [optional] 

## Methods

### NewPaginatedWaarschuwingsadresList

`func NewPaginatedWaarschuwingsadresList() *PaginatedWaarschuwingsadresList`

NewPaginatedWaarschuwingsadresList instantiates a new PaginatedWaarschuwingsadresList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedWaarschuwingsadresListWithDefaults

`func NewPaginatedWaarschuwingsadresListWithDefaults() *PaginatedWaarschuwingsadresList`

NewPaginatedWaarschuwingsadresListWithDefaults instantiates a new PaginatedWaarschuwingsadresList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedWaarschuwingsadresList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedWaarschuwingsadresList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedWaarschuwingsadresList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedWaarschuwingsadresList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedWaarschuwingsadresList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedWaarschuwingsadresList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedWaarschuwingsadresList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedWaarschuwingsadresList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedWaarschuwingsadresList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedWaarschuwingsadresList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedWaarschuwingsadresList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedWaarschuwingsadresList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedWaarschuwingsadresList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedWaarschuwingsadresList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedWaarschuwingsadresList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedWaarschuwingsadresList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedWaarschuwingsadresList) GetResults() []Waarschuwingsadres`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedWaarschuwingsadresList) GetResultsOk() (*[]Waarschuwingsadres, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedWaarschuwingsadresList) SetResults(v []Waarschuwingsadres)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedWaarschuwingsadresList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


