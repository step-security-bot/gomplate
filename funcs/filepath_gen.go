// Code generated by gencel. DO NOT EDIT.

package funcs

import "github.com/google/cel-go/cel"
import "github.com/google/cel-go/common/types"
import "github.com/google/cel-go/common/types/ref"

var filepathBaseGen = cel.Function("filepath.Base",
	cel.Overload("filepath.Base_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.StringType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x FilePathFuncs

			return types.DefaultTypeAdapter.NativeToValue(x.Base(args[0]))

		}),
	),
)

var filepathCleanGen = cel.Function("filepath.Clean",
	cel.Overload("filepath.Clean_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.StringType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x FilePathFuncs

			return types.DefaultTypeAdapter.NativeToValue(x.Clean(args[0]))

		}),
	),
)

var filepathDirGen = cel.Function("filepath.Dir",
	cel.Overload("filepath.Dir_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.StringType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x FilePathFuncs

			return types.DefaultTypeAdapter.NativeToValue(x.Dir(args[0]))

		}),
	),
)

var filepathExtGen = cel.Function("filepath.Ext",
	cel.Overload("filepath.Ext_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.StringType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x FilePathFuncs

			return types.DefaultTypeAdapter.NativeToValue(x.Ext(args[0]))

		}),
	),
)

var filepathFromSlashGen = cel.Function("filepath.FromSlash",
	cel.Overload("filepath.FromSlash_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.StringType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x FilePathFuncs

			return types.DefaultTypeAdapter.NativeToValue(x.FromSlash(args[0]))

		}),
	),
)

var filepathIsAbsGen = cel.Function("filepath.IsAbs",
	cel.Overload("filepath.IsAbs_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.BoolType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x FilePathFuncs

			return types.DefaultTypeAdapter.NativeToValue(x.IsAbs(args[0]))

		}),
	),
)

var filepathJoinGen = cel.Function("filepath.Join",
	cel.Overload("filepath.Join_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.StringType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x FilePathFuncs
			list, err := sliceToNative[interface{}](args[0].(ref.Val))
			if err != nil {
				return types.WrapErr(err)
			}


			return types.DefaultTypeAdapter.NativeToValue(x.Join(list...))

		}),
	),
)

var filepathMatchGen = cel.Function("filepath.Match",
	cel.Overload("filepath.Match_interface{}_interface{}",

		[]*cel.Type{
			cel.DynType, cel.DynType,
		},
		cel.DynType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x FilePathFuncs

			a0, err := x.Match(args[0], args[1])
			if err != nil {
				return types.WrapErr(err)
			}

			return types.DefaultTypeAdapter.NativeToValue(a0)
		}),
	),
)

var filepathRelGen = cel.Function("filepath.Rel",
	cel.Overload("filepath.Rel_interface{}_interface{}",

		[]*cel.Type{
			cel.DynType, cel.DynType,
		},
		cel.DynType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x FilePathFuncs

			a0, err := x.Rel(args[0], args[1])
			if err != nil {
				return types.WrapErr(err)
			}
			return types.DefaultTypeAdapter.NativeToValue(a0)
		}),
	),
)

var filepathSplitGen = cel.Function("filepath.Split",
	cel.Overload("filepath.Split_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.DynType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x FilePathFuncs

			return types.DefaultTypeAdapter.NativeToValue(x.Split(args[0]))

		}),
	),
)

var filepathToSlashGen = cel.Function("filepath.ToSlash",
	cel.Overload("filepath.ToSlash_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.StringType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x FilePathFuncs

			return types.DefaultTypeAdapter.NativeToValue(x.ToSlash(args[0]))

		}),
	),
)

var filepathVolumeNameGen = cel.Function("filepath.VolumeName",
	cel.Overload("filepath.VolumeName_interface{}",

		[]*cel.Type{
			cel.DynType,
		},
		cel.StringType,
		cel.FunctionBinding(func(args ...ref.Val) ref.Val {

			var x FilePathFuncs

			return types.DefaultTypeAdapter.NativeToValue(x.VolumeName(args[0]))

		}),
	),
)
