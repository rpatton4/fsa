package isirparser

import (
	"errors"
	"fmt"
	"github.com/rpatton4/fsa/pkg/isirmodels"
	"log/slog"
	"strings"
	"time"
)

const isirDateLayout2526 = "20060102"
const isirDateShortLayout2526 = "200601"

const totalISIRLength2526 int = 7704

// Field # 1
const yearIndicatorStartIndex2526 int = 1
const yearIndicatorLength2526 int = 1

// Field # 2
const fafsaUUIDStartIndex2526 int = 2
const fafsaUUIDLength2526 int = 36

// Field # 3
const transactionUUIDStartIndex2526 int = 38
const transactionUUIDLength2526 int = 36

// Field # 4
const personUUIDStartIndex2526 int = 74
const personUUIDLength2526 int = 36

// Field # 5
const transactionNumberStartIndex2526 int = 110
const transactionNumberLength2526 int = 2

// Field # 6
const dependencyModelStartIndex2526 int = 112
const dependencyModelLength2526 int = 1

// Field # 7
const applicationSourceStartIndex2526 int = 113
const applicationSourceLength2526 int = 1

// Field # 8
const applicationReceiptDateStartIndex2526 int = 114
const applicationReceiptDateLength2526 int = 8

// Field # 9
const transactionSourceStartIndex2526 int = 122
const transactionSourceLength2526 int = 1

// Field # 10
const transactionTypeStartIndex2526 int = 123
const transactionTypeLength2526 int = 1

// Field # 11
const transactionLanguageStartIndex2526 int = 124
const transactionLanguageLength2526 int = 1

// Field # 12
const transactionReceiptDateStartIndex2526 int = 125
const transactionReceiptDateLength2526 int = 8

// Field # 13
const transactionProcessedDateStartIndex2526 int = 133
const transactionProcessedDateLength2526 int = 8

// Field # 14
const transactionStatusStartIndex2526 int = 141
const transactionStatusLength2526 int = 30

// Field # 15
const renewalDataUsedStartIndex2526 int = 171
const renewalDataUsedLength2526 int = 3

// Field # 16
const fpsCorrectionReasonStartIndex2526 int = 174
const fpsCorrectionReasonLength2526 int = 1

// Field # 17
const saiChangeFlagStartIndex2526 int = 175
const saiChangeFlagLength2526 int = 1

// Field # 18
const saiStartIndex2526 int = 176
const saiLength2526 int = 6

// Field # 19
const provisionalSAIStartIndex2526 int = 182
const provisionalSAILength2526 int = 6

// Field # 20
const saiFormulaStartIndex2526 int = 188
const saiFormulaLength2526 int = 1

// Field # 21
const saiComputationTypeStartIndex2526 int = 189
const saiComputationTypeLength2526 int = 2

// Field # 22
const maxPellIndicatorStartIndex2526 int = 191
const maxPellIndicatorLength2526 int = 1

// Field # 23
const minimumPellIndicatorStartIndex2526 int = 192
const minimumPellIndicatorLength2526 int = 1

// Field # 25
const studentFirstNameStartIndex2526 int = 243
const studentFirstNameLength2526 int = 35

// Field # 26
const studentMiddleNameStartIndex2526 int = 278
const studentMiddleNameLength2526 int = 15

// Field # 27
const studentLastNameStartIndex2526 int = 293
const studentLastNameLength2526 int = 35

// Field # 28
const studentSuffixStartIndex2526 int = 328
const studentSuffixLength2526 int = 10

// Field # 29
const studentDateOfBirthStartIndex2526 int = 338
const studentDateOfBirthLength2526 int = 8

// Field # 30
const studentSSNStartIndex2526 int = 346
const studentSSNLength2526 int = 9

// Field # 31
const studentITINStartIndex2526 int = 355
const studentITINLength2526 int = 9

// Field # 32
const studentPhoneNumberStartIndex2526 int = 364
const studentPhoneNumberLength2526 int = 10

// Field # 33
const studentEmailAddressStartIndex2526 int = 374
const studentEmailAddressLength2526 int = 50

// Field # 34
const studentStreetAddressStartIndex2526 int = 424
const studentStreetAddressLength2526 int = 40

// Field # 35
const studentCityStartIndex2526 int = 464
const studentCityLength2526 int = 30

// Field # 36
const studentStateStartIndex2526 int = 494
const studentStateLength2526 int = 2

// Field # 37
const studentZipCodeStartIndex2526 int = 496
const studentZipCodeLength2526 int = 10

// Field # 38
const studentCountryStartIndex2526 int = 506
const studentCountryLength2526 int = 2

// Field # 40
const studentMaritalStatusStartIndex2526 int = 558
const studentMaritalStatusLength2526 int = 1

// Field # 41
const studentGradeLevelStartIndex2526 int = 559
const studentGradeLevelLength2526 int = 1

// Field # 42
const studentFirstBachelorsDegreeBefore2526StartIndex2526 int = 560
const studentFirstBachelorsDegreeBefore2526Length2526 int = 1

// Field # 43
const studentPursuingTeacherCertificationStartIndex2526 int = 561
const studentPursuingTeacherCertificationLength2526 int = 1

// Field # 44
const studentActiveDutyStartIndex2526 int = 562
const studentActiveDutyLength2526 int = 1

// Field # 45
const studentVeteranStartIndex2526 int = 563
const studentVeteranLength2526 int = 1

// Field # 46
const studentChildOrOtherDependentsStartIndex2526 int = 564
const studentChildOrOtherDependentsLength2526 int = 1

// Field # 47
const studentParentsDeceasedStartIndex2526 int = 565
const studentParentsDeceasedLength2526 int = 1

// Field # 48
const studentWardOfCourtStartIndex2526 int = 566
const studentWardOfCourtLength2526 int = 1

// Field # 49
const studentInFosterCareStartIndex2526 int = 567
const studentInFosterCareLength2526 int = 1

// Field # 50
const studentEmancipatedMinorStartIndex2526 int = 568
const studentEmancipatedMinorLength2526 int = 1

// Field # 51
const studentLegalGuardianshipStartIndex2526 int = 569
const studentLegalGuardianshipLength2526 int = 1

// Field # 52
const studentPersonalCircumstancesNoneOfTheAboveStartIndex2526 int = 570
const studentPersonalCircumstancesNoneOfTheAboveLength2526 int = 1

// Field # 53
const studentUnaccompaniedHomelessYouthAndSelfSupportingStartIndex2526 int = 571
const studentUnaccompaniedHomelessYouthAndSelfSupportingLength2526 int = 1

// Field # 54
const studentUnaccompaniedHomelessGeneralStartIndex2526 int = 572
const studentUnaccompaniedHomelessGeneralLength2526 int = 1

// Field # 55
const studentUnaccompaniedHomelessHSStartIndex2526 int = 573
const studentUnaccompaniedHomelessHSLength2526 int = 1

// Field # 56
const studentUnaccompaniedHomelessTRIOStartIndex2526 int = 574
const studentUnaccompaniedHomelessTRIOLength2526 int = 1

// Field # 57
const studentUnaccompaniedHomelessFAAStartIndex2526 int = 575
const studentUnaccompaniedHomelessFAALength2526 int = 1

// Field # 58
const studentHomelessnessNoneOfTheAboveStartIndex2526 int = 576
const studentHomelessnessNoneOfTheAboveLength2526 int = 1

// Field # 59
const studentUnusualCircumstanceStartIndex2526 int = 577
const studentUnusualCircumstanceLength2526 int = 1

// Field # 60
const studentUnsubOnlyStartIndex2526 int = 578
const studentUnsubOnlyLength2526 int = 1

// Field # 61
const studentUpdatedFamilySizeStartIndex2526 int = 579
const studentUpdatedFamilySizeLength2526 int = 2

// Field # 62
const studentNumberInCollegeStartIndex2526 int = 581
const studentNumberInCollegeLength2526 int = 2

// Field # 63
const studentCitizenshipStatusStartIndex2526 int = 583
const studentCitizenshipStatusLength2526 int = 1

// Field # 64
const studentANumberStartIndex2526 int = 584
const studentANumberLength2526 int = 9

// Field # 65
const studentStateOfLegalResidenceStartIndex2526 int = 593
const studentStateOfLegalResidenceLength2526 int = 2

// Field # 66
const studentLegalResidenceDateStartIndex2526 int = 595
const studentLegalResidenceDateLength2526 int = 6

// Field # 67
const studentEitherParentAttendCollegeStartIndex2526 int = 601
const studentEitherParentAttendCollegeLength2526 int = 1

// Field # 68
const studentParentKilledInTheLineOfDutyStartIndex2526 int = 602
const studentParentKilledInTheLineOfDutyLength2526 int = 1

// Field # 69
const studentHighSchoolCompletionStatusStartIndex2526 int = 603
const studentHighSchoolCompletionStatusLength2526 int = 1

// Field # 70
const studentHighSchoolNameStartIndex2526 int = 604
const studentHighSchoolNameLength2526 int = 60

// Field # 71
const studentHighSchoolCityStartIndex2526 int = 664
const studentHighSchoolCityLength2526 int = 28

// Field # 72
const studentHighSchoolStateStartIndex2526 int = 692
const studentHighSchoolStateLength2526 int = 2

// Field # 73
const studentHighSchoolEquivalentDiplomaNameStartIndex2526 int = 694
const studentHighSchoolEquivalentDiplomaNameLength2526 int = 1

// Field # 74
const studentHighSchoolEquivalentDiplomaStateStartIndex2526 int = 695
const studentHighSchoolEquivalentDiplomaStateLength2526 int = 2

// Field # 75
const studentManuallyEnteredReceivedEITCStartIndex2526 int = 697
const studentManuallyEnteredReceivedEITCLength2526 int = 1

// Field # 76
const studentManuallyEnteredReceivedFederalHousingAssistanceStartIndex2526 int = 698
const studentManuallyEnteredReceivedFederalHousingAssistanceLength2526 int = 1

// Field # 77
const studentManuallyEnteredReceivedFreeReducedPriceLunchStartIndex2526 int = 699
const studentManuallyEnteredReceivedFreeReducedPriceLunchLength2526 int = 1

// Field # 78
const studentManuallyEnteredReceivedMedicaidStartIndex2526 int = 700
const studentManuallyEnteredReceivedMedicaidLength2526 int = 1

// Field # 79
const studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanStartIndex2526 int = 701
const studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanLength2526 int = 1

// Field # 80
const studentManuallyEnteredReceivedSNAPStartIndex2526 int = 702
const studentManuallyEnteredReceivedSNAPLength2526 int = 1

// Field # 81
const studentManuallyEnteredReceivedSupplementalSecurityIncomeStartIndex2526 int = 703
const studentManuallyEnteredReceivedSupplementalSecurityIncomeLength2526 int = 1

// Field # 82
const studentManuallyEnteredReceivedTANFStartIndex2526 int = 704
const studentManuallyEnteredReceivedTANFLength2526 int = 1

// Field # 83
const studentManuallyEnteredReceivedWICStartIndex2526 int = 705
const studentManuallyEnteredReceivedWICLength2526 int = 1

// Field # 84
const studentManuallyEnteredFederalBenefitsNoneOfTheAboveStartIndex2526 int = 706
const studentManuallyEnteredFederalBenefitsNoneOfTheAboveLength2526 int = 1

// Field # 85
const studentManuallyEnteredFiled1040Or1040NRStartIndex2526 int = 707
const studentManuallyEnteredFiled1040Or1040NRLength2526 int = 1

// Field # 86
const studentManuallyEnteredFiledNonUSTaxReturnStartIndex2526 int = 708
const studentManuallyEnteredFiledNonUSTaxReturnLength2526 int = 1

// Field # 87
const studentManuallyEnteredFiledJointReturnWithCurrentSpouseStartIndex2526 int = 709
const studentManuallyEnteredFiledJointReturnWithCurrentSpouseLength2526 int = 1

// Field # 88
const studentManuallyEnteredTaxReturnFilingStatusStartIndex2526 int = 710
const studentManuallyEnteredTaxReturnFilingStatusLength2526 int = 1

// Field # 89
const studentManuallyEnteredIncomeEarnedFromWorkStartIndex2526 int = 711
const studentManuallyEnteredIncomeEarnedFromWorkLength2526 int = 11

// Field # 90
const studentManuallyEnteredTaxExemptInterestIncomeStartIndex2526 int = 722
const studentManuallyEnteredTaxExemptInterestIncomeLength2526 int = 11

// Field # 91
const studentManuallyEnteredUntaxedPortionsOfIRADistributionsStartIndex2526 int = 733
const studentManuallyEnteredUntaxedPortionsOfIRADistributionsLength2526 int = 11

// Field # 92
const studentManuallyEnteredIRARolloverStartIndex2526 int = 744
const studentManuallyEnteredIRARolloverLength2526 int = 11

// Field # 93
const studentManuallyEnteredUntaxedPortionsOfPensionsStartIndex2526 int = 755
const studentManuallyEnteredUntaxedPortionsOfPensionsLength2526 int = 11

// Field # 94
const studentManuallyEnteredPensionRolloverStartIndex2526 int = 766
const studentManuallyEnteredPensionRolloverLength2526 int = 11

// Field # 95
const studentManuallyEnteredAdjustedGrossIncomeStartIndex2526 int = 777
const studentManuallyEnteredAdjustedGrossIncomeLength2526 int = 10

// Field # 96
const studentManuallyEnteredIncomeTaxPaidStartIndex2526 int = 787
const studentManuallyEnteredIncomeTaxPaidLength2526 int = 9

// Field # 97
const studentManuallyEnteredEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex2526 int = 796
const studentManuallyEnteredEarnedIncomeTaxCreditReceivedDuringTaxYearLength2526 int = 1

// Field # 98
const studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherStartIndex2526 int = 797
const studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherLength2526 int = 11

// Field # 99
const studentManuallyEnteredEducationCreditsStartIndex2526 int = 808
const studentManuallyEnteredEducationCreditsLength2526 int = 9

// Field # 100
const studentManuallyEnteredFiledScheduleABDEFHStartIndex2526 int = 817
const studentManuallyEnteredFiledScheduleABDEFHLength2526 int = 1

// Field # 101
const studentManuallyEnteredScheduleCAmountStartIndex2526 int = 818
const studentManuallyEnteredScheduleCAmountLength2526 int = 12

// Field # 102
const studentManuallyEnteredCollegeGrantAndScholarshipAidStartIndex2526 int = 830
const studentManuallyEnteredCollegeGrantAndScholarshipAidLength2526 int = 7

// Field # 103
const studentManuallyEnteredForeignEarnedIncomeExclusionStartIndex2526 int = 837
const studentManuallyEnteredForeignEarnedIncomeExclusionLength2526 int = 10

// Field # 104
const studentManuallyEnteredChildSupportReceivedStartIndex2526 int = 847
const studentManuallyEnteredChildSupportReceivedLength2526 int = 7

// Field # 105
const studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsStartIndex2526 int = 854
const studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsLength2526 int = 7

// Field # 106
const studentManuallyEnteredNetWorthOfCurrentInvestmentsStartIndex2526 int = 861
const studentManuallyEnteredNetWorthOfCurrentInvestmentsLength2526 int = 7

// Field # 107
const studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsStartIndex2526 int = 868
const studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsLength2526 int = 7

// Field # 108
const studentCollege1StartIndex2526 int = 875
const studentCollege1Length2526 int = 6

// Field # 109
const studentCollege2StartIndex2526 int = 881
const studentCollege2Length2526 int = 6

// Field # 110
const studentCollege3StartIndex2526 int = 887
const studentCollege3Length2526 int = 6

// Field # 111
const studentCollege4StartIndex2526 int = 893
const studentCollege4Length2526 int = 6

// Field # 112
const studentCollege5StartIndex2526 int = 899
const studentCollege5Length2526 int = 6

// Field # 113
const studentCollege6StartIndex2526 int = 905
const studentCollege6Length2526 int = 6

// Field # 114
const studentCollege7StartIndex2526 int = 911
const studentCollege7Length2526 int = 6

// Field # 115
const studentCollege8StartIndex2526 int = 917
const studentCollege8Length2526 int = 6

// Field # 116
const studentCollege9StartIndex2526 int = 923
const studentCollege9Length2526 int = 6

// Field # 117
const studentCollege10StartIndex2526 int = 929
const studentCollege10Length2526 int = 6

// Field # 118
const studentCollege11StartIndex2526 int = 935
const studentCollege11Length2526 int = 6

// Field # 119
const studentCollege12StartIndex2526 int = 941
const studentCollege12Length2526 int = 6

// Field # 120
const studentCollege13StartIndex2526 int = 947
const studentCollege13Length2526 int = 6

// Field # 121
const studentCollege14StartIndex2526 int = 953
const studentCollege14Length2526 int = 6

// Field # 122
const studentCollege15StartIndex2526 int = 959
const studentCollege15Length2526 int = 6

// Field # 123
const studentCollege16StartIndex2526 int = 965
const studentCollege16Length2526 int = 6

// Field # 124
const studentCollege17StartIndex2526 int = 971
const studentCollege17Length2526 int = 6

// Field # 125
const studentCollege18StartIndex2526 int = 977
const studentCollege18Length2526 int = 6

// Field # 126
const studentCollege19StartIndex2526 int = 983
const studentCollege19Length2526 int = 6

// Field # 127
const studentCollege20StartIndex2526 int = 989
const studentCollege20Length2526 int = 6

// Field # 128
const studentConsentToRetrieveAndDiscloseFTIStartIndex2526 int = 995
const studentConsentToRetrieveAndDiscloseFTILength2526 int = 1

// Field # 129
const studentSignatureStartIndex2526 int = 996
const studentSignatureLength2526 int = 1

// Field # 130
const studentSignatureDateStartIndex2526 int = 997
const studentSignatureDateLength2526 int = 8

// Field # 132
const studentSpouseFirstNameStartIndex2526 int = 1055
const studentSpouseFirstNameLength2526 int = 35

// Field # 133
const studentSpouseMiddleNameStartIndex2526 int = 1090
const studentSpouseMiddleNameLength2526 int = 15

// Field # 134
const studentSpouseLastNameStartIndex2526 int = 1105
const studentSpouseLastNameLength2526 int = 35

// Field # 135
const studentSpouseSuffixStartIndex2526 int = 1140
const studentSpouseSuffixLength2526 int = 10

// Field # 136
const studentSpouseDateOfBirthStartIndex2526 int = 1150
const studentSpouseDateOfBirthLength2526 int = 8

// Field # 137
const studentSpouseSSNStartIndex2526 int = 1158
const studentSpouseSSNLength2526 int = 9

// Field # 138
const studentSpouseITINStartIndex2526 int = 1167
const studentSpouseITINLength2526 int = 9

// Field # 139
const studentSpousePhoneNumberStartIndex2526 int = 1176
const studentSpousePhoneNumberLength2526 int = 10

// Field # 140
const studentSpouseEmailAddressStartIndex2526 int = 1186
const studentSpouseEmailAddressLength2526 int = 50

// Field # 141
const studentSpouseStreetAddressStartIndex2526 int = 1236
const studentSpouseStreetAddressLength2526 int = 40

// Field # 142
const studentSpouseCityStartIndex2526 int = 1276
const studentSpouseCityLength2526 int = 30

// Field # 143
const studentSpouseStateStartIndex2526 int = 1306
const studentSpouseStateLength2526 int = 2

// Field # 144
const studentSpouseZipCodeStartIndex2526 int = 1308
const studentSpouseZipCodeLength2526 int = 10

// Field # 145
const studentSpouseCountryStartIndex2526 int = 1318
const studentSpouseCountryLength2526 int = 2

// Field # 146
const studentSpouseFiled1040Or1040NRStartIndex2526 int = 1320
const studentSpouseFiled1040Or1040NRLength2526 int = 1

// Field # 147
const studentSpouseFiledNonUSTaxReturnStartIndex2526 int = 1321
const studentSpouseFiledNonUSTaxReturnLength2526 int = 1

// Field # 148
const studentSpouseTaxReturnFilingStatusStartIndex2526 int = 1322
const studentSpouseTaxReturnFilingStatusLength2526 int = 1

// Field # 149
const studentSpouseIncomeEarnedFromWorkStartIndex2526 int = 1323
const studentSpouseIncomeEarnedFromWorkLength2526 int = 11

// Field # 150
const studentSpouseTaxExemptInterestIncomeStartIndex2526 int = 1334
const studentSpouseTaxExemptInterestIncomeLength2526 int = 11

// Field # 151
const studentSpouseUntaxedPortionsOfIRADistributionsStartIndex2526 int = 1345
const studentSpouseUntaxedPortionsOfIRADistributionsLength2526 int = 11

// Field # 152
const studentSpouseIRARolloverStartIndex2526 int = 1356
const studentSpouseIRARolloverLength2526 int = 11

// Field # 153
const studentSpouseUntaxedPortionsOfPensionsStartIndex2526 int = 1367
const studentSpouseUntaxedPortionsOfPensionsLength2526 int = 11

// Field # 154
const studentSpousePensionRolloverStartIndex2526 int = 1378
const studentSpousePensionRolloverLength2526 int = 11

// Field # 155
const studentSpouseAdjustedGrossIncomeStartIndex2526 int = 1389
const studentSpouseAdjustedGrossIncomeLength2526 int = 10

// Field # 156
const studentSpouseIncomeTaxPaidStartIndex2526 int = 1399
const studentSpouseIncomeTaxPaidLength2526 int = 9

// Field # 157
const studentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex2526 int = 1408
const studentSpouseDeductiblePaymentsToIRAKeoghOtherLength2526 int = 11

// Field # 158
const studentSpouseEducationCreditsStartIndex2526 int = 1419
const studentSpouseEducationCreditsLength2526 int = 9

// Field # 159
const studentSpouseFiledScheduleABDEFHStartIndex2526 int = 1428
const studentSpouseFiledScheduleABDEFHLength2526 int = 1

// Field # 160
const studentSpouseScheduleCAmountStartIndex2526 int = 1429
const studentSpouseScheduleCAmountLength2526 int = 12

// Field # 161
const studentSpouseForeignEarnedIncomeExclusionStartIndex2526 int = 1441
const studentSpouseForeignEarnedIncomeExclusionLength2526 int = 10

// Field # 162
const studentSpouseConsentToRetrieveAndDiscloseFTIStartIndex2526 int = 1451
const studentSpouseConsentToRetrieveAndDiscloseFTILength2526 int = 1

// Field # 163
const studentSpouseSignatureStartIndex2526 int = 1452
const studentSpouseSignatureLength2526 int = 1

// Field # 164
const studentSpouseSignatureDateStartIndex2526 int = 1453
const studentSpouseSignatureDateLength2526 int = 8

// Field # 166
const parentFirstNameStartIndex2526 int = 1511
const parentFirstNameLength2526 int = 35

// Field # 167
const parentMiddleNameStartIndex2526 int = 1546
const parentMiddleNameLength2526 int = 15

// Field # 168
const parentLastNameStartIndex2526 int = 1561
const parentLastNameLength2526 int = 35

// Field # 169
const parentSuffixStartIndex2526 int = 1596
const parentSuffixLength2526 int = 10

// Field # 170
const parentDateOfBirthStartIndex2526 int = 1606
const parentDateOfBirthLength2526 int = 8

// Field # 171
const parentSSNStartIndex2526 int = 1614
const parentSSNLength2526 int = 9

// Field # 172
const parentITINStartIndex2526 int = 1623
const parentITINLength2526 int = 9

// Field # 173
const parentPhoneNumberStartIndex2526 int = 1632
const parentPhoneNumberLength2526 int = 10

// Field # 174
const parentEmailAddressStartIndex2526 int = 1642
const parentEmailAddressLength2526 int = 50

// Field # 175
const parentStreetAddressStartIndex2526 int = 1692
const parentStreetAddressLength2526 int = 40

// Field # 176
const parentCityStartIndex2526 int = 1732
const parentCityLength2526 int = 30

// Field # 177
const parentStateStartIndex2526 int = 1762
const parentStateLength2526 int = 2

// Field # 178
const parentZipCodeStartIndex2526 int = 1764
const parentZipCodeLength2526 int = 10

// Field # 179
const parentCountryStartIndex2526 int = 1774
const parentCountryLength2526 int = 2

// Field # 180
const parentMaritalStatusStartIndex2526 int = 1776
const parentMaritalStatusLength2526 int = 1

// Field # 181
const parentStateOfLegalResidenceStartIndex2526 int = 1777
const parentStateOfLegalResidenceLength2526 int = 2

// Field # 182
const parentLegalResidenceDateStartIndex2526 int = 1779
const parentLegalResidenceDateLength2526 int = 6

// Field # 183
const parentUpdatedFamilySizeStartIndex2526 int = 1785
const parentUpdatedFamilySizeLength2526 int = 2

// Field # 184
const parentNumberInCollegeStartIndex2526 int = 1787
const parentNumberInCollegeLength2526 int = 2

// Field # 185
const parentReceivedEITCStartIndex2526 int = 1789
const parentReceivedEITCLength2526 int = 1

// Field # 186
const parentReceivedFederalHousingAssistanceStartIndex2526 int = 1790
const parentReceivedFederalHousingAssistanceLength2526 int = 1

// Field # 187
const parentReceivedFreeReducedPriceLunchStartIndex2526 int = 1791
const parentReceivedFreeReducedPriceLunchLength2526 int = 1

// Field # 188
const parentReceivedMedicaidStartIndex2526 int = 1792
const parentReceivedMedicaidLength2526 int = 1

// Field # 189
const parentReceivedRefundableCreditFor36BHealthPlanStartIndex2526 int = 1793
const parentReceivedRefundableCreditFor36BHealthPlanLength2526 int = 1

// Field # 190
const parentReceivedSNAPStartIndex2526 int = 1794
const parentReceivedSNAPLength2526 int = 1

// Field # 191
const parentReceivedSupplementalSecurityIncomeStartIndex2526 int = 1795
const parentReceivedSupplementalSecurityIncomeLength2526 int = 1

// Field # 192
const parentReceivedTANFStartIndex2526 int = 1796
const parentReceivedTANFLength2526 int = 1

// Field # 193
const parentReceivedWICStartIndex2526 int = 1797
const parentReceivedWICLength2526 int = 1

// Field # 194
const parentFederalBenefitsNoneOfTheAboveStartIndex2526 int = 1798
const parentFederalBenefitsNoneOfTheAboveLength2526 int = 1

// Field # 195
const parentFiled1040Or1040NRStartIndex2526 int = 1799
const parentFiled1040Or1040NRLength2526 int = 1

// Field # 196
const parentFileNonUSTaxReturnStartIndex2526 int = 1800
const parentFileNonUSTaxReturnLength2526 int = 1

// Field # 197
const parentFiledJointReturnWithCurrentSpouseStartIndex2526 int = 1801
const parentFiledJointReturnWithCurrentSpouseLength2526 int = 1

// Field # 198
const parentTaxReturnFilingStatusStartIndex2526 int = 1802
const parentTaxReturnFilingStatusLength2526 int = 1

// Field # 199
const parentIncomeEarnedFromWorkStartIndex2526 int = 1803
const parentIncomeEarnedFromWorkLength2526 int = 11

// Field # 200
const parentTaxExemptInterestIncomeStartIndex2526 int = 1814
const parentTaxExemptInterestIncomeLength2526 int = 11

// Field # 201
const parentUntaxedPortionsOfIRADistributionsStartIndex2526 int = 1825
const parentUntaxedPortionsOfIRADistributionsLength2526 int = 11

// Field # 202
const parentIRARolloverStartIndex2526 int = 1836
const parentIRARolloverLength2526 int = 11

// Field # 203
const parentUntaxedPortionsOfPensionsStartIndex2526 int = 1847
const parentUntaxedPortionsOfPensionsLength2526 int = 11

// Field # 204
const parentPensionRolloverStartIndex2526 int = 1858
const parentPensionRolloverLength2526 int = 11

// Field # 205
const parentAdjustedGrossIncomeStartIndex2526 int = 1869
const parentAdjustedGrossIncomeLength2526 int = 10

// Field # 206
const parentIncomeTaxPaidStartIndex2526 int = 1879
const parentIncomeTaxPaidLength2526 int = 9

// Field # 207
const parentEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex2526 int = 1888
const parentEarnedIncomeTaxCreditReceivedDuringTaxYearLength2526 int = 1

// Field # 208
const parentDeductiblePaymentsToIRAKeoghOtherStartIndex2526 int = 1889
const parentDeductiblePaymentsToIRAKeoghOtherLength2526 int = 11

// Field # 209
const parentEducationCreditsStartIndex2526 int = 1900
const parentEducationCreditsLength2526 int = 9

// Field # 210
const parentFiledScheduleABDEFHStartIndex2526 int = 1909
const parentFiledScheduleABDEFHLength2526 int = 1

// Field # 211
const parentScheduleCAmountStartIndex2526 int = 1910
const parentScheduleCAmountLength2526 int = 12

// Field # 212
const parentCollegeGrantAndScholarshipAidStartIndex2526 int = 1922
const parentCollegeGrantAndScholarshipAidLength2526 int = 7

// Field # 213
const parentForeignEarnedIncomeExclusionStartIndex2526 int = 1929
const parentForeignEarnedIncomeExclusionLength2526 int = 10

// Field # 214
const parentChildSupportReceivedStartIndex2526 int = 1939
const parentChildSupportReceivedLength2526 int = 7

// Field # 215
const parentTotalOfCashSavingsAndCheckingAccountsStartIndex2526 int = 1946
const parentTotalOfCashSavingsAndCheckingAccountsLength2526 int = 7

// Field # 216
const parentNetWorthOfCurrentInvestmentsStartIndex2526 int = 1953
const parentNetWorthOfCurrentInvestmentsLength2526 int = 7

// Field # 217
const parentNetWorthOfBusinessesAndInvestmentFarmsStartIndex2526 int = 1960
const parentNetWorthOfBusinessesAndInvestmentFarmsLength2526 int = 7

// Field # 218
const parentConsentToRetrieveAndDiscloseFTIStartIndex2526 int = 1967
const parentConsentToRetrieveAndDiscloseFTILength2526 int = 1

// Field # 219
const parentSignatureStartIndex2526 int = 1968
const parentSignatureLength2526 int = 1

// Field # 220
const parentSignatureDateStartIndex2526 int = 1969
const parentSignatureDateLength2526 int = 8

// Field # 222
const parentSpouseFirstNameStartIndex2526 int = 2027
const parentSpouseFirstNameLength2526 int = 35

// Field # 223
const parentSpouseMiddleNameStartIndex2526 int = 2062
const parentSpouseMiddleNameLength2526 int = 15

// Field # 224
const parentSpouseLastNameStartIndex2526 int = 2077
const parentSpouseLastNameLength2526 int = 35

// Field # 225
const parentSpouseSuffixStartIndex2526 int = 2112
const parentSpouseSuffixLength2526 int = 10

// Field # 226
const parentSpouseDateOfBirthStartIndex2526 int = 2122
const parentSpouseDateOfBirthLength2526 int = 8

// Field # 227
const parentSpouseSSNStartIndex2526 int = 2130
const parentSpouseSSNLength2526 int = 9

// Field # 228
const parentSpouseITINStartIndex2526 int = 2139
const parentSpouseITINLength2526 int = 9

// Field # 229
const parentSpousePhoneNumberStartIndex2526 int = 2148
const parentSpousePhoneNumberLength2526 int = 10

// Field # 230
const parentSpouseEmailAddressStartIndex2526 int = 2158
const parentSpouseEmailAddressLength2526 int = 50

// Field # 231
const parentSpouseStreetAddressStartIndex2526 int = 2208
const parentSpouseStreetAddressLength2526 int = 40

// Field # 232
const parentSpouseCityStartIndex2526 int = 2248
const parentSpouseCityLength2526 int = 30

// Field # 233
const parentSpouseStateStartIndex2526 int = 2278
const parentSpouseStateLength2526 int = 2

// Field # 234
const parentSpouseZipCodeStartIndex2526 int = 2280
const parentSpouseZipCodeLength2526 int = 10

// Field # 235
const parentSpouseCountryStartIndex2526 int = 2290
const parentSpouseCountryLength2526 int = 2

// Field # 236
const parentSpouseFiled1040Or1040NRStartIndex2526 int = 2292
const parentSpouseFiled1040Or1040NRLength2526 int = 1

// Field # 237
const parentSpouseFileNonUSTaxReturnStartIndex2526 int = 2293
const parentSpouseFileNonUSTaxReturnLength2526 int = 1

// Field # 238
const parentSpouseTaxReturnFilingStatusStartIndex2526 int = 2294
const parentSpouseTaxReturnFilingStatusLength2526 int = 1

// Field # 239
const parentSpouseIncomeEarnedFromWorkStartIndex2526 int = 2295
const parentSpouseIncomeEarnedFromWorkLength2526 int = 11

// Field # 240
const parentSpouseTaxExemptInterestIncomeStartIndex2526 int = 2306
const parentSpouseTaxExemptInterestIncomeLength2526 int = 11

// Field # 241
const parentSpouseUntaxedPortionsOfIRADistributionsStartIndex2526 int = 2317
const parentSpouseUntaxedPortionsOfIRADistributionsLength2526 int = 11

// Field # 242
const parentSpouseIRARolloverStartIndex2526 int = 2328
const parentSpouseIRARolloverLength2526 int = 11

// Field # 243
const parentSpouseUntaxedPortionsOfPensionsStartIndex2526 int = 2339
const parentSpouseUntaxedPortionsOfPensionsLength2526 int = 11

// Field # 244
const parentSpousePensionRolloverStartIndex2526 int = 2350
const parentSpousePensionRolloverLength2526 int = 11

// Field # 245
const parentSpouseAdjustedGrossIncomeStartIndex2526 int = 2361
const parentSpouseAdjustedGrossIncomeLength2526 int = 10

// Field # 246
const parentSpouseIncomeTaxPaidStartIndex2526 int = 2371
const parentSpouseIncomeTaxPaidLength2526 int = 9

// Field # 247
const parentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex2526 int = 2380
const parentSpouseDeductiblePaymentsToIRAKeoghOtherLength2526 int = 11

// Field # 248
const parentSpouseEducationCreditsStartIndex2526 int = 2391
const parentSpouseEducationCreditsLength2526 int = 9

// Field # 249
const parentSpouseFiledScheduleABDEFHStartIndex2526 int = 2400
const parentSpouseFiledScheduleABDEFHLength2526 int = 1

// Field # 250
const parentSpouseScheduleCAmountStartIndex2526 int = 2401
const parentSpouseScheduleCAmountLength2526 int = 12

// Field # 251
const parentSpouseForeignEarnedIncomeExclusionStartIndex2526 int = 2413
const parentSpouseForeignEarnedIncomeExclusionLength2526 int = 10

// Field # 252
const parentSpouseConsentToRetrieveAndDiscloseFTIStartIndex2526 int = 2423
const parentSpouseConsentToRetrieveAndDiscloseFTILength2526 int = 1

// Field # 253
const parentSpouseSignatureStartIndex2526 int = 2424
const parentSpouseSignatureLength2526 int = 1

// Field # 254
const parentSpouseSignatureDateStartIndex2526 int = 2425
const parentSpouseSignatureDateLength2526 int = 8

// Field # 256
const preparerFirstNameStartIndex2526 int = 2483
const preparerFirstNameLength2526 int = 35

// Field # 257
const preparerLastNameStartIndex2526 int = 2518
const preparerLastNameLength2526 int = 35

// Field # 258
const preparerSSNStartIndex2526 int = 2553
const preparerSSNLength2526 int = 9

// Field # 259
const preparerEINStartIndex2526 int = 2562
const preparerEINLength2526 int = 9

// Field # 260
const preparerAffiliationStartIndex2526 int = 2571
const preparerAffiliationLength2526 int = 30

// Field # 261
const preparerStreetAddressStartIndex2526 int = 2601
const preparerStreetAddressLength2526 int = 40

// Field # 262
const preparerCityStartIndex2526 int = 2641
const preparerCityLength2526 int = 30

// Field # 263
const preparerStateStartIndex2526 int = 2671
const preparerStateLength2526 int = 2

// Field # 264
const preparerZipCodeStartIndex2526 int = 2673
const preparerZipCodeLength2526 int = 10

// Field # 265
const preparerSignatureStartIndex2526 int = 2683
const preparerSignatureLength2526 int = 1

// Field # 266
const preparerSignatureDateStartIndex2526 int = 2684
const preparerSignatureDateLength2526 int = 8

// Field # 268
const studentAffirmationStatusStartIndex2526 int = 2742
const studentAffirmationStatusLength2526 int = 1

// Field # 269
const studentSpouseAffirmationStatusStartIndex2526 int = 2743
const studentSpouseAffirmationStatusLength2526 int = 1

// Field # 270
const parentAffirmationStatusStartIndex2526 int = 2744
const parentAffirmationStatusLength2526 int = 1

// Field # 271
const parentSpouseOrPartnerAffirmationStatusStartIndex2526 int = 2745
const parentSpouseOrPartnerAffirmationStatusLength2526 int = 1

// Field # 272
const studentDateConsentGrantedStartIndex2526 int = 2746
const studentDateConsentGrantedLength2526 int = 8

// Field # 273
const studentSpouseDateConsentGrantedStartIndex2526 int = 2754
const studentSpouseDateConsentGrantedLength2526 int = 8

// Field # 274
const parentDateConsentGrantedStartIndex2526 int = 2762
const parentDateConsentGrantedLength2526 int = 8

// Field # 275
const parentSpouseOrPartnerDateConsentGrantedStartIndex2526 int = 2770
const parentSpouseOrPartnerDateConsentGrantedLength2526 int = 8

// Field # 276
const studentTransunionMatchStatusStartIndex2526 int = 2778
const studentTransunionMatchStatusLength2526 int = 1

// Field # 277
const studentSpouseTransunionMatchStatusStartIndex2526 int = 2779
const studentSpouseTransunionMatchStatusLength2526 int = 1

// Field # 278
const studentParentTransunionMatchStatusStartIndex2526 int = 2780
const studentParentTransunionMatchStatusLength2526 int = 1

// Field # 279
const studentParentSpouseTransunionMatchStatusStartIndex2526 int = 2781
const studentParentSpouseTransunionMatchStatusLength2526 int = 1

// Field # 280
const correctionAppliedAgainstTransactionNumberStartIndex2526 int = 2782
const correctionAppliedAgainstTransactionNumberLength2526 int = 2

// Field # 281
const professionalJudgementStartIndex2526 int = 2784
const professionalJudgementLength2526 int = 1

// Field # 282
const dependencyOverrideIndicatorStartIndex2526 int = 2785
const dependencyOverrideIndicatorLength2526 int = 1

// Field # 283
const fAAFederalSchoolCodeStartIndex2526 int = 2786
const fAAFederalSchoolCodeLength2526 int = 6

// Field # 284
const fAASignatureStartIndex2526 int = 2792
const fAASignatureLength2526 int = 1

// Field # 285
const iASGIndicatorStartIndex2526 int = 2793
const iASGIndicatorLength2526 int = 1

// Field # 286
const childrenOfFallenHeroesIndicatorStartIndex2526 int = 2794
const childrenOfFallenHeroesIndicatorLength2526 int = 1

// Field # 287
const electronicTransactionIndicatorDestinationNumberStartIndex2526 int = 2795
const electronicTransactionIndicatorDestinationNumberLength2526 int = 7

// Field # 288
const studentSignatureSourceStartIndex2526 int = 2802
const studentSignatureSourceLength2526 int = 1

// Field # 289
const studentSpouseSignatureSourceStartIndex2526 int = 2803
const studentSpouseSignatureSourceLength2526 int = 1

// Field # 290
const parentSignatureSourceStartIndex2526 int = 2804
const parentSignatureSourceLength2526 int = 1

// Field # 291
const parentSpouseOrPartnerSignatureSourceStartIndex2526 int = 2805
const parentSpouseOrPartnerSignatureSourceLength2526 int = 1

// Field # 292
const specialHandlingIndicatorStartIndex2526 int = 2806
const specialHandlingIndicatorLength2526 int = 1

// Field # 293
const addressOnlyChangeFlagStartIndex2526 int = 2807
const addressOnlyChangeFlagLength2526 int = 1

// Field # 294
const fpsPushedISIRFlagStartIndex2526 int = 2808
const fpsPushedISIRFlagLength2526 int = 1

// Field # 295
const rejectStatusChangeFlagStartIndex2526 int = 2809
const rejectStatusChangeFlagLength2526 int = 1

// Field # 296
const verificationTrackingFlagStartIndex2526 int = 2810
const verificationTrackingFlagLength2526 int = 2

// Field # 297
const studentSelectedForVerificationStartIndex2526 int = 2812
const studentSelectedForVerificationLength2526 int = 1

// Field # 298
const incarceratedApplicantFlagStartIndex2526 int = 2813
const incarceratedApplicantFlagLength2526 int = 1

// Field # 299
const nsldsTransactionNumberStartIndex2526 int = 2814
const nsldsTransactionNumberLength2526 int = 2

// Field # 300
const nsldsDatabaseResultsFlagStartIndex2526 int = 2816
const nsldsDatabaseResultsFlagLength2526 int = 1

// Field # 301
const highSchoolFlagStartIndex2526 int = 2817
const highSchoolFlagLength2526 int = 1

// Field # 302
const studentTotalFederalWorkStudyEarningsStartIndex2526 int = 2818
const studentTotalFederalWorkStudyEarningsLength2526 int = 12

// Field # 303
const studentSpouseTotalFederalWorkStudyEarningsStartIndex2526 int = 2830
const studentSpouseTotalFederalWorkStudyEarningsLength2526 int = 12

// Field # 304
const parentTotalFederalWorkStudyEarningsStartIndex2526 int = 2842
const parentTotalFederalWorkStudyEarningsLength2526 int = 12

// Field # 305
const parentSpouseOrPartnerTotalFederalWorkStudyEarningsStartIndex2526 int = 2854
const parentSpouseOrPartnerTotalFederalWorkStudyEarningsLength2526 int = 12

// Field # 306
const totalParentAllowancesAgainstIncomeStartIndex2526 int = 2866
const totalParentAllowancesAgainstIncomeLength2526 int = 15

// Field # 307
const parentPayrollTaxAllowanceStartIndex2526 int = 2881
const parentPayrollTaxAllowanceLength2526 int = 15

// Field # 308
const parentIncomeProtectionAllowanceStartIndex2526 int = 2896
const parentIncomeProtectionAllowanceLength2526 int = 15

// Field # 309
const parentEmploymentExpenseAllowanceStartIndex2526 int = 2911
const parentEmploymentExpenseAllowanceLength2526 int = 15

// Field # 310
const parentAvailableIncomeStartIndex2526 int = 2926
const parentAvailableIncomeLength2526 int = 15

// Field # 311
const parentAdjustedAvailableIncomeStartIndex2526 int = 2941
const parentAdjustedAvailableIncomeLength2526 int = 15

// Field # 312
const parentContributionStartIndex2526 int = 2956
const parentContributionLength2526 int = 15

// Field # 313
const studentPayrollTaxAllowanceStartIndex2526 int = 2971
const studentPayrollTaxAllowanceLength2526 int = 15

// Field # 314
const studentIncomeProtectionAllowanceStartIndex2526 int = 2986
const studentIncomeProtectionAllowanceLength2526 int = 15

// Field # 315
const studentAllowanceForParentsNegativeAdjustedAvailableIncomeStartIndex2526 int = 3001
const studentAllowanceForParentsNegativeAdjustedAvailableIncomeLength2526 int = 15

// Field # 316
const studentEmploymentExpenseAllowanceStartIndex2526 int = 3016
const studentEmploymentExpenseAllowanceLength2526 int = 15

// Field # 317
const totalStudentAllowancesAgainstIncomeStartIndex2526 int = 3031
const totalStudentAllowancesAgainstIncomeLength2526 int = 15

// Field # 318
const studentAvailableIncomeStartIndex2526 int = 3046
const studentAvailableIncomeLength2526 int = 15

// Field # 319
const studentContributionFromIncomeStartIndex2526 int = 3061
const studentContributionFromIncomeLength2526 int = 15

// Field # 320
const studentAdjustedAvailableIncomeStartIndex2526 int = 3076
const studentAdjustedAvailableIncomeLength2526 int = 15

// Field # 321
const totalStudentContributionFromSAAIStartIndex2526 int = 3091
const totalStudentContributionFromSAAILength2526 int = 15

// Field # 322
const parentDiscretionaryNetWorthStartIndex2526 int = 3106
const parentDiscretionaryNetWorthLength2526 int = 7

// Field # 323
const parentNetWorthStartIndex2526 int = 3113
const parentNetWorthLength2526 int = 7

// Field # 324
const parentAssetProtectionAllowanceStartIndex2526 int = 3120
const parentAssetProtectionAllowanceLength2526 int = 12

// Field # 325
const parentContributionFromAssetsStartIndex2526 int = 3132
const parentContributionFromAssetsLength2526 int = 12

// Field # 326
const studentNetWorthStartIndex2526 int = 3144
const studentNetWorthLength2526 int = 7

// Field # 327
const studentAssetProtectionAllowanceStartIndex2526 int = 3151
const studentAssetProtectionAllowanceLength2526 int = 12

// Field # 328
const studentContributionFromAssetsStartIndex2526 int = 3163
const studentContributionFromAssetsLength2526 int = 12

// Field # 329
const assumedStudentFamilySizeStartIndex2526 int = 3175
const assumedStudentFamilySizeLength2526 int = 3

// Field # 330
const assumedParentFamilySizeStartIndex2526 int = 3178
const assumedParentFamilySizeLength2526 int = 3

// Field # 331
const studentFirstNameCHVFlagsStartIndex2526 int = 3181
const studentFirstNameCHVFlagsLength2526 int = 3

// Field # 332
const studentMiddleNameCHVFlagsStartIndex2526 int = 3184
const studentMiddleNameCHVFlagsLength2526 int = 3

// Field # 333
const studentLastNameCHVFLagsStartIndex2526 int = 3187
const studentLastNameCHVFLagsLength2526 int = 3

// Field # 334
const studentSuffixCHVFLagsStartIndex2526 int = 3190
const studentSuffixCHVFLagsLength2526 int = 3

// Field # 335
const studentDateOfBirthCHVFLagsStartIndex2526 int = 3193
const studentDateOfBirthCHVFLagsLength2526 int = 3

// Field # 336
const studentSSNCHVFlagsStartIndex2526 int = 3196
const studentSSNCHVFlagsLength2526 int = 3

// Field # 337
const studentITINCHVFLagsStartIndex2526 int = 3199
const studentITINCHVFLagsLength2526 int = 3

// Field # 338
const studentPhoneNumberCHVFlagsStartIndex2526 int = 3202
const studentPhoneNumberCHVFlagsLength2526 int = 3

// Field # 339
const studentEmailAddressCHVFlagsStartIndex2526 int = 3205
const studentEmailAddressCHVFlagsLength2526 int = 3

// Field # 340
const studentStreetAddressCHVFlagsStartIndex2526 int = 3208
const studentStreetAddressCHVFlagsLength2526 int = 3

// Field # 341
const studentCityCHVFLagsStartIndex2526 int = 3211
const studentCityCHVFLagsLength2526 int = 3

// Field # 342
const studentStateCHVFlagsStartIndex2526 int = 3214
const studentStateCHVFlagsLength2526 int = 3

// Field # 343
const studentZipCodeCHVFlagsStartIndex2526 int = 3217
const studentZipCodeCHVFlagsLength2526 int = 3

// Field # 344
const studentCountryCHVFlagsStartIndex2526 int = 3220
const studentCountryCHVFlagsLength2526 int = 3

// Field # 345
const studentMaritalStatusCHVFlagsStartIndex2526 int = 3223
const studentMaritalStatusCHVFlagsLength2526 int = 3

// Field # 346
const studentGradeLevelInCollegeCHVFlagsStartIndex2526 int = 3226
const studentGradeLevelInCollegeCHVFlagsLength2526 int = 3

// Field # 347
const studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsStartIndex2526 int = 3229
const studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsLength2526 int = 3

// Field # 348
const studentPursuingTeacherCertificationCHVFlagsStartIndex2526 int = 3232
const studentPursuingTeacherCertificationCHVFlagsLength2526 int = 3

// Field # 349
const studentActiveDutyCHVFlagsStartIndex2526 int = 3235
const studentActiveDutyCHVFlagsLength2526 int = 3

// Field # 350
const studentVeteranCHVFlagsStartIndex2526 int = 3238
const studentVeteranCHVFlagsLength2526 int = 3

// Field # 351
const studentChildOrOtherDependentsCHVFlagsStartIndex2526 int = 3241
const studentChildOrOtherDependentsCHVFlagsLength2526 int = 3

// Field # 352
const studentParentsDeceasedCHVFlagsStartIndex2526 int = 3244
const studentParentsDeceasedCHVFlagsLength2526 int = 3

// Field # 353
const studentWardOfCourtCHVFlagsStartIndex2526 int = 3247
const studentWardOfCourtCHVFlagsLength2526 int = 3

// Field # 354
const studentInFosterCareCHVFlagsStartIndex2526 int = 3250
const studentInFosterCareCHVFlagsLength2526 int = 3

// Field # 355
const studentEmancipatedMinorCHVFlagsStartIndex2526 int = 3253
const studentEmancipatedMinorCHVFlagsLength2526 int = 3

// Field # 356
const studentLegalGuardianshipCHVFlagsStartIndex2526 int = 3256
const studentLegalGuardianshipCHVFlagsLength2526 int = 3

// Field # 357
const studentPersonalCircumstancesNoneOfTheAboveCHVFlagsStartIndex2526 int = 3259
const studentPersonalCircumstancesNoneOfTheAboveCHVFlagsLength2526 int = 3

// Field # 358
const studentUnaccompaniedHomelessYouthOrIsUnaccompaniedAtRIskSelfSupportingCHVFlagsStartIndex2526 int = 3262
const studentUnaccompaniedHomelessYouthOrIsUnaccompaniedAtRIskSelfSupportingCHVFlagsLength2526 int = 3

// Field # 359
const studentUnaccompaniedAndHomelessGeneralCHVFlagsStartIndex2526 int = 3265
const studentUnaccompaniedAndHomelessGeneralCHVFlagsLength2526 int = 3

// Field # 360
const studentUnaccompaniedAndHomelessHSCHVFlagsStartIndex2526 int = 3268
const studentUnaccompaniedAndHomelessHSCHVFlagsLength2526 int = 3

// Field # 361
const studentUnaccompaniedAndHomelessTRIOCHVFlagsStartIndex2526 int = 3271
const studentUnaccompaniedAndHomelessTRIOCHVFlagsLength2526 int = 3

// Field # 362
const studentUnaccompaniedAndHomelessFAACHVFlagsStartIndex2526 int = 3274
const studentUnaccompaniedAndHomelessFAACHVFlagsLength2526 int = 3

// Field # 363
const studentHomelessnessNoneOfTheAboveCHVFlagsStartIndex2526 int = 3277
const studentHomelessnessNoneOfTheAboveCHVFlagsLength2526 int = 3

// Field # 364
const studentHasUnusualCircumstanceCHVFlagsStartIndex2526 int = 3280
const studentHasUnusualCircumstanceCHVFlagsLength2526 int = 3

// Field # 365
const studentUnsubOnlyCHVFlagsStartIndex2526 int = 3283
const studentUnsubOnlyCHVFlagsLength2526 int = 3

// Field # 366
const studentUpdatedFamilySizeCHVFlagsStartIndex2526 int = 3286
const studentUpdatedFamilySizeCHVFlagsLength2526 int = 3

// Field # 367
const studentNumberInCollegeCorrectionCHVFlagsStartIndex2526 int = 3289
const studentNumberInCollegeCorrectionCHVFlagsLength2526 int = 3

// Field # 368
const studentCitizenshipStatusCorrectionCHVFlagsStartIndex2526 int = 3292
const studentCitizenshipStatusCorrectionCHVFlagsLength2526 int = 3

// Field # 369
const studentANumberCHVFlagsStartIndex2526 int = 3295
const studentANumberCHVFlagsLength2526 int = 3

// Field # 370
const studentStateOfLegalResidenceCHVFlagsStartIndex2526 int = 3298
const studentStateOfLegalResidenceCHVFlagsLength2526 int = 3

// Field # 371
const studentLegalResidenceDateCHVFlagsStartIndex2526 int = 3301
const studentLegalResidenceDateCHVFlagsLength2526 int = 3

// Field # 372
const studentEitherParentAttendCollegeCHVFlagsStartIndex2526 int = 3304
const studentEitherParentAttendCollegeCHVFlagsLength2526 int = 3

// Field # 373
const studentParentKilledInTheLineOfDutyCHVFlagsStartIndex2526 int = 3307
const studentParentKilledInTheLineOfDutyCHVFlagsLength2526 int = 3

// Field # 374
const studentHighSchoolCompletionStatusCHVFlagsStartIndex2526 int = 3310
const studentHighSchoolCompletionStatusCHVFlagsLength2526 int = 3

// Field # 375
const studentHighSchoolNameCHVFlagsStartIndex2526 int = 3313
const studentHighSchoolNameCHVFlagsLength2526 int = 3

// Field # 376
const studentHighSchoolCityCHVFlagsStartIndex2526 int = 3316
const studentHighSchoolCityCHVFlagsLength2526 int = 3

// Field # 377
const studentHighSchoolStateCHVFlagsStartIndex2526 int = 3319
const studentHighSchoolStateCHVFlagsLength2526 int = 3

// Field # 378
const studentHighSchoolEquivalentDiplomaNameCHVFlagsStartIndex2526 int = 3322
const studentHighSchoolEquivalentDiplomaNameCHVFlagsLength2526 int = 3

// Field # 379
const studentHighSchoolEquivalentDiplomaStateCHVFlagsStartIndex2526 int = 3325
const studentHighSchoolEquivalentDiplomaStateCHVFlagsLength2526 int = 3

// Field # 380
const studentReceivedEITCCHVFlagsStartIndex2526 int = 3328
const studentReceivedEITCCHVFlagsLength2526 int = 3

// Field # 381
const studentReceivedFederalHousingAssistanceCHVFlagsStartIndex2526 int = 3331
const studentReceivedFederalHousingAssistanceCHVFlagsLength2526 int = 3

// Field # 382
const studentReceivedFreeReducedPriceLunchCHVFlagsStartIndex2526 int = 3334
const studentReceivedFreeReducedPriceLunchCHVFlagsLength2526 int = 3

// Field # 383
const studentReceivedMedicaidCHVFlagsStartIndex2526 int = 3337
const studentReceivedMedicaidCHVFlagsLength2526 int = 3

// Field # 384
const studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex2526 int = 3340
const studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsLength2526 int = 3

// Field # 385
const studentReceivedSNAPCHVFlagsStartIndex2526 int = 3343
const studentReceivedSNAPCHVFlagsLength2526 int = 3

// Field # 386
const studentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex2526 int = 3346
const studentReceivedSupplementalSecurityIncomeCHVFlagsLength2526 int = 3

// Field # 387
const studentReceivedTANFCHVFlagsStartIndex2526 int = 3349
const studentReceivedTANFCHVFlagsLength2526 int = 3

// Field # 388
const studentReceivedWICCHVFlagsStartIndex2526 int = 3352
const studentReceivedWICCHVFlagsLength2526 int = 3

// Field # 389
const studentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex2526 int = 3355
const studentFederalBenefitsNoneOfTheAboveCHVFlagsLength2526 int = 3

// Field # 390
const studentFiled1040Or1040NRCHVFlagsStartIndex2526 int = 3358
const studentFiled1040Or1040NRCHVFlagsLength2526 int = 3

// Field # 391
const studentFiledNonUSTaxReturnCHVFlagsStartIndex2526 int = 3361
const studentFiledNonUSTaxReturnCHVFlagsLength2526 int = 3

// Field # 392
const studentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex2526 int = 3364
const studentFiledJointReturnWithCurrentSpouseCHVFlagsLength2526 int = 3

// Field # 393
const studentTaxReturnFilingStatusCHVFlagsStartIndex2526 int = 3367
const studentTaxReturnFilingStatusCHVFlagsLength2526 int = 3

// Field # 394
const studentIncomeEarnedFromWorkCorrectionCHVFlagsStartIndex2526 int = 3370
const studentIncomeEarnedFromWorkCorrectionCHVFlagsLength2526 int = 3

// Field # 395
const studentTaxExemptInterestIncomeCHVFlagsStartIndex2526 int = 3373
const studentTaxExemptInterestIncomeCHVFlagsLength2526 int = 3

// Field # 396
const studentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2526 int = 3376
const studentUntaxedPortionsOfIRADistributionsCHVFlagsLength2526 int = 3

// Field # 397
const studentIRARolloverCHVFlagsStartIndex2526 int = 3379
const studentIRARolloverCHVFlagsLength2526 int = 3

// Field # 398
const studentUntaxedPortionsOfPensionsCHVFlagsStartIndex2526 int = 3382
const studentUntaxedPortionsOfPensionsCHVFlagsLength2526 int = 3

// Field # 399
const studentPensionRolloverCHVFlagsStartIndex2526 int = 3385
const studentPensionRolloverCHVFlagsLength2526 int = 3

// Field # 400
const studentAdjustedGrossIncomeCHVFlagsStartIndex2526 int = 3388
const studentAdjustedGrossIncomeCHVFlagsLength2526 int = 3

// Field # 401
const studentIncomeTaxPaidCHVFlagsStartIndex2526 int = 3391
const studentIncomeTaxPaidCHVFlagsLength2526 int = 3

// Field # 402
const studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex2526 int = 3394
const studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsLength2526 int = 3

// Field # 403
const studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2526 int = 3397
const studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2526 int = 3

// Field # 404
const studentEducationCreditsCHVFlagsStartIndex2526 int = 3400
const studentEducationCreditsCHVFlagsLength2526 int = 3

// Field # 405
const studentFiledScheduleABDEFHCHVFlagsStartIndex2526 int = 3403
const studentFiledScheduleABDEFHCHVFlagsLength2526 int = 3

// Field # 406
const studentScheduleCAmountCHVFlagsStartIndex2526 int = 3406
const studentScheduleCAmountCHVFlagsLength2526 int = 3

// Field # 407
const studentCollegeGrantAndScholarshipAidCHVFlagsStartIndex2526 int = 3409
const studentCollegeGrantAndScholarshipAidCHVFlagsLength2526 int = 3

// Field # 408
const studentForeignEarnedIncomeExclusionCHVFlagsStartIndex2526 int = 3412
const studentForeignEarnedIncomeExclusionCHVFlagsLength2526 int = 3

// Field # 409
const studentChildSupportReceivedCHVFlagsStartIndex2526 int = 3415
const studentChildSupportReceivedCHVFlagsLength2526 int = 3

// Field # 410
const studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex2526 int = 3418
const studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsLength2526 int = 3

// Field # 411
const studentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex2526 int = 3421
const studentNetWorthOfCurrentInvestmentsCHVFlagsLength2526 int = 3

// Field # 412
const studentTotalOfCashSavingsAndCheckingCHVFlagsStartIndex2526 int = 3424
const studentTotalOfCashSavingsAndCheckingCHVFlagsLength2526 int = 3

// Field # 413
const studentCollege1CHVFlagsStartIndex2526 int = 3427
const studentCollege1CHVFlagsLength2526 int = 3

// Field # 414
const studentCollege2CHVFlagsStartIndex2526 int = 3430
const studentCollege2CHVFlagsLength2526 int = 3

// Field # 415
const studentCollege3CHVFlagsStartIndex2526 int = 3433
const studentCollege3CHVFlagsLength2526 int = 3

// Field # 416
const studentCollege4CHVFlagsStartIndex2526 int = 3436
const studentCollege4CHVFlagsLength2526 int = 3

// Field # 417
const studentCollege5CHVFlagsStartIndex2526 int = 3439
const studentCollege5CHVFlagsLength2526 int = 3

// Field # 418
const studentCollege6CHVFlagsStartIndex2526 int = 3442
const studentCollege6CHVFlagsLength2526 int = 3

// Field # 419
const studentCollege7CHVFlagsStartIndex2526 int = 3445
const studentCollege7CHVFlagsLength2526 int = 3

// Field # 420
const studentCollege8CHVFlagsStartIndex2526 int = 3448
const studentCollege8CHVFlagsLength2526 int = 3

// Field # 421
const studentCollege9CHVFlagsStartIndex2526 int = 3451
const studentCollege9CHVFlagsLength2526 int = 3

// Field # 422
const studentCollege10CHVFlagsStartIndex2526 int = 3454
const studentCollege10CHVFlagsLength2526 int = 3

// Field # 423
const studentCollege11CHVFlagsStartIndex2526 int = 3457
const studentCollege11CHVFlagsLength2526 int = 3

// Field # 424
const studentCollege12CHVFlagsStartIndex2526 int = 3460
const studentCollege12CHVFlagsLength2526 int = 3

// Field # 425
const studentCollege13CHVFlagsStartIndex2526 int = 3463
const studentCollege13CHVFlagsLength2526 int = 3

// Field # 426
const studentCollege14CHVFlagsStartIndex2526 int = 3466
const studentCollege14CHVFlagsLength2526 int = 3

// Field # 427
const studentCollege15CHVFlagsStartIndex2526 int = 3469
const studentCollege15CHVFlagsLength2526 int = 3

// Field # 428
const studentCollege16CHVFlagsStartIndex2526 int = 3472
const studentCollege16CHVFlagsLength2526 int = 3

// Field # 429
const studentCollege17CHVFlagsStartIndex2526 int = 3475
const studentCollege17CHVFlagsLength2526 int = 3

// Field # 430
const studentCollege18CHVFlagsStartIndex2526 int = 3478
const studentCollege18CHVFlagsLength2526 int = 3

// Field # 431
const studentCollege19CHVFlagsStartIndex2526 int = 3481
const studentCollege19CHVFlagsLength2526 int = 3

// Field # 432
const studentCollege20CHVFlagsStartIndex2526 int = 3484
const studentCollege20CHVFlagsLength2526 int = 3

// Field # 433
const studentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2526 int = 3487
const studentConsentToRetrieveAndDiscloseFTICHVFlagsLength2526 int = 3

// Field # 434
const studentSignatureCHVFlagsStartIndex2526 int = 3490
const studentSignatureCHVFlagsLength2526 int = 3

// Field # 435
const studentSignatureDateCHVFlagsStartIndex2526 int = 3493
const studentSignatureDateCHVFlagsLength2526 int = 3

// Field # 436
const studentSpouseFirstNameCHVFlagsStartIndex2526 int = 3496
const studentSpouseFirstNameCHVFlagsLength2526 int = 3

// Field # 437
const studentSpouseMiddleNameCHVFlagsStartIndex2526 int = 3499
const studentSpouseMiddleNameCHVFlagsLength2526 int = 3

// Field # 438
const studentSpouseLastNameCHVFlagsStartIndex2526 int = 3502
const studentSpouseLastNameCHVFlagsLength2526 int = 3

// Field # 439
const studentSpouseSuffixCHVFlagsStartIndex2526 int = 3505
const studentSpouseSuffixCHVFlagsLength2526 int = 3

// Field # 440
const studentSpouseDateOfBirthCHVFlagsStartIndex2526 int = 3508
const studentSpouseDateOfBirthCHVFlagsLength2526 int = 3

// Field # 441
const studentSpouseSSNCHVFlagsStartIndex2526 int = 3511
const studentSpouseSSNCHVFlagsLength2526 int = 3

// Field # 442
const studentSpouseITINCHVFlagsStartIndex2526 int = 3514
const studentSpouseITINCHVFlagsLength2526 int = 3

// Field # 443
const studentSpousePhoneNumberCHVFlagsStartIndex2526 int = 3517
const studentSpousePhoneNumberCHVFlagsLength2526 int = 3

// Field # 444
const studentSpouseEmailAddressCHVFlagsStartIndex2526 int = 3520
const studentSpouseEmailAddressCHVFlagsLength2526 int = 3

// Field # 445
const studentSpouseStreetAddressCHVFlagsStartIndex2526 int = 3523
const studentSpouseStreetAddressCHVFlagsLength2526 int = 3

// Field # 446
const studentSpouseCityCHVFlagsStartIndex2526 int = 3526
const studentSpouseCityCHVFlagsLength2526 int = 3

// Field # 447
const studentSpouseStateCHVFlagsStartIndex2526 int = 3529
const studentSpouseStateCHVFlagsLength2526 int = 3

// Field # 448
const studentSpouseZipCodeCHVFlagsStartIndex2526 int = 3532
const studentSpouseZipCodeCHVFlagsLength2526 int = 3

// Field # 449
const studentSpouseCountryCHVFlagsStartIndex2526 int = 3535
const studentSpouseCountryCHVFlagsLength2526 int = 3

// Field # 450
const studentSpouseFiled1040Or1040NRCHVFlagsStartIndex2526 int = 3538
const studentSpouseFiled1040Or1040NRCHVFlagsLength2526 int = 3

// Field # 451
const studentSpouseFiledNonUSTaxReturnCHVFlagsStartIndex2526 int = 3541
const studentSpouseFiledNonUSTaxReturnCHVFlagsLength2526 int = 3

// Field # 452
const studentSpouseTaxReturnFilingStatusCHVFlagsStartIndex2526 int = 3544
const studentSpouseTaxReturnFilingStatusCHVFlagsLength2526 int = 3

// Field # 453
const studentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex2526 int = 3547
const studentSpouseIncomeEarnedFromWorkCHVFlagsLength2526 int = 3

// Field # 454
const studentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex2526 int = 3550
const studentSpouseTaxExemptInterestIncomeCHVFlagsLength2526 int = 3

// Field # 455
const studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2526 int = 3553
const studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsLength2526 int = 3

// Field # 456
const studentSpouseIRARolloverCHVFlagsStartIndex2526 int = 3556
const studentSpouseIRARolloverCHVFlagsLength2526 int = 3

// Field # 457
const studentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex2526 int = 3559
const studentSpouseUntaxedPortionsOfPensionsCHVFlagsLength2526 int = 3

// Field # 458
const studentSpousePensionRolloverCHVFlagsStartIndex2526 int = 3562
const studentSpousePensionRolloverCHVFlagsLength2526 int = 3

// Field # 459
const studentSpouseAdjustedGrossIncomeCHVFlagsStartIndex2526 int = 3565
const studentSpouseAdjustedGrossIncomeCHVFlagsLength2526 int = 3

// Field # 460
const studentSpouseIncomeTaxPaidCHVFlagsStartIndex2526 int = 3568
const studentSpouseIncomeTaxPaidCHVFlagsLength2526 int = 3

// Field # 461
const studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2526 int = 3571
const studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2526 int = 3

// Field # 462
const studentSpouseEducationCreditsCHVFlagsStartIndex2526 int = 3574
const studentSpouseEducationCreditsCHVFlagsLength2526 int = 3

// Field # 463
const studentSpouseFiledScheduleABDEFHCHVFlagsStartIndex2526 int = 3577
const studentSpouseFiledScheduleABDEFHCHVFlagsLength2526 int = 3

// Field # 464
const studentSpouseScheduleCAmountCHVFlagsStartIndex2526 int = 3580
const studentSpouseScheduleCAmountCHVFlagsLength2526 int = 3

// Field # 465
const studentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex2526 int = 3583
const studentSpouseForeignEarnedIncomeExclusionCHVFlagsLength2526 int = 3

// Field # 466
const studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2526 int = 3586
const studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsLength2526 int = 3

// Field # 467
const studentSpouseSignatureCHVFlagsStartIndex2526 int = 3589
const studentSpouseSignatureCHVFlagsLength2526 int = 3

// Field # 468
const studentSpouseSignatureDateCHVFlagsStartIndex2526 int = 3592
const studentSpouseSignatureDateCHVFlagsLength2526 int = 3

// Field # 469
const parentFirstNameCHVFlagsStartIndex2526 int = 3595
const parentFirstNameCHVFlagsLength2526 int = 3

// Field # 470
const parentMiddleNameCHVFlagsStartIndex2526 int = 3598
const parentMiddleNameCHVFlagsLength2526 int = 3

// Field # 471
const parentLastNameCHVFlagsStartIndex2526 int = 3601
const parentLastNameCHVFlagsLength2526 int = 3

// Field # 472
const parentSuffixCHVFlagsStartIndex2526 int = 3604
const parentSuffixCHVFlagsLength2526 int = 3

// Field # 473
const parentDateOfBirthCHVFlagsStartIndex2526 int = 3607
const parentDateOfBirthCHVFlagsLength2526 int = 3

// Field # 474
const parentSSNCHVFlagsStartIndex2526 int = 3610
const parentSSNCHVFlagsLength2526 int = 3

// Field # 475
const parentITINCHVFlagsStartIndex2526 int = 3613
const parentITINCHVFlagsLength2526 int = 3

// Field # 476
const parentPhoneNumberCHVFlagsStartIndex2526 int = 3616
const parentPhoneNumberCHVFlagsLength2526 int = 3

// Field # 477
const parentEmailAddressCHVFlagsStartIndex2526 int = 3619
const parentEmailAddressCHVFlagsLength2526 int = 3

// Field # 478
const parentStreetAddressCHVFlagsStartIndex2526 int = 3622
const parentStreetAddressCHVFlagsLength2526 int = 3

// Field # 479
const parentCityCHVFlagsStartIndex2526 int = 3625
const parentCityCHVFlagsLength2526 int = 3

// Field # 480
const parentStateCHVFlagsStartIndex2526 int = 3628
const parentStateCHVFlagsLength2526 int = 3

// Field # 481
const parentZipCodeCHVFlagsStartIndex2526 int = 3631
const parentZipCodeCHVFlagsLength2526 int = 3

// Field # 482
const parentCountryCHVFlagsStartIndex2526 int = 3634
const parentCountryCHVFlagsLength2526 int = 3

// Field # 483
const parentMaritalStatusCHVFlagsStartIndex2526 int = 3637
const parentMaritalStatusCHVFlagsLength2526 int = 3

// Field # 484
const parentStateOfLegalResidenceCHVFlagsStartIndex2526 int = 3640
const parentStateOfLegalResidenceCHVFlagsLength2526 int = 3

// Field # 485
const parentLegalResidenceDateCHVFlagsStartIndex2526 int = 3643
const parentLegalResidenceDateCHVFlagsLength2526 int = 3

// Field # 486
const parentUpdatedFamilySizeCHVFlagsStartIndex2526 int = 3646
const parentUpdatedFamilySizeCHVFlagsLength2526 int = 3

// Field # 487
const parentNumberInCollegeCHVFlagsStartIndex2526 int = 3649
const parentNumberInCollegeCHVFlagsLength2526 int = 3

// Field # 488
const parentReceivedEITCCHVFlagsStartIndex2526 int = 3652
const parentReceivedEITCCHVFlagsLength2526 int = 3

// Field # 489
const parentReceivedFederalHousingAssistanceCHVFlagsStartIndex2526 int = 3655
const parentReceivedFederalHousingAssistanceCHVFlagsLength2526 int = 3

// Field # 490
const parentReceivedFreeReducedPriceLunchCHVFlagsStartIndex2526 int = 3658
const parentReceivedFreeReducedPriceLunchCHVFlagsLength2526 int = 3

// Field # 491
const parentReceivedMedicaidCHVFlagsStartIndex2526 int = 3661
const parentReceivedMedicaidCHVFlagsLength2526 int = 3

// Field # 492
const parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex2526 int = 3664
const parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsLength2526 int = 3

// Field # 493
const parentReceivedSNAPCHVFlagsStartIndex2526 int = 3667
const parentReceivedSNAPCHVFlagsLength2526 int = 3

// Field # 494
const parentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex2526 int = 3670
const parentReceivedSupplementalSecurityIncomeCHVFlagsLength2526 int = 3

// Field # 495
const parentReceivedTANFCHVFlagsStartIndex2526 int = 3673
const parentReceivedTANFCHVFlagsLength2526 int = 3

// Field # 496
const parentReceivedWICCHVFlagsStartIndex2526 int = 3676
const parentReceivedWICCHVFlagsLength2526 int = 3

// Field # 497
const parentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex2526 int = 3679
const parentFederalBenefitsNoneOfTheAboveCHVFlagsLength2526 int = 3

// Field # 498
const parentFiled1040Or1040NRCHVFlagsStartIndex2526 int = 3682
const parentFiled1040Or1040NRCHVFlagsLength2526 int = 3

// Field # 499
const parentFileNonUSTaxReturnCHVFlagsStartIndex2526 int = 3685
const parentFileNonUSTaxReturnCHVFlagsLength2526 int = 3

// Field # 500
const parentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex2526 int = 3688
const parentFiledJointReturnWithCurrentSpouseCHVFlagsLength2526 int = 3

// Field # 501
const parentTaxReturnFilingStatusCHVFlagsStartIndex2526 int = 3691
const parentTaxReturnFilingStatusCHVFlagsLength2526 int = 3

// Field # 502
const parentIncomeEarnedFromWorkCHVFlagsStartIndex2526 int = 3694
const parentIncomeEarnedFromWorkCHVFlagsLength2526 int = 3

// Field # 503
const parentTaxExemptInterestIncomeCHVFlagsStartIndex2526 int = 3697
const parentTaxExemptInterestIncomeCHVFlagsLength2526 int = 3

// Field # 504
const parentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2526 int = 3700
const parentUntaxedPortionsOfIRADistributionsCHVFlagsLength2526 int = 3

// Field # 505
const parentIRARolloverCHVFlagsStartIndex2526 int = 3703
const parentIRARolloverCHVFlagsLength2526 int = 3

// Field # 506
const parentUntaxedPortionsOfPensionsCHVFlagsStartIndex2526 int = 3706
const parentUntaxedPortionsOfPensionsCHVFlagsLength2526 int = 3

// Field # 507
const parentPensionRolloverCHVFlagsStartIndex2526 int = 3709
const parentPensionRolloverCHVFlagsLength2526 int = 3

// Field # 508
const parentAdjustedGrossIncomeCHVFlagsStartIndex2526 int = 3712
const parentAdjustedGrossIncomeCHVFlagsLength2526 int = 3

// Field # 509
const parentIncomeTaxPaidCHVFlagsStartIndex2526 int = 3715
const parentIncomeTaxPaidCHVFlagsLength2526 int = 3

// Field # 510
const parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex2526 int = 3718
const parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsLength2526 int = 3

// Field # 511
const parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2526 int = 3721
const parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2526 int = 3

// Field # 512
const parentEducationCreditsCHVFlagsStartIndex2526 int = 3724
const parentEducationCreditsCHVFlagsLength2526 int = 3

// Field # 513
const parentFiledScheduleABDEFHCHVFlagsStartIndex2526 int = 3727
const parentFiledScheduleABDEFHCHVFlagsLength2526 int = 3

// Field # 514
const parentScheduleCAmountCHVFlagsStartIndex2526 int = 3730
const parentScheduleCAmountCHVFlagsLength2526 int = 3

// Field # 515
const parentCollegeGrantAndScholarshipAidCHVFlagsStartIndex2526 int = 3733
const parentCollegeGrantAndScholarshipAidCHVFlagsLength2526 int = 3

// Field # 516
const parentForeignEarnedIncomeExclusionCHVFlagsStartIndex2526 int = 3736
const parentForeignEarnedIncomeExclusionCHVFlagsLength2526 int = 3

// Field # 517
const parentChildSupportReceivedCHVFlagsStartIndex2526 int = 3739
const parentChildSupportReceivedCHVFlagsLength2526 int = 3

// Field # 518
const parentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex2526 int = 3742
const parentNetWorthOfCurrentInvestmentsCHVFlagsLength2526 int = 3

// Field # 519
const parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsStartIndex2526 int = 3745
const parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsLength2526 int = 3

// Field # 520
const parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex2526 int = 3748
const parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsLength2526 int = 3

// Field # 521
const parentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2526 int = 3751
const parentConsentToRetrieveAndDiscloseFTICHVFlagsLength2526 int = 3

// Field # 522
const parentSignatureCHVFlagsStartIndex2526 int = 3754
const parentSignatureCHVFlagsLength2526 int = 3

// Field # 523
const parentSignatureDateCHVFlagsStartIndex2526 int = 3757
const parentSignatureDateCHVFlagsLength2526 int = 3

// Field # 524
const parentSpouseFirstNameCHVFlagsStartIndex2526 int = 3760
const parentSpouseFirstNameCHVFlagsLength2526 int = 3

// Field # 525
const parentSpouseMiddleNameCHVFlagsStartIndex2526 int = 3763
const parentSpouseMiddleNameCHVFlagsLength2526 int = 3

// Field # 526
const parentSpouseLastNameCHVFlagsStartIndex2526 int = 3766
const parentSpouseLastNameCHVFlagsLength2526 int = 3

// Field # 527
const parentSpouseSuffixCHVFlagsStartIndex2526 int = 3769
const parentSpouseSuffixCHVFlagsLength2526 int = 3

// Field # 528
const parentSpouseDateOfBirthCHVFlagsStartIndex2526 int = 3772
const parentSpouseDateOfBirthCHVFlagsLength2526 int = 3

// Field # 529
const parentSpouseSSNCHVFlagsStartIndex2526 int = 3775
const parentSpouseSSNCHVFlagsLength2526 int = 3

// Field # 530
const parentSpouseITINCHVFlagsStartIndex2526 int = 3778
const parentSpouseITINCHVFlagsLength2526 int = 3

// Field # 531
const parentSpousePhoneNumberCHVFlagsStartIndex2526 int = 3781
const parentSpousePhoneNumberCHVFlagsLength2526 int = 3

// Field # 532
const parentSpouseEmailAddressCHVFlagsStartIndex2526 int = 3784
const parentSpouseEmailAddressCHVFlagsLength2526 int = 3

// Field # 533
const parentSpouseStreetAddressCHVFlagsStartIndex2526 int = 3787
const parentSpouseStreetAddressCHVFlagsLength2526 int = 3

// Field # 534
const parentSpouseCityCHVFlagsStartIndex2526 int = 3790
const parentSpouseCityCHVFlagsLength2526 int = 3

// Field # 535
const parentSpouseStateCHVFlagsStartIndex2526 int = 3793
const parentSpouseStateCHVFlagsLength2526 int = 3

// Field # 536
const parentSpouseZipCodeCHVFlagsStartIndex2526 int = 3796
const parentSpouseZipCodeCHVFlagsLength2526 int = 3

// Field # 537
const parentSpouseCountryCHVFlagsStartIndex2526 int = 3799
const parentSpouseCountryCHVFlagsLength2526 int = 3

// Field # 538
const parentSpouseFiled1040Or1040NRCHVFlagsStartIndex2526 int = 3802
const parentSpouseFiled1040Or1040NRCHVFlagsLength2526 int = 3

// Field # 539
const parentSpouseFileNonUSTaxReturnCHVFlagsStartIndex2526 int = 3805
const parentSpouseFileNonUSTaxReturnCHVFlagsLength2526 int = 3

// Field # 540
const parentSpouseTaxReturnFilingStatusCHVFlagsStartIndex2526 int = 3808
const parentSpouseTaxReturnFilingStatusCHVFlagsLength2526 int = 3

// Field # 541
const parentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex2526 int = 3811
const parentSpouseIncomeEarnedFromWorkCHVFlagsLength2526 int = 3

// Field # 542
const parentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex2526 int = 3814
const parentSpouseTaxExemptInterestIncomeCHVFlagsLength2526 int = 3

// Field # 543
const parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2526 int = 3817
const parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsLength2526 int = 3

// Field # 544
const parentSpouseIRARolloverCHVFlagsStartIndex2526 int = 3820
const parentSpouseIRARolloverCHVFlagsLength2526 int = 3

// Field # 545
const parentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex2526 int = 3823
const parentSpouseUntaxedPortionsOfPensionsCHVFlagsLength2526 int = 3

// Field # 546
const parentSpousePensionRolloverCHVFlagsStartIndex2526 int = 3826
const parentSpousePensionRolloverCHVFlagsLength2526 int = 3

// Field # 547
const parentSpouseAdjustedGrossIncomeCHVFlagsStartIndex2526 int = 3829
const parentSpouseAdjustedGrossIncomeCHVFlagsLength2526 int = 3

// Field # 548
const parentSpouseIncomeTaxPaidCHVFlagsStartIndex2526 int = 3832
const parentSpouseIncomeTaxPaidCHVFlagsLength2526 int = 3

// Field # 549
const parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2526 int = 3835
const parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2526 int = 3

// Field # 550
const parentSpouseEducationCreditsCHVFlagsStartIndex2526 int = 3838
const parentSpouseEducationCreditsCHVFlagsLength2526 int = 3

// Field # 551
const parentSpouseFiledScheduleABDEFHCHVFlagsStartIndex2526 int = 3841
const parentSpouseFiledScheduleABDEFHCHVFlagsLength2526 int = 3

// Field # 552
const parentSpouseScheduleCAmountCHVFlagsStartIndex2526 int = 3844
const parentSpouseScheduleCAmountCHVFlagsLength2526 int = 3

// Field # 553
const parentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex2526 int = 3847
const parentSpouseForeignEarnedIncomeExclusionCHVFlagsLength2526 int = 3

// Field # 554
const parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2526 int = 3850
const parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsLength2526 int = 3

// Field # 555
const parentSpouseSignatureCHVFlagsStartIndex2526 int = 3853
const parentSpouseSignatureCHVFlagsLength2526 int = 3

// Field # 556
const parentSpouseSignatureDateCHVFlagsStartIndex2526 int = 3856
const parentSpouseSignatureDateCHVFlagsLength2526 int = 3

// Field # 557
const dHSPrimaryMatchStatusStartIndex2526 int = 3859
const dHSPrimaryMatchStatusLength2526 int = 1

// Field # 559
const dHSCaseNumberStartIndex2526 int = 3861
const dHSCaseNumberLength2526 int = 15

// Field # 560
const nsldsMatchStatusStartIndex2526 int = 3876
const nsldsMatchStatusLength2526 int = 1

// Field # 561
const nsldsPostscreeningReasonCodeStartIndex2526 int = 3877
const nsldsPostscreeningReasonCodeLength2526 int = 6

// Field # 562
const studentSSACitizenshipFlagResultsStartIndex2526 int = 3883
const studentSSACitizenshipFlagResultsLength2526 int = 1

// Field # 563
const studentSSAMatchStatusStartIndex2526 int = 3884
const studentSSAMatchStatusLength2526 int = 1

// Field # 564
const studentSpouseSSAMatchStatusStartIndex2526 int = 3885
const studentSpouseSSAMatchStatusLength2526 int = 1

// Field # 565
const parentSSAMatchStatusStartIndex2526 int = 3886
const parentSSAMatchStatusLength2526 int = 1

// Field # 566
const parentSpouseOrPartnerSSAMatchStatusStartIndex2526 int = 3887
const parentSpouseOrPartnerSSAMatchStatusLength2526 int = 1

// Field # 567
const vAMatchFlagStartIndex2526 int = 3888
const vAMatchFlagLength2526 int = 1

// Field # 568
const commentCodesStartIndex2526 int = 3889
const commentCodesLength2526 int = 60

// Field # 569
const drugAbuseHoldIndicatorStartIndex2526 int = 3949
const drugAbuseHoldIndicatorLength2526 int = 1

// Field # 570
const graduateFlagStartIndex2526 int = 3950
const graduateFlagLength2526 int = 1

// Field # 571
const pellGrantEligibilityFlagStartIndex2526 int = 3951
const pellGrantEligibilityFlagLength2526 int = 1

// Field # 572
const reprocessedReasonCodeStartIndex2526 int = 3952
const reprocessedReasonCodeLength2526 int = 2

// Field # 573
const fpsCFlagStartIndex2526 int = 3954
const fpsCFlagLength2526 int = 1

// Field # 574
const fpsCChangeFlagStartIndex2526 int = 3955
const fpsCChangeFlagLength2526 int = 1

// Field # 575
const electronicFederalSchoolCodeIndicatorStartIndex2526 int = 3956
const electronicFederalSchoolCodeIndicatorLength2526 int = 2

// Field # 576
const rejectReasonCodesStartIndex2526 int = 3958
const rejectReasonCodesLength2526 int = 110

// Field # 577
const electronicTransactionIndicatorFlagStartIndex2526 int = 4068
const electronicTransactionIndicatorFlagLength2526 int = 1

// Field # 578
const studentLastNameSSNChangeFlagStartIndex2526 int = 4069
const studentLastNameSSNChangeFlagLength2526 int = 1

// Field # 579
const highSchoolCodeStartIndex2526 int = 4070
const highSchoolCodeLength2526 int = 12

// Field # 580
const verificationSelectionChangeFlagStartIndex2526 int = 4082
const verificationSelectionChangeFlagLength2526 int = 1

// Field # 581
const useUserProvidedDataOnlyStartIndex2526 int = 4083
const useUserProvidedDataOnlyLength2526 int = 5

// Field # 583
const nsldsPellOverpaymentFlagStartIndex2526 int = 4449
const nsldsPellOverpaymentFlagLength2526 int = 1

// Field # 584
const nsldsPellOverpaymentContactStartIndex2526 int = 4450
const nsldsPellOverpaymentContactLength2526 int = 8

// Field # 585
const nsldsFSEOGOverpaymentFlagStartIndex2526 int = 4458
const nsldsFSEOGOverpaymentFlagLength2526 int = 1

// Field # 586
const nsldsFSEOGOverpaymentContactStartIndex2526 int = 4459
const nsldsFSEOGOverpaymentContactLength2526 int = 8

// Field # 587
const nsldsPerkinsOverpaymentFlagStartIndex2526 int = 4467
const nsldsPerkinsOverpaymentFlagLength2526 int = 1

// Field # 588
const nsldsPerkinsOverpaymentContactStartIndex2526 int = 4468
const nsldsPerkinsOverpaymentContactLength2526 int = 8

// Field # 589
const nsldsTEACHGrantOverpaymentFlagStartIndex2526 int = 4476
const nsldsTEACHGrantOverpaymentFlagLength2526 int = 1

// Field # 590
const nsldsTEACHGrantOverpaymentContactStartIndex2526 int = 4477
const nsldsTEACHGrantOverpaymentContactLength2526 int = 8

// Field # 591
const nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagStartIndex2526 int = 4485
const nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagLength2526 int = 1

// Field # 592
const nsldsIraqAndAfghanistanServiceGrantOverpaymentContactStartIndex2526 int = 4486
const nsldsIraqAndAfghanistanServiceGrantOverpaymentContactLength2526 int = 8

// Field # 593
const nsldsDefaultedLoanFlagStartIndex2526 int = 4494
const nsldsDefaultedLoanFlagLength2526 int = 1

// Field # 594
const nsldsDischargedLoanFlagStartIndex2526 int = 4495
const nsldsDischargedLoanFlagLength2526 int = 1

// Field # 595
const nsldsFraudLoanFlagStartIndex2526 int = 4496
const nsldsFraudLoanFlagLength2526 int = 1

// Field # 596
const nsldsSatisfactoryArrangementsFlagStartIndex2526 int = 4497
const nsldsSatisfactoryArrangementsFlagLength2526 int = 1

// Field # 597
const nsldsActiveBankruptcyFlagStartIndex2526 int = 4498
const nsldsActiveBankruptcyFlagLength2526 int = 1

// Field # 598
const nsldsTEACHGrantConvertedToLoanFlagStartIndex2526 int = 4499
const nsldsTEACHGrantConvertedToLoanFlagLength2526 int = 1

// Field # 599
const nsldsAggregateSubsidizedOutstandingPrincipalBalanceStartIndex2526 int = 4500
const nsldsAggregateSubsidizedOutstandingPrincipalBalanceLength2526 int = 6

// Field # 600
const nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceStartIndex2526 int = 4506
const nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceLength2526 int = 6

// Field # 601
const nsldsAggregateCombinedOutstandingPrincipalBalanceStartIndex2526 int = 4512
const nsldsAggregateCombinedOutstandingPrincipalBalanceLength2526 int = 6

// Field # 602
const nsldsAggregateUnallocatedConsolidatedOutstandingPrincipalBalanceStartIndex2526 int = 4518
const nsldsAggregateUnallocatedConsolidatedOutstandingPrincipalBalanceLength2526 int = 6

// Field # 603
const nsldsAggregateTEACHLoanPrincipalBalanceStartIndex2526 int = 4524
const nsldsAggregateTEACHLoanPrincipalBalanceLength2526 int = 6

// Field # 604
const nsldsAggregateSubsidizedPendingDisbursementStartIndex2526 int = 4530
const nsldsAggregateSubsidizedPendingDisbursementLength2526 int = 6

// Field # 605
const nsldsAggregateUnsubsidizedPendingDisbursementStartIndex2526 int = 4536
const nsldsAggregateUnsubsidizedPendingDisbursementLength2526 int = 6

// Field # 606
const nsldsAggregateCombinedPendingDisbursementStartIndex2526 int = 4542
const nsldsAggregateCombinedPendingDisbursementLength2526 int = 6

// Field # 607
const nsldsAggregateSubsidizedTotalStartIndex2526 int = 4548
const nsldsAggregateSubsidizedTotalLength2526 int = 6

// Field # 608
const nsldsAggregateUnsubsidizedTotalStartIndex2526 int = 4554
const nsldsAggregateUnsubsidizedTotalLength2526 int = 6

// Field # 609
const nsldsAggregateCombinedTotalStartIndex2526 int = 4560
const nsldsAggregateCombinedTotalLength2526 int = 6

// Field # 610
const nsldsUnallocatedConsolidatedTotalStartIndex2526 int = 4566
const nsldsUnallocatedConsolidatedTotalLength2526 int = 6

// Field # 611
const nsldsTEACHLoanTotalStartIndex2526 int = 4572
const nsldsTEACHLoanTotalLength2526 int = 6

// Field # 612
const nsldsPerkinsTotalDisbursementsStartIndex2526 int = 4578
const nsldsPerkinsTotalDisbursementsLength2526 int = 6

// Field # 613
const nsldsPerkinsCurrentYearDisbursementAmountStartIndex2526 int = 4584
const nsldsPerkinsCurrentYearDisbursementAmountLength2526 int = 6

// Field # 614
const nsldsAggregateTEACHGrantUndergraduateDisbursedTotalStartIndex2526 int = 4590
const nsldsAggregateTEACHGrantUndergraduateDisbursedTotalLength2526 int = 6

// Field # 615
const nsldsAggregateTEACHGraduateDisbursementAmountStartIndex2526 int = 4596
const nsldsAggregateTEACHGraduateDisbursementAmountLength2526 int = 6

// Field # 616
const nsldsDefaultedLoanChangeFlagStartIndex2526 int = 4602
const nsldsDefaultedLoanChangeFlagLength2526 int = 1

// Field # 617
const nsldsFraudLoanChangeFlagStartIndex2526 int = 4603
const nsldsFraudLoanChangeFlagLength2526 int = 1

// Field # 618
const nsldsDischargedLoanChangeFlagStartIndex2526 int = 4604
const nsldsDischargedLoanChangeFlagLength2526 int = 1

// Field # 619
const nsldsLoanSatisfactoryRepaymentChangeFlagStartIndex2526 int = 4605
const nsldsLoanSatisfactoryRepaymentChangeFlagLength2526 int = 1

// Field # 620
const nsldsActiveBankruptcyChangeFlagStartIndex2526 int = 4606
const nsldsActiveBankruptcyChangeFlagLength2526 int = 1

// Field # 621
const nsldsTEACHGrantToLoanConversionChangeFlagStartIndex2526 int = 4607
const nsldsTEACHGrantToLoanConversionChangeFlagLength2526 int = 1

// Field # 622
const nsldsOverpaymentsChangeFlagStartIndex2526 int = 4608
const nsldsOverpaymentsChangeFlagLength2526 int = 1

// Field # 623
const nsldsAggregateLoanChangeFlagStartIndex2526 int = 4609
const nsldsAggregateLoanChangeFlagLength2526 int = 1

// Field # 624
const nsldsPerkinsLoanChangeFlagStartIndex2526 int = 4610
const nsldsPerkinsLoanChangeFlagLength2526 int = 1

// Field # 625
const nsldsPellPaymentChangeFlagStartIndex2526 int = 4611
const nsldsPellPaymentChangeFlagLength2526 int = 1

// Field # 626
const nsldsTEACHGrantChangeFlagStartIndex2526 int = 4612
const nsldsTEACHGrantChangeFlagLength2526 int = 1

// Field # 627
const nsldsAdditionalPellFlagStartIndex2526 int = 4613
const nsldsAdditionalPellFlagLength2526 int = 1

// Field # 628
const nsldsAdditionalLoansFlagStartIndex2526 int = 4614
const nsldsAdditionalLoansFlagLength2526 int = 1

// Field # 629
const nsldsAdditionalTEACHGrantFlagStartIndex2526 int = 4615
const nsldsAdditionalTEACHGrantFlagLength2526 int = 1

// Field # 630
const nsldsDirectLoanMasterPromNoteFlagStartIndex2526 int = 4616
const nsldsDirectLoanMasterPromNoteFlagLength2526 int = 1

// Field # 631
const nsldsDirectLoanPLUSMasterPromNoteFlagStartIndex2526 int = 4617
const nsldsDirectLoanPLUSMasterPromNoteFlagLength2526 int = 1

// Field # 632
const nsldsDirectLoanGraduatePLUSMasterPromNoteFlagStartIndex2526 int = 4618
const nsldsDirectLoanGraduatePLUSMasterPromNoteFlagLength2526 int = 1

// Field # 633
const nsldsUndergraduateSubsidizedLoanLimitFlagStartIndex2526 int = 4619
const nsldsUndergraduateSubsidizedLoanLimitFlagLength2526 int = 1

// Field # 634
const nsldsUndergraduateCombinedLoanLimitFlagStartIndex2526 int = 4620
const nsldsUndergraduateCombinedLoanLimitFlagLength2526 int = 1

// Field # 635
const nsldsGraduateSubsidizedLoanLimitFlagStartIndex2526 int = 4621
const nsldsGraduateSubsidizedLoanLimitFlagLength2526 int = 1

// Field # 636
const nsldsGraduateCombinedLoanLimitFlagStartIndex2526 int = 4622
const nsldsGraduateCombinedLoanLimitFlagLength2526 int = 1

// Field # 637
const nsldsLEULimitIndicatorStartIndex2526 int = 4623
const nsldsLEULimitIndicatorLength2526 int = 1

// Field # 638
const nsldsPellLifetimeEligibilityUsedStartIndex2526 int = 4624
const nsldsPellLifetimeEligibilityUsedLength2526 int = 7

// Field # 639
const nsldsSULAFlagStartIndex2526 int = 4631
const nsldsSULAFlagLength2526 int = 1

// Field # 640
const nsldsSubsidizedLimitEligibilityFlagStartIndex2526 int = 4632
const nsldsSubsidizedLimitEligibilityFlagLength2526 int = 6

// Field # 641
const nsldsUnusualEnrollmentHistoryFlagStartIndex2526 int = 4638
const nsldsUnusualEnrollmentHistoryFlagLength2526 int = 1

// Field # 643
const nsldsPellSequenceNumber1StartIndex2526 int = 4659
const nsldsPellSequenceNumber1Length2526 int = 2

// Field # 644
const nsldsPellVerificationFlag1StartIndex2526 int = 4661
const nsldsPellVerificationFlag1Length2526 int = 3

// Field # 645
const nsldsSAI1StartIndex2526 int = 4664
const nsldsSAI1Length2526 int = 6

// Field # 646
const nsldsPellSchoolCode1StartIndex2526 int = 4670
const nsldsPellSchoolCode1Length2526 int = 8

// Field # 647
const nsldsPellTransactionNumber1StartIndex2526 int = 4678
const nsldsPellTransactionNumber1Length2526 int = 2

// Field # 648
const nsldsPellLastDisbursementDate1StartIndex2526 int = 4680
const nsldsPellLastDisbursementDate1Length2526 int = 8

// Field # 649
const nsldsPellScheduledAmount1StartIndex2526 int = 4688
const nsldsPellScheduledAmount1Length2526 int = 6

// Field # 650
const nsldsPellAmountPaidToDate1StartIndex2526 int = 4694
const nsldsPellAmountPaidToDate1Length2526 int = 6

// Field # 651
const nsldsPellPercentEligibilityUsedDecimal1StartIndex2526 int = 4700
const nsldsPellPercentEligibilityUsedDecimal1Length2526 int = 7

// Field # 652
const nsldsPellAwardAmount1StartIndex2526 int = 4707
const nsldsPellAwardAmount1Length2526 int = 6

// Field # 653
const nsldsAdditionalEligibilityIndicator1StartIndex2526 int = 4713
const nsldsAdditionalEligibilityIndicator1Length2526 int = 1

// Field # 655
const nsldsPellSequenceNumber2StartIndex2526 int = 4734
const nsldsPellSequenceNumber2Length2526 int = 2

// Field # 656
const nsldsPellVerificationFlag2StartIndex2526 int = 4736
const nsldsPellVerificationFlag2Length2526 int = 3

// Field # 657
const nsldsSAI2StartIndex2526 int = 4739
const nsldsSAI2Length2526 int = 6

// Field # 658
const nsldsPellSchoolCode2StartIndex2526 int = 4745
const nsldsPellSchoolCode2Length2526 int = 8

// Field # 659
const nsldsPellTransactionNumber2StartIndex2526 int = 4753
const nsldsPellTransactionNumber2Length2526 int = 2

// Field # 660
const nsldsPellLastDisbursementDate2StartIndex2526 int = 4755
const nsldsPellLastDisbursementDate2Length2526 int = 8

// Field # 661
const nsldsPellScheduledAmount2StartIndex2526 int = 4763
const nsldsPellScheduledAmount2Length2526 int = 6

// Field # 662
const nsldsPellAmountPaidToDate2StartIndex2526 int = 4769
const nsldsPellAmountPaidToDate2Length2526 int = 6

// Field # 663
const nsldsPellPercentEligibilityUsedDecimal2StartIndex2526 int = 4775
const nsldsPellPercentEligibilityUsedDecimal2Length2526 int = 7

// Field # 664
const nsldsPellAwardAmount2StartIndex2526 int = 4782
const nsldsPellAwardAmount2Length2526 int = 6

// Field # 665
const nsldsAdditionalEligibilityIndicator2StartIndex2526 int = 4788
const nsldsAdditionalEligibilityIndicator2Length2526 int = 1

// Field # 667
const nsldsPellSequenceNumber3StartIndex2526 int = 4809
const nsldsPellSequenceNumber3Length2526 int = 2

// Field # 668
const nsldsPellVerificationFlag3StartIndex2526 int = 4811
const nsldsPellVerificationFlag3Length2526 int = 3

// Field # 669
const nsldsSAI3StartIndex2526 int = 4814
const nsldsSAI3Length2526 int = 6

// Field # 670
const nsldsPellSchoolCode3StartIndex2526 int = 4820
const nsldsPellSchoolCode3Length2526 int = 8

// Field # 671
const nsldsPellTransactionNumber3StartIndex2526 int = 4828
const nsldsPellTransactionNumber3Length2526 int = 2

// Field # 672
const nsldsPellLastDisbursementDate3StartIndex2526 int = 4830
const nsldsPellLastDisbursementDate3Length2526 int = 8

// Field # 673
const nsldsPellScheduledAmount3StartIndex2526 int = 4838
const nsldsPellScheduledAmount3Length2526 int = 6

// Field # 674
const nsldsPellAmountPaidToDate3StartIndex2526 int = 4844
const nsldsPellAmountPaidToDate3Length2526 int = 6

// Field # 675
const nsldsPellPercentEligibilityUsedDecimal3StartIndex2526 int = 4850
const nsldsPellPercentEligibilityUsedDecimal3Length2526 int = 7

// Field # 676
const nsldsPellAwardAmount3StartIndex2526 int = 4857
const nsldsPellAwardAmount3Length2526 int = 6

// Field # 677
const nsldsAdditionalEligibilityIndicator3StartIndex2526 int = 4863
const nsldsAdditionalEligibilityIndicator3Length2526 int = 1

// Field # 679
const nsldsTEACHGrantSequence1StartIndex2526 int = 4884
const nsldsTEACHGrantSequence1Length2526 int = 2

// Field # 680
const nsldsTEACHGrantSchoolCode1StartIndex2526 int = 4886
const nsldsTEACHGrantSchoolCode1Length2526 int = 8

// Field # 681
const nsldsTEACHGrantTransactionNumber1StartIndex2526 int = 4894
const nsldsTEACHGrantTransactionNumber1Length2526 int = 2

// Field # 682
const nsldsTEACHGrantLastDisbursementDate1StartIndex2526 int = 4896
const nsldsTEACHGrantLastDisbursementDate1Length2526 int = 8

// Field # 683
const nsldsTEACHGrantScheduledAmount1StartIndex2526 int = 4904
const nsldsTEACHGrantScheduledAmount1Length2526 int = 6

// Field # 684
const nsldsTEACHGrantAmountPaidToDate1StartIndex2526 int = 4910
const nsldsTEACHGrantAmountPaidToDate1Length2526 int = 6

// Field # 685
const nsldsTEACHGrantAwardAmount1StartIndex2526 int = 4916
const nsldsTEACHGrantAwardAmount1Length2526 int = 6

// Field # 686
const nsldsTEACHGrantAcademicYearLevel1StartIndex2526 int = 4922
const nsldsTEACHGrantAcademicYearLevel1Length2526 int = 1

// Field # 687
const nsldsTEACHGrantAwardYear1StartIndex2526 int = 4923
const nsldsTEACHGrantAwardYear1Length2526 int = 4

// Field # 688
const nsldsTEACHGrantLoanConversionFlag1StartIndex2526 int = 4927
const nsldsTEACHGrantLoanConversionFlag1Length2526 int = 1

// Field # 689
const nsldsTEACHGrantDischargeCode1StartIndex2526 int = 4928
const nsldsTEACHGrantDischargeCode1Length2526 int = 4

// Field # 690
const nsldsTEACHGrantDischargeAmount1StartIndex2526 int = 4932
const nsldsTEACHGrantDischargeAmount1Length2526 int = 6

// Field # 691
const nsldsTEACHGrantAdjustedDisbursement1StartIndex2526 int = 4938
const nsldsTEACHGrantAdjustedDisbursement1Length2526 int = 6

// Field # 693
const nsldsTEACHGrantSequence2StartIndex2526 int = 4964
const nsldsTEACHGrantSequence2Length2526 int = 2

// Field # 694
const nsldsTEACHGrantSchoolCode2StartIndex2526 int = 4966
const nsldsTEACHGrantSchoolCode2Length2526 int = 8

// Field # 695
const nsldsTEACHGrantTransactionNumber2StartIndex2526 int = 4974
const nsldsTEACHGrantTransactionNumber2Length2526 int = 2

// Field # 696
const nsldsTEACHGrantLastDisbursementDate2StartIndex2526 int = 4976
const nsldsTEACHGrantLastDisbursementDate2Length2526 int = 8

// Field # 697
const nsldsTEACHGrantScheduledAmount2StartIndex2526 int = 4984
const nsldsTEACHGrantScheduledAmount2Length2526 int = 6

// Field # 698
const nsldsTEACHGrantAmountPaidToDate2StartIndex2526 int = 4990
const nsldsTEACHGrantAmountPaidToDate2Length2526 int = 6

// Field # 699
const nsldsTEACHGrantAwardAmount2StartIndex2526 int = 4996
const nsldsTEACHGrantAwardAmount2Length2526 int = 6

// Field # 700
const nsldsTEACHGrantAcademicYearLevel2StartIndex2526 int = 5002
const nsldsTEACHGrantAcademicYearLevel2Length2526 int = 1

// Field # 701
const nsldsTEACHGrantAwardYear2StartIndex2526 int = 5003
const nsldsTEACHGrantAwardYear2Length2526 int = 4

// Field # 702
const nsldsTEACHGrantLoanConversionFlag2StartIndex2526 int = 5007
const nsldsTEACHGrantLoanConversionFlag2Length2526 int = 1

// Field # 703
const nsldsTEACHGrantDischargeCode2StartIndex2526 int = 5008
const nsldsTEACHGrantDischargeCode2Length2526 int = 4

// Field # 704
const nsldsTEACHGrantDischargeAmount2StartIndex2526 int = 5012
const nsldsTEACHGrantDischargeAmount2Length2526 int = 6

// Field # 705
const nsldsTEACHGrantAdjustedDisbursement2StartIndex2526 int = 5018
const nsldsTEACHGrantAdjustedDisbursement2Length2526 int = 6

// Field # 707
const nsldsTEACHGrantSequence3StartIndex2526 int = 5044
const nsldsTEACHGrantSequence3Length2526 int = 2

// Field # 708
const nsldsTEACHGrantSchoolCode3StartIndex2526 int = 5046
const nsldsTEACHGrantSchoolCode3Length2526 int = 8

// Field # 709
const nsldsTEACHGrantTransactionNumber3StartIndex2526 int = 5054
const nsldsTEACHGrantTransactionNumber3Length2526 int = 2

// Field # 710
const nsldsTEACHGrantLastDisbursementDate3StartIndex2526 int = 5056
const nsldsTEACHGrantLastDisbursementDate3Length2526 int = 8

// Field # 711
const nsldsTEACHGrantScheduledAmount3StartIndex2526 int = 5064
const nsldsTEACHGrantScheduledAmount3Length2526 int = 6

// Field # 712
const nsldsTEACHGrantAmountPaidToDate3StartIndex2526 int = 5070
const nsldsTEACHGrantAmountPaidToDate3Length2526 int = 6

// Field # 713
const nsldsTEACHGrantAwardAmount3StartIndex2526 int = 5076
const nsldsTEACHGrantAwardAmount3Length2526 int = 6

// Field # 714
const nsldsTEACHGrantAcademicYearLevel3StartIndex2526 int = 5082
const nsldsTEACHGrantAcademicYearLevel3Length2526 int = 1

// Field # 715
const nsldsTEACHGrantAwardYear3StartIndex2526 int = 5083
const nsldsTEACHGrantAwardYear3Length2526 int = 4

// Field # 716
const nsldsTEACHGrantLoanConversionFlag3StartIndex2526 int = 5087
const nsldsTEACHGrantLoanConversionFlag3Length2526 int = 1

// Field # 717
const nsldsTEACHGrantDischargeCode3StartIndex2526 int = 5088
const nsldsTEACHGrantDischargeCode3Length2526 int = 4

// Field # 718
const nsldsTEACHGrantDischargeAmount3StartIndex2526 int = 5092
const nsldsTEACHGrantDischargeAmount3Length2526 int = 6

// Field # 719
const nsldsTEACHGrantAdjustedDisbursement3StartIndex2526 int = 5098
const nsldsTEACHGrantAdjustedDisbursement3Length2526 int = 6

// Field # 721
const nsldsLoanSequenceNumber1StartIndex2526 int = 5124
const nsldsLoanSequenceNumber1Length2526 int = 2

// Field # 722
const nsldsLoanDefaultedRecentIndicator1StartIndex2526 int = 5126
const nsldsLoanDefaultedRecentIndicator1Length2526 int = 1

// Field # 723
const nsldsLoanChangeFlag1StartIndex2526 int = 5127
const nsldsLoanChangeFlag1Length2526 int = 1

// Field # 724
const nsldsLoanTypeCode1StartIndex2526 int = 5128
const nsldsLoanTypeCode1Length2526 int = 2

// Field # 725
const nsldsLoanNetAmount1StartIndex2526 int = 5130
const nsldsLoanNetAmount1Length2526 int = 6

// Field # 726
const nsldsLoanCurrentStatusCode1StartIndex2526 int = 5136
const nsldsLoanCurrentStatusCode1Length2526 int = 2

// Field # 727
const nsldsLoanCurrentStatusDate1StartIndex2526 int = 5138
const nsldsLoanCurrentStatusDate1Length2526 int = 8

// Field # 728
const nsldsLoanOutstandingPrincipalBalance1StartIndex2526 int = 5146
const nsldsLoanOutstandingPrincipalBalance1Length2526 int = 6

// Field # 729
const nsldsLoanOutstandingPrincipalBalanceDate1StartIndex2526 int = 5152
const nsldsLoanOutstandingPrincipalBalanceDate1Length2526 int = 8

// Field # 730
const nsldsLoanPeriodBeginDate1StartIndex2526 int = 5160
const nsldsLoanPeriodBeginDate1Length2526 int = 8

// Field # 731
const nsldsLoanPeriodEndDate1StartIndex2526 int = 5168
const nsldsLoanPeriodEndDate1Length2526 int = 8

// Field # 732
const nsldsLoanGuarantyAgencyCode1StartIndex2526 int = 5176
const nsldsLoanGuarantyAgencyCode1Length2526 int = 3

// Field # 733
const nsldsLoanContactType1StartIndex2526 int = 5179
const nsldsLoanContactType1Length2526 int = 3

// Field # 734
const nsldsLoanSchoolCode1StartIndex2526 int = 5182
const nsldsLoanSchoolCode1Length2526 int = 8

// Field # 735
const nsldsLoanContactCode1StartIndex2526 int = 5190
const nsldsLoanContactCode1Length2526 int = 8

// Field # 736
const nsldsLoanGradeLevel1StartIndex2526 int = 5198
const nsldsLoanGradeLevel1Length2526 int = 3

// Field # 737
const nsldsLoanAdditionalUnsubsidizedFlag1StartIndex2526 int = 5201
const nsldsLoanAdditionalUnsubsidizedFlag1Length2526 int = 1

// Field # 738
const nsldsLoanCapitalizedInterestFlag1StartIndex2526 int = 5202
const nsldsLoanCapitalizedInterestFlag1Length2526 int = 1

// Field # 739
const nsldsLoanDisbursementAmount1StartIndex2526 int = 5203
const nsldsLoanDisbursementAmount1Length2526 int = 6

// Field # 740
const nsldsLoanDisbursementDate1StartIndex2526 int = 5209
const nsldsLoanDisbursementDate1Length2526 int = 8

// Field # 741
const nsldsLoanConfirmedLoanSubsidyStatus1StartIndex2526 int = 5217
const nsldsLoanConfirmedLoanSubsidyStatus1Length2526 int = 1

// Field # 742
const nsldsLoanSubsidyStatusDate1StartIndex2526 int = 5218
const nsldsLoanSubsidyStatusDate1Length2526 int = 8

// Field # 744
const nsldsLoanSequenceNumber2StartIndex2526 int = 5246
const nsldsLoanSequenceNumber2Length2526 int = 2

// Field # 745
const nsldsLoanDefaultedRecentIndicator2StartIndex2526 int = 5248
const nsldsLoanDefaultedRecentIndicator2Length2526 int = 1

// Field # 746
const nsldsLoanChangeFlag2StartIndex2526 int = 5249
const nsldsLoanChangeFlag2Length2526 int = 1

// Field # 747
const nsldsLoanTypeCode2StartIndex2526 int = 5250
const nsldsLoanTypeCode2Length2526 int = 2

// Field # 748
const nsldsLoanNetAmount2StartIndex2526 int = 5252
const nsldsLoanNetAmount2Length2526 int = 6

// Field # 749
const nsldsLoanCurrentStatusCode2StartIndex2526 int = 5258
const nsldsLoanCurrentStatusCode2Length2526 int = 2

// Field # 750
const nsldsLoanCurrentStatusDate2StartIndex2526 int = 5260
const nsldsLoanCurrentStatusDate2Length2526 int = 8

// Field # 751
const nsldsLoanOutstandingPrincipalBalance2StartIndex2526 int = 5268
const nsldsLoanOutstandingPrincipalBalance2Length2526 int = 6

// Field # 752
const nsldsLoanOutstandingPrincipalBalanceDate2StartIndex2526 int = 5274
const nsldsLoanOutstandingPrincipalBalanceDate2Length2526 int = 8

// Field # 753
const nsldsLoanPeriodBeginDate2StartIndex2526 int = 5282
const nsldsLoanPeriodBeginDate2Length2526 int = 8

// Field # 754
const nsldsLoanPeriodEndDate2StartIndex2526 int = 5290
const nsldsLoanPeriodEndDate2Length2526 int = 8

// Field # 755
const nsldsLoanGuarantyAgencyCode2StartIndex2526 int = 5298
const nsldsLoanGuarantyAgencyCode2Length2526 int = 3

// Field # 756
const nsldsLoanContactType2StartIndex2526 int = 5301
const nsldsLoanContactType2Length2526 int = 3

// Field # 757
const nsldsLoanSchoolCode2StartIndex2526 int = 5304
const nsldsLoanSchoolCode2Length2526 int = 8

// Field # 758
const nsldsLoanContactCode2StartIndex2526 int = 5312
const nsldsLoanContactCode2Length2526 int = 8

// Field # 759
const nsldsLoanGradeLevel2StartIndex2526 int = 5320
const nsldsLoanGradeLevel2Length2526 int = 3

// Field # 760
const nsldsLoanAdditionalUnsubsidizedFlag2StartIndex2526 int = 5323
const nsldsLoanAdditionalUnsubsidizedFlag2Length2526 int = 1

// Field # 761
const nsldsLoanCapitalizedInterestFlag2StartIndex2526 int = 5324
const nsldsLoanCapitalizedInterestFlag2Length2526 int = 1

// Field # 762
const nsldsLoanDisbursementAmount2StartIndex2526 int = 5325
const nsldsLoanDisbursementAmount2Length2526 int = 6

// Field # 763
const nsldsLoanDisbursementDate2StartIndex2526 int = 5331
const nsldsLoanDisbursementDate2Length2526 int = 8

// Field # 764
const nsldsLoanConfirmedLoanSubsidyStatus2StartIndex2526 int = 5339
const nsldsLoanConfirmedLoanSubsidyStatus2Length2526 int = 1

// Field # 765
const nsldsLoanSubsidyStatusDate2StartIndex2526 int = 5340
const nsldsLoanSubsidyStatusDate2Length2526 int = 8

// Field # 767
const nsldsLoanSequenceNumber3StartIndex2526 int = 5368
const nsldsLoanSequenceNumber3Length2526 int = 2

// Field # 768
const nsldsLoanDefaultedRecentIndicator3StartIndex2526 int = 5370
const nsldsLoanDefaultedRecentIndicator3Length2526 int = 1

// Field # 769
const nsldsLoanChangeFlag3StartIndex2526 int = 5371
const nsldsLoanChangeFlag3Length2526 int = 1

// Field # 770
const nsldsLoanTypeCode3StartIndex2526 int = 5372
const nsldsLoanTypeCode3Length2526 int = 2

// Field # 771
const nsldsLoanNetAmount3StartIndex2526 int = 5374
const nsldsLoanNetAmount3Length2526 int = 6

// Field # 772
const nsldsLoanCurrentStatusCode3StartIndex2526 int = 5380
const nsldsLoanCurrentStatusCode3Length2526 int = 2

// Field # 773
const nsldsLoanCurrentStatusDate3StartIndex2526 int = 5382
const nsldsLoanCurrentStatusDate3Length2526 int = 8

// Field # 774
const nsldsLoanOutstandingPrincipalBalance3StartIndex2526 int = 5390
const nsldsLoanOutstandingPrincipalBalance3Length2526 int = 6

// Field # 775
const nsldsLoanOutstandingPrincipalBalanceDate3StartIndex2526 int = 5396
const nsldsLoanOutstandingPrincipalBalanceDate3Length2526 int = 8

// Field # 776
const nsldsLoanPeriodBeginDate3StartIndex2526 int = 5404
const nsldsLoanPeriodBeginDate3Length2526 int = 8

// Field # 777
const nsldsLoanPeriodEndDate3StartIndex2526 int = 5412
const nsldsLoanPeriodEndDate3Length2526 int = 8

// Field # 778
const nsldsLoanGuarantyAgencyCode3StartIndex2526 int = 5420
const nsldsLoanGuarantyAgencyCode3Length2526 int = 3

// Field # 779
const nsldsLoanContactType3StartIndex2526 int = 5423
const nsldsLoanContactType3Length2526 int = 3

// Field # 780
const nsldsLoanSchoolCode3StartIndex2526 int = 5426
const nsldsLoanSchoolCode3Length2526 int = 8

// Field # 781
const nsldsLoanContactCode3StartIndex2526 int = 5434
const nsldsLoanContactCode3Length2526 int = 8

// Field # 782
const nsldsLoanGradeLevel3StartIndex2526 int = 5442
const nsldsLoanGradeLevel3Length2526 int = 3

// Field # 783
const nsldsLoanAdditionalUnsubsidizedFlag3StartIndex2526 int = 5445
const nsldsLoanAdditionalUnsubsidizedFlag3Length2526 int = 1

// Field # 784
const nsldsLoanCapitalizedInterestFlag3StartIndex2526 int = 5446
const nsldsLoanCapitalizedInterestFlag3Length2526 int = 1

// Field # 785
const nsldsLoanDisbursementAmount3StartIndex2526 int = 5447
const nsldsLoanDisbursementAmount3Length2526 int = 6

// Field # 786
const nsldsLoanDisbursementDate3StartIndex2526 int = 5453
const nsldsLoanDisbursementDate3Length2526 int = 8

// Field # 787
const nsldsLoanConfirmedLoanSubsidyStatus3StartIndex2526 int = 5461
const nsldsLoanConfirmedLoanSubsidyStatus3Length2526 int = 1

// Field # 788
const nsldsLoanSubsidyStatusDate3StartIndex2526 int = 5462
const nsldsLoanSubsidyStatusDate3Length2526 int = 8

// Field # 790
const nsldsLoanSequenceNumber4StartIndex2526 int = 5490
const nsldsLoanSequenceNumber4Length2526 int = 2

// Field # 791
const nsldsLoanDefaultedRecentIndicator4StartIndex2526 int = 5492
const nsldsLoanDefaultedRecentIndicator4Length2526 int = 1

// Field # 792
const nsldsLoanChangeFlag4StartIndex2526 int = 5493
const nsldsLoanChangeFlag4Length2526 int = 1

// Field # 793
const nsldsLoanTypeCode4StartIndex2526 int = 5494
const nsldsLoanTypeCode4Length2526 int = 2

// Field # 794
const nsldsLoanNetAmount4StartIndex2526 int = 5496
const nsldsLoanNetAmount4Length2526 int = 6

// Field # 795
const nsldsLoanCurrentStatusCode4StartIndex2526 int = 5502
const nsldsLoanCurrentStatusCode4Length2526 int = 2

// Field # 796
const nsldsLoanCurrentStatusDate4StartIndex2526 int = 5504
const nsldsLoanCurrentStatusDate4Length2526 int = 8

// Field # 797
const nsldsLoanOutstandingPrincipalBalance4StartIndex2526 int = 5512
const nsldsLoanOutstandingPrincipalBalance4Length2526 int = 6

// Field # 798
const nsldsLoanOutstandingPrincipalBalanceDate4StartIndex2526 int = 5518
const nsldsLoanOutstandingPrincipalBalanceDate4Length2526 int = 8

// Field # 799
const nsldsLoanPeriodBeginDate4StartIndex2526 int = 5526
const nsldsLoanPeriodBeginDate4Length2526 int = 8

// Field # 800
const nsldsLoanPeriodEndDate4StartIndex2526 int = 5534
const nsldsLoanPeriodEndDate4Length2526 int = 8

// Field # 801
const nsldsLoanGuarantyAgencyCode4StartIndex2526 int = 5542
const nsldsLoanGuarantyAgencyCode4Length2526 int = 3

// Field # 802
const nsldsLoanContactType4StartIndex2526 int = 5545
const nsldsLoanContactType4Length2526 int = 3

// Field # 803
const nsldsLoanSchoolCode4StartIndex2526 int = 5548
const nsldsLoanSchoolCode4Length2526 int = 8

// Field # 804
const nsldsLoanContactCode4StartIndex2526 int = 5556
const nsldsLoanContactCode4Length2526 int = 8

// Field # 805
const nsldsLoanGradeLevel4StartIndex2526 int = 5564
const nsldsLoanGradeLevel4Length2526 int = 3

// Field # 806
const nsldsLoanAdditionalUnsubsidizedFlag4StartIndex2526 int = 5567
const nsldsLoanAdditionalUnsubsidizedFlag4Length2526 int = 1

// Field # 807
const nsldsLoanCapitalizedInterestFlag4StartIndex2526 int = 5568
const nsldsLoanCapitalizedInterestFlag4Length2526 int = 1

// Field # 808
const nsldsLoanDisbursementAmount4StartIndex2526 int = 5569
const nsldsLoanDisbursementAmount4Length2526 int = 6

// Field # 809
const nsldsLoanDisbursementDate4StartIndex2526 int = 5575
const nsldsLoanDisbursementDate4Length2526 int = 8

// Field # 810
const nsldsLoanConfirmedLoanSubsidyStatus4StartIndex2526 int = 5583
const nsldsLoanConfirmedLoanSubsidyStatus4Length2526 int = 1

// Field # 811
const nsldsLoanSubsidyStatusDate4StartIndex2526 int = 5584
const nsldsLoanSubsidyStatusDate4Length2526 int = 8

// Field # 813
const nsldsLoanSequenceNumber5StartIndex2526 int = 5612
const nsldsLoanSequenceNumber5Length2526 int = 2

// Field # 814
const nsldsLoanDefaultedRecentIndicator5StartIndex2526 int = 5614
const nsldsLoanDefaultedRecentIndicator5Length2526 int = 1

// Field # 815
const nsldsLoanChangeFlag5StartIndex2526 int = 5615
const nsldsLoanChangeFlag5Length2526 int = 1

// Field # 816
const nsldsLoanTypeCode5StartIndex2526 int = 5616
const nsldsLoanTypeCode5Length2526 int = 2

// Field # 817
const nsldsLoanNetAmount5StartIndex2526 int = 5618
const nsldsLoanNetAmount5Length2526 int = 6

// Field # 818
const nsldsLoanCurrentStatusCode5StartIndex2526 int = 5624
const nsldsLoanCurrentStatusCode5Length2526 int = 2

// Field # 819
const nsldsLoanCurrentStatusDate5StartIndex2526 int = 5626
const nsldsLoanCurrentStatusDate5Length2526 int = 8

// Field # 820
const nsldsLoanOutstandingPrincipalBalance5StartIndex2526 int = 5634
const nsldsLoanOutstandingPrincipalBalance5Length2526 int = 6

// Field # 821
const nsldsLoanOutstandingPrincipalBalanceDate5StartIndex2526 int = 5640
const nsldsLoanOutstandingPrincipalBalanceDate5Length2526 int = 8

// Field # 822
const nsldsLoanPeriodBeginDate5StartIndex2526 int = 5648
const nsldsLoanPeriodBeginDate5Length2526 int = 8

// Field # 823
const nsldsLoanPeriodEndDate5StartIndex2526 int = 5656
const nsldsLoanPeriodEndDate5Length2526 int = 8

// Field # 824
const nsldsLoanGuarantyAgencyCode5StartIndex2526 int = 5664
const nsldsLoanGuarantyAgencyCode5Length2526 int = 3

// Field # 825
const nsldsLoanContactType5StartIndex2526 int = 5667
const nsldsLoanContactType5Length2526 int = 3

// Field # 826
const nsldsLoanSchoolCode5StartIndex2526 int = 5670
const nsldsLoanSchoolCode5Length2526 int = 8

// Field # 827
const nsldsLoanContactCode5StartIndex2526 int = 5678
const nsldsLoanContactCode5Length2526 int = 8

// Field # 828
const nsldsLoanGradeLevel5StartIndex2526 int = 5686
const nsldsLoanGradeLevel5Length2526 int = 3

// Field # 829
const nsldsLoanAdditionalUnsubsidizedFlag5StartIndex2526 int = 5689
const nsldsLoanAdditionalUnsubsidizedFlag5Length2526 int = 1

// Field # 830
const nsldsLoanCapitalizedInterestFlag5StartIndex2526 int = 5690
const nsldsLoanCapitalizedInterestFlag5Length2526 int = 1

// Field # 831
const nsldsLoanDisbursementAmount5StartIndex2526 int = 5691
const nsldsLoanDisbursementAmount5Length2526 int = 6

// Field # 832
const nsldsLoanDisbursementDate5StartIndex2526 int = 5697
const nsldsLoanDisbursementDate5Length2526 int = 8

// Field # 833
const nsldsLoanConfirmedLoanSubsidyStatus5StartIndex2526 int = 5705
const nsldsLoanConfirmedLoanSubsidyStatus5Length2526 int = 1

// Field # 834
const nsldsLoanSubsidyStatusDate5StartIndex2526 int = 5706
const nsldsLoanSubsidyStatusDate5Length2526 int = 8

// Field # 836
const nsldsLoanSequenceNumber6StartIndex2526 int = 5734
const nsldsLoanSequenceNumber6Length2526 int = 2

// Field # 837
const nsldsLoanDefaultedRecentIndicator6StartIndex2526 int = 5736
const nsldsLoanDefaultedRecentIndicator6Length2526 int = 1

// Field # 838
const nsldsLoanChangeFlag6StartIndex2526 int = 5737
const nsldsLoanChangeFlag6Length2526 int = 1

// Field # 839
const nsldsLoanTypeCode6StartIndex2526 int = 5738
const nsldsLoanTypeCode6Length2526 int = 2

// Field # 840
const nsldsLoanNetAmount6StartIndex2526 int = 5740
const nsldsLoanNetAmount6Length2526 int = 6

// Field # 841
const nsldsLoanCurrentStatusCode6StartIndex2526 int = 5746
const nsldsLoanCurrentStatusCode6Length2526 int = 2

// Field # 842
const nsldsLoanCurrentStatusDate6StartIndex2526 int = 5748
const nsldsLoanCurrentStatusDate6Length2526 int = 8

// Field # 843
const nsldsLoanOutstandingPrincipalBalance6StartIndex2526 int = 5756
const nsldsLoanOutstandingPrincipalBalance6Length2526 int = 6

// Field # 844
const nsldsLoanOutstandingPrincipalBalanceDate6StartIndex2526 int = 5762
const nsldsLoanOutstandingPrincipalBalanceDate6Length2526 int = 8

// Field # 845
const nsldsLoanPeriodBeginDate6StartIndex2526 int = 5770
const nsldsLoanPeriodBeginDate6Length2526 int = 8

// Field # 846
const nsldsLoanPeriodEndDate6StartIndex2526 int = 5778
const nsldsLoanPeriodEndDate6Length2526 int = 8

// Field # 847
const nsldsLoanGuarantyAgencyCode6StartIndex2526 int = 5786
const nsldsLoanGuarantyAgencyCode6Length2526 int = 3

// Field # 848
const nsldsLoanContactType6StartIndex2526 int = 5789
const nsldsLoanContactType6Length2526 int = 3

// Field # 849
const nsldsLoanSchoolCode6StartIndex2526 int = 5792
const nsldsLoanSchoolCode6Length2526 int = 8

// Field # 850
const nsldsLoanContactCode6StartIndex2526 int = 5800
const nsldsLoanContactCode6Length2526 int = 8

// Field # 851
const nsldsLoanGradeLevel6StartIndex2526 int = 5808
const nsldsLoanGradeLevel6Length2526 int = 3

// Field # 852
const nsldsLoanAdditionalUnsubsidizedFlag6StartIndex2526 int = 5811
const nsldsLoanAdditionalUnsubsidizedFlag6Length2526 int = 1

// Field # 853
const nsldsLoanCapitalizedInterestFlag6StartIndex2526 int = 5812
const nsldsLoanCapitalizedInterestFlag6Length2526 int = 1

// Field # 854
const nsldsLoanDisbursementAmount6StartIndex2526 int = 5813
const nsldsLoanDisbursementAmount6Length2526 int = 6

// Field # 855
const nsldsLoanDisbursementDate6StartIndex2526 int = 5819
const nsldsLoanDisbursementDate6Length2526 int = 8

// Field # 856
const nsldsLoanConfirmedLoanSubsidyStatus6StartIndex2526 int = 5827
const nsldsLoanConfirmedLoanSubsidyStatus6Length2526 int = 1

// Field # 857
const nsldsLoanSubsidyStatusDate6StartIndex2526 int = 5828
const nsldsLoanSubsidyStatusDate6Length2526 int = 8

// Field # 861
const ftiLabelStartStartIndex2526 int = 7086
const ftiLabelStartLength2526 int = 11

// Field # 862
const studentFTIMReturnedTaxYearStartIndex2526 int = 7097
const studentFTIMReturnedTaxYearLength2526 int = 4

// Field # 863
const studentFTIMFilingStatusCodeStartIndex2526 int = 7101
const studentFTIMFilingStatusCodeLength2526 int = 1

// Field # 864
const studentFTIMAdjustedGrossIncomeStartIndex2526 int = 7102
const studentFTIMAdjustedGrossIncomeLength2526 int = 10

// Field # 865
const studentFTIMNumberOfExemptionsStartIndex2526 int = 7112
const studentFTIMNumberOfExemptionsLength2526 int = 2

// Field # 866
const studentFTIMNumberOfDependentsStartIndex2526 int = 7114
const studentFTIMNumberOfDependentsLength2526 int = 2

// Field # 867
const studentFTIMTotalIncomeEarnedAmountStartIndex2526 int = 7116
const studentFTIMTotalIncomeEarnedAmountLength2526 int = 11

// Field # 868
const studentFTIMIncomeTaxPaidStartIndex2526 int = 7127
const studentFTIMIncomeTaxPaidLength2526 int = 9

// Field # 869
const studentFTIMEducationCreditsStartIndex2526 int = 7136
const studentFTIMEducationCreditsLength2526 int = 9

// Field # 870
const studentFTIMUntaxedIRADistributionsStartIndex2526 int = 7145
const studentFTIMUntaxedIRADistributionsLength2526 int = 11

// Field # 871
const studentFTIMIRADeductibleAndPaymentsStartIndex2526 int = 7156
const studentFTIMIRADeductibleAndPaymentsLength2526 int = 11

// Field # 872
const studentFTIMTaxExemptInterestStartIndex2526 int = 7167
const studentFTIMTaxExemptInterestLength2526 int = 11

// Field # 873
const studentFTIMUntaxedPensionsAmountStartIndex2526 int = 7178
const studentFTIMUntaxedPensionsAmountLength2526 int = 11

// Field # 874
const studentFTIMScheduleCNetProfitLossStartIndex2526 int = 7189
const studentFTIMScheduleCNetProfitLossLength2526 int = 12

// Field # 875
const studentFTIMScheduleAIndicatorStartIndex2526 int = 7201
const studentFTIMScheduleAIndicatorLength2526 int = 1

// Field # 876
const studentFTIMScheduleBIndicatorStartIndex2526 int = 7202
const studentFTIMScheduleBIndicatorLength2526 int = 1

// Field # 877
const studentFTIMScheduleDIndicatorStartIndex2526 int = 7203
const studentFTIMScheduleDIndicatorLength2526 int = 1

// Field # 878
const studentFTIMScheduleEIndicatorStartIndex2526 int = 7204
const studentFTIMScheduleEIndicatorLength2526 int = 1

// Field # 879
const studentFTIMScheduleFIndicatorStartIndex2526 int = 7205
const studentFTIMScheduleFIndicatorLength2526 int = 1

// Field # 880
const studentFTIMScheduleHIndicatorStartIndex2526 int = 7206
const studentFTIMScheduleHIndicatorLength2526 int = 1

// Field # 881
const studentFTIMIRSResponseCodeStartIndex2526 int = 7207
const studentFTIMIRSResponseCodeLength2526 int = 3

// Field # 882
const studentFTIMSpouseReturnedTaxYearStartIndex2526 int = 7210
const studentFTIMSpouseReturnedTaxYearLength2526 int = 4

// Field # 883
const studentFTIMSpouseFilingStatusCodeStartIndex2526 int = 7214
const studentFTIMSpouseFilingStatusCodeLength2526 int = 1

// Field # 884
const studentFTIMSpouseAdjustedGrossIncomeStartIndex2526 int = 7215
const studentFTIMSpouseAdjustedGrossIncomeLength2526 int = 10

// Field # 885
const studentFTIMSpouseNumberOfExemptionsStartIndex2526 int = 7225
const studentFTIMSpouseNumberOfExemptionsLength2526 int = 2

// Field # 886
const studentFTIMSpouseNumberOfDependentsStartIndex2526 int = 7227
const studentFTIMSpouseNumberOfDependentsLength2526 int = 2

// Field # 887
const studentFTIMSpouseTotalIncomeEarnedAmountStartIndex2526 int = 7229
const studentFTIMSpouseTotalIncomeEarnedAmountLength2526 int = 11

// Field # 888
const studentFTIMSpouseIncomeTaxPaidStartIndex2526 int = 7240
const studentFTIMSpouseIncomeTaxPaidLength2526 int = 9

// Field # 889
const studentFTIMSpouseEducationCreditsStartIndex2526 int = 7249
const studentFTIMSpouseEducationCreditsLength2526 int = 9

// Field # 890
const studentFTIMSpouseUntaxedIRADistributionsStartIndex2526 int = 7258
const studentFTIMSpouseUntaxedIRADistributionsLength2526 int = 11

// Field # 891
const studentFTIMSpouseIRADeductibleAndPaymentsStartIndex2526 int = 7269
const studentFTIMSpouseIRADeductibleAndPaymentsLength2526 int = 11

// Field # 892
const studentFTIMSpouseTaxExemptInterestStartIndex2526 int = 7280
const studentFTIMSpouseTaxExemptInterestLength2526 int = 11

// Field # 893
const studentFTIMSpouseUntaxedPensionsAmountStartIndex2526 int = 7291
const studentFTIMSpouseUntaxedPensionsAmountLength2526 int = 11

// Field # 894
const studentFTIMSpouseScheduleCNetProfitLossStartIndex2526 int = 7302
const studentFTIMSpouseScheduleCNetProfitLossLength2526 int = 12

// Field # 895
const studentFTIMSpouseScheduleAIndicatorStartIndex2526 int = 7314
const studentFTIMSpouseScheduleAIndicatorLength2526 int = 1

// Field # 896
const studentFTIMSpouseScheduleBIndicatorStartIndex2526 int = 7315
const studentFTIMSpouseScheduleBIndicatorLength2526 int = 1

// Field # 897
const studentFTIMSpouseScheduleDIndicatorStartIndex2526 int = 7316
const studentFTIMSpouseScheduleDIndicatorLength2526 int = 1

// Field # 898
const studentFTIMSpouseScheduleEIndicatorStartIndex2526 int = 7317
const studentFTIMSpouseScheduleEIndicatorLength2526 int = 1

// Field # 899
const studentFTIMSpouseScheduleFIndicatorStartIndex2526 int = 7318
const studentFTIMSpouseScheduleFIndicatorLength2526 int = 1

// Field # 900
const studentFTIMSpouseScheduleHIndicatorStartIndex2526 int = 7319
const studentFTIMSpouseScheduleHIndicatorLength2526 int = 1

// Field # 901
const studentFTIMSpouseIRSResponseCodeStartIndex2526 int = 7320
const studentFTIMSpouseIRSResponseCodeLength2526 int = 3

// Field # 902
const parentFTIMReturnedTaxYearStartIndex2526 int = 7323
const parentFTIMReturnedTaxYearLength2526 int = 4

// Field # 903
const parentFTIMFilingStatusCodeStartIndex2526 int = 7327
const parentFTIMFilingStatusCodeLength2526 int = 1

// Field # 904
const parentFTIMAdjustedGrossIncomeStartIndex2526 int = 7328
const parentFTIMAdjustedGrossIncomeLength2526 int = 10

// Field # 905
const parentFTIMNumberOfExemptionsStartIndex2526 int = 7338
const parentFTIMNumberOfExemptionsLength2526 int = 2

// Field # 906
const parentFTIMNumberOfDependentsStartIndex2526 int = 7340
const parentFTIMNumberOfDependentsLength2526 int = 2

// Field # 907
const parentFTIMTotalIncomeEarnedAmountStartIndex2526 int = 7342
const parentFTIMTotalIncomeEarnedAmountLength2526 int = 11

// Field # 908
const parentFTIMIncomeTaxPaidStartIndex2526 int = 7353
const parentFTIMIncomeTaxPaidLength2526 int = 9

// Field # 909
const parentFTIMEducationCreditsStartIndex2526 int = 7362
const parentFTIMEducationCreditsLength2526 int = 9

// Field # 910
const parentFTIMUntaxedIRADistributionsStartIndex2526 int = 7371
const parentFTIMUntaxedIRADistributionsLength2526 int = 11

// Field # 911
const parentFTIMIRADeductibleAndPaymentsStartIndex2526 int = 7382
const parentFTIMIRADeductibleAndPaymentsLength2526 int = 11

// Field # 912
const parentFTIMTaxExemptInterestStartIndex2526 int = 7393
const parentFTIMTaxExemptInterestLength2526 int = 11

// Field # 913
const parentFTIMUntaxedPensionsAmountStartIndex2526 int = 7404
const parentFTIMUntaxedPensionsAmountLength2526 int = 11

// Field # 914
const parentFTIMScheduleCNetProfitLossStartIndex2526 int = 7415
const parentFTIMScheduleCNetProfitLossLength2526 int = 12

// Field # 915
const parentFTIMScheduleAIndicatorStartIndex2526 int = 7427
const parentFTIMScheduleAIndicatorLength2526 int = 1

// Field # 916
const parentFTIMScheduleBIndicatorStartIndex2526 int = 7428
const parentFTIMScheduleBIndicatorLength2526 int = 1

// Field # 917
const parentFTIMScheduleDIndicatorStartIndex2526 int = 7429
const parentFTIMScheduleDIndicatorLength2526 int = 1

// Field # 918
const parentFTIMScheduleEIndicatorStartIndex2526 int = 7430
const parentFTIMScheduleEIndicatorLength2526 int = 1

// Field # 919
const parentFTIMScheduleFIndicatorStartIndex2526 int = 7431
const parentFTIMScheduleFIndicatorLength2526 int = 1

// Field # 920
const parentFTIMScheduleHIndicatorStartIndex2526 int = 7432
const parentFTIMScheduleHIndicatorLength2526 int = 1

// Field # 921
const parentFTIMIRSResponseCodeStartIndex2526 int = 7433
const parentFTIMIRSResponseCodeLength2526 int = 3

// Field # 922
const parentFTIMSpouseReturnedTaxYearStartIndex2526 int = 7436
const parentFTIMSpouseReturnedTaxYearLength2526 int = 4

// Field # 923
const parentFTIMSpouseFilingStatusCodeStartIndex2526 int = 7440
const parentFTIMSpouseFilingStatusCodeLength2526 int = 1

// Field # 924
const parentFTIMSpouseAdjustedGrossIncomeStartIndex2526 int = 7441
const parentFTIMSpouseAdjustedGrossIncomeLength2526 int = 10

// Field # 925
const parentFTIMSpouseNumberOfExemptionsStartIndex2526 int = 7451
const parentFTIMSpouseNumberOfExemptionsLength2526 int = 2

// Field # 926
const parentFTIMSpouseNumberOfDependentsStartIndex2526 int = 7453
const parentFTIMSpouseNumberOfDependentsLength2526 int = 2

// Field # 927
const parentFTIMSpouseTotalIncomeEarnedAmountStartIndex2526 int = 7455
const parentFTIMSpouseTotalIncomeEarnedAmountLength2526 int = 11

// Field # 928
const parentFTIMSpouseIncomeTaxPaidStartIndex2526 int = 7466
const parentFTIMSpouseIncomeTaxPaidLength2526 int = 9

// Field # 929
const parentFTIMSpouseEducationCreditsStartIndex2526 int = 7475
const parentFTIMSpouseEducationCreditsLength2526 int = 9

// Field # 930
const parentFTIMSpouseUntaxedIRADistributionsStartIndex2526 int = 7484
const parentFTIMSpouseUntaxedIRADistributionsLength2526 int = 11

// Field # 931
const parentFTIMSpouseIRADeductibleAndPaymentsStartIndex2526 int = 7495
const parentFTIMSpouseIRADeductibleAndPaymentsLength2526 int = 11

// Field # 932
const parentFTIMSpouseTaxExemptInterestStartIndex2526 int = 7506
const parentFTIMSpouseTaxExemptInterestLength2526 int = 11

// Field # 933
const parentFTIMSpouseUntaxedPensionsAmountStartIndex2526 int = 7517
const parentFTIMSpouseUntaxedPensionsAmountLength2526 int = 11

// Field # 934
const parentFTIMSpouseScheduleCNetProfitLossStartIndex2526 int = 7528
const parentFTIMSpouseScheduleCNetProfitLossLength2526 int = 12

// Field # 935
const parentFTIMSpouseScheduleAIndicatorStartIndex2526 int = 7540
const parentFTIMSpouseScheduleAIndicatorLength2526 int = 1

// Field # 936
const parentFTIMSpouseScheduleBIndicatorStartIndex2526 int = 7541
const parentFTIMSpouseScheduleBIndicatorLength2526 int = 1

// Field # 937
const parentFTIMSpouseScheduleDIndicatorStartIndex2526 int = 7542
const parentFTIMSpouseScheduleDIndicatorLength2526 int = 1

// Field # 938
const parentFTIMSpouseScheduleEIndicatorStartIndex2526 int = 7543
const parentFTIMSpouseScheduleEIndicatorLength2526 int = 1

// Field # 939
const parentFTIMSpouseScheduleFIndicatorStartIndex2526 int = 7544
const parentFTIMSpouseScheduleFIndicatorLength2526 int = 1

// Field # 940
const parentFTIMSpouseScheduleHIndicatorStartIndex2526 int = 7545
const parentFTIMSpouseScheduleHIndicatorLength2526 int = 1

// Field # 941
const parentFTIMSpouseIRSResponseCodeStartIndex2526 int = 7546
const parentFTIMSpouseIRSResponseCodeLength2526 int = 3

// Field # 942
const ftiLabelEndStartIndex2526 int = 7549
const ftiLabelEndLength2526 int = 11

// Field # 944
const studentTotalIncomeStartIndex2526 int = 7610
const studentTotalIncomeLength2526 int = 15

// Field # 945
const parentTotalIncomeStartIndex2526 int = 7625
const parentTotalIncomeLength2526 int = 15

// Field # 946
const fisapTotalIncomeStartIndex2526 int = 7640
const fisapTotalIncomeLength2526 int = 15

type ISIRParser2526 struct {
}

func (parser *ISIRParser2526) ParseISIR(record string) (isirmodels.ISIRecord, error) {
	slog.Debug("Parsing an expected ISIR record from fixed format")
	if len(record) != totalISIRLength2526 {
		slog.Error(fmt.Sprintf("Expected ISIR to be length %d, received string with length %d", totalISIRLength2526, len(record)))
		return isirmodels.ISIRecord{}, errors.New(fmt.Sprintf("Input ISIR string is the incorrect length, expected %d and received %d", totalISIRLength2526, len(record)))
	}

	slog.Info("Parsing record", "FAFSAUUID", strings.TrimSpace(record[fafsaUUIDStartIndex2526-1:(fafsaUUIDStartIndex2526-1)+fafsaUUIDLength2526]),
		"TransactionUUID", strings.TrimSpace(record[transactionUUIDStartIndex2526-1:(transactionUUIDStartIndex2526-1)+transactionUUIDLength2526]),
		"PersonUUID", strings.TrimSpace(record[transactionUUIDStartIndex2526-1:(transactionUUIDStartIndex2526-1)+transactionUUIDLength2526]))

	r := isirmodels.ISIRecord{
		YearIndicator: strings.TrimSpace(record[yearIndicatorStartIndex2526-1 : (yearIndicatorStartIndex2526-1)+yearIndicatorLength2526]), // Field # 1

		FAFSAUUID: strings.TrimSpace(record[fafsaUUIDStartIndex2526-1 : (fafsaUUIDStartIndex2526-1)+fafsaUUIDLength2526]), // Field # 2

		TransactionUUID: strings.TrimSpace(record[transactionUUIDStartIndex2526-1 : (transactionUUIDStartIndex2526-1)+transactionUUIDLength2526]), // Field # 3

		PersonUUID: strings.TrimSpace(record[personUUIDStartIndex2526-1 : (personUUIDStartIndex2526-1)+personUUIDLength2526]), // Field # 4

		TransactionNumber: strings.TrimSpace(record[transactionNumberStartIndex2526-1 : (transactionNumberStartIndex2526-1)+transactionNumberLength2526]), // Field # 5

		DependencyModel: strings.TrimSpace(record[dependencyModelStartIndex2526-1 : (dependencyModelStartIndex2526-1)+dependencyModelLength2526]), // Field # 6

		ApplicationSource: strings.TrimSpace(record[applicationSourceStartIndex2526-1 : (applicationSourceStartIndex2526-1)+applicationSourceLength2526]), // Field # 7

		ApplicationReceiptDate: parseISIRDate2526(strings.TrimSpace(record[applicationReceiptDateStartIndex2526-1 : (applicationReceiptDateStartIndex2526-1)+applicationReceiptDateLength2526])), // Field # 8

		TransactionSource: strings.TrimSpace(record[transactionSourceStartIndex2526-1 : (transactionSourceStartIndex2526-1)+transactionSourceLength2526]), // Field # 9

		TransactionType: strings.TrimSpace(record[transactionTypeStartIndex2526-1 : (transactionTypeStartIndex2526-1)+transactionTypeLength2526]), // Field # 10

		TransactionLanguage: strings.TrimSpace(record[transactionLanguageStartIndex2526-1 : (transactionLanguageStartIndex2526-1)+transactionLanguageLength2526]), // Field # 11

		TransactionReceiptDate: parseISIRDate2526(strings.TrimSpace(record[transactionReceiptDateStartIndex2526-1 : (transactionReceiptDateStartIndex2526-1)+transactionReceiptDateLength2526])), // Field # 12

		TransactionProcessedDate: parseISIRDate2526(strings.TrimSpace(record[transactionProcessedDateStartIndex2526-1 : (transactionProcessedDateStartIndex2526-1)+transactionProcessedDateLength2526])), // Field # 13

		TransactionStatus: strings.TrimSpace(record[transactionStatusStartIndex2526-1 : (transactionStatusStartIndex2526-1)+transactionStatusLength2526]), // Field # 14

		RenewalDataUsed: strings.TrimSpace(record[renewalDataUsedStartIndex2526-1 : (renewalDataUsedStartIndex2526-1)+renewalDataUsedLength2526]), // Field # 15

		FPSCorrectionReason: strings.TrimSpace(record[fpsCorrectionReasonStartIndex2526-1 : (fpsCorrectionReasonStartIndex2526-1)+fpsCorrectionReasonLength2526]), // Field # 16

		SAIChangeFlag: strings.TrimSpace(record[saiChangeFlagStartIndex2526-1 : (saiChangeFlagStartIndex2526-1)+saiChangeFlagLength2526]), // Field # 17

		SAI: strings.TrimSpace(record[saiStartIndex2526-1 : (saiStartIndex2526-1)+saiLength2526]), // Field # 18

		ProvisionalSAI: strings.TrimSpace(record[provisionalSAIStartIndex2526-1 : (provisionalSAIStartIndex2526-1)+provisionalSAILength2526]), // Field # 19

		SAIFormula: strings.TrimSpace(record[saiFormulaStartIndex2526-1 : (saiFormulaStartIndex2526-1)+saiFormulaLength2526]), // Field # 20

		SAIComputationType: strings.TrimSpace(record[saiComputationTypeStartIndex2526-1 : (saiComputationTypeStartIndex2526-1)+saiComputationTypeLength2526]), // Field # 21

		MaxPellIndicator: strings.TrimSpace(record[maxPellIndicatorStartIndex2526-1 : (maxPellIndicatorStartIndex2526-1)+maxPellIndicatorLength2526]), // Field # 22

		MinimumPellIndicator: strings.TrimSpace(record[minimumPellIndicatorStartIndex2526-1 : (minimumPellIndicatorStartIndex2526-1)+minimumPellIndicatorLength2526]), // Field # 23

		StudentFirstName: strings.TrimSpace(record[studentFirstNameStartIndex2526-1 : (studentFirstNameStartIndex2526-1)+studentFirstNameLength2526]), // Field # 25

		StudentMiddleName: strings.TrimSpace(record[studentMiddleNameStartIndex2526-1 : (studentMiddleNameStartIndex2526-1)+studentMiddleNameLength2526]), // Field # 26

		StudentLastName: strings.TrimSpace(record[studentLastNameStartIndex2526-1 : (studentLastNameStartIndex2526-1)+studentLastNameLength2526]), // Field # 27

		StudentSuffix: strings.TrimSpace(record[studentSuffixStartIndex2526-1 : (studentSuffixStartIndex2526-1)+studentSuffixLength2526]), // Field # 28

		StudentDateOfBirth: parseISIRDate2526(strings.TrimSpace(record[studentDateOfBirthStartIndex2526-1 : (studentDateOfBirthStartIndex2526-1)+studentDateOfBirthLength2526])), // Field # 29

		StudentSSN: strings.TrimSpace(record[studentSSNStartIndex2526-1 : (studentSSNStartIndex2526-1)+studentSSNLength2526]), // Field # 30

		StudentITIN: strings.TrimSpace(record[studentITINStartIndex2526-1 : (studentITINStartIndex2526-1)+studentITINLength2526]), // Field # 31

		StudentPhoneNumber: strings.TrimSpace(record[studentPhoneNumberStartIndex2526-1 : (studentPhoneNumberStartIndex2526-1)+studentPhoneNumberLength2526]), // Field # 32

		StudentEmailAddress: strings.TrimSpace(record[studentEmailAddressStartIndex2526-1 : (studentEmailAddressStartIndex2526-1)+studentEmailAddressLength2526]), // Field # 33

		StudentStreetAddress: strings.TrimSpace(record[studentStreetAddressStartIndex2526-1 : (studentStreetAddressStartIndex2526-1)+studentStreetAddressLength2526]), // Field # 34

		StudentCity: strings.TrimSpace(record[studentCityStartIndex2526-1 : (studentCityStartIndex2526-1)+studentCityLength2526]), // Field # 35

		StudentState: strings.TrimSpace(record[studentStateStartIndex2526-1 : (studentStateStartIndex2526-1)+studentStateLength2526]), // Field # 36

		StudentZipCode: strings.TrimSpace(record[studentZipCodeStartIndex2526-1 : (studentZipCodeStartIndex2526-1)+studentZipCodeLength2526]), // Field # 37

		StudentCountry: strings.TrimSpace(record[studentCountryStartIndex2526-1 : (studentCountryStartIndex2526-1)+studentCountryLength2526]), // Field # 38

		StudentMaritalStatus: strings.TrimSpace(record[studentMaritalStatusStartIndex2526-1 : (studentMaritalStatusStartIndex2526-1)+studentMaritalStatusLength2526]), // Field # 40

		StudentGradeLevel: strings.TrimSpace(record[studentGradeLevelStartIndex2526-1 : (studentGradeLevelStartIndex2526-1)+studentGradeLevelLength2526]), // Field # 41

		StudentFirstBachelorsDegreeBefore2526: strings.TrimSpace(record[studentFirstBachelorsDegreeBefore2526StartIndex2526-1 : (studentFirstBachelorsDegreeBefore2526StartIndex2526-1)+studentFirstBachelorsDegreeBefore2526Length2526]), // Field # 42

		StudentPursuingTeacherCertification: strings.TrimSpace(record[studentPursuingTeacherCertificationStartIndex2526-1 : (studentPursuingTeacherCertificationStartIndex2526-1)+studentPursuingTeacherCertificationLength2526]), // Field # 43

		StudentActiveDuty: strings.TrimSpace(record[studentActiveDutyStartIndex2526-1 : (studentActiveDutyStartIndex2526-1)+studentActiveDutyLength2526]), // Field # 44

		StudentVeteran: strings.TrimSpace(record[studentVeteranStartIndex2526-1 : (studentVeteranStartIndex2526-1)+studentVeteranLength2526]), // Field # 45

		StudentChildOrOtherDependents: strings.TrimSpace(record[studentChildOrOtherDependentsStartIndex2526-1 : (studentChildOrOtherDependentsStartIndex2526-1)+studentChildOrOtherDependentsLength2526]), // Field # 46

		StudentParentsDeceased: strings.TrimSpace(record[studentParentsDeceasedStartIndex2526-1 : (studentParentsDeceasedStartIndex2526-1)+studentParentsDeceasedLength2526]), // Field # 47

		StudentWardOfCourt: strings.TrimSpace(record[studentWardOfCourtStartIndex2526-1 : (studentWardOfCourtStartIndex2526-1)+studentWardOfCourtLength2526]), // Field # 48

		StudentInFosterCare: strings.TrimSpace(record[studentInFosterCareStartIndex2526-1 : (studentInFosterCareStartIndex2526-1)+studentInFosterCareLength2526]), // Field # 49

		StudentEmancipatedMinor: strings.TrimSpace(record[studentEmancipatedMinorStartIndex2526-1 : (studentEmancipatedMinorStartIndex2526-1)+studentEmancipatedMinorLength2526]), // Field # 50

		StudentLegalGuardianship: strings.TrimSpace(record[studentLegalGuardianshipStartIndex2526-1 : (studentLegalGuardianshipStartIndex2526-1)+studentLegalGuardianshipLength2526]), // Field # 51

		StudentPersonalCircumstancesNoneOfTheAbove: strings.TrimSpace(record[studentPersonalCircumstancesNoneOfTheAboveStartIndex2526-1 : (studentPersonalCircumstancesNoneOfTheAboveStartIndex2526-1)+studentPersonalCircumstancesNoneOfTheAboveLength2526]), // Field # 52

		StudentUnaccompaniedHomelessYouthAndSelfSupporting: strings.TrimSpace(record[studentUnaccompaniedHomelessYouthAndSelfSupportingStartIndex2526-1 : (studentUnaccompaniedHomelessYouthAndSelfSupportingStartIndex2526-1)+studentUnaccompaniedHomelessYouthAndSelfSupportingLength2526]), // Field # 53

		StudentUnaccompaniedHomelessGeneral: strings.TrimSpace(record[studentUnaccompaniedHomelessGeneralStartIndex2526-1 : (studentUnaccompaniedHomelessGeneralStartIndex2526-1)+studentUnaccompaniedHomelessGeneralLength2526]), // Field # 54

		StudentUnaccompaniedHomelessHS: strings.TrimSpace(record[studentUnaccompaniedHomelessHSStartIndex2526-1 : (studentUnaccompaniedHomelessHSStartIndex2526-1)+studentUnaccompaniedHomelessHSLength2526]), // Field # 55

		StudentUnaccompaniedHomelessTRIO: strings.TrimSpace(record[studentUnaccompaniedHomelessTRIOStartIndex2526-1 : (studentUnaccompaniedHomelessTRIOStartIndex2526-1)+studentUnaccompaniedHomelessTRIOLength2526]), // Field # 56

		StudentUnaccompaniedHomelessFAA: strings.TrimSpace(record[studentUnaccompaniedHomelessFAAStartIndex2526-1 : (studentUnaccompaniedHomelessFAAStartIndex2526-1)+studentUnaccompaniedHomelessFAALength2526]), // Field # 57

		StudentHomelessnessNoneOfTheAbove: strings.TrimSpace(record[studentHomelessnessNoneOfTheAboveStartIndex2526-1 : (studentHomelessnessNoneOfTheAboveStartIndex2526-1)+studentHomelessnessNoneOfTheAboveLength2526]), // Field # 58

		StudentUnusualCircumstance: strings.TrimSpace(record[studentUnusualCircumstanceStartIndex2526-1 : (studentUnusualCircumstanceStartIndex2526-1)+studentUnusualCircumstanceLength2526]), // Field # 59

		StudentUnsubOnly: strings.TrimSpace(record[studentUnsubOnlyStartIndex2526-1 : (studentUnsubOnlyStartIndex2526-1)+studentUnsubOnlyLength2526]), // Field # 60

		StudentUpdatedFamilySize: strings.TrimSpace(record[studentUpdatedFamilySizeStartIndex2526-1 : (studentUpdatedFamilySizeStartIndex2526-1)+studentUpdatedFamilySizeLength2526]), // Field # 61

		StudentNumberInCollege: strings.TrimSpace(record[studentNumberInCollegeStartIndex2526-1 : (studentNumberInCollegeStartIndex2526-1)+studentNumberInCollegeLength2526]), // Field # 62

		StudentCitizenshipStatus: strings.TrimSpace(record[studentCitizenshipStatusStartIndex2526-1 : (studentCitizenshipStatusStartIndex2526-1)+studentCitizenshipStatusLength2526]), // Field # 63

		StudentANumber: strings.TrimSpace(record[studentANumberStartIndex2526-1 : (studentANumberStartIndex2526-1)+studentANumberLength2526]), // Field # 64

		StudentStateOfLegalResidence: strings.TrimSpace(record[studentStateOfLegalResidenceStartIndex2526-1 : (studentStateOfLegalResidenceStartIndex2526-1)+studentStateOfLegalResidenceLength2526]), // Field # 65

		StudentLegalResidenceDate: parseISIRDateShort2526(strings.TrimSpace(record[studentLegalResidenceDateStartIndex2526-1 : (studentLegalResidenceDateStartIndex2526-1)+studentLegalResidenceDateLength2526])), // Field # 66

		StudentEitherParentAttendCollege: strings.TrimSpace(record[studentEitherParentAttendCollegeStartIndex2526-1 : (studentEitherParentAttendCollegeStartIndex2526-1)+studentEitherParentAttendCollegeLength2526]), // Field # 67

		StudentParentKilledInTheLineOfDuty: strings.TrimSpace(record[studentParentKilledInTheLineOfDutyStartIndex2526-1 : (studentParentKilledInTheLineOfDutyStartIndex2526-1)+studentParentKilledInTheLineOfDutyLength2526]), // Field # 68

		StudentHighSchoolCompletionStatus: strings.TrimSpace(record[studentHighSchoolCompletionStatusStartIndex2526-1 : (studentHighSchoolCompletionStatusStartIndex2526-1)+studentHighSchoolCompletionStatusLength2526]), // Field # 69

		StudentHighSchoolName: strings.TrimSpace(record[studentHighSchoolNameStartIndex2526-1 : (studentHighSchoolNameStartIndex2526-1)+studentHighSchoolNameLength2526]), // Field # 70

		StudentHighSchoolCity: strings.TrimSpace(record[studentHighSchoolCityStartIndex2526-1 : (studentHighSchoolCityStartIndex2526-1)+studentHighSchoolCityLength2526]), // Field # 71

		StudentHighSchoolState: strings.TrimSpace(record[studentHighSchoolStateStartIndex2526-1 : (studentHighSchoolStateStartIndex2526-1)+studentHighSchoolStateLength2526]), // Field # 72

		StudentHighSchoolEquivalentDiplomaName: strings.TrimSpace(record[studentHighSchoolEquivalentDiplomaNameStartIndex2526-1 : (studentHighSchoolEquivalentDiplomaNameStartIndex2526-1)+studentHighSchoolEquivalentDiplomaNameLength2526]), // Field # 73

		StudentHighSchoolEquivalentDiplomaState: strings.TrimSpace(record[studentHighSchoolEquivalentDiplomaStateStartIndex2526-1 : (studentHighSchoolEquivalentDiplomaStateStartIndex2526-1)+studentHighSchoolEquivalentDiplomaStateLength2526]), // Field # 74

		StudentManuallyEnteredReceivedEITC: strings.TrimSpace(record[studentManuallyEnteredReceivedEITCStartIndex2526-1 : (studentManuallyEnteredReceivedEITCStartIndex2526-1)+studentManuallyEnteredReceivedEITCLength2526]), // Field # 75

		StudentManuallyEnteredReceivedFederalHousingAssistance: strings.TrimSpace(record[studentManuallyEnteredReceivedFederalHousingAssistanceStartIndex2526-1 : (studentManuallyEnteredReceivedFederalHousingAssistanceStartIndex2526-1)+studentManuallyEnteredReceivedFederalHousingAssistanceLength2526]), // Field # 76

		StudentManuallyEnteredReceivedFreeReducedPriceLunch: strings.TrimSpace(record[studentManuallyEnteredReceivedFreeReducedPriceLunchStartIndex2526-1 : (studentManuallyEnteredReceivedFreeReducedPriceLunchStartIndex2526-1)+studentManuallyEnteredReceivedFreeReducedPriceLunchLength2526]), // Field # 77

		StudentManuallyEnteredReceivedMedicaid: strings.TrimSpace(record[studentManuallyEnteredReceivedMedicaidStartIndex2526-1 : (studentManuallyEnteredReceivedMedicaidStartIndex2526-1)+studentManuallyEnteredReceivedMedicaidLength2526]), // Field # 78

		StudentManuallyEnteredReceivedRefundableCreditFor36BHealthPlan: strings.TrimSpace(record[studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanStartIndex2526-1 : (studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanStartIndex2526-1)+studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanLength2526]), // Field # 79

		StudentManuallyEnteredReceivedSNAP: strings.TrimSpace(record[studentManuallyEnteredReceivedSNAPStartIndex2526-1 : (studentManuallyEnteredReceivedSNAPStartIndex2526-1)+studentManuallyEnteredReceivedSNAPLength2526]), // Field # 80

		StudentManuallyEnteredReceivedSupplementalSecurityIncome: strings.TrimSpace(record[studentManuallyEnteredReceivedSupplementalSecurityIncomeStartIndex2526-1 : (studentManuallyEnteredReceivedSupplementalSecurityIncomeStartIndex2526-1)+studentManuallyEnteredReceivedSupplementalSecurityIncomeLength2526]), // Field # 81

		StudentManuallyEnteredReceivedTANF: strings.TrimSpace(record[studentManuallyEnteredReceivedTANFStartIndex2526-1 : (studentManuallyEnteredReceivedTANFStartIndex2526-1)+studentManuallyEnteredReceivedTANFLength2526]), // Field # 82

		StudentManuallyEnteredReceivedWIC: strings.TrimSpace(record[studentManuallyEnteredReceivedWICStartIndex2526-1 : (studentManuallyEnteredReceivedWICStartIndex2526-1)+studentManuallyEnteredReceivedWICLength2526]), // Field # 83

		StudentManuallyEnteredFederalBenefitsNoneOfTheAbove: strings.TrimSpace(record[studentManuallyEnteredFederalBenefitsNoneOfTheAboveStartIndex2526-1 : (studentManuallyEnteredFederalBenefitsNoneOfTheAboveStartIndex2526-1)+studentManuallyEnteredFederalBenefitsNoneOfTheAboveLength2526]), // Field # 84

		StudentManuallyEnteredFiled1040Or1040NR: strings.TrimSpace(record[studentManuallyEnteredFiled1040Or1040NRStartIndex2526-1 : (studentManuallyEnteredFiled1040Or1040NRStartIndex2526-1)+studentManuallyEnteredFiled1040Or1040NRLength2526]), // Field # 85

		StudentManuallyEnteredFiledNonUSTaxReturn: strings.TrimSpace(record[studentManuallyEnteredFiledNonUSTaxReturnStartIndex2526-1 : (studentManuallyEnteredFiledNonUSTaxReturnStartIndex2526-1)+studentManuallyEnteredFiledNonUSTaxReturnLength2526]), // Field # 86

		StudentManuallyEnteredFiledJointReturnWithCurrentSpouse: strings.TrimSpace(record[studentManuallyEnteredFiledJointReturnWithCurrentSpouseStartIndex2526-1 : (studentManuallyEnteredFiledJointReturnWithCurrentSpouseStartIndex2526-1)+studentManuallyEnteredFiledJointReturnWithCurrentSpouseLength2526]), // Field # 87

		StudentManuallyEnteredTaxReturnFilingStatus: strings.TrimSpace(record[studentManuallyEnteredTaxReturnFilingStatusStartIndex2526-1 : (studentManuallyEnteredTaxReturnFilingStatusStartIndex2526-1)+studentManuallyEnteredTaxReturnFilingStatusLength2526]), // Field # 88

		StudentManuallyEnteredIncomeEarnedFromWork: strings.TrimSpace(record[studentManuallyEnteredIncomeEarnedFromWorkStartIndex2526-1 : (studentManuallyEnteredIncomeEarnedFromWorkStartIndex2526-1)+studentManuallyEnteredIncomeEarnedFromWorkLength2526]), // Field # 89

		StudentManuallyEnteredTaxExemptInterestIncome: strings.TrimSpace(record[studentManuallyEnteredTaxExemptInterestIncomeStartIndex2526-1 : (studentManuallyEnteredTaxExemptInterestIncomeStartIndex2526-1)+studentManuallyEnteredTaxExemptInterestIncomeLength2526]), // Field # 90

		StudentManuallyEnteredUntaxedPortionsOfIRADistributions: strings.TrimSpace(record[studentManuallyEnteredUntaxedPortionsOfIRADistributionsStartIndex2526-1 : (studentManuallyEnteredUntaxedPortionsOfIRADistributionsStartIndex2526-1)+studentManuallyEnteredUntaxedPortionsOfIRADistributionsLength2526]), // Field # 91

		StudentManuallyEnteredIRARollover: strings.TrimSpace(record[studentManuallyEnteredIRARolloverStartIndex2526-1 : (studentManuallyEnteredIRARolloverStartIndex2526-1)+studentManuallyEnteredIRARolloverLength2526]), // Field # 92

		StudentManuallyEnteredUntaxedPortionsOfPensions: strings.TrimSpace(record[studentManuallyEnteredUntaxedPortionsOfPensionsStartIndex2526-1 : (studentManuallyEnteredUntaxedPortionsOfPensionsStartIndex2526-1)+studentManuallyEnteredUntaxedPortionsOfPensionsLength2526]), // Field # 93

		StudentManuallyEnteredPensionRollover: strings.TrimSpace(record[studentManuallyEnteredPensionRolloverStartIndex2526-1 : (studentManuallyEnteredPensionRolloverStartIndex2526-1)+studentManuallyEnteredPensionRolloverLength2526]), // Field # 94

		StudentManuallyEnteredAdjustedGrossIncome: strings.TrimSpace(record[studentManuallyEnteredAdjustedGrossIncomeStartIndex2526-1 : (studentManuallyEnteredAdjustedGrossIncomeStartIndex2526-1)+studentManuallyEnteredAdjustedGrossIncomeLength2526]), // Field # 95

		StudentManuallyEnteredIncomeTaxPaid: strings.TrimSpace(record[studentManuallyEnteredIncomeTaxPaidStartIndex2526-1 : (studentManuallyEnteredIncomeTaxPaidStartIndex2526-1)+studentManuallyEnteredIncomeTaxPaidLength2526]), // Field # 96

		StudentManuallyEnteredEarnedIncomeTaxCreditReceivedDuringTaxYear: strings.TrimSpace(record[studentManuallyEnteredEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex2526-1 : (studentManuallyEnteredEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex2526-1)+studentManuallyEnteredEarnedIncomeTaxCreditReceivedDuringTaxYearLength2526]), // Field # 97

		StudentManuallyEnteredDeductiblePaymentsToIRAKeoghOther: strings.TrimSpace(record[studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherStartIndex2526-1 : (studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherStartIndex2526-1)+studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherLength2526]), // Field # 98

		StudentManuallyEnteredEducationCredits: strings.TrimSpace(record[studentManuallyEnteredEducationCreditsStartIndex2526-1 : (studentManuallyEnteredEducationCreditsStartIndex2526-1)+studentManuallyEnteredEducationCreditsLength2526]), // Field # 99

		StudentManuallyEnteredFiledScheduleABDEFH: strings.TrimSpace(record[studentManuallyEnteredFiledScheduleABDEFHStartIndex2526-1 : (studentManuallyEnteredFiledScheduleABDEFHStartIndex2526-1)+studentManuallyEnteredFiledScheduleABDEFHLength2526]), // Field # 100

		StudentManuallyEnteredScheduleCAmount: strings.TrimSpace(record[studentManuallyEnteredScheduleCAmountStartIndex2526-1 : (studentManuallyEnteredScheduleCAmountStartIndex2526-1)+studentManuallyEnteredScheduleCAmountLength2526]), // Field # 101

		StudentManuallyEnteredCollegeGrantAndScholarshipAid: strings.TrimSpace(record[studentManuallyEnteredCollegeGrantAndScholarshipAidStartIndex2526-1 : (studentManuallyEnteredCollegeGrantAndScholarshipAidStartIndex2526-1)+studentManuallyEnteredCollegeGrantAndScholarshipAidLength2526]), // Field # 102

		StudentManuallyEnteredForeignEarnedIncomeExclusion: strings.TrimSpace(record[studentManuallyEnteredForeignEarnedIncomeExclusionStartIndex2526-1 : (studentManuallyEnteredForeignEarnedIncomeExclusionStartIndex2526-1)+studentManuallyEnteredForeignEarnedIncomeExclusionLength2526]), // Field # 103

		StudentManuallyEnteredChildSupportReceived: strings.TrimSpace(record[studentManuallyEnteredChildSupportReceivedStartIndex2526-1 : (studentManuallyEnteredChildSupportReceivedStartIndex2526-1)+studentManuallyEnteredChildSupportReceivedLength2526]), // Field # 104

		StudentManuallyEnteredTotalOfCashSavingsAndCheckingAccounts: strings.TrimSpace(record[studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsStartIndex2526-1 : (studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsStartIndex2526-1)+studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsLength2526]), // Field # 105

		StudentManuallyEnteredNetWorthOfCurrentInvestments: strings.TrimSpace(record[studentManuallyEnteredNetWorthOfCurrentInvestmentsStartIndex2526-1 : (studentManuallyEnteredNetWorthOfCurrentInvestmentsStartIndex2526-1)+studentManuallyEnteredNetWorthOfCurrentInvestmentsLength2526]), // Field # 106

		StudentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarms: strings.TrimSpace(record[studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsStartIndex2526-1 : (studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsStartIndex2526-1)+studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsLength2526]), // Field # 107

		StudentCollege1: strings.TrimSpace(record[studentCollege1StartIndex2526-1 : (studentCollege1StartIndex2526-1)+studentCollege1Length2526]), // Field # 108

		StudentCollege2: strings.TrimSpace(record[studentCollege2StartIndex2526-1 : (studentCollege2StartIndex2526-1)+studentCollege2Length2526]), // Field # 109

		StudentCollege3: strings.TrimSpace(record[studentCollege3StartIndex2526-1 : (studentCollege3StartIndex2526-1)+studentCollege3Length2526]), // Field # 110

		StudentCollege4: strings.TrimSpace(record[studentCollege4StartIndex2526-1 : (studentCollege4StartIndex2526-1)+studentCollege4Length2526]), // Field # 111

		StudentCollege5: strings.TrimSpace(record[studentCollege5StartIndex2526-1 : (studentCollege5StartIndex2526-1)+studentCollege5Length2526]), // Field # 112

		StudentCollege6: strings.TrimSpace(record[studentCollege6StartIndex2526-1 : (studentCollege6StartIndex2526-1)+studentCollege6Length2526]), // Field # 113

		StudentCollege7: strings.TrimSpace(record[studentCollege7StartIndex2526-1 : (studentCollege7StartIndex2526-1)+studentCollege7Length2526]), // Field # 114

		StudentCollege8: strings.TrimSpace(record[studentCollege8StartIndex2526-1 : (studentCollege8StartIndex2526-1)+studentCollege8Length2526]), // Field # 115

		StudentCollege9: strings.TrimSpace(record[studentCollege9StartIndex2526-1 : (studentCollege9StartIndex2526-1)+studentCollege9Length2526]), // Field # 116

		StudentCollege10: strings.TrimSpace(record[studentCollege10StartIndex2526-1 : (studentCollege10StartIndex2526-1)+studentCollege10Length2526]), // Field # 117

		StudentCollege11: strings.TrimSpace(record[studentCollege11StartIndex2526-1 : (studentCollege11StartIndex2526-1)+studentCollege11Length2526]), // Field # 118

		StudentCollege12: strings.TrimSpace(record[studentCollege12StartIndex2526-1 : (studentCollege12StartIndex2526-1)+studentCollege12Length2526]), // Field # 119

		StudentCollege13: strings.TrimSpace(record[studentCollege13StartIndex2526-1 : (studentCollege13StartIndex2526-1)+studentCollege13Length2526]), // Field # 120

		StudentCollege14: strings.TrimSpace(record[studentCollege14StartIndex2526-1 : (studentCollege14StartIndex2526-1)+studentCollege14Length2526]), // Field # 121

		StudentCollege15: strings.TrimSpace(record[studentCollege15StartIndex2526-1 : (studentCollege15StartIndex2526-1)+studentCollege15Length2526]), // Field # 122

		StudentCollege16: strings.TrimSpace(record[studentCollege16StartIndex2526-1 : (studentCollege16StartIndex2526-1)+studentCollege16Length2526]), // Field # 123

		StudentCollege17: strings.TrimSpace(record[studentCollege17StartIndex2526-1 : (studentCollege17StartIndex2526-1)+studentCollege17Length2526]), // Field # 124

		StudentCollege18: strings.TrimSpace(record[studentCollege18StartIndex2526-1 : (studentCollege18StartIndex2526-1)+studentCollege18Length2526]), // Field # 125

		StudentCollege19: strings.TrimSpace(record[studentCollege19StartIndex2526-1 : (studentCollege19StartIndex2526-1)+studentCollege19Length2526]), // Field # 126

		StudentCollege20: strings.TrimSpace(record[studentCollege20StartIndex2526-1 : (studentCollege20StartIndex2526-1)+studentCollege20Length2526]), // Field # 127

		StudentConsentToRetrieveAndDiscloseFTI: strings.TrimSpace(record[studentConsentToRetrieveAndDiscloseFTIStartIndex2526-1 : (studentConsentToRetrieveAndDiscloseFTIStartIndex2526-1)+studentConsentToRetrieveAndDiscloseFTILength2526]), // Field # 128

		StudentSignature: strings.TrimSpace(record[studentSignatureStartIndex2526-1 : (studentSignatureStartIndex2526-1)+studentSignatureLength2526]), // Field # 129

		StudentSignatureDate: parseISIRDate2526(strings.TrimSpace(record[studentSignatureDateStartIndex2526-1 : (studentSignatureDateStartIndex2526-1)+studentSignatureDateLength2526])), // Field # 130

		StudentSpouseFirstName: strings.TrimSpace(record[studentSpouseFirstNameStartIndex2526-1 : (studentSpouseFirstNameStartIndex2526-1)+studentSpouseFirstNameLength2526]), // Field # 132

		StudentSpouseMiddleName: strings.TrimSpace(record[studentSpouseMiddleNameStartIndex2526-1 : (studentSpouseMiddleNameStartIndex2526-1)+studentSpouseMiddleNameLength2526]), // Field # 133

		StudentSpouseLastName: strings.TrimSpace(record[studentSpouseLastNameStartIndex2526-1 : (studentSpouseLastNameStartIndex2526-1)+studentSpouseLastNameLength2526]), // Field # 134

		StudentSpouseSuffix: strings.TrimSpace(record[studentSpouseSuffixStartIndex2526-1 : (studentSpouseSuffixStartIndex2526-1)+studentSpouseSuffixLength2526]), // Field # 135

		StudentSpouseDateOfBirth: parseISIRDate2526(strings.TrimSpace(record[studentSpouseDateOfBirthStartIndex2526-1 : (studentSpouseDateOfBirthStartIndex2526-1)+studentSpouseDateOfBirthLength2526])), // Field # 136

		StudentSpouseSSN: strings.TrimSpace(record[studentSpouseSSNStartIndex2526-1 : (studentSpouseSSNStartIndex2526-1)+studentSpouseSSNLength2526]), // Field # 137

		StudentSpouseITIN: strings.TrimSpace(record[studentSpouseITINStartIndex2526-1 : (studentSpouseITINStartIndex2526-1)+studentSpouseITINLength2526]), // Field # 138

		StudentSpousePhoneNumber: strings.TrimSpace(record[studentSpousePhoneNumberStartIndex2526-1 : (studentSpousePhoneNumberStartIndex2526-1)+studentSpousePhoneNumberLength2526]), // Field # 139

		StudentSpouseEmailAddress: strings.TrimSpace(record[studentSpouseEmailAddressStartIndex2526-1 : (studentSpouseEmailAddressStartIndex2526-1)+studentSpouseEmailAddressLength2526]), // Field # 140

		StudentSpouseStreetAddress: strings.TrimSpace(record[studentSpouseStreetAddressStartIndex2526-1 : (studentSpouseStreetAddressStartIndex2526-1)+studentSpouseStreetAddressLength2526]), // Field # 141

		StudentSpouseCity: strings.TrimSpace(record[studentSpouseCityStartIndex2526-1 : (studentSpouseCityStartIndex2526-1)+studentSpouseCityLength2526]), // Field # 142

		StudentSpouseState: strings.TrimSpace(record[studentSpouseStateStartIndex2526-1 : (studentSpouseStateStartIndex2526-1)+studentSpouseStateLength2526]), // Field # 143

		StudentSpouseZipCode: strings.TrimSpace(record[studentSpouseZipCodeStartIndex2526-1 : (studentSpouseZipCodeStartIndex2526-1)+studentSpouseZipCodeLength2526]), // Field # 144

		StudentSpouseCountry: strings.TrimSpace(record[studentSpouseCountryStartIndex2526-1 : (studentSpouseCountryStartIndex2526-1)+studentSpouseCountryLength2526]), // Field # 145

		StudentSpouseFiled1040Or1040NR: strings.TrimSpace(record[studentSpouseFiled1040Or1040NRStartIndex2526-1 : (studentSpouseFiled1040Or1040NRStartIndex2526-1)+studentSpouseFiled1040Or1040NRLength2526]), // Field # 146

		StudentSpouseFiledNonUSTaxReturn: strings.TrimSpace(record[studentSpouseFiledNonUSTaxReturnStartIndex2526-1 : (studentSpouseFiledNonUSTaxReturnStartIndex2526-1)+studentSpouseFiledNonUSTaxReturnLength2526]), // Field # 147

		StudentSpouseTaxReturnFilingStatus: strings.TrimSpace(record[studentSpouseTaxReturnFilingStatusStartIndex2526-1 : (studentSpouseTaxReturnFilingStatusStartIndex2526-1)+studentSpouseTaxReturnFilingStatusLength2526]), // Field # 148

		StudentSpouseIncomeEarnedFromWork: strings.TrimSpace(record[studentSpouseIncomeEarnedFromWorkStartIndex2526-1 : (studentSpouseIncomeEarnedFromWorkStartIndex2526-1)+studentSpouseIncomeEarnedFromWorkLength2526]), // Field # 149

		StudentSpouseTaxExemptInterestIncome: strings.TrimSpace(record[studentSpouseTaxExemptInterestIncomeStartIndex2526-1 : (studentSpouseTaxExemptInterestIncomeStartIndex2526-1)+studentSpouseTaxExemptInterestIncomeLength2526]), // Field # 150

		StudentSpouseUntaxedPortionsOfIRADistributions: strings.TrimSpace(record[studentSpouseUntaxedPortionsOfIRADistributionsStartIndex2526-1 : (studentSpouseUntaxedPortionsOfIRADistributionsStartIndex2526-1)+studentSpouseUntaxedPortionsOfIRADistributionsLength2526]), // Field # 151

		StudentSpouseIRARollover: strings.TrimSpace(record[studentSpouseIRARolloverStartIndex2526-1 : (studentSpouseIRARolloverStartIndex2526-1)+studentSpouseIRARolloverLength2526]), // Field # 152

		StudentSpouseUntaxedPortionsOfPensions: strings.TrimSpace(record[studentSpouseUntaxedPortionsOfPensionsStartIndex2526-1 : (studentSpouseUntaxedPortionsOfPensionsStartIndex2526-1)+studentSpouseUntaxedPortionsOfPensionsLength2526]), // Field # 153

		StudentSpousePensionRollover: strings.TrimSpace(record[studentSpousePensionRolloverStartIndex2526-1 : (studentSpousePensionRolloverStartIndex2526-1)+studentSpousePensionRolloverLength2526]), // Field # 154

		StudentSpouseAdjustedGrossIncome: strings.TrimSpace(record[studentSpouseAdjustedGrossIncomeStartIndex2526-1 : (studentSpouseAdjustedGrossIncomeStartIndex2526-1)+studentSpouseAdjustedGrossIncomeLength2526]), // Field # 155

		StudentSpouseIncomeTaxPaid: strings.TrimSpace(record[studentSpouseIncomeTaxPaidStartIndex2526-1 : (studentSpouseIncomeTaxPaidStartIndex2526-1)+studentSpouseIncomeTaxPaidLength2526]), // Field # 156

		StudentSpouseDeductiblePaymentsToIRAKeoghOther: strings.TrimSpace(record[studentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex2526-1 : (studentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex2526-1)+studentSpouseDeductiblePaymentsToIRAKeoghOtherLength2526]), // Field # 157

		StudentSpouseEducationCredits: strings.TrimSpace(record[studentSpouseEducationCreditsStartIndex2526-1 : (studentSpouseEducationCreditsStartIndex2526-1)+studentSpouseEducationCreditsLength2526]), // Field # 158

		StudentSpouseFiledScheduleABDEFH: strings.TrimSpace(record[studentSpouseFiledScheduleABDEFHStartIndex2526-1 : (studentSpouseFiledScheduleABDEFHStartIndex2526-1)+studentSpouseFiledScheduleABDEFHLength2526]), // Field # 159

		StudentSpouseScheduleCAmount: strings.TrimSpace(record[studentSpouseScheduleCAmountStartIndex2526-1 : (studentSpouseScheduleCAmountStartIndex2526-1)+studentSpouseScheduleCAmountLength2526]), // Field # 160

		StudentSpouseForeignEarnedIncomeExclusion: strings.TrimSpace(record[studentSpouseForeignEarnedIncomeExclusionStartIndex2526-1 : (studentSpouseForeignEarnedIncomeExclusionStartIndex2526-1)+studentSpouseForeignEarnedIncomeExclusionLength2526]), // Field # 161

		StudentSpouseConsentToRetrieveAndDiscloseFTI: strings.TrimSpace(record[studentSpouseConsentToRetrieveAndDiscloseFTIStartIndex2526-1 : (studentSpouseConsentToRetrieveAndDiscloseFTIStartIndex2526-1)+studentSpouseConsentToRetrieveAndDiscloseFTILength2526]), // Field # 162

		StudentSpouseSignature: strings.TrimSpace(record[studentSpouseSignatureStartIndex2526-1 : (studentSpouseSignatureStartIndex2526-1)+studentSpouseSignatureLength2526]), // Field # 163

		StudentSpouseSignatureDate: parseISIRDate2526(strings.TrimSpace(record[studentSpouseSignatureDateStartIndex2526-1 : (studentSpouseSignatureDateStartIndex2526-1)+studentSpouseSignatureDateLength2526])), // Field # 164

		ParentFirstName: strings.TrimSpace(record[parentFirstNameStartIndex2526-1 : (parentFirstNameStartIndex2526-1)+parentFirstNameLength2526]), // Field # 166

		ParentMiddleName: strings.TrimSpace(record[parentMiddleNameStartIndex2526-1 : (parentMiddleNameStartIndex2526-1)+parentMiddleNameLength2526]), // Field # 167

		ParentLastName: strings.TrimSpace(record[parentLastNameStartIndex2526-1 : (parentLastNameStartIndex2526-1)+parentLastNameLength2526]), // Field # 168

		ParentSuffix: strings.TrimSpace(record[parentSuffixStartIndex2526-1 : (parentSuffixStartIndex2526-1)+parentSuffixLength2526]), // Field # 169

		ParentDateOfBirth: parseISIRDate2526(strings.TrimSpace(record[parentDateOfBirthStartIndex2526-1 : (parentDateOfBirthStartIndex2526-1)+parentDateOfBirthLength2526])), // Field # 170

		ParentSSN: strings.TrimSpace(record[parentSSNStartIndex2526-1 : (parentSSNStartIndex2526-1)+parentSSNLength2526]), // Field # 171

		ParentITIN: strings.TrimSpace(record[parentITINStartIndex2526-1 : (parentITINStartIndex2526-1)+parentITINLength2526]), // Field # 172

		ParentPhoneNumber: strings.TrimSpace(record[parentPhoneNumberStartIndex2526-1 : (parentPhoneNumberStartIndex2526-1)+parentPhoneNumberLength2526]), // Field # 173

		ParentEmailAddress: strings.TrimSpace(record[parentEmailAddressStartIndex2526-1 : (parentEmailAddressStartIndex2526-1)+parentEmailAddressLength2526]), // Field # 174

		ParentStreetAddress: strings.TrimSpace(record[parentStreetAddressStartIndex2526-1 : (parentStreetAddressStartIndex2526-1)+parentStreetAddressLength2526]), // Field # 175

		ParentCity: strings.TrimSpace(record[parentCityStartIndex2526-1 : (parentCityStartIndex2526-1)+parentCityLength2526]), // Field # 176

		ParentState: strings.TrimSpace(record[parentStateStartIndex2526-1 : (parentStateStartIndex2526-1)+parentStateLength2526]), // Field # 177

		ParentZipCode: strings.TrimSpace(record[parentZipCodeStartIndex2526-1 : (parentZipCodeStartIndex2526-1)+parentZipCodeLength2526]), // Field # 178

		ParentCountry: strings.TrimSpace(record[parentCountryStartIndex2526-1 : (parentCountryStartIndex2526-1)+parentCountryLength2526]), // Field # 179

		ParentMaritalStatus: strings.TrimSpace(record[parentMaritalStatusStartIndex2526-1 : (parentMaritalStatusStartIndex2526-1)+parentMaritalStatusLength2526]), // Field # 180

		ParentStateOfLegalResidence: strings.TrimSpace(record[parentStateOfLegalResidenceStartIndex2526-1 : (parentStateOfLegalResidenceStartIndex2526-1)+parentStateOfLegalResidenceLength2526]), // Field # 181

		ParentLegalResidenceDate: parseISIRDateShort2526(strings.TrimSpace(record[parentLegalResidenceDateStartIndex2526-1 : (parentLegalResidenceDateStartIndex2526-1)+parentLegalResidenceDateLength2526])), // Field # 182

		ParentUpdatedFamilySize: strings.TrimSpace(record[parentUpdatedFamilySizeStartIndex2526-1 : (parentUpdatedFamilySizeStartIndex2526-1)+parentUpdatedFamilySizeLength2526]), // Field # 183

		ParentNumberInCollege: strings.TrimSpace(record[parentNumberInCollegeStartIndex2526-1 : (parentNumberInCollegeStartIndex2526-1)+parentNumberInCollegeLength2526]), // Field # 184

		ParentReceivedEITC: strings.TrimSpace(record[parentReceivedEITCStartIndex2526-1 : (parentReceivedEITCStartIndex2526-1)+parentReceivedEITCLength2526]), // Field # 185

		ParentReceivedFederalHousingAssistance: strings.TrimSpace(record[parentReceivedFederalHousingAssistanceStartIndex2526-1 : (parentReceivedFederalHousingAssistanceStartIndex2526-1)+parentReceivedFederalHousingAssistanceLength2526]), // Field # 186

		ParentReceivedFreeReducedPriceLunch: strings.TrimSpace(record[parentReceivedFreeReducedPriceLunchStartIndex2526-1 : (parentReceivedFreeReducedPriceLunchStartIndex2526-1)+parentReceivedFreeReducedPriceLunchLength2526]), // Field # 187

		ParentReceivedMedicaid: strings.TrimSpace(record[parentReceivedMedicaidStartIndex2526-1 : (parentReceivedMedicaidStartIndex2526-1)+parentReceivedMedicaidLength2526]), // Field # 188

		ParentReceivedRefundableCreditFor36BHealthPlan: strings.TrimSpace(record[parentReceivedRefundableCreditFor36BHealthPlanStartIndex2526-1 : (parentReceivedRefundableCreditFor36BHealthPlanStartIndex2526-1)+parentReceivedRefundableCreditFor36BHealthPlanLength2526]), // Field # 189

		ParentReceivedSNAP: strings.TrimSpace(record[parentReceivedSNAPStartIndex2526-1 : (parentReceivedSNAPStartIndex2526-1)+parentReceivedSNAPLength2526]), // Field # 190

		ParentReceivedSupplementalSecurityIncome: strings.TrimSpace(record[parentReceivedSupplementalSecurityIncomeStartIndex2526-1 : (parentReceivedSupplementalSecurityIncomeStartIndex2526-1)+parentReceivedSupplementalSecurityIncomeLength2526]), // Field # 191

		ParentReceivedTANF: strings.TrimSpace(record[parentReceivedTANFStartIndex2526-1 : (parentReceivedTANFStartIndex2526-1)+parentReceivedTANFLength2526]), // Field # 192

		ParentReceivedWIC: strings.TrimSpace(record[parentReceivedWICStartIndex2526-1 : (parentReceivedWICStartIndex2526-1)+parentReceivedWICLength2526]), // Field # 193

		ParentFederalBenefitsNoneOfTheAbove: strings.TrimSpace(record[parentFederalBenefitsNoneOfTheAboveStartIndex2526-1 : (parentFederalBenefitsNoneOfTheAboveStartIndex2526-1)+parentFederalBenefitsNoneOfTheAboveLength2526]), // Field # 194

		ParentFiled1040Or1040NR: strings.TrimSpace(record[parentFiled1040Or1040NRStartIndex2526-1 : (parentFiled1040Or1040NRStartIndex2526-1)+parentFiled1040Or1040NRLength2526]), // Field # 195

		ParentFileNonUSTaxReturn: strings.TrimSpace(record[parentFileNonUSTaxReturnStartIndex2526-1 : (parentFileNonUSTaxReturnStartIndex2526-1)+parentFileNonUSTaxReturnLength2526]), // Field # 196

		ParentFiledJointReturnWithCurrentSpouse: strings.TrimSpace(record[parentFiledJointReturnWithCurrentSpouseStartIndex2526-1 : (parentFiledJointReturnWithCurrentSpouseStartIndex2526-1)+parentFiledJointReturnWithCurrentSpouseLength2526]), // Field # 197

		ParentTaxReturnFilingStatus: strings.TrimSpace(record[parentTaxReturnFilingStatusStartIndex2526-1 : (parentTaxReturnFilingStatusStartIndex2526-1)+parentTaxReturnFilingStatusLength2526]), // Field # 198

		ParentIncomeEarnedFromWork: strings.TrimSpace(record[parentIncomeEarnedFromWorkStartIndex2526-1 : (parentIncomeEarnedFromWorkStartIndex2526-1)+parentIncomeEarnedFromWorkLength2526]), // Field # 199

		ParentTaxExemptInterestIncome: strings.TrimSpace(record[parentTaxExemptInterestIncomeStartIndex2526-1 : (parentTaxExemptInterestIncomeStartIndex2526-1)+parentTaxExemptInterestIncomeLength2526]), // Field # 200

		ParentUntaxedPortionsOfIRADistributions: strings.TrimSpace(record[parentUntaxedPortionsOfIRADistributionsStartIndex2526-1 : (parentUntaxedPortionsOfIRADistributionsStartIndex2526-1)+parentUntaxedPortionsOfIRADistributionsLength2526]), // Field # 201

		ParentIRARollover: strings.TrimSpace(record[parentIRARolloverStartIndex2526-1 : (parentIRARolloverStartIndex2526-1)+parentIRARolloverLength2526]), // Field # 202

		ParentUntaxedPortionsOfPensions: strings.TrimSpace(record[parentUntaxedPortionsOfPensionsStartIndex2526-1 : (parentUntaxedPortionsOfPensionsStartIndex2526-1)+parentUntaxedPortionsOfPensionsLength2526]), // Field # 203

		ParentPensionRollover: strings.TrimSpace(record[parentPensionRolloverStartIndex2526-1 : (parentPensionRolloverStartIndex2526-1)+parentPensionRolloverLength2526]), // Field # 204

		ParentAdjustedGrossIncome: strings.TrimSpace(record[parentAdjustedGrossIncomeStartIndex2526-1 : (parentAdjustedGrossIncomeStartIndex2526-1)+parentAdjustedGrossIncomeLength2526]), // Field # 205

		ParentIncomeTaxPaid: strings.TrimSpace(record[parentIncomeTaxPaidStartIndex2526-1 : (parentIncomeTaxPaidStartIndex2526-1)+parentIncomeTaxPaidLength2526]), // Field # 206

		ParentEarnedIncomeTaxCreditReceivedDuringTaxYear: strings.TrimSpace(record[parentEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex2526-1 : (parentEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex2526-1)+parentEarnedIncomeTaxCreditReceivedDuringTaxYearLength2526]), // Field # 207

		ParentDeductiblePaymentsToIRAKeoghOther: strings.TrimSpace(record[parentDeductiblePaymentsToIRAKeoghOtherStartIndex2526-1 : (parentDeductiblePaymentsToIRAKeoghOtherStartIndex2526-1)+parentDeductiblePaymentsToIRAKeoghOtherLength2526]), // Field # 208

		ParentEducationCredits: strings.TrimSpace(record[parentEducationCreditsStartIndex2526-1 : (parentEducationCreditsStartIndex2526-1)+parentEducationCreditsLength2526]), // Field # 209

		ParentFiledScheduleABDEFH: strings.TrimSpace(record[parentFiledScheduleABDEFHStartIndex2526-1 : (parentFiledScheduleABDEFHStartIndex2526-1)+parentFiledScheduleABDEFHLength2526]), // Field # 210

		ParentScheduleCAmount: strings.TrimSpace(record[parentScheduleCAmountStartIndex2526-1 : (parentScheduleCAmountStartIndex2526-1)+parentScheduleCAmountLength2526]), // Field # 211

		ParentCollegeGrantAndScholarshipAid: strings.TrimSpace(record[parentCollegeGrantAndScholarshipAidStartIndex2526-1 : (parentCollegeGrantAndScholarshipAidStartIndex2526-1)+parentCollegeGrantAndScholarshipAidLength2526]), // Field # 212

		ParentForeignEarnedIncomeExclusion: strings.TrimSpace(record[parentForeignEarnedIncomeExclusionStartIndex2526-1 : (parentForeignEarnedIncomeExclusionStartIndex2526-1)+parentForeignEarnedIncomeExclusionLength2526]), // Field # 213

		ParentChildSupportReceived: strings.TrimSpace(record[parentChildSupportReceivedStartIndex2526-1 : (parentChildSupportReceivedStartIndex2526-1)+parentChildSupportReceivedLength2526]), // Field # 214

		ParentTotalOfCashSavingsAndCheckingAccounts: strings.TrimSpace(record[parentTotalOfCashSavingsAndCheckingAccountsStartIndex2526-1 : (parentTotalOfCashSavingsAndCheckingAccountsStartIndex2526-1)+parentTotalOfCashSavingsAndCheckingAccountsLength2526]), // Field # 215

		ParentNetWorthOfCurrentInvestments: strings.TrimSpace(record[parentNetWorthOfCurrentInvestmentsStartIndex2526-1 : (parentNetWorthOfCurrentInvestmentsStartIndex2526-1)+parentNetWorthOfCurrentInvestmentsLength2526]), // Field # 216

		ParentNetWorthOfBusinessesAndInvestmentFarms: strings.TrimSpace(record[parentNetWorthOfBusinessesAndInvestmentFarmsStartIndex2526-1 : (parentNetWorthOfBusinessesAndInvestmentFarmsStartIndex2526-1)+parentNetWorthOfBusinessesAndInvestmentFarmsLength2526]), // Field # 217

		ParentConsentToRetrieveAndDiscloseFTI: strings.TrimSpace(record[parentConsentToRetrieveAndDiscloseFTIStartIndex2526-1 : (parentConsentToRetrieveAndDiscloseFTIStartIndex2526-1)+parentConsentToRetrieveAndDiscloseFTILength2526]), // Field # 218

		ParentSignature: strings.TrimSpace(record[parentSignatureStartIndex2526-1 : (parentSignatureStartIndex2526-1)+parentSignatureLength2526]), // Field # 219

		ParentSignatureDate: parseISIRDate2526(strings.TrimSpace(record[parentSignatureDateStartIndex2526-1 : (parentSignatureDateStartIndex2526-1)+parentSignatureDateLength2526])), // Field # 220

		ParentSpouseFirstName: strings.TrimSpace(record[parentSpouseFirstNameStartIndex2526-1 : (parentSpouseFirstNameStartIndex2526-1)+parentSpouseFirstNameLength2526]), // Field # 222

		ParentSpouseMiddleName: strings.TrimSpace(record[parentSpouseMiddleNameStartIndex2526-1 : (parentSpouseMiddleNameStartIndex2526-1)+parentSpouseMiddleNameLength2526]), // Field # 223

		ParentSpouseLastName: strings.TrimSpace(record[parentSpouseLastNameStartIndex2526-1 : (parentSpouseLastNameStartIndex2526-1)+parentSpouseLastNameLength2526]), // Field # 224

		ParentSpouseSuffix: strings.TrimSpace(record[parentSpouseSuffixStartIndex2526-1 : (parentSpouseSuffixStartIndex2526-1)+parentSpouseSuffixLength2526]), // Field # 225

		ParentSpouseDateOfBirth: parseISIRDate2526(strings.TrimSpace(record[parentSpouseDateOfBirthStartIndex2526-1 : (parentSpouseDateOfBirthStartIndex2526-1)+parentSpouseDateOfBirthLength2526])), // Field # 226

		ParentSpouseSSN: strings.TrimSpace(record[parentSpouseSSNStartIndex2526-1 : (parentSpouseSSNStartIndex2526-1)+parentSpouseSSNLength2526]), // Field # 227

		ParentSpouseITIN: strings.TrimSpace(record[parentSpouseITINStartIndex2526-1 : (parentSpouseITINStartIndex2526-1)+parentSpouseITINLength2526]), // Field # 228

		ParentSpousePhoneNumber: strings.TrimSpace(record[parentSpousePhoneNumberStartIndex2526-1 : (parentSpousePhoneNumberStartIndex2526-1)+parentSpousePhoneNumberLength2526]), // Field # 229

		ParentSpouseEmailAddress: strings.TrimSpace(record[parentSpouseEmailAddressStartIndex2526-1 : (parentSpouseEmailAddressStartIndex2526-1)+parentSpouseEmailAddressLength2526]), // Field # 230

		ParentSpouseStreetAddress: strings.TrimSpace(record[parentSpouseStreetAddressStartIndex2526-1 : (parentSpouseStreetAddressStartIndex2526-1)+parentSpouseStreetAddressLength2526]), // Field # 231

		ParentSpouseCity: strings.TrimSpace(record[parentSpouseCityStartIndex2526-1 : (parentSpouseCityStartIndex2526-1)+parentSpouseCityLength2526]), // Field # 232

		ParentSpouseState: strings.TrimSpace(record[parentSpouseStateStartIndex2526-1 : (parentSpouseStateStartIndex2526-1)+parentSpouseStateLength2526]), // Field # 233

		ParentSpouseZipCode: strings.TrimSpace(record[parentSpouseZipCodeStartIndex2526-1 : (parentSpouseZipCodeStartIndex2526-1)+parentSpouseZipCodeLength2526]), // Field # 234

		ParentSpouseCountry: strings.TrimSpace(record[parentSpouseCountryStartIndex2526-1 : (parentSpouseCountryStartIndex2526-1)+parentSpouseCountryLength2526]), // Field # 235

		ParentSpouseFiled1040Or1040NR: strings.TrimSpace(record[parentSpouseFiled1040Or1040NRStartIndex2526-1 : (parentSpouseFiled1040Or1040NRStartIndex2526-1)+parentSpouseFiled1040Or1040NRLength2526]), // Field # 236

		ParentSpouseFileNonUSTaxReturn: strings.TrimSpace(record[parentSpouseFileNonUSTaxReturnStartIndex2526-1 : (parentSpouseFileNonUSTaxReturnStartIndex2526-1)+parentSpouseFileNonUSTaxReturnLength2526]), // Field # 237

		ParentSpouseTaxReturnFilingStatus: strings.TrimSpace(record[parentSpouseTaxReturnFilingStatusStartIndex2526-1 : (parentSpouseTaxReturnFilingStatusStartIndex2526-1)+parentSpouseTaxReturnFilingStatusLength2526]), // Field # 238

		ParentSpouseIncomeEarnedFromWork: strings.TrimSpace(record[parentSpouseIncomeEarnedFromWorkStartIndex2526-1 : (parentSpouseIncomeEarnedFromWorkStartIndex2526-1)+parentSpouseIncomeEarnedFromWorkLength2526]), // Field # 239

		ParentSpouseTaxExemptInterestIncome: strings.TrimSpace(record[parentSpouseTaxExemptInterestIncomeStartIndex2526-1 : (parentSpouseTaxExemptInterestIncomeStartIndex2526-1)+parentSpouseTaxExemptInterestIncomeLength2526]), // Field # 240

		ParentSpouseUntaxedPortionsOfIRADistributions: strings.TrimSpace(record[parentSpouseUntaxedPortionsOfIRADistributionsStartIndex2526-1 : (parentSpouseUntaxedPortionsOfIRADistributionsStartIndex2526-1)+parentSpouseUntaxedPortionsOfIRADistributionsLength2526]), // Field # 241

		ParentSpouseIRARollover: strings.TrimSpace(record[parentSpouseIRARolloverStartIndex2526-1 : (parentSpouseIRARolloverStartIndex2526-1)+parentSpouseIRARolloverLength2526]), // Field # 242

		ParentSpouseUntaxedPortionsOfPensions: strings.TrimSpace(record[parentSpouseUntaxedPortionsOfPensionsStartIndex2526-1 : (parentSpouseUntaxedPortionsOfPensionsStartIndex2526-1)+parentSpouseUntaxedPortionsOfPensionsLength2526]), // Field # 243

		ParentSpousePensionRollover: strings.TrimSpace(record[parentSpousePensionRolloverStartIndex2526-1 : (parentSpousePensionRolloverStartIndex2526-1)+parentSpousePensionRolloverLength2526]), // Field # 244

		ParentSpouseAdjustedGrossIncome: strings.TrimSpace(record[parentSpouseAdjustedGrossIncomeStartIndex2526-1 : (parentSpouseAdjustedGrossIncomeStartIndex2526-1)+parentSpouseAdjustedGrossIncomeLength2526]), // Field # 245

		ParentSpouseIncomeTaxPaid: strings.TrimSpace(record[parentSpouseIncomeTaxPaidStartIndex2526-1 : (parentSpouseIncomeTaxPaidStartIndex2526-1)+parentSpouseIncomeTaxPaidLength2526]), // Field # 246

		ParentSpouseDeductiblePaymentsToIRAKeoghOther: strings.TrimSpace(record[parentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex2526-1 : (parentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex2526-1)+parentSpouseDeductiblePaymentsToIRAKeoghOtherLength2526]), // Field # 247

		ParentSpouseEducationCredits: strings.TrimSpace(record[parentSpouseEducationCreditsStartIndex2526-1 : (parentSpouseEducationCreditsStartIndex2526-1)+parentSpouseEducationCreditsLength2526]), // Field # 248

		ParentSpouseFiledScheduleABDEFH: strings.TrimSpace(record[parentSpouseFiledScheduleABDEFHStartIndex2526-1 : (parentSpouseFiledScheduleABDEFHStartIndex2526-1)+parentSpouseFiledScheduleABDEFHLength2526]), // Field # 249

		ParentSpouseScheduleCAmount: strings.TrimSpace(record[parentSpouseScheduleCAmountStartIndex2526-1 : (parentSpouseScheduleCAmountStartIndex2526-1)+parentSpouseScheduleCAmountLength2526]), // Field # 250

		ParentSpouseForeignEarnedIncomeExclusion: strings.TrimSpace(record[parentSpouseForeignEarnedIncomeExclusionStartIndex2526-1 : (parentSpouseForeignEarnedIncomeExclusionStartIndex2526-1)+parentSpouseForeignEarnedIncomeExclusionLength2526]), // Field # 251

		ParentSpouseConsentToRetrieveAndDiscloseFTI: strings.TrimSpace(record[parentSpouseConsentToRetrieveAndDiscloseFTIStartIndex2526-1 : (parentSpouseConsentToRetrieveAndDiscloseFTIStartIndex2526-1)+parentSpouseConsentToRetrieveAndDiscloseFTILength2526]), // Field # 252

		ParentSpouseSignature: strings.TrimSpace(record[parentSpouseSignatureStartIndex2526-1 : (parentSpouseSignatureStartIndex2526-1)+parentSpouseSignatureLength2526]), // Field # 253

		ParentSpouseSignatureDate: parseISIRDate2526(strings.TrimSpace(record[parentSpouseSignatureDateStartIndex2526-1 : (parentSpouseSignatureDateStartIndex2526-1)+parentSpouseSignatureDateLength2526])), // Field # 254

		PreparerFirstName: strings.TrimSpace(record[preparerFirstNameStartIndex2526-1 : (preparerFirstNameStartIndex2526-1)+preparerFirstNameLength2526]), // Field # 256

		PreparerLastName: strings.TrimSpace(record[preparerLastNameStartIndex2526-1 : (preparerLastNameStartIndex2526-1)+preparerLastNameLength2526]), // Field # 257

		PreparerSSN: strings.TrimSpace(record[preparerSSNStartIndex2526-1 : (preparerSSNStartIndex2526-1)+preparerSSNLength2526]), // Field # 258

		PreparerEIN: strings.TrimSpace(record[preparerEINStartIndex2526-1 : (preparerEINStartIndex2526-1)+preparerEINLength2526]), // Field # 259

		PreparerAffiliation: strings.TrimSpace(record[preparerAffiliationStartIndex2526-1 : (preparerAffiliationStartIndex2526-1)+preparerAffiliationLength2526]), // Field # 260

		PreparerStreetAddress: strings.TrimSpace(record[preparerStreetAddressStartIndex2526-1 : (preparerStreetAddressStartIndex2526-1)+preparerStreetAddressLength2526]), // Field # 261

		PreparerCity: strings.TrimSpace(record[preparerCityStartIndex2526-1 : (preparerCityStartIndex2526-1)+preparerCityLength2526]), // Field # 262

		PreparerState: strings.TrimSpace(record[preparerStateStartIndex2526-1 : (preparerStateStartIndex2526-1)+preparerStateLength2526]), // Field # 263

		PreparerZipCode: strings.TrimSpace(record[preparerZipCodeStartIndex2526-1 : (preparerZipCodeStartIndex2526-1)+preparerZipCodeLength2526]), // Field # 264

		PreparerSignature: strings.TrimSpace(record[preparerSignatureStartIndex2526-1 : (preparerSignatureStartIndex2526-1)+preparerSignatureLength2526]), // Field # 265

		PreparerSignatureDate: parseISIRDate2526(strings.TrimSpace(record[preparerSignatureDateStartIndex2526-1 : (preparerSignatureDateStartIndex2526-1)+preparerSignatureDateLength2526])), // Field # 266

		StudentAffirmationStatus: strings.TrimSpace(record[studentAffirmationStatusStartIndex2526-1 : (studentAffirmationStatusStartIndex2526-1)+studentAffirmationStatusLength2526]), // Field # 268

		StudentSpouseAffirmationStatus: strings.TrimSpace(record[studentSpouseAffirmationStatusStartIndex2526-1 : (studentSpouseAffirmationStatusStartIndex2526-1)+studentSpouseAffirmationStatusLength2526]), // Field # 269

		ParentAffirmationStatus: strings.TrimSpace(record[parentAffirmationStatusStartIndex2526-1 : (parentAffirmationStatusStartIndex2526-1)+parentAffirmationStatusLength2526]), // Field # 270

		ParentSpouseOrPartnerAffirmationStatus: strings.TrimSpace(record[parentSpouseOrPartnerAffirmationStatusStartIndex2526-1 : (parentSpouseOrPartnerAffirmationStatusStartIndex2526-1)+parentSpouseOrPartnerAffirmationStatusLength2526]), // Field # 271

		StudentDateConsentGranted: parseISIRDate2526(strings.TrimSpace(record[studentDateConsentGrantedStartIndex2526-1 : (studentDateConsentGrantedStartIndex2526-1)+studentDateConsentGrantedLength2526])), // Field # 272

		StudentSpouseDateConsentGranted: parseISIRDate2526(strings.TrimSpace(record[studentSpouseDateConsentGrantedStartIndex2526-1 : (studentSpouseDateConsentGrantedStartIndex2526-1)+studentSpouseDateConsentGrantedLength2526])), // Field # 273

		ParentDateConsentGranted: parseISIRDate2526(strings.TrimSpace(record[parentDateConsentGrantedStartIndex2526-1 : (parentDateConsentGrantedStartIndex2526-1)+parentDateConsentGrantedLength2526])), // Field # 274

		ParentSpouseOrPartnerDateConsentGranted: parseISIRDate2526(strings.TrimSpace(record[parentSpouseOrPartnerDateConsentGrantedStartIndex2526-1 : (parentSpouseOrPartnerDateConsentGrantedStartIndex2526-1)+parentSpouseOrPartnerDateConsentGrantedLength2526])), // Field # 275

		StudentTransunionMatchStatus: strings.TrimSpace(record[studentTransunionMatchStatusStartIndex2526-1 : (studentTransunionMatchStatusStartIndex2526-1)+studentTransunionMatchStatusLength2526]), // Field # 276

		StudentSpouseTransunionMatchStatus: strings.TrimSpace(record[studentSpouseTransunionMatchStatusStartIndex2526-1 : (studentSpouseTransunionMatchStatusStartIndex2526-1)+studentSpouseTransunionMatchStatusLength2526]), // Field # 277

		StudentParentTransunionMatchStatus: strings.TrimSpace(record[studentParentTransunionMatchStatusStartIndex2526-1 : (studentParentTransunionMatchStatusStartIndex2526-1)+studentParentTransunionMatchStatusLength2526]), // Field # 278

		StudentParentSpouseTransunionMatchStatus: strings.TrimSpace(record[studentParentSpouseTransunionMatchStatusStartIndex2526-1 : (studentParentSpouseTransunionMatchStatusStartIndex2526-1)+studentParentSpouseTransunionMatchStatusLength2526]), // Field # 279

		CorrectionAppliedAgainstTransactionNumber: strings.TrimSpace(record[correctionAppliedAgainstTransactionNumberStartIndex2526-1 : (correctionAppliedAgainstTransactionNumberStartIndex2526-1)+correctionAppliedAgainstTransactionNumberLength2526]), // Field # 280

		ProfessionalJudgement: strings.TrimSpace(record[professionalJudgementStartIndex2526-1 : (professionalJudgementStartIndex2526-1)+professionalJudgementLength2526]), // Field # 281

		DependencyOverrideIndicator: strings.TrimSpace(record[dependencyOverrideIndicatorStartIndex2526-1 : (dependencyOverrideIndicatorStartIndex2526-1)+dependencyOverrideIndicatorLength2526]), // Field # 282

		FAAFederalSchoolCode: strings.TrimSpace(record[fAAFederalSchoolCodeStartIndex2526-1 : (fAAFederalSchoolCodeStartIndex2526-1)+fAAFederalSchoolCodeLength2526]), // Field # 283

		FAASignature: strings.TrimSpace(record[fAASignatureStartIndex2526-1 : (fAASignatureStartIndex2526-1)+fAASignatureLength2526]), // Field # 284

		IASGIndicator: strings.TrimSpace(record[iASGIndicatorStartIndex2526-1 : (iASGIndicatorStartIndex2526-1)+iASGIndicatorLength2526]), // Field # 285

		ChildrenOfFallenHeroesIndicator: strings.TrimSpace(record[childrenOfFallenHeroesIndicatorStartIndex2526-1 : (childrenOfFallenHeroesIndicatorStartIndex2526-1)+childrenOfFallenHeroesIndicatorLength2526]), // Field # 286

		ElectronicTransactionIndicatorDestinationNumber: strings.TrimSpace(record[electronicTransactionIndicatorDestinationNumberStartIndex2526-1 : (electronicTransactionIndicatorDestinationNumberStartIndex2526-1)+electronicTransactionIndicatorDestinationNumberLength2526]), // Field # 287

		StudentSignatureSource: strings.TrimSpace(record[studentSignatureSourceStartIndex2526-1 : (studentSignatureSourceStartIndex2526-1)+studentSignatureSourceLength2526]), // Field # 288

		StudentSpouseSignatureSource: strings.TrimSpace(record[studentSpouseSignatureSourceStartIndex2526-1 : (studentSpouseSignatureSourceStartIndex2526-1)+studentSpouseSignatureSourceLength2526]), // Field # 289

		ParentSignatureSource: strings.TrimSpace(record[parentSignatureSourceStartIndex2526-1 : (parentSignatureSourceStartIndex2526-1)+parentSignatureSourceLength2526]), // Field # 290

		ParentSpouseOrPartnerSignatureSource: strings.TrimSpace(record[parentSpouseOrPartnerSignatureSourceStartIndex2526-1 : (parentSpouseOrPartnerSignatureSourceStartIndex2526-1)+parentSpouseOrPartnerSignatureSourceLength2526]), // Field # 291

		SpecialHandlingIndicator: strings.TrimSpace(record[specialHandlingIndicatorStartIndex2526-1 : (specialHandlingIndicatorStartIndex2526-1)+specialHandlingIndicatorLength2526]), // Field # 292

		AddressOnlyChangeFlag: strings.TrimSpace(record[addressOnlyChangeFlagStartIndex2526-1 : (addressOnlyChangeFlagStartIndex2526-1)+addressOnlyChangeFlagLength2526]), // Field # 293

		FPSPushedISIRFlag: strings.TrimSpace(record[fpsPushedISIRFlagStartIndex2526-1 : (fpsPushedISIRFlagStartIndex2526-1)+fpsPushedISIRFlagLength2526]), // Field # 294

		RejectStatusChangeFlag: strings.TrimSpace(record[rejectStatusChangeFlagStartIndex2526-1 : (rejectStatusChangeFlagStartIndex2526-1)+rejectStatusChangeFlagLength2526]), // Field # 295

		VerificationTrackingFlag: strings.TrimSpace(record[verificationTrackingFlagStartIndex2526-1 : (verificationTrackingFlagStartIndex2526-1)+verificationTrackingFlagLength2526]), // Field # 296

		StudentSelectedForVerification: strings.TrimSpace(record[studentSelectedForVerificationStartIndex2526-1 : (studentSelectedForVerificationStartIndex2526-1)+studentSelectedForVerificationLength2526]), // Field # 297

		IncarceratedApplicantFlag: strings.TrimSpace(record[incarceratedApplicantFlagStartIndex2526-1 : (incarceratedApplicantFlagStartIndex2526-1)+incarceratedApplicantFlagLength2526]), // Field # 298

		NSLDSTransactionNumber: strings.TrimSpace(record[nsldsTransactionNumberStartIndex2526-1 : (nsldsTransactionNumberStartIndex2526-1)+nsldsTransactionNumberLength2526]), // Field # 299

		NSLDSDatabaseResultsFlag: strings.TrimSpace(record[nsldsDatabaseResultsFlagStartIndex2526-1 : (nsldsDatabaseResultsFlagStartIndex2526-1)+nsldsDatabaseResultsFlagLength2526]), // Field # 300

		HighSchoolFlag: strings.TrimSpace(record[highSchoolFlagStartIndex2526-1 : (highSchoolFlagStartIndex2526-1)+highSchoolFlagLength2526]), // Field # 301

		StudentTotalFederalWorkStudyEarnings: strings.TrimSpace(record[studentTotalFederalWorkStudyEarningsStartIndex2526-1 : (studentTotalFederalWorkStudyEarningsStartIndex2526-1)+studentTotalFederalWorkStudyEarningsLength2526]), // Field # 302

		StudentSpouseTotalFederalWorkStudyEarnings: strings.TrimSpace(record[studentSpouseTotalFederalWorkStudyEarningsStartIndex2526-1 : (studentSpouseTotalFederalWorkStudyEarningsStartIndex2526-1)+studentSpouseTotalFederalWorkStudyEarningsLength2526]), // Field # 303

		ParentTotalFederalWorkStudyEarnings: strings.TrimSpace(record[parentTotalFederalWorkStudyEarningsStartIndex2526-1 : (parentTotalFederalWorkStudyEarningsStartIndex2526-1)+parentTotalFederalWorkStudyEarningsLength2526]), // Field # 304

		ParentSpouseOrPartnerTotalFederalWorkStudyEarnings: strings.TrimSpace(record[parentSpouseOrPartnerTotalFederalWorkStudyEarningsStartIndex2526-1 : (parentSpouseOrPartnerTotalFederalWorkStudyEarningsStartIndex2526-1)+parentSpouseOrPartnerTotalFederalWorkStudyEarningsLength2526]), // Field # 305

		TotalParentAllowancesAgainstIncome: strings.TrimSpace(record[totalParentAllowancesAgainstIncomeStartIndex2526-1 : (totalParentAllowancesAgainstIncomeStartIndex2526-1)+totalParentAllowancesAgainstIncomeLength2526]), // Field # 306

		ParentPayrollTaxAllowance: strings.TrimSpace(record[parentPayrollTaxAllowanceStartIndex2526-1 : (parentPayrollTaxAllowanceStartIndex2526-1)+parentPayrollTaxAllowanceLength2526]), // Field # 307

		ParentIncomeProtectionAllowance: strings.TrimSpace(record[parentIncomeProtectionAllowanceStartIndex2526-1 : (parentIncomeProtectionAllowanceStartIndex2526-1)+parentIncomeProtectionAllowanceLength2526]), // Field # 308

		ParentEmploymentExpenseAllowance: strings.TrimSpace(record[parentEmploymentExpenseAllowanceStartIndex2526-1 : (parentEmploymentExpenseAllowanceStartIndex2526-1)+parentEmploymentExpenseAllowanceLength2526]), // Field # 309

		ParentAvailableIncome: strings.TrimSpace(record[parentAvailableIncomeStartIndex2526-1 : (parentAvailableIncomeStartIndex2526-1)+parentAvailableIncomeLength2526]), // Field # 310

		ParentAdjustedAvailableIncome: strings.TrimSpace(record[parentAdjustedAvailableIncomeStartIndex2526-1 : (parentAdjustedAvailableIncomeStartIndex2526-1)+parentAdjustedAvailableIncomeLength2526]), // Field # 311

		ParentContribution: strings.TrimSpace(record[parentContributionStartIndex2526-1 : (parentContributionStartIndex2526-1)+parentContributionLength2526]), // Field # 312

		StudentPayrollTaxAllowance: strings.TrimSpace(record[studentPayrollTaxAllowanceStartIndex2526-1 : (studentPayrollTaxAllowanceStartIndex2526-1)+studentPayrollTaxAllowanceLength2526]), // Field # 313

		StudentIncomeProtectionAllowance: strings.TrimSpace(record[studentIncomeProtectionAllowanceStartIndex2526-1 : (studentIncomeProtectionAllowanceStartIndex2526-1)+studentIncomeProtectionAllowanceLength2526]), // Field # 314

		StudentAllowanceForParentsNegativeAdjustedAvailableIncome: strings.TrimSpace(record[studentAllowanceForParentsNegativeAdjustedAvailableIncomeStartIndex2526-1 : (studentAllowanceForParentsNegativeAdjustedAvailableIncomeStartIndex2526-1)+studentAllowanceForParentsNegativeAdjustedAvailableIncomeLength2526]), // Field # 315

		StudentEmploymentExpenseAllowance: strings.TrimSpace(record[studentEmploymentExpenseAllowanceStartIndex2526-1 : (studentEmploymentExpenseAllowanceStartIndex2526-1)+studentEmploymentExpenseAllowanceLength2526]), // Field # 316

		TotalStudentAllowancesAgainstIncome: strings.TrimSpace(record[totalStudentAllowancesAgainstIncomeStartIndex2526-1 : (totalStudentAllowancesAgainstIncomeStartIndex2526-1)+totalStudentAllowancesAgainstIncomeLength2526]), // Field # 317

		StudentAvailableIncome: strings.TrimSpace(record[studentAvailableIncomeStartIndex2526-1 : (studentAvailableIncomeStartIndex2526-1)+studentAvailableIncomeLength2526]), // Field # 318

		StudentContributionFromIncome: strings.TrimSpace(record[studentContributionFromIncomeStartIndex2526-1 : (studentContributionFromIncomeStartIndex2526-1)+studentContributionFromIncomeLength2526]), // Field # 319

		StudentAdjustedAvailableIncome: strings.TrimSpace(record[studentAdjustedAvailableIncomeStartIndex2526-1 : (studentAdjustedAvailableIncomeStartIndex2526-1)+studentAdjustedAvailableIncomeLength2526]), // Field # 320

		TotalStudentContributionFromSAAI: strings.TrimSpace(record[totalStudentContributionFromSAAIStartIndex2526-1 : (totalStudentContributionFromSAAIStartIndex2526-1)+totalStudentContributionFromSAAILength2526]), // Field # 321

		ParentDiscretionaryNetWorth: strings.TrimSpace(record[parentDiscretionaryNetWorthStartIndex2526-1 : (parentDiscretionaryNetWorthStartIndex2526-1)+parentDiscretionaryNetWorthLength2526]), // Field # 322

		ParentNetWorth: strings.TrimSpace(record[parentNetWorthStartIndex2526-1 : (parentNetWorthStartIndex2526-1)+parentNetWorthLength2526]), // Field # 323

		ParentAssetProtectionAllowance: strings.TrimSpace(record[parentAssetProtectionAllowanceStartIndex2526-1 : (parentAssetProtectionAllowanceStartIndex2526-1)+parentAssetProtectionAllowanceLength2526]), // Field # 324

		ParentContributionFromAssets: strings.TrimSpace(record[parentContributionFromAssetsStartIndex2526-1 : (parentContributionFromAssetsStartIndex2526-1)+parentContributionFromAssetsLength2526]), // Field # 325

		StudentNetWorth: strings.TrimSpace(record[studentNetWorthStartIndex2526-1 : (studentNetWorthStartIndex2526-1)+studentNetWorthLength2526]), // Field # 326

		StudentAssetProtectionAllowance: strings.TrimSpace(record[studentAssetProtectionAllowanceStartIndex2526-1 : (studentAssetProtectionAllowanceStartIndex2526-1)+studentAssetProtectionAllowanceLength2526]), // Field # 327

		StudentContributionFromAssets: strings.TrimSpace(record[studentContributionFromAssetsStartIndex2526-1 : (studentContributionFromAssetsStartIndex2526-1)+studentContributionFromAssetsLength2526]), // Field # 328

		AssumedStudentFamilySize: strings.TrimSpace(record[assumedStudentFamilySizeStartIndex2526-1 : (assumedStudentFamilySizeStartIndex2526-1)+assumedStudentFamilySizeLength2526]), // Field # 329

		AssumedParentFamilySize: strings.TrimSpace(record[assumedParentFamilySizeStartIndex2526-1 : (assumedParentFamilySizeStartIndex2526-1)+assumedParentFamilySizeLength2526]), // Field # 330

		StudentFirstNameCHVFlags: strings.TrimSpace(record[studentFirstNameCHVFlagsStartIndex2526-1 : (studentFirstNameCHVFlagsStartIndex2526-1)+studentFirstNameCHVFlagsLength2526]), // Field # 331

		StudentMiddleNameCHVFlags: strings.TrimSpace(record[studentMiddleNameCHVFlagsStartIndex2526-1 : (studentMiddleNameCHVFlagsStartIndex2526-1)+studentMiddleNameCHVFlagsLength2526]), // Field # 332

		StudentLastNameCHVFLags: strings.TrimSpace(record[studentLastNameCHVFLagsStartIndex2526-1 : (studentLastNameCHVFLagsStartIndex2526-1)+studentLastNameCHVFLagsLength2526]), // Field # 333

		StudentSuffixCHVFLags: strings.TrimSpace(record[studentSuffixCHVFLagsStartIndex2526-1 : (studentSuffixCHVFLagsStartIndex2526-1)+studentSuffixCHVFLagsLength2526]), // Field # 334

		StudentDateOfBirthCHVFLags: strings.TrimSpace(record[studentDateOfBirthCHVFLagsStartIndex2526-1 : (studentDateOfBirthCHVFLagsStartIndex2526-1)+studentDateOfBirthCHVFLagsLength2526]), // Field # 335

		StudentSSNCHVFlags: strings.TrimSpace(record[studentSSNCHVFlagsStartIndex2526-1 : (studentSSNCHVFlagsStartIndex2526-1)+studentSSNCHVFlagsLength2526]), // Field # 336

		StudentITINCHVFLags: strings.TrimSpace(record[studentITINCHVFLagsStartIndex2526-1 : (studentITINCHVFLagsStartIndex2526-1)+studentITINCHVFLagsLength2526]), // Field # 337

		StudentPhoneNumberCHVFlags: strings.TrimSpace(record[studentPhoneNumberCHVFlagsStartIndex2526-1 : (studentPhoneNumberCHVFlagsStartIndex2526-1)+studentPhoneNumberCHVFlagsLength2526]), // Field # 338

		StudentEmailAddressCHVFlags: strings.TrimSpace(record[studentEmailAddressCHVFlagsStartIndex2526-1 : (studentEmailAddressCHVFlagsStartIndex2526-1)+studentEmailAddressCHVFlagsLength2526]), // Field # 339

		StudentStreetAddressCHVFlags: strings.TrimSpace(record[studentStreetAddressCHVFlagsStartIndex2526-1 : (studentStreetAddressCHVFlagsStartIndex2526-1)+studentStreetAddressCHVFlagsLength2526]), // Field # 340

		StudentCityCHVFLags: strings.TrimSpace(record[studentCityCHVFLagsStartIndex2526-1 : (studentCityCHVFLagsStartIndex2526-1)+studentCityCHVFLagsLength2526]), // Field # 341

		StudentStateCHVFlags: strings.TrimSpace(record[studentStateCHVFlagsStartIndex2526-1 : (studentStateCHVFlagsStartIndex2526-1)+studentStateCHVFlagsLength2526]), // Field # 342

		StudentZipCodeCHVFlags: strings.TrimSpace(record[studentZipCodeCHVFlagsStartIndex2526-1 : (studentZipCodeCHVFlagsStartIndex2526-1)+studentZipCodeCHVFlagsLength2526]), // Field # 343

		StudentCountryCHVFlags: strings.TrimSpace(record[studentCountryCHVFlagsStartIndex2526-1 : (studentCountryCHVFlagsStartIndex2526-1)+studentCountryCHVFlagsLength2526]), // Field # 344

		StudentMaritalStatusCHVFlags: strings.TrimSpace(record[studentMaritalStatusCHVFlagsStartIndex2526-1 : (studentMaritalStatusCHVFlagsStartIndex2526-1)+studentMaritalStatusCHVFlagsLength2526]), // Field # 345

		StudentGradeLevelInCollegeCHVFlags: strings.TrimSpace(record[studentGradeLevelInCollegeCHVFlagsStartIndex2526-1 : (studentGradeLevelInCollegeCHVFlagsStartIndex2526-1)+studentGradeLevelInCollegeCHVFlagsLength2526]), // Field # 346

		StudentFirstBachelorsDegreeBeforeSchoolYearCHVFlags: strings.TrimSpace(record[studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsStartIndex2526-1 : (studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsStartIndex2526-1)+studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsLength2526]), // Field # 347

		StudentPursuingTeacherCertificationCHVFlags: strings.TrimSpace(record[studentPursuingTeacherCertificationCHVFlagsStartIndex2526-1 : (studentPursuingTeacherCertificationCHVFlagsStartIndex2526-1)+studentPursuingTeacherCertificationCHVFlagsLength2526]), // Field # 348

		StudentActiveDutyCHVFlags: strings.TrimSpace(record[studentActiveDutyCHVFlagsStartIndex2526-1 : (studentActiveDutyCHVFlagsStartIndex2526-1)+studentActiveDutyCHVFlagsLength2526]), // Field # 349

		StudentVeteranCHVFlags: strings.TrimSpace(record[studentVeteranCHVFlagsStartIndex2526-1 : (studentVeteranCHVFlagsStartIndex2526-1)+studentVeteranCHVFlagsLength2526]), // Field # 350

		StudentChildOrOtherDependentsCHVFlags: strings.TrimSpace(record[studentChildOrOtherDependentsCHVFlagsStartIndex2526-1 : (studentChildOrOtherDependentsCHVFlagsStartIndex2526-1)+studentChildOrOtherDependentsCHVFlagsLength2526]), // Field # 351

		StudentParentsDeceasedCHVFlags: strings.TrimSpace(record[studentParentsDeceasedCHVFlagsStartIndex2526-1 : (studentParentsDeceasedCHVFlagsStartIndex2526-1)+studentParentsDeceasedCHVFlagsLength2526]), // Field # 352

		StudentWardOfCourtCHVFlags: strings.TrimSpace(record[studentWardOfCourtCHVFlagsStartIndex2526-1 : (studentWardOfCourtCHVFlagsStartIndex2526-1)+studentWardOfCourtCHVFlagsLength2526]), // Field # 353

		StudentInFosterCareCHVFlags: strings.TrimSpace(record[studentInFosterCareCHVFlagsStartIndex2526-1 : (studentInFosterCareCHVFlagsStartIndex2526-1)+studentInFosterCareCHVFlagsLength2526]), // Field # 354

		StudentEmancipatedMinorCHVFlags: strings.TrimSpace(record[studentEmancipatedMinorCHVFlagsStartIndex2526-1 : (studentEmancipatedMinorCHVFlagsStartIndex2526-1)+studentEmancipatedMinorCHVFlagsLength2526]), // Field # 355

		StudentLegalGuardianshipCHVFlags: strings.TrimSpace(record[studentLegalGuardianshipCHVFlagsStartIndex2526-1 : (studentLegalGuardianshipCHVFlagsStartIndex2526-1)+studentLegalGuardianshipCHVFlagsLength2526]), // Field # 356

		StudentPersonalCircumstancesNoneOfTheAboveCHVFlags: strings.TrimSpace(record[studentPersonalCircumstancesNoneOfTheAboveCHVFlagsStartIndex2526-1 : (studentPersonalCircumstancesNoneOfTheAboveCHVFlagsStartIndex2526-1)+studentPersonalCircumstancesNoneOfTheAboveCHVFlagsLength2526]), // Field # 357

		StudentUnaccompaniedHomelessYouthOrIsUnaccompaniedAtRiskSelfSupportingCHVFlags: strings.TrimSpace(record[studentUnaccompaniedHomelessYouthOrIsUnaccompaniedAtRIskSelfSupportingCHVFlagsStartIndex2526-1 : (studentUnaccompaniedHomelessYouthOrIsUnaccompaniedAtRIskSelfSupportingCHVFlagsStartIndex2526-1)+studentUnaccompaniedHomelessYouthOrIsUnaccompaniedAtRIskSelfSupportingCHVFlagsLength2526]), // Field # 358

		StudentUnaccompaniedAndHomelessGeneralCHVFlags: strings.TrimSpace(record[studentUnaccompaniedAndHomelessGeneralCHVFlagsStartIndex2526-1 : (studentUnaccompaniedAndHomelessGeneralCHVFlagsStartIndex2526-1)+studentUnaccompaniedAndHomelessGeneralCHVFlagsLength2526]), // Field # 359

		StudentUnaccompaniedAndHomelessHSCHVFlags: strings.TrimSpace(record[studentUnaccompaniedAndHomelessHSCHVFlagsStartIndex2526-1 : (studentUnaccompaniedAndHomelessHSCHVFlagsStartIndex2526-1)+studentUnaccompaniedAndHomelessHSCHVFlagsLength2526]), // Field # 360

		StudentUnaccompaniedAndHomelessTRIOCHVFlags: strings.TrimSpace(record[studentUnaccompaniedAndHomelessTRIOCHVFlagsStartIndex2526-1 : (studentUnaccompaniedAndHomelessTRIOCHVFlagsStartIndex2526-1)+studentUnaccompaniedAndHomelessTRIOCHVFlagsLength2526]), // Field # 361

		StudentUnaccompaniedAndHomelessFAACHVFlags: strings.TrimSpace(record[studentUnaccompaniedAndHomelessFAACHVFlagsStartIndex2526-1 : (studentUnaccompaniedAndHomelessFAACHVFlagsStartIndex2526-1)+studentUnaccompaniedAndHomelessFAACHVFlagsLength2526]), // Field # 362

		StudentHomelessnessNoneOfTheAboveCHVFlags: strings.TrimSpace(record[studentHomelessnessNoneOfTheAboveCHVFlagsStartIndex2526-1 : (studentHomelessnessNoneOfTheAboveCHVFlagsStartIndex2526-1)+studentHomelessnessNoneOfTheAboveCHVFlagsLength2526]), // Field # 363

		StudentHasUnusualCircumstanceCHVFlags: strings.TrimSpace(record[studentHasUnusualCircumstanceCHVFlagsStartIndex2526-1 : (studentHasUnusualCircumstanceCHVFlagsStartIndex2526-1)+studentHasUnusualCircumstanceCHVFlagsLength2526]), // Field # 364

		StudentUnsubOnlyCHVFlags: strings.TrimSpace(record[studentUnsubOnlyCHVFlagsStartIndex2526-1 : (studentUnsubOnlyCHVFlagsStartIndex2526-1)+studentUnsubOnlyCHVFlagsLength2526]), // Field # 365

		StudentUpdatedFamilySizeCHVFlags: strings.TrimSpace(record[studentUpdatedFamilySizeCHVFlagsStartIndex2526-1 : (studentUpdatedFamilySizeCHVFlagsStartIndex2526-1)+studentUpdatedFamilySizeCHVFlagsLength2526]), // Field # 366

		StudentNumberInCollegeCorrectionCHVFlags: strings.TrimSpace(record[studentNumberInCollegeCorrectionCHVFlagsStartIndex2526-1 : (studentNumberInCollegeCorrectionCHVFlagsStartIndex2526-1)+studentNumberInCollegeCorrectionCHVFlagsLength2526]), // Field # 367

		StudentCitizenshipStatusCorrectionCHVFlags: strings.TrimSpace(record[studentCitizenshipStatusCorrectionCHVFlagsStartIndex2526-1 : (studentCitizenshipStatusCorrectionCHVFlagsStartIndex2526-1)+studentCitizenshipStatusCorrectionCHVFlagsLength2526]), // Field # 368

		StudentANumberCHVFlags: strings.TrimSpace(record[studentANumberCHVFlagsStartIndex2526-1 : (studentANumberCHVFlagsStartIndex2526-1)+studentANumberCHVFlagsLength2526]), // Field # 369

		StudentStateOfLegalResidenceCHVFlags: strings.TrimSpace(record[studentStateOfLegalResidenceCHVFlagsStartIndex2526-1 : (studentStateOfLegalResidenceCHVFlagsStartIndex2526-1)+studentStateOfLegalResidenceCHVFlagsLength2526]), // Field # 370

		StudentLegalResidenceDateCHVFlags: strings.TrimSpace(record[studentLegalResidenceDateCHVFlagsStartIndex2526-1 : (studentLegalResidenceDateCHVFlagsStartIndex2526-1)+studentLegalResidenceDateCHVFlagsLength2526]), // Field # 371

		StudentEitherParentAttendCollegeCHVFlags: strings.TrimSpace(record[studentEitherParentAttendCollegeCHVFlagsStartIndex2526-1 : (studentEitherParentAttendCollegeCHVFlagsStartIndex2526-1)+studentEitherParentAttendCollegeCHVFlagsLength2526]), // Field # 372

		StudentParentKilledInTheLineOfDutyCHVFlags: strings.TrimSpace(record[studentParentKilledInTheLineOfDutyCHVFlagsStartIndex2526-1 : (studentParentKilledInTheLineOfDutyCHVFlagsStartIndex2526-1)+studentParentKilledInTheLineOfDutyCHVFlagsLength2526]), // Field # 373

		StudentHighSchoolCompletionStatusCHVFlags: strings.TrimSpace(record[studentHighSchoolCompletionStatusCHVFlagsStartIndex2526-1 : (studentHighSchoolCompletionStatusCHVFlagsStartIndex2526-1)+studentHighSchoolCompletionStatusCHVFlagsLength2526]), // Field # 374

		StudentHighSchoolNameCHVFlags: strings.TrimSpace(record[studentHighSchoolNameCHVFlagsStartIndex2526-1 : (studentHighSchoolNameCHVFlagsStartIndex2526-1)+studentHighSchoolNameCHVFlagsLength2526]), // Field # 375

		StudentHighSchoolCityCHVFlags: strings.TrimSpace(record[studentHighSchoolCityCHVFlagsStartIndex2526-1 : (studentHighSchoolCityCHVFlagsStartIndex2526-1)+studentHighSchoolCityCHVFlagsLength2526]), // Field # 376

		StudentHighSchoolStateCHVFlags: strings.TrimSpace(record[studentHighSchoolStateCHVFlagsStartIndex2526-1 : (studentHighSchoolStateCHVFlagsStartIndex2526-1)+studentHighSchoolStateCHVFlagsLength2526]), // Field # 377

		StudentHighSchoolEquivalentDiplomaNameCHVFlags: strings.TrimSpace(record[studentHighSchoolEquivalentDiplomaNameCHVFlagsStartIndex2526-1 : (studentHighSchoolEquivalentDiplomaNameCHVFlagsStartIndex2526-1)+studentHighSchoolEquivalentDiplomaNameCHVFlagsLength2526]), // Field # 378

		StudentHighSchoolEquivalentDiplomaStateCHVFlags: strings.TrimSpace(record[studentHighSchoolEquivalentDiplomaStateCHVFlagsStartIndex2526-1 : (studentHighSchoolEquivalentDiplomaStateCHVFlagsStartIndex2526-1)+studentHighSchoolEquivalentDiplomaStateCHVFlagsLength2526]), // Field # 379

		StudentReceivedEITCCHVFlags: strings.TrimSpace(record[studentReceivedEITCCHVFlagsStartIndex2526-1 : (studentReceivedEITCCHVFlagsStartIndex2526-1)+studentReceivedEITCCHVFlagsLength2526]), // Field # 380

		StudentReceivedFederalHousingAssistanceCHVFlags: strings.TrimSpace(record[studentReceivedFederalHousingAssistanceCHVFlagsStartIndex2526-1 : (studentReceivedFederalHousingAssistanceCHVFlagsStartIndex2526-1)+studentReceivedFederalHousingAssistanceCHVFlagsLength2526]), // Field # 381

		StudentReceivedFreeReducedPriceLunchCHVFlags: strings.TrimSpace(record[studentReceivedFreeReducedPriceLunchCHVFlagsStartIndex2526-1 : (studentReceivedFreeReducedPriceLunchCHVFlagsStartIndex2526-1)+studentReceivedFreeReducedPriceLunchCHVFlagsLength2526]), // Field # 382

		StudentReceivedMedicaidCHVFlags: strings.TrimSpace(record[studentReceivedMedicaidCHVFlagsStartIndex2526-1 : (studentReceivedMedicaidCHVFlagsStartIndex2526-1)+studentReceivedMedicaidCHVFlagsLength2526]), // Field # 383

		StudentReceivedRefundableCreditFor36BHealthPlanCHVFlags: strings.TrimSpace(record[studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex2526-1 : (studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex2526-1)+studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsLength2526]), // Field # 384

		StudentReceivedSNAPCHVFlags: strings.TrimSpace(record[studentReceivedSNAPCHVFlagsStartIndex2526-1 : (studentReceivedSNAPCHVFlagsStartIndex2526-1)+studentReceivedSNAPCHVFlagsLength2526]), // Field # 385

		StudentReceivedSupplementalSecurityIncomeCHVFlags: strings.TrimSpace(record[studentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex2526-1 : (studentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex2526-1)+studentReceivedSupplementalSecurityIncomeCHVFlagsLength2526]), // Field # 386

		StudentReceivedTANFCHVFlags: strings.TrimSpace(record[studentReceivedTANFCHVFlagsStartIndex2526-1 : (studentReceivedTANFCHVFlagsStartIndex2526-1)+studentReceivedTANFCHVFlagsLength2526]), // Field # 387

		StudentReceivedWICCHVFlags: strings.TrimSpace(record[studentReceivedWICCHVFlagsStartIndex2526-1 : (studentReceivedWICCHVFlagsStartIndex2526-1)+studentReceivedWICCHVFlagsLength2526]), // Field # 388

		StudentFederalBenefitsNoneOfTheAboveCHVFlags: strings.TrimSpace(record[studentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex2526-1 : (studentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex2526-1)+studentFederalBenefitsNoneOfTheAboveCHVFlagsLength2526]), // Field # 389

		StudentFiled1040Or1040NRCHVFlags: strings.TrimSpace(record[studentFiled1040Or1040NRCHVFlagsStartIndex2526-1 : (studentFiled1040Or1040NRCHVFlagsStartIndex2526-1)+studentFiled1040Or1040NRCHVFlagsLength2526]), // Field # 390

		StudentFiledNonUSTaxReturnCHVFlags: strings.TrimSpace(record[studentFiledNonUSTaxReturnCHVFlagsStartIndex2526-1 : (studentFiledNonUSTaxReturnCHVFlagsStartIndex2526-1)+studentFiledNonUSTaxReturnCHVFlagsLength2526]), // Field # 391

		StudentFiledJointReturnWithCurrentSpouseCHVFlags: strings.TrimSpace(record[studentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex2526-1 : (studentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex2526-1)+studentFiledJointReturnWithCurrentSpouseCHVFlagsLength2526]), // Field # 392

		StudentTaxReturnFilingStatusCHVFlags: strings.TrimSpace(record[studentTaxReturnFilingStatusCHVFlagsStartIndex2526-1 : (studentTaxReturnFilingStatusCHVFlagsStartIndex2526-1)+studentTaxReturnFilingStatusCHVFlagsLength2526]), // Field # 393

		StudentIncomeEarnedFromWorkCorrectionCHVFlags: strings.TrimSpace(record[studentIncomeEarnedFromWorkCorrectionCHVFlagsStartIndex2526-1 : (studentIncomeEarnedFromWorkCorrectionCHVFlagsStartIndex2526-1)+studentIncomeEarnedFromWorkCorrectionCHVFlagsLength2526]), // Field # 394

		StudentTaxExemptInterestIncomeCHVFlags: strings.TrimSpace(record[studentTaxExemptInterestIncomeCHVFlagsStartIndex2526-1 : (studentTaxExemptInterestIncomeCHVFlagsStartIndex2526-1)+studentTaxExemptInterestIncomeCHVFlagsLength2526]), // Field # 395

		StudentUntaxedPortionsOfIRADistributionsCHVFlags: strings.TrimSpace(record[studentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2526-1 : (studentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2526-1)+studentUntaxedPortionsOfIRADistributionsCHVFlagsLength2526]), // Field # 396

		StudentIRARolloverCHVFlags: strings.TrimSpace(record[studentIRARolloverCHVFlagsStartIndex2526-1 : (studentIRARolloverCHVFlagsStartIndex2526-1)+studentIRARolloverCHVFlagsLength2526]), // Field # 397

		StudentUntaxedPortionsOfPensionsCHVFlags: strings.TrimSpace(record[studentUntaxedPortionsOfPensionsCHVFlagsStartIndex2526-1 : (studentUntaxedPortionsOfPensionsCHVFlagsStartIndex2526-1)+studentUntaxedPortionsOfPensionsCHVFlagsLength2526]), // Field # 398

		StudentPensionRolloverCHVFlags: strings.TrimSpace(record[studentPensionRolloverCHVFlagsStartIndex2526-1 : (studentPensionRolloverCHVFlagsStartIndex2526-1)+studentPensionRolloverCHVFlagsLength2526]), // Field # 399

		StudentAdjustedGrossIncomeCHVFlags: strings.TrimSpace(record[studentAdjustedGrossIncomeCHVFlagsStartIndex2526-1 : (studentAdjustedGrossIncomeCHVFlagsStartIndex2526-1)+studentAdjustedGrossIncomeCHVFlagsLength2526]), // Field # 400

		StudentIncomeTaxPaidCHVFlags: strings.TrimSpace(record[studentIncomeTaxPaidCHVFlagsStartIndex2526-1 : (studentIncomeTaxPaidCHVFlagsStartIndex2526-1)+studentIncomeTaxPaidCHVFlagsLength2526]), // Field # 401

		StudentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlags: strings.TrimSpace(record[studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex2526-1 : (studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex2526-1)+studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsLength2526]), // Field # 402

		StudentDeductiblePaymentsToIRAKeoghOtherCHVFlags: strings.TrimSpace(record[studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2526-1 : (studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2526-1)+studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2526]), // Field # 403

		StudentEducationCreditsCHVFlags: strings.TrimSpace(record[studentEducationCreditsCHVFlagsStartIndex2526-1 : (studentEducationCreditsCHVFlagsStartIndex2526-1)+studentEducationCreditsCHVFlagsLength2526]), // Field # 404

		StudentFiledScheduleABDEFHCHVFlags: strings.TrimSpace(record[studentFiledScheduleABDEFHCHVFlagsStartIndex2526-1 : (studentFiledScheduleABDEFHCHVFlagsStartIndex2526-1)+studentFiledScheduleABDEFHCHVFlagsLength2526]), // Field # 405

		StudentScheduleCAmountCHVFlags: strings.TrimSpace(record[studentScheduleCAmountCHVFlagsStartIndex2526-1 : (studentScheduleCAmountCHVFlagsStartIndex2526-1)+studentScheduleCAmountCHVFlagsLength2526]), // Field # 406

		StudentCollegeGrantAndScholarshipAidCHVFlags: strings.TrimSpace(record[studentCollegeGrantAndScholarshipAidCHVFlagsStartIndex2526-1 : (studentCollegeGrantAndScholarshipAidCHVFlagsStartIndex2526-1)+studentCollegeGrantAndScholarshipAidCHVFlagsLength2526]), // Field # 407

		StudentForeignEarnedIncomeExclusionCHVFlags: strings.TrimSpace(record[studentForeignEarnedIncomeExclusionCHVFlagsStartIndex2526-1 : (studentForeignEarnedIncomeExclusionCHVFlagsStartIndex2526-1)+studentForeignEarnedIncomeExclusionCHVFlagsLength2526]), // Field # 408

		StudentChildSupportReceivedCHVFlags: strings.TrimSpace(record[studentChildSupportReceivedCHVFlagsStartIndex2526-1 : (studentChildSupportReceivedCHVFlagsStartIndex2526-1)+studentChildSupportReceivedCHVFlagsLength2526]), // Field # 409

		StudentNetWorthOfBusinessesAndInvestmentFarmsCHVFlags: strings.TrimSpace(record[studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex2526-1 : (studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex2526-1)+studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsLength2526]), // Field # 410

		StudentNetWorthOfCurrentInvestmentsCHVFlags: strings.TrimSpace(record[studentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex2526-1 : (studentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex2526-1)+studentNetWorthOfCurrentInvestmentsCHVFlagsLength2526]), // Field # 411

		StudentTotalOfCashSavingsAndCheckingCHVFlags: strings.TrimSpace(record[studentTotalOfCashSavingsAndCheckingCHVFlagsStartIndex2526-1 : (studentTotalOfCashSavingsAndCheckingCHVFlagsStartIndex2526-1)+studentTotalOfCashSavingsAndCheckingCHVFlagsLength2526]), // Field # 412

		StudentCollege1CHVFlags: strings.TrimSpace(record[studentCollege1CHVFlagsStartIndex2526-1 : (studentCollege1CHVFlagsStartIndex2526-1)+studentCollege1CHVFlagsLength2526]), // Field # 413

		StudentCollege2CHVFlags: strings.TrimSpace(record[studentCollege2CHVFlagsStartIndex2526-1 : (studentCollege2CHVFlagsStartIndex2526-1)+studentCollege2CHVFlagsLength2526]), // Field # 414

		StudentCollege3CHVFlags: strings.TrimSpace(record[studentCollege3CHVFlagsStartIndex2526-1 : (studentCollege3CHVFlagsStartIndex2526-1)+studentCollege3CHVFlagsLength2526]), // Field # 415

		StudentCollege4CHVFlags: strings.TrimSpace(record[studentCollege4CHVFlagsStartIndex2526-1 : (studentCollege4CHVFlagsStartIndex2526-1)+studentCollege4CHVFlagsLength2526]), // Field # 416

		StudentCollege5CHVFlags: strings.TrimSpace(record[studentCollege5CHVFlagsStartIndex2526-1 : (studentCollege5CHVFlagsStartIndex2526-1)+studentCollege5CHVFlagsLength2526]), // Field # 417

		StudentCollege6CHVFlags: strings.TrimSpace(record[studentCollege6CHVFlagsStartIndex2526-1 : (studentCollege6CHVFlagsStartIndex2526-1)+studentCollege6CHVFlagsLength2526]), // Field # 418

		StudentCollege7CHVFlags: strings.TrimSpace(record[studentCollege7CHVFlagsStartIndex2526-1 : (studentCollege7CHVFlagsStartIndex2526-1)+studentCollege7CHVFlagsLength2526]), // Field # 419

		StudentCollege8CHVFlags: strings.TrimSpace(record[studentCollege8CHVFlagsStartIndex2526-1 : (studentCollege8CHVFlagsStartIndex2526-1)+studentCollege8CHVFlagsLength2526]), // Field # 420

		StudentCollege9CHVFlags: strings.TrimSpace(record[studentCollege9CHVFlagsStartIndex2526-1 : (studentCollege9CHVFlagsStartIndex2526-1)+studentCollege9CHVFlagsLength2526]), // Field # 421

		StudentCollege10CHVFlags: strings.TrimSpace(record[studentCollege10CHVFlagsStartIndex2526-1 : (studentCollege10CHVFlagsStartIndex2526-1)+studentCollege10CHVFlagsLength2526]), // Field # 422

		StudentCollege11CHVFlags: strings.TrimSpace(record[studentCollege11CHVFlagsStartIndex2526-1 : (studentCollege11CHVFlagsStartIndex2526-1)+studentCollege11CHVFlagsLength2526]), // Field # 423

		StudentCollege12CHVFlags: strings.TrimSpace(record[studentCollege12CHVFlagsStartIndex2526-1 : (studentCollege12CHVFlagsStartIndex2526-1)+studentCollege12CHVFlagsLength2526]), // Field # 424

		StudentCollege13CHVFlags: strings.TrimSpace(record[studentCollege13CHVFlagsStartIndex2526-1 : (studentCollege13CHVFlagsStartIndex2526-1)+studentCollege13CHVFlagsLength2526]), // Field # 425

		StudentCollege14CHVFlags: strings.TrimSpace(record[studentCollege14CHVFlagsStartIndex2526-1 : (studentCollege14CHVFlagsStartIndex2526-1)+studentCollege14CHVFlagsLength2526]), // Field # 426

		StudentCollege15CHVFlags: strings.TrimSpace(record[studentCollege15CHVFlagsStartIndex2526-1 : (studentCollege15CHVFlagsStartIndex2526-1)+studentCollege15CHVFlagsLength2526]), // Field # 427

		StudentCollege16CHVFlags: strings.TrimSpace(record[studentCollege16CHVFlagsStartIndex2526-1 : (studentCollege16CHVFlagsStartIndex2526-1)+studentCollege16CHVFlagsLength2526]), // Field # 428

		StudentCollege17CHVFlags: strings.TrimSpace(record[studentCollege17CHVFlagsStartIndex2526-1 : (studentCollege17CHVFlagsStartIndex2526-1)+studentCollege17CHVFlagsLength2526]), // Field # 429

		StudentCollege18CHVFlags: strings.TrimSpace(record[studentCollege18CHVFlagsStartIndex2526-1 : (studentCollege18CHVFlagsStartIndex2526-1)+studentCollege18CHVFlagsLength2526]), // Field # 430

		StudentCollege19CHVFlags: strings.TrimSpace(record[studentCollege19CHVFlagsStartIndex2526-1 : (studentCollege19CHVFlagsStartIndex2526-1)+studentCollege19CHVFlagsLength2526]), // Field # 431

		StudentCollege20CHVFlags: strings.TrimSpace(record[studentCollege20CHVFlagsStartIndex2526-1 : (studentCollege20CHVFlagsStartIndex2526-1)+studentCollege20CHVFlagsLength2526]), // Field # 432

		StudentConsentToRetrieveAndDiscloseFTICHVFlags: strings.TrimSpace(record[studentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2526-1 : (studentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2526-1)+studentConsentToRetrieveAndDiscloseFTICHVFlagsLength2526]), // Field # 433

		StudentSignatureCHVFlags: strings.TrimSpace(record[studentSignatureCHVFlagsStartIndex2526-1 : (studentSignatureCHVFlagsStartIndex2526-1)+studentSignatureCHVFlagsLength2526]), // Field # 434

		StudentSignatureDateCHVFlags: strings.TrimSpace(record[studentSignatureDateCHVFlagsStartIndex2526-1 : (studentSignatureDateCHVFlagsStartIndex2526-1)+studentSignatureDateCHVFlagsLength2526]), // Field # 435

		StudentSpouseFirstNameCHVFlags: strings.TrimSpace(record[studentSpouseFirstNameCHVFlagsStartIndex2526-1 : (studentSpouseFirstNameCHVFlagsStartIndex2526-1)+studentSpouseFirstNameCHVFlagsLength2526]), // Field # 436

		StudentSpouseMiddleNameCHVFlags: strings.TrimSpace(record[studentSpouseMiddleNameCHVFlagsStartIndex2526-1 : (studentSpouseMiddleNameCHVFlagsStartIndex2526-1)+studentSpouseMiddleNameCHVFlagsLength2526]), // Field # 437

		StudentSpouseLastNameCHVFlags: strings.TrimSpace(record[studentSpouseLastNameCHVFlagsStartIndex2526-1 : (studentSpouseLastNameCHVFlagsStartIndex2526-1)+studentSpouseLastNameCHVFlagsLength2526]), // Field # 438

		StudentSpouseSuffixCHVFlags: strings.TrimSpace(record[studentSpouseSuffixCHVFlagsStartIndex2526-1 : (studentSpouseSuffixCHVFlagsStartIndex2526-1)+studentSpouseSuffixCHVFlagsLength2526]), // Field # 439

		StudentSpouseDateOfBirthCHVFlags: strings.TrimSpace(record[studentSpouseDateOfBirthCHVFlagsStartIndex2526-1 : (studentSpouseDateOfBirthCHVFlagsStartIndex2526-1)+studentSpouseDateOfBirthCHVFlagsLength2526]), // Field # 440

		StudentSpouseSSNCHVFlags: strings.TrimSpace(record[studentSpouseSSNCHVFlagsStartIndex2526-1 : (studentSpouseSSNCHVFlagsStartIndex2526-1)+studentSpouseSSNCHVFlagsLength2526]), // Field # 441

		StudentSpouseITINCHVFlags: strings.TrimSpace(record[studentSpouseITINCHVFlagsStartIndex2526-1 : (studentSpouseITINCHVFlagsStartIndex2526-1)+studentSpouseITINCHVFlagsLength2526]), // Field # 442

		StudentSpousePhoneNumberCHVFlags: strings.TrimSpace(record[studentSpousePhoneNumberCHVFlagsStartIndex2526-1 : (studentSpousePhoneNumberCHVFlagsStartIndex2526-1)+studentSpousePhoneNumberCHVFlagsLength2526]), // Field # 443

		StudentSpouseEmailAddressCHVFlags: strings.TrimSpace(record[studentSpouseEmailAddressCHVFlagsStartIndex2526-1 : (studentSpouseEmailAddressCHVFlagsStartIndex2526-1)+studentSpouseEmailAddressCHVFlagsLength2526]), // Field # 444

		StudentSpouseStreetAddressCHVFlags: strings.TrimSpace(record[studentSpouseStreetAddressCHVFlagsStartIndex2526-1 : (studentSpouseStreetAddressCHVFlagsStartIndex2526-1)+studentSpouseStreetAddressCHVFlagsLength2526]), // Field # 445

		StudentSpouseCityCHVFlags: strings.TrimSpace(record[studentSpouseCityCHVFlagsStartIndex2526-1 : (studentSpouseCityCHVFlagsStartIndex2526-1)+studentSpouseCityCHVFlagsLength2526]), // Field # 446

		StudentSpouseStateCHVFlags: strings.TrimSpace(record[studentSpouseStateCHVFlagsStartIndex2526-1 : (studentSpouseStateCHVFlagsStartIndex2526-1)+studentSpouseStateCHVFlagsLength2526]), // Field # 447

		StudentSpouseZipCodeCHVFlags: strings.TrimSpace(record[studentSpouseZipCodeCHVFlagsStartIndex2526-1 : (studentSpouseZipCodeCHVFlagsStartIndex2526-1)+studentSpouseZipCodeCHVFlagsLength2526]), // Field # 448

		StudentSpouseCountryCHVFlags: strings.TrimSpace(record[studentSpouseCountryCHVFlagsStartIndex2526-1 : (studentSpouseCountryCHVFlagsStartIndex2526-1)+studentSpouseCountryCHVFlagsLength2526]), // Field # 449

		StudentSpouseFiled1040Or1040NRCHVFlags: strings.TrimSpace(record[studentSpouseFiled1040Or1040NRCHVFlagsStartIndex2526-1 : (studentSpouseFiled1040Or1040NRCHVFlagsStartIndex2526-1)+studentSpouseFiled1040Or1040NRCHVFlagsLength2526]), // Field # 450

		StudentSpouseFiledNonUSTaxReturnCHVFlags: strings.TrimSpace(record[studentSpouseFiledNonUSTaxReturnCHVFlagsStartIndex2526-1 : (studentSpouseFiledNonUSTaxReturnCHVFlagsStartIndex2526-1)+studentSpouseFiledNonUSTaxReturnCHVFlagsLength2526]), // Field # 451

		StudentSpouseTaxReturnFilingStatusCHVFlags: strings.TrimSpace(record[studentSpouseTaxReturnFilingStatusCHVFlagsStartIndex2526-1 : (studentSpouseTaxReturnFilingStatusCHVFlagsStartIndex2526-1)+studentSpouseTaxReturnFilingStatusCHVFlagsLength2526]), // Field # 452

		StudentSpouseIncomeEarnedFromWorkCHVFlags: strings.TrimSpace(record[studentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex2526-1 : (studentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex2526-1)+studentSpouseIncomeEarnedFromWorkCHVFlagsLength2526]), // Field # 453

		StudentSpouseTaxExemptInterestIncomeCHVFlags: strings.TrimSpace(record[studentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex2526-1 : (studentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex2526-1)+studentSpouseTaxExemptInterestIncomeCHVFlagsLength2526]), // Field # 454

		StudentSpouseUntaxedPortionsOfIRADistributionsCHVFlags: strings.TrimSpace(record[studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2526-1 : (studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2526-1)+studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsLength2526]), // Field # 455

		StudentSpouseIRARolloverCHVFlags: strings.TrimSpace(record[studentSpouseIRARolloverCHVFlagsStartIndex2526-1 : (studentSpouseIRARolloverCHVFlagsStartIndex2526-1)+studentSpouseIRARolloverCHVFlagsLength2526]), // Field # 456

		StudentSpouseUntaxedPortionsOfPensionsCHVFlags: strings.TrimSpace(record[studentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex2526-1 : (studentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex2526-1)+studentSpouseUntaxedPortionsOfPensionsCHVFlagsLength2526]), // Field # 457

		StudentSpousePensionRolloverCHVFlags: strings.TrimSpace(record[studentSpousePensionRolloverCHVFlagsStartIndex2526-1 : (studentSpousePensionRolloverCHVFlagsStartIndex2526-1)+studentSpousePensionRolloverCHVFlagsLength2526]), // Field # 458

		StudentSpouseAdjustedGrossIncomeCHVFlags: strings.TrimSpace(record[studentSpouseAdjustedGrossIncomeCHVFlagsStartIndex2526-1 : (studentSpouseAdjustedGrossIncomeCHVFlagsStartIndex2526-1)+studentSpouseAdjustedGrossIncomeCHVFlagsLength2526]), // Field # 459

		StudentSpouseIncomeTaxPaidCHVFlags: strings.TrimSpace(record[studentSpouseIncomeTaxPaidCHVFlagsStartIndex2526-1 : (studentSpouseIncomeTaxPaidCHVFlagsStartIndex2526-1)+studentSpouseIncomeTaxPaidCHVFlagsLength2526]), // Field # 460

		StudentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlags: strings.TrimSpace(record[studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2526-1 : (studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2526-1)+studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2526]), // Field # 461

		StudentSpouseEducationCreditsCHVFlags: strings.TrimSpace(record[studentSpouseEducationCreditsCHVFlagsStartIndex2526-1 : (studentSpouseEducationCreditsCHVFlagsStartIndex2526-1)+studentSpouseEducationCreditsCHVFlagsLength2526]), // Field # 462

		StudentSpouseFiledScheduleABDEFHCHVFlags: strings.TrimSpace(record[studentSpouseFiledScheduleABDEFHCHVFlagsStartIndex2526-1 : (studentSpouseFiledScheduleABDEFHCHVFlagsStartIndex2526-1)+studentSpouseFiledScheduleABDEFHCHVFlagsLength2526]), // Field # 463

		StudentSpouseScheduleCAmountCHVFlags: strings.TrimSpace(record[studentSpouseScheduleCAmountCHVFlagsStartIndex2526-1 : (studentSpouseScheduleCAmountCHVFlagsStartIndex2526-1)+studentSpouseScheduleCAmountCHVFlagsLength2526]), // Field # 464

		StudentSpouseForeignEarnedIncomeExclusionCHVFlags: strings.TrimSpace(record[studentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex2526-1 : (studentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex2526-1)+studentSpouseForeignEarnedIncomeExclusionCHVFlagsLength2526]), // Field # 465

		StudentSpouseConsentToRetrieveAndDiscloseFTICHVFlags: strings.TrimSpace(record[studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2526-1 : (studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2526-1)+studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsLength2526]), // Field # 466

		StudentSpouseSignatureCHVFlags: strings.TrimSpace(record[studentSpouseSignatureCHVFlagsStartIndex2526-1 : (studentSpouseSignatureCHVFlagsStartIndex2526-1)+studentSpouseSignatureCHVFlagsLength2526]), // Field # 467

		StudentSpouseSignatureDateCHVFlags: strings.TrimSpace(record[studentSpouseSignatureDateCHVFlagsStartIndex2526-1 : (studentSpouseSignatureDateCHVFlagsStartIndex2526-1)+studentSpouseSignatureDateCHVFlagsLength2526]), // Field # 468

		ParentFirstNameCHVFlags: strings.TrimSpace(record[parentFirstNameCHVFlagsStartIndex2526-1 : (parentFirstNameCHVFlagsStartIndex2526-1)+parentFirstNameCHVFlagsLength2526]), // Field # 469

		ParentMiddleNameCHVFlags: strings.TrimSpace(record[parentMiddleNameCHVFlagsStartIndex2526-1 : (parentMiddleNameCHVFlagsStartIndex2526-1)+parentMiddleNameCHVFlagsLength2526]), // Field # 470

		ParentLastNameCHVFlags: strings.TrimSpace(record[parentLastNameCHVFlagsStartIndex2526-1 : (parentLastNameCHVFlagsStartIndex2526-1)+parentLastNameCHVFlagsLength2526]), // Field # 471

		ParentSuffixCHVFlags: strings.TrimSpace(record[parentSuffixCHVFlagsStartIndex2526-1 : (parentSuffixCHVFlagsStartIndex2526-1)+parentSuffixCHVFlagsLength2526]), // Field # 472

		ParentDateOfBirthCHVFlags: strings.TrimSpace(record[parentDateOfBirthCHVFlagsStartIndex2526-1 : (parentDateOfBirthCHVFlagsStartIndex2526-1)+parentDateOfBirthCHVFlagsLength2526]), // Field # 473

		ParentSSNCHVFlags: strings.TrimSpace(record[parentSSNCHVFlagsStartIndex2526-1 : (parentSSNCHVFlagsStartIndex2526-1)+parentSSNCHVFlagsLength2526]), // Field # 474

		ParentITINCHVFlags: strings.TrimSpace(record[parentITINCHVFlagsStartIndex2526-1 : (parentITINCHVFlagsStartIndex2526-1)+parentITINCHVFlagsLength2526]), // Field # 475

		ParentPhoneNumberCHVFlags: strings.TrimSpace(record[parentPhoneNumberCHVFlagsStartIndex2526-1 : (parentPhoneNumberCHVFlagsStartIndex2526-1)+parentPhoneNumberCHVFlagsLength2526]), // Field # 476

		ParentEmailAddressCHVFlags: strings.TrimSpace(record[parentEmailAddressCHVFlagsStartIndex2526-1 : (parentEmailAddressCHVFlagsStartIndex2526-1)+parentEmailAddressCHVFlagsLength2526]), // Field # 477

		ParentStreetAddressCHVFlags: strings.TrimSpace(record[parentStreetAddressCHVFlagsStartIndex2526-1 : (parentStreetAddressCHVFlagsStartIndex2526-1)+parentStreetAddressCHVFlagsLength2526]), // Field # 478

		ParentCityCHVFlags: strings.TrimSpace(record[parentCityCHVFlagsStartIndex2526-1 : (parentCityCHVFlagsStartIndex2526-1)+parentCityCHVFlagsLength2526]), // Field # 479

		ParentStateCHVFlags: strings.TrimSpace(record[parentStateCHVFlagsStartIndex2526-1 : (parentStateCHVFlagsStartIndex2526-1)+parentStateCHVFlagsLength2526]), // Field # 480

		ParentZipCodeCHVFlags: strings.TrimSpace(record[parentZipCodeCHVFlagsStartIndex2526-1 : (parentZipCodeCHVFlagsStartIndex2526-1)+parentZipCodeCHVFlagsLength2526]), // Field # 481

		ParentCountryCHVFlags: strings.TrimSpace(record[parentCountryCHVFlagsStartIndex2526-1 : (parentCountryCHVFlagsStartIndex2526-1)+parentCountryCHVFlagsLength2526]), // Field # 482

		ParentMaritalStatusCHVFlags: strings.TrimSpace(record[parentMaritalStatusCHVFlagsStartIndex2526-1 : (parentMaritalStatusCHVFlagsStartIndex2526-1)+parentMaritalStatusCHVFlagsLength2526]), // Field # 483

		ParentStateOfLegalResidenceCHVFlags: strings.TrimSpace(record[parentStateOfLegalResidenceCHVFlagsStartIndex2526-1 : (parentStateOfLegalResidenceCHVFlagsStartIndex2526-1)+parentStateOfLegalResidenceCHVFlagsLength2526]), // Field # 484

		ParentLegalResidenceDateCHVFlags: strings.TrimSpace(record[parentLegalResidenceDateCHVFlagsStartIndex2526-1 : (parentLegalResidenceDateCHVFlagsStartIndex2526-1)+parentLegalResidenceDateCHVFlagsLength2526]), // Field # 485

		ParentUpdatedFamilySizeCHVFlags: strings.TrimSpace(record[parentUpdatedFamilySizeCHVFlagsStartIndex2526-1 : (parentUpdatedFamilySizeCHVFlagsStartIndex2526-1)+parentUpdatedFamilySizeCHVFlagsLength2526]), // Field # 486

		ParentNumberInCollegeCHVFlags: strings.TrimSpace(record[parentNumberInCollegeCHVFlagsStartIndex2526-1 : (parentNumberInCollegeCHVFlagsStartIndex2526-1)+parentNumberInCollegeCHVFlagsLength2526]), // Field # 487

		ParentReceivedEITCCHVFlags: strings.TrimSpace(record[parentReceivedEITCCHVFlagsStartIndex2526-1 : (parentReceivedEITCCHVFlagsStartIndex2526-1)+parentReceivedEITCCHVFlagsLength2526]), // Field # 488

		ParentReceivedFederalHousingAssistanceCHVFlags: strings.TrimSpace(record[parentReceivedFederalHousingAssistanceCHVFlagsStartIndex2526-1 : (parentReceivedFederalHousingAssistanceCHVFlagsStartIndex2526-1)+parentReceivedFederalHousingAssistanceCHVFlagsLength2526]), // Field # 489

		ParentReceivedFreeReducedPriceLunchCHVFlags: strings.TrimSpace(record[parentReceivedFreeReducedPriceLunchCHVFlagsStartIndex2526-1 : (parentReceivedFreeReducedPriceLunchCHVFlagsStartIndex2526-1)+parentReceivedFreeReducedPriceLunchCHVFlagsLength2526]), // Field # 490

		ParentReceivedMedicaidCHVFlags: strings.TrimSpace(record[parentReceivedMedicaidCHVFlagsStartIndex2526-1 : (parentReceivedMedicaidCHVFlagsStartIndex2526-1)+parentReceivedMedicaidCHVFlagsLength2526]), // Field # 491

		ParentReceivedRefundableCreditFor36BHealthPlanCHVFlags: strings.TrimSpace(record[parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex2526-1 : (parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex2526-1)+parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsLength2526]), // Field # 492

		ParentReceivedSNAPCHVFlags: strings.TrimSpace(record[parentReceivedSNAPCHVFlagsStartIndex2526-1 : (parentReceivedSNAPCHVFlagsStartIndex2526-1)+parentReceivedSNAPCHVFlagsLength2526]), // Field # 493

		ParentReceivedSupplementalSecurityIncomeCHVFlags: strings.TrimSpace(record[parentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex2526-1 : (parentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex2526-1)+parentReceivedSupplementalSecurityIncomeCHVFlagsLength2526]), // Field # 494

		ParentReceivedTANFCHVFlags: strings.TrimSpace(record[parentReceivedTANFCHVFlagsStartIndex2526-1 : (parentReceivedTANFCHVFlagsStartIndex2526-1)+parentReceivedTANFCHVFlagsLength2526]), // Field # 495

		ParentReceivedWICCHVFlags: strings.TrimSpace(record[parentReceivedWICCHVFlagsStartIndex2526-1 : (parentReceivedWICCHVFlagsStartIndex2526-1)+parentReceivedWICCHVFlagsLength2526]), // Field # 496

		ParentFederalBenefitsNoneOfTheAboveCHVFlags: strings.TrimSpace(record[parentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex2526-1 : (parentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex2526-1)+parentFederalBenefitsNoneOfTheAboveCHVFlagsLength2526]), // Field # 497

		ParentFiled1040Or1040NRCHVFlags: strings.TrimSpace(record[parentFiled1040Or1040NRCHVFlagsStartIndex2526-1 : (parentFiled1040Or1040NRCHVFlagsStartIndex2526-1)+parentFiled1040Or1040NRCHVFlagsLength2526]), // Field # 498

		ParentFileNonUSTaxReturnCHVFlags: strings.TrimSpace(record[parentFileNonUSTaxReturnCHVFlagsStartIndex2526-1 : (parentFileNonUSTaxReturnCHVFlagsStartIndex2526-1)+parentFileNonUSTaxReturnCHVFlagsLength2526]), // Field # 499

		ParentFiledJointReturnWithCurrentSpouseCHVFlags: strings.TrimSpace(record[parentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex2526-1 : (parentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex2526-1)+parentFiledJointReturnWithCurrentSpouseCHVFlagsLength2526]), // Field # 500

		ParentTaxReturnFilingStatusCHVFlags: strings.TrimSpace(record[parentTaxReturnFilingStatusCHVFlagsStartIndex2526-1 : (parentTaxReturnFilingStatusCHVFlagsStartIndex2526-1)+parentTaxReturnFilingStatusCHVFlagsLength2526]), // Field # 501

		ParentIncomeEarnedFromWorkCHVFlags: strings.TrimSpace(record[parentIncomeEarnedFromWorkCHVFlagsStartIndex2526-1 : (parentIncomeEarnedFromWorkCHVFlagsStartIndex2526-1)+parentIncomeEarnedFromWorkCHVFlagsLength2526]), // Field # 502

		ParentTaxExemptInterestIncomeCHVFlags: strings.TrimSpace(record[parentTaxExemptInterestIncomeCHVFlagsStartIndex2526-1 : (parentTaxExemptInterestIncomeCHVFlagsStartIndex2526-1)+parentTaxExemptInterestIncomeCHVFlagsLength2526]), // Field # 503

		ParentUntaxedPortionsOfIRADistributionsCHVFlags: strings.TrimSpace(record[parentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2526-1 : (parentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2526-1)+parentUntaxedPortionsOfIRADistributionsCHVFlagsLength2526]), // Field # 504

		ParentIRARolloverCHVFlags: strings.TrimSpace(record[parentIRARolloverCHVFlagsStartIndex2526-1 : (parentIRARolloverCHVFlagsStartIndex2526-1)+parentIRARolloverCHVFlagsLength2526]), // Field # 505

		ParentUntaxedPortionsOfPensionsCHVFlags: strings.TrimSpace(record[parentUntaxedPortionsOfPensionsCHVFlagsStartIndex2526-1 : (parentUntaxedPortionsOfPensionsCHVFlagsStartIndex2526-1)+parentUntaxedPortionsOfPensionsCHVFlagsLength2526]), // Field # 506

		ParentPensionRolloverCHVFlags: strings.TrimSpace(record[parentPensionRolloverCHVFlagsStartIndex2526-1 : (parentPensionRolloverCHVFlagsStartIndex2526-1)+parentPensionRolloverCHVFlagsLength2526]), // Field # 507

		ParentAdjustedGrossIncomeCHVFlags: strings.TrimSpace(record[parentAdjustedGrossIncomeCHVFlagsStartIndex2526-1 : (parentAdjustedGrossIncomeCHVFlagsStartIndex2526-1)+parentAdjustedGrossIncomeCHVFlagsLength2526]), // Field # 508

		ParentIncomeTaxPaidCHVFlags: strings.TrimSpace(record[parentIncomeTaxPaidCHVFlagsStartIndex2526-1 : (parentIncomeTaxPaidCHVFlagsStartIndex2526-1)+parentIncomeTaxPaidCHVFlagsLength2526]), // Field # 509

		ParentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlags: strings.TrimSpace(record[parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex2526-1 : (parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex2526-1)+parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsLength2526]), // Field # 510

		ParentDeductiblePaymentsToIRAKeoghOtherCHVFlags: strings.TrimSpace(record[parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2526-1 : (parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2526-1)+parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2526]), // Field # 511

		ParentEducationCreditsCHVFlags: strings.TrimSpace(record[parentEducationCreditsCHVFlagsStartIndex2526-1 : (parentEducationCreditsCHVFlagsStartIndex2526-1)+parentEducationCreditsCHVFlagsLength2526]), // Field # 512

		ParentFiledScheduleABDEFHCHVFlags: strings.TrimSpace(record[parentFiledScheduleABDEFHCHVFlagsStartIndex2526-1 : (parentFiledScheduleABDEFHCHVFlagsStartIndex2526-1)+parentFiledScheduleABDEFHCHVFlagsLength2526]), // Field # 513

		ParentScheduleCAmountCHVFlags: strings.TrimSpace(record[parentScheduleCAmountCHVFlagsStartIndex2526-1 : (parentScheduleCAmountCHVFlagsStartIndex2526-1)+parentScheduleCAmountCHVFlagsLength2526]), // Field # 514

		ParentCollegeGrantAndScholarshipAidCHVFlags: strings.TrimSpace(record[parentCollegeGrantAndScholarshipAidCHVFlagsStartIndex2526-1 : (parentCollegeGrantAndScholarshipAidCHVFlagsStartIndex2526-1)+parentCollegeGrantAndScholarshipAidCHVFlagsLength2526]), // Field # 515

		ParentForeignEarnedIncomeExclusionCHVFlags: strings.TrimSpace(record[parentForeignEarnedIncomeExclusionCHVFlagsStartIndex2526-1 : (parentForeignEarnedIncomeExclusionCHVFlagsStartIndex2526-1)+parentForeignEarnedIncomeExclusionCHVFlagsLength2526]), // Field # 516

		ParentChildSupportReceivedCHVFlags: strings.TrimSpace(record[parentChildSupportReceivedCHVFlagsStartIndex2526-1 : (parentChildSupportReceivedCHVFlagsStartIndex2526-1)+parentChildSupportReceivedCHVFlagsLength2526]), // Field # 517

		ParentNetWorthOfCurrentInvestmentsCHVFlags: strings.TrimSpace(record[parentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex2526-1 : (parentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex2526-1)+parentNetWorthOfCurrentInvestmentsCHVFlagsLength2526]), // Field # 518

		ParentTotalOfCashSavingsAndCheckingAccountsCHVFlags: strings.TrimSpace(record[parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsStartIndex2526-1 : (parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsStartIndex2526-1)+parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsLength2526]), // Field # 519

		ParentNetWorthOfBusinessesAndInvestmentFarmsCHVFlags: strings.TrimSpace(record[parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex2526-1 : (parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex2526-1)+parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsLength2526]), // Field # 520

		ParentConsentToRetrieveAndDiscloseFTICHVFlags: strings.TrimSpace(record[parentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2526-1 : (parentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2526-1)+parentConsentToRetrieveAndDiscloseFTICHVFlagsLength2526]), // Field # 521

		ParentSignatureCHVFlags: strings.TrimSpace(record[parentSignatureCHVFlagsStartIndex2526-1 : (parentSignatureCHVFlagsStartIndex2526-1)+parentSignatureCHVFlagsLength2526]), // Field # 522

		ParentSignatureDateCHVFlags: strings.TrimSpace(record[parentSignatureDateCHVFlagsStartIndex2526-1 : (parentSignatureDateCHVFlagsStartIndex2526-1)+parentSignatureDateCHVFlagsLength2526]), // Field # 523

		ParentSpouseFirstNameCHVFlags: strings.TrimSpace(record[parentSpouseFirstNameCHVFlagsStartIndex2526-1 : (parentSpouseFirstNameCHVFlagsStartIndex2526-1)+parentSpouseFirstNameCHVFlagsLength2526]), // Field # 524

		ParentSpouseMiddleNameCHVFlags: strings.TrimSpace(record[parentSpouseMiddleNameCHVFlagsStartIndex2526-1 : (parentSpouseMiddleNameCHVFlagsStartIndex2526-1)+parentSpouseMiddleNameCHVFlagsLength2526]), // Field # 525

		ParentSpouseLastNameCHVFlags: strings.TrimSpace(record[parentSpouseLastNameCHVFlagsStartIndex2526-1 : (parentSpouseLastNameCHVFlagsStartIndex2526-1)+parentSpouseLastNameCHVFlagsLength2526]), // Field # 526

		ParentSpouseSuffixCHVFlags: strings.TrimSpace(record[parentSpouseSuffixCHVFlagsStartIndex2526-1 : (parentSpouseSuffixCHVFlagsStartIndex2526-1)+parentSpouseSuffixCHVFlagsLength2526]), // Field # 527

		ParentSpouseDateOfBirthCHVFlags: strings.TrimSpace(record[parentSpouseDateOfBirthCHVFlagsStartIndex2526-1 : (parentSpouseDateOfBirthCHVFlagsStartIndex2526-1)+parentSpouseDateOfBirthCHVFlagsLength2526]), // Field # 528

		ParentSpouseSSNCHVFlags: strings.TrimSpace(record[parentSpouseSSNCHVFlagsStartIndex2526-1 : (parentSpouseSSNCHVFlagsStartIndex2526-1)+parentSpouseSSNCHVFlagsLength2526]), // Field # 529

		ParentSpouseITINCHVFlags: strings.TrimSpace(record[parentSpouseITINCHVFlagsStartIndex2526-1 : (parentSpouseITINCHVFlagsStartIndex2526-1)+parentSpouseITINCHVFlagsLength2526]), // Field # 530

		ParentSpousePhoneNumberCHVFlags: strings.TrimSpace(record[parentSpousePhoneNumberCHVFlagsStartIndex2526-1 : (parentSpousePhoneNumberCHVFlagsStartIndex2526-1)+parentSpousePhoneNumberCHVFlagsLength2526]), // Field # 531

		ParentSpouseEmailAddressCHVFlags: strings.TrimSpace(record[parentSpouseEmailAddressCHVFlagsStartIndex2526-1 : (parentSpouseEmailAddressCHVFlagsStartIndex2526-1)+parentSpouseEmailAddressCHVFlagsLength2526]), // Field # 532

		ParentSpouseStreetAddressCHVFlags: strings.TrimSpace(record[parentSpouseStreetAddressCHVFlagsStartIndex2526-1 : (parentSpouseStreetAddressCHVFlagsStartIndex2526-1)+parentSpouseStreetAddressCHVFlagsLength2526]), // Field # 533

		ParentSpouseCityCHVFlags: strings.TrimSpace(record[parentSpouseCityCHVFlagsStartIndex2526-1 : (parentSpouseCityCHVFlagsStartIndex2526-1)+parentSpouseCityCHVFlagsLength2526]), // Field # 534

		ParentSpouseStateCHVFlags: strings.TrimSpace(record[parentSpouseStateCHVFlagsStartIndex2526-1 : (parentSpouseStateCHVFlagsStartIndex2526-1)+parentSpouseStateCHVFlagsLength2526]), // Field # 535

		ParentSpouseZipCodeCHVFlags: strings.TrimSpace(record[parentSpouseZipCodeCHVFlagsStartIndex2526-1 : (parentSpouseZipCodeCHVFlagsStartIndex2526-1)+parentSpouseZipCodeCHVFlagsLength2526]), // Field # 536

		ParentSpouseCountryCHVFlags: strings.TrimSpace(record[parentSpouseCountryCHVFlagsStartIndex2526-1 : (parentSpouseCountryCHVFlagsStartIndex2526-1)+parentSpouseCountryCHVFlagsLength2526]), // Field # 537

		ParentSpouseFiled1040Or1040NRCHVFlags: strings.TrimSpace(record[parentSpouseFiled1040Or1040NRCHVFlagsStartIndex2526-1 : (parentSpouseFiled1040Or1040NRCHVFlagsStartIndex2526-1)+parentSpouseFiled1040Or1040NRCHVFlagsLength2526]), // Field # 538

		ParentSpouseFileNonUSTaxReturnCHVFlags: strings.TrimSpace(record[parentSpouseFileNonUSTaxReturnCHVFlagsStartIndex2526-1 : (parentSpouseFileNonUSTaxReturnCHVFlagsStartIndex2526-1)+parentSpouseFileNonUSTaxReturnCHVFlagsLength2526]), // Field # 539

		ParentSpouseTaxReturnFilingStatusCHVFlags: strings.TrimSpace(record[parentSpouseTaxReturnFilingStatusCHVFlagsStartIndex2526-1 : (parentSpouseTaxReturnFilingStatusCHVFlagsStartIndex2526-1)+parentSpouseTaxReturnFilingStatusCHVFlagsLength2526]), // Field # 540

		ParentSpouseIncomeEarnedFromWorkCHVFlags: strings.TrimSpace(record[parentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex2526-1 : (parentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex2526-1)+parentSpouseIncomeEarnedFromWorkCHVFlagsLength2526]), // Field # 541

		ParentSpouseTaxExemptInterestIncomeCHVFlags: strings.TrimSpace(record[parentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex2526-1 : (parentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex2526-1)+parentSpouseTaxExemptInterestIncomeCHVFlagsLength2526]), // Field # 542

		ParentSpouseUntaxedPortionsOfIRADistributionsCHVFlags: strings.TrimSpace(record[parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2526-1 : (parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex2526-1)+parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsLength2526]), // Field # 543

		ParentSpouseIRARolloverCHVFlags: strings.TrimSpace(record[parentSpouseIRARolloverCHVFlagsStartIndex2526-1 : (parentSpouseIRARolloverCHVFlagsStartIndex2526-1)+parentSpouseIRARolloverCHVFlagsLength2526]), // Field # 544

		ParentSpouseUntaxedPortionsOfPensionsCHVFlags: strings.TrimSpace(record[parentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex2526-1 : (parentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex2526-1)+parentSpouseUntaxedPortionsOfPensionsCHVFlagsLength2526]), // Field # 545

		ParentSpousePensionRolloverCHVFlags: strings.TrimSpace(record[parentSpousePensionRolloverCHVFlagsStartIndex2526-1 : (parentSpousePensionRolloverCHVFlagsStartIndex2526-1)+parentSpousePensionRolloverCHVFlagsLength2526]), // Field # 546

		ParentSpouseAdjustedGrossIncomeCHVFlags: strings.TrimSpace(record[parentSpouseAdjustedGrossIncomeCHVFlagsStartIndex2526-1 : (parentSpouseAdjustedGrossIncomeCHVFlagsStartIndex2526-1)+parentSpouseAdjustedGrossIncomeCHVFlagsLength2526]), // Field # 547

		ParentSpouseIncomeTaxPaidCHVFlags: strings.TrimSpace(record[parentSpouseIncomeTaxPaidCHVFlagsStartIndex2526-1 : (parentSpouseIncomeTaxPaidCHVFlagsStartIndex2526-1)+parentSpouseIncomeTaxPaidCHVFlagsLength2526]), // Field # 548

		ParentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlags: strings.TrimSpace(record[parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2526-1 : (parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex2526-1)+parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength2526]), // Field # 549

		ParentSpouseEducationCreditsCHVFlags: strings.TrimSpace(record[parentSpouseEducationCreditsCHVFlagsStartIndex2526-1 : (parentSpouseEducationCreditsCHVFlagsStartIndex2526-1)+parentSpouseEducationCreditsCHVFlagsLength2526]), // Field # 550

		ParentSpouseFiledScheduleABDEFHCHVFlags: strings.TrimSpace(record[parentSpouseFiledScheduleABDEFHCHVFlagsStartIndex2526-1 : (parentSpouseFiledScheduleABDEFHCHVFlagsStartIndex2526-1)+parentSpouseFiledScheduleABDEFHCHVFlagsLength2526]), // Field # 551

		ParentSpouseScheduleCAmountCHVFlags: strings.TrimSpace(record[parentSpouseScheduleCAmountCHVFlagsStartIndex2526-1 : (parentSpouseScheduleCAmountCHVFlagsStartIndex2526-1)+parentSpouseScheduleCAmountCHVFlagsLength2526]), // Field # 552

		ParentSpouseForeignEarnedIncomeExclusionCHVFlags: strings.TrimSpace(record[parentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex2526-1 : (parentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex2526-1)+parentSpouseForeignEarnedIncomeExclusionCHVFlagsLength2526]), // Field # 553

		ParentSpouseConsentToRetrieveAndDiscloseFTICHVFlags: strings.TrimSpace(record[parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2526-1 : (parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex2526-1)+parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsLength2526]), // Field # 554

		ParentSpouseSignatureCHVFlags: strings.TrimSpace(record[parentSpouseSignatureCHVFlagsStartIndex2526-1 : (parentSpouseSignatureCHVFlagsStartIndex2526-1)+parentSpouseSignatureCHVFlagsLength2526]), // Field # 555

		ParentSpouseSignatureDateCHVFlags: strings.TrimSpace(record[parentSpouseSignatureDateCHVFlagsStartIndex2526-1 : (parentSpouseSignatureDateCHVFlagsStartIndex2526-1)+parentSpouseSignatureDateCHVFlagsLength2526]), // Field # 556

		DHSPrimaryMatchStatus: strings.TrimSpace(record[dHSPrimaryMatchStatusStartIndex2526-1 : (dHSPrimaryMatchStatusStartIndex2526-1)+dHSPrimaryMatchStatusLength2526]), // Field # 557

		DHSCaseNumber: strings.TrimSpace(record[dHSCaseNumberStartIndex2526-1 : (dHSCaseNumberStartIndex2526-1)+dHSCaseNumberLength2526]), // Field # 559

		NSLDSMatchStatus: strings.TrimSpace(record[nsldsMatchStatusStartIndex2526-1 : (nsldsMatchStatusStartIndex2526-1)+nsldsMatchStatusLength2526]), // Field # 560

		NSLDSPostscreeningReasonCode: strings.TrimSpace(record[nsldsPostscreeningReasonCodeStartIndex2526-1 : (nsldsPostscreeningReasonCodeStartIndex2526-1)+nsldsPostscreeningReasonCodeLength2526]), // Field # 561

		StudentSSACitizenshipFlagResults: strings.TrimSpace(record[studentSSACitizenshipFlagResultsStartIndex2526-1 : (studentSSACitizenshipFlagResultsStartIndex2526-1)+studentSSACitizenshipFlagResultsLength2526]), // Field # 562

		StudentSSAMatchStatus: strings.TrimSpace(record[studentSSAMatchStatusStartIndex2526-1 : (studentSSAMatchStatusStartIndex2526-1)+studentSSAMatchStatusLength2526]), // Field # 563

		StudentSpouseSSAMatchStatus: strings.TrimSpace(record[studentSpouseSSAMatchStatusStartIndex2526-1 : (studentSpouseSSAMatchStatusStartIndex2526-1)+studentSpouseSSAMatchStatusLength2526]), // Field # 564

		ParentSSAMatchStatus: strings.TrimSpace(record[parentSSAMatchStatusStartIndex2526-1 : (parentSSAMatchStatusStartIndex2526-1)+parentSSAMatchStatusLength2526]), // Field # 565

		ParentSpouseOrPartnerSSAMatchStatus: strings.TrimSpace(record[parentSpouseOrPartnerSSAMatchStatusStartIndex2526-1 : (parentSpouseOrPartnerSSAMatchStatusStartIndex2526-1)+parentSpouseOrPartnerSSAMatchStatusLength2526]), // Field # 566

		VAMatchFlag: strings.TrimSpace(record[vAMatchFlagStartIndex2526-1 : (vAMatchFlagStartIndex2526-1)+vAMatchFlagLength2526]), // Field # 567

		CommentCodes: strings.TrimSpace(record[commentCodesStartIndex2526-1 : (commentCodesStartIndex2526-1)+commentCodesLength2526]), // Field # 568

		DrugAbuseHoldIndicator: strings.TrimSpace(record[drugAbuseHoldIndicatorStartIndex2526-1 : (drugAbuseHoldIndicatorStartIndex2526-1)+drugAbuseHoldIndicatorLength2526]), // Field # 569

		GraduateFlag: strings.TrimSpace(record[graduateFlagStartIndex2526-1 : (graduateFlagStartIndex2526-1)+graduateFlagLength2526]), // Field # 570

		PellGrantEligibilityFlag: strings.TrimSpace(record[pellGrantEligibilityFlagStartIndex2526-1 : (pellGrantEligibilityFlagStartIndex2526-1)+pellGrantEligibilityFlagLength2526]), // Field # 571

		ReprocessedReasonCode: strings.TrimSpace(record[reprocessedReasonCodeStartIndex2526-1 : (reprocessedReasonCodeStartIndex2526-1)+reprocessedReasonCodeLength2526]), // Field # 572

		FPSCFlag: strings.TrimSpace(record[fpsCFlagStartIndex2526-1 : (fpsCFlagStartIndex2526-1)+fpsCFlagLength2526]), // Field # 573

		FPSCChangeFlag: strings.TrimSpace(record[fpsCChangeFlagStartIndex2526-1 : (fpsCChangeFlagStartIndex2526-1)+fpsCChangeFlagLength2526]), // Field # 574

		ElectronicFederalSchoolCodeIndicator: strings.TrimSpace(record[electronicFederalSchoolCodeIndicatorStartIndex2526-1 : (electronicFederalSchoolCodeIndicatorStartIndex2526-1)+electronicFederalSchoolCodeIndicatorLength2526]), // Field # 575

		RejectReasonCodes: strings.TrimSpace(record[rejectReasonCodesStartIndex2526-1 : (rejectReasonCodesStartIndex2526-1)+rejectReasonCodesLength2526]), // Field # 576

		ElectronicTransactionIndicatorFlag: strings.TrimSpace(record[electronicTransactionIndicatorFlagStartIndex2526-1 : (electronicTransactionIndicatorFlagStartIndex2526-1)+electronicTransactionIndicatorFlagLength2526]), // Field # 577

		StudentLastNameSSNChangeFlag: strings.TrimSpace(record[studentLastNameSSNChangeFlagStartIndex2526-1 : (studentLastNameSSNChangeFlagStartIndex2526-1)+studentLastNameSSNChangeFlagLength2526]), // Field # 578

		HighSchoolCode: strings.TrimSpace(record[highSchoolCodeStartIndex2526-1 : (highSchoolCodeStartIndex2526-1)+highSchoolCodeLength2526]), // Field # 579

		VerificationSelectionChangeFlag: strings.TrimSpace(record[verificationSelectionChangeFlagStartIndex2526-1 : (verificationSelectionChangeFlagStartIndex2526-1)+verificationSelectionChangeFlagLength2526]), // Field # 580

		UseUserProvidedDataOnly: strings.TrimSpace(record[useUserProvidedDataOnlyStartIndex2526-1 : (useUserProvidedDataOnlyStartIndex2526-1)+useUserProvidedDataOnlyLength2526]), // Field # 581

		NSLDSPellOverpaymentFlag: strings.TrimSpace(record[nsldsPellOverpaymentFlagStartIndex2526-1 : (nsldsPellOverpaymentFlagStartIndex2526-1)+nsldsPellOverpaymentFlagLength2526]), // Field # 583

		NSLDSPellOverpaymentContact: strings.TrimSpace(record[nsldsPellOverpaymentContactStartIndex2526-1 : (nsldsPellOverpaymentContactStartIndex2526-1)+nsldsPellOverpaymentContactLength2526]), // Field # 584

		NSLDSFSEOGOverpaymentFlag: strings.TrimSpace(record[nsldsFSEOGOverpaymentFlagStartIndex2526-1 : (nsldsFSEOGOverpaymentFlagStartIndex2526-1)+nsldsFSEOGOverpaymentFlagLength2526]), // Field # 585

		NSLDSFSEOGOverpaymentContact: strings.TrimSpace(record[nsldsFSEOGOverpaymentContactStartIndex2526-1 : (nsldsFSEOGOverpaymentContactStartIndex2526-1)+nsldsFSEOGOverpaymentContactLength2526]), // Field # 586

		NSLDSPerkinsOverpaymentFlag: strings.TrimSpace(record[nsldsPerkinsOverpaymentFlagStartIndex2526-1 : (nsldsPerkinsOverpaymentFlagStartIndex2526-1)+nsldsPerkinsOverpaymentFlagLength2526]), // Field # 587

		NSLDSPerkinsOverpaymentContact: strings.TrimSpace(record[nsldsPerkinsOverpaymentContactStartIndex2526-1 : (nsldsPerkinsOverpaymentContactStartIndex2526-1)+nsldsPerkinsOverpaymentContactLength2526]), // Field # 588

		NSLDSTEACHGrantOverpaymentFlag: strings.TrimSpace(record[nsldsTEACHGrantOverpaymentFlagStartIndex2526-1 : (nsldsTEACHGrantOverpaymentFlagStartIndex2526-1)+nsldsTEACHGrantOverpaymentFlagLength2526]), // Field # 589

		NSLDSTEACHGrantOverpaymentContact: strings.TrimSpace(record[nsldsTEACHGrantOverpaymentContactStartIndex2526-1 : (nsldsTEACHGrantOverpaymentContactStartIndex2526-1)+nsldsTEACHGrantOverpaymentContactLength2526]), // Field # 590

		NSLDSIraqAndAfghanistanServiceGrantOverpaymentFlag: strings.TrimSpace(record[nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagStartIndex2526-1 : (nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagStartIndex2526-1)+nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagLength2526]), // Field # 591

		NSLDSIraqAndAfghanistanServiceGrantOverpaymentContact: strings.TrimSpace(record[nsldsIraqAndAfghanistanServiceGrantOverpaymentContactStartIndex2526-1 : (nsldsIraqAndAfghanistanServiceGrantOverpaymentContactStartIndex2526-1)+nsldsIraqAndAfghanistanServiceGrantOverpaymentContactLength2526]), // Field # 592

		NSLDSDefaultedLoanFlag: strings.TrimSpace(record[nsldsDefaultedLoanFlagStartIndex2526-1 : (nsldsDefaultedLoanFlagStartIndex2526-1)+nsldsDefaultedLoanFlagLength2526]), // Field # 593

		NSLDSDischargedLoanFlag: strings.TrimSpace(record[nsldsDischargedLoanFlagStartIndex2526-1 : (nsldsDischargedLoanFlagStartIndex2526-1)+nsldsDischargedLoanFlagLength2526]), // Field # 594

		NSLDSFraudLoanFlag: strings.TrimSpace(record[nsldsFraudLoanFlagStartIndex2526-1 : (nsldsFraudLoanFlagStartIndex2526-1)+nsldsFraudLoanFlagLength2526]), // Field # 595

		NSLDSSatisfactoryArrangementsFlag: strings.TrimSpace(record[nsldsSatisfactoryArrangementsFlagStartIndex2526-1 : (nsldsSatisfactoryArrangementsFlagStartIndex2526-1)+nsldsSatisfactoryArrangementsFlagLength2526]), // Field # 596

		NSLDSActiveBankruptcyFlag: strings.TrimSpace(record[nsldsActiveBankruptcyFlagStartIndex2526-1 : (nsldsActiveBankruptcyFlagStartIndex2526-1)+nsldsActiveBankruptcyFlagLength2526]), // Field # 597

		NSLDSTEACHGrantConvertedToLoanFlag: strings.TrimSpace(record[nsldsTEACHGrantConvertedToLoanFlagStartIndex2526-1 : (nsldsTEACHGrantConvertedToLoanFlagStartIndex2526-1)+nsldsTEACHGrantConvertedToLoanFlagLength2526]), // Field # 598

		NSLDSAggregateSubsidizedOutstandingPrincipalBalance: strings.TrimSpace(record[nsldsAggregateSubsidizedOutstandingPrincipalBalanceStartIndex2526-1 : (nsldsAggregateSubsidizedOutstandingPrincipalBalanceStartIndex2526-1)+nsldsAggregateSubsidizedOutstandingPrincipalBalanceLength2526]), // Field # 599

		NSLDSAggregateUnsubsidizedOutstandingPrincipalBalance: strings.TrimSpace(record[nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceStartIndex2526-1 : (nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceStartIndex2526-1)+nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceLength2526]), // Field # 600

		NSLDSAggregateCombinedOutstandingPrincipalBalance: strings.TrimSpace(record[nsldsAggregateCombinedOutstandingPrincipalBalanceStartIndex2526-1 : (nsldsAggregateCombinedOutstandingPrincipalBalanceStartIndex2526-1)+nsldsAggregateCombinedOutstandingPrincipalBalanceLength2526]), // Field # 601

		NSLDSAggregateUnallocatedConsolidatedOutstandingPrincipalBalance: strings.TrimSpace(record[nsldsAggregateUnallocatedConsolidatedOutstandingPrincipalBalanceStartIndex2526-1 : (nsldsAggregateUnallocatedConsolidatedOutstandingPrincipalBalanceStartIndex2526-1)+nsldsAggregateUnallocatedConsolidatedOutstandingPrincipalBalanceLength2526]), // Field # 602

		NSLDSAggregateTEACHLoanPrincipalBalance: strings.TrimSpace(record[nsldsAggregateTEACHLoanPrincipalBalanceStartIndex2526-1 : (nsldsAggregateTEACHLoanPrincipalBalanceStartIndex2526-1)+nsldsAggregateTEACHLoanPrincipalBalanceLength2526]), // Field # 603

		NSLDSAggregateSubsidizedPendingDisbursement: strings.TrimSpace(record[nsldsAggregateSubsidizedPendingDisbursementStartIndex2526-1 : (nsldsAggregateSubsidizedPendingDisbursementStartIndex2526-1)+nsldsAggregateSubsidizedPendingDisbursementLength2526]), // Field # 604

		NSLDSAggregateUnsubsidizedPendingDisbursement: strings.TrimSpace(record[nsldsAggregateUnsubsidizedPendingDisbursementStartIndex2526-1 : (nsldsAggregateUnsubsidizedPendingDisbursementStartIndex2526-1)+nsldsAggregateUnsubsidizedPendingDisbursementLength2526]), // Field # 605

		NSLDSAggregateCombinedPendingDisbursement: strings.TrimSpace(record[nsldsAggregateCombinedPendingDisbursementStartIndex2526-1 : (nsldsAggregateCombinedPendingDisbursementStartIndex2526-1)+nsldsAggregateCombinedPendingDisbursementLength2526]), // Field # 606

		NSLDSAggregateSubsidizedTotal: strings.TrimSpace(record[nsldsAggregateSubsidizedTotalStartIndex2526-1 : (nsldsAggregateSubsidizedTotalStartIndex2526-1)+nsldsAggregateSubsidizedTotalLength2526]), // Field # 607

		NSLDSAggregateUnsubsidizedTotal: strings.TrimSpace(record[nsldsAggregateUnsubsidizedTotalStartIndex2526-1 : (nsldsAggregateUnsubsidizedTotalStartIndex2526-1)+nsldsAggregateUnsubsidizedTotalLength2526]), // Field # 608

		NSLDSAggregateCombinedTotal: strings.TrimSpace(record[nsldsAggregateCombinedTotalStartIndex2526-1 : (nsldsAggregateCombinedTotalStartIndex2526-1)+nsldsAggregateCombinedTotalLength2526]), // Field # 609

		NSLDSUnallocatedConsolidatedTotal: strings.TrimSpace(record[nsldsUnallocatedConsolidatedTotalStartIndex2526-1 : (nsldsUnallocatedConsolidatedTotalStartIndex2526-1)+nsldsUnallocatedConsolidatedTotalLength2526]), // Field # 610

		NSLDSTEACHLoanTotal: strings.TrimSpace(record[nsldsTEACHLoanTotalStartIndex2526-1 : (nsldsTEACHLoanTotalStartIndex2526-1)+nsldsTEACHLoanTotalLength2526]), // Field # 611

		NSLDSPerkinsTotalDisbursements: strings.TrimSpace(record[nsldsPerkinsTotalDisbursementsStartIndex2526-1 : (nsldsPerkinsTotalDisbursementsStartIndex2526-1)+nsldsPerkinsTotalDisbursementsLength2526]), // Field # 612

		NSLDSPerkinsCurrentYearDisbursementAmount: strings.TrimSpace(record[nsldsPerkinsCurrentYearDisbursementAmountStartIndex2526-1 : (nsldsPerkinsCurrentYearDisbursementAmountStartIndex2526-1)+nsldsPerkinsCurrentYearDisbursementAmountLength2526]), // Field # 613

		NSLDSAggregateTEACHGrantUndergraduateDisbursedTotal: strings.TrimSpace(record[nsldsAggregateTEACHGrantUndergraduateDisbursedTotalStartIndex2526-1 : (nsldsAggregateTEACHGrantUndergraduateDisbursedTotalStartIndex2526-1)+nsldsAggregateTEACHGrantUndergraduateDisbursedTotalLength2526]), // Field # 614

		NSLDSAggregateTEACHGraduateDisbursementAmount: strings.TrimSpace(record[nsldsAggregateTEACHGraduateDisbursementAmountStartIndex2526-1 : (nsldsAggregateTEACHGraduateDisbursementAmountStartIndex2526-1)+nsldsAggregateTEACHGraduateDisbursementAmountLength2526]), // Field # 615

		NSLDSDefaultedLoanChangeFlag: strings.TrimSpace(record[nsldsDefaultedLoanChangeFlagStartIndex2526-1 : (nsldsDefaultedLoanChangeFlagStartIndex2526-1)+nsldsDefaultedLoanChangeFlagLength2526]), // Field # 616

		NSLDSFraudLoanChangeFlag: strings.TrimSpace(record[nsldsFraudLoanChangeFlagStartIndex2526-1 : (nsldsFraudLoanChangeFlagStartIndex2526-1)+nsldsFraudLoanChangeFlagLength2526]), // Field # 617

		NSLDSDischargedLoanChangeFlag: strings.TrimSpace(record[nsldsDischargedLoanChangeFlagStartIndex2526-1 : (nsldsDischargedLoanChangeFlagStartIndex2526-1)+nsldsDischargedLoanChangeFlagLength2526]), // Field # 618

		NSLDSLoanSatisfactoryRepaymentChangeFlag: strings.TrimSpace(record[nsldsLoanSatisfactoryRepaymentChangeFlagStartIndex2526-1 : (nsldsLoanSatisfactoryRepaymentChangeFlagStartIndex2526-1)+nsldsLoanSatisfactoryRepaymentChangeFlagLength2526]), // Field # 619

		NSLDSActiveBankruptcyChangeFlag: strings.TrimSpace(record[nsldsActiveBankruptcyChangeFlagStartIndex2526-1 : (nsldsActiveBankruptcyChangeFlagStartIndex2526-1)+nsldsActiveBankruptcyChangeFlagLength2526]), // Field # 620

		NSLDSTEACHGrantToLoanConversionChangeFlag: strings.TrimSpace(record[nsldsTEACHGrantToLoanConversionChangeFlagStartIndex2526-1 : (nsldsTEACHGrantToLoanConversionChangeFlagStartIndex2526-1)+nsldsTEACHGrantToLoanConversionChangeFlagLength2526]), // Field # 621

		NSLDSOverpaymentsChangeFlag: strings.TrimSpace(record[nsldsOverpaymentsChangeFlagStartIndex2526-1 : (nsldsOverpaymentsChangeFlagStartIndex2526-1)+nsldsOverpaymentsChangeFlagLength2526]), // Field # 622

		NSLDSAggregateLoanChangeFlag: strings.TrimSpace(record[nsldsAggregateLoanChangeFlagStartIndex2526-1 : (nsldsAggregateLoanChangeFlagStartIndex2526-1)+nsldsAggregateLoanChangeFlagLength2526]), // Field # 623

		NSLDSPerkinsLoanChangeFlag: strings.TrimSpace(record[nsldsPerkinsLoanChangeFlagStartIndex2526-1 : (nsldsPerkinsLoanChangeFlagStartIndex2526-1)+nsldsPerkinsLoanChangeFlagLength2526]), // Field # 624

		NSLDSPellPaymentChangeFlag: strings.TrimSpace(record[nsldsPellPaymentChangeFlagStartIndex2526-1 : (nsldsPellPaymentChangeFlagStartIndex2526-1)+nsldsPellPaymentChangeFlagLength2526]), // Field # 625

		NSLDSTEACHGrantChangeFlag: strings.TrimSpace(record[nsldsTEACHGrantChangeFlagStartIndex2526-1 : (nsldsTEACHGrantChangeFlagStartIndex2526-1)+nsldsTEACHGrantChangeFlagLength2526]), // Field # 626

		NSLDSAdditionalPellFlag: strings.TrimSpace(record[nsldsAdditionalPellFlagStartIndex2526-1 : (nsldsAdditionalPellFlagStartIndex2526-1)+nsldsAdditionalPellFlagLength2526]), // Field # 627

		NSLDSAdditionalLoansFlag: strings.TrimSpace(record[nsldsAdditionalLoansFlagStartIndex2526-1 : (nsldsAdditionalLoansFlagStartIndex2526-1)+nsldsAdditionalLoansFlagLength2526]), // Field # 628

		NSLDSAdditionalTEACHGrantFlag: strings.TrimSpace(record[nsldsAdditionalTEACHGrantFlagStartIndex2526-1 : (nsldsAdditionalTEACHGrantFlagStartIndex2526-1)+nsldsAdditionalTEACHGrantFlagLength2526]), // Field # 629

		NSLDSDirectLoanMasterPromNoteFlag: strings.TrimSpace(record[nsldsDirectLoanMasterPromNoteFlagStartIndex2526-1 : (nsldsDirectLoanMasterPromNoteFlagStartIndex2526-1)+nsldsDirectLoanMasterPromNoteFlagLength2526]), // Field # 630

		NSLDSDirectLoanPLUSMasterPromNoteFlag: strings.TrimSpace(record[nsldsDirectLoanPLUSMasterPromNoteFlagStartIndex2526-1 : (nsldsDirectLoanPLUSMasterPromNoteFlagStartIndex2526-1)+nsldsDirectLoanPLUSMasterPromNoteFlagLength2526]), // Field # 631

		NSLDSDirectLoanGraduatePLUSMasterPromNoteFlag: strings.TrimSpace(record[nsldsDirectLoanGraduatePLUSMasterPromNoteFlagStartIndex2526-1 : (nsldsDirectLoanGraduatePLUSMasterPromNoteFlagStartIndex2526-1)+nsldsDirectLoanGraduatePLUSMasterPromNoteFlagLength2526]), // Field # 632

		NSLDSUndergraduateSubsidizedLoanLimitFlag: strings.TrimSpace(record[nsldsUndergraduateSubsidizedLoanLimitFlagStartIndex2526-1 : (nsldsUndergraduateSubsidizedLoanLimitFlagStartIndex2526-1)+nsldsUndergraduateSubsidizedLoanLimitFlagLength2526]), // Field # 633

		NSLDSUndergraduateCombinedLoanLimitFlag: strings.TrimSpace(record[nsldsUndergraduateCombinedLoanLimitFlagStartIndex2526-1 : (nsldsUndergraduateCombinedLoanLimitFlagStartIndex2526-1)+nsldsUndergraduateCombinedLoanLimitFlagLength2526]), // Field # 634

		NSLDSGraduateSubsidizedLoanLimitFlag: strings.TrimSpace(record[nsldsGraduateSubsidizedLoanLimitFlagStartIndex2526-1 : (nsldsGraduateSubsidizedLoanLimitFlagStartIndex2526-1)+nsldsGraduateSubsidizedLoanLimitFlagLength2526]), // Field # 635

		NSLDSGraduateCombinedLoanLimitFlag: strings.TrimSpace(record[nsldsGraduateCombinedLoanLimitFlagStartIndex2526-1 : (nsldsGraduateCombinedLoanLimitFlagStartIndex2526-1)+nsldsGraduateCombinedLoanLimitFlagLength2526]), // Field # 636

		NSLDSLEULimitIndicator: strings.TrimSpace(record[nsldsLEULimitIndicatorStartIndex2526-1 : (nsldsLEULimitIndicatorStartIndex2526-1)+nsldsLEULimitIndicatorLength2526]), // Field # 637

		NSLDSPellLifetimeEligibilityUsed: strings.TrimSpace(record[nsldsPellLifetimeEligibilityUsedStartIndex2526-1 : (nsldsPellLifetimeEligibilityUsedStartIndex2526-1)+nsldsPellLifetimeEligibilityUsedLength2526]), // Field # 638

		NSLDSSULAFlag: strings.TrimSpace(record[nsldsSULAFlagStartIndex2526-1 : (nsldsSULAFlagStartIndex2526-1)+nsldsSULAFlagLength2526]), // Field # 639

		NSLDSSubsidizedLimitEligibilityFlag: strings.TrimSpace(record[nsldsSubsidizedLimitEligibilityFlagStartIndex2526-1 : (nsldsSubsidizedLimitEligibilityFlagStartIndex2526-1)+nsldsSubsidizedLimitEligibilityFlagLength2526]), // Field # 640

		NSLDSUnusualEnrollmentHistoryFlag: strings.TrimSpace(record[nsldsUnusualEnrollmentHistoryFlagStartIndex2526-1 : (nsldsUnusualEnrollmentHistoryFlagStartIndex2526-1)+nsldsUnusualEnrollmentHistoryFlagLength2526]), // Field # 641

		NSLDSPellSequenceNumber1: strings.TrimSpace(record[nsldsPellSequenceNumber1StartIndex2526-1 : (nsldsPellSequenceNumber1StartIndex2526-1)+nsldsPellSequenceNumber1Length2526]), // Field # 643

		NSLDSPellVerificationFlag1: strings.TrimSpace(record[nsldsPellVerificationFlag1StartIndex2526-1 : (nsldsPellVerificationFlag1StartIndex2526-1)+nsldsPellVerificationFlag1Length2526]), // Field # 644

		NSLDSSAI1: strings.TrimSpace(record[nsldsSAI1StartIndex2526-1 : (nsldsSAI1StartIndex2526-1)+nsldsSAI1Length2526]), // Field # 645

		NSLDSPellSchoolCode1: strings.TrimSpace(record[nsldsPellSchoolCode1StartIndex2526-1 : (nsldsPellSchoolCode1StartIndex2526-1)+nsldsPellSchoolCode1Length2526]), // Field # 646

		NSLDSPellTransactionNumber1: strings.TrimSpace(record[nsldsPellTransactionNumber1StartIndex2526-1 : (nsldsPellTransactionNumber1StartIndex2526-1)+nsldsPellTransactionNumber1Length2526]), // Field # 647

		NSLDSPellLastDisbursementDate1: parseISIRDate2526(strings.TrimSpace(record[nsldsPellLastDisbursementDate1StartIndex2526-1 : (nsldsPellLastDisbursementDate1StartIndex2526-1)+nsldsPellLastDisbursementDate1Length2526])), // Field # 648

		NSLDSPellScheduledAmount1: strings.TrimSpace(record[nsldsPellScheduledAmount1StartIndex2526-1 : (nsldsPellScheduledAmount1StartIndex2526-1)+nsldsPellScheduledAmount1Length2526]), // Field # 649

		NSLDSPellAmountPaidToDate1: parseISIRDate2526(strings.TrimSpace(record[nsldsPellAmountPaidToDate1StartIndex2526-1 : (nsldsPellAmountPaidToDate1StartIndex2526-1)+nsldsPellAmountPaidToDate1Length2526])), // Field # 650

		NSLDSPellPercentEligibilityUsedDecimal1: strings.TrimSpace(record[nsldsPellPercentEligibilityUsedDecimal1StartIndex2526-1 : (nsldsPellPercentEligibilityUsedDecimal1StartIndex2526-1)+nsldsPellPercentEligibilityUsedDecimal1Length2526]), // Field # 651

		NSLDSPellAwardAmount1: strings.TrimSpace(record[nsldsPellAwardAmount1StartIndex2526-1 : (nsldsPellAwardAmount1StartIndex2526-1)+nsldsPellAwardAmount1Length2526]), // Field # 652

		NSLDSAdditionalEligibilityIndicator1: strings.TrimSpace(record[nsldsAdditionalEligibilityIndicator1StartIndex2526-1 : (nsldsAdditionalEligibilityIndicator1StartIndex2526-1)+nsldsAdditionalEligibilityIndicator1Length2526]), // Field # 653

		NSLDSPellSequenceNumber2: strings.TrimSpace(record[nsldsPellSequenceNumber2StartIndex2526-1 : (nsldsPellSequenceNumber2StartIndex2526-1)+nsldsPellSequenceNumber2Length2526]), // Field # 655

		NSLDSPellVerificationFlag2: strings.TrimSpace(record[nsldsPellVerificationFlag2StartIndex2526-1 : (nsldsPellVerificationFlag2StartIndex2526-1)+nsldsPellVerificationFlag2Length2526]), // Field # 656

		NSLDSSAI2: strings.TrimSpace(record[nsldsSAI2StartIndex2526-1 : (nsldsSAI2StartIndex2526-1)+nsldsSAI2Length2526]), // Field # 657

		NSLDSPellSchoolCode2: strings.TrimSpace(record[nsldsPellSchoolCode2StartIndex2526-1 : (nsldsPellSchoolCode2StartIndex2526-1)+nsldsPellSchoolCode2Length2526]), // Field # 658

		NSLDSPellTransactionNumber2: strings.TrimSpace(record[nsldsPellTransactionNumber2StartIndex2526-1 : (nsldsPellTransactionNumber2StartIndex2526-1)+nsldsPellTransactionNumber2Length2526]), // Field # 659

		NSLDSPellLastDisbursementDate2: parseISIRDate2526(strings.TrimSpace(record[nsldsPellLastDisbursementDate2StartIndex2526-1 : (nsldsPellLastDisbursementDate2StartIndex2526-1)+nsldsPellLastDisbursementDate2Length2526])), // Field # 660

		NSLDSPellScheduledAmount2: strings.TrimSpace(record[nsldsPellScheduledAmount2StartIndex2526-1 : (nsldsPellScheduledAmount2StartIndex2526-1)+nsldsPellScheduledAmount2Length2526]), // Field # 661

		NSLDSPellAmountPaidToDate2: parseISIRDate2526(strings.TrimSpace(record[nsldsPellAmountPaidToDate2StartIndex2526-1 : (nsldsPellAmountPaidToDate2StartIndex2526-1)+nsldsPellAmountPaidToDate2Length2526])), // Field # 662

		NSLDSPellPercentEligibilityUsedDecimal2: strings.TrimSpace(record[nsldsPellPercentEligibilityUsedDecimal2StartIndex2526-1 : (nsldsPellPercentEligibilityUsedDecimal2StartIndex2526-1)+nsldsPellPercentEligibilityUsedDecimal2Length2526]), // Field # 663

		NSLDSPellAwardAmount2: strings.TrimSpace(record[nsldsPellAwardAmount2StartIndex2526-1 : (nsldsPellAwardAmount2StartIndex2526-1)+nsldsPellAwardAmount2Length2526]), // Field # 664

		NSLDSAdditionalEligibilityIndicator2: strings.TrimSpace(record[nsldsAdditionalEligibilityIndicator2StartIndex2526-1 : (nsldsAdditionalEligibilityIndicator2StartIndex2526-1)+nsldsAdditionalEligibilityIndicator2Length2526]), // Field # 665

		NSLDSPellSequenceNumber3: strings.TrimSpace(record[nsldsPellSequenceNumber3StartIndex2526-1 : (nsldsPellSequenceNumber3StartIndex2526-1)+nsldsPellSequenceNumber3Length2526]), // Field # 667

		NSLDSPellVerificationFlag3: strings.TrimSpace(record[nsldsPellVerificationFlag3StartIndex2526-1 : (nsldsPellVerificationFlag3StartIndex2526-1)+nsldsPellVerificationFlag3Length2526]), // Field # 668

		NSLDSSAI3: strings.TrimSpace(record[nsldsSAI3StartIndex2526-1 : (nsldsSAI3StartIndex2526-1)+nsldsSAI3Length2526]), // Field # 669

		NSLDSPellSchoolCode3: strings.TrimSpace(record[nsldsPellSchoolCode3StartIndex2526-1 : (nsldsPellSchoolCode3StartIndex2526-1)+nsldsPellSchoolCode3Length2526]), // Field # 670

		NSLDSPellTransactionNumber3: strings.TrimSpace(record[nsldsPellTransactionNumber3StartIndex2526-1 : (nsldsPellTransactionNumber3StartIndex2526-1)+nsldsPellTransactionNumber3Length2526]), // Field # 671

		NSLDSPellLastDisbursementDate3: parseISIRDate2526(strings.TrimSpace(record[nsldsPellLastDisbursementDate3StartIndex2526-1 : (nsldsPellLastDisbursementDate3StartIndex2526-1)+nsldsPellLastDisbursementDate3Length2526])), // Field # 672

		NSLDSPellScheduledAmount3: strings.TrimSpace(record[nsldsPellScheduledAmount3StartIndex2526-1 : (nsldsPellScheduledAmount3StartIndex2526-1)+nsldsPellScheduledAmount3Length2526]), // Field # 673

		NSLDSPellAmountPaidToDate3: parseISIRDate2526(strings.TrimSpace(record[nsldsPellAmountPaidToDate3StartIndex2526-1 : (nsldsPellAmountPaidToDate3StartIndex2526-1)+nsldsPellAmountPaidToDate3Length2526])), // Field # 674

		NSLDSPellPercentEligibilityUsedDecimal3: strings.TrimSpace(record[nsldsPellPercentEligibilityUsedDecimal3StartIndex2526-1 : (nsldsPellPercentEligibilityUsedDecimal3StartIndex2526-1)+nsldsPellPercentEligibilityUsedDecimal3Length2526]), // Field # 675

		NSLDSPellAwardAmount3: strings.TrimSpace(record[nsldsPellAwardAmount3StartIndex2526-1 : (nsldsPellAwardAmount3StartIndex2526-1)+nsldsPellAwardAmount3Length2526]), // Field # 676

		NSLDSAdditionalEligibilityIndicator3: strings.TrimSpace(record[nsldsAdditionalEligibilityIndicator3StartIndex2526-1 : (nsldsAdditionalEligibilityIndicator3StartIndex2526-1)+nsldsAdditionalEligibilityIndicator3Length2526]), // Field # 677

		NSLDSTEACHGrantSequence1: strings.TrimSpace(record[nsldsTEACHGrantSequence1StartIndex2526-1 : (nsldsTEACHGrantSequence1StartIndex2526-1)+nsldsTEACHGrantSequence1Length2526]), // Field # 679

		NSLDSTEACHGrantSchoolCode1: strings.TrimSpace(record[nsldsTEACHGrantSchoolCode1StartIndex2526-1 : (nsldsTEACHGrantSchoolCode1StartIndex2526-1)+nsldsTEACHGrantSchoolCode1Length2526]), // Field # 680

		NSLDSTEACHGrantTransactionNumber1: strings.TrimSpace(record[nsldsTEACHGrantTransactionNumber1StartIndex2526-1 : (nsldsTEACHGrantTransactionNumber1StartIndex2526-1)+nsldsTEACHGrantTransactionNumber1Length2526]), // Field # 681

		NSLDSTEACHGrantLastDisbursementDate1: parseISIRDate2526(strings.TrimSpace(record[nsldsTEACHGrantLastDisbursementDate1StartIndex2526-1 : (nsldsTEACHGrantLastDisbursementDate1StartIndex2526-1)+nsldsTEACHGrantLastDisbursementDate1Length2526])), // Field # 682

		NSLDSTEACHGrantScheduledAmount1: strings.TrimSpace(record[nsldsTEACHGrantScheduledAmount1StartIndex2526-1 : (nsldsTEACHGrantScheduledAmount1StartIndex2526-1)+nsldsTEACHGrantScheduledAmount1Length2526]), // Field # 683

		NSLDSTEACHGrantAmountPaidToDate1: parseISIRDate2526(strings.TrimSpace(record[nsldsTEACHGrantAmountPaidToDate1StartIndex2526-1 : (nsldsTEACHGrantAmountPaidToDate1StartIndex2526-1)+nsldsTEACHGrantAmountPaidToDate1Length2526])), // Field # 684

		NSLDSTEACHGrantAwardAmount1: strings.TrimSpace(record[nsldsTEACHGrantAwardAmount1StartIndex2526-1 : (nsldsTEACHGrantAwardAmount1StartIndex2526-1)+nsldsTEACHGrantAwardAmount1Length2526]), // Field # 685

		NSLDSTEACHGrantAcademicYearLevel1: strings.TrimSpace(record[nsldsTEACHGrantAcademicYearLevel1StartIndex2526-1 : (nsldsTEACHGrantAcademicYearLevel1StartIndex2526-1)+nsldsTEACHGrantAcademicYearLevel1Length2526]), // Field # 686

		NSLDSTEACHGrantAwardYear1: strings.TrimSpace(record[nsldsTEACHGrantAwardYear1StartIndex2526-1 : (nsldsTEACHGrantAwardYear1StartIndex2526-1)+nsldsTEACHGrantAwardYear1Length2526]), // Field # 687

		NSLDSTEACHGrantLoanConversionFlag1: strings.TrimSpace(record[nsldsTEACHGrantLoanConversionFlag1StartIndex2526-1 : (nsldsTEACHGrantLoanConversionFlag1StartIndex2526-1)+nsldsTEACHGrantLoanConversionFlag1Length2526]), // Field # 688

		NSLDSTEACHGrantDischargeCode1: strings.TrimSpace(record[nsldsTEACHGrantDischargeCode1StartIndex2526-1 : (nsldsTEACHGrantDischargeCode1StartIndex2526-1)+nsldsTEACHGrantDischargeCode1Length2526]), // Field # 689

		NSLDSTEACHGrantDischargeAmount1: strings.TrimSpace(record[nsldsTEACHGrantDischargeAmount1StartIndex2526-1 : (nsldsTEACHGrantDischargeAmount1StartIndex2526-1)+nsldsTEACHGrantDischargeAmount1Length2526]), // Field # 690

		NSLDSTEACHGrantAdjustedDisbursement1: strings.TrimSpace(record[nsldsTEACHGrantAdjustedDisbursement1StartIndex2526-1 : (nsldsTEACHGrantAdjustedDisbursement1StartIndex2526-1)+nsldsTEACHGrantAdjustedDisbursement1Length2526]), // Field # 691

		NSLDSTEACHGrantSequence2: strings.TrimSpace(record[nsldsTEACHGrantSequence2StartIndex2526-1 : (nsldsTEACHGrantSequence2StartIndex2526-1)+nsldsTEACHGrantSequence2Length2526]), // Field # 693

		NSLDSTEACHGrantSchoolCode2: strings.TrimSpace(record[nsldsTEACHGrantSchoolCode2StartIndex2526-1 : (nsldsTEACHGrantSchoolCode2StartIndex2526-1)+nsldsTEACHGrantSchoolCode2Length2526]), // Field # 694

		NSLDSTEACHGrantTransactionNumber2: strings.TrimSpace(record[nsldsTEACHGrantTransactionNumber2StartIndex2526-1 : (nsldsTEACHGrantTransactionNumber2StartIndex2526-1)+nsldsTEACHGrantTransactionNumber2Length2526]), // Field # 695

		NSLDSTEACHGrantLastDisbursementDate2: parseISIRDate2526(strings.TrimSpace(record[nsldsTEACHGrantLastDisbursementDate2StartIndex2526-1 : (nsldsTEACHGrantLastDisbursementDate2StartIndex2526-1)+nsldsTEACHGrantLastDisbursementDate2Length2526])), // Field # 696

		NSLDSTEACHGrantScheduledAmount2: strings.TrimSpace(record[nsldsTEACHGrantScheduledAmount2StartIndex2526-1 : (nsldsTEACHGrantScheduledAmount2StartIndex2526-1)+nsldsTEACHGrantScheduledAmount2Length2526]), // Field # 697

		NSLDSTEACHGrantAmountPaidToDate2: parseISIRDate2526(strings.TrimSpace(record[nsldsTEACHGrantAmountPaidToDate2StartIndex2526-1 : (nsldsTEACHGrantAmountPaidToDate2StartIndex2526-1)+nsldsTEACHGrantAmountPaidToDate2Length2526])), // Field # 698

		NSLDSTEACHGrantAwardAmount2: strings.TrimSpace(record[nsldsTEACHGrantAwardAmount2StartIndex2526-1 : (nsldsTEACHGrantAwardAmount2StartIndex2526-1)+nsldsTEACHGrantAwardAmount2Length2526]), // Field # 699

		NSLDSTEACHGrantAcademicYearLevel2: strings.TrimSpace(record[nsldsTEACHGrantAcademicYearLevel2StartIndex2526-1 : (nsldsTEACHGrantAcademicYearLevel2StartIndex2526-1)+nsldsTEACHGrantAcademicYearLevel2Length2526]), // Field # 700

		NSLDSTEACHGrantAwardYear2: strings.TrimSpace(record[nsldsTEACHGrantAwardYear2StartIndex2526-1 : (nsldsTEACHGrantAwardYear2StartIndex2526-1)+nsldsTEACHGrantAwardYear2Length2526]), // Field # 701

		NSLDSTEACHGrantLoanConversionFlag2: strings.TrimSpace(record[nsldsTEACHGrantLoanConversionFlag2StartIndex2526-1 : (nsldsTEACHGrantLoanConversionFlag2StartIndex2526-1)+nsldsTEACHGrantLoanConversionFlag2Length2526]), // Field # 702

		NSLDSTEACHGrantDischargeCode2: strings.TrimSpace(record[nsldsTEACHGrantDischargeCode2StartIndex2526-1 : (nsldsTEACHGrantDischargeCode2StartIndex2526-1)+nsldsTEACHGrantDischargeCode2Length2526]), // Field # 703

		NSLDSTEACHGrantDischargeAmount2: strings.TrimSpace(record[nsldsTEACHGrantDischargeAmount2StartIndex2526-1 : (nsldsTEACHGrantDischargeAmount2StartIndex2526-1)+nsldsTEACHGrantDischargeAmount2Length2526]), // Field # 704

		NSLDSTEACHGrantAdjustedDisbursement2: strings.TrimSpace(record[nsldsTEACHGrantAdjustedDisbursement2StartIndex2526-1 : (nsldsTEACHGrantAdjustedDisbursement2StartIndex2526-1)+nsldsTEACHGrantAdjustedDisbursement2Length2526]), // Field # 705

		NSLDSTEACHGrantSequence3: strings.TrimSpace(record[nsldsTEACHGrantSequence3StartIndex2526-1 : (nsldsTEACHGrantSequence3StartIndex2526-1)+nsldsTEACHGrantSequence3Length2526]), // Field # 707

		NSLDSTEACHGrantSchoolCode3: strings.TrimSpace(record[nsldsTEACHGrantSchoolCode3StartIndex2526-1 : (nsldsTEACHGrantSchoolCode3StartIndex2526-1)+nsldsTEACHGrantSchoolCode3Length2526]), // Field # 708

		NSLDSTEACHGrantTransactionNumber3: strings.TrimSpace(record[nsldsTEACHGrantTransactionNumber3StartIndex2526-1 : (nsldsTEACHGrantTransactionNumber3StartIndex2526-1)+nsldsTEACHGrantTransactionNumber3Length2526]), // Field # 709

		NSLDSTEACHGrantLastDisbursementDate3: parseISIRDate2526(strings.TrimSpace(record[nsldsTEACHGrantLastDisbursementDate3StartIndex2526-1 : (nsldsTEACHGrantLastDisbursementDate3StartIndex2526-1)+nsldsTEACHGrantLastDisbursementDate3Length2526])), // Field # 710

		NSLDSTEACHGrantScheduledAmount3: strings.TrimSpace(record[nsldsTEACHGrantScheduledAmount3StartIndex2526-1 : (nsldsTEACHGrantScheduledAmount3StartIndex2526-1)+nsldsTEACHGrantScheduledAmount3Length2526]), // Field # 711

		NSLDSTEACHGrantAmountPaidToDate3: parseISIRDate2526(strings.TrimSpace(record[nsldsTEACHGrantAmountPaidToDate3StartIndex2526-1 : (nsldsTEACHGrantAmountPaidToDate3StartIndex2526-1)+nsldsTEACHGrantAmountPaidToDate3Length2526])), // Field # 712

		NSLDSTEACHGrantAwardAmount3: strings.TrimSpace(record[nsldsTEACHGrantAwardAmount3StartIndex2526-1 : (nsldsTEACHGrantAwardAmount3StartIndex2526-1)+nsldsTEACHGrantAwardAmount3Length2526]), // Field # 713

		NSLDSTEACHGrantAcademicYearLevel3: strings.TrimSpace(record[nsldsTEACHGrantAcademicYearLevel3StartIndex2526-1 : (nsldsTEACHGrantAcademicYearLevel3StartIndex2526-1)+nsldsTEACHGrantAcademicYearLevel3Length2526]), // Field # 714

		NSLDSTEACHGrantAwardYear3: strings.TrimSpace(record[nsldsTEACHGrantAwardYear3StartIndex2526-1 : (nsldsTEACHGrantAwardYear3StartIndex2526-1)+nsldsTEACHGrantAwardYear3Length2526]), // Field # 715

		NSLDSTEACHGrantLoanConversionFlag3: strings.TrimSpace(record[nsldsTEACHGrantLoanConversionFlag3StartIndex2526-1 : (nsldsTEACHGrantLoanConversionFlag3StartIndex2526-1)+nsldsTEACHGrantLoanConversionFlag3Length2526]), // Field # 716

		NSLDSTEACHGrantDischargeCode3: strings.TrimSpace(record[nsldsTEACHGrantDischargeCode3StartIndex2526-1 : (nsldsTEACHGrantDischargeCode3StartIndex2526-1)+nsldsTEACHGrantDischargeCode3Length2526]), // Field # 717

		NSLDSTEACHGrantDischargeAmount3: strings.TrimSpace(record[nsldsTEACHGrantDischargeAmount3StartIndex2526-1 : (nsldsTEACHGrantDischargeAmount3StartIndex2526-1)+nsldsTEACHGrantDischargeAmount3Length2526]), // Field # 718

		NSLDSTEACHGrantAdjustedDisbursement3: strings.TrimSpace(record[nsldsTEACHGrantAdjustedDisbursement3StartIndex2526-1 : (nsldsTEACHGrantAdjustedDisbursement3StartIndex2526-1)+nsldsTEACHGrantAdjustedDisbursement3Length2526]), // Field # 719

		NSLDSLoanSequenceNumber1: strings.TrimSpace(record[nsldsLoanSequenceNumber1StartIndex2526-1 : (nsldsLoanSequenceNumber1StartIndex2526-1)+nsldsLoanSequenceNumber1Length2526]), // Field # 721

		NSLDSLoanDefaultedRecentIndicator1: strings.TrimSpace(record[nsldsLoanDefaultedRecentIndicator1StartIndex2526-1 : (nsldsLoanDefaultedRecentIndicator1StartIndex2526-1)+nsldsLoanDefaultedRecentIndicator1Length2526]), // Field # 722

		NSLDSLoanChangeFlag1: strings.TrimSpace(record[nsldsLoanChangeFlag1StartIndex2526-1 : (nsldsLoanChangeFlag1StartIndex2526-1)+nsldsLoanChangeFlag1Length2526]), // Field # 723

		NSLDSLoanTypeCode1: strings.TrimSpace(record[nsldsLoanTypeCode1StartIndex2526-1 : (nsldsLoanTypeCode1StartIndex2526-1)+nsldsLoanTypeCode1Length2526]), // Field # 724

		NSLDSLoanNetAmount1: strings.TrimSpace(record[nsldsLoanNetAmount1StartIndex2526-1 : (nsldsLoanNetAmount1StartIndex2526-1)+nsldsLoanNetAmount1Length2526]), // Field # 725

		NSLDSLoanCurrentStatusCode1: strings.TrimSpace(record[nsldsLoanCurrentStatusCode1StartIndex2526-1 : (nsldsLoanCurrentStatusCode1StartIndex2526-1)+nsldsLoanCurrentStatusCode1Length2526]), // Field # 726

		NSLDSLoanCurrentStatusDate1: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanCurrentStatusDate1StartIndex2526-1 : (nsldsLoanCurrentStatusDate1StartIndex2526-1)+nsldsLoanCurrentStatusDate1Length2526])), // Field # 727

		NSLDSLoanOutstandingPrincipalBalance1: strings.TrimSpace(record[nsldsLoanOutstandingPrincipalBalance1StartIndex2526-1 : (nsldsLoanOutstandingPrincipalBalance1StartIndex2526-1)+nsldsLoanOutstandingPrincipalBalance1Length2526]), // Field # 728

		NSLDSLoanOutstandingPrincipalBalanceDate1: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanOutstandingPrincipalBalanceDate1StartIndex2526-1 : (nsldsLoanOutstandingPrincipalBalanceDate1StartIndex2526-1)+nsldsLoanOutstandingPrincipalBalanceDate1Length2526])), // Field # 729

		NSLDSLoanPeriodBeginDate1: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanPeriodBeginDate1StartIndex2526-1 : (nsldsLoanPeriodBeginDate1StartIndex2526-1)+nsldsLoanPeriodBeginDate1Length2526])), // Field # 730

		NSLDSLoanPeriodEndDate1: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanPeriodEndDate1StartIndex2526-1 : (nsldsLoanPeriodEndDate1StartIndex2526-1)+nsldsLoanPeriodEndDate1Length2526])), // Field # 731

		NSLDSLoanGuarantyAgencyCode1: strings.TrimSpace(record[nsldsLoanGuarantyAgencyCode1StartIndex2526-1 : (nsldsLoanGuarantyAgencyCode1StartIndex2526-1)+nsldsLoanGuarantyAgencyCode1Length2526]), // Field # 732

		NSLDSLoanContactType1: strings.TrimSpace(record[nsldsLoanContactType1StartIndex2526-1 : (nsldsLoanContactType1StartIndex2526-1)+nsldsLoanContactType1Length2526]), // Field # 733

		NSLDSLoanSchoolCode1: strings.TrimSpace(record[nsldsLoanSchoolCode1StartIndex2526-1 : (nsldsLoanSchoolCode1StartIndex2526-1)+nsldsLoanSchoolCode1Length2526]), // Field # 734

		NSLDSLoanContactCode1: strings.TrimSpace(record[nsldsLoanContactCode1StartIndex2526-1 : (nsldsLoanContactCode1StartIndex2526-1)+nsldsLoanContactCode1Length2526]), // Field # 735

		NSLDSLoanGradeLevel1: strings.TrimSpace(record[nsldsLoanGradeLevel1StartIndex2526-1 : (nsldsLoanGradeLevel1StartIndex2526-1)+nsldsLoanGradeLevel1Length2526]), // Field # 736

		NSLDSLoanAdditionalUnsubsidizedFlag1: strings.TrimSpace(record[nsldsLoanAdditionalUnsubsidizedFlag1StartIndex2526-1 : (nsldsLoanAdditionalUnsubsidizedFlag1StartIndex2526-1)+nsldsLoanAdditionalUnsubsidizedFlag1Length2526]), // Field # 737

		NSLDSLoanCapitalizedInterestFlag1: strings.TrimSpace(record[nsldsLoanCapitalizedInterestFlag1StartIndex2526-1 : (nsldsLoanCapitalizedInterestFlag1StartIndex2526-1)+nsldsLoanCapitalizedInterestFlag1Length2526]), // Field # 738

		NSLDSLoanDisbursementAmount1: strings.TrimSpace(record[nsldsLoanDisbursementAmount1StartIndex2526-1 : (nsldsLoanDisbursementAmount1StartIndex2526-1)+nsldsLoanDisbursementAmount1Length2526]), // Field # 739

		NSLDSLoanDisbursementDate1: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanDisbursementDate1StartIndex2526-1 : (nsldsLoanDisbursementDate1StartIndex2526-1)+nsldsLoanDisbursementDate1Length2526])), // Field # 740

		NSLDSLoanConfirmedLoanSubsidyStatus1: strings.TrimSpace(record[nsldsLoanConfirmedLoanSubsidyStatus1StartIndex2526-1 : (nsldsLoanConfirmedLoanSubsidyStatus1StartIndex2526-1)+nsldsLoanConfirmedLoanSubsidyStatus1Length2526]), // Field # 741

		NSLDSLoanSubsidyStatusDate1: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanSubsidyStatusDate1StartIndex2526-1 : (nsldsLoanSubsidyStatusDate1StartIndex2526-1)+nsldsLoanSubsidyStatusDate1Length2526])), // Field # 742

		NSLDSLoanSequenceNumber2: strings.TrimSpace(record[nsldsLoanSequenceNumber2StartIndex2526-1 : (nsldsLoanSequenceNumber2StartIndex2526-1)+nsldsLoanSequenceNumber2Length2526]), // Field # 744

		NSLDSLoanDefaultedRecentIndicator2: strings.TrimSpace(record[nsldsLoanDefaultedRecentIndicator2StartIndex2526-1 : (nsldsLoanDefaultedRecentIndicator2StartIndex2526-1)+nsldsLoanDefaultedRecentIndicator2Length2526]), // Field # 745

		NSLDSLoanChangeFlag2: strings.TrimSpace(record[nsldsLoanChangeFlag2StartIndex2526-1 : (nsldsLoanChangeFlag2StartIndex2526-1)+nsldsLoanChangeFlag2Length2526]), // Field # 746

		NSLDSLoanTypeCode2: strings.TrimSpace(record[nsldsLoanTypeCode2StartIndex2526-1 : (nsldsLoanTypeCode2StartIndex2526-1)+nsldsLoanTypeCode2Length2526]), // Field # 747

		NSLDSLoanNetAmount2: strings.TrimSpace(record[nsldsLoanNetAmount2StartIndex2526-1 : (nsldsLoanNetAmount2StartIndex2526-1)+nsldsLoanNetAmount2Length2526]), // Field # 748

		NSLDSLoanCurrentStatusCode2: strings.TrimSpace(record[nsldsLoanCurrentStatusCode2StartIndex2526-1 : (nsldsLoanCurrentStatusCode2StartIndex2526-1)+nsldsLoanCurrentStatusCode2Length2526]), // Field # 749

		NSLDSLoanCurrentStatusDate2: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanCurrentStatusDate2StartIndex2526-1 : (nsldsLoanCurrentStatusDate2StartIndex2526-1)+nsldsLoanCurrentStatusDate2Length2526])), // Field # 750

		NSLDSLoanOutstandingPrincipalBalance2: strings.TrimSpace(record[nsldsLoanOutstandingPrincipalBalance2StartIndex2526-1 : (nsldsLoanOutstandingPrincipalBalance2StartIndex2526-1)+nsldsLoanOutstandingPrincipalBalance2Length2526]), // Field # 751

		NSLDSLoanOutstandingPrincipalBalanceDate2: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanOutstandingPrincipalBalanceDate2StartIndex2526-1 : (nsldsLoanOutstandingPrincipalBalanceDate2StartIndex2526-1)+nsldsLoanOutstandingPrincipalBalanceDate2Length2526])), // Field # 752

		NSLDSLoanPeriodBeginDate2: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanPeriodBeginDate2StartIndex2526-1 : (nsldsLoanPeriodBeginDate2StartIndex2526-1)+nsldsLoanPeriodBeginDate2Length2526])), // Field # 753

		NSLDSLoanPeriodEndDate2: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanPeriodEndDate2StartIndex2526-1 : (nsldsLoanPeriodEndDate2StartIndex2526-1)+nsldsLoanPeriodEndDate2Length2526])), // Field # 754

		NSLDSLoanGuarantyAgencyCode2: strings.TrimSpace(record[nsldsLoanGuarantyAgencyCode2StartIndex2526-1 : (nsldsLoanGuarantyAgencyCode2StartIndex2526-1)+nsldsLoanGuarantyAgencyCode2Length2526]), // Field # 755

		NSLDSLoanContactType2: strings.TrimSpace(record[nsldsLoanContactType2StartIndex2526-1 : (nsldsLoanContactType2StartIndex2526-1)+nsldsLoanContactType2Length2526]), // Field # 756

		NSLDSLoanSchoolCode2: strings.TrimSpace(record[nsldsLoanSchoolCode2StartIndex2526-1 : (nsldsLoanSchoolCode2StartIndex2526-1)+nsldsLoanSchoolCode2Length2526]), // Field # 757

		NSLDSLoanContactCode2: strings.TrimSpace(record[nsldsLoanContactCode2StartIndex2526-1 : (nsldsLoanContactCode2StartIndex2526-1)+nsldsLoanContactCode2Length2526]), // Field # 758

		NSLDSLoanGradeLevel2: strings.TrimSpace(record[nsldsLoanGradeLevel2StartIndex2526-1 : (nsldsLoanGradeLevel2StartIndex2526-1)+nsldsLoanGradeLevel2Length2526]), // Field # 759

		NSLDSLoanAdditionalUnsubsidizedFlag2: strings.TrimSpace(record[nsldsLoanAdditionalUnsubsidizedFlag2StartIndex2526-1 : (nsldsLoanAdditionalUnsubsidizedFlag2StartIndex2526-1)+nsldsLoanAdditionalUnsubsidizedFlag2Length2526]), // Field # 760

		NSLDSLoanCapitalizedInterestFlag2: strings.TrimSpace(record[nsldsLoanCapitalizedInterestFlag2StartIndex2526-1 : (nsldsLoanCapitalizedInterestFlag2StartIndex2526-1)+nsldsLoanCapitalizedInterestFlag2Length2526]), // Field # 761

		NSLDSLoanDisbursementAmount2: strings.TrimSpace(record[nsldsLoanDisbursementAmount2StartIndex2526-1 : (nsldsLoanDisbursementAmount2StartIndex2526-1)+nsldsLoanDisbursementAmount2Length2526]), // Field # 762

		NSLDSLoanDisbursementDate2: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanDisbursementDate2StartIndex2526-1 : (nsldsLoanDisbursementDate2StartIndex2526-1)+nsldsLoanDisbursementDate2Length2526])), // Field # 763

		NSLDSLoanConfirmedLoanSubsidyStatus2: strings.TrimSpace(record[nsldsLoanConfirmedLoanSubsidyStatus2StartIndex2526-1 : (nsldsLoanConfirmedLoanSubsidyStatus2StartIndex2526-1)+nsldsLoanConfirmedLoanSubsidyStatus2Length2526]), // Field # 764

		NSLDSLoanSubsidyStatusDate2: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanSubsidyStatusDate2StartIndex2526-1 : (nsldsLoanSubsidyStatusDate2StartIndex2526-1)+nsldsLoanSubsidyStatusDate2Length2526])), // Field # 765

		NSLDSLoanSequenceNumber3: strings.TrimSpace(record[nsldsLoanSequenceNumber3StartIndex2526-1 : (nsldsLoanSequenceNumber3StartIndex2526-1)+nsldsLoanSequenceNumber3Length2526]), // Field # 767

		NSLDSLoanDefaultedRecentIndicator3: strings.TrimSpace(record[nsldsLoanDefaultedRecentIndicator3StartIndex2526-1 : (nsldsLoanDefaultedRecentIndicator3StartIndex2526-1)+nsldsLoanDefaultedRecentIndicator3Length2526]), // Field # 768

		NSLDSLoanChangeFlag3: strings.TrimSpace(record[nsldsLoanChangeFlag3StartIndex2526-1 : (nsldsLoanChangeFlag3StartIndex2526-1)+nsldsLoanChangeFlag3Length2526]), // Field # 769

		NSLDSLoanTypeCode3: strings.TrimSpace(record[nsldsLoanTypeCode3StartIndex2526-1 : (nsldsLoanTypeCode3StartIndex2526-1)+nsldsLoanTypeCode3Length2526]), // Field # 770

		NSLDSLoanNetAmount3: strings.TrimSpace(record[nsldsLoanNetAmount3StartIndex2526-1 : (nsldsLoanNetAmount3StartIndex2526-1)+nsldsLoanNetAmount3Length2526]), // Field # 771

		NSLDSLoanCurrentStatusCode3: strings.TrimSpace(record[nsldsLoanCurrentStatusCode3StartIndex2526-1 : (nsldsLoanCurrentStatusCode3StartIndex2526-1)+nsldsLoanCurrentStatusCode3Length2526]), // Field # 772

		NSLDSLoanCurrentStatusDate3: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanCurrentStatusDate3StartIndex2526-1 : (nsldsLoanCurrentStatusDate3StartIndex2526-1)+nsldsLoanCurrentStatusDate3Length2526])), // Field # 773

		NSLDSLoanOutstandingPrincipalBalance3: strings.TrimSpace(record[nsldsLoanOutstandingPrincipalBalance3StartIndex2526-1 : (nsldsLoanOutstandingPrincipalBalance3StartIndex2526-1)+nsldsLoanOutstandingPrincipalBalance3Length2526]), // Field # 774

		NSLDSLoanOutstandingPrincipalBalanceDate3: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanOutstandingPrincipalBalanceDate3StartIndex2526-1 : (nsldsLoanOutstandingPrincipalBalanceDate3StartIndex2526-1)+nsldsLoanOutstandingPrincipalBalanceDate3Length2526])), // Field # 775

		NSLDSLoanPeriodBeginDate3: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanPeriodBeginDate3StartIndex2526-1 : (nsldsLoanPeriodBeginDate3StartIndex2526-1)+nsldsLoanPeriodBeginDate3Length2526])), // Field # 776

		NSLDSLoanPeriodEndDate3: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanPeriodEndDate3StartIndex2526-1 : (nsldsLoanPeriodEndDate3StartIndex2526-1)+nsldsLoanPeriodEndDate3Length2526])), // Field # 777

		NSLDSLoanGuarantyAgencyCode3: strings.TrimSpace(record[nsldsLoanGuarantyAgencyCode3StartIndex2526-1 : (nsldsLoanGuarantyAgencyCode3StartIndex2526-1)+nsldsLoanGuarantyAgencyCode3Length2526]), // Field # 778

		NSLDSLoanContactType3: strings.TrimSpace(record[nsldsLoanContactType3StartIndex2526-1 : (nsldsLoanContactType3StartIndex2526-1)+nsldsLoanContactType3Length2526]), // Field # 779

		NSLDSLoanSchoolCode3: strings.TrimSpace(record[nsldsLoanSchoolCode3StartIndex2526-1 : (nsldsLoanSchoolCode3StartIndex2526-1)+nsldsLoanSchoolCode3Length2526]), // Field # 780

		NSLDSLoanContactCode3: strings.TrimSpace(record[nsldsLoanContactCode3StartIndex2526-1 : (nsldsLoanContactCode3StartIndex2526-1)+nsldsLoanContactCode3Length2526]), // Field # 781

		NSLDSLoanGradeLevel3: strings.TrimSpace(record[nsldsLoanGradeLevel3StartIndex2526-1 : (nsldsLoanGradeLevel3StartIndex2526-1)+nsldsLoanGradeLevel3Length2526]), // Field # 782

		NSLDSLoanAdditionalUnsubsidizedFlag3: strings.TrimSpace(record[nsldsLoanAdditionalUnsubsidizedFlag3StartIndex2526-1 : (nsldsLoanAdditionalUnsubsidizedFlag3StartIndex2526-1)+nsldsLoanAdditionalUnsubsidizedFlag3Length2526]), // Field # 783

		NSLDSLoanCapitalizedInterestFlag3: strings.TrimSpace(record[nsldsLoanCapitalizedInterestFlag3StartIndex2526-1 : (nsldsLoanCapitalizedInterestFlag3StartIndex2526-1)+nsldsLoanCapitalizedInterestFlag3Length2526]), // Field # 784

		NSLDSLoanDisbursementAmount3: strings.TrimSpace(record[nsldsLoanDisbursementAmount3StartIndex2526-1 : (nsldsLoanDisbursementAmount3StartIndex2526-1)+nsldsLoanDisbursementAmount3Length2526]), // Field # 785

		NSLDSLoanDisbursementDate3: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanDisbursementDate3StartIndex2526-1 : (nsldsLoanDisbursementDate3StartIndex2526-1)+nsldsLoanDisbursementDate3Length2526])), // Field # 786

		NSLDSLoanConfirmedLoanSubsidyStatus3: strings.TrimSpace(record[nsldsLoanConfirmedLoanSubsidyStatus3StartIndex2526-1 : (nsldsLoanConfirmedLoanSubsidyStatus3StartIndex2526-1)+nsldsLoanConfirmedLoanSubsidyStatus3Length2526]), // Field # 787

		NSLDSLoanSubsidyStatusDate3: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanSubsidyStatusDate3StartIndex2526-1 : (nsldsLoanSubsidyStatusDate3StartIndex2526-1)+nsldsLoanSubsidyStatusDate3Length2526])), // Field # 788

		NSLDSLoanSequenceNumber4: strings.TrimSpace(record[nsldsLoanSequenceNumber4StartIndex2526-1 : (nsldsLoanSequenceNumber4StartIndex2526-1)+nsldsLoanSequenceNumber4Length2526]), // Field # 790

		NSLDSLoanDefaultedRecentIndicator4: strings.TrimSpace(record[nsldsLoanDefaultedRecentIndicator4StartIndex2526-1 : (nsldsLoanDefaultedRecentIndicator4StartIndex2526-1)+nsldsLoanDefaultedRecentIndicator4Length2526]), // Field # 791

		NSLDSLoanChangeFlag4: strings.TrimSpace(record[nsldsLoanChangeFlag4StartIndex2526-1 : (nsldsLoanChangeFlag4StartIndex2526-1)+nsldsLoanChangeFlag4Length2526]), // Field # 792

		NSLDSLoanTypeCode4: strings.TrimSpace(record[nsldsLoanTypeCode4StartIndex2526-1 : (nsldsLoanTypeCode4StartIndex2526-1)+nsldsLoanTypeCode4Length2526]), // Field # 793

		NSLDSLoanNetAmount4: strings.TrimSpace(record[nsldsLoanNetAmount4StartIndex2526-1 : (nsldsLoanNetAmount4StartIndex2526-1)+nsldsLoanNetAmount4Length2526]), // Field # 794

		NSLDSLoanCurrentStatusCode4: strings.TrimSpace(record[nsldsLoanCurrentStatusCode4StartIndex2526-1 : (nsldsLoanCurrentStatusCode4StartIndex2526-1)+nsldsLoanCurrentStatusCode4Length2526]), // Field # 795

		NSLDSLoanCurrentStatusDate4: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanCurrentStatusDate4StartIndex2526-1 : (nsldsLoanCurrentStatusDate4StartIndex2526-1)+nsldsLoanCurrentStatusDate4Length2526])), // Field # 796

		NSLDSLoanOutstandingPrincipalBalance4: strings.TrimSpace(record[nsldsLoanOutstandingPrincipalBalance4StartIndex2526-1 : (nsldsLoanOutstandingPrincipalBalance4StartIndex2526-1)+nsldsLoanOutstandingPrincipalBalance4Length2526]), // Field # 797

		NSLDSLoanOutstandingPrincipalBalanceDate4: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanOutstandingPrincipalBalanceDate4StartIndex2526-1 : (nsldsLoanOutstandingPrincipalBalanceDate4StartIndex2526-1)+nsldsLoanOutstandingPrincipalBalanceDate4Length2526])), // Field # 798

		NSLDSLoanPeriodBeginDate4: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanPeriodBeginDate4StartIndex2526-1 : (nsldsLoanPeriodBeginDate4StartIndex2526-1)+nsldsLoanPeriodBeginDate4Length2526])), // Field # 799

		NSLDSLoanPeriodEndDate4: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanPeriodEndDate4StartIndex2526-1 : (nsldsLoanPeriodEndDate4StartIndex2526-1)+nsldsLoanPeriodEndDate4Length2526])), // Field # 800

		NSLDSLoanGuarantyAgencyCode4: strings.TrimSpace(record[nsldsLoanGuarantyAgencyCode4StartIndex2526-1 : (nsldsLoanGuarantyAgencyCode4StartIndex2526-1)+nsldsLoanGuarantyAgencyCode4Length2526]), // Field # 801

		NSLDSLoanContactType4: strings.TrimSpace(record[nsldsLoanContactType4StartIndex2526-1 : (nsldsLoanContactType4StartIndex2526-1)+nsldsLoanContactType4Length2526]), // Field # 802

		NSLDSLoanSchoolCode4: strings.TrimSpace(record[nsldsLoanSchoolCode4StartIndex2526-1 : (nsldsLoanSchoolCode4StartIndex2526-1)+nsldsLoanSchoolCode4Length2526]), // Field # 803

		NSLDSLoanContactCode4: strings.TrimSpace(record[nsldsLoanContactCode4StartIndex2526-1 : (nsldsLoanContactCode4StartIndex2526-1)+nsldsLoanContactCode4Length2526]), // Field # 804

		NSLDSLoanGradeLevel4: strings.TrimSpace(record[nsldsLoanGradeLevel4StartIndex2526-1 : (nsldsLoanGradeLevel4StartIndex2526-1)+nsldsLoanGradeLevel4Length2526]), // Field # 805

		NSLDSLoanAdditionalUnsubsidizedFlag4: strings.TrimSpace(record[nsldsLoanAdditionalUnsubsidizedFlag4StartIndex2526-1 : (nsldsLoanAdditionalUnsubsidizedFlag4StartIndex2526-1)+nsldsLoanAdditionalUnsubsidizedFlag4Length2526]), // Field # 806

		NSLDSLoanCapitalizedInterestFlag4: strings.TrimSpace(record[nsldsLoanCapitalizedInterestFlag4StartIndex2526-1 : (nsldsLoanCapitalizedInterestFlag4StartIndex2526-1)+nsldsLoanCapitalizedInterestFlag4Length2526]), // Field # 807

		NSLDSLoanDisbursementAmount4: strings.TrimSpace(record[nsldsLoanDisbursementAmount4StartIndex2526-1 : (nsldsLoanDisbursementAmount4StartIndex2526-1)+nsldsLoanDisbursementAmount4Length2526]), // Field # 808

		NSLDSLoanDisbursementDate4: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanDisbursementDate4StartIndex2526-1 : (nsldsLoanDisbursementDate4StartIndex2526-1)+nsldsLoanDisbursementDate4Length2526])), // Field # 809

		NSLDSLoanConfirmedLoanSubsidyStatus4: strings.TrimSpace(record[nsldsLoanConfirmedLoanSubsidyStatus4StartIndex2526-1 : (nsldsLoanConfirmedLoanSubsidyStatus4StartIndex2526-1)+nsldsLoanConfirmedLoanSubsidyStatus4Length2526]), // Field # 810

		NSLDSLoanSubsidyStatusDate4: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanSubsidyStatusDate4StartIndex2526-1 : (nsldsLoanSubsidyStatusDate4StartIndex2526-1)+nsldsLoanSubsidyStatusDate4Length2526])), // Field # 811

		NSLDSLoanSequenceNumber5: strings.TrimSpace(record[nsldsLoanSequenceNumber5StartIndex2526-1 : (nsldsLoanSequenceNumber5StartIndex2526-1)+nsldsLoanSequenceNumber5Length2526]), // Field # 813

		NSLDSLoanDefaultedRecentIndicator5: strings.TrimSpace(record[nsldsLoanDefaultedRecentIndicator5StartIndex2526-1 : (nsldsLoanDefaultedRecentIndicator5StartIndex2526-1)+nsldsLoanDefaultedRecentIndicator5Length2526]), // Field # 814

		NSLDSLoanChangeFlag5: strings.TrimSpace(record[nsldsLoanChangeFlag5StartIndex2526-1 : (nsldsLoanChangeFlag5StartIndex2526-1)+nsldsLoanChangeFlag5Length2526]), // Field # 815

		NSLDSLoanTypeCode5: strings.TrimSpace(record[nsldsLoanTypeCode5StartIndex2526-1 : (nsldsLoanTypeCode5StartIndex2526-1)+nsldsLoanTypeCode5Length2526]), // Field # 816

		NSLDSLoanNetAmount5: strings.TrimSpace(record[nsldsLoanNetAmount5StartIndex2526-1 : (nsldsLoanNetAmount5StartIndex2526-1)+nsldsLoanNetAmount5Length2526]), // Field # 817

		NSLDSLoanCurrentStatusCode5: strings.TrimSpace(record[nsldsLoanCurrentStatusCode5StartIndex2526-1 : (nsldsLoanCurrentStatusCode5StartIndex2526-1)+nsldsLoanCurrentStatusCode5Length2526]), // Field # 818

		NSLDSLoanCurrentStatusDate5: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanCurrentStatusDate5StartIndex2526-1 : (nsldsLoanCurrentStatusDate5StartIndex2526-1)+nsldsLoanCurrentStatusDate5Length2526])), // Field # 819

		NSLDSLoanOutstandingPrincipalBalance5: strings.TrimSpace(record[nsldsLoanOutstandingPrincipalBalance5StartIndex2526-1 : (nsldsLoanOutstandingPrincipalBalance5StartIndex2526-1)+nsldsLoanOutstandingPrincipalBalance5Length2526]), // Field # 820

		NSLDSLoanOutstandingPrincipalBalanceDate5: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanOutstandingPrincipalBalanceDate5StartIndex2526-1 : (nsldsLoanOutstandingPrincipalBalanceDate5StartIndex2526-1)+nsldsLoanOutstandingPrincipalBalanceDate5Length2526])), // Field # 821

		NSLDSLoanPeriodBeginDate5: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanPeriodBeginDate5StartIndex2526-1 : (nsldsLoanPeriodBeginDate5StartIndex2526-1)+nsldsLoanPeriodBeginDate5Length2526])), // Field # 822

		NSLDSLoanPeriodEndDate5: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanPeriodEndDate5StartIndex2526-1 : (nsldsLoanPeriodEndDate5StartIndex2526-1)+nsldsLoanPeriodEndDate5Length2526])), // Field # 823

		NSLDSLoanGuarantyAgencyCode5: strings.TrimSpace(record[nsldsLoanGuarantyAgencyCode5StartIndex2526-1 : (nsldsLoanGuarantyAgencyCode5StartIndex2526-1)+nsldsLoanGuarantyAgencyCode5Length2526]), // Field # 824

		NSLDSLoanContactType5: strings.TrimSpace(record[nsldsLoanContactType5StartIndex2526-1 : (nsldsLoanContactType5StartIndex2526-1)+nsldsLoanContactType5Length2526]), // Field # 825

		NSLDSLoanSchoolCode5: strings.TrimSpace(record[nsldsLoanSchoolCode5StartIndex2526-1 : (nsldsLoanSchoolCode5StartIndex2526-1)+nsldsLoanSchoolCode5Length2526]), // Field # 826

		NSLDSLoanContactCode5: strings.TrimSpace(record[nsldsLoanContactCode5StartIndex2526-1 : (nsldsLoanContactCode5StartIndex2526-1)+nsldsLoanContactCode5Length2526]), // Field # 827

		NSLDSLoanGradeLevel5: strings.TrimSpace(record[nsldsLoanGradeLevel5StartIndex2526-1 : (nsldsLoanGradeLevel5StartIndex2526-1)+nsldsLoanGradeLevel5Length2526]), // Field # 828

		NSLDSLoanAdditionalUnsubsidizedFlag5: strings.TrimSpace(record[nsldsLoanAdditionalUnsubsidizedFlag5StartIndex2526-1 : (nsldsLoanAdditionalUnsubsidizedFlag5StartIndex2526-1)+nsldsLoanAdditionalUnsubsidizedFlag5Length2526]), // Field # 829

		NSLDSLoanCapitalizedInterestFlag5: strings.TrimSpace(record[nsldsLoanCapitalizedInterestFlag5StartIndex2526-1 : (nsldsLoanCapitalizedInterestFlag5StartIndex2526-1)+nsldsLoanCapitalizedInterestFlag5Length2526]), // Field # 830

		NSLDSLoanDisbursementAmount5: strings.TrimSpace(record[nsldsLoanDisbursementAmount5StartIndex2526-1 : (nsldsLoanDisbursementAmount5StartIndex2526-1)+nsldsLoanDisbursementAmount5Length2526]), // Field # 831

		NSLDSLoanDisbursementDate5: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanDisbursementDate5StartIndex2526-1 : (nsldsLoanDisbursementDate5StartIndex2526-1)+nsldsLoanDisbursementDate5Length2526])), // Field # 832

		NSLDSLoanConfirmedLoanSubsidyStatus5: strings.TrimSpace(record[nsldsLoanConfirmedLoanSubsidyStatus5StartIndex2526-1 : (nsldsLoanConfirmedLoanSubsidyStatus5StartIndex2526-1)+nsldsLoanConfirmedLoanSubsidyStatus5Length2526]), // Field # 833

		NSLDSLoanSubsidyStatusDate5: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanSubsidyStatusDate5StartIndex2526-1 : (nsldsLoanSubsidyStatusDate5StartIndex2526-1)+nsldsLoanSubsidyStatusDate5Length2526])), // Field # 834

		NSLDSLoanSequenceNumber6: strings.TrimSpace(record[nsldsLoanSequenceNumber6StartIndex2526-1 : (nsldsLoanSequenceNumber6StartIndex2526-1)+nsldsLoanSequenceNumber6Length2526]), // Field # 836

		NSLDSLoanDefaultedRecentIndicator6: strings.TrimSpace(record[nsldsLoanDefaultedRecentIndicator6StartIndex2526-1 : (nsldsLoanDefaultedRecentIndicator6StartIndex2526-1)+nsldsLoanDefaultedRecentIndicator6Length2526]), // Field # 837

		NSLDSLoanChangeFlag6: strings.TrimSpace(record[nsldsLoanChangeFlag6StartIndex2526-1 : (nsldsLoanChangeFlag6StartIndex2526-1)+nsldsLoanChangeFlag6Length2526]), // Field # 838

		NSLDSLoanTypeCode6: strings.TrimSpace(record[nsldsLoanTypeCode6StartIndex2526-1 : (nsldsLoanTypeCode6StartIndex2526-1)+nsldsLoanTypeCode6Length2526]), // Field # 839

		NSLDSLoanNetAmount6: strings.TrimSpace(record[nsldsLoanNetAmount6StartIndex2526-1 : (nsldsLoanNetAmount6StartIndex2526-1)+nsldsLoanNetAmount6Length2526]), // Field # 840

		NSLDSLoanCurrentStatusCode6: strings.TrimSpace(record[nsldsLoanCurrentStatusCode6StartIndex2526-1 : (nsldsLoanCurrentStatusCode6StartIndex2526-1)+nsldsLoanCurrentStatusCode6Length2526]), // Field # 841

		NSLDSLoanCurrentStatusDate6: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanCurrentStatusDate6StartIndex2526-1 : (nsldsLoanCurrentStatusDate6StartIndex2526-1)+nsldsLoanCurrentStatusDate6Length2526])), // Field # 842

		NSLDSLoanOutstandingPrincipalBalance6: strings.TrimSpace(record[nsldsLoanOutstandingPrincipalBalance6StartIndex2526-1 : (nsldsLoanOutstandingPrincipalBalance6StartIndex2526-1)+nsldsLoanOutstandingPrincipalBalance6Length2526]), // Field # 843

		NSLDSLoanOutstandingPrincipalBalanceDate6: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanOutstandingPrincipalBalanceDate6StartIndex2526-1 : (nsldsLoanOutstandingPrincipalBalanceDate6StartIndex2526-1)+nsldsLoanOutstandingPrincipalBalanceDate6Length2526])), // Field # 844

		NSLDSLoanPeriodBeginDate6: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanPeriodBeginDate6StartIndex2526-1 : (nsldsLoanPeriodBeginDate6StartIndex2526-1)+nsldsLoanPeriodBeginDate6Length2526])), // Field # 845

		NSLDSLoanPeriodEndDate6: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanPeriodEndDate6StartIndex2526-1 : (nsldsLoanPeriodEndDate6StartIndex2526-1)+nsldsLoanPeriodEndDate6Length2526])), // Field # 846

		NSLDSLoanGuarantyAgencyCode6: strings.TrimSpace(record[nsldsLoanGuarantyAgencyCode6StartIndex2526-1 : (nsldsLoanGuarantyAgencyCode6StartIndex2526-1)+nsldsLoanGuarantyAgencyCode6Length2526]), // Field # 847

		NSLDSLoanContactType6: strings.TrimSpace(record[nsldsLoanContactType6StartIndex2526-1 : (nsldsLoanContactType6StartIndex2526-1)+nsldsLoanContactType6Length2526]), // Field # 848

		NSLDSLoanSchoolCode6: strings.TrimSpace(record[nsldsLoanSchoolCode6StartIndex2526-1 : (nsldsLoanSchoolCode6StartIndex2526-1)+nsldsLoanSchoolCode6Length2526]), // Field # 849

		NSLDSLoanContactCode6: strings.TrimSpace(record[nsldsLoanContactCode6StartIndex2526-1 : (nsldsLoanContactCode6StartIndex2526-1)+nsldsLoanContactCode6Length2526]), // Field # 850

		NSLDSLoanGradeLevel6: strings.TrimSpace(record[nsldsLoanGradeLevel6StartIndex2526-1 : (nsldsLoanGradeLevel6StartIndex2526-1)+nsldsLoanGradeLevel6Length2526]), // Field # 851

		NSLDSLoanAdditionalUnsubsidizedFlag6: strings.TrimSpace(record[nsldsLoanAdditionalUnsubsidizedFlag6StartIndex2526-1 : (nsldsLoanAdditionalUnsubsidizedFlag6StartIndex2526-1)+nsldsLoanAdditionalUnsubsidizedFlag6Length2526]), // Field # 852

		NSLDSLoanCapitalizedInterestFlag6: strings.TrimSpace(record[nsldsLoanCapitalizedInterestFlag6StartIndex2526-1 : (nsldsLoanCapitalizedInterestFlag6StartIndex2526-1)+nsldsLoanCapitalizedInterestFlag6Length2526]), // Field # 853

		NSLDSLoanDisbursementAmount6: strings.TrimSpace(record[nsldsLoanDisbursementAmount6StartIndex2526-1 : (nsldsLoanDisbursementAmount6StartIndex2526-1)+nsldsLoanDisbursementAmount6Length2526]), // Field # 854

		NSLDSLoanDisbursementDate6: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanDisbursementDate6StartIndex2526-1 : (nsldsLoanDisbursementDate6StartIndex2526-1)+nsldsLoanDisbursementDate6Length2526])), // Field # 855

		NSLDSLoanConfirmedLoanSubsidyStatus6: strings.TrimSpace(record[nsldsLoanConfirmedLoanSubsidyStatus6StartIndex2526-1 : (nsldsLoanConfirmedLoanSubsidyStatus6StartIndex2526-1)+nsldsLoanConfirmedLoanSubsidyStatus6Length2526]), // Field # 856

		NSLDSLoanSubsidyStatusDate6: parseISIRDate2526(strings.TrimSpace(record[nsldsLoanSubsidyStatusDate6StartIndex2526-1 : (nsldsLoanSubsidyStatusDate6StartIndex2526-1)+nsldsLoanSubsidyStatusDate6Length2526])), // Field # 857

		FTILabelStart: strings.TrimSpace(record[ftiLabelStartStartIndex2526-1 : (ftiLabelStartStartIndex2526-1)+ftiLabelStartLength2526]), // Field # 861

		StudentFTIMReturnedTaxYear: strings.TrimSpace(record[studentFTIMReturnedTaxYearStartIndex2526-1 : (studentFTIMReturnedTaxYearStartIndex2526-1)+studentFTIMReturnedTaxYearLength2526]), // Field # 862

		StudentFTIMFilingStatusCode: strings.TrimSpace(record[studentFTIMFilingStatusCodeStartIndex2526-1 : (studentFTIMFilingStatusCodeStartIndex2526-1)+studentFTIMFilingStatusCodeLength2526]), // Field # 863

		StudentFTIMAdjustedGrossIncome: strings.TrimSpace(record[studentFTIMAdjustedGrossIncomeStartIndex2526-1 : (studentFTIMAdjustedGrossIncomeStartIndex2526-1)+studentFTIMAdjustedGrossIncomeLength2526]), // Field # 864

		StudentFTIMNumberOfExemptions: strings.TrimSpace(record[studentFTIMNumberOfExemptionsStartIndex2526-1 : (studentFTIMNumberOfExemptionsStartIndex2526-1)+studentFTIMNumberOfExemptionsLength2526]), // Field # 865

		StudentFTIMNumberOfDependents: strings.TrimSpace(record[studentFTIMNumberOfDependentsStartIndex2526-1 : (studentFTIMNumberOfDependentsStartIndex2526-1)+studentFTIMNumberOfDependentsLength2526]), // Field # 866

		StudentFTIMTotalIncomeEarnedAmount: strings.TrimSpace(record[studentFTIMTotalIncomeEarnedAmountStartIndex2526-1 : (studentFTIMTotalIncomeEarnedAmountStartIndex2526-1)+studentFTIMTotalIncomeEarnedAmountLength2526]), // Field # 867

		StudentFTIMIncomeTaxPaid: strings.TrimSpace(record[studentFTIMIncomeTaxPaidStartIndex2526-1 : (studentFTIMIncomeTaxPaidStartIndex2526-1)+studentFTIMIncomeTaxPaidLength2526]), // Field # 868

		StudentFTIMEducationCredits: strings.TrimSpace(record[studentFTIMEducationCreditsStartIndex2526-1 : (studentFTIMEducationCreditsStartIndex2526-1)+studentFTIMEducationCreditsLength2526]), // Field # 869

		StudentFTIMUntaxedIRADistributions: strings.TrimSpace(record[studentFTIMUntaxedIRADistributionsStartIndex2526-1 : (studentFTIMUntaxedIRADistributionsStartIndex2526-1)+studentFTIMUntaxedIRADistributionsLength2526]), // Field # 870

		StudentFTIMIRADeductibleAndPayments: strings.TrimSpace(record[studentFTIMIRADeductibleAndPaymentsStartIndex2526-1 : (studentFTIMIRADeductibleAndPaymentsStartIndex2526-1)+studentFTIMIRADeductibleAndPaymentsLength2526]), // Field # 871

		StudentFTIMTaxExemptInterest: strings.TrimSpace(record[studentFTIMTaxExemptInterestStartIndex2526-1 : (studentFTIMTaxExemptInterestStartIndex2526-1)+studentFTIMTaxExemptInterestLength2526]), // Field # 872

		StudentFTIMUntaxedPensionsAmount: strings.TrimSpace(record[studentFTIMUntaxedPensionsAmountStartIndex2526-1 : (studentFTIMUntaxedPensionsAmountStartIndex2526-1)+studentFTIMUntaxedPensionsAmountLength2526]), // Field # 873

		StudentFTIMScheduleCNetProfitLoss: strings.TrimSpace(record[studentFTIMScheduleCNetProfitLossStartIndex2526-1 : (studentFTIMScheduleCNetProfitLossStartIndex2526-1)+studentFTIMScheduleCNetProfitLossLength2526]), // Field # 874

		StudentFTIMScheduleAIndicator: strings.TrimSpace(record[studentFTIMScheduleAIndicatorStartIndex2526-1 : (studentFTIMScheduleAIndicatorStartIndex2526-1)+studentFTIMScheduleAIndicatorLength2526]), // Field # 875

		StudentFTIMScheduleBIndicator: strings.TrimSpace(record[studentFTIMScheduleBIndicatorStartIndex2526-1 : (studentFTIMScheduleBIndicatorStartIndex2526-1)+studentFTIMScheduleBIndicatorLength2526]), // Field # 876

		StudentFTIMScheduleDIndicator: strings.TrimSpace(record[studentFTIMScheduleDIndicatorStartIndex2526-1 : (studentFTIMScheduleDIndicatorStartIndex2526-1)+studentFTIMScheduleDIndicatorLength2526]), // Field # 877

		StudentFTIMScheduleEIndicator: strings.TrimSpace(record[studentFTIMScheduleEIndicatorStartIndex2526-1 : (studentFTIMScheduleEIndicatorStartIndex2526-1)+studentFTIMScheduleEIndicatorLength2526]), // Field # 878

		StudentFTIMScheduleFIndicator: strings.TrimSpace(record[studentFTIMScheduleFIndicatorStartIndex2526-1 : (studentFTIMScheduleFIndicatorStartIndex2526-1)+studentFTIMScheduleFIndicatorLength2526]), // Field # 879

		StudentFTIMScheduleHIndicator: strings.TrimSpace(record[studentFTIMScheduleHIndicatorStartIndex2526-1 : (studentFTIMScheduleHIndicatorStartIndex2526-1)+studentFTIMScheduleHIndicatorLength2526]), // Field # 880

		StudentFTIMIRSResponseCode: strings.TrimSpace(record[studentFTIMIRSResponseCodeStartIndex2526-1 : (studentFTIMIRSResponseCodeStartIndex2526-1)+studentFTIMIRSResponseCodeLength2526]), // Field # 881

		StudentFTIMSpouseReturnedTaxYear: strings.TrimSpace(record[studentFTIMSpouseReturnedTaxYearStartIndex2526-1 : (studentFTIMSpouseReturnedTaxYearStartIndex2526-1)+studentFTIMSpouseReturnedTaxYearLength2526]), // Field # 882

		StudentFTIMSpouseFilingStatusCode: strings.TrimSpace(record[studentFTIMSpouseFilingStatusCodeStartIndex2526-1 : (studentFTIMSpouseFilingStatusCodeStartIndex2526-1)+studentFTIMSpouseFilingStatusCodeLength2526]), // Field # 883

		StudentFTIMSpouseAdjustedGrossIncome: strings.TrimSpace(record[studentFTIMSpouseAdjustedGrossIncomeStartIndex2526-1 : (studentFTIMSpouseAdjustedGrossIncomeStartIndex2526-1)+studentFTIMSpouseAdjustedGrossIncomeLength2526]), // Field # 884

		StudentFTIMSpouseNumberOfExemptions: strings.TrimSpace(record[studentFTIMSpouseNumberOfExemptionsStartIndex2526-1 : (studentFTIMSpouseNumberOfExemptionsStartIndex2526-1)+studentFTIMSpouseNumberOfExemptionsLength2526]), // Field # 885

		StudentFTIMSpouseNumberOfDependents: strings.TrimSpace(record[studentFTIMSpouseNumberOfDependentsStartIndex2526-1 : (studentFTIMSpouseNumberOfDependentsStartIndex2526-1)+studentFTIMSpouseNumberOfDependentsLength2526]), // Field # 886

		StudentFTIMSpouseTotalIncomeEarnedAmount: strings.TrimSpace(record[studentFTIMSpouseTotalIncomeEarnedAmountStartIndex2526-1 : (studentFTIMSpouseTotalIncomeEarnedAmountStartIndex2526-1)+studentFTIMSpouseTotalIncomeEarnedAmountLength2526]), // Field # 887

		StudentFTIMSpouseIncomeTaxPaid: strings.TrimSpace(record[studentFTIMSpouseIncomeTaxPaidStartIndex2526-1 : (studentFTIMSpouseIncomeTaxPaidStartIndex2526-1)+studentFTIMSpouseIncomeTaxPaidLength2526]), // Field # 888

		StudentFTIMSpouseEducationCredits: strings.TrimSpace(record[studentFTIMSpouseEducationCreditsStartIndex2526-1 : (studentFTIMSpouseEducationCreditsStartIndex2526-1)+studentFTIMSpouseEducationCreditsLength2526]), // Field # 889

		StudentFTIMSpouseUntaxedIRADistributions: strings.TrimSpace(record[studentFTIMSpouseUntaxedIRADistributionsStartIndex2526-1 : (studentFTIMSpouseUntaxedIRADistributionsStartIndex2526-1)+studentFTIMSpouseUntaxedIRADistributionsLength2526]), // Field # 890

		StudentFTIMSpouseIRADeductibleAndPayments: strings.TrimSpace(record[studentFTIMSpouseIRADeductibleAndPaymentsStartIndex2526-1 : (studentFTIMSpouseIRADeductibleAndPaymentsStartIndex2526-1)+studentFTIMSpouseIRADeductibleAndPaymentsLength2526]), // Field # 891

		StudentFTIMSpouseTaxExemptInterest: strings.TrimSpace(record[studentFTIMSpouseTaxExemptInterestStartIndex2526-1 : (studentFTIMSpouseTaxExemptInterestStartIndex2526-1)+studentFTIMSpouseTaxExemptInterestLength2526]), // Field # 892

		StudentFTIMSpouseUntaxedPensionsAmount: strings.TrimSpace(record[studentFTIMSpouseUntaxedPensionsAmountStartIndex2526-1 : (studentFTIMSpouseUntaxedPensionsAmountStartIndex2526-1)+studentFTIMSpouseUntaxedPensionsAmountLength2526]), // Field # 893

		StudentFTIMSpouseScheduleCNetProfitLoss: strings.TrimSpace(record[studentFTIMSpouseScheduleCNetProfitLossStartIndex2526-1 : (studentFTIMSpouseScheduleCNetProfitLossStartIndex2526-1)+studentFTIMSpouseScheduleCNetProfitLossLength2526]), // Field # 894

		StudentFTIMSpouseScheduleAIndicator: strings.TrimSpace(record[studentFTIMSpouseScheduleAIndicatorStartIndex2526-1 : (studentFTIMSpouseScheduleAIndicatorStartIndex2526-1)+studentFTIMSpouseScheduleAIndicatorLength2526]), // Field # 895

		StudentFTIMSpouseScheduleBIndicator: strings.TrimSpace(record[studentFTIMSpouseScheduleBIndicatorStartIndex2526-1 : (studentFTIMSpouseScheduleBIndicatorStartIndex2526-1)+studentFTIMSpouseScheduleBIndicatorLength2526]), // Field # 896

		StudentFTIMSpouseScheduleDIndicator: strings.TrimSpace(record[studentFTIMSpouseScheduleDIndicatorStartIndex2526-1 : (studentFTIMSpouseScheduleDIndicatorStartIndex2526-1)+studentFTIMSpouseScheduleDIndicatorLength2526]), // Field # 897

		StudentFTIMSpouseScheduleEIndicator: strings.TrimSpace(record[studentFTIMSpouseScheduleEIndicatorStartIndex2526-1 : (studentFTIMSpouseScheduleEIndicatorStartIndex2526-1)+studentFTIMSpouseScheduleEIndicatorLength2526]), // Field # 898

		StudentFTIMSpouseScheduleFIndicator: strings.TrimSpace(record[studentFTIMSpouseScheduleFIndicatorStartIndex2526-1 : (studentFTIMSpouseScheduleFIndicatorStartIndex2526-1)+studentFTIMSpouseScheduleFIndicatorLength2526]), // Field # 899

		StudentFTIMSpouseScheduleHIndicator: strings.TrimSpace(record[studentFTIMSpouseScheduleHIndicatorStartIndex2526-1 : (studentFTIMSpouseScheduleHIndicatorStartIndex2526-1)+studentFTIMSpouseScheduleHIndicatorLength2526]), // Field # 900

		StudentFTIMSpouseIRSResponseCode: strings.TrimSpace(record[studentFTIMSpouseIRSResponseCodeStartIndex2526-1 : (studentFTIMSpouseIRSResponseCodeStartIndex2526-1)+studentFTIMSpouseIRSResponseCodeLength2526]), // Field # 901

		ParentFTIMReturnedTaxYear: strings.TrimSpace(record[parentFTIMReturnedTaxYearStartIndex2526-1 : (parentFTIMReturnedTaxYearStartIndex2526-1)+parentFTIMReturnedTaxYearLength2526]), // Field # 902

		ParentFTIMFilingStatusCode: strings.TrimSpace(record[parentFTIMFilingStatusCodeStartIndex2526-1 : (parentFTIMFilingStatusCodeStartIndex2526-1)+parentFTIMFilingStatusCodeLength2526]), // Field # 903

		ParentFTIMAdjustedGrossIncome: strings.TrimSpace(record[parentFTIMAdjustedGrossIncomeStartIndex2526-1 : (parentFTIMAdjustedGrossIncomeStartIndex2526-1)+parentFTIMAdjustedGrossIncomeLength2526]), // Field # 904

		ParentFTIMNumberOfExemptions: strings.TrimSpace(record[parentFTIMNumberOfExemptionsStartIndex2526-1 : (parentFTIMNumberOfExemptionsStartIndex2526-1)+parentFTIMNumberOfExemptionsLength2526]), // Field # 905

		ParentFTIMNumberOfDependents: strings.TrimSpace(record[parentFTIMNumberOfDependentsStartIndex2526-1 : (parentFTIMNumberOfDependentsStartIndex2526-1)+parentFTIMNumberOfDependentsLength2526]), // Field # 906

		ParentFTIMTotalIncomeEarnedAmount: strings.TrimSpace(record[parentFTIMTotalIncomeEarnedAmountStartIndex2526-1 : (parentFTIMTotalIncomeEarnedAmountStartIndex2526-1)+parentFTIMTotalIncomeEarnedAmountLength2526]), // Field # 907

		ParentFTIMIncomeTaxPaid: strings.TrimSpace(record[parentFTIMIncomeTaxPaidStartIndex2526-1 : (parentFTIMIncomeTaxPaidStartIndex2526-1)+parentFTIMIncomeTaxPaidLength2526]), // Field # 908

		ParentFTIMEducationCredits: strings.TrimSpace(record[parentFTIMEducationCreditsStartIndex2526-1 : (parentFTIMEducationCreditsStartIndex2526-1)+parentFTIMEducationCreditsLength2526]), // Field # 909

		ParentFTIMUntaxedIRADistributions: strings.TrimSpace(record[parentFTIMUntaxedIRADistributionsStartIndex2526-1 : (parentFTIMUntaxedIRADistributionsStartIndex2526-1)+parentFTIMUntaxedIRADistributionsLength2526]), // Field # 910

		ParentFTIMIRADeductibleAndPayments: strings.TrimSpace(record[parentFTIMIRADeductibleAndPaymentsStartIndex2526-1 : (parentFTIMIRADeductibleAndPaymentsStartIndex2526-1)+parentFTIMIRADeductibleAndPaymentsLength2526]), // Field # 911

		ParentFTIMTaxExemptInterest: strings.TrimSpace(record[parentFTIMTaxExemptInterestStartIndex2526-1 : (parentFTIMTaxExemptInterestStartIndex2526-1)+parentFTIMTaxExemptInterestLength2526]), // Field # 912

		ParentFTIMUntaxedPensionsAmount: strings.TrimSpace(record[parentFTIMUntaxedPensionsAmountStartIndex2526-1 : (parentFTIMUntaxedPensionsAmountStartIndex2526-1)+parentFTIMUntaxedPensionsAmountLength2526]), // Field # 913

		ParentFTIMScheduleCNetProfitLoss: strings.TrimSpace(record[parentFTIMScheduleCNetProfitLossStartIndex2526-1 : (parentFTIMScheduleCNetProfitLossStartIndex2526-1)+parentFTIMScheduleCNetProfitLossLength2526]), // Field # 914

		ParentFTIMScheduleAIndicator: strings.TrimSpace(record[parentFTIMScheduleAIndicatorStartIndex2526-1 : (parentFTIMScheduleAIndicatorStartIndex2526-1)+parentFTIMScheduleAIndicatorLength2526]), // Field # 915

		ParentFTIMScheduleBIndicator: strings.TrimSpace(record[parentFTIMScheduleBIndicatorStartIndex2526-1 : (parentFTIMScheduleBIndicatorStartIndex2526-1)+parentFTIMScheduleBIndicatorLength2526]), // Field # 916

		ParentFTIMScheduleDIndicator: strings.TrimSpace(record[parentFTIMScheduleDIndicatorStartIndex2526-1 : (parentFTIMScheduleDIndicatorStartIndex2526-1)+parentFTIMScheduleDIndicatorLength2526]), // Field # 917

		ParentFTIMScheduleEIndicator: strings.TrimSpace(record[parentFTIMScheduleEIndicatorStartIndex2526-1 : (parentFTIMScheduleEIndicatorStartIndex2526-1)+parentFTIMScheduleEIndicatorLength2526]), // Field # 918

		ParentFTIMScheduleFIndicator: strings.TrimSpace(record[parentFTIMScheduleFIndicatorStartIndex2526-1 : (parentFTIMScheduleFIndicatorStartIndex2526-1)+parentFTIMScheduleFIndicatorLength2526]), // Field # 919

		ParentFTIMScheduleHIndicator: strings.TrimSpace(record[parentFTIMScheduleHIndicatorStartIndex2526-1 : (parentFTIMScheduleHIndicatorStartIndex2526-1)+parentFTIMScheduleHIndicatorLength2526]), // Field # 920

		ParentFTIMIRSResponseCode: strings.TrimSpace(record[parentFTIMIRSResponseCodeStartIndex2526-1 : (parentFTIMIRSResponseCodeStartIndex2526-1)+parentFTIMIRSResponseCodeLength2526]), // Field # 921

		ParentFTIMSpouseReturnedTaxYear: strings.TrimSpace(record[parentFTIMSpouseReturnedTaxYearStartIndex2526-1 : (parentFTIMSpouseReturnedTaxYearStartIndex2526-1)+parentFTIMSpouseReturnedTaxYearLength2526]), // Field # 922

		ParentFTIMSpouseFilingStatusCode: strings.TrimSpace(record[parentFTIMSpouseFilingStatusCodeStartIndex2526-1 : (parentFTIMSpouseFilingStatusCodeStartIndex2526-1)+parentFTIMSpouseFilingStatusCodeLength2526]), // Field # 923

		ParentFTIMSpouseAdjustedGrossIncome: strings.TrimSpace(record[parentFTIMSpouseAdjustedGrossIncomeStartIndex2526-1 : (parentFTIMSpouseAdjustedGrossIncomeStartIndex2526-1)+parentFTIMSpouseAdjustedGrossIncomeLength2526]), // Field # 924

		ParentFTIMSpouseNumberOfExemptions: strings.TrimSpace(record[parentFTIMSpouseNumberOfExemptionsStartIndex2526-1 : (parentFTIMSpouseNumberOfExemptionsStartIndex2526-1)+parentFTIMSpouseNumberOfExemptionsLength2526]), // Field # 925

		ParentFTIMSpouseNumberOfDependents: strings.TrimSpace(record[parentFTIMSpouseNumberOfDependentsStartIndex2526-1 : (parentFTIMSpouseNumberOfDependentsStartIndex2526-1)+parentFTIMSpouseNumberOfDependentsLength2526]), // Field # 926

		ParentFTIMSpouseTotalIncomeEarnedAmount: strings.TrimSpace(record[parentFTIMSpouseTotalIncomeEarnedAmountStartIndex2526-1 : (parentFTIMSpouseTotalIncomeEarnedAmountStartIndex2526-1)+parentFTIMSpouseTotalIncomeEarnedAmountLength2526]), // Field # 927

		ParentFTIMSpouseIncomeTaxPaid: strings.TrimSpace(record[parentFTIMSpouseIncomeTaxPaidStartIndex2526-1 : (parentFTIMSpouseIncomeTaxPaidStartIndex2526-1)+parentFTIMSpouseIncomeTaxPaidLength2526]), // Field # 928

		ParentFTIMSpouseEducationCredits: strings.TrimSpace(record[parentFTIMSpouseEducationCreditsStartIndex2526-1 : (parentFTIMSpouseEducationCreditsStartIndex2526-1)+parentFTIMSpouseEducationCreditsLength2526]), // Field # 929

		ParentFTIMSpouseUntaxedIRADistributions: strings.TrimSpace(record[parentFTIMSpouseUntaxedIRADistributionsStartIndex2526-1 : (parentFTIMSpouseUntaxedIRADistributionsStartIndex2526-1)+parentFTIMSpouseUntaxedIRADistributionsLength2526]), // Field # 930

		ParentFTIMSpouseIRADeductibleAndPayments: strings.TrimSpace(record[parentFTIMSpouseIRADeductibleAndPaymentsStartIndex2526-1 : (parentFTIMSpouseIRADeductibleAndPaymentsStartIndex2526-1)+parentFTIMSpouseIRADeductibleAndPaymentsLength2526]), // Field # 931

		ParentFTIMSpouseTaxExemptInterest: strings.TrimSpace(record[parentFTIMSpouseTaxExemptInterestStartIndex2526-1 : (parentFTIMSpouseTaxExemptInterestStartIndex2526-1)+parentFTIMSpouseTaxExemptInterestLength2526]), // Field # 932

		ParentFTIMSpouseUntaxedPensionsAmount: strings.TrimSpace(record[parentFTIMSpouseUntaxedPensionsAmountStartIndex2526-1 : (parentFTIMSpouseUntaxedPensionsAmountStartIndex2526-1)+parentFTIMSpouseUntaxedPensionsAmountLength2526]), // Field # 933

		ParentFTIMSpouseScheduleCNetProfitLoss: strings.TrimSpace(record[parentFTIMSpouseScheduleCNetProfitLossStartIndex2526-1 : (parentFTIMSpouseScheduleCNetProfitLossStartIndex2526-1)+parentFTIMSpouseScheduleCNetProfitLossLength2526]), // Field # 934

		ParentFTIMSpouseScheduleAIndicator: strings.TrimSpace(record[parentFTIMSpouseScheduleAIndicatorStartIndex2526-1 : (parentFTIMSpouseScheduleAIndicatorStartIndex2526-1)+parentFTIMSpouseScheduleAIndicatorLength2526]), // Field # 935

		ParentFTIMSpouseScheduleBIndicator: strings.TrimSpace(record[parentFTIMSpouseScheduleBIndicatorStartIndex2526-1 : (parentFTIMSpouseScheduleBIndicatorStartIndex2526-1)+parentFTIMSpouseScheduleBIndicatorLength2526]), // Field # 936

		ParentFTIMSpouseScheduleDIndicator: strings.TrimSpace(record[parentFTIMSpouseScheduleDIndicatorStartIndex2526-1 : (parentFTIMSpouseScheduleDIndicatorStartIndex2526-1)+parentFTIMSpouseScheduleDIndicatorLength2526]), // Field # 937

		ParentFTIMSpouseScheduleEIndicator: strings.TrimSpace(record[parentFTIMSpouseScheduleEIndicatorStartIndex2526-1 : (parentFTIMSpouseScheduleEIndicatorStartIndex2526-1)+parentFTIMSpouseScheduleEIndicatorLength2526]), // Field # 938

		ParentFTIMSpouseScheduleFIndicator: strings.TrimSpace(record[parentFTIMSpouseScheduleFIndicatorStartIndex2526-1 : (parentFTIMSpouseScheduleFIndicatorStartIndex2526-1)+parentFTIMSpouseScheduleFIndicatorLength2526]), // Field # 939

		ParentFTIMSpouseScheduleHIndicator: strings.TrimSpace(record[parentFTIMSpouseScheduleHIndicatorStartIndex2526-1 : (parentFTIMSpouseScheduleHIndicatorStartIndex2526-1)+parentFTIMSpouseScheduleHIndicatorLength2526]), // Field # 940

		ParentFTIMSpouseIRSResponseCode: strings.TrimSpace(record[parentFTIMSpouseIRSResponseCodeStartIndex2526-1 : (parentFTIMSpouseIRSResponseCodeStartIndex2526-1)+parentFTIMSpouseIRSResponseCodeLength2526]), // Field # 941

		FTILabelEnd: strings.TrimSpace(record[ftiLabelEndStartIndex2526-1 : (ftiLabelEndStartIndex2526-1)+ftiLabelEndLength2526]), // Field # 942

		StudentTotalIncome: strings.TrimSpace(record[studentTotalIncomeStartIndex2526-1 : (studentTotalIncomeStartIndex2526-1)+studentTotalIncomeLength2526]), // Field # 944

		ParentTotalIncome: strings.TrimSpace(record[parentTotalIncomeStartIndex2526-1 : (parentTotalIncomeStartIndex2526-1)+parentTotalIncomeLength2526]), // Field # 945

		FISAPTotalIncome: strings.TrimSpace(record[fisapTotalIncomeStartIndex2526-1 : (fisapTotalIncomeStartIndex2526-1)+fisapTotalIncomeLength2526]), // Field # 946

	}

	return r, nil
}

func parseISIRDate2526(s string) time.Time {
	parsedDate, err := time.Parse(isirDateLayout2526, s)

	if err != nil {
		return time.Time{}
	}

	return parsedDate
}

func parseISIRDateShort2526(s string) time.Time {
	parsedDate, err := time.Parse(isirDateShortLayout2526, s)

	if err != nil {
		return time.Time{}
	}

	return parsedDate
}
