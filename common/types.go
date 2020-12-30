package common

// ReverseTarget represents a reverse target
type ReverseTarget struct {
	Type          string   `yaml:"type"`
	IncludeTables []string `yaml:"include_tables"`
	ExcludeTables []string `yaml:"exclude_tables"`
	TableMapper   string   `yaml:"table_mapper"`
	ColumnMapper  string   `yaml:"column_mapper"`
	TemplatePath  string   `yaml:"template_path"`
	Template      string   `yaml:"template"`
	MultipleFiles bool     `yaml:"multiple_files"`
	OutputDir     string   `yaml:"output_dir"`
	TablePrefix   string   `yaml:"table_prefix"`
	Language      string   `yaml:"language"`
	TableName     bool     `yaml:"table_name"`
	JsonTag       bool     `yaml:"json_tag"`
	JsonTagMapper string   `yaml:"json_tag_mapper"`

	Funcs     map[string]string `yaml:"funcs"`
	Formatter string            `yaml:"formatter"`
	Importter string            `yaml:"importter"`
	ExtName   string            `yaml:"ext_name"`
}

// ReverseConfig represents a reverse configuration
type ReverseConfig struct {
	Kind    string          `yaml:"kind"`
	Name    string          `yaml:"name"`
	Source  ReverseSource   `yaml:"source"`
	Targets []ReverseTarget `yaml:"targets"`
}

// ReverseSource represents a reverse source which should be a database connection
type ReverseSource struct {
	Database string `yaml:"database"`
	ConnStr  string `yaml:"conn_str"`
}
