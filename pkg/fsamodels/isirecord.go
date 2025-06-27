// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package fsamodels

import "time"

type ISIRecord struct {
	// Field # 1
	YearIndicator string

	// Field # 2
	FAFSAUUID string

	// Field # 3
	TransactionUUID string

	// Field # 4
	PersonUUID string

	// Field # 5
	TransactionNumber string

	// Field # 6
	DependencyModel string

	// Field # 7
	ApplicationSource string

	// Field # 8
	ApplicationReceiptDate time.Time

	// Field # 9
	TransactionSource string

	// Field # 10
	TransactionType string

	// Field # 11
	TransactionLanguage string

	// Field # 12
	TransactionReceiptDate time.Time

	// Field # 13
	TransactionProcessedDate time.Time

	// Field # 14
	TransactionStatus string

	// Field # 15
	RenewalDataUsed string

	// Field # 16
	FPSCorrectionReason string

	// Field # 17
	SAIChangeFlag string

	// Field # 18
	SAI string

	// Field # 19
	ProvisionalSAI string

	// Field # 20
	SAIFormula string

	// Field # 21
	SAIComputationType string

	// Field # 22
	MaxPellIndicator string

	// Field # 23
	MinimumPellIndicator string

	// Field # 25
	StudentFirstName string

	// Field # 26
	StudentMiddleName string

	// Field # 27
	StudentLastName string

	// Field # 28
	StudentSuffix string

	// Field # 29
	StudentDateOfBirth time.Time

	// Field # 30
	StudentSSN string

	// Field # 31
	StudentITIN string

	// Field # 32
	StudentPhoneNumber string

	// Field # 33
	StudentEmailAddress string

	// Field # 34
	StudentStreetAddress string

	// Field # 35
	StudentCity string

	// Field # 36
	StudentState string

	// Field # 37
	StudentZipCode string

	// Field # 38
	StudentCountry string

	// Field # 40
	StudentMaritalStatus string

	// Field # 41
	StudentGradeLevel string

	// Field # 42
	StudentFirstBachelorsDegreeBefore2526 string

	// Field # 43
	StudentPursuingTeacherCertification string

	// Field # 44
	StudentActiveDuty string

	// Field # 45
	StudentVeteran string

	// Field # 46
	StudentChildOrOtherDependents string

	// Field # 47
	StudentParentsDeceased string

	// Field # 48
	StudentWardOfCourt string

	// Field # 49
	StudentInFosterCare string

	// Field # 50
	StudentEmancipatedMinor string

	// Field # 51
	StudentLegalGuardianship string

	// Field # 52
	StudentPersonalCircumstancesNoneOfTheAbove string

	// Field # 53
	StudentUnaccompaniedHomelessYouthAndSelfSupporting string

	// Field # 54
	StudentUnaccompaniedHomelessGeneral string

	// Field # 55
	StudentUnaccompaniedHomelessHS string

	// Field # 56
	StudentUnaccompaniedHomelessTRIO string

	// Field # 57
	StudentUnaccompaniedHomelessFAA string

	// Field # 58
	StudentHomelessnessNoneOfTheAbove string

	// Field # 59
	StudentUnusualCircumstance string

	// Field # 60
	StudentUnsubOnly string

	// Field # 61
	StudentUpdatedFamilySize string

	// Field # 62
	StudentNumberInCollege string

	// Field # 63
	StudentCitizenshipStatus string

	// Field # 64
	StudentANumber string

	// Field # 65
	StudentStateOfLegalResidence string

	// Field # 66
	StudentLegalResidenceDate time.Time

	// Field # 67
	StudentEitherParentAttendCollege string

	// Field # 68
	StudentParentKilledInTheLineOfDuty string

	// Field # 69
	StudentHighSchoolCompletionStatus string

	// Field # 70
	StudentHighSchoolName string

	// Field # 71
	StudentHighSchoolCity string

	// Field # 72
	StudentHighSchoolState string

	// Field # 73
	StudentHighSchoolEquivalentDiplomaName string

	// Field # 74
	StudentHighSchoolEquivalentDiplomaState string

	// Field # 75
	StudentManuallyEnteredReceivedEITC string

	// Field # 76
	StudentManuallyEnteredReceivedFederalHousingAssistance string

	// Field # 77
	StudentManuallyEnteredReceivedFreeReducedPriceLunch string

	// Field # 78
	StudentManuallyEnteredReceivedMedicaid string

	// Field # 79
	StudentManuallyEnteredReceivedRefundableCreditFor36BHealthPlan string

	// Field # 80
	StudentManuallyEnteredReceivedSNAP string

	// Field # 81
	StudentManuallyEnteredReceivedSupplementalSecurityIncome string

	// Field # 82
	StudentManuallyEnteredReceivedTANF string

	// Field # 83
	StudentManuallyEnteredReceivedWIC string

	// Field # 84
	StudentManuallyEnteredFederalBenefitsNoneOfTheAbove string

	// Field # 85
	StudentManuallyEnteredFiled1040Or1040NR string

	// Field # 86
	StudentManuallyEnteredFiledNonUSTaxReturn string

	// Field # 87
	StudentManuallyEnteredFiledJointReturnWithCurrentSpouse string

	// Field # 88
	StudentManuallyEnteredTaxReturnFilingStatus string

	// Field # 89
	StudentManuallyEnteredIncomeEarnedFromWork string

	// Field # 90
	StudentManuallyEnteredTaxExemptInterestIncome string

	// Field # 91
	StudentManuallyEnteredUntaxedPortionsOfIRADistributions string

	// Field # 92
	StudentManuallyEnteredIRARollover string

	// Field # 93
	StudentManuallyEnteredUntaxedPortionsOfPensions string

	// Field # 94
	StudentManuallyEnteredPensionRollover string

	// Field # 95
	StudentManuallyEnteredAdjustedGrossIncome string

	// Field # 96
	StudentManuallyEnteredIncomeTaxPaid string

	// Field # 97
	StudentManuallyEnteredEarnedIncomeTaxCreditReceivedDuringTaxYear string

	// Field # 98
	StudentManuallyEnteredDeductiblePaymentsToIRAKeoghOther string

	// Field # 99
	StudentManuallyEnteredEducationCredits string

	// Field # 100
	StudentManuallyEnteredFiledScheduleABDEFH string

	// Field # 101
	StudentManuallyEnteredScheduleCAmount string

	// Field # 102
	StudentManuallyEnteredCollegeGrantAndScholarshipAid string

	// Field # 103
	StudentManuallyEnteredForeignEarnedIncomeExclusion string

	// Field # 104
	StudentManuallyEnteredChildSupportReceived string

	// Field # 105
	StudentManuallyEnteredTotalOfCashSavingsAndCheckingAccounts string

	// Field # 106
	StudentManuallyEnteredNetWorthOfCurrentInvestments string

	// Field # 107
	StudentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarms string

	// Field # 108
	StudentCollege1 string

	// Field # 109
	StudentCollege2 string

	// Field # 110
	StudentCollege3 string

	// Field # 111
	StudentCollege4 string

	// Field # 112
	StudentCollege5 string

	// Field # 113
	StudentCollege6 string

	// Field # 114
	StudentCollege7 string

	// Field # 115
	StudentCollege8 string

	// Field # 116
	StudentCollege9 string

	// Field # 117
	StudentCollege10 string

	// Field # 118
	StudentCollege11 string

	// Field # 119
	StudentCollege12 string

	// Field # 120
	StudentCollege13 string

	// Field # 121
	StudentCollege14 string

	// Field # 122
	StudentCollege15 string

	// Field # 123
	StudentCollege16 string

	// Field # 124
	StudentCollege17 string

	// Field # 125
	StudentCollege18 string

	// Field # 126
	StudentCollege19 string

	// Field # 127
	StudentCollege20 string

	// Field # 128
	StudentConsentToRetrieveAndDiscloseFTI string

	// Field # 129
	StudentSignature string

	// Field # 130
	StudentSignatureDate time.Time

	// Field # 132
	StudentSpouseFirstName string

	// Field # 133
	StudentSpouseMiddleName string

	// Field # 134
	StudentSpouseLastName string

	// Field # 135
	StudentSpouseSuffix string

	// Field # 136
	StudentSpouseDateOfBirth time.Time

	// Field # 137
	StudentSpouseSSN string

	// Field # 138
	StudentSpouseITIN string

	// Field # 139
	StudentSpousePhoneNumber string

	// Field # 140
	StudentSpouseEmailAddress string

	// Field # 141
	StudentSpouseStreetAddress string

	// Field # 142
	StudentSpouseCity string

	// Field # 143
	StudentSpouseState string

	// Field # 144
	StudentSpouseZipCode string

	// Field # 145
	StudentSpouseCountry string

	// Field # 146
	StudentSpouseFiled1040Or1040NR string

	// Field # 147
	StudentSpouseFiledNonUSTaxReturn string

	// Field # 148
	StudentSpouseTaxReturnFilingStatus string

	// Field # 149
	StudentSpouseIncomeEarnedFromWork string

	// Field # 150
	StudentSpouseTaxExemptInterestIncome string

	// Field # 151
	StudentSpouseUntaxedPortionsOfIRADistributions string

	// Field # 152
	StudentSpouseIRARollover string

	// Field # 153
	StudentSpouseUntaxedPortionsOfPensions string

	// Field # 154
	StudentSpousePensionRollover string

	// Field # 155
	StudentSpouseAdjustedGrossIncome string

	// Field # 156
	StudentSpouseIncomeTaxPaid string

	// Field # 157
	StudentSpouseDeductiblePaymentsToIRAKeoghOther string

	// Field # 158
	StudentSpouseEducationCredits string

	// Field # 159
	StudentSpouseFiledScheduleABDEFH string

	// Field # 160
	StudentSpouseScheduleCAmount string

	// Field # 161
	StudentSpouseForeignEarnedIncomeExclusion string

	// Field # 162
	StudentSpouseConsentToRetrieveAndDiscloseFTI string

	// Field # 163
	StudentSpouseSignature string

	// Field # 164
	StudentSpouseSignatureDate time.Time

	// Field # 166
	ParentFirstName string

	// Field # 167
	ParentMiddleName string

	// Field # 168
	ParentLastName string

	// Field # 169
	ParentSuffix string

	// Field # 170
	ParentDateOfBirth time.Time

	// Field # 171
	ParentSSN string

	// Field # 172
	ParentITIN string

	// Field # 173
	ParentPhoneNumber string

	// Field # 174
	ParentEmailAddress string

	// Field # 175
	ParentStreetAddress string

	// Field # 176
	ParentCity string

	// Field # 177
	ParentState string

	// Field # 178
	ParentZipCode string

	// Field # 179
	ParentCountry string

	// Field # 180
	ParentMaritalStatus string

	// Field # 181
	ParentStateOfLegalResidence string

	// Field # 182
	ParentLegalResidenceDate time.Time

	// Field # 183
	ParentUpdatedFamilySize string

	// Field # 184
	ParentNumberInCollege string

	// Field # 185
	ParentReceivedEITC string

	// Field # 186
	ParentReceivedFederalHousingAssistance string

	// Field # 187
	ParentReceivedFreeReducedPriceLunch string

	// Field # 188
	ParentReceivedMedicaid string

	// Field # 189
	ParentReceivedRefundableCreditFor36BHealthPlan string

	// Field # 190
	ParentReceivedSNAP string

	// Field # 191
	ParentReceivedSupplementalSecurityIncome string

	// Field # 192
	ParentReceivedTANF string

	// Field # 193
	ParentReceivedWIC string

	// Field # 194
	ParentFederalBenefitsNoneOfTheAbove string

	// Field # 195
	ParentFiled1040Or1040NR string

	// Field # 196
	ParentFileNonUSTaxReturn string

	// Field # 197
	ParentFiledJointReturnWithCurrentSpouse string

	// Field # 198
	ParentTaxReturnFilingStatus string

	// Field # 199
	ParentIncomeEarnedFromWork string

	// Field # 200
	ParentTaxExemptInterestIncome string

	// Field # 201
	ParentUntaxedPortionsOfIRADistributions string

	// Field # 202
	ParentIRARollover string

	// Field # 203
	ParentUntaxedPortionsOfPensions string

	// Field # 204
	ParentPensionRollover string

	// Field # 205
	ParentAdjustedGrossIncome string

	// Field # 206
	ParentIncomeTaxPaid string

	// Field # 207
	ParentEarnedIncomeTaxCreditReceivedDuringTaxYear string

	// Field # 208
	ParentDeductiblePaymentsToIRAKeoghOther string

	// Field # 209
	ParentEducationCredits string

	// Field # 210
	ParentFiledScheduleABDEFH string

	// Field # 211
	ParentScheduleCAmount string

	// Field # 212
	ParentCollegeGrantAndScholarshipAid string

	// Field # 213
	ParentForeignEarnedIncomeExclusion string

	// Field # 214
	ParentChildSupportReceived string

	// Field # 215
	ParentTotalOfCashSavingsAndCheckingAccounts string

	// Field # 216
	ParentNetWorthOfCurrentInvestments string

	// Field # 217
	ParentNetWorthOfBusinessesAndInvestmentFarms string

	// Field # 218
	ParentConsentToRetrieveAndDiscloseFTI string

	// Field # 219
	ParentSignature string

	// Field # 220
	ParentSignatureDate time.Time

	// Field # 222
	ParentSpouseFirstName string

	// Field # 223
	ParentSpouseMiddleName string

	// Field # 224
	ParentSpouseLastName string

	// Field # 225
	ParentSpouseSuffix string

	// Field # 226
	ParentSpouseDateOfBirth time.Time

	// Field # 227
	ParentSpouseSSN string

	// Field # 228
	ParentSpouseITIN string

	// Field # 229
	ParentSpousePhoneNumber string

	// Field # 230
	ParentSpouseEmailAddress string

	// Field # 231
	ParentSpouseStreetAddress string

	// Field # 232
	ParentSpouseCity string

	// Field # 233
	ParentSpouseState string

	// Field # 234
	ParentSpouseZipCode string

	// Field # 235
	ParentSpouseCountry string

	// Field # 236
	ParentSpouseFiled1040Or1040NR string

	// Field # 237
	ParentSpouseFileNonUSTaxReturn string

	// Field # 238
	ParentSpouseTaxReturnFilingStatus string

	// Field # 239
	ParentSpouseIncomeEarnedFromWork string

	// Field # 240
	ParentSpouseTaxExemptInterestIncome string

	// Field # 241
	ParentSpouseUntaxedPortionsOfIRADistributions string

	// Field # 242
	ParentSpouseIRARollover string

	// Field # 243
	ParentSpouseUntaxedPortionsOfPensions string

	// Field # 244
	ParentSpousePensionRollover string

	// Field # 245
	ParentSpouseAdjustedGrossIncome string

	// Field # 246
	ParentSpouseIncomeTaxPaid string

	// Field # 247
	ParentSpouseDeductiblePaymentsToIRAKeoghOther string

	// Field # 248
	ParentSpouseEducationCredits string

	// Field # 249
	ParentSpouseFiledScheduleABDEFH string

	// Field # 250
	ParentSpouseScheduleCAmount string

	// Field # 251
	ParentSpouseForeignEarnedIncomeExclusion string

	// Field # 252
	ParentSpouseConsentToRetrieveAndDiscloseFTI string

	// Field # 253
	ParentSpouseSignature string

	// Field # 254
	ParentSpouseSignatureDate time.Time

	// Field # 256
	PreparerFirstName string

	// Field # 257
	PreparerLastName string

	// Field # 258
	PreparerSSN string

	// Field # 259
	PreparerEIN string

	// Field # 260
	PreparerAffiliation string

	// Field # 261
	PreparerStreetAddress string

	// Field # 262
	PreparerCity string

	// Field # 263
	PreparerState string

	// Field # 264
	PreparerZipCode string

	// Field # 265
	PreparerSignature string

	// Field # 266
	PreparerSignatureDate time.Time

	// Field # 268
	StudentAffirmationStatus string

	// Field # 269
	StudentSpouseAffirmationStatus string

	// Field # 270
	ParentAffirmationStatus string

	// Field # 271
	ParentSpouseOrPartnerAffirmationStatus string

	// Field # 272
	StudentDateConsentGranted time.Time

	// Field # 273
	StudentSpouseDateConsentGranted time.Time

	// Field # 274
	ParentDateConsentGranted time.Time

	// Field # 275
	ParentSpouseOrPartnerDateConsentGranted time.Time

	// Field # 276
	StudentTransunionMatchStatus string

	// Field # 277
	StudentSpouseTransunionMatchStatus string

	// Field # 278
	StudentParentTransunionMatchStatus string

	// Field # 279
	StudentParentSpouseTransunionMatchStatus string

	// Field # 280
	CorrectionAppliedAgainstTransactionNumber string

	// Field # 281
	ProfessionalJudgement string

	// Field # 282
	DependencyOverrideIndicator string

	// Field # 283
	FAAFederalSchoolCode string

	// Field # 284
	FAASignature string

	// Field # 285
	IASGIndicator string

	// Field # 286
	ChildrenOfFallenHeroesIndicator string

	// Field # 287
	ElectronicTransactionIndicatorDestinationNumber string

	// Field # 288
	StudentSignatureSource string

	// Field # 289
	StudentSpouseSignatureSource string

	// Field # 290
	ParentSignatureSource string

	// Field # 291
	ParentSpouseOrPartnerSignatureSource string

	// Field # 292
	SpecialHandlingIndicator string

	// Field # 293
	AddressOnlyChangeFlag string

	// Field # 294
	FPSPushedISIRFlag string

	// Field # 295
	RejectStatusChangeFlag string

	// Field # 296
	VerificationTrackingFlag string

	// Field # 297
	StudentSelectedForVerification string

	// Field # 298
	IncarceratedApplicantFlag string

	// Field # 299
	NSLDSTransactionNumber string

	// Field # 300
	NSLDSDatabaseResultsFlag string

	// Field # 301
	HighSchoolFlag string

	// Field # 302
	StudentTotalFederalWorkStudyEarnings string

	// Field # 303
	StudentSpouseTotalFederalWorkStudyEarnings string

	// Field # 304
	ParentTotalFederalWorkStudyEarnings string

	// Field # 305
	ParentSpouseOrPartnerTotalFederalWorkStudyEarnings string

	// Field # 306
	TotalParentAllowancesAgainstIncome string

	// Field # 307
	ParentPayrollTaxAllowance string

	// Field # 308
	ParentIncomeProtectionAllowance string

	// Field # 309
	ParentEmploymentExpenseAllowance string

	// Field # 310
	ParentAvailableIncome string

	// Field # 311
	ParentAdjustedAvailableIncome string

	// Field # 312
	ParentContribution string

	// Field # 313
	StudentPayrollTaxAllowance string

	// Field # 314
	StudentIncomeProtectionAllowance string

	// Field # 315
	StudentAllowanceForParentsNegativeAdjustedAvailableIncome string

	// Field # 316
	StudentEmploymentExpenseAllowance string

	// Field # 317
	TotalStudentAllowancesAgainstIncome string

	// Field # 318
	StudentAvailableIncome string

	// Field # 319
	StudentContributionFromIncome string

	// Field # 320
	StudentAdjustedAvailableIncome string

	// Field # 321
	TotalStudentContributionFromSAAI string

	// Field # 322
	ParentDiscretionaryNetWorth string

	// Field # 323
	ParentNetWorth string

	// Field # 324
	ParentAssetProtectionAllowance string

	// Field # 325
	ParentContributionFromAssets string

	// Field # 326
	StudentNetWorth string

	// Field # 327
	StudentAssetProtectionAllowance string

	// Field # 328
	StudentContributionFromAssets string

	// Field # 329
	AssumedStudentFamilySize string

	// Field # 330
	AssumedParentFamilySize string

	// Field # 331
	StudentFirstNameCHVFlags string

	// Field # 332
	StudentMiddleNameCHVFlags string

	// Field # 333
	StudentLastNameCHVFLags string

	// Field # 334
	StudentSuffixCHVFLags string

	// Field # 335
	StudentDateOfBirthCHVFLags string

	// Field # 336
	StudentSSNCHVFlags string

	// Field # 337
	StudentITINCHVFLags string

	// Field # 338
	StudentPhoneNumberCHVFlags string

	// Field # 339
	StudentEmailAddressCHVFlags string

	// Field # 340
	StudentStreetAddressCHVFlags string

	// Field # 341
	StudentCityCHVFLags string

	// Field # 342
	StudentStateCHVFlags string

	// Field # 343
	StudentZipCodeCHVFlags string

	// Field # 344
	StudentCountryCHVFlags string

	// Field # 345
	StudentMaritalStatusCHVFlags string

	// Field # 346
	StudentGradeLevelInCollegeCHVFlags string

	// Field # 347
	StudentFirstBachelorsDegreeBeforeSchoolYearCHVFlags string

	// Field # 348
	StudentPursuingTeacherCertificationCHVFlags string

	// Field # 349
	StudentActiveDutyCHVFlags string

	// Field # 350
	StudentVeteranCHVFlags string

	// Field # 351
	StudentChildOrOtherDependentsCHVFlags string

	// Field # 352
	StudentParentsDeceasedCHVFlags string

	// Field # 353
	StudentWardOfCourtCHVFlags string

	// Field # 354
	StudentInFosterCareCHVFlags string

	// Field # 355
	StudentEmancipatedMinorCHVFlags string

	// Field # 356
	StudentLegalGuardianshipCHVFlags string

	// Field # 357
	StudentPersonalCircumstancesNoneOfTheAboveCHVFlags string

	// Field # 358
	StudentUnaccompaniedHomelessYouthOrIsUnaccompaniedAtRiskSelfSupportingCHVFlags string

	// Field # 359
	StudentUnaccompaniedAndHomelessGeneralCHVFlags string

	// Field # 360
	StudentUnaccompaniedAndHomelessHSCHVFlags string

	// Field # 361
	StudentUnaccompaniedAndHomelessTRIOCHVFlags string

	// Field # 362
	StudentUnaccompaniedAndHomelessFAACHVFlags string

	// Field # 363
	StudentHomelessnessNoneOfTheAboveCHVFlags string

	// Field # 364
	StudentHasUnusualCircumstanceCHVFlags string

	// Field # 365
	StudentUnsubOnlyCHVFlags string

	// Field # 366
	StudentUpdatedFamilySizeCHVFlags string

	// Field # 367
	StudentNumberInCollegeCorrectionCHVFlags string

	// Field # 368
	StudentCitizenshipStatusCorrectionCHVFlags string

	// Field # 369
	StudentANumberCHVFlags string

	// Field # 370
	StudentStateOfLegalResidenceCHVFlags string

	// Field # 371
	StudentLegalResidenceDateCHVFlags string

	// Field # 372
	StudentEitherParentAttendCollegeCHVFlags string

	// Field # 373
	StudentParentKilledInTheLineOfDutyCHVFlags string

	// Field # 374
	StudentHighSchoolCompletionStatusCHVFlags string

	// Field # 375
	StudentHighSchoolNameCHVFlags string

	// Field # 376
	StudentHighSchoolCityCHVFlags string

	// Field # 377
	StudentHighSchoolStateCHVFlags string

	// Field # 378
	StudentHighSchoolEquivalentDiplomaNameCHVFlags string

	// Field # 379
	StudentHighSchoolEquivalentDiplomaStateCHVFlags string

	// Field # 380
	StudentReceivedEITCCHVFlags string

	// Field # 381
	StudentReceivedFederalHousingAssistanceCHVFlags string

	// Field # 382
	StudentReceivedFreeReducedPriceLunchCHVFlags string

	// Field # 383
	StudentReceivedMedicaidCHVFlags string

	// Field # 384
	StudentReceivedRefundableCreditFor36BHealthPlanCHVFlags string

	// Field # 385
	StudentReceivedSNAPCHVFlags string

	// Field # 386
	StudentReceivedSupplementalSecurityIncomeCHVFlags string

	// Field # 387
	StudentReceivedTANFCHVFlags string

	// Field # 388
	StudentReceivedWICCHVFlags string

	// Field # 389
	StudentFederalBenefitsNoneOfTheAboveCHVFlags string

	// Field # 390
	StudentFiled1040Or1040NRCHVFlags string

	// Field # 391
	StudentFiledNonUSTaxReturnCHVFlags string

	// Field # 392
	StudentFiledJointReturnWithCurrentSpouseCHVFlags string

	// Field # 393
	StudentTaxReturnFilingStatusCHVFlags string

	// Field # 394
	StudentIncomeEarnedFromWorkCorrectionCHVFlags string

	// Field # 395
	StudentTaxExemptInterestIncomeCHVFlags string

	// Field # 396
	StudentUntaxedPortionsOfIRADistributionsCHVFlags string

	// Field # 397
	StudentIRARolloverCHVFlags string

	// Field # 398
	StudentUntaxedPortionsOfPensionsCHVFlags string

	// Field # 399
	StudentPensionRolloverCHVFlags string

	// Field # 400
	StudentAdjustedGrossIncomeCHVFlags string

	// Field # 401
	StudentIncomeTaxPaidCHVFlags string

	// Field # 402
	StudentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlags string

	// Field # 403
	StudentDeductiblePaymentsToIRAKeoghOtherCHVFlags string

	// Field # 404
	StudentEducationCreditsCHVFlags string

	// Field # 405
	StudentFiledScheduleABDEFHCHVFlags string

	// Field # 406
	StudentScheduleCAmountCHVFlags string

	// Field # 407
	StudentCollegeGrantAndScholarshipAidCHVFlags string

	// Field # 408
	StudentForeignEarnedIncomeExclusionCHVFlags string

	// Field # 409
	StudentChildSupportReceivedCHVFlags string

	// Field # 410
	StudentNetWorthOfBusinessesAndInvestmentFarmsCHVFlags string

	// Field # 411
	StudentNetWorthOfCurrentInvestmentsCHVFlags string

	// Field # 412
	StudentTotalOfCashSavingsAndCheckingCHVFlags string

	// Field # 413
	StudentCollege1CHVFlags string

	// Field # 414
	StudentCollege2CHVFlags string

	// Field # 415
	StudentCollege3CHVFlags string

	// Field # 416
	StudentCollege4CHVFlags string

	// Field # 417
	StudentCollege5CHVFlags string

	// Field # 418
	StudentCollege6CHVFlags string

	// Field # 419
	StudentCollege7CHVFlags string

	// Field # 420
	StudentCollege8CHVFlags string

	// Field # 421
	StudentCollege9CHVFlags string

	// Field # 422
	StudentCollege10CHVFlags string

	// Field # 423
	StudentCollege11CHVFlags string

	// Field # 424
	StudentCollege12CHVFlags string

	// Field # 425
	StudentCollege13CHVFlags string

	// Field # 426
	StudentCollege14CHVFlags string

	// Field # 427
	StudentCollege15CHVFlags string

	// Field # 428
	StudentCollege16CHVFlags string

	// Field # 429
	StudentCollege17CHVFlags string

	// Field # 430
	StudentCollege18CHVFlags string

	// Field # 431
	StudentCollege19CHVFlags string

	// Field # 432
	StudentCollege20CHVFlags string

	// Field # 433
	StudentConsentToRetrieveAndDiscloseFTICHVFlags string

	// Field # 434
	StudentSignatureCHVFlags string

	// Field # 435
	StudentSignatureDateCHVFlags string

	// Field # 436
	StudentSpouseFirstNameCHVFlags string

	// Field # 437
	StudentSpouseMiddleNameCHVFlags string

	// Field # 438
	StudentSpouseLastNameCHVFlags string

	// Field # 439
	StudentSpouseSuffixCHVFlags string

	// Field # 440
	StudentSpouseDateOfBirthCHVFlags string

	// Field # 441
	StudentSpouseSSNCHVFlags string

	// Field # 442
	StudentSpouseITINCHVFlags string

	// Field # 443
	StudentSpousePhoneNumberCHVFlags string

	// Field # 444
	StudentSpouseEmailAddressCHVFlags string

	// Field # 445
	StudentSpouseStreetAddressCHVFlags string

	// Field # 446
	StudentSpouseCityCHVFlags string

	// Field # 447
	StudentSpouseStateCHVFlags string

	// Field # 448
	StudentSpouseZipCodeCHVFlags string

	// Field # 449
	StudentSpouseCountryCHVFlags string

	// Field # 450
	StudentSpouseFiled1040Or1040NRCHVFlags string

	// Field # 451
	StudentSpouseFiledNonUSTaxReturnCHVFlags string

	// Field # 452
	StudentSpouseTaxReturnFilingStatusCHVFlags string

	// Field # 453
	StudentSpouseIncomeEarnedFromWorkCHVFlags string

	// Field # 454
	StudentSpouseTaxExemptInterestIncomeCHVFlags string

	// Field # 455
	StudentSpouseUntaxedPortionsOfIRADistributionsCHVFlags string

	// Field # 456
	StudentSpouseIRARolloverCHVFlags string

	// Field # 457
	StudentSpouseUntaxedPortionsOfPensionsCHVFlags string

	// Field # 458
	StudentSpousePensionRolloverCHVFlags string

	// Field # 459
	StudentSpouseAdjustedGrossIncomeCHVFlags string

	// Field # 460
	StudentSpouseIncomeTaxPaidCHVFlags string

	// Field # 461
	StudentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlags string

	// Field # 462
	StudentSpouseEducationCreditsCHVFlags string

	// Field # 463
	StudentSpouseFiledScheduleABDEFHCHVFlags string

	// Field # 464
	StudentSpouseScheduleCAmountCHVFlags string

	// Field # 465
	StudentSpouseForeignEarnedIncomeExclusionCHVFlags string

	// Field # 466
	StudentSpouseConsentToRetrieveAndDiscloseFTICHVFlags string

	// Field # 467
	StudentSpouseSignatureCHVFlags string

	// Field # 468
	StudentSpouseSignatureDateCHVFlags string

	// Field # 469
	ParentFirstNameCHVFlags string

	// Field # 470
	ParentMiddleNameCHVFlags string

	// Field # 471
	ParentLastNameCHVFlags string

	// Field # 472
	ParentSuffixCHVFlags string

	// Field # 473
	ParentDateOfBirthCHVFlags string

	// Field # 474
	ParentSSNCHVFlags string

	// Field # 475
	ParentITINCHVFlags string

	// Field # 476
	ParentPhoneNumberCHVFlags string

	// Field # 477
	ParentEmailAddressCHVFlags string

	// Field # 478
	ParentStreetAddressCHVFlags string

	// Field # 479
	ParentCityCHVFlags string

	// Field # 480
	ParentStateCHVFlags string

	// Field # 481
	ParentZipCodeCHVFlags string

	// Field # 482
	ParentCountryCHVFlags string

	// Field # 483
	ParentMaritalStatusCHVFlags string

	// Field # 484
	ParentStateOfLegalResidenceCHVFlags string

	// Field # 485
	ParentLegalResidenceDateCHVFlags string

	// Field # 486
	ParentUpdatedFamilySizeCHVFlags string

	// Field # 487
	ParentNumberInCollegeCHVFlags string

	// Field # 488
	ParentReceivedEITCCHVFlags string

	// Field # 489
	ParentReceivedFederalHousingAssistanceCHVFlags string

	// Field # 490
	ParentReceivedFreeReducedPriceLunchCHVFlags string

	// Field # 491
	ParentReceivedMedicaidCHVFlags string

	// Field # 492
	ParentReceivedRefundableCreditFor36BHealthPlanCHVFlags string

	// Field # 493
	ParentReceivedSNAPCHVFlags string

	// Field # 494
	ParentReceivedSupplementalSecurityIncomeCHVFlags string

	// Field # 495
	ParentReceivedTANFCHVFlags string

	// Field # 496
	ParentReceivedWICCHVFlags string

	// Field # 497
	ParentFederalBenefitsNoneOfTheAboveCHVFlags string

	// Field # 498
	ParentFiled1040Or1040NRCHVFlags string

	// Field # 499
	ParentFileNonUSTaxReturnCHVFlags string

	// Field # 500
	ParentFiledJointReturnWithCurrentSpouseCHVFlags string

	// Field # 501
	ParentTaxReturnFilingStatusCHVFlags string

	// Field # 502
	ParentIncomeEarnedFromWorkCHVFlags string

	// Field # 503
	ParentTaxExemptInterestIncomeCHVFlags string

	// Field # 504
	ParentUntaxedPortionsOfIRADistributionsCHVFlags string

	// Field # 505
	ParentIRARolloverCHVFlags string

	// Field # 506
	ParentUntaxedPortionsOfPensionsCHVFlags string

	// Field # 507
	ParentPensionRolloverCHVFlags string

	// Field # 508
	ParentAdjustedGrossIncomeCHVFlags string

	// Field # 509
	ParentIncomeTaxPaidCHVFlags string

	// Field # 510
	ParentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlags string

	// Field # 511
	ParentDeductiblePaymentsToIRAKeoghOtherCHVFlags string

	// Field # 512
	ParentEducationCreditsCHVFlags string

	// Field # 513
	ParentFiledScheduleABDEFHCHVFlags string

	// Field # 514
	ParentScheduleCAmountCHVFlags string

	// Field # 515
	ParentCollegeGrantAndScholarshipAidCHVFlags string

	// Field # 516
	ParentForeignEarnedIncomeExclusionCHVFlags string

	// Field # 517
	ParentChildSupportReceivedCHVFlags string

	// Field # 518
	ParentNetWorthOfCurrentInvestmentsCHVFlags string

	// Field # 519
	ParentTotalOfCashSavingsAndCheckingAccountsCHVFlags string

	// Field # 520
	ParentNetWorthOfBusinessesAndInvestmentFarmsCHVFlags string

	// Field # 521
	ParentConsentToRetrieveAndDiscloseFTICHVFlags string

	// Field # 522
	ParentSignatureCHVFlags string

	// Field # 523
	ParentSignatureDateCHVFlags string

	// Field # 524
	ParentSpouseFirstNameCHVFlags string

	// Field # 525
	ParentSpouseMiddleNameCHVFlags string

	// Field # 526
	ParentSpouseLastNameCHVFlags string

	// Field # 527
	ParentSpouseSuffixCHVFlags string

	// Field # 528
	ParentSpouseDateOfBirthCHVFlags string

	// Field # 529
	ParentSpouseSSNCHVFlags string

	// Field # 530
	ParentSpouseITINCHVFlags string

	// Field # 531
	ParentSpousePhoneNumberCHVFlags string

	// Field # 532
	ParentSpouseEmailAddressCHVFlags string

	// Field # 533
	ParentSpouseStreetAddressCHVFlags string

	// Field # 534
	ParentSpouseCityCHVFlags string

	// Field # 535
	ParentSpouseStateCHVFlags string

	// Field # 536
	ParentSpouseZipCodeCHVFlags string

	// Field # 537
	ParentSpouseCountryCHVFlags string

	// Field # 538
	ParentSpouseFiled1040Or1040NRCHVFlags string

	// Field # 539
	ParentSpouseFileNonUSTaxReturnCHVFlags string

	// Field # 540
	ParentSpouseTaxReturnFilingStatusCHVFlags string

	// Field # 541
	ParentSpouseIncomeEarnedFromWorkCHVFlags string

	// Field # 542
	ParentSpouseTaxExemptInterestIncomeCHVFlags string

	// Field # 543
	ParentSpouseUntaxedPortionsOfIRADistributionsCHVFlags string

	// Field # 544
	ParentSpouseIRARolloverCHVFlags string

	// Field # 545
	ParentSpouseUntaxedPortionsOfPensionsCHVFlags string

	// Field # 546
	ParentSpousePensionRolloverCHVFlags string

	// Field # 547
	ParentSpouseAdjustedGrossIncomeCHVFlags string

	// Field # 548
	ParentSpouseIncomeTaxPaidCHVFlags string

	// Field # 549
	ParentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlags string

	// Field # 550
	ParentSpouseEducationCreditsCHVFlags string

	// Field # 551
	ParentSpouseFiledScheduleABDEFHCHVFlags string

	// Field # 552
	ParentSpouseScheduleCAmountCHVFlags string

	// Field # 553
	ParentSpouseForeignEarnedIncomeExclusionCHVFlags string

	// Field # 554
	ParentSpouseConsentToRetrieveAndDiscloseFTICHVFlags string

	// Field # 555
	ParentSpouseSignatureCHVFlags string

	// Field # 556
	ParentSpouseSignatureDateCHVFlags string

	// Field # 557
	DHSPrimaryMatchStatus string

	// Field # 559
	DHSCaseNumber string

	// Field # 560
	NSLDSMatchStatus string

	// Field # 561
	NSLDSPostscreeningReasonCode string

	// Field # 562
	StudentSSACitizenshipFlagResults string

	// Field # 563
	StudentSSAMatchStatus string

	// Field # 564
	StudentSpouseSSAMatchStatus string

	// Field # 565
	ParentSSAMatchStatus string

	// Field # 566
	ParentSpouseOrPartnerSSAMatchStatus string

	// Field # 567
	VAMatchFlag string

	// Field # 568
	CommentCodes string

	// Field # 569
	DrugAbuseHoldIndicator string

	// Field # 570
	GraduateFlag string

	// Field # 571
	PellGrantEligibilityFlag string

	// Field # 572
	ReprocessedReasonCode string

	// Field # 573
	FPSCFlag string

	// Field # 574
	FPSCChangeFlag string

	// Field # 575
	ElectronicFederalSchoolCodeIndicator string

	// Field # 576
	RejectReasonCodes string

	// Field # 577
	ElectronicTransactionIndicatorFlag string

	// Field # 578
	StudentLastNameSSNChangeFlag string

	// Field # 579
	HighSchoolCode string

	// Field # 580
	VerificationSelectionChangeFlag string

	// Field # 581
	UseUserProvidedDataOnly string

	// Field # 583
	NSLDSPellOverpaymentFlag string

	// Field # 584
	NSLDSPellOverpaymentContact string

	// Field # 585
	NSLDSFSEOGOverpaymentFlag string

	// Field # 586
	NSLDSFSEOGOverpaymentContact string

	// Field # 587
	NSLDSPerkinsOverpaymentFlag string

	// Field # 588
	NSLDSPerkinsOverpaymentContact string

	// Field # 589
	NSLDSTEACHGrantOverpaymentFlag string

	// Field # 590
	NSLDSTEACHGrantOverpaymentContact string

	// Field # 591
	NSLDSIraqAndAfghanistanServiceGrantOverpaymentFlag string

	// Field # 592
	NSLDSIraqAndAfghanistanServiceGrantOverpaymentContact string

	// Field # 593
	NSLDSDefaultedLoanFlag string

	// Field # 594
	NSLDSDischargedLoanFlag string

	// Field # 595
	NSLDSFraudLoanFlag string

	// Field # 596
	NSLDSSatisfactoryArrangementsFlag string

	// Field # 597
	NSLDSActiveBankruptcyFlag string

	// Field # 598
	NSLDSTEACHGrantConvertedToLoanFlag string

	// Field # 599
	NSLDSAggregateSubsidizedOutstandingPrincipalBalance string

	// Field # 600
	NSLDSAggregateUnsubsidizedOutstandingPrincipalBalance string

	// Field # 601
	NSLDSAggregateCombinedOutstandingPrincipalBalance string

	// Field # 602
	NSLDSAggregateUnallocatedConsolidatedOutstandingPrincipalBalance string

	// Field # 603
	NSLDSAggregateTEACHLoanPrincipalBalance string

	// Field # 604
	NSLDSAggregateSubsidizedPendingDisbursement string

	// Field # 605
	NSLDSAggregateUnsubsidizedPendingDisbursement string

	// Field # 606
	NSLDSAggregateCombinedPendingDisbursement string

	// Field # 607
	NSLDSAggregateSubsidizedTotal string

	// Field # 608
	NSLDSAggregateUnsubsidizedTotal string

	// Field # 609
	NSLDSAggregateCombinedTotal string

	// Field # 610
	NSLDSUnallocatedConsolidatedTotal string

	// Field # 611
	NSLDSTEACHLoanTotal string

	// Field # 612
	NSLDSPerkinsTotalDisbursements string

	// Field # 613
	NSLDSPerkinsCurrentYearDisbursementAmount string

	// Field # 614
	NSLDSAggregateTEACHGrantUndergraduateDisbursedTotal string

	// Field # 615
	NSLDSAggregateTEACHGraduateDisbursementAmount string

	// Field # 616
	NSLDSDefaultedLoanChangeFlag string

	// Field # 617
	NSLDSFraudLoanChangeFlag string

	// Field # 618
	NSLDSDischargedLoanChangeFlag string

	// Field # 619
	NSLDSLoanSatisfactoryRepaymentChangeFlag string

	// Field # 620
	NSLDSActiveBankruptcyChangeFlag string

	// Field # 621
	NSLDSTEACHGrantToLoanConversionChangeFlag string

	// Field # 622
	NSLDSOverpaymentsChangeFlag string

	// Field # 623
	NSLDSAggregateLoanChangeFlag string

	// Field # 624
	NSLDSPerkinsLoanChangeFlag string

	// Field # 625
	NSLDSPellPaymentChangeFlag string

	// Field # 626
	NSLDSTEACHGrantChangeFlag string

	// Field # 627
	NSLDSAdditionalPellFlag string

	// Field # 628
	NSLDSAdditionalLoansFlag string

	// Field # 629
	NSLDSAdditionalTEACHGrantFlag string

	// Field # 630
	NSLDSDirectLoanMasterPromNoteFlag string

	// Field # 631
	NSLDSDirectLoanPLUSMasterPromNoteFlag string

	// Field # 632
	NSLDSDirectLoanGraduatePLUSMasterPromNoteFlag string

	// Field # 633
	NSLDSUndergraduateSubsidizedLoanLimitFlag string

	// Field # 634
	NSLDSUndergraduateCombinedLoanLimitFlag string

	// Field # 635
	NSLDSGraduateSubsidizedLoanLimitFlag string

	// Field # 636
	NSLDSGraduateCombinedLoanLimitFlag string

	// Field # 637
	NSLDSLEULimitIndicator string

	// Field # 638
	NSLDSPellLifetimeEligibilityUsed string

	// Field # 639
	NSLDSSULAFlag string

	// Field # 640
	NSLDSSubsidizedLimitEligibilityFlag string

	// Field # 641
	NSLDSUnusualEnrollmentHistoryFlag string

	// Field # 643
	NSLDSPellSequenceNumber1 string

	// Field # 644
	NSLDSPellVerificationFlag1 string

	// Field # 645
	NSLDSSAI1 string

	// Field # 646
	NSLDSPellSchoolCode1 string

	// Field # 647
	NSLDSPellTransactionNumber1 string

	// Field # 648
	NSLDSPellLastDisbursementDate1 time.Time

	// Field # 649
	NSLDSPellScheduledAmount1 string

	// Field # 650
	NSLDSPellAmountPaidToDate1 time.Time

	// Field # 651
	NSLDSPellPercentEligibilityUsedDecimal1 string

	// Field # 652
	NSLDSPellAwardAmount1 string

	// Field # 653
	NSLDSAdditionalEligibilityIndicator1 string

	// Field # 655
	NSLDSPellSequenceNumber2 string

	// Field # 656
	NSLDSPellVerificationFlag2 string

	// Field # 657
	NSLDSSAI2 string

	// Field # 658
	NSLDSPellSchoolCode2 string

	// Field # 659
	NSLDSPellTransactionNumber2 string

	// Field # 660
	NSLDSPellLastDisbursementDate2 time.Time

	// Field # 661
	NSLDSPellScheduledAmount2 string

	// Field # 662
	NSLDSPellAmountPaidToDate2 time.Time

	// Field # 663
	NSLDSPellPercentEligibilityUsedDecimal2 string

	// Field # 664
	NSLDSPellAwardAmount2 string

	// Field # 665
	NSLDSAdditionalEligibilityIndicator2 string

	// Field # 667
	NSLDSPellSequenceNumber3 string

	// Field # 668
	NSLDSPellVerificationFlag3 string

	// Field # 669
	NSLDSSAI3 string

	// Field # 670
	NSLDSPellSchoolCode3 string

	// Field # 671
	NSLDSPellTransactionNumber3 string

	// Field # 672
	NSLDSPellLastDisbursementDate3 time.Time

	// Field # 673
	NSLDSPellScheduledAmount3 string

	// Field # 674
	NSLDSPellAmountPaidToDate3 time.Time

	// Field # 675
	NSLDSPellPercentEligibilityUsedDecimal3 string

	// Field # 676
	NSLDSPellAwardAmount3 string

	// Field # 677
	NSLDSAdditionalEligibilityIndicator3 string

	// Field # 679
	NSLDSTEACHGrantSequence1 string

	// Field # 680
	NSLDSTEACHGrantSchoolCode1 string

	// Field # 681
	NSLDSTEACHGrantTransactionNumber1 string

	// Field # 682
	NSLDSTEACHGrantLastDisbursementDate1 time.Time

	// Field # 683
	NSLDSTEACHGrantScheduledAmount1 string

	// Field # 684
	NSLDSTEACHGrantAmountPaidToDate1 time.Time

	// Field # 685
	NSLDSTEACHGrantAwardAmount1 string

	// Field # 686
	NSLDSTEACHGrantAcademicYearLevel1 string

	// Field # 687
	NSLDSTEACHGrantAwardYear1 string

	// Field # 688
	NSLDSTEACHGrantLoanConversionFlag1 string

	// Field # 689
	NSLDSTEACHGrantDischargeCode1 string

	// Field # 690
	NSLDSTEACHGrantDischargeAmount1 string

	// Field # 691
	NSLDSTEACHGrantAdjustedDisbursement1 string

	// Field # 693
	NSLDSTEACHGrantSequence2 string

	// Field # 694
	NSLDSTEACHGrantSchoolCode2 string

	// Field # 695
	NSLDSTEACHGrantTransactionNumber2 string

	// Field # 696
	NSLDSTEACHGrantLastDisbursementDate2 time.Time

	// Field # 697
	NSLDSTEACHGrantScheduledAmount2 string

	// Field # 698
	NSLDSTEACHGrantAmountPaidToDate2 time.Time

	// Field # 699
	NSLDSTEACHGrantAwardAmount2 string

	// Field # 700
	NSLDSTEACHGrantAcademicYearLevel2 string

	// Field # 701
	NSLDSTEACHGrantAwardYear2 string

	// Field # 702
	NSLDSTEACHGrantLoanConversionFlag2 string

	// Field # 703
	NSLDSTEACHGrantDischargeCode2 string

	// Field # 704
	NSLDSTEACHGrantDischargeAmount2 string

	// Field # 705
	NSLDSTEACHGrantAdjustedDisbursement2 string

	// Field # 707
	NSLDSTEACHGrantSequence3 string

	// Field # 708
	NSLDSTEACHGrantSchoolCode3 string

	// Field # 709
	NSLDSTEACHGrantTransactionNumber3 string

	// Field # 710
	NSLDSTEACHGrantLastDisbursementDate3 time.Time

	// Field # 711
	NSLDSTEACHGrantScheduledAmount3 string

	// Field # 712
	NSLDSTEACHGrantAmountPaidToDate3 time.Time

	// Field # 713
	NSLDSTEACHGrantAwardAmount3 string

	// Field # 714
	NSLDSTEACHGrantAcademicYearLevel3 string

	// Field # 715
	NSLDSTEACHGrantAwardYear3 string

	// Field # 716
	NSLDSTEACHGrantLoanConversionFlag3 string

	// Field # 717
	NSLDSTEACHGrantDischargeCode3 string

	// Field # 718
	NSLDSTEACHGrantDischargeAmount3 string

	// Field # 719
	NSLDSTEACHGrantAdjustedDisbursement3 string

	// Field # 721
	NSLDSLoanSequenceNumber1 string

	// Field # 722
	NSLDSLoanDefaultedRecentIndicator1 string

	// Field # 723
	NSLDSLoanChangeFlag1 string

	// Field # 724
	NSLDSLoanTypeCode1 string

	// Field # 725
	NSLDSLoanNetAmount1 string

	// Field # 726
	NSLDSLoanCurrentStatusCode1 string

	// Field # 727
	NSLDSLoanCurrentStatusDate1 time.Time

	// Field # 728
	NSLDSLoanOutstandingPrincipalBalance1 string

	// Field # 729
	NSLDSLoanOutstandingPrincipalBalanceDate1 time.Time

	// Field # 730
	NSLDSLoanPeriodBeginDate1 time.Time

	// Field # 731
	NSLDSLoanPeriodEndDate1 time.Time

	// Field # 732
	NSLDSLoanGuarantyAgencyCode1 string

	// Field # 733
	NSLDSLoanContactType1 string

	// Field # 734
	NSLDSLoanSchoolCode1 string

	// Field # 735
	NSLDSLoanContactCode1 string

	// Field # 736
	NSLDSLoanGradeLevel1 string

	// Field # 737
	NSLDSLoanAdditionalUnsubsidizedFlag1 string

	// Field # 738
	NSLDSLoanCapitalizedInterestFlag1 string

	// Field # 739
	NSLDSLoanDisbursementAmount1 string

	// Field # 740
	NSLDSLoanDisbursementDate1 time.Time

	// Field # 741
	NSLDSLoanConfirmedLoanSubsidyStatus1 string

	// Field # 742
	NSLDSLoanSubsidyStatusDate1 time.Time

	// Field # 744
	NSLDSLoanSequenceNumber2 string

	// Field # 745
	NSLDSLoanDefaultedRecentIndicator2 string

	// Field # 746
	NSLDSLoanChangeFlag2 string

	// Field # 747
	NSLDSLoanTypeCode2 string

	// Field # 748
	NSLDSLoanNetAmount2 string

	// Field # 749
	NSLDSLoanCurrentStatusCode2 string

	// Field # 750
	NSLDSLoanCurrentStatusDate2 time.Time

	// Field # 751
	NSLDSLoanOutstandingPrincipalBalance2 string

	// Field # 752
	NSLDSLoanOutstandingPrincipalBalanceDate2 time.Time

	// Field # 753
	NSLDSLoanPeriodBeginDate2 time.Time

	// Field # 754
	NSLDSLoanPeriodEndDate2 time.Time

	// Field # 755
	NSLDSLoanGuarantyAgencyCode2 string

	// Field # 756
	NSLDSLoanContactType2 string

	// Field # 757
	NSLDSLoanSchoolCode2 string

	// Field # 758
	NSLDSLoanContactCode2 string

	// Field # 759
	NSLDSLoanGradeLevel2 string

	// Field # 760
	NSLDSLoanAdditionalUnsubsidizedFlag2 string

	// Field # 761
	NSLDSLoanCapitalizedInterestFlag2 string

	// Field # 762
	NSLDSLoanDisbursementAmount2 string

	// Field # 763
	NSLDSLoanDisbursementDate2 time.Time

	// Field # 764
	NSLDSLoanConfirmedLoanSubsidyStatus2 string

	// Field # 765
	NSLDSLoanSubsidyStatusDate2 time.Time

	// Field # 767
	NSLDSLoanSequenceNumber3 string

	// Field # 768
	NSLDSLoanDefaultedRecentIndicator3 string

	// Field # 769
	NSLDSLoanChangeFlag3 string

	// Field # 770
	NSLDSLoanTypeCode3 string

	// Field # 771
	NSLDSLoanNetAmount3 string

	// Field # 772
	NSLDSLoanCurrentStatusCode3 string

	// Field # 773
	NSLDSLoanCurrentStatusDate3 time.Time

	// Field # 774
	NSLDSLoanOutstandingPrincipalBalance3 string

	// Field # 775
	NSLDSLoanOutstandingPrincipalBalanceDate3 time.Time

	// Field # 776
	NSLDSLoanPeriodBeginDate3 time.Time

	// Field # 777
	NSLDSLoanPeriodEndDate3 time.Time

	// Field # 778
	NSLDSLoanGuarantyAgencyCode3 string

	// Field # 779
	NSLDSLoanContactType3 string

	// Field # 780
	NSLDSLoanSchoolCode3 string

	// Field # 781
	NSLDSLoanContactCode3 string

	// Field # 782
	NSLDSLoanGradeLevel3 string

	// Field # 783
	NSLDSLoanAdditionalUnsubsidizedFlag3 string

	// Field # 784
	NSLDSLoanCapitalizedInterestFlag3 string

	// Field # 785
	NSLDSLoanDisbursementAmount3 string

	// Field # 786
	NSLDSLoanDisbursementDate3 time.Time

	// Field # 787
	NSLDSLoanConfirmedLoanSubsidyStatus3 string

	// Field # 788
	NSLDSLoanSubsidyStatusDate3 time.Time

	// Field # 790
	NSLDSLoanSequenceNumber4 string

	// Field # 791
	NSLDSLoanDefaultedRecentIndicator4 string

	// Field # 792
	NSLDSLoanChangeFlag4 string

	// Field # 793
	NSLDSLoanTypeCode4 string

	// Field # 794
	NSLDSLoanNetAmount4 string

	// Field # 795
	NSLDSLoanCurrentStatusCode4 string

	// Field # 796
	NSLDSLoanCurrentStatusDate4 time.Time

	// Field # 797
	NSLDSLoanOutstandingPrincipalBalance4 string

	// Field # 798
	NSLDSLoanOutstandingPrincipalBalanceDate4 time.Time

	// Field # 799
	NSLDSLoanPeriodBeginDate4 time.Time

	// Field # 800
	NSLDSLoanPeriodEndDate4 time.Time

	// Field # 801
	NSLDSLoanGuarantyAgencyCode4 string

	// Field # 802
	NSLDSLoanContactType4 string

	// Field # 803
	NSLDSLoanSchoolCode4 string

	// Field # 804
	NSLDSLoanContactCode4 string

	// Field # 805
	NSLDSLoanGradeLevel4 string

	// Field # 806
	NSLDSLoanAdditionalUnsubsidizedFlag4 string

	// Field # 807
	NSLDSLoanCapitalizedInterestFlag4 string

	// Field # 808
	NSLDSLoanDisbursementAmount4 string

	// Field # 809
	NSLDSLoanDisbursementDate4 time.Time

	// Field # 810
	NSLDSLoanConfirmedLoanSubsidyStatus4 string

	// Field # 811
	NSLDSLoanSubsidyStatusDate4 time.Time

	// Field # 813
	NSLDSLoanSequenceNumber5 string

	// Field # 814
	NSLDSLoanDefaultedRecentIndicator5 string

	// Field # 815
	NSLDSLoanChangeFlag5 string

	// Field # 816
	NSLDSLoanTypeCode5 string

	// Field # 817
	NSLDSLoanNetAmount5 string

	// Field # 818
	NSLDSLoanCurrentStatusCode5 string

	// Field # 819
	NSLDSLoanCurrentStatusDate5 time.Time

	// Field # 820
	NSLDSLoanOutstandingPrincipalBalance5 string

	// Field # 821
	NSLDSLoanOutstandingPrincipalBalanceDate5 time.Time

	// Field # 822
	NSLDSLoanPeriodBeginDate5 time.Time

	// Field # 823
	NSLDSLoanPeriodEndDate5 time.Time

	// Field # 824
	NSLDSLoanGuarantyAgencyCode5 string

	// Field # 825
	NSLDSLoanContactType5 string

	// Field # 826
	NSLDSLoanSchoolCode5 string

	// Field # 827
	NSLDSLoanContactCode5 string

	// Field # 828
	NSLDSLoanGradeLevel5 string

	// Field # 829
	NSLDSLoanAdditionalUnsubsidizedFlag5 string

	// Field # 830
	NSLDSLoanCapitalizedInterestFlag5 string

	// Field # 831
	NSLDSLoanDisbursementAmount5 string

	// Field # 832
	NSLDSLoanDisbursementDate5 time.Time

	// Field # 833
	NSLDSLoanConfirmedLoanSubsidyStatus5 string

	// Field # 834
	NSLDSLoanSubsidyStatusDate5 time.Time

	// Field # 836
	NSLDSLoanSequenceNumber6 string

	// Field # 837
	NSLDSLoanDefaultedRecentIndicator6 string

	// Field # 838
	NSLDSLoanChangeFlag6 string

	// Field # 839
	NSLDSLoanTypeCode6 string

	// Field # 840
	NSLDSLoanNetAmount6 string

	// Field # 841
	NSLDSLoanCurrentStatusCode6 string

	// Field # 842
	NSLDSLoanCurrentStatusDate6 time.Time

	// Field # 843
	NSLDSLoanOutstandingPrincipalBalance6 string

	// Field # 844
	NSLDSLoanOutstandingPrincipalBalanceDate6 time.Time

	// Field # 845
	NSLDSLoanPeriodBeginDate6 time.Time

	// Field # 846
	NSLDSLoanPeriodEndDate6 time.Time

	// Field # 847
	NSLDSLoanGuarantyAgencyCode6 string

	// Field # 848
	NSLDSLoanContactType6 string

	// Field # 849
	NSLDSLoanSchoolCode6 string

	// Field # 850
	NSLDSLoanContactCode6 string

	// Field # 851
	NSLDSLoanGradeLevel6 string

	// Field # 852
	NSLDSLoanAdditionalUnsubsidizedFlag6 string

	// Field # 853
	NSLDSLoanCapitalizedInterestFlag6 string

	// Field # 854
	NSLDSLoanDisbursementAmount6 string

	// Field # 855
	NSLDSLoanDisbursementDate6 time.Time

	// Field # 856
	NSLDSLoanConfirmedLoanSubsidyStatus6 string

	// Field # 857
	NSLDSLoanSubsidyStatusDate6 time.Time

	// Field # 861
	FTILabelStart string

	// Field # 862
	StudentFTIMReturnedTaxYear string

	// Field # 863
	StudentFTIMFilingStatusCode string

	// Field # 864
	StudentFTIMAdjustedGrossIncome string

	// Field # 865
	StudentFTIMNumberOfExemptions string

	// Field # 866
	StudentFTIMNumberOfDependents string

	// Field # 867
	StudentFTIMTotalIncomeEarnedAmount string

	// Field # 868
	StudentFTIMIncomeTaxPaid string

	// Field # 869
	StudentFTIMEducationCredits string

	// Field # 870
	StudentFTIMUntaxedIRADistributions string

	// Field # 871
	StudentFTIMIRADeductibleAndPayments string

	// Field # 872
	StudentFTIMTaxExemptInterest string

	// Field # 873
	StudentFTIMUntaxedPensionsAmount string

	// Field # 874
	StudentFTIMScheduleCNetProfitLoss string

	// Field # 875
	StudentFTIMScheduleAIndicator string

	// Field # 876
	StudentFTIMScheduleBIndicator string

	// Field # 877
	StudentFTIMScheduleDIndicator string

	// Field # 878
	StudentFTIMScheduleEIndicator string

	// Field # 879
	StudentFTIMScheduleFIndicator string

	// Field # 880
	StudentFTIMScheduleHIndicator string

	// Field # 881
	StudentFTIMIRSResponseCode string

	// Field # 882
	StudentFTIMSpouseReturnedTaxYear string

	// Field # 883
	StudentFTIMSpouseFilingStatusCode string

	// Field # 884
	StudentFTIMSpouseAdjustedGrossIncome string

	// Field # 885
	StudentFTIMSpouseNumberOfExemptions string

	// Field # 886
	StudentFTIMSpouseNumberOfDependents string

	// Field # 887
	StudentFTIMSpouseTotalIncomeEarnedAmount string

	// Field # 888
	StudentFTIMSpouseIncomeTaxPaid string

	// Field # 889
	StudentFTIMSpouseEducationCredits string

	// Field # 890
	StudentFTIMSpouseUntaxedIRADistributions string

	// Field # 891
	StudentFTIMSpouseIRADeductibleAndPayments string

	// Field # 892
	StudentFTIMSpouseTaxExemptInterest string

	// Field # 893
	StudentFTIMSpouseUntaxedPensionsAmount string

	// Field # 894
	StudentFTIMSpouseScheduleCNetProfitLoss string

	// Field # 895
	StudentFTIMSpouseScheduleAIndicator string

	// Field # 896
	StudentFTIMSpouseScheduleBIndicator string

	// Field # 897
	StudentFTIMSpouseScheduleDIndicator string

	// Field # 898
	StudentFTIMSpouseScheduleEIndicator string

	// Field # 899
	StudentFTIMSpouseScheduleFIndicator string

	// Field # 900
	StudentFTIMSpouseScheduleHIndicator string

	// Field # 901
	StudentFTIMSpouseIRSResponseCode string

	// Field # 902
	ParentFTIMReturnedTaxYear string

	// Field # 903
	ParentFTIMFilingStatusCode string

	// Field # 904
	ParentFTIMAdjustedGrossIncome string

	// Field # 905
	ParentFTIMNumberOfExemptions string

	// Field # 906
	ParentFTIMNumberOfDependents string

	// Field # 907
	ParentFTIMTotalIncomeEarnedAmount string

	// Field # 908
	ParentFTIMIncomeTaxPaid string

	// Field # 909
	ParentFTIMEducationCredits string

	// Field # 910
	ParentFTIMUntaxedIRADistributions string

	// Field # 911
	ParentFTIMIRADeductibleAndPayments string

	// Field # 912
	ParentFTIMTaxExemptInterest string

	// Field # 913
	ParentFTIMUntaxedPensionsAmount string

	// Field # 914
	ParentFTIMScheduleCNetProfitLoss string

	// Field # 915
	ParentFTIMScheduleAIndicator string

	// Field # 916
	ParentFTIMScheduleBIndicator string

	// Field # 917
	ParentFTIMScheduleDIndicator string

	// Field # 918
	ParentFTIMScheduleEIndicator string

	// Field # 919
	ParentFTIMScheduleFIndicator string

	// Field # 920
	ParentFTIMScheduleHIndicator string

	// Field # 921
	ParentFTIMIRSResponseCode string

	// Field # 922
	ParentFTIMSpouseReturnedTaxYear string

	// Field # 923
	ParentFTIMSpouseFilingStatusCode string

	// Field # 924
	ParentFTIMSpouseAdjustedGrossIncome string

	// Field # 925
	ParentFTIMSpouseNumberOfExemptions string

	// Field # 926
	ParentFTIMSpouseNumberOfDependents string

	// Field # 927
	ParentFTIMSpouseTotalIncomeEarnedAmount string

	// Field # 928
	ParentFTIMSpouseIncomeTaxPaid string

	// Field # 929
	ParentFTIMSpouseEducationCredits string

	// Field # 930
	ParentFTIMSpouseUntaxedIRADistributions string

	// Field # 931
	ParentFTIMSpouseIRADeductibleAndPayments string

	// Field # 932
	ParentFTIMSpouseTaxExemptInterest string

	// Field # 933
	ParentFTIMSpouseUntaxedPensionsAmount string

	// Field # 934
	ParentFTIMSpouseScheduleCNetProfitLoss string

	// Field # 935
	ParentFTIMSpouseScheduleAIndicator string

	// Field # 936
	ParentFTIMSpouseScheduleBIndicator string

	// Field # 937
	ParentFTIMSpouseScheduleDIndicator string

	// Field # 938
	ParentFTIMSpouseScheduleEIndicator string

	// Field # 939
	ParentFTIMSpouseScheduleFIndicator string

	// Field # 940
	ParentFTIMSpouseScheduleHIndicator string

	// Field # 941
	ParentFTIMSpouseIRSResponseCode string

	// Field # 942
	FTILabelEnd string

	// Field # 944
	StudentTotalIncome string

	// Field # 945
	ParentTotalIncome string

	// Field # 946
	FISAPTotalIncome string
}
