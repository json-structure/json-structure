# JSON Structure

JSON Structure (Structure) is a validation language for JSON objects. 
It is inspired by, and borrows heavily from,
[JSON Schema](http://json-schema.org/). Structure has been designed
to take advantage of host languages with static type systems. Other
primary design goals include the capacity to compose multiple
structures with ease, and the relative ease of writing property-based
generators for QuickCheck-type testing of JSON objects.

## High Level Overview

Below is a comment-annotated JSON Structure that describes the
primary features of a JSON Structure. Because the document
has comments it is not a valid JSON object. _Sshhh..._ I won't tell
if you won't tell.

```
{
    "title": "Voynich manuscript",
    "description": "Not a hoax",

    // fragments are pieces of type declarations.
    // They cannot be used directly but they
    // can be composed into other fragments or types.
    // See the README section on composition.

    "fragments": {
        "foo": { ... },
        "bar": { ... }
    },

    // types is a mapping of type names to
    // type declarations. Type declarations are
    // the primary components of JSON Structures.

    "types": {
       "a": { ... },
       "b": { ... },
       "c": { ... }
    },

    // main is the type declaration that is used
    // for validation against JSON values.

    "main": { ... }
}
```

## An Example

Let's work through a concrete example to learn about what can
be specified using JSON Structures.

```
{
    "title": "Linked List",
    "description": "Too many linked lists",

    "fragments": {
        "positive": {
            "exclusiveMinimum": 0
        },
        "even": {
            "multipleOf": 2
        }
    },

    "types": {
        "node": {
            "type": "struct",
            "nullable": true,
            "fields": {
                "data": {
                    "type": "integer",
                    "\u0ADD": ["positive", "even"]
                },
                "next": {
                    "type": "node"
                }
            }
        }
    },

    "main": {
        "type": "node"
    }
}
```

The following JSON object validates against the previous structure.

```
{ "data": 2, "next": { "data": 4, "next": { "data": 6, "next": null }}}
```

## Composition

The composition mechanism allows for the reuse of structures without
imposing a specific hierarchy among types. Composition is closer
in behavior to macro expansion than it is to object inheritance.
But composition has been designed to easily model object
inheritance.

The compose property can be declared in any JSON object in
a JSON Structure. Composition is applied to transform the JSON
Structure based on the compose properties. The compose
property is written using the Unicode unassigned code point "\u0ADD".

Composition is applied to all JSON objects in a JSON Structure.
To compose a JSON object, first recursively apply the compose
operation to all children of the JSON object. Next, look to
see whether the "\u0ADD" property is defined. If it is not
defined then terminate the composition.

If "\u0ADD" is defined, the composition will be recursively
merging a series of JSON objects. Begin the composition
with a destination JSON object that is an empty object. 
Traverse in order through the names in the "\u0ADD" array.
Each name must correspond to a defined fragment or type.
After locating the fragment or type, apply the composition
operation on that JSON object. Then recursively merge that
object into the destination object.

A recursive merge writes all the properties of
the source object into the destination object. If the child
property is itself a JSON object then apply the recursive
merge to the child property.

After the traversal of the "\u0ADD" array then complete
the compose operation by recursively merging the remaining
properties of the original JSON object into the destination
object. Replace the original object with the destination object.

Below is the same algorithm described in pseudocode rather
than prose. For simplicitly the error paths have been omitted.

![compose algorithm](images/compose-scaled.png#cachebuster)

![merge algorithm](images/merge-scaled.png#cachebuster)

Fragments exist solely for the purpose of composition. Fragments
are not required to be valid type definitions so they cannot be
used in the types section unless they are used in a compose
property.

## A Larger Example

A JSON Structure specification for JSON Structures
[is provided](https://github.com/json-structure/json-structure/blob/master/specification/json-structure.json).
Notice how the "\u0ADD" property is nested to facilitate
the reuse of declarations. The specification is a little longer
than 250 lines. The [expanded version](https://github.com/json-structure/json-structure/blob/master/specification/json-structure-composed.json) with composition sections expanded to their full form is 670 lines.

As with more specification languages, not all constraints of a valid
JSON Structure can be described with a JSON structure. The goal
is to describe a significant amount of the semantics and rely
on each application to fill in the missing pieces.

## Primitive Types

JSON Structure supports the following primitive types.

| Type    |  Description |
| -----   |  ----------- |
| boolean | true or false |
| integer | mathematical integer  |
| number  | JSON number (mathematical rational number) |
| string  | JSON string |
| json    | raw json values |
| struct  | structure or record |
| array   | sequence of values of the same type |
| set     | typed collection of unique values |
| map     | typed associative array of strings to values |
| union   | sum type |

## Type Declarations

JSON Structures are composed primarily of type declarations.
Here is a description of the properties that define a
type declaration. "type" is the only property that is
required for all type declarations. Some primitive types
("struct", "array", "set", "map", "union") have required properties.
The remaining properties are optional.

| Property      | JSON |  Description |
| ------------- | ---- |  ----------- |
| type     | string | Required. Must be either the name of a primitive type or a name defined in the "types" map. |

**Properties Shared By All Primitive Types**

| Property      | JSON |  Description |
| ------------- | ---- |  ----------- |
| format   | string |  Defines additional semantic validation |
| nullable | boolean|  If true then allow the null value. Default is false |
| optional | boolean | If true then allow the value to be missing. Requires default value |
| default  | JSON value | Use this value when none is provided. |
| enum | list of JSON values | Value must be one of the provided elements. Optional |

**Number and Integer Properties**

| Property      | JSON |  Description |
| ------------- | ---- |  ----------- |
| multipleOf | number | See JSON Schema |
| minimum | number | See JSON Schema |
| maximum | number | See JSON Schema |
| exclusiveMinimum | number | See JSON Schema |
| exclusiveMaximum | number | See JSON Schema |

**String Properties**

| Property      | JSON |  Description |
| ------------- | ---- |  ----------- |
| pattern | string | See JSON Schema |
| minLength | integer | See JSON Schema |
| maxLength | integer | See JSON Schema |

**Struct Properties**

| Property      | JSON |  Description |
| ------------- | ---- |  ----------- |
| fields | map | Required. Mapping of field names to type declaration. |

**Array Properties**

| Property      | JSON |  Description |
| ------------- | ---- |  ----------- |
| item | map | Required. Type declaration applied to each item. |
| minItems | integer | See JSON Schema |
| maxItems | integer | See JSON Schema |

**Set Properties**

| Property      | JSON |  Description |
| ------------- | ---- |  ----------- |
| item | map | Required. Type declaration applied to each item. |
| minItems | integer | See JSON Schema |
| maxItems | integer | See JSON Schema |

**Map Properties**

| Property      | JSON |  Description |
| ------------- | ---- |  ----------- |
| item | map | Required. Type declaration applied to each item. |
| minItems | integer | See JSON Schema |
| maxItems | integer | See JSON Schema |

**Union Properties**

The keys of the types map are not using during object validation.
Types is a JSON object and therefore recursively composed
when composition is applied.

| Property      | JSON |  Description |
| ------------- | ---- |  ----------- |
| types | map | Required. Mapping of names to type declaration. |

## Pattern Property

A JSON Structure validator must be able to test whether
all 'pattern' properties fall into the restricted regular
expression grammar from JSON Schema section 3.3. We
provide an [ANTLRv4 Grammar](simpleregexp/SimpleRegexp.g4)
to implement this requirement. A JSON Structure validator
may provide an option to parse the JSON Structure with
the host language's native regular expression dialect.

## Declarations and Cycles

A type declaration and a fragment declaration cannot have the same name.
Types and fragments cannot have the name of primitive types.

A JSON Structure validator must check for cycles in 'compose'
properties and report an error when compose cycles are detected.
A JSON Structure validator must check for cycles in type
aliasing and report an error when type cycles are detected.
This is an example of a type alias.

```
   "types": {
        "foo": { "type": "bar" }
   }
```

## Contributing

JSON Structure is licensed under the Apache Public License
Version 2. Please open an issue before submitting a pull
request. We tend to be conservative with new feature requests,
so opening an issue lets us talk about the proposal before
you spend the effort writing code.

Another source of concern is feature interaction.
In general there should only be one way to do a thing.

## FAQ

1\. How do I specify a struct field as optional or required?

If a field's type declaration has an 'default' property then the
field is optional. Otherwise the field is required.

2\. How do I specify a type declaration for extra fields in a struct?

You can't. Consider using the 'map' type instead of the 'struct' type.

3\. Why are you throwing shade on JSON Schema?

We aren't. There are many good use cases for JSON Schema. And there
probably a few things you can do using JSON Schema that you can't
do using JSON Structure. The JSON Structure project started as
an exercise in quickly generating property-based generators for
JSON objects. Along the way we found that we could simplify the
design if we started from scratch.
