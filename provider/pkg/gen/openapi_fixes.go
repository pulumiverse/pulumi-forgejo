package gen

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// FixOpenAPIDoc applies patches to the raw OpenAPI spec
// before passing it to pulschema.
func FixOpenAPIDoc(openAPIDoc *openapi3.T) error {
	// fixCreateTeamOptionExample(openAPIDoc)
	// return errors.New("Abandon ship")
	return nil
}

// fixCreateTeamOptionExample fixes the example to be an object rather than a
// string
func fixCreateTeamOptionExample(openAPIDoc *openapi3.T) {
	createTeamOptionProps, ok := openAPIDoc.Components.Schemas["CreateTeamOption"]
	contract.Assertf(ok, "Expected to find CreateTeamOption type")
	unitsMapProp, ok := createTeamOptionProps.Value.Properties["units_map"]
	contract.Assertf(ok, "Expected to find 'units_map' property in CreateTeamOption type")
	exampleProp := unitsMapProp.Value.Example
	unitsMapProp.Value.Example = map[string]*openapi3.SchemaRef{
		"first": openapi3.NewStringSchema().NewRef(),
		"last":  openapi3.NewStringSchema().NewRef(),
		"next":  openapi3.NewStringSchema().NewRef(),
		"prev":  openapi3.NewStringSchema().NewRef(),
	}
	fmt.Printf("Example: %v\n", exampleProp)
	fmt.Printf("Example fixed: %v\n", unitsMapProp.Value.Example)
	fmt.Printf("Type: %v\n", unitsMapProp.Value.Type)
}
