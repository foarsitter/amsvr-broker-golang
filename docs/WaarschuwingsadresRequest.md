# WaarschuwingsadresRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Telefoonnummer** | **int32** |  | 
**Prioriteit** | Pointer to **int32** | De contactpersoon met de hoogste prioriteit wordt als eerste opgeroepen | [optional] 

## Methods

### NewWaarschuwingsadresRequest

`func NewWaarschuwingsadresRequest(telefoonnummer int32, ) *WaarschuwingsadresRequest`

NewWaarschuwingsadresRequest instantiates a new WaarschuwingsadresRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaarschuwingsadresRequestWithDefaults

`func NewWaarschuwingsadresRequestWithDefaults() *WaarschuwingsadresRequest`

NewWaarschuwingsadresRequestWithDefaults instantiates a new WaarschuwingsadresRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTelefoonnummer

`func (o *WaarschuwingsadresRequest) GetTelefoonnummer() int32`

GetTelefoonnummer returns the Telefoonnummer field if non-nil, zero value otherwise.

### GetTelefoonnummerOk

`func (o *WaarschuwingsadresRequest) GetTelefoonnummerOk() (*int32, bool)`

GetTelefoonnummerOk returns a tuple with the Telefoonnummer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelefoonnummer

`func (o *WaarschuwingsadresRequest) SetTelefoonnummer(v int32)`

SetTelefoonnummer sets Telefoonnummer field to given value.


### GetPrioriteit

`func (o *WaarschuwingsadresRequest) GetPrioriteit() int32`

GetPrioriteit returns the Prioriteit field if non-nil, zero value otherwise.

### GetPrioriteitOk

`func (o *WaarschuwingsadresRequest) GetPrioriteitOk() (*int32, bool)`

GetPrioriteitOk returns a tuple with the Prioriteit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioriteit

`func (o *WaarschuwingsadresRequest) SetPrioriteit(v int32)`

SetPrioriteit sets Prioriteit field to given value.

### HasPrioriteit

`func (o *WaarschuwingsadresRequest) HasPrioriteit() bool`

HasPrioriteit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


