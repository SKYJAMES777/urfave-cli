package cli

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// SliceFlag is a flag with a slice value.
type SliceFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	EnvVars     []string
	FilePath    string
	Required    bool
	Hidden      bool
	HasBeenSet  bool
	Value       Slice
	DefaultText string
	Destination *Slice
	TakesFile   bool
}

// IsSet returns whether the flag has been set.
func (f *SliceFlag) IsSet() bool {
	return f.HasBeenSet
}

// String returns a readable representation of the flag.
func (f *SliceFlag) String() string {
	return FlagStringer(f)
}

// Names returns the names of the flag.
func (f *SliceFlag) Names() []string {
	return flagNames(f.Name, f.Aliases)
}

// IsRequired returns whether the flag is required.
func (f *SliceFlag) IsRequired() bool {
	return f.Required
}

// IsVisible returns whether the flag is visible.
func (f *SliceFlag) IsVisible() bool {
	return !f.Hidden
}

// GetDefaultText returns the default text for the flag.
func (f *SliceFlag) GetDefaultText() string {
	if f.DefaultText != "" {
		return f.DefaultText
	}
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}

// GetUsage returns the usage string for the flag.
func (f *SliceFlag) GetUsage() string {
	return f.Usage
}

// GetValue returns the flags value as string representation and an empty string if the flag takes no value.
func (f *SliceFlag) GetValue() string {
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}

// RunAction executes the action for the flag.
func (f *SliceFlag) RunAction(c *Context) error {
	return nil
}

// Apply applies the flag to the flag set.
func (f *SliceFlag) Apply(set *flag.FlagSet) error {
	if f.Destination == nil {
		f.Destination = new(Slice)
	}

	// Initialize destination with default value if present
	if f.Value != nil {
		*f.Destination = *f.Value
	}

	// Check environment variables first (they take precedence over defaults)
	envVal := f.lookupEnvVars()
	if envVal != "" {
		// Parse environment variable value and set destination
		parsed := f.parseEnvValue(envVal)
		if parsed != nil {
			*f.Destination = *parsed
			f.HasBeenSet = true
		}
	}

	// Register the flag with the flag set
	set.Var(f.Destination, f.Name, f.Usage)
	for _, alias := range f.Aliases {
		set.Var(f.Destination, alias, f.Usage)
	}

	return nil
}

// lookupEnvVars checks environment variables in order and returns the first non-empty value.
func (f *SliceFlag) lookupEnvVars() string {
	for _, envVar := range f.EnvVars {
		if val, ok := os.LookupEnv(envVar); ok && val != "" {
			return val
		}
	}
	return ""
}

// parseEnvValue parses an environment variable string into a Slice.
func (f *SliceFlag) parseEnvValue(val string) *Slice {
	// Use CSV reader to handle quoted strings and commas
	r := csv.NewReader(strings.NewReader(val))
	records, err := r.Read()
	if err != nil {
		return nil
	}
	result := &Slice{}
	for _, record := range records {
		result.Set(record)
	}
	return result
}

// Slice is a flag.Value that holds a slice of strings.
type Slice []string

// String returns a string representation of the slice.
func (s *Slice) String() string {
	return strings.Join(*s, ",")
}

// Set appends a value to the slice.
func (s *Slice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

// Get returns the slice as an interface{}.
func (s *Slice) Get() interface{} {
	return []string(*s)
}

// StringSliceFlag is a flag with a string slice value.
type StringSliceFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	EnvVars     []string
	FilePath    string
	Required    bool
	Hidden      bool
	HasBeenSet  bool
	Value       *StringSlice
	DefaultText string
	Destination *StringSlice
	TakesFile   bool
}

// IsSet returns whether the flag has been set.
func (f *StringSliceFlag) IsSet() bool {
	return f.HasBeenSet
}

// String returns a readable representation of the flag.
func (f *StringSliceFlag) String() string {
	return FlagStringer(f)
}

// Names returns the names of the flag.
func (f *StringSliceFlag) Names() []string {
	return flagNames(f.Name, f.Aliases)
}

// IsRequired returns whether the flag is required.
func (f *StringSliceFlag) IsRequired() bool {
	return f.Required
}

// IsVisible returns whether the flag is visible.
func (f *StringSliceFlag) IsVisible() bool {
	return !f.Hidden
}

// GetDefaultText returns the default text for the flag.
func (f *StringSliceFlag) GetDefaultText() string {
	if f.DefaultText != "" {
		return f.DefaultText
	}
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}

