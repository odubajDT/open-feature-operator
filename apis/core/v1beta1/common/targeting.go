package common

const TARGETING = `
{
	"$id": "https://flagd.dev/schema/v0/targeting.json",
	"$schema": "http://json-schema.org/draft-07/schema#",
	"title": "flagd Targeting",
	"description": "Defines targeting logic for flagd; a extension of JSONLogic, including purpose-built feature-flagging operations.",
	"type": "object",
	"$defs": {
	  "targeting": {
		"title": "Targeting",
		"description": "An expression returning a value which is coerced to a string to be used as a targeting key, or null (to fall back to defaultVariant). If targeting returns a value which is not a variant key, it's considered an error.",
		"anyOf": [
		  {
			"$comment": "we need this to support empty targeting",
			"type": "object",
			"additionalProperties": false,
			"properties": {}
		  },
		  {
			"$ref": "#/$defs/anyRule"
		  }
		]
	  },
	  "primitive": {
		"oneOf": [
		  {
			"description": "When returned from rules, a null value \"exits\", the targeting, and the \"defaultValue\" is returned, with the reason indicating the targeting did not match.",
			"type": "null"
		  },
		  {
			"description": "When returned from rules, booleans are converted to strings (\"true\"/\"false\"), and used to as keys to retrieve the associated value from the \"variants\" object. Be sure that the returned string is present as a key in the variants!",
			"type": "boolean"
		  },
		  {
			"description": "When returned from rules, the behavior of numbers is not defined.",
			"type": "number"
		  },
		  {
			"description": "When returned from rules, strings are used to as keys to retrieve the associated value from the \"variants\" object. Be sure that the returned string is present as a key in the variants!.",
			"type": "string"
		  },
		  {
			"description": "When returned from rules, strings are used to as keys to retrieve the associated value from the \"variants\" object. Be sure that the returned string is present as a key in the variants!.",
			"type": "array"
		  }
		]
	  },
	  "varRule": {
		"title": "Var Operation",
		"description": "Retrieve data from the provided data object.",
		"type": "object",
		"additionalProperties": false,
		"properties": {
		  "var": {
			"anyOf": [
			  {
				"type": "string",
				"description": "flagd automatically injects \"$flagd.timestamp\" (unix epoch) and \"$flagd.flagKey\" (the key of the flag in evaluation) into the context.",
				"pattern": "^\\$flagd\\.((timestamp)|(flagKey))$"
			  },
			  {
				"not": {
				  "$comment": "this is a negated (not) match of \"$flagd.{some-key}\", which is faster and more compatible that a negative lookahead regex",
				  "type": "string",
				  "description": "flagd automatically injects \"$flagd.timestamp\" (unix epoch) and \"$flagd.flagKey\" (the key of the flag in evaluation) into the context.",
				  "pattern": "^\\$flagd\\..*$"
				}
			  },
			  {
				"type": "array",
				"$comment": "this is to support the form of var with a default... there seems to be a bug here, where ajv gives a warning (not an error) because maxItems doesn't equal the number of entries in items, though this is valid in this case",
				"minItems": 1,
				"items": [
				  {
					"type": "string"
				  }
				],
				"additionalItems": {
				  "anyOf": [
					{
					  "type": "null"
					},
					{
					  "type": "boolean"
					},
					{
					  "type": "string"
					},
					{
					  "type": "number"
					}
				  ]
				}
			  }
			]
		  }
		}
	  },
	  "missingRule": {
		"title": "Missing Operation",
		"description": "Takes an array of data keys to search for (same format as var). Returns an array of any keys that are missing from the data object, or an empty array.",
		"type": "object",
		"additionalProperties": false,
		"properties": {
		  "missing": {
			"type": "array",
			"items": {
			  "type": "string"
			}
		  }
		}
	  },
	  "missingSomeRule": {
		"title": "Missing-Some Operation",
		"description": "Takes a minimum number of data keys that are required, and an array of keys to search for (same format as var or missing). Returns an empty array if the minimum is met, or an array of the missing keys otherwise.",
		"type": "object",
		"additionalProperties": false,
		"properties": {
		  "missing_some": {
			"minItems": 2,
			"maxItems": 2,
			"type": "array",
			"items": [
			  {
				"type": "number"
			  },
			  {
				"type": "array",
				"items": {
				  "type": "string"
				}
			  }
			]
		  }
		}
	  },
	  "binaryOrTernaryOp": {
		"type": "array",
		"minItems": 2,
		"maxItems": 3,
		"items": {
		  "$ref": "#/$defs/args"
		}
	  },
	  "binaryOrTernaryRule": {
		"type": "object",
		"additionalProperties": false,
		"properties": {
		  "substr": {
			"title": "Substring Operation",
			"description": "Get a portion of a string. Give a positive start position to return everything beginning at that index. Give a negative start position to work backwards from the end of the string, then return everything. Give a positive length to express how many characters to return.",
			"$ref": "#/$defs/binaryOrTernaryOp"
		  },
		  "<": {
			"title": "Less-Than/Between Operation. Can be used to test that one value is between two others.",
			"$ref": "#/$defs/binaryOrTernaryOp"
		  },
		  "<=": {
			"title": "Less-Than-Or-Equal-To/Between Operation. Can be used to test that one value is between two others.",
			"$ref": "#/$defs/binaryOrTernaryOp"
		  }
		}
	  },
	  "binaryOp": {
		"type": "array",
		"minItems": 2,
		"maxItems": 2,
		"items": {
		  "$ref": "#/$defs/args"
		}
	  },
	  "binaryRule": {
		"title": "Binary Operation",
		"description": "Any primitive JSONLogic operation with 2 operands.",
		"type": "object",
		"additionalProperties": false,
		"properties": {
		  "if": {
			"title": "If Operator",
			"description": "The if statement takes 1 or more arguments: a condition (\"if\"), what to do if its true (\"then\", optional, defaults to returning true), and what to do if its false (\"else\", optional, defaults to returning false). Note that the else condition can be used as an else-if statement by adding additional arguments.",
			"$ref": "#/$defs/variadicOp"
		  },
		  "==": {
			"title": "Lose Equality Operation",
			"description": "Tests equality, with type coercion. Requires two arguments.",
			"$ref": "#/$defs/binaryOp"
		  },
		  "===": {
			"title": "Strict Equality Operation",
			"description": "Tests strict equality. Requires two arguments.",
			"$ref": "#/$defs/binaryOp"
		  },
		  "!=": {
			"title": "Lose Inequality Operation",
			"description": "Tests not-equal, with type coercion.",
			"$ref": "#/$defs/binaryOp"
		  },
		  "!==": {
			"title": "Strict Inequality Operation",
			"description": "Tests strict not-equal.",
			"$ref": "#/$defs/binaryOp"
		  },
		  ">": {
			"title": "Greater-Than Operation",
			"$ref": "#/$defs/binaryOp"
		  },
		  ">=": {
			"title": "Greater-Than-Or-Equal-To Operation",
			"$ref": "#/$defs/binaryOp"
		  },
		  "%": {
			"title": "Modulo Operation",
			"description": "Finds the remainder after the first argument is divided by the second argument.",
			"$ref": "#/$defs/binaryOp"
		  },
		  "/": {
			"title": "Division Operation",
			"$ref": "#/$defs/binaryOp"
		  },
		  "map": {
			"title": "Map Operation",
			"description": "Perform an action on every member of an array. Note, that inside the logic being used to map, var operations are relative to the array element being worked on.",
			"$ref": "#/$defs/binaryOp"
		  },
		  "filter": {
			"title": "Filter Operation",
			"description": "Keep only elements of the array that pass a test. Note, that inside the logic being used to filter, var operations are relative to the array element being worked on.",
			"$ref": "#/$defs/binaryOp"
		  },
		  "all": {
			"title": "All Operation",
			"description": "Perform a test on each member of that array, returning true if all pass. Inside the test code, var operations are relative to the array element being tested.",
			"$ref": "#/$defs/binaryOp"
		  },
		  "none": {
			"title": "None Operation",
			"description": "Perform a test on each member of that array, returning true if none pass. Inside the test code, var operations are relative to the array element being tested.",
			"$ref": "#/$defs/binaryOp"
		  },
		  "some": {
			"title": "Some Operation",
			"description": "Perform a test on each member of that array, returning true if some pass. Inside the test code, var operations are relative to the array element being tested.",
			"$ref": "#/$defs/binaryOp"
		  },
		  "in": {
			"title": "In Operation",
			"description": "If the second argument is an array, tests that the first argument is a member of the array.",
			"$ref": "#/$defs/binaryOp"
		  }
		}
	  },
	  "reduceRule": {
		"type": "object",
		"additionalProperties": false,
		"properties": {
		  "reduce": {
			"title": "Reduce Operation",
			"description": "Combine all the elements in an array into a single value, like adding up a list of numbers. Note, that inside the logic being used to reduce, var operations only have access to an object with a \"current\" and a \"accumulator\".",
			"type": "array",
			"minItems": 3,
			"maxItems": 3,
			"items": {
			  "$ref": "#/$defs/args"
			}
		  }
		}
	  },
	  "associativeOp": {
		"type": "array",
		"minItems": 2,
		"items": {
		  "$ref": "#/$defs/args"
		}
	  },
	  "associativeRule": {
		"title": "Mathematically Associative Operation",
		"description": "Operation applicable to 2 or more parameters.",
		"type": "object",
		"additionalProperties": false,
		"properties": {
		  "*": {
			"title": "Multiplication Operation",
			"description": "Multiplication; associative, will accept and unlimited amount of arguments.",
			"$ref": "#/$defs/associativeOp"
		  }
		}
	  },
	  "unaryOp": {
		"type": "array",
		"minItems": 1,
		"maxItems": 1,
		"items": {
		  "$ref": "#/$defs/args"
		}
	  },
	  "unaryRule": {
		"title": "Unary Operation",
		"description": "Any primitive JSONLogic operation with 1 operands.",
		"type": "object",
		"additionalProperties": false,
		"properties": {
		  "!": {
			"title": "Negation Operation",
			"description": "Logical negation (“not”). Takes just one argument.",
			"$ref": "#/$defs/unaryOp"
		  },
		  "!!": {
			"title": "Double Negation Operation",
			"description": "Double negation, or 'cast to a boolean'. Takes a single argument.",
			"$ref": "#/$defs/unaryOp"
		  }
		}
	  },
	  "variadicOp": {
		"type": "array",
		"minItems": 1,
		"items": {
		  "$ref": "#/$defs/args"
		}
	  },
	  "variadicRule": {
		"$comment": "note < and <= can be used with up to 3 ops (between)",
		"type": "object",
		"additionalProperties": false,
		"properties": {
		  "or": {
			"title": "Or Operation",
			"description": "Simple boolean test, with 1 or more arguments. At a more sophisticated level, \"or\" returns the first truthy argument, or the last argument.",
			"$ref": "#/$defs/variadicOp"
		  },
		  "and": {
			"title": "",
			"description": "Simple boolean test, with 1 or more arguments. At a more sophisticated level, \"and\" returns the first falsy argument, or the last argument.",
			"$ref": "#/$defs/variadicOp"
		  },
		  "+": {
			"title": "Addition Operation",
			"description": "Addition; associative, will accept and unlimited amount of arguments.",
			"$ref": "#/$defs/variadicOp"
		  },
		  "-": {
			"title": "Subtraction Operation",
			"$ref": "#/$defs/variadicOp"
		  },
		  "max": {
			"title": "Maximum Operation",
			"description": "Return the maximum from a list of values.",
			"$ref": "#/$defs/variadicOp"
		  },
		  "min": {
			"title": "Minimum Operation",
			"description": "Return the minimum from a list of values.",
			"$ref": "#/$defs/variadicOp"
		  },
		  "merge": {
			"title": "Merge Operation",
			"description": "Takes one or more arrays, and merges them into one array. If arguments aren't arrays, they get cast to arrays.",
			"$ref": "#/$defs/variadicOp"
		  },
		  "cat": {
			"title": "Concatenate Operation",
			"description": "Concatenate all the supplied arguments. Note that this is not a join or implode operation, there is no “glue” string.",
			"$ref": "#/$defs/variadicOp"
		  }
		}
	  },
	  "stringCompareArg": {
		"oneOf": [
		  {
			"type": "string"
		  },
		  {
			"$ref": "#/$defs/anyRule"
		  }
		]
	  },
	  "stringCompareArgs": {
		"type": "array",
		"minItems": 2,
		"maxItems": 2,
		"items": {
		  "$ref": "#/$defs/stringCompareArg"
		}
	  },
	  "stringCompareRule": {
		"type": "object",
		"additionalProperties": false,
		"properties": {
		  "starts_with": {
			"title": "Starts-With Operation",
			"description": "The string attribute starts with the specified string value.",
			"$ref": "#/$defs/stringCompareArgs"
		  },
		  "ends_with": {
			"title": "Ends-With Operation",
			"description": "The string attribute ends with the specified string value.",
			"$ref": "#/$defs/stringCompareArgs"
		  }
		}
	  },
	  "semVerString": {
		"title": "Semantic Version String",
		"description": "A string representing a valid semantic version expression as per https://semver.org/.",
		"type": "string",
		"pattern": "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$"
	  },
	  "ruleSemVer": {
		"type": "object",
		"additionalProperties": false,
		"properties": {
		  "sem_ver": {
			"title": "Semantic Version Operation",
			"description": "Attribute matches a semantic version condition. Accepts \"npm-style\" range specifiers: \"=\", \"!=\", \">\", \"<\", \">=\", \"<=\", \"~\" (match minor version), \"^\" (match major version).",
			"type": "array",
			"minItems": 3,
			"maxItems": 3,
			"items": [
			  {
				"oneOf": [
				  {
					"$ref": "#/$defs/semVerString"
				  },
				  {
					"$ref": "#/$defs/varRule"
				  }
				]
			  },
			  {
				"description": "Range specifiers: \"=\", \"!=\", \">\", \"<\", \">=\", \"<=\", \"~\" (match minor version), \"^\" (match major version).",
				"enum": [
				  "=",
				  "!=",
				  ">",
				  "<",
				  ">=",
				  "<=",
				  "~",
				  "^"
				]
			  },
			  {
				"oneOf": [
				  {
					"$ref": "#/$defs/semVerString"
				  },
				  {
					"$ref": "#/$defs/varRule"
				  }
				]
			  }
			]
		  }
		}
	  },
	  "fractionalWeightArg": {
		"$comment": "if we remove the \"sum to 100\" restriction, update the descriptions below!",
		"description": "Distribution for all possible variants, with their associated weighting out of 100.",
		"type": "array",
		"minItems": 2,
		"maxItems": 2,
		"items": [
		  {
			"description": "If this bucket is randomly selected, this string is used to as a key to retrieve the associated value from the \"variants\" object.",
			"type": "string"
		  },
		  {
			"description": "Weighted distribution for this variant key (must sum to 100).",
			"type": "number"
		  }
		]
	  },
	  "fractionalOp": {
		"type": "array",
		"minItems": 3,
		"$comment": "there seems to be a bug here, where ajv gives a warning (not an error) because maxItems doesn't equal the number of entries in items, though this is valid in this case",
		"items": [
		  {
			"description": "Bucketing value used in pseudorandom assignment; should be unique and stable for each subject of flag evaluation. Defaults to a concatenation of the flagKey and targetingKey.",
			"$ref": "#/$defs/anyRule"
		  },
		  {
			"$ref": "#/$defs/fractionalWeightArg"
		  },
		  {
			"$ref": "#/$defs/fractionalWeightArg"
		  }
		],
		"additionalItems": {
		  "$ref": "#/$defs/fractionalWeightArg"
		}
	  },
	  "fractionalShorthandOp": {
		"type": "array",
		"minItems": 2,
		"items": {
		  "$ref": "#/$defs/fractionalWeightArg"
		}
	  },
	  "fractionalRule": {
		"type": "object",
		"additionalProperties": false,
		"properties": {
		  "fractional": {
			"title": "Fractional Operation",
			"description": "Deterministic, pseudorandom fractional distribution.",
			"oneOf": [
			  {
				"$ref": "#/$defs/fractionalOp"
			  },
			  {
				"$ref": "#/$defs/fractionalShorthandOp"
			  }
			]
		  }
		}
	  },
	  "reference": {
		"additionalProperties": false,
		"type": "object",
		"$comment": "patternProperties here is a bit of a hack to prevent this definition from being dereferenced early.",
		"patternProperties": {
		  "^\\$ref$": {
			"title": "Reference",
			"description": "A reference to another entity, used for $evaluators (shared rules).",
			"type": "string"
		  }
		}
	  },
	  "args": {
		"oneOf": [
		  {
			"$ref": "#/$defs/reference"
		  },
		  {
			"$ref": "#/$defs/anyRule"
		  },
		  {
			"$ref": "#/$defs/primitive"
		  }
		]
	  },
	  "anyRule": {
		"anyOf": [
		  {
			"$ref": "#/$defs/varRule"
		  },
		  {
			"$ref": "#/$defs/missingRule"
		  },
		  {
			"$ref": "#/$defs/missingSomeRule"
		  },
		  {
			"$ref": "#/$defs/binaryRule"
		  },
		  {
			"$ref": "#/$defs/binaryOrTernaryRule"
		  },
		  {
			"$ref": "#/$defs/associativeRule"
		  },
		  {
			"$ref": "#/$defs/unaryRule"
		  },
		  {
			"$ref": "#/$defs/variadicRule"
		  },
		  {
			"$ref": "#/$defs/reduceRule"
		  },
		  {
			"$ref": "#/$defs/stringCompareRule"
		  },
		  {
			"$ref": "#/$defs/ruleSemVer"
		  },
		  {
			"$ref": "#/$defs/fractionalRule"
		  }
		]
	  }
	}
  }  
`
