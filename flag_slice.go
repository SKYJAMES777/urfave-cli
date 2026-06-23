package cli

import (
	"encoding/csv"
	"reflect"
	"strconv"
	"strings"
)

// SliceFlag is a flag with a slice value.
type SliceFlag struct {
	Name      string
	Aliases   []string
	Usage     string
	Required  bool
	Hidden    bool
	Value     interface{}
	DefaultText string
	EnvVars   []string
	FilePath  string
	Category  string
	Action    func(ctx *Context, value interface{}) error
}

// Apply applies the flag to the flag set.
func (f *SliceFlag) Apply(set *flag.FlagSet) {
	if f.EnvVars != nil {
		for _, envVar := range f.EnvVars {
			if val, ok := lookupEnv(envVar); ok {
				f.Value = parseSliceFromEnv(f.Value, val)
				break
			}
		}
	}

	if f.Value == nil {
		f.Value = f.getDefault()
	}

	switch v := f.Value.(type) {
	case StringSlice:
		f.applyStringSlice(set, v)
	case IntSlice:
		f.applyIntSlice(set, v)
	case Float64Slice:
		f.applyFloat64Slice(set, v)
	case BoolSlice:
		f.applyBoolSlice(set, v)
	default:
		panic("unsupported slice type")
	}
}

func (f *SliceFlag) getDefault() interface{} {
	switch f.Name {
	case "stringSlice":
		return StringSlice{}
	case "intSlice":
		return IntSlice{}
	case "float64Slice":
		return Float64Slice{}
	case "boolSlice":
		return BoolSlice{}
	default:
		return nil
	}
}

func parseSliceFromEnv(value interface{}, envVal string) interface{} {
	switch v := value.(type) {
	case StringSlice:
		return parseStringSliceFromEnv(v, envVal)
	case IntSlice:
		return parseIntSliceFromEnv(v, envVal)
	case Float64Slice:
		return parseFloat64SliceFromEnv(v, envVal)
	case BoolSlice:
		return parseBoolSliceFromEnv(v, envVal)
	default:
		return value
	}
}

func parseStringSliceFromEnv(existing StringSlice, envVal string) StringSlice {
	parts := splitEnvVar(envVal)
	for _, p := range parts {
		existing.Set(p)
	}
	return existing
}

func parseIntSliceFromEnv(existing IntSlice, envVal string) IntSlice {
	parts := splitEnvVar(envVal)
	for _, p := range parts {
		i, err := strconv.Atoi(p)
		if err == nil {
			existing.Set(i)
		}
	}
	return existing
}

func parseFloat64SliceFromEnv(existing Float64Slice, envVal string) Float64Slice {
	parts := splitEnvVar(envVal)
	for _, p := range parts {
		f, err := strconv.ParseFloat(p, 64)
		if err == nil {
			existing.Set(f)
		}
	}
	return existing
}

func parseBoolSliceFromEnv(existing BoolSlice, envVal string) BoolSlice {
	parts := splitEnvVar(envVal)
	for _, p := range parts {
		b, err := strconv.ParseBool(p)
		if err == nil {
			existing.Set(b)
		}
	}
	return existing
}

func splitEnvVar(envVal string) []string {
	r := csv.NewReader(strings.NewReader(envVal))
	records, err := r.Read()
	if err != nil {
		return []string{envVal}
	}
	return records
}

func (f *SliceFlag) applyStringSlice(set *flag.FlagSet, val StringSlice) {
	f.Value = &val
	set.Var(&val, f.Name, f.Usage)
}

func (f *SliceFlag) applyIntSlice(set *flag.FlagSet, val IntSlice) {
	f.Value = &val
	set.Var(&val, f.Name, f.Usage)
}

func (f *SliceFlag) applyFloat64Slice(set *flag.FlagSet, val Float64Slice) {
	f.Value = &val
	set.Var(&val, f.Name, f.Usage)
}

func (f *SliceFlag) applyBoolSlice(set *flag.FlagSet, val BoolSlice) {
	f.Value = &val
	set.Var(&val, f.Name, f.Usage)
}

// lookupEnv is a wrapper for os.LookupEnv to allow testing.
var lookupEnv = func(key string) (string, bool) {
	return os.LookupEnv(key)
}

// SetLookupEnv allows overriding the environment lookup function for testing.
func SetLookupEnv(fn func(string) (string, bool)) {
	lookupEnv = fn
}
