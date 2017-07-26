package internal

// enum output_order_type {raw,wesn,wens};
type output_order_type int

const (
	raw  = 0
	wesn = 1
	wens = 2
)

const (
	UNDEFINED      = 9.999e20
	UNDEFINED_LOW  = 9.9989e20
	UNDEFINED_HIGH = 9.9991e20
)

// enum output_order_type output_order, output_order_wanted;
var output_order output_order_type = wesn
var output_order_wanted output_order_type = wesn

// #define DEFINED_VAL(x) ((x) < UNDEFINED_LOW || (x) > UNDEFINED_HIGH)
func DEFINED_VAL(x float) bool {
	return (x) < UNDEFINED_LOW || (x) > UNDEFINED_HIGH
}

type gribtable_s struct {
	disc      int /* Section 0 Discipline                                */
	mtab_set  int /* Section 1 Master Tables Version Number used by set_var      */
	mtab_low  int /* Section 1 Master Tables Version Number low range of tables  */
	mtab_high int /* Section 1 Master Tables Version Number high range of tables */
	cntr      int /* Section 1 originating centre, used for local tables */
	ltab      int /* Section 1 Local Tables Version Number               */
	pcat      int /* Section 4 Template 4.0 Parameter category           */
	pnum      int /* Section 4 Template 4.0 Parameter number             */
	name      string
	desc      string
	unit      string
}