// GetUsage returns the usage string for the flag.
func (f *StringSliceFlag) GetUsage() string {
	return f.Usage
}

// GetValue returns the flags value as string representation and an empty string if the flag takes no value.
func (f *StringSliceFlag) GetValue() string {
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}

// RunAction executes the action for the flag.
func (f *StringSliceFlag) RunAction(c *Context) error {
	return nil
}

// Apply applies the flag to the flag set.
func (f *StringSliceFlag) Apply(set *flag.FlagSet) error {
	if f.Destination == nil {
		f.Destination = new(StringSlice)
	}

	// Initialize destination with default value if present
	if f.Value != nil {
		*f.Destination = *f.Value
	}

	// Check environment variables first (they take precedence over defaults)
	envVal := f.lookupEnvVars()
	if envVal != "" {
		// Parse environment variable value and set destination
		parsed := f.parseEnvValue(envVal)
		if parsed != nil {
			*f.Destination = *parsed
			f.HasBeenSet = true
		}
	}

	// Register the flag with the flag set
	set.Var(f.Destination, f.Name, f.Usage)
	for _, alias := range f.Aliases {
		set.Var(f.Destination, alias, f.Usage)
	}

	return nil
}

// lookupEnvVars checks environment variables in order and returns the first non-empty value.
func (f *StringSliceFlag) lookupEnvVars() string {
	for _, envVar := range f.EnvVars {
		if val, ok := os.LookupEnv(envVar); ok && val != "" {
			return val
		}
	}
	return ""
}

// parseEnvValue parses an environment variable string into a StringSlice.
func (f *StringSliceFlag) parseEnvValue(val string) *StringSlice {
	// Use CSV reader to handle quoted strings and commas
	r := csv.NewReader(strings.NewReader(val))
	records, err := r.Read()
	if err != nil {
		return nil
	}
	result := &StringSlice{}
	for _, record := range records {
		result.Set(record)
	}
	return result
}

// StringSlice is a flag.Value that holds a slice of strings.
type StringSlice []string

// String returns a string representation of the slice.
func (s *StringSlice) String() string {
	return strings.Join(*s, ",")
}

// Set appends a value to the slice.
func (s *StringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

// Get returns the slice as an interface{}.
func (s *StringSlice) Get() interface{} {
	return []string(*s)
}

// IntSliceFlag is a flag with an int slice value.
type IntSliceFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	EnvVars     []string
	FilePath    string
	Required    bool
	Hidden      bool
	HasBeenSet  bool
	Value       *IntSlice
	DefaultText string
	Destination *IntSlice
	TakesFile   bool
}

// IsSet returns whether the flag has been set.
func (f *IntSliceFlag) IsSet() bool {
	return f.HasBeenSet
}

// String returns a readable representation of the flag.
func (f *IntSliceFlag) String() string {
	return FlagStringer(f)
}

// Names returns the names of the flag.
func (f *IntSliceFlag) Names() []string {
	return flagNames(f.Name, f.Aliases)
}

// IsRequired returns whether the flag is required.
func (f *IntSliceFlag) IsRequired() bool {
	return f.Required
}

// IsVisible returns whether the flag is visible.
func (f *IntSliceFlag) IsVisible() bool {
	return !f.Hidden
}

// GetDefaultText returns the default text for the flag.
func (f *IntSliceFlag) GetDefaultText() string {
	if f.DefaultText != "" {
		return f.DefaultText
	}
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}

// GetUsage returns the usage string for the flag.
func (f *IntSliceFlag) GetUsage() string {
	return f.Usage
}

// GetValue returns the flags value as string representation and an empty string if the flag takes no value.
func (f *IntSliceFlag) GetValue() string {
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}

// RunAction executes the action for the flag.
func (f *IntSliceFlag) RunAction(c *Context) error {
	return nil
}

// Apply applies the flag to the flag set.
func (f *IntSliceFlag) Apply(set *flag.FlagSet) error {
	if f.Destination == nil {
		f.Destination = new(IntSlice)
	}

	// Initialize destination with default value if present
	if f.Value != nil {
		*f.Destination = *f.Value
	}

	// Check environment variables first (they take precedence over defaults)
	envVal := f.lookupEnvVars()
	if envVal != "" {
		// Parse environment variable value and set destination
		parsed := f.parseEnvValue(envVal)
		if parsed != nil {
			*f.Destination = *parsed
			f.HasBeenSet = true
		}
	}

	// Register the flag with the flag set
	set.Var(f.Destination, f.Name, f.Usage)
	for _, alias := range f.Aliases {
		set.Var(f.Destination, alias, f.Usage)
	}

	return nil
}

