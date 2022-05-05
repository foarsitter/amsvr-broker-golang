# ValidationObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | [**NullableObjectTypeEnum**](ObjectTypeEnum.md) |  | [readonly] 
**ObjectId** | **int32** |  | [readonly] 
**Results** | [**[]ValidationRule**](ValidationRule.md) |  | 

## Methods

### NewValidationObject

`func NewValidationObject(objectType NullableObjectTypeEnum, objectId int32, results []ValidationRule, ) *ValidationObject`

NewValidationObject instantiates a new ValidationObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationObjectWithDefaults

`func NewValidationObjectWithDefaults() *ValidationObject`

NewValidationObjectWithDefaults instantiates a new ValidationObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *ValidationObject) GetObjectType() ObjectTypeEnum`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ValidationObject) GetObjectTypeOk() (*ObjectTypeEnum, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ValidationObject) SetObjectType(v ObjectTypeEnum)`

SetObjectType sets ObjectType field to given value.


### SetObjectTypeNil

`func (o *ValidationObject) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *ValidationObject) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetObjectId

`func (o *ValidationObject) GetObjectId() int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ValidationObject) GetObjectIdOk() (*int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ValidationObject) SetObjectId(v int32)`

SetObjectId sets ObjectId field to given value.


### GetResults

`func (o *ValidationObject) GetResults() []ValidationRule`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ValidationObject) GetResultsOk() (*[]ValidationRule, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ValidationObject) SetResults(v []ValidationRule)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


