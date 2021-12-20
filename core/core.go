package core

// Operator defines the valid actions one is allowed to do on an endpoint
type Operator string

//Crud Operatios
const (
	OperatorCreate Operator = "CREATE"
	OperatorRead   Operator = "READ"
	OperatorUpdate Operator = "UPDATE"
	OperatorDelete Operator = "DELETE"
)

// CrudEndpoint defines a single endpoint of a crud API
type CrudEndpoint struct {
	Package   string     `json:"package,omitempty"`
	Imports   []string   `json:"imports,omitempty"`
	DataType  string     `json:"data_type,omitempty"`
	BasePath  string     `json:"base_path,omitempty"`
	Name      string     `json:"name,omitempty"`
	UpperName string     `json:"upper_name,omitempty"`
	Operators []Operator `json:"operators,omitempty"`
}

// Error definition for errors
type Error struct {
	Name       string   `json:"name,omitempty"`
	Short      string   `json:"short,omitempty"`
	Message    string   `json:"message,omitempty"`
	Parameters []string `json:"parameters,omitempty"`
}

// Errors data for errors
type Errors struct {
	Package string  `json:"package,omitempty"`
	Errors  []Error `json:"errors,omitempty"`
}

// CrudConfig is the top-level configuration for a crud API
type CrudConfig struct {
	Framework string         `json:"framework,omitempty"`
	Package   string         `json:"package,omitempty"`
	Endpoints []CrudEndpoint `json:"endpoints,omitempty"`
	Errors    *Errors        `json:"errors,omitempty"`
}

// Meta holds additional information about the invoked command
type Meta struct {
	CLIInput string
}
