package api

// First quandry; How do we verify input digests without making that the job's responsibility.
// In my mind, this is a vital lift. Job author says "I need this input, and I can trust that the system will capture its signature"
// Maybe the difference between "Env Vars" and these "digestable inputs" is that they are "Externals".

const Digestable_Value = "Value"
const Digestable_OciRegistry = "OCIRegistry"
const Digestable_UrlBody = "URLBody"

type Mapping struct {
}

type Input struct {
	// Name the input. This is also used as the default mapping to Environment variables, see Mapping.
	Name string `json:"name"`

	// Description is a free text definition of the input.
	Description string `json:"description"`

	// Type is a free text definition
	Type string `json:"type"`

	// Optional is true if this value does not need to be provided
	Optional bool `json:"optional"`

	// Default value that can be either string or number
	Default Value `json:"default"`

	// Mapping defines how the input is passed into the container
	Mapping Mapping `json:"mapping"`
}

type Inputs []Input

type Spec struct {
	Inputs Inputs `json:"env"`
}
