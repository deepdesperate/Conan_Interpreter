package evaluator

import (
	"github.com/deepdesperate/Conan_Interpreter/object"
)

var builtins = map[string] *object.Builtin{
	"len": object.GetBuiltinsByName("len"),

	"first": object.GetBuiltinsByName("first"),

	"last": object.GetBuiltinsByName("last"),

	"rest": object.GetBuiltinsByName("rest"),

	"push": object.GetBuiltinsByName("push"),

	"puts": object.GetBuiltinsByName("puts"),
}