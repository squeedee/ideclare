package job

// First quandry; How do we verify input digests without making that the job's responsibility.
// In my mind, this is a vital lift. Job author says "I need this input, and I can trust that the system will capture its signature"
// Maybe the difference between "Env Vars" and these "digestable inputs" is that they are "Externals".

const Digestable_Config = "config"
const Digestable_Input = "input"

type EnvVar struct {
	Name   string `json:"name"`
	Desc   string `json:"description"`
	Type   string `json:"type"` // Validate enum, default to config
	Digest string `json:"digest"`
}

type Env []EnvVar

type Spec struct {
	// Env records all the environment variables
	Env Env `json:"env"`
}
