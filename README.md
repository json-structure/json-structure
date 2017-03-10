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
    // the primary components of JSON structures.

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
                    "compose": ["positive", "even"],
                    "type": "integer"
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

The 'compose' property can be declared in any JSON object in
a JSON Structure. Composition is applied to transform the JSON
Structure based on the 'compose' properties.

Composition is applied to all JSON objects in a JSON Structure.
To compose a JSON object, first recursively apply the compose
operation to all children of the JSON object. Next, look to
see whether the 'compose' property is defined. If it is not
defined then terminate.

If 'compose' is defined, the composition will be recursively
merging a series of JSON objects. Begin the composition
with a destination JSON object that is an empty object. 
Traverse in order through the names in the 'compose' array.
Each name must correspond to a defined fragment or type.
After locating the fragment or type, apply the 'compose'
operation on that JSON object. Then merge that object
into the destination object.

A recursive merge writes all the properties of
the source object into the destination object. If the child
property is itself a JSON object then apply the recursive
merge to the child property.

After the traversal of the 'compose' array then complete
the compose operation by recursively merging the remaining
properties of the original JSON object into the destination
object. Replace the original object with the destination object.

Fragments exist solely for the purpose of composition. Fragments
are not required to be valid type definitions so they cannot be
used in the types section. Fragments can be composed with
other fragments or with types. Types can be composed with
other types.

**TODO:** show psuedocode that implements what was described

## Primitive Types

JSON Structure supports the following primitive types.

| Type    |  Description |
| -----   |  ----------- |
| boolean | true or false |
| integer | mathematical integer  |
| number  | JSON number (rational number) |
| string  | JSON string |
| struct  | structure or record |
| array   | sequence of values of the same type |
| json    | Raw json values. Coming soon |
| union   | Tagged union. Coming soon |
| set     | Collection type. Coming soon |
| map     | Collection type. Coming soon |

## Declaration Names

A type name and a fragment name cannot have the same name.
Types and fragments cannot have the name of primitive types.
A JSON Structure validator must check for cycles in 'compose'
properties and report an error when cycles are detected.

The following is valid.

```
   "types": {
        "foo": { "type": "bar" }
   }
```

"foo" is defined as a type alias for "bar". A JSON Structure
validator must check for cycles in type aliasing and report
an error when cycles are detected.

## Type Declarations

JSON Structures are composed primarily of type declarations.
Here is a description of the properties that define a
type declaration.


| Property      | JSON |  Description |
| ------------- | ---- |  ----------- |
| type     | string |  Required. Must be either the name of a primitive type or a name defined in the "types" map. |

**Properties Shared By All Primitive Types**

| default  | JSON value |  Use this value when none is provided |
| format   | string |  Defines additional semantic validation |
| nullable | boolean|  If true then allow the null value. Default is false. |


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

## FAQ

1\. How do I specify a struct field as optional or required?

If a field's type declaration has an 'default' property then the
field is optional. Otherwise the field is required.

2\. How do I specify a type declaration for extra fields in a struct?

You can't.
