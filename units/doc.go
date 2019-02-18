/*

The units package represents valid units of measure (the Unit type) and
conversions between them (the Mult type). It also provides a MultSetter type
that satisfies the param.Setter interface. This can be used to specify a Mult
through a command line parameter.

The Mult type can be used to convert between a value expressed in the base
unit and some other chosen unit. The Mult values show how many of the base
unit are in the alternative unit. To convert a value held in the base units
into a value expressed in the alternative units you should divide by the Mult
Val field.

*/
package units
