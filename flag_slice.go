package cli

import (
	"encoding/csv"
	"reflect"
	"strconv"
	"strings"
)

// SliceFlag is a flag with a slice value.
type SliceFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	DefaultText string
	Required    bool
	Hidden      bool
	HasBeenSet  bool
	Value       Slice
	Destination *Slice
	EnvVars     []string
	FilePath    string
	Category    string
	Sources     []string
	Config      SliceConfig
}

// SliceConfig holds configuration for slice flags.
type SliceConfig struct {
	Separator string
	TrimSpace bool
}

// Slice is an interface for slice values.
type Slice interface {
	Set(string) error
	String() string
	Get() interface{}
	Append(string) error
	Replace([]string) error
	Clear()
	Len() int
	Slice() []string
}

// StringSlice is a slice of strings.
type StringSlice struct {
	values []string
}

func NewStringSlice(defaults ...string) *StringSlice {
	return &StringSlice{values: defaults}
}

func (s *StringSlice) Set(val string) error {
	s.values = append(s.values, val)
	return nil
}

func (s *StringSlice) String() string {
	return strings.Join(s.values, ",")
}

func (s *StringSlice) Get() interface{} {
	return s.values
}

func (s *StringSlice) Append(val string) error {
	s.values = append(s.values, val)
	return nil
}

func (s *StringSlice) Replace(vals []string) error {
	s.values = vals
	return nil
}

func (s *StringSlice) Clear() {
	s.values = []string{}
}

func (s *StringSlice) Len() int {
	return len(s.values)
}

func (s *StringSlice) Slice() []string {
	return s.values
}

// IntSlice is a slice of ints.
type IntSlice struct {
	values []int
}

func NewIntSlice(defaults ...int) *IntSlice {
	return &IntSlice{values: defaults}
}

func (s *IntSlice) Set(val string) error {
	i, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	s.values = append(s.values, i)
	return nil
}

func (s *IntSlice) String() string {
	strs := make([]string, len(s.values))
	for i, v := range s.values {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, ",")
}

func (s *IntSlice) Get() interface{} {
	return s.values
}

func (s *IntSlice) Append(val string) error {
	i, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	s.values = append(s.values, i)
	return nil
}

func (s *IntSlice) Replace(vals []string) error {
	ints := make([]int, len(vals))
	for i, v := range vals {
		iv, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		ints[i] = iv
	}
	s.values = ints
	return nil
}

func (s *IntSlice) Clear() {
	s.values = []int{}
}

func (s *IntSlice) Len() int {
	return len(s.values)
}

func (s *IntSlice) Slice() []string {
	strs := make([]string, len(s.values))
	for i, v := range s.values {
		strs[i] = strconv.Itoa(v)
	}
	return strs
}

// Float64Slice is a slice of float64s.
type Float64Slice struct {
	values []float64
}

func NewFloat64Slice(defaults ...float64) *Float64Slice {
	return &Float64Slice{values: defaults}
}

func (s *Float64Slice) Set(val string) error {
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}
	s.values = append(s.values, f)
	return nil
}

