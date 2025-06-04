// The isirparser package contains behavior for reading an Institutional Student
// Information Record (ISIR) from the standard fixed format provided by the
// Department of Education and populating an isirecord struct with the data
//
// The fixed format does change over time, and the changes coincide with new
// award years (AY) so this package is broken into functionaly by AY.
//
// The bulk of the logic is intended to be entirely internal to this package,
// while the ISIRParser interface and implementations are the external entry
// points available outside of the package.
package isirparser

import "github.com/rpatton4/fsa/pkg/isirmodels"

type ISIRParser interface {
	ParseISIR(record string) (isirmodels.ISIRecord, error)
}