// lookupEnvVars checks environment variables in order and returns the first non-empty value.
func (f *IntSliceFlag) lookupEnvVars() string {
	for _, envVar := range f.EnvVars {
		if val, ok := os.LookupEnv(envVar); ok && val != "" {
			return val
		}
	}
	return ""
}

// parseEnvValue parses an environment variable string into an IntSlice.
func (f *IntSliceFlag) parseEnvValue(val string) *IntSlice {
	// Use CSV reader to handle quoted strings and commas
	r := csv.NewReader(strings.NewReader(val))
	records, err := r.Read()
	if err != nil {
		return nil
	}
	result := &IntSlice{}
	for _, record := range records {
		result.Set(record)
	}
	return result
}

// IntSlice is a flag.Value that holds a slice of ints.
type IntSlice []int

// String returns a string representation of the slice.
func (s *IntSlice) String() string {
	parts := make([]string, len(*s))
	for i, v := range *s {
		parts[i] = strconv.Itoa(v)
	}
	return strings.Join(parts, ",")
}

// Set appends a value to the slice.
func (s *IntSlice) Set(value string) error {
	v, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	*s = append(*s, v)
	return nil
}

// Get returns the slice as an interface{}.
func (s *IntSlice) Get() interface{} {
	return []int(*s)
}

// Float64SliceFlag is a flag with a float64 slice value.
type Float64SliceFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	EnvVars     []string
	FilePath    string
	Required    bool
	Hidden      bool
	HasBeenSet  bool
	Value       *Float64Slice
	DefaultText string
	Destination *Float64Slice
	TakesFile   bool
}

// IsSet returns whether the flag has been set.
func (f *Float64SliceFlag) IsSet() bool {
	return f.HasBeenSet
}

// String returns a readable representation of the flag.
func (f *Float64SliceFlag) String() string {
	return FlagStringer(f)
}

// Names returns the names of the flag.
func (f *Float64SliceFlag) Names() []string {
	return flagNames(f.Name, f.Aliases)
}

// IsRequired returns whether the flag is required.
func (f *Float64SliceFlag) IsRequired() bool {
	return f.Required
}

// IsVisible returns whether the flag is visible.
func (f *Float64SliceFlag) IsVisible() bool {
	return !f.Hidden
}

// GetDefaultText returns the default text for the flag.
func (f *Float64SliceFlag) GetDefaultText() string {
	if f.DefaultText != "" {
		return f.DefaultText
	}
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}

// GetUsage returns the usage string for the flag.
func (f *Float64SliceFlag) GetUsage() string {
	return f.Usage
}

// GetValue returns the flags value as string representation and an empty string if the flag takes no value.
func (f *Float64SliceFlag) GetValue() string {
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}

// RunAction executes the action for the flag.
func (f *Float64SliceFlag) RunAction(c *Context) error {
	return nil
}

// Apply applies the flag to the flag set.
func (f *Float64SliceFlag) Apply(set *flag.FlagSet) error {
	if f.Destination == nil {
		f.Destination = new(Float64Slice)
	}

	// Initialize destination with default value if present
	if f.Value != nil {
		*f.Destination = *f.Value
	}

	// Check environment variables first (they take precedence over defaults)
	envVal := f.lookupEnvVars()
	if envVal != "" {
		// Parse environment variable value and set destination
		parsed := f.parseEnvValue(envVal)
		if parsed != nil {
			*f.Destination = *parsed
			f.HasBeenSet = true
		}
	}

	// Register the flag with the flag set
	set.Var(f.Destination, f.Name, f.Usage)
	for _, alias := range f.Aliases {
		set.Var(f.Destination, alias, f.Usage)
	}

	return nil
}

// lookupEnvVars checks environment variables in order and returns the first non-empty value.
func (f *Float64SliceFlag) lookupEnvVars() string {
	for _, envVar := range f.EnvVars {
		if val, ok := os.LookupEnv(envVar); ok && val != "" {
			return val
		}
	}
	return ""
}

