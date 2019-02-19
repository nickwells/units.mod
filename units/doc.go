/*

Package units represents valid units of measure (the Unit type) and families
of related units (the Family type). It also provides a UnitSetter type that
satisfies the param.Setter interface. This can be used to specify a Unit
through a command line parameter.

The Unit type can be used to convert between a value expressed in the base
unit and some other chosen unit. The Unit values show how many of the base
unit are in the alternative unit. To convert a value held in the base units
into a value expressed in the alternative units you should divide by the Unit
Factor field.

*/
package units
