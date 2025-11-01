// SPDX-FileCopyrightText: © 2025 Robert Patton robpatton@infiniteskye.com
// SPDX-License-Identifier: Apache-2.0

package isirparser

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rpatton4/fsa/pkg/fsaerrors"
	"github.com/rpatton4/fsa/pkg/fsamodels"
)

// <editor-fold desc="Field Definition Constants">
const isirDateLayout2627 = "20060102"    // CCYYMMDD
const isirDateShortLayout2627 = "200601" // CCYYMM

const totalISIRLength2627 int = 7704

// Field # 1
const yearIndicatorStartIndex2627 int = 1
const yearIndicatorLength2627 int = 1

// Field # 2
const fafsaUUIDStartIndex2627 int = 2
const fafsaUUIDLength2627 int = 36

// Field # 3
const transactionUUIDStartIndex2627 int = 38
const transactionUUIDLength2627 int = 36

// Field # 4
const personUUIDStartIndex2627 int = 74
const personUUIDLength2627 int = 36

// Field # 5
const transactionNumberStartIndex2627 int = 110
const transactionNumberLength2627 int = 2

// Field # 6
const dependencyModelStartIndex2627 int = 112
const dependencyModelLength2627 int = 1

// Field # 7
const applicationSourceStartIndex2627 int = 113
const applicationSourceLength2627 int = 1

// Field # 8
const applicationReceiptDateStartIndex2627 int = 114
const applicationReceiptDateLength2627 int = 8

// Field # 9
const transactionSourceStartIndex2627 int = 122
const transactionSourceLength2627 int = 1

// Field # 10
const transactionTypeStartIndex2627 int = 123
const transactionTypeLength2627 int = 1

// Field # 11
const transactionLanguageStartIndex2627 int = 124
const transactionLanguageLength2627 int = 1

// Field # 12
const transactionReceiptDateStartIndex2627 int = 125
const transactionReceiptDateLength2627 int = 8

// Field # 13
const transactionProcessedDateStartIndex2627 int = 133
const transactionProcessedDateLength2627 int = 8

// Field # 14
const transactionStatusStartIndex2627 int = 141
const transactionStatusLength2627 int = 30

// Field # 15
const renewalDataUsedStartIndex2627 int = 171
const renewalDataUsedLength2627 int = 3

// Field # 16
const fpsCorrectionReasonStartIndex2627 int = 174
const fpsCorrectionReasonLength2627 int = 1

// Field # 17
const saiChangeFlagStartIndex2627 int = 175
const saiChangeFlagLength2627 int = 1

// Field # 18
const saiStartIndex2627 int = 176
const saiLength2627 int = 6

// Field # 19
const provisionalSAIStartIndex2627 int = 182
const provisionalSAILength2627 int = 6

// Field # 20
const saiFormulaStartIndex2627 int = 188
const saiFormulaLength2627 int = 1

// Field # 21
const saiComputationTypeStartIndex2627 int = 189
const saiComputationTypeLength2627 int = 2

// Field # 22
const maxPellIndicatorStartIndex2627 int = 191
const maxPellIndicatorLength2627 int = 1

// Field # 23
const minimumPellIndicatorStartIndex2627 int = 192
const minimumPellIndicatorLength2627 int = 1

// Field # 25
const studentFirstNameStartIndex2627 int = 243
const studentFirstNameLength2627 int = 35

// Field # 26
const studentMiddleNameStartIndex2627 int = 278
const studentMiddleNameLength2627 int = 15

// Field # 27
const studentLastNameStartIndex2627 int = 293
const studentLastNameLength2627 int = 35

// Field # 28
const studentSuffixStartIndex2627 int = 328
const studentSuffixLength2627 int = 10

// Field # 29
const studentDateOfBirthStartIndex2627 int = 338
const studentDateOfBirthLength2627 int = 8

// Field # 30
const studentSSNStartIndex2627 int = 346
const studentSSNLength2627 int = 9

// Field # 31
const studentITINStartIndex2627 int = 355
const studentITINLength2627 int = 9

// Field # 32
const studentPhoneNumberStartIndex2627 int = 364
const studentPhoneNumberLength2627 int = 10

// Field # 33
const studentEmailAddressStartIndex2627 int = 374
const studentEmailAddressLength2627 int = 50

// Field # 34
const studentStreetAddressStartIndex2627 int = 424
const studentStreetAddressLength2627 int = 40

// Field # 35
const studentCityStartIndex2627 int = 464
const studentCityLength2627 int = 30

// Field # 36
const studentStateStartIndex2627 int = 494
const studentStateLength2627 int = 2

// Field # 37
const studentZipCodeStartIndex2627 int = 496
const studentZipCodeLength2627 int = 10

// Field # 38
const studentCountryStartIndex2627 int = 506
const studentCountryLength2627 int = 2

// Field # 40
const studentMaritalStatusStartIndex2627 int = 558
const studentMaritalStatusLength2627 int = 1

// Field # 41
const studentGradeLevelStartIndex2627 int = 559
const studentGradeLevelLength2627 int = 1

// Field # 42
// TODO rename so it is not year specific
const studentFirstBachelorsDegreeBefore2627StartIndex2627 int = 560
const studentFirstBachelorsDegreeBefore2627Length2627 int = 1

// Field # 43
const studentPursuingTeacherCertificationStartIndex2627 int = 561
const studentPursuingTeacherCertificationLength2627 int = 1

// Field # 44
const studentActiveDutyStartIndex2627 int = 562
const studentActiveDutyLength2627 int = 1

// Field # 45
const studentVeteranStartIndex2627 int = 563
const studentVeteranLength2627 int = 1

// Field # 46
const studentChildOrOtherDependentsStartIndex2627 int = 564
const studentChildOrOtherDependentsLength2627 int = 1

// Field # 47
const studentParentsDeceasedStartIndex2627 int = 565
const studentParentsDeceasedLength2627 int = 1

// Field # 48
const studentWardOfCourtStartIndex2627 int = 566
const studentWardOfCourtLength2627 int = 1

// Field # 49
const studentInFosterCareStartIndex2627 int = 567
const studentInFosterCareLength2627 int = 1

// Field # 50
const studentEmancipatedMinorStartIndex2627 int = 568
const studentEmancipatedMinorLength2627 int = 1

// Field # 51
const studentLegalGuardianshipStartIndex2627 int = 569
const studentLegalGuardianshipLength2627 int = 1

// Field # 52
const studentPersonalCircumstancesNoneOfTheAboveStartIndex2627 int = 570
const studentPersonalCircumstancesNoneOfTheAboveLength2627 int = 1

// Field # 53
const studentUnaccompaniedHomelessYouthAndSelfSupportingStartIndex2627 int = 571
const studentUnaccompaniedHomelessYouthAndSelfSupportingLength2627 int = 1

// Field # 54
const studentUnaccompaniedHomelessGeneralStartIndex2627 int = 572
const studentUnaccompaniedHomelessGeneralLength2627 int = 1

// Field # 55
const studentUnaccompaniedHomelessHSStartIndex2627 int = 573
const studentUnaccompaniedHomelessHSLength2627 int = 1

// Field # 56
const studentUnaccompaniedHomelessTRIOStartIndex2627 int = 574
const studentUnaccompaniedHomelessTRIOLength2627 int = 1

// Field # 57
const studentUnaccompaniedHomelessFAAStartIndex2627 int = 575
const studentUnaccompaniedHomelessFAALength2627 int = 1

// Field # 58
const studentHomelessnessNoneOfTheAboveStartIndex2627 int = 576
const studentHomelessnessNoneOfTheAboveLength2627 int = 1

// Field # 59
const studentUnusualCircumstanceStartIndex2627 int = 577
const studentUnusualCircumstanceLength2627 int = 1

// Field # 60
const studentUnsubOnlyStartIndex2627 int = 578
const studentUnsubOnlyLength2627 int = 1

// Field # 61
const studentUpdatedFamilySizeStartIndex2627 int = 579
const studentUpdatedFamilySizeLength2627 int = 2

// Field # 62
const studentNumberInCollegeStartIndex2627 int = 581
const studentNumberInCollegeLength2627 int = 2

// Field # 63
const studentCitizenshipStatusStartIndex2627 int = 583
const studentCitizenshipStatusLength2627 int = 1

// Field # 64
const studentANumberStartIndex2627 int = 584
const studentANumberLength2627 int = 9

// Field # 65
const studentStateOfLegalResidenceStartIndex2627 int = 593
const studentStateOfLegalResidenceLength2627 int = 2

// Field # 66
const studentLegalResidenceDateStartIndex2627 int = 595
const studentLegalResidenceDateLength2627 int = 6

// Field # 67
const studentEitherParentAttendCollegeStartIndex2627 int = 601
const studentEitherParentAttendCollegeLength2627 int = 1

// Field # 68
const studentParentKilledInTheLineOfDutyStartIndex2627 int = 602
const studentParentKilledInTheLineOfDutyLength2627 int = 1

// Field # 69
const studentHighSchoolCompletionStatusStartIndex2627 int = 603
const studentHighSchoolCompletionStatusLength2627 int = 1

// Field # 70
const studentHighSchoolNameStartIndex2627 int = 604
const studentHighSchoolNameLength2627 int = 60

// Field # 71
const studentHighSchoolCityStartIndex2627 int = 664
const studentHighSchoolCityLength2627 int = 28

// Field # 72
const studentHighSchoolStateStartIndex2627 int = 692
const studentHighSchoolStateLength2627 int = 2

// Field # 73
const studentHighSchoolEquivalentDiplomaNameStartIndex2627 int = 694
const studentHighSchoolEquivalentDiplomaNameLength2627 int = 1

// Field # 74
const studentHighSchoolEquivalentDiplomaStateStartIndex2627 int = 695
const studentHighSchoolEquivalentDiplomaStateLength2627 int = 2

// Field # 75
const studentManuallyEnteredReceivedEITCStartIndex2627 int = 697
const studentManuallyEnteredReceivedEITCLength2627 int = 1

// Field # 76
const studentManuallyEnteredReceivedFederalHousingAssistanceStartIndex2627 int = 698
const studentManuallyEnteredReceivedFederalHousingAssistanceLength2627 int = 1

// Field # 77
const studentManuallyEnteredReceivedFreeReducedPriceLunchStartIndex2627 int = 699
const studentManuallyEnteredReceivedFreeReducedPriceLunchLength2627 int = 1

// Field # 78
const studentManuallyEnteredReceivedMedicaidStartIndex2627 int = 700
const studentManuallyEnteredReceivedMedicaidLength2627 int = 1

// Field # 79
const studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanStartIndex2627 int = 701
const studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanLength2627 int = 1

// Field # 80
const studentManuallyEnteredReceivedSNAPStartIndex2627 int = 702
const studentManuallyEnteredReceivedSNAPLength2627 int = 1

// Field # 81
const studentManuallyEnteredReceivedSupplementalSecurityIncomeStartIndex2627 int = 703
const studentManuallyEnteredReceivedSupplementalSecurityIncomeLength2627 int = 1

// Field # 82
const studentManuallyEnteredReceivedTANFStartIndex2627 int = 704
const studentManuallyEnteredReceivedTANFLength2627 int = 1

// Field # 83
const studentManuallyEnteredReceivedWICStartIndex2627 int = 705
const studentManuallyEnteredReceivedWICLength2627 int = 1

// Field # 84
const studentManuallyEnteredFederalBenefitsNoneOfTheAboveStartIndex2627 int = 706
const studentManuallyEnteredFederalBenefitsNoneOfTheAboveLength2627 int = 1

// Field # 85
const studentManuallyEnteredFiled1040Or1040NRStartIndex2627 int = 707
const studentManuallyEnteredFiled1040Or1040NRLength2627 int = 1

// Field # 86
const studentManuallyEnteredFiledNonUSTaxReturnStartIndex2627 int = 708
const studentManuallyEnteredFiledNonUSTaxReturnLength2627 int = 1

// Field # 87
const studentManuallyEnteredFiledJointReturnWithCurrentSpouseStartIndex2627 int = 709
const studentManuallyEnteredFiledJointReturnWithCurrentSpouseLength2627 int = 1

// Field # 88
const studentManuallyEnteredTaxReturnFilingStatusStartIndex2627 int = 710
const studentManuallyEnteredTaxReturnFilingStatusLength2627 int = 1

// Field # 89
const studentManuallyEnteredIncomeEarnedFromWorkStartIndex2627 int = 711
const studentManuallyEnteredIncomeEarnedFromWorkLength2627 int = 11

// Field # 90
const studentManuallyEnteredTaxExemptInterestIncomeStartIndex2627 int = 722
const studentManuallyEnteredTaxExemptInterestIncomeLength2627 int = 11

// Field # 91
const studentManuallyEnteredUntaxedPortionsOfIRADistributionsStartIndex2627 int = 733
const studentManuallyEnteredUntaxedPortionsOfIRADistributionsLength2627 int = 11

// Field # 92
const studentManuallyEnteredIRARolloverStartIndex2627 int = 744
const studentManuallyEnteredIRARolloverLength2627 int = 11

// Field # 93
const studentManuallyEnteredUntaxedPortionsOfPensionsStartIndex2627 int = 755
const studentManuallyEnteredUntaxedPortionsOfPensionsLength2627 int = 11

// Field # 94
const studentManuallyEnteredPensionRolloverStartIndex2627 int = 766
const studentManuallyEnteredPensionRolloverLength2627 int = 11

// Field # 95
const studentManuallyEnteredAdjustedGrossIncomeStartIndex2627 int = 777
const studentManuallyEnteredAdjustedGrossIncomeLength2627 int = 10

// Field # 96
const studentManuallyEnteredIncomeTaxPaidStartIndex2627 int = 787
const studentManuallyEnteredIncomeTaxPaidLength2627 int = 9

// Field # 97
const studentManuallyEnteredEITCReceivedDuringTaxYearStartIndex2627 int = 796
const studentManuallyEnteredEITCReceivedDuringTaxYearLength2627 int = 1

// Field # 98
const studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherStartIndex2627 int = 797
const studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherLength2627 int = 11

// Field # 99
const studentManuallyEnteredEducationCreditsStartIndex2627 int = 808
const studentManuallyEnteredEducationCreditsLength2627 int = 9

// Field # 100
const studentManuallyEnteredFiledScheduleABDEFHStartIndex2627 int = 817
const studentManuallyEnteredFiledScheduleABDEFHLength2627 int = 1

// Field # 101
const studentManuallyEnteredScheduleCAmountStartIndex2627 int = 818
const studentManuallyEnteredScheduleCAmountLength2627 int = 12

// Field # 102
const studentManuallyEnteredCollegeGrantAndScholarshipAidStartIndex2627 int = 830
const studentManuallyEnteredCollegeGrantAndScholarshipAidLength2627 int = 7

// Field # 103
const studentManuallyEnteredForeignEarnedIncomeExclusionStartIndex2627 int = 837
const studentManuallyEnteredForeignEarnedIncomeExclusionLength2627 int = 10

// Field # 104
const studentManuallyEnteredChildSupportReceivedStartIndex2627 int = 847
const studentManuallyEnteredChildSupportReceivedLength2627 int = 7

// Field # 105
const studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsStartIndex2627 int = 854
const studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsLength2627 int = 7

// Field # 106
const studentManuallyEnteredNetWorthOfCurrentInvestmentsStartIndex2627 int = 861
const studentManuallyEnteredNetWorthOfCurrentInvestmentsLength2627 int = 7

// Field # 107
const studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsStartIndex2627 int = 868
const studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsLength2627 int = 7

// Field # 108
const studentCollege1StartIndex2627 int = 875
const studentCollege1Length2627 int = 6

// Field # 109
const studentCollege2StartIndex2627 int = 881
const studentCollege2Length2627 int = 6

// Field # 110
const studentCollege3StartIndex2627 int = 887
const studentCollege3Length2627 int = 6

// Field # 111
const studentCollege4StartIndex2627 int = 893
const studentCollege4Length2627 int = 6

// Field # 112
const studentCollege5StartIndex2627 int = 899
const studentCollege5Length2627 int = 6

// Field # 113
const studentCollege6StartIndex2627 int = 905
const studentCollege6Length2627 int = 6

// Field # 114
const studentCollege7StartIndex2627 int = 911
const studentCollege7Length2627 int = 6

// Field # 115
const studentCollege8StartIndex2627 int = 917
const studentCollege8Length2627 int = 6

// Field # 116
const studentCollege9StartIndex2627 int = 923
const studentCollege9Length2627 int = 6

// Field # 117
const studentCollege10StartIndex2627 int = 929
const studentCollege10Length2627 int = 6

// Field # 118
const studentCollege11StartIndex2627 int = 935
const studentCollege11Length2627 int = 6

// Field # 119
const studentCollege12StartIndex2627 int = 941
const studentCollege12Length2627 int = 6

// Field # 120
const studentCollege13StartIndex2627 int = 947
const studentCollege13Length2627 int = 6

// Field # 121
const studentCollege14StartIndex2627 int = 953
const studentCollege14Length2627 int = 6

// Field # 122
const studentCollege15StartIndex2627 int = 959
const studentCollege15Length2627 int = 6

// Field # 123
const studentCollege16StartIndex2627 int = 965
const studentCollege16Length2627 int = 6

// Field # 124
const studentCollege17StartIndex2627 int = 971
const studentCollege17Length2627 int = 6

// Field # 125
const studentCollege18StartIndex2627 int = 977
const studentCollege18Length2627 int = 6

// Field # 126
const studentCollege19StartIndex2627 int = 983
const studentCollege19Length2627 int = 6

// Field # 127
const studentCollege20StartIndex2627 int = 989
const studentCollege20Length2627 int = 6

// Field # 128
const studentConsentToRetrieveAndDiscloseFTIStartIndex2627 int = 995
const studentConsentToRetrieveAndDiscloseFTILength2627 int = 1

// Field # 129
const studentSignatureStartIndex2627 int = 996
const studentSignatureLength2627 int = 1

// Field # 130
const studentSignatureDateStartIndex2627 int = 997
const studentSignatureDateLength2627 int = 8

// Field # 132
const studentSpouseFirstNameStartIndex2627 int = 1055
const studentSpouseFirstNameLength2627 int = 35

// Field # 133
const studentSpouseMiddleNameStartIndex2627 int = 1090
const studentSpouseMiddleNameLength2627 int = 15

// Field # 134
const studentSpouseLastNameStartIndex2627 int = 1105
const studentSpouseLastNameLength2627 int = 35

// Field # 135
const studentSpouseSuffixStartIndex2627 int = 1140
const studentSpouseSuffixLength2627 int = 10

// Field # 136
const studentSpouseDateOfBirthStartIndex2627 int = 1150
const studentSpouseDateOfBirthLength2627 int = 8

// Field # 137
const studentSpouseSSNStartIndex2627 int = 1158
const studentSpouseSSNLength2627 int = 9

// Field # 138
const studentSpouseITINStartIndex2627 int = 1167
const studentSpouseITINLength2627 int = 9

// Field # 139
const studentSpousePhoneNumberStartIndex2627 int = 1176
const studentSpousePhoneNumberLength2627 int = 10

// Field # 140
const studentSpouseEmailAddressStartIndex2627 int = 1186
const studentSpouseEmailAddressLength2627 int = 50

// Field # 141
const studentSpouseStreetAddressStartIndex2627 int = 1236
const studentSpouseStreetAddressLength2627 int = 40

// Field # 142
const studentSpouseCityStartIndex2627 int = 1276
const studentSpouseCityLength2627 int = 30

// Field # 143
const studentSpouseStateStartIndex2627 int = 1306
const studentSpouseStateLength2627 int = 2

// Field # 144
const studentSpouseZipCodeStartIndex2627 int = 1308
const studentSpouseZipCodeLength2627 int = 10

// Field # 145
const studentSpouseCountryStartIndex2627 int = 1318
const studentSpouseCountryLength2627 int = 2

// Field # 146
const studentSpouseFiled1040Or1040NRStartIndex2627 int = 1320
const studentSpouseFiled1040Or1040NRLength2627 int = 1

// Field # 147
const studentSpouseFiledNonUSTaxReturnStartIndex2627 int = 1321
const studentSpouseFiledNonUSTaxReturnLength2627 int = 1

// Field # 148
const studentSpouseTaxReturnFilingStatusStartIndex2627 int = 1322
const studentSpouseTaxReturnFilingStatusLength2627 int = 1

// Field # 149
const studentSpouseIncomeEarnedFromWorkStartIndex2627 int = 1323
const studentSpouseIncomeEarnedFromWorkLength2627 int = 11

// Field # 150
const studentSpouseTaxExemptInterestIncomeStartIndex2627 int = 1334
const studentSpouseTaxExemptInterestIncomeLength2627 int = 11

// Field # 151
const studentSpouseUntaxedPortionsOfIRADistributionsStartIndex2627 int = 1345
const studentSpouseUntaxedPortionsOfIRADistributionsLength2627 int = 11

// Field # 152
const studentSpouseIRARolloverStartIndex2627 int = 1356
const studentSpouseIRARolloverLength2627 int = 11

// Field # 153
const studentSpouseUntaxedPortionsOfPensionsStartIndex2627 int = 1367
const studentSpouseUntaxedPortionsOfPensionsLength2627 int = 11

// Field # 154
const studentSpousePensionRolloverStartIndex2627 int = 1378
const studentSpousePensionRolloverLength2627 int = 11

// Field # 155
const studentSpouseAdjustedGrossIncomeStartIndex2627 int = 1389
const studentSpouseAdjustedGrossIncomeLength2627 int = 10

// Field # 156
const studentSpouseIncomeTaxPaidStartIndex2627 int = 1399
const studentSpouseIncomeTaxPaidLength2627 int = 9

// Field # 157
const studentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex2627 int = 1408
const studentSpouseDeductiblePaymentsToIRAKeoghOtherLength2627 int = 11

// Field # 158
const studentSpouseEducationCreditsStartIndex2627 int = 1419
const studentSpouseEducationCreditsLength2627 int = 9

// Field # 159
const studentSpouseFiledScheduleABDEFHStartIndex2627 int = 1428
const studentSpouseFiledScheduleABDEFHLength2627 int = 1

// Field # 160
const studentSpouseScheduleCAmountStartIndex2627 int = 1429
const studentSpouseScheduleCAmountLength2627 int = 12

// Field # 161
const studentSpouseForeignEarnedIncomeExclusionStartIndex2627 int = 1441
const studentSpouseForeignEarnedIncomeExclusionLength2627 int = 10

// Field # 162
const studentSpouseConsentToRetrieveAndDiscloseFTIStartIndex2627 int = 1451
const studentSpouseConsentToRetrieveAndDiscloseFTILength2627 int = 1

// Field # 163
const studentSpouseSignatureStartIndex2627 int = 1452
const studentSpouseSignatureLength2627 int = 1

// Field # 164
const studentSpouseSignatureDateStartIndex2627 int = 1453
const studentSpouseSignatureDateLength2627 int = 8

// Field # 166
const parentFirstNameStartIndex2627 int = 1511
const parentFirstNameLength2627 int = 35

// Field # 167
const parentMiddleNameStartIndex2627 int = 1546
const parentMiddleNameLength2627 int = 15

// Field # 168
const parentLastNameStartIndex2627 int = 1561
const parentLastNameLength2627 int = 35

// Field # 169
const parentSuffixStartIndex2627 int = 1596
const parentSuffixLength2627 int = 10

// Field # 170
const parentDateOfBirthStartIndex2627 int = 1606
const parentDateOfBirthLength2627 int = 8

// Field # 171
const parentSSNStartIndex2627 int = 1614
const parentSSNLength2627 int = 9

// Field # 172
const parentITINStartIndex2627 int = 1623
const parentITINLength2627 int = 9

// Field # 173
const parentPhoneNumberStartIndex2627 int = 1632
const parentPhoneNumberLength2627 int = 10

// Field # 174
const parentEmailAddressStartIndex2627 int = 1642
const parentEmailAddressLength2627 int = 50

// Field # 175
const parentStreetAddressStartIndex2627 int = 1692
const parentStreetAddressLength2627 int = 40

// Field # 176
const parentCityStartIndex2627 int = 1732
const parentCityLength2627 int = 30

// Field # 177
const parentStateStartIndex2627 int = 1762
const parentStateLength2627 int = 2

// Field # 178
const parentZipCodeStartIndex2627 int = 1764
const parentZipCodeLength2627 int = 10

// Field # 179
const parentCountryStartIndex2627 int = 1774
const parentCountryLength2627 int = 2

// Field # 180
const parentMaritalStatusStartIndex2627 int = 1776
const parentMaritalStatusLength2627 int = 1

// Field # 181
const parentStateOfLegalResidenceStartIndex2627 int = 1777
const parentStateOfLegalResidenceLength2627 int = 2

// Field # 182
const parentLegalResidenceDateStartIndex2627 int = 1779
const parentLegalResidenceDateLength2627 int = 6

// Field # 183
const parentUpdatedFamilySizeStartIndex2627 int = 1785
const parentUpdatedFamilySizeLength2627 int = 2

// Field # 184
const parentNumberInCollegeStartIndex2627 int = 1787
const parentNumberInCollegeLength2627 int = 2

// Field # 185
const parentReceivedEITCStartIndex2627 int = 1789
const parentReceivedEITCLength2627 int = 1

// Field # 186
const parentReceivedFederalHousingAssistanceStartIndex2627 int = 1790
const parentReceivedFederalHousingAssistanceLength2627 int = 1

// Field # 187
const parentReceivedFreeReducedPriceLunchStartIndex2627 int = 1791
const parentReceivedFreeReducedPriceLunchLength2627 int = 1

// Field # 188
const parentReceivedMedicaidStartIndex2627 int = 1792
const parentReceivedMedicaidLength2627 int = 1

// Field # 189
const parentReceivedRefundableCreditFor36BHealthPlanStartIndex2627 int = 1793
const parentReceivedRefundableCreditFor36BHealthPlanLength2627 int = 1

// Field # 190
const parentReceivedSNAPStartIndex2627 int = 1794
const parentReceivedSNAPLength2627 int = 1

// Field # 191
const parentReceivedSupplementalSecurityIncomeStartIndex2627 int = 1795
const parentReceivedSupplementalSecurityIncomeLength2627 int = 1

// Field # 192
const parentReceivedTANFStartIndex2627 int = 1796
const parentReceivedTANFLength2627 int = 1

// Field # 193
const parentReceivedWICStartIndex2627 int = 1797
const parentReceivedWICLength2627 int = 1

// Field # 194
const parentFederalBenefitsNoneOfTheAboveStartIndex2627 int = 1798
const parentFederalBenefitsNoneOfTheAboveLength2627 int = 1

// Field # 195
const parentFiled1040Or1040NRStartIndex2627 int = 1799
const parentFiled1040Or1040NRLength2627 int = 1

// Field # 196
const parentFileNonUSTaxReturnStartIndex2627 int = 1800
const parentFileNonUSTaxReturnLength2627 int = 1

// Field # 197
const parentFiledJointReturnWithCurrentSpouseStartIndex2627 int = 1801
const parentFiledJointReturnWithCurrentSpouseLength2627 int = 1

// Field # 198
const parentTaxReturnFilingStatusStartIndex2627 int = 1802
const parentTaxReturnFilingStatusLength2627 int = 1

// Field # 199
const parentIncomeEarnedFromWorkStartIndex2627 int = 1803
const parentIncomeEarnedFromWorkLength2627 int = 11

// Field # 200
const parentTaxExemptInterestIncomeStartIndex2627 int = 1814
const parentTaxExemptInterestIncomeLength2627 int = 11

// Field # 201
const parentUntaxedPortionsOfIRADistributionsStartIndex2627 int = 1825
const parentUntaxedPortionsOfIRADistributionsLength2627 int = 11

// Field # 202
const parentIRARolloverStartIndex2627 int = 1836
const parentIRARolloverLength2627 int = 11

// Field # 203
const parentUntaxedPortionsOfPensionsStartIndex2627 int = 1847
const parentUntaxedPortionsOfPensionsLength2627 int = 11

// Field # 204
const parentPensionRolloverStartIndex2627 int = 1858
const parentPensionRolloverLength2627 int = 11

// Field # 205
const parentAdjustedGrossIncomeStartIndex2627 int = 1869
const parentAdjustedGrossIncomeLength2627 int = 10

// Field # 206
const parentIncomeTaxPaidStartIndex2627 int = 1879
const parentIncomeTaxPaidLength2627 int = 9

// Field # 207
const parentEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex2627 int = 1888
const parentEarnedIncomeTaxCreditReceivedDuringTaxYearLength2627 int = 1

// Field # 208
const parentDeductiblePaymentsToIRAKeoghOtherStartIndex2627 int = 1889
const parentDeductiblePaymentsToIRAKeoghOtherLength2627 int = 11

// Field # 209
const parentEducationCreditsStartIndex2627 int = 1900
const parentEducationCreditsLength2627 int = 9

// Field # 210
const parentFiledScheduleABDEFHStartIndex2627 int = 1909
const parentFiledScheduleABDEFHLength2627 int = 1

// Field # 211
const parentScheduleCAmountStartIndex2627 int = 1910
const parentScheduleCAmountLength2627 int = 12

// Field # 212
const parentCollegeGrantAndScholarshipAidStartIndex2627 int = 1922
const parentCollegeGrantAndScholarshipAidLength2627 int = 7

// Field # 213
const parentForeignEarnedIncomeExclusionStartIndex2627 int = 1929
const parentForeignEarnedIncomeExclusionLength2627 int = 10

// Field # 214
const parentChildSupportReceivedStartIndex2627 int = 1939
const parentChildSupportReceivedLength2627 int = 7

// Field # 215
const parentTotalOfCashSavingsAndCheckingAccountsStartIndex2627 int = 1946
const parentTotalOfCashSavingsAndCheckingAccountsLength2627 int = 7

// Field # 216
const parentNetWorthOfCurrentInvestmentsStartIndex2627 int = 1953
const parentNetWorthOfCurrentInvestmentsLength2627 int = 7

// Field # 217
const parentNetWorthOfBusinessesAndInvestmentFarmsStartIndex2627 int = 1960
const parentNetWorthOfBusinessesAndInvestmentFarmsLength2627 int = 7

// Field # 218
const parentConsentToRetrieveAndDiscloseFTIStartIndex2627 int = 1967
const parentConsentToRetrieveAndDiscloseFTILength2627 int = 1

// Field # 219
const parentSignatureStartIndex2627 int = 1968
const parentSignatureLength2627 int = 1

// Field # 220
const parentSignatureDateStartIndex2627 int = 1969
const parentSignatureDateLength2627 int = 8

// Field # 222
const parentSpouseFirstNameStartIndex2627 int = 2027
const parentSpouseFirstNameLength2627 int = 35

// Field # 223
const parentSpouseMiddleNameStartIndex2627 int = 2062
const parentSpouseMiddleNameLength2627 int = 15

// Field # 224
const parentSpouseLastNameStartIndex2627 int = 2077
const parentSpouseLastNameLength2627 int = 35

// Field # 225
const parentSpouseSuffixStartIndex2627 int = 2112
const parentSpouseSuffixLength2627 int = 10

// Field # 226
const parentSpouseDateOfBirthStartIndex2627 int = 2122
const parentSpouseDateOfBirthLength2627 int = 8

// Field # 227
const parentSpouseSSNStartIndex2627 int = 2130
const parentSpouseSSNLength2627 int = 9

// Field # 228
const parentSpouseITINStartIndex2627 int = 2139
const parentSpouseITINLength2627 int = 9

// Field # 229
const parentSpousePhoneNumberStartIndex2627 int = 2148
const parentSpousePhoneNumberLength2627 int = 10

// Field # 230
const parentSpouseEmailAddressStartIndex2627 int = 2158
const parentSpouseEmailAddressLength2627 int = 50

// Field # 231
const parentSpouseStreetAddressStartIndex2627 int = 2208
const parentSpouseStreetAddressLength2627 int = 40

// Field # 232
const parentSpouseCityStartIndex2627 int = 2248
const parentSpouseCityLength2627 int = 30

// Field # 233
const parentSpouseStateStartIndex2627 int = 2278
const parentSpouseStateLength2627 int = 2

// Field # 234
const parentSpouseZipCodeStartIndex2627 int = 2280
const parentSpouseZipCodeLength2627 int = 10

// Field # 235
const parentSpouseCountryStartIndex2627 int = 2290
const parentSpouseCountryLength2627 int = 2

// Field # 236
const parentSpouseFiled1040Or1040NRStartIndex2627 int = 2292
const parentSpouseFiled1040Or1040NRLength2627 int = 1

// Field # 237
const parentSpouseFileNonUSTaxReturnStartIndex2627 int = 2293
const parentSpouseFileNonUSTaxReturnLength2627 int = 1

// Field # 238
const parentSpouseTaxReturnFilingStatusStartIndex2627 int = 2294
const parentSpouseTaxReturnFilingStatusLength2627 int = 1

// Field # 239
const parentSpouseIncomeEarnedFromWorkStartIndex2627 int = 2295
const parentSpouseIncomeEarnedFromWorkLength2627 int = 11

// Field # 240
const parentSpouseTaxExemptInterestIncomeStartIndex2627 int = 2306
const parentSpouseTaxExemptInterestIncomeLength2627 int = 11

// Field # 241
const parentSpouseUntaxedPortionsOfIRADistributionsStartIndex2627 int = 2317
const parentSpouseUntaxedPortionsOfIRADistributionsLength2627 int = 11

// Field # 242
const parentSpouseIRARolloverStartIndex2627 int = 2328
const parentSpouseIRARolloverLength2627 int = 11

// Field # 243
const parentSpouseUntaxedPortionsOfPensionsStartIndex2627 int = 2339
const parentSpouseUntaxedPortionsOfPensionsLength2627 int = 11

// Field # 244
const parentSpousePensionRolloverStartIndex2627 int = 2350
const parentSpousePensionRolloverLength2627 int = 11

// Field # 245
const parentSpouseAdjustedGrossIncomeStartIndex2627 int = 2361
const parentSpouseAdjustedGrossIncomeLength2627 int = 10

// Field # 246
const parentSpouseIncomeTaxPaidStartIndex2627 int = 2371
const parentSpouseIncomeTaxPaidLength2627 int = 9

// Field # 247
const parentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex2627 int = 2380
const parentSpouseDeductiblePaymentsToIRAKeoghOtherLength2627 int = 11

// Field # 248
const parentSpouseEducationCreditsStartIndex2627 int = 2391
const parentSpouseEducationCreditsLength2627 int = 9

// Field # 249
const parentSpouseFiledScheduleABDEFHStartIndex2627 int = 2400
const parentSpouseFiledScheduleABDEFHLength2627 int = 1

// Field # 250
const parentSpouseScheduleCAmountStartIndex2627 int = 2401
const parentSpouseScheduleCAmountLength2627 int = 12

// Field # 251
const parentSpouseForeignEarnedIncomeExclusionStartIndex2627 int = 2413
const parentSpouseForeignEarnedIncomeExclusionLength2627 int = 10

// Field # 252
const parentSpouseConsentToRetrieveAndDiscloseFTIStartIndex2627 int = 2423
const parentSpouseConsentToRetrieveAndDiscloseFTILength2627 int = 1

// Field # 253
const parentSpouseSignatureStartIndex2627 int = 2424
const parentSpouseSignatureLength2627 int = 1

// Field # 254
const parentSpouseSignatureDateStartIndex2627 int = 2425
const parentSpouseSignatureDateLength2627 int = 8

// Field # 256
const preparerFirstNameStartIndex2627 int = 2483
const preparerFirstNameLength2627 int = 35

// Field # 257
const preparerLastNameStartIndex2627 int = 2518
const preparerLastNameLength2627 int = 35

// Field # 258
const preparerSSNStartIndex2627 int = 2553
const preparerSSNLength2627 int = 9

// Field # 259
const preparerEINStartIndex2627 int = 2562
const preparerEINLength2627 int = 9

// Field # 260
const preparerAffiliationStartIndex2627 int = 2571
const preparerAffiliationLength2627 int = 30

// Field # 261
const preparerStreetAddressStartIndex2627 int = 2601
const preparerStreetAddressLength2627 int = 40

// Field # 262
const preparerCityStartIndex2627 int = 2641
const preparerCityLength2627 int = 30

// Field # 263
const preparerStateStartIndex2627 int = 2671
const preparerStateLength2627 int = 2

// Field # 264
const preparerZipCodeStartIndex2627 int = 2673
const preparerZipCodeLength2627 int = 10

// Field # 265
const preparerSignatureStartIndex2627 int = 2683
const preparerSignatureLength2627 int = 1

// Field # 266
const preparerSignatureDateStartIndex2627 int = 2684
const preparerSignatureDateLength2627 int = 8

// Field # 268
const studentAffirmationStatusStartIndex2627 int = 2742
const studentAffirmationStatusLength2627 int = 1

// Field # 269
const studentSpouseAffirmationStatusStartIndex2627 int = 2743
const studentSpouseAffirmationStatusLength2627 int = 1

// Field # 270
const parentAffirmationStatusStartIndex2627 int = 2744
const parentAffirmationStatusLength2627 int = 1

// Field # 271
const parentSpouseOrPartnerAffirmationStatusStartIndex2627 int = 2745
const parentSpouseOrPartnerAffirmationStatusLength2627 int = 1

// Field # 272
const studentDateConsentGrantedStartIndex2627 int = 2746
const studentDateConsentGrantedLength2627 int = 8

// Field # 273
const studentSpouseDateConsentGrantedStartIndex2627 int = 2754
const studentSpouseDateConsentGrantedLength2627 int = 8

// Field # 274
const parentDateConsentGrantedStartIndex2627 int = 2762
const parentDateConsentGrantedLength2627 int = 8

// Field # 275
const parentSpouseOrPartnerDateConsentGrantedStartIndex2627 int = 2770
const parentSpouseOrPartnerDateConsentGrantedLength2627 int = 8

// Field # 276
const studentTransunionMatchStatusStartIndex2627 int = 2778
const studentTransunionMatchStatusLength2627 int = 1

// Field # 277
const studentSpouseTransunionMatchStatusStartIndex2627 int = 2779
const studentSpouseTransunionMatchStatusLength2627 int = 1

// Field # 278
const studentParentTransunionMatchStatusStartIndex2627 int = 2780
const studentParentTransunionMatchStatusLength2627 int = 1

// Field # 279
const studentParentSpouseTransunionMatchStatusStartIndex2627 int = 2781
const studentParentSpouseTransunionMatchStatusLength2627 int = 1

// Field # 280
const correctionAppliedAgainstTransactionNumberStartIndex2627 int = 2782
const correctionAppliedAgainstTransactionNumberLength2627 int = 2

// Field # 281
const professionalJudgementStartIndex2627 int = 2784
const professionalJudgementLength2627 int = 1

// Field # 282
const dependencyOverrideIndicatorStartIndex2627 int = 2785
const dependencyOverrideIndicatorLength2627 int = 1

// Field # 283
const fAAFederalSchoolCodeStartIndex2627 int = 2786
const fAAFederalSchoolCodeLength2627 int = 6

// Field # 284
const fAASignatureStartIndex2627 int = 2792
const fAASignatureLength2627 int = 1

// Field # 285
const iASGIndicatorStartIndex2627 int = 2793
const iASGIndicatorLength2627 int = 1

// Field # 286
const childrenOfFallenHeroesIndicatorStartIndex2627 int = 2794
const childrenOfFallenHeroesIndicatorLength2627 int = 1

// Field # 287
const electronicTransactionIndicatorDestinationNumberStartIndex2627 int = 2795
const electronicTransactionIndicatorDestinationNumberLength2627 int = 7

// Field # 288
const studentSignatureSourceStartIndex2627 int = 2802
const studentSignatureSourceLength2627 int = 1

// Field # 289
const studentSpouseSignatureSourceStartIndex2627 int = 2803
const studentSpouseSignatureSourceLength2627 int = 1

// Field # 290
const parentSignatureSourceStartIndex2627 int = 2804
const parentSignatureSourceLength2627 int = 1

// Field # 291
const parentSpouseOrPartnerSignatureSourceStartIndex2627 int = 2805
const parentSpouseOrPartnerSignatureSourceLength2627 int = 1

// Field # 292
const specialHandlingIndicatorStartIndex2627 int = 2806
const specialHandlingIndicatorLength2627 int = 1

// Field # 293
const addressOnlyChangeFlagStartIndex2627 int = 2807
const addressOnlyChangeFlagLength2627 int = 1

// Field # 294
const fpsPushedISIRFlagStartIndex2627 int = 2808
const fpsPushedISIRFlagLength2627 int = 1

// Field # 295
const rejectStatusChangeFlagStartIndex2627 int = 2809
const rejectStatusChangeFlagLength2627 int = 1

// Field # 296
const verificationTrackingFlagStartIndex2627 int = 2810
const verificationTrackingFlagLength2627 int = 2

// Field # 297
const studentSelectedForVerificationStartIndex2627 int = 2812
const studentSelectedForVerificationLength2627 int = 1

// Field # 298
const incarceratedApplicantFlagStartIndex2627 int = 2813
const incarceratedApplicantFlagLength2627 int = 1

// Field # 299
const nsldsTransactionNumberStartIndex2627 int = 2814
const nsldsTransactionNumberLength2627 int = 2

// Field # 300
const nsldsDatabaseResultsFlagStartIndex2627 int = 2816
const nsldsDatabaseResultsFlagLength2627 int = 1

// Field # 301
const highSchoolFlagStartIndex2627 int = 2817
const highSchoolFlagLength2627 int = 1

// Field # 302
const studentTotalFederalWorkStudyEarningsStartIndex2627 int = 2818
const studentTotalFederalWorkStudyEarningsLength2627 int = 12

// Field # 303
const studentSpouseTotalFederalWorkStudyEarningsStartIndex2627 int = 2830
const studentSpouseTotalFederalWorkStudyEarningsLength2627 int = 12

// Field # 304
const parentTotalFederalWorkStudyEarningsStartIndex2627 int = 2842
const parentTotalFederalWorkStudyEarningsLength2627 int = 12

// Field # 305
const parentSpouseOrPartnerTotalFederalWorkStudyEarningsStartIndex2627 int = 2854
const parentSpouseOrPartnerTotalFederalWorkStudyEarningsLength2627 int = 12

// FILLER BEGINS HERE NOW 2866 to 3105 length 240

// END NEW FILLER

// Field # 307
const parentDiscretionaryNetWorthStartIndex2627 int = 3106
const parentDiscretionaryNetWorthLength2627 int = 7

// Field # 308
const parentNetWorthStartIndex2627 int = 3113
const parentNetWorthLength2627 int = 7

// Field # 309
const parentAssetProtectionAllowanceStartIndex2627 int = 3120
const parentAssetProtectionAllowanceLength2627 int = 12

// Field # 310
const parentContributionFromAssetsStartIndex2627 int = 3132
const parentContributionFromAssetsLength2627 int = 12

// Field # 311
const studentNetWorthStartIndex2627 int = 3144
const studentNetWorthLength2627 int = 7

// Field # 312
const studentAssetProtectionAllowanceStartIndex2627 int = 3151
const studentAssetProtectionAllowanceLength2627 int = 12

// Field # 313
const studentContributionFromAssetsStartIndex2627 int = 3163
const studentContributionFromAssetsLength2627 int = 12

// Field # 314
const assumedStudentFamilySizeStartIndex2627 int = 3175
const assumedStudentFamilySizeLength2627 int = 3

// Field # 315
const assumedParentFamilySizeStartIndex2627 int = 3178
const assumedParentFamilySizeLength2627 int = 3

// Field # 316
const studentFirstNameCHVFlagsStartIndex2627 int = 3181
const studentFirstNameCHVFlagsLength2627 int = 3

// Field # 317
const studentMiddleNameCHVFlagsStartIndex2627 int = 3184
const studentMiddleNameCHVFlagsLength2627 int = 3

// Field # 318
const studentLastNameCHVFLagsStartIndex2627 int = 3187
const studentLastNameCHVFLagsLength2627 int = 3

// Field # 319
const studentSuffixCHVFLagsStartIndex2627 int = 3190
const studentSuffixCHVFLagsLength2627 int = 3

// Field # 320
const studentDateOfBirthCHVFLagsStartIndex2627 int = 3193
const studentDateOfBirthCHVFLagsLength2627 int = 3

// Field # 321
const studentSSNCHVFlagsStartIndex2627 int = 3196
const studentSSNCHVFlagsLength2627 int = 3

// Field # 322
const studentITINCHVFLagsStartIndex2627 int = 3199
const studentITINCHVFLagsLength2627 int = 3

// Field # 323
const studentPhoneNumberCHVFlagsStartIndex2627 int = 3202
const studentPhoneNumberCHVFlagsLength2627 int = 3

// Field # 324
const studentEmailAddressCHVFlagsStartIndex2627 int = 3205
const studentEmailAddressCHVFlagsLength2627 int = 3

// Field # 325
const studentStreetAddressCHVFlagsStartIndex2627 int = 3208
const studentStreetAddressCHVFlagsLength2627 int = 3

// Field # 326
const studentCityCHVFLagsStartIndex2627 int = 3211
const studentCityCHVFLagsLength2627 int = 3

// Field # 327
const studentStateCHVFlagsStartIndex2627 int = 3214
const studentStateCHVFlagsLength2627 int = 3

// Field # 328
const studentZipCodeCHVFlagsStartIndex2627 int = 3217
const studentZipCodeCHVFlagsLength2627 int = 3

// Field # 329
const studentCountryCHVFlagsStartIndex2627 int = 3220
const studentCountryCHVFlagsLength2627 int = 3

// Field # 330
const studentMaritalStatusCHVFlagsStartIndex2627 int = 3223
const studentMaritalStatusCHVFlagsLength2627 int = 3

// Field # 331
const studentGradeLevelInCollegeCHVFlagsStartIndex2627 int = 3226
const studentGradeLevelInCollegeCHVFlagsLength2627 int = 3

// Field # 332
const studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsStartIndex2627 int = 3229
const studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsLength2627 int = 3

// Field # 333
const studentPursuingTeacherCertificationCHVFlagsStartIndex2627 int = 3232
const studentPursuingTeacherCertificationCHVFlagsLength2627 int = 3

// Field # 334
const studentActiveDutyCHVFlagsStartIndex2627 int = 3235
const studentActiveDutyCHVFlagsLength2627 int = 3

// Field # 335
const studentVeteranCHVFlagsStartIndex2627 int = 3238
const studentVeteranCHVFlagsLength2627 int = 3

// Field # 336
const studentChildOrOtherDependentsCHVFlagsStartIndex2627 int = 3241
const studentChildOrOtherDependentsCHVFlagsLength2627 int = 3

// Field # 337
const studentParentsDeceasedCHVFlagsStartIndex2627 int = 3244
const studentParentsDeceasedCHVFlagsLength2627 int = 3

// Field # 338
const studentWardOfCourtCHVFlagsStartIndex2627 int = 3247
const studentWardOfCourtCHVFlagsLength2627 int = 3

// Field # 339
const studentInFosterCareCHVFlagsStartIndex2627 int = 3250
const studentInFosterCareCHVFlagsLength2627 int = 3

// Field # 340
const studentEmancipatedMinorCHVFlagsStartIndex2627 int = 3253
const studentEmancipatedMinorCHVFlagsLength2627 int = 3

// Field # 341
const studentLegalGuardianshipCHVFlagsStartIndex2627 int = 3256
const studentLegalGuardianshipCHVFlagsLength2627 int = 3

// Field # 342
const studentPersonalCircumstancesNoneOfTheAboveCHVFlagsStartIndex2627 int = 3259
const studentPersonalCircumstancesNoneOfTheAboveCHVFlagsLength2627 int = 3

// Field # 343
const studentUnaccompaniedHomelessOrIsUnaccompaniedCHVFlagsStartIndex2627 int = 3262
const studentUnaccompaniedHomelessOrIsUnaccompaniedCHVFlagsLength2627 int = 3

// Field # 344
const studentUnaccompaniedAndHomelessGeneralCHVFlagsStartIndex2627 int = 3265
const studentUnaccompaniedAndHomelessGeneralCHVFlagsLength2627 int = 3

// Field # 345
const studentUnaccompaniedAndHomelessHSCHVFlagsStartIndex2627 int = 3268
const studentUnaccompaniedAndHomelessHSCHVFlagsLength2627 int = 3

// Field # 346
const studentUnaccompaniedAndHomelessTRIOCHVFlagsStartIndex2627 int = 3271
const studentUnaccompaniedAndHomelessTRIOCHVFlagsLength2627 int = 3

// Field # 347
const studentUnaccompaniedAndHomelessFAACHVFlagsStartIndex2627 int = 3274
const studentUnaccompaniedAndHomelessFAACHVFlagsLength2627 int = 3

// Field # 348
const studentHomelessnessNoneOfTheAboveCHVFlagsStartIndex2627 int = 3277
const studentHomelessnessNoneOfTheAboveCHVFlagsLength2627 int = 3

// Field # 349
const studentHasUnusualCircumstanceCHVFlagsStartIndex2627 int = 3280
const studentHasUnusualCircumstanceCHVFlagsLength2627 int = 3

// Field # 350
const studentUnsubOnlyCHVFlagsStartIndex2627 int = 3283
const studentUnsubOnlyCHVFlagsLength2627 int = 3

// Field # 351
const studentUpdatedFamilySizeCHVFlagsStartIndex2627 int = 3286
const studentUpdatedFamilySizeCHVFlagsLength2627 int = 3

// Field # 352
const studentNumberInCollegeCorrectionCHVFlagsStartIndex2627 int = 3289
const studentNumberInCollegeCorrectionCHVFlagsLength2627 int = 3

// Field # 353
const studentCitizenshipStatusCorrectionCHVFlagsStartIndex2627 int = 3292
const studentCitizenshipStatusCorrectionCHVFlagsLength2627 int = 3

// Field # 354
const studentANumberCHVFlagsStartIndex2627 int = 3295
const studentANumberCHVFlagsLength2627 int = 3

// Field # 355
const studentStateOfLegalResidenceCHVFlagsStartIndex2627 int = 3298
const studentStateOfLegalResidenceCHVFlagsLength2627 int = 3

// Field # 356
const studentLegalResidenceDateCHVFlagsStartIndex2627 int = 3301
const studentLegalResidenceDateCHVFlagsLength2627 int = 3

// Field # 357
const studentEitherParentAttendCollegeCHVFlagsStartIndex2627 int = 3304
const studentEitherParentAttendCollegeCHVFlagsLength2627 int = 3

// Field # 358
const studentParentKilledInTheLineOfDutyCHVFlagsStartIndex2627 int = 3307
const studentParentKilledInTheLineOfDutyCHVFlagsLength2627 int = 3

// Field # 359
const studentHighSchoolCompletionStatusCHVFlagsStartIndex2627 int = 3310
const studentHighSchoolCompletionStatusCHVFlagsLength2627 int = 3

// Field # 360
const studentHighSchoolNameCHVFlagsStartIndex2627 int = 3313
const studentHighSchoolNameCHVFlagsLength2627 int = 3

// Field # 361
const studentHighSchoolCityCHVFlagsStartIndex2627 int = 3316
const studentHighSchoolCityCHVFlagsLength2627 int = 3

// Field # 362
const studentHighSchoolStateCHVFlagsStartIndex2627 int = 3319
const studentHighSchoolStateCHVFlagsLength2627 int = 3

// Field # 363
const studentHighSchoolEquivalentDiplomaNameCHVFlagsStartIndex2627 int = 3322
const studentHighSchoolEquivalentDiplomaNameCHVFlagsLength2627 int = 3

// Field # 364
const studentHighSchoolEquivalentDiplomaStateCHVFlagsStartIndex2627 int = 3325
const studentHighSchoolEquivalentDiplomaStateCHVFlagsLength2627 int = 3

// Field # 365
const studentReceivedEITCCHVFlagsStartIndex2627 int = 3328
const studentReceivedEITCCHVFlagsLength2627 int = 3

// Field # 366
const studentReceivedFederalHousingAssistanceCHVFlagsStartIndex2627 int = 3331
const studentReceivedFederalHousingAssistanceCHVFlagsLength2627 int = 3

// Field # 367
const studentReceivedFreeReducedPriceLunchCHVFlagsStartIndex2627 int = 3334
const studentReceivedFreeReducedPriceLunchCHVFlagsLength2627 int = 3

// Field # 368
const studentReceivedMedicaidCHVFlagsStartIndex2627 int = 3337
const studentReceivedMedicaidCHVFlagsLength2627 int = 3

// Field # 369
const studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex2627 int = 3340
const studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsLength2627 int = 3

// Field # 370
const studentReceivedSNAPCHVFlagsStartIndex2627 int = 3343
const studentReceivedSNAPCHVFlagsLength2627 int = 3

// Field # 371
const studentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex2627 int = 3346
const studentReceivedSupplementalSecurityIncomeCHVFlagsLength2627 int = 3

// Field # 372
const studentReceivedTANFCHVFlagsStartIndex2627 int = 3349
const studentReceivedTANFCHVFlagsLength2627 int = 3

// Field # 373
const studentReceivedWICCHVFlagsStartIndex2627 int = 3352
const studentReceivedWICCHVFlagsLength2627 int = 3

// Field # 374
const studentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex2627 int = 3355
const studentFederalBenefitsNoneOfTheAboveCHVFlagsLength2627 int = 3

// Field # 375
const studentFiled1040Or1040NRCHVFlagsStartIndex2627 int = 3358
const studentFiled1040Or1040NRCHVFlagsLength2627 int = 3

// Field # 376
const studentFiledNonUSTaxReturnCHVFlagsStartIndex2627 int = 3361
const studentFiledNonUSTaxReturnCHVFlagsLength2627 int = 3

// Field # 377
const studentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex2627 int = 3364
const studentFiledJointReturnWithCurrentSpouseCHVFlagsLength2627 int = 3

// Field # 378
const studentTaxReturnFilingStatusCHVFlagsStartIndex2627 int = 3367
const studentTaxReturnFilingStatusCHVFlagsLength2627 int = 3

// Field # 379
const studentIncomeEarnedFromWorkCorrectionCHVFlagsStartIndex2627 int = 3370
const studentIncomeEarnedFromWorkCorrectionCHVFlagsLength2627 int = 3

// Field # 380
const studentTaxExemptInterestIncomeCHVFlagsStartIndex2627 int = 3373
const studentTaxExemptInterestIncomeCHVFlagsLength2627 int = 3

// Field # 381
const studentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2627 int = 3376
const studentUntaxedPortionsOfIRADistributionsCHVFlagsLength2627 int = 3

// Field # 382
const studentIRARolloverCHVFlagsStartIndex2627 int = 3379
const studentIRARolloverCHVFlagsLength2627 int = 3

// Field # 383
const studentUntaxedPortionsOfPensionsCHVFlagsStartIndex2627 int = 3382
const studentUntaxedPortionsOfPensionsCHVFlagsLength2627 int = 3

// Field # 384
const studentPensionRolloverCHVFlagsStartIndex2627 int = 3385
const studentPensionRolloverCHVFlagsLength2627 int = 3

// Field # 385
const studentAdjustedGrossIncomeCHVFlagsStartIndex2627 int = 3388
const studentAdjustedGrossIncomeCHVFlagsLength2627 int = 3

// Field # 386
const studentIncomeTaxPaidCHVFlagsStartIndex2627 int = 3391
const studentIncomeTaxPaidCHVFlagsLength2627 int = 3

// Field # 387
const studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex2627 int = 3394
const studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsLength2627 int = 3

// Field # 388
const studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2627 int = 3397
const studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2627 int = 3

// Field # 389
const studentEducationCreditsCHVFlagsStartIndex2627 int = 3400
const studentEducationCreditsCHVFlagsLength2627 int = 3

// Field # 390
const studentFiledScheduleABDEFHCHVFlagsStartIndex2627 int = 3403
const studentFiledScheduleABDEFHCHVFlagsLength2627 int = 3

// Field # 391
const studentScheduleCAmountCHVFlagsStartIndex2627 int = 3406
const studentScheduleCAmountCHVFlagsLength2627 int = 3

// Field # 392
const studentCollegeGrantAndScholarshipAidCHVFlagsStartIndex2627 int = 3409
const studentCollegeGrantAndScholarshipAidCHVFlagsLength2627 int = 3

// Field # 393
const studentForeignEarnedIncomeExclusionCHVFlagsStartIndex2627 int = 3412
const studentForeignEarnedIncomeExclusionCHVFlagsLength2627 int = 3

// Field # 394
const studentChildSupportReceivedCHVFlagsStartIndex2627 int = 3415
const studentChildSupportReceivedCHVFlagsLength2627 int = 3

// Field # 395
const studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex2627 int = 3418
const studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsLength2627 int = 3

// Field # 396
const studentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex2627 int = 3421
const studentNetWorthOfCurrentInvestmentsCHVFlagsLength2627 int = 3

// Field # 397
const studentTotalOfCashSavingsAndCheckingCHVFlagsStartIndex2627 int = 3424
const studentTotalOfCashSavingsAndCheckingCHVFlagsLength2627 int = 3

// Field # 398
const studentCollege1CHVFlagsStartIndex2627 int = 3427
const studentCollege1CHVFlagsLength2627 int = 3

// Field # 399
const studentCollege2CHVFlagsStartIndex2627 int = 3430
const studentCollege2CHVFlagsLength2627 int = 3

// Field # 400
const studentCollege3CHVFlagsStartIndex2627 int = 3433
const studentCollege3CHVFlagsLength2627 int = 3

// Field # 401
const studentCollege4CHVFlagsStartIndex2627 int = 3436
const studentCollege4CHVFlagsLength2627 int = 3

// Field # 402
const studentCollege5CHVFlagsStartIndex2627 int = 3439
const studentCollege5CHVFlagsLength2627 int = 3

// Field # 403
const studentCollege6CHVFlagsStartIndex2627 int = 3442
const studentCollege6CHVFlagsLength2627 int = 3

// Field # 404
const studentCollege7CHVFlagsStartIndex2627 int = 3445
const studentCollege7CHVFlagsLength2627 int = 3

// Field # 405
const studentCollege8CHVFlagsStartIndex2627 int = 3448
const studentCollege8CHVFlagsLength2627 int = 3

// Field # 406
const studentCollege9CHVFlagsStartIndex2627 int = 3451
const studentCollege9CHVFlagsLength2627 int = 3

// Field # 407
const studentCollege10CHVFlagsStartIndex2627 int = 3454
const studentCollege10CHVFlagsLength2627 int = 3

// Field # 408
const studentCollege11CHVFlagsStartIndex2627 int = 3457
const studentCollege11CHVFlagsLength2627 int = 3

// Field # 409
const studentCollege12CHVFlagsStartIndex2627 int = 3460
const studentCollege12CHVFlagsLength2627 int = 3

// Field # 410
const studentCollege13CHVFlagsStartIndex2627 int = 3463
const studentCollege13CHVFlagsLength2627 int = 3

// Field # 411
const studentCollege14CHVFlagsStartIndex2627 int = 3466
const studentCollege14CHVFlagsLength2627 int = 3

// Field # 412
const studentCollege15CHVFlagsStartIndex2627 int = 3469
const studentCollege15CHVFlagsLength2627 int = 3

// Field # 413
const studentCollege16CHVFlagsStartIndex2627 int = 3472
const studentCollege16CHVFlagsLength2627 int = 3

// Field # 414
const studentCollege17CHVFlagsStartIndex2627 int = 3475
const studentCollege17CHVFlagsLength2627 int = 3

// Field # 415
const studentCollege18CHVFlagsStartIndex2627 int = 3478
const studentCollege18CHVFlagsLength2627 int = 3

// Field # 416
const studentCollege19CHVFlagsStartIndex2627 int = 3481
const studentCollege19CHVFlagsLength2627 int = 3

// Field # 417
const studentCollege20CHVFlagsStartIndex2627 int = 3484
const studentCollege20CHVFlagsLength2627 int = 3

// Field # 418
const studentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2627 int = 3487
const studentConsentToRetrieveAndDiscloseFTICHVFlagsLength2627 int = 3

// Field # 419
const studentSignatureCHVFlagsStartIndex2627 int = 3490
const studentSignatureCHVFlagsLength2627 int = 3

// Field # 420
const studentSignatureDateCHVFlagsStartIndex2627 int = 3493
const studentSignatureDateCHVFlagsLength2627 int = 3

// Field # 421
const studentSpouseFirstNameCHVFlagsStartIndex2627 int = 3496
const studentSpouseFirstNameCHVFlagsLength2627 int = 3

// Field # 422
const studentSpouseMiddleNameCHVFlagsStartIndex2627 int = 3499
const studentSpouseMiddleNameCHVFlagsLength2627 int = 3

// Field # 423
const studentSpouseLastNameCHVFlagsStartIndex2627 int = 3502
const studentSpouseLastNameCHVFlagsLength2627 int = 3

// Field # 424
const studentSpouseSuffixCHVFlagsStartIndex2627 int = 3505
const studentSpouseSuffixCHVFlagsLength2627 int = 3

// Field # 425
const studentSpouseDateOfBirthCHVFlagsStartIndex2627 int = 3508
const studentSpouseDateOfBirthCHVFlagsLength2627 int = 3

// Field # 426
const studentSpouseSSNCHVFlagsStartIndex2627 int = 3511
const studentSpouseSSNCHVFlagsLength2627 int = 3

// Field # 427
const studentSpouseITINCHVFlagsStartIndex2627 int = 3514
const studentSpouseITINCHVFlagsLength2627 int = 3

// Field # 428
const studentSpousePhoneNumberCHVFlagsStartIndex2627 int = 3517
const studentSpousePhoneNumberCHVFlagsLength2627 int = 3

// Field # 429
const studentSpouseEmailAddressCHVFlagsStartIndex2627 int = 3520
const studentSpouseEmailAddressCHVFlagsLength2627 int = 3

// Field # 430
const studentSpouseStreetAddressCHVFlagsStartIndex2627 int = 3523
const studentSpouseStreetAddressCHVFlagsLength2627 int = 3

// Field # 431
const studentSpouseCityCHVFlagsStartIndex2627 int = 3526
const studentSpouseCityCHVFlagsLength2627 int = 3

// Field # 432
const studentSpouseStateCHVFlagsStartIndex2627 int = 3529
const studentSpouseStateCHVFlagsLength2627 int = 3

// Field # 433
const studentSpouseZipCodeCHVFlagsStartIndex2627 int = 3532
const studentSpouseZipCodeCHVFlagsLength2627 int = 3

// Field # 434
const studentSpouseCountryCHVFlagsStartIndex2627 int = 3535
const studentSpouseCountryCHVFlagsLength2627 int = 3

// Field # 435
const studentSpouseFiled1040Or1040NRCHVFlagsStartIndex2627 int = 3538
const studentSpouseFiled1040Or1040NRCHVFlagsLength2627 int = 3

// Field # 436
const studentSpouseFiledNonUSTaxReturnCHVFlagsStartIndex2627 int = 3541
const studentSpouseFiledNonUSTaxReturnCHVFlagsLength2627 int = 3

// Field # 437
const studentSpouseTaxReturnFilingStatusCHVFlagsStartIndex2627 int = 3544
const studentSpouseTaxReturnFilingStatusCHVFlagsLength2627 int = 3

// Field # 438
const studentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex2627 int = 3547
const studentSpouseIncomeEarnedFromWorkCHVFlagsLength2627 int = 3

// Field # 439
const studentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex2627 int = 3550
const studentSpouseTaxExemptInterestIncomeCHVFlagsLength2627 int = 3

// Field # 440
const studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2627 int = 3553
const studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsLength2627 int = 3

// Field # 441
const studentSpouseIRARolloverCHVFlagsStartIndex2627 int = 3556
const studentSpouseIRARolloverCHVFlagsLength2627 int = 3

// Field # 442
const studentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex2627 int = 3559
const studentSpouseUntaxedPortionsOfPensionsCHVFlagsLength2627 int = 3

// Field # 443
const studentSpousePensionRolloverCHVFlagsStartIndex2627 int = 3562
const studentSpousePensionRolloverCHVFlagsLength2627 int = 3

// Field # 444
const studentSpouseAdjustedGrossIncomeCHVFlagsStartIndex2627 int = 3565
const studentSpouseAdjustedGrossIncomeCHVFlagsLength2627 int = 3

// Field # 445
const studentSpouseIncomeTaxPaidCHVFlagsStartIndex2627 int = 3568
const studentSpouseIncomeTaxPaidCHVFlagsLength2627 int = 3

// Field # 446
const studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2627 int = 3571
const studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2627 int = 3

// Field # 447
const studentSpouseEducationCreditsCHVFlagsStartIndex2627 int = 3574
const studentSpouseEducationCreditsCHVFlagsLength2627 int = 3

// Field # 448
const studentSpouseFiledScheduleABDEFHCHVFlagsStartIndex2627 int = 3577
const studentSpouseFiledScheduleABDEFHCHVFlagsLength2627 int = 3

// Field # 449
const studentSpouseScheduleCAmountCHVFlagsStartIndex2627 int = 3580
const studentSpouseScheduleCAmountCHVFlagsLength2627 int = 3

// Field # 450
const studentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex2627 int = 3583
const studentSpouseForeignEarnedIncomeExclusionCHVFlagsLength2627 int = 3

// Field # 451
const studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2627 int = 3586
const studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsLength2627 int = 3

// Field # 452
const studentSpouseSignatureCHVFlagsStartIndex2627 int = 3589
const studentSpouseSignatureCHVFlagsLength2627 int = 3

// Field # 453
const studentSpouseSignatureDateCHVFlagsStartIndex2627 int = 3592
const studentSpouseSignatureDateCHVFlagsLength2627 int = 3

// Field # 454
const parentFirstNameCHVFlagsStartIndex2627 int = 3595
const parentFirstNameCHVFlagsLength2627 int = 3

// Field # 455
const parentMiddleNameCHVFlagsStartIndex2627 int = 3598
const parentMiddleNameCHVFlagsLength2627 int = 3

// Field # 456
const parentLastNameCHVFlagsStartIndex2627 int = 3601
const parentLastNameCHVFlagsLength2627 int = 3

// Field # 457
const parentSuffixCHVFlagsStartIndex2627 int = 3604
const parentSuffixCHVFlagsLength2627 int = 3

// Field # 458
const parentDateOfBirthCHVFlagsStartIndex2627 int = 3607
const parentDateOfBirthCHVFlagsLength2627 int = 3

// Field # 459
const parentSSNCHVFlagsStartIndex2627 int = 3610
const parentSSNCHVFlagsLength2627 int = 3

// Field # 460
const parentITINCHVFlagsStartIndex2627 int = 3613
const parentITINCHVFlagsLength2627 int = 3

// Field # 461
const parentPhoneNumberCHVFlagsStartIndex2627 int = 3616
const parentPhoneNumberCHVFlagsLength2627 int = 3

// Field # 462
const parentEmailAddressCHVFlagsStartIndex2627 int = 3619
const parentEmailAddressCHVFlagsLength2627 int = 3

// Field # 463
const parentStreetAddressCHVFlagsStartIndex2627 int = 3622
const parentStreetAddressCHVFlagsLength2627 int = 3

// Field # 464
const parentCityCHVFlagsStartIndex2627 int = 3625
const parentCityCHVFlagsLength2627 int = 3

// Field # 465
const parentStateCHVFlagsStartIndex2627 int = 3628
const parentStateCHVFlagsLength2627 int = 3

// Field # 466
const parentZipCodeCHVFlagsStartIndex2627 int = 3631
const parentZipCodeCHVFlagsLength2627 int = 3

// Field # 467
const parentCountryCHVFlagsStartIndex2627 int = 3634
const parentCountryCHVFlagsLength2627 int = 3

// Field # 468
const parentMaritalStatusCHVFlagsStartIndex2627 int = 3637
const parentMaritalStatusCHVFlagsLength2627 int = 3

// Field # 469
const parentStateOfLegalResidenceCHVFlagsStartIndex2627 int = 3640
const parentStateOfLegalResidenceCHVFlagsLength2627 int = 3

// Field # 470
const parentLegalResidenceDateCHVFlagsStartIndex2627 int = 3643
const parentLegalResidenceDateCHVFlagsLength2627 int = 3

// Field # 471
const parentUpdatedFamilySizeCHVFlagsStartIndex2627 int = 3646
const parentUpdatedFamilySizeCHVFlagsLength2627 int = 3

// Field # 472
const parentNumberInCollegeCHVFlagsStartIndex2627 int = 3649
const parentNumberInCollegeCHVFlagsLength2627 int = 3

// Field # 473
const parentReceivedEITCCHVFlagsStartIndex2627 int = 3652
const parentReceivedEITCCHVFlagsLength2627 int = 3

// Field # 474
const parentReceivedFederalHousingAssistanceCHVFlagsStartIndex2627 int = 3655
const parentReceivedFederalHousingAssistanceCHVFlagsLength2627 int = 3

// Field # 475
const parentReceivedFreeReducedPriceLunchCHVFlagsStartIndex2627 int = 3658
const parentReceivedFreeReducedPriceLunchCHVFlagsLength2627 int = 3

// Field # 476
const parentReceivedMedicaidCHVFlagsStartIndex2627 int = 3661
const parentReceivedMedicaidCHVFlagsLength2627 int = 3

// Field # 477
const parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex2627 int = 3664
const parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsLength2627 int = 3

// Field # 478
const parentReceivedSNAPCHVFlagsStartIndex2627 int = 3667
const parentReceivedSNAPCHVFlagsLength2627 int = 3

// Field # 479
const parentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex2627 int = 3670
const parentReceivedSupplementalSecurityIncomeCHVFlagsLength2627 int = 3

// Field # 480
const parentReceivedTANFCHVFlagsStartIndex2627 int = 3673
const parentReceivedTANFCHVFlagsLength2627 int = 3

// Field # 481
const parentReceivedWICCHVFlagsStartIndex2627 int = 3676
const parentReceivedWICCHVFlagsLength2627 int = 3

// Field # 482
const parentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex2627 int = 3679
const parentFederalBenefitsNoneOfTheAboveCHVFlagsLength2627 int = 3

// Field # 483
const parentFiled1040Or1040NRCHVFlagsStartIndex2627 int = 3682
const parentFiled1040Or1040NRCHVFlagsLength2627 int = 3

// Field # 484
const parentFileNonUSTaxReturnCHVFlagsStartIndex2627 int = 3685
const parentFileNonUSTaxReturnCHVFlagsLength2627 int = 3

// Field # 485
const parentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex2627 int = 3688
const parentFiledJointReturnWithCurrentSpouseCHVFlagsLength2627 int = 3

// Field # 486
const parentTaxReturnFilingStatusCHVFlagsStartIndex2627 int = 3691
const parentTaxReturnFilingStatusCHVFlagsLength2627 int = 3

// Field # 487
const parentIncomeEarnedFromWorkCHVFlagsStartIndex2627 int = 3694
const parentIncomeEarnedFromWorkCHVFlagsLength2627 int = 3

// Field # 488
const parentTaxExemptInterestIncomeCHVFlagsStartIndex2627 int = 3697
const parentTaxExemptInterestIncomeCHVFlagsLength2627 int = 3

// Field # 489
const parentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2627 int = 3700
const parentUntaxedPortionsOfIRADistributionsCHVFlagsLength2627 int = 3

// Field # 490
const parentIRARolloverCHVFlagsStartIndex2627 int = 3703
const parentIRARolloverCHVFlagsLength2627 int = 3

// Field # 491
const parentUntaxedPortionsOfPensionsCHVFlagsStartIndex2627 int = 3706
const parentUntaxedPortionsOfPensionsCHVFlagsLength2627 int = 3

// Field # 492
const parentPensionRolloverCHVFlagsStartIndex2627 int = 3709
const parentPensionRolloverCHVFlagsLength2627 int = 3

// Field # 493
const parentAdjustedGrossIncomeCHVFlagsStartIndex2627 int = 3712
const parentAdjustedGrossIncomeCHVFlagsLength2627 int = 3

// Field # 494
const parentIncomeTaxPaidCHVFlagsStartIndex2627 int = 3715
const parentIncomeTaxPaidCHVFlagsLength2627 int = 3

// Field # 495
const parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex2627 int = 3718
const parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsLength2627 int = 3

// Field # 496
const parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2627 int = 3721
const parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2627 int = 3

// Field # 497
const parentEducationCreditsCHVFlagsStartIndex2627 int = 3724
const parentEducationCreditsCHVFlagsLength2627 int = 3

// Field # 498
const parentFiledScheduleABDEFHCHVFlagsStartIndex2627 int = 3727
const parentFiledScheduleABDEFHCHVFlagsLength2627 int = 3

// Field # 499
const parentScheduleCAmountCHVFlagsStartIndex2627 int = 3730
const parentScheduleCAmountCHVFlagsLength2627 int = 3

// Field # 500
const parentCollegeGrantAndScholarshipAidCHVFlagsStartIndex2627 int = 3733
const parentCollegeGrantAndScholarshipAidCHVFlagsLength2627 int = 3

// Field # 501
const parentForeignEarnedIncomeExclusionCHVFlagsStartIndex2627 int = 3736
const parentForeignEarnedIncomeExclusionCHVFlagsLength2627 int = 3

// Field # 502
const parentChildSupportReceivedCHVFlagsStartIndex2627 int = 3739
const parentChildSupportReceivedCHVFlagsLength2627 int = 3

// Field # 503
const parentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex2627 int = 3742
const parentNetWorthOfCurrentInvestmentsCHVFlagsLength2627 int = 3

// Field # 504
const parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsStartIndex2627 int = 3745
const parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsLength2627 int = 3

// Field # 505
const parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex2627 int = 3748
const parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsLength2627 int = 3

// Field # 506
const parentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2627 int = 3751
const parentConsentToRetrieveAndDiscloseFTICHVFlagsLength2627 int = 3

// Field # 507
const parentSignatureCHVFlagsStartIndex2627 int = 3754
const parentSignatureCHVFlagsLength2627 int = 3

// Field # 508
const parentSignatureDateCHVFlagsStartIndex2627 int = 3757
const parentSignatureDateCHVFlagsLength2627 int = 3

// Field # 509
const parentSpouseFirstNameCHVFlagsStartIndex2627 int = 3760
const parentSpouseFirstNameCHVFlagsLength2627 int = 3

// Field # 510
const parentSpouseMiddleNameCHVFlagsStartIndex2627 int = 3763
const parentSpouseMiddleNameCHVFlagsLength2627 int = 3

// Field # 511
const parentSpouseLastNameCHVFlagsStartIndex2627 int = 3766
const parentSpouseLastNameCHVFlagsLength2627 int = 3

// Field # 512
const parentSpouseSuffixCHVFlagsStartIndex2627 int = 3769
const parentSpouseSuffixCHVFlagsLength2627 int = 3

// Field # 513
const parentSpouseDateOfBirthCHVFlagsStartIndex2627 int = 3772
const parentSpouseDateOfBirthCHVFlagsLength2627 int = 3

// Field # 514
const parentSpouseSSNCHVFlagsStartIndex2627 int = 3775
const parentSpouseSSNCHVFlagsLength2627 int = 3

// Field # 515
const parentSpouseITINCHVFlagsStartIndex2627 int = 3778
const parentSpouseITINCHVFlagsLength2627 int = 3

// Field # 516
const parentSpousePhoneNumberCHVFlagsStartIndex2627 int = 3781
const parentSpousePhoneNumberCHVFlagsLength2627 int = 3

// Field # 517
const parentSpouseEmailAddressCHVFlagsStartIndex2627 int = 3784
const parentSpouseEmailAddressCHVFlagsLength2627 int = 3

// Field # 518
const parentSpouseStreetAddressCHVFlagsStartIndex2627 int = 3787
const parentSpouseStreetAddressCHVFlagsLength2627 int = 3

// Field # 519
const parentSpouseCityCHVFlagsStartIndex2627 int = 3790
const parentSpouseCityCHVFlagsLength2627 int = 3

// Field # 520
const parentSpouseStateCHVFlagsStartIndex2627 int = 3793
const parentSpouseStateCHVFlagsLength2627 int = 3

// Field # 521
const parentSpouseZipCodeCHVFlagsStartIndex2627 int = 3796
const parentSpouseZipCodeCHVFlagsLength2627 int = 3

// Field # 522
const parentSpouseCountryCHVFlagsStartIndex2627 int = 3799
const parentSpouseCountryCHVFlagsLength2627 int = 3

// Field # 523
const parentSpouseFiled1040Or1040NRCHVFlagsStartIndex2627 int = 3802
const parentSpouseFiled1040Or1040NRCHVFlagsLength2627 int = 3

// Field # 524
const parentSpouseFileNonUSTaxReturnCHVFlagsStartIndex2627 int = 3805
const parentSpouseFileNonUSTaxReturnCHVFlagsLength2627 int = 3

// Field # 525
const parentSpouseTaxReturnFilingStatusCHVFlagsStartIndex2627 int = 3808
const parentSpouseTaxReturnFilingStatusCHVFlagsLength2627 int = 3

// Field # 526
const parentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex2627 int = 3811
const parentSpouseIncomeEarnedFromWorkCHVFlagsLength2627 int = 3

// Field # 527
const parentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex2627 int = 3814
const parentSpouseTaxExemptInterestIncomeCHVFlagsLength2627 int = 3

// Field # 528
const parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2627 int = 3817
const parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsLength2627 int = 3

// Field # 529
const parentSpouseIRARolloverCHVFlagsStartIndex2627 int = 3820
const parentSpouseIRARolloverCHVFlagsLength2627 int = 3

// Field # 530
const parentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex2627 int = 3823
const parentSpouseUntaxedPortionsOfPensionsCHVFlagsLength2627 int = 3

// Field # 531
const parentSpousePensionRolloverCHVFlagsStartIndex2627 int = 3826
const parentSpousePensionRolloverCHVFlagsLength2627 int = 3

// Field # 532
const parentSpouseAdjustedGrossIncomeCHVFlagsStartIndex2627 int = 3829
const parentSpouseAdjustedGrossIncomeCHVFlagsLength2627 int = 3

// Field # 533
const parentSpouseIncomeTaxPaidCHVFlagsStartIndex2627 int = 3832
const parentSpouseIncomeTaxPaidCHVFlagsLength2627 int = 3

// Field # 534
const parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2627 int = 3835
const parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2627 int = 3

// Field # 535
const parentSpouseEducationCreditsCHVFlagsStartIndex2627 int = 3838
const parentSpouseEducationCreditsCHVFlagsLength2627 int = 3

// Field # 536
const parentSpouseFiledScheduleABDEFHCHVFlagsStartIndex2627 int = 3841
const parentSpouseFiledScheduleABDEFHCHVFlagsLength2627 int = 3

// Field # 537
const parentSpouseScheduleCAmountCHVFlagsStartIndex2627 int = 3844
const parentSpouseScheduleCAmountCHVFlagsLength2627 int = 3

// Field # 538
const parentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex2627 int = 3847
const parentSpouseForeignEarnedIncomeExclusionCHVFlagsLength2627 int = 3

// Field # 539
const parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2627 int = 3850
const parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsLength2627 int = 3

// Field # 540
const parentSpouseSignatureCHVFlagsStartIndex2627 int = 3853
const parentSpouseSignatureCHVFlagsLength2627 int = 3

// Field # 541
const parentSpouseSignatureDateCHVFlagsStartIndex2627 int = 3856
const parentSpouseSignatureDateCHVFlagsLength2627 int = 3

// Field # 542
const dHSPrimaryMatchStatusStartIndex2627 int = 3859
const dHSPrimaryMatchStatusLength2627 int = 1

// Field # 544
const dHSCaseNumberStartIndex2627 int = 3861
const dHSCaseNumberLength2627 int = 15

// Field # 545
const nsldsMatchStatusStartIndex2627 int = 3876
const nsldsMatchStatusLength2627 int = 1

// Field # 546
const nsldsPostscreeningReasonCodeStartIndex2627 int = 3877
const nsldsPostscreeningReasonCodeLength2627 int = 6

// Field # 547
const studentSSACitizenshipFlagResultsStartIndex2627 int = 3883
const studentSSACitizenshipFlagResultsLength2627 int = 1

// Field # 548
const studentSSAMatchStatusStartIndex2627 int = 3884
const studentSSAMatchStatusLength2627 int = 1

// Field # 549
const studentSpouseSSAMatchStatusStartIndex2627 int = 3885
const studentSpouseSSAMatchStatusLength2627 int = 1

// Field # 550
const parentSSAMatchStatusStartIndex2627 int = 3886
const parentSSAMatchStatusLength2627 int = 1

// Field # 551
const parentSpouseOrPartnerSSAMatchStatusStartIndex2627 int = 3887
const parentSpouseOrPartnerSSAMatchStatusLength2627 int = 1

// Field # 552
const vAMatchFlagStartIndex2627 int = 3888
const vAMatchFlagLength2627 int = 1

// Field # 553
const commentCodesStartIndex2627 int = 3889
const commentCodesLength2627 int = 60

// Field # 554
const drugAbuseHoldIndicatorStartIndex2627 int = 3949
const drugAbuseHoldIndicatorLength2627 int = 1

// Field # 555
const graduateFlagStartIndex2627 int = 3950
const graduateFlagLength2627 int = 1

// Field # 556
const pellGrantEligibilityFlagStartIndex2627 int = 3951
const pellGrantEligibilityFlagLength2627 int = 1

// Field # 557
const reprocessedReasonCodeStartIndex2627 int = 3952
const reprocessedReasonCodeLength2627 int = 2

// Field # 558
const fpsCFlagStartIndex2627 int = 3954
const fpsCFlagLength2627 int = 1

// Field # 559
const fpsCChangeFlagStartIndex2627 int = 3955
const fpsCChangeFlagLength2627 int = 1

// Field # 560
const electronicFederalSchoolCodeIndicatorStartIndex2627 int = 3956
const electronicFederalSchoolCodeIndicatorLength2627 int = 2

// Field # 561
const rejectReasonCodesStartIndex2627 int = 3958
const rejectReasonCodesLength2627 int = 110

// Field # 562
const electronicTransactionIndicatorFlagStartIndex2627 int = 4068
const electronicTransactionIndicatorFlagLength2627 int = 1

// Field # 563
const studentLastNameSSNChangeFlagStartIndex2627 int = 4069
const studentLastNameSSNChangeFlagLength2627 int = 1

// Field # 564
const highSchoolCodeStartIndex2627 int = 4070
const highSchoolCodeLength2627 int = 12

// Field # 565
const verificationSelectionChangeFlagStartIndex2627 int = 4082
const verificationSelectionChangeFlagLength2627 int = 1

// Field # 566 removed in 26-27
//const useUserProvidedDataOnlyStartIndex2627 int = 4083
//const useUserProvidedDataOnlyLength2627 int = 5

// Field # 568
const nsldsPellOverpaymentFlagStartIndex2627 int = 4449
const nsldsPellOverpaymentFlagLength2627 int = 1

// Field # 569
const nsldsPellOverpaymentContactStartIndex2627 int = 4450
const nsldsPellOverpaymentContactLength2627 int = 8

// Field # 570
const nsldsFSEOGOverpaymentFlagStartIndex2627 int = 4458
const nsldsFSEOGOverpaymentFlagLength2627 int = 1

// Field # 571
const nsldsFSEOGOverpaymentContactStartIndex2627 int = 4459
const nsldsFSEOGOverpaymentContactLength2627 int = 8

// Field # 572
const nsldsPerkinsOverpaymentFlagStartIndex2627 int = 4467
const nsldsPerkinsOverpaymentFlagLength2627 int = 1

// Field # 573
const nsldsPerkinsOverpaymentContactStartIndex2627 int = 4468
const nsldsPerkinsOverpaymentContactLength2627 int = 8

// Field # 574
const nsldsTEACHGrantOverpaymentFlagStartIndex2627 int = 4476
const nsldsTEACHGrantOverpaymentFlagLength2627 int = 1

// Field # 575
const nsldsTEACHGrantOverpaymentContactStartIndex2627 int = 4477
const nsldsTEACHGrantOverpaymentContactLength2627 int = 8

// Field # 576
const nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagStartIndex2627 int = 4485
const nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagLength2627 int = 1

// Field # 577
const nsldsIraqAndAfghanistanServiceGrantOverpaymentContactStartIndex2627 int = 4486
const nsldsIraqAndAfghanistanServiceGrantOverpaymentContactLength2627 int = 8

// Field # 578
const nsldsDefaultedLoanFlagStartIndex2627 int = 4494
const nsldsDefaultedLoanFlagLength2627 int = 1

// Field # 579
const nsldsDischargedLoanFlagStartIndex2627 int = 4495
const nsldsDischargedLoanFlagLength2627 int = 1

// Field # 580
const nsldsFraudLoanFlagStartIndex2627 int = 4496
const nsldsFraudLoanFlagLength2627 int = 1

// Field # 581
const nsldsSatisfactoryArrangementsFlagStartIndex2627 int = 4497
const nsldsSatisfactoryArrangementsFlagLength2627 int = 1

// Field # 582
const nsldsActiveBankruptcyFlagStartIndex2627 int = 4498
const nsldsActiveBankruptcyFlagLength2627 int = 1

// Field # 583
const nsldsTEACHGrantConvertedToLoanFlagStartIndex2627 int = 4499
const nsldsTEACHGrantConvertedToLoanFlagLength2627 int = 1

// Field # 584
const nsldsAggregateSubsidizedOutstandingPrincipalBalanceStartIndex2627 int = 4500
const nsldsAggregateSubsidizedOutstandingPrincipalBalanceLength2627 int = 6

// Field # 585
const nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceStartIndex2627 int = 4506
const nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceLength2627 int = 6

// Field # 586
const nsldsAggregateCombinedOutstandingPrincipalBalanceStartIndex2627 int = 4512
const nsldsAggregateCombinedOutstandingPrincipalBalanceLength2627 int = 6

// Field # 587
const nsldsAggregateUnallocConsolidatedOutstandingPrincipalBalanceStartIndex2627 int = 4518
const nsldsAggregateUnallocConsolidatedOutstandingPrincipalBalanceLength2627 int = 6

// Field # 588
const nsldsAggregateTEACHLoanPrincipalBalanceStartIndex2627 int = 4524
const nsldsAggregateTEACHLoanPrincipalBalanceLength2627 int = 6

// Field # 589
const nsldsAggregateSubsidizedPendingDisbursementStartIndex2627 int = 4530
const nsldsAggregateSubsidizedPendingDisbursementLength2627 int = 6

// Field # 590
const nsldsAggregateUnsubsidizedPendingDisbursementStartIndex2627 int = 4536
const nsldsAggregateUnsubsidizedPendingDisbursementLength2627 int = 6

// Field # 591
const nsldsAggregateCombinedPendingDisbursementStartIndex2627 int = 4542
const nsldsAggregateCombinedPendingDisbursementLength2627 int = 6

// Field # 592
const nsldsAggregateSubsidizedTotalStartIndex2627 int = 4548
const nsldsAggregateSubsidizedTotalLength2627 int = 6

// Field # 593
const nsldsAggregateUnsubsidizedTotalStartIndex2627 int = 4554
const nsldsAggregateUnsubsidizedTotalLength2627 int = 6

// Field # 594
const nsldsAggregateCombinedTotalStartIndex2627 int = 4560
const nsldsAggregateCombinedTotalLength2627 int = 6

// Field # 595
const nsldsUnallocatedConsolidatedTotalStartIndex2627 int = 4566
const nsldsUnallocatedConsolidatedTotalLength2627 int = 6

// Field # 596
const nsldsTEACHLoanTotalStartIndex2627 int = 4572
const nsldsTEACHLoanTotalLength2627 int = 6

// Field # 597
const nsldsPerkinsTotalDisbursementsStartIndex2627 int = 4578
const nsldsPerkinsTotalDisbursementsLength2627 int = 6

// Field # 598
const nsldsPerkinsCurrentYearDisbursementAmountStartIndex2627 int = 4584
const nsldsPerkinsCurrentYearDisbursementAmountLength2627 int = 6

// Field # 599
const nsldsAggregateTEACHGrantUndergraduateDisbursedTotalStartIndex2627 int = 4590
const nsldsAggregateTEACHGrantUndergraduateDisbursedTotalLength2627 int = 6

// Field # 600
const nsldsAggregateTEACHGraduateDisbursementAmountStartIndex2627 int = 4596
const nsldsAggregateTEACHGraduateDisbursementAmountLength2627 int = 6

// Field # 601
const nsldsDefaultedLoanChangeFlagStartIndex2627 int = 4602
const nsldsDefaultedLoanChangeFlagLength2627 int = 1

// Field # 602
const nsldsFraudLoanChangeFlagStartIndex2627 int = 4603
const nsldsFraudLoanChangeFlagLength2627 int = 1

// Field # 603
const nsldsDischargedLoanChangeFlagStartIndex2627 int = 4604
const nsldsDischargedLoanChangeFlagLength2627 int = 1

// Field # 604
const nsldsLoanSatisfactoryRepaymentChangeFlagStartIndex2627 int = 4605
const nsldsLoanSatisfactoryRepaymentChangeFlagLength2627 int = 1

// Field # 605
const nsldsActiveBankruptcyChangeFlagStartIndex2627 int = 4606
const nsldsActiveBankruptcyChangeFlagLength2627 int = 1

// Field # 606
const nsldsTEACHGrantToLoanConversionChangeFlagStartIndex2627 int = 4607
const nsldsTEACHGrantToLoanConversionChangeFlagLength2627 int = 1

// Field # 607
const nsldsOverpaymentsChangeFlagStartIndex2627 int = 4608
const nsldsOverpaymentsChangeFlagLength2627 int = 1

// Field # 608
const nsldsAggregateLoanChangeFlagStartIndex2627 int = 4609
const nsldsAggregateLoanChangeFlagLength2627 int = 1

// Field # 609
const nsldsPerkinsLoanChangeFlagStartIndex2627 int = 4610
const nsldsPerkinsLoanChangeFlagLength2627 int = 1

// Field # 610
const nsldsPellPaymentChangeFlagStartIndex2627 int = 4611
const nsldsPellPaymentChangeFlagLength2627 int = 1

// Field # 611
const nsldsTEACHGrantChangeFlagStartIndex2627 int = 4612
const nsldsTEACHGrantChangeFlagLength2627 int = 1

// Field # 612
const nsldsAdditionalPellFlagStartIndex2627 int = 4613
const nsldsAdditionalPellFlagLength2627 int = 1

// Field # 613
const nsldsAdditionalLoansFlagStartIndex2627 int = 4614
const nsldsAdditionalLoansFlagLength2627 int = 1

// Field # 614
const nsldsAdditionalTEACHGrantFlagStartIndex2627 int = 4615
const nsldsAdditionalTEACHGrantFlagLength2627 int = 1

// Field # 615
const nsldsDirectLoanMasterPromNoteFlagStartIndex2627 int = 4616
const nsldsDirectLoanMasterPromNoteFlagLength2627 int = 1

// Field # 616
const nsldsDirectLoanPLUSMasterPromNoteFlagStartIndex2627 int = 4617
const nsldsDirectLoanPLUSMasterPromNoteFlagLength2627 int = 1

// Field # 617
const nsldsDirectLoanGraduatePLUSMasterPromNoteFlagStartIndex2627 int = 4618
const nsldsDirectLoanGraduatePLUSMasterPromNoteFlagLength2627 int = 1

// Field # 618
const nsldsUndergraduateSubsidizedLoanLimitFlagStartIndex2627 int = 4619
const nsldsUndergraduateSubsidizedLoanLimitFlagLength2627 int = 1

// Field # 619
const nsldsUndergraduateCombinedLoanLimitFlagStartIndex2627 int = 4620
const nsldsUndergraduateCombinedLoanLimitFlagLength2627 int = 1

// Field # 620
const nsldsGraduateSubsidizedLoanLimitFlagStartIndex2627 int = 4621
const nsldsGraduateSubsidizedLoanLimitFlagLength2627 int = 1

// Field # 621
const nsldsGraduateCombinedLoanLimitFlagStartIndex2627 int = 4622
const nsldsGraduateCombinedLoanLimitFlagLength2627 int = 1

// Field # 622
const nsldsPellLifetimeLimitFlagStartIndex2627 int = 4623
const nsldsPellLifetimeLimitFlagLength2627 int = 1

// Renamed: unsure why these were mapped over as LEU and no "Flag" but it was a mistake
//const nsldsLEULimitIndicatorStartIndex2627 int = 4623
//const nsldsLEULimitIndicatorLength2627 int = 1

// Field # 623
const nsldsPellLifetimeEligibilityUsedStartIndex2627 int = 4624
const nsldsPellLifetimeEligibilityUsedLength2627 int = 7

// Field # 624
const nsldsSULAFlagStartIndex2627 int = 4631
const nsldsSULAFlagLength2627 int = 1

// Field # 625
const nsldsSubsidizedLimitEligibilityUsedStartIndex2627 int = 4632
const nsldsSubsidizedLimitEligibilityUsedLength2627 int = 6

// Field # 626
const nsldsUnusualEnrollmentHistoryFlagStartIndex2627 int = 4638
const nsldsUnusualEnrollmentHistoryFlagLength2627 int = 1

// Field # 628
const nsldsPellSequenceNumber1StartIndex2627 int = 4659
const nsldsPellSequenceNumber1Length2627 int = 2

// Field # 629
const nsldsPellVerificationFlag1StartIndex2627 int = 4661
const nsldsPellVerificationFlag1Length2627 int = 3

// Field # 630
const nsldsSAI1StartIndex2627 int = 4664
const nsldsSAI1Length2627 int = 6

// Field # 631
const nsldsPellSchoolCode1StartIndex2627 int = 4670
const nsldsPellSchoolCode1Length2627 int = 8

// Field # 632
const nsldsPellTransactionNumber1StartIndex2627 int = 4678
const nsldsPellTransactionNumber1Length2627 int = 2

// Field # 633
const nsldsPellDisbursementDate1StartIndex2627 int = 4680
const nsldsPellDisbursementDate1Length2627 int = 8

// Field # 634
const nsldsPellScheduledAmount1StartIndex2627 int = 4688
const nsldsPellScheduledAmount1Length2627 int = 6

// Field # 635
const nsldsPellAmountPaidToDate1StartIndex2627 int = 4694
const nsldsPellAmountPaidToDate1Length2627 int = 6

// Field # 636
const nsldsPellPercentEligibilityUsedDecimal1StartIndex2627 int = 4700
const nsldsPellPercentEligibilityUsedDecimal1Length2627 int = 7

// Field # 637
const nsldsPellAwardAmount1StartIndex2627 int = 4707
const nsldsPellAwardAmount1Length2627 int = 6

// Field # 638
const nsldsAdditionalEligibilityIndicator1StartIndex2627 int = 4713
const nsldsAdditionalEligibilityIndicator1Length2627 int = 1

// Field # 640
const nsldsPellSequenceNumber2StartIndex2627 int = 4734
const nsldsPellSequenceNumber2Length2627 int = 2

// Field # 641
const nsldsPellVerificationFlag2StartIndex2627 int = 4736
const nsldsPellVerificationFlag2Length2627 int = 3

// Field # 642
const nsldsSAI2StartIndex2627 int = 4739
const nsldsSAI2Length2627 int = 6

// Field # 643
const nsldsPellSchoolCode2StartIndex2627 int = 4745
const nsldsPellSchoolCode2Length2627 int = 8

// Field # 644
const nsldsPellTransactionNumber2StartIndex2627 int = 4753
const nsldsPellTransactionNumber2Length2627 int = 2

// Field # 645
const nsldsPellLastDisbursementDate2StartIndex2627 int = 4755
const nsldsPellLastDisbursementDate2Length2627 int = 8

// Field # 646
const nsldsPellScheduledAmount2StartIndex2627 int = 4763
const nsldsPellScheduledAmount2Length2627 int = 6

// Field # 647
const nsldsPellAmountPaidToDate2StartIndex2627 int = 4769
const nsldsPellAmountPaidToDate2Length2627 int = 6

// Field # 648
const nsldsPellPercentEligibilityUsedDecimal2StartIndex2627 int = 4775
const nsldsPellPercentEligibilityUsedDecimal2Length2627 int = 7

// Field # 649
const nsldsPellAwardAmount2StartIndex2627 int = 4782
const nsldsPellAwardAmount2Length2627 int = 6

// Field # 650
const nsldsAdditionalEligibilityIndicator2StartIndex2627 int = 4788
const nsldsAdditionalEligibilityIndicator2Length2627 int = 1

// Field # 652
const nsldsPellSequenceNumber3StartIndex2627 int = 4809
const nsldsPellSequenceNumber3Length2627 int = 2

// Field # 653
const nsldsPellVerificationFlag3StartIndex2627 int = 4811
const nsldsPellVerificationFlag3Length2627 int = 3

// Field # 654
const nsldsSAI3StartIndex2627 int = 4814
const nsldsSAI3Length2627 int = 6

// Field # 655
const nsldsPellSchoolCode3StartIndex2627 int = 4820
const nsldsPellSchoolCode3Length2627 int = 8

// Field # 656
const nsldsPellTransactionNumber3StartIndex2627 int = 4828
const nsldsPellTransactionNumber3Length2627 int = 2

// Field # 657
const nsldsPellLastDisbursementDate3StartIndex2627 int = 4830
const nsldsPellLastDisbursementDate3Length2627 int = 8

// Field # 658
const nsldsPellScheduledAmount3StartIndex2627 int = 4838
const nsldsPellScheduledAmount3Length2627 int = 6

// Field # 659
const nsldsPellAmountPaidToDate3StartIndex2627 int = 4844
const nsldsPellAmountPaidToDate3Length2627 int = 6

// Field # 660
const nsldsPellPercentEligibilityUsedDecimal3StartIndex2627 int = 4850
const nsldsPellPercentEligibilityUsedDecimal3Length2627 int = 7

// Field # 661
const nsldsPellAwardAmount3StartIndex2627 int = 4857
const nsldsPellAwardAmount3Length2627 int = 6

// Field # 662
const nsldsAdditionalEligibilityIndicator3StartIndex2627 int = 4863
const nsldsAdditionalEligibilityIndicator3Length2627 int = 1

// Field # 664
const nsldsTEACHGrantSequence1StartIndex2627 int = 4884
const nsldsTEACHGrantSequence1Length2627 int = 2

// Field # 665
const nsldsTEACHGrantSchoolCode1StartIndex2627 int = 4886
const nsldsTEACHGrantSchoolCode1Length2627 int = 8

// Field # 666
const nsldsTEACHGrantTransactionNumber1StartIndex2627 int = 4894
const nsldsTEACHGrantTransactionNumber1Length2627 int = 2

// Field # 667
const nsldsTEACHGrantLastDisbursementDate1StartIndex2627 int = 4896
const nsldsTEACHGrantLastDisbursementDate1Length2627 int = 8

// Field # 668
const nsldsTEACHGrantScheduledAmount1StartIndex2627 int = 4904
const nsldsTEACHGrantScheduledAmount1Length2627 int = 6

// Field # 669
const nsldsTEACHGrantAmountPaidToDate1StartIndex2627 int = 4910
const nsldsTEACHGrantAmountPaidToDate1Length2627 int = 6

// Field # 670
const nsldsTEACHGrantAwardAmount1StartIndex2627 int = 4916
const nsldsTEACHGrantAwardAmount1Length2627 int = 6

// Field # 671
const nsldsTEACHGrantAcademicYearLevel1StartIndex2627 int = 4922
const nsldsTEACHGrantAcademicYearLevel1Length2627 int = 1

// Field # 672
const nsldsTEACHGrantAwardYear1StartIndex2627 int = 4923
const nsldsTEACHGrantAwardYear1Length2627 int = 4

// Field # 673
const nsldsTEACHGrantLoanConversionFlag1StartIndex2627 int = 4927
const nsldsTEACHGrantLoanConversionFlag1Length2627 int = 1

// Field # 674
const nsldsTEACHGrantDischargeCode1StartIndex2627 int = 4928
const nsldsTEACHGrantDischargeCode1Length2627 int = 4

// Field # 675
const nsldsTEACHGrantDischargeAmount1StartIndex2627 int = 4932
const nsldsTEACHGrantDischargeAmount1Length2627 int = 6

// Field # 676
const nsldsTEACHGrantAdjustedDisbursement1StartIndex2627 int = 4938
const nsldsTEACHGrantAdjustedDisbursement1Length2627 int = 6

// Field # 678
const nsldsTEACHGrantSequence2StartIndex2627 int = 4964
const nsldsTEACHGrantSequence2Length2627 int = 2

// Field # 679
const nsldsTEACHGrantSchoolCode2StartIndex2627 int = 4966
const nsldsTEACHGrantSchoolCode2Length2627 int = 8

// Field # 680
const nsldsTEACHGrantTransactionNumber2StartIndex2627 int = 4974
const nsldsTEACHGrantTransactionNumber2Length2627 int = 2

// Field # 681
const nsldsTEACHGrantLastDisbursementDate2StartIndex2627 int = 4976
const nsldsTEACHGrantLastDisbursementDate2Length2627 int = 8

// Field # 682
const nsldsTEACHGrantScheduledAmount2StartIndex2627 int = 4984
const nsldsTEACHGrantScheduledAmount2Length2627 int = 6

// Field # 683
const nsldsTEACHGrantAmountPaidToDate2StartIndex2627 int = 4990
const nsldsTEACHGrantAmountPaidToDate2Length2627 int = 6

// Field # 684
const nsldsTEACHGrantAwardAmount2StartIndex2627 int = 4996
const nsldsTEACHGrantAwardAmount2Length2627 int = 6

// Field # 685
const nsldsTEACHGrantAcademicYearLevel2StartIndex2627 int = 5002
const nsldsTEACHGrantAcademicYearLevel2Length2627 int = 1

// Field # 686
const nsldsTEACHGrantAwardYear2StartIndex2627 int = 5003
const nsldsTEACHGrantAwardYear2Length2627 int = 4

// Field # 687
const nsldsTEACHGrantLoanConversionFlag2StartIndex2627 int = 5007
const nsldsTEACHGrantLoanConversionFlag2Length2627 int = 1

// Field # 688
const nsldsTEACHGrantDischargeCode2StartIndex2627 int = 5008
const nsldsTEACHGrantDischargeCode2Length2627 int = 4

// Field # 689
const nsldsTEACHGrantDischargeAmount2StartIndex2627 int = 5012
const nsldsTEACHGrantDischargeAmount2Length2627 int = 6

// Field # 690
const nsldsTEACHGrantAdjustedDisbursement2StartIndex2627 int = 5018
const nsldsTEACHGrantAdjustedDisbursement2Length2627 int = 6

// Field # 692
const nsldsTEACHGrantSequence3StartIndex2627 int = 5044
const nsldsTEACHGrantSequence3Length2627 int = 2

// Field # 693
const nsldsTEACHGrantSchoolCode3StartIndex2627 int = 5046
const nsldsTEACHGrantSchoolCode3Length2627 int = 8

// Field # 694
const nsldsTEACHGrantTransactionNumber3StartIndex2627 int = 5054
const nsldsTEACHGrantTransactionNumber3Length2627 int = 2

// Field # 695
const nsldsTEACHGrantLastDisbursementDate3StartIndex2627 int = 5056
const nsldsTEACHGrantLastDisbursementDate3Length2627 int = 8

// Field # 696
const nsldsTEACHGrantScheduledAmount3StartIndex2627 int = 5064
const nsldsTEACHGrantScheduledAmount3Length2627 int = 6

// Field # 697
const nsldsTEACHGrantAmountPaidToDate3StartIndex2627 int = 5070
const nsldsTEACHGrantAmountPaidToDate3Length2627 int = 6

// Field # 698
const nsldsTEACHGrantAwardAmount3StartIndex2627 int = 5076
const nsldsTEACHGrantAwardAmount3Length2627 int = 6

// Field # 699
const nsldsTEACHGrantAcademicYearLevel3StartIndex2627 int = 5082
const nsldsTEACHGrantAcademicYearLevel3Length2627 int = 1

// Field # 700
const nsldsTEACHGrantAwardYear3StartIndex2627 int = 5083
const nsldsTEACHGrantAwardYear3Length2627 int = 4

// Field # 701
const nsldsTEACHGrantLoanConversionFlag3StartIndex2627 int = 5087
const nsldsTEACHGrantLoanConversionFlag3Length2627 int = 1

// Field # 702
const nsldsTEACHGrantDischargeCode3StartIndex2627 int = 5088
const nsldsTEACHGrantDischargeCode3Length2627 int = 4

// Field # 703
const nsldsTEACHGrantDischargeAmount3StartIndex2627 int = 5092
const nsldsTEACHGrantDischargeAmount3Length2627 int = 6

// Field # 704
const nsldsTEACHGrantAdjustedDisbursement3StartIndex2627 int = 5098
const nsldsTEACHGrantAdjustedDisbursement3Length2627 int = 6

// Field # 706
const nsldsLoanSequenceNumber1StartIndex2627 int = 5124
const nsldsLoanSequenceNumber1Length2627 int = 2

// Field # 707
const nsldsLoanDefaultedRecentIndicator1StartIndex2627 int = 5126
const nsldsLoanDefaultedRecentIndicator1Length2627 int = 1

// Field # 708
const nsldsLoanChangeFlag1StartIndex2627 int = 5127
const nsldsLoanChangeFlag1Length2627 int = 1

// Field # 709
const nsldsLoanTypeCode1StartIndex2627 int = 5128
const nsldsLoanTypeCode1Length2627 int = 2

// Field # 710
const nsldsLoanNetAmount1StartIndex2627 int = 5130
const nsldsLoanNetAmount1Length2627 int = 6

// Field # 711
const nsldsLoanCurrentStatusCode1StartIndex2627 int = 5136
const nsldsLoanCurrentStatusCode1Length2627 int = 2

// Field # 712
const nsldsLoanCurrentStatusDate1StartIndex2627 int = 5138
const nsldsLoanCurrentStatusDate1Length2627 int = 8

// Field # 713
const nsldsLoanOutstandingPrincipalBalance1StartIndex2627 int = 5146
const nsldsLoanOutstandingPrincipalBalance1Length2627 int = 6

// Field # 714
const nsldsLoanOutstandingPrincipalBalanceDate1StartIndex2627 int = 5152
const nsldsLoanOutstandingPrincipalBalanceDate1Length2627 int = 8

// Field # 715
const nsldsLoanPeriodBeginDate1StartIndex2627 int = 5160
const nsldsLoanPeriodBeginDate1Length2627 int = 8

// Field # 716
const nsldsLoanPeriodEndDate1StartIndex2627 int = 5168
const nsldsLoanPeriodEndDate1Length2627 int = 8

// Field # 717
const nsldsLoanGuarantyAgencyCode1StartIndex2627 int = 5176
const nsldsLoanGuarantyAgencyCode1Length2627 int = 3

// Field # 718
const nsldsLoanContactType1StartIndex2627 int = 5179
const nsldsLoanContactType1Length2627 int = 3

// Field # 719
const nsldsLoanSchoolCode1StartIndex2627 int = 5182
const nsldsLoanSchoolCode1Length2627 int = 8

// Field # 720
const nsldsLoanContactCode1StartIndex2627 int = 5190
const nsldsLoanContactCode1Length2627 int = 8

// Field # 721
const nsldsLoanGradeLevel1StartIndex2627 int = 5198
const nsldsLoanGradeLevel1Length2627 int = 3

// Field # 722
const nsldsLoanAdditionalUnsubsidizedFlag1StartIndex2627 int = 5201
const nsldsLoanAdditionalUnsubsidizedFlag1Length2627 int = 1

// Field # 723
const nsldsLoanCapitalizedInterestFlag1StartIndex2627 int = 5202
const nsldsLoanCapitalizedInterestFlag1Length2627 int = 1

// Field # 724
const nsldsLoanDisbursementAmount1StartIndex2627 int = 5203
const nsldsLoanDisbursementAmount1Length2627 int = 6

// Field # 725
const nsldsLoanDisbursementDate1StartIndex2627 int = 5209
const nsldsLoanDisbursementDate1Length2627 int = 8

// Field # 726
const nsldsLoanConfirmedLoanSubsidyStatus1StartIndex2627 int = 5217
const nsldsLoanConfirmedLoanSubsidyStatus1Length2627 int = 1

// Field # 727
const nsldsLoanConfirmedLoanSubsidyStatusDate1StartIndex2627 int = 5218
const nsldsLoanConfirmedLoanSubsidyStatusDate1Length2627 int = 8

// Field # 728 FILLER

// Field # 729
const nsldsLoanSequenceNumber2StartIndex2627 int = 5246
const nsldsLoanSequenceNumber2Length2627 int = 2

// Field # 730
const nsldsLoanDefaultedRecentIndicator2StartIndex2627 int = 5248
const nsldsLoanDefaultedRecentIndicator2Length2627 int = 1

// Field # 731
const nsldsLoanChangeFlag2StartIndex2627 int = 5249
const nsldsLoanChangeFlag2Length2627 int = 1

// Field # 732
const nsldsLoanTypeCode2StartIndex2627 int = 5250
const nsldsLoanTypeCode2Length2627 int = 2

// Field # 733
const nsldsLoanNetAmount2StartIndex2627 int = 5252
const nsldsLoanNetAmount2Length2627 int = 6

// Field # 734
const nsldsLoanCurrentStatusCode2StartIndex2627 int = 5258
const nsldsLoanCurrentStatusCode2Length2627 int = 2

// Field # 735
const nsldsLoanCurrentStatusDate2StartIndex2627 int = 5260
const nsldsLoanCurrentStatusDate2Length2627 int = 8

// Field # 736
const nsldsLoanOutstandingPrincipalBalance2StartIndex2627 int = 5268
const nsldsLoanOutstandingPrincipalBalance2Length2627 int = 6

// Field # 737
const nsldsLoanOutstandingPrincipalBalanceDate2StartIndex2627 int = 5274
const nsldsLoanOutstandingPrincipalBalanceDate2Length2627 int = 8

// Field # 738
const nsldsLoanPeriodBeginDate2StartIndex2627 int = 5282
const nsldsLoanPeriodBeginDate2Length2627 int = 8

// Field # 739
const nsldsLoanPeriodEndDate2StartIndex2627 int = 5290
const nsldsLoanPeriodEndDate2Length2627 int = 8

// Field # 740
const nsldsLoanGuarantyAgencyCode2StartIndex2627 int = 5298
const nsldsLoanGuarantyAgencyCode2Length2627 int = 3

// Field # 741
const nsldsLoanContactType2StartIndex2627 int = 5301
const nsldsLoanContactType2Length2627 int = 3

// Field # 742
const nsldsLoanSchoolCode2StartIndex2627 int = 5304
const nsldsLoanSchoolCode2Length2627 int = 8

// Field # 743
const nsldsLoanContactCode2StartIndex2627 int = 5312
const nsldsLoanContactCode2Length2627 int = 8

// Field # 744
const nsldsLoanGradeLevel2StartIndex2627 int = 5320
const nsldsLoanGradeLevel2Length2627 int = 3

// Field # 745
const nsldsLoanAdditionalUnsubsidizedFlag2StartIndex2627 int = 5323
const nsldsLoanAdditionalUnsubsidizedFlag2Length2627 int = 1

// Field # 746
const nsldsLoanCapitalizedInterestFlag2StartIndex2627 int = 5324
const nsldsLoanCapitalizedInterestFlag2Length2627 int = 1

// Field # 747
const nsldsLoanDisbursementAmount2StartIndex2627 int = 5325
const nsldsLoanDisbursementAmount2Length2627 int = 6

// Field # 748
const nsldsLoanDisbursementDate2StartIndex2627 int = 5331
const nsldsLoanDisbursementDate2Length2627 int = 8

// Field # 749
const nsldsLoanConfirmedLoanSubsidyStatus2StartIndex2627 int = 5339
const nsldsLoanConfirmedLoanSubsidyStatus2Length2627 int = 1

// Field # 750
const nsldsLoanConfirmedLoanSubsidyStatusDate2StartIndex2627 int = 5340
const nsldsLoanConfirmedLoanSubsidyStatusDate2Length2627 int = 8

// Field # 751 FILLER

// Field # 752
const nsldsLoanSequenceNumber3StartIndex2627 int = 5368
const nsldsLoanSequenceNumber3Length2627 int = 2

// Field # 753
const nsldsLoanDefaultedRecentIndicator3StartIndex2627 int = 5370
const nsldsLoanDefaultedRecentIndicator3Length2627 int = 1

// Field # 754
const nsldsLoanChangeFlag3StartIndex2627 int = 5371
const nsldsLoanChangeFlag3Length2627 int = 1

// Field # 755
const nsldsLoanTypeCode3StartIndex2627 int = 5372
const nsldsLoanTypeCode3Length2627 int = 2

// Field # 756
const nsldsLoanNetAmount3StartIndex2627 int = 5374
const nsldsLoanNetAmount3Length2627 int = 6

// Field # 757
const nsldsLoanCurrentStatusCode3StartIndex2627 int = 5380
const nsldsLoanCurrentStatusCode3Length2627 int = 2

// Field # 758
const nsldsLoanCurrentStatusDate3StartIndex2627 int = 5382
const nsldsLoanCurrentStatusDate3Length2627 int = 8

// Field # 759
const nsldsLoanOutstandingPrincipalBalance3StartIndex2627 int = 5390
const nsldsLoanOutstandingPrincipalBalance3Length2627 int = 6

// Field # 760
const nsldsLoanOutstandingPrincipalBalanceDate3StartIndex2627 int = 5396
const nsldsLoanOutstandingPrincipalBalanceDate3Length2627 int = 8

// Field # 761
const nsldsLoanPeriodBeginDate3StartIndex2627 int = 5404
const nsldsLoanPeriodBeginDate3Length2627 int = 8

// Field # 762
const nsldsLoanPeriodEndDate3StartIndex2627 int = 5412
const nsldsLoanPeriodEndDate3Length2627 int = 8

// Field # 763
const nsldsLoanGuarantyAgencyCode3StartIndex2627 int = 5420
const nsldsLoanGuarantyAgencyCode3Length2627 int = 3

// Field # 764
const nsldsLoanContactType3StartIndex2627 int = 5423
const nsldsLoanContactType3Length2627 int = 3

// Field # 765
const nsldsLoanSchoolCode3StartIndex2627 int = 5426
const nsldsLoanSchoolCode3Length2627 int = 8

// Field # 766
const nsldsLoanContactCode3StartIndex2627 int = 5434
const nsldsLoanContactCode3Length2627 int = 8

// Field # 767
const nsldsLoanGradeLevel3StartIndex2627 int = 5442
const nsldsLoanGradeLevel3Length2627 int = 3

// Field # 768
const nsldsLoanAdditionalUnsubsidizedFlag3StartIndex2627 int = 5445
const nsldsLoanAdditionalUnsubsidizedFlag3Length2627 int = 1

// Field # 769
const nsldsLoanCapitalizedInterestFlag3StartIndex2627 int = 5446
const nsldsLoanCapitalizedInterestFlag3Length2627 int = 1

// Field # 770
const nsldsLoanDisbursementAmount3StartIndex2627 int = 5447
const nsldsLoanDisbursementAmount3Length2627 int = 6

// Field # 771
const nsldsLoanDisbursementDate3StartIndex2627 int = 5453
const nsldsLoanDisbursementDate3Length2627 int = 8

// Field # 772
const nsldsLoanConfirmedLoanSubsidyStatus3StartIndex2627 int = 5461
const nsldsLoanConfirmedLoanSubsidyStatus3Length2627 int = 1

// Field # 773
const nsldsLoanConfirmedLoanSubsidyStatusDate3StartIndex2627 int = 5462
const nsldsLoanConfirmedLoanSubsidyStatusDate3Length2627 int = 8

// Field # 774 FILLER

// Field # 775
const nsldsLoanSequenceNumber4StartIndex2627 int = 5490
const nsldsLoanSequenceNumber4Length2627 int = 2

// Field # 776
const nsldsLoanDefaultedRecentIndicator4StartIndex2627 int = 5492
const nsldsLoanDefaultedRecentIndicator4Length2627 int = 1

// Field # 777
const nsldsLoanChangeFlag4StartIndex2627 int = 5493
const nsldsLoanChangeFlag4Length2627 int = 1

// Field # 778
const nsldsLoanTypeCode4StartIndex2627 int = 5494
const nsldsLoanTypeCode4Length2627 int = 2

// Field # 779
const nsldsLoanNetAmount4StartIndex2627 int = 5496
const nsldsLoanNetAmount4Length2627 int = 6

// Field # 780
const nsldsLoanCurrentStatusCode4StartIndex2627 int = 5502
const nsldsLoanCurrentStatusCode4Length2627 int = 2

// Field # 781
const nsldsLoanCurrentStatusDate4StartIndex2627 int = 5504
const nsldsLoanCurrentStatusDate4Length2627 int = 8

// Field # 782
const nsldsLoanOutstandingPrincipalBalance4StartIndex2627 int = 5512
const nsldsLoanOutstandingPrincipalBalance4Length2627 int = 6

// Field # 783
const nsldsLoanOutstandingPrincipalBalanceDate4StartIndex2627 int = 5518
const nsldsLoanOutstandingPrincipalBalanceDate4Length2627 int = 8

// Field # 784
const nsldsLoanPeriodBeginDate4StartIndex2627 int = 5526
const nsldsLoanPeriodBeginDate4Length2627 int = 8

// Field # 785
const nsldsLoanPeriodEndDate4StartIndex2627 int = 5534
const nsldsLoanPeriodEndDate4Length2627 int = 8

// Field # 786
const nsldsLoanGuarantyAgencyCode4StartIndex2627 int = 5542
const nsldsLoanGuarantyAgencyCode4Length2627 int = 3

// Field # 787
const nsldsLoanContactType4StartIndex2627 int = 5545
const nsldsLoanContactType4Length2627 int = 3

// Field # 788
const nsldsLoanSchoolCode4StartIndex2627 int = 5548
const nsldsLoanSchoolCode4Length2627 int = 8

// Field # 789
const nsldsLoanContactCode4StartIndex2627 int = 5556
const nsldsLoanContactCode4Length2627 int = 8

// Field # 790
const nsldsLoanGradeLevel4StartIndex2627 int = 5564
const nsldsLoanGradeLevel4Length2627 int = 3

// Field # 791
const nsldsLoanAdditionalUnsubsidizedFlag4StartIndex2627 int = 5567
const nsldsLoanAdditionalUnsubsidizedFlag4Length2627 int = 1

// Field # 792
const nsldsLoanCapitalizedInterestFlag4StartIndex2627 int = 5568
const nsldsLoanCapitalizedInterestFlag4Length2627 int = 1

// Field # 793
const nsldsLoanDisbursementAmount4StartIndex2627 int = 5569
const nsldsLoanDisbursementAmount4Length2627 int = 6

// Field # 794
const nsldsLoanDisbursementDate4StartIndex2627 int = 5575
const nsldsLoanDisbursementDate4Length2627 int = 8

// Field # 795
const nsldsLoanConfirmedLoanSubsidyStatus4StartIndex2627 int = 5583
const nsldsLoanConfirmedLoanSubsidyStatus4Length2627 int = 1

// Field # 796
const nsldsLoanSubsidyStatusDate4StartIndex2627 int = 5584
const nsldsLoanSubsidyStatusDate4Length2627 int = 8

// Field # 797 FILLER

// Field # 798
const nsldsLoanSequenceNumber5StartIndex2627 int = 5612
const nsldsLoanSequenceNumber5Length2627 int = 2

// Field # 799
const nsldsLoanDefaultedRecentIndicator5StartIndex2627 int = 5614
const nsldsLoanDefaultedRecentIndicator5Length2627 int = 1

// Field # 800
const nsldsLoanChangeFlag5StartIndex2627 int = 5615
const nsldsLoanChangeFlag5Length2627 int = 1

// Field # 801
const nsldsLoanTypeCode5StartIndex2627 int = 5616
const nsldsLoanTypeCode5Length2627 int = 2

// Field # 802
const nsldsLoanNetAmount5StartIndex2627 int = 5618
const nsldsLoanNetAmount5Length2627 int = 6

// Field # 803
const nsldsLoanCurrentStatusCode5StartIndex2627 int = 5624
const nsldsLoanCurrentStatusCode5Length2627 int = 2

// Field # 804
const nsldsLoanCurrentStatusDate5StartIndex2627 int = 5626
const nsldsLoanCurrentStatusDate5Length2627 int = 8

// Field # 805
const nsldsLoanOutstandingPrincipalBalance5StartIndex2627 int = 5634
const nsldsLoanOutstandingPrincipalBalance5Length2627 int = 6

// Field # 806
const nsldsLoanOutstandingPrincipalBalanceDate5StartIndex2627 int = 5640
const nsldsLoanOutstandingPrincipalBalanceDate5Length2627 int = 8

// Field # 807
const nsldsLoanPeriodBeginDate5StartIndex2627 int = 5648
const nsldsLoanPeriodBeginDate5Length2627 int = 8

// Field # 808
const nsldsLoanPeriodEndDate5StartIndex2627 int = 5656
const nsldsLoanPeriodEndDate5Length2627 int = 8

// Field # 809
const nsldsLoanGuarantyAgencyCode5StartIndex2627 int = 5664
const nsldsLoanGuarantyAgencyCode5Length2627 int = 3

// Field # 810
const nsldsLoanContactType5StartIndex2627 int = 5667
const nsldsLoanContactType5Length2627 int = 3

// Field # 811
const nsldsLoanSchoolCode5StartIndex2627 int = 5670
const nsldsLoanSchoolCode5Length2627 int = 8

// Field # 812
const nsldsLoanContactCode5StartIndex2627 int = 5678
const nsldsLoanContactCode5Length2627 int = 8

// Field # 813
const nsldsLoanGradeLevel5StartIndex2627 int = 5686
const nsldsLoanGradeLevel5Length2627 int = 3

// Field # 814
const nsldsLoanAdditionalUnsubsidizedFlag5StartIndex2627 int = 5689
const nsldsLoanAdditionalUnsubsidizedFlag5Length2627 int = 1

// Field # 815
const nsldsLoanCapitalizedInterestFlag5StartIndex2627 int = 5690
const nsldsLoanCapitalizedInterestFlag5Length2627 int = 1

// Field # 816
const nsldsLoanDisbursementAmount5StartIndex2627 int = 5691
const nsldsLoanDisbursementAmount5Length2627 int = 6

// Field # 817
const nsldsLoanDisbursementDate5StartIndex2627 int = 5697
const nsldsLoanDisbursementDate5Length2627 int = 8

// Field # 818
const nsldsLoanConfirmedLoanSubsidyStatus5StartIndex2627 int = 5705
const nsldsLoanConfirmedLoanSubsidyStatus5Length2627 int = 1

// Field # 819
const nsldsLoanSubsidyStatusDate5StartIndex2627 int = 5706
const nsldsLoanSubsidyStatusDate5Length2627 int = 8

// Field #820 FILLER

// Field # 821
const nsldsLoanSequenceNumber6StartIndex2627 int = 5734
const nsldsLoanSequenceNumber6Length2627 int = 2

// Field # 822
const nsldsLoanDefaultedRecentIndicator6StartIndex2627 int = 5736
const nsldsLoanDefaultedRecentIndicator6Length2627 int = 1

// Field # 823
const nsldsLoanChangeFlag6StartIndex2627 int = 5737
const nsldsLoanChangeFlag6Length2627 int = 1

// Field # 824
const nsldsLoanTypeCode6StartIndex2627 int = 5738
const nsldsLoanTypeCode6Length2627 int = 2

// Field # 825
const nsldsLoanNetAmount6StartIndex2627 int = 5740
const nsldsLoanNetAmount6Length2627 int = 6

// Field # 826
const nsldsLoanCurrentStatusCode6StartIndex2627 int = 5746
const nsldsLoanCurrentStatusCode6Length2627 int = 2

// Field # 827
const nsldsLoanCurrentStatusDate6StartIndex2627 int = 5748
const nsldsLoanCurrentStatusDate6Length2627 int = 8

// Field # 828
const nsldsLoanOutstandingPrincipalBalance6StartIndex2627 int = 5756
const nsldsLoanOutstandingPrincipalBalance6Length2627 int = 6

// Field # 829
const nsldsLoanOutstandingPrincipalBalanceDate6StartIndex2627 int = 5762
const nsldsLoanOutstandingPrincipalBalanceDate6Length2627 int = 8

// Field # 830
const nsldsLoanPeriodBeginDate6StartIndex2627 int = 5770
const nsldsLoanPeriodBeginDate6Length2627 int = 8

// Field # 831
const nsldsLoanPeriodEndDate6StartIndex2627 int = 5778
const nsldsLoanPeriodEndDate6Length2627 int = 8

// Field # 832
const nsldsLoanGuarantyAgencyCode6StartIndex2627 int = 5786
const nsldsLoanGuarantyAgencyCode6Length2627 int = 3

// Field # 833
const nsldsLoanContactType6StartIndex2627 int = 5789
const nsldsLoanContactType6Length2627 int = 3

// Field # 834
const nsldsLoanSchoolCode6StartIndex2627 int = 5792
const nsldsLoanSchoolCode6Length2627 int = 8

// Field # 835
const nsldsLoanContactCode6StartIndex2627 int = 5800
const nsldsLoanContactCode6Length2627 int = 8

// Field # 836
const nsldsLoanGradeLevel6StartIndex2627 int = 5808
const nsldsLoanGradeLevel6Length2627 int = 3

// Field # 837
const nsldsLoanAdditionalUnsubsidizedFlag6StartIndex2627 int = 5811
const nsldsLoanAdditionalUnsubsidizedFlag6Length2627 int = 1

// Field # 838
const nsldsLoanCapitalizedInterestFlag6StartIndex2627 int = 5812
const nsldsLoanCapitalizedInterestFlag6Length2627 int = 1

// Field # 839
const nsldsLoanDisbursementAmount6StartIndex2627 int = 5813
const nsldsLoanDisbursementAmount6Length2627 int = 6

// Field # 840
const nsldsLoanDisbursementDate6StartIndex2627 int = 5819
const nsldsLoanDisbursementDate6Length2627 int = 8

// Field # 841
const nsldsLoanConfirmedLoanSubsidyStatus6StartIndex2627 int = 5827
const nsldsLoanConfirmedLoanSubsidyStatus6Length2627 int = 1

// Field # 842
const nsldsLoanSubsidyStatusDate6StartIndex2627 int = 5828
const nsldsLoanSubsidyStatusDate6Length2627 int = 8

// Field # 843 FILLER

// Field # 844 FILLER

// Field # 845 FILLER

// Field # 846
//const ftiLabelStartStartIndex2627 int = 7086
//const ftiLabelStartLength2627 int = 11

// Field # 847
const studentFTIMReturnedTaxYearStartIndex2627 int = 7097
const studentFTIMReturnedTaxYearLength2627 int = 4

// Field # 848
const studentFTIMFilingStatusCodeStartIndex2627 int = 7101
const studentFTIMFilingStatusCodeLength2627 int = 1

// Field # 849
const studentFTIMAdjustedGrossIncomeStartIndex2627 int = 7102
const studentFTIMAdjustedGrossIncomeLength2627 int = 10

// Field # 850
const studentFTIMNumberOfExemptionsStartIndex2627 int = 7112
const studentFTIMNumberOfExemptionsLength2627 int = 2

// Field # 851
const studentFTIMNumberOfDependentsStartIndex2627 int = 7114
const studentFTIMNumberOfDependentsLength2627 int = 2

// Field # 852
const studentFTIMTotalIncomeEarnedAmountStartIndex2627 int = 7116
const studentFTIMTotalIncomeEarnedAmountLength2627 int = 11

// Field # 853
const studentFTIMIncomeTaxPaidStartIndex2627 int = 7127
const studentFTIMIncomeTaxPaidLength2627 int = 9

// Field # 854
const studentFTIMEducationCreditsStartIndex2627 int = 7136
const studentFTIMEducationCreditsLength2627 int = 9

// Field # 855
const studentFTIMUntaxedIRADistributionsStartIndex2627 int = 7145
const studentFTIMUntaxedIRADistributionsLength2627 int = 11

// Field # 856
const studentFTIMIRADeductibleAndPaymentsStartIndex2627 int = 7156
const studentFTIMIRADeductibleAndPaymentsLength2627 int = 11

// Field # 857
const studentFTIMTaxExemptInterestStartIndex2627 int = 7167
const studentFTIMTaxExemptInterestLength2627 int = 11

// Field # 858
const studentFTIMUntaxedPensionsAmountStartIndex2627 int = 7178
const studentFTIMUntaxedPensionsAmountLength2627 int = 11

// Field # 859
const studentFTIMScheduleCNetProfitLossStartIndex2627 int = 7189
const studentFTIMScheduleCNetProfitLossLength2627 int = 12

// Field # 860
const studentFTIMScheduleAIndicatorStartIndex2627 int = 7201
const studentFTIMScheduleAIndicatorLength2627 int = 1

// Field # 861
const studentFTIMScheduleBIndicatorStartIndex2627 int = 7202
const studentFTIMScheduleBIndicatorLength2627 int = 1

// Field # 862
const studentFTIMScheduleDIndicatorStartIndex2627 int = 7203
const studentFTIMScheduleDIndicatorLength2627 int = 1

// Field # 863
const studentFTIMScheduleEIndicatorStartIndex2627 int = 7204
const studentFTIMScheduleEIndicatorLength2627 int = 1

// Field # 864
const studentFTIMScheduleFIndicatorStartIndex2627 int = 7205
const studentFTIMScheduleFIndicatorLength2627 int = 1

// Field # 865
const studentFTIMScheduleHIndicatorStartIndex2627 int = 7206
const studentFTIMScheduleHIndicatorLength2627 int = 1

// Field # 866
const studentFTIMIRSResponseCodeStartIndex2627 int = 7207
const studentFTIMIRSResponseCodeLength2627 int = 3

// Field # 867
const studentFTIMSpouseReturnedTaxYearStartIndex2627 int = 7210
const studentFTIMSpouseReturnedTaxYearLength2627 int = 4

// Field # 868
const studentFTIMSpouseFilingStatusCodeStartIndex2627 int = 7214
const studentFTIMSpouseFilingStatusCodeLength2627 int = 1

// Field # 869
const studentFTIMSpouseAdjustedGrossIncomeStartIndex2627 int = 7215
const studentFTIMSpouseAdjustedGrossIncomeLength2627 int = 10

// Field # 870
const studentFTIMSpouseNumberOfExemptionsStartIndex2627 int = 7225
const studentFTIMSpouseNumberOfExemptionsLength2627 int = 2

// Field # 871
const studentFTIMSpouseNumberOfDependentsStartIndex2627 int = 7227
const studentFTIMSpouseNumberOfDependentsLength2627 int = 2

// Field # 872
const studentFTIMSpouseTotalIncomeEarnedAmountStartIndex2627 int = 7229
const studentFTIMSpouseTotalIncomeEarnedAmountLength2627 int = 11

// Field # 873
const studentFTIMSpouseIncomeTaxPaidStartIndex2627 int = 7240
const studentFTIMSpouseIncomeTaxPaidLength2627 int = 9

// Field # 874
const studentFTIMSpouseEducationCreditsStartIndex2627 int = 7249
const studentFTIMSpouseEducationCreditsLength2627 int = 9

// Field # 875
const studentFTIMSpouseUntaxedIRADistributionsStartIndex2627 int = 7258
const studentFTIMSpouseUntaxedIRADistributionsLength2627 int = 11

// Field # 876
const studentFTIMSpouseIRADeductibleAndPaymentsStartIndex2627 int = 7269
const studentFTIMSpouseIRADeductibleAndPaymentsLength2627 int = 11

// Field # 877
const studentFTIMSpouseTaxExemptInterestStartIndex2627 int = 7280
const studentFTIMSpouseTaxExemptInterestLength2627 int = 11

// Field # 878
const studentFTIMSpouseUntaxedPensionsAmountStartIndex2627 int = 7291
const studentFTIMSpouseUntaxedPensionsAmountLength2627 int = 11

// Field # 879
const studentFTIMSpouseScheduleCNetProfitLossStartIndex2627 int = 7302
const studentFTIMSpouseScheduleCNetProfitLossLength2627 int = 12

// Field # 880
const studentFTIMSpouseScheduleAIndicatorStartIndex2627 int = 7314
const studentFTIMSpouseScheduleAIndicatorLength2627 int = 1

// Field # 881
const studentFTIMSpouseScheduleBIndicatorStartIndex2627 int = 7315
const studentFTIMSpouseScheduleBIndicatorLength2627 int = 1

// Field # 882
const studentFTIMSpouseScheduleDIndicatorStartIndex2627 int = 7316
const studentFTIMSpouseScheduleDIndicatorLength2627 int = 1

// Field # 883
const studentFTIMSpouseScheduleEIndicatorStartIndex2627 int = 7317
const studentFTIMSpouseScheduleEIndicatorLength2627 int = 1

// Field # 884
const studentFTIMSpouseScheduleFIndicatorStartIndex2627 int = 7318
const studentFTIMSpouseScheduleFIndicatorLength2627 int = 1

// Field # 885
const studentFTIMSpouseScheduleHIndicatorStartIndex2627 int = 7319
const studentFTIMSpouseScheduleHIndicatorLength2627 int = 1

// Field # 886
const studentFTIMSpouseIRSResponseCodeStartIndex2627 int = 7320
const studentFTIMSpouseIRSResponseCodeLength2627 int = 3

// Field # 887
const parentFTIMReturnedTaxYearStartIndex2627 int = 7323
const parentFTIMReturnedTaxYearLength2627 int = 4

// Field # 888
const parentFTIMFilingStatusCodeStartIndex2627 int = 7327
const parentFTIMFilingStatusCodeLength2627 int = 1

// Field # 889
const parentFTIMAdjustedGrossIncomeStartIndex2627 int = 7328
const parentFTIMAdjustedGrossIncomeLength2627 int = 10

// Field # 890
const parentFTIMNumberOfExemptionsStartIndex2627 int = 7338
const parentFTIMNumberOfExemptionsLength2627 int = 2

// Field # 891
const parentFTIMNumberOfDependentsStartIndex2627 int = 7340
const parentFTIMNumberOfDependentsLength2627 int = 2

// Field # 892
const parentFTIMTotalIncomeEarnedAmountStartIndex2627 int = 7342
const parentFTIMTotalIncomeEarnedAmountLength2627 int = 11

// Field # 893
const parentFTIMIncomeTaxPaidStartIndex2627 int = 7353
const parentFTIMIncomeTaxPaidLength2627 int = 9

// Field # 894
const parentFTIMEducationCreditsStartIndex2627 int = 7362
const parentFTIMEducationCreditsLength2627 int = 9

// Field # 895
const parentFTIMUntaxedIRADistributionsStartIndex2627 int = 7371
const parentFTIMUntaxedIRADistributionsLength2627 int = 11

// Field # 896
const parentFTIMIRADeductibleAndPaymentsStartIndex2627 int = 7382
const parentFTIMIRADeductibleAndPaymentsLength2627 int = 11

// Field # 897
const parentFTIMTaxExemptInterestStartIndex2627 int = 7393
const parentFTIMTaxExemptInterestLength2627 int = 11

// Field # 898
const parentFTIMUntaxedPensionsAmountStartIndex2627 int = 7404
const parentFTIMUntaxedPensionsAmountLength2627 int = 11

// Field # 899
const parentFTIMScheduleCNetProfitLossStartIndex2627 int = 7415
const parentFTIMScheduleCNetProfitLossLength2627 int = 12

// Field # 900
const parentFTIMScheduleAIndicatorStartIndex2627 int = 7427
const parentFTIMScheduleAIndicatorLength2627 int = 1

// Field # 901
const parentFTIMScheduleBIndicatorStartIndex2627 int = 7428
const parentFTIMScheduleBIndicatorLength2627 int = 1

// Field # 902
const parentFTIMScheduleDIndicatorStartIndex2627 int = 7429
const parentFTIMScheduleDIndicatorLength2627 int = 1

// Field # 903
const parentFTIMScheduleEIndicatorStartIndex2627 int = 7430
const parentFTIMScheduleEIndicatorLength2627 int = 1

// Field # 904
const parentFTIMScheduleFIndicatorStartIndex2627 int = 7431
const parentFTIMScheduleFIndicatorLength2627 int = 1

// Field # 905
const parentFTIMScheduleHIndicatorStartIndex2627 int = 7432
const parentFTIMScheduleHIndicatorLength2627 int = 1

// Field # 906
const parentFTIMIRSResponseCodeStartIndex2627 int = 7433
const parentFTIMIRSResponseCodeLength2627 int = 3

// Field # 907
const parentFTIMSpouseReturnedTaxYearStartIndex2627 int = 7436
const parentFTIMSpouseReturnedTaxYearLength2627 int = 4

// Field # 908
const parentFTIMSpouseFilingStatusCodeStartIndex2627 int = 7440
const parentFTIMSpouseFilingStatusCodeLength2627 int = 1

// Field # 909
const parentFTIMSpouseAdjustedGrossIncomeStartIndex2627 int = 7441
const parentFTIMSpouseAdjustedGrossIncomeLength2627 int = 10

// Field # 910
const parentFTIMSpouseNumberOfExemptionsStartIndex2627 int = 7451
const parentFTIMSpouseNumberOfExemptionsLength2627 int = 2

// Field # 911
const parentFTIMSpouseNumberOfDependentsStartIndex2627 int = 7453
const parentFTIMSpouseNumberOfDependentsLength2627 int = 2

// Field # 912
const parentFTIMSpouseTotalIncomeEarnedAmountStartIndex2627 int = 7455
const parentFTIMSpouseTotalIncomeEarnedAmountLength2627 int = 11

// Field # 913
const parentFTIMSpouseIncomeTaxPaidStartIndex2627 int = 7466
const parentFTIMSpouseIncomeTaxPaidLength2627 int = 9

// Field # 914
const parentFTIMSpouseEducationCreditsStartIndex2627 int = 7475
const parentFTIMSpouseEducationCreditsLength2627 int = 9

// Field # 915
const parentFTIMSpouseUntaxedIRADistributionsStartIndex2627 int = 7484
const parentFTIMSpouseUntaxedIRADistributionsLength2627 int = 11

// Field # 916
const parentFTIMSpouseIRADeductibleAndPaymentsStartIndex2627 int = 7495
const parentFTIMSpouseIRADeductibleAndPaymentsLength2627 int = 11

// Field # 917
const parentFTIMSpouseTaxExemptInterestStartIndex2627 int = 7506
const parentFTIMSpouseTaxExemptInterestLength2627 int = 11

// Field # 918
const parentFTIMSpouseUntaxedPensionsAmountStartIndex2627 int = 7517
const parentFTIMSpouseUntaxedPensionsAmountLength2627 int = 11

// Field # 919
const parentFTIMSpouseScheduleCNetProfitLossStartIndex2627 int = 7528
const parentFTIMSpouseScheduleCNetProfitLossLength2627 int = 12

// Field # 920
const parentFTIMSpouseScheduleAIndicatorStartIndex2627 int = 7540
const parentFTIMSpouseScheduleAIndicatorLength2627 int = 1

// Field # 921
const parentFTIMSpouseScheduleBIndicatorStartIndex2627 int = 7541
const parentFTIMSpouseScheduleBIndicatorLength2627 int = 1

// Field # 922
const parentFTIMSpouseScheduleDIndicatorStartIndex2627 int = 7542
const parentFTIMSpouseScheduleDIndicatorLength2627 int = 1

// Field # 923
const parentFTIMSpouseScheduleEIndicatorStartIndex2627 int = 7543
const parentFTIMSpouseScheduleEIndicatorLength2627 int = 1

// Field # 924
const parentFTIMSpouseScheduleFIndicatorStartIndex2627 int = 7544
const parentFTIMSpouseScheduleFIndicatorLength2627 int = 1

// Field # 925
const parentFTIMSpouseScheduleHIndicatorStartIndex2627 int = 7545
const parentFTIMSpouseScheduleHIndicatorLength2627 int = 1

// Field # 926
const parentFTIMSpouseIRSResponseCodeStartIndex2627 int = 7546
const parentFTIMSpouseIRSResponseCodeLength2627 int = 3

// Field # 927 FILLER

// Field # 92
//const ftiLabelEndStartIndex2627 int = 7549
//const ftiLabelEndLength2627 int = 11

// FTIM Processing Information

// Field # 928
const studentTotalIncomeStartIndex2627 int = 7610
const studentTotalIncomeLength2627 int = 15

// Field # 929
const parentTotalIncomeStartIndex2627 int = 7625
const parentTotalIncomeLength2627 int = 15

// Field # 930
const fisapTotalIncomeStartIndex2627 int = 7640
const fisapTotalIncomeLength2627 int = 15

// Fields below moved from above for 26-27

// Field # 931
const totalParentAllowancesAgainstIncomeStartIndex2627 int = 7655
const totalParentAllowancesAgainstIncomeLength2627 int = 15

// Field # 932
const parentPayrollTaxAllowanceStartIndex2627 int = 7670
const parentPayrollTaxAllowanceLength2627 int = 15

// Field # 933
const parentIncomeProtectionAllowanceStartIndex2627 int = 7685
const parentIncomeProtectionAllowanceLength2627 int = 15

// Field # 934
const parentEmploymentExpenseAllowanceStartIndex2627 int = 7700
const parentEmploymentExpenseAllowanceLength2627 int = 15

// Field # 935
const parentAvailableIncomeStartIndex2627 int = 7715
const parentAvailableIncomeLength2627 int = 15

// Field # 936
const parentAdjustedAvailableIncomeStartIndex2627 int = 7730
const parentAdjustedAvailableIncomeLength2627 int = 15

// Field # 937
const parentContributionStartIndex2627 int = 7745
const parentContributionLength2627 int = 15

// Field # 938
const studentPayrollTaxAllowanceStartIndex2627 int = 7760
const studentPayrollTaxAllowanceLength2627 int = 15

// Field # 939
const studentIncomeProtectionAllowanceStartIndex2627 int = 7775
const studentIncomeProtectionAllowanceLength2627 int = 15

// Field # 940
const studentAllowanceForParentsNegativeAdjustedAvailableIncomeStartIndex2627 int = 7790
const studentAllowanceForParentsNegativeAdjustedAvailableIncomeLength2627 int = 15

// Field # 941
const studentEmploymentExpenseAllowanceStartIndex2627 int = 7805
const studentEmploymentExpenseAllowanceLength2627 int = 15

// Field # 942
const totalStudentAllowancesAgainstIncomeStartIndex2627 int = 7820
const totalStudentAllowancesAgainstIncomeLength2627 int = 15

// Field # 943
const studentAvailableIncomeStartIndex2627 int = 7835
const studentAvailableIncomeLength2627 int = 15

// Field # 944
const studentContributionFromIncomeStartIndex2627 int = 7850
const studentContributionFromIncomeLength2627 int = 15

// Field # 945
const studentAdjustedAvailableIncomeStartIndex2627 int = 7865
const studentAdjustedAvailableIncomeLength2627 int = 15

// Field # 946
const totalStudentContributionFromSAAIStartIndex2627 int = 7880
const totalStudentContributionFromSAAILength2627 int = 15

//</editor-fold>

type ISIRParser2627 struct {
}

func (parser *ISIRParser2627) ParseISIR(record string, cid uuid.UUID) (fsamodels.ISIRecord, *fsaerrors.Error) {
	if len(record) != totalISIRLength2627 {
		return fsamodels.ISIRecord{}, &fsaerrors.Error{
			Code:          fsaerrors.ISIRParseError,
			Message:       fmt.Sprintf("input ISIR string is the incorrect length, expected %d and received %d, correlation id='%s'", totalISIRLength2627, len(record), cid.String()),
			Func:          "ParseISIR()",
			Record:        record,
			CorrelationID: cid.String(),
		}
	}

	// <editor-fold desc="Parsing Fields">
	r := fsamodels.ISIRecord{
		YearIndicator: preprocessString2627(record[yearIndicatorStartIndex2627-1 : (yearIndicatorStartIndex2627-1)+yearIndicatorLength2627]), // Field # 1

		FAFSAUUID: preprocessString2627(record[fafsaUUIDStartIndex2627-1 : (fafsaUUIDStartIndex2627-1)+fafsaUUIDLength2627]), // Field # 2

		TransactionUUID: preprocessString2627(record[transactionUUIDStartIndex2627-1 : (transactionUUIDStartIndex2627-1)+transactionUUIDLength2627]), // Field # 3

		PersonUUID: preprocessString2627(record[personUUIDStartIndex2627-1 : (personUUIDStartIndex2627-1)+personUUIDLength2627]), // Field # 4

		TransactionNumber: preprocessString2627(record[transactionNumberStartIndex2627-1 : (transactionNumberStartIndex2627-1)+transactionNumberLength2627]), // Field # 5

		DependencyModel: preprocessString2627(record[dependencyModelStartIndex2627-1 : (dependencyModelStartIndex2627-1)+dependencyModelLength2627]), // Field # 6

		ApplicationSource: preprocessString2627(record[applicationSourceStartIndex2627-1 : (applicationSourceStartIndex2627-1)+applicationSourceLength2627]), // Field # 7

		ApplicationReceiptDate: parseISIRDate2627(preprocessString2627(record[applicationReceiptDateStartIndex2627-1 : (applicationReceiptDateStartIndex2627-1)+applicationReceiptDateLength2627])), // Field # 8

		TransactionSource: preprocessString2627(record[transactionSourceStartIndex2627-1 : (transactionSourceStartIndex2627-1)+transactionSourceLength2627]), // Field # 9

		TransactionType: preprocessString2627(record[transactionTypeStartIndex2627-1 : (transactionTypeStartIndex2627-1)+transactionTypeLength2627]), // Field # 10

		TransactionLanguage: preprocessString2627(record[transactionLanguageStartIndex2627-1 : (transactionLanguageStartIndex2627-1)+transactionLanguageLength2627]), // Field # 11

		TransactionReceiptDate: parseISIRDate2627(preprocessString2627(record[transactionReceiptDateStartIndex2627-1 : (transactionReceiptDateStartIndex2627-1)+transactionReceiptDateLength2627])), // Field # 12

		TransactionProcessedDate: parseISIRDate2627(preprocessString2627(record[transactionProcessedDateStartIndex2627-1 : (transactionProcessedDateStartIndex2627-1)+transactionProcessedDateLength2627])), // Field # 13

		TransactionStatus: preprocessString2627(record[transactionStatusStartIndex2627-1 : (transactionStatusStartIndex2627-1)+transactionStatusLength2627]), // Field # 14

		RenewalDataUsed: preprocessString2627(record[renewalDataUsedStartIndex2627-1 : (renewalDataUsedStartIndex2627-1)+renewalDataUsedLength2627]), // Field # 15

		FPSCorrectionReason: preprocessString2627(record[fpsCorrectionReasonStartIndex2627-1 : (fpsCorrectionReasonStartIndex2627-1)+fpsCorrectionReasonLength2627]), // Field # 16

		SAIChangeFlag: preprocessString2627(record[saiChangeFlagStartIndex2627-1 : (saiChangeFlagStartIndex2627-1)+saiChangeFlagLength2627]), // Field # 17

		SAI: preprocessString2627(record[saiStartIndex2627-1 : (saiStartIndex2627-1)+saiLength2627]), // Field # 18

		ProvisionalSAI: preprocessString2627(record[provisionalSAIStartIndex2627-1 : (provisionalSAIStartIndex2627-1)+provisionalSAILength2627]), // Field # 19

		SAIFormula: preprocessString2627(record[saiFormulaStartIndex2627-1 : (saiFormulaStartIndex2627-1)+saiFormulaLength2627]), // Field # 20

		SAIComputationType: preprocessString2627(record[saiComputationTypeStartIndex2627-1 : (saiComputationTypeStartIndex2627-1)+saiComputationTypeLength2627]), // Field # 21

		MaxPellIndicator: preprocessString2627(record[maxPellIndicatorStartIndex2627-1 : (maxPellIndicatorStartIndex2627-1)+maxPellIndicatorLength2627]), // Field # 22

		MinimumPellIndicator: preprocessString2627(record[minimumPellIndicatorStartIndex2627-1 : (minimumPellIndicatorStartIndex2627-1)+minimumPellIndicatorLength2627]), // Field # 23

		StudentFirstName: preprocessString2627(record[studentFirstNameStartIndex2627-1 : (studentFirstNameStartIndex2627-1)+studentFirstNameLength2627]), // Field # 25

		StudentMiddleName: preprocessString2627(record[studentMiddleNameStartIndex2627-1 : (studentMiddleNameStartIndex2627-1)+studentMiddleNameLength2627]), // Field # 26

		StudentLastName: preprocessString2627(record[studentLastNameStartIndex2627-1 : (studentLastNameStartIndex2627-1)+studentLastNameLength2627]), // Field # 27

		StudentSuffix: preprocessString2627(record[studentSuffixStartIndex2627-1 : (studentSuffixStartIndex2627-1)+studentSuffixLength2627]), // Field # 28

		StudentDateOfBirth: parseISIRDate2627(preprocessString2627(record[studentDateOfBirthStartIndex2627-1 : (studentDateOfBirthStartIndex2627-1)+studentDateOfBirthLength2627])), // Field # 29

		StudentSSN: preprocessString2627(record[studentSSNStartIndex2627-1 : (studentSSNStartIndex2627-1)+studentSSNLength2627]), // Field # 30

		StudentITIN: preprocessString2627(record[studentITINStartIndex2627-1 : (studentITINStartIndex2627-1)+studentITINLength2627]), // Field # 31

		StudentPhoneNumber: preprocessString2627(record[studentPhoneNumberStartIndex2627-1 : (studentPhoneNumberStartIndex2627-1)+studentPhoneNumberLength2627]), // Field # 32

		StudentEmailAddress: preprocessString2627(record[studentEmailAddressStartIndex2627-1 : (studentEmailAddressStartIndex2627-1)+studentEmailAddressLength2627]), // Field # 33

		StudentStreetAddress: preprocessString2627(record[studentStreetAddressStartIndex2627-1 : (studentStreetAddressStartIndex2627-1)+studentStreetAddressLength2627]), // Field # 34

		StudentCity: preprocessString2627(record[studentCityStartIndex2627-1 : (studentCityStartIndex2627-1)+studentCityLength2627]), // Field # 35

		StudentState: preprocessString2627(record[studentStateStartIndex2627-1 : (studentStateStartIndex2627-1)+studentStateLength2627]), // Field # 36

		StudentZipCode: preprocessString2627(record[studentZipCodeStartIndex2627-1 : (studentZipCodeStartIndex2627-1)+studentZipCodeLength2627]), // Field # 37

		StudentCountry: preprocessString2627(record[studentCountryStartIndex2627-1 : (studentCountryStartIndex2627-1)+studentCountryLength2627]), // Field # 38

		StudentMaritalStatus: preprocessString2627(record[studentMaritalStatusStartIndex2627-1 : (studentMaritalStatusStartIndex2627-1)+studentMaritalStatusLength2627]), // Field # 40

		StudentGradeLevel: preprocessString2627(record[studentGradeLevelStartIndex2627-1 : (studentGradeLevelStartIndex2627-1)+studentGradeLevelLength2627]), // Field # 41

		StudentHasBachelorsDegree: preprocessString2627(record[studentFirstBachelorsDegreeBefore2627StartIndex2627-1 : (studentFirstBachelorsDegreeBefore2627StartIndex2627-1)+studentFirstBachelorsDegreeBefore2627Length2627]), // Field # 42

		StudentPursuingTeacherCertification: preprocessString2627(record[studentPursuingTeacherCertificationStartIndex2627-1 : (studentPursuingTeacherCertificationStartIndex2627-1)+studentPursuingTeacherCertificationLength2627]), // Field # 43

		StudentActiveDuty: preprocessString2627(record[studentActiveDutyStartIndex2627-1 : (studentActiveDutyStartIndex2627-1)+studentActiveDutyLength2627]), // Field # 44

		StudentVeteran: preprocessString2627(record[studentVeteranStartIndex2627-1 : (studentVeteranStartIndex2627-1)+studentVeteranLength2627]), // Field # 45

		StudentChildOrOtherDependents: preprocessString2627(record[studentChildOrOtherDependentsStartIndex2627-1 : (studentChildOrOtherDependentsStartIndex2627-1)+studentChildOrOtherDependentsLength2627]), // Field # 46

		StudentParentsDeceased: preprocessString2627(record[studentParentsDeceasedStartIndex2627-1 : (studentParentsDeceasedStartIndex2627-1)+studentParentsDeceasedLength2627]), // Field # 47

		StudentWardOfCourt: preprocessString2627(record[studentWardOfCourtStartIndex2627-1 : (studentWardOfCourtStartIndex2627-1)+studentWardOfCourtLength2627]), // Field # 48

		StudentInFosterCare: preprocessString2627(record[studentInFosterCareStartIndex2627-1 : (studentInFosterCareStartIndex2627-1)+studentInFosterCareLength2627]), // Field # 49

		StudentEmancipatedMinor: preprocessString2627(record[studentEmancipatedMinorStartIndex2627-1 : (studentEmancipatedMinorStartIndex2627-1)+studentEmancipatedMinorLength2627]), // Field # 50

		StudentLegalGuardianship: preprocessString2627(record[studentLegalGuardianshipStartIndex2627-1 : (studentLegalGuardianshipStartIndex2627-1)+studentLegalGuardianshipLength2627]), // Field # 51

		StudentPersonalCircumstancesNoneOfTheAbove: preprocessString2627(record[studentPersonalCircumstancesNoneOfTheAboveStartIndex2627-1 : (studentPersonalCircumstancesNoneOfTheAboveStartIndex2627-1)+studentPersonalCircumstancesNoneOfTheAboveLength2627]), // Field # 52

		StudentUnaccompaniedHomelessYouthAndSelfSupporting: preprocessString2627(record[studentUnaccompaniedHomelessYouthAndSelfSupportingStartIndex2627-1 : (studentUnaccompaniedHomelessYouthAndSelfSupportingStartIndex2627-1)+studentUnaccompaniedHomelessYouthAndSelfSupportingLength2627]), // Field # 53

		StudentUnaccompaniedHomelessGeneral: preprocessString2627(record[studentUnaccompaniedHomelessGeneralStartIndex2627-1 : (studentUnaccompaniedHomelessGeneralStartIndex2627-1)+studentUnaccompaniedHomelessGeneralLength2627]), // Field # 54

		StudentUnaccompaniedHomelessHS: preprocessString2627(record[studentUnaccompaniedHomelessHSStartIndex2627-1 : (studentUnaccompaniedHomelessHSStartIndex2627-1)+studentUnaccompaniedHomelessHSLength2627]), // Field # 55

		StudentUnaccompaniedHomelessTRIO: preprocessString2627(record[studentUnaccompaniedHomelessTRIOStartIndex2627-1 : (studentUnaccompaniedHomelessTRIOStartIndex2627-1)+studentUnaccompaniedHomelessTRIOLength2627]), // Field # 56

		StudentUnaccompaniedHomelessFAA: preprocessString2627(record[studentUnaccompaniedHomelessFAAStartIndex2627-1 : (studentUnaccompaniedHomelessFAAStartIndex2627-1)+studentUnaccompaniedHomelessFAALength2627]), // Field # 57

		StudentHomelessnessNoneOfTheAbove: preprocessString2627(record[studentHomelessnessNoneOfTheAboveStartIndex2627-1 : (studentHomelessnessNoneOfTheAboveStartIndex2627-1)+studentHomelessnessNoneOfTheAboveLength2627]), // Field # 58

		StudentUnusualCircumstance: preprocessString2627(record[studentUnusualCircumstanceStartIndex2627-1 : (studentUnusualCircumstanceStartIndex2627-1)+studentUnusualCircumstanceLength2627]), // Field # 59

		StudentUnsubOnly: preprocessString2627(record[studentUnsubOnlyStartIndex2627-1 : (studentUnsubOnlyStartIndex2627-1)+studentUnsubOnlyLength2627]), // Field # 60

		StudentUpdatedFamilySize: preprocessString2627(record[studentUpdatedFamilySizeStartIndex2627-1 : (studentUpdatedFamilySizeStartIndex2627-1)+studentUpdatedFamilySizeLength2627]), // Field # 61

		StudentNumberInCollege: preprocessString2627(record[studentNumberInCollegeStartIndex2627-1 : (studentNumberInCollegeStartIndex2627-1)+studentNumberInCollegeLength2627]), // Field # 62

		StudentCitizenshipStatus: preprocessString2627(record[studentCitizenshipStatusStartIndex2627-1 : (studentCitizenshipStatusStartIndex2627-1)+studentCitizenshipStatusLength2627]), // Field # 63

		StudentANumber: preprocessString2627(record[studentANumberStartIndex2627-1 : (studentANumberStartIndex2627-1)+studentANumberLength2627]), // Field # 64

		StudentStateOfLegalResidence: preprocessString2627(record[studentStateOfLegalResidenceStartIndex2627-1 : (studentStateOfLegalResidenceStartIndex2627-1)+studentStateOfLegalResidenceLength2627]), // Field # 65

		StudentLegalResidenceDate: parseISIRDateShort2627(preprocessString2627(record[studentLegalResidenceDateStartIndex2627-1 : (studentLegalResidenceDateStartIndex2627-1)+studentLegalResidenceDateLength2627])), // Field # 66

		StudentEitherParentAttendCollege: preprocessString2627(record[studentEitherParentAttendCollegeStartIndex2627-1 : (studentEitherParentAttendCollegeStartIndex2627-1)+studentEitherParentAttendCollegeLength2627]), // Field # 67

		StudentParentKilledInTheLineOfDuty: preprocessString2627(record[studentParentKilledInTheLineOfDutyStartIndex2627-1 : (studentParentKilledInTheLineOfDutyStartIndex2627-1)+studentParentKilledInTheLineOfDutyLength2627]), // Field # 68

		StudentHighSchoolCompletionStatus: preprocessString2627(record[studentHighSchoolCompletionStatusStartIndex2627-1 : (studentHighSchoolCompletionStatusStartIndex2627-1)+studentHighSchoolCompletionStatusLength2627]), // Field # 69

		StudentHighSchoolName: preprocessString2627(record[studentHighSchoolNameStartIndex2627-1 : (studentHighSchoolNameStartIndex2627-1)+studentHighSchoolNameLength2627]), // Field # 70

		StudentHighSchoolCity: preprocessString2627(record[studentHighSchoolCityStartIndex2627-1 : (studentHighSchoolCityStartIndex2627-1)+studentHighSchoolCityLength2627]), // Field # 71

		StudentHighSchoolState: preprocessString2627(record[studentHighSchoolStateStartIndex2627-1 : (studentHighSchoolStateStartIndex2627-1)+studentHighSchoolStateLength2627]), // Field # 72

		StudentHighSchoolEquivalentDiplomaName: preprocessString2627(record[studentHighSchoolEquivalentDiplomaNameStartIndex2627-1 : (studentHighSchoolEquivalentDiplomaNameStartIndex2627-1)+studentHighSchoolEquivalentDiplomaNameLength2627]), // Field # 73

		StudentHighSchoolEquivalentDiplomaState: preprocessString2627(record[studentHighSchoolEquivalentDiplomaStateStartIndex2627-1 : (studentHighSchoolEquivalentDiplomaStateStartIndex2627-1)+studentHighSchoolEquivalentDiplomaStateLength2627]), // Field # 74

		StudentManuallyEnteredReceivedEITC: preprocessString2627(record[studentManuallyEnteredReceivedEITCStartIndex2627-1 : (studentManuallyEnteredReceivedEITCStartIndex2627-1)+studentManuallyEnteredReceivedEITCLength2627]), // Field # 75

		StudentManuallyEnteredReceivedFederalHousingAssistance: preprocessString2627(record[studentManuallyEnteredReceivedFederalHousingAssistanceStartIndex2627-1 : (studentManuallyEnteredReceivedFederalHousingAssistanceStartIndex2627-1)+studentManuallyEnteredReceivedFederalHousingAssistanceLength2627]), // Field # 76

		StudentManuallyEnteredReceivedFreeReducedPriceLunch: preprocessString2627(record[studentManuallyEnteredReceivedFreeReducedPriceLunchStartIndex2627-1 : (studentManuallyEnteredReceivedFreeReducedPriceLunchStartIndex2627-1)+studentManuallyEnteredReceivedFreeReducedPriceLunchLength2627]), // Field # 77

		StudentManuallyEnteredReceivedMedicaid: preprocessString2627(record[studentManuallyEnteredReceivedMedicaidStartIndex2627-1 : (studentManuallyEnteredReceivedMedicaidStartIndex2627-1)+studentManuallyEnteredReceivedMedicaidLength2627]), // Field # 78

		StudentManuallyEnteredReceivedRefundableCreditFor36BHealthPlan: preprocessString2627(record[studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanStartIndex2627-1 : (studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanStartIndex2627-1)+studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanLength2627]), // Field # 79

		StudentManuallyEnteredReceivedSNAP: preprocessString2627(record[studentManuallyEnteredReceivedSNAPStartIndex2627-1 : (studentManuallyEnteredReceivedSNAPStartIndex2627-1)+studentManuallyEnteredReceivedSNAPLength2627]), // Field # 80

		StudentManuallyEnteredReceivedSupplementalSecurityIncome: preprocessString2627(record[studentManuallyEnteredReceivedSupplementalSecurityIncomeStartIndex2627-1 : (studentManuallyEnteredReceivedSupplementalSecurityIncomeStartIndex2627-1)+studentManuallyEnteredReceivedSupplementalSecurityIncomeLength2627]), // Field # 81

		StudentManuallyEnteredReceivedTANF: preprocessString2627(record[studentManuallyEnteredReceivedTANFStartIndex2627-1 : (studentManuallyEnteredReceivedTANFStartIndex2627-1)+studentManuallyEnteredReceivedTANFLength2627]), // Field # 82

		StudentManuallyEnteredReceivedWIC: preprocessString2627(record[studentManuallyEnteredReceivedWICStartIndex2627-1 : (studentManuallyEnteredReceivedWICStartIndex2627-1)+studentManuallyEnteredReceivedWICLength2627]), // Field # 83

		StudentManuallyEnteredFederalBenefitsNoneOfTheAbove: preprocessString2627(record[studentManuallyEnteredFederalBenefitsNoneOfTheAboveStartIndex2627-1 : (studentManuallyEnteredFederalBenefitsNoneOfTheAboveStartIndex2627-1)+studentManuallyEnteredFederalBenefitsNoneOfTheAboveLength2627]), // Field # 84

		StudentManuallyEnteredFiled1040Or1040NR: preprocessString2627(record[studentManuallyEnteredFiled1040Or1040NRStartIndex2627-1 : (studentManuallyEnteredFiled1040Or1040NRStartIndex2627-1)+studentManuallyEnteredFiled1040Or1040NRLength2627]), // Field # 85

		StudentManuallyEnteredFiledNonUSTaxReturn: preprocessString2627(record[studentManuallyEnteredFiledNonUSTaxReturnStartIndex2627-1 : (studentManuallyEnteredFiledNonUSTaxReturnStartIndex2627-1)+studentManuallyEnteredFiledNonUSTaxReturnLength2627]), // Field # 86

		StudentManuallyEnteredFiledJointReturnWithCurrentSpouse: preprocessString2627(record[studentManuallyEnteredFiledJointReturnWithCurrentSpouseStartIndex2627-1 : (studentManuallyEnteredFiledJointReturnWithCurrentSpouseStartIndex2627-1)+studentManuallyEnteredFiledJointReturnWithCurrentSpouseLength2627]), // Field # 87

		StudentManuallyEnteredTaxReturnFilingStatus: preprocessString2627(record[studentManuallyEnteredTaxReturnFilingStatusStartIndex2627-1 : (studentManuallyEnteredTaxReturnFilingStatusStartIndex2627-1)+studentManuallyEnteredTaxReturnFilingStatusLength2627]), // Field # 88

		StudentManuallyEnteredIncomeEarnedFromWork: preprocessString2627(record[studentManuallyEnteredIncomeEarnedFromWorkStartIndex2627-1 : (studentManuallyEnteredIncomeEarnedFromWorkStartIndex2627-1)+studentManuallyEnteredIncomeEarnedFromWorkLength2627]), // Field # 89

		StudentManuallyEnteredTaxExemptInterestIncome: preprocessString2627(record[studentManuallyEnteredTaxExemptInterestIncomeStartIndex2627-1 : (studentManuallyEnteredTaxExemptInterestIncomeStartIndex2627-1)+studentManuallyEnteredTaxExemptInterestIncomeLength2627]), // Field # 90

		StudentManuallyEnteredUntaxedPortionsOfIRADistributions: preprocessString2627(record[studentManuallyEnteredUntaxedPortionsOfIRADistributionsStartIndex2627-1 : (studentManuallyEnteredUntaxedPortionsOfIRADistributionsStartIndex2627-1)+studentManuallyEnteredUntaxedPortionsOfIRADistributionsLength2627]), // Field # 91

		StudentManuallyEnteredIRARollover: preprocessString2627(record[studentManuallyEnteredIRARolloverStartIndex2627-1 : (studentManuallyEnteredIRARolloverStartIndex2627-1)+studentManuallyEnteredIRARolloverLength2627]), // Field # 92

		StudentManuallyEnteredUntaxedPortionsOfPensions: preprocessString2627(record[studentManuallyEnteredUntaxedPortionsOfPensionsStartIndex2627-1 : (studentManuallyEnteredUntaxedPortionsOfPensionsStartIndex2627-1)+studentManuallyEnteredUntaxedPortionsOfPensionsLength2627]), // Field # 93

		StudentManuallyEnteredPensionRollover: preprocessString2627(record[studentManuallyEnteredPensionRolloverStartIndex2627-1 : (studentManuallyEnteredPensionRolloverStartIndex2627-1)+studentManuallyEnteredPensionRolloverLength2627]), // Field # 94

		StudentManuallyEnteredAdjustedGrossIncome: preprocessString2627(record[studentManuallyEnteredAdjustedGrossIncomeStartIndex2627-1 : (studentManuallyEnteredAdjustedGrossIncomeStartIndex2627-1)+studentManuallyEnteredAdjustedGrossIncomeLength2627]), // Field # 95

		StudentManuallyEnteredIncomeTaxPaid: preprocessString2627(record[studentManuallyEnteredIncomeTaxPaidStartIndex2627-1 : (studentManuallyEnteredIncomeTaxPaidStartIndex2627-1)+studentManuallyEnteredIncomeTaxPaidLength2627]), // Field # 96

		StudentManuallyEnteredEITCReceivedDuringTaxYear: preprocessString2627(record[studentManuallyEnteredEITCReceivedDuringTaxYearStartIndex2627-1 : (studentManuallyEnteredEITCReceivedDuringTaxYearStartIndex2627-1)+studentManuallyEnteredEITCReceivedDuringTaxYearLength2627]), // Field # 97

		StudentManuallyEnteredDeductiblePaymentsToIRAKeoghOther: preprocessString2627(record[studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherStartIndex2627-1 : (studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherStartIndex2627-1)+studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherLength2627]), // Field # 98

		StudentManuallyEnteredEducationCredits: preprocessString2627(record[studentManuallyEnteredEducationCreditsStartIndex2627-1 : (studentManuallyEnteredEducationCreditsStartIndex2627-1)+studentManuallyEnteredEducationCreditsLength2627]), // Field # 99

		StudentManuallyEnteredFiledScheduleABDEFH: preprocessString2627(record[studentManuallyEnteredFiledScheduleABDEFHStartIndex2627-1 : (studentManuallyEnteredFiledScheduleABDEFHStartIndex2627-1)+studentManuallyEnteredFiledScheduleABDEFHLength2627]), // Field # 100

		StudentManuallyEnteredScheduleCAmount: preprocessString2627(record[studentManuallyEnteredScheduleCAmountStartIndex2627-1 : (studentManuallyEnteredScheduleCAmountStartIndex2627-1)+studentManuallyEnteredScheduleCAmountLength2627]), // Field # 101

		StudentManuallyEnteredCollegeGrantAndScholarshipAid: preprocessString2627(record[studentManuallyEnteredCollegeGrantAndScholarshipAidStartIndex2627-1 : (studentManuallyEnteredCollegeGrantAndScholarshipAidStartIndex2627-1)+studentManuallyEnteredCollegeGrantAndScholarshipAidLength2627]), // Field # 102

		StudentManuallyEnteredForeignEarnedIncomeExclusion: preprocessString2627(record[studentManuallyEnteredForeignEarnedIncomeExclusionStartIndex2627-1 : (studentManuallyEnteredForeignEarnedIncomeExclusionStartIndex2627-1)+studentManuallyEnteredForeignEarnedIncomeExclusionLength2627]), // Field # 103

		StudentManuallyEnteredChildSupportReceived: preprocessString2627(record[studentManuallyEnteredChildSupportReceivedStartIndex2627-1 : (studentManuallyEnteredChildSupportReceivedStartIndex2627-1)+studentManuallyEnteredChildSupportReceivedLength2627]), // Field # 104

		StudentManuallyEnteredTotalOfCashSavingsAndCheckingAccounts: preprocessString2627(record[studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsStartIndex2627-1 : (studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsStartIndex2627-1)+studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsLength2627]), // Field # 105

		StudentManuallyEnteredNetWorthOfCurrentInvestments: preprocessString2627(record[studentManuallyEnteredNetWorthOfCurrentInvestmentsStartIndex2627-1 : (studentManuallyEnteredNetWorthOfCurrentInvestmentsStartIndex2627-1)+studentManuallyEnteredNetWorthOfCurrentInvestmentsLength2627]), // Field # 106

		StudentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarms: preprocessString2627(record[studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsStartIndex2627-1 : (studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsStartIndex2627-1)+studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsLength2627]), // Field # 107

		StudentCollege1: preprocessString2627(record[studentCollege1StartIndex2627-1 : (studentCollege1StartIndex2627-1)+studentCollege1Length2627]), // Field # 108

		StudentCollege2: preprocessString2627(record[studentCollege2StartIndex2627-1 : (studentCollege2StartIndex2627-1)+studentCollege2Length2627]), // Field # 109

		StudentCollege3: preprocessString2627(record[studentCollege3StartIndex2627-1 : (studentCollege3StartIndex2627-1)+studentCollege3Length2627]), // Field # 110

		StudentCollege4: preprocessString2627(record[studentCollege4StartIndex2627-1 : (studentCollege4StartIndex2627-1)+studentCollege4Length2627]), // Field # 111

		StudentCollege5: preprocessString2627(record[studentCollege5StartIndex2627-1 : (studentCollege5StartIndex2627-1)+studentCollege5Length2627]), // Field # 112

		StudentCollege6: preprocessString2627(record[studentCollege6StartIndex2627-1 : (studentCollege6StartIndex2627-1)+studentCollege6Length2627]), // Field # 113

		StudentCollege7: preprocessString2627(record[studentCollege7StartIndex2627-1 : (studentCollege7StartIndex2627-1)+studentCollege7Length2627]), // Field # 114

		StudentCollege8: preprocessString2627(record[studentCollege8StartIndex2627-1 : (studentCollege8StartIndex2627-1)+studentCollege8Length2627]), // Field # 115

		StudentCollege9: preprocessString2627(record[studentCollege9StartIndex2627-1 : (studentCollege9StartIndex2627-1)+studentCollege9Length2627]), // Field # 116

		StudentCollege10: preprocessString2627(record[studentCollege10StartIndex2627-1 : (studentCollege10StartIndex2627-1)+studentCollege10Length2627]), // Field # 117

		StudentCollege11: preprocessString2627(record[studentCollege11StartIndex2627-1 : (studentCollege11StartIndex2627-1)+studentCollege11Length2627]), // Field # 118

		StudentCollege12: preprocessString2627(record[studentCollege12StartIndex2627-1 : (studentCollege12StartIndex2627-1)+studentCollege12Length2627]), // Field # 119

		StudentCollege13: preprocessString2627(record[studentCollege13StartIndex2627-1 : (studentCollege13StartIndex2627-1)+studentCollege13Length2627]), // Field # 120

		StudentCollege14: preprocessString2627(record[studentCollege14StartIndex2627-1 : (studentCollege14StartIndex2627-1)+studentCollege14Length2627]), // Field # 121

		StudentCollege15: preprocessString2627(record[studentCollege15StartIndex2627-1 : (studentCollege15StartIndex2627-1)+studentCollege15Length2627]), // Field # 122

		StudentCollege16: preprocessString2627(record[studentCollege16StartIndex2627-1 : (studentCollege16StartIndex2627-1)+studentCollege16Length2627]), // Field # 123

		StudentCollege17: preprocessString2627(record[studentCollege17StartIndex2627-1 : (studentCollege17StartIndex2627-1)+studentCollege17Length2627]), // Field # 124

		StudentCollege18: preprocessString2627(record[studentCollege18StartIndex2627-1 : (studentCollege18StartIndex2627-1)+studentCollege18Length2627]), // Field # 125

		StudentCollege19: preprocessString2627(record[studentCollege19StartIndex2627-1 : (studentCollege19StartIndex2627-1)+studentCollege19Length2627]), // Field # 126

		StudentCollege20: preprocessString2627(record[studentCollege20StartIndex2627-1 : (studentCollege20StartIndex2627-1)+studentCollege20Length2627]), // Field # 127

		StudentConsentToRetrieveAndDiscloseFTI: preprocessString2627(record[studentConsentToRetrieveAndDiscloseFTIStartIndex2627-1 : (studentConsentToRetrieveAndDiscloseFTIStartIndex2627-1)+studentConsentToRetrieveAndDiscloseFTILength2627]), // Field # 128

		StudentSignature: preprocessString2627(record[studentSignatureStartIndex2627-1 : (studentSignatureStartIndex2627-1)+studentSignatureLength2627]), // Field # 129

		StudentSignatureDate: parseISIRDate2627(preprocessString2627(record[studentSignatureDateStartIndex2627-1 : (studentSignatureDateStartIndex2627-1)+studentSignatureDateLength2627])), // Field # 130

		StudentSpouseFirstName: preprocessString2627(record[studentSpouseFirstNameStartIndex2627-1 : (studentSpouseFirstNameStartIndex2627-1)+studentSpouseFirstNameLength2627]), // Field # 132

		StudentSpouseMiddleName: preprocessString2627(record[studentSpouseMiddleNameStartIndex2627-1 : (studentSpouseMiddleNameStartIndex2627-1)+studentSpouseMiddleNameLength2627]), // Field # 133

		StudentSpouseLastName: preprocessString2627(record[studentSpouseLastNameStartIndex2627-1 : (studentSpouseLastNameStartIndex2627-1)+studentSpouseLastNameLength2627]), // Field # 134

		StudentSpouseSuffix: preprocessString2627(record[studentSpouseSuffixStartIndex2627-1 : (studentSpouseSuffixStartIndex2627-1)+studentSpouseSuffixLength2627]), // Field # 135

		StudentSpouseDateOfBirth: parseISIRDate2627(preprocessString2627(record[studentSpouseDateOfBirthStartIndex2627-1 : (studentSpouseDateOfBirthStartIndex2627-1)+studentSpouseDateOfBirthLength2627])), // Field # 136

		StudentSpouseSSN: preprocessString2627(record[studentSpouseSSNStartIndex2627-1 : (studentSpouseSSNStartIndex2627-1)+studentSpouseSSNLength2627]), // Field # 137

		StudentSpouseITIN: preprocessString2627(record[studentSpouseITINStartIndex2627-1 : (studentSpouseITINStartIndex2627-1)+studentSpouseITINLength2627]), // Field # 138

		StudentSpousePhoneNumber: preprocessString2627(record[studentSpousePhoneNumberStartIndex2627-1 : (studentSpousePhoneNumberStartIndex2627-1)+studentSpousePhoneNumberLength2627]), // Field # 139

		StudentSpouseEmailAddress: preprocessString2627(record[studentSpouseEmailAddressStartIndex2627-1 : (studentSpouseEmailAddressStartIndex2627-1)+studentSpouseEmailAddressLength2627]), // Field # 140

		StudentSpouseStreetAddress: preprocessString2627(record[studentSpouseStreetAddressStartIndex2627-1 : (studentSpouseStreetAddressStartIndex2627-1)+studentSpouseStreetAddressLength2627]), // Field # 141

		StudentSpouseCity: preprocessString2627(record[studentSpouseCityStartIndex2627-1 : (studentSpouseCityStartIndex2627-1)+studentSpouseCityLength2627]), // Field # 142

		StudentSpouseState: preprocessString2627(record[studentSpouseStateStartIndex2627-1 : (studentSpouseStateStartIndex2627-1)+studentSpouseStateLength2627]), // Field # 143

		StudentSpouseZipCode: preprocessString2627(record[studentSpouseZipCodeStartIndex2627-1 : (studentSpouseZipCodeStartIndex2627-1)+studentSpouseZipCodeLength2627]), // Field # 144

		StudentSpouseCountry: preprocessString2627(record[studentSpouseCountryStartIndex2627-1 : (studentSpouseCountryStartIndex2627-1)+studentSpouseCountryLength2627]), // Field # 145

		StudentSpouseFiled1040Or1040NR: preprocessString2627(record[studentSpouseFiled1040Or1040NRStartIndex2627-1 : (studentSpouseFiled1040Or1040NRStartIndex2627-1)+studentSpouseFiled1040Or1040NRLength2627]), // Field # 146

		StudentSpouseFiledNonUSTaxReturn: preprocessString2627(record[studentSpouseFiledNonUSTaxReturnStartIndex2627-1 : (studentSpouseFiledNonUSTaxReturnStartIndex2627-1)+studentSpouseFiledNonUSTaxReturnLength2627]), // Field # 147

		StudentSpouseTaxReturnFilingStatus: preprocessString2627(record[studentSpouseTaxReturnFilingStatusStartIndex2627-1 : (studentSpouseTaxReturnFilingStatusStartIndex2627-1)+studentSpouseTaxReturnFilingStatusLength2627]), // Field # 148

		StudentSpouseIncomeEarnedFromWork: preprocessString2627(record[studentSpouseIncomeEarnedFromWorkStartIndex2627-1 : (studentSpouseIncomeEarnedFromWorkStartIndex2627-1)+studentSpouseIncomeEarnedFromWorkLength2627]), // Field # 149

		StudentSpouseTaxExemptInterestIncome: preprocessString2627(record[studentSpouseTaxExemptInterestIncomeStartIndex2627-1 : (studentSpouseTaxExemptInterestIncomeStartIndex2627-1)+studentSpouseTaxExemptInterestIncomeLength2627]), // Field # 150

		StudentSpouseUntaxedPortionsOfIRADistributions: preprocessString2627(record[studentSpouseUntaxedPortionsOfIRADistributionsStartIndex2627-1 : (studentSpouseUntaxedPortionsOfIRADistributionsStartIndex2627-1)+studentSpouseUntaxedPortionsOfIRADistributionsLength2627]), // Field # 151

		StudentSpouseIRARollover: preprocessString2627(record[studentSpouseIRARolloverStartIndex2627-1 : (studentSpouseIRARolloverStartIndex2627-1)+studentSpouseIRARolloverLength2627]), // Field # 152

		StudentSpouseUntaxedPortionsOfPensions: preprocessString2627(record[studentSpouseUntaxedPortionsOfPensionsStartIndex2627-1 : (studentSpouseUntaxedPortionsOfPensionsStartIndex2627-1)+studentSpouseUntaxedPortionsOfPensionsLength2627]), // Field # 153

		StudentSpousePensionRollover: preprocessString2627(record[studentSpousePensionRolloverStartIndex2627-1 : (studentSpousePensionRolloverStartIndex2627-1)+studentSpousePensionRolloverLength2627]), // Field # 154

		StudentSpouseAdjustedGrossIncome: preprocessString2627(record[studentSpouseAdjustedGrossIncomeStartIndex2627-1 : (studentSpouseAdjustedGrossIncomeStartIndex2627-1)+studentSpouseAdjustedGrossIncomeLength2627]), // Field # 155

		StudentSpouseIncomeTaxPaid: preprocessString2627(record[studentSpouseIncomeTaxPaidStartIndex2627-1 : (studentSpouseIncomeTaxPaidStartIndex2627-1)+studentSpouseIncomeTaxPaidLength2627]), // Field # 156

		StudentSpouseDeductiblePaymentsToIRAKeoghOther: preprocessString2627(record[studentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex2627-1 : (studentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex2627-1)+studentSpouseDeductiblePaymentsToIRAKeoghOtherLength2627]), // Field # 157

		StudentSpouseEducationCredits: preprocessString2627(record[studentSpouseEducationCreditsStartIndex2627-1 : (studentSpouseEducationCreditsStartIndex2627-1)+studentSpouseEducationCreditsLength2627]), // Field # 158

		StudentSpouseFiledScheduleABDEFH: preprocessString2627(record[studentSpouseFiledScheduleABDEFHStartIndex2627-1 : (studentSpouseFiledScheduleABDEFHStartIndex2627-1)+studentSpouseFiledScheduleABDEFHLength2627]), // Field # 159

		StudentSpouseScheduleCAmount: preprocessString2627(record[studentSpouseScheduleCAmountStartIndex2627-1 : (studentSpouseScheduleCAmountStartIndex2627-1)+studentSpouseScheduleCAmountLength2627]), // Field # 160

		StudentSpouseForeignEarnedIncomeExclusion: preprocessString2627(record[studentSpouseForeignEarnedIncomeExclusionStartIndex2627-1 : (studentSpouseForeignEarnedIncomeExclusionStartIndex2627-1)+studentSpouseForeignEarnedIncomeExclusionLength2627]), // Field # 161

		StudentSpouseConsentToRetrieveAndDiscloseFTI: preprocessString2627(record[studentSpouseConsentToRetrieveAndDiscloseFTIStartIndex2627-1 : (studentSpouseConsentToRetrieveAndDiscloseFTIStartIndex2627-1)+studentSpouseConsentToRetrieveAndDiscloseFTILength2627]), // Field # 162

		StudentSpouseSignature: preprocessString2627(record[studentSpouseSignatureStartIndex2627-1 : (studentSpouseSignatureStartIndex2627-1)+studentSpouseSignatureLength2627]), // Field # 163

		StudentSpouseSignatureDate: parseISIRDate2627(preprocessString2627(record[studentSpouseSignatureDateStartIndex2627-1 : (studentSpouseSignatureDateStartIndex2627-1)+studentSpouseSignatureDateLength2627])), // Field # 164

		ParentFirstName: preprocessString2627(record[parentFirstNameStartIndex2627-1 : (parentFirstNameStartIndex2627-1)+parentFirstNameLength2627]), // Field # 166

		ParentMiddleName: preprocessString2627(record[parentMiddleNameStartIndex2627-1 : (parentMiddleNameStartIndex2627-1)+parentMiddleNameLength2627]), // Field # 167

		ParentLastName: preprocessString2627(record[parentLastNameStartIndex2627-1 : (parentLastNameStartIndex2627-1)+parentLastNameLength2627]), // Field # 168

		ParentSuffix: preprocessString2627(record[parentSuffixStartIndex2627-1 : (parentSuffixStartIndex2627-1)+parentSuffixLength2627]), // Field # 169

		ParentDateOfBirth: parseISIRDate2627(preprocessString2627(record[parentDateOfBirthStartIndex2627-1 : (parentDateOfBirthStartIndex2627-1)+parentDateOfBirthLength2627])), // Field # 170

		ParentSSN: preprocessString2627(record[parentSSNStartIndex2627-1 : (parentSSNStartIndex2627-1)+parentSSNLength2627]), // Field # 171

		ParentITIN: preprocessString2627(record[parentITINStartIndex2627-1 : (parentITINStartIndex2627-1)+parentITINLength2627]), // Field # 172

		ParentPhoneNumber: preprocessString2627(record[parentPhoneNumberStartIndex2627-1 : (parentPhoneNumberStartIndex2627-1)+parentPhoneNumberLength2627]), // Field # 173

		ParentEmailAddress: preprocessString2627(record[parentEmailAddressStartIndex2627-1 : (parentEmailAddressStartIndex2627-1)+parentEmailAddressLength2627]), // Field # 174

		ParentStreetAddress: preprocessString2627(record[parentStreetAddressStartIndex2627-1 : (parentStreetAddressStartIndex2627-1)+parentStreetAddressLength2627]), // Field # 175

		ParentCity: preprocessString2627(record[parentCityStartIndex2627-1 : (parentCityStartIndex2627-1)+parentCityLength2627]), // Field # 176

		ParentState: preprocessString2627(record[parentStateStartIndex2627-1 : (parentStateStartIndex2627-1)+parentStateLength2627]), // Field # 177

		ParentZipCode: preprocessString2627(record[parentZipCodeStartIndex2627-1 : (parentZipCodeStartIndex2627-1)+parentZipCodeLength2627]), // Field # 178

		ParentCountry: preprocessString2627(record[parentCountryStartIndex2627-1 : (parentCountryStartIndex2627-1)+parentCountryLength2627]), // Field # 179

		ParentMaritalStatus: preprocessString2627(record[parentMaritalStatusStartIndex2627-1 : (parentMaritalStatusStartIndex2627-1)+parentMaritalStatusLength2627]), // Field # 180

		ParentStateOfLegalResidence: preprocessString2627(record[parentStateOfLegalResidenceStartIndex2627-1 : (parentStateOfLegalResidenceStartIndex2627-1)+parentStateOfLegalResidenceLength2627]), // Field # 181

		ParentLegalResidenceDate: parseISIRDateShort2627(preprocessString2627(record[parentLegalResidenceDateStartIndex2627-1 : (parentLegalResidenceDateStartIndex2627-1)+parentLegalResidenceDateLength2627])), // Field # 182

		ParentUpdatedFamilySize: preprocessString2627(record[parentUpdatedFamilySizeStartIndex2627-1 : (parentUpdatedFamilySizeStartIndex2627-1)+parentUpdatedFamilySizeLength2627]), // Field # 183

		ParentNumberInCollege: preprocessString2627(record[parentNumberInCollegeStartIndex2627-1 : (parentNumberInCollegeStartIndex2627-1)+parentNumberInCollegeLength2627]), // Field # 184

		ParentReceivedEITC: preprocessString2627(record[parentReceivedEITCStartIndex2627-1 : (parentReceivedEITCStartIndex2627-1)+parentReceivedEITCLength2627]), // Field # 185

		ParentReceivedFederalHousingAssistance: preprocessString2627(record[parentReceivedFederalHousingAssistanceStartIndex2627-1 : (parentReceivedFederalHousingAssistanceStartIndex2627-1)+parentReceivedFederalHousingAssistanceLength2627]), // Field # 186

		ParentReceivedFreeReducedPriceLunch: preprocessString2627(record[parentReceivedFreeReducedPriceLunchStartIndex2627-1 : (parentReceivedFreeReducedPriceLunchStartIndex2627-1)+parentReceivedFreeReducedPriceLunchLength2627]), // Field # 187

		ParentReceivedMedicaid: preprocessString2627(record[parentReceivedMedicaidStartIndex2627-1 : (parentReceivedMedicaidStartIndex2627-1)+parentReceivedMedicaidLength2627]), // Field # 188

		ParentReceivedRefundableCreditFor36BHealthPlan: preprocessString2627(record[parentReceivedRefundableCreditFor36BHealthPlanStartIndex2627-1 : (parentReceivedRefundableCreditFor36BHealthPlanStartIndex2627-1)+parentReceivedRefundableCreditFor36BHealthPlanLength2627]), // Field # 189

		ParentReceivedSNAP: preprocessString2627(record[parentReceivedSNAPStartIndex2627-1 : (parentReceivedSNAPStartIndex2627-1)+parentReceivedSNAPLength2627]), // Field # 190

		ParentReceivedSupplementalSecurityIncome: preprocessString2627(record[parentReceivedSupplementalSecurityIncomeStartIndex2627-1 : (parentReceivedSupplementalSecurityIncomeStartIndex2627-1)+parentReceivedSupplementalSecurityIncomeLength2627]), // Field # 191

		ParentReceivedTANF: preprocessString2627(record[parentReceivedTANFStartIndex2627-1 : (parentReceivedTANFStartIndex2627-1)+parentReceivedTANFLength2627]), // Field # 192

		ParentReceivedWIC: preprocessString2627(record[parentReceivedWICStartIndex2627-1 : (parentReceivedWICStartIndex2627-1)+parentReceivedWICLength2627]), // Field # 193

		ParentFederalBenefitsNoneOfTheAbove: preprocessString2627(record[parentFederalBenefitsNoneOfTheAboveStartIndex2627-1 : (parentFederalBenefitsNoneOfTheAboveStartIndex2627-1)+parentFederalBenefitsNoneOfTheAboveLength2627]), // Field # 194

		ParentFiled1040Or1040NR: preprocessString2627(record[parentFiled1040Or1040NRStartIndex2627-1 : (parentFiled1040Or1040NRStartIndex2627-1)+parentFiled1040Or1040NRLength2627]), // Field # 195

		ParentFileNonUSTaxReturn: preprocessString2627(record[parentFileNonUSTaxReturnStartIndex2627-1 : (parentFileNonUSTaxReturnStartIndex2627-1)+parentFileNonUSTaxReturnLength2627]), // Field # 196

		ParentFiledJointReturnWithCurrentSpouse: preprocessString2627(record[parentFiledJointReturnWithCurrentSpouseStartIndex2627-1 : (parentFiledJointReturnWithCurrentSpouseStartIndex2627-1)+parentFiledJointReturnWithCurrentSpouseLength2627]), // Field # 197

		ParentTaxReturnFilingStatus: preprocessString2627(record[parentTaxReturnFilingStatusStartIndex2627-1 : (parentTaxReturnFilingStatusStartIndex2627-1)+parentTaxReturnFilingStatusLength2627]), // Field # 198

		ParentIncomeEarnedFromWork: preprocessString2627(record[parentIncomeEarnedFromWorkStartIndex2627-1 : (parentIncomeEarnedFromWorkStartIndex2627-1)+parentIncomeEarnedFromWorkLength2627]), // Field # 199

		ParentTaxExemptInterestIncome: preprocessString2627(record[parentTaxExemptInterestIncomeStartIndex2627-1 : (parentTaxExemptInterestIncomeStartIndex2627-1)+parentTaxExemptInterestIncomeLength2627]), // Field # 200

		ParentUntaxedPortionsOfIRADistributions: preprocessString2627(record[parentUntaxedPortionsOfIRADistributionsStartIndex2627-1 : (parentUntaxedPortionsOfIRADistributionsStartIndex2627-1)+parentUntaxedPortionsOfIRADistributionsLength2627]), // Field # 201

		ParentIRARollover: preprocessString2627(record[parentIRARolloverStartIndex2627-1 : (parentIRARolloverStartIndex2627-1)+parentIRARolloverLength2627]), // Field # 202

		ParentUntaxedPortionsOfPensions: preprocessString2627(record[parentUntaxedPortionsOfPensionsStartIndex2627-1 : (parentUntaxedPortionsOfPensionsStartIndex2627-1)+parentUntaxedPortionsOfPensionsLength2627]), // Field # 203

		ParentPensionRollover: preprocessString2627(record[parentPensionRolloverStartIndex2627-1 : (parentPensionRolloverStartIndex2627-1)+parentPensionRolloverLength2627]), // Field # 204

		ParentAdjustedGrossIncome: preprocessString2627(record[parentAdjustedGrossIncomeStartIndex2627-1 : (parentAdjustedGrossIncomeStartIndex2627-1)+parentAdjustedGrossIncomeLength2627]), // Field # 205

		ParentIncomeTaxPaid: preprocessString2627(record[parentIncomeTaxPaidStartIndex2627-1 : (parentIncomeTaxPaidStartIndex2627-1)+parentIncomeTaxPaidLength2627]), // Field # 206

		ParentEarnedIncomeTaxCreditReceivedDuringTaxYear: preprocessString2627(record[parentEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex2627-1 : (parentEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex2627-1)+parentEarnedIncomeTaxCreditReceivedDuringTaxYearLength2627]), // Field # 207

		ParentDeductiblePaymentsToIRAKeoghOther: preprocessString2627(record[parentDeductiblePaymentsToIRAKeoghOtherStartIndex2627-1 : (parentDeductiblePaymentsToIRAKeoghOtherStartIndex2627-1)+parentDeductiblePaymentsToIRAKeoghOtherLength2627]), // Field # 208

		ParentEducationCredits: preprocessString2627(record[parentEducationCreditsStartIndex2627-1 : (parentEducationCreditsStartIndex2627-1)+parentEducationCreditsLength2627]), // Field # 209

		ParentFiledScheduleABDEFH: preprocessString2627(record[parentFiledScheduleABDEFHStartIndex2627-1 : (parentFiledScheduleABDEFHStartIndex2627-1)+parentFiledScheduleABDEFHLength2627]), // Field # 210

		ParentScheduleCAmount: preprocessString2627(record[parentScheduleCAmountStartIndex2627-1 : (parentScheduleCAmountStartIndex2627-1)+parentScheduleCAmountLength2627]), // Field # 211

		ParentCollegeGrantAndScholarshipAid: preprocessString2627(record[parentCollegeGrantAndScholarshipAidStartIndex2627-1 : (parentCollegeGrantAndScholarshipAidStartIndex2627-1)+parentCollegeGrantAndScholarshipAidLength2627]), // Field # 212

		ParentForeignEarnedIncomeExclusion: preprocessString2627(record[parentForeignEarnedIncomeExclusionStartIndex2627-1 : (parentForeignEarnedIncomeExclusionStartIndex2627-1)+parentForeignEarnedIncomeExclusionLength2627]), // Field # 213

		ParentChildSupportReceived: preprocessString2627(record[parentChildSupportReceivedStartIndex2627-1 : (parentChildSupportReceivedStartIndex2627-1)+parentChildSupportReceivedLength2627]), // Field # 214

		ParentTotalOfCashSavingsAndCheckingAccounts: preprocessString2627(record[parentTotalOfCashSavingsAndCheckingAccountsStartIndex2627-1 : (parentTotalOfCashSavingsAndCheckingAccountsStartIndex2627-1)+parentTotalOfCashSavingsAndCheckingAccountsLength2627]), // Field # 215

		ParentNetWorthOfCurrentInvestments: preprocessString2627(record[parentNetWorthOfCurrentInvestmentsStartIndex2627-1 : (parentNetWorthOfCurrentInvestmentsStartIndex2627-1)+parentNetWorthOfCurrentInvestmentsLength2627]), // Field # 216

		ParentNetWorthOfBusinessesAndInvestmentFarms: preprocessString2627(record[parentNetWorthOfBusinessesAndInvestmentFarmsStartIndex2627-1 : (parentNetWorthOfBusinessesAndInvestmentFarmsStartIndex2627-1)+parentNetWorthOfBusinessesAndInvestmentFarmsLength2627]), // Field # 217

		ParentConsentToRetrieveAndDiscloseFTI: preprocessString2627(record[parentConsentToRetrieveAndDiscloseFTIStartIndex2627-1 : (parentConsentToRetrieveAndDiscloseFTIStartIndex2627-1)+parentConsentToRetrieveAndDiscloseFTILength2627]), // Field # 218

		ParentSignature: preprocessString2627(record[parentSignatureStartIndex2627-1 : (parentSignatureStartIndex2627-1)+parentSignatureLength2627]), // Field # 219

		ParentSignatureDate: parseISIRDate2627(preprocessString2627(record[parentSignatureDateStartIndex2627-1 : (parentSignatureDateStartIndex2627-1)+parentSignatureDateLength2627])), // Field # 220

		ParentSpouseFirstName: preprocessString2627(record[parentSpouseFirstNameStartIndex2627-1 : (parentSpouseFirstNameStartIndex2627-1)+parentSpouseFirstNameLength2627]), // Field # 222

		ParentSpouseMiddleName: preprocessString2627(record[parentSpouseMiddleNameStartIndex2627-1 : (parentSpouseMiddleNameStartIndex2627-1)+parentSpouseMiddleNameLength2627]), // Field # 223

		ParentSpouseLastName: preprocessString2627(record[parentSpouseLastNameStartIndex2627-1 : (parentSpouseLastNameStartIndex2627-1)+parentSpouseLastNameLength2627]), // Field # 224

		ParentSpouseSuffix: preprocessString2627(record[parentSpouseSuffixStartIndex2627-1 : (parentSpouseSuffixStartIndex2627-1)+parentSpouseSuffixLength2627]), // Field # 225

		ParentSpouseDateOfBirth: parseISIRDate2627(preprocessString2627(record[parentSpouseDateOfBirthStartIndex2627-1 : (parentSpouseDateOfBirthStartIndex2627-1)+parentSpouseDateOfBirthLength2627])), // Field # 226

		ParentSpouseSSN: preprocessString2627(record[parentSpouseSSNStartIndex2627-1 : (parentSpouseSSNStartIndex2627-1)+parentSpouseSSNLength2627]), // Field # 227

		ParentSpouseITIN: preprocessString2627(record[parentSpouseITINStartIndex2627-1 : (parentSpouseITINStartIndex2627-1)+parentSpouseITINLength2627]), // Field # 228

		ParentSpousePhoneNumber: preprocessString2627(record[parentSpousePhoneNumberStartIndex2627-1 : (parentSpousePhoneNumberStartIndex2627-1)+parentSpousePhoneNumberLength2627]), // Field # 229

		ParentSpouseEmailAddress: preprocessString2627(record[parentSpouseEmailAddressStartIndex2627-1 : (parentSpouseEmailAddressStartIndex2627-1)+parentSpouseEmailAddressLength2627]), // Field # 230

		ParentSpouseStreetAddress: preprocessString2627(record[parentSpouseStreetAddressStartIndex2627-1 : (parentSpouseStreetAddressStartIndex2627-1)+parentSpouseStreetAddressLength2627]), // Field # 231

		ParentSpouseCity: preprocessString2627(record[parentSpouseCityStartIndex2627-1 : (parentSpouseCityStartIndex2627-1)+parentSpouseCityLength2627]), // Field # 232

		ParentSpouseState: preprocessString2627(record[parentSpouseStateStartIndex2627-1 : (parentSpouseStateStartIndex2627-1)+parentSpouseStateLength2627]), // Field # 233

		ParentSpouseZipCode: preprocessString2627(record[parentSpouseZipCodeStartIndex2627-1 : (parentSpouseZipCodeStartIndex2627-1)+parentSpouseZipCodeLength2627]), // Field # 234

		ParentSpouseCountry: preprocessString2627(record[parentSpouseCountryStartIndex2627-1 : (parentSpouseCountryStartIndex2627-1)+parentSpouseCountryLength2627]), // Field # 235

		ParentSpouseFiled1040Or1040NR: preprocessString2627(record[parentSpouseFiled1040Or1040NRStartIndex2627-1 : (parentSpouseFiled1040Or1040NRStartIndex2627-1)+parentSpouseFiled1040Or1040NRLength2627]), // Field # 236

		ParentSpouseFileNonUSTaxReturn: preprocessString2627(record[parentSpouseFileNonUSTaxReturnStartIndex2627-1 : (parentSpouseFileNonUSTaxReturnStartIndex2627-1)+parentSpouseFileNonUSTaxReturnLength2627]), // Field # 237

		ParentSpouseTaxReturnFilingStatus: preprocessString2627(record[parentSpouseTaxReturnFilingStatusStartIndex2627-1 : (parentSpouseTaxReturnFilingStatusStartIndex2627-1)+parentSpouseTaxReturnFilingStatusLength2627]), // Field # 238

		ParentSpouseIncomeEarnedFromWork: preprocessString2627(record[parentSpouseIncomeEarnedFromWorkStartIndex2627-1 : (parentSpouseIncomeEarnedFromWorkStartIndex2627-1)+parentSpouseIncomeEarnedFromWorkLength2627]), // Field # 239

		ParentSpouseTaxExemptInterestIncome: preprocessString2627(record[parentSpouseTaxExemptInterestIncomeStartIndex2627-1 : (parentSpouseTaxExemptInterestIncomeStartIndex2627-1)+parentSpouseTaxExemptInterestIncomeLength2627]), // Field # 240

		ParentSpouseUntaxedPortionsOfIRADistributions: preprocessString2627(record[parentSpouseUntaxedPortionsOfIRADistributionsStartIndex2627-1 : (parentSpouseUntaxedPortionsOfIRADistributionsStartIndex2627-1)+parentSpouseUntaxedPortionsOfIRADistributionsLength2627]), // Field # 241

		ParentSpouseIRARollover: preprocessString2627(record[parentSpouseIRARolloverStartIndex2627-1 : (parentSpouseIRARolloverStartIndex2627-1)+parentSpouseIRARolloverLength2627]), // Field # 242

		ParentSpouseUntaxedPortionsOfPensions: preprocessString2627(record[parentSpouseUntaxedPortionsOfPensionsStartIndex2627-1 : (parentSpouseUntaxedPortionsOfPensionsStartIndex2627-1)+parentSpouseUntaxedPortionsOfPensionsLength2627]), // Field # 243

		ParentSpousePensionRollover: preprocessString2627(record[parentSpousePensionRolloverStartIndex2627-1 : (parentSpousePensionRolloverStartIndex2627-1)+parentSpousePensionRolloverLength2627]), // Field # 244

		ParentSpouseAdjustedGrossIncome: preprocessString2627(record[parentSpouseAdjustedGrossIncomeStartIndex2627-1 : (parentSpouseAdjustedGrossIncomeStartIndex2627-1)+parentSpouseAdjustedGrossIncomeLength2627]), // Field # 245

		ParentSpouseIncomeTaxPaid: preprocessString2627(record[parentSpouseIncomeTaxPaidStartIndex2627-1 : (parentSpouseIncomeTaxPaidStartIndex2627-1)+parentSpouseIncomeTaxPaidLength2627]), // Field # 246

		ParentSpouseDeductiblePaymentsToIRAKeoghOther: preprocessString2627(record[parentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex2627-1 : (parentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex2627-1)+parentSpouseDeductiblePaymentsToIRAKeoghOtherLength2627]), // Field # 247

		ParentSpouseEducationCredits: preprocessString2627(record[parentSpouseEducationCreditsStartIndex2627-1 : (parentSpouseEducationCreditsStartIndex2627-1)+parentSpouseEducationCreditsLength2627]), // Field # 248

		ParentSpouseFiledScheduleABDEFH: preprocessString2627(record[parentSpouseFiledScheduleABDEFHStartIndex2627-1 : (parentSpouseFiledScheduleABDEFHStartIndex2627-1)+parentSpouseFiledScheduleABDEFHLength2627]), // Field # 249

		ParentSpouseScheduleCAmount: preprocessString2627(record[parentSpouseScheduleCAmountStartIndex2627-1 : (parentSpouseScheduleCAmountStartIndex2627-1)+parentSpouseScheduleCAmountLength2627]), // Field # 250

		ParentSpouseForeignEarnedIncomeExclusion: preprocessString2627(record[parentSpouseForeignEarnedIncomeExclusionStartIndex2627-1 : (parentSpouseForeignEarnedIncomeExclusionStartIndex2627-1)+parentSpouseForeignEarnedIncomeExclusionLength2627]), // Field # 251

		ParentSpouseConsentToRetrieveAndDiscloseFTI: preprocessString2627(record[parentSpouseConsentToRetrieveAndDiscloseFTIStartIndex2627-1 : (parentSpouseConsentToRetrieveAndDiscloseFTIStartIndex2627-1)+parentSpouseConsentToRetrieveAndDiscloseFTILength2627]), // Field # 252

		ParentSpouseSignature: preprocessString2627(record[parentSpouseSignatureStartIndex2627-1 : (parentSpouseSignatureStartIndex2627-1)+parentSpouseSignatureLength2627]), // Field # 253

		ParentSpouseSignatureDate: parseISIRDate2627(preprocessString2627(record[parentSpouseSignatureDateStartIndex2627-1 : (parentSpouseSignatureDateStartIndex2627-1)+parentSpouseSignatureDateLength2627])), // Field # 254

		PreparerFirstName: preprocessString2627(record[preparerFirstNameStartIndex2627-1 : (preparerFirstNameStartIndex2627-1)+preparerFirstNameLength2627]), // Field # 256

		PreparerLastName: preprocessString2627(record[preparerLastNameStartIndex2627-1 : (preparerLastNameStartIndex2627-1)+preparerLastNameLength2627]), // Field # 257

		PreparerSSN: preprocessString2627(record[preparerSSNStartIndex2627-1 : (preparerSSNStartIndex2627-1)+preparerSSNLength2627]), // Field # 258

		PreparerEIN: preprocessString2627(record[preparerEINStartIndex2627-1 : (preparerEINStartIndex2627-1)+preparerEINLength2627]), // Field # 259

		PreparerAffiliation: preprocessString2627(record[preparerAffiliationStartIndex2627-1 : (preparerAffiliationStartIndex2627-1)+preparerAffiliationLength2627]), // Field # 260

		PreparerStreetAddress: preprocessString2627(record[preparerStreetAddressStartIndex2627-1 : (preparerStreetAddressStartIndex2627-1)+preparerStreetAddressLength2627]), // Field # 261

		PreparerCity: preprocessString2627(record[preparerCityStartIndex2627-1 : (preparerCityStartIndex2627-1)+preparerCityLength2627]), // Field # 262

		PreparerState: preprocessString2627(record[preparerStateStartIndex2627-1 : (preparerStateStartIndex2627-1)+preparerStateLength2627]), // Field # 263

		PreparerZipCode: preprocessString2627(record[preparerZipCodeStartIndex2627-1 : (preparerZipCodeStartIndex2627-1)+preparerZipCodeLength2627]), // Field # 264

		PreparerSignature: preprocessString2627(record[preparerSignatureStartIndex2627-1 : (preparerSignatureStartIndex2627-1)+preparerSignatureLength2627]), // Field # 265

		PreparerSignatureDate: parseISIRDate2627(preprocessString2627(record[preparerSignatureDateStartIndex2627-1 : (preparerSignatureDateStartIndex2627-1)+preparerSignatureDateLength2627])), // Field # 266

		StudentAffirmationStatus: preprocessString2627(record[studentAffirmationStatusStartIndex2627-1 : (studentAffirmationStatusStartIndex2627-1)+studentAffirmationStatusLength2627]), // Field # 268

		StudentSpouseAffirmationStatus: preprocessString2627(record[studentSpouseAffirmationStatusStartIndex2627-1 : (studentSpouseAffirmationStatusStartIndex2627-1)+studentSpouseAffirmationStatusLength2627]), // Field # 269

		ParentAffirmationStatus: preprocessString2627(record[parentAffirmationStatusStartIndex2627-1 : (parentAffirmationStatusStartIndex2627-1)+parentAffirmationStatusLength2627]), // Field # 270

		ParentSpouseOrPartnerAffirmationStatus: preprocessString2627(record[parentSpouseOrPartnerAffirmationStatusStartIndex2627-1 : (parentSpouseOrPartnerAffirmationStatusStartIndex2627-1)+parentSpouseOrPartnerAffirmationStatusLength2627]), // Field # 271

		StudentDateConsentGranted: parseISIRDate2627(preprocessString2627(record[studentDateConsentGrantedStartIndex2627-1 : (studentDateConsentGrantedStartIndex2627-1)+studentDateConsentGrantedLength2627])), // Field # 272

		StudentSpouseDateConsentGranted: parseISIRDate2627(preprocessString2627(record[studentSpouseDateConsentGrantedStartIndex2627-1 : (studentSpouseDateConsentGrantedStartIndex2627-1)+studentSpouseDateConsentGrantedLength2627])), // Field # 273

		ParentDateConsentGranted: parseISIRDate2627(preprocessString2627(record[parentDateConsentGrantedStartIndex2627-1 : (parentDateConsentGrantedStartIndex2627-1)+parentDateConsentGrantedLength2627])), // Field # 274

		ParentSpouseOrPartnerDateConsentGranted: parseISIRDate2627(preprocessString2627(record[parentSpouseOrPartnerDateConsentGrantedStartIndex2627-1 : (parentSpouseOrPartnerDateConsentGrantedStartIndex2627-1)+parentSpouseOrPartnerDateConsentGrantedLength2627])), // Field # 275

		StudentTransunionMatchStatus: preprocessString2627(record[studentTransunionMatchStatusStartIndex2627-1 : (studentTransunionMatchStatusStartIndex2627-1)+studentTransunionMatchStatusLength2627]), // Field # 276

		StudentSpouseTransunionMatchStatus: preprocessString2627(record[studentSpouseTransunionMatchStatusStartIndex2627-1 : (studentSpouseTransunionMatchStatusStartIndex2627-1)+studentSpouseTransunionMatchStatusLength2627]), // Field # 277

		StudentParentTransunionMatchStatus: preprocessString2627(record[studentParentTransunionMatchStatusStartIndex2627-1 : (studentParentTransunionMatchStatusStartIndex2627-1)+studentParentTransunionMatchStatusLength2627]), // Field # 278

		StudentParentSpouseTransunionMatchStatus: preprocessString2627(record[studentParentSpouseTransunionMatchStatusStartIndex2627-1 : (studentParentSpouseTransunionMatchStatusStartIndex2627-1)+studentParentSpouseTransunionMatchStatusLength2627]), // Field # 279

		CorrectionAppliedAgainstTransactionNumber: preprocessString2627(record[correctionAppliedAgainstTransactionNumberStartIndex2627-1 : (correctionAppliedAgainstTransactionNumberStartIndex2627-1)+correctionAppliedAgainstTransactionNumberLength2627]), // Field # 280

		ProfessionalJudgement: preprocessString2627(record[professionalJudgementStartIndex2627-1 : (professionalJudgementStartIndex2627-1)+professionalJudgementLength2627]), // Field # 281

		DependencyOverrideIndicator: preprocessString2627(record[dependencyOverrideIndicatorStartIndex2627-1 : (dependencyOverrideIndicatorStartIndex2627-1)+dependencyOverrideIndicatorLength2627]), // Field # 282

		FAAFederalSchoolCode: preprocessString2627(record[fAAFederalSchoolCodeStartIndex2627-1 : (fAAFederalSchoolCodeStartIndex2627-1)+fAAFederalSchoolCodeLength2627]), // Field # 283

		FAASignature: preprocessString2627(record[fAASignatureStartIndex2627-1 : (fAASignatureStartIndex2627-1)+fAASignatureLength2627]), // Field # 284

		IASGIndicator: preprocessString2627(record[iASGIndicatorStartIndex2627-1 : (iASGIndicatorStartIndex2627-1)+iASGIndicatorLength2627]), // Field # 285

		ChildrenOfFallenHeroesIndicator: preprocessString2627(record[childrenOfFallenHeroesIndicatorStartIndex2627-1 : (childrenOfFallenHeroesIndicatorStartIndex2627-1)+childrenOfFallenHeroesIndicatorLength2627]), // Field # 286

		ElectronicTransactionIndicatorDestinationNumber: preprocessString2627(record[electronicTransactionIndicatorDestinationNumberStartIndex2627-1 : (electronicTransactionIndicatorDestinationNumberStartIndex2627-1)+electronicTransactionIndicatorDestinationNumberLength2627]), // Field # 287

		StudentSignatureSource: preprocessString2627(record[studentSignatureSourceStartIndex2627-1 : (studentSignatureSourceStartIndex2627-1)+studentSignatureSourceLength2627]), // Field # 288

		StudentSpouseSignatureSource: preprocessString2627(record[studentSpouseSignatureSourceStartIndex2627-1 : (studentSpouseSignatureSourceStartIndex2627-1)+studentSpouseSignatureSourceLength2627]), // Field # 289

		ParentSignatureSource: preprocessString2627(record[parentSignatureSourceStartIndex2627-1 : (parentSignatureSourceStartIndex2627-1)+parentSignatureSourceLength2627]), // Field # 290

		ParentSpouseOrPartnerSignatureSource: preprocessString2627(record[parentSpouseOrPartnerSignatureSourceStartIndex2627-1 : (parentSpouseOrPartnerSignatureSourceStartIndex2627-1)+parentSpouseOrPartnerSignatureSourceLength2627]), // Field # 291

		SpecialHandlingIndicator: preprocessString2627(record[specialHandlingIndicatorStartIndex2627-1 : (specialHandlingIndicatorStartIndex2627-1)+specialHandlingIndicatorLength2627]), // Field # 292

		AddressOnlyChangeFlag: preprocessString2627(record[addressOnlyChangeFlagStartIndex2627-1 : (addressOnlyChangeFlagStartIndex2627-1)+addressOnlyChangeFlagLength2627]), // Field # 293

		FPSPushedISIRFlag: preprocessString2627(record[fpsPushedISIRFlagStartIndex2627-1 : (fpsPushedISIRFlagStartIndex2627-1)+fpsPushedISIRFlagLength2627]), // Field # 294

		RejectStatusChangeFlag: preprocessString2627(record[rejectStatusChangeFlagStartIndex2627-1 : (rejectStatusChangeFlagStartIndex2627-1)+rejectStatusChangeFlagLength2627]), // Field # 295

		VerificationTrackingFlag: preprocessString2627(record[verificationTrackingFlagStartIndex2627-1 : (verificationTrackingFlagStartIndex2627-1)+verificationTrackingFlagLength2627]), // Field # 296

		StudentSelectedForVerification: preprocessString2627(record[studentSelectedForVerificationStartIndex2627-1 : (studentSelectedForVerificationStartIndex2627-1)+studentSelectedForVerificationLength2627]), // Field # 297

		IncarceratedApplicantFlag: preprocessString2627(record[incarceratedApplicantFlagStartIndex2627-1 : (incarceratedApplicantFlagStartIndex2627-1)+incarceratedApplicantFlagLength2627]), // Field # 298

		NSLDSTransactionNumber: preprocessString2627(record[nsldsTransactionNumberStartIndex2627-1 : (nsldsTransactionNumberStartIndex2627-1)+nsldsTransactionNumberLength2627]), // Field # 299

		NSLDSDatabaseResultsFlag: preprocessString2627(record[nsldsDatabaseResultsFlagStartIndex2627-1 : (nsldsDatabaseResultsFlagStartIndex2627-1)+nsldsDatabaseResultsFlagLength2627]), // Field # 300

		HighSchoolFlag: preprocessString2627(record[highSchoolFlagStartIndex2627-1 : (highSchoolFlagStartIndex2627-1)+highSchoolFlagLength2627]), // Field # 301

		StudentTotalFederalWorkStudyEarnings: preprocessString2627(record[studentTotalFederalWorkStudyEarningsStartIndex2627-1 : (studentTotalFederalWorkStudyEarningsStartIndex2627-1)+studentTotalFederalWorkStudyEarningsLength2627]), // Field # 302

		StudentSpouseTotalFederalWorkStudyEarnings: preprocessString2627(record[studentSpouseTotalFederalWorkStudyEarningsStartIndex2627-1 : (studentSpouseTotalFederalWorkStudyEarningsStartIndex2627-1)+studentSpouseTotalFederalWorkStudyEarningsLength2627]), // Field # 303

		ParentTotalFederalWorkStudyEarnings: preprocessString2627(record[parentTotalFederalWorkStudyEarningsStartIndex2627-1 : (parentTotalFederalWorkStudyEarningsStartIndex2627-1)+parentTotalFederalWorkStudyEarningsLength2627]), // Field # 304

		ParentSpouseOrPartnerTotalFederalWorkStudyEarnings: preprocessString2627(record[parentSpouseOrPartnerTotalFederalWorkStudyEarningsStartIndex2627-1 : (parentSpouseOrPartnerTotalFederalWorkStudyEarningsStartIndex2627-1)+parentSpouseOrPartnerTotalFederalWorkStudyEarningsLength2627]), // Field # 305

		TotalParentAllowancesAgainstIncome: preprocessString2627(record[totalParentAllowancesAgainstIncomeStartIndex2627-1 : (totalParentAllowancesAgainstIncomeStartIndex2627-1)+totalParentAllowancesAgainstIncomeLength2627]), // Field # 306

		ParentPayrollTaxAllowance: preprocessString2627(record[parentPayrollTaxAllowanceStartIndex2627-1 : (parentPayrollTaxAllowanceStartIndex2627-1)+parentPayrollTaxAllowanceLength2627]), // Field # 307

		ParentIncomeProtectionAllowance: preprocessString2627(record[parentIncomeProtectionAllowanceStartIndex2627-1 : (parentIncomeProtectionAllowanceStartIndex2627-1)+parentIncomeProtectionAllowanceLength2627]), // Field # 308

		ParentEmploymentExpenseAllowance: preprocessString2627(record[parentEmploymentExpenseAllowanceStartIndex2627-1 : (parentEmploymentExpenseAllowanceStartIndex2627-1)+parentEmploymentExpenseAllowanceLength2627]), // Field # 309

		ParentAvailableIncome: preprocessString2627(record[parentAvailableIncomeStartIndex2627-1 : (parentAvailableIncomeStartIndex2627-1)+parentAvailableIncomeLength2627]), // Field # 310

		ParentAdjustedAvailableIncome: preprocessString2627(record[parentAdjustedAvailableIncomeStartIndex2627-1 : (parentAdjustedAvailableIncomeStartIndex2627-1)+parentAdjustedAvailableIncomeLength2627]), // Field # 311

		ParentContribution: preprocessString2627(record[parentContributionStartIndex2627-1 : (parentContributionStartIndex2627-1)+parentContributionLength2627]), // Field # 312

		StudentPayrollTaxAllowance: preprocessString2627(record[studentPayrollTaxAllowanceStartIndex2627-1 : (studentPayrollTaxAllowanceStartIndex2627-1)+studentPayrollTaxAllowanceLength2627]), // Field # 313

		StudentIncomeProtectionAllowance: preprocessString2627(record[studentIncomeProtectionAllowanceStartIndex2627-1 : (studentIncomeProtectionAllowanceStartIndex2627-1)+studentIncomeProtectionAllowanceLength2627]), // Field # 314

		StudentAllowanceForParentsNegativeAdjustedAvailableIncome: preprocessString2627(record[studentAllowanceForParentsNegativeAdjustedAvailableIncomeStartIndex2627-1 : (studentAllowanceForParentsNegativeAdjustedAvailableIncomeStartIndex2627-1)+studentAllowanceForParentsNegativeAdjustedAvailableIncomeLength2627]), // Field # 315

		StudentEmploymentExpenseAllowance: preprocessString2627(record[studentEmploymentExpenseAllowanceStartIndex2627-1 : (studentEmploymentExpenseAllowanceStartIndex2627-1)+studentEmploymentExpenseAllowanceLength2627]), // Field # 316

		TotalStudentAllowancesAgainstIncome: preprocessString2627(record[totalStudentAllowancesAgainstIncomeStartIndex2627-1 : (totalStudentAllowancesAgainstIncomeStartIndex2627-1)+totalStudentAllowancesAgainstIncomeLength2627]), // Field # 317

		StudentAvailableIncome: preprocessString2627(record[studentAvailableIncomeStartIndex2627-1 : (studentAvailableIncomeStartIndex2627-1)+studentAvailableIncomeLength2627]), // Field # 318

		StudentContributionFromIncome: preprocessString2627(record[studentContributionFromIncomeStartIndex2627-1 : (studentContributionFromIncomeStartIndex2627-1)+studentContributionFromIncomeLength2627]), // Field # 319

		StudentAdjustedAvailableIncome: preprocessString2627(record[studentAdjustedAvailableIncomeStartIndex2627-1 : (studentAdjustedAvailableIncomeStartIndex2627-1)+studentAdjustedAvailableIncomeLength2627]), // Field # 320

		TotalStudentContributionFromSAAI: preprocessString2627(record[totalStudentContributionFromSAAIStartIndex2627-1 : (totalStudentContributionFromSAAIStartIndex2627-1)+totalStudentContributionFromSAAILength2627]), // Field # 321

		ParentDiscretionaryNetWorth: preprocessString2627(record[parentDiscretionaryNetWorthStartIndex2627-1 : (parentDiscretionaryNetWorthStartIndex2627-1)+parentDiscretionaryNetWorthLength2627]), // Field # 322

		ParentNetWorth: preprocessString2627(record[parentNetWorthStartIndex2627-1 : (parentNetWorthStartIndex2627-1)+parentNetWorthLength2627]), // Field # 323

		ParentAssetProtectionAllowance: preprocessString2627(record[parentAssetProtectionAllowanceStartIndex2627-1 : (parentAssetProtectionAllowanceStartIndex2627-1)+parentAssetProtectionAllowanceLength2627]), // Field # 324

		ParentContributionFromAssets: preprocessString2627(record[parentContributionFromAssetsStartIndex2627-1 : (parentContributionFromAssetsStartIndex2627-1)+parentContributionFromAssetsLength2627]), // Field # 325

		StudentNetWorth: preprocessString2627(record[studentNetWorthStartIndex2627-1 : (studentNetWorthStartIndex2627-1)+studentNetWorthLength2627]), // Field # 326

		StudentAssetProtectionAllowance: preprocessString2627(record[studentAssetProtectionAllowanceStartIndex2627-1 : (studentAssetProtectionAllowanceStartIndex2627-1)+studentAssetProtectionAllowanceLength2627]), // Field # 327

		StudentContributionFromAssets: preprocessString2627(record[studentContributionFromAssetsStartIndex2627-1 : (studentContributionFromAssetsStartIndex2627-1)+studentContributionFromAssetsLength2627]), // Field # 328

		AssumedStudentFamilySize: preprocessString2627(record[assumedStudentFamilySizeStartIndex2627-1 : (assumedStudentFamilySizeStartIndex2627-1)+assumedStudentFamilySizeLength2627]), // Field # 329

		AssumedParentFamilySize: preprocessString2627(record[assumedParentFamilySizeStartIndex2627-1 : (assumedParentFamilySizeStartIndex2627-1)+assumedParentFamilySizeLength2627]), // Field # 330

		StudentFirstNameCHVFlags: preprocessString2627(record[studentFirstNameCHVFlagsStartIndex2627-1 : (studentFirstNameCHVFlagsStartIndex2627-1)+studentFirstNameCHVFlagsLength2627]), // Field # 331

		StudentMiddleNameCHVFlags: preprocessString2627(record[studentMiddleNameCHVFlagsStartIndex2627-1 : (studentMiddleNameCHVFlagsStartIndex2627-1)+studentMiddleNameCHVFlagsLength2627]), // Field # 332

		StudentLastNameCHVFLags: preprocessString2627(record[studentLastNameCHVFLagsStartIndex2627-1 : (studentLastNameCHVFLagsStartIndex2627-1)+studentLastNameCHVFLagsLength2627]), // Field # 333

		StudentSuffixCHVFLags: preprocessString2627(record[studentSuffixCHVFLagsStartIndex2627-1 : (studentSuffixCHVFLagsStartIndex2627-1)+studentSuffixCHVFLagsLength2627]), // Field # 334

		StudentDateOfBirthCHVFLags: preprocessString2627(record[studentDateOfBirthCHVFLagsStartIndex2627-1 : (studentDateOfBirthCHVFLagsStartIndex2627-1)+studentDateOfBirthCHVFLagsLength2627]), // Field # 335

		StudentSSNCHVFlags: preprocessString2627(record[studentSSNCHVFlagsStartIndex2627-1 : (studentSSNCHVFlagsStartIndex2627-1)+studentSSNCHVFlagsLength2627]), // Field # 336

		StudentITINCHVFLags: preprocessString2627(record[studentITINCHVFLagsStartIndex2627-1 : (studentITINCHVFLagsStartIndex2627-1)+studentITINCHVFLagsLength2627]), // Field # 337

		StudentPhoneNumberCHVFlags: preprocessString2627(record[studentPhoneNumberCHVFlagsStartIndex2627-1 : (studentPhoneNumberCHVFlagsStartIndex2627-1)+studentPhoneNumberCHVFlagsLength2627]), // Field # 338

		StudentEmailAddressCHVFlags: preprocessString2627(record[studentEmailAddressCHVFlagsStartIndex2627-1 : (studentEmailAddressCHVFlagsStartIndex2627-1)+studentEmailAddressCHVFlagsLength2627]), // Field # 339

		StudentStreetAddressCHVFlags: preprocessString2627(record[studentStreetAddressCHVFlagsStartIndex2627-1 : (studentStreetAddressCHVFlagsStartIndex2627-1)+studentStreetAddressCHVFlagsLength2627]), // Field # 340

		StudentCityCHVFLags: preprocessString2627(record[studentCityCHVFLagsStartIndex2627-1 : (studentCityCHVFLagsStartIndex2627-1)+studentCityCHVFLagsLength2627]), // Field # 341

		StudentStateCHVFlags: preprocessString2627(record[studentStateCHVFlagsStartIndex2627-1 : (studentStateCHVFlagsStartIndex2627-1)+studentStateCHVFlagsLength2627]), // Field # 342

		StudentZipCodeCHVFlags: preprocessString2627(record[studentZipCodeCHVFlagsStartIndex2627-1 : (studentZipCodeCHVFlagsStartIndex2627-1)+studentZipCodeCHVFlagsLength2627]), // Field # 343

		StudentCountryCHVFlags: preprocessString2627(record[studentCountryCHVFlagsStartIndex2627-1 : (studentCountryCHVFlagsStartIndex2627-1)+studentCountryCHVFlagsLength2627]), // Field # 344

		StudentMaritalStatusCHVFlags: preprocessString2627(record[studentMaritalStatusCHVFlagsStartIndex2627-1 : (studentMaritalStatusCHVFlagsStartIndex2627-1)+studentMaritalStatusCHVFlagsLength2627]), // Field # 345

		StudentGradeLevelInCollegeCHVFlags: preprocessString2627(record[studentGradeLevelInCollegeCHVFlagsStartIndex2627-1 : (studentGradeLevelInCollegeCHVFlagsStartIndex2627-1)+studentGradeLevelInCollegeCHVFlagsLength2627]), // Field # 346

		StudentHasBachelorsDegreeCHVFlags: preprocessString2627(record[studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsStartIndex2627-1 : (studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsStartIndex2627-1)+studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsLength2627]), // Field # 347

		StudentPursuingTeacherCertificationCHVFlags: preprocessString2627(record[studentPursuingTeacherCertificationCHVFlagsStartIndex2627-1 : (studentPursuingTeacherCertificationCHVFlagsStartIndex2627-1)+studentPursuingTeacherCertificationCHVFlagsLength2627]), // Field # 348

		StudentActiveDutyCHVFlags: preprocessString2627(record[studentActiveDutyCHVFlagsStartIndex2627-1 : (studentActiveDutyCHVFlagsStartIndex2627-1)+studentActiveDutyCHVFlagsLength2627]), // Field # 349

		StudentVeteranCHVFlags: preprocessString2627(record[studentVeteranCHVFlagsStartIndex2627-1 : (studentVeteranCHVFlagsStartIndex2627-1)+studentVeteranCHVFlagsLength2627]), // Field # 350

		StudentChildOrOtherDependentsCHVFlags: preprocessString2627(record[studentChildOrOtherDependentsCHVFlagsStartIndex2627-1 : (studentChildOrOtherDependentsCHVFlagsStartIndex2627-1)+studentChildOrOtherDependentsCHVFlagsLength2627]), // Field # 351

		StudentParentsDeceasedCHVFlags: preprocessString2627(record[studentParentsDeceasedCHVFlagsStartIndex2627-1 : (studentParentsDeceasedCHVFlagsStartIndex2627-1)+studentParentsDeceasedCHVFlagsLength2627]), // Field # 352

		StudentWardOfCourtCHVFlags: preprocessString2627(record[studentWardOfCourtCHVFlagsStartIndex2627-1 : (studentWardOfCourtCHVFlagsStartIndex2627-1)+studentWardOfCourtCHVFlagsLength2627]), // Field # 353

		StudentInFosterCareCHVFlags: preprocessString2627(record[studentInFosterCareCHVFlagsStartIndex2627-1 : (studentInFosterCareCHVFlagsStartIndex2627-1)+studentInFosterCareCHVFlagsLength2627]), // Field # 354

		StudentEmancipatedMinorCHVFlags: preprocessString2627(record[studentEmancipatedMinorCHVFlagsStartIndex2627-1 : (studentEmancipatedMinorCHVFlagsStartIndex2627-1)+studentEmancipatedMinorCHVFlagsLength2627]), // Field # 355

		StudentLegalGuardianshipCHVFlags: preprocessString2627(record[studentLegalGuardianshipCHVFlagsStartIndex2627-1 : (studentLegalGuardianshipCHVFlagsStartIndex2627-1)+studentLegalGuardianshipCHVFlagsLength2627]), // Field # 356

		StudentPersonalCircumstancesNoneOfTheAboveCHVFlags: preprocessString2627(record[studentPersonalCircumstancesNoneOfTheAboveCHVFlagsStartIndex2627-1 : (studentPersonalCircumstancesNoneOfTheAboveCHVFlagsStartIndex2627-1)+studentPersonalCircumstancesNoneOfTheAboveCHVFlagsLength2627]), // Field # 357

		StudentUnaccompaniedHomelessOrIsUnaccompaniedCHVFlags: preprocessString2627(record[studentUnaccompaniedHomelessOrIsUnaccompaniedCHVFlagsStartIndex2627-1 : (studentUnaccompaniedHomelessOrIsUnaccompaniedCHVFlagsStartIndex2627-1)+studentUnaccompaniedHomelessOrIsUnaccompaniedCHVFlagsLength2627]), // Field # 358

		StudentUnaccompaniedAndHomelessGeneralCHVFlags: preprocessString2627(record[studentUnaccompaniedAndHomelessGeneralCHVFlagsStartIndex2627-1 : (studentUnaccompaniedAndHomelessGeneralCHVFlagsStartIndex2627-1)+studentUnaccompaniedAndHomelessGeneralCHVFlagsLength2627]), // Field # 359

		StudentUnaccompaniedAndHomelessHSCHVFlags: preprocessString2627(record[studentUnaccompaniedAndHomelessHSCHVFlagsStartIndex2627-1 : (studentUnaccompaniedAndHomelessHSCHVFlagsStartIndex2627-1)+studentUnaccompaniedAndHomelessHSCHVFlagsLength2627]), // Field # 360

		StudentUnaccompaniedAndHomelessTRIOCHVFlags: preprocessString2627(record[studentUnaccompaniedAndHomelessTRIOCHVFlagsStartIndex2627-1 : (studentUnaccompaniedAndHomelessTRIOCHVFlagsStartIndex2627-1)+studentUnaccompaniedAndHomelessTRIOCHVFlagsLength2627]), // Field # 361

		StudentUnaccompaniedAndHomelessFAACHVFlags: preprocessString2627(record[studentUnaccompaniedAndHomelessFAACHVFlagsStartIndex2627-1 : (studentUnaccompaniedAndHomelessFAACHVFlagsStartIndex2627-1)+studentUnaccompaniedAndHomelessFAACHVFlagsLength2627]), // Field # 362

		StudentHomelessnessNoneOfTheAboveCHVFlags: preprocessString2627(record[studentHomelessnessNoneOfTheAboveCHVFlagsStartIndex2627-1 : (studentHomelessnessNoneOfTheAboveCHVFlagsStartIndex2627-1)+studentHomelessnessNoneOfTheAboveCHVFlagsLength2627]), // Field # 363

		StudentHasUnusualCircumstanceCHVFlags: preprocessString2627(record[studentHasUnusualCircumstanceCHVFlagsStartIndex2627-1 : (studentHasUnusualCircumstanceCHVFlagsStartIndex2627-1)+studentHasUnusualCircumstanceCHVFlagsLength2627]), // Field # 364

		StudentUnsubOnlyCHVFlags: preprocessString2627(record[studentUnsubOnlyCHVFlagsStartIndex2627-1 : (studentUnsubOnlyCHVFlagsStartIndex2627-1)+studentUnsubOnlyCHVFlagsLength2627]), // Field # 365

		StudentUpdatedFamilySizeCHVFlags: preprocessString2627(record[studentUpdatedFamilySizeCHVFlagsStartIndex2627-1 : (studentUpdatedFamilySizeCHVFlagsStartIndex2627-1)+studentUpdatedFamilySizeCHVFlagsLength2627]), // Field # 366

		StudentNumberInCollegeCorrectionCHVFlags: preprocessString2627(record[studentNumberInCollegeCorrectionCHVFlagsStartIndex2627-1 : (studentNumberInCollegeCorrectionCHVFlagsStartIndex2627-1)+studentNumberInCollegeCorrectionCHVFlagsLength2627]), // Field # 367

		StudentCitizenshipStatusCorrectionCHVFlags: preprocessString2627(record[studentCitizenshipStatusCorrectionCHVFlagsStartIndex2627-1 : (studentCitizenshipStatusCorrectionCHVFlagsStartIndex2627-1)+studentCitizenshipStatusCorrectionCHVFlagsLength2627]), // Field # 368

		StudentANumberCHVFlags: preprocessString2627(record[studentANumberCHVFlagsStartIndex2627-1 : (studentANumberCHVFlagsStartIndex2627-1)+studentANumberCHVFlagsLength2627]), // Field # 369

		StudentStateOfLegalResidenceCHVFlags: preprocessString2627(record[studentStateOfLegalResidenceCHVFlagsStartIndex2627-1 : (studentStateOfLegalResidenceCHVFlagsStartIndex2627-1)+studentStateOfLegalResidenceCHVFlagsLength2627]), // Field # 370

		StudentLegalResidenceDateCHVFlags: preprocessString2627(record[studentLegalResidenceDateCHVFlagsStartIndex2627-1 : (studentLegalResidenceDateCHVFlagsStartIndex2627-1)+studentLegalResidenceDateCHVFlagsLength2627]), // Field # 371

		StudentEitherParentAttendCollegeCHVFlags: preprocessString2627(record[studentEitherParentAttendCollegeCHVFlagsStartIndex2627-1 : (studentEitherParentAttendCollegeCHVFlagsStartIndex2627-1)+studentEitherParentAttendCollegeCHVFlagsLength2627]), // Field # 372

		StudentParentKilledInTheLineOfDutyCHVFlags: preprocessString2627(record[studentParentKilledInTheLineOfDutyCHVFlagsStartIndex2627-1 : (studentParentKilledInTheLineOfDutyCHVFlagsStartIndex2627-1)+studentParentKilledInTheLineOfDutyCHVFlagsLength2627]), // Field # 373

		StudentHighSchoolCompletionStatusCHVFlags: preprocessString2627(record[studentHighSchoolCompletionStatusCHVFlagsStartIndex2627-1 : (studentHighSchoolCompletionStatusCHVFlagsStartIndex2627-1)+studentHighSchoolCompletionStatusCHVFlagsLength2627]), // Field # 374

		StudentHighSchoolNameCHVFlags: preprocessString2627(record[studentHighSchoolNameCHVFlagsStartIndex2627-1 : (studentHighSchoolNameCHVFlagsStartIndex2627-1)+studentHighSchoolNameCHVFlagsLength2627]), // Field # 375

		StudentHighSchoolCityCHVFlags: preprocessString2627(record[studentHighSchoolCityCHVFlagsStartIndex2627-1 : (studentHighSchoolCityCHVFlagsStartIndex2627-1)+studentHighSchoolCityCHVFlagsLength2627]), // Field # 376

		StudentHighSchoolStateCHVFlags: preprocessString2627(record[studentHighSchoolStateCHVFlagsStartIndex2627-1 : (studentHighSchoolStateCHVFlagsStartIndex2627-1)+studentHighSchoolStateCHVFlagsLength2627]), // Field # 377

		StudentHighSchoolEquivalentDiplomaNameCHVFlags: preprocessString2627(record[studentHighSchoolEquivalentDiplomaNameCHVFlagsStartIndex2627-1 : (studentHighSchoolEquivalentDiplomaNameCHVFlagsStartIndex2627-1)+studentHighSchoolEquivalentDiplomaNameCHVFlagsLength2627]), // Field # 378

		StudentHighSchoolEquivalentDiplomaStateCHVFlags: preprocessString2627(record[studentHighSchoolEquivalentDiplomaStateCHVFlagsStartIndex2627-1 : (studentHighSchoolEquivalentDiplomaStateCHVFlagsStartIndex2627-1)+studentHighSchoolEquivalentDiplomaStateCHVFlagsLength2627]), // Field # 379

		StudentReceivedEITCCHVFlags: preprocessString2627(record[studentReceivedEITCCHVFlagsStartIndex2627-1 : (studentReceivedEITCCHVFlagsStartIndex2627-1)+studentReceivedEITCCHVFlagsLength2627]), // Field # 380

		StudentReceivedFederalHousingAssistanceCHVFlags: preprocessString2627(record[studentReceivedFederalHousingAssistanceCHVFlagsStartIndex2627-1 : (studentReceivedFederalHousingAssistanceCHVFlagsStartIndex2627-1)+studentReceivedFederalHousingAssistanceCHVFlagsLength2627]), // Field # 381

		StudentReceivedFreeReducedPriceLunchCHVFlags: preprocessString2627(record[studentReceivedFreeReducedPriceLunchCHVFlagsStartIndex2627-1 : (studentReceivedFreeReducedPriceLunchCHVFlagsStartIndex2627-1)+studentReceivedFreeReducedPriceLunchCHVFlagsLength2627]), // Field # 382

		StudentReceivedMedicaidCHVFlags: preprocessString2627(record[studentReceivedMedicaidCHVFlagsStartIndex2627-1 : (studentReceivedMedicaidCHVFlagsStartIndex2627-1)+studentReceivedMedicaidCHVFlagsLength2627]), // Field # 383

		StudentReceivedRefundableCreditFor36BHealthPlanCHVFlags: preprocessString2627(record[studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex2627-1 : (studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex2627-1)+studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsLength2627]), // Field # 384

		StudentReceivedSNAPCHVFlags: preprocessString2627(record[studentReceivedSNAPCHVFlagsStartIndex2627-1 : (studentReceivedSNAPCHVFlagsStartIndex2627-1)+studentReceivedSNAPCHVFlagsLength2627]), // Field # 385

		StudentReceivedSupplementalSecurityIncomeCHVFlags: preprocessString2627(record[studentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex2627-1 : (studentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex2627-1)+studentReceivedSupplementalSecurityIncomeCHVFlagsLength2627]), // Field # 386

		StudentReceivedTANFCHVFlags: preprocessString2627(record[studentReceivedTANFCHVFlagsStartIndex2627-1 : (studentReceivedTANFCHVFlagsStartIndex2627-1)+studentReceivedTANFCHVFlagsLength2627]), // Field # 387

		StudentReceivedWICCHVFlags: preprocessString2627(record[studentReceivedWICCHVFlagsStartIndex2627-1 : (studentReceivedWICCHVFlagsStartIndex2627-1)+studentReceivedWICCHVFlagsLength2627]), // Field # 388

		StudentFederalBenefitsNoneOfTheAboveCHVFlags: preprocessString2627(record[studentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex2627-1 : (studentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex2627-1)+studentFederalBenefitsNoneOfTheAboveCHVFlagsLength2627]), // Field # 389

		StudentFiled1040Or1040NRCHVFlags: preprocessString2627(record[studentFiled1040Or1040NRCHVFlagsStartIndex2627-1 : (studentFiled1040Or1040NRCHVFlagsStartIndex2627-1)+studentFiled1040Or1040NRCHVFlagsLength2627]), // Field # 390

		StudentFiledNonUSTaxReturnCHVFlags: preprocessString2627(record[studentFiledNonUSTaxReturnCHVFlagsStartIndex2627-1 : (studentFiledNonUSTaxReturnCHVFlagsStartIndex2627-1)+studentFiledNonUSTaxReturnCHVFlagsLength2627]), // Field # 391

		StudentFiledJointReturnWithCurrentSpouseCHVFlags: preprocessString2627(record[studentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex2627-1 : (studentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex2627-1)+studentFiledJointReturnWithCurrentSpouseCHVFlagsLength2627]), // Field # 392

		StudentTaxReturnFilingStatusCHVFlags: preprocessString2627(record[studentTaxReturnFilingStatusCHVFlagsStartIndex2627-1 : (studentTaxReturnFilingStatusCHVFlagsStartIndex2627-1)+studentTaxReturnFilingStatusCHVFlagsLength2627]), // Field # 393

		StudentIncomeEarnedFromWorkCorrectionCHVFlags: preprocessString2627(record[studentIncomeEarnedFromWorkCorrectionCHVFlagsStartIndex2627-1 : (studentIncomeEarnedFromWorkCorrectionCHVFlagsStartIndex2627-1)+studentIncomeEarnedFromWorkCorrectionCHVFlagsLength2627]), // Field # 394

		StudentTaxExemptInterestIncomeCHVFlags: preprocessString2627(record[studentTaxExemptInterestIncomeCHVFlagsStartIndex2627-1 : (studentTaxExemptInterestIncomeCHVFlagsStartIndex2627-1)+studentTaxExemptInterestIncomeCHVFlagsLength2627]), // Field # 395

		StudentUntaxedPortionsOfIRADistributionsCHVFlags: preprocessString2627(record[studentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2627-1 : (studentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2627-1)+studentUntaxedPortionsOfIRADistributionsCHVFlagsLength2627]), // Field # 396

		StudentIRARolloverCHVFlags: preprocessString2627(record[studentIRARolloverCHVFlagsStartIndex2627-1 : (studentIRARolloverCHVFlagsStartIndex2627-1)+studentIRARolloverCHVFlagsLength2627]), // Field # 397

		StudentUntaxedPortionsOfPensionsCHVFlags: preprocessString2627(record[studentUntaxedPortionsOfPensionsCHVFlagsStartIndex2627-1 : (studentUntaxedPortionsOfPensionsCHVFlagsStartIndex2627-1)+studentUntaxedPortionsOfPensionsCHVFlagsLength2627]), // Field # 398

		StudentPensionRolloverCHVFlags: preprocessString2627(record[studentPensionRolloverCHVFlagsStartIndex2627-1 : (studentPensionRolloverCHVFlagsStartIndex2627-1)+studentPensionRolloverCHVFlagsLength2627]), // Field # 399

		StudentAdjustedGrossIncomeCHVFlags: preprocessString2627(record[studentAdjustedGrossIncomeCHVFlagsStartIndex2627-1 : (studentAdjustedGrossIncomeCHVFlagsStartIndex2627-1)+studentAdjustedGrossIncomeCHVFlagsLength2627]), // Field # 400

		StudentIncomeTaxPaidCHVFlags: preprocessString2627(record[studentIncomeTaxPaidCHVFlagsStartIndex2627-1 : (studentIncomeTaxPaidCHVFlagsStartIndex2627-1)+studentIncomeTaxPaidCHVFlagsLength2627]), // Field # 401

		StudentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlags: preprocessString2627(record[studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex2627-1 : (studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex2627-1)+studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsLength2627]), // Field # 402

		StudentDeductiblePaymentsToIRAKeoghOtherCHVFlags: preprocessString2627(record[studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2627-1 : (studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2627-1)+studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2627]), // Field # 403

		StudentEducationCreditsCHVFlags: preprocessString2627(record[studentEducationCreditsCHVFlagsStartIndex2627-1 : (studentEducationCreditsCHVFlagsStartIndex2627-1)+studentEducationCreditsCHVFlagsLength2627]), // Field # 404

		StudentFiledScheduleABDEFHCHVFlags: preprocessString2627(record[studentFiledScheduleABDEFHCHVFlagsStartIndex2627-1 : (studentFiledScheduleABDEFHCHVFlagsStartIndex2627-1)+studentFiledScheduleABDEFHCHVFlagsLength2627]), // Field # 405

		StudentScheduleCAmountCHVFlags: preprocessString2627(record[studentScheduleCAmountCHVFlagsStartIndex2627-1 : (studentScheduleCAmountCHVFlagsStartIndex2627-1)+studentScheduleCAmountCHVFlagsLength2627]), // Field # 406

		StudentCollegeGrantAndScholarshipAidCHVFlags: preprocessString2627(record[studentCollegeGrantAndScholarshipAidCHVFlagsStartIndex2627-1 : (studentCollegeGrantAndScholarshipAidCHVFlagsStartIndex2627-1)+studentCollegeGrantAndScholarshipAidCHVFlagsLength2627]), // Field # 407

		StudentForeignEarnedIncomeExclusionCHVFlags: preprocessString2627(record[studentForeignEarnedIncomeExclusionCHVFlagsStartIndex2627-1 : (studentForeignEarnedIncomeExclusionCHVFlagsStartIndex2627-1)+studentForeignEarnedIncomeExclusionCHVFlagsLength2627]), // Field # 408

		StudentChildSupportReceivedCHVFlags: preprocessString2627(record[studentChildSupportReceivedCHVFlagsStartIndex2627-1 : (studentChildSupportReceivedCHVFlagsStartIndex2627-1)+studentChildSupportReceivedCHVFlagsLength2627]), // Field # 409

		StudentNetWorthOfBusinessesAndInvestmentFarmsCHVFlags: preprocessString2627(record[studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex2627-1 : (studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex2627-1)+studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsLength2627]), // Field # 410

		StudentNetWorthOfCurrentInvestmentsCHVFlags: preprocessString2627(record[studentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex2627-1 : (studentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex2627-1)+studentNetWorthOfCurrentInvestmentsCHVFlagsLength2627]), // Field # 411

		StudentTotalOfCashSavingsAndCheckingCHVFlags: preprocessString2627(record[studentTotalOfCashSavingsAndCheckingCHVFlagsStartIndex2627-1 : (studentTotalOfCashSavingsAndCheckingCHVFlagsStartIndex2627-1)+studentTotalOfCashSavingsAndCheckingCHVFlagsLength2627]), // Field # 412

		StudentCollege1CHVFlags: preprocessString2627(record[studentCollege1CHVFlagsStartIndex2627-1 : (studentCollege1CHVFlagsStartIndex2627-1)+studentCollege1CHVFlagsLength2627]), // Field # 413

		StudentCollege2CHVFlags: preprocessString2627(record[studentCollege2CHVFlagsStartIndex2627-1 : (studentCollege2CHVFlagsStartIndex2627-1)+studentCollege2CHVFlagsLength2627]), // Field # 414

		StudentCollege3CHVFlags: preprocessString2627(record[studentCollege3CHVFlagsStartIndex2627-1 : (studentCollege3CHVFlagsStartIndex2627-1)+studentCollege3CHVFlagsLength2627]), // Field # 415

		StudentCollege4CHVFlags: preprocessString2627(record[studentCollege4CHVFlagsStartIndex2627-1 : (studentCollege4CHVFlagsStartIndex2627-1)+studentCollege4CHVFlagsLength2627]), // Field # 416

		StudentCollege5CHVFlags: preprocessString2627(record[studentCollege5CHVFlagsStartIndex2627-1 : (studentCollege5CHVFlagsStartIndex2627-1)+studentCollege5CHVFlagsLength2627]), // Field # 417

		StudentCollege6CHVFlags: preprocessString2627(record[studentCollege6CHVFlagsStartIndex2627-1 : (studentCollege6CHVFlagsStartIndex2627-1)+studentCollege6CHVFlagsLength2627]), // Field # 418

		StudentCollege7CHVFlags: preprocessString2627(record[studentCollege7CHVFlagsStartIndex2627-1 : (studentCollege7CHVFlagsStartIndex2627-1)+studentCollege7CHVFlagsLength2627]), // Field # 419

		StudentCollege8CHVFlags: preprocessString2627(record[studentCollege8CHVFlagsStartIndex2627-1 : (studentCollege8CHVFlagsStartIndex2627-1)+studentCollege8CHVFlagsLength2627]), // Field # 420

		StudentCollege9CHVFlags: preprocessString2627(record[studentCollege9CHVFlagsStartIndex2627-1 : (studentCollege9CHVFlagsStartIndex2627-1)+studentCollege9CHVFlagsLength2627]), // Field # 421

		StudentCollege10CHVFlags: preprocessString2627(record[studentCollege10CHVFlagsStartIndex2627-1 : (studentCollege10CHVFlagsStartIndex2627-1)+studentCollege10CHVFlagsLength2627]), // Field # 422

		StudentCollege11CHVFlags: preprocessString2627(record[studentCollege11CHVFlagsStartIndex2627-1 : (studentCollege11CHVFlagsStartIndex2627-1)+studentCollege11CHVFlagsLength2627]), // Field # 423

		StudentCollege12CHVFlags: preprocessString2627(record[studentCollege12CHVFlagsStartIndex2627-1 : (studentCollege12CHVFlagsStartIndex2627-1)+studentCollege12CHVFlagsLength2627]), // Field # 424

		StudentCollege13CHVFlags: preprocessString2627(record[studentCollege13CHVFlagsStartIndex2627-1 : (studentCollege13CHVFlagsStartIndex2627-1)+studentCollege13CHVFlagsLength2627]), // Field # 425

		StudentCollege14CHVFlags: preprocessString2627(record[studentCollege14CHVFlagsStartIndex2627-1 : (studentCollege14CHVFlagsStartIndex2627-1)+studentCollege14CHVFlagsLength2627]), // Field # 426

		StudentCollege15CHVFlags: preprocessString2627(record[studentCollege15CHVFlagsStartIndex2627-1 : (studentCollege15CHVFlagsStartIndex2627-1)+studentCollege15CHVFlagsLength2627]), // Field # 427

		StudentCollege16CHVFlags: preprocessString2627(record[studentCollege16CHVFlagsStartIndex2627-1 : (studentCollege16CHVFlagsStartIndex2627-1)+studentCollege16CHVFlagsLength2627]), // Field # 428

		StudentCollege17CHVFlags: preprocessString2627(record[studentCollege17CHVFlagsStartIndex2627-1 : (studentCollege17CHVFlagsStartIndex2627-1)+studentCollege17CHVFlagsLength2627]), // Field # 429

		StudentCollege18CHVFlags: preprocessString2627(record[studentCollege18CHVFlagsStartIndex2627-1 : (studentCollege18CHVFlagsStartIndex2627-1)+studentCollege18CHVFlagsLength2627]), // Field # 430

		StudentCollege19CHVFlags: preprocessString2627(record[studentCollege19CHVFlagsStartIndex2627-1 : (studentCollege19CHVFlagsStartIndex2627-1)+studentCollege19CHVFlagsLength2627]), // Field # 431

		StudentCollege20CHVFlags: preprocessString2627(record[studentCollege20CHVFlagsStartIndex2627-1 : (studentCollege20CHVFlagsStartIndex2627-1)+studentCollege20CHVFlagsLength2627]), // Field # 432

		StudentConsentToRetrieveAndDiscloseFTICHVFlags: preprocessString2627(record[studentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2627-1 : (studentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2627-1)+studentConsentToRetrieveAndDiscloseFTICHVFlagsLength2627]), // Field # 433

		StudentSignatureCHVFlags: preprocessString2627(record[studentSignatureCHVFlagsStartIndex2627-1 : (studentSignatureCHVFlagsStartIndex2627-1)+studentSignatureCHVFlagsLength2627]), // Field # 434

		StudentSignatureDateCHVFlags: preprocessString2627(record[studentSignatureDateCHVFlagsStartIndex2627-1 : (studentSignatureDateCHVFlagsStartIndex2627-1)+studentSignatureDateCHVFlagsLength2627]), // Field # 435

		StudentSpouseFirstNameCHVFlags: preprocessString2627(record[studentSpouseFirstNameCHVFlagsStartIndex2627-1 : (studentSpouseFirstNameCHVFlagsStartIndex2627-1)+studentSpouseFirstNameCHVFlagsLength2627]), // Field # 436

		StudentSpouseMiddleNameCHVFlags: preprocessString2627(record[studentSpouseMiddleNameCHVFlagsStartIndex2627-1 : (studentSpouseMiddleNameCHVFlagsStartIndex2627-1)+studentSpouseMiddleNameCHVFlagsLength2627]), // Field # 437

		StudentSpouseLastNameCHVFlags: preprocessString2627(record[studentSpouseLastNameCHVFlagsStartIndex2627-1 : (studentSpouseLastNameCHVFlagsStartIndex2627-1)+studentSpouseLastNameCHVFlagsLength2627]), // Field # 438

		StudentSpouseSuffixCHVFlags: preprocessString2627(record[studentSpouseSuffixCHVFlagsStartIndex2627-1 : (studentSpouseSuffixCHVFlagsStartIndex2627-1)+studentSpouseSuffixCHVFlagsLength2627]), // Field # 439

		StudentSpouseDateOfBirthCHVFlags: preprocessString2627(record[studentSpouseDateOfBirthCHVFlagsStartIndex2627-1 : (studentSpouseDateOfBirthCHVFlagsStartIndex2627-1)+studentSpouseDateOfBirthCHVFlagsLength2627]), // Field # 440

		StudentSpouseSSNCHVFlags: preprocessString2627(record[studentSpouseSSNCHVFlagsStartIndex2627-1 : (studentSpouseSSNCHVFlagsStartIndex2627-1)+studentSpouseSSNCHVFlagsLength2627]), // Field # 441

		StudentSpouseITINCHVFlags: preprocessString2627(record[studentSpouseITINCHVFlagsStartIndex2627-1 : (studentSpouseITINCHVFlagsStartIndex2627-1)+studentSpouseITINCHVFlagsLength2627]), // Field # 442

		StudentSpousePhoneNumberCHVFlags: preprocessString2627(record[studentSpousePhoneNumberCHVFlagsStartIndex2627-1 : (studentSpousePhoneNumberCHVFlagsStartIndex2627-1)+studentSpousePhoneNumberCHVFlagsLength2627]), // Field # 443

		StudentSpouseEmailAddressCHVFlags: preprocessString2627(record[studentSpouseEmailAddressCHVFlagsStartIndex2627-1 : (studentSpouseEmailAddressCHVFlagsStartIndex2627-1)+studentSpouseEmailAddressCHVFlagsLength2627]), // Field # 444

		StudentSpouseStreetAddressCHVFlags: preprocessString2627(record[studentSpouseStreetAddressCHVFlagsStartIndex2627-1 : (studentSpouseStreetAddressCHVFlagsStartIndex2627-1)+studentSpouseStreetAddressCHVFlagsLength2627]), // Field # 445

		StudentSpouseCityCHVFlags: preprocessString2627(record[studentSpouseCityCHVFlagsStartIndex2627-1 : (studentSpouseCityCHVFlagsStartIndex2627-1)+studentSpouseCityCHVFlagsLength2627]), // Field # 446

		StudentSpouseStateCHVFlags: preprocessString2627(record[studentSpouseStateCHVFlagsStartIndex2627-1 : (studentSpouseStateCHVFlagsStartIndex2627-1)+studentSpouseStateCHVFlagsLength2627]), // Field # 447

		StudentSpouseZipCodeCHVFlags: preprocessString2627(record[studentSpouseZipCodeCHVFlagsStartIndex2627-1 : (studentSpouseZipCodeCHVFlagsStartIndex2627-1)+studentSpouseZipCodeCHVFlagsLength2627]), // Field # 448

		StudentSpouseCountryCHVFlags: preprocessString2627(record[studentSpouseCountryCHVFlagsStartIndex2627-1 : (studentSpouseCountryCHVFlagsStartIndex2627-1)+studentSpouseCountryCHVFlagsLength2627]), // Field # 449

		StudentSpouseFiled1040Or1040NRCHVFlags: preprocessString2627(record[studentSpouseFiled1040Or1040NRCHVFlagsStartIndex2627-1 : (studentSpouseFiled1040Or1040NRCHVFlagsStartIndex2627-1)+studentSpouseFiled1040Or1040NRCHVFlagsLength2627]), // Field # 450

		StudentSpouseFiledNonUSTaxReturnCHVFlags: preprocessString2627(record[studentSpouseFiledNonUSTaxReturnCHVFlagsStartIndex2627-1 : (studentSpouseFiledNonUSTaxReturnCHVFlagsStartIndex2627-1)+studentSpouseFiledNonUSTaxReturnCHVFlagsLength2627]), // Field # 451

		StudentSpouseTaxReturnFilingStatusCHVFlags: preprocessString2627(record[studentSpouseTaxReturnFilingStatusCHVFlagsStartIndex2627-1 : (studentSpouseTaxReturnFilingStatusCHVFlagsStartIndex2627-1)+studentSpouseTaxReturnFilingStatusCHVFlagsLength2627]), // Field # 452

		StudentSpouseIncomeEarnedFromWorkCHVFlags: preprocessString2627(record[studentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex2627-1 : (studentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex2627-1)+studentSpouseIncomeEarnedFromWorkCHVFlagsLength2627]), // Field # 453

		StudentSpouseTaxExemptInterestIncomeCHVFlags: preprocessString2627(record[studentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex2627-1 : (studentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex2627-1)+studentSpouseTaxExemptInterestIncomeCHVFlagsLength2627]), // Field # 454

		StudentSpouseUntaxedPortionsOfIRADistributionsCHVFlags: preprocessString2627(record[studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2627-1 : (studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2627-1)+studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsLength2627]), // Field # 455

		StudentSpouseIRARolloverCHVFlags: preprocessString2627(record[studentSpouseIRARolloverCHVFlagsStartIndex2627-1 : (studentSpouseIRARolloverCHVFlagsStartIndex2627-1)+studentSpouseIRARolloverCHVFlagsLength2627]), // Field # 456

		StudentSpouseUntaxedPortionsOfPensionsCHVFlags: preprocessString2627(record[studentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex2627-1 : (studentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex2627-1)+studentSpouseUntaxedPortionsOfPensionsCHVFlagsLength2627]), // Field # 457

		StudentSpousePensionRolloverCHVFlags: preprocessString2627(record[studentSpousePensionRolloverCHVFlagsStartIndex2627-1 : (studentSpousePensionRolloverCHVFlagsStartIndex2627-1)+studentSpousePensionRolloverCHVFlagsLength2627]), // Field # 458

		StudentSpouseAdjustedGrossIncomeCHVFlags: preprocessString2627(record[studentSpouseAdjustedGrossIncomeCHVFlagsStartIndex2627-1 : (studentSpouseAdjustedGrossIncomeCHVFlagsStartIndex2627-1)+studentSpouseAdjustedGrossIncomeCHVFlagsLength2627]), // Field # 459

		StudentSpouseIncomeTaxPaidCHVFlags: preprocessString2627(record[studentSpouseIncomeTaxPaidCHVFlagsStartIndex2627-1 : (studentSpouseIncomeTaxPaidCHVFlagsStartIndex2627-1)+studentSpouseIncomeTaxPaidCHVFlagsLength2627]), // Field # 460

		StudentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlags: preprocessString2627(record[studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2627-1 : (studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2627-1)+studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2627]), // Field # 461

		StudentSpouseEducationCreditsCHVFlags: preprocessString2627(record[studentSpouseEducationCreditsCHVFlagsStartIndex2627-1 : (studentSpouseEducationCreditsCHVFlagsStartIndex2627-1)+studentSpouseEducationCreditsCHVFlagsLength2627]), // Field # 462

		StudentSpouseFiledScheduleABDEFHCHVFlags: preprocessString2627(record[studentSpouseFiledScheduleABDEFHCHVFlagsStartIndex2627-1 : (studentSpouseFiledScheduleABDEFHCHVFlagsStartIndex2627-1)+studentSpouseFiledScheduleABDEFHCHVFlagsLength2627]), // Field # 463

		StudentSpouseScheduleCAmountCHVFlags: preprocessString2627(record[studentSpouseScheduleCAmountCHVFlagsStartIndex2627-1 : (studentSpouseScheduleCAmountCHVFlagsStartIndex2627-1)+studentSpouseScheduleCAmountCHVFlagsLength2627]), // Field # 464

		StudentSpouseForeignEarnedIncomeExclusionCHVFlags: preprocessString2627(record[studentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex2627-1 : (studentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex2627-1)+studentSpouseForeignEarnedIncomeExclusionCHVFlagsLength2627]), // Field # 465

		StudentSpouseConsentToRetrieveAndDiscloseFTICHVFlags: preprocessString2627(record[studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2627-1 : (studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2627-1)+studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsLength2627]), // Field # 466

		StudentSpouseSignatureCHVFlags: preprocessString2627(record[studentSpouseSignatureCHVFlagsStartIndex2627-1 : (studentSpouseSignatureCHVFlagsStartIndex2627-1)+studentSpouseSignatureCHVFlagsLength2627]), // Field # 467

		StudentSpouseSignatureDateCHVFlags: preprocessString2627(record[studentSpouseSignatureDateCHVFlagsStartIndex2627-1 : (studentSpouseSignatureDateCHVFlagsStartIndex2627-1)+studentSpouseSignatureDateCHVFlagsLength2627]), // Field # 468

		ParentFirstNameCHVFlags: preprocessString2627(record[parentFirstNameCHVFlagsStartIndex2627-1 : (parentFirstNameCHVFlagsStartIndex2627-1)+parentFirstNameCHVFlagsLength2627]), // Field # 469

		ParentMiddleNameCHVFlags: preprocessString2627(record[parentMiddleNameCHVFlagsStartIndex2627-1 : (parentMiddleNameCHVFlagsStartIndex2627-1)+parentMiddleNameCHVFlagsLength2627]), // Field # 470

		ParentLastNameCHVFlags: preprocessString2627(record[parentLastNameCHVFlagsStartIndex2627-1 : (parentLastNameCHVFlagsStartIndex2627-1)+parentLastNameCHVFlagsLength2627]), // Field # 471

		ParentSuffixCHVFlags: preprocessString2627(record[parentSuffixCHVFlagsStartIndex2627-1 : (parentSuffixCHVFlagsStartIndex2627-1)+parentSuffixCHVFlagsLength2627]), // Field # 472

		ParentDateOfBirthCHVFlags: preprocessString2627(record[parentDateOfBirthCHVFlagsStartIndex2627-1 : (parentDateOfBirthCHVFlagsStartIndex2627-1)+parentDateOfBirthCHVFlagsLength2627]), // Field # 473

		ParentSSNCHVFlags: preprocessString2627(record[parentSSNCHVFlagsStartIndex2627-1 : (parentSSNCHVFlagsStartIndex2627-1)+parentSSNCHVFlagsLength2627]), // Field # 474

		ParentITINCHVFlags: preprocessString2627(record[parentITINCHVFlagsStartIndex2627-1 : (parentITINCHVFlagsStartIndex2627-1)+parentITINCHVFlagsLength2627]), // Field # 475

		ParentPhoneNumberCHVFlags: preprocessString2627(record[parentPhoneNumberCHVFlagsStartIndex2627-1 : (parentPhoneNumberCHVFlagsStartIndex2627-1)+parentPhoneNumberCHVFlagsLength2627]), // Field # 476

		ParentEmailAddressCHVFlags: preprocessString2627(record[parentEmailAddressCHVFlagsStartIndex2627-1 : (parentEmailAddressCHVFlagsStartIndex2627-1)+parentEmailAddressCHVFlagsLength2627]), // Field # 477

		ParentStreetAddressCHVFlags: preprocessString2627(record[parentStreetAddressCHVFlagsStartIndex2627-1 : (parentStreetAddressCHVFlagsStartIndex2627-1)+parentStreetAddressCHVFlagsLength2627]), // Field # 478

		ParentCityCHVFlags: preprocessString2627(record[parentCityCHVFlagsStartIndex2627-1 : (parentCityCHVFlagsStartIndex2627-1)+parentCityCHVFlagsLength2627]), // Field # 479

		ParentStateCHVFlags: preprocessString2627(record[parentStateCHVFlagsStartIndex2627-1 : (parentStateCHVFlagsStartIndex2627-1)+parentStateCHVFlagsLength2627]), // Field # 480

		ParentZipCodeCHVFlags: preprocessString2627(record[parentZipCodeCHVFlagsStartIndex2627-1 : (parentZipCodeCHVFlagsStartIndex2627-1)+parentZipCodeCHVFlagsLength2627]), // Field # 481

		ParentCountryCHVFlags: preprocessString2627(record[parentCountryCHVFlagsStartIndex2627-1 : (parentCountryCHVFlagsStartIndex2627-1)+parentCountryCHVFlagsLength2627]), // Field # 482

		ParentMaritalStatusCHVFlags: preprocessString2627(record[parentMaritalStatusCHVFlagsStartIndex2627-1 : (parentMaritalStatusCHVFlagsStartIndex2627-1)+parentMaritalStatusCHVFlagsLength2627]), // Field # 483

		ParentStateOfLegalResidenceCHVFlags: preprocessString2627(record[parentStateOfLegalResidenceCHVFlagsStartIndex2627-1 : (parentStateOfLegalResidenceCHVFlagsStartIndex2627-1)+parentStateOfLegalResidenceCHVFlagsLength2627]), // Field # 484

		ParentLegalResidenceDateCHVFlags: preprocessString2627(record[parentLegalResidenceDateCHVFlagsStartIndex2627-1 : (parentLegalResidenceDateCHVFlagsStartIndex2627-1)+parentLegalResidenceDateCHVFlagsLength2627]), // Field # 485

		ParentUpdatedFamilySizeCHVFlags: preprocessString2627(record[parentUpdatedFamilySizeCHVFlagsStartIndex2627-1 : (parentUpdatedFamilySizeCHVFlagsStartIndex2627-1)+parentUpdatedFamilySizeCHVFlagsLength2627]), // Field # 486

		ParentNumberInCollegeCHVFlags: preprocessString2627(record[parentNumberInCollegeCHVFlagsStartIndex2627-1 : (parentNumberInCollegeCHVFlagsStartIndex2627-1)+parentNumberInCollegeCHVFlagsLength2627]), // Field # 487

		ParentReceivedEITCCHVFlags: preprocessString2627(record[parentReceivedEITCCHVFlagsStartIndex2627-1 : (parentReceivedEITCCHVFlagsStartIndex2627-1)+parentReceivedEITCCHVFlagsLength2627]), // Field # 488

		ParentReceivedFederalHousingAssistanceCHVFlags: preprocessString2627(record[parentReceivedFederalHousingAssistanceCHVFlagsStartIndex2627-1 : (parentReceivedFederalHousingAssistanceCHVFlagsStartIndex2627-1)+parentReceivedFederalHousingAssistanceCHVFlagsLength2627]), // Field # 489

		ParentReceivedFreeReducedPriceLunchCHVFlags: preprocessString2627(record[parentReceivedFreeReducedPriceLunchCHVFlagsStartIndex2627-1 : (parentReceivedFreeReducedPriceLunchCHVFlagsStartIndex2627-1)+parentReceivedFreeReducedPriceLunchCHVFlagsLength2627]), // Field # 490

		ParentReceivedMedicaidCHVFlags: preprocessString2627(record[parentReceivedMedicaidCHVFlagsStartIndex2627-1 : (parentReceivedMedicaidCHVFlagsStartIndex2627-1)+parentReceivedMedicaidCHVFlagsLength2627]), // Field # 491

		ParentReceivedRefundableCreditFor36BHealthPlanCHVFlags: preprocessString2627(record[parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex2627-1 : (parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex2627-1)+parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsLength2627]), // Field # 492

		ParentReceivedSNAPCHVFlags: preprocessString2627(record[parentReceivedSNAPCHVFlagsStartIndex2627-1 : (parentReceivedSNAPCHVFlagsStartIndex2627-1)+parentReceivedSNAPCHVFlagsLength2627]), // Field # 493

		ParentReceivedSupplementalSecurityIncomeCHVFlags: preprocessString2627(record[parentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex2627-1 : (parentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex2627-1)+parentReceivedSupplementalSecurityIncomeCHVFlagsLength2627]), // Field # 494

		ParentReceivedTANFCHVFlags: preprocessString2627(record[parentReceivedTANFCHVFlagsStartIndex2627-1 : (parentReceivedTANFCHVFlagsStartIndex2627-1)+parentReceivedTANFCHVFlagsLength2627]), // Field # 495

		ParentReceivedWICCHVFlags: preprocessString2627(record[parentReceivedWICCHVFlagsStartIndex2627-1 : (parentReceivedWICCHVFlagsStartIndex2627-1)+parentReceivedWICCHVFlagsLength2627]), // Field # 496

		ParentFederalBenefitsNoneOfTheAboveCHVFlags: preprocessString2627(record[parentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex2627-1 : (parentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex2627-1)+parentFederalBenefitsNoneOfTheAboveCHVFlagsLength2627]), // Field # 497

		ParentFiled1040Or1040NRCHVFlags: preprocessString2627(record[parentFiled1040Or1040NRCHVFlagsStartIndex2627-1 : (parentFiled1040Or1040NRCHVFlagsStartIndex2627-1)+parentFiled1040Or1040NRCHVFlagsLength2627]), // Field # 498

		ParentFileNonUSTaxReturnCHVFlags: preprocessString2627(record[parentFileNonUSTaxReturnCHVFlagsStartIndex2627-1 : (parentFileNonUSTaxReturnCHVFlagsStartIndex2627-1)+parentFileNonUSTaxReturnCHVFlagsLength2627]), // Field # 499

		ParentFiledJointReturnWithCurrentSpouseCHVFlags: preprocessString2627(record[parentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex2627-1 : (parentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex2627-1)+parentFiledJointReturnWithCurrentSpouseCHVFlagsLength2627]), // Field # 500

		ParentTaxReturnFilingStatusCHVFlags: preprocessString2627(record[parentTaxReturnFilingStatusCHVFlagsStartIndex2627-1 : (parentTaxReturnFilingStatusCHVFlagsStartIndex2627-1)+parentTaxReturnFilingStatusCHVFlagsLength2627]), // Field # 501

		ParentIncomeEarnedFromWorkCHVFlags: preprocessString2627(record[parentIncomeEarnedFromWorkCHVFlagsStartIndex2627-1 : (parentIncomeEarnedFromWorkCHVFlagsStartIndex2627-1)+parentIncomeEarnedFromWorkCHVFlagsLength2627]), // Field # 502

		ParentTaxExemptInterestIncomeCHVFlags: preprocessString2627(record[parentTaxExemptInterestIncomeCHVFlagsStartIndex2627-1 : (parentTaxExemptInterestIncomeCHVFlagsStartIndex2627-1)+parentTaxExemptInterestIncomeCHVFlagsLength2627]), // Field # 503

		ParentUntaxedPortionsOfIRADistributionsCHVFlags: preprocessString2627(record[parentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2627-1 : (parentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2627-1)+parentUntaxedPortionsOfIRADistributionsCHVFlagsLength2627]), // Field # 504

		ParentIRARolloverCHVFlags: preprocessString2627(record[parentIRARolloverCHVFlagsStartIndex2627-1 : (parentIRARolloverCHVFlagsStartIndex2627-1)+parentIRARolloverCHVFlagsLength2627]), // Field # 505

		ParentUntaxedPortionsOfPensionsCHVFlags: preprocessString2627(record[parentUntaxedPortionsOfPensionsCHVFlagsStartIndex2627-1 : (parentUntaxedPortionsOfPensionsCHVFlagsStartIndex2627-1)+parentUntaxedPortionsOfPensionsCHVFlagsLength2627]), // Field # 506

		ParentPensionRolloverCHVFlags: preprocessString2627(record[parentPensionRolloverCHVFlagsStartIndex2627-1 : (parentPensionRolloverCHVFlagsStartIndex2627-1)+parentPensionRolloverCHVFlagsLength2627]), // Field # 507

		ParentAdjustedGrossIncomeCHVFlags: preprocessString2627(record[parentAdjustedGrossIncomeCHVFlagsStartIndex2627-1 : (parentAdjustedGrossIncomeCHVFlagsStartIndex2627-1)+parentAdjustedGrossIncomeCHVFlagsLength2627]), // Field # 508

		ParentIncomeTaxPaidCHVFlags: preprocessString2627(record[parentIncomeTaxPaidCHVFlagsStartIndex2627-1 : (parentIncomeTaxPaidCHVFlagsStartIndex2627-1)+parentIncomeTaxPaidCHVFlagsLength2627]), // Field # 509

		ParentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlags: preprocessString2627(record[parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex2627-1 : (parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex2627-1)+parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsLength2627]), // Field # 510

		ParentDeductiblePaymentsToIRAKeoghOtherCHVFlags: preprocessString2627(record[parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2627-1 : (parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2627-1)+parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2627]), // Field # 511

		ParentEducationCreditsCHVFlags: preprocessString2627(record[parentEducationCreditsCHVFlagsStartIndex2627-1 : (parentEducationCreditsCHVFlagsStartIndex2627-1)+parentEducationCreditsCHVFlagsLength2627]), // Field # 512

		ParentFiledScheduleABDEFHCHVFlags: preprocessString2627(record[parentFiledScheduleABDEFHCHVFlagsStartIndex2627-1 : (parentFiledScheduleABDEFHCHVFlagsStartIndex2627-1)+parentFiledScheduleABDEFHCHVFlagsLength2627]), // Field # 513

		ParentScheduleCAmountCHVFlags: preprocessString2627(record[parentScheduleCAmountCHVFlagsStartIndex2627-1 : (parentScheduleCAmountCHVFlagsStartIndex2627-1)+parentScheduleCAmountCHVFlagsLength2627]), // Field # 514

		ParentCollegeGrantAndScholarshipAidCHVFlags: preprocessString2627(record[parentCollegeGrantAndScholarshipAidCHVFlagsStartIndex2627-1 : (parentCollegeGrantAndScholarshipAidCHVFlagsStartIndex2627-1)+parentCollegeGrantAndScholarshipAidCHVFlagsLength2627]), // Field # 515

		ParentForeignEarnedIncomeExclusionCHVFlags: preprocessString2627(record[parentForeignEarnedIncomeExclusionCHVFlagsStartIndex2627-1 : (parentForeignEarnedIncomeExclusionCHVFlagsStartIndex2627-1)+parentForeignEarnedIncomeExclusionCHVFlagsLength2627]), // Field # 516

		ParentChildSupportReceivedCHVFlags: preprocessString2627(record[parentChildSupportReceivedCHVFlagsStartIndex2627-1 : (parentChildSupportReceivedCHVFlagsStartIndex2627-1)+parentChildSupportReceivedCHVFlagsLength2627]), // Field # 517

		ParentNetWorthOfCurrentInvestmentsCHVFlags: preprocessString2627(record[parentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex2627-1 : (parentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex2627-1)+parentNetWorthOfCurrentInvestmentsCHVFlagsLength2627]), // Field # 518

		ParentTotalOfCashSavingsAndCheckingAccountsCHVFlags: preprocessString2627(record[parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsStartIndex2627-1 : (parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsStartIndex2627-1)+parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsLength2627]), // Field # 519

		ParentNetWorthOfBusinessesAndInvestmentFarmsCHVFlags: preprocessString2627(record[parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex2627-1 : (parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex2627-1)+parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsLength2627]), // Field # 520

		ParentConsentToRetrieveAndDiscloseFTICHVFlags: preprocessString2627(record[parentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2627-1 : (parentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2627-1)+parentConsentToRetrieveAndDiscloseFTICHVFlagsLength2627]), // Field # 521

		ParentSignatureCHVFlags: preprocessString2627(record[parentSignatureCHVFlagsStartIndex2627-1 : (parentSignatureCHVFlagsStartIndex2627-1)+parentSignatureCHVFlagsLength2627]), // Field # 522

		ParentSignatureDateCHVFlags: preprocessString2627(record[parentSignatureDateCHVFlagsStartIndex2627-1 : (parentSignatureDateCHVFlagsStartIndex2627-1)+parentSignatureDateCHVFlagsLength2627]), // Field # 523

		ParentSpouseFirstNameCHVFlags: preprocessString2627(record[parentSpouseFirstNameCHVFlagsStartIndex2627-1 : (parentSpouseFirstNameCHVFlagsStartIndex2627-1)+parentSpouseFirstNameCHVFlagsLength2627]), // Field # 524

		ParentSpouseMiddleNameCHVFlags: preprocessString2627(record[parentSpouseMiddleNameCHVFlagsStartIndex2627-1 : (parentSpouseMiddleNameCHVFlagsStartIndex2627-1)+parentSpouseMiddleNameCHVFlagsLength2627]), // Field # 525

		ParentSpouseLastNameCHVFlags: preprocessString2627(record[parentSpouseLastNameCHVFlagsStartIndex2627-1 : (parentSpouseLastNameCHVFlagsStartIndex2627-1)+parentSpouseLastNameCHVFlagsLength2627]), // Field # 526

		ParentSpouseSuffixCHVFlags: preprocessString2627(record[parentSpouseSuffixCHVFlagsStartIndex2627-1 : (parentSpouseSuffixCHVFlagsStartIndex2627-1)+parentSpouseSuffixCHVFlagsLength2627]), // Field # 527

		ParentSpouseDateOfBirthCHVFlags: preprocessString2627(record[parentSpouseDateOfBirthCHVFlagsStartIndex2627-1 : (parentSpouseDateOfBirthCHVFlagsStartIndex2627-1)+parentSpouseDateOfBirthCHVFlagsLength2627]), // Field # 528

		ParentSpouseSSNCHVFlags: preprocessString2627(record[parentSpouseSSNCHVFlagsStartIndex2627-1 : (parentSpouseSSNCHVFlagsStartIndex2627-1)+parentSpouseSSNCHVFlagsLength2627]), // Field # 529

		ParentSpouseITINCHVFlags: preprocessString2627(record[parentSpouseITINCHVFlagsStartIndex2627-1 : (parentSpouseITINCHVFlagsStartIndex2627-1)+parentSpouseITINCHVFlagsLength2627]), // Field # 530

		ParentSpousePhoneNumberCHVFlags: preprocessString2627(record[parentSpousePhoneNumberCHVFlagsStartIndex2627-1 : (parentSpousePhoneNumberCHVFlagsStartIndex2627-1)+parentSpousePhoneNumberCHVFlagsLength2627]), // Field # 531

		ParentSpouseEmailAddressCHVFlags: preprocessString2627(record[parentSpouseEmailAddressCHVFlagsStartIndex2627-1 : (parentSpouseEmailAddressCHVFlagsStartIndex2627-1)+parentSpouseEmailAddressCHVFlagsLength2627]), // Field # 532

		ParentSpouseStreetAddressCHVFlags: preprocessString2627(record[parentSpouseStreetAddressCHVFlagsStartIndex2627-1 : (parentSpouseStreetAddressCHVFlagsStartIndex2627-1)+parentSpouseStreetAddressCHVFlagsLength2627]), // Field # 533

		ParentSpouseCityCHVFlags: preprocessString2627(record[parentSpouseCityCHVFlagsStartIndex2627-1 : (parentSpouseCityCHVFlagsStartIndex2627-1)+parentSpouseCityCHVFlagsLength2627]), // Field # 534

		ParentSpouseStateCHVFlags: preprocessString2627(record[parentSpouseStateCHVFlagsStartIndex2627-1 : (parentSpouseStateCHVFlagsStartIndex2627-1)+parentSpouseStateCHVFlagsLength2627]), // Field # 535

		ParentSpouseZipCodeCHVFlags: preprocessString2627(record[parentSpouseZipCodeCHVFlagsStartIndex2627-1 : (parentSpouseZipCodeCHVFlagsStartIndex2627-1)+parentSpouseZipCodeCHVFlagsLength2627]), // Field # 536

		ParentSpouseCountryCHVFlags: preprocessString2627(record[parentSpouseCountryCHVFlagsStartIndex2627-1 : (parentSpouseCountryCHVFlagsStartIndex2627-1)+parentSpouseCountryCHVFlagsLength2627]), // Field # 537

		ParentSpouseFiled1040Or1040NRCHVFlags: preprocessString2627(record[parentSpouseFiled1040Or1040NRCHVFlagsStartIndex2627-1 : (parentSpouseFiled1040Or1040NRCHVFlagsStartIndex2627-1)+parentSpouseFiled1040Or1040NRCHVFlagsLength2627]), // Field # 538

		ParentSpouseFileNonUSTaxReturnCHVFlags: preprocessString2627(record[parentSpouseFileNonUSTaxReturnCHVFlagsStartIndex2627-1 : (parentSpouseFileNonUSTaxReturnCHVFlagsStartIndex2627-1)+parentSpouseFileNonUSTaxReturnCHVFlagsLength2627]), // Field # 539

		ParentSpouseTaxReturnFilingStatusCHVFlags: preprocessString2627(record[parentSpouseTaxReturnFilingStatusCHVFlagsStartIndex2627-1 : (parentSpouseTaxReturnFilingStatusCHVFlagsStartIndex2627-1)+parentSpouseTaxReturnFilingStatusCHVFlagsLength2627]), // Field # 540

		ParentSpouseIncomeEarnedFromWorkCHVFlags: preprocessString2627(record[parentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex2627-1 : (parentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex2627-1)+parentSpouseIncomeEarnedFromWorkCHVFlagsLength2627]), // Field # 541

		ParentSpouseTaxExemptInterestIncomeCHVFlags: preprocessString2627(record[parentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex2627-1 : (parentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex2627-1)+parentSpouseTaxExemptInterestIncomeCHVFlagsLength2627]), // Field # 542

		ParentSpouseUntaxedPortionsOfIRADistributionsCHVFlags: preprocessString2627(record[parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2627-1 : (parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2627-1)+parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsLength2627]), // Field # 543

		ParentSpouseIRARolloverCHVFlags: preprocessString2627(record[parentSpouseIRARolloverCHVFlagsStartIndex2627-1 : (parentSpouseIRARolloverCHVFlagsStartIndex2627-1)+parentSpouseIRARolloverCHVFlagsLength2627]), // Field # 544

		ParentSpouseUntaxedPortionsOfPensionsCHVFlags: preprocessString2627(record[parentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex2627-1 : (parentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex2627-1)+parentSpouseUntaxedPortionsOfPensionsCHVFlagsLength2627]), // Field # 545

		ParentSpousePensionRolloverCHVFlags: preprocessString2627(record[parentSpousePensionRolloverCHVFlagsStartIndex2627-1 : (parentSpousePensionRolloverCHVFlagsStartIndex2627-1)+parentSpousePensionRolloverCHVFlagsLength2627]), // Field # 546

		ParentSpouseAdjustedGrossIncomeCHVFlags: preprocessString2627(record[parentSpouseAdjustedGrossIncomeCHVFlagsStartIndex2627-1 : (parentSpouseAdjustedGrossIncomeCHVFlagsStartIndex2627-1)+parentSpouseAdjustedGrossIncomeCHVFlagsLength2627]), // Field # 547

		ParentSpouseIncomeTaxPaidCHVFlags: preprocessString2627(record[parentSpouseIncomeTaxPaidCHVFlagsStartIndex2627-1 : (parentSpouseIncomeTaxPaidCHVFlagsStartIndex2627-1)+parentSpouseIncomeTaxPaidCHVFlagsLength2627]), // Field # 548

		ParentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlags: preprocessString2627(record[parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2627-1 : (parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2627-1)+parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2627]), // Field # 549

		ParentSpouseEducationCreditsCHVFlags: preprocessString2627(record[parentSpouseEducationCreditsCHVFlagsStartIndex2627-1 : (parentSpouseEducationCreditsCHVFlagsStartIndex2627-1)+parentSpouseEducationCreditsCHVFlagsLength2627]), // Field # 550

		ParentSpouseFiledScheduleABDEFHCHVFlags: preprocessString2627(record[parentSpouseFiledScheduleABDEFHCHVFlagsStartIndex2627-1 : (parentSpouseFiledScheduleABDEFHCHVFlagsStartIndex2627-1)+parentSpouseFiledScheduleABDEFHCHVFlagsLength2627]), // Field # 551

		ParentSpouseScheduleCAmountCHVFlags: preprocessString2627(record[parentSpouseScheduleCAmountCHVFlagsStartIndex2627-1 : (parentSpouseScheduleCAmountCHVFlagsStartIndex2627-1)+parentSpouseScheduleCAmountCHVFlagsLength2627]), // Field # 552

		ParentSpouseForeignEarnedIncomeExclusionCHVFlags: preprocessString2627(record[parentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex2627-1 : (parentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex2627-1)+parentSpouseForeignEarnedIncomeExclusionCHVFlagsLength2627]), // Field # 553

		ParentSpouseConsentToRetrieveAndDiscloseFTICHVFlags: preprocessString2627(record[parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2627-1 : (parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2627-1)+parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsLength2627]), // Field # 554

		ParentSpouseSignatureCHVFlags: preprocessString2627(record[parentSpouseSignatureCHVFlagsStartIndex2627-1 : (parentSpouseSignatureCHVFlagsStartIndex2627-1)+parentSpouseSignatureCHVFlagsLength2627]), // Field # 555

		ParentSpouseSignatureDateCHVFlags: preprocessString2627(record[parentSpouseSignatureDateCHVFlagsStartIndex2627-1 : (parentSpouseSignatureDateCHVFlagsStartIndex2627-1)+parentSpouseSignatureDateCHVFlagsLength2627]), // Field # 556

		DHSPrimaryMatchStatus: preprocessString2627(record[dHSPrimaryMatchStatusStartIndex2627-1 : (dHSPrimaryMatchStatusStartIndex2627-1)+dHSPrimaryMatchStatusLength2627]), // Field # 557

		DHSCaseNumber: preprocessString2627(record[dHSCaseNumberStartIndex2627-1 : (dHSCaseNumberStartIndex2627-1)+dHSCaseNumberLength2627]), // Field # 559

		NSLDSMatchStatus: preprocessString2627(record[nsldsMatchStatusStartIndex2627-1 : (nsldsMatchStatusStartIndex2627-1)+nsldsMatchStatusLength2627]), // Field # 560

		NSLDSPostscreeningReasonCode: preprocessString2627(record[nsldsPostscreeningReasonCodeStartIndex2627-1 : (nsldsPostscreeningReasonCodeStartIndex2627-1)+nsldsPostscreeningReasonCodeLength2627]), // Field # 561

		StudentSSACitizenshipFlagResults: preprocessString2627(record[studentSSACitizenshipFlagResultsStartIndex2627-1 : (studentSSACitizenshipFlagResultsStartIndex2627-1)+studentSSACitizenshipFlagResultsLength2627]), // Field # 562

		StudentSSAMatchStatus: preprocessString2627(record[studentSSAMatchStatusStartIndex2627-1 : (studentSSAMatchStatusStartIndex2627-1)+studentSSAMatchStatusLength2627]), // Field # 563

		StudentSpouseSSAMatchStatus: preprocessString2627(record[studentSpouseSSAMatchStatusStartIndex2627-1 : (studentSpouseSSAMatchStatusStartIndex2627-1)+studentSpouseSSAMatchStatusLength2627]), // Field # 564

		ParentSSAMatchStatus: preprocessString2627(record[parentSSAMatchStatusStartIndex2627-1 : (parentSSAMatchStatusStartIndex2627-1)+parentSSAMatchStatusLength2627]), // Field # 565

		ParentSpouseOrPartnerSSAMatchStatus: preprocessString2627(record[parentSpouseOrPartnerSSAMatchStatusStartIndex2627-1 : (parentSpouseOrPartnerSSAMatchStatusStartIndex2627-1)+parentSpouseOrPartnerSSAMatchStatusLength2627]), // Field # 566

		VAMatchFlag: preprocessString2627(record[vAMatchFlagStartIndex2627-1 : (vAMatchFlagStartIndex2627-1)+vAMatchFlagLength2627]), // Field # 567

		CommentCodes: preprocessString2627(record[commentCodesStartIndex2627-1 : (commentCodesStartIndex2627-1)+commentCodesLength2627]), // Field # 568

		DrugAbuseHoldIndicator: preprocessString2627(record[drugAbuseHoldIndicatorStartIndex2627-1 : (drugAbuseHoldIndicatorStartIndex2627-1)+drugAbuseHoldIndicatorLength2627]), // Field # 569

		GraduateFlag: preprocessString2627(record[graduateFlagStartIndex2627-1 : (graduateFlagStartIndex2627-1)+graduateFlagLength2627]), // Field # 570

		PellGrantEligibilityFlag: preprocessString2627(record[pellGrantEligibilityFlagStartIndex2627-1 : (pellGrantEligibilityFlagStartIndex2627-1)+pellGrantEligibilityFlagLength2627]), // Field # 571

		ReprocessedReasonCode: preprocessString2627(record[reprocessedReasonCodeStartIndex2627-1 : (reprocessedReasonCodeStartIndex2627-1)+reprocessedReasonCodeLength2627]), // Field # 572

		FPSCFlag: preprocessString2627(record[fpsCFlagStartIndex2627-1 : (fpsCFlagStartIndex2627-1)+fpsCFlagLength2627]), // Field # 573

		FPSCChangeFlag: preprocessString2627(record[fpsCChangeFlagStartIndex2627-1 : (fpsCChangeFlagStartIndex2627-1)+fpsCChangeFlagLength2627]), // Field # 574

		ElectronicFederalSchoolCodeIndicator: preprocessString2627(record[electronicFederalSchoolCodeIndicatorStartIndex2627-1 : (electronicFederalSchoolCodeIndicatorStartIndex2627-1)+electronicFederalSchoolCodeIndicatorLength2627]), // Field # 575

		RejectReasonCodes: preprocessString2627(record[rejectReasonCodesStartIndex2627-1 : (rejectReasonCodesStartIndex2627-1)+rejectReasonCodesLength2627]), // Field # 576

		ElectronicTransactionIndicatorFlag: preprocessString2627(record[electronicTransactionIndicatorFlagStartIndex2627-1 : (electronicTransactionIndicatorFlagStartIndex2627-1)+electronicTransactionIndicatorFlagLength2627]), // Field # 577

		StudentLastNameSSNChangeFlag: preprocessString2627(record[studentLastNameSSNChangeFlagStartIndex2627-1 : (studentLastNameSSNChangeFlagStartIndex2627-1)+studentLastNameSSNChangeFlagLength2627]), // Field # 578

		HighSchoolCode: preprocessString2627(record[highSchoolCodeStartIndex2627-1 : (highSchoolCodeStartIndex2627-1)+highSchoolCodeLength2627]), // Field # 579

		VerificationSelectionChangeFlag: preprocessString2627(record[verificationSelectionChangeFlagStartIndex2627-1 : (verificationSelectionChangeFlagStartIndex2627-1)+verificationSelectionChangeFlagLength2627]), // Field # 580

		//UseUserProvidedDataOnly: preprocessString2627(record[useUserProvidedDataOnlyStartIndex2627-1 : (useUserProvidedDataOnlyStartIndex2627-1)+useUserProvidedDataOnlyLength2627]), // Field # 581

		NSLDSPellOverpaymentFlag: preprocessString2627(record[nsldsPellOverpaymentFlagStartIndex2627-1 : (nsldsPellOverpaymentFlagStartIndex2627-1)+nsldsPellOverpaymentFlagLength2627]), // Field # 583

		NSLDSPellOverpaymentContact: preprocessString2627(record[nsldsPellOverpaymentContactStartIndex2627-1 : (nsldsPellOverpaymentContactStartIndex2627-1)+nsldsPellOverpaymentContactLength2627]), // Field # 584

		NSLDSFSEOGOverpaymentFlag: preprocessString2627(record[nsldsFSEOGOverpaymentFlagStartIndex2627-1 : (nsldsFSEOGOverpaymentFlagStartIndex2627-1)+nsldsFSEOGOverpaymentFlagLength2627]), // Field # 585

		NSLDSFSEOGOverpaymentContact: preprocessString2627(record[nsldsFSEOGOverpaymentContactStartIndex2627-1 : (nsldsFSEOGOverpaymentContactStartIndex2627-1)+nsldsFSEOGOverpaymentContactLength2627]), // Field # 586

		NSLDSPerkinsOverpaymentFlag: preprocessString2627(record[nsldsPerkinsOverpaymentFlagStartIndex2627-1 : (nsldsPerkinsOverpaymentFlagStartIndex2627-1)+nsldsPerkinsOverpaymentFlagLength2627]), // Field # 587

		NSLDSPerkinsOverpaymentContact: preprocessString2627(record[nsldsPerkinsOverpaymentContactStartIndex2627-1 : (nsldsPerkinsOverpaymentContactStartIndex2627-1)+nsldsPerkinsOverpaymentContactLength2627]), // Field # 588

		NSLDSTEACHGrantOverpaymentFlag: preprocessString2627(record[nsldsTEACHGrantOverpaymentFlagStartIndex2627-1 : (nsldsTEACHGrantOverpaymentFlagStartIndex2627-1)+nsldsTEACHGrantOverpaymentFlagLength2627]), // Field # 589

		NSLDSTEACHGrantOverpaymentContact: preprocessString2627(record[nsldsTEACHGrantOverpaymentContactStartIndex2627-1 : (nsldsTEACHGrantOverpaymentContactStartIndex2627-1)+nsldsTEACHGrantOverpaymentContactLength2627]), // Field # 590

		NSLDSIraqAndAfghanistanServiceGrantOverpaymentFlag: preprocessString2627(record[nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagStartIndex2627-1 : (nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagStartIndex2627-1)+nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagLength2627]), // Field # 591

		NSLDSIraqAndAfghanistanServiceGrantOverpaymentContact: preprocessString2627(record[nsldsIraqAndAfghanistanServiceGrantOverpaymentContactStartIndex2627-1 : (nsldsIraqAndAfghanistanServiceGrantOverpaymentContactStartIndex2627-1)+nsldsIraqAndAfghanistanServiceGrantOverpaymentContactLength2627]), // Field # 592

		NSLDSDefaultedLoanFlag: preprocessString2627(record[nsldsDefaultedLoanFlagStartIndex2627-1 : (nsldsDefaultedLoanFlagStartIndex2627-1)+nsldsDefaultedLoanFlagLength2627]), // Field # 593

		NSLDSDischargedLoanFlag: preprocessString2627(record[nsldsDischargedLoanFlagStartIndex2627-1 : (nsldsDischargedLoanFlagStartIndex2627-1)+nsldsDischargedLoanFlagLength2627]), // Field # 594

		NSLDSFraudLoanFlag: preprocessString2627(record[nsldsFraudLoanFlagStartIndex2627-1 : (nsldsFraudLoanFlagStartIndex2627-1)+nsldsFraudLoanFlagLength2627]), // Field # 595

		NSLDSSatisfactoryArrangementsFlag: preprocessString2627(record[nsldsSatisfactoryArrangementsFlagStartIndex2627-1 : (nsldsSatisfactoryArrangementsFlagStartIndex2627-1)+nsldsSatisfactoryArrangementsFlagLength2627]), // Field # 596

		NSLDSActiveBankruptcyFlag: preprocessString2627(record[nsldsActiveBankruptcyFlagStartIndex2627-1 : (nsldsActiveBankruptcyFlagStartIndex2627-1)+nsldsActiveBankruptcyFlagLength2627]), // Field # 597

		NSLDSTEACHGrantConvertedToLoanFlag: preprocessString2627(record[nsldsTEACHGrantConvertedToLoanFlagStartIndex2627-1 : (nsldsTEACHGrantConvertedToLoanFlagStartIndex2627-1)+nsldsTEACHGrantConvertedToLoanFlagLength2627]), // Field # 598

		NSLDSAggregateSubsidizedOutstandingPrincipalBalance: preprocessString2627(record[nsldsAggregateSubsidizedOutstandingPrincipalBalanceStartIndex2627-1 : (nsldsAggregateSubsidizedOutstandingPrincipalBalanceStartIndex2627-1)+nsldsAggregateSubsidizedOutstandingPrincipalBalanceLength2627]), // Field # 599

		NSLDSAggregateUnsubsidizedOutstandingPrincipalBalance: preprocessString2627(record[nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceStartIndex2627-1 : (nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceStartIndex2627-1)+nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceLength2627]), // Field # 600

		NSLDSAggregateCombinedOutstandingPrincipalBalance: preprocessString2627(record[nsldsAggregateCombinedOutstandingPrincipalBalanceStartIndex2627-1 : (nsldsAggregateCombinedOutstandingPrincipalBalanceStartIndex2627-1)+nsldsAggregateCombinedOutstandingPrincipalBalanceLength2627]), // Field # 601

		NSLDSAggregateUnallocConsolidatedOutstandingPrincipalBalance: preprocessString2627(record[nsldsAggregateUnallocConsolidatedOutstandingPrincipalBalanceStartIndex2627-1 : (nsldsAggregateUnallocConsolidatedOutstandingPrincipalBalanceStartIndex2627-1)+nsldsAggregateUnallocConsolidatedOutstandingPrincipalBalanceLength2627]), // Field # 602

		NSLDSAggregateTEACHLoanPrincipalBalance: preprocessString2627(record[nsldsAggregateTEACHLoanPrincipalBalanceStartIndex2627-1 : (nsldsAggregateTEACHLoanPrincipalBalanceStartIndex2627-1)+nsldsAggregateTEACHLoanPrincipalBalanceLength2627]), // Field # 603

		NSLDSAggregateSubsidizedPendingDisbursement: preprocessString2627(record[nsldsAggregateSubsidizedPendingDisbursementStartIndex2627-1 : (nsldsAggregateSubsidizedPendingDisbursementStartIndex2627-1)+nsldsAggregateSubsidizedPendingDisbursementLength2627]), // Field # 604

		NSLDSAggregateUnsubsidizedPendingDisbursement: preprocessString2627(record[nsldsAggregateUnsubsidizedPendingDisbursementStartIndex2627-1 : (nsldsAggregateUnsubsidizedPendingDisbursementStartIndex2627-1)+nsldsAggregateUnsubsidizedPendingDisbursementLength2627]), // Field # 605

		NSLDSAggregateCombinedPendingDisbursement: preprocessString2627(record[nsldsAggregateCombinedPendingDisbursementStartIndex2627-1 : (nsldsAggregateCombinedPendingDisbursementStartIndex2627-1)+nsldsAggregateCombinedPendingDisbursementLength2627]), // Field # 606

		NSLDSAggregateSubsidizedTotal: preprocessString2627(record[nsldsAggregateSubsidizedTotalStartIndex2627-1 : (nsldsAggregateSubsidizedTotalStartIndex2627-1)+nsldsAggregateSubsidizedTotalLength2627]), // Field # 607

		NSLDSAggregateUnsubsidizedTotal: preprocessString2627(record[nsldsAggregateUnsubsidizedTotalStartIndex2627-1 : (nsldsAggregateUnsubsidizedTotalStartIndex2627-1)+nsldsAggregateUnsubsidizedTotalLength2627]), // Field # 608

		NSLDSAggregateCombinedTotal: preprocessString2627(record[nsldsAggregateCombinedTotalStartIndex2627-1 : (nsldsAggregateCombinedTotalStartIndex2627-1)+nsldsAggregateCombinedTotalLength2627]), // Field # 609

		NSLDSUnallocatedConsolidatedTotal: preprocessString2627(record[nsldsUnallocatedConsolidatedTotalStartIndex2627-1 : (nsldsUnallocatedConsolidatedTotalStartIndex2627-1)+nsldsUnallocatedConsolidatedTotalLength2627]), // Field # 610

		NSLDSTEACHLoanTotal: preprocessString2627(record[nsldsTEACHLoanTotalStartIndex2627-1 : (nsldsTEACHLoanTotalStartIndex2627-1)+nsldsTEACHLoanTotalLength2627]), // Field # 611

		NSLDSPerkinsTotalDisbursements: preprocessString2627(record[nsldsPerkinsTotalDisbursementsStartIndex2627-1 : (nsldsPerkinsTotalDisbursementsStartIndex2627-1)+nsldsPerkinsTotalDisbursementsLength2627]), // Field # 612

		NSLDSPerkinsCurrentYearDisbursementAmount: preprocessString2627(record[nsldsPerkinsCurrentYearDisbursementAmountStartIndex2627-1 : (nsldsPerkinsCurrentYearDisbursementAmountStartIndex2627-1)+nsldsPerkinsCurrentYearDisbursementAmountLength2627]), // Field # 613

		NSLDSAggregateTEACHGrantUndergraduateDisbursedTotal: preprocessString2627(record[nsldsAggregateTEACHGrantUndergraduateDisbursedTotalStartIndex2627-1 : (nsldsAggregateTEACHGrantUndergraduateDisbursedTotalStartIndex2627-1)+nsldsAggregateTEACHGrantUndergraduateDisbursedTotalLength2627]), // Field # 614

		NSLDSAggregateTEACHGraduateDisbursementAmount: preprocessString2627(record[nsldsAggregateTEACHGraduateDisbursementAmountStartIndex2627-1 : (nsldsAggregateTEACHGraduateDisbursementAmountStartIndex2627-1)+nsldsAggregateTEACHGraduateDisbursementAmountLength2627]), // Field # 615

		NSLDSDefaultedLoanChangeFlag: preprocessString2627(record[nsldsDefaultedLoanChangeFlagStartIndex2627-1 : (nsldsDefaultedLoanChangeFlagStartIndex2627-1)+nsldsDefaultedLoanChangeFlagLength2627]), // Field # 616

		NSLDSFraudLoanChangeFlag: preprocessString2627(record[nsldsFraudLoanChangeFlagStartIndex2627-1 : (nsldsFraudLoanChangeFlagStartIndex2627-1)+nsldsFraudLoanChangeFlagLength2627]), // Field # 617

		NSLDSDischargedLoanChangeFlag: preprocessString2627(record[nsldsDischargedLoanChangeFlagStartIndex2627-1 : (nsldsDischargedLoanChangeFlagStartIndex2627-1)+nsldsDischargedLoanChangeFlagLength2627]), // Field # 618

		NSLDSLoanSatisfactoryRepaymentChangeFlag: preprocessString2627(record[nsldsLoanSatisfactoryRepaymentChangeFlagStartIndex2627-1 : (nsldsLoanSatisfactoryRepaymentChangeFlagStartIndex2627-1)+nsldsLoanSatisfactoryRepaymentChangeFlagLength2627]), // Field # 619

		NSLDSActiveBankruptcyChangeFlag: preprocessString2627(record[nsldsActiveBankruptcyChangeFlagStartIndex2627-1 : (nsldsActiveBankruptcyChangeFlagStartIndex2627-1)+nsldsActiveBankruptcyChangeFlagLength2627]), // Field # 620

		NSLDSTEACHGrantToLoanConversionChangeFlag: preprocessString2627(record[nsldsTEACHGrantToLoanConversionChangeFlagStartIndex2627-1 : (nsldsTEACHGrantToLoanConversionChangeFlagStartIndex2627-1)+nsldsTEACHGrantToLoanConversionChangeFlagLength2627]), // Field # 621

		NSLDSOverpaymentsChangeFlag: preprocessString2627(record[nsldsOverpaymentsChangeFlagStartIndex2627-1 : (nsldsOverpaymentsChangeFlagStartIndex2627-1)+nsldsOverpaymentsChangeFlagLength2627]), // Field # 622

		NSLDSAggregateLoanChangeFlag: preprocessString2627(record[nsldsAggregateLoanChangeFlagStartIndex2627-1 : (nsldsAggregateLoanChangeFlagStartIndex2627-1)+nsldsAggregateLoanChangeFlagLength2627]), // Field # 623

		NSLDSPerkinsLoanChangeFlag: preprocessString2627(record[nsldsPerkinsLoanChangeFlagStartIndex2627-1 : (nsldsPerkinsLoanChangeFlagStartIndex2627-1)+nsldsPerkinsLoanChangeFlagLength2627]), // Field # 624

		NSLDSPellPaymentChangeFlag: preprocessString2627(record[nsldsPellPaymentChangeFlagStartIndex2627-1 : (nsldsPellPaymentChangeFlagStartIndex2627-1)+nsldsPellPaymentChangeFlagLength2627]), // Field # 625

		NSLDSTEACHGrantChangeFlag: preprocessString2627(record[nsldsTEACHGrantChangeFlagStartIndex2627-1 : (nsldsTEACHGrantChangeFlagStartIndex2627-1)+nsldsTEACHGrantChangeFlagLength2627]), // Field # 626

		NSLDSAdditionalPellFlag: preprocessString2627(record[nsldsAdditionalPellFlagStartIndex2627-1 : (nsldsAdditionalPellFlagStartIndex2627-1)+nsldsAdditionalPellFlagLength2627]), // Field # 627

		NSLDSAdditionalLoansFlag: preprocessString2627(record[nsldsAdditionalLoansFlagStartIndex2627-1 : (nsldsAdditionalLoansFlagStartIndex2627-1)+nsldsAdditionalLoansFlagLength2627]), // Field # 628

		NSLDSAdditionalTEACHGrantFlag: preprocessString2627(record[nsldsAdditionalTEACHGrantFlagStartIndex2627-1 : (nsldsAdditionalTEACHGrantFlagStartIndex2627-1)+nsldsAdditionalTEACHGrantFlagLength2627]), // Field # 629

		NSLDSDirectLoanMasterPromNoteFlag: preprocessString2627(record[nsldsDirectLoanMasterPromNoteFlagStartIndex2627-1 : (nsldsDirectLoanMasterPromNoteFlagStartIndex2627-1)+nsldsDirectLoanMasterPromNoteFlagLength2627]), // Field # 630

		NSLDSDirectLoanPLUSMasterPromNoteFlag: preprocessString2627(record[nsldsDirectLoanPLUSMasterPromNoteFlagStartIndex2627-1 : (nsldsDirectLoanPLUSMasterPromNoteFlagStartIndex2627-1)+nsldsDirectLoanPLUSMasterPromNoteFlagLength2627]), // Field # 631

		NSLDSDirectLoanGraduatePLUSMasterPromNoteFlag: preprocessString2627(record[nsldsDirectLoanGraduatePLUSMasterPromNoteFlagStartIndex2627-1 : (nsldsDirectLoanGraduatePLUSMasterPromNoteFlagStartIndex2627-1)+nsldsDirectLoanGraduatePLUSMasterPromNoteFlagLength2627]), // Field # 632

		NSLDSUndergraduateSubsidizedLoanLimitFlag: preprocessString2627(record[nsldsUndergraduateSubsidizedLoanLimitFlagStartIndex2627-1 : (nsldsUndergraduateSubsidizedLoanLimitFlagStartIndex2627-1)+nsldsUndergraduateSubsidizedLoanLimitFlagLength2627]), // Field # 633

		NSLDSUndergraduateCombinedLoanLimitFlag: preprocessString2627(record[nsldsUndergraduateCombinedLoanLimitFlagStartIndex2627-1 : (nsldsUndergraduateCombinedLoanLimitFlagStartIndex2627-1)+nsldsUndergraduateCombinedLoanLimitFlagLength2627]), // Field # 634

		NSLDSGraduateSubsidizedLoanLimitFlag: preprocessString2627(record[nsldsGraduateSubsidizedLoanLimitFlagStartIndex2627-1 : (nsldsGraduateSubsidizedLoanLimitFlagStartIndex2627-1)+nsldsGraduateSubsidizedLoanLimitFlagLength2627]), // Field # 635

		NSLDSGraduateCombinedLoanLimitFlag: preprocessString2627(record[nsldsGraduateCombinedLoanLimitFlagStartIndex2627-1 : (nsldsGraduateCombinedLoanLimitFlagStartIndex2627-1)+nsldsGraduateCombinedLoanLimitFlagLength2627]), // Field # 636

		NSLDSPellLifetimeLimitFlag: preprocessString2627(record[nsldsPellLifetimeLimitFlagStartIndex2627-1 : (nsldsPellLifetimeLimitFlagStartIndex2627-1)+nsldsPellLifetimeLimitFlagLength2627]), // Field # 637

		NSLDSPellLifetimeEligibilityUsed: preprocessString2627(record[nsldsPellLifetimeEligibilityUsedStartIndex2627-1 : (nsldsPellLifetimeEligibilityUsedStartIndex2627-1)+nsldsPellLifetimeEligibilityUsedLength2627]), // Field # 638

		NSLDSSULAFlag: preprocessString2627(record[nsldsSULAFlagStartIndex2627-1 : (nsldsSULAFlagStartIndex2627-1)+nsldsSULAFlagLength2627]), // Field # 639

		NSLDSSubsidizedLimitEligibilityFlag: preprocessString2627(record[nsldsSubsidizedLimitEligibilityUsedStartIndex2627-1 : (nsldsSubsidizedLimitEligibilityUsedStartIndex2627-1)+nsldsSubsidizedLimitEligibilityUsedLength2627]), // Field # 640

		NSLDSUnusualEnrollmentHistoryFlag: preprocessString2627(record[nsldsUnusualEnrollmentHistoryFlagStartIndex2627-1 : (nsldsUnusualEnrollmentHistoryFlagStartIndex2627-1)+nsldsUnusualEnrollmentHistoryFlagLength2627]), // Field # 641

		NSLDSPellSequenceNumber1: preprocessString2627(record[nsldsPellSequenceNumber1StartIndex2627-1 : (nsldsPellSequenceNumber1StartIndex2627-1)+nsldsPellSequenceNumber1Length2627]), // Field # 643

		NSLDSPellVerificationFlag1: preprocessString2627(record[nsldsPellVerificationFlag1StartIndex2627-1 : (nsldsPellVerificationFlag1StartIndex2627-1)+nsldsPellVerificationFlag1Length2627]), // Field # 644

		NSLDSSAI1: preprocessString2627(record[nsldsSAI1StartIndex2627-1 : (nsldsSAI1StartIndex2627-1)+nsldsSAI1Length2627]), // Field # 645

		NSLDSPellSchoolCode1: preprocessString2627(record[nsldsPellSchoolCode1StartIndex2627-1 : (nsldsPellSchoolCode1StartIndex2627-1)+nsldsPellSchoolCode1Length2627]), // Field # 646

		NSLDSPellTransactionNumber1: preprocessString2627(record[nsldsPellTransactionNumber1StartIndex2627-1 : (nsldsPellTransactionNumber1StartIndex2627-1)+nsldsPellTransactionNumber1Length2627]), // Field # 647

		NSLDSPellDisbursementDate1: parseISIRDate2627(preprocessString2627(record[nsldsPellDisbursementDate1StartIndex2627-1 : (nsldsPellDisbursementDate1StartIndex2627-1)+nsldsPellDisbursementDate1Length2627])), // Field # 648

		NSLDSPellScheduledAmount1: preprocessString2627(record[nsldsPellScheduledAmount1StartIndex2627-1 : (nsldsPellScheduledAmount1StartIndex2627-1)+nsldsPellScheduledAmount1Length2627]), // Field # 649

		NSLDSPellAmountPaidToDate1: parseISIRDate2627(preprocessString2627(record[nsldsPellAmountPaidToDate1StartIndex2627-1 : (nsldsPellAmountPaidToDate1StartIndex2627-1)+nsldsPellAmountPaidToDate1Length2627])), // Field # 650

		NSLDSPellPercentEligibilityUsedDecimal1: preprocessString2627(record[nsldsPellPercentEligibilityUsedDecimal1StartIndex2627-1 : (nsldsPellPercentEligibilityUsedDecimal1StartIndex2627-1)+nsldsPellPercentEligibilityUsedDecimal1Length2627]), // Field # 651

		NSLDSPellAwardAmount1: preprocessString2627(record[nsldsPellAwardAmount1StartIndex2627-1 : (nsldsPellAwardAmount1StartIndex2627-1)+nsldsPellAwardAmount1Length2627]), // Field # 652

		NSLDSAdditionalEligibilityIndicator1: preprocessString2627(record[nsldsAdditionalEligibilityIndicator1StartIndex2627-1 : (nsldsAdditionalEligibilityIndicator1StartIndex2627-1)+nsldsAdditionalEligibilityIndicator1Length2627]), // Field # 653

		NSLDSPellSequenceNumber2: preprocessString2627(record[nsldsPellSequenceNumber2StartIndex2627-1 : (nsldsPellSequenceNumber2StartIndex2627-1)+nsldsPellSequenceNumber2Length2627]), // Field # 655

		NSLDSPellVerificationFlag2: preprocessString2627(record[nsldsPellVerificationFlag2StartIndex2627-1 : (nsldsPellVerificationFlag2StartIndex2627-1)+nsldsPellVerificationFlag2Length2627]), // Field # 656

		NSLDSSAI2: preprocessString2627(record[nsldsSAI2StartIndex2627-1 : (nsldsSAI2StartIndex2627-1)+nsldsSAI2Length2627]), // Field # 657

		NSLDSPellSchoolCode2: preprocessString2627(record[nsldsPellSchoolCode2StartIndex2627-1 : (nsldsPellSchoolCode2StartIndex2627-1)+nsldsPellSchoolCode2Length2627]), // Field # 658

		NSLDSPellTransactionNumber2: preprocessString2627(record[nsldsPellTransactionNumber2StartIndex2627-1 : (nsldsPellTransactionNumber2StartIndex2627-1)+nsldsPellTransactionNumber2Length2627]), // Field # 659

		NSLDSPellLastDisbursementDate2: parseISIRDate2627(preprocessString2627(record[nsldsPellLastDisbursementDate2StartIndex2627-1 : (nsldsPellLastDisbursementDate2StartIndex2627-1)+nsldsPellLastDisbursementDate2Length2627])), // Field # 660

		NSLDSPellScheduledAmount2: preprocessString2627(record[nsldsPellScheduledAmount2StartIndex2627-1 : (nsldsPellScheduledAmount2StartIndex2627-1)+nsldsPellScheduledAmount2Length2627]), // Field # 661

		NSLDSPellAmountPaidToDate2: parseISIRDate2627(preprocessString2627(record[nsldsPellAmountPaidToDate2StartIndex2627-1 : (nsldsPellAmountPaidToDate2StartIndex2627-1)+nsldsPellAmountPaidToDate2Length2627])), // Field # 662

		NSLDSPellPercentEligibilityUsedDecimal2: preprocessString2627(record[nsldsPellPercentEligibilityUsedDecimal2StartIndex2627-1 : (nsldsPellPercentEligibilityUsedDecimal2StartIndex2627-1)+nsldsPellPercentEligibilityUsedDecimal2Length2627]), // Field # 663

		NSLDSPellAwardAmount2: preprocessString2627(record[nsldsPellAwardAmount2StartIndex2627-1 : (nsldsPellAwardAmount2StartIndex2627-1)+nsldsPellAwardAmount2Length2627]), // Field # 664

		NSLDSAdditionalEligibilityIndicator2: preprocessString2627(record[nsldsAdditionalEligibilityIndicator2StartIndex2627-1 : (nsldsAdditionalEligibilityIndicator2StartIndex2627-1)+nsldsAdditionalEligibilityIndicator2Length2627]), // Field # 665

		NSLDSPellSequenceNumber3: preprocessString2627(record[nsldsPellSequenceNumber3StartIndex2627-1 : (nsldsPellSequenceNumber3StartIndex2627-1)+nsldsPellSequenceNumber3Length2627]), // Field # 667

		NSLDSPellVerificationFlag3: preprocessString2627(record[nsldsPellVerificationFlag3StartIndex2627-1 : (nsldsPellVerificationFlag3StartIndex2627-1)+nsldsPellVerificationFlag3Length2627]), // Field # 668

		NSLDSSAI3: preprocessString2627(record[nsldsSAI3StartIndex2627-1 : (nsldsSAI3StartIndex2627-1)+nsldsSAI3Length2627]), // Field # 669

		NSLDSPellSchoolCode3: preprocessString2627(record[nsldsPellSchoolCode3StartIndex2627-1 : (nsldsPellSchoolCode3StartIndex2627-1)+nsldsPellSchoolCode3Length2627]), // Field # 670

		NSLDSPellTransactionNumber3: preprocessString2627(record[nsldsPellTransactionNumber3StartIndex2627-1 : (nsldsPellTransactionNumber3StartIndex2627-1)+nsldsPellTransactionNumber3Length2627]), // Field # 671

		NSLDSPellLastDisbursementDate3: parseISIRDate2627(preprocessString2627(record[nsldsPellLastDisbursementDate3StartIndex2627-1 : (nsldsPellLastDisbursementDate3StartIndex2627-1)+nsldsPellLastDisbursementDate3Length2627])), // Field # 672

		NSLDSPellScheduledAmount3: preprocessString2627(record[nsldsPellScheduledAmount3StartIndex2627-1 : (nsldsPellScheduledAmount3StartIndex2627-1)+nsldsPellScheduledAmount3Length2627]), // Field # 673

		NSLDSPellAmountPaidToDate3: parseISIRDate2627(preprocessString2627(record[nsldsPellAmountPaidToDate3StartIndex2627-1 : (nsldsPellAmountPaidToDate3StartIndex2627-1)+nsldsPellAmountPaidToDate3Length2627])), // Field # 674

		NSLDSPellPercentEligibilityUsedDecimal3: preprocessString2627(record[nsldsPellPercentEligibilityUsedDecimal3StartIndex2627-1 : (nsldsPellPercentEligibilityUsedDecimal3StartIndex2627-1)+nsldsPellPercentEligibilityUsedDecimal3Length2627]), // Field # 675

		NSLDSPellAwardAmount3: preprocessString2627(record[nsldsPellAwardAmount3StartIndex2627-1 : (nsldsPellAwardAmount3StartIndex2627-1)+nsldsPellAwardAmount3Length2627]), // Field # 676

		NSLDSAdditionalEligibilityIndicator3: preprocessString2627(record[nsldsAdditionalEligibilityIndicator3StartIndex2627-1 : (nsldsAdditionalEligibilityIndicator3StartIndex2627-1)+nsldsAdditionalEligibilityIndicator3Length2627]), // Field # 677

		NSLDSTEACHGrantSequence1: preprocessString2627(record[nsldsTEACHGrantSequence1StartIndex2627-1 : (nsldsTEACHGrantSequence1StartIndex2627-1)+nsldsTEACHGrantSequence1Length2627]), // Field # 679

		NSLDSTEACHGrantSchoolCode1: preprocessString2627(record[nsldsTEACHGrantSchoolCode1StartIndex2627-1 : (nsldsTEACHGrantSchoolCode1StartIndex2627-1)+nsldsTEACHGrantSchoolCode1Length2627]), // Field # 680

		NSLDSTEACHGrantTransactionNumber1: preprocessString2627(record[nsldsTEACHGrantTransactionNumber1StartIndex2627-1 : (nsldsTEACHGrantTransactionNumber1StartIndex2627-1)+nsldsTEACHGrantTransactionNumber1Length2627]), // Field # 681

		NSLDSTEACHGrantLastDisbursementDate1: parseISIRDate2627(preprocessString2627(record[nsldsTEACHGrantLastDisbursementDate1StartIndex2627-1 : (nsldsTEACHGrantLastDisbursementDate1StartIndex2627-1)+nsldsTEACHGrantLastDisbursementDate1Length2627])), // Field # 682

		NSLDSTEACHGrantScheduledAmount1: preprocessString2627(record[nsldsTEACHGrantScheduledAmount1StartIndex2627-1 : (nsldsTEACHGrantScheduledAmount1StartIndex2627-1)+nsldsTEACHGrantScheduledAmount1Length2627]), // Field # 683

		NSLDSTEACHGrantAmountPaidToDate1: parseISIRDate2627(preprocessString2627(record[nsldsTEACHGrantAmountPaidToDate1StartIndex2627-1 : (nsldsTEACHGrantAmountPaidToDate1StartIndex2627-1)+nsldsTEACHGrantAmountPaidToDate1Length2627])), // Field # 684

		NSLDSTEACHGrantAwardAmount1: preprocessString2627(record[nsldsTEACHGrantAwardAmount1StartIndex2627-1 : (nsldsTEACHGrantAwardAmount1StartIndex2627-1)+nsldsTEACHGrantAwardAmount1Length2627]), // Field # 685

		NSLDSTEACHGrantAcademicYearLevel1: preprocessString2627(record[nsldsTEACHGrantAcademicYearLevel1StartIndex2627-1 : (nsldsTEACHGrantAcademicYearLevel1StartIndex2627-1)+nsldsTEACHGrantAcademicYearLevel1Length2627]), // Field # 686

		NSLDSTEACHGrantAwardYear1: preprocessString2627(record[nsldsTEACHGrantAwardYear1StartIndex2627-1 : (nsldsTEACHGrantAwardYear1StartIndex2627-1)+nsldsTEACHGrantAwardYear1Length2627]), // Field # 687

		NSLDSTEACHGrantLoanConversionFlag1: preprocessString2627(record[nsldsTEACHGrantLoanConversionFlag1StartIndex2627-1 : (nsldsTEACHGrantLoanConversionFlag1StartIndex2627-1)+nsldsTEACHGrantLoanConversionFlag1Length2627]), // Field # 688

		NSLDSTEACHGrantDischargeCode1: preprocessString2627(record[nsldsTEACHGrantDischargeCode1StartIndex2627-1 : (nsldsTEACHGrantDischargeCode1StartIndex2627-1)+nsldsTEACHGrantDischargeCode1Length2627]), // Field # 689

		NSLDSTEACHGrantDischargeAmount1: preprocessString2627(record[nsldsTEACHGrantDischargeAmount1StartIndex2627-1 : (nsldsTEACHGrantDischargeAmount1StartIndex2627-1)+nsldsTEACHGrantDischargeAmount1Length2627]), // Field # 690

		NSLDSTEACHGrantAdjustedDisbursement1: preprocessString2627(record[nsldsTEACHGrantAdjustedDisbursement1StartIndex2627-1 : (nsldsTEACHGrantAdjustedDisbursement1StartIndex2627-1)+nsldsTEACHGrantAdjustedDisbursement1Length2627]), // Field # 691

		NSLDSTEACHGrantSequence2: preprocessString2627(record[nsldsTEACHGrantSequence2StartIndex2627-1 : (nsldsTEACHGrantSequence2StartIndex2627-1)+nsldsTEACHGrantSequence2Length2627]), // Field # 693

		NSLDSTEACHGrantSchoolCode2: preprocessString2627(record[nsldsTEACHGrantSchoolCode2StartIndex2627-1 : (nsldsTEACHGrantSchoolCode2StartIndex2627-1)+nsldsTEACHGrantSchoolCode2Length2627]), // Field # 694

		NSLDSTEACHGrantTransactionNumber2: preprocessString2627(record[nsldsTEACHGrantTransactionNumber2StartIndex2627-1 : (nsldsTEACHGrantTransactionNumber2StartIndex2627-1)+nsldsTEACHGrantTransactionNumber2Length2627]), // Field # 695

		NSLDSTEACHGrantLastDisbursementDate2: parseISIRDate2627(preprocessString2627(record[nsldsTEACHGrantLastDisbursementDate2StartIndex2627-1 : (nsldsTEACHGrantLastDisbursementDate2StartIndex2627-1)+nsldsTEACHGrantLastDisbursementDate2Length2627])), // Field # 696

		NSLDSTEACHGrantScheduledAmount2: preprocessString2627(record[nsldsTEACHGrantScheduledAmount2StartIndex2627-1 : (nsldsTEACHGrantScheduledAmount2StartIndex2627-1)+nsldsTEACHGrantScheduledAmount2Length2627]), // Field # 697

		NSLDSTEACHGrantAmountPaidToDate2: parseISIRDate2627(preprocessString2627(record[nsldsTEACHGrantAmountPaidToDate2StartIndex2627-1 : (nsldsTEACHGrantAmountPaidToDate2StartIndex2627-1)+nsldsTEACHGrantAmountPaidToDate2Length2627])), // Field # 698

		NSLDSTEACHGrantAwardAmount2: preprocessString2627(record[nsldsTEACHGrantAwardAmount2StartIndex2627-1 : (nsldsTEACHGrantAwardAmount2StartIndex2627-1)+nsldsTEACHGrantAwardAmount2Length2627]), // Field # 699

		NSLDSTEACHGrantAcademicYearLevel2: preprocessString2627(record[nsldsTEACHGrantAcademicYearLevel2StartIndex2627-1 : (nsldsTEACHGrantAcademicYearLevel2StartIndex2627-1)+nsldsTEACHGrantAcademicYearLevel2Length2627]), // Field # 700

		NSLDSTEACHGrantAwardYear2: preprocessString2627(record[nsldsTEACHGrantAwardYear2StartIndex2627-1 : (nsldsTEACHGrantAwardYear2StartIndex2627-1)+nsldsTEACHGrantAwardYear2Length2627]), // Field # 701

		NSLDSTEACHGrantLoanConversionFlag2: preprocessString2627(record[nsldsTEACHGrantLoanConversionFlag2StartIndex2627-1 : (nsldsTEACHGrantLoanConversionFlag2StartIndex2627-1)+nsldsTEACHGrantLoanConversionFlag2Length2627]), // Field # 702

		NSLDSTEACHGrantDischargeCode2: preprocessString2627(record[nsldsTEACHGrantDischargeCode2StartIndex2627-1 : (nsldsTEACHGrantDischargeCode2StartIndex2627-1)+nsldsTEACHGrantDischargeCode2Length2627]), // Field # 703

		NSLDSTEACHGrantDischargeAmount2: preprocessString2627(record[nsldsTEACHGrantDischargeAmount2StartIndex2627-1 : (nsldsTEACHGrantDischargeAmount2StartIndex2627-1)+nsldsTEACHGrantDischargeAmount2Length2627]), // Field # 704

		NSLDSTEACHGrantAdjustedDisbursement2: preprocessString2627(record[nsldsTEACHGrantAdjustedDisbursement2StartIndex2627-1 : (nsldsTEACHGrantAdjustedDisbursement2StartIndex2627-1)+nsldsTEACHGrantAdjustedDisbursement2Length2627]), // Field # 705

		NSLDSTEACHGrantSequence3: preprocessString2627(record[nsldsTEACHGrantSequence3StartIndex2627-1 : (nsldsTEACHGrantSequence3StartIndex2627-1)+nsldsTEACHGrantSequence3Length2627]), // Field # 707

		NSLDSTEACHGrantSchoolCode3: preprocessString2627(record[nsldsTEACHGrantSchoolCode3StartIndex2627-1 : (nsldsTEACHGrantSchoolCode3StartIndex2627-1)+nsldsTEACHGrantSchoolCode3Length2627]), // Field # 708

		NSLDSTEACHGrantTransactionNumber3: preprocessString2627(record[nsldsTEACHGrantTransactionNumber3StartIndex2627-1 : (nsldsTEACHGrantTransactionNumber3StartIndex2627-1)+nsldsTEACHGrantTransactionNumber3Length2627]), // Field # 709

		NSLDSTEACHGrantLastDisbursementDate3: parseISIRDate2627(preprocessString2627(record[nsldsTEACHGrantLastDisbursementDate3StartIndex2627-1 : (nsldsTEACHGrantLastDisbursementDate3StartIndex2627-1)+nsldsTEACHGrantLastDisbursementDate3Length2627])), // Field # 710

		NSLDSTEACHGrantScheduledAmount3: preprocessString2627(record[nsldsTEACHGrantScheduledAmount3StartIndex2627-1 : (nsldsTEACHGrantScheduledAmount3StartIndex2627-1)+nsldsTEACHGrantScheduledAmount3Length2627]), // Field # 711

		NSLDSTEACHGrantAmountPaidToDate3: parseISIRDate2627(preprocessString2627(record[nsldsTEACHGrantAmountPaidToDate3StartIndex2627-1 : (nsldsTEACHGrantAmountPaidToDate3StartIndex2627-1)+nsldsTEACHGrantAmountPaidToDate3Length2627])), // Field # 712

		NSLDSTEACHGrantAwardAmount3: preprocessString2627(record[nsldsTEACHGrantAwardAmount3StartIndex2627-1 : (nsldsTEACHGrantAwardAmount3StartIndex2627-1)+nsldsTEACHGrantAwardAmount3Length2627]), // Field # 713

		NSLDSTEACHGrantAcademicYearLevel3: preprocessString2627(record[nsldsTEACHGrantAcademicYearLevel3StartIndex2627-1 : (nsldsTEACHGrantAcademicYearLevel3StartIndex2627-1)+nsldsTEACHGrantAcademicYearLevel3Length2627]), // Field # 714

		NSLDSTEACHGrantAwardYear3: preprocessString2627(record[nsldsTEACHGrantAwardYear3StartIndex2627-1 : (nsldsTEACHGrantAwardYear3StartIndex2627-1)+nsldsTEACHGrantAwardYear3Length2627]), // Field # 715

		NSLDSTEACHGrantLoanConversionFlag3: preprocessString2627(record[nsldsTEACHGrantLoanConversionFlag3StartIndex2627-1 : (nsldsTEACHGrantLoanConversionFlag3StartIndex2627-1)+nsldsTEACHGrantLoanConversionFlag3Length2627]), // Field # 716

		NSLDSTEACHGrantDischargeCode3: preprocessString2627(record[nsldsTEACHGrantDischargeCode3StartIndex2627-1 : (nsldsTEACHGrantDischargeCode3StartIndex2627-1)+nsldsTEACHGrantDischargeCode3Length2627]), // Field # 717

		NSLDSTEACHGrantDischargeAmount3: preprocessString2627(record[nsldsTEACHGrantDischargeAmount3StartIndex2627-1 : (nsldsTEACHGrantDischargeAmount3StartIndex2627-1)+nsldsTEACHGrantDischargeAmount3Length2627]), // Field # 718

		NSLDSTEACHGrantAdjustedDisbursement3: preprocessString2627(record[nsldsTEACHGrantAdjustedDisbursement3StartIndex2627-1 : (nsldsTEACHGrantAdjustedDisbursement3StartIndex2627-1)+nsldsTEACHGrantAdjustedDisbursement3Length2627]), // Field # 719

		NSLDSLoanSequenceNumber1: preprocessString2627(record[nsldsLoanSequenceNumber1StartIndex2627-1 : (nsldsLoanSequenceNumber1StartIndex2627-1)+nsldsLoanSequenceNumber1Length2627]), // Field # 721

		NSLDSLoanDefaultedRecentIndicator1: preprocessString2627(record[nsldsLoanDefaultedRecentIndicator1StartIndex2627-1 : (nsldsLoanDefaultedRecentIndicator1StartIndex2627-1)+nsldsLoanDefaultedRecentIndicator1Length2627]), // Field # 722

		NSLDSLoanChangeFlag1: preprocessString2627(record[nsldsLoanChangeFlag1StartIndex2627-1 : (nsldsLoanChangeFlag1StartIndex2627-1)+nsldsLoanChangeFlag1Length2627]), // Field # 723

		NSLDSLoanTypeCode1: preprocessString2627(record[nsldsLoanTypeCode1StartIndex2627-1 : (nsldsLoanTypeCode1StartIndex2627-1)+nsldsLoanTypeCode1Length2627]), // Field # 724

		NSLDSLoanNetAmount1: preprocessString2627(record[nsldsLoanNetAmount1StartIndex2627-1 : (nsldsLoanNetAmount1StartIndex2627-1)+nsldsLoanNetAmount1Length2627]), // Field # 725

		NSLDSLoanCurrentStatusCode1: preprocessString2627(record[nsldsLoanCurrentStatusCode1StartIndex2627-1 : (nsldsLoanCurrentStatusCode1StartIndex2627-1)+nsldsLoanCurrentStatusCode1Length2627]), // Field # 726

		NSLDSLoanCurrentStatusDate1: parseISIRDate2627(preprocessString2627(record[nsldsLoanCurrentStatusDate1StartIndex2627-1 : (nsldsLoanCurrentStatusDate1StartIndex2627-1)+nsldsLoanCurrentStatusDate1Length2627])), // Field # 727

		NSLDSLoanOutstandingPrincipalBalance1: preprocessString2627(record[nsldsLoanOutstandingPrincipalBalance1StartIndex2627-1 : (nsldsLoanOutstandingPrincipalBalance1StartIndex2627-1)+nsldsLoanOutstandingPrincipalBalance1Length2627]), // Field # 728

		NSLDSLoanOutstandingPrincipalBalanceDate1: parseISIRDate2627(preprocessString2627(record[nsldsLoanOutstandingPrincipalBalanceDate1StartIndex2627-1 : (nsldsLoanOutstandingPrincipalBalanceDate1StartIndex2627-1)+nsldsLoanOutstandingPrincipalBalanceDate1Length2627])), // Field # 729

		NSLDSLoanPeriodBeginDate1: parseISIRDate2627(preprocessString2627(record[nsldsLoanPeriodBeginDate1StartIndex2627-1 : (nsldsLoanPeriodBeginDate1StartIndex2627-1)+nsldsLoanPeriodBeginDate1Length2627])), // Field # 730

		NSLDSLoanPeriodEndDate1: parseISIRDate2627(preprocessString2627(record[nsldsLoanPeriodEndDate1StartIndex2627-1 : (nsldsLoanPeriodEndDate1StartIndex2627-1)+nsldsLoanPeriodEndDate1Length2627])), // Field # 731

		NSLDSLoanGuarantyAgencyCode1: preprocessString2627(record[nsldsLoanGuarantyAgencyCode1StartIndex2627-1 : (nsldsLoanGuarantyAgencyCode1StartIndex2627-1)+nsldsLoanGuarantyAgencyCode1Length2627]), // Field # 732

		NSLDSLoanContactType1: preprocessString2627(record[nsldsLoanContactType1StartIndex2627-1 : (nsldsLoanContactType1StartIndex2627-1)+nsldsLoanContactType1Length2627]), // Field # 733

		NSLDSLoanSchoolCode1: preprocessString2627(record[nsldsLoanSchoolCode1StartIndex2627-1 : (nsldsLoanSchoolCode1StartIndex2627-1)+nsldsLoanSchoolCode1Length2627]), // Field # 734

		NSLDSLoanContactCode1: preprocessString2627(record[nsldsLoanContactCode1StartIndex2627-1 : (nsldsLoanContactCode1StartIndex2627-1)+nsldsLoanContactCode1Length2627]), // Field # 735

		NSLDSLoanGradeLevel1: preprocessString2627(record[nsldsLoanGradeLevel1StartIndex2627-1 : (nsldsLoanGradeLevel1StartIndex2627-1)+nsldsLoanGradeLevel1Length2627]), // Field # 736

		NSLDSLoanAdditionalUnsubsidizedFlag1: preprocessString2627(record[nsldsLoanAdditionalUnsubsidizedFlag1StartIndex2627-1 : (nsldsLoanAdditionalUnsubsidizedFlag1StartIndex2627-1)+nsldsLoanAdditionalUnsubsidizedFlag1Length2627]), // Field # 737

		NSLDSLoanCapitalizedInterestFlag1: preprocessString2627(record[nsldsLoanCapitalizedInterestFlag1StartIndex2627-1 : (nsldsLoanCapitalizedInterestFlag1StartIndex2627-1)+nsldsLoanCapitalizedInterestFlag1Length2627]), // Field # 738

		NSLDSLoanDisbursementAmount1: preprocessString2627(record[nsldsLoanDisbursementAmount1StartIndex2627-1 : (nsldsLoanDisbursementAmount1StartIndex2627-1)+nsldsLoanDisbursementAmount1Length2627]), // Field # 739

		NSLDSLoanDisbursementDate1: parseISIRDate2627(preprocessString2627(record[nsldsLoanDisbursementDate1StartIndex2627-1 : (nsldsLoanDisbursementDate1StartIndex2627-1)+nsldsLoanDisbursementDate1Length2627])), // Field # 740

		NSLDSLoanConfirmedLoanSubsidyStatus1: preprocessString2627(record[nsldsLoanConfirmedLoanSubsidyStatus1StartIndex2627-1 : (nsldsLoanConfirmedLoanSubsidyStatus1StartIndex2627-1)+nsldsLoanConfirmedLoanSubsidyStatus1Length2627]), // Field # 741

		NSLDSLoanConfirmedLoanSubsidyStatusDate1: parseISIRDate2627(preprocessString2627(record[nsldsLoanConfirmedLoanSubsidyStatusDate1StartIndex2627-1 : (nsldsLoanConfirmedLoanSubsidyStatusDate1StartIndex2627-1)+nsldsLoanConfirmedLoanSubsidyStatusDate1Length2627])), // Field # 742

		NSLDSLoanSequenceNumber2: preprocessString2627(record[nsldsLoanSequenceNumber2StartIndex2627-1 : (nsldsLoanSequenceNumber2StartIndex2627-1)+nsldsLoanSequenceNumber2Length2627]), // Field # 744

		NSLDSLoanDefaultedRecentIndicator2: preprocessString2627(record[nsldsLoanDefaultedRecentIndicator2StartIndex2627-1 : (nsldsLoanDefaultedRecentIndicator2StartIndex2627-1)+nsldsLoanDefaultedRecentIndicator2Length2627]), // Field # 745

		NSLDSLoanChangeFlag2: preprocessString2627(record[nsldsLoanChangeFlag2StartIndex2627-1 : (nsldsLoanChangeFlag2StartIndex2627-1)+nsldsLoanChangeFlag2Length2627]), // Field # 746

		NSLDSLoanTypeCode2: preprocessString2627(record[nsldsLoanTypeCode2StartIndex2627-1 : (nsldsLoanTypeCode2StartIndex2627-1)+nsldsLoanTypeCode2Length2627]), // Field # 747

		NSLDSLoanNetAmount2: preprocessString2627(record[nsldsLoanNetAmount2StartIndex2627-1 : (nsldsLoanNetAmount2StartIndex2627-1)+nsldsLoanNetAmount2Length2627]), // Field # 748

		NSLDSLoanCurrentStatusCode2: preprocessString2627(record[nsldsLoanCurrentStatusCode2StartIndex2627-1 : (nsldsLoanCurrentStatusCode2StartIndex2627-1)+nsldsLoanCurrentStatusCode2Length2627]), // Field # 749

		NSLDSLoanCurrentStatusDate2: parseISIRDate2627(preprocessString2627(record[nsldsLoanCurrentStatusDate2StartIndex2627-1 : (nsldsLoanCurrentStatusDate2StartIndex2627-1)+nsldsLoanCurrentStatusDate2Length2627])), // Field # 750

		NSLDSLoanOutstandingPrincipalBalance2: preprocessString2627(record[nsldsLoanOutstandingPrincipalBalance2StartIndex2627-1 : (nsldsLoanOutstandingPrincipalBalance2StartIndex2627-1)+nsldsLoanOutstandingPrincipalBalance2Length2627]), // Field # 751

		NSLDSLoanOutstandingPrincipalBalanceDate2: parseISIRDate2627(preprocessString2627(record[nsldsLoanOutstandingPrincipalBalanceDate2StartIndex2627-1 : (nsldsLoanOutstandingPrincipalBalanceDate2StartIndex2627-1)+nsldsLoanOutstandingPrincipalBalanceDate2Length2627])), // Field # 752

		NSLDSLoanPeriodBeginDate2: parseISIRDate2627(preprocessString2627(record[nsldsLoanPeriodBeginDate2StartIndex2627-1 : (nsldsLoanPeriodBeginDate2StartIndex2627-1)+nsldsLoanPeriodBeginDate2Length2627])), // Field # 753

		NSLDSLoanPeriodEndDate2: parseISIRDate2627(preprocessString2627(record[nsldsLoanPeriodEndDate2StartIndex2627-1 : (nsldsLoanPeriodEndDate2StartIndex2627-1)+nsldsLoanPeriodEndDate2Length2627])), // Field # 754

		NSLDSLoanGuarantyAgencyCode2: preprocessString2627(record[nsldsLoanGuarantyAgencyCode2StartIndex2627-1 : (nsldsLoanGuarantyAgencyCode2StartIndex2627-1)+nsldsLoanGuarantyAgencyCode2Length2627]), // Field # 755

		NSLDSLoanContactType2: preprocessString2627(record[nsldsLoanContactType2StartIndex2627-1 : (nsldsLoanContactType2StartIndex2627-1)+nsldsLoanContactType2Length2627]), // Field # 756

		NSLDSLoanSchoolCode2: preprocessString2627(record[nsldsLoanSchoolCode2StartIndex2627-1 : (nsldsLoanSchoolCode2StartIndex2627-1)+nsldsLoanSchoolCode2Length2627]), // Field # 757

		NSLDSLoanContactCode2: preprocessString2627(record[nsldsLoanContactCode2StartIndex2627-1 : (nsldsLoanContactCode2StartIndex2627-1)+nsldsLoanContactCode2Length2627]), // Field # 758

		NSLDSLoanGradeLevel2: preprocessString2627(record[nsldsLoanGradeLevel2StartIndex2627-1 : (nsldsLoanGradeLevel2StartIndex2627-1)+nsldsLoanGradeLevel2Length2627]), // Field # 759

		NSLDSLoanAdditionalUnsubsidizedFlag2: preprocessString2627(record[nsldsLoanAdditionalUnsubsidizedFlag2StartIndex2627-1 : (nsldsLoanAdditionalUnsubsidizedFlag2StartIndex2627-1)+nsldsLoanAdditionalUnsubsidizedFlag2Length2627]), // Field # 760

		NSLDSLoanCapitalizedInterestFlag2: preprocessString2627(record[nsldsLoanCapitalizedInterestFlag2StartIndex2627-1 : (nsldsLoanCapitalizedInterestFlag2StartIndex2627-1)+nsldsLoanCapitalizedInterestFlag2Length2627]), // Field # 761

		NSLDSLoanDisbursementAmount2: preprocessString2627(record[nsldsLoanDisbursementAmount2StartIndex2627-1 : (nsldsLoanDisbursementAmount2StartIndex2627-1)+nsldsLoanDisbursementAmount2Length2627]), // Field # 762

		NSLDSLoanDisbursementDate2: parseISIRDate2627(preprocessString2627(record[nsldsLoanDisbursementDate2StartIndex2627-1 : (nsldsLoanDisbursementDate2StartIndex2627-1)+nsldsLoanDisbursementDate2Length2627])), // Field # 763

		NSLDSLoanConfirmedLoanSubsidyStatus2: preprocessString2627(record[nsldsLoanConfirmedLoanSubsidyStatus2StartIndex2627-1 : (nsldsLoanConfirmedLoanSubsidyStatus2StartIndex2627-1)+nsldsLoanConfirmedLoanSubsidyStatus2Length2627]), // Field # 764

		NSLDSLoanConfirmedLoanSubsidyStatusDate2: parseISIRDate2627(preprocessString2627(record[nsldsLoanConfirmedLoanSubsidyStatusDate2StartIndex2627-1 : (nsldsLoanConfirmedLoanSubsidyStatusDate2StartIndex2627-1)+nsldsLoanConfirmedLoanSubsidyStatusDate2Length2627])), // Field # 765

		NSLDSLoanSequenceNumber3: preprocessString2627(record[nsldsLoanSequenceNumber3StartIndex2627-1 : (nsldsLoanSequenceNumber3StartIndex2627-1)+nsldsLoanSequenceNumber3Length2627]), // Field # 767

		NSLDSLoanDefaultedRecentIndicator3: preprocessString2627(record[nsldsLoanDefaultedRecentIndicator3StartIndex2627-1 : (nsldsLoanDefaultedRecentIndicator3StartIndex2627-1)+nsldsLoanDefaultedRecentIndicator3Length2627]), // Field # 768

		NSLDSLoanChangeFlag3: preprocessString2627(record[nsldsLoanChangeFlag3StartIndex2627-1 : (nsldsLoanChangeFlag3StartIndex2627-1)+nsldsLoanChangeFlag3Length2627]), // Field # 769

		NSLDSLoanTypeCode3: preprocessString2627(record[nsldsLoanTypeCode3StartIndex2627-1 : (nsldsLoanTypeCode3StartIndex2627-1)+nsldsLoanTypeCode3Length2627]), // Field # 770

		NSLDSLoanNetAmount3: preprocessString2627(record[nsldsLoanNetAmount3StartIndex2627-1 : (nsldsLoanNetAmount3StartIndex2627-1)+nsldsLoanNetAmount3Length2627]), // Field # 771

		NSLDSLoanCurrentStatusCode3: preprocessString2627(record[nsldsLoanCurrentStatusCode3StartIndex2627-1 : (nsldsLoanCurrentStatusCode3StartIndex2627-1)+nsldsLoanCurrentStatusCode3Length2627]), // Field # 772

		NSLDSLoanCurrentStatusDate3: parseISIRDate2627(preprocessString2627(record[nsldsLoanCurrentStatusDate3StartIndex2627-1 : (nsldsLoanCurrentStatusDate3StartIndex2627-1)+nsldsLoanCurrentStatusDate3Length2627])), // Field # 773

		NSLDSLoanOutstandingPrincipalBalance3: preprocessString2627(record[nsldsLoanOutstandingPrincipalBalance3StartIndex2627-1 : (nsldsLoanOutstandingPrincipalBalance3StartIndex2627-1)+nsldsLoanOutstandingPrincipalBalance3Length2627]), // Field # 774

		NSLDSLoanOutstandingPrincipalBalanceDate3: parseISIRDate2627(preprocessString2627(record[nsldsLoanOutstandingPrincipalBalanceDate3StartIndex2627-1 : (nsldsLoanOutstandingPrincipalBalanceDate3StartIndex2627-1)+nsldsLoanOutstandingPrincipalBalanceDate3Length2627])), // Field # 775

		NSLDSLoanPeriodBeginDate3: parseISIRDate2627(preprocessString2627(record[nsldsLoanPeriodBeginDate3StartIndex2627-1 : (nsldsLoanPeriodBeginDate3StartIndex2627-1)+nsldsLoanPeriodBeginDate3Length2627])), // Field # 776

		NSLDSLoanPeriodEndDate3: parseISIRDate2627(preprocessString2627(record[nsldsLoanPeriodEndDate3StartIndex2627-1 : (nsldsLoanPeriodEndDate3StartIndex2627-1)+nsldsLoanPeriodEndDate3Length2627])), // Field # 777

		NSLDSLoanGuarantyAgencyCode3: preprocessString2627(record[nsldsLoanGuarantyAgencyCode3StartIndex2627-1 : (nsldsLoanGuarantyAgencyCode3StartIndex2627-1)+nsldsLoanGuarantyAgencyCode3Length2627]), // Field # 778

		NSLDSLoanContactType3: preprocessString2627(record[nsldsLoanContactType3StartIndex2627-1 : (nsldsLoanContactType3StartIndex2627-1)+nsldsLoanContactType3Length2627]), // Field # 779

		NSLDSLoanSchoolCode3: preprocessString2627(record[nsldsLoanSchoolCode3StartIndex2627-1 : (nsldsLoanSchoolCode3StartIndex2627-1)+nsldsLoanSchoolCode3Length2627]), // Field # 780

		NSLDSLoanContactCode3: preprocessString2627(record[nsldsLoanContactCode3StartIndex2627-1 : (nsldsLoanContactCode3StartIndex2627-1)+nsldsLoanContactCode3Length2627]), // Field # 781

		NSLDSLoanGradeLevel3: preprocessString2627(record[nsldsLoanGradeLevel3StartIndex2627-1 : (nsldsLoanGradeLevel3StartIndex2627-1)+nsldsLoanGradeLevel3Length2627]), // Field # 782

		NSLDSLoanAdditionalUnsubsidizedFlag3: preprocessString2627(record[nsldsLoanAdditionalUnsubsidizedFlag3StartIndex2627-1 : (nsldsLoanAdditionalUnsubsidizedFlag3StartIndex2627-1)+nsldsLoanAdditionalUnsubsidizedFlag3Length2627]), // Field # 783

		NSLDSLoanCapitalizedInterestFlag3: preprocessString2627(record[nsldsLoanCapitalizedInterestFlag3StartIndex2627-1 : (nsldsLoanCapitalizedInterestFlag3StartIndex2627-1)+nsldsLoanCapitalizedInterestFlag3Length2627]), // Field # 784

		NSLDSLoanDisbursementAmount3: preprocessString2627(record[nsldsLoanDisbursementAmount3StartIndex2627-1 : (nsldsLoanDisbursementAmount3StartIndex2627-1)+nsldsLoanDisbursementAmount3Length2627]), // Field # 785

		NSLDSLoanDisbursementDate3: parseISIRDate2627(preprocessString2627(record[nsldsLoanDisbursementDate3StartIndex2627-1 : (nsldsLoanDisbursementDate3StartIndex2627-1)+nsldsLoanDisbursementDate3Length2627])), // Field # 786

		NSLDSLoanConfirmedLoanSubsidyStatus3: preprocessString2627(record[nsldsLoanConfirmedLoanSubsidyStatus3StartIndex2627-1 : (nsldsLoanConfirmedLoanSubsidyStatus3StartIndex2627-1)+nsldsLoanConfirmedLoanSubsidyStatus3Length2627]), // Field # 787

		NSLDSLoanConfirmedLoanSubsidyStatusDate3: parseISIRDate2627(preprocessString2627(record[nsldsLoanConfirmedLoanSubsidyStatusDate3StartIndex2627-1 : (nsldsLoanConfirmedLoanSubsidyStatusDate3StartIndex2627-1)+nsldsLoanConfirmedLoanSubsidyStatusDate3Length2627])), // Field # 788

		NSLDSLoanSequenceNumber4: preprocessString2627(record[nsldsLoanSequenceNumber4StartIndex2627-1 : (nsldsLoanSequenceNumber4StartIndex2627-1)+nsldsLoanSequenceNumber4Length2627]), // Field # 790

		NSLDSLoanDefaultedRecentIndicator4: preprocessString2627(record[nsldsLoanDefaultedRecentIndicator4StartIndex2627-1 : (nsldsLoanDefaultedRecentIndicator4StartIndex2627-1)+nsldsLoanDefaultedRecentIndicator4Length2627]), // Field # 791

		NSLDSLoanChangeFlag4: preprocessString2627(record[nsldsLoanChangeFlag4StartIndex2627-1 : (nsldsLoanChangeFlag4StartIndex2627-1)+nsldsLoanChangeFlag4Length2627]), // Field # 792

		NSLDSLoanTypeCode4: preprocessString2627(record[nsldsLoanTypeCode4StartIndex2627-1 : (nsldsLoanTypeCode4StartIndex2627-1)+nsldsLoanTypeCode4Length2627]), // Field # 793

		NSLDSLoanNetAmount4: preprocessString2627(record[nsldsLoanNetAmount4StartIndex2627-1 : (nsldsLoanNetAmount4StartIndex2627-1)+nsldsLoanNetAmount4Length2627]), // Field # 794

		NSLDSLoanCurrentStatusCode4: preprocessString2627(record[nsldsLoanCurrentStatusCode4StartIndex2627-1 : (nsldsLoanCurrentStatusCode4StartIndex2627-1)+nsldsLoanCurrentStatusCode4Length2627]), // Field # 795

		NSLDSLoanCurrentStatusDate4: parseISIRDate2627(preprocessString2627(record[nsldsLoanCurrentStatusDate4StartIndex2627-1 : (nsldsLoanCurrentStatusDate4StartIndex2627-1)+nsldsLoanCurrentStatusDate4Length2627])), // Field # 796

		NSLDSLoanOutstandingPrincipalBalance4: preprocessString2627(record[nsldsLoanOutstandingPrincipalBalance4StartIndex2627-1 : (nsldsLoanOutstandingPrincipalBalance4StartIndex2627-1)+nsldsLoanOutstandingPrincipalBalance4Length2627]), // Field # 797

		NSLDSLoanOutstandingPrincipalBalanceDate4: parseISIRDate2627(preprocessString2627(record[nsldsLoanOutstandingPrincipalBalanceDate4StartIndex2627-1 : (nsldsLoanOutstandingPrincipalBalanceDate4StartIndex2627-1)+nsldsLoanOutstandingPrincipalBalanceDate4Length2627])), // Field # 798

		NSLDSLoanPeriodBeginDate4: parseISIRDate2627(preprocessString2627(record[nsldsLoanPeriodBeginDate4StartIndex2627-1 : (nsldsLoanPeriodBeginDate4StartIndex2627-1)+nsldsLoanPeriodBeginDate4Length2627])), // Field # 799

		NSLDSLoanPeriodEndDate4: parseISIRDate2627(preprocessString2627(record[nsldsLoanPeriodEndDate4StartIndex2627-1 : (nsldsLoanPeriodEndDate4StartIndex2627-1)+nsldsLoanPeriodEndDate4Length2627])), // Field # 800

		NSLDSLoanGuarantyAgencyCode4: preprocessString2627(record[nsldsLoanGuarantyAgencyCode4StartIndex2627-1 : (nsldsLoanGuarantyAgencyCode4StartIndex2627-1)+nsldsLoanGuarantyAgencyCode4Length2627]), // Field # 801

		NSLDSLoanContactType4: preprocessString2627(record[nsldsLoanContactType4StartIndex2627-1 : (nsldsLoanContactType4StartIndex2627-1)+nsldsLoanContactType4Length2627]), // Field # 802

		NSLDSLoanSchoolCode4: preprocessString2627(record[nsldsLoanSchoolCode4StartIndex2627-1 : (nsldsLoanSchoolCode4StartIndex2627-1)+nsldsLoanSchoolCode4Length2627]), // Field # 803

		NSLDSLoanContactCode4: preprocessString2627(record[nsldsLoanContactCode4StartIndex2627-1 : (nsldsLoanContactCode4StartIndex2627-1)+nsldsLoanContactCode4Length2627]), // Field # 804

		NSLDSLoanGradeLevel4: preprocessString2627(record[nsldsLoanGradeLevel4StartIndex2627-1 : (nsldsLoanGradeLevel4StartIndex2627-1)+nsldsLoanGradeLevel4Length2627]), // Field # 805

		NSLDSLoanAdditionalUnsubsidizedFlag4: preprocessString2627(record[nsldsLoanAdditionalUnsubsidizedFlag4StartIndex2627-1 : (nsldsLoanAdditionalUnsubsidizedFlag4StartIndex2627-1)+nsldsLoanAdditionalUnsubsidizedFlag4Length2627]), // Field # 806

		NSLDSLoanCapitalizedInterestFlag4: preprocessString2627(record[nsldsLoanCapitalizedInterestFlag4StartIndex2627-1 : (nsldsLoanCapitalizedInterestFlag4StartIndex2627-1)+nsldsLoanCapitalizedInterestFlag4Length2627]), // Field # 807

		NSLDSLoanDisbursementAmount4: preprocessString2627(record[nsldsLoanDisbursementAmount4StartIndex2627-1 : (nsldsLoanDisbursementAmount4StartIndex2627-1)+nsldsLoanDisbursementAmount4Length2627]), // Field # 808

		NSLDSLoanDisbursementDate4: parseISIRDate2627(preprocessString2627(record[nsldsLoanDisbursementDate4StartIndex2627-1 : (nsldsLoanDisbursementDate4StartIndex2627-1)+nsldsLoanDisbursementDate4Length2627])), // Field # 809

		NSLDSLoanConfirmedLoanSubsidyStatus4: preprocessString2627(record[nsldsLoanConfirmedLoanSubsidyStatus4StartIndex2627-1 : (nsldsLoanConfirmedLoanSubsidyStatus4StartIndex2627-1)+nsldsLoanConfirmedLoanSubsidyStatus4Length2627]), // Field # 810

		NSLDSLoanSubsidyStatusDate4: parseISIRDate2627(preprocessString2627(record[nsldsLoanSubsidyStatusDate4StartIndex2627-1 : (nsldsLoanSubsidyStatusDate4StartIndex2627-1)+nsldsLoanSubsidyStatusDate4Length2627])), // Field # 811

		NSLDSLoanSequenceNumber5: preprocessString2627(record[nsldsLoanSequenceNumber5StartIndex2627-1 : (nsldsLoanSequenceNumber5StartIndex2627-1)+nsldsLoanSequenceNumber5Length2627]), // Field # 813

		NSLDSLoanDefaultedRecentIndicator5: preprocessString2627(record[nsldsLoanDefaultedRecentIndicator5StartIndex2627-1 : (nsldsLoanDefaultedRecentIndicator5StartIndex2627-1)+nsldsLoanDefaultedRecentIndicator5Length2627]), // Field # 814

		NSLDSLoanChangeFlag5: preprocessString2627(record[nsldsLoanChangeFlag5StartIndex2627-1 : (nsldsLoanChangeFlag5StartIndex2627-1)+nsldsLoanChangeFlag5Length2627]), // Field # 815

		NSLDSLoanTypeCode5: preprocessString2627(record[nsldsLoanTypeCode5StartIndex2627-1 : (nsldsLoanTypeCode5StartIndex2627-1)+nsldsLoanTypeCode5Length2627]), // Field # 816

		NSLDSLoanNetAmount5: preprocessString2627(record[nsldsLoanNetAmount5StartIndex2627-1 : (nsldsLoanNetAmount5StartIndex2627-1)+nsldsLoanNetAmount5Length2627]), // Field # 817

		NSLDSLoanCurrentStatusCode5: preprocessString2627(record[nsldsLoanCurrentStatusCode5StartIndex2627-1 : (nsldsLoanCurrentStatusCode5StartIndex2627-1)+nsldsLoanCurrentStatusCode5Length2627]), // Field # 818

		NSLDSLoanCurrentStatusDate5: parseISIRDate2627(preprocessString2627(record[nsldsLoanCurrentStatusDate5StartIndex2627-1 : (nsldsLoanCurrentStatusDate5StartIndex2627-1)+nsldsLoanCurrentStatusDate5Length2627])), // Field # 819

		NSLDSLoanOutstandingPrincipalBalance5: preprocessString2627(record[nsldsLoanOutstandingPrincipalBalance5StartIndex2627-1 : (nsldsLoanOutstandingPrincipalBalance5StartIndex2627-1)+nsldsLoanOutstandingPrincipalBalance5Length2627]), // Field # 820

		NSLDSLoanOutstandingPrincipalBalanceDate5: parseISIRDate2627(preprocessString2627(record[nsldsLoanOutstandingPrincipalBalanceDate5StartIndex2627-1 : (nsldsLoanOutstandingPrincipalBalanceDate5StartIndex2627-1)+nsldsLoanOutstandingPrincipalBalanceDate5Length2627])), // Field # 821

		NSLDSLoanPeriodBeginDate5: parseISIRDate2627(preprocessString2627(record[nsldsLoanPeriodBeginDate5StartIndex2627-1 : (nsldsLoanPeriodBeginDate5StartIndex2627-1)+nsldsLoanPeriodBeginDate5Length2627])), // Field # 822

		NSLDSLoanPeriodEndDate5: parseISIRDate2627(preprocessString2627(record[nsldsLoanPeriodEndDate5StartIndex2627-1 : (nsldsLoanPeriodEndDate5StartIndex2627-1)+nsldsLoanPeriodEndDate5Length2627])), // Field # 823

		NSLDSLoanGuarantyAgencyCode5: preprocessString2627(record[nsldsLoanGuarantyAgencyCode5StartIndex2627-1 : (nsldsLoanGuarantyAgencyCode5StartIndex2627-1)+nsldsLoanGuarantyAgencyCode5Length2627]), // Field # 824

		NSLDSLoanContactType5: preprocessString2627(record[nsldsLoanContactType5StartIndex2627-1 : (nsldsLoanContactType5StartIndex2627-1)+nsldsLoanContactType5Length2627]), // Field # 825

		NSLDSLoanSchoolCode5: preprocessString2627(record[nsldsLoanSchoolCode5StartIndex2627-1 : (nsldsLoanSchoolCode5StartIndex2627-1)+nsldsLoanSchoolCode5Length2627]), // Field # 826

		NSLDSLoanContactCode5: preprocessString2627(record[nsldsLoanContactCode5StartIndex2627-1 : (nsldsLoanContactCode5StartIndex2627-1)+nsldsLoanContactCode5Length2627]), // Field # 827

		NSLDSLoanGradeLevel5: preprocessString2627(record[nsldsLoanGradeLevel5StartIndex2627-1 : (nsldsLoanGradeLevel5StartIndex2627-1)+nsldsLoanGradeLevel5Length2627]), // Field # 828

		NSLDSLoanAdditionalUnsubsidizedFlag5: preprocessString2627(record[nsldsLoanAdditionalUnsubsidizedFlag5StartIndex2627-1 : (nsldsLoanAdditionalUnsubsidizedFlag5StartIndex2627-1)+nsldsLoanAdditionalUnsubsidizedFlag5Length2627]), // Field # 829

		NSLDSLoanCapitalizedInterestFlag5: preprocessString2627(record[nsldsLoanCapitalizedInterestFlag5StartIndex2627-1 : (nsldsLoanCapitalizedInterestFlag5StartIndex2627-1)+nsldsLoanCapitalizedInterestFlag5Length2627]), // Field # 830

		NSLDSLoanDisbursementAmount5: preprocessString2627(record[nsldsLoanDisbursementAmount5StartIndex2627-1 : (nsldsLoanDisbursementAmount5StartIndex2627-1)+nsldsLoanDisbursementAmount5Length2627]), // Field # 831

		NSLDSLoanDisbursementDate5: parseISIRDate2627(preprocessString2627(record[nsldsLoanDisbursementDate5StartIndex2627-1 : (nsldsLoanDisbursementDate5StartIndex2627-1)+nsldsLoanDisbursementDate5Length2627])), // Field # 832

		NSLDSLoanConfirmedLoanSubsidyStatus5: preprocessString2627(record[nsldsLoanConfirmedLoanSubsidyStatus5StartIndex2627-1 : (nsldsLoanConfirmedLoanSubsidyStatus5StartIndex2627-1)+nsldsLoanConfirmedLoanSubsidyStatus5Length2627]), // Field # 833

		NSLDSLoanSubsidyStatusDate5: parseISIRDate2627(preprocessString2627(record[nsldsLoanSubsidyStatusDate5StartIndex2627-1 : (nsldsLoanSubsidyStatusDate5StartIndex2627-1)+nsldsLoanSubsidyStatusDate5Length2627])), // Field # 834

		NSLDSLoanSequenceNumber6: preprocessString2627(record[nsldsLoanSequenceNumber6StartIndex2627-1 : (nsldsLoanSequenceNumber6StartIndex2627-1)+nsldsLoanSequenceNumber6Length2627]), // Field # 836

		NSLDSLoanDefaultedRecentIndicator6: preprocessString2627(record[nsldsLoanDefaultedRecentIndicator6StartIndex2627-1 : (nsldsLoanDefaultedRecentIndicator6StartIndex2627-1)+nsldsLoanDefaultedRecentIndicator6Length2627]), // Field # 837

		NSLDSLoanChangeFlag6: preprocessString2627(record[nsldsLoanChangeFlag6StartIndex2627-1 : (nsldsLoanChangeFlag6StartIndex2627-1)+nsldsLoanChangeFlag6Length2627]), // Field # 838

		NSLDSLoanTypeCode6: preprocessString2627(record[nsldsLoanTypeCode6StartIndex2627-1 : (nsldsLoanTypeCode6StartIndex2627-1)+nsldsLoanTypeCode6Length2627]), // Field # 839

		NSLDSLoanNetAmount6: preprocessString2627(record[nsldsLoanNetAmount6StartIndex2627-1 : (nsldsLoanNetAmount6StartIndex2627-1)+nsldsLoanNetAmount6Length2627]), // Field # 840

		NSLDSLoanCurrentStatusCode6: preprocessString2627(record[nsldsLoanCurrentStatusCode6StartIndex2627-1 : (nsldsLoanCurrentStatusCode6StartIndex2627-1)+nsldsLoanCurrentStatusCode6Length2627]), // Field # 841

		NSLDSLoanCurrentStatusDate6: parseISIRDate2627(preprocessString2627(record[nsldsLoanCurrentStatusDate6StartIndex2627-1 : (nsldsLoanCurrentStatusDate6StartIndex2627-1)+nsldsLoanCurrentStatusDate6Length2627])), // Field # 842

		NSLDSLoanOutstandingPrincipalBalance6: preprocessString2627(record[nsldsLoanOutstandingPrincipalBalance6StartIndex2627-1 : (nsldsLoanOutstandingPrincipalBalance6StartIndex2627-1)+nsldsLoanOutstandingPrincipalBalance6Length2627]), // Field # 843

		NSLDSLoanOutstandingPrincipalBalanceDate6: parseISIRDate2627(preprocessString2627(record[nsldsLoanOutstandingPrincipalBalanceDate6StartIndex2627-1 : (nsldsLoanOutstandingPrincipalBalanceDate6StartIndex2627-1)+nsldsLoanOutstandingPrincipalBalanceDate6Length2627])), // Field # 844

		NSLDSLoanPeriodBeginDate6: parseISIRDate2627(preprocessString2627(record[nsldsLoanPeriodBeginDate6StartIndex2627-1 : (nsldsLoanPeriodBeginDate6StartIndex2627-1)+nsldsLoanPeriodBeginDate6Length2627])), // Field # 845

		NSLDSLoanPeriodEndDate6: parseISIRDate2627(preprocessString2627(record[nsldsLoanPeriodEndDate6StartIndex2627-1 : (nsldsLoanPeriodEndDate6StartIndex2627-1)+nsldsLoanPeriodEndDate6Length2627])), // Field # 846

		NSLDSLoanGuarantyAgencyCode6: preprocessString2627(record[nsldsLoanGuarantyAgencyCode6StartIndex2627-1 : (nsldsLoanGuarantyAgencyCode6StartIndex2627-1)+nsldsLoanGuarantyAgencyCode6Length2627]), // Field # 847

		NSLDSLoanContactType6: preprocessString2627(record[nsldsLoanContactType6StartIndex2627-1 : (nsldsLoanContactType6StartIndex2627-1)+nsldsLoanContactType6Length2627]), // Field # 848

		NSLDSLoanSchoolCode6: preprocessString2627(record[nsldsLoanSchoolCode6StartIndex2627-1 : (nsldsLoanSchoolCode6StartIndex2627-1)+nsldsLoanSchoolCode6Length2627]), // Field # 849

		NSLDSLoanContactCode6: preprocessString2627(record[nsldsLoanContactCode6StartIndex2627-1 : (nsldsLoanContactCode6StartIndex2627-1)+nsldsLoanContactCode6Length2627]), // Field # 850

		NSLDSLoanGradeLevel6: preprocessString2627(record[nsldsLoanGradeLevel6StartIndex2627-1 : (nsldsLoanGradeLevel6StartIndex2627-1)+nsldsLoanGradeLevel6Length2627]), // Field # 851

		NSLDSLoanAdditionalUnsubsidizedFlag6: preprocessString2627(record[nsldsLoanAdditionalUnsubsidizedFlag6StartIndex2627-1 : (nsldsLoanAdditionalUnsubsidizedFlag6StartIndex2627-1)+nsldsLoanAdditionalUnsubsidizedFlag6Length2627]), // Field # 852

		NSLDSLoanCapitalizedInterestFlag6: preprocessString2627(record[nsldsLoanCapitalizedInterestFlag6StartIndex2627-1 : (nsldsLoanCapitalizedInterestFlag6StartIndex2627-1)+nsldsLoanCapitalizedInterestFlag6Length2627]), // Field # 853

		NSLDSLoanDisbursementAmount6: preprocessString2627(record[nsldsLoanDisbursementAmount6StartIndex2627-1 : (nsldsLoanDisbursementAmount6StartIndex2627-1)+nsldsLoanDisbursementAmount6Length2627]), // Field # 854

		NSLDSLoanDisbursementDate6: parseISIRDate2627(preprocessString2627(record[nsldsLoanDisbursementDate6StartIndex2627-1 : (nsldsLoanDisbursementDate6StartIndex2627-1)+nsldsLoanDisbursementDate6Length2627])), // Field # 855

		NSLDSLoanConfirmedLoanSubsidyStatus6: preprocessString2627(record[nsldsLoanConfirmedLoanSubsidyStatus6StartIndex2627-1 : (nsldsLoanConfirmedLoanSubsidyStatus6StartIndex2627-1)+nsldsLoanConfirmedLoanSubsidyStatus6Length2627]), // Field # 856

		NSLDSLoanSubsidyStatusDate6: parseISIRDate2627(preprocessString2627(record[nsldsLoanSubsidyStatusDate6StartIndex2627-1 : (nsldsLoanSubsidyStatusDate6StartIndex2627-1)+nsldsLoanSubsidyStatusDate6Length2627])), // Field # 857

		//FTILabelStart: preprocessString2627(record[ftiLabelStartStartIndex2627-1 : (ftiLabelStartStartIndex2627-1)+ftiLabelStartLength2627]), // Field # 861

		StudentFTIMReturnedTaxYear: preprocessString2627(record[studentFTIMReturnedTaxYearStartIndex2627-1 : (studentFTIMReturnedTaxYearStartIndex2627-1)+studentFTIMReturnedTaxYearLength2627]), // Field # 862

		StudentFTIMFilingStatusCode: preprocessString2627(record[studentFTIMFilingStatusCodeStartIndex2627-1 : (studentFTIMFilingStatusCodeStartIndex2627-1)+studentFTIMFilingStatusCodeLength2627]), // Field # 863

		StudentFTIMAdjustedGrossIncome: preprocessString2627(record[studentFTIMAdjustedGrossIncomeStartIndex2627-1 : (studentFTIMAdjustedGrossIncomeStartIndex2627-1)+studentFTIMAdjustedGrossIncomeLength2627]), // Field # 864

		StudentFTIMNumberOfExemptions: preprocessString2627(record[studentFTIMNumberOfExemptionsStartIndex2627-1 : (studentFTIMNumberOfExemptionsStartIndex2627-1)+studentFTIMNumberOfExemptionsLength2627]), // Field # 865

		StudentFTIMNumberOfDependents: preprocessString2627(record[studentFTIMNumberOfDependentsStartIndex2627-1 : (studentFTIMNumberOfDependentsStartIndex2627-1)+studentFTIMNumberOfDependentsLength2627]), // Field # 866

		StudentFTIMTotalIncomeEarnedAmount: preprocessString2627(record[studentFTIMTotalIncomeEarnedAmountStartIndex2627-1 : (studentFTIMTotalIncomeEarnedAmountStartIndex2627-1)+studentFTIMTotalIncomeEarnedAmountLength2627]), // Field # 867

		StudentFTIMIncomeTaxPaid: preprocessString2627(record[studentFTIMIncomeTaxPaidStartIndex2627-1 : (studentFTIMIncomeTaxPaidStartIndex2627-1)+studentFTIMIncomeTaxPaidLength2627]), // Field # 868

		StudentFTIMEducationCredits: preprocessString2627(record[studentFTIMEducationCreditsStartIndex2627-1 : (studentFTIMEducationCreditsStartIndex2627-1)+studentFTIMEducationCreditsLength2627]), // Field # 869

		StudentFTIMUntaxedIRADistributions: preprocessString2627(record[studentFTIMUntaxedIRADistributionsStartIndex2627-1 : (studentFTIMUntaxedIRADistributionsStartIndex2627-1)+studentFTIMUntaxedIRADistributionsLength2627]), // Field # 870

		StudentFTIMIRADeductibleAndPayments: preprocessString2627(record[studentFTIMIRADeductibleAndPaymentsStartIndex2627-1 : (studentFTIMIRADeductibleAndPaymentsStartIndex2627-1)+studentFTIMIRADeductibleAndPaymentsLength2627]), // Field # 871

		StudentFTIMTaxExemptInterest: preprocessString2627(record[studentFTIMTaxExemptInterestStartIndex2627-1 : (studentFTIMTaxExemptInterestStartIndex2627-1)+studentFTIMTaxExemptInterestLength2627]), // Field # 872

		StudentFTIMUntaxedPensionsAmount: preprocessString2627(record[studentFTIMUntaxedPensionsAmountStartIndex2627-1 : (studentFTIMUntaxedPensionsAmountStartIndex2627-1)+studentFTIMUntaxedPensionsAmountLength2627]), // Field # 873

		StudentFTIMScheduleCNetProfitLoss: preprocessString2627(record[studentFTIMScheduleCNetProfitLossStartIndex2627-1 : (studentFTIMScheduleCNetProfitLossStartIndex2627-1)+studentFTIMScheduleCNetProfitLossLength2627]), // Field # 874

		StudentFTIMScheduleAIndicator: preprocessString2627(record[studentFTIMScheduleAIndicatorStartIndex2627-1 : (studentFTIMScheduleAIndicatorStartIndex2627-1)+studentFTIMScheduleAIndicatorLength2627]), // Field # 875

		StudentFTIMScheduleBIndicator: preprocessString2627(record[studentFTIMScheduleBIndicatorStartIndex2627-1 : (studentFTIMScheduleBIndicatorStartIndex2627-1)+studentFTIMScheduleBIndicatorLength2627]), // Field # 876

		StudentFTIMScheduleDIndicator: preprocessString2627(record[studentFTIMScheduleDIndicatorStartIndex2627-1 : (studentFTIMScheduleDIndicatorStartIndex2627-1)+studentFTIMScheduleDIndicatorLength2627]), // Field # 877

		StudentFTIMScheduleEIndicator: preprocessString2627(record[studentFTIMScheduleEIndicatorStartIndex2627-1 : (studentFTIMScheduleEIndicatorStartIndex2627-1)+studentFTIMScheduleEIndicatorLength2627]), // Field # 878

		StudentFTIMScheduleFIndicator: preprocessString2627(record[studentFTIMScheduleFIndicatorStartIndex2627-1 : (studentFTIMScheduleFIndicatorStartIndex2627-1)+studentFTIMScheduleFIndicatorLength2627]), // Field # 879

		StudentFTIMScheduleHIndicator: preprocessString2627(record[studentFTIMScheduleHIndicatorStartIndex2627-1 : (studentFTIMScheduleHIndicatorStartIndex2627-1)+studentFTIMScheduleHIndicatorLength2627]), // Field # 880

		StudentFTIMIRSResponseCode: preprocessString2627(record[studentFTIMIRSResponseCodeStartIndex2627-1 : (studentFTIMIRSResponseCodeStartIndex2627-1)+studentFTIMIRSResponseCodeLength2627]), // Field # 881

		StudentFTIMSpouseReturnedTaxYear: preprocessString2627(record[studentFTIMSpouseReturnedTaxYearStartIndex2627-1 : (studentFTIMSpouseReturnedTaxYearStartIndex2627-1)+studentFTIMSpouseReturnedTaxYearLength2627]), // Field # 882

		StudentFTIMSpouseFilingStatusCode: preprocessString2627(record[studentFTIMSpouseFilingStatusCodeStartIndex2627-1 : (studentFTIMSpouseFilingStatusCodeStartIndex2627-1)+studentFTIMSpouseFilingStatusCodeLength2627]), // Field # 883

		StudentFTIMSpouseAdjustedGrossIncome: preprocessString2627(record[studentFTIMSpouseAdjustedGrossIncomeStartIndex2627-1 : (studentFTIMSpouseAdjustedGrossIncomeStartIndex2627-1)+studentFTIMSpouseAdjustedGrossIncomeLength2627]), // Field # 884

		StudentFTIMSpouseNumberOfExemptions: preprocessString2627(record[studentFTIMSpouseNumberOfExemptionsStartIndex2627-1 : (studentFTIMSpouseNumberOfExemptionsStartIndex2627-1)+studentFTIMSpouseNumberOfExemptionsLength2627]), // Field # 885

		StudentFTIMSpouseNumberOfDependents: preprocessString2627(record[studentFTIMSpouseNumberOfDependentsStartIndex2627-1 : (studentFTIMSpouseNumberOfDependentsStartIndex2627-1)+studentFTIMSpouseNumberOfDependentsLength2627]), // Field # 886

		StudentFTIMSpouseTotalIncomeEarnedAmount: preprocessString2627(record[studentFTIMSpouseTotalIncomeEarnedAmountStartIndex2627-1 : (studentFTIMSpouseTotalIncomeEarnedAmountStartIndex2627-1)+studentFTIMSpouseTotalIncomeEarnedAmountLength2627]), // Field # 887

		StudentFTIMSpouseIncomeTaxPaid: preprocessString2627(record[studentFTIMSpouseIncomeTaxPaidStartIndex2627-1 : (studentFTIMSpouseIncomeTaxPaidStartIndex2627-1)+studentFTIMSpouseIncomeTaxPaidLength2627]), // Field # 888

		StudentFTIMSpouseEducationCredits: preprocessString2627(record[studentFTIMSpouseEducationCreditsStartIndex2627-1 : (studentFTIMSpouseEducationCreditsStartIndex2627-1)+studentFTIMSpouseEducationCreditsLength2627]), // Field # 889

		StudentFTIMSpouseUntaxedIRADistributions: preprocessString2627(record[studentFTIMSpouseUntaxedIRADistributionsStartIndex2627-1 : (studentFTIMSpouseUntaxedIRADistributionsStartIndex2627-1)+studentFTIMSpouseUntaxedIRADistributionsLength2627]), // Field # 890

		StudentFTIMSpouseIRADeductibleAndPayments: preprocessString2627(record[studentFTIMSpouseIRADeductibleAndPaymentsStartIndex2627-1 : (studentFTIMSpouseIRADeductibleAndPaymentsStartIndex2627-1)+studentFTIMSpouseIRADeductibleAndPaymentsLength2627]), // Field # 891

		StudentFTIMSpouseTaxExemptInterest: preprocessString2627(record[studentFTIMSpouseTaxExemptInterestStartIndex2627-1 : (studentFTIMSpouseTaxExemptInterestStartIndex2627-1)+studentFTIMSpouseTaxExemptInterestLength2627]), // Field # 892

		StudentFTIMSpouseUntaxedPensionsAmount: preprocessString2627(record[studentFTIMSpouseUntaxedPensionsAmountStartIndex2627-1 : (studentFTIMSpouseUntaxedPensionsAmountStartIndex2627-1)+studentFTIMSpouseUntaxedPensionsAmountLength2627]), // Field # 893

		StudentFTIMSpouseScheduleCNetProfitLoss: preprocessString2627(record[studentFTIMSpouseScheduleCNetProfitLossStartIndex2627-1 : (studentFTIMSpouseScheduleCNetProfitLossStartIndex2627-1)+studentFTIMSpouseScheduleCNetProfitLossLength2627]), // Field # 894

		StudentFTIMSpouseScheduleAIndicator: preprocessString2627(record[studentFTIMSpouseScheduleAIndicatorStartIndex2627-1 : (studentFTIMSpouseScheduleAIndicatorStartIndex2627-1)+studentFTIMSpouseScheduleAIndicatorLength2627]), // Field # 895

		StudentFTIMSpouseScheduleBIndicator: preprocessString2627(record[studentFTIMSpouseScheduleBIndicatorStartIndex2627-1 : (studentFTIMSpouseScheduleBIndicatorStartIndex2627-1)+studentFTIMSpouseScheduleBIndicatorLength2627]), // Field # 896

		StudentFTIMSpouseScheduleDIndicator: preprocessString2627(record[studentFTIMSpouseScheduleDIndicatorStartIndex2627-1 : (studentFTIMSpouseScheduleDIndicatorStartIndex2627-1)+studentFTIMSpouseScheduleDIndicatorLength2627]), // Field # 897

		StudentFTIMSpouseScheduleEIndicator: preprocessString2627(record[studentFTIMSpouseScheduleEIndicatorStartIndex2627-1 : (studentFTIMSpouseScheduleEIndicatorStartIndex2627-1)+studentFTIMSpouseScheduleEIndicatorLength2627]), // Field # 898

		StudentFTIMSpouseScheduleFIndicator: preprocessString2627(record[studentFTIMSpouseScheduleFIndicatorStartIndex2627-1 : (studentFTIMSpouseScheduleFIndicatorStartIndex2627-1)+studentFTIMSpouseScheduleFIndicatorLength2627]), // Field # 899

		StudentFTIMSpouseScheduleHIndicator: preprocessString2627(record[studentFTIMSpouseScheduleHIndicatorStartIndex2627-1 : (studentFTIMSpouseScheduleHIndicatorStartIndex2627-1)+studentFTIMSpouseScheduleHIndicatorLength2627]), // Field # 900

		StudentFTIMSpouseIRSResponseCode: preprocessString2627(record[studentFTIMSpouseIRSResponseCodeStartIndex2627-1 : (studentFTIMSpouseIRSResponseCodeStartIndex2627-1)+studentFTIMSpouseIRSResponseCodeLength2627]), // Field # 901

		ParentFTIMReturnedTaxYear: preprocessString2627(record[parentFTIMReturnedTaxYearStartIndex2627-1 : (parentFTIMReturnedTaxYearStartIndex2627-1)+parentFTIMReturnedTaxYearLength2627]), // Field # 902

		ParentFTIMFilingStatusCode: preprocessString2627(record[parentFTIMFilingStatusCodeStartIndex2627-1 : (parentFTIMFilingStatusCodeStartIndex2627-1)+parentFTIMFilingStatusCodeLength2627]), // Field # 903

		ParentFTIMAdjustedGrossIncome: preprocessString2627(record[parentFTIMAdjustedGrossIncomeStartIndex2627-1 : (parentFTIMAdjustedGrossIncomeStartIndex2627-1)+parentFTIMAdjustedGrossIncomeLength2627]), // Field # 904

		ParentFTIMNumberOfExemptions: preprocessString2627(record[parentFTIMNumberOfExemptionsStartIndex2627-1 : (parentFTIMNumberOfExemptionsStartIndex2627-1)+parentFTIMNumberOfExemptionsLength2627]), // Field # 905

		ParentFTIMNumberOfDependents: preprocessString2627(record[parentFTIMNumberOfDependentsStartIndex2627-1 : (parentFTIMNumberOfDependentsStartIndex2627-1)+parentFTIMNumberOfDependentsLength2627]), // Field # 906

		ParentFTIMTotalIncomeEarnedAmount: preprocessString2627(record[parentFTIMTotalIncomeEarnedAmountStartIndex2627-1 : (parentFTIMTotalIncomeEarnedAmountStartIndex2627-1)+parentFTIMTotalIncomeEarnedAmountLength2627]), // Field # 907

		ParentFTIMIncomeTaxPaid: preprocessString2627(record[parentFTIMIncomeTaxPaidStartIndex2627-1 : (parentFTIMIncomeTaxPaidStartIndex2627-1)+parentFTIMIncomeTaxPaidLength2627]), // Field # 908

		ParentFTIMEducationCredits: preprocessString2627(record[parentFTIMEducationCreditsStartIndex2627-1 : (parentFTIMEducationCreditsStartIndex2627-1)+parentFTIMEducationCreditsLength2627]), // Field # 909

		ParentFTIMUntaxedIRADistributions: preprocessString2627(record[parentFTIMUntaxedIRADistributionsStartIndex2627-1 : (parentFTIMUntaxedIRADistributionsStartIndex2627-1)+parentFTIMUntaxedIRADistributionsLength2627]), // Field # 910

		ParentFTIMIRADeductibleAndPayments: preprocessString2627(record[parentFTIMIRADeductibleAndPaymentsStartIndex2627-1 : (parentFTIMIRADeductibleAndPaymentsStartIndex2627-1)+parentFTIMIRADeductibleAndPaymentsLength2627]), // Field # 911

		ParentFTIMTaxExemptInterest: preprocessString2627(record[parentFTIMTaxExemptInterestStartIndex2627-1 : (parentFTIMTaxExemptInterestStartIndex2627-1)+parentFTIMTaxExemptInterestLength2627]), // Field # 912

		ParentFTIMUntaxedPensionsAmount: preprocessString2627(record[parentFTIMUntaxedPensionsAmountStartIndex2627-1 : (parentFTIMUntaxedPensionsAmountStartIndex2627-1)+parentFTIMUntaxedPensionsAmountLength2627]), // Field # 913

		ParentFTIMScheduleCNetProfitLoss: preprocessString2627(record[parentFTIMScheduleCNetProfitLossStartIndex2627-1 : (parentFTIMScheduleCNetProfitLossStartIndex2627-1)+parentFTIMScheduleCNetProfitLossLength2627]), // Field # 914

		ParentFTIMScheduleAIndicator: preprocessString2627(record[parentFTIMScheduleAIndicatorStartIndex2627-1 : (parentFTIMScheduleAIndicatorStartIndex2627-1)+parentFTIMScheduleAIndicatorLength2627]), // Field # 915

		ParentFTIMScheduleBIndicator: preprocessString2627(record[parentFTIMScheduleBIndicatorStartIndex2627-1 : (parentFTIMScheduleBIndicatorStartIndex2627-1)+parentFTIMScheduleBIndicatorLength2627]), // Field # 916

		ParentFTIMScheduleDIndicator: preprocessString2627(record[parentFTIMScheduleDIndicatorStartIndex2627-1 : (parentFTIMScheduleDIndicatorStartIndex2627-1)+parentFTIMScheduleDIndicatorLength2627]), // Field # 917

		ParentFTIMScheduleEIndicator: preprocessString2627(record[parentFTIMScheduleEIndicatorStartIndex2627-1 : (parentFTIMScheduleEIndicatorStartIndex2627-1)+parentFTIMScheduleEIndicatorLength2627]), // Field # 918

		ParentFTIMScheduleFIndicator: preprocessString2627(record[parentFTIMScheduleFIndicatorStartIndex2627-1 : (parentFTIMScheduleFIndicatorStartIndex2627-1)+parentFTIMScheduleFIndicatorLength2627]), // Field # 919

		ParentFTIMScheduleHIndicator: preprocessString2627(record[parentFTIMScheduleHIndicatorStartIndex2627-1 : (parentFTIMScheduleHIndicatorStartIndex2627-1)+parentFTIMScheduleHIndicatorLength2627]), // Field # 920

		ParentFTIMIRSResponseCode: preprocessString2627(record[parentFTIMIRSResponseCodeStartIndex2627-1 : (parentFTIMIRSResponseCodeStartIndex2627-1)+parentFTIMIRSResponseCodeLength2627]), // Field # 921

		ParentFTIMSpouseReturnedTaxYear: preprocessString2627(record[parentFTIMSpouseReturnedTaxYearStartIndex2627-1 : (parentFTIMSpouseReturnedTaxYearStartIndex2627-1)+parentFTIMSpouseReturnedTaxYearLength2627]), // Field # 922

		ParentFTIMSpouseFilingStatusCode: preprocessString2627(record[parentFTIMSpouseFilingStatusCodeStartIndex2627-1 : (parentFTIMSpouseFilingStatusCodeStartIndex2627-1)+parentFTIMSpouseFilingStatusCodeLength2627]), // Field # 923

		ParentFTIMSpouseAdjustedGrossIncome: preprocessString2627(record[parentFTIMSpouseAdjustedGrossIncomeStartIndex2627-1 : (parentFTIMSpouseAdjustedGrossIncomeStartIndex2627-1)+parentFTIMSpouseAdjustedGrossIncomeLength2627]), // Field # 924

		ParentFTIMSpouseNumberOfExemptions: preprocessString2627(record[parentFTIMSpouseNumberOfExemptionsStartIndex2627-1 : (parentFTIMSpouseNumberOfExemptionsStartIndex2627-1)+parentFTIMSpouseNumberOfExemptionsLength2627]), // Field # 925

		ParentFTIMSpouseNumberOfDependents: preprocessString2627(record[parentFTIMSpouseNumberOfDependentsStartIndex2627-1 : (parentFTIMSpouseNumberOfDependentsStartIndex2627-1)+parentFTIMSpouseNumberOfDependentsLength2627]), // Field # 926

		ParentFTIMSpouseTotalIncomeEarnedAmount: preprocessString2627(record[parentFTIMSpouseTotalIncomeEarnedAmountStartIndex2627-1 : (parentFTIMSpouseTotalIncomeEarnedAmountStartIndex2627-1)+parentFTIMSpouseTotalIncomeEarnedAmountLength2627]), // Field # 927

		ParentFTIMSpouseIncomeTaxPaid: preprocessString2627(record[parentFTIMSpouseIncomeTaxPaidStartIndex2627-1 : (parentFTIMSpouseIncomeTaxPaidStartIndex2627-1)+parentFTIMSpouseIncomeTaxPaidLength2627]), // Field # 928

		ParentFTIMSpouseEducationCredits: preprocessString2627(record[parentFTIMSpouseEducationCreditsStartIndex2627-1 : (parentFTIMSpouseEducationCreditsStartIndex2627-1)+parentFTIMSpouseEducationCreditsLength2627]), // Field # 929

		ParentFTIMSpouseUntaxedIRADistributions: preprocessString2627(record[parentFTIMSpouseUntaxedIRADistributionsStartIndex2627-1 : (parentFTIMSpouseUntaxedIRADistributionsStartIndex2627-1)+parentFTIMSpouseUntaxedIRADistributionsLength2627]), // Field # 930

		ParentFTIMSpouseIRADeductibleAndPayments: preprocessString2627(record[parentFTIMSpouseIRADeductibleAndPaymentsStartIndex2627-1 : (parentFTIMSpouseIRADeductibleAndPaymentsStartIndex2627-1)+parentFTIMSpouseIRADeductibleAndPaymentsLength2627]), // Field # 931

		ParentFTIMSpouseTaxExemptInterest: preprocessString2627(record[parentFTIMSpouseTaxExemptInterestStartIndex2627-1 : (parentFTIMSpouseTaxExemptInterestStartIndex2627-1)+parentFTIMSpouseTaxExemptInterestLength2627]), // Field # 932

		ParentFTIMSpouseUntaxedPensionsAmount: preprocessString2627(record[parentFTIMSpouseUntaxedPensionsAmountStartIndex2627-1 : (parentFTIMSpouseUntaxedPensionsAmountStartIndex2627-1)+parentFTIMSpouseUntaxedPensionsAmountLength2627]), // Field # 933

		ParentFTIMSpouseScheduleCNetProfitLoss: preprocessString2627(record[parentFTIMSpouseScheduleCNetProfitLossStartIndex2627-1 : (parentFTIMSpouseScheduleCNetProfitLossStartIndex2627-1)+parentFTIMSpouseScheduleCNetProfitLossLength2627]), // Field # 934

		ParentFTIMSpouseScheduleAIndicator: preprocessString2627(record[parentFTIMSpouseScheduleAIndicatorStartIndex2627-1 : (parentFTIMSpouseScheduleAIndicatorStartIndex2627-1)+parentFTIMSpouseScheduleAIndicatorLength2627]), // Field # 935

		ParentFTIMSpouseScheduleBIndicator: preprocessString2627(record[parentFTIMSpouseScheduleBIndicatorStartIndex2627-1 : (parentFTIMSpouseScheduleBIndicatorStartIndex2627-1)+parentFTIMSpouseScheduleBIndicatorLength2627]), // Field # 936

		ParentFTIMSpouseScheduleDIndicator: preprocessString2627(record[parentFTIMSpouseScheduleDIndicatorStartIndex2627-1 : (parentFTIMSpouseScheduleDIndicatorStartIndex2627-1)+parentFTIMSpouseScheduleDIndicatorLength2627]), // Field # 937

		ParentFTIMSpouseScheduleEIndicator: preprocessString2627(record[parentFTIMSpouseScheduleEIndicatorStartIndex2627-1 : (parentFTIMSpouseScheduleEIndicatorStartIndex2627-1)+parentFTIMSpouseScheduleEIndicatorLength2627]), // Field # 938

		ParentFTIMSpouseScheduleFIndicator: preprocessString2627(record[parentFTIMSpouseScheduleFIndicatorStartIndex2627-1 : (parentFTIMSpouseScheduleFIndicatorStartIndex2627-1)+parentFTIMSpouseScheduleFIndicatorLength2627]), // Field # 939

		ParentFTIMSpouseScheduleHIndicator: preprocessString2627(record[parentFTIMSpouseScheduleHIndicatorStartIndex2627-1 : (parentFTIMSpouseScheduleHIndicatorStartIndex2627-1)+parentFTIMSpouseScheduleHIndicatorLength2627]), // Field # 940

		ParentFTIMSpouseIRSResponseCode: preprocessString2627(record[parentFTIMSpouseIRSResponseCodeStartIndex2627-1 : (parentFTIMSpouseIRSResponseCodeStartIndex2627-1)+parentFTIMSpouseIRSResponseCodeLength2627]), // Field # 941

		//FTILabelEnd: preprocessString2627(record[ftiLabelEndStartIndex2627-1 : (ftiLabelEndStartIndex2627-1)+ftiLabelEndLength2627]), // Field # 942

		StudentTotalIncome: preprocessString2627(record[studentTotalIncomeStartIndex2627-1 : (studentTotalIncomeStartIndex2627-1)+studentTotalIncomeLength2627]), // Field # 944

		ParentTotalIncome: preprocessString2627(record[parentTotalIncomeStartIndex2627-1 : (parentTotalIncomeStartIndex2627-1)+parentTotalIncomeLength2627]), // Field # 945

		FISAPTotalIncome: preprocessString2627(record[fisapTotalIncomeStartIndex2627-1 : (fisapTotalIncomeStartIndex2627-1)+fisapTotalIncomeLength2627]), // Field # 946
	}
	//</editor-fold>
	return r, nil
}

// preprocessString2627 trims whitespace and checks for any indications of a null or missing value from FSA, returning an empty string if so
// in order for the zero value to be used in the struct.  For example, "N/A" is a common value used by FSA to indicate a null value.
func preprocessString2627(s string) string {
	p := strings.TrimSpace(s)
	if p == "N/A" {
		p = ""
	}

	return p
}

func parseISIRDate2627(s string) *time.Time {
	parsedDate, err := time.Parse(isirDateLayout2627, s)

	if err != nil {
		return nil
	}

	return &parsedDate
}

func parseISIRDateShort2627(s string) *time.Time {
	parsedDate, err := time.Parse(isirDateShortLayout2627, s)

	if err != nil {
		return nil
	}

	return &parsedDate
}