// parseEnvValue parses an environment variable string into a Float64Slice.
func (f *Float64SliceFlag) parseEnvValue(val string) *Float64Slice {
	// Use CSV reader to handle quoted strings and commas
	r := csv.NewReader(strings.NewReader(val))
	records, err := r.Read()
	if err != nil {
		return nil
	}
	result := &Float64Slice{}
	for _, record := range records {
		result.Set(record)
	}
	return result
}

// Float64Slice is a flag.Value that holds a slice of float64s.
type Float64Slice []float64

// String returns a string representation of the slice.
func (s *Float64Slice) String() string {
	parts := make([]string, len(*s))
	for i, v := range *s {
		parts[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return strings.Join(parts, ",")
}

// Set appends a value to the slice.
func (s *Float64Slice) Set(value string) error {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	*s = append(*s, v)
	return nil
}

// Get returns the slice as an interface{}.
func (s *Float64Slice) Get() interface{} {
	return []float64(*s)
}

// GenericSliceFlag is a flag with a generic slice value.
type GenericSliceFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	EnvVars     []string
	FilePath    string
	Required    bool
	Hidden      bool
	HasBeenSet  bool
	Value       Generic
	DefaultText string
	Destination *Generic
	TakesFile   bool
}

// IsSet returns whether the flag has been set.
func (f *GenericSliceFlag) IsSet() bool {
	return f.HasBeenSet
}

// String returns a readable representation of the flag.
func (f *GenericSliceFlag) String() string {
	return FlagStringer(f)
}

// Names returns the names of the flag.
func (f *GenericSliceFlag) Names() []string {
	return flagNames(f.Name, f.Aliases)
}

// IsRequired returns whether the flag is required.
func (f *GenericSliceFlag) IsRequired() bool {
	return f.Required
}

// IsVisible returns whether the flag is visible.
func (f *GenericSliceFlag) IsVisible() bool {
	return !f.Hidden
}

// GetDefaultText returns the default text for the flag.
func (f *GenericSliceFlag) GetDefaultText() string {
	if f.DefaultText != "" {
		return f.DefaultText
	}
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}

// GetUsage returns the usage string for the flag.
func (f *GenericSliceFlag) GetUsage() string {
	return f.Usage
}

// GetValue returns the flags value as string representation and an empty string if the flag takes no value.
func (f *GenericSliceFlag) GetValue() string {
	if f.Value != nil {
		return f.Value.String()
	}
	return ""
}

// RunAction executes the action for the flag.
func (f *GenericSliceFlag) RunAction(c *Context) error {
	return nil
}

// Apply applies the flag to the flag set.
func (f *GenericSliceFlag) Apply(set *flag.FlagSet) error {
	if f.Destination == nil {
		f.Destination = new(Generic)
	}

	// Initialize destination with default value if present
	if f.Value != nil {
		*f.Destination = f.Value
	}

	// Check environment variables first (they take precedence over defaults)
	envVal := f.lookupEnvVars()
	if envVal != "" {
		// Parse environment variable value and set destination
		parsed := f.parseEnvValue(envVal)
		if parsed != nil {
			*f.Destination = *parsed
			f.HasBeenSet = true
		}
	}

	// Register the flag with the flag set
	set.Var(f.Destination, f.Name, f.Usage)
	for _, alias := range f.Aliases {
		set.Var(f.Destination, alias, f.Usage)
	}

	return nil
}

// lookupEnvVars checks environment variables in order and returns the first non-empty value.
func (f *GenericSliceFlag) lookupEnvVars() string {
	for _, envVar := range f.EnvVars {
		if val, ok := os.LookupEnv(envVar); ok && val != "" {
			return val
		}
	}
	return ""
}

// parseEnvValue parses an environment variable string into a Generic.
func (f *GenericSliceFlag) parseEnvValue(val string) *Generic {
	// Use CSV reader to handle quoted strings and commas
	r := csv.NewReader(strings.NewReader(val))
	records, err := r.Read()
	if err != nil {
		return nil
	}
	result := &Generic{}
	for _, record := range records {
		result.Set(record)
	}
	return result
}

// Generic is a flag.Value that holds a generic value.
type Generic struct {
	value interface{}
}

// String returns a string representation of the generic value.
func (g *Generic) String() string {
	return fmt.Sprintf("%v", g.value)
}

// Set sets the generic value.
func (g *Generic) Set(value string) error {
	g.value = value
	return nil
}

// Get returns the generic value.
func (g *Generic) Get() interface{} {
	return g.value
}
