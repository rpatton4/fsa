// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsamodels

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ISIRecord struct {
	// Field # 1
	YearIndicator string `json:",omitempty"`

	// Field # 2
	FAFSAUUID string `json:",omitempty"`

	// Field # 3
	TransactionUUID string `json:",omitempty"`

	// Field # 4
	PersonUUID string `json:",omitempty"`

	// Field # 5
	TransactionNumber string `json:",omitempty"`

	// Field # 6
	DependencyModel string `json:",omitempty"`

	// Field # 7
	ApplicationSource string `json:",omitempty"`

	// Field # 8
	ApplicationReceiptDate *time.Time `json:",omitempty"`

	// Field # 9
	TransactionSource string `json:",omitempty"`

	// Field # 10
	TransactionType string `json:",omitempty"`

	// Field # 11
	TransactionLanguage string `json:",omitempty"`

	// Field # 12
	TransactionReceiptDate *time.Time `json:",omitempty"`

	// Field # 13
	TransactionProcessedDate *time.Time `json:",omitempty"`

	// Field # 14
	TransactionStatus string `json:",omitempty"`

	// Field # 15
	RenewalDataUsed string `json:",omitempty"`

	// Field # 16
	FPSCorrectionReason string `json:",omitempty"`

	// Field # 17
	SAIChangeFlag string `json:",omitempty"`

	// Field # 18
	SAI string `json:",omitempty"`

	// Field # 19
	ProvisionalSAI string `json:",omitempty"`

	// Field # 20
	SAIFormula string `json:",omitempty"`

	// Field # 21
	SAIComputationType string `json:",omitempty"`

	// Field # 22
	MaxPellIndicator string `json:",omitempty"`

	// Field # 23
	MinimumPellIndicator string `json:",omitempty"`

	// Field # 25
	StudentFirstName string `json:",omitempty"`

	// Field # 26
	StudentMiddleName string `json:",omitempty"`

	// Field # 27
	StudentLastName string `json:",omitempty"`

	// Field # 28
	StudentSuffix string `json:",omitempty"`

	// Field # 29
	StudentDateOfBirth *time.Time `json:",omitempty"`

	// Field # 30
	StudentSSN string `json:",omitempty"`

	// Field # 31
	StudentITIN string `json:",omitempty"`

	// Field # 32
	StudentPhoneNumber string `json:",omitempty"`

	// Field # 33
	StudentEmailAddress string `json:",omitempty"`

	// Field # 34
	StudentStreetAddress string `json:",omitempty"`

	// Field # 35
	StudentCity string `json:",omitempty"`

	// Field # 36
	StudentState string `json:",omitempty"`

	// Field # 37
	StudentZipCode string `json:",omitempty"`

	// Field # 38
	StudentCountry string `json:",omitempty"`

	// Field # 40
	StudentMaritalStatus string `json:",omitempty"`

	// Field # 41
	StudentGradeLevel string `json:",omitempty"`

	// Field # 42
	StudentHasBachelorsDegree string `json:",omitempty"`

	// Field # 43
	StudentPursuingTeacherCertification string `json:",omitempty"`

	// Field # 44
	StudentActiveDuty string `json:",omitempty"`

	// Field # 45
	StudentVeteran string `json:",omitempty"`

	// Field # 46
	StudentChildOrOtherDependents string `json:",omitempty"`

	// Field # 47
	StudentParentsDeceased string `json:",omitempty"`

	// Field # 48
	StudentWardOfCourt string `json:",omitempty"`

	// Field # 49
	StudentInFosterCare string `json:",omitempty"`

	// Field # 50
	StudentEmancipatedMinor string `json:",omitempty"`

	// Field # 51
	StudentLegalGuardianship string `json:",omitempty"`

	// Field # 52
	StudentPersonalCircumstancesNoneOfTheAbove string `json:",omitempty"`

	// Field # 53
	StudentUnaccompaniedHomelessYouthAndSelfSupporting string `json:",omitempty"`

	// Field # 54
	StudentUnaccompaniedHomelessGeneral string `json:",omitempty"`

	// Field # 55
	StudentUnaccompaniedHomelessHS string `json:",omitempty"`

	// Field # 56
	StudentUnaccompaniedHomelessTRIO string `json:",omitempty"`

	// Field # 57
	StudentUnaccompaniedHomelessFAA string `json:",omitempty"`

	// Field # 58
	StudentHomelessnessNoneOfTheAbove string `json:",omitempty"`

	// Field # 59
	StudentUnusualCircumstance string `json:",omitempty"`

	// Field # 60
	StudentUnsubOnly string `json:",omitempty"`

	// Field # 61
	StudentUpdatedFamilySize string `json:",omitempty"`

	// Field # 62
	StudentNumberInCollege string `json:",omitempty"`

	// Field # 63
	StudentCitizenshipStatus string `json:",omitempty"`

	// Field # 64
	StudentANumber string `json:",omitempty"`

	// Field # 65
	StudentStateOfLegalResidence string `json:",omitempty"`

	// Field # 66
	StudentLegalResidenceDate *time.Time `json:",omitempty"`

	// Field # 67
	StudentEitherParentAttendCollege string `json:",omitempty"`

	// Field # 68
	StudentParentKilledInTheLineOfDuty string `json:",omitempty"`

	// Field # 69
	StudentHighSchoolCompletionStatus string `json:",omitempty"`

	// Field # 70
	StudentHighSchoolName string `json:",omitempty"`

	// Field # 71
	StudentHighSchoolCity string `json:",omitempty"`

	// Field # 72
	StudentHighSchoolState string `json:",omitempty"`

	// Field # 73
	StudentHighSchoolEquivalentDiplomaName string `json:",omitempty"`

	// Field # 74
	StudentHighSchoolEquivalentDiplomaState string `json:",omitempty"`

	// Field # 75
	StudentManuallyEnteredReceivedEITC string `json:",omitempty"`

	// Field # 76
	StudentManuallyEnteredReceivedFederalHousingAssistance string `json:",omitempty"`

	// Field # 77
	StudentManuallyEnteredReceivedFreeReducedPriceLunch string `json:",omitempty"`

	// Field # 78
	StudentManuallyEnteredReceivedMedicaid string `json:",omitempty"`

	// Field # 79
	StudentManuallyEnteredReceivedRefundableCreditFor36BHealthPlan string `json:",omitempty"`

	// Field # 80
	StudentManuallyEnteredReceivedSNAP string `json:",omitempty"`

	// Field # 81
	StudentManuallyEnteredReceivedSupplementalSecurityIncome string `json:",omitempty"`

	// Field # 82
	StudentManuallyEnteredReceivedTANF string `json:",omitempty"`

	// Field # 83
	StudentManuallyEnteredReceivedWIC string `json:",omitempty"`

	// Field # 84
	StudentManuallyEnteredFederalBenefitsNoneOfTheAbove string `json:",omitempty"`

	// Field # 85
	StudentManuallyEnteredFiled1040Or1040NR string `json:",omitempty"`

	// Field # 86
	StudentManuallyEnteredFiledNonUSTaxReturn string `json:",omitempty"`

	// Field # 87
	StudentManuallyEnteredFiledJointReturnWithCurrentSpouse string `json:",omitempty"`

	// Field # 88
	StudentManuallyEnteredTaxReturnFilingStatus string `json:",omitempty"`

	// Field # 89
	StudentManuallyEnteredIncomeEarnedFromWork string `json:",omitempty"`

	// Field # 90
	StudentManuallyEnteredTaxExemptInterestIncome string `json:",omitempty"`

	// Field # 91
	StudentManuallyEnteredUntaxedPortionsOfIRADistributions string `json:",omitempty"`

	// Field # 92
	StudentManuallyEnteredIRARollover string `json:",omitempty"`

	// Field # 93
	StudentManuallyEnteredUntaxedPortionsOfPensions string `json:",omitempty"`

	// Field # 94
	StudentManuallyEnteredPensionRollover string `json:",omitempty"`

	// Field # 95
	StudentManuallyEnteredAdjustedGrossIncome string `json:",omitempty"`

	// Field # 96
	StudentManuallyEnteredIncomeTaxPaid string `json:",omitempty"`

	// Field # 97
	StudentManuallyEnteredEITCReceivedDuringTaxYear string `json:",omitempty"`

	// Field # 98
	StudentManuallyEnteredDeductiblePaymentsToIRAKeoghOther string `json:",omitempty"`

	// Field # 99
	StudentManuallyEnteredEducationCredits string `json:",omitempty"`

	// Field # 100
	StudentManuallyEnteredFiledScheduleABDEFH string `json:",omitempty"`

	// Field # 101
	StudentManuallyEnteredScheduleCAmount string `json:",omitempty"`

	// Field # 102
	StudentManuallyEnteredCollegeGrantAndScholarshipAid string `json:",omitempty"`

	// Field # 103
	StudentManuallyEnteredForeignEarnedIncomeExclusion string `json:",omitempty"`

	// Field # 104
	StudentManuallyEnteredChildSupportReceived string `json:",omitempty"`

	// Field # 105
	StudentManuallyEnteredTotalOfCashSavingsAndCheckingAccounts string `json:",omitempty"`

	// Field # 106
	StudentManuallyEnteredNetWorthOfCurrentInvestments string `json:",omitempty"`

	// Field # 107
	StudentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarms string `json:",omitempty"`

	// Field # 108
	StudentCollege1 string `json:",omitempty"`

	// Field # 109
	StudentCollege2 string `json:",omitempty"`

	// Field # 110
	StudentCollege3 string `json:",omitempty"`

	// Field # 111
	StudentCollege4 string `json:",omitempty"`

	// Field # 112
	StudentCollege5 string `json:",omitempty"`

	// Field # 113
	StudentCollege6 string `json:",omitempty"`

	// Field # 114
	StudentCollege7 string `json:",omitempty"`

	// Field # 115
	StudentCollege8 string `json:",omitempty"`

	// Field # 116
	StudentCollege9 string `json:",omitempty"`

	// Field # 117
	StudentCollege10 string `json:",omitempty"`

	// Field # 118
	StudentCollege11 string `json:",omitempty"`

	// Field # 119
	StudentCollege12 string `json:",omitempty"`

	// Field # 120
	StudentCollege13 string `json:",omitempty"`

	// Field # 121
	StudentCollege14 string `json:",omitempty"`

	// Field # 122
	StudentCollege15 string `json:",omitempty"`

	// Field # 123
	StudentCollege16 string `json:",omitempty"`

	// Field # 124
	StudentCollege17 string `json:",omitempty"`

	// Field # 125
	StudentCollege18 string `json:",omitempty"`

	// Field # 126
	StudentCollege19 string `json:",omitempty"`

	// Field # 127
	StudentCollege20 string `json:",omitempty"`

	// Field # 128
	StudentConsentToRetrieveAndDiscloseFTI string `json:",omitempty"`

	// Field # 129
	StudentSignature string `json:",omitempty"`

	// Field # 130
	StudentSignatureDate *time.Time `json:",omitempty"`

	// Field # 132
	StudentSpouseFirstName string `json:",omitempty"`

	// Field # 133
	StudentSpouseMiddleName string `json:",omitempty"`

	// Field # 134
	StudentSpouseLastName string `json:",omitempty"`

	// Field # 135
	StudentSpouseSuffix string `json:",omitempty"`

	// Field # 136
	StudentSpouseDateOfBirth *time.Time `json:",omitempty"`

	// Field # 137
	StudentSpouseSSN string `json:",omitempty"`

	// Field # 138
	StudentSpouseITIN string `json:",omitempty"`

	// Field # 139
	StudentSpousePhoneNumber string `json:",omitempty"`

	// Field # 140
	StudentSpouseEmailAddress string `json:",omitempty"`

	// Field # 141
	StudentSpouseStreetAddress string `json:",omitempty"`

	// Field # 142
	StudentSpouseCity string `json:",omitempty"`

	// Field # 143
	StudentSpouseState string `json:",omitempty"`

	// Field # 144
	StudentSpouseZipCode string `json:",omitempty"`

	// Field # 145
	StudentSpouseCountry string `json:",omitempty"`

	// Field # 146
	StudentSpouseFiled1040Or1040NR string `json:",omitempty"`

	// Field # 147
	StudentSpouseFiledNonUSTaxReturn string `json:",omitempty"`

	// Field # 148
	StudentSpouseTaxReturnFilingStatus string `json:",omitempty"`

	// Field # 149
	StudentSpouseIncomeEarnedFromWork string `json:",omitempty"`

	// Field # 150
	StudentSpouseTaxExemptInterestIncome string `json:",omitempty"`

	// Field # 151
	StudentSpouseUntaxedPortionsOfIRADistributions string `json:",omitempty"`

	// Field # 152
	StudentSpouseIRARollover string `json:",omitempty"`

	// Field # 153
	StudentSpouseUntaxedPortionsOfPensions string `json:",omitempty"`

	// Field # 154
	StudentSpousePensionRollover string `json:",omitempty"`

	// Field # 155
	StudentSpouseAdjustedGrossIncome string `json:",omitempty"`

	// Field # 156
	StudentSpouseIncomeTaxPaid string `json:",omitempty"`

	// Field # 157
	StudentSpouseDeductiblePaymentsToIRAKeoghOther string `json:",omitempty"`

	// Field # 158
	StudentSpouseEducationCredits string `json:",omitempty"`

	// Field # 159
	StudentSpouseFiledScheduleABDEFH string `json:",omitempty"`

	// Field # 160
	StudentSpouseScheduleCAmount string `json:",omitempty"`

	// Field # 161
	StudentSpouseForeignEarnedIncomeExclusion string `json:",omitempty"`

	// Field # 162
	StudentSpouseConsentToRetrieveAndDiscloseFTI string `json:",omitempty"`

	// Field # 163
	StudentSpouseSignature string `json:",omitempty"`

	// Field # 164
	StudentSpouseSignatureDate *time.Time `json:",omitempty"`

	// Field # 166
	ParentFirstName string `json:",omitempty"`

	// Field # 167
	ParentMiddleName string `json:",omitempty"`

	// Field # 168
	ParentLastName string `json:",omitempty"`

	// Field # 169
	ParentSuffix string `json:",omitempty"`

	// Field # 170
	ParentDateOfBirth *time.Time `json:",omitempty"`

	// Field # 171
	ParentSSN string `json:",omitempty"`

	// Field # 172
	ParentITIN string `json:",omitempty"`

	// Field # 173
	ParentPhoneNumber string `json:",omitempty"`

	// Field # 174
	ParentEmailAddress string `json:",omitempty"`

	// Field # 175
	ParentStreetAddress string `json:",omitempty"`

	// Field # 176
	ParentCity string `json:",omitempty"`

	// Field # 177
	ParentState string `json:",omitempty"`

	// Field # 178
	ParentZipCode string `json:",omitempty"`

	// Field # 179
	ParentCountry string `json:",omitempty"`

	// Field # 180
	ParentMaritalStatus string `json:",omitempty"`

	// Field # 181
	ParentStateOfLegalResidence string `json:",omitempty"`

	// Field # 182
	ParentLegalResidenceDate *time.Time `json:",omitempty"`

	// Field # 183
	ParentUpdatedFamilySize string `json:",omitempty"`

	// Field # 184
	ParentNumberInCollege string `json:",omitempty"`

	// Field # 185
	ParentReceivedEITC string `json:",omitempty"`

	// Field # 186
	ParentReceivedFederalHousingAssistance string `json:",omitempty"`

	// Field # 187
	ParentReceivedFreeReducedPriceLunch string `json:",omitempty"`

	// Field # 188
	ParentReceivedMedicaid string `json:",omitempty"`

	// Field # 189
	ParentReceivedRefundableCreditFor36BHealthPlan string `json:",omitempty"`

	// Field # 190
	ParentReceivedSNAP string `json:",omitempty"`

	// Field # 191
	ParentReceivedSupplementalSecurityIncome string `json:",omitempty"`

	// Field # 192
	ParentReceivedTANF string `json:",omitempty"`

	// Field # 193
	ParentReceivedWIC string `json:",omitempty"`

	// Field # 194
	ParentFederalBenefitsNoneOfTheAbove string `json:",omitempty"`

	// Field # 195
	ParentFiled1040Or1040NR string `json:",omitempty"`

	// Field # 196
	ParentFileNonUSTaxReturn string `json:",omitempty"`

	// Field # 197
	ParentFiledJointReturnWithCurrentSpouse string `json:",omitempty"`

	// Field # 198
	ParentTaxReturnFilingStatus string `json:",omitempty"`

	// Field # 199
	ParentIncomeEarnedFromWork string `json:",omitempty"`

	// Field # 200
	ParentTaxExemptInterestIncome string `json:",omitempty"`

	// Field # 201
	ParentUntaxedPortionsOfIRADistributions string `json:",omitempty"`

	// Field # 202
	ParentIRARollover string `json:",omitempty"`

	// Field # 203
	ParentUntaxedPortionsOfPensions string `json:",omitempty"`

	// Field # 204
	ParentPensionRollover string `json:",omitempty"`

	// Field # 205
	ParentAdjustedGrossIncome string `json:",omitempty"`

	// Field # 206
	ParentIncomeTaxPaid string `json:",omitempty"`

	// Field # 207
	ParentEarnedIncomeTaxCreditReceivedDuringTaxYear string `json:",omitempty"`

	// Field # 208
	ParentDeductiblePaymentsToIRAKeoghOther string `json:",omitempty"`

	// Field # 209
	ParentEducationCredits string `json:",omitempty"`

	// Field # 210
	ParentFiledScheduleABDEFH string `json:",omitempty"`

	// Field # 211
	ParentScheduleCAmount string `json:",omitempty"`

	// Field # 212
	ParentCollegeGrantAndScholarshipAid string `json:",omitempty"`

	// Field # 213
	ParentForeignEarnedIncomeExclusion string `json:",omitempty"`

	// Field # 214
	ParentChildSupportReceived string `json:",omitempty"`

	// Field # 215
	ParentTotalOfCashSavingsAndCheckingAccounts string `json:",omitempty"`

	// Field # 216
	ParentNetWorthOfCurrentInvestments string `json:",omitempty"`

	// Field # 217
	ParentNetWorthOfBusinessesAndInvestmentFarms string `json:",omitempty"`

	// Field # 218
	ParentConsentToRetrieveAndDiscloseFTI string `json:",omitempty"`

	// Field # 219
	ParentSignature string `json:",omitempty"`

	// Field # 220
	ParentSignatureDate *time.Time `json:",omitempty"`

	// Field # 222
	ParentSpouseFirstName string `json:",omitempty"`

	// Field # 223
	ParentSpouseMiddleName string `json:",omitempty"`

	// Field # 224
	ParentSpouseLastName string `json:",omitempty"`

	// Field # 225
	ParentSpouseSuffix string `json:",omitempty"`

	// Field # 226
	ParentSpouseDateOfBirth *time.Time `json:",omitempty"`

	// Field # 227
	ParentSpouseSSN string `json:",omitempty"`

	// Field # 228
	ParentSpouseITIN string `json:",omitempty"`

	// Field # 229
	ParentSpousePhoneNumber string `json:",omitempty"`

	// Field # 230
	ParentSpouseEmailAddress string `json:",omitempty"`

	// Field # 231
	ParentSpouseStreetAddress string `json:",omitempty"`

	// Field # 232
	ParentSpouseCity string `json:",omitempty"`

	// Field # 233
	ParentSpouseState string `json:",omitempty"`

	// Field # 234
	ParentSpouseZipCode string `json:",omitempty"`

	// Field # 235
	ParentSpouseCountry string `json:",omitempty"`

	// Field # 236
	ParentSpouseFiled1040Or1040NR string `json:",omitempty"`

	// Field # 237
	ParentSpouseFileNonUSTaxReturn string `json:",omitempty"`

	// Field # 238
	ParentSpouseTaxReturnFilingStatus string `json:",omitempty"`

	// Field # 239
	ParentSpouseIncomeEarnedFromWork string `json:",omitempty"`

	// Field # 240
	ParentSpouseTaxExemptInterestIncome string `json:",omitempty"`

	// Field # 241
	ParentSpouseUntaxedPortionsOfIRADistributions string `json:",omitempty"`

	// Field # 242
	ParentSpouseIRARollover string `json:",omitempty"`

	// Field # 243
	ParentSpouseUntaxedPortionsOfPensions string `json:",omitempty"`

	// Field # 244
	ParentSpousePensionRollover string `json:",omitempty"`

	// Field # 245
	ParentSpouseAdjustedGrossIncome string `json:",omitempty"`

	// Field # 246
	ParentSpouseIncomeTaxPaid string `json:",omitempty"`

	// Field # 247
	ParentSpouseDeductiblePaymentsToIRAKeoghOther string `json:",omitempty"`

	// Field # 248
	ParentSpouseEducationCredits string `json:",omitempty"`

	// Field # 249
	ParentSpouseFiledScheduleABDEFH string `json:",omitempty"`

	// Field # 250
	ParentSpouseScheduleCAmount string `json:",omitempty"`

	// Field # 251
	ParentSpouseForeignEarnedIncomeExclusion string `json:",omitempty"`

	// Field # 252
	ParentSpouseConsentToRetrieveAndDiscloseFTI string `json:",omitempty"`

	// Field # 253
	ParentSpouseSignature string `json:",omitempty"`

	// Field # 254
	ParentSpouseSignatureDate *time.Time `json:",omitempty"`

	// Field # 256
	PreparerFirstName string `json:",omitempty"`

	// Field # 257
	PreparerLastName string `json:",omitempty"`

	// Field # 258
	PreparerSSN string `json:",omitempty"`

	// Field # 259
	PreparerEIN string `json:",omitempty"`

	// Field # 260
	PreparerAffiliation string `json:",omitempty"`

	// Field # 261
	PreparerStreetAddress string `json:",omitempty"`

	// Field # 262
	PreparerCity string `json:",omitempty"`

	// Field # 263
	PreparerState string `json:",omitempty"`

	// Field # 264
	PreparerZipCode string `json:",omitempty"`

	// Field # 265
	PreparerSignature string `json:",omitempty"`

	// Field # 266
	PreparerSignatureDate *time.Time `json:",omitempty"`

	// Field # 268
	StudentAffirmationStatus string `json:",omitempty"`

	// Field # 269
	StudentSpouseAffirmationStatus string `json:",omitempty"`

	// Field # 270
	ParentAffirmationStatus string `json:",omitempty"`

	// Field # 271
	ParentSpouseOrPartnerAffirmationStatus string `json:",omitempty"`

	// Field # 272
	StudentDateConsentGranted *time.Time `json:",omitempty"`

	// Field # 273
	StudentSpouseDateConsentGranted *time.Time `json:",omitempty"`

	// Field # 274
	ParentDateConsentGranted *time.Time `json:",omitempty"`

	// Field # 275
	ParentSpouseOrPartnerDateConsentGranted *time.Time `json:",omitempty"`

	// Field # 276
	StudentTransunionMatchStatus string `json:",omitempty"`

	// Field # 277
	StudentSpouseTransunionMatchStatus string `json:",omitempty"`

	// Field # 278
	StudentParentTransunionMatchStatus string `json:",omitempty"`

	// Field # 279
	StudentParentSpouseTransunionMatchStatus string `json:",omitempty"`

	// Field # 280
	CorrectionAppliedAgainstTransactionNumber string `json:",omitempty"`

	// Field # 281
	ProfessionalJudgement string `json:",omitempty"`

	// Field # 282
	DependencyOverrideIndicator string `json:",omitempty"`

	// Field # 283
	FAAFederalSchoolCode string `json:",omitempty"`

	// Field # 284
	FAASignature string `json:",omitempty"`

	// Field # 285
	IASGIndicator string `json:",omitempty"`

	// Field # 286
	ChildrenOfFallenHeroesIndicator string `json:",omitempty"`

	// Field # 287
	ElectronicTransactionIndicatorDestinationNumber string `json:",omitempty"`

	// Field # 288
	StudentSignatureSource string `json:",omitempty"`

	// Field # 289
	StudentSpouseSignatureSource string `json:",omitempty"`

	// Field # 290
	ParentSignatureSource string `json:",omitempty"`

	// Field # 291
	ParentSpouseOrPartnerSignatureSource string `json:",omitempty"`

	// Field # 292
	SpecialHandlingIndicator string `json:",omitempty"`

	// Field # 293
	AddressOnlyChangeFlag string `json:",omitempty"`

	// Field # 294
	FPSPushedISIRFlag string `json:",omitempty"`

	// Field # 295
	RejectStatusChangeFlag string `json:",omitempty"`

	// Field # 296
	VerificationTrackingFlag string `json:",omitempty"`

	// Field # 297
	StudentSelectedForVerification string `json:",omitempty"`

	// Field # 298
	IncarceratedApplicantFlag string `json:",omitempty"`

	// Field # 299
	NSLDSTransactionNumber string `json:",omitempty"`

	// Field # 300
	NSLDSDatabaseResultsFlag string `json:",omitempty"`

	// Field # 301
	HighSchoolFlag string `json:",omitempty"`

	// Field # 302
	StudentTotalFederalWorkStudyEarnings string `json:",omitempty"`

	// Field # 303
	StudentSpouseTotalFederalWorkStudyEarnings string `json:",omitempty"`

	// Field # 304
	ParentTotalFederalWorkStudyEarnings string `json:",omitempty"`

	// Field # 305
	ParentSpouseOrPartnerTotalFederalWorkStudyEarnings string `json:",omitempty"`

	// Field # 306
	TotalParentAllowancesAgainstIncome string `json:",omitempty"`

	// Field # 307
	ParentPayrollTaxAllowance string `json:",omitempty"`

	// Field # 308
	ParentIncomeProtectionAllowance string `json:",omitempty"`

	// Field # 309
	ParentEmploymentExpenseAllowance string `json:",omitempty"`

	// Field # 310
	ParentAvailableIncome string `json:",omitempty"`

	// Field # 311
	ParentAdjustedAvailableIncome string `json:",omitempty"`

	// Field # 312
	ParentContribution string `json:",omitempty"`

	// Field # 313
	StudentPayrollTaxAllowance string `json:",omitempty"`

	// Field # 314
	StudentIncomeProtectionAllowance string `json:",omitempty"`

	// Field # 315
	StudentAllowanceForParentsNegativeAdjustedAvailableIncome string `json:",omitempty"`

	// Field # 316
	StudentEmploymentExpenseAllowance string `json:",omitempty"`

	// Field # 317
	TotalStudentAllowancesAgainstIncome string `json:",omitempty"`

	// Field # 318
	StudentAvailableIncome string `json:",omitempty"`

	// Field # 319
	StudentContributionFromIncome string `json:",omitempty"`

	// Field # 320
	StudentAdjustedAvailableIncome string `json:",omitempty"`

	// Field # 321
	TotalStudentContributionFromSAAI string `json:",omitempty"`

	// Field # 322
	ParentDiscretionaryNetWorth string `json:",omitempty"`

	// Field # 323
	ParentNetWorth string `json:",omitempty"`

	// Field # 324
	ParentAssetProtectionAllowance string `json:",omitempty"`

	// Field # 325
	ParentContributionFromAssets string `json:",omitempty"`

	// Field # 326
	StudentNetWorth string `json:",omitempty"`

	// Field # 327
	StudentAssetProtectionAllowance string `json:",omitempty"`

	// Field # 328
	StudentContributionFromAssets string `json:",omitempty"`

	// Field # 329
	AssumedStudentFamilySize string `json:",omitempty"`

	// Field # 330
	AssumedParentFamilySize string `json:",omitempty"`

	// Field # 331
	StudentFirstNameCHVFlags string `json:",omitempty"`

	// Field # 332
	StudentMiddleNameCHVFlags string `json:",omitempty"`

	// Field # 333
	StudentLastNameCHVFlags string `json:",omitempty"`

	// Field # 334
	StudentSuffixCHVFlags string `json:",omitempty"`

	// Field # 335
	StudentDateOfBirthCHVFlags string `json:",omitempty"`

	// Field # 336
	StudentSSNCHVFlags string `json:",omitempty"`

	// Field # 337
	StudentITINCHVFlags string `json:",omitempty"`

	// Field # 338
	StudentPhoneNumberCHVFlags string `json:",omitempty"`

	// Field # 339
	StudentEmailAddressCHVFlags string `json:",omitempty"`

	// Field # 340
	StudentStreetAddressCHVFlags string `json:",omitempty"`

	// Field # 341
	StudentCityCHVFlags string `json:",omitempty"`

	// Field # 342
	StudentStateCHVFlags string `json:",omitempty"`

	// Field # 343
	StudentZipCodeCHVFlags string `json:",omitempty"`

	// Field # 344
	StudentCountryCHVFlags string `json:",omitempty"`

	// Field # 345
	StudentMaritalStatusCHVFlags string `json:",omitempty"`

	// Field # 346
	StudentGradeLevelInCollegeCHVFlags string `json:",omitempty"`

	// Field # 347
	StudentHasBachelorsDegreeCHVFlags string `json:",omitempty"`

	// Field # 348
	StudentPursuingTeacherCertificationCHVFlags string `json:",omitempty"`

	// Field # 349
	StudentActiveDutyCHVFlags string `json:",omitempty"`

	// Field # 350
	StudentVeteranCHVFlags string `json:",omitempty"`

	// Field # 351
	StudentChildOrOtherDependentsCHVFlags string `json:",omitempty"`

	// Field # 352
	StudentParentsDeceasedCHVFlags string `json:",omitempty"`

	// Field # 353
	StudentWardOfCourtCHVFlags string `json:",omitempty"`

	// Field # 354
	StudentInFosterCareCHVFlags string `json:",omitempty"`

	// Field # 355
	StudentEmancipatedMinorCHVFlags string `json:",omitempty"`

	// Field # 356
	StudentLegalGuardianshipCHVFlags string `json:",omitempty"`

	// Field # 357
	StudentPersonalCircumstancesNoneOfTheAboveCHVFlags string `json:",omitempty"`

	// Field # 358
	StudentUnaccompaniedHomelessOrIsUnaccompaniedCHVFlags string `json:",omitempty"`

	// Field # 359
	StudentUnaccompaniedAndHomelessGeneralCHVFlags string `json:",omitempty"`

	// Field # 360
	StudentUnaccompaniedAndHomelessHSCHVFlags string `json:",omitempty"`

	// Field # 361
	StudentUnaccompaniedAndHomelessTRIOCHVFlags string `json:",omitempty"`

	// Field # 362
	StudentUnaccompaniedAndHomelessFAACHVFlags string `json:",omitempty"`

	// Field # 363
	StudentHomelessnessNoneOfTheAboveCHVFlags string `json:",omitempty"`

	// Field # 364
	StudentHasUnusualCircumstanceCHVFlags string `json:",omitempty"`

	// Field # 365
	StudentUnsubOnlyCHVFlags string `json:",omitempty"`

	// Field # 366
	StudentUpdatedFamilySizeCHVFlags string `json:",omitempty"`

	// Field # 367
	StudentNumberInCollegeCorrectionCHVFlags string `json:",omitempty"`

	// Field # 368
	StudentCitizenshipStatusCorrectionCHVFlags string `json:",omitempty"`

	// Field # 369
	StudentANumberCHVFlags string `json:",omitempty"`

	// Field # 370
	StudentStateOfLegalResidenceCHVFlags string `json:",omitempty"`

	// Field # 371
	StudentLegalResidenceDateCHVFlags string `json:",omitempty"`

	// Field # 372
	StudentEitherParentAttendCollegeCHVFlags string `json:",omitempty"`

	// Field # 373
	StudentParentKilledInTheLineOfDutyCHVFlags string `json:",omitempty"`

	// Field # 374
	StudentHighSchoolCompletionStatusCHVFlags string `json:",omitempty"`

	// Field # 375
	StudentHighSchoolNameCHVFlags string `json:",omitempty"`

	// Field # 376
	StudentHighSchoolCityCHVFlags string `json:",omitempty"`

	// Field # 377
	StudentHighSchoolStateCHVFlags string `json:",omitempty"`

	// Field # 378
	StudentHighSchoolEquivalentDiplomaNameCHVFlags string `json:",omitempty"`

	// Field # 379
	StudentHighSchoolEquivalentDiplomaStateCHVFlags string `json:",omitempty"`

	// Field # 380
	StudentReceivedEITCCHVFlags string `json:",omitempty"`

	// Field # 381
	StudentReceivedFederalHousingAssistanceCHVFlags string `json:",omitempty"`

	// Field # 382
	StudentReceivedFreeReducedPriceLunchCHVFlags string `json:",omitempty"`

	// Field # 383
	StudentReceivedMedicaidCHVFlags string `json:",omitempty"`

	// Field # 384
	StudentReceivedRefundableCreditFor36BHealthPlanCHVFlags string `json:",omitempty"`

	// Field # 385
	StudentReceivedSNAPCHVFlags string `json:",omitempty"`

	// Field # 386
	StudentReceivedSupplementalSecurityIncomeCHVFlags string `json:",omitempty"`

	// Field # 387
	StudentReceivedTANFCHVFlags string `json:",omitempty"`

	// Field # 388
	StudentReceivedWICCHVFlags string `json:",omitempty"`

	// Field # 389
	StudentFederalBenefitsNoneOfTheAboveCHVFlags string `json:",omitempty"`

	// Field # 390
	StudentFiled1040Or1040NRCHVFlags string `json:",omitempty"`

	// Field # 391
	StudentFiledNonUSTaxReturnCHVFlags string `json:",omitempty"`

	// Field # 392
	StudentFiledJointReturnWithCurrentSpouseCHVFlags string `json:",omitempty"`

	// Field # 393
	StudentTaxReturnFilingStatusCHVFlags string `json:",omitempty"`

	// Field # 394
	StudentIncomeEarnedFromWorkCorrectionCHVFlags string `json:",omitempty"`

	// Field # 395
	StudentTaxExemptInterestIncomeCHVFlags string `json:",omitempty"`

	// Field # 396
	StudentUntaxedPortionsOfIRADistributionsCHVFlags string `json:",omitempty"`

	// Field # 397
	StudentIRARolloverCHVFlags string `json:",omitempty"`

	// Field # 398
	StudentUntaxedPortionsOfPensionsCHVFlags string `json:",omitempty"`

	// Field # 399
	StudentPensionRolloverCHVFlags string `json:",omitempty"`

	// Field # 400
	StudentAdjustedGrossIncomeCHVFlags string `json:",omitempty"`

	// Field # 401
	StudentIncomeTaxPaidCHVFlags string `json:",omitempty"`

	// Field # 402
	StudentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlags string `json:",omitempty"`

	// Field # 403
	StudentDeductiblePaymentsToIRAKeoghOtherCHVFlags string `json:",omitempty"`

	// Field # 404
	StudentEducationCreditsCHVFlags string `json:",omitempty"`

	// Field # 405
	StudentFiledScheduleABDEFHCHVFlags string `json:",omitempty"`

	// Field # 406
	StudentScheduleCAmountCHVFlags string `json:",omitempty"`

	// Field # 407
	StudentCollegeGrantAndScholarshipAidCHVFlags string `json:",omitempty"`

	// Field # 408
	StudentForeignEarnedIncomeExclusionCHVFlags string `json:",omitempty"`

	// Field # 409
	StudentChildSupportReceivedCHVFlags string `json:",omitempty"`

	// Field # 410
	StudentNetWorthOfBusinessesAndInvestmentFarmsCHVFlags string `json:",omitempty"`

	// Field # 411
	StudentNetWorthOfCurrentInvestmentsCHVFlags string `json:",omitempty"`

	// Field # 412
	StudentTotalOfCashSavingsAndCheckingCHVFlags string `json:",omitempty"`

	// Field # 413
	StudentCollege1CHVFlags string `json:",omitempty"`

	// Field # 414
	StudentCollege2CHVFlags string `json:",omitempty"`

	// Field # 415
	StudentCollege3CHVFlags string `json:",omitempty"`

	// Field # 416
	StudentCollege4CHVFlags string `json:",omitempty"`

	// Field # 417
	StudentCollege5CHVFlags string `json:",omitempty"`

	// Field # 418
	StudentCollege6CHVFlags string `json:",omitempty"`

	// Field # 419
	StudentCollege7CHVFlags string `json:",omitempty"`

	// Field # 420
	StudentCollege8CHVFlags string `json:",omitempty"`

	// Field # 421
	StudentCollege9CHVFlags string `json:",omitempty"`

	// Field # 422
	StudentCollege10CHVFlags string `json:",omitempty"`

	// Field # 423
	StudentCollege11CHVFlags string `json:",omitempty"`

	// Field # 424
	StudentCollege12CHVFlags string `json:",omitempty"`

	// Field # 425
	StudentCollege13CHVFlags string `json:",omitempty"`

	// Field # 426
	StudentCollege14CHVFlags string `json:",omitempty"`

	// Field # 427
	StudentCollege15CHVFlags string `json:",omitempty"`

	// Field # 428
	StudentCollege16CHVFlags string `json:",omitempty"`

	// Field # 429
	StudentCollege17CHVFlags string `json:",omitempty"`

	// Field # 430
	StudentCollege18CHVFlags string `json:",omitempty"`

	// Field # 431
	StudentCollege19CHVFlags string `json:",omitempty"`

	// Field # 432
	StudentCollege20CHVFlags string `json:",omitempty"`

	// Field # 433
	StudentConsentToRetrieveAndDiscloseFTICHVFlags string `json:",omitempty"`

	// Field # 434
	StudentSignatureCHVFlags string `json:",omitempty"`

	// Field # 435
	StudentSignatureDateCHVFlags string `json:",omitempty"`

	// Field # 436
	StudentSpouseFirstNameCHVFlags string `json:",omitempty"`

	// Field # 437
	StudentSpouseMiddleNameCHVFlags string `json:",omitempty"`

	// Field # 438
	StudentSpouseLastNameCHVFlags string `json:",omitempty"`

	// Field # 439
	StudentSpouseSuffixCHVFlags string `json:",omitempty"`

	// Field # 440
	StudentSpouseDateOfBirthCHVFlags string `json:",omitempty"`

	// Field # 441
	StudentSpouseSSNCHVFlags string `json:",omitempty"`

	// Field # 442
	StudentSpouseITINCHVFlags string `json:",omitempty"`

	// Field # 443
	StudentSpousePhoneNumberCHVFlags string `json:",omitempty"`

	// Field # 444
	StudentSpouseEmailAddressCHVFlags string `json:",omitempty"`

	// Field # 445
	StudentSpouseStreetAddressCHVFlags string `json:",omitempty"`

	// Field # 446
	StudentSpouseCityCHVFlags string `json:",omitempty"`

	// Field # 447
	StudentSpouseStateCHVFlags string `json:",omitempty"`

	// Field # 448
	StudentSpouseZipCodeCHVFlags string `json:",omitempty"`

	// Field # 449
	StudentSpouseCountryCHVFlags string `json:",omitempty"`

	// Field # 450
	StudentSpouseFiled1040Or1040NRCHVFlags string `json:",omitempty"`

	// Field # 451
	StudentSpouseFiledNonUSTaxReturnCHVFlags string `json:",omitempty"`

	// Field # 452
	StudentSpouseTaxReturnFilingStatusCHVFlags string `json:",omitempty"`

	// Field # 453
	StudentSpouseIncomeEarnedFromWorkCHVFlags string `json:",omitempty"`

	// Field # 454
	StudentSpouseTaxExemptInterestIncomeCHVFlags string `json:",omitempty"`

	// Field # 455
	StudentSpouseUntaxedPortionsOfIRADistributionsCHVFlags string `json:",omitempty"`

	// Field # 456
	StudentSpouseIRARolloverCHVFlags string `json:",omitempty"`

	// Field # 457
	StudentSpouseUntaxedPortionsOfPensionsCHVFlags string `json:",omitempty"`

	// Field # 458
	StudentSpousePensionRolloverCHVFlags string `json:",omitempty"`

	// Field # 459
	StudentSpouseAdjustedGrossIncomeCHVFlags string `json:",omitempty"`

	// Field # 460
	StudentSpouseIncomeTaxPaidCHVFlags string `json:",omitempty"`

	// Field # 461
	StudentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlags string `json:",omitempty"`

	// Field # 462
	StudentSpouseEducationCreditsCHVFlags string `json:",omitempty"`

	// Field # 463
	StudentSpouseFiledScheduleABDEFHCHVFlags string `json:",omitempty"`

	// Field # 464
	StudentSpouseScheduleCAmountCHVFlags string `json:",omitempty"`

	// Field # 465
	StudentSpouseForeignEarnedIncomeExclusionCHVFlags string `json:",omitempty"`

	// Field # 466
	StudentSpouseConsentToRetrieveAndDiscloseFTICHVFlags string `json:",omitempty"`

	// Field # 467
	StudentSpouseSignatureCHVFlags string `json:",omitempty"`

	// Field # 468
	StudentSpouseSignatureDateCHVFlags string `json:",omitempty"`

	// Field # 469
	ParentFirstNameCHVFlags string `json:",omitempty"`

	// Field # 470
	ParentMiddleNameCHVFlags string `json:",omitempty"`

	// Field # 471
	ParentLastNameCHVFlags string `json:",omitempty"`

	// Field # 472
	ParentSuffixCHVFlags string `json:",omitempty"`

	// Field # 473
	ParentDateOfBirthCHVFlags string `json:",omitempty"`

	// Field # 474
	ParentSSNCHVFlags string `json:",omitempty"`

	// Field # 475
	ParentITINCHVFlags string `json:",omitempty"`

	// Field # 476
	ParentPhoneNumberCHVFlags string `json:",omitempty"`

	// Field # 477
	ParentEmailAddressCHVFlags string `json:",omitempty"`

	// Field # 478
	ParentStreetAddressCHVFlags string `json:",omitempty"`

	// Field # 479
	ParentCityCHVFlags string `json:",omitempty"`

	// Field # 480
	ParentStateCHVFlags string `json:",omitempty"`

	// Field # 481
	ParentZipCodeCHVFlags string `json:",omitempty"`

	// Field # 482
	ParentCountryCHVFlags string `json:",omitempty"`

	// Field # 483
	ParentMaritalStatusCHVFlags string `json:",omitempty"`

	// Field # 484
	ParentStateOfLegalResidenceCHVFlags string `json:",omitempty"`

	// Field # 485
	ParentLegalResidenceDateCHVFlags string `json:",omitempty"`

	// Field # 486
	ParentUpdatedFamilySizeCHVFlags string `json:",omitempty"`

	// Field # 487
	ParentNumberInCollegeCHVFlags string `json:",omitempty"`

	// Field # 488
	ParentReceivedEITCCHVFlags string `json:",omitempty"`

	// Field # 489
	ParentReceivedFederalHousingAssistanceCHVFlags string `json:",omitempty"`

	// Field # 490
	ParentReceivedFreeReducedPriceLunchCHVFlags string `json:",omitempty"`

	// Field # 491
	ParentReceivedMedicaidCHVFlags string `json:",omitempty"`

	// Field # 492
	ParentReceivedRefundableCreditFor36BHealthPlanCHVFlags string `json:",omitempty"`

	// Field # 493
	ParentReceivedSNAPCHVFlags string `json:",omitempty"`

	// Field # 494
	ParentReceivedSupplementalSecurityIncomeCHVFlags string `json:",omitempty"`

	// Field # 495
	ParentReceivedTANFCHVFlags string `json:",omitempty"`

	// Field # 496
	ParentReceivedWICCHVFlags string `json:",omitempty"`

	// Field # 497
	ParentFederalBenefitsNoneOfTheAboveCHVFlags string `json:",omitempty"`

	// Field # 498
	ParentFiled1040Or1040NRCHVFlags string `json:",omitempty"`

	// Field # 499
	ParentFileNonUSTaxReturnCHVFlags string `json:",omitempty"`

	// Field # 500
	ParentFiledJointReturnWithCurrentSpouseCHVFlags string `json:",omitempty"`

	// Field # 501
	ParentTaxReturnFilingStatusCHVFlags string `json:",omitempty"`

	// Field # 502
	ParentIncomeEarnedFromWorkCHVFlags string `json:",omitempty"`

	// Field # 503
	ParentTaxExemptInterestIncomeCHVFlags string `json:",omitempty"`

	// Field # 504
	ParentUntaxedPortionsOfIRADistributionsCHVFlags string `json:",omitempty"`

	// Field # 505
	ParentIRARolloverCHVFlags string `json:",omitempty"`

	// Field # 506
	ParentUntaxedPortionsOfPensionsCHVFlags string `json:",omitempty"`

	// Field # 507
	ParentPensionRolloverCHVFlags string `json:",omitempty"`

	// Field # 508
	ParentAdjustedGrossIncomeCHVFlags string `json:",omitempty"`

	// Field # 509
	ParentIncomeTaxPaidCHVFlags string `json:",omitempty"`

	// Field # 510
	ParentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlags string `json:",omitempty"`

	// Field # 511
	ParentDeductiblePaymentsToIRAKeoghOtherCHVFlags string `json:",omitempty"`

	// Field # 512
	ParentEducationCreditsCHVFlags string `json:",omitempty"`

	// Field # 513
	ParentFiledScheduleABDEFHCHVFlags string `json:",omitempty"`

	// Field # 514
	ParentScheduleCAmountCHVFlags string `json:",omitempty"`

	// Field # 515
	ParentCollegeGrantAndScholarshipAidCHVFlags string `json:",omitempty"`

	// Field # 516
	ParentForeignEarnedIncomeExclusionCHVFlags string `json:",omitempty"`

	// Field # 517
	ParentChildSupportReceivedCHVFlags string `json:",omitempty"`

	// Field # 518
	ParentNetWorthOfCurrentInvestmentsCHVFlags string `json:",omitempty"`

	// Field # 519
	ParentTotalOfCashSavingsAndCheckingAccountsCHVFlags string `json:",omitempty"`

	// Field # 520
	ParentNetWorthOfBusinessesAndInvestmentFarmsCHVFlags string `json:",omitempty"`

	// Field # 521
	ParentConsentToRetrieveAndDiscloseFTICHVFlags string `json:",omitempty"`

	// Field # 522
	ParentSignatureCHVFlags string `json:",omitempty"`

	// Field # 523
	ParentSignatureDateCHVFlags string `json:",omitempty"`

	// Field # 524
	ParentSpouseFirstNameCHVFlags string `json:",omitempty"`

	// Field # 525
	ParentSpouseMiddleNameCHVFlags string `json:",omitempty"`

	// Field # 526
	ParentSpouseLastNameCHVFlags string `json:",omitempty"`

	// Field # 527
	ParentSpouseSuffixCHVFlags string `json:",omitempty"`

	// Field # 528
	ParentSpouseDateOfBirthCHVFlags string `json:",omitempty"`

	// Field # 529
	ParentSpouseSSNCHVFlags string `json:",omitempty"`

	// Field # 530
	ParentSpouseITINCHVFlags string `json:",omitempty"`

	// Field # 531
	ParentSpousePhoneNumberCHVFlags string `json:",omitempty"`

	// Field # 532
	ParentSpouseEmailAddressCHVFlags string `json:",omitempty"`

	// Field # 533
	ParentSpouseStreetAddressCHVFlags string `json:",omitempty"`

	// Field # 534
	ParentSpouseCityCHVFlags string `json:",omitempty"`

	// Field # 535
	ParentSpouseStateCHVFlags string `json:",omitempty"`

	// Field # 536
	ParentSpouseZipCodeCHVFlags string `json:",omitempty"`

	// Field # 537
	ParentSpouseCountryCHVFlags string `json:",omitempty"`

	// Field # 538
	ParentSpouseFiled1040Or1040NRCHVFlags string `json:",omitempty"`

	// Field # 539
	ParentSpouseFileNonUSTaxReturnCHVFlags string `json:",omitempty"`

	// Field # 540
	ParentSpouseTaxReturnFilingStatusCHVFlags string `json:",omitempty"`

	// Field # 541
	ParentSpouseIncomeEarnedFromWorkCHVFlags string `json:",omitempty"`

	// Field # 542
	ParentSpouseTaxExemptInterestIncomeCHVFlags string `json:",omitempty"`

	// Field # 543
	ParentSpouseUntaxedPortionsOfIRADistributionsCHVFlags string `json:",omitempty"`

	// Field # 544
	ParentSpouseIRARolloverCHVFlags string `json:",omitempty"`

	// Field # 545
	ParentSpouseUntaxedPortionsOfPensionsCHVFlags string `json:",omitempty"`

	// Field # 546
	ParentSpousePensionRolloverCHVFlags string `json:",omitempty"`

	// Field # 547
	ParentSpouseAdjustedGrossIncomeCHVFlags string `json:",omitempty"`

	// Field # 548
	ParentSpouseIncomeTaxPaidCHVFlags string `json:",omitempty"`

	// Field # 549
	ParentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlags string `json:",omitempty"`

	// Field # 550
	ParentSpouseEducationCreditsCHVFlags string `json:",omitempty"`

	// Field # 551
	ParentSpouseFiledScheduleABDEFHCHVFlags string `json:",omitempty"`

	// Field # 552
	ParentSpouseScheduleCAmountCHVFlags string `json:",omitempty"`

	// Field # 553
	ParentSpouseForeignEarnedIncomeExclusionCHVFlags string `json:",omitempty"`

	// Field # 554
	ParentSpouseConsentToRetrieveAndDiscloseFTICHVFlags string `json:",omitempty"`

	// Field # 555
	ParentSpouseSignatureCHVFlags string `json:",omitempty"`

	// Field # 556
	ParentSpouseSignatureDateCHVFlags string `json:",omitempty"`

	// Field # 557
	DHSPrimaryMatchStatus string `json:",omitempty"`

	// Field # 559
	DHSCaseNumber string `json:",omitempty"`

	// Field # 560
	NSLDSMatchStatus string `json:",omitempty"`

	// Field # 561
	NSLDSPostscreeningReasonCode string `json:",omitempty"`

	// Field # 562
	StudentSSACitizenshipFlagResults string `json:",omitempty"`

	// Field # 563
	StudentSSAMatchStatus string `json:",omitempty"`

	// Field # 564
	StudentSpouseSSAMatchStatus string `json:",omitempty"`

	// Field # 565
	ParentSSAMatchStatus string `json:",omitempty"`

	// Field # 566
	ParentSpouseOrPartnerSSAMatchStatus string `json:",omitempty"`

	// Field # 567
	VAMatchFlag string `json:",omitempty"`

	// Field # 568
	CommentCodes string `json:",omitempty"`

	// Field # 569
	DrugAbuseHoldIndicator string `json:",omitempty"`

	// Field # 570
	GraduateFlag string `json:",omitempty"`

	// Field # 571
	PellGrantEligibilityFlag string `json:",omitempty"`

	// Field # 572
	ReprocessedReasonCode string `json:",omitempty"`

	// Field # 573
	FPSCFlag string `json:",omitempty"`

	// Field # 574
	FPSCChangeFlag string `json:",omitempty"`

	// Field # 575
	ElectronicFederalSchoolCodeIndicator string `json:",omitempty"`

	// Field # 576
	RejectReasonCodes string `json:",omitempty"`

	// Field # 577
	ElectronicTransactionIndicatorFlag string `json:",omitempty"`

	// Field # 578
	StudentLastNameSSNChangeFlag string `json:",omitempty"`

	// Field # 579
	HighSchoolCode string `json:",omitempty"`

	// Field # 580
	VerificationSelectionChangeFlag string `json:",omitempty"`

	// Field # 581
	UseUserProvidedDataOnly string `json:",omitempty"`

	// Field # 583
	NSLDSPellOverpaymentFlag string `json:",omitempty"`

	// Field # 584
	NSLDSPellOverpaymentContact string `json:",omitempty"`

	// Field # 585
	NSLDSFSEOGOverpaymentFlag string `json:",omitempty"`

	// Field # 586
	NSLDSFSEOGOverpaymentContact string `json:",omitempty"`

	// Field # 587
	NSLDSPerkinsOverpaymentFlag string `json:",omitempty"`

	// Field # 588
	NSLDSPerkinsOverpaymentContact string `json:",omitempty"`

	// Field # 589
	NSLDSTEACHGrantOverpaymentFlag string `json:",omitempty"`

	// Field # 590
	NSLDSTEACHGrantOverpaymentContact string `json:",omitempty"`

	// Field # 591
	NSLDSIraqAndAfghanistanServiceGrantOverpaymentFlag string `json:",omitempty"`

	// Field # 592
	NSLDSIraqAndAfghanistanServiceGrantOverpaymentContact string `json:",omitempty"`

	// Field # 593
	NSLDSDefaultedLoanFlag string `json:",omitempty"`

	// Field # 594
	NSLDSDischargedLoanFlag string `json:",omitempty"`

	// Field # 595
	NSLDSFraudLoanFlag string `json:",omitempty"`

	// Field # 596
	NSLDSSatisfactoryArrangementsFlag string `json:",omitempty"`

	// Field # 597
	NSLDSActiveBankruptcyFlag string `json:",omitempty"`

	// Field # 598
	NSLDSTEACHGrantConvertedToLoanFlag string `json:",omitempty"`

	// Field # 599
	NSLDSAggregateSubsidizedOutstandingPrincipalBalance string `json:",omitempty"`

	// Field # 600
	NSLDSAggregateUnsubsidizedOutstandingPrincipalBalance string `json:",omitempty"`

	// Field # 601
	NSLDSAggregateCombinedOutstandingPrincipalBalance string `json:",omitempty"`

	// Field # 602
	NSLDSAggregateUnallocConsolidatedOutstandingPrincipalBalance string `json:",omitempty"`

	// Field # 603
	NSLDSAggregateTEACHLoanPrincipalBalance string `json:",omitempty"`

	// Field # 604
	NSLDSAggregateSubsidizedPendingDisbursement string `json:",omitempty"`

	// Field # 605
	NSLDSAggregateUnsubsidizedPendingDisbursement string `json:",omitempty"`

	// Field # 606
	NSLDSAggregateCombinedPendingDisbursement string `json:",omitempty"`

	// Field # 607
	NSLDSAggregateSubsidizedTotal string `json:",omitempty"`

	// Field # 608
	NSLDSAggregateUnsubsidizedTotal string `json:",omitempty"`

	// Field # 609
	NSLDSAggregateCombinedTotal string `json:",omitempty"`

	// Field # 610
	NSLDSUnallocatedConsolidatedTotal string `json:",omitempty"`

	// Field # 611
	NSLDSTEACHLoanTotal string `json:",omitempty"`

	// Field # 612
	NSLDSPerkinsTotalDisbursements string `json:",omitempty"`

	// Field # 613
	NSLDSPerkinsCurrentYearDisbursementAmount string `json:",omitempty"`

	// Field # 614
	NSLDSAggregateTEACHGrantUndergraduateDisbursedTotal string `json:",omitempty"`

	// Field # 615
	NSLDSAggregateTEACHGraduateDisbursementAmount string `json:",omitempty"`

	// Field # 616
	NSLDSDefaultedLoanChangeFlag string `json:",omitempty"`

	// Field # 617
	NSLDSFraudLoanChangeFlag string `json:",omitempty"`

	// Field # 618
	NSLDSDischargedLoanChangeFlag string `json:",omitempty"`

	// Field # 619
	NSLDSLoanSatisfactoryRepaymentChangeFlag string `json:",omitempty"`

	// Field # 620
	NSLDSActiveBankruptcyChangeFlag string `json:",omitempty"`

	// Field # 621
	NSLDSTEACHGrantToLoanConversionChangeFlag string `json:",omitempty"`

	// Field # 622
	NSLDSOverpaymentsChangeFlag string `json:",omitempty"`

	// Field # 623
	NSLDSAggregateLoanChangeFlag string `json:",omitempty"`

	// Field # 624
	NSLDSPerkinsLoanChangeFlag string `json:",omitempty"`

	// Field # 625
	NSLDSPellPaymentChangeFlag string `json:",omitempty"`

	// Field # 626
	NSLDSTEACHGrantChangeFlag string `json:",omitempty"`

	// Field # 627
	NSLDSAdditionalPellFlag string `json:",omitempty"`

	// Field # 628
	NSLDSAdditionalLoansFlag string `json:",omitempty"`

	// Field # 629
	NSLDSAdditionalTEACHGrantFlag string `json:",omitempty"`

	// Field # 630
	NSLDSDirectLoanMasterPromNoteFlag string `json:",omitempty"`

	// Field # 631
	NSLDSDirectLoanPLUSMasterPromNoteFlag string `json:",omitempty"`

	// Field # 632
	NSLDSDirectLoanGraduatePLUSMasterPromNoteFlag string `json:",omitempty"`

	// Field # 633
	NSLDSUndergraduateSubsidizedLoanLimitFlag string `json:",omitempty"`

	// Field # 634
	NSLDSUndergraduateCombinedLoanLimitFlag string `json:",omitempty"`

	// Field # 635
	NSLDSGraduateSubsidizedLoanLimitFlag string `json:",omitempty"`

	// Field # 636
	NSLDSGraduateCombinedLoanLimitFlag string `json:",omitempty"`

	// Field # 637
	NSLDSPellLifetimeLimitFlag string `json:",omitempty"`

	// Field # 638
	NSLDSPellLifetimeEligibilityUsed string `json:",omitempty"`

	// Field # 639
	NSLDSSULAFlag string `json:",omitempty"`

	// Field # 640
	NSLDSSubsidizedLimitEligibilityUsed string `json:",omitempty"`

	// Field # 641
	NSLDSUnusualEnrollmentHistoryFlag string `json:",omitempty"`

	// Field # 643
	NSLDSPellSequenceNumber1 string `json:",omitempty"`

	// Field # 644
	NSLDSPellVerificationFlag1 string `json:",omitempty"`

	// Field # 645
	NSLDSSAI1 string `json:",omitempty"`

	// Field # 646
	NSLDSPellSchoolCode1 string `json:",omitempty"`

	// Field # 647
	NSLDSPellTransactionNumber1 string `json:",omitempty"`

	// Field # 648
	NSLDSPellDisbursementDate1 *time.Time `json:",omitempty"`

	// Field # 649
	NSLDSPellScheduledAmount1 string `json:",omitempty"`

	// Field # 650
	NSLDSPellAmountPaidToDate1 *time.Time `json:",omitempty"`

	// Field # 651
	NSLDSPellPercentEligibilityUsedDecimal1 string `json:",omitempty"`

	// Field # 652
	NSLDSPellAwardAmount1 string `json:",omitempty"`

	// Field # 653
	NSLDSAdditionalEligibilityIndicator1 string `json:",omitempty"`

	// Field # 655
	NSLDSPellSequenceNumber2 string `json:",omitempty"`

	// Field # 656
	NSLDSPellVerificationFlag2 string `json:",omitempty"`

	// Field # 657
	NSLDSSAI2 string `json:",omitempty"`

	// Field # 658
	NSLDSPellSchoolCode2 string `json:",omitempty"`

	// Field # 659
	NSLDSPellTransactionNumber2 string `json:",omitempty"`

	// Field # 660
	NSLDSPellLastDisbursementDate2 *time.Time `json:",omitempty"`

	// Field # 661
	NSLDSPellScheduledAmount2 string `json:",omitempty"`

	// Field # 662
	NSLDSPellAmountPaidToDate2 *time.Time `json:",omitempty"`

	// Field # 663
	NSLDSPellPercentEligibilityUsedDecimal2 string `json:",omitempty"`

	// Field # 664
	NSLDSPellAwardAmount2 string `json:",omitempty"`

	// Field # 665
	NSLDSAdditionalEligibilityIndicator2 string `json:",omitempty"`

	// Field # 667
	NSLDSPellSequenceNumber3 string `json:",omitempty"`

	// Field # 668
	NSLDSPellVerificationFlag3 string `json:",omitempty"`

	// Field # 669
	NSLDSSAI3 string `json:",omitempty"`

	// Field # 670
	NSLDSPellSchoolCode3 string `json:",omitempty"`

	// Field # 671
	NSLDSPellTransactionNumber3 string `json:",omitempty"`

	// Field # 672
	NSLDSPellLastDisbursementDate3 *time.Time `json:",omitempty"`

	// Field # 673
	NSLDSPellScheduledAmount3 string `json:",omitempty"`

	// Field # 674
	NSLDSPellAmountPaidToDate3 *time.Time `json:",omitempty"`

	// Field # 675
	NSLDSPellPercentEligibilityUsedDecimal3 string `json:",omitempty"`

	// Field # 676
	NSLDSPellAwardAmount3 string `json:",omitempty"`

	// Field # 677
	NSLDSAdditionalEligibilityIndicator3 string `json:",omitempty"`

	// Field # 679
	NSLDSTEACHGrantSequence1 string `json:",omitempty"`

	// Field # 680
	NSLDSTEACHGrantSchoolCode1 string `json:",omitempty"`

	// Field # 681
	NSLDSTEACHGrantTransactionNumber1 string `json:",omitempty"`

	// Field # 682
	NSLDSTEACHGrantLastDisbursementDate1 *time.Time `json:",omitempty"`

	// Field # 683
	NSLDSTEACHGrantScheduledAmount1 string `json:",omitempty"`

	// Field # 684
	NSLDSTEACHGrantAmountPaidToDate1 *time.Time `json:",omitempty"`

	// Field # 685
	NSLDSTEACHGrantAwardAmount1 string `json:",omitempty"`

	// Field # 686
	NSLDSTEACHGrantAcademicYearLevel1 string `json:",omitempty"`

	// Field # 687
	NSLDSTEACHGrantAwardYear1 string `json:",omitempty"`

	// Field # 688
	NSLDSTEACHGrantLoanConversionFlag1 string `json:",omitempty"`

	// Field # 689
	NSLDSTEACHGrantDischargeCode1 string `json:",omitempty"`

	// Field # 690
	NSLDSTEACHGrantDischargeAmount1 string `json:",omitempty"`

	// Field # 691
	NSLDSTEACHGrantAdjustedDisbursement1 string `json:",omitempty"`

	// Field # 693
	NSLDSTEACHGrantSequence2 string `json:",omitempty"`

	// Field # 694
	NSLDSTEACHGrantSchoolCode2 string `json:",omitempty"`

	// Field # 695
	NSLDSTEACHGrantTransactionNumber2 string `json:",omitempty"`

	// Field # 696
	NSLDSTEACHGrantLastDisbursementDate2 *time.Time `json:",omitempty"`

	// Field # 697
	NSLDSTEACHGrantScheduledAmount2 string `json:",omitempty"`

	// Field # 698
	NSLDSTEACHGrantAmountPaidToDate2 *time.Time `json:",omitempty"`

	// Field # 699
	NSLDSTEACHGrantAwardAmount2 string `json:",omitempty"`

	// Field # 700
	NSLDSTEACHGrantAcademicYearLevel2 string `json:",omitempty"`

	// Field # 701
	NSLDSTEACHGrantAwardYear2 string `json:",omitempty"`

	// Field # 702
	NSLDSTEACHGrantLoanConversionFlag2 string `json:",omitempty"`

	// Field # 703
	NSLDSTEACHGrantDischargeCode2 string `json:",omitempty"`

	// Field # 704
	NSLDSTEACHGrantDischargeAmount2 string `json:",omitempty"`

	// Field # 705
	NSLDSTEACHGrantAdjustedDisbursement2 string `json:",omitempty"`

	// Field # 707
	NSLDSTEACHGrantSequence3 string `json:",omitempty"`

	// Field # 708
	NSLDSTEACHGrantSchoolCode3 string `json:",omitempty"`

	// Field # 709
	NSLDSTEACHGrantTransactionNumber3 string `json:",omitempty"`

	// Field # 710
	NSLDSTEACHGrantLastDisbursementDate3 *time.Time `json:",omitempty"`

	// Field # 711
	NSLDSTEACHGrantScheduledAmount3 string `json:",omitempty"`

	// Field # 712
	NSLDSTEACHGrantAmountPaidToDate3 *time.Time `json:",omitempty"`

	// Field # 713
	NSLDSTEACHGrantAwardAmount3 string `json:",omitempty"`

	// Field # 714
	NSLDSTEACHGrantAcademicYearLevel3 string `json:",omitempty"`

	// Field # 715
	NSLDSTEACHGrantAwardYear3 string `json:",omitempty"`

	// Field # 716
	NSLDSTEACHGrantLoanConversionFlag3 string `json:",omitempty"`

	// Field # 717
	NSLDSTEACHGrantDischargeCode3 string `json:",omitempty"`

	// Field # 718
	NSLDSTEACHGrantDischargeAmount3 string `json:",omitempty"`

	// Field # 719
	NSLDSTEACHGrantAdjustedDisbursement3 string `json:",omitempty"`

	// Field # 721
	NSLDSLoanSequenceNumber1 string `json:",omitempty"`

	// Field # 722
	NSLDSLoanDefaultedRecentIndicator1 string `json:",omitempty"`

	// Field # 723
	NSLDSLoanChangeFlag1 string `json:",omitempty"`

	// Field # 724
	NSLDSLoanTypeCode1 string `json:",omitempty"`

	// Field # 725
	NSLDSLoanNetAmount1 string `json:",omitempty"`

	// Field # 726
	NSLDSLoanCurrentStatusCode1 string `json:",omitempty"`

	// Field # 727
	NSLDSLoanCurrentStatusDate1 *time.Time `json:",omitempty"`

	// Field # 728
	NSLDSLoanOutstandingPrincipalBalance1 string `json:",omitempty"`

	// Field # 729
	NSLDSLoanOutstandingPrincipalBalanceDate1 *time.Time `json:",omitempty"`

	// Field # 730
	NSLDSLoanPeriodBeginDate1 *time.Time `json:",omitempty"`

	// Field # 731
	NSLDSLoanPeriodEndDate1 *time.Time `json:",omitempty"`

	// Field # 732
	NSLDSLoanGuarantyAgencyCode1 string `json:",omitempty"`

	// Field # 733
	NSLDSLoanContactType1 string `json:",omitempty"`

	// Field # 734
	NSLDSLoanSchoolCode1 string `json:",omitempty"`

	// Field # 735
	NSLDSLoanContactCode1 string `json:",omitempty"`

	// Field # 736
	NSLDSLoanGradeLevel1 string `json:",omitempty"`

	// Field # 737
	NSLDSLoanAdditionalUnsubsidizedFlag1 string `json:",omitempty"`

	// Field # 738
	NSLDSLoanCapitalizedInterestFlag1 string `json:",omitempty"`

	// Field # 739
	NSLDSLoanDisbursementAmount1 string `json:",omitempty"`

	// Field # 740
	NSLDSLoanDisbursementDate1 *time.Time `json:",omitempty"`

	// Field # 741
	NSLDSLoanConfirmedLoanSubsidyStatus1 string `json:",omitempty"`

	// Field # 742
	NSLDSLoanConfirmedLoanSubsidyStatusDate1 *time.Time `json:",omitempty"`

	// Field # 744
	NSLDSLoanSequenceNumber2 string `json:",omitempty"`

	// Field # 745
	NSLDSLoanDefaultedRecentIndicator2 string `json:",omitempty"`

	// Field # 746
	NSLDSLoanChangeFlag2 string `json:",omitempty"`

	// Field # 747
	NSLDSLoanTypeCode2 string `json:",omitempty"`

	// Field # 748
	NSLDSLoanNetAmount2 string `json:",omitempty"`

	// Field # 749
	NSLDSLoanCurrentStatusCode2 string `json:",omitempty"`

	// Field # 750
	NSLDSLoanCurrentStatusDate2 *time.Time `json:",omitempty"`

	// Field # 751
	NSLDSLoanOutstandingPrincipalBalance2 string `json:",omitempty"`

	// Field # 752
	NSLDSLoanOutstandingPrincipalBalanceDate2 *time.Time `json:",omitempty"`

	// Field # 753
	NSLDSLoanPeriodBeginDate2 *time.Time `json:",omitempty"`

	// Field # 754
	NSLDSLoanPeriodEndDate2 *time.Time `json:",omitempty"`

	// Field # 755
	NSLDSLoanGuarantyAgencyCode2 string `json:",omitempty"`

	// Field # 756
	NSLDSLoanContactType2 string `json:",omitempty"`

	// Field # 757
	NSLDSLoanSchoolCode2 string `json:",omitempty"`

	// Field # 758
	NSLDSLoanContactCode2 string `json:",omitempty"`

	// Field # 759
	NSLDSLoanGradeLevel2 string `json:",omitempty"`

	// Field # 760
	NSLDSLoanAdditionalUnsubsidizedFlag2 string `json:",omitempty"`

	// Field # 761
	NSLDSLoanCapitalizedInterestFlag2 string `json:",omitempty"`

	// Field # 762
	NSLDSLoanDisbursementAmount2 string `json:",omitempty"`

	// Field # 763
	NSLDSLoanDisbursementDate2 *time.Time `json:",omitempty"`

	// Field # 764
	NSLDSLoanConfirmedLoanSubsidyStatus2 string `json:",omitempty"`

	// Field # 765
	NSLDSLoanConfirmedLoanSubsidyStatusDate2 *time.Time `json:",omitempty"`

	// Field # 767
	NSLDSLoanSequenceNumber3 string `json:",omitempty"`

	// Field # 768
	NSLDSLoanDefaultedRecentIndicator3 string `json:",omitempty"`

	// Field # 769
	NSLDSLoanChangeFlag3 string `json:",omitempty"`

	// Field # 770
	NSLDSLoanTypeCode3 string `json:",omitempty"`

	// Field # 771
	NSLDSLoanNetAmount3 string `json:",omitempty"`

	// Field # 772
	NSLDSLoanCurrentStatusCode3 string `json:",omitempty"`

	// Field # 773
	NSLDSLoanCurrentStatusDate3 *time.Time `json:",omitempty"`

	// Field # 774
	NSLDSLoanOutstandingPrincipalBalance3 string `json:",omitempty"`

	// Field # 775
	NSLDSLoanOutstandingPrincipalBalanceDate3 *time.Time `json:",omitempty"`

	// Field # 776
	NSLDSLoanPeriodBeginDate3 *time.Time `json:",omitempty"`

	// Field # 777
	NSLDSLoanPeriodEndDate3 *time.Time `json:",omitempty"`

	// Field # 778
	NSLDSLoanGuarantyAgencyCode3 string `json:",omitempty"`

	// Field # 779
	NSLDSLoanContactType3 string `json:",omitempty"`

	// Field # 780
	NSLDSLoanSchoolCode3 string `json:",omitempty"`

	// Field # 781
	NSLDSLoanContactCode3 string `json:",omitempty"`

	// Field # 782
	NSLDSLoanGradeLevel3 string `json:",omitempty"`

	// Field # 783
	NSLDSLoanAdditionalUnsubsidizedFlag3 string `json:",omitempty"`

	// Field # 784
	NSLDSLoanCapitalizedInterestFlag3 string `json:",omitempty"`

	// Field # 785
	NSLDSLoanDisbursementAmount3 string `json:",omitempty"`

	// Field # 786
	NSLDSLoanDisbursementDate3 *time.Time `json:",omitempty"`

	// Field # 787
	NSLDSLoanConfirmedLoanSubsidyStatus3 string `json:",omitempty"`

	// Field # 788
	NSLDSLoanConfirmedLoanSubsidyStatusDate3 *time.Time `json:",omitempty"`

	// Field # 790
	NSLDSLoanSequenceNumber4 string `json:",omitempty"`

	// Field # 791
	NSLDSLoanDefaultedRecentIndicator4 string `json:",omitempty"`

	// Field # 792
	NSLDSLoanChangeFlag4 string `json:",omitempty"`

	// Field # 793
	NSLDSLoanTypeCode4 string `json:",omitempty"`

	// Field # 794
	NSLDSLoanNetAmount4 string `json:",omitempty"`

	// Field # 795
	NSLDSLoanCurrentStatusCode4 string `json:",omitempty"`

	// Field # 796
	NSLDSLoanCurrentStatusDate4 *time.Time `json:",omitempty"`

	// Field # 797
	NSLDSLoanOutstandingPrincipalBalance4 string `json:",omitempty"`

	// Field # 798
	NSLDSLoanOutstandingPrincipalBalanceDate4 *time.Time `json:",omitempty"`

	// Field # 799
	NSLDSLoanPeriodBeginDate4 *time.Time `json:",omitempty"`

	// Field # 800
	NSLDSLoanPeriodEndDate4 *time.Time `json:",omitempty"`

	// Field # 801
	NSLDSLoanGuarantyAgencyCode4 string `json:",omitempty"`

	// Field # 802
	NSLDSLoanContactType4 string `json:",omitempty"`

	// Field # 803
	NSLDSLoanSchoolCode4 string `json:",omitempty"`

	// Field # 804
	NSLDSLoanContactCode4 string `json:",omitempty"`

	// Field # 805
	NSLDSLoanGradeLevel4 string `json:",omitempty"`

	// Field # 806
	NSLDSLoanAdditionalUnsubsidizedFlag4 string `json:",omitempty"`

	// Field # 807
	NSLDSLoanCapitalizedInterestFlag4 string `json:",omitempty"`

	// Field # 808
	NSLDSLoanDisbursementAmount4 string `json:",omitempty"`

	// Field # 809
	NSLDSLoanDisbursementDate4 *time.Time `json:",omitempty"`

	// Field # 810
	NSLDSLoanConfirmedLoanSubsidyStatus4 string `json:",omitempty"`

	// Field # 811
	NSLDSLoanSubsidyStatusDate4 *time.Time `json:",omitempty"`

	// Field # 813
	NSLDSLoanSequenceNumber5 string `json:",omitempty"`

	// Field # 814
	NSLDSLoanDefaultedRecentIndicator5 string `json:",omitempty"`

	// Field # 815
	NSLDSLoanChangeFlag5 string `json:",omitempty"`

	// Field # 816
	NSLDSLoanTypeCode5 string `json:",omitempty"`

	// Field # 817
	NSLDSLoanNetAmount5 string `json:",omitempty"`

	// Field # 818
	NSLDSLoanCurrentStatusCode5 string `json:",omitempty"`

	// Field # 819
	NSLDSLoanCurrentStatusDate5 *time.Time `json:",omitempty"`

	// Field # 820
	NSLDSLoanOutstandingPrincipalBalance5 string `json:",omitempty"`

	// Field # 821
	NSLDSLoanOutstandingPrincipalBalanceDate5 *time.Time `json:",omitempty"`

	// Field # 822
	NSLDSLoanPeriodBeginDate5 *time.Time `json:",omitempty"`

	// Field # 823
	NSLDSLoanPeriodEndDate5 *time.Time `json:",omitempty"`

	// Field # 824
	NSLDSLoanGuarantyAgencyCode5 string `json:",omitempty"`

	// Field # 825
	NSLDSLoanContactType5 string `json:",omitempty"`

	// Field # 826
	NSLDSLoanSchoolCode5 string `json:",omitempty"`

	// Field # 827
	NSLDSLoanContactCode5 string `json:",omitempty"`

	// Field # 828
	NSLDSLoanGradeLevel5 string `json:",omitempty"`

	// Field # 829
	NSLDSLoanAdditionalUnsubsidizedFlag5 string `json:",omitempty"`

	// Field # 830
	NSLDSLoanCapitalizedInterestFlag5 string `json:",omitempty"`

	// Field # 831
	NSLDSLoanDisbursementAmount5 string `json:",omitempty"`

	// Field # 832
	NSLDSLoanDisbursementDate5 *time.Time `json:",omitempty"`

	// Field # 833
	NSLDSLoanConfirmedLoanSubsidyStatus5 string `json:",omitempty"`

	// Field # 834
	NSLDSLoanSubsidyStatusDate5 *time.Time `json:",omitempty"`

	// Field # 836
	NSLDSLoanSequenceNumber6 string `json:",omitempty"`

	// Field # 837
	NSLDSLoanDefaultedRecentIndicator6 string `json:",omitempty"`

	// Field # 838
	NSLDSLoanChangeFlag6 string `json:",omitempty"`

	// Field # 839
	NSLDSLoanTypeCode6 string `json:",omitempty"`

	// Field # 840
	NSLDSLoanNetAmount6 string `json:",omitempty"`

	// Field # 841
	NSLDSLoanCurrentStatusCode6 string `json:",omitempty"`

	// Field # 842
	NSLDSLoanCurrentStatusDate6 *time.Time `json:",omitempty"`

	// Field # 843
	NSLDSLoanOutstandingPrincipalBalance6 string `json:",omitempty"`

	// Field # 844
	NSLDSLoanOutstandingPrincipalBalanceDate6 *time.Time `json:",omitempty"`

	// Field # 845
	NSLDSLoanPeriodBeginDate6 *time.Time `json:",omitempty"`

	// Field # 846
	NSLDSLoanPeriodEndDate6 *time.Time `json:",omitempty"`

	// Field # 847
	NSLDSLoanGuarantyAgencyCode6 string `json:",omitempty"`

	// Field # 848
	NSLDSLoanContactType6 string `json:",omitempty"`

	// Field # 849
	NSLDSLoanSchoolCode6 string `json:",omitempty"`

	// Field # 850
	NSLDSLoanContactCode6 string `json:",omitempty"`

	// Field # 851
	NSLDSLoanGradeLevel6 string `json:",omitempty"`

	// Field # 852
	NSLDSLoanAdditionalUnsubsidizedFlag6 string `json:",omitempty"`

	// Field # 853
	NSLDSLoanCapitalizedInterestFlag6 string `json:",omitempty"`

	// Field # 854
	NSLDSLoanDisbursementAmount6 string `json:",omitempty"`

	// Field # 855
	NSLDSLoanDisbursementDate6 *time.Time `json:",omitempty"`

	// Field # 856
	NSLDSLoanConfirmedLoanSubsidyStatus6 string `json:",omitempty"`

	// Field # 857
	NSLDSLoanSubsidyStatusDate6 *time.Time `json:",omitempty"`

	// Field # 861
	//FTILabelStart string `json:",omitempty"`

	// Field # 862
	StudentFTIMReturnedTaxYear string `json:",omitempty"`

	// Field # 863
	StudentFTIMFilingStatusCode string `json:",omitempty"`

	// Field # 864
	StudentFTIMAdjustedGrossIncome string `json:",omitempty"`

	// Field # 865
	StudentFTIMNumberOfExemptions string `json:",omitempty"`

	// Field # 866
	StudentFTIMNumberOfDependents string `json:",omitempty"`

	// Field # 867
	StudentFTIMTotalIncomeEarnedAmount string `json:",omitempty"`

	// Field # 868
	StudentFTIMIncomeTaxPaid string `json:",omitempty"`

	// Field # 869
	StudentFTIMEducationCredits string `json:",omitempty"`

	// Field # 870
	StudentFTIMUntaxedIRADistributions string `json:",omitempty"`

	// Field # 871
	StudentFTIMIRADeductibleAndPayments string `json:",omitempty"`

	// Field # 872
	StudentFTIMTaxExemptInterest string `json:",omitempty"`

	// Field # 873
	StudentFTIMUntaxedPensionsAmount string `json:",omitempty"`

	// Field # 874
	StudentFTIMScheduleCNetProfitLoss string `json:",omitempty"`

	// Field # 875
	StudentFTIMScheduleAIndicator string `json:",omitempty"`

	// Field # 876
	StudentFTIMScheduleBIndicator string `json:",omitempty"`

	// Field # 877
	StudentFTIMScheduleDIndicator string `json:",omitempty"`

	// Field # 878
	StudentFTIMScheduleEIndicator string `json:",omitempty"`

	// Field # 879
	StudentFTIMScheduleFIndicator string `json:",omitempty"`

	// Field # 880
	StudentFTIMScheduleHIndicator string `json:",omitempty"`

	// Field # 881
	StudentFTIMIRSResponseCode string `json:",omitempty"`

	// Field # 882
	StudentFTIMSpouseReturnedTaxYear string `json:",omitempty"`

	// Field # 883
	StudentFTIMSpouseFilingStatusCode string `json:",omitempty"`

	// Field # 884
	StudentFTIMSpouseAdjustedGrossIncome string `json:",omitempty"`

	// Field # 885
	StudentFTIMSpouseNumberOfExemptions string `json:",omitempty"`

	// Field # 886
	StudentFTIMSpouseNumberOfDependents string `json:",omitempty"`

	// Field # 887
	StudentFTIMSpouseTotalIncomeEarnedAmount string `json:",omitempty"`

	// Field # 888
	StudentFTIMSpouseIncomeTaxPaid string `json:",omitempty"`

	// Field # 889
	StudentFTIMSpouseEducationCredits string `json:",omitempty"`

	// Field # 890
	StudentFTIMSpouseUntaxedIRADistributions string `json:",omitempty"`

	// Field # 891
	StudentFTIMSpouseIRADeductibleAndPayments string `json:",omitempty"`

	// Field # 892
	StudentFTIMSpouseTaxExemptInterest string `json:",omitempty"`

	// Field # 893
	StudentFTIMSpouseUntaxedPensionsAmount string `json:",omitempty"`

	// Field # 894
	StudentFTIMSpouseScheduleCNetProfitLoss string `json:",omitempty"`

	// Field # 895
	StudentFTIMSpouseScheduleAIndicator string `json:",omitempty"`

	// Field # 896
	StudentFTIMSpouseScheduleBIndicator string `json:",omitempty"`

	// Field # 897
	StudentFTIMSpouseScheduleDIndicator string `json:",omitempty"`

	// Field # 898
	StudentFTIMSpouseScheduleEIndicator string `json:",omitempty"`

	// Field # 899
	StudentFTIMSpouseScheduleFIndicator string `json:",omitempty"`

	// Field # 900
	StudentFTIMSpouseScheduleHIndicator string `json:",omitempty"`

	// Field # 901
	StudentFTIMSpouseIRSResponseCode string `json:",omitempty"`

	// Field # 902
	ParentFTIMReturnedTaxYear string `json:",omitempty"`

	// Field # 903
	ParentFTIMFilingStatusCode string `json:",omitempty"`

	// Field # 904
	ParentFTIMAdjustedGrossIncome string `json:",omitempty"`

	// Field # 905
	ParentFTIMNumberOfExemptions string `json:",omitempty"`

	// Field # 906
	ParentFTIMNumberOfDependents string `json:",omitempty"`

	// Field # 907
	ParentFTIMTotalIncomeEarnedAmount string `json:",omitempty"`

	// Field # 908
	ParentFTIMIncomeTaxPaid string `json:",omitempty"`

	// Field # 909
	ParentFTIMEducationCredits string `json:",omitempty"`

	// Field # 910
	ParentFTIMUntaxedIRADistributions string `json:",omitempty"`

	// Field # 911
	ParentFTIMIRADeductibleAndPayments string `json:",omitempty"`

	// Field # 912
	ParentFTIMTaxExemptInterest string `json:",omitempty"`

	// Field # 913
	ParentFTIMUntaxedPensionsAmount string `json:",omitempty"`

	// Field # 914
	ParentFTIMScheduleCNetProfitLoss string `json:",omitempty"`

	// Field # 915
	ParentFTIMScheduleAIndicator string `json:",omitempty"`

	// Field # 916
	ParentFTIMScheduleBIndicator string `json:",omitempty"`

	// Field # 917
	ParentFTIMScheduleDIndicator string `json:",omitempty"`

	// Field # 918
	ParentFTIMScheduleEIndicator string `json:",omitempty"`

	// Field # 919
	ParentFTIMScheduleFIndicator string `json:",omitempty"`

	// Field # 920
	ParentFTIMScheduleHIndicator string `json:",omitempty"`

	// Field # 921
	ParentFTIMIRSResponseCode string `json:",omitempty"`

	// Field # 922
	ParentFTIMSpouseReturnedTaxYear string `json:",omitempty"`

	// Field # 923
	ParentFTIMSpouseFilingStatusCode string `json:",omitempty"`

	// Field # 924
	ParentFTIMSpouseAdjustedGrossIncome string `json:",omitempty"`

	// Field # 925
	ParentFTIMSpouseNumberOfExemptions string `json:",omitempty"`

	// Field # 926
	ParentFTIMSpouseNumberOfDependents string `json:",omitempty"`

	// Field # 927
	ParentFTIMSpouseTotalIncomeEarnedAmount string `json:",omitempty"`

	// Field # 928
	ParentFTIMSpouseIncomeTaxPaid string `json:",omitempty"`

	// Field # 929
	ParentFTIMSpouseEducationCredits string `json:",omitempty"`

	// Field # 930
	ParentFTIMSpouseUntaxedIRADistributions string `json:",omitempty"`

	// Field # 931
	ParentFTIMSpouseIRADeductibleAndPayments string `json:",omitempty"`

	// Field # 932
	ParentFTIMSpouseTaxExemptInterest string `json:",omitempty"`

	// Field # 933
	ParentFTIMSpouseUntaxedPensionsAmount string `json:",omitempty"`

	// Field # 934
	ParentFTIMSpouseScheduleCNetProfitLoss string `json:",omitempty"`

	// Field # 935
	ParentFTIMSpouseScheduleAIndicator string `json:",omitempty"`

	// Field # 936
	ParentFTIMSpouseScheduleBIndicator string `json:",omitempty"`

	// Field # 937
	ParentFTIMSpouseScheduleDIndicator string `json:",omitempty"`

	// Field # 938
	ParentFTIMSpouseScheduleEIndicator string `json:",omitempty"`

	// Field # 939
	ParentFTIMSpouseScheduleFIndicator string `json:",omitempty"`

	// Field # 940
	ParentFTIMSpouseScheduleHIndicator string `json:",omitempty"`

	// Field # 941
	ParentFTIMSpouseIRSResponseCode string `json:",omitempty"`

	// Field # 942
	FTILabelEnd string `json:",omitempty"`

	// Field # 944
	StudentTotalIncome string `json:",omitempty"`

	// Field # 945
	ParentTotalIncome string `json:",omitempty"`

	// Field # 946
	FISAPTotalIncome string `json:",omitempty"`
}

// JsonString marshalls the ISIRecord to a string, emitting an error message but returning an empty string if an error occurs.
// Each field in the struct which has a value will be included in the JSON output with the same exact name,
// while fields with zero values will be omitted.
func (r *ISIRecord) JsonString(cid uuid.UUID) string {
	j, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return fmt.Sprintf("{\"errorMessage\":\"%s\",\"correlationID\":\"%s\"}", err.Error(), cid.String())
	}
	return string(j)
}
