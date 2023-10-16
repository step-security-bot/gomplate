// Code generated by gencel. DO NOT EDIT.

package funcs

import "github.com/google/cel-go/cel"
import "github.com/google/cel-go/common/types"
import "github.com/google/cel-go/common/types/ref"
import "reflect"
import "sort"

var collSliceGen = cel.Function("Slice",
	cel.Overload("Slice_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.DynType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x CollFuncs
			list := transferSlice[interface{}](args[0].(ref.Val))

			return types.DefaultTypeAdapter.NativeToValue(x.Slice(list...))

		}),
	),
)

var collHasGen = cel.Function("Has",
	cel.Overload("Has_interface{}_string",

		[]*cel.Type{
			cel.DynType, cel.StringType,
		},
		cel.BoolType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x CollFuncs

			return types.DefaultTypeAdapter.NativeToValue(x.Has(args[0], args[1].Value().(string)))

		}),
	),
)

var collDictGen = cel.Function("Dict",
	cel.Overload("Dict_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.DynType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x CollFuncs
			list := transferSlice[interface{}](args[0].(ref.Val))

			result, err := x.Dict(list...)
			if err != nil {
				return types.NewErr(err.Error())
			}
			return types.DefaultTypeAdapter.NativeToValue(result)
		}),
	),
)

var collKeysGen = cel.Function("keys",
	cel.MemberOverload("map_keys",

		[]*cel.Type{
			cel.MapType(cel.StringType, cel.AnyType),
		},
		cel.ListType(cel.StringType),
		cel.UnaryBinding(func(arg ref.Val) ref.Val {
			var x CollFuncs
			_map, err := convertMap[string,any](arg)
			if err != nil {
				return types.NewErr(err.Error())
			}
			result := x.Keys(_map)
			sort.Strings(result)
			return types.DefaultTypeAdapter.NativeToValue(result)
		}),
	),
)

var collValuesGen = cel.Function("values",
	cel.MemberOverload("map_values",

	[]*cel.Type{
		cel.MapType(cel.StringType, cel.AnyType),
	},
	cel.ListType(cel.AnyType),
		cel.UnaryBinding(func(arg ref.Val) ref.Val {

			var x CollFuncs

			_map, err := convertMap[string,any](arg)
			if err != nil {
				return types.NewErr(err.Error())
			}

			result := x.Values(_map)
			return types.DefaultTypeAdapter.NativeToValue(result)
		}),
	),
)

var collAppendGen = cel.Function("Append",
	cel.Overload("Append_interface{}_interface{}",

		[]*cel.Type{
			cel.DynType, cel.DynType,
		},
		cel.DynType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x CollFuncs

			result, err := x.Append(args[0], args[1])
			if err != nil {
				return types.NewErr(err.Error())
			}
			return types.DefaultTypeAdapter.NativeToValue(result)
		}),
	),
)

var collPrependGen = cel.Function("Prepend",
	cel.Overload("Prepend_interface{}_interface{}",

		[]*cel.Type{
			cel.DynType, cel.DynType,
		},
		cel.DynType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x CollFuncs

			result, err := x.Prepend(args[0], args[1])
			if err != nil {
				return types.NewErr(err.Error())
			}
			return types.DefaultTypeAdapter.NativeToValue(result)

		}),
	),
)

var collUniqGen = cel.Function("uniq",
	cel.Overload("Uniq_interface{}",

		[]*cel.Type{
			cel.ListType(cel.StringType),
		},
		cel.ListType(cel.StringType),
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x CollFuncs

			list, _ := args[0].ConvertToNative(reflect.TypeOf([]string{}))

			result, err := x.Uniq(list)
			if err != nil {
				return types.NewErr(err.Error())
			}
			return types.DefaultTypeAdapter.NativeToValue(result)

		}),
	),
)

var collReverseGen = cel.Function("Reverse",
	cel.Overload("Reverse_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.DynType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x CollFuncs

			result, err := x.Reverse(args[0])
			if err != nil {
				return types.NewErr(err.Error())
			}
			return types.DefaultTypeAdapter.NativeToValue(result)

		}),
	),
)

var collMergeGen = cel.Function("Merge",
	cel.Overload("Merge_map[string]interface{}_map[string]interface{}",

		[]*cel.Type{
			cel.DynType, cel.DynType,
		},
		cel.DynType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x CollFuncs
			list := transferSlice[map[string]interface{}](args[1].(ref.Val))

			result, err := x.Merge(args[0].Value().(map[string]interface{}), list...)
			if err != nil {
				return types.NewErr(err.Error())
			}
			return types.DefaultTypeAdapter.NativeToValue(result)

		}),
	),
)

var collSortGen = cel.Function("Sort",
	cel.Overload("Sort_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.DynType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x CollFuncs
			list := transferSlice[interface{}](args[0].(ref.Val))

			result, err := x.Sort(list...)
			if err != nil {
				return types.NewErr(err.Error())
			}
			return types.DefaultTypeAdapter.NativeToValue(result)

		}),
	),
)

var collJQGen = cel.Function("jq",
	cel.Overload("jq_string_interface{}",

		[]*cel.Type{
			cel.StringType, cel.DynType,
		},
		cel.DynType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x CollFuncs

			result, err := x.JQ(args[0].Value().(string), args[1].Value())
			if err != nil {
				return types.NewErr(err.Error())
			}
			return types.DefaultTypeAdapter.NativeToValue(result)

		}),
	),
)

var collFlattenGen = cel.Function("Flatten",
	cel.Overload("Flatten_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.DynType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x CollFuncs
			list := transferSlice[interface{}](args[0].(ref.Val))

			result, err := x.Flatten(list...)
			return types.DefaultTypeAdapter.NativeToValue([]any{
				result, err,
			})

		}),
	),
)

var collPickGen = cel.Function("Pick",
	cel.Overload("Pick_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.DynType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x CollFuncs
			list := transferSlice[interface{}](args[0].(ref.Val))

			result, err := x.Pick(list...)
			if err != nil {
				return types.NewErr(err.Error())
			}
			return types.DefaultTypeAdapter.NativeToValue(result)

		}),
	),
)

var collOmitGen = cel.Function("Omit",
	cel.Overload("Omit_interface{}",
		[]*cel.Type{
			cel.DynType,
		},
		cel.DynType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x CollFuncs
			list := transferSlice[interface{}](args[0].(ref.Val))

			result, err := x.Omit(list...)
			if err != nil {
				return types.NewErr(err.Error())
			}
			return types.DefaultTypeAdapter.NativeToValue(result)
		}),
	),
)