func (s *Float64Slice) String() string {
	strs := make([]string, len(s.values))
	for i, v := range s.values {
		strs[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return strings.Join(strs, ",")
}

func (s *Float64Slice) Get() interface{} {
	return s.values
}

func (s *Float64Slice) Append(val string) error {
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}
	s.values = append(s.values, f)
	return nil
}

func (s *Float64Slice) Replace(vals []string) error {
	floats := make([]float64, len(vals))
	for i, v := range vals {
		fv, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		floats[i] = fv
	}
	s.values = floats
	return nil
}

func (s *Float64Slice) Clear() {
	s.values = []float64{}
}

func (s *Float64Slice) Len() int {
	return len(s.values)
}

func (s *Float64Slice) Slice() []string {
	strs := make([]string, len(s.values))
	for i, v := range s.values {
		strs[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return strs
}

// BoolSlice is a slice of bools.
type BoolSlice struct {
	values []bool
}

func NewBoolSlice(defaults ...bool) *BoolSlice {
	return &BoolSlice{values: defaults}
}

func (s *BoolSlice) Set(val string) error {
	b, err := strconv.ParseBool(val)
	if err != nil {
		return err
	}
	s.values = append(s.values, b)
	return nil
}

func (s *BoolSlice) String() string {
	strs := make([]string, len(s.values))
	for i, v := range s.values {
		strs[i] = strconv.FormatBool(v)
	}
	return strings.Join(strs, ",")
}

func (s *BoolSlice) Get() interface{} {
	return s.values
}

func (s *BoolSlice) Append(val string) error {
	b, err := strconv.ParseBool(val)
	if err != nil {
		return err
	}
	s.values = append(s.values, b)
	return nil
}

func (s *BoolSlice) Replace(vals []string) error {
	bools := make([]bool, len(vals))
	for i, v := range vals {
		bv, err := strconv.ParseBool(v)
		if err != nil {
			return err
		}
		bools[i] = bv
	}
	s.values = bools
	return nil
}

func (s *BoolSlice) Clear() {
	s.values = []bool{}
}

func (s *BoolSlice) Len() int {
	return len(s.values)
}

func (s *BoolSlice) Slice() []string {
	strs := make([]string, len(s.values))
	for i, v := range s.values {
		strs[i] = strconv.FormatBool(v)
	}
	return strs
}

// parseSliceFromEnv parses a slice value from environment variables.
// It supports splitting by a configurable separator (default comma).
func parseSliceFromEnv(envVars []string, config SliceConfig) ([]string, bool) {
	for _, envVar := range envVars {
		val, ok := lookupEnv(envVar)
		if ok {
			separator := ","
			if config.Separator != "" {
				separator = config.Separator
			}
			parts := strings.Split(val, separator)
			if config.TrimSpace {
				for i, part := range parts {
					parts[i] = strings.TrimSpace(part)
				}
			}
			return parts, true
		}
	}
	return nil, false
}

// lookupEnv is a wrapper for os.LookupEnv to allow testing.
var lookupEnv = func(key string) (string, bool) {
	return os.LookupEnv(key)
}

// applySliceFlagDefaults applies the default value for a slice flag.
// It clears the slice and sets it from the default values.
func applySliceFlagDefaults(f *SliceFlag) {
	if f.Value != nil {
		f.Value.Clear()
		for _, val := range f.Value.Slice() {
			f.Value.Append(val)
		}
	}
}

// applySliceFlagEnv applies environment variable values to a slice flag.
// It overrides any existing values.
func applySliceFlagEnv(f *SliceFlag) {
	if len(f.EnvVars) > 0 {
		vals, ok := parseSliceFromEnv(f.EnvVars, f.Config)
		if ok {
			f.Value.Clear()
			for _, val := range vals {
				f.Value.Append(val)
			}
			f.HasBeenSet = true
		}
	}
}

// applySliceFlagArgs applies command-line arguments to a slice flag.
// It appends to existing values.
func applySliceFlagArgs(f *SliceFlag, args []string) {
	for _, arg := range args {
		f.Value.Append(arg)
	}
	f.HasBeenSet = true
}

// resolveSliceFlag resolves the final value of a slice flag based on precedence:
// 1. Command-line arguments (highest)
// 2. Environment variables
// 3. Default values (lowest)
func resolveSliceFlag(f *SliceFlag, args []string) {
	// Start with defaults
	applySliceFlagDefaults(f)

	// Override with environment variables if present
	applySliceFlagEnv(f)

	// Override with command-line arguments if present
	if len(args) > 0 {
		applySliceFlagArgs(f, args)
	}
}

// StringSliceFlag is a flag with a string slice value.
type StringSliceFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	DefaultText string
	Required    bool
	Hidden      bool
	HasBeenSet  bool
	Value       *StringSlice
	Destination *StringSlice
	EnvVars     []string
	FilePath    string
	Category    string
	Sources     []string
	Config      SliceConfig
}

func (f *StringSliceFlag) Apply(set *flag.FlagSet) error {
	if f.Value == nil {
		f.Value = NewStringSlice()
	}
	if f.Destination != nil {
		f.Destination = f.Value
	}

	// Resolve value based on precedence
	resolveSliceFlag(&SliceFlag{
		Value:   f.Value,
		EnvVars: f.EnvVars,
		Config:  f.Config,
	}, f.Sources)

	return nil
}

// IntSliceFlag is a flag with an int slice value.
type IntSliceFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	DefaultText string
	Required    bool
	Hidden      bool
	HasBeenSet  bool
	Value       *IntSlice
	Destination *IntSlice
	EnvVars     []string
	FilePath    string
	Category    string
	Sources     []string
	Config      SliceConfig
}

func (f *IntSliceFlag) Apply(set *flag.FlagSet) error {
	if f.Value == nil {
		f.Value = NewIntSlice()
	}
	if f.Destination != nil {
		f.Destination = f.Value
	}

	resolveSliceFlag(&SliceFlag{
		Value:   f.Value,
		EnvVars: f.EnvVars,
		Config:  f.Config,
	}, f.Sources)

	return nil
}

// Float64SliceFlag is a flag with a float64 slice value.
type Float64SliceFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	DefaultText string
	Required    bool
	Hidden      bool
	HasBeenSet  bool
	Value       *Float64Slice
	Destination *Float64Slice
	EnvVars     []string
	FilePath    string
	Category    string
	Sources     []string
	Config      SliceConfig
}

func (f *Float64SliceFlag) Apply(set *flag.FlagSet) error {
	if f.Value == nil {
		f.Value = NewFloat64Slice()
	}
	if f.Destination != nil {
		f.Destination = f.Value
	}

	resolveSliceFlag(&SliceFlag{
		Value:   f.Value,
		EnvVars: f.EnvVars,
		Config:  f.Config,
	}, f.Sources)

	return nil
}

// BoolSliceFlag is a flag with a bool slice value.
type BoolSliceFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	DefaultText string
	Required    bool
	Hidden      bool
	HasBeenSet  bool
	Value       *BoolSlice
	Destination *BoolSlice
	EnvVars     []string
	FilePath    string
	Category    string
	Sources     []string
	Config      SliceConfig
}

func (f *BoolSliceFlag) Apply(set *flag.FlagSet) error {
	if f.Value == nil {
		f.Value = NewBoolSlice()
	}
	if f.Destination != nil {
		f.Destination = f.Value
	}

	resolveSliceFlag(&SliceFlag{
		Value:   f.Value,
		EnvVars: f.EnvVars,
		Config:  f.Config,
	}, f.Sources)

	return nil
}
