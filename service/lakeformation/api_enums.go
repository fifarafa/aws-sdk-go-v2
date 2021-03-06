// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorEq          ComparisonOperator = "EQ"
	ComparisonOperatorNe          ComparisonOperator = "NE"
	ComparisonOperatorLe          ComparisonOperator = "LE"
	ComparisonOperatorLt          ComparisonOperator = "LT"
	ComparisonOperatorGe          ComparisonOperator = "GE"
	ComparisonOperatorGt          ComparisonOperator = "GT"
	ComparisonOperatorContains    ComparisonOperator = "CONTAINS"
	ComparisonOperatorNotContains ComparisonOperator = "NOT_CONTAINS"
	ComparisonOperatorBeginsWith  ComparisonOperator = "BEGINS_WITH"
	ComparisonOperatorIn          ComparisonOperator = "IN"
	ComparisonOperatorBetween     ComparisonOperator = "BETWEEN"
)

func (enum ComparisonOperator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ComparisonOperator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DataLakeResourceType string

// Enum values for DataLakeResourceType
const (
	DataLakeResourceTypeCatalog      DataLakeResourceType = "CATALOG"
	DataLakeResourceTypeDatabase     DataLakeResourceType = "DATABASE"
	DataLakeResourceTypeTable        DataLakeResourceType = "TABLE"
	DataLakeResourceTypeDataLocation DataLakeResourceType = "DATA_LOCATION"
)

func (enum DataLakeResourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DataLakeResourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FieldNameString string

// Enum values for FieldNameString
const (
	FieldNameStringResourceArn  FieldNameString = "RESOURCE_ARN"
	FieldNameStringRoleArn      FieldNameString = "ROLE_ARN"
	FieldNameStringLastModified FieldNameString = "LAST_MODIFIED"
)

func (enum FieldNameString) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FieldNameString) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Permission string

// Enum values for Permission
const (
	PermissionAll                Permission = "ALL"
	PermissionSelect             Permission = "SELECT"
	PermissionAlter              Permission = "ALTER"
	PermissionDrop               Permission = "DROP"
	PermissionDelete             Permission = "DELETE"
	PermissionInsert             Permission = "INSERT"
	PermissionCreateDatabase     Permission = "CREATE_DATABASE"
	PermissionCreateTable        Permission = "CREATE_TABLE"
	PermissionDataLocationAccess Permission = "DATA_LOCATION_ACCESS"
)

func (enum Permission) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Permission) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
