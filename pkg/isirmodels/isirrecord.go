package isirmodels

import "time"

type ISIRRecord struct {
	YearIndicator            string    // 1
	FAFSAUUID                string    // 2
	TransactionUUID          string    // 3
	PersonUUID               string    // 4
	TransactionNumber        int       // 5
	DependencyModel          string    // 6 - const
	ApplicationSource        int       // 7 - enum
	ApplicationReceiptDate   time.Time // 8
	TransactionSource        int       // 9  - enum
	TransactionType          string    // 10 - const
	TransactionLanguage      string    // 11 - const
	TransactionReceiptDate   time.Time // 12
	TransactionProcessedDate time.Time // 13
	TransactionStatus        string    // 14
	RenewalDataUsed          bool      // 15
	FPSCorrectionReason      string    // 16
	SAIChangeFlag            string    // 17
	SAI                      int       // 18
	ProvisionalSAI           int       // 19
	SAIFormula               string    // 20
	SAIComputationType       int       // 21
	MaxPellIndicator         string    // 22
	MinimumPellIndicator     string    // 23
	// 24 skipped, it is filler
	FirstName     string    // 25
	MiddleName    string    // 26
	LastName      string    // 27
	Suffix        string    // 28
	DateOfBirth   time.Time // 29
	SSN           string    // 30
	ITIN          string    // 31
	PhoneNumber   string    // 32
	EmailAddress  string    // 33
	StreetAddress string    // 34
	City          string    // 35
	State         string    // 36
	ZipCode       string    // 37
	Country       string    // 38
	// 39 skipped, it is filler
	MaritalStatus                               string // 40
	GradeLevel                                  string // 41
	FirstBachelorsDegreeBefore2526              string // 42
	PursuingTeacherCertification                string // 43
	ActiveDuty                                  string // 44
	Veteran                                     string // 45
	ChildOrOtherDependents                      string // 46
	ParentsDeceased                             string // 47
	WardOfCourt                                 string // 48
	InFosterCare                                string // 49
	EmancipatedMinor                            string // 50
	LegalGuardianship                           string // 51
	PersonalCircumstancesNoneOfTheAbove         string // 52
	UnaccompaniedHomelessYouthAndSelfSupporting string // 53
	UnaccompaniedHomelessGeneral                string // 54
	UnaccompaniedHomelessHS                     string // 55
	UnaccompaniedHomelessTRIO                   string // 56
	UnaccompaniedHomelessFAA                    string // 57
	StudentHomelessnessNoneOfTheAbove           string // 58
	UnusualCircumstance                         string // 59
	UnsubOnly                                   string // 60
	UpdatedFamilySize                           string // 61
	NumberInCollege                             string // 62
	CitizenshipStatus                           string // 63
	ANumber                                     string // 64
	StateOfLegalResidence                       string // 65
	LegalResidenceDate                          string // 66
	EitherParentAttendCollege                   string // 67
	ParentKilledInTheLineOfDuty                 string // 68
	HighSchoolCompletionStatus                  string // 69

}
