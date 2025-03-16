package isirconv

import (
	"errors"
	"fmt"
	"github.com/rpatton4/fsa/pkg/isirmodels"
	"log/slog"
	"strings"
	"time"
)

const isirDateLayout = "20060102"

const totalISIRLength int = 7704

// Field # 1
const yearIndicatorStartIndex int = 1
const yearIndicatorLength int = 1

// Field # 2
const fafsaUUIDStartIndex int = 2
const fafsaUUIDLength int = 36

// Field # 3
const transactionUUIDStartIndex int = 38
const transactionUUIDLength int = 36

// Field # 4
const personUUIDStartIndex int = 74
const personUUIDLength int = 36

// Field # 5
const transactionNumberStartIndex int = 110
const transactionNumberLength int = 2

// Field # 6
const dependencyModelStartIndex int = 112
const dependencyModelLength int = 1

// Field # 7
const applicationSourceStartIndex int = 113
const applicationSourceLength int = 1

// Field # 8
const applicationReceiptDateStartIndex int = 114
const applicationReceiptDateLength int = 8

// Field # 9
const transactionSourceStartIndex int = 122
const transactionSourceLength int = 1

// Field # 10
const transactionTypeStartIndex int = 123
const transactionTypeLength int = 1

// Field # 11
const transactionLanguageStartIndex int = 124
const transactionLanguageLength int = 1

// Field # 12
const transactionReceiptDateStartIndex int = 125
const transactionReceiptDateLength int = 8

// Field # 13
const transactionProcessedDateStartIndex int = 133
const transactionProcessedDateLength int = 8

// Field # 14
const transactionStatusStartIndex int = 141
const transactionStatusLength int = 30

// Field # 15
const renewalDataUsedStartIndex int = 171
const renewalDataUsedLength int = 3

// Field # 16
const fpsCorrectionReasonStartIndex int = 174
const fpsCorrectionReasonLength int = 1

// Field # 17
const saiChangeFlagStartIndex int = 175
const saiChangeFlagLength int = 1

// Field # 18
const saiStartIndex int = 176
const saiLength int = 6

// Field # 19
const provisionalSAIStartIndex int = 182
const provisionalSAILength int = 6

// Field # 20
const saiFormulaStartIndex int = 188
const saiFormulaLength int = 1

// Field # 21
const saiComputationTypeStartIndex int = 189
const saiComputationTypeLength int = 2

// Field # 22
const maxPellIndicatorStartIndex int = 191
const maxPellIndicatorLength int = 1

// Field # 23
const minimumPellIndicatorStartIndex int = 192
const minimumPellIndicatorLength int = 1

// Field # 25
const studentFirstNameStartIndex int = 243
const studentFirstNameLength int = 35

// Field # 26
const studentMiddleNameStartIndex int = 278
const studentMiddleNameLength int = 15

// Field # 27
const studentLastNameStartIndex int = 293
const studentLastNameLength int = 35

// Field # 28
const studentSuffixStartIndex int = 328
const studentSuffixLength int = 10

// Field # 29
const studentDateOfBirthStartIndex int = 338
const studentDateOfBirthLength int = 8

// Field # 30
const studentSSNStartIndex int = 346
const studentSSNLength int = 9

// Field # 31
const studentITINStartIndex int = 355
const studentITINLength int = 9

// Field # 32
const studentPhoneNumberStartIndex int = 364
const studentPhoneNumberLength int = 10

// Field # 33
const studentEmailAddressStartIndex int = 374
const studentEmailAddressLength int = 50

// Field # 34
const studentStreetAddressStartIndex int = 424
const studentStreetAddressLength int = 40

// Field # 35
const studentCityStartIndex int = 464
const studentCityLength int = 30

// Field # 36
const studentStateStartIndex int = 494
const studentStateLength int = 2

// Field # 37
const studentZipCodeStartIndex int = 496
const studentZipCodeLength int = 10

// Field # 38
const studentCountryStartIndex int = 506
const studentCountryLength int = 2

// Field # 40
const studentMaritalStatusStartIndex int = 558
const studentMaritalStatusLength int = 1

// Field # 41
const studentGradeLevelStartIndex int = 559
const studentGradeLevelLength int = 1

// Field # 42
const studentFirstBachelorsDegreeBefore2526StartIndex int = 560
const studentFirstBachelorsDegreeBefore2526Length int = 1

// Field # 43
const studentPursuingTeacherCertificationStartIndex int = 561
const studentPursuingTeacherCertificationLength int = 1

// Field # 44
const studentActiveDutyStartIndex int = 562
const studentActiveDutyLength int = 1

// Field # 45
const studentVeteranStartIndex int = 563
const studentVeteranLength int = 1

// Field # 46
const studentChildOrOtherDependentsStartIndex int = 564
const studentChildOrOtherDependentsLength int = 1

// Field # 47
const studentParentsDeceasedStartIndex int = 565
const studentParentsDeceasedLength int = 1

// Field # 48
const studentWardOfCourtStartIndex int = 566
const studentWardOfCourtLength int = 1

// Field # 49
const studentInFosterCareStartIndex int = 567
const studentInFosterCareLength int = 1

// Field # 50
const studentEmancipatedMinorStartIndex int = 568
const studentEmancipatedMinorLength int = 1

// Field # 51
const studentLegalGuardianshipStartIndex int = 569
const studentLegalGuardianshipLength int = 1

// Field # 52
const studentPersonalCircumstancesNoneOfTheAboveStartIndex int = 570
const studentPersonalCircumstancesNoneOfTheAboveLength int = 1

// Field # 53
const studentUnaccompaniedHomelessYouthAndSelfSupportingStartIndex int = 571
const studentUnaccompaniedHomelessYouthAndSelfSupportingLength int = 1

// Field # 54
const studentUnaccompaniedHomelessGeneralStartIndex int = 572
const studentUnaccompaniedHomelessGeneralLength int = 1

// Field # 55
const studentUnaccompaniedHomelessHSStartIndex int = 573
const studentUnaccompaniedHomelessHSLength int = 1

// Field # 56
const studentUnaccompaniedHomelessTRIOStartIndex int = 574
const studentUnaccompaniedHomelessTRIOLength int = 1

// Field # 57
const studentUnaccompaniedHomelessFAAStartIndex int = 575
const studentUnaccompaniedHomelessFAALength int = 1

// Field # 58
const studentHomelessnessNoneOfTheAboveStartIndex int = 576
const studentHomelessnessNoneOfTheAboveLength int = 1

// Field # 59
const studentUnusualCircumstanceStartIndex int = 577
const studentUnusualCircumstanceLength int = 1

// Field # 60
const studentUnsubOnlyStartIndex int = 578
const studentUnsubOnlyLength int = 1

// Field # 61
const studentUpdatedFamilySizeStartIndex int = 579
const studentUpdatedFamilySizeLength int = 2

// Field # 62
const studentNumberInCollegeStartIndex int = 581
const studentNumberInCollegeLength int = 2

// Field # 63
const studentCitizenshipStatusStartIndex int = 583
const studentCitizenshipStatusLength int = 1

// Field # 64
const studentANumberStartIndex int = 584
const studentANumberLength int = 9

// Field # 65
const studentStateOfLegalResidenceStartIndex int = 593
const studentStateOfLegalResidenceLength int = 2

// Field # 66
const studentLegalResidenceDateStartIndex int = 595
const studentLegalResidenceDateLength int = 6

// Field # 67
const studentEitherParentAttendCollegeStartIndex int = 601
const studentEitherParentAttendCollegeLength int = 1

// Field # 68
const studentParentKilledInTheLineOfDutyStartIndex int = 602
const studentParentKilledInTheLineOfDutyLength int = 1

// Field # 69
const studentHighSchoolCompletionStatusStartIndex int = 603
const studentHighSchoolCompletionStatusLength int = 1

// Field # 70
const studentHighSchoolNameStartIndex int = 604
const studentHighSchoolNameLength int = 60

// Field # 71
const studentHighSchoolCityStartIndex int = 664
const studentHighSchoolCityLength int = 28

// Field # 72
const studentHighSchoolStateStartIndex int = 692
const studentHighSchoolStateLength int = 2

// Field # 73
const studentHighSchoolEquivalentDiplomaNameStartIndex int = 694
const studentHighSchoolEquivalentDiplomaNameLength int = 1

// Field # 74
const studentHighSchoolEquivalentDiplomaStateStartIndex int = 695
const studentHighSchoolEquivalentDiplomaStateLength int = 2

// Field # 75
const studentManuallyEnteredReceivedEITCStartIndex int = 697
const studentManuallyEnteredReceivedEITCLength int = 1

// Field # 76
const studentManuallyEnteredReceivedFederalHousingAssistanceStartIndex int = 698
const studentManuallyEnteredReceivedFederalHousingAssistanceLength int = 1

// Field # 77
const studentManuallyEnteredReceivedFreeReducedPriceLunchStartIndex int = 699
const studentManuallyEnteredReceivedFreeReducedPriceLunchLength int = 1

// Field # 78
const studentManuallyEnteredReceivedMedicaidStartIndex int = 700
const studentManuallyEnteredReceivedMedicaidLength int = 1

// Field # 79
const studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanStartIndex int = 701
const studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanLength int = 1

// Field # 80
const studentManuallyEnteredReceivedSNAPStartIndex int = 702
const studentManuallyEnteredReceivedSNAPLength int = 1

// Field # 81
const studentManuallyEnteredReceivedSupplementalSecurityIncomeStartIndex int = 703
const studentManuallyEnteredReceivedSupplementalSecurityIncomeLength int = 1

// Field # 82
const studentManuallyEnteredReceivedTANFStartIndex int = 704
const studentManuallyEnteredReceivedTANFLength int = 1

// Field # 83
const studentManuallyEnteredReceivedWICStartIndex int = 705
const studentManuallyEnteredReceivedWICLength int = 1

// Field # 84
const studentManuallyEnteredFederalBenefitsNoneOfTheAboveStartIndex int = 706
const studentManuallyEnteredFederalBenefitsNoneOfTheAboveLength int = 1

// Field # 85
const studentManuallyEnteredFiled1040Or1040NRStartIndex int = 707
const studentManuallyEnteredFiled1040Or1040NRLength int = 1

// Field # 86
const studentManuallyEnteredFiledNonUSTaxReturnStartIndex int = 708
const studentManuallyEnteredFiledNonUSTaxReturnLength int = 1

// Field # 87
const studentManuallyEnteredFiledJointReturnWithCurrentSpouseStartIndex int = 709
const studentManuallyEnteredFiledJointReturnWithCurrentSpouseLength int = 1

// Field # 88
const studentManuallyEnteredTaxReturnFilingStatusStartIndex int = 710
const studentManuallyEnteredTaxReturnFilingStatusLength int = 1

// Field # 89
const studentManuallyEnteredIncomeEarnedFromWorkStartIndex int = 711
const studentManuallyEnteredIncomeEarnedFromWorkLength int = 11

// Field # 90
const studentManuallyEnteredTaxExemptInterestIncomeStartIndex int = 722
const studentManuallyEnteredTaxExemptInterestIncomeLength int = 11

// Field # 91
const studentManuallyEnteredUntaxedPortionsOfIRADistributionsStartIndex int = 733
const studentManuallyEnteredUntaxedPortionsOfIRADistributionsLength int = 11

// Field # 92
const studentManuallyEnteredIRARolloverStartIndex int = 744
const studentManuallyEnteredIRARolloverLength int = 11

// Field # 93
const studentManuallyEnteredUntaxedPortionsOfPensionsStartIndex int = 755
const studentManuallyEnteredUntaxedPortionsOfPensionsLength int = 11

// Field # 94
const studentManuallyEnteredPensionRolloverStartIndex int = 766
const studentManuallyEnteredPensionRolloverLength int = 11

// Field # 95
const studentManuallyEnteredAdjustedGrossIncomeStartIndex int = 777
const studentManuallyEnteredAdjustedGrossIncomeLength int = 10

// Field # 96
const studentManuallyEnteredIncomeTaxPaidStartIndex int = 787
const studentManuallyEnteredIncomeTaxPaidLength int = 9

// Field # 97
const studentManuallyEnteredEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex int = 796
const studentManuallyEnteredEarnedIncomeTaxCreditReceivedDuringTaxYearLength int = 1

// Field # 98
const studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherStartIndex int = 797
const studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherLength int = 11

// Field # 99
const studentManuallyEnteredEducationCreditsStartIndex int = 808
const studentManuallyEnteredEducationCreditsLength int = 9

// Field # 100
const studentManuallyEnteredFiledScheduleABDEFHStartIndex int = 817
const studentManuallyEnteredFiledScheduleABDEFHLength int = 1

// Field # 101
const studentManuallyEnteredScheduleCAmountStartIndex int = 818
const studentManuallyEnteredScheduleCAmountLength int = 12

// Field # 102
const studentManuallyEnteredCollegeGrantAndScholarshipAidStartIndex int = 830
const studentManuallyEnteredCollegeGrantAndScholarshipAidLength int = 7

// Field # 103
const studentManuallyEnteredForeignEarnedIncomeExclusionStartIndex int = 837
const studentManuallyEnteredForeignEarnedIncomeExclusionLength int = 10

// Field # 104
const studentManuallyEnteredChildSupportReceivedStartIndex int = 847
const studentManuallyEnteredChildSupportReceivedLength int = 7

// Field # 105
const studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsStartIndex int = 854
const studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsLength int = 7

// Field # 106
const studentManuallyEnteredNetWorthOfCurrentInvestmentsStartIndex int = 861
const studentManuallyEnteredNetWorthOfCurrentInvestmentsLength int = 7

// Field # 107
const studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsStartIndex int = 868
const studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsLength int = 7

// Field # 108
const studentCollege1StartIndex int = 875
const studentCollege1Length int = 6

// Field # 109
const studentCollege2StartIndex int = 881
const studentCollege2Length int = 6

// Field # 110
const studentCollege3StartIndex int = 887
const studentCollege3Length int = 6

// Field # 111
const studentCollege4StartIndex int = 893
const studentCollege4Length int = 6

// Field # 112
const studentCollege5StartIndex int = 899
const studentCollege5Length int = 6

// Field # 113
const studentCollege6StartIndex int = 905
const studentCollege6Length int = 6

// Field # 114
const studentCollege7StartIndex int = 911
const studentCollege7Length int = 6

// Field # 115
const studentCollege8StartIndex int = 917
const studentCollege8Length int = 6

// Field # 116
const studentCollege9StartIndex int = 923
const studentCollege9Length int = 6

// Field # 117
const studentCollege10StartIndex int = 929
const studentCollege10Length int = 6

// Field # 118
const studentCollege11StartIndex int = 935
const studentCollege11Length int = 6

// Field # 119
const studentCollege12StartIndex int = 941
const studentCollege12Length int = 6

// Field # 120
const studentCollege13StartIndex int = 947
const studentCollege13Length int = 6

// Field # 121
const studentCollege14StartIndex int = 953
const studentCollege14Length int = 6

// Field # 122
const studentCollege15StartIndex int = 959
const studentCollege15Length int = 6

// Field # 123
const studentCollege16StartIndex int = 965
const studentCollege16Length int = 6

// Field # 124
const studentCollege17StartIndex int = 971
const studentCollege17Length int = 6

// Field # 125
const studentCollege18StartIndex int = 977
const studentCollege18Length int = 6

// Field # 126
const studentCollege19StartIndex int = 983
const studentCollege19Length int = 6

// Field # 127
const studentCollege20StartIndex int = 989
const studentCollege20Length int = 6

// Field # 128
const studentConsentToRetrieveAndDiscloseFTIStartIndex int = 995
const studentConsentToRetrieveAndDiscloseFTILength int = 1

// Field # 129
const studentSignatureStartIndex int = 996
const studentSignatureLength int = 1

// Field # 130
const studentSignatureDateStartIndex int = 997
const studentSignatureDateLength int = 8

// Field # 132
const studentSpouseFirstNameStartIndex int = 1055
const studentSpouseFirstNameLength int = 35

// Field # 133
const studentSpouseMiddleNameStartIndex int = 1090
const studentSpouseMiddleNameLength int = 15

// Field # 134
const studentSpouseLastNameStartIndex int = 1105
const studentSpouseLastNameLength int = 35

// Field # 135
const studentSpouseSuffixStartIndex int = 1140
const studentSpouseSuffixLength int = 10

// Field # 136
const studentSpouseDateOfBirthStartIndex int = 1150
const studentSpouseDateOfBirthLength int = 8

// Field # 137
const studentSpouseSSNStartIndex int = 1158
const studentSpouseSSNLength int = 9

// Field # 138
const studentSpouseITINStartIndex int = 1167
const studentSpouseITINLength int = 9

// Field # 139
const studentSpousePhoneNumberStartIndex int = 1176
const studentSpousePhoneNumberLength int = 10

// Field # 140
const studentSpouseEmailAddressStartIndex int = 1186
const studentSpouseEmailAddressLength int = 50

// Field # 141
const studentSpouseStreetAddressStartIndex int = 1236
const studentSpouseStreetAddressLength int = 40

// Field # 142
const studentSpouseCityStartIndex int = 1276
const studentSpouseCityLength int = 30

// Field # 143
const studentSpouseStateStartIndex int = 1306
const studentSpouseStateLength int = 2

// Field # 144
const studentSpouseZipCodeStartIndex int = 1308
const studentSpouseZipCodeLength int = 10

// Field # 145
const studentSpouseCountryStartIndex int = 1318
const studentSpouseCountryLength int = 2

// Field # 146
const studentSpouseFiled1040Or1040NRStartIndex int = 1320
const studentSpouseFiled1040Or1040NRLength int = 1

// Field # 147
const studentSpouseFiledNonUSTaxReturnStartIndex int = 1321
const studentSpouseFiledNonUSTaxReturnLength int = 1

// Field # 148
const studentSpouseTaxReturnFilingStatusStartIndex int = 1322
const studentSpouseTaxReturnFilingStatusLength int = 1

// Field # 149
const studentSpouseIncomeEarnedFromWorkStartIndex int = 1323
const studentSpouseIncomeEarnedFromWorkLength int = 11

// Field # 150
const studentSpouseTaxExemptInterestIncomeStartIndex int = 1334
const studentSpouseTaxExemptInterestIncomeLength int = 11

// Field # 151
const studentSpouseUntaxedPortionsOfIRADistributionsStartIndex int = 1345
const studentSpouseUntaxedPortionsOfIRADistributionsLength int = 11

// Field # 152
const studentSpouseIRARolloverStartIndex int = 1356
const studentSpouseIRARolloverLength int = 11

// Field # 153
const studentSpouseUntaxedPortionsOfPensionsStartIndex int = 1367
const studentSpouseUntaxedPortionsOfPensionsLength int = 11

// Field # 154
const studentSpousePensionRolloverStartIndex int = 1378
const studentSpousePensionRolloverLength int = 11

// Field # 155
const studentSpouseAdjustedGrossIncomeStartIndex int = 1389
const studentSpouseAdjustedGrossIncomeLength int = 10

// Field # 156
const studentSpouseIncomeTaxPaidStartIndex int = 1399
const studentSpouseIncomeTaxPaidLength int = 9

// Field # 157
const studentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex int = 1408
const studentSpouseDeductiblePaymentsToIRAKeoghOtherLength int = 11

// Field # 158
const studentSpouseEducationCreditsStartIndex int = 1419
const studentSpouseEducationCreditsLength int = 9

// Field # 159
const studentSpouseFiledScheduleABDEFHStartIndex int = 1428
const studentSpouseFiledScheduleABDEFHLength int = 1

// Field # 160
const studentSpouseScheduleCAmountStartIndex int = 1429
const studentSpouseScheduleCAmountLength int = 12

// Field # 161
const studentSpouseForeignEarnedIncomeExclusionStartIndex int = 1441
const studentSpouseForeignEarnedIncomeExclusionLength int = 10

// Field # 162
const studentSpouseConsentToRetrieveAndDiscloseFTIStartIndex int = 1451
const studentSpouseConsentToRetrieveAndDiscloseFTILength int = 1

// Field # 163
const studentSpouseSignatureStartIndex int = 1452
const studentSpouseSignatureLength int = 1

// Field # 164
const studentSpouseSignatureDateStartIndex int = 1453
const studentSpouseSignatureDateLength int = 8

// Field # 166
const parentFirstNameStartIndex int = 1511
const parentFirstNameLength int = 35

// Field # 167
const parentMiddleNameStartIndex int = 1546
const parentMiddleNameLength int = 15

// Field # 168
const parentLastNameStartIndex int = 1561
const parentLastNameLength int = 35

// Field # 169
const parentSuffixStartIndex int = 1596
const parentSuffixLength int = 10

// Field # 170
const parentDateOfBirthStartIndex int = 1606
const parentDateOfBirthLength int = 8

// Field # 171
const parentSSNStartIndex int = 1614
const parentSSNLength int = 9

// Field # 172
const parentITINStartIndex int = 1623
const parentITINLength int = 9

// Field # 173
const parentPhoneNumberStartIndex int = 1632
const parentPhoneNumberLength int = 10

// Field # 174
const parentEmailAddressStartIndex int = 1642
const parentEmailAddressLength int = 50

// Field # 175
const parentStreetAddressStartIndex int = 1692
const parentStreetAddressLength int = 40

// Field # 176
const parentCityStartIndex int = 1732
const parentCityLength int = 30

// Field # 177
const parentStateStartIndex int = 1762
const parentStateLength int = 2

// Field # 178
const parentZipCodeStartIndex int = 1764
const parentZipCodeLength int = 10

// Field # 179
const parentCountryStartIndex int = 1774
const parentCountryLength int = 2

// Field # 180
const parentMaritalStatusStartIndex int = 1776
const parentMaritalStatusLength int = 1

// Field # 181
const parentStateOfLegalResidenceStartIndex int = 1777
const parentStateOfLegalResidenceLength int = 2

// Field # 182
const parentLegalResidenceDateStartIndex int = 1779
const parentLegalResidenceDateLength int = 6

// Field # 183
const parentUpdatedFamilySizeStartIndex int = 1785
const parentUpdatedFamilySizeLength int = 2

// Field # 184
const parentNumberInCollegeStartIndex int = 1787
const parentNumberInCollegeLength int = 2

// Field # 185
const parentReceivedEITCStartIndex int = 1789
const parentReceivedEITCLength int = 1

// Field # 186
const parentReceivedFederalHousingAssistanceStartIndex int = 1790
const parentReceivedFederalHousingAssistanceLength int = 1

// Field # 187
const parentReceivedFreeReducedPriceLunchStartIndex int = 1791
const parentReceivedFreeReducedPriceLunchLength int = 1

// Field # 188
const parentReceivedMedicaidStartIndex int = 1792
const parentReceivedMedicaidLength int = 1

// Field # 189
const parentReceivedRefundableCreditFor36BHealthPlanStartIndex int = 1793
const parentReceivedRefundableCreditFor36BHealthPlanLength int = 1

// Field # 190
const parentReceivedSNAPStartIndex int = 1794
const parentReceivedSNAPLength int = 1

// Field # 191
const parentReceivedSupplementalSecurityIncomeStartIndex int = 1795
const parentReceivedSupplementalSecurityIncomeLength int = 1

// Field # 192
const parentReceivedTANFStartIndex int = 1796
const parentReceivedTANFLength int = 1

// Field # 193
const parentReceivedWICStartIndex int = 1797
const parentReceivedWICLength int = 1

// Field # 194
const parentFederalBenefitsNoneOfTheAboveStartIndex int = 1798
const parentFederalBenefitsNoneOfTheAboveLength int = 1

// Field # 195
const parentFiled1040Or1040NRStartIndex int = 1799
const parentFiled1040Or1040NRLength int = 1

// Field # 196
const parentFileNonUSTaxReturnStartIndex int = 1800
const parentFileNonUSTaxReturnLength int = 1

// Field # 197
const parentFiledJointReturnWithCurrentSpouseStartIndex int = 1801
const parentFiledJointReturnWithCurrentSpouseLength int = 1

// Field # 198
const parentTaxReturnFilingStatusStartIndex int = 1802
const parentTaxReturnFilingStatusLength int = 1

// Field # 199
const parentIncomeEarnedFromWorkStartIndex int = 1803
const parentIncomeEarnedFromWorkLength int = 11

// Field # 200
const parentTaxExemptInterestIncomeStartIndex int = 1814
const parentTaxExemptInterestIncomeLength int = 11

// Field # 201
const parentUntaxedPortionsOfIRADistributionsStartIndex int = 1825
const parentUntaxedPortionsOfIRADistributionsLength int = 11

// Field # 202
const parentIRARolloverStartIndex int = 1836
const parentIRARolloverLength int = 11

// Field # 203
const parentUntaxedPortionsOfPensionsStartIndex int = 1847
const parentUntaxedPortionsOfPensionsLength int = 11

// Field # 204
const parentPensionRolloverStartIndex int = 1858
const parentPensionRolloverLength int = 11

// Field # 205
const parentAdjustedGrossIncomeStartIndex int = 1869
const parentAdjustedGrossIncomeLength int = 10

// Field # 206
const parentIncomeTaxPaidStartIndex int = 1879
const parentIncomeTaxPaidLength int = 9

// Field # 207
const parentEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex int = 1888
const parentEarnedIncomeTaxCreditReceivedDuringTaxYearLength int = 1

// Field # 208
const parentDeductiblePaymentsToIRAKeoghOtherStartIndex int = 1889
const parentDeductiblePaymentsToIRAKeoghOtherLength int = 11

// Field # 209
const parentEducationCreditsStartIndex int = 1900
const parentEducationCreditsLength int = 9

// Field # 210
const parentFiledScheduleABDEFHStartIndex int = 1909
const parentFiledScheduleABDEFHLength int = 1

// Field # 211
const parentScheduleCAmountStartIndex int = 1910
const parentScheduleCAmountLength int = 12

// Field # 212
const parentCollegeGrantAndScholarshipAidStartIndex int = 1922
const parentCollegeGrantAndScholarshipAidLength int = 7

// Field # 213
const parentForeignEarnedIncomeExclusionStartIndex int = 1929
const parentForeignEarnedIncomeExclusionLength int = 10

// Field # 214
const parentChildSupportReceivedStartIndex int = 1939
const parentChildSupportReceivedLength int = 7

// Field # 215
const parentTotalOfCashSavingsAndCheckingAccountsStartIndex int = 1946
const parentTotalOfCashSavingsAndCheckingAccountsLength int = 7

// Field # 216
const parentNetWorthOfCurrentInvestmentsStartIndex int = 1953
const parentNetWorthOfCurrentInvestmentsLength int = 7

// Field # 217
const parentNetWorthOfBusinessesAndInvestmentFarmsStartIndex int = 1960
const parentNetWorthOfBusinessesAndInvestmentFarmsLength int = 7

// Field # 218
const parentConsentToRetrieveAndDiscloseFTIStartIndex int = 1967
const parentConsentToRetrieveAndDiscloseFTILength int = 1

// Field # 219
const parentSignatureStartIndex int = 1968
const parentSignatureLength int = 1

// Field # 220
const parentSignatureDateStartIndex int = 1969
const parentSignatureDateLength int = 8

// Field # 222
const parentSpouseFirstNameStartIndex int = 2027
const parentSpouseFirstNameLength int = 35

// Field # 223
const parentSpouseMiddleNameStartIndex int = 2062
const parentSpouseMiddleNameLength int = 15

// Field # 224
const parentSpouseLastNameStartIndex int = 2077
const parentSpouseLastNameLength int = 35

// Field # 225
const parentSpouseSuffixStartIndex int = 2112
const parentSpouseSuffixLength int = 10

// Field # 226
const parentSpouseDateOfBirthStartIndex int = 2122
const parentSpouseDateOfBirthLength int = 8

// Field # 227
const parentSpouseSSNStartIndex int = 2130
const parentSpouseSSNLength int = 9

// Field # 228
const parentSpouseITINStartIndex int = 2139
const parentSpouseITINLength int = 9

// Field # 229
const parentSpousePhoneNumberStartIndex int = 2148
const parentSpousePhoneNumberLength int = 10

// Field # 230
const parentSpouseEmailAddressStartIndex int = 2158
const parentSpouseEmailAddressLength int = 50

// Field # 231
const parentSpouseStreetAddressStartIndex int = 2208
const parentSpouseStreetAddressLength int = 40

// Field # 232
const parentSpouseCityStartIndex int = 2248
const parentSpouseCityLength int = 30

// Field # 233
const parentSpouseStateStartIndex int = 2278
const parentSpouseStateLength int = 2

// Field # 234
const parentSpouseZipCodeStartIndex int = 2280
const parentSpouseZipCodeLength int = 10

// Field # 235
const parentSpouseCountryStartIndex int = 2290
const parentSpouseCountryLength int = 2

// Field # 236
const parentSpouseFiled1040Or1040NRStartIndex int = 2292
const parentSpouseFiled1040Or1040NRLength int = 1

// Field # 237
const parentSpouseFileNonUSTaxReturnStartIndex int = 2293
const parentSpouseFileNonUSTaxReturnLength int = 1

// Field # 238
const parentSpouseTaxReturnFilingStatusStartIndex int = 2294
const parentSpouseTaxReturnFilingStatusLength int = 1

// Field # 239
const parentSpouseIncomeEarnedFromWorkStartIndex int = 2295
const parentSpouseIncomeEarnedFromWorkLength int = 11

// Field # 240
const parentSpouseTaxExemptInterestIncomeStartIndex int = 2306
const parentSpouseTaxExemptInterestIncomeLength int = 11

// Field # 241
const parentSpouseUntaxedPortionsOfIRADistributionsStartIndex int = 2317
const parentSpouseUntaxedPortionsOfIRADistributionsLength int = 11

// Field # 242
const parentSpouseIRARolloverStartIndex int = 2328
const parentSpouseIRARolloverLength int = 11

// Field # 243
const parentSpouseUntaxedPortionsOfPensionsStartIndex int = 2339
const parentSpouseUntaxedPortionsOfPensionsLength int = 11

// Field # 244
const parentSpousePensionRolloverStartIndex int = 2350
const parentSpousePensionRolloverLength int = 11

// Field # 245
const parentSpouseAdjustedGrossIncomeStartIndex int = 2361
const parentSpouseAdjustedGrossIncomeLength int = 10

// Field # 246
const parentSpouseIncomeTaxPaidStartIndex int = 2371
const parentSpouseIncomeTaxPaidLength int = 9

// Field # 247
const parentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex int = 2380
const parentSpouseDeductiblePaymentsToIRAKeoghOtherLength int = 11

// Field # 248
const parentSpouseEducationCreditsStartIndex int = 2391
const parentSpouseEducationCreditsLength int = 9

// Field # 249
const parentSpouseFiledScheduleABDEFHStartIndex int = 2400
const parentSpouseFiledScheduleABDEFHLength int = 1

// Field # 250
const parentSpouseScheduleCAmountStartIndex int = 2401
const parentSpouseScheduleCAmountLength int = 12

// Field # 251
const parentSpouseForeignEarnedIncomeExclusionStartIndex int = 2413
const parentSpouseForeignEarnedIncomeExclusionLength int = 10

// Field # 252
const parentSpouseConsentToRetrieveAndDiscloseFTIStartIndex int = 2423
const parentSpouseConsentToRetrieveAndDiscloseFTILength int = 1

// Field # 253
const parentSpouseSignatureStartIndex int = 2424
const parentSpouseSignatureLength int = 1

// Field # 254
const parentSpouseSignatureDateStartIndex int = 2425
const parentSpouseSignatureDateLength int = 8

// Field # 256
const preparerFirstNameStartIndex int = 2483
const preparerFirstNameLength int = 35

// Field # 257
const preparerLastNameStartIndex int = 2518
const preparerLastNameLength int = 35

// Field # 258
const preparerSSNStartIndex int = 2553
const preparerSSNLength int = 9

// Field # 259
const preparerEINStartIndex int = 2562
const preparerEINLength int = 9

// Field # 260
const preparerAffiliationStartIndex int = 2571
const preparerAffiliationLength int = 30

// Field # 261
const preparerStreetAddressStartIndex int = 2601
const preparerStreetAddressLength int = 40

// Field # 262
const preparerCityStartIndex int = 2641
const preparerCityLength int = 30

// Field # 263
const preparerStateStartIndex int = 2671
const preparerStateLength int = 2

// Field # 264
const preparerZipCodeStartIndex int = 2673
const preparerZipCodeLength int = 10

// Field # 265
const preparerSignatureStartIndex int = 2683
const preparerSignatureLength int = 1

// Field # 266
const preparerSignatureDateStartIndex int = 2684
const preparerSignatureDateLength int = 8

// Field # 268
const studentAffirmationStatusStartIndex int = 2742
const studentAffirmationStatusLength int = 1

// Field # 269
const studentSpouseAffirmationStatusStartIndex int = 2743
const studentSpouseAffirmationStatusLength int = 1

// Field # 270
const parentAffirmationStatusStartIndex int = 2744
const parentAffirmationStatusLength int = 1

// Field # 271
const parentSpouseOrPartnerAffirmationStatusStartIndex int = 2745
const parentSpouseOrPartnerAffirmationStatusLength int = 1

// Field # 272
const studentDateConsentGrantedStartIndex int = 2746
const studentDateConsentGrantedLength int = 8

// Field # 273
const studentSpouseDateConsentGrantedStartIndex int = 2754
const studentSpouseDateConsentGrantedLength int = 8

// Field # 274
const parentDateConsentGrantedStartIndex int = 2762
const parentDateConsentGrantedLength int = 8

// Field # 275
const parentSpouseOrPartnerDateConsentGrantedStartIndex int = 2770
const parentSpouseOrPartnerDateConsentGrantedLength int = 8

// Field # 276
const studentTransunionMatchStatusStartIndex int = 2778
const studentTransunionMatchStatusLength int = 1

// Field # 277
const studentSpouseTransunionMatchStatusStartIndex int = 2779
const studentSpouseTransunionMatchStatusLength int = 1

// Field # 278
const studentParentTransunionMatchStatusStartIndex int = 2780
const studentParentTransunionMatchStatusLength int = 1

// Field # 279
const studentParentSpouseTransunionMatchStatusStartIndex int = 2781
const studentParentSpouseTransunionMatchStatusLength int = 1

// Field # 280
const correctionAppliedAgainstTransactionNumberStartIndex int = 2782
const correctionAppliedAgainstTransactionNumberLength int = 2

// Field # 281
const professionalJudgementStartIndex int = 2784
const professionalJudgementLength int = 1

// Field # 282
const dependencyOverrideIndicatorStartIndex int = 2785
const dependencyOverrideIndicatorLength int = 1

// Field # 283
const fAAFederalSchoolCodeStartIndex int = 2786
const fAAFederalSchoolCodeLength int = 6

// Field # 284
const fAASignatureStartIndex int = 2792
const fAASignatureLength int = 1

// Field # 285
const iASGIndicatorStartIndex int = 2793
const iASGIndicatorLength int = 1

// Field # 286
const childrenOfFallenHeroesIndicatorStartIndex int = 2794
const childrenOfFallenHeroesIndicatorLength int = 1

// Field # 287
const electronicTransactionIndicatorDestinationNumberStartIndex int = 2795
const electronicTransactionIndicatorDestinationNumberLength int = 7

// Field # 288
const studentSignatureSourceStartIndex int = 2802
const studentSignatureSourceLength int = 1

// Field # 289
const studentSpouseSignatureSourceStartIndex int = 2803
const studentSpouseSignatureSourceLength int = 1

// Field # 290
const parentSignatureSourceStartIndex int = 2804
const parentSignatureSourceLength int = 1

// Field # 291
const parentSpouseOrPartnerSignatureSourceStartIndex int = 2805
const parentSpouseOrPartnerSignatureSourceLength int = 1

// Field # 292
const specialHandlingIndicatorStartIndex int = 2806
const specialHandlingIndicatorLength int = 1

// Field # 293
const addressOnlyChangeFlagStartIndex int = 2807
const addressOnlyChangeFlagLength int = 1

// Field # 294
const fpsPushedISIRFlagStartIndex int = 2808
const fpsPushedISIRFlagLength int = 1

// Field # 295
const rejectStatusChangeFlagStartIndex int = 2809
const rejectStatusChangeFlagLength int = 1

// Field # 296
const verificationTrackingFlagStartIndex int = 2810
const verificationTrackingFlagLength int = 2

// Field # 297
const studentSelectedForVerificationStartIndex int = 2812
const studentSelectedForVerificationLength int = 1

// Field # 298
const incarceratedApplicantFlagStartIndex int = 2813
const incarceratedApplicantFlagLength int = 1

// Field # 299
const nsldsTransactionNumberStartIndex int = 2814
const nsldsTransactionNumberLength int = 2

// Field # 300
const nsldsDatabaseResultsFlagStartIndex int = 2816
const nsldsDatabaseResultsFlagLength int = 1

// Field # 301
const highSchoolFlagStartIndex int = 2817
const highSchoolFlagLength int = 1

// Field # 302
const studentTotalFederalWorkStudyEarningsStartIndex int = 2818
const studentTotalFederalWorkStudyEarningsLength int = 12

// Field # 303
const studentSpouseTotalFederalWorkStudyEarningsStartIndex int = 2830
const studentSpouseTotalFederalWorkStudyEarningsLength int = 12

// Field # 304
const parentTotalFederalWorkStudyEarningsStartIndex int = 2842
const parentTotalFederalWorkStudyEarningsLength int = 12

// Field # 305
const parentSpouseOrPartnerTotalFederalWorkStudyEarningsStartIndex int = 2854
const parentSpouseOrPartnerTotalFederalWorkStudyEarningsLength int = 12

// Field # 306
const totalParentAllowancesAgainstIncomeStartIndex int = 2866
const totalParentAllowancesAgainstIncomeLength int = 15

// Field # 307
const parentPayrollTaxAllowanceStartIndex int = 2881
const parentPayrollTaxAllowanceLength int = 15

// Field # 308
const parentIncomeProtectionAllowanceStartIndex int = 2896
const parentIncomeProtectionAllowanceLength int = 15

// Field # 309
const parentEmploymentExpenseAllowanceStartIndex int = 2911
const parentEmploymentExpenseAllowanceLength int = 15

// Field # 310
const parentAvailableIncomeStartIndex int = 2926
const parentAvailableIncomeLength int = 15

// Field # 311
const parentAdjustedAvailableIncomeStartIndex int = 2941
const parentAdjustedAvailableIncomeLength int = 15

// Field # 312
const parentContributionStartIndex int = 2956
const parentContributionLength int = 15

// Field # 313
const studentPayrollTaxAllowanceStartIndex int = 2971
const studentPayrollTaxAllowanceLength int = 15

// Field # 314
const studentIncomeProtectionAllowanceStartIndex int = 2986
const studentIncomeProtectionAllowanceLength int = 15

// Field # 315
const studentAllowanceForParentsNegativeAdjustedAvailableIncomeStartIndex int = 3001
const studentAllowanceForParentsNegativeAdjustedAvailableIncomeLength int = 15

// Field # 316
const studentEmploymentExpenseAllowanceStartIndex int = 3016
const studentEmploymentExpenseAllowanceLength int = 15

// Field # 317
const totalStudentAllowancesAgainstIncomeStartIndex int = 3031
const totalStudentAllowancesAgainstIncomeLength int = 15

// Field # 318
const studentAvailableIncomeStartIndex int = 3046
const studentAvailableIncomeLength int = 15

// Field # 319
const studentContributionFromIncomeStartIndex int = 3061
const studentContributionFromIncomeLength int = 15

// Field # 320
const studentAdjustedAvailableIncomeStartIndex int = 3076
const studentAdjustedAvailableIncomeLength int = 15

// Field # 321
const totalStudentContributionFromSAAIStartIndex int = 3091
const totalStudentContributionFromSAAILength int = 15

// Field # 322
const parentDiscretionaryNetWorthStartIndex int = 3106
const parentDiscretionaryNetWorthLength int = 7

// Field # 323
const parentNetWorthStartIndex int = 3113
const parentNetWorthLength int = 7

// Field # 324
const parentAssetProtectionAllowanceStartIndex int = 3120
const parentAssetProtectionAllowanceLength int = 12

// Field # 325
const parentContributionFromAssetsStartIndex int = 3132
const parentContributionFromAssetsLength int = 12

// Field # 326
const studentNetWorthStartIndex int = 3144
const studentNetWorthLength int = 7

// Field # 327
const studentAssetProtectionAllowanceStartIndex int = 3151
const studentAssetProtectionAllowanceLength int = 12

// Field # 328
const studentContributionFromAssetsStartIndex int = 3163
const studentContributionFromAssetsLength int = 12

// Field # 329
const assumedStudentFamilySizeStartIndex int = 3175
const assumedStudentFamilySizeLength int = 3

// Field # 330
const assumedParentFamilySizeStartIndex int = 3178
const assumedParentFamilySizeLength int = 3

// Field # 331
const studentFirstNameCHVFlagsStartIndex int = 3181
const studentFirstNameCHVFlagsLength int = 3

// Field # 332
const studentMiddleNameCHVFlagsStartIndex int = 3184
const studentMiddleNameCHVFlagsLength int = 3

// Field # 333
const studentLastNameCHVFLagsStartIndex int = 3187
const studentLastNameCHVFLagsLength int = 3

// Field # 334
const studentSuffixCHVFLagsStartIndex int = 3190
const studentSuffixCHVFLagsLength int = 3

// Field # 335
const studentDateOfBirthCHVFLagsStartIndex int = 3193
const studentDateOfBirthCHVFLagsLength int = 3

// Field # 336
const studentSSNCHVFlagsStartIndex int = 3196
const studentSSNCHVFlagsLength int = 3

// Field # 337
const studentITINCHVFLagsStartIndex int = 3199
const studentITINCHVFLagsLength int = 3

// Field # 338
const studentPhoneNumberCHVFlagsStartIndex int = 3202
const studentPhoneNumberCHVFlagsLength int = 3

// Field # 339
const studentEmailAddressCHVFlagsStartIndex int = 3205
const studentEmailAddressCHVFlagsLength int = 3

// Field # 340
const studentStreetAddressCHVFlagsStartIndex int = 3208
const studentStreetAddressCHVFlagsLength int = 3

// Field # 341
const studentCityCHVFLagsStartIndex int = 3211
const studentCityCHVFLagsLength int = 3

// Field # 342
const studentStateCHVFlagsStartIndex int = 3214
const studentStateCHVFlagsLength int = 3

// Field # 343
const studentZipCodeCHVFlagsStartIndex int = 3217
const studentZipCodeCHVFlagsLength int = 3

// Field # 344
const studentCountryCHVFlagsStartIndex int = 3220
const studentCountryCHVFlagsLength int = 3

// Field # 345
const studentMaritalStatusCHVFlagsStartIndex int = 3223
const studentMaritalStatusCHVFlagsLength int = 3

// Field # 346
const studentGradeLevelInCollegeCHVFlagsStartIndex int = 3226
const studentGradeLevelInCollegeCHVFlagsLength int = 3

// Field # 347
const studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsStartIndex int = 3229
const studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsLength int = 3

// Field # 348
const studentPursuingTeacherCertificationCHVFlagsStartIndex int = 3232
const studentPursuingTeacherCertificationCHVFlagsLength int = 3

// Field # 349
const studentActiveDutyCHVFlagsStartIndex int = 3235
const studentActiveDutyCHVFlagsLength int = 3

// Field # 350
const studentVeteranCHVFlagsStartIndex int = 3238
const studentVeteranCHVFlagsLength int = 3

// Field # 351
const studentChildOrOtherDependentsCHVFlagsStartIndex int = 3241
const studentChildOrOtherDependentsCHVFlagsLength int = 3

// Field # 352
const studentParentsDeceasedCHVFlagsStartIndex int = 3244
const studentParentsDeceasedCHVFlagsLength int = 3

// Field # 353
const studentWardOfCourtCHVFlagsStartIndex int = 3247
const studentWardOfCourtCHVFlagsLength int = 3

// Field # 354
const studentInFosterCareCHVFlagsStartIndex int = 3250
const studentInFosterCareCHVFlagsLength int = 3

// Field # 355
const studentEmancipatedMinorCHVFlagsStartIndex int = 3253
const studentEmancipatedMinorCHVFlagsLength int = 3

// Field # 356
const studentLegalGuardianshipCHVFlagsStartIndex int = 3256
const studentLegalGuardianshipCHVFlagsLength int = 3

// Field # 357
const studentPersonalCircumstancesNoneOfTheAboveCHVFlagsStartIndex int = 3259
const studentPersonalCircumstancesNoneOfTheAboveCHVFlagsLength int = 3

// Field # 358
const studentUnaccompaniedHomelessYouthOrIsUnaccompaniedAtRIskSelfSupportingCHVFlagsStartIndex int = 3262
const studentUnaccompaniedHomelessYouthOrIsUnaccompaniedAtRIskSelfSupportingCHVFlagsLength int = 3

// Field # 359
const studentUnaccompaniedAndHomelessGeneralCHVFlagsStartIndex int = 3265
const studentUnaccompaniedAndHomelessGeneralCHVFlagsLength int = 3

// Field # 360
const studentUnaccompaniedAndHomelessHSCHVFlagsStartIndex int = 3268
const studentUnaccompaniedAndHomelessHSCHVFlagsLength int = 3

// Field # 361
const studentUnaccompaniedAndHomelessTRIOCHVFlagsStartIndex int = 3271
const studentUnaccompaniedAndHomelessTRIOCHVFlagsLength int = 3

// Field # 362
const studentUnaccompaniedAndHomelessFAACHVFlagsStartIndex int = 3274
const studentUnaccompaniedAndHomelessFAACHVFlagsLength int = 3

// Field # 363
const studentHomelessnessNoneOfTheAboveCHVFlagsStartIndex int = 3277
const studentHomelessnessNoneOfTheAboveCHVFlagsLength int = 3

// Field # 364
const studentHasUnusualCircumstanceCHVFlagsStartIndex int = 3280
const studentHasUnusualCircumstanceCHVFlagsLength int = 3

// Field # 365
const studentUnsubOnlyCHVFlagsStartIndex int = 3283
const studentUnsubOnlyCHVFlagsLength int = 3

// Field # 366
const studentUpdatedFamilySizeCHVFlagsStartIndex int = 3286
const studentUpdatedFamilySizeCHVFlagsLength int = 3

// Field # 367
const studentNumberInCollegeCorrectionCHVFlagsStartIndex int = 3289
const studentNumberInCollegeCorrectionCHVFlagsLength int = 3

// Field # 368
const studentCitizenshipStatusCorrectionCHVFlagsStartIndex int = 3292
const studentCitizenshipStatusCorrectionCHVFlagsLength int = 3

// Field # 369
const studentANumberCHVFlagsStartIndex int = 3295
const studentANumberCHVFlagsLength int = 3

// Field # 370
const studentStateOfLegalResidenceCHVFlagsStartIndex int = 3298
const studentStateOfLegalResidenceCHVFlagsLength int = 3

// Field # 371
const studentLegalResidenceDateCHVFlagsStartIndex int = 3301
const studentLegalResidenceDateCHVFlagsLength int = 3

// Field # 372
const studentEitherParentAttendCollegeCHVFlagsStartIndex int = 3304
const studentEitherParentAttendCollegeCHVFlagsLength int = 3

// Field # 373
const studentParentKilledInTheLineOfDutyCHVFlagsStartIndex int = 3307
const studentParentKilledInTheLineOfDutyCHVFlagsLength int = 3

// Field # 374
const studentHighSchoolCompletionStatusCHVFlagsStartIndex int = 3310
const studentHighSchoolCompletionStatusCHVFlagsLength int = 3

// Field # 375
const studentHighSchoolNameCHVFlagsStartIndex int = 3313
const studentHighSchoolNameCHVFlagsLength int = 3

// Field # 376
const studentHighSchoolCityCHVFlagsStartIndex int = 3316
const studentHighSchoolCityCHVFlagsLength int = 3

// Field # 377
const studentHighSchoolStateCHVFlagsStartIndex int = 3319
const studentHighSchoolStateCHVFlagsLength int = 3

// Field # 378
const studentHighSchoolEquivalentDiplomaNameCHVFlagsStartIndex int = 3322
const studentHighSchoolEquivalentDiplomaNameCHVFlagsLength int = 3

// Field # 379
const studentHighSchoolEquivalentDiplomaStateCHVFlagsStartIndex int = 3325
const studentHighSchoolEquivalentDiplomaStateCHVFlagsLength int = 3

// Field # 380
const studentReceivedEITCCHVFlagsStartIndex int = 3328
const studentReceivedEITCCHVFlagsLength int = 3

// Field # 381
const studentReceivedFederalHousingAssistanceCHVFlagsStartIndex int = 3331
const studentReceivedFederalHousingAssistanceCHVFlagsLength int = 3

// Field # 382
const studentReceivedFreeReducedPriceLunchCHVFlagsStartIndex int = 3334
const studentReceivedFreeReducedPriceLunchCHVFlagsLength int = 3

// Field # 383
const studentReceivedMedicaidCHVFlagsStartIndex int = 3337
const studentReceivedMedicaidCHVFlagsLength int = 3

// Field # 384
const studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex int = 3340
const studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsLength int = 3

// Field # 385
const studentReceivedSNAPCHVFlagsStartIndex int = 3343
const studentReceivedSNAPCHVFlagsLength int = 3

// Field # 386
const studentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex int = 3346
const studentReceivedSupplementalSecurityIncomeCHVFlagsLength int = 3

// Field # 387
const studentReceivedTANFCHVFlagsStartIndex int = 3349
const studentReceivedTANFCHVFlagsLength int = 3

// Field # 388
const studentReceivedWICCHVFlagsStartIndex int = 3352
const studentReceivedWICCHVFlagsLength int = 3

// Field # 389
const studentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex int = 3355
const studentFederalBenefitsNoneOfTheAboveCHVFlagsLength int = 3

// Field # 390
const studentFiled1040Or1040NRCHVFlagsStartIndex int = 3358
const studentFiled1040Or1040NRCHVFlagsLength int = 3

// Field # 391
const studentFiledNonUSTaxReturnCHVFlagsStartIndex int = 3361
const studentFiledNonUSTaxReturnCHVFlagsLength int = 3

// Field # 392
const studentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex int = 3364
const studentFiledJointReturnWithCurrentSpouseCHVFlagsLength int = 3

// Field # 393
const studentTaxReturnFilingStatusCHVFlagsStartIndex int = 3367
const studentTaxReturnFilingStatusCHVFlagsLength int = 3

// Field # 394
const studentIncomeEarnedFromWorkCorrectionCHVFlagsStartIndex int = 3370
const studentIncomeEarnedFromWorkCorrectionCHVFlagsLength int = 3

// Field # 395
const studentTaxExemptInterestIncomeCHVFlagsStartIndex int = 3373
const studentTaxExemptInterestIncomeCHVFlagsLength int = 3

// Field # 396
const studentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex int = 3376
const studentUntaxedPortionsOfIRADistributionsCHVFlagsLength int = 3

// Field # 397
const studentIRARolloverCHVFlagsStartIndex int = 3379
const studentIRARolloverCHVFlagsLength int = 3

// Field # 398
const studentUntaxedPortionsOfPensionsCHVFlagsStartIndex int = 3382
const studentUntaxedPortionsOfPensionsCHVFlagsLength int = 3

// Field # 399
const studentPensionRolloverCHVFlagsStartIndex int = 3385
const studentPensionRolloverCHVFlagsLength int = 3

// Field # 400
const studentAdjustedGrossIncomeCHVFlagsStartIndex int = 3388
const studentAdjustedGrossIncomeCHVFlagsLength int = 3

// Field # 401
const studentIncomeTaxPaidCHVFlagsStartIndex int = 3391
const studentIncomeTaxPaidCHVFlagsLength int = 3

// Field # 402
const studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex int = 3394
const studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsLength int = 3

// Field # 403
const studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex int = 3397
const studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength int = 3

// Field # 404
const studentEducationCreditsCHVFlagsStartIndex int = 3400
const studentEducationCreditsCHVFlagsLength int = 3

// Field # 405
const studentFiledScheduleABDEFHCHVFlagsStartIndex int = 3403
const studentFiledScheduleABDEFHCHVFlagsLength int = 3

// Field # 406
const studentScheduleCAmountCHVFlagsStartIndex int = 3406
const studentScheduleCAmountCHVFlagsLength int = 3

// Field # 407
const studentCollegeGrantAndScholarshipAidCHVFlagsStartIndex int = 3409
const studentCollegeGrantAndScholarshipAidCHVFlagsLength int = 3

// Field # 408
const studentForeignEarnedIncomeExclusionCHVFlagsStartIndex int = 3412
const studentForeignEarnedIncomeExclusionCHVFlagsLength int = 3

// Field # 409
const studentChildSupportReceivedCHVFlagsStartIndex int = 3415
const studentChildSupportReceivedCHVFlagsLength int = 3

// Field # 410
const studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex int = 3418
const studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsLength int = 3

// Field # 411
const studentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex int = 3421
const studentNetWorthOfCurrentInvestmentsCHVFlagsLength int = 3

// Field # 412
const studentTotalOfCashSavingsAndCheckingCHVFlagsStartIndex int = 3424
const studentTotalOfCashSavingsAndCheckingCHVFlagsLength int = 3

// Field # 413
const studentCollege1CHVFlagsStartIndex int = 3427
const studentCollege1CHVFlagsLength int = 3

// Field # 414
const studentCollege2CHVFlagsStartIndex int = 3430
const studentCollege2CHVFlagsLength int = 3

// Field # 415
const studentCollege3CHVFlagsStartIndex int = 3433
const studentCollege3CHVFlagsLength int = 3

// Field # 416
const studentCollege4CHVFlagsStartIndex int = 3436
const studentCollege4CHVFlagsLength int = 3

// Field # 417
const studentCollege5CHVFlagsStartIndex int = 3439
const studentCollege5CHVFlagsLength int = 3

// Field # 418
const studentCollege6CHVFlagsStartIndex int = 3442
const studentCollege6CHVFlagsLength int = 3

// Field # 419
const studentCollege7CHVFlagsStartIndex int = 3445
const studentCollege7CHVFlagsLength int = 3

// Field # 420
const studentCollege8CHVFlagsStartIndex int = 3448
const studentCollege8CHVFlagsLength int = 3

// Field # 421
const studentCollege9CHVFlagsStartIndex int = 3451
const studentCollege9CHVFlagsLength int = 3

// Field # 422
const studentCollege10CHVFlagsStartIndex int = 3454
const studentCollege10CHVFlagsLength int = 3

// Field # 423
const studentCollege11CHVFlagsStartIndex int = 3457
const studentCollege11CHVFlagsLength int = 3

// Field # 424
const studentCollege12CHVFlagsStartIndex int = 3460
const studentCollege12CHVFlagsLength int = 3

// Field # 425
const studentCollege13CHVFlagsStartIndex int = 3463
const studentCollege13CHVFlagsLength int = 3

// Field # 426
const studentCollege14CHVFlagsStartIndex int = 3466
const studentCollege14CHVFlagsLength int = 3

// Field # 427
const studentCollege15CHVFlagsStartIndex int = 3469
const studentCollege15CHVFlagsLength int = 3

// Field # 428
const studentCollege16CHVFlagsStartIndex int = 3472
const studentCollege16CHVFlagsLength int = 3

// Field # 429
const studentCollege17CHVFlagsStartIndex int = 3475
const studentCollege17CHVFlagsLength int = 3

// Field # 430
const studentCollege18CHVFlagsStartIndex int = 3478
const studentCollege18CHVFlagsLength int = 3

// Field # 431
const studentCollege19CHVFlagsStartIndex int = 3481
const studentCollege19CHVFlagsLength int = 3

// Field # 432
const studentCollege20CHVFlagsStartIndex int = 3484
const studentCollege20CHVFlagsLength int = 3

// Field # 433
const studentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex int = 3487
const studentConsentToRetrieveAndDiscloseFTICHVFlagsLength int = 3

// Field # 434
const studentSignatureCHVFlagsStartIndex int = 3490
const studentSignatureCHVFlagsLength int = 3

// Field # 435
const studentSignatureDateCHVFlagsStartIndex int = 3493
const studentSignatureDateCHVFlagsLength int = 3

// Field # 436
const studentSpouseFirstNameCHVFlagsStartIndex int = 3496
const studentSpouseFirstNameCHVFlagsLength int = 3

// Field # 437
const studentSpouseMiddleNameCHVFlagsStartIndex int = 3499
const studentSpouseMiddleNameCHVFlagsLength int = 3

// Field # 438
const studentSpouseLastNameCHVFlagsStartIndex int = 3502
const studentSpouseLastNameCHVFlagsLength int = 3

// Field # 439
const studentSpouseSuffixCHVFlagsStartIndex int = 3505
const studentSpouseSuffixCHVFlagsLength int = 3

// Field # 440
const studentSpouseDateOfBirthCHVFlagsStartIndex int = 3508
const studentSpouseDateOfBirthCHVFlagsLength int = 3

// Field # 441
const studentSpouseSSNCHVFlagsStartIndex int = 3511
const studentSpouseSSNCHVFlagsLength int = 3

// Field # 442
const studentSpouseITINCHVFlagsStartIndex int = 3514
const studentSpouseITINCHVFlagsLength int = 3

// Field # 443
const studentSpousePhoneNumberCHVFlagsStartIndex int = 3517
const studentSpousePhoneNumberCHVFlagsLength int = 3

// Field # 444
const studentSpouseEmailAddressCHVFlagsStartIndex int = 3520
const studentSpouseEmailAddressCHVFlagsLength int = 3

// Field # 445
const studentSpouseStreetAddressCHVFlagsStartIndex int = 3523
const studentSpouseStreetAddressCHVFlagsLength int = 3

// Field # 446
const studentSpouseCityCHVFlagsStartIndex int = 3526
const studentSpouseCityCHVFlagsLength int = 3

// Field # 447
const studentSpouseStateCHVFlagsStartIndex int = 3529
const studentSpouseStateCHVFlagsLength int = 3

// Field # 448
const studentSpouseZipCodeCHVFlagsStartIndex int = 3532
const studentSpouseZipCodeCHVFlagsLength int = 3

// Field # 449
const studentSpouseCountryCHVFlagsStartIndex int = 3535
const studentSpouseCountryCHVFlagsLength int = 3

// Field # 450
const studentSpouseFiled1040Or1040NRCHVFlagsStartIndex int = 3538
const studentSpouseFiled1040Or1040NRCHVFlagsLength int = 3

// Field # 451
const studentSpouseFiledNonUSTaxReturnCHVFlagsStartIndex int = 3541
const studentSpouseFiledNonUSTaxReturnCHVFlagsLength int = 3

// Field # 452
const studentSpouseTaxReturnFilingStatusCHVFlagsStartIndex int = 3544
const studentSpouseTaxReturnFilingStatusCHVFlagsLength int = 3

// Field # 453
const studentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex int = 3547
const studentSpouseIncomeEarnedFromWorkCHVFlagsLength int = 3

// Field # 454
const studentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex int = 3550
const studentSpouseTaxExemptInterestIncomeCHVFlagsLength int = 3

// Field # 455
const studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex int = 3553
const studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsLength int = 3

// Field # 456
const studentSpouseIRARolloverCHVFlagsStartIndex int = 3556
const studentSpouseIRARolloverCHVFlagsLength int = 3

// Field # 457
const studentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex int = 3559
const studentSpouseUntaxedPortionsOfPensionsCHVFlagsLength int = 3

// Field # 458
const studentSpousePensionRolloverCHVFlagsStartIndex int = 3562
const studentSpousePensionRolloverCHVFlagsLength int = 3

// Field # 459
const studentSpouseAdjustedGrossIncomeCHVFlagsStartIndex int = 3565
const studentSpouseAdjustedGrossIncomeCHVFlagsLength int = 3

// Field # 460
const studentSpouseIncomeTaxPaidCHVFlagsStartIndex int = 3568
const studentSpouseIncomeTaxPaidCHVFlagsLength int = 3

// Field # 461
const studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex int = 3571
const studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength int = 3

// Field # 462
const studentSpouseEducationCreditsCHVFlagsStartIndex int = 3574
const studentSpouseEducationCreditsCHVFlagsLength int = 3

// Field # 463
const studentSpouseFiledScheduleABDEFHCHVFlagsStartIndex int = 3577
const studentSpouseFiledScheduleABDEFHCHVFlagsLength int = 3

// Field # 464
const studentSpouseScheduleCAmountCHVFlagsStartIndex int = 3580
const studentSpouseScheduleCAmountCHVFlagsLength int = 3

// Field # 465
const studentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex int = 3583
const studentSpouseForeignEarnedIncomeExclusionCHVFlagsLength int = 3

// Field # 466
const studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex int = 3586
const studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsLength int = 3

// Field # 467
const studentSpouseSignatureCHVFlagsStartIndex int = 3589
const studentSpouseSignatureCHVFlagsLength int = 3

// Field # 468
const studentSpouseSignatureDateCHVFlagsStartIndex int = 3592
const studentSpouseSignatureDateCHVFlagsLength int = 3

// Field # 469
const parentFirstNameCHVFlagsStartIndex int = 3595
const parentFirstNameCHVFlagsLength int = 3

// Field # 470
const parentMiddleNameCHVFlagsStartIndex int = 3598
const parentMiddleNameCHVFlagsLength int = 3

// Field # 471
const parentLastNameCHVFlagsStartIndex int = 3601
const parentLastNameCHVFlagsLength int = 3

// Field # 472
const parentSuffixCHVFlagsStartIndex int = 3604
const parentSuffixCHVFlagsLength int = 3

// Field # 473
const parentDateOfBirthCHVFlagsStartIndex int = 3607
const parentDateOfBirthCHVFlagsLength int = 3

// Field # 474
const parentSSNCHVFlagsStartIndex int = 3610
const parentSSNCHVFlagsLength int = 3

// Field # 475
const parentITINCHVFlagsStartIndex int = 3613
const parentITINCHVFlagsLength int = 3

// Field # 476
const parentPhoneNumberCHVFlagsStartIndex int = 3616
const parentPhoneNumberCHVFlagsLength int = 3

// Field # 477
const parentEmailAddressCHVFlagsStartIndex int = 3619
const parentEmailAddressCHVFlagsLength int = 3

// Field # 478
const parentStreetAddressCHVFlagsStartIndex int = 3622
const parentStreetAddressCHVFlagsLength int = 3

// Field # 479
const parentCityCHVFlagsStartIndex int = 3625
const parentCityCHVFlagsLength int = 3

// Field # 480
const parentStateCHVFlagsStartIndex int = 3628
const parentStateCHVFlagsLength int = 3

// Field # 481
const parentZipCodeCHVFlagsStartIndex int = 3631
const parentZipCodeCHVFlagsLength int = 3

// Field # 482
const parentCountryCHVFlagsStartIndex int = 3634
const parentCountryCHVFlagsLength int = 3

// Field # 483
const parentMaritalStatusCHVFlagsStartIndex int = 3637
const parentMaritalStatusCHVFlagsLength int = 3

// Field # 484
const parentStateOfLegalResidenceCHVFlagsStartIndex int = 3640
const parentStateOfLegalResidenceCHVFlagsLength int = 3

// Field # 485
const parentLegalResidenceDateCHVFlagsStartIndex int = 3643
const parentLegalResidenceDateCHVFlagsLength int = 3

// Field # 486
const parentUpdatedFamilySizeCHVFlagsStartIndex int = 3646
const parentUpdatedFamilySizeCHVFlagsLength int = 3

// Field # 487
const parentNumberInCollegeCHVFlagsStartIndex int = 3649
const parentNumberInCollegeCHVFlagsLength int = 3

// Field # 488
const parentReceivedEITCCHVFlagsStartIndex int = 3652
const parentReceivedEITCCHVFlagsLength int = 3

// Field # 489
const parentReceivedFederalHousingAssistanceCHVFlagsStartIndex int = 3655
const parentReceivedFederalHousingAssistanceCHVFlagsLength int = 3

// Field # 490
const parentReceivedFreeReducedPriceLunchCHVFlagsStartIndex int = 3658
const parentReceivedFreeReducedPriceLunchCHVFlagsLength int = 3

// Field # 491
const parentReceivedMedicaidCHVFlagsStartIndex int = 3661
const parentReceivedMedicaidCHVFlagsLength int = 3

// Field # 492
const parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex int = 3664
const parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsLength int = 3

// Field # 493
const parentReceivedSNAPCHVFlagsStartIndex int = 3667
const parentReceivedSNAPCHVFlagsLength int = 3

// Field # 494
const parentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex int = 3670
const parentReceivedSupplementalSecurityIncomeCHVFlagsLength int = 3

// Field # 495
const parentReceivedTANFCHVFlagsStartIndex int = 3673
const parentReceivedTANFCHVFlagsLength int = 3

// Field # 496
const parentReceivedWICCHVFlagsStartIndex int = 3676
const parentReceivedWICCHVFlagsLength int = 3

// Field # 497
const parentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex int = 3679
const parentFederalBenefitsNoneOfTheAboveCHVFlagsLength int = 3

// Field # 498
const parentFiled1040Or1040NRCHVFlagsStartIndex int = 3682
const parentFiled1040Or1040NRCHVFlagsLength int = 3

// Field # 499
const parentFileNonUSTaxReturnCHVFlagsStartIndex int = 3685
const parentFileNonUSTaxReturnCHVFlagsLength int = 3

// Field # 500
const parentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex int = 3688
const parentFiledJointReturnWithCurrentSpouseCHVFlagsLength int = 3

// Field # 501
const parentTaxReturnFilingStatusCHVFlagsStartIndex int = 3691
const parentTaxReturnFilingStatusCHVFlagsLength int = 3

// Field # 502
const parentIncomeEarnedFromWorkCHVFlagsStartIndex int = 3694
const parentIncomeEarnedFromWorkCHVFlagsLength int = 3

// Field # 503
const parentTaxExemptInterestIncomeCHVFlagsStartIndex int = 3697
const parentTaxExemptInterestIncomeCHVFlagsLength int = 3

// Field # 504
const parentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex int = 3700
const parentUntaxedPortionsOfIRADistributionsCHVFlagsLength int = 3

// Field # 505
const parentIRARolloverCHVFlagsStartIndex int = 3703
const parentIRARolloverCHVFlagsLength int = 3

// Field # 506
const parentUntaxedPortionsOfPensionsCHVFlagsStartIndex int = 3706
const parentUntaxedPortionsOfPensionsCHVFlagsLength int = 3

// Field # 507
const parentPensionRolloverCHVFlagsStartIndex int = 3709
const parentPensionRolloverCHVFlagsLength int = 3

// Field # 508
const parentAdjustedGrossIncomeCHVFlagsStartIndex int = 3712
const parentAdjustedGrossIncomeCHVFlagsLength int = 3

// Field # 509
const parentIncomeTaxPaidCHVFlagsStartIndex int = 3715
const parentIncomeTaxPaidCHVFlagsLength int = 3

// Field # 510
const parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex int = 3718
const parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsLength int = 3

// Field # 511
const parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex int = 3721
const parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength int = 3

// Field # 512
const parentEducationCreditsCHVFlagsStartIndex int = 3724
const parentEducationCreditsCHVFlagsLength int = 3

// Field # 513
const parentFiledScheduleABDEFHCHVFlagsStartIndex int = 3727
const parentFiledScheduleABDEFHCHVFlagsLength int = 3

// Field # 514
const parentScheduleCAmountCHVFlagsStartIndex int = 3730
const parentScheduleCAmountCHVFlagsLength int = 3

// Field # 515
const parentCollegeGrantAndScholarshipAidCHVFlagsStartIndex int = 3733
const parentCollegeGrantAndScholarshipAidCHVFlagsLength int = 3

// Field # 516
const parentForeignEarnedIncomeExclusionCHVFlagsStartIndex int = 3736
const parentForeignEarnedIncomeExclusionCHVFlagsLength int = 3

// Field # 517
const parentChildSupportReceivedCHVFlagsStartIndex int = 3739
const parentChildSupportReceivedCHVFlagsLength int = 3

// Field # 518
const parentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex int = 3742
const parentNetWorthOfCurrentInvestmentsCHVFlagsLength int = 3

// Field # 519
const parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsStartIndex int = 3745
const parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsLength int = 3

// Field # 520
const parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex int = 3748
const parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsLength int = 3

// Field # 521
const parentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex int = 3751
const parentConsentToRetrieveAndDiscloseFTICHVFlagsLength int = 3

// Field # 522
const parentSignatureCHVFlagsStartIndex int = 3754
const parentSignatureCHVFlagsLength int = 3

// Field # 523
const parentSignatureDateCHVFlagsStartIndex int = 3757
const parentSignatureDateCHVFlagsLength int = 3

// Field # 524
const parentSpouseFirstNameCHVFlagsStartIndex int = 3760
const parentSpouseFirstNameCHVFlagsLength int = 3

// Field # 525
const parentSpouseMiddleNameCHVFlagsStartIndex int = 3763
const parentSpouseMiddleNameCHVFlagsLength int = 3

// Field # 526
const parentSpouseLastNameCHVFlagsStartIndex int = 3766
const parentSpouseLastNameCHVFlagsLength int = 3

// Field # 527
const parentSpouseSuffixCHVFlagsStartIndex int = 3769
const parentSpouseSuffixCHVFlagsLength int = 3

// Field # 528
const parentSpouseDateOfBirthCHVFlagsStartIndex int = 3772
const parentSpouseDateOfBirthCHVFlagsLength int = 3

// Field # 529
const parentSpouseSSNCHVFlagsStartIndex int = 3775
const parentSpouseSSNCHVFlagsLength int = 3

// Field # 530
const parentSpouseITINCHVFlagsStartIndex int = 3778
const parentSpouseITINCHVFlagsLength int = 3

// Field # 531
const parentSpousePhoneNumberCHVFlagsStartIndex int = 3781
const parentSpousePhoneNumberCHVFlagsLength int = 3

// Field # 532
const parentSpouseEmailAddressCHVFlagsStartIndex int = 3784
const parentSpouseEmailAddressCHVFlagsLength int = 3

// Field # 533
const parentSpouseStreetAddressCHVFlagsStartIndex int = 3787
const parentSpouseStreetAddressCHVFlagsLength int = 3

// Field # 534
const parentSpouseCityCHVFlagsStartIndex int = 3790
const parentSpouseCityCHVFlagsLength int = 3

// Field # 535
const parentSpouseStateCHVFlagsStartIndex int = 3793
const parentSpouseStateCHVFlagsLength int = 3

// Field # 536
const parentSpouseZipCodeCHVFlagsStartIndex int = 3796
const parentSpouseZipCodeCHVFlagsLength int = 3

// Field # 537
const parentSpouseCountryCHVFlagsStartIndex int = 3799
const parentSpouseCountryCHVFlagsLength int = 3

// Field # 538
const parentSpouseFiled1040Or1040NRCHVFlagsStartIndex int = 3802
const parentSpouseFiled1040Or1040NRCHVFlagsLength int = 3

// Field # 539
const parentSpouseFileNonUSTaxReturnCHVFlagsStartIndex int = 3805
const parentSpouseFileNonUSTaxReturnCHVFlagsLength int = 3

// Field # 540
const parentSpouseTaxReturnFilingStatusCHVFlagsStartIndex int = 3808
const parentSpouseTaxReturnFilingStatusCHVFlagsLength int = 3

// Field # 541
const parentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex int = 3811
const parentSpouseIncomeEarnedFromWorkCHVFlagsLength int = 3

// Field # 542
const parentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex int = 3814
const parentSpouseTaxExemptInterestIncomeCHVFlagsLength int = 3

// Field # 543
const parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex int = 3817
const parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsLength int = 3

// Field # 544
const parentSpouseIRARolloverCHVFlagsStartIndex int = 3820
const parentSpouseIRARolloverCHVFlagsLength int = 3

// Field # 545
const parentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex int = 3823
const parentSpouseUntaxedPortionsOfPensionsCHVFlagsLength int = 3

// Field # 546
const parentSpousePensionRolloverCHVFlagsStartIndex int = 3826
const parentSpousePensionRolloverCHVFlagsLength int = 3

// Field # 547
const parentSpouseAdjustedGrossIncomeCHVFlagsStartIndex int = 3829
const parentSpouseAdjustedGrossIncomeCHVFlagsLength int = 3

// Field # 548
const parentSpouseIncomeTaxPaidCHVFlagsStartIndex int = 3832
const parentSpouseIncomeTaxPaidCHVFlagsLength int = 3

// Field # 549
const parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex int = 3835
const parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength int = 3

// Field # 550
const parentSpouseEducationCreditsCHVFlagsStartIndex int = 3838
const parentSpouseEducationCreditsCHVFlagsLength int = 3

// Field # 551
const parentSpouseFiledScheduleABDEFHCHVFlagsStartIndex int = 3841
const parentSpouseFiledScheduleABDEFHCHVFlagsLength int = 3

// Field # 552
const parentSpouseScheduleCAmountCHVFlagsStartIndex int = 3844
const parentSpouseScheduleCAmountCHVFlagsLength int = 3

// Field # 553
const parentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex int = 3847
const parentSpouseForeignEarnedIncomeExclusionCHVFlagsLength int = 3

// Field # 554
const parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex int = 3850
const parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsLength int = 3

// Field # 555
const parentSpouseSignatureCHVFlagsStartIndex int = 3853
const parentSpouseSignatureCHVFlagsLength int = 3

// Field # 556
const parentSpouseSignatureDateCHVFlagsStartIndex int = 3856
const parentSpouseSignatureDateCHVFlagsLength int = 3

// Field # 557
const dHSPrimaryMatchStatusStartIndex int = 3859
const dHSPrimaryMatchStatusLength int = 1

// Field # 559
const dHSCaseNumberStartIndex int = 3861
const dHSCaseNumberLength int = 15

// Field # 560
const nsldsMatchStatusStartIndex int = 3876
const nsldsMatchStatusLength int = 1

// Field # 561
const nsldsPostscreeningReasonCodeStartIndex int = 3877
const nsldsPostscreeningReasonCodeLength int = 6

// Field # 562
const studentSSACitizenshipFlagResultsStartIndex int = 3883
const studentSSACitizenshipFlagResultsLength int = 1

// Field # 563
const studentSSAMatchStatusStartIndex int = 3884
const studentSSAMatchStatusLength int = 1

// Field # 564
const studentSpouseSSAMatchStatusStartIndex int = 3885
const studentSpouseSSAMatchStatusLength int = 1

// Field # 565
const parentSSAMatchStatusStartIndex int = 3886
const parentSSAMatchStatusLength int = 1

// Field # 566
const parentSpouseOrPartnerSSAMatchStatusStartIndex int = 3887
const parentSpouseOrPartnerSSAMatchStatusLength int = 1

// Field # 567
const vAMatchFlagStartIndex int = 3888
const vAMatchFlagLength int = 1

// Field # 568
const commentCodesStartIndex int = 3889
const commentCodesLength int = 60

// Field # 569
const drugAbuseHoldIndicatorStartIndex int = 3949
const drugAbuseHoldIndicatorLength int = 1

// Field # 570
const graduateFlagStartIndex int = 3950
const graduateFlagLength int = 1

// Field # 571
const pellGrantEligibilityFlagStartIndex int = 3951
const pellGrantEligibilityFlagLength int = 1

// Field # 572
const reprocessedReasonCodeStartIndex int = 3952
const reprocessedReasonCodeLength int = 2

// Field # 573
const fpsCFlagStartIndex int = 3954
const fpsCFlagLength int = 1

// Field # 574
const fpsCChangeFlagStartIndex int = 3955
const fpsCChangeFlagLength int = 1

// Field # 575
const electronicFederalSchoolCodeIndicatorStartIndex int = 3956
const electronicFederalSchoolCodeIndicatorLength int = 2

// Field # 576
const rejectReasonCodesStartIndex int = 3958
const rejectReasonCodesLength int = 110

// Field # 577
const electronicTransactionIndicatorFlagStartIndex int = 4068
const electronicTransactionIndicatorFlagLength int = 1

// Field # 578
const studentLastNameSSNChangeFlagStartIndex int = 4069
const studentLastNameSSNChangeFlagLength int = 1

// Field # 579
const highSchoolCodeStartIndex int = 4070
const highSchoolCodeLength int = 12

// Field # 580
const verificationSelectionChangeFlagStartIndex int = 4082
const verificationSelectionChangeFlagLength int = 1

// Field # 581
const useUserProvidedDataOnlyStartIndex int = 4083
const useUserProvidedDataOnlyLength int = 5

// Field # 583
const nsldsPellOverpaymentFlagStartIndex int = 4449
const nsldsPellOverpaymentFlagLength int = 1

// Field # 584
const nsldsPellOverpaymentContactStartIndex int = 4450
const nsldsPellOverpaymentContactLength int = 8

// Field # 585
const nsldsFSEOGOverpaymentFlagStartIndex int = 4458
const nsldsFSEOGOverpaymentFlagLength int = 1

// Field # 586
const nsldsFSEOGOverpaymentContactStartIndex int = 4459
const nsldsFSEOGOverpaymentContactLength int = 8

// Field # 587
const nsldsPerkinsOverpaymentFlagStartIndex int = 4467
const nsldsPerkinsOverpaymentFlagLength int = 1

// Field # 588
const nsldsPerkinsOverpaymentContactStartIndex int = 4468
const nsldsPerkinsOverpaymentContactLength int = 8

// Field # 589
const nsldsTEACHGrantOverpaymentFlagStartIndex int = 4476
const nsldsTEACHGrantOverpaymentFlagLength int = 1

// Field # 590
const nsldsTEACHGrantOverpaymentContactStartIndex int = 4477
const nsldsTEACHGrantOverpaymentContactLength int = 8

// Field # 591
const nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagStartIndex int = 4485
const nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagLength int = 1

// Field # 592
const nsldsIraqAndAfghanistanServiceGrantOverpaymentContactStartIndex int = 4486
const nsldsIraqAndAfghanistanServiceGrantOverpaymentContactLength int = 8

// Field # 593
const nsldsDefaultedLoanFlagStartIndex int = 4494
const nsldsDefaultedLoanFlagLength int = 1

// Field # 594
const nsldsDischargedLoanFlagStartIndex int = 4495
const nsldsDischargedLoanFlagLength int = 1

// Field # 595
const nsldsFraudLoanFlagStartIndex int = 4496
const nsldsFraudLoanFlagLength int = 1

// Field # 596
const nsldsSatisfactoryArrangementsFlagStartIndex int = 4497
const nsldsSatisfactoryArrangementsFlagLength int = 1

// Field # 597
const nsldsActiveBankruptcyFlagStartIndex int = 4498
const nsldsActiveBankruptcyFlagLength int = 1

// Field # 598
const nsldsTEACHGrantConvertedToLoanFlagStartIndex int = 4499
const nsldsTEACHGrantConvertedToLoanFlagLength int = 1

// Field # 599
const nsldsAggregateSubsidizedOutstandingPrincipalBalanceStartIndex int = 4500
const nsldsAggregateSubsidizedOutstandingPrincipalBalanceLength int = 6

// Field # 600
const nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceStartIndex int = 4506
const nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceLength int = 6

// Field # 601
const nsldsAggregateCombinedOutstandingPrincipalBalanceStartIndex int = 4512
const nsldsAggregateCombinedOutstandingPrincipalBalanceLength int = 6

// Field # 602
const nsldsAggregateUnallocatedConsolidatedOutstandingPrincipalBalanceStartIndex int = 4518
const nsldsAggregateUnallocatedConsolidatedOutstandingPrincipalBalanceLength int = 6

// Field # 603
const nsldsAggregateTEACHLoanPrincipalBalanceStartIndex int = 4524
const nsldsAggregateTEACHLoanPrincipalBalanceLength int = 6

// Field # 604
const nsldsAggregateSubsidizedPendingDisbursementStartIndex int = 4530
const nsldsAggregateSubsidizedPendingDisbursementLength int = 6

// Field # 605
const nsldsAggregateUnsubsidizedPendingDisbursementStartIndex int = 4536
const nsldsAggregateUnsubsidizedPendingDisbursementLength int = 6

// Field # 606
const nsldsAggregateCombinedPendingDisbursementStartIndex int = 4542
const nsldsAggregateCombinedPendingDisbursementLength int = 6

// Field # 607
const nsldsAggregateSubsidizedTotalStartIndex int = 4548
const nsldsAggregateSubsidizedTotalLength int = 6

// Field # 608
const nsldsAggregateUnsubsidizedTotalStartIndex int = 4554
const nsldsAggregateUnsubsidizedTotalLength int = 6

// Field # 609
const nsldsAggregateCombinedTotalStartIndex int = 4560
const nsldsAggregateCombinedTotalLength int = 6

// Field # 610
const nsldsUnallocatedConsolidatedTotalStartIndex int = 4566
const nsldsUnallocatedConsolidatedTotalLength int = 6

// Field # 611
const nsldsTEACHLoanTotalStartIndex int = 4572
const nsldsTEACHLoanTotalLength int = 6

// Field # 612
const nsldsPerkinsTotalDisbursementsStartIndex int = 4578
const nsldsPerkinsTotalDisbursementsLength int = 6

// Field # 613
const nsldsPerkinsCurrentYearDisbursementAmountStartIndex int = 4584
const nsldsPerkinsCurrentYearDisbursementAmountLength int = 6

// Field # 614
const nsldsAggregateTEACHGrantUndergraduateDisbursedTotalStartIndex int = 4590
const nsldsAggregateTEACHGrantUndergraduateDisbursedTotalLength int = 6

// Field # 615
const nsldsAggregateTEACHGraduateDisbursementAmountStartIndex int = 4596
const nsldsAggregateTEACHGraduateDisbursementAmountLength int = 6

// Field # 616
const nsldsDefaultedLoanChangeFlagStartIndex int = 4602
const nsldsDefaultedLoanChangeFlagLength int = 1

// Field # 617
const nsldsFraudLoanChangeFlagStartIndex int = 4603
const nsldsFraudLoanChangeFlagLength int = 1

// Field # 618
const nsldsDischargedLoanChangeFlagStartIndex int = 4604
const nsldsDischargedLoanChangeFlagLength int = 1

// Field # 619
const nsldsLoanSatisfactoryRepaymentChangeFlagStartIndex int = 4605
const nsldsLoanSatisfactoryRepaymentChangeFlagLength int = 1

// Field # 620
const nsldsActiveBankruptcyChangeFlagStartIndex int = 4606
const nsldsActiveBankruptcyChangeFlagLength int = 1

// Field # 621
const nsldsTEACHGrantToLoanConversionChangeFlagStartIndex int = 4607
const nsldsTEACHGrantToLoanConversionChangeFlagLength int = 1

// Field # 622
const nsldsOverpaymentsChangeFlagStartIndex int = 4608
const nsldsOverpaymentsChangeFlagLength int = 1

// Field # 623
const nsldsAggregateLoanChangeFlagStartIndex int = 4609
const nsldsAggregateLoanChangeFlagLength int = 1

// Field # 624
const nsldsPerkinsLoanChangeFlagStartIndex int = 4610
const nsldsPerkinsLoanChangeFlagLength int = 1

// Field # 625
const nsldsPellPaymentChangeFlagStartIndex int = 4611
const nsldsPellPaymentChangeFlagLength int = 1

// Field # 626
const nsldsTEACHGrantChangeFlagStartIndex int = 4612
const nsldsTEACHGrantChangeFlagLength int = 1

// Field # 627
const nsldsAdditionalPellFlagStartIndex int = 4613
const nsldsAdditionalPellFlagLength int = 1

// Field # 628
const nsldsAdditionalLoansFlagStartIndex int = 4614
const nsldsAdditionalLoansFlagLength int = 1

// Field # 629
const nsldsAdditionalTEACHGrantFlagStartIndex int = 4615
const nsldsAdditionalTEACHGrantFlagLength int = 1

// Field # 630
const nsldsDirectLoanMasterPromNoteFlagStartIndex int = 4616
const nsldsDirectLoanMasterPromNoteFlagLength int = 1

// Field # 631
const nsldsDirectLoanPLUSMasterPromNoteFlagStartIndex int = 4617
const nsldsDirectLoanPLUSMasterPromNoteFlagLength int = 1

// Field # 632
const nsldsDirectLoanGraduatePLUSMasterPromNoteFlagStartIndex int = 4618
const nsldsDirectLoanGraduatePLUSMasterPromNoteFlagLength int = 1

// Field # 633
const nsldsUndergraduateSubsidizedLoanLimitFlagStartIndex int = 4619
const nsldsUndergraduateSubsidizedLoanLimitFlagLength int = 1

// Field # 634
const nsldsUndergraduateCombinedLoanLimitFlagStartIndex int = 4620
const nsldsUndergraduateCombinedLoanLimitFlagLength int = 1

// Field # 635
const nsldsGraduateSubsidizedLoanLimitFlagStartIndex int = 4621
const nsldsGraduateSubsidizedLoanLimitFlagLength int = 1

// Field # 636
const nsldsGraduateCombinedLoanLimitFlagStartIndex int = 4622
const nsldsGraduateCombinedLoanLimitFlagLength int = 1

// Field # 637
const nsldsLEULimitIndicatorStartIndex int = 4623
const nsldsLEULimitIndicatorLength int = 1

// Field # 638
const nsldsPellLifetimeEligibilityUsedStartIndex int = 4624
const nsldsPellLifetimeEligibilityUsedLength int = 7

// Field # 639
const nsldsSULAFlagStartIndex int = 4631
const nsldsSULAFlagLength int = 1

// Field # 640
const nsldsSubsidizedLimitEligibilityFlagStartIndex int = 4632
const nsldsSubsidizedLimitEligibilityFlagLength int = 6

// Field # 641
const nsldsUnusualEnrollmentHistoryFlagStartIndex int = 4638
const nsldsUnusualEnrollmentHistoryFlagLength int = 1

// Field # 643
const nsldsPellSequenceNumber1StartIndex int = 4659
const nsldsPellSequenceNumber1Length int = 2

// Field # 644
const nsldsPellVerificationFlag1StartIndex int = 4661
const nsldsPellVerificationFlag1Length int = 3

// Field # 645
const nsldsSAI1StartIndex int = 4664
const nsldsSAI1Length int = 6

// Field # 646
const nsldsPellSchoolCode1StartIndex int = 4670
const nsldsPellSchoolCode1Length int = 8

// Field # 647
const nsldsPellTransactionNumber1StartIndex int = 4678
const nsldsPellTransactionNumber1Length int = 2

// Field # 648
const nsldsPellLastDisbursementDate1StartIndex int = 4680
const nsldsPellLastDisbursementDate1Length int = 8

// Field # 649
const nsldsPellScheduledAmount1StartIndex int = 4688
const nsldsPellScheduledAmount1Length int = 6

// Field # 650
const nsldsPellAmountPaidToDate1StartIndex int = 4694
const nsldsPellAmountPaidToDate1Length int = 6

// Field # 651
const nsldsPellPercentEligibilityUsedDecimal1StartIndex int = 4700
const nsldsPellPercentEligibilityUsedDecimal1Length int = 7

// Field # 652
const nsldsPellAwardAmount1StartIndex int = 4707
const nsldsPellAwardAmount1Length int = 6

// Field # 653
const nsldsAdditionalEligibilityIndicator1StartIndex int = 4713
const nsldsAdditionalEligibilityIndicator1Length int = 1

// Field # 655
const nsldsPellSequenceNumber2StartIndex int = 4734
const nsldsPellSequenceNumber2Length int = 2

// Field # 656
const nsldsPellVerificationFlag2StartIndex int = 4736
const nsldsPellVerificationFlag2Length int = 3

// Field # 657
const nsldsSAI2StartIndex int = 4739
const nsldsSAI2Length int = 6

// Field # 658
const nsldsPellSchoolCode2StartIndex int = 4745
const nsldsPellSchoolCode2Length int = 8

// Field # 659
const nsldsPellTransactionNumber2StartIndex int = 4753
const nsldsPellTransactionNumber2Length int = 2

// Field # 660
const nsldsPellLastDisbursementDate2StartIndex int = 4755
const nsldsPellLastDisbursementDate2Length int = 8

// Field # 661
const nsldsPellScheduledAmount2StartIndex int = 4763
const nsldsPellScheduledAmount2Length int = 6

// Field # 662
const nsldsPellAmountPaidToDate2StartIndex int = 4769
const nsldsPellAmountPaidToDate2Length int = 6

// Field # 663
const nsldsPellPercentEligibilityUsedDecimal2StartIndex int = 4775
const nsldsPellPercentEligibilityUsedDecimal2Length int = 7

// Field # 664
const nsldsPellAwardAmount2StartIndex int = 4782
const nsldsPellAwardAmount2Length int = 6

// Field # 665
const nsldsAdditionalEligibilityIndicator2StartIndex int = 4788
const nsldsAdditionalEligibilityIndicator2Length int = 1

// Field # 667
const nsldsPellSequenceNumber3StartIndex int = 4809
const nsldsPellSequenceNumber3Length int = 2

// Field # 668
const nsldsPellVerificationFlag3StartIndex int = 4811
const nsldsPellVerificationFlag3Length int = 3

// Field # 669
const nsldsSAI3StartIndex int = 4814
const nsldsSAI3Length int = 6

// Field # 670
const nsldsPellSchoolCode3StartIndex int = 4820
const nsldsPellSchoolCode3Length int = 8

// Field # 671
const nsldsPellTransactionNumber3StartIndex int = 4828
const nsldsPellTransactionNumber3Length int = 2

// Field # 672
const nsldsPellLastDisbursementDate3StartIndex int = 4830
const nsldsPellLastDisbursementDate3Length int = 8

// Field # 673
const nsldsPellScheduledAmount3StartIndex int = 4838
const nsldsPellScheduledAmount3Length int = 6

// Field # 674
const nsldsPellAmountPaidToDate3StartIndex int = 4844
const nsldsPellAmountPaidToDate3Length int = 6

// Field # 675
const nsldsPellPercentEligibilityUsedDecimal3StartIndex int = 4850
const nsldsPellPercentEligibilityUsedDecimal3Length int = 7

// Field # 676
const nsldsPellAwardAmount3StartIndex int = 4857
const nsldsPellAwardAmount3Length int = 6

// Field # 677
const nsldsAdditionalEligibilityIndicator3StartIndex int = 4863
const nsldsAdditionalEligibilityIndicator3Length int = 1

// Field # 679
const nsldsTEACHGrantSequence1StartIndex int = 4884
const nsldsTEACHGrantSequence1Length int = 2

// Field # 680
const nsldsTEACHGrantSchoolCode1StartIndex int = 4886
const nsldsTEACHGrantSchoolCode1Length int = 8

// Field # 681
const nsldsTEACHGrantTransactionNumber1StartIndex int = 4894
const nsldsTEACHGrantTransactionNumber1Length int = 2

// Field # 682
const nsldsTEACHGrantLastDisbursementDate1StartIndex int = 4896
const nsldsTEACHGrantLastDisbursementDate1Length int = 8

// Field # 683
const nsldsTEACHGrantScheduledAmount1StartIndex int = 4904
const nsldsTEACHGrantScheduledAmount1Length int = 6

// Field # 684
const nsldsTEACHGrantAmountPaidToDate1StartIndex int = 4910
const nsldsTEACHGrantAmountPaidToDate1Length int = 6

// Field # 685
const nsldsTEACHGrantAwardAmount1StartIndex int = 4916
const nsldsTEACHGrantAwardAmount1Length int = 6

// Field # 686
const nsldsTEACHGrantAcademicYearLevel1StartIndex int = 4922
const nsldsTEACHGrantAcademicYearLevel1Length int = 1

// Field # 687
const nsldsTEACHGrantAwardYear1StartIndex int = 4923
const nsldsTEACHGrantAwardYear1Length int = 4

// Field # 688
const nsldsTEACHGrantLoanConversionFlag1StartIndex int = 4927
const nsldsTEACHGrantLoanConversionFlag1Length int = 1

// Field # 689
const nsldsTEACHGrantDischargeCode1StartIndex int = 4928
const nsldsTEACHGrantDischargeCode1Length int = 4

// Field # 690
const nsldsTEACHGrantDischargeAmount1StartIndex int = 4932
const nsldsTEACHGrantDischargeAmount1Length int = 6

// Field # 691
const nsldsTEACHGrantAdjustedDisbursement1StartIndex int = 4938
const nsldsTEACHGrantAdjustedDisbursement1Length int = 6

// Field # 693
const nsldsTEACHGrantSequence2StartIndex int = 4964
const nsldsTEACHGrantSequence2Length int = 2

// Field # 694
const nsldsTEACHGrantSchoolCode2StartIndex int = 4966
const nsldsTEACHGrantSchoolCode2Length int = 8

// Field # 695
const nsldsTEACHGrantTransactionNumber2StartIndex int = 4974
const nsldsTEACHGrantTransactionNumber2Length int = 2

// Field # 696
const nsldsTEACHGrantLastDisbursementDate2StartIndex int = 4976
const nsldsTEACHGrantLastDisbursementDate2Length int = 8

// Field # 697
const nsldsTEACHGrantScheduledAmount2StartIndex int = 4984
const nsldsTEACHGrantScheduledAmount2Length int = 6

// Field # 698
const nsldsTEACHGrantAmountPaidToDate2StartIndex int = 4990
const nsldsTEACHGrantAmountPaidToDate2Length int = 6

// Field # 699
const nsldsTEACHGrantAwardAmount2StartIndex int = 4996
const nsldsTEACHGrantAwardAmount2Length int = 6

// Field # 700
const nsldsTEACHGrantAcademicYearLevel2StartIndex int = 5002
const nsldsTEACHGrantAcademicYearLevel2Length int = 1

// Field # 701
const nsldsTEACHGrantAwardYear2StartIndex int = 5003
const nsldsTEACHGrantAwardYear2Length int = 4

// Field # 702
const nsldsTEACHGrantLoanConversionFlag2StartIndex int = 5007
const nsldsTEACHGrantLoanConversionFlag2Length int = 1

// Field # 703
const nsldsTEACHGrantDischargeCode2StartIndex int = 5008
const nsldsTEACHGrantDischargeCode2Length int = 4

// Field # 704
const nsldsTEACHGrantDischargeAmount2StartIndex int = 5012
const nsldsTEACHGrantDischargeAmount2Length int = 6

// Field # 705
const nsldsTEACHGrantAdjustedDisbursement2StartIndex int = 5018
const nsldsTEACHGrantAdjustedDisbursement2Length int = 6

// Field # 707
const nsldsTEACHGrantSequence3StartIndex int = 5044
const nsldsTEACHGrantSequence3Length int = 2

// Field # 708
const nsldsTEACHGrantSchoolCode3StartIndex int = 5046
const nsldsTEACHGrantSchoolCode3Length int = 8

// Field # 709
const nsldsTEACHGrantTransactionNumber3StartIndex int = 5054
const nsldsTEACHGrantTransactionNumber3Length int = 2

// Field # 710
const nsldsTEACHGrantLastDisbursementDate3StartIndex int = 5056
const nsldsTEACHGrantLastDisbursementDate3Length int = 8

// Field # 711
const nsldsTEACHGrantScheduledAmount3StartIndex int = 5064
const nsldsTEACHGrantScheduledAmount3Length int = 6

// Field # 712
const nsldsTEACHGrantAmountPaidToDate3StartIndex int = 5070
const nsldsTEACHGrantAmountPaidToDate3Length int = 6

// Field # 713
const nsldsTEACHGrantAwardAmount3StartIndex int = 5076
const nsldsTEACHGrantAwardAmount3Length int = 6

// Field # 714
const nsldsTEACHGrantAcademicYearLevel3StartIndex int = 5082
const nsldsTEACHGrantAcademicYearLevel3Length int = 1

// Field # 715
const nsldsTEACHGrantAwardYear3StartIndex int = 5083
const nsldsTEACHGrantAwardYear3Length int = 4

// Field # 716
const nsldsTEACHGrantLoanConversionFlag3StartIndex int = 5087
const nsldsTEACHGrantLoanConversionFlag3Length int = 1

// Field # 717
const nsldsTEACHGrantDischargeCode3StartIndex int = 5088
const nsldsTEACHGrantDischargeCode3Length int = 4

// Field # 718
const nsldsTEACHGrantDischargeAmount3StartIndex int = 5092
const nsldsTEACHGrantDischargeAmount3Length int = 6

// Field # 719
const nsldsTEACHGrantAdjustedDisbursement3StartIndex int = 5098
const nsldsTEACHGrantAdjustedDisbursement3Length int = 6

// Field # 721
const nsldsLoanSequenceNumber1StartIndex int = 5124
const nsldsLoanSequenceNumber1Length int = 2

// Field # 722
const nsldsLoanDefaultedRecentIndicator1StartIndex int = 5126
const nsldsLoanDefaultedRecentIndicator1Length int = 1

// Field # 723
const nsldsLoanChangeFlag1StartIndex int = 5127
const nsldsLoanChangeFlag1Length int = 1

// Field # 724
const nsldsLoanTypeCode1StartIndex int = 5128
const nsldsLoanTypeCode1Length int = 2

// Field # 725
const nsldsLoanNetAmount1StartIndex int = 5130
const nsldsLoanNetAmount1Length int = 6

// Field # 726
const nsldsLoanCurrentStatusCode1StartIndex int = 5136
const nsldsLoanCurrentStatusCode1Length int = 2

// Field # 727
const nsldsLoanCurrentStatusDate1StartIndex int = 5138
const nsldsLoanCurrentStatusDate1Length int = 8

// Field # 728
const nsldsLoanOutstandingPrincipalBalance1StartIndex int = 5146
const nsldsLoanOutstandingPrincipalBalance1Length int = 6

// Field # 729
const nsldsLoanOutstandingPrincipalBalanceDate1StartIndex int = 5152
const nsldsLoanOutstandingPrincipalBalanceDate1Length int = 8

// Field # 730
const nsldsLoanPeriodBeginDate1StartIndex int = 5160
const nsldsLoanPeriodBeginDate1Length int = 8

// Field # 731
const nsldsLoanPeriodEndDate1StartIndex int = 5168
const nsldsLoanPeriodEndDate1Length int = 8

// Field # 732
const nsldsLoanGuarantyAgencyCode1StartIndex int = 5176
const nsldsLoanGuarantyAgencyCode1Length int = 3

// Field # 733
const nsldsLoanContactType1StartIndex int = 5179
const nsldsLoanContactType1Length int = 3

// Field # 734
const nsldsLoanSchoolCode1StartIndex int = 5182
const nsldsLoanSchoolCode1Length int = 8

// Field # 735
const nsldsLoanContactCode1StartIndex int = 5190
const nsldsLoanContactCode1Length int = 8

// Field # 736
const nsldsLoanGradeLevel1StartIndex int = 5198
const nsldsLoanGradeLevel1Length int = 3

// Field # 737
const nsldsLoanAdditionalUnsubsidizedFlag1StartIndex int = 5201
const nsldsLoanAdditionalUnsubsidizedFlag1Length int = 1

// Field # 738
const nsldsLoanCapitalizedInterestFlag1StartIndex int = 5202
const nsldsLoanCapitalizedInterestFlag1Length int = 1

// Field # 739
const nsldsLoanDisbursementAmount1StartIndex int = 5203
const nsldsLoanDisbursementAmount1Length int = 6

// Field # 740
const nsldsLoanDisbursementDate1StartIndex int = 5209
const nsldsLoanDisbursementDate1Length int = 8

// Field # 741
const nsldsLoanConfirmedLoanSubsidyStatus1StartIndex int = 5217
const nsldsLoanConfirmedLoanSubsidyStatus1Length int = 1

// Field # 742
const nsldsLoanSubsidyStatusDate1StartIndex int = 5218
const nsldsLoanSubsidyStatusDate1Length int = 8

// Field # 744
const nsldsLoanSequenceNumber2StartIndex int = 5246
const nsldsLoanSequenceNumber2Length int = 2

// Field # 745
const nsldsLoanDefaultedRecentIndicator2StartIndex int = 5248
const nsldsLoanDefaultedRecentIndicator2Length int = 1

// Field # 746
const nsldsLoanChangeFlag2StartIndex int = 5249
const nsldsLoanChangeFlag2Length int = 1

// Field # 747
const nsldsLoanTypeCode2StartIndex int = 5250
const nsldsLoanTypeCode2Length int = 2

// Field # 748
const nsldsLoanNetAmount2StartIndex int = 5252
const nsldsLoanNetAmount2Length int = 6

// Field # 749
const nsldsLoanCurrentStatusCode2StartIndex int = 5258
const nsldsLoanCurrentStatusCode2Length int = 2

// Field # 750
const nsldsLoanCurrentStatusDate2StartIndex int = 5260
const nsldsLoanCurrentStatusDate2Length int = 8

// Field # 751
const nsldsLoanOutstandingPrincipalBalance2StartIndex int = 5268
const nsldsLoanOutstandingPrincipalBalance2Length int = 6

// Field # 752
const nsldsLoanOutstandingPrincipalBalanceDate2StartIndex int = 5274
const nsldsLoanOutstandingPrincipalBalanceDate2Length int = 8

// Field # 753
const nsldsLoanPeriodBeginDate2StartIndex int = 5282
const nsldsLoanPeriodBeginDate2Length int = 8

// Field # 754
const nsldsLoanPeriodEndDate2StartIndex int = 5290
const nsldsLoanPeriodEndDate2Length int = 8

// Field # 755
const nsldsLoanGuarantyAgencyCode2StartIndex int = 5298
const nsldsLoanGuarantyAgencyCode2Length int = 3

// Field # 756
const nsldsLoanContactType2StartIndex int = 5301
const nsldsLoanContactType2Length int = 3

// Field # 757
const nsldsLoanSchoolCode2StartIndex int = 5304
const nsldsLoanSchoolCode2Length int = 8

// Field # 758
const nsldsLoanContactCode2StartIndex int = 5312
const nsldsLoanContactCode2Length int = 8

// Field # 759
const nsldsLoanGradeLevel2StartIndex int = 5320
const nsldsLoanGradeLevel2Length int = 3

// Field # 760
const nsldsLoanAdditionalUnsubsidizedFlag2StartIndex int = 5323
const nsldsLoanAdditionalUnsubsidizedFlag2Length int = 1

// Field # 761
const nsldsLoanCapitalizedInterestFlag2StartIndex int = 5324
const nsldsLoanCapitalizedInterestFlag2Length int = 1

// Field # 762
const nsldsLoanDisbursementAmount2StartIndex int = 5325
const nsldsLoanDisbursementAmount2Length int = 6

// Field # 763
const nsldsLoanDisbursementDate2StartIndex int = 5331
const nsldsLoanDisbursementDate2Length int = 8

// Field # 764
const nsldsLoanConfirmedLoanSubsidyStatus2StartIndex int = 5339
const nsldsLoanConfirmedLoanSubsidyStatus2Length int = 1

// Field # 765
const nsldsLoanSubsidyStatusDate2StartIndex int = 5340
const nsldsLoanSubsidyStatusDate2Length int = 8

// Field # 767
const nsldsLoanSequenceNumber3StartIndex int = 5368
const nsldsLoanSequenceNumber3Length int = 2

// Field # 768
const nsldsLoanDefaultedRecentIndicator3StartIndex int = 5370
const nsldsLoanDefaultedRecentIndicator3Length int = 1

// Field # 769
const nsldsLoanChangeFlag3StartIndex int = 5371
const nsldsLoanChangeFlag3Length int = 1

// Field # 770
const nsldsLoanTypeCode3StartIndex int = 5372
const nsldsLoanTypeCode3Length int = 2

// Field # 771
const nsldsLoanNetAmount3StartIndex int = 5374
const nsldsLoanNetAmount3Length int = 6

// Field # 772
const nsldsLoanCurrentStatusCode3StartIndex int = 5380
const nsldsLoanCurrentStatusCode3Length int = 2

// Field # 773
const nsldsLoanCurrentStatusDate3StartIndex int = 5382
const nsldsLoanCurrentStatusDate3Length int = 8

// Field # 774
const nsldsLoanOutstandingPrincipalBalance3StartIndex int = 5390
const nsldsLoanOutstandingPrincipalBalance3Length int = 6

// Field # 775
const nsldsLoanOutstandingPrincipalBalanceDate3StartIndex int = 5396
const nsldsLoanOutstandingPrincipalBalanceDate3Length int = 8

// Field # 776
const nsldsLoanPeriodBeginDate3StartIndex int = 5404
const nsldsLoanPeriodBeginDate3Length int = 8

// Field # 777
const nsldsLoanPeriodEndDate3StartIndex int = 5412
const nsldsLoanPeriodEndDate3Length int = 8

// Field # 778
const nsldsLoanGuarantyAgencyCode3StartIndex int = 5420
const nsldsLoanGuarantyAgencyCode3Length int = 3

// Field # 779
const nsldsLoanContactType3StartIndex int = 5423
const nsldsLoanContactType3Length int = 3

// Field # 780
const nsldsLoanSchoolCode3StartIndex int = 5426
const nsldsLoanSchoolCode3Length int = 8

// Field # 781
const nsldsLoanContactCode3StartIndex int = 5434
const nsldsLoanContactCode3Length int = 8

// Field # 782
const nsldsLoanGradeLevel3StartIndex int = 5442
const nsldsLoanGradeLevel3Length int = 3

// Field # 783
const nsldsLoanAdditionalUnsubsidizedFlag3StartIndex int = 5445
const nsldsLoanAdditionalUnsubsidizedFlag3Length int = 1

// Field # 784
const nsldsLoanCapitalizedInterestFlag3StartIndex int = 5446
const nsldsLoanCapitalizedInterestFlag3Length int = 1

// Field # 785
const nsldsLoanDisbursementAmount3StartIndex int = 5447
const nsldsLoanDisbursementAmount3Length int = 6

// Field # 786
const nsldsLoanDisbursementDate3StartIndex int = 5453
const nsldsLoanDisbursementDate3Length int = 8

// Field # 787
const nsldsLoanConfirmedLoanSubsidyStatus3StartIndex int = 5461
const nsldsLoanConfirmedLoanSubsidyStatus3Length int = 1

// Field # 788
const nsldsLoanSubsidyStatusDate3StartIndex int = 5462
const nsldsLoanSubsidyStatusDate3Length int = 8

// Field # 790
const nsldsLoanSequenceNumber4StartIndex int = 5490
const nsldsLoanSequenceNumber4Length int = 2

// Field # 791
const nsldsLoanDefaultedRecentIndicator4StartIndex int = 5492
const nsldsLoanDefaultedRecentIndicator4Length int = 1

// Field # 792
const nsldsLoanChangeFlag4StartIndex int = 5493
const nsldsLoanChangeFlag4Length int = 1

// Field # 793
const nsldsLoanTypeCode4StartIndex int = 5494
const nsldsLoanTypeCode4Length int = 2

// Field # 794
const nsldsLoanNetAmount4StartIndex int = 5496
const nsldsLoanNetAmount4Length int = 6

// Field # 795
const nsldsLoanCurrentStatusCode4StartIndex int = 5502
const nsldsLoanCurrentStatusCode4Length int = 2

// Field # 796
const nsldsLoanCurrentStatusDate4StartIndex int = 5504
const nsldsLoanCurrentStatusDate4Length int = 8

// Field # 797
const nsldsLoanOutstandingPrincipalBalance4StartIndex int = 5512
const nsldsLoanOutstandingPrincipalBalance4Length int = 6

// Field # 798
const nsldsLoanOutstandingPrincipalBalanceDate4StartIndex int = 5518
const nsldsLoanOutstandingPrincipalBalanceDate4Length int = 8

// Field # 799
const nsldsLoanPeriodBeginDate4StartIndex int = 5526
const nsldsLoanPeriodBeginDate4Length int = 8

// Field # 800
const nsldsLoanPeriodEndDate4StartIndex int = 5534
const nsldsLoanPeriodEndDate4Length int = 8

// Field # 801
const nsldsLoanGuarantyAgencyCode4StartIndex int = 5542
const nsldsLoanGuarantyAgencyCode4Length int = 3

// Field # 802
const nsldsLoanContactType4StartIndex int = 5545
const nsldsLoanContactType4Length int = 3

// Field # 803
const nsldsLoanSchoolCode4StartIndex int = 5548
const nsldsLoanSchoolCode4Length int = 8

// Field # 804
const nsldsLoanContactCode4StartIndex int = 5556
const nsldsLoanContactCode4Length int = 8

// Field # 805
const nsldsLoanGradeLevel4StartIndex int = 5564
const nsldsLoanGradeLevel4Length int = 3

// Field # 806
const nsldsLoanAdditionalUnsubsidizedFlag4StartIndex int = 5567
const nsldsLoanAdditionalUnsubsidizedFlag4Length int = 1

// Field # 807
const nsldsLoanCapitalizedInterestFlag4StartIndex int = 5568
const nsldsLoanCapitalizedInterestFlag4Length int = 1

// Field # 808
const nsldsLoanDisbursementAmount4StartIndex int = 5569
const nsldsLoanDisbursementAmount4Length int = 6

// Field # 809
const nsldsLoanDisbursementDate4StartIndex int = 5575
const nsldsLoanDisbursementDate4Length int = 8

// Field # 810
const nsldsLoanConfirmedLoanSubsidyStatus4StartIndex int = 5583
const nsldsLoanConfirmedLoanSubsidyStatus4Length int = 1

// Field # 811
const nsldsLoanSubsidyStatusDate4StartIndex int = 5584
const nsldsLoanSubsidyStatusDate4Length int = 8

// Field # 813
const nsldsLoanSequenceNumber5StartIndex int = 5612
const nsldsLoanSequenceNumber5Length int = 2

// Field # 814
const nsldsLoanDefaultedRecentIndicator5StartIndex int = 5614
const nsldsLoanDefaultedRecentIndicator5Length int = 1

// Field # 815
const nsldsLoanChangeFlag5StartIndex int = 5615
const nsldsLoanChangeFlag5Length int = 1

// Field # 816
const nsldsLoanTypeCode5StartIndex int = 5616
const nsldsLoanTypeCode5Length int = 2

// Field # 817
const nsldsLoanNetAmount5StartIndex int = 5618
const nsldsLoanNetAmount5Length int = 6

// Field # 818
const nsldsLoanCurrentStatusCode5StartIndex int = 5624
const nsldsLoanCurrentStatusCode5Length int = 2

// Field # 819
const nsldsLoanCurrentStatusDate5StartIndex int = 5626
const nsldsLoanCurrentStatusDate5Length int = 8

// Field # 820
const nsldsLoanOutstandingPrincipalBalance5StartIndex int = 5634
const nsldsLoanOutstandingPrincipalBalance5Length int = 6

// Field # 821
const nsldsLoanOutstandingPrincipalBalanceDate5StartIndex int = 5640
const nsldsLoanOutstandingPrincipalBalanceDate5Length int = 8

// Field # 822
const nsldsLoanPeriodBeginDate5StartIndex int = 5648
const nsldsLoanPeriodBeginDate5Length int = 8

// Field # 823
const nsldsLoanPeriodEndDate5StartIndex int = 5656
const nsldsLoanPeriodEndDate5Length int = 8

// Field # 824
const nsldsLoanGuarantyAgencyCode5StartIndex int = 5664
const nsldsLoanGuarantyAgencyCode5Length int = 3

// Field # 825
const nsldsLoanContactType5StartIndex int = 5667
const nsldsLoanContactType5Length int = 3

// Field # 826
const nsldsLoanSchoolCode5StartIndex int = 5670
const nsldsLoanSchoolCode5Length int = 8

// Field # 827
const nsldsLoanContactCode5StartIndex int = 5678
const nsldsLoanContactCode5Length int = 8

// Field # 828
const nsldsLoanGradeLevel5StartIndex int = 5686
const nsldsLoanGradeLevel5Length int = 3

// Field # 829
const nsldsLoanAdditionalUnsubsidizedFlag5StartIndex int = 5689
const nsldsLoanAdditionalUnsubsidizedFlag5Length int = 1

// Field # 830
const nsldsLoanCapitalizedInterestFlag5StartIndex int = 5690
const nsldsLoanCapitalizedInterestFlag5Length int = 1

// Field # 831
const nsldsLoanDisbursementAmount5StartIndex int = 5691
const nsldsLoanDisbursementAmount5Length int = 6

// Field # 832
const nsldsLoanDisbursementDate5StartIndex int = 5697
const nsldsLoanDisbursementDate5Length int = 8

// Field # 833
const nsldsLoanConfirmedLoanSubsidyStatus5StartIndex int = 5705
const nsldsLoanConfirmedLoanSubsidyStatus5Length int = 1

// Field # 834
const nsldsLoanSubsidyStatusDate5StartIndex int = 5706
const nsldsLoanSubsidyStatusDate5Length int = 8

// Field # 836
const nsldsLoanSequenceNumber6StartIndex int = 5734
const nsldsLoanSequenceNumber6Length int = 2

// Field # 837
const nsldsLoanDefaultedRecentIndicator6StartIndex int = 5736
const nsldsLoanDefaultedRecentIndicator6Length int = 1

// Field # 838
const nsldsLoanChangeFlag6StartIndex int = 5737
const nsldsLoanChangeFlag6Length int = 1

// Field # 839
const nsldsLoanTypeCode6StartIndex int = 5738
const nsldsLoanTypeCode6Length int = 2

// Field # 840
const nsldsLoanNetAmount6StartIndex int = 5740
const nsldsLoanNetAmount6Length int = 6

// Field # 841
const nsldsLoanCurrentStatusCode6StartIndex int = 5746
const nsldsLoanCurrentStatusCode6Length int = 2

// Field # 842
const nsldsLoanCurrentStatusDate6StartIndex int = 5748
const nsldsLoanCurrentStatusDate6Length int = 8

// Field # 843
const nsldsLoanOutstandingPrincipalBalance6StartIndex int = 5756
const nsldsLoanOutstandingPrincipalBalance6Length int = 6

// Field # 844
const nsldsLoanOutstandingPrincipalBalanceDate6StartIndex int = 5762
const nsldsLoanOutstandingPrincipalBalanceDate6Length int = 8

// Field # 845
const nsldsLoanPeriodBeginDate6StartIndex int = 5770
const nsldsLoanPeriodBeginDate6Length int = 8

// Field # 846
const nsldsLoanPeriodEndDate6StartIndex int = 5778
const nsldsLoanPeriodEndDate6Length int = 8

// Field # 847
const nsldsLoanGuarantyAgencyCode6StartIndex int = 5786
const nsldsLoanGuarantyAgencyCode6Length int = 3

// Field # 848
const nsldsLoanContactType6StartIndex int = 5789
const nsldsLoanContactType6Length int = 3

// Field # 849
const nsldsLoanSchoolCode6StartIndex int = 5792
const nsldsLoanSchoolCode6Length int = 8

// Field # 850
const nsldsLoanContactCode6StartIndex int = 5800
const nsldsLoanContactCode6Length int = 8

// Field # 851
const nsldsLoanGradeLevel6StartIndex int = 5808
const nsldsLoanGradeLevel6Length int = 3

// Field # 852
const nsldsLoanAdditionalUnsubsidizedFlag6StartIndex int = 5811
const nsldsLoanAdditionalUnsubsidizedFlag6Length int = 1

// Field # 853
const nsldsLoanCapitalizedInterestFlag6StartIndex int = 5812
const nsldsLoanCapitalizedInterestFlag6Length int = 1

// Field # 854
const nsldsLoanDisbursementAmount6StartIndex int = 5813
const nsldsLoanDisbursementAmount6Length int = 6

// Field # 855
const nsldsLoanDisbursementDate6StartIndex int = 5819
const nsldsLoanDisbursementDate6Length int = 8

// Field # 856
const nsldsLoanConfirmedLoanSubsidyStatus6StartIndex int = 5827
const nsldsLoanConfirmedLoanSubsidyStatus6Length int = 1

// Field # 857
const nsldsLoanSubsidyStatusDate6StartIndex int = 5828
const nsldsLoanSubsidyStatusDate6Length int = 8

// Field # 861
const ftiLabelStartStartIndex int = 7086
const ftiLabelStartLength int = 11

// Field # 862
const studentFTIMReturnedTaxYearStartIndex int = 7097
const studentFTIMReturnedTaxYearLength int = 4

// Field # 863
const studentFTIMFilingStatusCodeStartIndex int = 7101
const studentFTIMFilingStatusCodeLength int = 1

// Field # 864
const studentFTIMAdjustedGrossIncomeStartIndex int = 7102
const studentFTIMAdjustedGrossIncomeLength int = 10

// Field # 865
const studentFTIMNumberOfExemptionsStartIndex int = 7112
const studentFTIMNumberOfExemptionsLength int = 2

// Field # 866
const studentFTIMNumberOfDependentsStartIndex int = 7114
const studentFTIMNumberOfDependentsLength int = 2

// Field # 867
const studentFTIMTotalIncomeEarnedAmountStartIndex int = 7116
const studentFTIMTotalIncomeEarnedAmountLength int = 11

// Field # 868
const studentFTIMIncomeTaxPaidStartIndex int = 7127
const studentFTIMIncomeTaxPaidLength int = 9

// Field # 869
const studentFTIMEducationCreditsStartIndex int = 7136
const studentFTIMEducationCreditsLength int = 9

// Field # 870
const studentFTIMUntaxedIRADistributionsStartIndex int = 7145
const studentFTIMUntaxedIRADistributionsLength int = 11

// Field # 871
const studentFTIMIRADeductibleAndPaymentsStartIndex int = 7156
const studentFTIMIRADeductibleAndPaymentsLength int = 11

// Field # 872
const studentFTIMTaxExemptInterestStartIndex int = 7167
const studentFTIMTaxExemptInterestLength int = 11

// Field # 873
const studentFTIMUntaxedPensionsAmountStartIndex int = 7178
const studentFTIMUntaxedPensionsAmountLength int = 11

// Field # 874
const studentFTIMScheduleCNetProfitLossStartIndex int = 7189
const studentFTIMScheduleCNetProfitLossLength int = 12

// Field # 875
const studentFTIMScheduleAIndicatorStartIndex int = 7201
const studentFTIMScheduleAIndicatorLength int = 1

// Field # 876
const studentFTIMScheduleBIndicatorStartIndex int = 7202
const studentFTIMScheduleBIndicatorLength int = 1

// Field # 877
const studentFTIMScheduleDIndicatorStartIndex int = 7203
const studentFTIMScheduleDIndicatorLength int = 1

// Field # 878
const studentFTIMScheduleEIndicatorStartIndex int = 7204
const studentFTIMScheduleEIndicatorLength int = 1

// Field # 879
const studentFTIMScheduleFIndicatorStartIndex int = 7205
const studentFTIMScheduleFIndicatorLength int = 1

// Field # 880
const studentFTIMScheduleHIndicatorStartIndex int = 7206
const studentFTIMScheduleHIndicatorLength int = 1

// Field # 881
const studentFTIMIRSResponseCodeStartIndex int = 7207
const studentFTIMIRSResponseCodeLength int = 3

// Field # 882
const studentFTIMSpouseReturnedTaxYearStartIndex int = 7210
const studentFTIMSpouseReturnedTaxYearLength int = 4

// Field # 883
const studentFTIMSpouseFilingStatusCodeStartIndex int = 7214
const studentFTIMSpouseFilingStatusCodeLength int = 1

// Field # 884
const studentFTIMSpouseAdjustedGrossIncomeStartIndex int = 7215
const studentFTIMSpouseAdjustedGrossIncomeLength int = 10

// Field # 885
const studentFTIMSpouseNumberOfExemptionsStartIndex int = 7225
const studentFTIMSpouseNumberOfExemptionsLength int = 2

// Field # 886
const studentFTIMSpouseNumberOfDependentsStartIndex int = 7227
const studentFTIMSpouseNumberOfDependentsLength int = 2

// Field # 887
const studentFTIMSpouseTotalIncomeEarnedAmountStartIndex int = 7229
const studentFTIMSpouseTotalIncomeEarnedAmountLength int = 11

// Field # 888
const studentFTIMSpouseIncomeTaxPaidStartIndex int = 7240
const studentFTIMSpouseIncomeTaxPaidLength int = 9

// Field # 889
const studentFTIMSpouseEducationCreditsStartIndex int = 7249
const studentFTIMSpouseEducationCreditsLength int = 9

// Field # 890
const studentFTIMSpouseUntaxedIRADistributionsStartIndex int = 7258
const studentFTIMSpouseUntaxedIRADistributionsLength int = 11

// Field # 891
const studentFTIMSpouseIRADeductibleAndPaymentsStartIndex int = 7269
const studentFTIMSpouseIRADeductibleAndPaymentsLength int = 11

// Field # 892
const studentFTIMSpouseTaxExemptInterestStartIndex int = 7280
const studentFTIMSpouseTaxExemptInterestLength int = 11

// Field # 893
const studentFTIMSpouseUntaxedPensionsAmountStartIndex int = 7291
const studentFTIMSpouseUntaxedPensionsAmountLength int = 11

// Field # 894
const studentFTIMSpouseScheduleCNetProfitLossStartIndex int = 7302
const studentFTIMSpouseScheduleCNetProfitLossLength int = 12

// Field # 895
const studentFTIMSpouseScheduleAIndicatorStartIndex int = 7314
const studentFTIMSpouseScheduleAIndicatorLength int = 1

// Field # 896
const studentFTIMSpouseScheduleBIndicatorStartIndex int = 7315
const studentFTIMSpouseScheduleBIndicatorLength int = 1

// Field # 897
const studentFTIMSpouseScheduleDIndicatorStartIndex int = 7316
const studentFTIMSpouseScheduleDIndicatorLength int = 1

// Field # 898
const studentFTIMSpouseScheduleEIndicatorStartIndex int = 7317
const studentFTIMSpouseScheduleEIndicatorLength int = 1

// Field # 899
const studentFTIMSpouseScheduleFIndicatorStartIndex int = 7318
const studentFTIMSpouseScheduleFIndicatorLength int = 1

// Field # 900
const studentFTIMSpouseScheduleHIndicatorStartIndex int = 7319
const studentFTIMSpouseScheduleHIndicatorLength int = 1

// Field # 901
const studentFTIMSpouseIRSResponseCodeStartIndex int = 7320
const studentFTIMSpouseIRSResponseCodeLength int = 3

// Field # 902
const parentFTIMReturnedTaxYearStartIndex int = 7323
const parentFTIMReturnedTaxYearLength int = 4

// Field # 903
const parentFTIMFilingStatusCodeStartIndex int = 7327
const parentFTIMFilingStatusCodeLength int = 1

// Field # 904
const parentFTIMAdjustedGrossIncomeStartIndex int = 7328
const parentFTIMAdjustedGrossIncomeLength int = 10

// Field # 905
const parentFTIMNumberOfExemptionsStartIndex int = 7338
const parentFTIMNumberOfExemptionsLength int = 2

// Field # 906
const parentFTIMNumberOfDependentsStartIndex int = 7340
const parentFTIMNumberOfDependentsLength int = 2

// Field # 907
const parentFTIMTotalIncomeEarnedAmountStartIndex int = 7342
const parentFTIMTotalIncomeEarnedAmountLength int = 11

// Field # 908
const parentFTIMIncomeTaxPaidStartIndex int = 7353
const parentFTIMIncomeTaxPaidLength int = 9

// Field # 909
const parentFTIMEducationCreditsStartIndex int = 7362
const parentFTIMEducationCreditsLength int = 9

// Field # 910
const parentFTIMUntaxedIRADistributionsStartIndex int = 7371
const parentFTIMUntaxedIRADistributionsLength int = 11

// Field # 911
const parentFTIMIRADeductibleAndPaymentsStartIndex int = 7382
const parentFTIMIRADeductibleAndPaymentsLength int = 11

// Field # 912
const parentFTIMTaxExemptInterestStartIndex int = 7393
const parentFTIMTaxExemptInterestLength int = 11

// Field # 913
const parentFTIMUntaxedPensionsAmountStartIndex int = 7404
const parentFTIMUntaxedPensionsAmountLength int = 11

// Field # 914
const parentFTIMScheduleCNetProfitLossStartIndex int = 7415
const parentFTIMScheduleCNetProfitLossLength int = 12

// Field # 915
const parentFTIMScheduleAIndicatorStartIndex int = 7427
const parentFTIMScheduleAIndicatorLength int = 1

// Field # 916
const parentFTIMScheduleBIndicatorStartIndex int = 7428
const parentFTIMScheduleBIndicatorLength int = 1

// Field # 917
const parentFTIMScheduleDIndicatorStartIndex int = 7429
const parentFTIMScheduleDIndicatorLength int = 1

// Field # 918
const parentFTIMScheduleEIndicatorStartIndex int = 7430
const parentFTIMScheduleEIndicatorLength int = 1

// Field # 919
const parentFTIMScheduleFIndicatorStartIndex int = 7431
const parentFTIMScheduleFIndicatorLength int = 1

// Field # 920
const parentFTIMScheduleHIndicatorStartIndex int = 7432
const parentFTIMScheduleHIndicatorLength int = 1

// Field # 921
const parentFTIMIRSResponseCodeStartIndex int = 7433
const parentFTIMIRSResponseCodeLength int = 3

// Field # 922
const parentFTIMSpouseReturnedTaxYearStartIndex int = 7436
const parentFTIMSpouseReturnedTaxYearLength int = 4

// Field # 923
const parentFTIMSpouseFilingStatusCodeStartIndex int = 7440
const parentFTIMSpouseFilingStatusCodeLength int = 1

// Field # 924
const parentFTIMSpouseAdjustedGrossIncomeStartIndex int = 7441
const parentFTIMSpouseAdjustedGrossIncomeLength int = 10

// Field # 925
const parentFTIMSpouseNumberOfExemptionsStartIndex int = 7451
const parentFTIMSpouseNumberOfExemptionsLength int = 2

// Field # 926
const parentFTIMSpouseNumberOfDependentsStartIndex int = 7453
const parentFTIMSpouseNumberOfDependentsLength int = 2

// Field # 927
const parentFTIMSpouseTotalIncomeEarnedAmountStartIndex int = 7455
const parentFTIMSpouseTotalIncomeEarnedAmountLength int = 11

// Field # 928
const parentFTIMSpouseIncomeTaxPaidStartIndex int = 7466
const parentFTIMSpouseIncomeTaxPaidLength int = 9

// Field # 929
const parentFTIMSpouseEducationCreditsStartIndex int = 7475
const parentFTIMSpouseEducationCreditsLength int = 9

// Field # 930
const parentFTIMSpouseUntaxedIRADistributionsStartIndex int = 7484
const parentFTIMSpouseUntaxedIRADistributionsLength int = 11

// Field # 931
const parentFTIMSpouseIRADeductibleAndPaymentsStartIndex int = 7495
const parentFTIMSpouseIRADeductibleAndPaymentsLength int = 11

// Field # 932
const parentFTIMSpouseTaxExemptInterestStartIndex int = 7506
const parentFTIMSpouseTaxExemptInterestLength int = 11

// Field # 933
const parentFTIMSpouseUntaxedPensionsAmountStartIndex int = 7517
const parentFTIMSpouseUntaxedPensionsAmountLength int = 11

// Field # 934
const parentFTIMSpouseScheduleCNetProfitLossStartIndex int = 7528
const parentFTIMSpouseScheduleCNetProfitLossLength int = 12

// Field # 935
const parentFTIMSpouseScheduleAIndicatorStartIndex int = 7540
const parentFTIMSpouseScheduleAIndicatorLength int = 1

// Field # 936
const parentFTIMSpouseScheduleBIndicatorStartIndex int = 7541
const parentFTIMSpouseScheduleBIndicatorLength int = 1

// Field # 937
const parentFTIMSpouseScheduleDIndicatorStartIndex int = 7542
const parentFTIMSpouseScheduleDIndicatorLength int = 1

// Field # 938
const parentFTIMSpouseScheduleEIndicatorStartIndex int = 7543
const parentFTIMSpouseScheduleEIndicatorLength int = 1

// Field # 939
const parentFTIMSpouseScheduleFIndicatorStartIndex int = 7544
const parentFTIMSpouseScheduleFIndicatorLength int = 1

// Field # 940
const parentFTIMSpouseScheduleHIndicatorStartIndex int = 7545
const parentFTIMSpouseScheduleHIndicatorLength int = 1

// Field # 941
const parentFTIMSpouseIRSResponseCodeStartIndex int = 7546
const parentFTIMSpouseIRSResponseCodeLength int = 3

// Field # 942
const ftiLabelEndStartIndex int = 7549
const ftiLabelEndLength int = 11

// Field # 944
const studentTotalIncomeStartIndex int = 7610
const studentTotalIncomeLength int = 15

// Field # 945
const parentTotalIncomeStartIndex int = 7625
const parentTotalIncomeLength int = 15

// Field # 946
const fisapTotalIncomeStartIndex int = 7640
const fisapTotalIncomeLength int = 15

func ParseISIR(s string) (isirmodels.ISIRecord, error) {
	slog.Debug("Parsing an expected ISIR record from fixed format")
	if len(s) != totalISIRLength {
		slog.Error("Expected ISIR to be length %d, received string with length %d", totalISIRLength, len(s))
		return isirmodels.ISIRecord{}, errors.New(fmt.Sprintf("Input ISIR string is the incorrect length, expected %d and received %d", totalISIRLength, len(s)))
	}

	slog.Info("Parsing record", "FAFSAUUID", strings.TrimSpace(s[fafsaUUIDStartIndex-1:(fafsaUUIDStartIndex-1)+fafsaUUIDLength]),
		"TransactionUUID", strings.TrimSpace(s[transactionUUIDStartIndex-1:(transactionUUIDStartIndex-1)+transactionUUIDLength]),
		"PersonUUID", strings.TrimSpace(s[transactionUUIDStartIndex-1:(transactionUUIDStartIndex-1)+transactionUUIDLength]))

	r := isirmodels.ISIRecord{
		YearIndicator: strings.TrimSpace(s[yearIndicatorStartIndex-1 : (yearIndicatorStartIndex-1)+yearIndicatorLength]), // Field # 1

		FAFSAUUID: strings.TrimSpace(s[fafsaUUIDStartIndex-1 : (fafsaUUIDStartIndex-1)+fafsaUUIDLength]), // Field # 2

		TransactionUUID: strings.TrimSpace(s[transactionUUIDStartIndex-1 : (transactionUUIDStartIndex-1)+transactionUUIDLength]), // Field # 3

		PersonUUID: strings.TrimSpace(s[personUUIDStartIndex-1 : (personUUIDStartIndex-1)+personUUIDLength]), // Field # 4

		TransactionNumber: strings.TrimSpace(s[transactionNumberStartIndex-1 : (transactionNumberStartIndex-1)+transactionNumberLength]), // Field # 5

		DependencyModel: strings.TrimSpace(s[dependencyModelStartIndex-1 : (dependencyModelStartIndex-1)+dependencyModelLength]), // Field # 6

		ApplicationSource: strings.TrimSpace(s[applicationSourceStartIndex-1 : (applicationSourceStartIndex-1)+applicationSourceLength]), // Field # 7

		ApplicationReceiptDate: parseISIRDate(strings.TrimSpace(s[applicationReceiptDateStartIndex-1 : (applicationReceiptDateStartIndex-1)+applicationReceiptDateLength])), // Field # 8

		TransactionSource: strings.TrimSpace(s[transactionSourceStartIndex-1 : (transactionSourceStartIndex-1)+transactionSourceLength]), // Field # 9

		TransactionType: strings.TrimSpace(s[transactionTypeStartIndex-1 : (transactionTypeStartIndex-1)+transactionTypeLength]), // Field # 10

		TransactionLanguage: strings.TrimSpace(s[transactionLanguageStartIndex-1 : (transactionLanguageStartIndex-1)+transactionLanguageLength]), // Field # 11

		TransactionReceiptDate: parseISIRDate(strings.TrimSpace(s[transactionReceiptDateStartIndex-1 : (transactionReceiptDateStartIndex-1)+transactionReceiptDateLength])), // Field # 12

		TransactionProcessedDate: parseISIRDate(strings.TrimSpace(s[transactionProcessedDateStartIndex-1 : (transactionProcessedDateStartIndex-1)+transactionProcessedDateLength])), // Field # 13

		TransactionStatus: strings.TrimSpace(s[transactionStatusStartIndex-1 : (transactionStatusStartIndex-1)+transactionStatusLength]), // Field # 14

		RenewalDataUsed: strings.TrimSpace(s[renewalDataUsedStartIndex-1 : (renewalDataUsedStartIndex-1)+renewalDataUsedLength]), // Field # 15

		FPSCorrectionReason: strings.TrimSpace(s[fpsCorrectionReasonStartIndex-1 : (fpsCorrectionReasonStartIndex-1)+fpsCorrectionReasonLength]), // Field # 16

		SAIChangeFlag: strings.TrimSpace(s[saiChangeFlagStartIndex-1 : (saiChangeFlagStartIndex-1)+saiChangeFlagLength]), // Field # 17

		SAI: strings.TrimSpace(s[saiStartIndex-1 : (saiStartIndex-1)+saiLength]), // Field # 18

		ProvisionalSAI: strings.TrimSpace(s[provisionalSAIStartIndex-1 : (provisionalSAIStartIndex-1)+provisionalSAILength]), // Field # 19

		SAIFormula: strings.TrimSpace(s[saiFormulaStartIndex-1 : (saiFormulaStartIndex-1)+saiFormulaLength]), // Field # 20

		SAIComputationType: strings.TrimSpace(s[saiComputationTypeStartIndex-1 : (saiComputationTypeStartIndex-1)+saiComputationTypeLength]), // Field # 21

		MaxPellIndicator: strings.TrimSpace(s[maxPellIndicatorStartIndex-1 : (maxPellIndicatorStartIndex-1)+maxPellIndicatorLength]), // Field # 22

		MinimumPellIndicator: strings.TrimSpace(s[minimumPellIndicatorStartIndex-1 : (minimumPellIndicatorStartIndex-1)+minimumPellIndicatorLength]), // Field # 23

		StudentFirstName: strings.TrimSpace(s[studentFirstNameStartIndex-1 : (studentFirstNameStartIndex-1)+studentFirstNameLength]), // Field # 25

		StudentMiddleName: strings.TrimSpace(s[studentMiddleNameStartIndex-1 : (studentMiddleNameStartIndex-1)+studentMiddleNameLength]), // Field # 26

		StudentLastName: strings.TrimSpace(s[studentLastNameStartIndex-1 : (studentLastNameStartIndex-1)+studentLastNameLength]), // Field # 27

		StudentSuffix: strings.TrimSpace(s[studentSuffixStartIndex-1 : (studentSuffixStartIndex-1)+studentSuffixLength]), // Field # 28

		StudentDateOfBirth: parseISIRDate(strings.TrimSpace(s[studentDateOfBirthStartIndex-1 : (studentDateOfBirthStartIndex-1)+studentDateOfBirthLength])), // Field # 29

		StudentSSN: strings.TrimSpace(s[studentSSNStartIndex-1 : (studentSSNStartIndex-1)+studentSSNLength]), // Field # 30

		StudentITIN: strings.TrimSpace(s[studentITINStartIndex-1 : (studentITINStartIndex-1)+studentITINLength]), // Field # 31

		StudentPhoneNumber: strings.TrimSpace(s[studentPhoneNumberStartIndex-1 : (studentPhoneNumberStartIndex-1)+studentPhoneNumberLength]), // Field # 32

		StudentEmailAddress: strings.TrimSpace(s[studentEmailAddressStartIndex-1 : (studentEmailAddressStartIndex-1)+studentEmailAddressLength]), // Field # 33

		StudentStreetAddress: strings.TrimSpace(s[studentStreetAddressStartIndex-1 : (studentStreetAddressStartIndex-1)+studentStreetAddressLength]), // Field # 34

		StudentCity: strings.TrimSpace(s[studentCityStartIndex-1 : (studentCityStartIndex-1)+studentCityLength]), // Field # 35

		StudentState: strings.TrimSpace(s[studentStateStartIndex-1 : (studentStateStartIndex-1)+studentStateLength]), // Field # 36

		StudentZipCode: strings.TrimSpace(s[studentZipCodeStartIndex-1 : (studentZipCodeStartIndex-1)+studentZipCodeLength]), // Field # 37

		StudentCountry: strings.TrimSpace(s[studentCountryStartIndex-1 : (studentCountryStartIndex-1)+studentCountryLength]), // Field # 38

		StudentMaritalStatus: strings.TrimSpace(s[studentMaritalStatusStartIndex-1 : (studentMaritalStatusStartIndex-1)+studentMaritalStatusLength]), // Field # 40

		StudentGradeLevel: strings.TrimSpace(s[studentGradeLevelStartIndex-1 : (studentGradeLevelStartIndex-1)+studentGradeLevelLength]), // Field # 41

		StudentFirstBachelorsDegreeBefore2526: strings.TrimSpace(s[studentFirstBachelorsDegreeBefore2526StartIndex-1 : (studentFirstBachelorsDegreeBefore2526StartIndex-1)+studentFirstBachelorsDegreeBefore2526Length]), // Field # 42

		StudentPursuingTeacherCertification: strings.TrimSpace(s[studentPursuingTeacherCertificationStartIndex-1 : (studentPursuingTeacherCertificationStartIndex-1)+studentPursuingTeacherCertificationLength]), // Field # 43

		StudentActiveDuty: strings.TrimSpace(s[studentActiveDutyStartIndex-1 : (studentActiveDutyStartIndex-1)+studentActiveDutyLength]), // Field # 44

		StudentVeteran: strings.TrimSpace(s[studentVeteranStartIndex-1 : (studentVeteranStartIndex-1)+studentVeteranLength]), // Field # 45

		StudentChildOrOtherDependents: strings.TrimSpace(s[studentChildOrOtherDependentsStartIndex-1 : (studentChildOrOtherDependentsStartIndex-1)+studentChildOrOtherDependentsLength]), // Field # 46

		StudentParentsDeceased: strings.TrimSpace(s[studentParentsDeceasedStartIndex-1 : (studentParentsDeceasedStartIndex-1)+studentParentsDeceasedLength]), // Field # 47

		StudentWardOfCourt: strings.TrimSpace(s[studentWardOfCourtStartIndex-1 : (studentWardOfCourtStartIndex-1)+studentWardOfCourtLength]), // Field # 48

		StudentInFosterCare: strings.TrimSpace(s[studentInFosterCareStartIndex-1 : (studentInFosterCareStartIndex-1)+studentInFosterCareLength]), // Field # 49

		StudentEmancipatedMinor: strings.TrimSpace(s[studentEmancipatedMinorStartIndex-1 : (studentEmancipatedMinorStartIndex-1)+studentEmancipatedMinorLength]), // Field # 50

		StudentLegalGuardianship: strings.TrimSpace(s[studentLegalGuardianshipStartIndex-1 : (studentLegalGuardianshipStartIndex-1)+studentLegalGuardianshipLength]), // Field # 51

		StudentPersonalCircumstancesNoneOfTheAbove: strings.TrimSpace(s[studentPersonalCircumstancesNoneOfTheAboveStartIndex-1 : (studentPersonalCircumstancesNoneOfTheAboveStartIndex-1)+studentPersonalCircumstancesNoneOfTheAboveLength]), // Field # 52

		StudentUnaccompaniedHomelessYouthAndSelfSupporting: strings.TrimSpace(s[studentUnaccompaniedHomelessYouthAndSelfSupportingStartIndex-1 : (studentUnaccompaniedHomelessYouthAndSelfSupportingStartIndex-1)+studentUnaccompaniedHomelessYouthAndSelfSupportingLength]), // Field # 53

		StudentUnaccompaniedHomelessGeneral: strings.TrimSpace(s[studentUnaccompaniedHomelessGeneralStartIndex-1 : (studentUnaccompaniedHomelessGeneralStartIndex-1)+studentUnaccompaniedHomelessGeneralLength]), // Field # 54

		StudentUnaccompaniedHomelessHS: strings.TrimSpace(s[studentUnaccompaniedHomelessHSStartIndex-1 : (studentUnaccompaniedHomelessHSStartIndex-1)+studentUnaccompaniedHomelessHSLength]), // Field # 55

		StudentUnaccompaniedHomelessTRIO: strings.TrimSpace(s[studentUnaccompaniedHomelessTRIOStartIndex-1 : (studentUnaccompaniedHomelessTRIOStartIndex-1)+studentUnaccompaniedHomelessTRIOLength]), // Field # 56

		StudentUnaccompaniedHomelessFAA: strings.TrimSpace(s[studentUnaccompaniedHomelessFAAStartIndex-1 : (studentUnaccompaniedHomelessFAAStartIndex-1)+studentUnaccompaniedHomelessFAALength]), // Field # 57

		StudentHomelessnessNoneOfTheAbove: strings.TrimSpace(s[studentHomelessnessNoneOfTheAboveStartIndex-1 : (studentHomelessnessNoneOfTheAboveStartIndex-1)+studentHomelessnessNoneOfTheAboveLength]), // Field # 58

		StudentUnusualCircumstance: strings.TrimSpace(s[studentUnusualCircumstanceStartIndex-1 : (studentUnusualCircumstanceStartIndex-1)+studentUnusualCircumstanceLength]), // Field # 59

		StudentUnsubOnly: strings.TrimSpace(s[studentUnsubOnlyStartIndex-1 : (studentUnsubOnlyStartIndex-1)+studentUnsubOnlyLength]), // Field # 60

		StudentUpdatedFamilySize: strings.TrimSpace(s[studentUpdatedFamilySizeStartIndex-1 : (studentUpdatedFamilySizeStartIndex-1)+studentUpdatedFamilySizeLength]), // Field # 61

		StudentNumberInCollege: strings.TrimSpace(s[studentNumberInCollegeStartIndex-1 : (studentNumberInCollegeStartIndex-1)+studentNumberInCollegeLength]), // Field # 62

		StudentCitizenshipStatus: strings.TrimSpace(s[studentCitizenshipStatusStartIndex-1 : (studentCitizenshipStatusStartIndex-1)+studentCitizenshipStatusLength]), // Field # 63

		StudentANumber: strings.TrimSpace(s[studentANumberStartIndex-1 : (studentANumberStartIndex-1)+studentANumberLength]), // Field # 64

		StudentStateOfLegalResidence: strings.TrimSpace(s[studentStateOfLegalResidenceStartIndex-1 : (studentStateOfLegalResidenceStartIndex-1)+studentStateOfLegalResidenceLength]), // Field # 65

		StudentLegalResidenceDate: parseISIRDate(strings.TrimSpace(s[studentLegalResidenceDateStartIndex-1 : (studentLegalResidenceDateStartIndex-1)+studentLegalResidenceDateLength])), // Field # 66

		StudentEitherParentAttendCollege: strings.TrimSpace(s[studentEitherParentAttendCollegeStartIndex-1 : (studentEitherParentAttendCollegeStartIndex-1)+studentEitherParentAttendCollegeLength]), // Field # 67

		StudentParentKilledInTheLineOfDuty: strings.TrimSpace(s[studentParentKilledInTheLineOfDutyStartIndex-1 : (studentParentKilledInTheLineOfDutyStartIndex-1)+studentParentKilledInTheLineOfDutyLength]), // Field # 68

		StudentHighSchoolCompletionStatus: strings.TrimSpace(s[studentHighSchoolCompletionStatusStartIndex-1 : (studentHighSchoolCompletionStatusStartIndex-1)+studentHighSchoolCompletionStatusLength]), // Field # 69

		StudentHighSchoolName: strings.TrimSpace(s[studentHighSchoolNameStartIndex-1 : (studentHighSchoolNameStartIndex-1)+studentHighSchoolNameLength]), // Field # 70

		StudentHighSchoolCity: strings.TrimSpace(s[studentHighSchoolCityStartIndex-1 : (studentHighSchoolCityStartIndex-1)+studentHighSchoolCityLength]), // Field # 71

		StudentHighSchoolState: strings.TrimSpace(s[studentHighSchoolStateStartIndex-1 : (studentHighSchoolStateStartIndex-1)+studentHighSchoolStateLength]), // Field # 72

		StudentHighSchoolEquivalentDiplomaName: strings.TrimSpace(s[studentHighSchoolEquivalentDiplomaNameStartIndex-1 : (studentHighSchoolEquivalentDiplomaNameStartIndex-1)+studentHighSchoolEquivalentDiplomaNameLength]), // Field # 73

		StudentHighSchoolEquivalentDiplomaState: strings.TrimSpace(s[studentHighSchoolEquivalentDiplomaStateStartIndex-1 : (studentHighSchoolEquivalentDiplomaStateStartIndex-1)+studentHighSchoolEquivalentDiplomaStateLength]), // Field # 74

		StudentManuallyEnteredReceivedEITC: strings.TrimSpace(s[studentManuallyEnteredReceivedEITCStartIndex-1 : (studentManuallyEnteredReceivedEITCStartIndex-1)+studentManuallyEnteredReceivedEITCLength]), // Field # 75

		StudentManuallyEnteredReceivedFederalHousingAssistance: strings.TrimSpace(s[studentManuallyEnteredReceivedFederalHousingAssistanceStartIndex-1 : (studentManuallyEnteredReceivedFederalHousingAssistanceStartIndex-1)+studentManuallyEnteredReceivedFederalHousingAssistanceLength]), // Field # 76

		StudentManuallyEnteredReceivedFreeReducedPriceLunch: strings.TrimSpace(s[studentManuallyEnteredReceivedFreeReducedPriceLunchStartIndex-1 : (studentManuallyEnteredReceivedFreeReducedPriceLunchStartIndex-1)+studentManuallyEnteredReceivedFreeReducedPriceLunchLength]), // Field # 77

		StudentManuallyEnteredReceivedMedicaid: strings.TrimSpace(s[studentManuallyEnteredReceivedMedicaidStartIndex-1 : (studentManuallyEnteredReceivedMedicaidStartIndex-1)+studentManuallyEnteredReceivedMedicaidLength]), // Field # 78

		StudentManuallyEnteredReceivedRefundableCreditFor36BHealthPlan: strings.TrimSpace(s[studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanStartIndex-1 : (studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanStartIndex-1)+studentManuallyEnteredReceivedRefundableCreditFor36BHealthPlanLength]), // Field # 79

		StudentManuallyEnteredReceivedSNAP: strings.TrimSpace(s[studentManuallyEnteredReceivedSNAPStartIndex-1 : (studentManuallyEnteredReceivedSNAPStartIndex-1)+studentManuallyEnteredReceivedSNAPLength]), // Field # 80

		StudentManuallyEnteredReceivedSupplementalSecurityIncome: strings.TrimSpace(s[studentManuallyEnteredReceivedSupplementalSecurityIncomeStartIndex-1 : (studentManuallyEnteredReceivedSupplementalSecurityIncomeStartIndex-1)+studentManuallyEnteredReceivedSupplementalSecurityIncomeLength]), // Field # 81

		StudentManuallyEnteredReceivedTANF: strings.TrimSpace(s[studentManuallyEnteredReceivedTANFStartIndex-1 : (studentManuallyEnteredReceivedTANFStartIndex-1)+studentManuallyEnteredReceivedTANFLength]), // Field # 82

		StudentManuallyEnteredReceivedWIC: strings.TrimSpace(s[studentManuallyEnteredReceivedWICStartIndex-1 : (studentManuallyEnteredReceivedWICStartIndex-1)+studentManuallyEnteredReceivedWICLength]), // Field # 83

		StudentManuallyEnteredFederalBenefitsNoneOfTheAbove: strings.TrimSpace(s[studentManuallyEnteredFederalBenefitsNoneOfTheAboveStartIndex-1 : (studentManuallyEnteredFederalBenefitsNoneOfTheAboveStartIndex-1)+studentManuallyEnteredFederalBenefitsNoneOfTheAboveLength]), // Field # 84

		StudentManuallyEnteredFiled1040Or1040NR: strings.TrimSpace(s[studentManuallyEnteredFiled1040Or1040NRStartIndex-1 : (studentManuallyEnteredFiled1040Or1040NRStartIndex-1)+studentManuallyEnteredFiled1040Or1040NRLength]), // Field # 85

		StudentManuallyEnteredFiledNonUSTaxReturn: strings.TrimSpace(s[studentManuallyEnteredFiledNonUSTaxReturnStartIndex-1 : (studentManuallyEnteredFiledNonUSTaxReturnStartIndex-1)+studentManuallyEnteredFiledNonUSTaxReturnLength]), // Field # 86

		StudentManuallyEnteredFiledJointReturnWithCurrentSpouse: strings.TrimSpace(s[studentManuallyEnteredFiledJointReturnWithCurrentSpouseStartIndex-1 : (studentManuallyEnteredFiledJointReturnWithCurrentSpouseStartIndex-1)+studentManuallyEnteredFiledJointReturnWithCurrentSpouseLength]), // Field # 87

		StudentManuallyEnteredTaxReturnFilingStatus: strings.TrimSpace(s[studentManuallyEnteredTaxReturnFilingStatusStartIndex-1 : (studentManuallyEnteredTaxReturnFilingStatusStartIndex-1)+studentManuallyEnteredTaxReturnFilingStatusLength]), // Field # 88

		StudentManuallyEnteredIncomeEarnedFromWork: strings.TrimSpace(s[studentManuallyEnteredIncomeEarnedFromWorkStartIndex-1 : (studentManuallyEnteredIncomeEarnedFromWorkStartIndex-1)+studentManuallyEnteredIncomeEarnedFromWorkLength]), // Field # 89

		StudentManuallyEnteredTaxExemptInterestIncome: strings.TrimSpace(s[studentManuallyEnteredTaxExemptInterestIncomeStartIndex-1 : (studentManuallyEnteredTaxExemptInterestIncomeStartIndex-1)+studentManuallyEnteredTaxExemptInterestIncomeLength]), // Field # 90

		StudentManuallyEnteredUntaxedPortionsOfIRADistributions: strings.TrimSpace(s[studentManuallyEnteredUntaxedPortionsOfIRADistributionsStartIndex-1 : (studentManuallyEnteredUntaxedPortionsOfIRADistributionsStartIndex-1)+studentManuallyEnteredUntaxedPortionsOfIRADistributionsLength]), // Field # 91

		StudentManuallyEnteredIRARollover: strings.TrimSpace(s[studentManuallyEnteredIRARolloverStartIndex-1 : (studentManuallyEnteredIRARolloverStartIndex-1)+studentManuallyEnteredIRARolloverLength]), // Field # 92

		StudentManuallyEnteredUntaxedPortionsOfPensions: strings.TrimSpace(s[studentManuallyEnteredUntaxedPortionsOfPensionsStartIndex-1 : (studentManuallyEnteredUntaxedPortionsOfPensionsStartIndex-1)+studentManuallyEnteredUntaxedPortionsOfPensionsLength]), // Field # 93

		StudentManuallyEnteredPensionRollover: strings.TrimSpace(s[studentManuallyEnteredPensionRolloverStartIndex-1 : (studentManuallyEnteredPensionRolloverStartIndex-1)+studentManuallyEnteredPensionRolloverLength]), // Field # 94

		StudentManuallyEnteredAdjustedGrossIncome: strings.TrimSpace(s[studentManuallyEnteredAdjustedGrossIncomeStartIndex-1 : (studentManuallyEnteredAdjustedGrossIncomeStartIndex-1)+studentManuallyEnteredAdjustedGrossIncomeLength]), // Field # 95

		StudentManuallyEnteredIncomeTaxPaid: strings.TrimSpace(s[studentManuallyEnteredIncomeTaxPaidStartIndex-1 : (studentManuallyEnteredIncomeTaxPaidStartIndex-1)+studentManuallyEnteredIncomeTaxPaidLength]), // Field # 96

		StudentManuallyEnteredEarnedIncomeTaxCreditReceivedDuringTaxYear: strings.TrimSpace(s[studentManuallyEnteredEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex-1 : (studentManuallyEnteredEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex-1)+studentManuallyEnteredEarnedIncomeTaxCreditReceivedDuringTaxYearLength]), // Field # 97

		StudentManuallyEnteredDeductiblePaymentsToIRAKeoghOther: strings.TrimSpace(s[studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherStartIndex-1 : (studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherStartIndex-1)+studentManuallyEnteredDeductiblePaymentsToIRAKeoghOtherLength]), // Field # 98

		StudentManuallyEnteredEducationCredits: strings.TrimSpace(s[studentManuallyEnteredEducationCreditsStartIndex-1 : (studentManuallyEnteredEducationCreditsStartIndex-1)+studentManuallyEnteredEducationCreditsLength]), // Field # 99

		StudentManuallyEnteredFiledScheduleABDEFH: strings.TrimSpace(s[studentManuallyEnteredFiledScheduleABDEFHStartIndex-1 : (studentManuallyEnteredFiledScheduleABDEFHStartIndex-1)+studentManuallyEnteredFiledScheduleABDEFHLength]), // Field # 100

		StudentManuallyEnteredScheduleCAmount: strings.TrimSpace(s[studentManuallyEnteredScheduleCAmountStartIndex-1 : (studentManuallyEnteredScheduleCAmountStartIndex-1)+studentManuallyEnteredScheduleCAmountLength]), // Field # 101

		StudentManuallyEnteredCollegeGrantAndScholarshipAid: strings.TrimSpace(s[studentManuallyEnteredCollegeGrantAndScholarshipAidStartIndex-1 : (studentManuallyEnteredCollegeGrantAndScholarshipAidStartIndex-1)+studentManuallyEnteredCollegeGrantAndScholarshipAidLength]), // Field # 102

		StudentManuallyEnteredForeignEarnedIncomeExclusion: strings.TrimSpace(s[studentManuallyEnteredForeignEarnedIncomeExclusionStartIndex-1 : (studentManuallyEnteredForeignEarnedIncomeExclusionStartIndex-1)+studentManuallyEnteredForeignEarnedIncomeExclusionLength]), // Field # 103

		StudentManuallyEnteredChildSupportReceived: strings.TrimSpace(s[studentManuallyEnteredChildSupportReceivedStartIndex-1 : (studentManuallyEnteredChildSupportReceivedStartIndex-1)+studentManuallyEnteredChildSupportReceivedLength]), // Field # 104

		StudentManuallyEnteredTotalOfCashSavingsAndCheckingAccounts: strings.TrimSpace(s[studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsStartIndex-1 : (studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsStartIndex-1)+studentManuallyEnteredTotalOfCashSavingsAndCheckingAccountsLength]), // Field # 105

		StudentManuallyEnteredNetWorthOfCurrentInvestments: strings.TrimSpace(s[studentManuallyEnteredNetWorthOfCurrentInvestmentsStartIndex-1 : (studentManuallyEnteredNetWorthOfCurrentInvestmentsStartIndex-1)+studentManuallyEnteredNetWorthOfCurrentInvestmentsLength]), // Field # 106

		StudentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarms: strings.TrimSpace(s[studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsStartIndex-1 : (studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsStartIndex-1)+studentManuallyEnteredNetWorthOfBusinessesAndInvestmentFarmsLength]), // Field # 107

		StudentCollege1: strings.TrimSpace(s[studentCollege1StartIndex-1 : (studentCollege1StartIndex-1)+studentCollege1Length]), // Field # 108

		StudentCollege2: strings.TrimSpace(s[studentCollege2StartIndex-1 : (studentCollege2StartIndex-1)+studentCollege2Length]), // Field # 109

		StudentCollege3: strings.TrimSpace(s[studentCollege3StartIndex-1 : (studentCollege3StartIndex-1)+studentCollege3Length]), // Field # 110

		StudentCollege4: strings.TrimSpace(s[studentCollege4StartIndex-1 : (studentCollege4StartIndex-1)+studentCollege4Length]), // Field # 111

		StudentCollege5: strings.TrimSpace(s[studentCollege5StartIndex-1 : (studentCollege5StartIndex-1)+studentCollege5Length]), // Field # 112

		StudentCollege6: strings.TrimSpace(s[studentCollege6StartIndex-1 : (studentCollege6StartIndex-1)+studentCollege6Length]), // Field # 113

		StudentCollege7: strings.TrimSpace(s[studentCollege7StartIndex-1 : (studentCollege7StartIndex-1)+studentCollege7Length]), // Field # 114

		StudentCollege8: strings.TrimSpace(s[studentCollege8StartIndex-1 : (studentCollege8StartIndex-1)+studentCollege8Length]), // Field # 115

		StudentCollege9: strings.TrimSpace(s[studentCollege9StartIndex-1 : (studentCollege9StartIndex-1)+studentCollege9Length]), // Field # 116

		StudentCollege10: strings.TrimSpace(s[studentCollege10StartIndex-1 : (studentCollege10StartIndex-1)+studentCollege10Length]), // Field # 117

		StudentCollege11: strings.TrimSpace(s[studentCollege11StartIndex-1 : (studentCollege11StartIndex-1)+studentCollege11Length]), // Field # 118

		StudentCollege12: strings.TrimSpace(s[studentCollege12StartIndex-1 : (studentCollege12StartIndex-1)+studentCollege12Length]), // Field # 119

		StudentCollege13: strings.TrimSpace(s[studentCollege13StartIndex-1 : (studentCollege13StartIndex-1)+studentCollege13Length]), // Field # 120

		StudentCollege14: strings.TrimSpace(s[studentCollege14StartIndex-1 : (studentCollege14StartIndex-1)+studentCollege14Length]), // Field # 121

		StudentCollege15: strings.TrimSpace(s[studentCollege15StartIndex-1 : (studentCollege15StartIndex-1)+studentCollege15Length]), // Field # 122

		StudentCollege16: strings.TrimSpace(s[studentCollege16StartIndex-1 : (studentCollege16StartIndex-1)+studentCollege16Length]), // Field # 123

		StudentCollege17: strings.TrimSpace(s[studentCollege17StartIndex-1 : (studentCollege17StartIndex-1)+studentCollege17Length]), // Field # 124

		StudentCollege18: strings.TrimSpace(s[studentCollege18StartIndex-1 : (studentCollege18StartIndex-1)+studentCollege18Length]), // Field # 125

		StudentCollege19: strings.TrimSpace(s[studentCollege19StartIndex-1 : (studentCollege19StartIndex-1)+studentCollege19Length]), // Field # 126

		StudentCollege20: strings.TrimSpace(s[studentCollege20StartIndex-1 : (studentCollege20StartIndex-1)+studentCollege20Length]), // Field # 127

		StudentConsentToRetrieveAndDiscloseFTI: strings.TrimSpace(s[studentConsentToRetrieveAndDiscloseFTIStartIndex-1 : (studentConsentToRetrieveAndDiscloseFTIStartIndex-1)+studentConsentToRetrieveAndDiscloseFTILength]), // Field # 128

		StudentSignature: strings.TrimSpace(s[studentSignatureStartIndex-1 : (studentSignatureStartIndex-1)+studentSignatureLength]), // Field # 129

		StudentSignatureDate: parseISIRDate(strings.TrimSpace(s[studentSignatureDateStartIndex-1 : (studentSignatureDateStartIndex-1)+studentSignatureDateLength])), // Field # 130

		StudentSpouseFirstName: strings.TrimSpace(s[studentSpouseFirstNameStartIndex-1 : (studentSpouseFirstNameStartIndex-1)+studentSpouseFirstNameLength]), // Field # 132

		StudentSpouseMiddleName: strings.TrimSpace(s[studentSpouseMiddleNameStartIndex-1 : (studentSpouseMiddleNameStartIndex-1)+studentSpouseMiddleNameLength]), // Field # 133

		StudentSpouseLastName: strings.TrimSpace(s[studentSpouseLastNameStartIndex-1 : (studentSpouseLastNameStartIndex-1)+studentSpouseLastNameLength]), // Field # 134

		StudentSpouseSuffix: strings.TrimSpace(s[studentSpouseSuffixStartIndex-1 : (studentSpouseSuffixStartIndex-1)+studentSpouseSuffixLength]), // Field # 135

		StudentSpouseDateOfBirth: parseISIRDate(strings.TrimSpace(s[studentSpouseDateOfBirthStartIndex-1 : (studentSpouseDateOfBirthStartIndex-1)+studentSpouseDateOfBirthLength])), // Field # 136

		StudentSpouseSSN: strings.TrimSpace(s[studentSpouseSSNStartIndex-1 : (studentSpouseSSNStartIndex-1)+studentSpouseSSNLength]), // Field # 137

		StudentSpouseITIN: strings.TrimSpace(s[studentSpouseITINStartIndex-1 : (studentSpouseITINStartIndex-1)+studentSpouseITINLength]), // Field # 138

		StudentSpousePhoneNumber: strings.TrimSpace(s[studentSpousePhoneNumberStartIndex-1 : (studentSpousePhoneNumberStartIndex-1)+studentSpousePhoneNumberLength]), // Field # 139

		StudentSpouseEmailAddress: strings.TrimSpace(s[studentSpouseEmailAddressStartIndex-1 : (studentSpouseEmailAddressStartIndex-1)+studentSpouseEmailAddressLength]), // Field # 140

		StudentSpouseStreetAddress: strings.TrimSpace(s[studentSpouseStreetAddressStartIndex-1 : (studentSpouseStreetAddressStartIndex-1)+studentSpouseStreetAddressLength]), // Field # 141

		StudentSpouseCity: strings.TrimSpace(s[studentSpouseCityStartIndex-1 : (studentSpouseCityStartIndex-1)+studentSpouseCityLength]), // Field # 142

		StudentSpouseState: strings.TrimSpace(s[studentSpouseStateStartIndex-1 : (studentSpouseStateStartIndex-1)+studentSpouseStateLength]), // Field # 143

		StudentSpouseZipCode: strings.TrimSpace(s[studentSpouseZipCodeStartIndex-1 : (studentSpouseZipCodeStartIndex-1)+studentSpouseZipCodeLength]), // Field # 144

		StudentSpouseCountry: strings.TrimSpace(s[studentSpouseCountryStartIndex-1 : (studentSpouseCountryStartIndex-1)+studentSpouseCountryLength]), // Field # 145

		StudentSpouseFiled1040Or1040NR: strings.TrimSpace(s[studentSpouseFiled1040Or1040NRStartIndex-1 : (studentSpouseFiled1040Or1040NRStartIndex-1)+studentSpouseFiled1040Or1040NRLength]), // Field # 146

		StudentSpouseFiledNonUSTaxReturn: strings.TrimSpace(s[studentSpouseFiledNonUSTaxReturnStartIndex-1 : (studentSpouseFiledNonUSTaxReturnStartIndex-1)+studentSpouseFiledNonUSTaxReturnLength]), // Field # 147

		StudentSpouseTaxReturnFilingStatus: strings.TrimSpace(s[studentSpouseTaxReturnFilingStatusStartIndex-1 : (studentSpouseTaxReturnFilingStatusStartIndex-1)+studentSpouseTaxReturnFilingStatusLength]), // Field # 148

		StudentSpouseIncomeEarnedFromWork: strings.TrimSpace(s[studentSpouseIncomeEarnedFromWorkStartIndex-1 : (studentSpouseIncomeEarnedFromWorkStartIndex-1)+studentSpouseIncomeEarnedFromWorkLength]), // Field # 149

		StudentSpouseTaxExemptInterestIncome: strings.TrimSpace(s[studentSpouseTaxExemptInterestIncomeStartIndex-1 : (studentSpouseTaxExemptInterestIncomeStartIndex-1)+studentSpouseTaxExemptInterestIncomeLength]), // Field # 150

		StudentSpouseUntaxedPortionsOfIRADistributions: strings.TrimSpace(s[studentSpouseUntaxedPortionsOfIRADistributionsStartIndex-1 : (studentSpouseUntaxedPortionsOfIRADistributionsStartIndex-1)+studentSpouseUntaxedPortionsOfIRADistributionsLength]), // Field # 151

		StudentSpouseIRARollover: strings.TrimSpace(s[studentSpouseIRARolloverStartIndex-1 : (studentSpouseIRARolloverStartIndex-1)+studentSpouseIRARolloverLength]), // Field # 152

		StudentSpouseUntaxedPortionsOfPensions: strings.TrimSpace(s[studentSpouseUntaxedPortionsOfPensionsStartIndex-1 : (studentSpouseUntaxedPortionsOfPensionsStartIndex-1)+studentSpouseUntaxedPortionsOfPensionsLength]), // Field # 153

		StudentSpousePensionRollover: strings.TrimSpace(s[studentSpousePensionRolloverStartIndex-1 : (studentSpousePensionRolloverStartIndex-1)+studentSpousePensionRolloverLength]), // Field # 154

		StudentSpouseAdjustedGrossIncome: strings.TrimSpace(s[studentSpouseAdjustedGrossIncomeStartIndex-1 : (studentSpouseAdjustedGrossIncomeStartIndex-1)+studentSpouseAdjustedGrossIncomeLength]), // Field # 155

		StudentSpouseIncomeTaxPaid: strings.TrimSpace(s[studentSpouseIncomeTaxPaidStartIndex-1 : (studentSpouseIncomeTaxPaidStartIndex-1)+studentSpouseIncomeTaxPaidLength]), // Field # 156

		StudentSpouseDeductiblePaymentsToIRAKeoghOther: strings.TrimSpace(s[studentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex-1 : (studentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex-1)+studentSpouseDeductiblePaymentsToIRAKeoghOtherLength]), // Field # 157

		StudentSpouseEducationCredits: strings.TrimSpace(s[studentSpouseEducationCreditsStartIndex-1 : (studentSpouseEducationCreditsStartIndex-1)+studentSpouseEducationCreditsLength]), // Field # 158

		StudentSpouseFiledScheduleABDEFH: strings.TrimSpace(s[studentSpouseFiledScheduleABDEFHStartIndex-1 : (studentSpouseFiledScheduleABDEFHStartIndex-1)+studentSpouseFiledScheduleABDEFHLength]), // Field # 159

		StudentSpouseScheduleCAmount: strings.TrimSpace(s[studentSpouseScheduleCAmountStartIndex-1 : (studentSpouseScheduleCAmountStartIndex-1)+studentSpouseScheduleCAmountLength]), // Field # 160

		StudentSpouseForeignEarnedIncomeExclusion: strings.TrimSpace(s[studentSpouseForeignEarnedIncomeExclusionStartIndex-1 : (studentSpouseForeignEarnedIncomeExclusionStartIndex-1)+studentSpouseForeignEarnedIncomeExclusionLength]), // Field # 161

		StudentSpouseConsentToRetrieveAndDiscloseFTI: strings.TrimSpace(s[studentSpouseConsentToRetrieveAndDiscloseFTIStartIndex-1 : (studentSpouseConsentToRetrieveAndDiscloseFTIStartIndex-1)+studentSpouseConsentToRetrieveAndDiscloseFTILength]), // Field # 162

		StudentSpouseSignature: strings.TrimSpace(s[studentSpouseSignatureStartIndex-1 : (studentSpouseSignatureStartIndex-1)+studentSpouseSignatureLength]), // Field # 163

		StudentSpouseSignatureDate: parseISIRDate(strings.TrimSpace(s[studentSpouseSignatureDateStartIndex-1 : (studentSpouseSignatureDateStartIndex-1)+studentSpouseSignatureDateLength])), // Field # 164

		ParentFirstName: strings.TrimSpace(s[parentFirstNameStartIndex-1 : (parentFirstNameStartIndex-1)+parentFirstNameLength]), // Field # 166

		ParentMiddleName: strings.TrimSpace(s[parentMiddleNameStartIndex-1 : (parentMiddleNameStartIndex-1)+parentMiddleNameLength]), // Field # 167

		ParentLastName: strings.TrimSpace(s[parentLastNameStartIndex-1 : (parentLastNameStartIndex-1)+parentLastNameLength]), // Field # 168

		ParentSuffix: strings.TrimSpace(s[parentSuffixStartIndex-1 : (parentSuffixStartIndex-1)+parentSuffixLength]), // Field # 169

		ParentDateOfBirth: parseISIRDate(strings.TrimSpace(s[parentDateOfBirthStartIndex-1 : (parentDateOfBirthStartIndex-1)+parentDateOfBirthLength])), // Field # 170

		ParentSSN: strings.TrimSpace(s[parentSSNStartIndex-1 : (parentSSNStartIndex-1)+parentSSNLength]), // Field # 171

		ParentITIN: strings.TrimSpace(s[parentITINStartIndex-1 : (parentITINStartIndex-1)+parentITINLength]), // Field # 172

		ParentPhoneNumber: strings.TrimSpace(s[parentPhoneNumberStartIndex-1 : (parentPhoneNumberStartIndex-1)+parentPhoneNumberLength]), // Field # 173

		ParentEmailAddress: strings.TrimSpace(s[parentEmailAddressStartIndex-1 : (parentEmailAddressStartIndex-1)+parentEmailAddressLength]), // Field # 174

		ParentStreetAddress: strings.TrimSpace(s[parentStreetAddressStartIndex-1 : (parentStreetAddressStartIndex-1)+parentStreetAddressLength]), // Field # 175

		ParentCity: strings.TrimSpace(s[parentCityStartIndex-1 : (parentCityStartIndex-1)+parentCityLength]), // Field # 176

		ParentState: strings.TrimSpace(s[parentStateStartIndex-1 : (parentStateStartIndex-1)+parentStateLength]), // Field # 177

		ParentZipCode: strings.TrimSpace(s[parentZipCodeStartIndex-1 : (parentZipCodeStartIndex-1)+parentZipCodeLength]), // Field # 178

		ParentCountry: strings.TrimSpace(s[parentCountryStartIndex-1 : (parentCountryStartIndex-1)+parentCountryLength]), // Field # 179

		ParentMaritalStatus: strings.TrimSpace(s[parentMaritalStatusStartIndex-1 : (parentMaritalStatusStartIndex-1)+parentMaritalStatusLength]), // Field # 180

		ParentStateOfLegalResidence: strings.TrimSpace(s[parentStateOfLegalResidenceStartIndex-1 : (parentStateOfLegalResidenceStartIndex-1)+parentStateOfLegalResidenceLength]), // Field # 181

		ParentLegalResidenceDate: parseISIRDate(strings.TrimSpace(s[parentLegalResidenceDateStartIndex-1 : (parentLegalResidenceDateStartIndex-1)+parentLegalResidenceDateLength])), // Field # 182

		ParentUpdatedFamilySize: strings.TrimSpace(s[parentUpdatedFamilySizeStartIndex-1 : (parentUpdatedFamilySizeStartIndex-1)+parentUpdatedFamilySizeLength]), // Field # 183

		ParentNumberInCollege: strings.TrimSpace(s[parentNumberInCollegeStartIndex-1 : (parentNumberInCollegeStartIndex-1)+parentNumberInCollegeLength]), // Field # 184

		ParentReceivedEITC: strings.TrimSpace(s[parentReceivedEITCStartIndex-1 : (parentReceivedEITCStartIndex-1)+parentReceivedEITCLength]), // Field # 185

		ParentReceivedFederalHousingAssistance: strings.TrimSpace(s[parentReceivedFederalHousingAssistanceStartIndex-1 : (parentReceivedFederalHousingAssistanceStartIndex-1)+parentReceivedFederalHousingAssistanceLength]), // Field # 186

		ParentReceivedFreeReducedPriceLunch: strings.TrimSpace(s[parentReceivedFreeReducedPriceLunchStartIndex-1 : (parentReceivedFreeReducedPriceLunchStartIndex-1)+parentReceivedFreeReducedPriceLunchLength]), // Field # 187

		ParentReceivedMedicaid: strings.TrimSpace(s[parentReceivedMedicaidStartIndex-1 : (parentReceivedMedicaidStartIndex-1)+parentReceivedMedicaidLength]), // Field # 188

		ParentReceivedRefundableCreditFor36BHealthPlan: strings.TrimSpace(s[parentReceivedRefundableCreditFor36BHealthPlanStartIndex-1 : (parentReceivedRefundableCreditFor36BHealthPlanStartIndex-1)+parentReceivedRefundableCreditFor36BHealthPlanLength]), // Field # 189

		ParentReceivedSNAP: strings.TrimSpace(s[parentReceivedSNAPStartIndex-1 : (parentReceivedSNAPStartIndex-1)+parentReceivedSNAPLength]), // Field # 190

		ParentReceivedSupplementalSecurityIncome: strings.TrimSpace(s[parentReceivedSupplementalSecurityIncomeStartIndex-1 : (parentReceivedSupplementalSecurityIncomeStartIndex-1)+parentReceivedSupplementalSecurityIncomeLength]), // Field # 191

		ParentReceivedTANF: strings.TrimSpace(s[parentReceivedTANFStartIndex-1 : (parentReceivedTANFStartIndex-1)+parentReceivedTANFLength]), // Field # 192

		ParentReceivedWIC: strings.TrimSpace(s[parentReceivedWICStartIndex-1 : (parentReceivedWICStartIndex-1)+parentReceivedWICLength]), // Field # 193

		ParentFederalBenefitsNoneOfTheAbove: strings.TrimSpace(s[parentFederalBenefitsNoneOfTheAboveStartIndex-1 : (parentFederalBenefitsNoneOfTheAboveStartIndex-1)+parentFederalBenefitsNoneOfTheAboveLength]), // Field # 194

		ParentFiled1040Or1040NR: strings.TrimSpace(s[parentFiled1040Or1040NRStartIndex-1 : (parentFiled1040Or1040NRStartIndex-1)+parentFiled1040Or1040NRLength]), // Field # 195

		ParentFileNonUSTaxReturn: strings.TrimSpace(s[parentFileNonUSTaxReturnStartIndex-1 : (parentFileNonUSTaxReturnStartIndex-1)+parentFileNonUSTaxReturnLength]), // Field # 196

		ParentFiledJointReturnWithCurrentSpouse: strings.TrimSpace(s[parentFiledJointReturnWithCurrentSpouseStartIndex-1 : (parentFiledJointReturnWithCurrentSpouseStartIndex-1)+parentFiledJointReturnWithCurrentSpouseLength]), // Field # 197

		ParentTaxReturnFilingStatus: strings.TrimSpace(s[parentTaxReturnFilingStatusStartIndex-1 : (parentTaxReturnFilingStatusStartIndex-1)+parentTaxReturnFilingStatusLength]), // Field # 198

		ParentIncomeEarnedFromWork: strings.TrimSpace(s[parentIncomeEarnedFromWorkStartIndex-1 : (parentIncomeEarnedFromWorkStartIndex-1)+parentIncomeEarnedFromWorkLength]), // Field # 199

		ParentTaxExemptInterestIncome: strings.TrimSpace(s[parentTaxExemptInterestIncomeStartIndex-1 : (parentTaxExemptInterestIncomeStartIndex-1)+parentTaxExemptInterestIncomeLength]), // Field # 200

		ParentUntaxedPortionsOfIRADistributions: strings.TrimSpace(s[parentUntaxedPortionsOfIRADistributionsStartIndex-1 : (parentUntaxedPortionsOfIRADistributionsStartIndex-1)+parentUntaxedPortionsOfIRADistributionsLength]), // Field # 201

		ParentIRARollover: strings.TrimSpace(s[parentIRARolloverStartIndex-1 : (parentIRARolloverStartIndex-1)+parentIRARolloverLength]), // Field # 202

		ParentUntaxedPortionsOfPensions: strings.TrimSpace(s[parentUntaxedPortionsOfPensionsStartIndex-1 : (parentUntaxedPortionsOfPensionsStartIndex-1)+parentUntaxedPortionsOfPensionsLength]), // Field # 203

		ParentPensionRollover: strings.TrimSpace(s[parentPensionRolloverStartIndex-1 : (parentPensionRolloverStartIndex-1)+parentPensionRolloverLength]), // Field # 204

		ParentAdjustedGrossIncome: strings.TrimSpace(s[parentAdjustedGrossIncomeStartIndex-1 : (parentAdjustedGrossIncomeStartIndex-1)+parentAdjustedGrossIncomeLength]), // Field # 205

		ParentIncomeTaxPaid: strings.TrimSpace(s[parentIncomeTaxPaidStartIndex-1 : (parentIncomeTaxPaidStartIndex-1)+parentIncomeTaxPaidLength]), // Field # 206

		ParentEarnedIncomeTaxCreditReceivedDuringTaxYear: strings.TrimSpace(s[parentEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex-1 : (parentEarnedIncomeTaxCreditReceivedDuringTaxYearStartIndex-1)+parentEarnedIncomeTaxCreditReceivedDuringTaxYearLength]), // Field # 207

		ParentDeductiblePaymentsToIRAKeoghOther: strings.TrimSpace(s[parentDeductiblePaymentsToIRAKeoghOtherStartIndex-1 : (parentDeductiblePaymentsToIRAKeoghOtherStartIndex-1)+parentDeductiblePaymentsToIRAKeoghOtherLength]), // Field # 208

		ParentEducationCredits: strings.TrimSpace(s[parentEducationCreditsStartIndex-1 : (parentEducationCreditsStartIndex-1)+parentEducationCreditsLength]), // Field # 209

		ParentFiledScheduleABDEFH: strings.TrimSpace(s[parentFiledScheduleABDEFHStartIndex-1 : (parentFiledScheduleABDEFHStartIndex-1)+parentFiledScheduleABDEFHLength]), // Field # 210

		ParentScheduleCAmount: strings.TrimSpace(s[parentScheduleCAmountStartIndex-1 : (parentScheduleCAmountStartIndex-1)+parentScheduleCAmountLength]), // Field # 211

		ParentCollegeGrantAndScholarshipAid: strings.TrimSpace(s[parentCollegeGrantAndScholarshipAidStartIndex-1 : (parentCollegeGrantAndScholarshipAidStartIndex-1)+parentCollegeGrantAndScholarshipAidLength]), // Field # 212

		ParentForeignEarnedIncomeExclusion: strings.TrimSpace(s[parentForeignEarnedIncomeExclusionStartIndex-1 : (parentForeignEarnedIncomeExclusionStartIndex-1)+parentForeignEarnedIncomeExclusionLength]), // Field # 213

		ParentChildSupportReceived: strings.TrimSpace(s[parentChildSupportReceivedStartIndex-1 : (parentChildSupportReceivedStartIndex-1)+parentChildSupportReceivedLength]), // Field # 214

		ParentTotalOfCashSavingsAndCheckingAccounts: strings.TrimSpace(s[parentTotalOfCashSavingsAndCheckingAccountsStartIndex-1 : (parentTotalOfCashSavingsAndCheckingAccountsStartIndex-1)+parentTotalOfCashSavingsAndCheckingAccountsLength]), // Field # 215

		ParentNetWorthOfCurrentInvestments: strings.TrimSpace(s[parentNetWorthOfCurrentInvestmentsStartIndex-1 : (parentNetWorthOfCurrentInvestmentsStartIndex-1)+parentNetWorthOfCurrentInvestmentsLength]), // Field # 216

		ParentNetWorthOfBusinessesAndInvestmentFarms: strings.TrimSpace(s[parentNetWorthOfBusinessesAndInvestmentFarmsStartIndex-1 : (parentNetWorthOfBusinessesAndInvestmentFarmsStartIndex-1)+parentNetWorthOfBusinessesAndInvestmentFarmsLength]), // Field # 217

		ParentConsentToRetrieveAndDiscloseFTI: strings.TrimSpace(s[parentConsentToRetrieveAndDiscloseFTIStartIndex-1 : (parentConsentToRetrieveAndDiscloseFTIStartIndex-1)+parentConsentToRetrieveAndDiscloseFTILength]), // Field # 218

		ParentSignature: strings.TrimSpace(s[parentSignatureStartIndex-1 : (parentSignatureStartIndex-1)+parentSignatureLength]), // Field # 219

		ParentSignatureDate: parseISIRDate(strings.TrimSpace(s[parentSignatureDateStartIndex-1 : (parentSignatureDateStartIndex-1)+parentSignatureDateLength])), // Field # 220

		ParentSpouseFirstName: strings.TrimSpace(s[parentSpouseFirstNameStartIndex-1 : (parentSpouseFirstNameStartIndex-1)+parentSpouseFirstNameLength]), // Field # 222

		ParentSpouseMiddleName: strings.TrimSpace(s[parentSpouseMiddleNameStartIndex-1 : (parentSpouseMiddleNameStartIndex-1)+parentSpouseMiddleNameLength]), // Field # 223

		ParentSpouseLastName: strings.TrimSpace(s[parentSpouseLastNameStartIndex-1 : (parentSpouseLastNameStartIndex-1)+parentSpouseLastNameLength]), // Field # 224

		ParentSpouseSuffix: strings.TrimSpace(s[parentSpouseSuffixStartIndex-1 : (parentSpouseSuffixStartIndex-1)+parentSpouseSuffixLength]), // Field # 225

		ParentSpouseDateOfBirth: parseISIRDate(strings.TrimSpace(s[parentSpouseDateOfBirthStartIndex-1 : (parentSpouseDateOfBirthStartIndex-1)+parentSpouseDateOfBirthLength])), // Field # 226

		ParentSpouseSSN: strings.TrimSpace(s[parentSpouseSSNStartIndex-1 : (parentSpouseSSNStartIndex-1)+parentSpouseSSNLength]), // Field # 227

		ParentSpouseITIN: strings.TrimSpace(s[parentSpouseITINStartIndex-1 : (parentSpouseITINStartIndex-1)+parentSpouseITINLength]), // Field # 228

		ParentSpousePhoneNumber: strings.TrimSpace(s[parentSpousePhoneNumberStartIndex-1 : (parentSpousePhoneNumberStartIndex-1)+parentSpousePhoneNumberLength]), // Field # 229

		ParentSpouseEmailAddress: strings.TrimSpace(s[parentSpouseEmailAddressStartIndex-1 : (parentSpouseEmailAddressStartIndex-1)+parentSpouseEmailAddressLength]), // Field # 230

		ParentSpouseStreetAddress: strings.TrimSpace(s[parentSpouseStreetAddressStartIndex-1 : (parentSpouseStreetAddressStartIndex-1)+parentSpouseStreetAddressLength]), // Field # 231

		ParentSpouseCity: strings.TrimSpace(s[parentSpouseCityStartIndex-1 : (parentSpouseCityStartIndex-1)+parentSpouseCityLength]), // Field # 232

		ParentSpouseState: strings.TrimSpace(s[parentSpouseStateStartIndex-1 : (parentSpouseStateStartIndex-1)+parentSpouseStateLength]), // Field # 233

		ParentSpouseZipCode: strings.TrimSpace(s[parentSpouseZipCodeStartIndex-1 : (parentSpouseZipCodeStartIndex-1)+parentSpouseZipCodeLength]), // Field # 234

		ParentSpouseCountry: strings.TrimSpace(s[parentSpouseCountryStartIndex-1 : (parentSpouseCountryStartIndex-1)+parentSpouseCountryLength]), // Field # 235

		ParentSpouseFiled1040Or1040NR: strings.TrimSpace(s[parentSpouseFiled1040Or1040NRStartIndex-1 : (parentSpouseFiled1040Or1040NRStartIndex-1)+parentSpouseFiled1040Or1040NRLength]), // Field # 236

		ParentSpouseFileNonUSTaxReturn: strings.TrimSpace(s[parentSpouseFileNonUSTaxReturnStartIndex-1 : (parentSpouseFileNonUSTaxReturnStartIndex-1)+parentSpouseFileNonUSTaxReturnLength]), // Field # 237

		ParentSpouseTaxReturnFilingStatus: strings.TrimSpace(s[parentSpouseTaxReturnFilingStatusStartIndex-1 : (parentSpouseTaxReturnFilingStatusStartIndex-1)+parentSpouseTaxReturnFilingStatusLength]), // Field # 238

		ParentSpouseIncomeEarnedFromWork: strings.TrimSpace(s[parentSpouseIncomeEarnedFromWorkStartIndex-1 : (parentSpouseIncomeEarnedFromWorkStartIndex-1)+parentSpouseIncomeEarnedFromWorkLength]), // Field # 239

		ParentSpouseTaxExemptInterestIncome: strings.TrimSpace(s[parentSpouseTaxExemptInterestIncomeStartIndex-1 : (parentSpouseTaxExemptInterestIncomeStartIndex-1)+parentSpouseTaxExemptInterestIncomeLength]), // Field # 240

		ParentSpouseUntaxedPortionsOfIRADistributions: strings.TrimSpace(s[parentSpouseUntaxedPortionsOfIRADistributionsStartIndex-1 : (parentSpouseUntaxedPortionsOfIRADistributionsStartIndex-1)+parentSpouseUntaxedPortionsOfIRADistributionsLength]), // Field # 241

		ParentSpouseIRARollover: strings.TrimSpace(s[parentSpouseIRARolloverStartIndex-1 : (parentSpouseIRARolloverStartIndex-1)+parentSpouseIRARolloverLength]), // Field # 242

		ParentSpouseUntaxedPortionsOfPensions: strings.TrimSpace(s[parentSpouseUntaxedPortionsOfPensionsStartIndex-1 : (parentSpouseUntaxedPortionsOfPensionsStartIndex-1)+parentSpouseUntaxedPortionsOfPensionsLength]), // Field # 243

		ParentSpousePensionRollover: strings.TrimSpace(s[parentSpousePensionRolloverStartIndex-1 : (parentSpousePensionRolloverStartIndex-1)+parentSpousePensionRolloverLength]), // Field # 244

		ParentSpouseAdjustedGrossIncome: strings.TrimSpace(s[parentSpouseAdjustedGrossIncomeStartIndex-1 : (parentSpouseAdjustedGrossIncomeStartIndex-1)+parentSpouseAdjustedGrossIncomeLength]), // Field # 245

		ParentSpouseIncomeTaxPaid: strings.TrimSpace(s[parentSpouseIncomeTaxPaidStartIndex-1 : (parentSpouseIncomeTaxPaidStartIndex-1)+parentSpouseIncomeTaxPaidLength]), // Field # 246

		ParentSpouseDeductiblePaymentsToIRAKeoghOther: strings.TrimSpace(s[parentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex-1 : (parentSpouseDeductiblePaymentsToIRAKeoghOtherStartIndex-1)+parentSpouseDeductiblePaymentsToIRAKeoghOtherLength]), // Field # 247

		ParentSpouseEducationCredits: strings.TrimSpace(s[parentSpouseEducationCreditsStartIndex-1 : (parentSpouseEducationCreditsStartIndex-1)+parentSpouseEducationCreditsLength]), // Field # 248

		ParentSpouseFiledScheduleABDEFH: strings.TrimSpace(s[parentSpouseFiledScheduleABDEFHStartIndex-1 : (parentSpouseFiledScheduleABDEFHStartIndex-1)+parentSpouseFiledScheduleABDEFHLength]), // Field # 249

		ParentSpouseScheduleCAmount: strings.TrimSpace(s[parentSpouseScheduleCAmountStartIndex-1 : (parentSpouseScheduleCAmountStartIndex-1)+parentSpouseScheduleCAmountLength]), // Field # 250

		ParentSpouseForeignEarnedIncomeExclusion: strings.TrimSpace(s[parentSpouseForeignEarnedIncomeExclusionStartIndex-1 : (parentSpouseForeignEarnedIncomeExclusionStartIndex-1)+parentSpouseForeignEarnedIncomeExclusionLength]), // Field # 251

		ParentSpouseConsentToRetrieveAndDiscloseFTI: strings.TrimSpace(s[parentSpouseConsentToRetrieveAndDiscloseFTIStartIndex-1 : (parentSpouseConsentToRetrieveAndDiscloseFTIStartIndex-1)+parentSpouseConsentToRetrieveAndDiscloseFTILength]), // Field # 252

		ParentSpouseSignature: strings.TrimSpace(s[parentSpouseSignatureStartIndex-1 : (parentSpouseSignatureStartIndex-1)+parentSpouseSignatureLength]), // Field # 253

		ParentSpouseSignatureDate: parseISIRDate(strings.TrimSpace(s[parentSpouseSignatureDateStartIndex-1 : (parentSpouseSignatureDateStartIndex-1)+parentSpouseSignatureDateLength])), // Field # 254

		PreparerFirstName: strings.TrimSpace(s[preparerFirstNameStartIndex-1 : (preparerFirstNameStartIndex-1)+preparerFirstNameLength]), // Field # 256

		PreparerLastName: strings.TrimSpace(s[preparerLastNameStartIndex-1 : (preparerLastNameStartIndex-1)+preparerLastNameLength]), // Field # 257

		PreparerSSN: strings.TrimSpace(s[preparerSSNStartIndex-1 : (preparerSSNStartIndex-1)+preparerSSNLength]), // Field # 258

		PreparerEIN: strings.TrimSpace(s[preparerEINStartIndex-1 : (preparerEINStartIndex-1)+preparerEINLength]), // Field # 259

		PreparerAffiliation: strings.TrimSpace(s[preparerAffiliationStartIndex-1 : (preparerAffiliationStartIndex-1)+preparerAffiliationLength]), // Field # 260

		PreparerStreetAddress: strings.TrimSpace(s[preparerStreetAddressStartIndex-1 : (preparerStreetAddressStartIndex-1)+preparerStreetAddressLength]), // Field # 261

		PreparerCity: strings.TrimSpace(s[preparerCityStartIndex-1 : (preparerCityStartIndex-1)+preparerCityLength]), // Field # 262

		PreparerState: strings.TrimSpace(s[preparerStateStartIndex-1 : (preparerStateStartIndex-1)+preparerStateLength]), // Field # 263

		PreparerZipCode: strings.TrimSpace(s[preparerZipCodeStartIndex-1 : (preparerZipCodeStartIndex-1)+preparerZipCodeLength]), // Field # 264

		PreparerSignature: strings.TrimSpace(s[preparerSignatureStartIndex-1 : (preparerSignatureStartIndex-1)+preparerSignatureLength]), // Field # 265

		PreparerSignatureDate: parseISIRDate(strings.TrimSpace(s[preparerSignatureDateStartIndex-1 : (preparerSignatureDateStartIndex-1)+preparerSignatureDateLength])), // Field # 266

		StudentAffirmationStatus: strings.TrimSpace(s[studentAffirmationStatusStartIndex-1 : (studentAffirmationStatusStartIndex-1)+studentAffirmationStatusLength]), // Field # 268

		StudentSpouseAffirmationStatus: strings.TrimSpace(s[studentSpouseAffirmationStatusStartIndex-1 : (studentSpouseAffirmationStatusStartIndex-1)+studentSpouseAffirmationStatusLength]), // Field # 269

		ParentAffirmationStatus: strings.TrimSpace(s[parentAffirmationStatusStartIndex-1 : (parentAffirmationStatusStartIndex-1)+parentAffirmationStatusLength]), // Field # 270

		ParentSpouseOrPartnerAffirmationStatus: strings.TrimSpace(s[parentSpouseOrPartnerAffirmationStatusStartIndex-1 : (parentSpouseOrPartnerAffirmationStatusStartIndex-1)+parentSpouseOrPartnerAffirmationStatusLength]), // Field # 271

		StudentDateConsentGranted: parseISIRDate(strings.TrimSpace(s[studentDateConsentGrantedStartIndex-1 : (studentDateConsentGrantedStartIndex-1)+studentDateConsentGrantedLength])), // Field # 272

		StudentSpouseDateConsentGranted: parseISIRDate(strings.TrimSpace(s[studentSpouseDateConsentGrantedStartIndex-1 : (studentSpouseDateConsentGrantedStartIndex-1)+studentSpouseDateConsentGrantedLength])), // Field # 273

		ParentDateConsentGranted: parseISIRDate(strings.TrimSpace(s[parentDateConsentGrantedStartIndex-1 : (parentDateConsentGrantedStartIndex-1)+parentDateConsentGrantedLength])), // Field # 274

		ParentSpouseOrPartnerDateConsentGranted: parseISIRDate(strings.TrimSpace(s[parentSpouseOrPartnerDateConsentGrantedStartIndex-1 : (parentSpouseOrPartnerDateConsentGrantedStartIndex-1)+parentSpouseOrPartnerDateConsentGrantedLength])), // Field # 275

		StudentTransunionMatchStatus: strings.TrimSpace(s[studentTransunionMatchStatusStartIndex-1 : (studentTransunionMatchStatusStartIndex-1)+studentTransunionMatchStatusLength]), // Field # 276

		StudentSpouseTransunionMatchStatus: strings.TrimSpace(s[studentSpouseTransunionMatchStatusStartIndex-1 : (studentSpouseTransunionMatchStatusStartIndex-1)+studentSpouseTransunionMatchStatusLength]), // Field # 277

		StudentParentTransunionMatchStatus: strings.TrimSpace(s[studentParentTransunionMatchStatusStartIndex-1 : (studentParentTransunionMatchStatusStartIndex-1)+studentParentTransunionMatchStatusLength]), // Field # 278

		StudentParentSpouseTransunionMatchStatus: strings.TrimSpace(s[studentParentSpouseTransunionMatchStatusStartIndex-1 : (studentParentSpouseTransunionMatchStatusStartIndex-1)+studentParentSpouseTransunionMatchStatusLength]), // Field # 279

		CorrectionAppliedAgainstTransactionNumber: strings.TrimSpace(s[correctionAppliedAgainstTransactionNumberStartIndex-1 : (correctionAppliedAgainstTransactionNumberStartIndex-1)+correctionAppliedAgainstTransactionNumberLength]), // Field # 280

		ProfessionalJudgement: strings.TrimSpace(s[professionalJudgementStartIndex-1 : (professionalJudgementStartIndex-1)+professionalJudgementLength]), // Field # 281

		DependencyOverrideIndicator: strings.TrimSpace(s[dependencyOverrideIndicatorStartIndex-1 : (dependencyOverrideIndicatorStartIndex-1)+dependencyOverrideIndicatorLength]), // Field # 282

		FAAFederalSchoolCode: strings.TrimSpace(s[fAAFederalSchoolCodeStartIndex-1 : (fAAFederalSchoolCodeStartIndex-1)+fAAFederalSchoolCodeLength]), // Field # 283

		FAASignature: strings.TrimSpace(s[fAASignatureStartIndex-1 : (fAASignatureStartIndex-1)+fAASignatureLength]), // Field # 284

		IASGIndicator: strings.TrimSpace(s[iASGIndicatorStartIndex-1 : (iASGIndicatorStartIndex-1)+iASGIndicatorLength]), // Field # 285

		ChildrenOfFallenHeroesIndicator: strings.TrimSpace(s[childrenOfFallenHeroesIndicatorStartIndex-1 : (childrenOfFallenHeroesIndicatorStartIndex-1)+childrenOfFallenHeroesIndicatorLength]), // Field # 286

		ElectronicTransactionIndicatorDestinationNumber: strings.TrimSpace(s[electronicTransactionIndicatorDestinationNumberStartIndex-1 : (electronicTransactionIndicatorDestinationNumberStartIndex-1)+electronicTransactionIndicatorDestinationNumberLength]), // Field # 287

		StudentSignatureSource: strings.TrimSpace(s[studentSignatureSourceStartIndex-1 : (studentSignatureSourceStartIndex-1)+studentSignatureSourceLength]), // Field # 288

		StudentSpouseSignatureSource: strings.TrimSpace(s[studentSpouseSignatureSourceStartIndex-1 : (studentSpouseSignatureSourceStartIndex-1)+studentSpouseSignatureSourceLength]), // Field # 289

		ParentSignatureSource: strings.TrimSpace(s[parentSignatureSourceStartIndex-1 : (parentSignatureSourceStartIndex-1)+parentSignatureSourceLength]), // Field # 290

		ParentSpouseOrPartnerSignatureSource: strings.TrimSpace(s[parentSpouseOrPartnerSignatureSourceStartIndex-1 : (parentSpouseOrPartnerSignatureSourceStartIndex-1)+parentSpouseOrPartnerSignatureSourceLength]), // Field # 291

		SpecialHandlingIndicator: strings.TrimSpace(s[specialHandlingIndicatorStartIndex-1 : (specialHandlingIndicatorStartIndex-1)+specialHandlingIndicatorLength]), // Field # 292

		AddressOnlyChangeFlag: strings.TrimSpace(s[addressOnlyChangeFlagStartIndex-1 : (addressOnlyChangeFlagStartIndex-1)+addressOnlyChangeFlagLength]), // Field # 293

		FPSPushedISIRFlag: strings.TrimSpace(s[fpsPushedISIRFlagStartIndex-1 : (fpsPushedISIRFlagStartIndex-1)+fpsPushedISIRFlagLength]), // Field # 294

		RejectStatusChangeFlag: strings.TrimSpace(s[rejectStatusChangeFlagStartIndex-1 : (rejectStatusChangeFlagStartIndex-1)+rejectStatusChangeFlagLength]), // Field # 295

		VerificationTrackingFlag: strings.TrimSpace(s[verificationTrackingFlagStartIndex-1 : (verificationTrackingFlagStartIndex-1)+verificationTrackingFlagLength]), // Field # 296

		StudentSelectedForVerification: strings.TrimSpace(s[studentSelectedForVerificationStartIndex-1 : (studentSelectedForVerificationStartIndex-1)+studentSelectedForVerificationLength]), // Field # 297

		IncarceratedApplicantFlag: strings.TrimSpace(s[incarceratedApplicantFlagStartIndex-1 : (incarceratedApplicantFlagStartIndex-1)+incarceratedApplicantFlagLength]), // Field # 298

		NSLDSTransactionNumber: strings.TrimSpace(s[nsldsTransactionNumberStartIndex-1 : (nsldsTransactionNumberStartIndex-1)+nsldsTransactionNumberLength]), // Field # 299

		NSLDSDatabaseResultsFlag: strings.TrimSpace(s[nsldsDatabaseResultsFlagStartIndex-1 : (nsldsDatabaseResultsFlagStartIndex-1)+nsldsDatabaseResultsFlagLength]), // Field # 300

		HighSchoolFlag: strings.TrimSpace(s[highSchoolFlagStartIndex-1 : (highSchoolFlagStartIndex-1)+highSchoolFlagLength]), // Field # 301

		StudentTotalFederalWorkStudyEarnings: strings.TrimSpace(s[studentTotalFederalWorkStudyEarningsStartIndex-1 : (studentTotalFederalWorkStudyEarningsStartIndex-1)+studentTotalFederalWorkStudyEarningsLength]), // Field # 302

		StudentSpouseTotalFederalWorkStudyEarnings: strings.TrimSpace(s[studentSpouseTotalFederalWorkStudyEarningsStartIndex-1 : (studentSpouseTotalFederalWorkStudyEarningsStartIndex-1)+studentSpouseTotalFederalWorkStudyEarningsLength]), // Field # 303

		ParentTotalFederalWorkStudyEarnings: strings.TrimSpace(s[parentTotalFederalWorkStudyEarningsStartIndex-1 : (parentTotalFederalWorkStudyEarningsStartIndex-1)+parentTotalFederalWorkStudyEarningsLength]), // Field # 304

		ParentSpouseOrPartnerTotalFederalWorkStudyEarnings: strings.TrimSpace(s[parentSpouseOrPartnerTotalFederalWorkStudyEarningsStartIndex-1 : (parentSpouseOrPartnerTotalFederalWorkStudyEarningsStartIndex-1)+parentSpouseOrPartnerTotalFederalWorkStudyEarningsLength]), // Field # 305

		TotalParentAllowancesAgainstIncome: strings.TrimSpace(s[totalParentAllowancesAgainstIncomeStartIndex-1 : (totalParentAllowancesAgainstIncomeStartIndex-1)+totalParentAllowancesAgainstIncomeLength]), // Field # 306

		ParentPayrollTaxAllowance: strings.TrimSpace(s[parentPayrollTaxAllowanceStartIndex-1 : (parentPayrollTaxAllowanceStartIndex-1)+parentPayrollTaxAllowanceLength]), // Field # 307

		ParentIncomeProtectionAllowance: strings.TrimSpace(s[parentIncomeProtectionAllowanceStartIndex-1 : (parentIncomeProtectionAllowanceStartIndex-1)+parentIncomeProtectionAllowanceLength]), // Field # 308

		ParentEmploymentExpenseAllowance: strings.TrimSpace(s[parentEmploymentExpenseAllowanceStartIndex-1 : (parentEmploymentExpenseAllowanceStartIndex-1)+parentEmploymentExpenseAllowanceLength]), // Field # 309

		ParentAvailableIncome: strings.TrimSpace(s[parentAvailableIncomeStartIndex-1 : (parentAvailableIncomeStartIndex-1)+parentAvailableIncomeLength]), // Field # 310

		ParentAdjustedAvailableIncome: strings.TrimSpace(s[parentAdjustedAvailableIncomeStartIndex-1 : (parentAdjustedAvailableIncomeStartIndex-1)+parentAdjustedAvailableIncomeLength]), // Field # 311

		ParentContribution: strings.TrimSpace(s[parentContributionStartIndex-1 : (parentContributionStartIndex-1)+parentContributionLength]), // Field # 312

		StudentPayrollTaxAllowance: strings.TrimSpace(s[studentPayrollTaxAllowanceStartIndex-1 : (studentPayrollTaxAllowanceStartIndex-1)+studentPayrollTaxAllowanceLength]), // Field # 313

		StudentIncomeProtectionAllowance: strings.TrimSpace(s[studentIncomeProtectionAllowanceStartIndex-1 : (studentIncomeProtectionAllowanceStartIndex-1)+studentIncomeProtectionAllowanceLength]), // Field # 314

		StudentAllowanceForParentsNegativeAdjustedAvailableIncome: strings.TrimSpace(s[studentAllowanceForParentsNegativeAdjustedAvailableIncomeStartIndex-1 : (studentAllowanceForParentsNegativeAdjustedAvailableIncomeStartIndex-1)+studentAllowanceForParentsNegativeAdjustedAvailableIncomeLength]), // Field # 315

		StudentEmploymentExpenseAllowance: strings.TrimSpace(s[studentEmploymentExpenseAllowanceStartIndex-1 : (studentEmploymentExpenseAllowanceStartIndex-1)+studentEmploymentExpenseAllowanceLength]), // Field # 316

		TotalStudentAllowancesAgainstIncome: strings.TrimSpace(s[totalStudentAllowancesAgainstIncomeStartIndex-1 : (totalStudentAllowancesAgainstIncomeStartIndex-1)+totalStudentAllowancesAgainstIncomeLength]), // Field # 317

		StudentAvailableIncome: strings.TrimSpace(s[studentAvailableIncomeStartIndex-1 : (studentAvailableIncomeStartIndex-1)+studentAvailableIncomeLength]), // Field # 318

		StudentContributionFromIncome: strings.TrimSpace(s[studentContributionFromIncomeStartIndex-1 : (studentContributionFromIncomeStartIndex-1)+studentContributionFromIncomeLength]), // Field # 319

		StudentAdjustedAvailableIncome: strings.TrimSpace(s[studentAdjustedAvailableIncomeStartIndex-1 : (studentAdjustedAvailableIncomeStartIndex-1)+studentAdjustedAvailableIncomeLength]), // Field # 320

		TotalStudentContributionFromSAAI: strings.TrimSpace(s[totalStudentContributionFromSAAIStartIndex-1 : (totalStudentContributionFromSAAIStartIndex-1)+totalStudentContributionFromSAAILength]), // Field # 321

		ParentDiscretionaryNetWorth: strings.TrimSpace(s[parentDiscretionaryNetWorthStartIndex-1 : (parentDiscretionaryNetWorthStartIndex-1)+parentDiscretionaryNetWorthLength]), // Field # 322

		ParentNetWorth: strings.TrimSpace(s[parentNetWorthStartIndex-1 : (parentNetWorthStartIndex-1)+parentNetWorthLength]), // Field # 323

		ParentAssetProtectionAllowance: strings.TrimSpace(s[parentAssetProtectionAllowanceStartIndex-1 : (parentAssetProtectionAllowanceStartIndex-1)+parentAssetProtectionAllowanceLength]), // Field # 324

		ParentContributionFromAssets: strings.TrimSpace(s[parentContributionFromAssetsStartIndex-1 : (parentContributionFromAssetsStartIndex-1)+parentContributionFromAssetsLength]), // Field # 325

		StudentNetWorth: strings.TrimSpace(s[studentNetWorthStartIndex-1 : (studentNetWorthStartIndex-1)+studentNetWorthLength]), // Field # 326

		StudentAssetProtectionAllowance: strings.TrimSpace(s[studentAssetProtectionAllowanceStartIndex-1 : (studentAssetProtectionAllowanceStartIndex-1)+studentAssetProtectionAllowanceLength]), // Field # 327

		StudentContributionFromAssets: strings.TrimSpace(s[studentContributionFromAssetsStartIndex-1 : (studentContributionFromAssetsStartIndex-1)+studentContributionFromAssetsLength]), // Field # 328

		AssumedStudentFamilySize: strings.TrimSpace(s[assumedStudentFamilySizeStartIndex-1 : (assumedStudentFamilySizeStartIndex-1)+assumedStudentFamilySizeLength]), // Field # 329

		AssumedParentFamilySize: strings.TrimSpace(s[assumedParentFamilySizeStartIndex-1 : (assumedParentFamilySizeStartIndex-1)+assumedParentFamilySizeLength]), // Field # 330

		StudentFirstNameCHVFlags: strings.TrimSpace(s[studentFirstNameCHVFlagsStartIndex-1 : (studentFirstNameCHVFlagsStartIndex-1)+studentFirstNameCHVFlagsLength]), // Field # 331

		StudentMiddleNameCHVFlags: strings.TrimSpace(s[studentMiddleNameCHVFlagsStartIndex-1 : (studentMiddleNameCHVFlagsStartIndex-1)+studentMiddleNameCHVFlagsLength]), // Field # 332

		StudentLastNameCHVFLags: strings.TrimSpace(s[studentLastNameCHVFLagsStartIndex-1 : (studentLastNameCHVFLagsStartIndex-1)+studentLastNameCHVFLagsLength]), // Field # 333

		StudentSuffixCHVFLags: strings.TrimSpace(s[studentSuffixCHVFLagsStartIndex-1 : (studentSuffixCHVFLagsStartIndex-1)+studentSuffixCHVFLagsLength]), // Field # 334

		StudentDateOfBirthCHVFLags: strings.TrimSpace(s[studentDateOfBirthCHVFLagsStartIndex-1 : (studentDateOfBirthCHVFLagsStartIndex-1)+studentDateOfBirthCHVFLagsLength]), // Field # 335

		StudentSSNCHVFlags: strings.TrimSpace(s[studentSSNCHVFlagsStartIndex-1 : (studentSSNCHVFlagsStartIndex-1)+studentSSNCHVFlagsLength]), // Field # 336

		StudentITINCHVFLags: strings.TrimSpace(s[studentITINCHVFLagsStartIndex-1 : (studentITINCHVFLagsStartIndex-1)+studentITINCHVFLagsLength]), // Field # 337

		StudentPhoneNumberCHVFlags: strings.TrimSpace(s[studentPhoneNumberCHVFlagsStartIndex-1 : (studentPhoneNumberCHVFlagsStartIndex-1)+studentPhoneNumberCHVFlagsLength]), // Field # 338

		StudentEmailAddressCHVFlags: strings.TrimSpace(s[studentEmailAddressCHVFlagsStartIndex-1 : (studentEmailAddressCHVFlagsStartIndex-1)+studentEmailAddressCHVFlagsLength]), // Field # 339

		StudentStreetAddressCHVFlags: strings.TrimSpace(s[studentStreetAddressCHVFlagsStartIndex-1 : (studentStreetAddressCHVFlagsStartIndex-1)+studentStreetAddressCHVFlagsLength]), // Field # 340

		StudentCityCHVFLags: strings.TrimSpace(s[studentCityCHVFLagsStartIndex-1 : (studentCityCHVFLagsStartIndex-1)+studentCityCHVFLagsLength]), // Field # 341

		StudentStateCHVFlags: strings.TrimSpace(s[studentStateCHVFlagsStartIndex-1 : (studentStateCHVFlagsStartIndex-1)+studentStateCHVFlagsLength]), // Field # 342

		StudentZipCodeCHVFlags: strings.TrimSpace(s[studentZipCodeCHVFlagsStartIndex-1 : (studentZipCodeCHVFlagsStartIndex-1)+studentZipCodeCHVFlagsLength]), // Field # 343

		StudentCountryCHVFlags: strings.TrimSpace(s[studentCountryCHVFlagsStartIndex-1 : (studentCountryCHVFlagsStartIndex-1)+studentCountryCHVFlagsLength]), // Field # 344

		StudentMaritalStatusCHVFlags: strings.TrimSpace(s[studentMaritalStatusCHVFlagsStartIndex-1 : (studentMaritalStatusCHVFlagsStartIndex-1)+studentMaritalStatusCHVFlagsLength]), // Field # 345

		StudentGradeLevelInCollegeCHVFlags: strings.TrimSpace(s[studentGradeLevelInCollegeCHVFlagsStartIndex-1 : (studentGradeLevelInCollegeCHVFlagsStartIndex-1)+studentGradeLevelInCollegeCHVFlagsLength]), // Field # 346

		StudentFirstBachelorsDegreeBeforeSchoolYearCHVFlags: strings.TrimSpace(s[studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsStartIndex-1 : (studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsStartIndex-1)+studentFirstBachelorsDegreeBeforeSchoolYearCHVFlagsLength]), // Field # 347

		StudentPursuingTeacherCertificationCHVFlags: strings.TrimSpace(s[studentPursuingTeacherCertificationCHVFlagsStartIndex-1 : (studentPursuingTeacherCertificationCHVFlagsStartIndex-1)+studentPursuingTeacherCertificationCHVFlagsLength]), // Field # 348

		StudentActiveDutyCHVFlags: strings.TrimSpace(s[studentActiveDutyCHVFlagsStartIndex-1 : (studentActiveDutyCHVFlagsStartIndex-1)+studentActiveDutyCHVFlagsLength]), // Field # 349

		StudentVeteranCHVFlags: strings.TrimSpace(s[studentVeteranCHVFlagsStartIndex-1 : (studentVeteranCHVFlagsStartIndex-1)+studentVeteranCHVFlagsLength]), // Field # 350

		StudentChildOrOtherDependentsCHVFlags: strings.TrimSpace(s[studentChildOrOtherDependentsCHVFlagsStartIndex-1 : (studentChildOrOtherDependentsCHVFlagsStartIndex-1)+studentChildOrOtherDependentsCHVFlagsLength]), // Field # 351

		StudentParentsDeceasedCHVFlags: strings.TrimSpace(s[studentParentsDeceasedCHVFlagsStartIndex-1 : (studentParentsDeceasedCHVFlagsStartIndex-1)+studentParentsDeceasedCHVFlagsLength]), // Field # 352

		StudentWardOfCourtCHVFlags: strings.TrimSpace(s[studentWardOfCourtCHVFlagsStartIndex-1 : (studentWardOfCourtCHVFlagsStartIndex-1)+studentWardOfCourtCHVFlagsLength]), // Field # 353

		StudentInFosterCareCHVFlags: strings.TrimSpace(s[studentInFosterCareCHVFlagsStartIndex-1 : (studentInFosterCareCHVFlagsStartIndex-1)+studentInFosterCareCHVFlagsLength]), // Field # 354

		StudentEmancipatedMinorCHVFlags: strings.TrimSpace(s[studentEmancipatedMinorCHVFlagsStartIndex-1 : (studentEmancipatedMinorCHVFlagsStartIndex-1)+studentEmancipatedMinorCHVFlagsLength]), // Field # 355

		StudentLegalGuardianshipCHVFlags: strings.TrimSpace(s[studentLegalGuardianshipCHVFlagsStartIndex-1 : (studentLegalGuardianshipCHVFlagsStartIndex-1)+studentLegalGuardianshipCHVFlagsLength]), // Field # 356

		StudentPersonalCircumstancesNoneOfTheAboveCHVFlags: strings.TrimSpace(s[studentPersonalCircumstancesNoneOfTheAboveCHVFlagsStartIndex-1 : (studentPersonalCircumstancesNoneOfTheAboveCHVFlagsStartIndex-1)+studentPersonalCircumstancesNoneOfTheAboveCHVFlagsLength]), // Field # 357

		StudentUnaccompaniedHomelessYouthOrIsUnaccompaniedAtRIskSelfSupportingCHVFlags: strings.TrimSpace(s[studentUnaccompaniedHomelessYouthOrIsUnaccompaniedAtRIskSelfSupportingCHVFlagsStartIndex-1 : (studentUnaccompaniedHomelessYouthOrIsUnaccompaniedAtRIskSelfSupportingCHVFlagsStartIndex-1)+studentUnaccompaniedHomelessYouthOrIsUnaccompaniedAtRIskSelfSupportingCHVFlagsLength]), // Field # 358

		StudentUnaccompaniedAndHomelessGeneralCHVFlags: strings.TrimSpace(s[studentUnaccompaniedAndHomelessGeneralCHVFlagsStartIndex-1 : (studentUnaccompaniedAndHomelessGeneralCHVFlagsStartIndex-1)+studentUnaccompaniedAndHomelessGeneralCHVFlagsLength]), // Field # 359

		StudentUnaccompaniedAndHomelessHSCHVFlags: strings.TrimSpace(s[studentUnaccompaniedAndHomelessHSCHVFlagsStartIndex-1 : (studentUnaccompaniedAndHomelessHSCHVFlagsStartIndex-1)+studentUnaccompaniedAndHomelessHSCHVFlagsLength]), // Field # 360

		StudentUnaccompaniedAndHomelessTRIOCHVFlags: strings.TrimSpace(s[studentUnaccompaniedAndHomelessTRIOCHVFlagsStartIndex-1 : (studentUnaccompaniedAndHomelessTRIOCHVFlagsStartIndex-1)+studentUnaccompaniedAndHomelessTRIOCHVFlagsLength]), // Field # 361

		StudentUnaccompaniedAndHomelessFAACHVFlags: strings.TrimSpace(s[studentUnaccompaniedAndHomelessFAACHVFlagsStartIndex-1 : (studentUnaccompaniedAndHomelessFAACHVFlagsStartIndex-1)+studentUnaccompaniedAndHomelessFAACHVFlagsLength]), // Field # 362

		StudentHomelessnessNoneOfTheAboveCHVFlags: strings.TrimSpace(s[studentHomelessnessNoneOfTheAboveCHVFlagsStartIndex-1 : (studentHomelessnessNoneOfTheAboveCHVFlagsStartIndex-1)+studentHomelessnessNoneOfTheAboveCHVFlagsLength]), // Field # 363

		StudentHasUnusualCircumstanceCHVFlags: strings.TrimSpace(s[studentHasUnusualCircumstanceCHVFlagsStartIndex-1 : (studentHasUnusualCircumstanceCHVFlagsStartIndex-1)+studentHasUnusualCircumstanceCHVFlagsLength]), // Field # 364

		StudentUnsubOnlyCHVFlags: strings.TrimSpace(s[studentUnsubOnlyCHVFlagsStartIndex-1 : (studentUnsubOnlyCHVFlagsStartIndex-1)+studentUnsubOnlyCHVFlagsLength]), // Field # 365

		StudentUpdatedFamilySizeCHVFlags: strings.TrimSpace(s[studentUpdatedFamilySizeCHVFlagsStartIndex-1 : (studentUpdatedFamilySizeCHVFlagsStartIndex-1)+studentUpdatedFamilySizeCHVFlagsLength]), // Field # 366

		StudentNumberInCollegeCorrectionCHVFlags: strings.TrimSpace(s[studentNumberInCollegeCorrectionCHVFlagsStartIndex-1 : (studentNumberInCollegeCorrectionCHVFlagsStartIndex-1)+studentNumberInCollegeCorrectionCHVFlagsLength]), // Field # 367

		StudentCitizenshipStatusCorrectionCHVFlags: strings.TrimSpace(s[studentCitizenshipStatusCorrectionCHVFlagsStartIndex-1 : (studentCitizenshipStatusCorrectionCHVFlagsStartIndex-1)+studentCitizenshipStatusCorrectionCHVFlagsLength]), // Field # 368

		StudentANumberCHVFlags: strings.TrimSpace(s[studentANumberCHVFlagsStartIndex-1 : (studentANumberCHVFlagsStartIndex-1)+studentANumberCHVFlagsLength]), // Field # 369

		StudentStateOfLegalResidenceCHVFlags: strings.TrimSpace(s[studentStateOfLegalResidenceCHVFlagsStartIndex-1 : (studentStateOfLegalResidenceCHVFlagsStartIndex-1)+studentStateOfLegalResidenceCHVFlagsLength]), // Field # 370

		StudentLegalResidenceDateCHVFlags: strings.TrimSpace(s[studentLegalResidenceDateCHVFlagsStartIndex-1 : (studentLegalResidenceDateCHVFlagsStartIndex-1)+studentLegalResidenceDateCHVFlagsLength]), // Field # 371

		StudentEitherParentAttendCollegeCHVFlags: strings.TrimSpace(s[studentEitherParentAttendCollegeCHVFlagsStartIndex-1 : (studentEitherParentAttendCollegeCHVFlagsStartIndex-1)+studentEitherParentAttendCollegeCHVFlagsLength]), // Field # 372

		StudentParentKilledInTheLineOfDutyCHVFlags: strings.TrimSpace(s[studentParentKilledInTheLineOfDutyCHVFlagsStartIndex-1 : (studentParentKilledInTheLineOfDutyCHVFlagsStartIndex-1)+studentParentKilledInTheLineOfDutyCHVFlagsLength]), // Field # 373

		StudentHighSchoolCompletionStatusCHVFlags: strings.TrimSpace(s[studentHighSchoolCompletionStatusCHVFlagsStartIndex-1 : (studentHighSchoolCompletionStatusCHVFlagsStartIndex-1)+studentHighSchoolCompletionStatusCHVFlagsLength]), // Field # 374

		StudentHighSchoolNameCHVFlags: strings.TrimSpace(s[studentHighSchoolNameCHVFlagsStartIndex-1 : (studentHighSchoolNameCHVFlagsStartIndex-1)+studentHighSchoolNameCHVFlagsLength]), // Field # 375

		StudentHighSchoolCityCHVFlags: strings.TrimSpace(s[studentHighSchoolCityCHVFlagsStartIndex-1 : (studentHighSchoolCityCHVFlagsStartIndex-1)+studentHighSchoolCityCHVFlagsLength]), // Field # 376

		StudentHighSchoolStateCHVFlags: strings.TrimSpace(s[studentHighSchoolStateCHVFlagsStartIndex-1 : (studentHighSchoolStateCHVFlagsStartIndex-1)+studentHighSchoolStateCHVFlagsLength]), // Field # 377

		StudentHighSchoolEquivalentDiplomaNameCHVFlags: strings.TrimSpace(s[studentHighSchoolEquivalentDiplomaNameCHVFlagsStartIndex-1 : (studentHighSchoolEquivalentDiplomaNameCHVFlagsStartIndex-1)+studentHighSchoolEquivalentDiplomaNameCHVFlagsLength]), // Field # 378

		StudentHighSchoolEquivalentDiplomaStateCHVFlags: strings.TrimSpace(s[studentHighSchoolEquivalentDiplomaStateCHVFlagsStartIndex-1 : (studentHighSchoolEquivalentDiplomaStateCHVFlagsStartIndex-1)+studentHighSchoolEquivalentDiplomaStateCHVFlagsLength]), // Field # 379

		StudentReceivedEITCCHVFlags: strings.TrimSpace(s[studentReceivedEITCCHVFlagsStartIndex-1 : (studentReceivedEITCCHVFlagsStartIndex-1)+studentReceivedEITCCHVFlagsLength]), // Field # 380

		StudentReceivedFederalHousingAssistanceCHVFlags: strings.TrimSpace(s[studentReceivedFederalHousingAssistanceCHVFlagsStartIndex-1 : (studentReceivedFederalHousingAssistanceCHVFlagsStartIndex-1)+studentReceivedFederalHousingAssistanceCHVFlagsLength]), // Field # 381

		StudentReceivedFreeReducedPriceLunchCHVFlags: strings.TrimSpace(s[studentReceivedFreeReducedPriceLunchCHVFlagsStartIndex-1 : (studentReceivedFreeReducedPriceLunchCHVFlagsStartIndex-1)+studentReceivedFreeReducedPriceLunchCHVFlagsLength]), // Field # 382

		StudentReceivedMedicaidCHVFlags: strings.TrimSpace(s[studentReceivedMedicaidCHVFlagsStartIndex-1 : (studentReceivedMedicaidCHVFlagsStartIndex-1)+studentReceivedMedicaidCHVFlagsLength]), // Field # 383

		StudentReceivedRefundableCreditFor36BHealthPlanCHVFlags: strings.TrimSpace(s[studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex-1 : (studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex-1)+studentReceivedRefundableCreditFor36BHealthPlanCHVFlagsLength]), // Field # 384

		StudentReceivedSNAPCHVFlags: strings.TrimSpace(s[studentReceivedSNAPCHVFlagsStartIndex-1 : (studentReceivedSNAPCHVFlagsStartIndex-1)+studentReceivedSNAPCHVFlagsLength]), // Field # 385

		StudentReceivedSupplementalSecurityIncomeCHVFlags: strings.TrimSpace(s[studentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex-1 : (studentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex-1)+studentReceivedSupplementalSecurityIncomeCHVFlagsLength]), // Field # 386

		StudentReceivedTANFCHVFlags: strings.TrimSpace(s[studentReceivedTANFCHVFlagsStartIndex-1 : (studentReceivedTANFCHVFlagsStartIndex-1)+studentReceivedTANFCHVFlagsLength]), // Field # 387

		StudentReceivedWICCHVFlags: strings.TrimSpace(s[studentReceivedWICCHVFlagsStartIndex-1 : (studentReceivedWICCHVFlagsStartIndex-1)+studentReceivedWICCHVFlagsLength]), // Field # 388

		StudentFederalBenefitsNoneOfTheAboveCHVFlags: strings.TrimSpace(s[studentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex-1 : (studentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex-1)+studentFederalBenefitsNoneOfTheAboveCHVFlagsLength]), // Field # 389

		StudentFiled1040Or1040NRCHVFlags: strings.TrimSpace(s[studentFiled1040Or1040NRCHVFlagsStartIndex-1 : (studentFiled1040Or1040NRCHVFlagsStartIndex-1)+studentFiled1040Or1040NRCHVFlagsLength]), // Field # 390

		StudentFiledNonUSTaxReturnCHVFlags: strings.TrimSpace(s[studentFiledNonUSTaxReturnCHVFlagsStartIndex-1 : (studentFiledNonUSTaxReturnCHVFlagsStartIndex-1)+studentFiledNonUSTaxReturnCHVFlagsLength]), // Field # 391

		StudentFiledJointReturnWithCurrentSpouseCHVFlags: strings.TrimSpace(s[studentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex-1 : (studentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex-1)+studentFiledJointReturnWithCurrentSpouseCHVFlagsLength]), // Field # 392

		StudentTaxReturnFilingStatusCHVFlags: strings.TrimSpace(s[studentTaxReturnFilingStatusCHVFlagsStartIndex-1 : (studentTaxReturnFilingStatusCHVFlagsStartIndex-1)+studentTaxReturnFilingStatusCHVFlagsLength]), // Field # 393

		StudentIncomeEarnedFromWorkCorrectionCHVFlags: strings.TrimSpace(s[studentIncomeEarnedFromWorkCorrectionCHVFlagsStartIndex-1 : (studentIncomeEarnedFromWorkCorrectionCHVFlagsStartIndex-1)+studentIncomeEarnedFromWorkCorrectionCHVFlagsLength]), // Field # 394

		StudentTaxExemptInterestIncomeCHVFlags: strings.TrimSpace(s[studentTaxExemptInterestIncomeCHVFlagsStartIndex-1 : (studentTaxExemptInterestIncomeCHVFlagsStartIndex-1)+studentTaxExemptInterestIncomeCHVFlagsLength]), // Field # 395

		StudentUntaxedPortionsOfIRADistributionsCHVFlags: strings.TrimSpace(s[studentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex-1 : (studentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex-1)+studentUntaxedPortionsOfIRADistributionsCHVFlagsLength]), // Field # 396

		StudentIRARolloverCHVFlags: strings.TrimSpace(s[studentIRARolloverCHVFlagsStartIndex-1 : (studentIRARolloverCHVFlagsStartIndex-1)+studentIRARolloverCHVFlagsLength]), // Field # 397

		StudentUntaxedPortionsOfPensionsCHVFlags: strings.TrimSpace(s[studentUntaxedPortionsOfPensionsCHVFlagsStartIndex-1 : (studentUntaxedPortionsOfPensionsCHVFlagsStartIndex-1)+studentUntaxedPortionsOfPensionsCHVFlagsLength]), // Field # 398

		StudentPensionRolloverCHVFlags: strings.TrimSpace(s[studentPensionRolloverCHVFlagsStartIndex-1 : (studentPensionRolloverCHVFlagsStartIndex-1)+studentPensionRolloverCHVFlagsLength]), // Field # 399

		StudentAdjustedGrossIncomeCHVFlags: strings.TrimSpace(s[studentAdjustedGrossIncomeCHVFlagsStartIndex-1 : (studentAdjustedGrossIncomeCHVFlagsStartIndex-1)+studentAdjustedGrossIncomeCHVFlagsLength]), // Field # 400

		StudentIncomeTaxPaidCHVFlags: strings.TrimSpace(s[studentIncomeTaxPaidCHVFlagsStartIndex-1 : (studentIncomeTaxPaidCHVFlagsStartIndex-1)+studentIncomeTaxPaidCHVFlagsLength]), // Field # 401

		StudentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlags: strings.TrimSpace(s[studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex-1 : (studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex-1)+studentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsLength]), // Field # 402

		StudentDeductiblePaymentsToIRAKeoghOtherCHVFlags: strings.TrimSpace(s[studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex-1 : (studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex-1)+studentDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength]), // Field # 403

		StudentEducationCreditsCHVFlags: strings.TrimSpace(s[studentEducationCreditsCHVFlagsStartIndex-1 : (studentEducationCreditsCHVFlagsStartIndex-1)+studentEducationCreditsCHVFlagsLength]), // Field # 404

		StudentFiledScheduleABDEFHCHVFlags: strings.TrimSpace(s[studentFiledScheduleABDEFHCHVFlagsStartIndex-1 : (studentFiledScheduleABDEFHCHVFlagsStartIndex-1)+studentFiledScheduleABDEFHCHVFlagsLength]), // Field # 405

		StudentScheduleCAmountCHVFlags: strings.TrimSpace(s[studentScheduleCAmountCHVFlagsStartIndex-1 : (studentScheduleCAmountCHVFlagsStartIndex-1)+studentScheduleCAmountCHVFlagsLength]), // Field # 406

		StudentCollegeGrantAndScholarshipAidCHVFlags: strings.TrimSpace(s[studentCollegeGrantAndScholarshipAidCHVFlagsStartIndex-1 : (studentCollegeGrantAndScholarshipAidCHVFlagsStartIndex-1)+studentCollegeGrantAndScholarshipAidCHVFlagsLength]), // Field # 407

		StudentForeignEarnedIncomeExclusionCHVFlags: strings.TrimSpace(s[studentForeignEarnedIncomeExclusionCHVFlagsStartIndex-1 : (studentForeignEarnedIncomeExclusionCHVFlagsStartIndex-1)+studentForeignEarnedIncomeExclusionCHVFlagsLength]), // Field # 408

		StudentChildSupportReceivedCHVFlags: strings.TrimSpace(s[studentChildSupportReceivedCHVFlagsStartIndex-1 : (studentChildSupportReceivedCHVFlagsStartIndex-1)+studentChildSupportReceivedCHVFlagsLength]), // Field # 409

		StudentNetWorthOfBusinessesAndInvestmentFarmsCHVFlags: strings.TrimSpace(s[studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex-1 : (studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex-1)+studentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsLength]), // Field # 410

		StudentNetWorthOfCurrentInvestmentsCHVFlags: strings.TrimSpace(s[studentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex-1 : (studentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex-1)+studentNetWorthOfCurrentInvestmentsCHVFlagsLength]), // Field # 411

		StudentTotalOfCashSavingsAndCheckingCHVFlags: strings.TrimSpace(s[studentTotalOfCashSavingsAndCheckingCHVFlagsStartIndex-1 : (studentTotalOfCashSavingsAndCheckingCHVFlagsStartIndex-1)+studentTotalOfCashSavingsAndCheckingCHVFlagsLength]), // Field # 412

		StudentCollege1CHVFlags: strings.TrimSpace(s[studentCollege1CHVFlagsStartIndex-1 : (studentCollege1CHVFlagsStartIndex-1)+studentCollege1CHVFlagsLength]), // Field # 413

		StudentCollege2CHVFlags: strings.TrimSpace(s[studentCollege2CHVFlagsStartIndex-1 : (studentCollege2CHVFlagsStartIndex-1)+studentCollege2CHVFlagsLength]), // Field # 414

		StudentCollege3CHVFlags: strings.TrimSpace(s[studentCollege3CHVFlagsStartIndex-1 : (studentCollege3CHVFlagsStartIndex-1)+studentCollege3CHVFlagsLength]), // Field # 415

		StudentCollege4CHVFlags: strings.TrimSpace(s[studentCollege4CHVFlagsStartIndex-1 : (studentCollege4CHVFlagsStartIndex-1)+studentCollege4CHVFlagsLength]), // Field # 416

		StudentCollege5CHVFlags: strings.TrimSpace(s[studentCollege5CHVFlagsStartIndex-1 : (studentCollege5CHVFlagsStartIndex-1)+studentCollege5CHVFlagsLength]), // Field # 417

		StudentCollege6CHVFlags: strings.TrimSpace(s[studentCollege6CHVFlagsStartIndex-1 : (studentCollege6CHVFlagsStartIndex-1)+studentCollege6CHVFlagsLength]), // Field # 418

		StudentCollege7CHVFlags: strings.TrimSpace(s[studentCollege7CHVFlagsStartIndex-1 : (studentCollege7CHVFlagsStartIndex-1)+studentCollege7CHVFlagsLength]), // Field # 419

		StudentCollege8CHVFlags: strings.TrimSpace(s[studentCollege8CHVFlagsStartIndex-1 : (studentCollege8CHVFlagsStartIndex-1)+studentCollege8CHVFlagsLength]), // Field # 420

		StudentCollege9CHVFlags: strings.TrimSpace(s[studentCollege9CHVFlagsStartIndex-1 : (studentCollege9CHVFlagsStartIndex-1)+studentCollege9CHVFlagsLength]), // Field # 421

		StudentCollege10CHVFlags: strings.TrimSpace(s[studentCollege10CHVFlagsStartIndex-1 : (studentCollege10CHVFlagsStartIndex-1)+studentCollege10CHVFlagsLength]), // Field # 422

		StudentCollege11CHVFlags: strings.TrimSpace(s[studentCollege11CHVFlagsStartIndex-1 : (studentCollege11CHVFlagsStartIndex-1)+studentCollege11CHVFlagsLength]), // Field # 423

		StudentCollege12CHVFlags: strings.TrimSpace(s[studentCollege12CHVFlagsStartIndex-1 : (studentCollege12CHVFlagsStartIndex-1)+studentCollege12CHVFlagsLength]), // Field # 424

		StudentCollege13CHVFlags: strings.TrimSpace(s[studentCollege13CHVFlagsStartIndex-1 : (studentCollege13CHVFlagsStartIndex-1)+studentCollege13CHVFlagsLength]), // Field # 425

		StudentCollege14CHVFlags: strings.TrimSpace(s[studentCollege14CHVFlagsStartIndex-1 : (studentCollege14CHVFlagsStartIndex-1)+studentCollege14CHVFlagsLength]), // Field # 426

		StudentCollege15CHVFlags: strings.TrimSpace(s[studentCollege15CHVFlagsStartIndex-1 : (studentCollege15CHVFlagsStartIndex-1)+studentCollege15CHVFlagsLength]), // Field # 427

		StudentCollege16CHVFlags: strings.TrimSpace(s[studentCollege16CHVFlagsStartIndex-1 : (studentCollege16CHVFlagsStartIndex-1)+studentCollege16CHVFlagsLength]), // Field # 428

		StudentCollege17CHVFlags: strings.TrimSpace(s[studentCollege17CHVFlagsStartIndex-1 : (studentCollege17CHVFlagsStartIndex-1)+studentCollege17CHVFlagsLength]), // Field # 429

		StudentCollege18CHVFlags: strings.TrimSpace(s[studentCollege18CHVFlagsStartIndex-1 : (studentCollege18CHVFlagsStartIndex-1)+studentCollege18CHVFlagsLength]), // Field # 430

		StudentCollege19CHVFlags: strings.TrimSpace(s[studentCollege19CHVFlagsStartIndex-1 : (studentCollege19CHVFlagsStartIndex-1)+studentCollege19CHVFlagsLength]), // Field # 431

		StudentCollege20CHVFlags: strings.TrimSpace(s[studentCollege20CHVFlagsStartIndex-1 : (studentCollege20CHVFlagsStartIndex-1)+studentCollege20CHVFlagsLength]), // Field # 432

		StudentConsentToRetrieveAndDiscloseFTICHVFlags: strings.TrimSpace(s[studentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex-1 : (studentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex-1)+studentConsentToRetrieveAndDiscloseFTICHVFlagsLength]), // Field # 433

		StudentSignatureCHVFlags: strings.TrimSpace(s[studentSignatureCHVFlagsStartIndex-1 : (studentSignatureCHVFlagsStartIndex-1)+studentSignatureCHVFlagsLength]), // Field # 434

		StudentSignatureDateCHVFlags: strings.TrimSpace(s[studentSignatureDateCHVFlagsStartIndex-1 : (studentSignatureDateCHVFlagsStartIndex-1)+studentSignatureDateCHVFlagsLength]), // Field # 435

		StudentSpouseFirstNameCHVFlags: strings.TrimSpace(s[studentSpouseFirstNameCHVFlagsStartIndex-1 : (studentSpouseFirstNameCHVFlagsStartIndex-1)+studentSpouseFirstNameCHVFlagsLength]), // Field # 436

		StudentSpouseMiddleNameCHVFlags: strings.TrimSpace(s[studentSpouseMiddleNameCHVFlagsStartIndex-1 : (studentSpouseMiddleNameCHVFlagsStartIndex-1)+studentSpouseMiddleNameCHVFlagsLength]), // Field # 437

		StudentSpouseLastNameCHVFlags: strings.TrimSpace(s[studentSpouseLastNameCHVFlagsStartIndex-1 : (studentSpouseLastNameCHVFlagsStartIndex-1)+studentSpouseLastNameCHVFlagsLength]), // Field # 438

		StudentSpouseSuffixCHVFlags: strings.TrimSpace(s[studentSpouseSuffixCHVFlagsStartIndex-1 : (studentSpouseSuffixCHVFlagsStartIndex-1)+studentSpouseSuffixCHVFlagsLength]), // Field # 439

		StudentSpouseDateOfBirthCHVFlags: strings.TrimSpace(s[studentSpouseDateOfBirthCHVFlagsStartIndex-1 : (studentSpouseDateOfBirthCHVFlagsStartIndex-1)+studentSpouseDateOfBirthCHVFlagsLength]), // Field # 440

		StudentSpouseSSNCHVFlags: strings.TrimSpace(s[studentSpouseSSNCHVFlagsStartIndex-1 : (studentSpouseSSNCHVFlagsStartIndex-1)+studentSpouseSSNCHVFlagsLength]), // Field # 441

		StudentSpouseITINCHVFlags: strings.TrimSpace(s[studentSpouseITINCHVFlagsStartIndex-1 : (studentSpouseITINCHVFlagsStartIndex-1)+studentSpouseITINCHVFlagsLength]), // Field # 442

		StudentSpousePhoneNumberCHVFlags: strings.TrimSpace(s[studentSpousePhoneNumberCHVFlagsStartIndex-1 : (studentSpousePhoneNumberCHVFlagsStartIndex-1)+studentSpousePhoneNumberCHVFlagsLength]), // Field # 443

		StudentSpouseEmailAddressCHVFlags: strings.TrimSpace(s[studentSpouseEmailAddressCHVFlagsStartIndex-1 : (studentSpouseEmailAddressCHVFlagsStartIndex-1)+studentSpouseEmailAddressCHVFlagsLength]), // Field # 444

		StudentSpouseStreetAddressCHVFlags: strings.TrimSpace(s[studentSpouseStreetAddressCHVFlagsStartIndex-1 : (studentSpouseStreetAddressCHVFlagsStartIndex-1)+studentSpouseStreetAddressCHVFlagsLength]), // Field # 445

		StudentSpouseCityCHVFlags: strings.TrimSpace(s[studentSpouseCityCHVFlagsStartIndex-1 : (studentSpouseCityCHVFlagsStartIndex-1)+studentSpouseCityCHVFlagsLength]), // Field # 446

		StudentSpouseStateCHVFlags: strings.TrimSpace(s[studentSpouseStateCHVFlagsStartIndex-1 : (studentSpouseStateCHVFlagsStartIndex-1)+studentSpouseStateCHVFlagsLength]), // Field # 447

		StudentSpouseZipCodeCHVFlags: strings.TrimSpace(s[studentSpouseZipCodeCHVFlagsStartIndex-1 : (studentSpouseZipCodeCHVFlagsStartIndex-1)+studentSpouseZipCodeCHVFlagsLength]), // Field # 448

		StudentSpouseCountryCHVFlags: strings.TrimSpace(s[studentSpouseCountryCHVFlagsStartIndex-1 : (studentSpouseCountryCHVFlagsStartIndex-1)+studentSpouseCountryCHVFlagsLength]), // Field # 449

		StudentSpouseFiled1040Or1040NRCHVFlags: strings.TrimSpace(s[studentSpouseFiled1040Or1040NRCHVFlagsStartIndex-1 : (studentSpouseFiled1040Or1040NRCHVFlagsStartIndex-1)+studentSpouseFiled1040Or1040NRCHVFlagsLength]), // Field # 450

		StudentSpouseFiledNonUSTaxReturnCHVFlags: strings.TrimSpace(s[studentSpouseFiledNonUSTaxReturnCHVFlagsStartIndex-1 : (studentSpouseFiledNonUSTaxReturnCHVFlagsStartIndex-1)+studentSpouseFiledNonUSTaxReturnCHVFlagsLength]), // Field # 451

		StudentSpouseTaxReturnFilingStatusCHVFlags: strings.TrimSpace(s[studentSpouseTaxReturnFilingStatusCHVFlagsStartIndex-1 : (studentSpouseTaxReturnFilingStatusCHVFlagsStartIndex-1)+studentSpouseTaxReturnFilingStatusCHVFlagsLength]), // Field # 452

		StudentSpouseIncomeEarnedFromWorkCHVFlags: strings.TrimSpace(s[studentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex-1 : (studentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex-1)+studentSpouseIncomeEarnedFromWorkCHVFlagsLength]), // Field # 453

		StudentSpouseTaxExemptInterestIncomeCHVFlags: strings.TrimSpace(s[studentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex-1 : (studentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex-1)+studentSpouseTaxExemptInterestIncomeCHVFlagsLength]), // Field # 454

		StudentSpouseUntaxedPortionsOfIRADistributionsCHVFlags: strings.TrimSpace(s[studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex-1 : (studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex-1)+studentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsLength]), // Field # 455

		StudentSpouseIRARolloverCHVFlags: strings.TrimSpace(s[studentSpouseIRARolloverCHVFlagsStartIndex-1 : (studentSpouseIRARolloverCHVFlagsStartIndex-1)+studentSpouseIRARolloverCHVFlagsLength]), // Field # 456

		StudentSpouseUntaxedPortionsOfPensionsCHVFlags: strings.TrimSpace(s[studentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex-1 : (studentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex-1)+studentSpouseUntaxedPortionsOfPensionsCHVFlagsLength]), // Field # 457

		StudentSpousePensionRolloverCHVFlags: strings.TrimSpace(s[studentSpousePensionRolloverCHVFlagsStartIndex-1 : (studentSpousePensionRolloverCHVFlagsStartIndex-1)+studentSpousePensionRolloverCHVFlagsLength]), // Field # 458

		StudentSpouseAdjustedGrossIncomeCHVFlags: strings.TrimSpace(s[studentSpouseAdjustedGrossIncomeCHVFlagsStartIndex-1 : (studentSpouseAdjustedGrossIncomeCHVFlagsStartIndex-1)+studentSpouseAdjustedGrossIncomeCHVFlagsLength]), // Field # 459

		StudentSpouseIncomeTaxPaidCHVFlags: strings.TrimSpace(s[studentSpouseIncomeTaxPaidCHVFlagsStartIndex-1 : (studentSpouseIncomeTaxPaidCHVFlagsStartIndex-1)+studentSpouseIncomeTaxPaidCHVFlagsLength]), // Field # 460

		StudentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlags: strings.TrimSpace(s[studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex-1 : (studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex-1)+studentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength]), // Field # 461

		StudentSpouseEducationCreditsCHVFlags: strings.TrimSpace(s[studentSpouseEducationCreditsCHVFlagsStartIndex-1 : (studentSpouseEducationCreditsCHVFlagsStartIndex-1)+studentSpouseEducationCreditsCHVFlagsLength]), // Field # 462

		StudentSpouseFiledScheduleABDEFHCHVFlags: strings.TrimSpace(s[studentSpouseFiledScheduleABDEFHCHVFlagsStartIndex-1 : (studentSpouseFiledScheduleABDEFHCHVFlagsStartIndex-1)+studentSpouseFiledScheduleABDEFHCHVFlagsLength]), // Field # 463

		StudentSpouseScheduleCAmountCHVFlags: strings.TrimSpace(s[studentSpouseScheduleCAmountCHVFlagsStartIndex-1 : (studentSpouseScheduleCAmountCHVFlagsStartIndex-1)+studentSpouseScheduleCAmountCHVFlagsLength]), // Field # 464

		StudentSpouseForeignEarnedIncomeExclusionCHVFlags: strings.TrimSpace(s[studentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex-1 : (studentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex-1)+studentSpouseForeignEarnedIncomeExclusionCHVFlagsLength]), // Field # 465

		StudentSpouseConsentToRetrieveAndDiscloseFTICHVFlags: strings.TrimSpace(s[studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex-1 : (studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex-1)+studentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsLength]), // Field # 466

		StudentSpouseSignatureCHVFlags: strings.TrimSpace(s[studentSpouseSignatureCHVFlagsStartIndex-1 : (studentSpouseSignatureCHVFlagsStartIndex-1)+studentSpouseSignatureCHVFlagsLength]), // Field # 467

		StudentSpouseSignatureDateCHVFlags: strings.TrimSpace(s[studentSpouseSignatureDateCHVFlagsStartIndex-1 : (studentSpouseSignatureDateCHVFlagsStartIndex-1)+studentSpouseSignatureDateCHVFlagsLength]), // Field # 468

		ParentFirstNameCHVFlags: strings.TrimSpace(s[parentFirstNameCHVFlagsStartIndex-1 : (parentFirstNameCHVFlagsStartIndex-1)+parentFirstNameCHVFlagsLength]), // Field # 469

		ParentMiddleNameCHVFlags: strings.TrimSpace(s[parentMiddleNameCHVFlagsStartIndex-1 : (parentMiddleNameCHVFlagsStartIndex-1)+parentMiddleNameCHVFlagsLength]), // Field # 470

		ParentLastNameCHVFlags: strings.TrimSpace(s[parentLastNameCHVFlagsStartIndex-1 : (parentLastNameCHVFlagsStartIndex-1)+parentLastNameCHVFlagsLength]), // Field # 471

		ParentSuffixCHVFlags: strings.TrimSpace(s[parentSuffixCHVFlagsStartIndex-1 : (parentSuffixCHVFlagsStartIndex-1)+parentSuffixCHVFlagsLength]), // Field # 472

		ParentDateOfBirthCHVFlags: strings.TrimSpace(s[parentDateOfBirthCHVFlagsStartIndex-1 : (parentDateOfBirthCHVFlagsStartIndex-1)+parentDateOfBirthCHVFlagsLength]), // Field # 473

		ParentSSNCHVFlags: strings.TrimSpace(s[parentSSNCHVFlagsStartIndex-1 : (parentSSNCHVFlagsStartIndex-1)+parentSSNCHVFlagsLength]), // Field # 474

		ParentITINCHVFlags: strings.TrimSpace(s[parentITINCHVFlagsStartIndex-1 : (parentITINCHVFlagsStartIndex-1)+parentITINCHVFlagsLength]), // Field # 475

		ParentPhoneNumberCHVFlags: strings.TrimSpace(s[parentPhoneNumberCHVFlagsStartIndex-1 : (parentPhoneNumberCHVFlagsStartIndex-1)+parentPhoneNumberCHVFlagsLength]), // Field # 476

		ParentEmailAddressCHVFlags: strings.TrimSpace(s[parentEmailAddressCHVFlagsStartIndex-1 : (parentEmailAddressCHVFlagsStartIndex-1)+parentEmailAddressCHVFlagsLength]), // Field # 477

		ParentStreetAddressCHVFlags: strings.TrimSpace(s[parentStreetAddressCHVFlagsStartIndex-1 : (parentStreetAddressCHVFlagsStartIndex-1)+parentStreetAddressCHVFlagsLength]), // Field # 478

		ParentCityCHVFlags: strings.TrimSpace(s[parentCityCHVFlagsStartIndex-1 : (parentCityCHVFlagsStartIndex-1)+parentCityCHVFlagsLength]), // Field # 479

		ParentStateCHVFlags: strings.TrimSpace(s[parentStateCHVFlagsStartIndex-1 : (parentStateCHVFlagsStartIndex-1)+parentStateCHVFlagsLength]), // Field # 480

		ParentZipCodeCHVFlags: strings.TrimSpace(s[parentZipCodeCHVFlagsStartIndex-1 : (parentZipCodeCHVFlagsStartIndex-1)+parentZipCodeCHVFlagsLength]), // Field # 481

		ParentCountryCHVFlags: strings.TrimSpace(s[parentCountryCHVFlagsStartIndex-1 : (parentCountryCHVFlagsStartIndex-1)+parentCountryCHVFlagsLength]), // Field # 482

		ParentMaritalStatusCHVFlags: strings.TrimSpace(s[parentMaritalStatusCHVFlagsStartIndex-1 : (parentMaritalStatusCHVFlagsStartIndex-1)+parentMaritalStatusCHVFlagsLength]), // Field # 483

		ParentStateOfLegalResidenceCHVFlags: strings.TrimSpace(s[parentStateOfLegalResidenceCHVFlagsStartIndex-1 : (parentStateOfLegalResidenceCHVFlagsStartIndex-1)+parentStateOfLegalResidenceCHVFlagsLength]), // Field # 484

		ParentLegalResidenceDateCHVFlags: strings.TrimSpace(s[parentLegalResidenceDateCHVFlagsStartIndex-1 : (parentLegalResidenceDateCHVFlagsStartIndex-1)+parentLegalResidenceDateCHVFlagsLength]), // Field # 485

		ParentUpdatedFamilySizeCHVFlags: strings.TrimSpace(s[parentUpdatedFamilySizeCHVFlagsStartIndex-1 : (parentUpdatedFamilySizeCHVFlagsStartIndex-1)+parentUpdatedFamilySizeCHVFlagsLength]), // Field # 486

		ParentNumberInCollegeCHVFlags: strings.TrimSpace(s[parentNumberInCollegeCHVFlagsStartIndex-1 : (parentNumberInCollegeCHVFlagsStartIndex-1)+parentNumberInCollegeCHVFlagsLength]), // Field # 487

		ParentReceivedEITCCHVFlags: strings.TrimSpace(s[parentReceivedEITCCHVFlagsStartIndex-1 : (parentReceivedEITCCHVFlagsStartIndex-1)+parentReceivedEITCCHVFlagsLength]), // Field # 488

		ParentReceivedFederalHousingAssistanceCHVFlags: strings.TrimSpace(s[parentReceivedFederalHousingAssistanceCHVFlagsStartIndex-1 : (parentReceivedFederalHousingAssistanceCHVFlagsStartIndex-1)+parentReceivedFederalHousingAssistanceCHVFlagsLength]), // Field # 489

		ParentReceivedFreeReducedPriceLunchCHVFlags: strings.TrimSpace(s[parentReceivedFreeReducedPriceLunchCHVFlagsStartIndex-1 : (parentReceivedFreeReducedPriceLunchCHVFlagsStartIndex-1)+parentReceivedFreeReducedPriceLunchCHVFlagsLength]), // Field # 490

		ParentReceivedMedicaidCHVFlags: strings.TrimSpace(s[parentReceivedMedicaidCHVFlagsStartIndex-1 : (parentReceivedMedicaidCHVFlagsStartIndex-1)+parentReceivedMedicaidCHVFlagsLength]), // Field # 491

		ParentReceivedRefundableCreditFor36BHealthPlanCHVFlags: strings.TrimSpace(s[parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex-1 : (parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsStartIndex-1)+parentReceivedRefundableCreditFor36BHealthPlanCHVFlagsLength]), // Field # 492

		ParentReceivedSNAPCHVFlags: strings.TrimSpace(s[parentReceivedSNAPCHVFlagsStartIndex-1 : (parentReceivedSNAPCHVFlagsStartIndex-1)+parentReceivedSNAPCHVFlagsLength]), // Field # 493

		ParentReceivedSupplementalSecurityIncomeCHVFlags: strings.TrimSpace(s[parentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex-1 : (parentReceivedSupplementalSecurityIncomeCHVFlagsStartIndex-1)+parentReceivedSupplementalSecurityIncomeCHVFlagsLength]), // Field # 494

		ParentReceivedTANFCHVFlags: strings.TrimSpace(s[parentReceivedTANFCHVFlagsStartIndex-1 : (parentReceivedTANFCHVFlagsStartIndex-1)+parentReceivedTANFCHVFlagsLength]), // Field # 495

		ParentReceivedWICCHVFlags: strings.TrimSpace(s[parentReceivedWICCHVFlagsStartIndex-1 : (parentReceivedWICCHVFlagsStartIndex-1)+parentReceivedWICCHVFlagsLength]), // Field # 496

		ParentFederalBenefitsNoneOfTheAboveCHVFlags: strings.TrimSpace(s[parentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex-1 : (parentFederalBenefitsNoneOfTheAboveCHVFlagsStartIndex-1)+parentFederalBenefitsNoneOfTheAboveCHVFlagsLength]), // Field # 497

		ParentFiled1040Or1040NRCHVFlags: strings.TrimSpace(s[parentFiled1040Or1040NRCHVFlagsStartIndex-1 : (parentFiled1040Or1040NRCHVFlagsStartIndex-1)+parentFiled1040Or1040NRCHVFlagsLength]), // Field # 498

		ParentFileNonUSTaxReturnCHVFlags: strings.TrimSpace(s[parentFileNonUSTaxReturnCHVFlagsStartIndex-1 : (parentFileNonUSTaxReturnCHVFlagsStartIndex-1)+parentFileNonUSTaxReturnCHVFlagsLength]), // Field # 499

		ParentFiledJointReturnWithCurrentSpouseCHVFlags: strings.TrimSpace(s[parentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex-1 : (parentFiledJointReturnWithCurrentSpouseCHVFlagsStartIndex-1)+parentFiledJointReturnWithCurrentSpouseCHVFlagsLength]), // Field # 500

		ParentTaxReturnFilingStatusCHVFlags: strings.TrimSpace(s[parentTaxReturnFilingStatusCHVFlagsStartIndex-1 : (parentTaxReturnFilingStatusCHVFlagsStartIndex-1)+parentTaxReturnFilingStatusCHVFlagsLength]), // Field # 501

		ParentIncomeEarnedFromWorkCHVFlags: strings.TrimSpace(s[parentIncomeEarnedFromWorkCHVFlagsStartIndex-1 : (parentIncomeEarnedFromWorkCHVFlagsStartIndex-1)+parentIncomeEarnedFromWorkCHVFlagsLength]), // Field # 502

		ParentTaxExemptInterestIncomeCHVFlags: strings.TrimSpace(s[parentTaxExemptInterestIncomeCHVFlagsStartIndex-1 : (parentTaxExemptInterestIncomeCHVFlagsStartIndex-1)+parentTaxExemptInterestIncomeCHVFlagsLength]), // Field # 503

		ParentUntaxedPortionsOfIRADistributionsCHVFlags: strings.TrimSpace(s[parentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex-1 : (parentUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex-1)+parentUntaxedPortionsOfIRADistributionsCHVFlagsLength]), // Field # 504

		ParentIRARolloverCHVFlags: strings.TrimSpace(s[parentIRARolloverCHVFlagsStartIndex-1 : (parentIRARolloverCHVFlagsStartIndex-1)+parentIRARolloverCHVFlagsLength]), // Field # 505

		ParentUntaxedPortionsOfPensionsCHVFlags: strings.TrimSpace(s[parentUntaxedPortionsOfPensionsCHVFlagsStartIndex-1 : (parentUntaxedPortionsOfPensionsCHVFlagsStartIndex-1)+parentUntaxedPortionsOfPensionsCHVFlagsLength]), // Field # 506

		ParentPensionRolloverCHVFlags: strings.TrimSpace(s[parentPensionRolloverCHVFlagsStartIndex-1 : (parentPensionRolloverCHVFlagsStartIndex-1)+parentPensionRolloverCHVFlagsLength]), // Field # 507

		ParentAdjustedGrossIncomeCHVFlags: strings.TrimSpace(s[parentAdjustedGrossIncomeCHVFlagsStartIndex-1 : (parentAdjustedGrossIncomeCHVFlagsStartIndex-1)+parentAdjustedGrossIncomeCHVFlagsLength]), // Field # 508

		ParentIncomeTaxPaidCHVFlags: strings.TrimSpace(s[parentIncomeTaxPaidCHVFlagsStartIndex-1 : (parentIncomeTaxPaidCHVFlagsStartIndex-1)+parentIncomeTaxPaidCHVFlagsLength]), // Field # 509

		ParentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlags: strings.TrimSpace(s[parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex-1 : (parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsStartIndex-1)+parentEarnedIncomeTaxCreditReceivedDuringTaxYearCHVFlagsLength]), // Field # 510

		ParentDeductiblePaymentsToIRAKeoghOtherCHVFlags: strings.TrimSpace(s[parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex-1 : (parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex-1)+parentDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength]), // Field # 511

		ParentEducationCreditsCHVFlags: strings.TrimSpace(s[parentEducationCreditsCHVFlagsStartIndex-1 : (parentEducationCreditsCHVFlagsStartIndex-1)+parentEducationCreditsCHVFlagsLength]), // Field # 512

		ParentFiledScheduleABDEFHCHVFlags: strings.TrimSpace(s[parentFiledScheduleABDEFHCHVFlagsStartIndex-1 : (parentFiledScheduleABDEFHCHVFlagsStartIndex-1)+parentFiledScheduleABDEFHCHVFlagsLength]), // Field # 513

		ParentScheduleCAmountCHVFlags: strings.TrimSpace(s[parentScheduleCAmountCHVFlagsStartIndex-1 : (parentScheduleCAmountCHVFlagsStartIndex-1)+parentScheduleCAmountCHVFlagsLength]), // Field # 514

		ParentCollegeGrantAndScholarshipAidCHVFlags: strings.TrimSpace(s[parentCollegeGrantAndScholarshipAidCHVFlagsStartIndex-1 : (parentCollegeGrantAndScholarshipAidCHVFlagsStartIndex-1)+parentCollegeGrantAndScholarshipAidCHVFlagsLength]), // Field # 515

		ParentForeignEarnedIncomeExclusionCHVFlags: strings.TrimSpace(s[parentForeignEarnedIncomeExclusionCHVFlagsStartIndex-1 : (parentForeignEarnedIncomeExclusionCHVFlagsStartIndex-1)+parentForeignEarnedIncomeExclusionCHVFlagsLength]), // Field # 516

		ParentChildSupportReceivedCHVFlags: strings.TrimSpace(s[parentChildSupportReceivedCHVFlagsStartIndex-1 : (parentChildSupportReceivedCHVFlagsStartIndex-1)+parentChildSupportReceivedCHVFlagsLength]), // Field # 517

		ParentNetWorthOfCurrentInvestmentsCHVFlags: strings.TrimSpace(s[parentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex-1 : (parentNetWorthOfCurrentInvestmentsCHVFlagsStartIndex-1)+parentNetWorthOfCurrentInvestmentsCHVFlagsLength]), // Field # 518

		ParentTotalOfCashSavingsAndCheckingAccountsCHVFlags: strings.TrimSpace(s[parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsStartIndex-1 : (parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsStartIndex-1)+parentTotalOfCashSavingsAndCheckingAccountsCHVFlagsLength]), // Field # 519

		ParentNetWorthOfBusinessesAndInvestmentFarmsCHVFlags: strings.TrimSpace(s[parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex-1 : (parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsStartIndex-1)+parentNetWorthOfBusinessesAndInvestmentFarmsCHVFlagsLength]), // Field # 520

		ParentConsentToRetrieveAndDiscloseFTICHVFlags: strings.TrimSpace(s[parentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex-1 : (parentConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex-1)+parentConsentToRetrieveAndDiscloseFTICHVFlagsLength]), // Field # 521

		ParentSignatureCHVFlags: strings.TrimSpace(s[parentSignatureCHVFlagsStartIndex-1 : (parentSignatureCHVFlagsStartIndex-1)+parentSignatureCHVFlagsLength]), // Field # 522

		ParentSignatureDateCHVFlags: strings.TrimSpace(s[parentSignatureDateCHVFlagsStartIndex-1 : (parentSignatureDateCHVFlagsStartIndex-1)+parentSignatureDateCHVFlagsLength]), // Field # 523

		ParentSpouseFirstNameCHVFlags: strings.TrimSpace(s[parentSpouseFirstNameCHVFlagsStartIndex-1 : (parentSpouseFirstNameCHVFlagsStartIndex-1)+parentSpouseFirstNameCHVFlagsLength]), // Field # 524

		ParentSpouseMiddleNameCHVFlags: strings.TrimSpace(s[parentSpouseMiddleNameCHVFlagsStartIndex-1 : (parentSpouseMiddleNameCHVFlagsStartIndex-1)+parentSpouseMiddleNameCHVFlagsLength]), // Field # 525

		ParentSpouseLastNameCHVFlags: strings.TrimSpace(s[parentSpouseLastNameCHVFlagsStartIndex-1 : (parentSpouseLastNameCHVFlagsStartIndex-1)+parentSpouseLastNameCHVFlagsLength]), // Field # 526

		ParentSpouseSuffixCHVFlags: strings.TrimSpace(s[parentSpouseSuffixCHVFlagsStartIndex-1 : (parentSpouseSuffixCHVFlagsStartIndex-1)+parentSpouseSuffixCHVFlagsLength]), // Field # 527

		ParentSpouseDateOfBirthCHVFlags: strings.TrimSpace(s[parentSpouseDateOfBirthCHVFlagsStartIndex-1 : (parentSpouseDateOfBirthCHVFlagsStartIndex-1)+parentSpouseDateOfBirthCHVFlagsLength]), // Field # 528

		ParentSpouseSSNCHVFlags: strings.TrimSpace(s[parentSpouseSSNCHVFlagsStartIndex-1 : (parentSpouseSSNCHVFlagsStartIndex-1)+parentSpouseSSNCHVFlagsLength]), // Field # 529

		ParentSpouseITINCHVFlags: strings.TrimSpace(s[parentSpouseITINCHVFlagsStartIndex-1 : (parentSpouseITINCHVFlagsStartIndex-1)+parentSpouseITINCHVFlagsLength]), // Field # 530

		ParentSpousePhoneNumberCHVFlags: strings.TrimSpace(s[parentSpousePhoneNumberCHVFlagsStartIndex-1 : (parentSpousePhoneNumberCHVFlagsStartIndex-1)+parentSpousePhoneNumberCHVFlagsLength]), // Field # 531

		ParentSpouseEmailAddressCHVFlags: strings.TrimSpace(s[parentSpouseEmailAddressCHVFlagsStartIndex-1 : (parentSpouseEmailAddressCHVFlagsStartIndex-1)+parentSpouseEmailAddressCHVFlagsLength]), // Field # 532

		ParentSpouseStreetAddressCHVFlags: strings.TrimSpace(s[parentSpouseStreetAddressCHVFlagsStartIndex-1 : (parentSpouseStreetAddressCHVFlagsStartIndex-1)+parentSpouseStreetAddressCHVFlagsLength]), // Field # 533

		ParentSpouseCityCHVFlags: strings.TrimSpace(s[parentSpouseCityCHVFlagsStartIndex-1 : (parentSpouseCityCHVFlagsStartIndex-1)+parentSpouseCityCHVFlagsLength]), // Field # 534

		ParentSpouseStateCHVFlags: strings.TrimSpace(s[parentSpouseStateCHVFlagsStartIndex-1 : (parentSpouseStateCHVFlagsStartIndex-1)+parentSpouseStateCHVFlagsLength]), // Field # 535

		ParentSpouseZipCodeCHVFlags: strings.TrimSpace(s[parentSpouseZipCodeCHVFlagsStartIndex-1 : (parentSpouseZipCodeCHVFlagsStartIndex-1)+parentSpouseZipCodeCHVFlagsLength]), // Field # 536

		ParentSpouseCountryCHVFlags: strings.TrimSpace(s[parentSpouseCountryCHVFlagsStartIndex-1 : (parentSpouseCountryCHVFlagsStartIndex-1)+parentSpouseCountryCHVFlagsLength]), // Field # 537

		ParentSpouseFiled1040Or1040NRCHVFlags: strings.TrimSpace(s[parentSpouseFiled1040Or1040NRCHVFlagsStartIndex-1 : (parentSpouseFiled1040Or1040NRCHVFlagsStartIndex-1)+parentSpouseFiled1040Or1040NRCHVFlagsLength]), // Field # 538

		ParentSpouseFileNonUSTaxReturnCHVFlags: strings.TrimSpace(s[parentSpouseFileNonUSTaxReturnCHVFlagsStartIndex-1 : (parentSpouseFileNonUSTaxReturnCHVFlagsStartIndex-1)+parentSpouseFileNonUSTaxReturnCHVFlagsLength]), // Field # 539

		ParentSpouseTaxReturnFilingStatusCHVFlags: strings.TrimSpace(s[parentSpouseTaxReturnFilingStatusCHVFlagsStartIndex-1 : (parentSpouseTaxReturnFilingStatusCHVFlagsStartIndex-1)+parentSpouseTaxReturnFilingStatusCHVFlagsLength]), // Field # 540

		ParentSpouseIncomeEarnedFromWorkCHVFlags: strings.TrimSpace(s[parentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex-1 : (parentSpouseIncomeEarnedFromWorkCHVFlagsStartIndex-1)+parentSpouseIncomeEarnedFromWorkCHVFlagsLength]), // Field # 541

		ParentSpouseTaxExemptInterestIncomeCHVFlags: strings.TrimSpace(s[parentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex-1 : (parentSpouseTaxExemptInterestIncomeCHVFlagsStartIndex-1)+parentSpouseTaxExemptInterestIncomeCHVFlagsLength]), // Field # 542

		ParentSpouseUntaxedPortionsOfIRADistributionsCHVFlags: strings.TrimSpace(s[parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex-1 : (parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsStartIndex-1)+parentSpouseUntaxedPortionsOfIRADistributionsCHVFlagsLength]), // Field # 543

		ParentSpouseIRARolloverCHVFlags: strings.TrimSpace(s[parentSpouseIRARolloverCHVFlagsStartIndex-1 : (parentSpouseIRARolloverCHVFlagsStartIndex-1)+parentSpouseIRARolloverCHVFlagsLength]), // Field # 544

		ParentSpouseUntaxedPortionsOfPensionsCHVFlags: strings.TrimSpace(s[parentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex-1 : (parentSpouseUntaxedPortionsOfPensionsCHVFlagsStartIndex-1)+parentSpouseUntaxedPortionsOfPensionsCHVFlagsLength]), // Field # 545

		ParentSpousePensionRolloverCHVFlags: strings.TrimSpace(s[parentSpousePensionRolloverCHVFlagsStartIndex-1 : (parentSpousePensionRolloverCHVFlagsStartIndex-1)+parentSpousePensionRolloverCHVFlagsLength]), // Field # 546

		ParentSpouseAdjustedGrossIncomeCHVFlags: strings.TrimSpace(s[parentSpouseAdjustedGrossIncomeCHVFlagsStartIndex-1 : (parentSpouseAdjustedGrossIncomeCHVFlagsStartIndex-1)+parentSpouseAdjustedGrossIncomeCHVFlagsLength]), // Field # 547

		ParentSpouseIncomeTaxPaidCHVFlags: strings.TrimSpace(s[parentSpouseIncomeTaxPaidCHVFlagsStartIndex-1 : (parentSpouseIncomeTaxPaidCHVFlagsStartIndex-1)+parentSpouseIncomeTaxPaidCHVFlagsLength]), // Field # 548

		ParentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlags: strings.TrimSpace(s[parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex-1 : (parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsStartIndex-1)+parentSpouseDeductiblePaymentsToIRAKeoghOtherCHVFlagsLength]), // Field # 549

		ParentSpouseEducationCreditsCHVFlags: strings.TrimSpace(s[parentSpouseEducationCreditsCHVFlagsStartIndex-1 : (parentSpouseEducationCreditsCHVFlagsStartIndex-1)+parentSpouseEducationCreditsCHVFlagsLength]), // Field # 550

		ParentSpouseFiledScheduleABDEFHCHVFlags: strings.TrimSpace(s[parentSpouseFiledScheduleABDEFHCHVFlagsStartIndex-1 : (parentSpouseFiledScheduleABDEFHCHVFlagsStartIndex-1)+parentSpouseFiledScheduleABDEFHCHVFlagsLength]), // Field # 551

		ParentSpouseScheduleCAmountCHVFlags: strings.TrimSpace(s[parentSpouseScheduleCAmountCHVFlagsStartIndex-1 : (parentSpouseScheduleCAmountCHVFlagsStartIndex-1)+parentSpouseScheduleCAmountCHVFlagsLength]), // Field # 552

		ParentSpouseForeignEarnedIncomeExclusionCHVFlags: strings.TrimSpace(s[parentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex-1 : (parentSpouseForeignEarnedIncomeExclusionCHVFlagsStartIndex-1)+parentSpouseForeignEarnedIncomeExclusionCHVFlagsLength]), // Field # 553

		ParentSpouseConsentToRetrieveAndDiscloseFTICHVFlags: strings.TrimSpace(s[parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex-1 : (parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsStartIndex-1)+parentSpouseConsentToRetrieveAndDiscloseFTICHVFlagsLength]), // Field # 554

		ParentSpouseSignatureCHVFlags: strings.TrimSpace(s[parentSpouseSignatureCHVFlagsStartIndex-1 : (parentSpouseSignatureCHVFlagsStartIndex-1)+parentSpouseSignatureCHVFlagsLength]), // Field # 555

		ParentSpouseSignatureDateCHVFlags: strings.TrimSpace(s[parentSpouseSignatureDateCHVFlagsStartIndex-1 : (parentSpouseSignatureDateCHVFlagsStartIndex-1)+parentSpouseSignatureDateCHVFlagsLength]), // Field # 556

		DHSPrimaryMatchStatus: strings.TrimSpace(s[dHSPrimaryMatchStatusStartIndex-1 : (dHSPrimaryMatchStatusStartIndex-1)+dHSPrimaryMatchStatusLength]), // Field # 557

		DHSCaseNumber: strings.TrimSpace(s[dHSCaseNumberStartIndex-1 : (dHSCaseNumberStartIndex-1)+dHSCaseNumberLength]), // Field # 559

		NSLDSMatchStatus: strings.TrimSpace(s[nsldsMatchStatusStartIndex-1 : (nsldsMatchStatusStartIndex-1)+nsldsMatchStatusLength]), // Field # 560

		NSLDSPostscreeningReasonCode: strings.TrimSpace(s[nsldsPostscreeningReasonCodeStartIndex-1 : (nsldsPostscreeningReasonCodeStartIndex-1)+nsldsPostscreeningReasonCodeLength]), // Field # 561

		StudentSSACitizenshipFlagResults: strings.TrimSpace(s[studentSSACitizenshipFlagResultsStartIndex-1 : (studentSSACitizenshipFlagResultsStartIndex-1)+studentSSACitizenshipFlagResultsLength]), // Field # 562

		StudentSSAMatchStatus: strings.TrimSpace(s[studentSSAMatchStatusStartIndex-1 : (studentSSAMatchStatusStartIndex-1)+studentSSAMatchStatusLength]), // Field # 563

		StudentSpouseSSAMatchStatus: strings.TrimSpace(s[studentSpouseSSAMatchStatusStartIndex-1 : (studentSpouseSSAMatchStatusStartIndex-1)+studentSpouseSSAMatchStatusLength]), // Field # 564

		ParentSSAMatchStatus: strings.TrimSpace(s[parentSSAMatchStatusStartIndex-1 : (parentSSAMatchStatusStartIndex-1)+parentSSAMatchStatusLength]), // Field # 565

		ParentSpouseOrPartnerSSAMatchStatus: strings.TrimSpace(s[parentSpouseOrPartnerSSAMatchStatusStartIndex-1 : (parentSpouseOrPartnerSSAMatchStatusStartIndex-1)+parentSpouseOrPartnerSSAMatchStatusLength]), // Field # 566

		VAMatchFlag: strings.TrimSpace(s[vAMatchFlagStartIndex-1 : (vAMatchFlagStartIndex-1)+vAMatchFlagLength]), // Field # 567

		CommentCodes: strings.TrimSpace(s[commentCodesStartIndex-1 : (commentCodesStartIndex-1)+commentCodesLength]), // Field # 568

		DrugAbuseHoldIndicator: strings.TrimSpace(s[drugAbuseHoldIndicatorStartIndex-1 : (drugAbuseHoldIndicatorStartIndex-1)+drugAbuseHoldIndicatorLength]), // Field # 569

		GraduateFlag: strings.TrimSpace(s[graduateFlagStartIndex-1 : (graduateFlagStartIndex-1)+graduateFlagLength]), // Field # 570

		PellGrantEligibilityFlag: strings.TrimSpace(s[pellGrantEligibilityFlagStartIndex-1 : (pellGrantEligibilityFlagStartIndex-1)+pellGrantEligibilityFlagLength]), // Field # 571

		ReprocessedReasonCode: strings.TrimSpace(s[reprocessedReasonCodeStartIndex-1 : (reprocessedReasonCodeStartIndex-1)+reprocessedReasonCodeLength]), // Field # 572

		FPSCFlag: strings.TrimSpace(s[fpsCFlagStartIndex-1 : (fpsCFlagStartIndex-1)+fpsCFlagLength]), // Field # 573

		FPSCChangeFlag: strings.TrimSpace(s[fpsCChangeFlagStartIndex-1 : (fpsCChangeFlagStartIndex-1)+fpsCChangeFlagLength]), // Field # 574

		ElectronicFederalSchoolCodeIndicator: strings.TrimSpace(s[electronicFederalSchoolCodeIndicatorStartIndex-1 : (electronicFederalSchoolCodeIndicatorStartIndex-1)+electronicFederalSchoolCodeIndicatorLength]), // Field # 575

		RejectReasonCodes: strings.TrimSpace(s[rejectReasonCodesStartIndex-1 : (rejectReasonCodesStartIndex-1)+rejectReasonCodesLength]), // Field # 576

		ElectronicTransactionIndicatorFlag: strings.TrimSpace(s[electronicTransactionIndicatorFlagStartIndex-1 : (electronicTransactionIndicatorFlagStartIndex-1)+electronicTransactionIndicatorFlagLength]), // Field # 577

		StudentLastNameSSNChangeFlag: strings.TrimSpace(s[studentLastNameSSNChangeFlagStartIndex-1 : (studentLastNameSSNChangeFlagStartIndex-1)+studentLastNameSSNChangeFlagLength]), // Field # 578

		HighSchoolCode: strings.TrimSpace(s[highSchoolCodeStartIndex-1 : (highSchoolCodeStartIndex-1)+highSchoolCodeLength]), // Field # 579

		VerificationSelectionChangeFlag: strings.TrimSpace(s[verificationSelectionChangeFlagStartIndex-1 : (verificationSelectionChangeFlagStartIndex-1)+verificationSelectionChangeFlagLength]), // Field # 580

		UseUserProvidedDataOnly: strings.TrimSpace(s[useUserProvidedDataOnlyStartIndex-1 : (useUserProvidedDataOnlyStartIndex-1)+useUserProvidedDataOnlyLength]), // Field # 581

		NSLDSPellOverpaymentFlag: strings.TrimSpace(s[nsldsPellOverpaymentFlagStartIndex-1 : (nsldsPellOverpaymentFlagStartIndex-1)+nsldsPellOverpaymentFlagLength]), // Field # 583

		NSLDSPellOverpaymentContact: strings.TrimSpace(s[nsldsPellOverpaymentContactStartIndex-1 : (nsldsPellOverpaymentContactStartIndex-1)+nsldsPellOverpaymentContactLength]), // Field # 584

		NSLDSFSEOGOverpaymentFlag: strings.TrimSpace(s[nsldsFSEOGOverpaymentFlagStartIndex-1 : (nsldsFSEOGOverpaymentFlagStartIndex-1)+nsldsFSEOGOverpaymentFlagLength]), // Field # 585

		NSLDSFSEOGOverpaymentContact: strings.TrimSpace(s[nsldsFSEOGOverpaymentContactStartIndex-1 : (nsldsFSEOGOverpaymentContactStartIndex-1)+nsldsFSEOGOverpaymentContactLength]), // Field # 586

		NSLDSPerkinsOverpaymentFlag: strings.TrimSpace(s[nsldsPerkinsOverpaymentFlagStartIndex-1 : (nsldsPerkinsOverpaymentFlagStartIndex-1)+nsldsPerkinsOverpaymentFlagLength]), // Field # 587

		NSLDSPerkinsOverpaymentContact: strings.TrimSpace(s[nsldsPerkinsOverpaymentContactStartIndex-1 : (nsldsPerkinsOverpaymentContactStartIndex-1)+nsldsPerkinsOverpaymentContactLength]), // Field # 588

		NSLDSTEACHGrantOverpaymentFlag: strings.TrimSpace(s[nsldsTEACHGrantOverpaymentFlagStartIndex-1 : (nsldsTEACHGrantOverpaymentFlagStartIndex-1)+nsldsTEACHGrantOverpaymentFlagLength]), // Field # 589

		NSLDSTEACHGrantOverpaymentContact: strings.TrimSpace(s[nsldsTEACHGrantOverpaymentContactStartIndex-1 : (nsldsTEACHGrantOverpaymentContactStartIndex-1)+nsldsTEACHGrantOverpaymentContactLength]), // Field # 590

		NSLDSIraqAndAfghanistanServiceGrantOverpaymentFlag: strings.TrimSpace(s[nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagStartIndex-1 : (nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagStartIndex-1)+nsldsIraqAndAfghanistanServiceGrantOverpaymentFlagLength]), // Field # 591

		NSLDSIraqAndAfghanistanServiceGrantOverpaymentContact: strings.TrimSpace(s[nsldsIraqAndAfghanistanServiceGrantOverpaymentContactStartIndex-1 : (nsldsIraqAndAfghanistanServiceGrantOverpaymentContactStartIndex-1)+nsldsIraqAndAfghanistanServiceGrantOverpaymentContactLength]), // Field # 592

		NSLDSDefaultedLoanFlag: strings.TrimSpace(s[nsldsDefaultedLoanFlagStartIndex-1 : (nsldsDefaultedLoanFlagStartIndex-1)+nsldsDefaultedLoanFlagLength]), // Field # 593

		NSLDSDischargedLoanFlag: strings.TrimSpace(s[nsldsDischargedLoanFlagStartIndex-1 : (nsldsDischargedLoanFlagStartIndex-1)+nsldsDischargedLoanFlagLength]), // Field # 594

		NSLDSFraudLoanFlag: strings.TrimSpace(s[nsldsFraudLoanFlagStartIndex-1 : (nsldsFraudLoanFlagStartIndex-1)+nsldsFraudLoanFlagLength]), // Field # 595

		NSLDSSatisfactoryArrangementsFlag: strings.TrimSpace(s[nsldsSatisfactoryArrangementsFlagStartIndex-1 : (nsldsSatisfactoryArrangementsFlagStartIndex-1)+nsldsSatisfactoryArrangementsFlagLength]), // Field # 596

		NSLDSActiveBankruptcyFlag: strings.TrimSpace(s[nsldsActiveBankruptcyFlagStartIndex-1 : (nsldsActiveBankruptcyFlagStartIndex-1)+nsldsActiveBankruptcyFlagLength]), // Field # 597

		NSLDSTEACHGrantConvertedToLoanFlag: strings.TrimSpace(s[nsldsTEACHGrantConvertedToLoanFlagStartIndex-1 : (nsldsTEACHGrantConvertedToLoanFlagStartIndex-1)+nsldsTEACHGrantConvertedToLoanFlagLength]), // Field # 598

		NSLDSAggregateSubsidizedOutstandingPrincipalBalance: strings.TrimSpace(s[nsldsAggregateSubsidizedOutstandingPrincipalBalanceStartIndex-1 : (nsldsAggregateSubsidizedOutstandingPrincipalBalanceStartIndex-1)+nsldsAggregateSubsidizedOutstandingPrincipalBalanceLength]), // Field # 599

		NSLDSAggregateUnsubsidizedOutstandingPrincipalBalance: strings.TrimSpace(s[nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceStartIndex-1 : (nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceStartIndex-1)+nsldsAggregateUnsubsidizedOutstandingPrincipalBalanceLength]), // Field # 600

		NSLDSAggregateCombinedOutstandingPrincipalBalance: strings.TrimSpace(s[nsldsAggregateCombinedOutstandingPrincipalBalanceStartIndex-1 : (nsldsAggregateCombinedOutstandingPrincipalBalanceStartIndex-1)+nsldsAggregateCombinedOutstandingPrincipalBalanceLength]), // Field # 601

		NSLDSAggregateUnallocatedConsolidatedOutstandingPrincipalBalance: strings.TrimSpace(s[nsldsAggregateUnallocatedConsolidatedOutstandingPrincipalBalanceStartIndex-1 : (nsldsAggregateUnallocatedConsolidatedOutstandingPrincipalBalanceStartIndex-1)+nsldsAggregateUnallocatedConsolidatedOutstandingPrincipalBalanceLength]), // Field # 602

		NSLDSAggregateTEACHLoanPrincipalBalance: strings.TrimSpace(s[nsldsAggregateTEACHLoanPrincipalBalanceStartIndex-1 : (nsldsAggregateTEACHLoanPrincipalBalanceStartIndex-1)+nsldsAggregateTEACHLoanPrincipalBalanceLength]), // Field # 603

		NSLDSAggregateSubsidizedPendingDisbursement: strings.TrimSpace(s[nsldsAggregateSubsidizedPendingDisbursementStartIndex-1 : (nsldsAggregateSubsidizedPendingDisbursementStartIndex-1)+nsldsAggregateSubsidizedPendingDisbursementLength]), // Field # 604

		NSLDSAggregateUnsubsidizedPendingDisbursement: strings.TrimSpace(s[nsldsAggregateUnsubsidizedPendingDisbursementStartIndex-1 : (nsldsAggregateUnsubsidizedPendingDisbursementStartIndex-1)+nsldsAggregateUnsubsidizedPendingDisbursementLength]), // Field # 605

		NSLDSAggregateCombinedPendingDisbursement: strings.TrimSpace(s[nsldsAggregateCombinedPendingDisbursementStartIndex-1 : (nsldsAggregateCombinedPendingDisbursementStartIndex-1)+nsldsAggregateCombinedPendingDisbursementLength]), // Field # 606

		NSLDSAggregateSubsidizedTotal: strings.TrimSpace(s[nsldsAggregateSubsidizedTotalStartIndex-1 : (nsldsAggregateSubsidizedTotalStartIndex-1)+nsldsAggregateSubsidizedTotalLength]), // Field # 607

		NSLDSAggregateUnsubsidizedTotal: strings.TrimSpace(s[nsldsAggregateUnsubsidizedTotalStartIndex-1 : (nsldsAggregateUnsubsidizedTotalStartIndex-1)+nsldsAggregateUnsubsidizedTotalLength]), // Field # 608

		NSLDSAggregateCombinedTotal: strings.TrimSpace(s[nsldsAggregateCombinedTotalStartIndex-1 : (nsldsAggregateCombinedTotalStartIndex-1)+nsldsAggregateCombinedTotalLength]), // Field # 609

		NSLDSUnallocatedConsolidatedTotal: strings.TrimSpace(s[nsldsUnallocatedConsolidatedTotalStartIndex-1 : (nsldsUnallocatedConsolidatedTotalStartIndex-1)+nsldsUnallocatedConsolidatedTotalLength]), // Field # 610

		NSLDSTEACHLoanTotal: strings.TrimSpace(s[nsldsTEACHLoanTotalStartIndex-1 : (nsldsTEACHLoanTotalStartIndex-1)+nsldsTEACHLoanTotalLength]), // Field # 611

		NSLDSPerkinsTotalDisbursements: strings.TrimSpace(s[nsldsPerkinsTotalDisbursementsStartIndex-1 : (nsldsPerkinsTotalDisbursementsStartIndex-1)+nsldsPerkinsTotalDisbursementsLength]), // Field # 612

		NSLDSPerkinsCurrentYearDisbursementAmount: strings.TrimSpace(s[nsldsPerkinsCurrentYearDisbursementAmountStartIndex-1 : (nsldsPerkinsCurrentYearDisbursementAmountStartIndex-1)+nsldsPerkinsCurrentYearDisbursementAmountLength]), // Field # 613

		NSLDSAggregateTEACHGrantUndergraduateDisbursedTotal: strings.TrimSpace(s[nsldsAggregateTEACHGrantUndergraduateDisbursedTotalStartIndex-1 : (nsldsAggregateTEACHGrantUndergraduateDisbursedTotalStartIndex-1)+nsldsAggregateTEACHGrantUndergraduateDisbursedTotalLength]), // Field # 614

		NSLDSAggregateTEACHGraduateDisbursementAmount: strings.TrimSpace(s[nsldsAggregateTEACHGraduateDisbursementAmountStartIndex-1 : (nsldsAggregateTEACHGraduateDisbursementAmountStartIndex-1)+nsldsAggregateTEACHGraduateDisbursementAmountLength]), // Field # 615

		NSLDSDefaultedLoanChangeFlag: strings.TrimSpace(s[nsldsDefaultedLoanChangeFlagStartIndex-1 : (nsldsDefaultedLoanChangeFlagStartIndex-1)+nsldsDefaultedLoanChangeFlagLength]), // Field # 616

		NSLDSFraudLoanChangeFlag: strings.TrimSpace(s[nsldsFraudLoanChangeFlagStartIndex-1 : (nsldsFraudLoanChangeFlagStartIndex-1)+nsldsFraudLoanChangeFlagLength]), // Field # 617

		NSLDSDischargedLoanChangeFlag: strings.TrimSpace(s[nsldsDischargedLoanChangeFlagStartIndex-1 : (nsldsDischargedLoanChangeFlagStartIndex-1)+nsldsDischargedLoanChangeFlagLength]), // Field # 618

		NSLDSLoanSatisfactoryRepaymentChangeFlag: strings.TrimSpace(s[nsldsLoanSatisfactoryRepaymentChangeFlagStartIndex-1 : (nsldsLoanSatisfactoryRepaymentChangeFlagStartIndex-1)+nsldsLoanSatisfactoryRepaymentChangeFlagLength]), // Field # 619

		NSLDSActiveBankruptcyChangeFlag: strings.TrimSpace(s[nsldsActiveBankruptcyChangeFlagStartIndex-1 : (nsldsActiveBankruptcyChangeFlagStartIndex-1)+nsldsActiveBankruptcyChangeFlagLength]), // Field # 620

		NSLDSTEACHGrantToLoanConversionChangeFlag: strings.TrimSpace(s[nsldsTEACHGrantToLoanConversionChangeFlagStartIndex-1 : (nsldsTEACHGrantToLoanConversionChangeFlagStartIndex-1)+nsldsTEACHGrantToLoanConversionChangeFlagLength]), // Field # 621

		NSLDSOverpaymentsChangeFlag: strings.TrimSpace(s[nsldsOverpaymentsChangeFlagStartIndex-1 : (nsldsOverpaymentsChangeFlagStartIndex-1)+nsldsOverpaymentsChangeFlagLength]), // Field # 622

		NSLDSAggregateLoanChangeFlag: strings.TrimSpace(s[nsldsAggregateLoanChangeFlagStartIndex-1 : (nsldsAggregateLoanChangeFlagStartIndex-1)+nsldsAggregateLoanChangeFlagLength]), // Field # 623

		NSLDSPerkinsLoanChangeFlag: strings.TrimSpace(s[nsldsPerkinsLoanChangeFlagStartIndex-1 : (nsldsPerkinsLoanChangeFlagStartIndex-1)+nsldsPerkinsLoanChangeFlagLength]), // Field # 624

		NSLDSPellPaymentChangeFlag: strings.TrimSpace(s[nsldsPellPaymentChangeFlagStartIndex-1 : (nsldsPellPaymentChangeFlagStartIndex-1)+nsldsPellPaymentChangeFlagLength]), // Field # 625

		NSLDSTEACHGrantChangeFlag: strings.TrimSpace(s[nsldsTEACHGrantChangeFlagStartIndex-1 : (nsldsTEACHGrantChangeFlagStartIndex-1)+nsldsTEACHGrantChangeFlagLength]), // Field # 626

		NSLDSAdditionalPellFlag: strings.TrimSpace(s[nsldsAdditionalPellFlagStartIndex-1 : (nsldsAdditionalPellFlagStartIndex-1)+nsldsAdditionalPellFlagLength]), // Field # 627

		NSLDSAdditionalLoansFlag: strings.TrimSpace(s[nsldsAdditionalLoansFlagStartIndex-1 : (nsldsAdditionalLoansFlagStartIndex-1)+nsldsAdditionalLoansFlagLength]), // Field # 628

		NSLDSAdditionalTEACHGrantFlag: strings.TrimSpace(s[nsldsAdditionalTEACHGrantFlagStartIndex-1 : (nsldsAdditionalTEACHGrantFlagStartIndex-1)+nsldsAdditionalTEACHGrantFlagLength]), // Field # 629

		NSLDSDirectLoanMasterPromNoteFlag: strings.TrimSpace(s[nsldsDirectLoanMasterPromNoteFlagStartIndex-1 : (nsldsDirectLoanMasterPromNoteFlagStartIndex-1)+nsldsDirectLoanMasterPromNoteFlagLength]), // Field # 630

		NSLDSDirectLoanPLUSMasterPromNoteFlag: strings.TrimSpace(s[nsldsDirectLoanPLUSMasterPromNoteFlagStartIndex-1 : (nsldsDirectLoanPLUSMasterPromNoteFlagStartIndex-1)+nsldsDirectLoanPLUSMasterPromNoteFlagLength]), // Field # 631

		NSLDSDirectLoanGraduatePLUSMasterPromNoteFlag: strings.TrimSpace(s[nsldsDirectLoanGraduatePLUSMasterPromNoteFlagStartIndex-1 : (nsldsDirectLoanGraduatePLUSMasterPromNoteFlagStartIndex-1)+nsldsDirectLoanGraduatePLUSMasterPromNoteFlagLength]), // Field # 632

		NSLDSUndergraduateSubsidizedLoanLimitFlag: strings.TrimSpace(s[nsldsUndergraduateSubsidizedLoanLimitFlagStartIndex-1 : (nsldsUndergraduateSubsidizedLoanLimitFlagStartIndex-1)+nsldsUndergraduateSubsidizedLoanLimitFlagLength]), // Field # 633

		NSLDSUndergraduateCombinedLoanLimitFlag: strings.TrimSpace(s[nsldsUndergraduateCombinedLoanLimitFlagStartIndex-1 : (nsldsUndergraduateCombinedLoanLimitFlagStartIndex-1)+nsldsUndergraduateCombinedLoanLimitFlagLength]), // Field # 634

		NSLDSGraduateSubsidizedLoanLimitFlag: strings.TrimSpace(s[nsldsGraduateSubsidizedLoanLimitFlagStartIndex-1 : (nsldsGraduateSubsidizedLoanLimitFlagStartIndex-1)+nsldsGraduateSubsidizedLoanLimitFlagLength]), // Field # 635

		NSLDSGraduateCombinedLoanLimitFlag: strings.TrimSpace(s[nsldsGraduateCombinedLoanLimitFlagStartIndex-1 : (nsldsGraduateCombinedLoanLimitFlagStartIndex-1)+nsldsGraduateCombinedLoanLimitFlagLength]), // Field # 636

		NSLDSLEULimitIndicator: strings.TrimSpace(s[nsldsLEULimitIndicatorStartIndex-1 : (nsldsLEULimitIndicatorStartIndex-1)+nsldsLEULimitIndicatorLength]), // Field # 637

		NSLDSPellLifetimeEligibilityUsed: strings.TrimSpace(s[nsldsPellLifetimeEligibilityUsedStartIndex-1 : (nsldsPellLifetimeEligibilityUsedStartIndex-1)+nsldsPellLifetimeEligibilityUsedLength]), // Field # 638

		NSLDSSULAFlag: strings.TrimSpace(s[nsldsSULAFlagStartIndex-1 : (nsldsSULAFlagStartIndex-1)+nsldsSULAFlagLength]), // Field # 639

		NSLDSSubsidizedLimitEligibilityFlag: strings.TrimSpace(s[nsldsSubsidizedLimitEligibilityFlagStartIndex-1 : (nsldsSubsidizedLimitEligibilityFlagStartIndex-1)+nsldsSubsidizedLimitEligibilityFlagLength]), // Field # 640

		NSLDSUnusualEnrollmentHistoryFlag: strings.TrimSpace(s[nsldsUnusualEnrollmentHistoryFlagStartIndex-1 : (nsldsUnusualEnrollmentHistoryFlagStartIndex-1)+nsldsUnusualEnrollmentHistoryFlagLength]), // Field # 641

		NSLDSPellSequenceNumber1: strings.TrimSpace(s[nsldsPellSequenceNumber1StartIndex-1 : (nsldsPellSequenceNumber1StartIndex-1)+nsldsPellSequenceNumber1Length]), // Field # 643

		NSLDSPellVerificationFlag1: strings.TrimSpace(s[nsldsPellVerificationFlag1StartIndex-1 : (nsldsPellVerificationFlag1StartIndex-1)+nsldsPellVerificationFlag1Length]), // Field # 644

		NSLDSSAI1: strings.TrimSpace(s[nsldsSAI1StartIndex-1 : (nsldsSAI1StartIndex-1)+nsldsSAI1Length]), // Field # 645

		NSLDSPellSchoolCode1: strings.TrimSpace(s[nsldsPellSchoolCode1StartIndex-1 : (nsldsPellSchoolCode1StartIndex-1)+nsldsPellSchoolCode1Length]), // Field # 646

		NSLDSPellTransactionNumber1: strings.TrimSpace(s[nsldsPellTransactionNumber1StartIndex-1 : (nsldsPellTransactionNumber1StartIndex-1)+nsldsPellTransactionNumber1Length]), // Field # 647

		NSLDSPellLastDisbursementDate1: parseISIRDate(strings.TrimSpace(s[nsldsPellLastDisbursementDate1StartIndex-1 : (nsldsPellLastDisbursementDate1StartIndex-1)+nsldsPellLastDisbursementDate1Length])), // Field # 648

		NSLDSPellScheduledAmount1: strings.TrimSpace(s[nsldsPellScheduledAmount1StartIndex-1 : (nsldsPellScheduledAmount1StartIndex-1)+nsldsPellScheduledAmount1Length]), // Field # 649

		NSLDSPellAmountPaidToDate1: parseISIRDate(strings.TrimSpace(s[nsldsPellAmountPaidToDate1StartIndex-1 : (nsldsPellAmountPaidToDate1StartIndex-1)+nsldsPellAmountPaidToDate1Length])), // Field # 650

		NSLDSPellPercentEligibilityUsedDecimal1: strings.TrimSpace(s[nsldsPellPercentEligibilityUsedDecimal1StartIndex-1 : (nsldsPellPercentEligibilityUsedDecimal1StartIndex-1)+nsldsPellPercentEligibilityUsedDecimal1Length]), // Field # 651

		NSLDSPellAwardAmount1: strings.TrimSpace(s[nsldsPellAwardAmount1StartIndex-1 : (nsldsPellAwardAmount1StartIndex-1)+nsldsPellAwardAmount1Length]), // Field # 652

		NSLDSAdditionalEligibilityIndicator1: strings.TrimSpace(s[nsldsAdditionalEligibilityIndicator1StartIndex-1 : (nsldsAdditionalEligibilityIndicator1StartIndex-1)+nsldsAdditionalEligibilityIndicator1Length]), // Field # 653

		NSLDSPellSequenceNumber2: strings.TrimSpace(s[nsldsPellSequenceNumber2StartIndex-1 : (nsldsPellSequenceNumber2StartIndex-1)+nsldsPellSequenceNumber2Length]), // Field # 655

		NSLDSPellVerificationFlag2: strings.TrimSpace(s[nsldsPellVerificationFlag2StartIndex-1 : (nsldsPellVerificationFlag2StartIndex-1)+nsldsPellVerificationFlag2Length]), // Field # 656

		NSLDSSAI2: strings.TrimSpace(s[nsldsSAI2StartIndex-1 : (nsldsSAI2StartIndex-1)+nsldsSAI2Length]), // Field # 657

		NSLDSPellSchoolCode2: strings.TrimSpace(s[nsldsPellSchoolCode2StartIndex-1 : (nsldsPellSchoolCode2StartIndex-1)+nsldsPellSchoolCode2Length]), // Field # 658

		NSLDSPellTransactionNumber2: strings.TrimSpace(s[nsldsPellTransactionNumber2StartIndex-1 : (nsldsPellTransactionNumber2StartIndex-1)+nsldsPellTransactionNumber2Length]), // Field # 659

		NSLDSPellLastDisbursementDate2: parseISIRDate(strings.TrimSpace(s[nsldsPellLastDisbursementDate2StartIndex-1 : (nsldsPellLastDisbursementDate2StartIndex-1)+nsldsPellLastDisbursementDate2Length])), // Field # 660

		NSLDSPellScheduledAmount2: strings.TrimSpace(s[nsldsPellScheduledAmount2StartIndex-1 : (nsldsPellScheduledAmount2StartIndex-1)+nsldsPellScheduledAmount2Length]), // Field # 661

		NSLDSPellAmountPaidToDate2: parseISIRDate(strings.TrimSpace(s[nsldsPellAmountPaidToDate2StartIndex-1 : (nsldsPellAmountPaidToDate2StartIndex-1)+nsldsPellAmountPaidToDate2Length])), // Field # 662

		NSLDSPellPercentEligibilityUsedDecimal2: strings.TrimSpace(s[nsldsPellPercentEligibilityUsedDecimal2StartIndex-1 : (nsldsPellPercentEligibilityUsedDecimal2StartIndex-1)+nsldsPellPercentEligibilityUsedDecimal2Length]), // Field # 663

		NSLDSPellAwardAmount2: strings.TrimSpace(s[nsldsPellAwardAmount2StartIndex-1 : (nsldsPellAwardAmount2StartIndex-1)+nsldsPellAwardAmount2Length]), // Field # 664

		NSLDSAdditionalEligibilityIndicator2: strings.TrimSpace(s[nsldsAdditionalEligibilityIndicator2StartIndex-1 : (nsldsAdditionalEligibilityIndicator2StartIndex-1)+nsldsAdditionalEligibilityIndicator2Length]), // Field # 665

		NSLDSPellSequenceNumber3: strings.TrimSpace(s[nsldsPellSequenceNumber3StartIndex-1 : (nsldsPellSequenceNumber3StartIndex-1)+nsldsPellSequenceNumber3Length]), // Field # 667

		NSLDSPellVerificationFlag3: strings.TrimSpace(s[nsldsPellVerificationFlag3StartIndex-1 : (nsldsPellVerificationFlag3StartIndex-1)+nsldsPellVerificationFlag3Length]), // Field # 668

		NSLDSSAI3: strings.TrimSpace(s[nsldsSAI3StartIndex-1 : (nsldsSAI3StartIndex-1)+nsldsSAI3Length]), // Field # 669

		NSLDSPellSchoolCode3: strings.TrimSpace(s[nsldsPellSchoolCode3StartIndex-1 : (nsldsPellSchoolCode3StartIndex-1)+nsldsPellSchoolCode3Length]), // Field # 670

		NSLDSPellTransactionNumber3: strings.TrimSpace(s[nsldsPellTransactionNumber3StartIndex-1 : (nsldsPellTransactionNumber3StartIndex-1)+nsldsPellTransactionNumber3Length]), // Field # 671

		NSLDSPellLastDisbursementDate3: parseISIRDate(strings.TrimSpace(s[nsldsPellLastDisbursementDate3StartIndex-1 : (nsldsPellLastDisbursementDate3StartIndex-1)+nsldsPellLastDisbursementDate3Length])), // Field # 672

		NSLDSPellScheduledAmount3: strings.TrimSpace(s[nsldsPellScheduledAmount3StartIndex-1 : (nsldsPellScheduledAmount3StartIndex-1)+nsldsPellScheduledAmount3Length]), // Field # 673

		NSLDSPellAmountPaidToDate3: parseISIRDate(strings.TrimSpace(s[nsldsPellAmountPaidToDate3StartIndex-1 : (nsldsPellAmountPaidToDate3StartIndex-1)+nsldsPellAmountPaidToDate3Length])), // Field # 674

		NSLDSPellPercentEligibilityUsedDecimal3: strings.TrimSpace(s[nsldsPellPercentEligibilityUsedDecimal3StartIndex-1 : (nsldsPellPercentEligibilityUsedDecimal3StartIndex-1)+nsldsPellPercentEligibilityUsedDecimal3Length]), // Field # 675

		NSLDSPellAwardAmount3: strings.TrimSpace(s[nsldsPellAwardAmount3StartIndex-1 : (nsldsPellAwardAmount3StartIndex-1)+nsldsPellAwardAmount3Length]), // Field # 676

		NSLDSAdditionalEligibilityIndicator3: strings.TrimSpace(s[nsldsAdditionalEligibilityIndicator3StartIndex-1 : (nsldsAdditionalEligibilityIndicator3StartIndex-1)+nsldsAdditionalEligibilityIndicator3Length]), // Field # 677

		NSLDSTEACHGrantSequence1: strings.TrimSpace(s[nsldsTEACHGrantSequence1StartIndex-1 : (nsldsTEACHGrantSequence1StartIndex-1)+nsldsTEACHGrantSequence1Length]), // Field # 679

		NSLDSTEACHGrantSchoolCode1: strings.TrimSpace(s[nsldsTEACHGrantSchoolCode1StartIndex-1 : (nsldsTEACHGrantSchoolCode1StartIndex-1)+nsldsTEACHGrantSchoolCode1Length]), // Field # 680

		NSLDSTEACHGrantTransactionNumber1: strings.TrimSpace(s[nsldsTEACHGrantTransactionNumber1StartIndex-1 : (nsldsTEACHGrantTransactionNumber1StartIndex-1)+nsldsTEACHGrantTransactionNumber1Length]), // Field # 681

		NSLDSTEACHGrantLastDisbursementDate1: parseISIRDate(strings.TrimSpace(s[nsldsTEACHGrantLastDisbursementDate1StartIndex-1 : (nsldsTEACHGrantLastDisbursementDate1StartIndex-1)+nsldsTEACHGrantLastDisbursementDate1Length])), // Field # 682

		NSLDSTEACHGrantScheduledAmount1: strings.TrimSpace(s[nsldsTEACHGrantScheduledAmount1StartIndex-1 : (nsldsTEACHGrantScheduledAmount1StartIndex-1)+nsldsTEACHGrantScheduledAmount1Length]), // Field # 683

		NSLDSTEACHGrantAmountPaidToDate1: parseISIRDate(strings.TrimSpace(s[nsldsTEACHGrantAmountPaidToDate1StartIndex-1 : (nsldsTEACHGrantAmountPaidToDate1StartIndex-1)+nsldsTEACHGrantAmountPaidToDate1Length])), // Field # 684

		NSLDSTEACHGrantAwardAmount1: strings.TrimSpace(s[nsldsTEACHGrantAwardAmount1StartIndex-1 : (nsldsTEACHGrantAwardAmount1StartIndex-1)+nsldsTEACHGrantAwardAmount1Length]), // Field # 685

		NSLDSTEACHGrantAcademicYearLevel1: strings.TrimSpace(s[nsldsTEACHGrantAcademicYearLevel1StartIndex-1 : (nsldsTEACHGrantAcademicYearLevel1StartIndex-1)+nsldsTEACHGrantAcademicYearLevel1Length]), // Field # 686

		NSLDSTEACHGrantAwardYear1: strings.TrimSpace(s[nsldsTEACHGrantAwardYear1StartIndex-1 : (nsldsTEACHGrantAwardYear1StartIndex-1)+nsldsTEACHGrantAwardYear1Length]), // Field # 687

		NSLDSTEACHGrantLoanConversionFlag1: strings.TrimSpace(s[nsldsTEACHGrantLoanConversionFlag1StartIndex-1 : (nsldsTEACHGrantLoanConversionFlag1StartIndex-1)+nsldsTEACHGrantLoanConversionFlag1Length]), // Field # 688

		NSLDSTEACHGrantDischargeCode1: strings.TrimSpace(s[nsldsTEACHGrantDischargeCode1StartIndex-1 : (nsldsTEACHGrantDischargeCode1StartIndex-1)+nsldsTEACHGrantDischargeCode1Length]), // Field # 689

		NSLDSTEACHGrantDischargeAmount1: strings.TrimSpace(s[nsldsTEACHGrantDischargeAmount1StartIndex-1 : (nsldsTEACHGrantDischargeAmount1StartIndex-1)+nsldsTEACHGrantDischargeAmount1Length]), // Field # 690

		NSLDSTEACHGrantAdjustedDisbursement1: strings.TrimSpace(s[nsldsTEACHGrantAdjustedDisbursement1StartIndex-1 : (nsldsTEACHGrantAdjustedDisbursement1StartIndex-1)+nsldsTEACHGrantAdjustedDisbursement1Length]), // Field # 691

		NSLDSTEACHGrantSequence2: strings.TrimSpace(s[nsldsTEACHGrantSequence2StartIndex-1 : (nsldsTEACHGrantSequence2StartIndex-1)+nsldsTEACHGrantSequence2Length]), // Field # 693

		NSLDSTEACHGrantSchoolCode2: strings.TrimSpace(s[nsldsTEACHGrantSchoolCode2StartIndex-1 : (nsldsTEACHGrantSchoolCode2StartIndex-1)+nsldsTEACHGrantSchoolCode2Length]), // Field # 694

		NSLDSTEACHGrantTransactionNumber2: strings.TrimSpace(s[nsldsTEACHGrantTransactionNumber2StartIndex-1 : (nsldsTEACHGrantTransactionNumber2StartIndex-1)+nsldsTEACHGrantTransactionNumber2Length]), // Field # 695

		NSLDSTEACHGrantLastDisbursementDate2: parseISIRDate(strings.TrimSpace(s[nsldsTEACHGrantLastDisbursementDate2StartIndex-1 : (nsldsTEACHGrantLastDisbursementDate2StartIndex-1)+nsldsTEACHGrantLastDisbursementDate2Length])), // Field # 696

		NSLDSTEACHGrantScheduledAmount2: strings.TrimSpace(s[nsldsTEACHGrantScheduledAmount2StartIndex-1 : (nsldsTEACHGrantScheduledAmount2StartIndex-1)+nsldsTEACHGrantScheduledAmount2Length]), // Field # 697

		NSLDSTEACHGrantAmountPaidToDate2: parseISIRDate(strings.TrimSpace(s[nsldsTEACHGrantAmountPaidToDate2StartIndex-1 : (nsldsTEACHGrantAmountPaidToDate2StartIndex-1)+nsldsTEACHGrantAmountPaidToDate2Length])), // Field # 698

		NSLDSTEACHGrantAwardAmount2: strings.TrimSpace(s[nsldsTEACHGrantAwardAmount2StartIndex-1 : (nsldsTEACHGrantAwardAmount2StartIndex-1)+nsldsTEACHGrantAwardAmount2Length]), // Field # 699

		NSLDSTEACHGrantAcademicYearLevel2: strings.TrimSpace(s[nsldsTEACHGrantAcademicYearLevel2StartIndex-1 : (nsldsTEACHGrantAcademicYearLevel2StartIndex-1)+nsldsTEACHGrantAcademicYearLevel2Length]), // Field # 700

		NSLDSTEACHGrantAwardYear2: strings.TrimSpace(s[nsldsTEACHGrantAwardYear2StartIndex-1 : (nsldsTEACHGrantAwardYear2StartIndex-1)+nsldsTEACHGrantAwardYear2Length]), // Field # 701

		NSLDSTEACHGrantLoanConversionFlag2: strings.TrimSpace(s[nsldsTEACHGrantLoanConversionFlag2StartIndex-1 : (nsldsTEACHGrantLoanConversionFlag2StartIndex-1)+nsldsTEACHGrantLoanConversionFlag2Length]), // Field # 702

		NSLDSTEACHGrantDischargeCode2: strings.TrimSpace(s[nsldsTEACHGrantDischargeCode2StartIndex-1 : (nsldsTEACHGrantDischargeCode2StartIndex-1)+nsldsTEACHGrantDischargeCode2Length]), // Field # 703

		NSLDSTEACHGrantDischargeAmount2: strings.TrimSpace(s[nsldsTEACHGrantDischargeAmount2StartIndex-1 : (nsldsTEACHGrantDischargeAmount2StartIndex-1)+nsldsTEACHGrantDischargeAmount2Length]), // Field # 704

		NSLDSTEACHGrantAdjustedDisbursement2: strings.TrimSpace(s[nsldsTEACHGrantAdjustedDisbursement2StartIndex-1 : (nsldsTEACHGrantAdjustedDisbursement2StartIndex-1)+nsldsTEACHGrantAdjustedDisbursement2Length]), // Field # 705

		NSLDSTEACHGrantSequence3: strings.TrimSpace(s[nsldsTEACHGrantSequence3StartIndex-1 : (nsldsTEACHGrantSequence3StartIndex-1)+nsldsTEACHGrantSequence3Length]), // Field # 707

		NSLDSTEACHGrantSchoolCode3: strings.TrimSpace(s[nsldsTEACHGrantSchoolCode3StartIndex-1 : (nsldsTEACHGrantSchoolCode3StartIndex-1)+nsldsTEACHGrantSchoolCode3Length]), // Field # 708

		NSLDSTEACHGrantTransactionNumber3: strings.TrimSpace(s[nsldsTEACHGrantTransactionNumber3StartIndex-1 : (nsldsTEACHGrantTransactionNumber3StartIndex-1)+nsldsTEACHGrantTransactionNumber3Length]), // Field # 709

		NSLDSTEACHGrantLastDisbursementDate3: parseISIRDate(strings.TrimSpace(s[nsldsTEACHGrantLastDisbursementDate3StartIndex-1 : (nsldsTEACHGrantLastDisbursementDate3StartIndex-1)+nsldsTEACHGrantLastDisbursementDate3Length])), // Field # 710

		NSLDSTEACHGrantScheduledAmount3: strings.TrimSpace(s[nsldsTEACHGrantScheduledAmount3StartIndex-1 : (nsldsTEACHGrantScheduledAmount3StartIndex-1)+nsldsTEACHGrantScheduledAmount3Length]), // Field # 711

		NSLDSTEACHGrantAmountPaidToDate3: parseISIRDate(strings.TrimSpace(s[nsldsTEACHGrantAmountPaidToDate3StartIndex-1 : (nsldsTEACHGrantAmountPaidToDate3StartIndex-1)+nsldsTEACHGrantAmountPaidToDate3Length])), // Field # 712

		NSLDSTEACHGrantAwardAmount3: strings.TrimSpace(s[nsldsTEACHGrantAwardAmount3StartIndex-1 : (nsldsTEACHGrantAwardAmount3StartIndex-1)+nsldsTEACHGrantAwardAmount3Length]), // Field # 713

		NSLDSTEACHGrantAcademicYearLevel3: strings.TrimSpace(s[nsldsTEACHGrantAcademicYearLevel3StartIndex-1 : (nsldsTEACHGrantAcademicYearLevel3StartIndex-1)+nsldsTEACHGrantAcademicYearLevel3Length]), // Field # 714

		NSLDSTEACHGrantAwardYear3: strings.TrimSpace(s[nsldsTEACHGrantAwardYear3StartIndex-1 : (nsldsTEACHGrantAwardYear3StartIndex-1)+nsldsTEACHGrantAwardYear3Length]), // Field # 715

		NSLDSTEACHGrantLoanConversionFlag3: strings.TrimSpace(s[nsldsTEACHGrantLoanConversionFlag3StartIndex-1 : (nsldsTEACHGrantLoanConversionFlag3StartIndex-1)+nsldsTEACHGrantLoanConversionFlag3Length]), // Field # 716

		NSLDSTEACHGrantDischargeCode3: strings.TrimSpace(s[nsldsTEACHGrantDischargeCode3StartIndex-1 : (nsldsTEACHGrantDischargeCode3StartIndex-1)+nsldsTEACHGrantDischargeCode3Length]), // Field # 717

		NSLDSTEACHGrantDischargeAmount3: strings.TrimSpace(s[nsldsTEACHGrantDischargeAmount3StartIndex-1 : (nsldsTEACHGrantDischargeAmount3StartIndex-1)+nsldsTEACHGrantDischargeAmount3Length]), // Field # 718

		NSLDSTEACHGrantAdjustedDisbursement3: strings.TrimSpace(s[nsldsTEACHGrantAdjustedDisbursement3StartIndex-1 : (nsldsTEACHGrantAdjustedDisbursement3StartIndex-1)+nsldsTEACHGrantAdjustedDisbursement3Length]), // Field # 719

		NSLDSLoanSequenceNumber1: strings.TrimSpace(s[nsldsLoanSequenceNumber1StartIndex-1 : (nsldsLoanSequenceNumber1StartIndex-1)+nsldsLoanSequenceNumber1Length]), // Field # 721

		NSLDSLoanDefaultedRecentIndicator1: strings.TrimSpace(s[nsldsLoanDefaultedRecentIndicator1StartIndex-1 : (nsldsLoanDefaultedRecentIndicator1StartIndex-1)+nsldsLoanDefaultedRecentIndicator1Length]), // Field # 722

		NSLDSLoanChangeFlag1: strings.TrimSpace(s[nsldsLoanChangeFlag1StartIndex-1 : (nsldsLoanChangeFlag1StartIndex-1)+nsldsLoanChangeFlag1Length]), // Field # 723

		NSLDSLoanTypeCode1: strings.TrimSpace(s[nsldsLoanTypeCode1StartIndex-1 : (nsldsLoanTypeCode1StartIndex-1)+nsldsLoanTypeCode1Length]), // Field # 724

		NSLDSLoanNetAmount1: strings.TrimSpace(s[nsldsLoanNetAmount1StartIndex-1 : (nsldsLoanNetAmount1StartIndex-1)+nsldsLoanNetAmount1Length]), // Field # 725

		NSLDSLoanCurrentStatusCode1: strings.TrimSpace(s[nsldsLoanCurrentStatusCode1StartIndex-1 : (nsldsLoanCurrentStatusCode1StartIndex-1)+nsldsLoanCurrentStatusCode1Length]), // Field # 726

		NSLDSLoanCurrentStatusDate1: parseISIRDate(strings.TrimSpace(s[nsldsLoanCurrentStatusDate1StartIndex-1 : (nsldsLoanCurrentStatusDate1StartIndex-1)+nsldsLoanCurrentStatusDate1Length])), // Field # 727

		NSLDSLoanOutstandingPrincipalBalance1: strings.TrimSpace(s[nsldsLoanOutstandingPrincipalBalance1StartIndex-1 : (nsldsLoanOutstandingPrincipalBalance1StartIndex-1)+nsldsLoanOutstandingPrincipalBalance1Length]), // Field # 728

		NSLDSLoanOutstandingPrincipalBalanceDate1: parseISIRDate(strings.TrimSpace(s[nsldsLoanOutstandingPrincipalBalanceDate1StartIndex-1 : (nsldsLoanOutstandingPrincipalBalanceDate1StartIndex-1)+nsldsLoanOutstandingPrincipalBalanceDate1Length])), // Field # 729

		NSLDSLoanPeriodBeginDate1: parseISIRDate(strings.TrimSpace(s[nsldsLoanPeriodBeginDate1StartIndex-1 : (nsldsLoanPeriodBeginDate1StartIndex-1)+nsldsLoanPeriodBeginDate1Length])), // Field # 730

		NSLDSLoanPeriodEndDate1: parseISIRDate(strings.TrimSpace(s[nsldsLoanPeriodEndDate1StartIndex-1 : (nsldsLoanPeriodEndDate1StartIndex-1)+nsldsLoanPeriodEndDate1Length])), // Field # 731

		NSLDSLoanGuarantyAgencyCode1: strings.TrimSpace(s[nsldsLoanGuarantyAgencyCode1StartIndex-1 : (nsldsLoanGuarantyAgencyCode1StartIndex-1)+nsldsLoanGuarantyAgencyCode1Length]), // Field # 732

		NSLDSLoanContactType1: strings.TrimSpace(s[nsldsLoanContactType1StartIndex-1 : (nsldsLoanContactType1StartIndex-1)+nsldsLoanContactType1Length]), // Field # 733

		NSLDSLoanSchoolCode1: strings.TrimSpace(s[nsldsLoanSchoolCode1StartIndex-1 : (nsldsLoanSchoolCode1StartIndex-1)+nsldsLoanSchoolCode1Length]), // Field # 734

		NSLDSLoanContactCode1: strings.TrimSpace(s[nsldsLoanContactCode1StartIndex-1 : (nsldsLoanContactCode1StartIndex-1)+nsldsLoanContactCode1Length]), // Field # 735

		NSLDSLoanGradeLevel1: strings.TrimSpace(s[nsldsLoanGradeLevel1StartIndex-1 : (nsldsLoanGradeLevel1StartIndex-1)+nsldsLoanGradeLevel1Length]), // Field # 736

		NSLDSLoanAdditionalUnsubsidizedFlag1: strings.TrimSpace(s[nsldsLoanAdditionalUnsubsidizedFlag1StartIndex-1 : (nsldsLoanAdditionalUnsubsidizedFlag1StartIndex-1)+nsldsLoanAdditionalUnsubsidizedFlag1Length]), // Field # 737

		NSLDSLoanCapitalizedInterestFlag1: strings.TrimSpace(s[nsldsLoanCapitalizedInterestFlag1StartIndex-1 : (nsldsLoanCapitalizedInterestFlag1StartIndex-1)+nsldsLoanCapitalizedInterestFlag1Length]), // Field # 738

		NSLDSLoanDisbursementAmount1: strings.TrimSpace(s[nsldsLoanDisbursementAmount1StartIndex-1 : (nsldsLoanDisbursementAmount1StartIndex-1)+nsldsLoanDisbursementAmount1Length]), // Field # 739

		NSLDSLoanDisbursementDate1: parseISIRDate(strings.TrimSpace(s[nsldsLoanDisbursementDate1StartIndex-1 : (nsldsLoanDisbursementDate1StartIndex-1)+nsldsLoanDisbursementDate1Length])), // Field # 740

		NSLDSLoanConfirmedLoanSubsidyStatus1: strings.TrimSpace(s[nsldsLoanConfirmedLoanSubsidyStatus1StartIndex-1 : (nsldsLoanConfirmedLoanSubsidyStatus1StartIndex-1)+nsldsLoanConfirmedLoanSubsidyStatus1Length]), // Field # 741

		NSLDSLoanSubsidyStatusDate1: parseISIRDate(strings.TrimSpace(s[nsldsLoanSubsidyStatusDate1StartIndex-1 : (nsldsLoanSubsidyStatusDate1StartIndex-1)+nsldsLoanSubsidyStatusDate1Length])), // Field # 742

		NSLDSLoanSequenceNumber2: strings.TrimSpace(s[nsldsLoanSequenceNumber2StartIndex-1 : (nsldsLoanSequenceNumber2StartIndex-1)+nsldsLoanSequenceNumber2Length]), // Field # 744

		NSLDSLoanDefaultedRecentIndicator2: strings.TrimSpace(s[nsldsLoanDefaultedRecentIndicator2StartIndex-1 : (nsldsLoanDefaultedRecentIndicator2StartIndex-1)+nsldsLoanDefaultedRecentIndicator2Length]), // Field # 745

		NSLDSLoanChangeFlag2: strings.TrimSpace(s[nsldsLoanChangeFlag2StartIndex-1 : (nsldsLoanChangeFlag2StartIndex-1)+nsldsLoanChangeFlag2Length]), // Field # 746

		NSLDSLoanTypeCode2: strings.TrimSpace(s[nsldsLoanTypeCode2StartIndex-1 : (nsldsLoanTypeCode2StartIndex-1)+nsldsLoanTypeCode2Length]), // Field # 747

		NSLDSLoanNetAmount2: strings.TrimSpace(s[nsldsLoanNetAmount2StartIndex-1 : (nsldsLoanNetAmount2StartIndex-1)+nsldsLoanNetAmount2Length]), // Field # 748

		NSLDSLoanCurrentStatusCode2: strings.TrimSpace(s[nsldsLoanCurrentStatusCode2StartIndex-1 : (nsldsLoanCurrentStatusCode2StartIndex-1)+nsldsLoanCurrentStatusCode2Length]), // Field # 749

		NSLDSLoanCurrentStatusDate2: parseISIRDate(strings.TrimSpace(s[nsldsLoanCurrentStatusDate2StartIndex-1 : (nsldsLoanCurrentStatusDate2StartIndex-1)+nsldsLoanCurrentStatusDate2Length])), // Field # 750

		NSLDSLoanOutstandingPrincipalBalance2: strings.TrimSpace(s[nsldsLoanOutstandingPrincipalBalance2StartIndex-1 : (nsldsLoanOutstandingPrincipalBalance2StartIndex-1)+nsldsLoanOutstandingPrincipalBalance2Length]), // Field # 751

		NSLDSLoanOutstandingPrincipalBalanceDate2: parseISIRDate(strings.TrimSpace(s[nsldsLoanOutstandingPrincipalBalanceDate2StartIndex-1 : (nsldsLoanOutstandingPrincipalBalanceDate2StartIndex-1)+nsldsLoanOutstandingPrincipalBalanceDate2Length])), // Field # 752

		NSLDSLoanPeriodBeginDate2: parseISIRDate(strings.TrimSpace(s[nsldsLoanPeriodBeginDate2StartIndex-1 : (nsldsLoanPeriodBeginDate2StartIndex-1)+nsldsLoanPeriodBeginDate2Length])), // Field # 753

		NSLDSLoanPeriodEndDate2: parseISIRDate(strings.TrimSpace(s[nsldsLoanPeriodEndDate2StartIndex-1 : (nsldsLoanPeriodEndDate2StartIndex-1)+nsldsLoanPeriodEndDate2Length])), // Field # 754

		NSLDSLoanGuarantyAgencyCode2: strings.TrimSpace(s[nsldsLoanGuarantyAgencyCode2StartIndex-1 : (nsldsLoanGuarantyAgencyCode2StartIndex-1)+nsldsLoanGuarantyAgencyCode2Length]), // Field # 755

		NSLDSLoanContactType2: strings.TrimSpace(s[nsldsLoanContactType2StartIndex-1 : (nsldsLoanContactType2StartIndex-1)+nsldsLoanContactType2Length]), // Field # 756

		NSLDSLoanSchoolCode2: strings.TrimSpace(s[nsldsLoanSchoolCode2StartIndex-1 : (nsldsLoanSchoolCode2StartIndex-1)+nsldsLoanSchoolCode2Length]), // Field # 757

		NSLDSLoanContactCode2: strings.TrimSpace(s[nsldsLoanContactCode2StartIndex-1 : (nsldsLoanContactCode2StartIndex-1)+nsldsLoanContactCode2Length]), // Field # 758

		NSLDSLoanGradeLevel2: strings.TrimSpace(s[nsldsLoanGradeLevel2StartIndex-1 : (nsldsLoanGradeLevel2StartIndex-1)+nsldsLoanGradeLevel2Length]), // Field # 759

		NSLDSLoanAdditionalUnsubsidizedFlag2: strings.TrimSpace(s[nsldsLoanAdditionalUnsubsidizedFlag2StartIndex-1 : (nsldsLoanAdditionalUnsubsidizedFlag2StartIndex-1)+nsldsLoanAdditionalUnsubsidizedFlag2Length]), // Field # 760

		NSLDSLoanCapitalizedInterestFlag2: strings.TrimSpace(s[nsldsLoanCapitalizedInterestFlag2StartIndex-1 : (nsldsLoanCapitalizedInterestFlag2StartIndex-1)+nsldsLoanCapitalizedInterestFlag2Length]), // Field # 761

		NSLDSLoanDisbursementAmount2: strings.TrimSpace(s[nsldsLoanDisbursementAmount2StartIndex-1 : (nsldsLoanDisbursementAmount2StartIndex-1)+nsldsLoanDisbursementAmount2Length]), // Field # 762

		NSLDSLoanDisbursementDate2: parseISIRDate(strings.TrimSpace(s[nsldsLoanDisbursementDate2StartIndex-1 : (nsldsLoanDisbursementDate2StartIndex-1)+nsldsLoanDisbursementDate2Length])), // Field # 763

		NSLDSLoanConfirmedLoanSubsidyStatus2: strings.TrimSpace(s[nsldsLoanConfirmedLoanSubsidyStatus2StartIndex-1 : (nsldsLoanConfirmedLoanSubsidyStatus2StartIndex-1)+nsldsLoanConfirmedLoanSubsidyStatus2Length]), // Field # 764

		NSLDSLoanSubsidyStatusDate2: parseISIRDate(strings.TrimSpace(s[nsldsLoanSubsidyStatusDate2StartIndex-1 : (nsldsLoanSubsidyStatusDate2StartIndex-1)+nsldsLoanSubsidyStatusDate2Length])), // Field # 765

		NSLDSLoanSequenceNumber3: strings.TrimSpace(s[nsldsLoanSequenceNumber3StartIndex-1 : (nsldsLoanSequenceNumber3StartIndex-1)+nsldsLoanSequenceNumber3Length]), // Field # 767

		NSLDSLoanDefaultedRecentIndicator3: strings.TrimSpace(s[nsldsLoanDefaultedRecentIndicator3StartIndex-1 : (nsldsLoanDefaultedRecentIndicator3StartIndex-1)+nsldsLoanDefaultedRecentIndicator3Length]), // Field # 768

		NSLDSLoanChangeFlag3: strings.TrimSpace(s[nsldsLoanChangeFlag3StartIndex-1 : (nsldsLoanChangeFlag3StartIndex-1)+nsldsLoanChangeFlag3Length]), // Field # 769

		NSLDSLoanTypeCode3: strings.TrimSpace(s[nsldsLoanTypeCode3StartIndex-1 : (nsldsLoanTypeCode3StartIndex-1)+nsldsLoanTypeCode3Length]), // Field # 770

		NSLDSLoanNetAmount3: strings.TrimSpace(s[nsldsLoanNetAmount3StartIndex-1 : (nsldsLoanNetAmount3StartIndex-1)+nsldsLoanNetAmount3Length]), // Field # 771

		NSLDSLoanCurrentStatusCode3: strings.TrimSpace(s[nsldsLoanCurrentStatusCode3StartIndex-1 : (nsldsLoanCurrentStatusCode3StartIndex-1)+nsldsLoanCurrentStatusCode3Length]), // Field # 772

		NSLDSLoanCurrentStatusDate3: parseISIRDate(strings.TrimSpace(s[nsldsLoanCurrentStatusDate3StartIndex-1 : (nsldsLoanCurrentStatusDate3StartIndex-1)+nsldsLoanCurrentStatusDate3Length])), // Field # 773

		NSLDSLoanOutstandingPrincipalBalance3: strings.TrimSpace(s[nsldsLoanOutstandingPrincipalBalance3StartIndex-1 : (nsldsLoanOutstandingPrincipalBalance3StartIndex-1)+nsldsLoanOutstandingPrincipalBalance3Length]), // Field # 774

		NSLDSLoanOutstandingPrincipalBalanceDate3: parseISIRDate(strings.TrimSpace(s[nsldsLoanOutstandingPrincipalBalanceDate3StartIndex-1 : (nsldsLoanOutstandingPrincipalBalanceDate3StartIndex-1)+nsldsLoanOutstandingPrincipalBalanceDate3Length])), // Field # 775

		NSLDSLoanPeriodBeginDate3: parseISIRDate(strings.TrimSpace(s[nsldsLoanPeriodBeginDate3StartIndex-1 : (nsldsLoanPeriodBeginDate3StartIndex-1)+nsldsLoanPeriodBeginDate3Length])), // Field # 776

		NSLDSLoanPeriodEndDate3: parseISIRDate(strings.TrimSpace(s[nsldsLoanPeriodEndDate3StartIndex-1 : (nsldsLoanPeriodEndDate3StartIndex-1)+nsldsLoanPeriodEndDate3Length])), // Field # 777

		NSLDSLoanGuarantyAgencyCode3: strings.TrimSpace(s[nsldsLoanGuarantyAgencyCode3StartIndex-1 : (nsldsLoanGuarantyAgencyCode3StartIndex-1)+nsldsLoanGuarantyAgencyCode3Length]), // Field # 778

		NSLDSLoanContactType3: strings.TrimSpace(s[nsldsLoanContactType3StartIndex-1 : (nsldsLoanContactType3StartIndex-1)+nsldsLoanContactType3Length]), // Field # 779

		NSLDSLoanSchoolCode3: strings.TrimSpace(s[nsldsLoanSchoolCode3StartIndex-1 : (nsldsLoanSchoolCode3StartIndex-1)+nsldsLoanSchoolCode3Length]), // Field # 780

		NSLDSLoanContactCode3: strings.TrimSpace(s[nsldsLoanContactCode3StartIndex-1 : (nsldsLoanContactCode3StartIndex-1)+nsldsLoanContactCode3Length]), // Field # 781

		NSLDSLoanGradeLevel3: strings.TrimSpace(s[nsldsLoanGradeLevel3StartIndex-1 : (nsldsLoanGradeLevel3StartIndex-1)+nsldsLoanGradeLevel3Length]), // Field # 782

		NSLDSLoanAdditionalUnsubsidizedFlag3: strings.TrimSpace(s[nsldsLoanAdditionalUnsubsidizedFlag3StartIndex-1 : (nsldsLoanAdditionalUnsubsidizedFlag3StartIndex-1)+nsldsLoanAdditionalUnsubsidizedFlag3Length]), // Field # 783

		NSLDSLoanCapitalizedInterestFlag3: strings.TrimSpace(s[nsldsLoanCapitalizedInterestFlag3StartIndex-1 : (nsldsLoanCapitalizedInterestFlag3StartIndex-1)+nsldsLoanCapitalizedInterestFlag3Length]), // Field # 784

		NSLDSLoanDisbursementAmount3: strings.TrimSpace(s[nsldsLoanDisbursementAmount3StartIndex-1 : (nsldsLoanDisbursementAmount3StartIndex-1)+nsldsLoanDisbursementAmount3Length]), // Field # 785

		NSLDSLoanDisbursementDate3: parseISIRDate(strings.TrimSpace(s[nsldsLoanDisbursementDate3StartIndex-1 : (nsldsLoanDisbursementDate3StartIndex-1)+nsldsLoanDisbursementDate3Length])), // Field # 786

		NSLDSLoanConfirmedLoanSubsidyStatus3: strings.TrimSpace(s[nsldsLoanConfirmedLoanSubsidyStatus3StartIndex-1 : (nsldsLoanConfirmedLoanSubsidyStatus3StartIndex-1)+nsldsLoanConfirmedLoanSubsidyStatus3Length]), // Field # 787

		NSLDSLoanSubsidyStatusDate3: parseISIRDate(strings.TrimSpace(s[nsldsLoanSubsidyStatusDate3StartIndex-1 : (nsldsLoanSubsidyStatusDate3StartIndex-1)+nsldsLoanSubsidyStatusDate3Length])), // Field # 788

		NSLDSLoanSequenceNumber4: strings.TrimSpace(s[nsldsLoanSequenceNumber4StartIndex-1 : (nsldsLoanSequenceNumber4StartIndex-1)+nsldsLoanSequenceNumber4Length]), // Field # 790

		NSLDSLoanDefaultedRecentIndicator4: strings.TrimSpace(s[nsldsLoanDefaultedRecentIndicator4StartIndex-1 : (nsldsLoanDefaultedRecentIndicator4StartIndex-1)+nsldsLoanDefaultedRecentIndicator4Length]), // Field # 791

		NSLDSLoanChangeFlag4: strings.TrimSpace(s[nsldsLoanChangeFlag4StartIndex-1 : (nsldsLoanChangeFlag4StartIndex-1)+nsldsLoanChangeFlag4Length]), // Field # 792

		NSLDSLoanTypeCode4: strings.TrimSpace(s[nsldsLoanTypeCode4StartIndex-1 : (nsldsLoanTypeCode4StartIndex-1)+nsldsLoanTypeCode4Length]), // Field # 793

		NSLDSLoanNetAmount4: strings.TrimSpace(s[nsldsLoanNetAmount4StartIndex-1 : (nsldsLoanNetAmount4StartIndex-1)+nsldsLoanNetAmount4Length]), // Field # 794

		NSLDSLoanCurrentStatusCode4: strings.TrimSpace(s[nsldsLoanCurrentStatusCode4StartIndex-1 : (nsldsLoanCurrentStatusCode4StartIndex-1)+nsldsLoanCurrentStatusCode4Length]), // Field # 795

		NSLDSLoanCurrentStatusDate4: parseISIRDate(strings.TrimSpace(s[nsldsLoanCurrentStatusDate4StartIndex-1 : (nsldsLoanCurrentStatusDate4StartIndex-1)+nsldsLoanCurrentStatusDate4Length])), // Field # 796

		NSLDSLoanOutstandingPrincipalBalance4: strings.TrimSpace(s[nsldsLoanOutstandingPrincipalBalance4StartIndex-1 : (nsldsLoanOutstandingPrincipalBalance4StartIndex-1)+nsldsLoanOutstandingPrincipalBalance4Length]), // Field # 797

		NSLDSLoanOutstandingPrincipalBalanceDate4: parseISIRDate(strings.TrimSpace(s[nsldsLoanOutstandingPrincipalBalanceDate4StartIndex-1 : (nsldsLoanOutstandingPrincipalBalanceDate4StartIndex-1)+nsldsLoanOutstandingPrincipalBalanceDate4Length])), // Field # 798

		NSLDSLoanPeriodBeginDate4: parseISIRDate(strings.TrimSpace(s[nsldsLoanPeriodBeginDate4StartIndex-1 : (nsldsLoanPeriodBeginDate4StartIndex-1)+nsldsLoanPeriodBeginDate4Length])), // Field # 799

		NSLDSLoanPeriodEndDate4: parseISIRDate(strings.TrimSpace(s[nsldsLoanPeriodEndDate4StartIndex-1 : (nsldsLoanPeriodEndDate4StartIndex-1)+nsldsLoanPeriodEndDate4Length])), // Field # 800

		NSLDSLoanGuarantyAgencyCode4: strings.TrimSpace(s[nsldsLoanGuarantyAgencyCode4StartIndex-1 : (nsldsLoanGuarantyAgencyCode4StartIndex-1)+nsldsLoanGuarantyAgencyCode4Length]), // Field # 801

		NSLDSLoanContactType4: strings.TrimSpace(s[nsldsLoanContactType4StartIndex-1 : (nsldsLoanContactType4StartIndex-1)+nsldsLoanContactType4Length]), // Field # 802

		NSLDSLoanSchoolCode4: strings.TrimSpace(s[nsldsLoanSchoolCode4StartIndex-1 : (nsldsLoanSchoolCode4StartIndex-1)+nsldsLoanSchoolCode4Length]), // Field # 803

		NSLDSLoanContactCode4: strings.TrimSpace(s[nsldsLoanContactCode4StartIndex-1 : (nsldsLoanContactCode4StartIndex-1)+nsldsLoanContactCode4Length]), // Field # 804

		NSLDSLoanGradeLevel4: strings.TrimSpace(s[nsldsLoanGradeLevel4StartIndex-1 : (nsldsLoanGradeLevel4StartIndex-1)+nsldsLoanGradeLevel4Length]), // Field # 805

		NSLDSLoanAdditionalUnsubsidizedFlag4: strings.TrimSpace(s[nsldsLoanAdditionalUnsubsidizedFlag4StartIndex-1 : (nsldsLoanAdditionalUnsubsidizedFlag4StartIndex-1)+nsldsLoanAdditionalUnsubsidizedFlag4Length]), // Field # 806

		NSLDSLoanCapitalizedInterestFlag4: strings.TrimSpace(s[nsldsLoanCapitalizedInterestFlag4StartIndex-1 : (nsldsLoanCapitalizedInterestFlag4StartIndex-1)+nsldsLoanCapitalizedInterestFlag4Length]), // Field # 807

		NSLDSLoanDisbursementAmount4: strings.TrimSpace(s[nsldsLoanDisbursementAmount4StartIndex-1 : (nsldsLoanDisbursementAmount4StartIndex-1)+nsldsLoanDisbursementAmount4Length]), // Field # 808

		NSLDSLoanDisbursementDate4: parseISIRDate(strings.TrimSpace(s[nsldsLoanDisbursementDate4StartIndex-1 : (nsldsLoanDisbursementDate4StartIndex-1)+nsldsLoanDisbursementDate4Length])), // Field # 809

		NSLDSLoanConfirmedLoanSubsidyStatus4: strings.TrimSpace(s[nsldsLoanConfirmedLoanSubsidyStatus4StartIndex-1 : (nsldsLoanConfirmedLoanSubsidyStatus4StartIndex-1)+nsldsLoanConfirmedLoanSubsidyStatus4Length]), // Field # 810

		NSLDSLoanSubsidyStatusDate4: parseISIRDate(strings.TrimSpace(s[nsldsLoanSubsidyStatusDate4StartIndex-1 : (nsldsLoanSubsidyStatusDate4StartIndex-1)+nsldsLoanSubsidyStatusDate4Length])), // Field # 811

		NSLDSLoanSequenceNumber5: strings.TrimSpace(s[nsldsLoanSequenceNumber5StartIndex-1 : (nsldsLoanSequenceNumber5StartIndex-1)+nsldsLoanSequenceNumber5Length]), // Field # 813

		NSLDSLoanDefaultedRecentIndicator5: strings.TrimSpace(s[nsldsLoanDefaultedRecentIndicator5StartIndex-1 : (nsldsLoanDefaultedRecentIndicator5StartIndex-1)+nsldsLoanDefaultedRecentIndicator5Length]), // Field # 814

		NSLDSLoanChangeFlag5: strings.TrimSpace(s[nsldsLoanChangeFlag5StartIndex-1 : (nsldsLoanChangeFlag5StartIndex-1)+nsldsLoanChangeFlag5Length]), // Field # 815

		NSLDSLoanTypeCode5: strings.TrimSpace(s[nsldsLoanTypeCode5StartIndex-1 : (nsldsLoanTypeCode5StartIndex-1)+nsldsLoanTypeCode5Length]), // Field # 816

		NSLDSLoanNetAmount5: strings.TrimSpace(s[nsldsLoanNetAmount5StartIndex-1 : (nsldsLoanNetAmount5StartIndex-1)+nsldsLoanNetAmount5Length]), // Field # 817

		NSLDSLoanCurrentStatusCode5: strings.TrimSpace(s[nsldsLoanCurrentStatusCode5StartIndex-1 : (nsldsLoanCurrentStatusCode5StartIndex-1)+nsldsLoanCurrentStatusCode5Length]), // Field # 818

		NSLDSLoanCurrentStatusDate5: parseISIRDate(strings.TrimSpace(s[nsldsLoanCurrentStatusDate5StartIndex-1 : (nsldsLoanCurrentStatusDate5StartIndex-1)+nsldsLoanCurrentStatusDate5Length])), // Field # 819

		NSLDSLoanOutstandingPrincipalBalance5: strings.TrimSpace(s[nsldsLoanOutstandingPrincipalBalance5StartIndex-1 : (nsldsLoanOutstandingPrincipalBalance5StartIndex-1)+nsldsLoanOutstandingPrincipalBalance5Length]), // Field # 820

		NSLDSLoanOutstandingPrincipalBalanceDate5: parseISIRDate(strings.TrimSpace(s[nsldsLoanOutstandingPrincipalBalanceDate5StartIndex-1 : (nsldsLoanOutstandingPrincipalBalanceDate5StartIndex-1)+nsldsLoanOutstandingPrincipalBalanceDate5Length])), // Field # 821

		NSLDSLoanPeriodBeginDate5: parseISIRDate(strings.TrimSpace(s[nsldsLoanPeriodBeginDate5StartIndex-1 : (nsldsLoanPeriodBeginDate5StartIndex-1)+nsldsLoanPeriodBeginDate5Length])), // Field # 822

		NSLDSLoanPeriodEndDate5: parseISIRDate(strings.TrimSpace(s[nsldsLoanPeriodEndDate5StartIndex-1 : (nsldsLoanPeriodEndDate5StartIndex-1)+nsldsLoanPeriodEndDate5Length])), // Field # 823

		NSLDSLoanGuarantyAgencyCode5: strings.TrimSpace(s[nsldsLoanGuarantyAgencyCode5StartIndex-1 : (nsldsLoanGuarantyAgencyCode5StartIndex-1)+nsldsLoanGuarantyAgencyCode5Length]), // Field # 824

		NSLDSLoanContactType5: strings.TrimSpace(s[nsldsLoanContactType5StartIndex-1 : (nsldsLoanContactType5StartIndex-1)+nsldsLoanContactType5Length]), // Field # 825

		NSLDSLoanSchoolCode5: strings.TrimSpace(s[nsldsLoanSchoolCode5StartIndex-1 : (nsldsLoanSchoolCode5StartIndex-1)+nsldsLoanSchoolCode5Length]), // Field # 826

		NSLDSLoanContactCode5: strings.TrimSpace(s[nsldsLoanContactCode5StartIndex-1 : (nsldsLoanContactCode5StartIndex-1)+nsldsLoanContactCode5Length]), // Field # 827

		NSLDSLoanGradeLevel5: strings.TrimSpace(s[nsldsLoanGradeLevel5StartIndex-1 : (nsldsLoanGradeLevel5StartIndex-1)+nsldsLoanGradeLevel5Length]), // Field # 828

		NSLDSLoanAdditionalUnsubsidizedFlag5: strings.TrimSpace(s[nsldsLoanAdditionalUnsubsidizedFlag5StartIndex-1 : (nsldsLoanAdditionalUnsubsidizedFlag5StartIndex-1)+nsldsLoanAdditionalUnsubsidizedFlag5Length]), // Field # 829

		NSLDSLoanCapitalizedInterestFlag5: strings.TrimSpace(s[nsldsLoanCapitalizedInterestFlag5StartIndex-1 : (nsldsLoanCapitalizedInterestFlag5StartIndex-1)+nsldsLoanCapitalizedInterestFlag5Length]), // Field # 830

		NSLDSLoanDisbursementAmount5: strings.TrimSpace(s[nsldsLoanDisbursementAmount5StartIndex-1 : (nsldsLoanDisbursementAmount5StartIndex-1)+nsldsLoanDisbursementAmount5Length]), // Field # 831

		NSLDSLoanDisbursementDate5: parseISIRDate(strings.TrimSpace(s[nsldsLoanDisbursementDate5StartIndex-1 : (nsldsLoanDisbursementDate5StartIndex-1)+nsldsLoanDisbursementDate5Length])), // Field # 832

		NSLDSLoanConfirmedLoanSubsidyStatus5: strings.TrimSpace(s[nsldsLoanConfirmedLoanSubsidyStatus5StartIndex-1 : (nsldsLoanConfirmedLoanSubsidyStatus5StartIndex-1)+nsldsLoanConfirmedLoanSubsidyStatus5Length]), // Field # 833

		NSLDSLoanSubsidyStatusDate5: parseISIRDate(strings.TrimSpace(s[nsldsLoanSubsidyStatusDate5StartIndex-1 : (nsldsLoanSubsidyStatusDate5StartIndex-1)+nsldsLoanSubsidyStatusDate5Length])), // Field # 834

		NSLDSLoanSequenceNumber6: strings.TrimSpace(s[nsldsLoanSequenceNumber6StartIndex-1 : (nsldsLoanSequenceNumber6StartIndex-1)+nsldsLoanSequenceNumber6Length]), // Field # 836

		NSLDSLoanDefaultedRecentIndicator6: strings.TrimSpace(s[nsldsLoanDefaultedRecentIndicator6StartIndex-1 : (nsldsLoanDefaultedRecentIndicator6StartIndex-1)+nsldsLoanDefaultedRecentIndicator6Length]), // Field # 837

		NSLDSLoanChangeFlag6: strings.TrimSpace(s[nsldsLoanChangeFlag6StartIndex-1 : (nsldsLoanChangeFlag6StartIndex-1)+nsldsLoanChangeFlag6Length]), // Field # 838

		NSLDSLoanTypeCode6: strings.TrimSpace(s[nsldsLoanTypeCode6StartIndex-1 : (nsldsLoanTypeCode6StartIndex-1)+nsldsLoanTypeCode6Length]), // Field # 839

		NSLDSLoanNetAmount6: strings.TrimSpace(s[nsldsLoanNetAmount6StartIndex-1 : (nsldsLoanNetAmount6StartIndex-1)+nsldsLoanNetAmount6Length]), // Field # 840

		NSLDSLoanCurrentStatusCode6: strings.TrimSpace(s[nsldsLoanCurrentStatusCode6StartIndex-1 : (nsldsLoanCurrentStatusCode6StartIndex-1)+nsldsLoanCurrentStatusCode6Length]), // Field # 841

		NSLDSLoanCurrentStatusDate6: parseISIRDate(strings.TrimSpace(s[nsldsLoanCurrentStatusDate6StartIndex-1 : (nsldsLoanCurrentStatusDate6StartIndex-1)+nsldsLoanCurrentStatusDate6Length])), // Field # 842

		NSLDSLoanOutstandingPrincipalBalance6: strings.TrimSpace(s[nsldsLoanOutstandingPrincipalBalance6StartIndex-1 : (nsldsLoanOutstandingPrincipalBalance6StartIndex-1)+nsldsLoanOutstandingPrincipalBalance6Length]), // Field # 843

		NSLDSLoanOutstandingPrincipalBalanceDate6: parseISIRDate(strings.TrimSpace(s[nsldsLoanOutstandingPrincipalBalanceDate6StartIndex-1 : (nsldsLoanOutstandingPrincipalBalanceDate6StartIndex-1)+nsldsLoanOutstandingPrincipalBalanceDate6Length])), // Field # 844

		NSLDSLoanPeriodBeginDate6: parseISIRDate(strings.TrimSpace(s[nsldsLoanPeriodBeginDate6StartIndex-1 : (nsldsLoanPeriodBeginDate6StartIndex-1)+nsldsLoanPeriodBeginDate6Length])), // Field # 845

		NSLDSLoanPeriodEndDate6: parseISIRDate(strings.TrimSpace(s[nsldsLoanPeriodEndDate6StartIndex-1 : (nsldsLoanPeriodEndDate6StartIndex-1)+nsldsLoanPeriodEndDate6Length])), // Field # 846

		NSLDSLoanGuarantyAgencyCode6: strings.TrimSpace(s[nsldsLoanGuarantyAgencyCode6StartIndex-1 : (nsldsLoanGuarantyAgencyCode6StartIndex-1)+nsldsLoanGuarantyAgencyCode6Length]), // Field # 847

		NSLDSLoanContactType6: strings.TrimSpace(s[nsldsLoanContactType6StartIndex-1 : (nsldsLoanContactType6StartIndex-1)+nsldsLoanContactType6Length]), // Field # 848

		NSLDSLoanSchoolCode6: strings.TrimSpace(s[nsldsLoanSchoolCode6StartIndex-1 : (nsldsLoanSchoolCode6StartIndex-1)+nsldsLoanSchoolCode6Length]), // Field # 849

		NSLDSLoanContactCode6: strings.TrimSpace(s[nsldsLoanContactCode6StartIndex-1 : (nsldsLoanContactCode6StartIndex-1)+nsldsLoanContactCode6Length]), // Field # 850

		NSLDSLoanGradeLevel6: strings.TrimSpace(s[nsldsLoanGradeLevel6StartIndex-1 : (nsldsLoanGradeLevel6StartIndex-1)+nsldsLoanGradeLevel6Length]), // Field # 851

		NSLDSLoanAdditionalUnsubsidizedFlag6: strings.TrimSpace(s[nsldsLoanAdditionalUnsubsidizedFlag6StartIndex-1 : (nsldsLoanAdditionalUnsubsidizedFlag6StartIndex-1)+nsldsLoanAdditionalUnsubsidizedFlag6Length]), // Field # 852

		NSLDSLoanCapitalizedInterestFlag6: strings.TrimSpace(s[nsldsLoanCapitalizedInterestFlag6StartIndex-1 : (nsldsLoanCapitalizedInterestFlag6StartIndex-1)+nsldsLoanCapitalizedInterestFlag6Length]), // Field # 853

		NSLDSLoanDisbursementAmount6: strings.TrimSpace(s[nsldsLoanDisbursementAmount6StartIndex-1 : (nsldsLoanDisbursementAmount6StartIndex-1)+nsldsLoanDisbursementAmount6Length]), // Field # 854

		NSLDSLoanDisbursementDate6: parseISIRDate(strings.TrimSpace(s[nsldsLoanDisbursementDate6StartIndex-1 : (nsldsLoanDisbursementDate6StartIndex-1)+nsldsLoanDisbursementDate6Length])), // Field # 855

		NSLDSLoanConfirmedLoanSubsidyStatus6: strings.TrimSpace(s[nsldsLoanConfirmedLoanSubsidyStatus6StartIndex-1 : (nsldsLoanConfirmedLoanSubsidyStatus6StartIndex-1)+nsldsLoanConfirmedLoanSubsidyStatus6Length]), // Field # 856

		NSLDSLoanSubsidyStatusDate6: parseISIRDate(strings.TrimSpace(s[nsldsLoanSubsidyStatusDate6StartIndex-1 : (nsldsLoanSubsidyStatusDate6StartIndex-1)+nsldsLoanSubsidyStatusDate6Length])), // Field # 857

		FTILabelStart: strings.TrimSpace(s[ftiLabelStartStartIndex-1 : (ftiLabelStartStartIndex-1)+ftiLabelStartLength]), // Field # 861

		StudentFTIMReturnedTaxYear: strings.TrimSpace(s[studentFTIMReturnedTaxYearStartIndex-1 : (studentFTIMReturnedTaxYearStartIndex-1)+studentFTIMReturnedTaxYearLength]), // Field # 862

		StudentFTIMFilingStatusCode: strings.TrimSpace(s[studentFTIMFilingStatusCodeStartIndex-1 : (studentFTIMFilingStatusCodeStartIndex-1)+studentFTIMFilingStatusCodeLength]), // Field # 863

		StudentFTIMAdjustedGrossIncome: strings.TrimSpace(s[studentFTIMAdjustedGrossIncomeStartIndex-1 : (studentFTIMAdjustedGrossIncomeStartIndex-1)+studentFTIMAdjustedGrossIncomeLength]), // Field # 864

		StudentFTIMNumberOfExemptions: strings.TrimSpace(s[studentFTIMNumberOfExemptionsStartIndex-1 : (studentFTIMNumberOfExemptionsStartIndex-1)+studentFTIMNumberOfExemptionsLength]), // Field # 865

		StudentFTIMNumberOfDependents: strings.TrimSpace(s[studentFTIMNumberOfDependentsStartIndex-1 : (studentFTIMNumberOfDependentsStartIndex-1)+studentFTIMNumberOfDependentsLength]), // Field # 866

		StudentFTIMTotalIncomeEarnedAmount: strings.TrimSpace(s[studentFTIMTotalIncomeEarnedAmountStartIndex-1 : (studentFTIMTotalIncomeEarnedAmountStartIndex-1)+studentFTIMTotalIncomeEarnedAmountLength]), // Field # 867

		StudentFTIMIncomeTaxPaid: strings.TrimSpace(s[studentFTIMIncomeTaxPaidStartIndex-1 : (studentFTIMIncomeTaxPaidStartIndex-1)+studentFTIMIncomeTaxPaidLength]), // Field # 868

		StudentFTIMEducationCredits: strings.TrimSpace(s[studentFTIMEducationCreditsStartIndex-1 : (studentFTIMEducationCreditsStartIndex-1)+studentFTIMEducationCreditsLength]), // Field # 869

		StudentFTIMUntaxedIRADistributions: strings.TrimSpace(s[studentFTIMUntaxedIRADistributionsStartIndex-1 : (studentFTIMUntaxedIRADistributionsStartIndex-1)+studentFTIMUntaxedIRADistributionsLength]), // Field # 870

		StudentFTIMIRADeductibleAndPayments: strings.TrimSpace(s[studentFTIMIRADeductibleAndPaymentsStartIndex-1 : (studentFTIMIRADeductibleAndPaymentsStartIndex-1)+studentFTIMIRADeductibleAndPaymentsLength]), // Field # 871

		StudentFTIMTaxExemptInterest: strings.TrimSpace(s[studentFTIMTaxExemptInterestStartIndex-1 : (studentFTIMTaxExemptInterestStartIndex-1)+studentFTIMTaxExemptInterestLength]), // Field # 872

		StudentFTIMUntaxedPensionsAmount: strings.TrimSpace(s[studentFTIMUntaxedPensionsAmountStartIndex-1 : (studentFTIMUntaxedPensionsAmountStartIndex-1)+studentFTIMUntaxedPensionsAmountLength]), // Field # 873

		StudentFTIMScheduleCNetProfitLoss: strings.TrimSpace(s[studentFTIMScheduleCNetProfitLossStartIndex-1 : (studentFTIMScheduleCNetProfitLossStartIndex-1)+studentFTIMScheduleCNetProfitLossLength]), // Field # 874

		StudentFTIMScheduleAIndicator: strings.TrimSpace(s[studentFTIMScheduleAIndicatorStartIndex-1 : (studentFTIMScheduleAIndicatorStartIndex-1)+studentFTIMScheduleAIndicatorLength]), // Field # 875

		StudentFTIMScheduleBIndicator: strings.TrimSpace(s[studentFTIMScheduleBIndicatorStartIndex-1 : (studentFTIMScheduleBIndicatorStartIndex-1)+studentFTIMScheduleBIndicatorLength]), // Field # 876

		StudentFTIMScheduleDIndicator: strings.TrimSpace(s[studentFTIMScheduleDIndicatorStartIndex-1 : (studentFTIMScheduleDIndicatorStartIndex-1)+studentFTIMScheduleDIndicatorLength]), // Field # 877

		StudentFTIMScheduleEIndicator: strings.TrimSpace(s[studentFTIMScheduleEIndicatorStartIndex-1 : (studentFTIMScheduleEIndicatorStartIndex-1)+studentFTIMScheduleEIndicatorLength]), // Field # 878

		StudentFTIMScheduleFIndicator: strings.TrimSpace(s[studentFTIMScheduleFIndicatorStartIndex-1 : (studentFTIMScheduleFIndicatorStartIndex-1)+studentFTIMScheduleFIndicatorLength]), // Field # 879

		StudentFTIMScheduleHIndicator: strings.TrimSpace(s[studentFTIMScheduleHIndicatorStartIndex-1 : (studentFTIMScheduleHIndicatorStartIndex-1)+studentFTIMScheduleHIndicatorLength]), // Field # 880

		StudentFTIMIRSResponseCode: strings.TrimSpace(s[studentFTIMIRSResponseCodeStartIndex-1 : (studentFTIMIRSResponseCodeStartIndex-1)+studentFTIMIRSResponseCodeLength]), // Field # 881

		StudentFTIMSpouseReturnedTaxYear: strings.TrimSpace(s[studentFTIMSpouseReturnedTaxYearStartIndex-1 : (studentFTIMSpouseReturnedTaxYearStartIndex-1)+studentFTIMSpouseReturnedTaxYearLength]), // Field # 882

		StudentFTIMSpouseFilingStatusCode: strings.TrimSpace(s[studentFTIMSpouseFilingStatusCodeStartIndex-1 : (studentFTIMSpouseFilingStatusCodeStartIndex-1)+studentFTIMSpouseFilingStatusCodeLength]), // Field # 883

		StudentFTIMSpouseAdjustedGrossIncome: strings.TrimSpace(s[studentFTIMSpouseAdjustedGrossIncomeStartIndex-1 : (studentFTIMSpouseAdjustedGrossIncomeStartIndex-1)+studentFTIMSpouseAdjustedGrossIncomeLength]), // Field # 884

		StudentFTIMSpouseNumberOfExemptions: strings.TrimSpace(s[studentFTIMSpouseNumberOfExemptionsStartIndex-1 : (studentFTIMSpouseNumberOfExemptionsStartIndex-1)+studentFTIMSpouseNumberOfExemptionsLength]), // Field # 885

		StudentFTIMSpouseNumberOfDependents: strings.TrimSpace(s[studentFTIMSpouseNumberOfDependentsStartIndex-1 : (studentFTIMSpouseNumberOfDependentsStartIndex-1)+studentFTIMSpouseNumberOfDependentsLength]), // Field # 886

		StudentFTIMSpouseTotalIncomeEarnedAmount: strings.TrimSpace(s[studentFTIMSpouseTotalIncomeEarnedAmountStartIndex-1 : (studentFTIMSpouseTotalIncomeEarnedAmountStartIndex-1)+studentFTIMSpouseTotalIncomeEarnedAmountLength]), // Field # 887

		StudentFTIMSpouseIncomeTaxPaid: strings.TrimSpace(s[studentFTIMSpouseIncomeTaxPaidStartIndex-1 : (studentFTIMSpouseIncomeTaxPaidStartIndex-1)+studentFTIMSpouseIncomeTaxPaidLength]), // Field # 888

		StudentFTIMSpouseEducationCredits: strings.TrimSpace(s[studentFTIMSpouseEducationCreditsStartIndex-1 : (studentFTIMSpouseEducationCreditsStartIndex-1)+studentFTIMSpouseEducationCreditsLength]), // Field # 889

		StudentFTIMSpouseUntaxedIRADistributions: strings.TrimSpace(s[studentFTIMSpouseUntaxedIRADistributionsStartIndex-1 : (studentFTIMSpouseUntaxedIRADistributionsStartIndex-1)+studentFTIMSpouseUntaxedIRADistributionsLength]), // Field # 890

		StudentFTIMSpouseIRADeductibleAndPayments: strings.TrimSpace(s[studentFTIMSpouseIRADeductibleAndPaymentsStartIndex-1 : (studentFTIMSpouseIRADeductibleAndPaymentsStartIndex-1)+studentFTIMSpouseIRADeductibleAndPaymentsLength]), // Field # 891

		StudentFTIMSpouseTaxExemptInterest: strings.TrimSpace(s[studentFTIMSpouseTaxExemptInterestStartIndex-1 : (studentFTIMSpouseTaxExemptInterestStartIndex-1)+studentFTIMSpouseTaxExemptInterestLength]), // Field # 892

		StudentFTIMSpouseUntaxedPensionsAmount: strings.TrimSpace(s[studentFTIMSpouseUntaxedPensionsAmountStartIndex-1 : (studentFTIMSpouseUntaxedPensionsAmountStartIndex-1)+studentFTIMSpouseUntaxedPensionsAmountLength]), // Field # 893

		StudentFTIMSpouseScheduleCNetProfitLoss: strings.TrimSpace(s[studentFTIMSpouseScheduleCNetProfitLossStartIndex-1 : (studentFTIMSpouseScheduleCNetProfitLossStartIndex-1)+studentFTIMSpouseScheduleCNetProfitLossLength]), // Field # 894

		StudentFTIMSpouseScheduleAIndicator: strings.TrimSpace(s[studentFTIMSpouseScheduleAIndicatorStartIndex-1 : (studentFTIMSpouseScheduleAIndicatorStartIndex-1)+studentFTIMSpouseScheduleAIndicatorLength]), // Field # 895

		StudentFTIMSpouseScheduleBIndicator: strings.TrimSpace(s[studentFTIMSpouseScheduleBIndicatorStartIndex-1 : (studentFTIMSpouseScheduleBIndicatorStartIndex-1)+studentFTIMSpouseScheduleBIndicatorLength]), // Field # 896

		StudentFTIMSpouseScheduleDIndicator: strings.TrimSpace(s[studentFTIMSpouseScheduleDIndicatorStartIndex-1 : (studentFTIMSpouseScheduleDIndicatorStartIndex-1)+studentFTIMSpouseScheduleDIndicatorLength]), // Field # 897

		StudentFTIMSpouseScheduleEIndicator: strings.TrimSpace(s[studentFTIMSpouseScheduleEIndicatorStartIndex-1 : (studentFTIMSpouseScheduleEIndicatorStartIndex-1)+studentFTIMSpouseScheduleEIndicatorLength]), // Field # 898

		StudentFTIMSpouseScheduleFIndicator: strings.TrimSpace(s[studentFTIMSpouseScheduleFIndicatorStartIndex-1 : (studentFTIMSpouseScheduleFIndicatorStartIndex-1)+studentFTIMSpouseScheduleFIndicatorLength]), // Field # 899

		StudentFTIMSpouseScheduleHIndicator: strings.TrimSpace(s[studentFTIMSpouseScheduleHIndicatorStartIndex-1 : (studentFTIMSpouseScheduleHIndicatorStartIndex-1)+studentFTIMSpouseScheduleHIndicatorLength]), // Field # 900

		StudentFTIMSpouseIRSResponseCode: strings.TrimSpace(s[studentFTIMSpouseIRSResponseCodeStartIndex-1 : (studentFTIMSpouseIRSResponseCodeStartIndex-1)+studentFTIMSpouseIRSResponseCodeLength]), // Field # 901

		ParentFTIMReturnedTaxYear: strings.TrimSpace(s[parentFTIMReturnedTaxYearStartIndex-1 : (parentFTIMReturnedTaxYearStartIndex-1)+parentFTIMReturnedTaxYearLength]), // Field # 902

		ParentFTIMFilingStatusCode: strings.TrimSpace(s[parentFTIMFilingStatusCodeStartIndex-1 : (parentFTIMFilingStatusCodeStartIndex-1)+parentFTIMFilingStatusCodeLength]), // Field # 903

		ParentFTIMAdjustedGrossIncome: strings.TrimSpace(s[parentFTIMAdjustedGrossIncomeStartIndex-1 : (parentFTIMAdjustedGrossIncomeStartIndex-1)+parentFTIMAdjustedGrossIncomeLength]), // Field # 904

		ParentFTIMNumberOfExemptions: strings.TrimSpace(s[parentFTIMNumberOfExemptionsStartIndex-1 : (parentFTIMNumberOfExemptionsStartIndex-1)+parentFTIMNumberOfExemptionsLength]), // Field # 905

		ParentFTIMNumberOfDependents: strings.TrimSpace(s[parentFTIMNumberOfDependentsStartIndex-1 : (parentFTIMNumberOfDependentsStartIndex-1)+parentFTIMNumberOfDependentsLength]), // Field # 906

		ParentFTIMTotalIncomeEarnedAmount: strings.TrimSpace(s[parentFTIMTotalIncomeEarnedAmountStartIndex-1 : (parentFTIMTotalIncomeEarnedAmountStartIndex-1)+parentFTIMTotalIncomeEarnedAmountLength]), // Field # 907

		ParentFTIMIncomeTaxPaid: strings.TrimSpace(s[parentFTIMIncomeTaxPaidStartIndex-1 : (parentFTIMIncomeTaxPaidStartIndex-1)+parentFTIMIncomeTaxPaidLength]), // Field # 908

		ParentFTIMEducationCredits: strings.TrimSpace(s[parentFTIMEducationCreditsStartIndex-1 : (parentFTIMEducationCreditsStartIndex-1)+parentFTIMEducationCreditsLength]), // Field # 909

		ParentFTIMUntaxedIRADistributions: strings.TrimSpace(s[parentFTIMUntaxedIRADistributionsStartIndex-1 : (parentFTIMUntaxedIRADistributionsStartIndex-1)+parentFTIMUntaxedIRADistributionsLength]), // Field # 910

		ParentFTIMIRADeductibleAndPayments: strings.TrimSpace(s[parentFTIMIRADeductibleAndPaymentsStartIndex-1 : (parentFTIMIRADeductibleAndPaymentsStartIndex-1)+parentFTIMIRADeductibleAndPaymentsLength]), // Field # 911

		ParentFTIMTaxExemptInterest: strings.TrimSpace(s[parentFTIMTaxExemptInterestStartIndex-1 : (parentFTIMTaxExemptInterestStartIndex-1)+parentFTIMTaxExemptInterestLength]), // Field # 912

		ParentFTIMUntaxedPensionsAmount: strings.TrimSpace(s[parentFTIMUntaxedPensionsAmountStartIndex-1 : (parentFTIMUntaxedPensionsAmountStartIndex-1)+parentFTIMUntaxedPensionsAmountLength]), // Field # 913

		ParentFTIMScheduleCNetProfitLoss: strings.TrimSpace(s[parentFTIMScheduleCNetProfitLossStartIndex-1 : (parentFTIMScheduleCNetProfitLossStartIndex-1)+parentFTIMScheduleCNetProfitLossLength]), // Field # 914

		ParentFTIMScheduleAIndicator: strings.TrimSpace(s[parentFTIMScheduleAIndicatorStartIndex-1 : (parentFTIMScheduleAIndicatorStartIndex-1)+parentFTIMScheduleAIndicatorLength]), // Field # 915

		ParentFTIMScheduleBIndicator: strings.TrimSpace(s[parentFTIMScheduleBIndicatorStartIndex-1 : (parentFTIMScheduleBIndicatorStartIndex-1)+parentFTIMScheduleBIndicatorLength]), // Field # 916

		ParentFTIMScheduleDIndicator: strings.TrimSpace(s[parentFTIMScheduleDIndicatorStartIndex-1 : (parentFTIMScheduleDIndicatorStartIndex-1)+parentFTIMScheduleDIndicatorLength]), // Field # 917

		ParentFTIMScheduleEIndicator: strings.TrimSpace(s[parentFTIMScheduleEIndicatorStartIndex-1 : (parentFTIMScheduleEIndicatorStartIndex-1)+parentFTIMScheduleEIndicatorLength]), // Field # 918

		ParentFTIMScheduleFIndicator: strings.TrimSpace(s[parentFTIMScheduleFIndicatorStartIndex-1 : (parentFTIMScheduleFIndicatorStartIndex-1)+parentFTIMScheduleFIndicatorLength]), // Field # 919

		ParentFTIMScheduleHIndicator: strings.TrimSpace(s[parentFTIMScheduleHIndicatorStartIndex-1 : (parentFTIMScheduleHIndicatorStartIndex-1)+parentFTIMScheduleHIndicatorLength]), // Field # 920

		ParentFTIMIRSResponseCode: strings.TrimSpace(s[parentFTIMIRSResponseCodeStartIndex-1 : (parentFTIMIRSResponseCodeStartIndex-1)+parentFTIMIRSResponseCodeLength]), // Field # 921

		ParentFTIMSpouseReturnedTaxYear: strings.TrimSpace(s[parentFTIMSpouseReturnedTaxYearStartIndex-1 : (parentFTIMSpouseReturnedTaxYearStartIndex-1)+parentFTIMSpouseReturnedTaxYearLength]), // Field # 922

		ParentFTIMSpouseFilingStatusCode: strings.TrimSpace(s[parentFTIMSpouseFilingStatusCodeStartIndex-1 : (parentFTIMSpouseFilingStatusCodeStartIndex-1)+parentFTIMSpouseFilingStatusCodeLength]), // Field # 923

		ParentFTIMSpouseAdjustedGrossIncome: strings.TrimSpace(s[parentFTIMSpouseAdjustedGrossIncomeStartIndex-1 : (parentFTIMSpouseAdjustedGrossIncomeStartIndex-1)+parentFTIMSpouseAdjustedGrossIncomeLength]), // Field # 924

		ParentFTIMSpouseNumberOfExemptions: strings.TrimSpace(s[parentFTIMSpouseNumberOfExemptionsStartIndex-1 : (parentFTIMSpouseNumberOfExemptionsStartIndex-1)+parentFTIMSpouseNumberOfExemptionsLength]), // Field # 925

		ParentFTIMSpouseNumberOfDependents: strings.TrimSpace(s[parentFTIMSpouseNumberOfDependentsStartIndex-1 : (parentFTIMSpouseNumberOfDependentsStartIndex-1)+parentFTIMSpouseNumberOfDependentsLength]), // Field # 926

		ParentFTIMSpouseTotalIncomeEarnedAmount: strings.TrimSpace(s[parentFTIMSpouseTotalIncomeEarnedAmountStartIndex-1 : (parentFTIMSpouseTotalIncomeEarnedAmountStartIndex-1)+parentFTIMSpouseTotalIncomeEarnedAmountLength]), // Field # 927

		ParentFTIMSpouseIncomeTaxPaid: strings.TrimSpace(s[parentFTIMSpouseIncomeTaxPaidStartIndex-1 : (parentFTIMSpouseIncomeTaxPaidStartIndex-1)+parentFTIMSpouseIncomeTaxPaidLength]), // Field # 928

		ParentFTIMSpouseEducationCredits: strings.TrimSpace(s[parentFTIMSpouseEducationCreditsStartIndex-1 : (parentFTIMSpouseEducationCreditsStartIndex-1)+parentFTIMSpouseEducationCreditsLength]), // Field # 929

		ParentFTIMSpouseUntaxedIRADistributions: strings.TrimSpace(s[parentFTIMSpouseUntaxedIRADistributionsStartIndex-1 : (parentFTIMSpouseUntaxedIRADistributionsStartIndex-1)+parentFTIMSpouseUntaxedIRADistributionsLength]), // Field # 930

		ParentFTIMSpouseIRADeductibleAndPayments: strings.TrimSpace(s[parentFTIMSpouseIRADeductibleAndPaymentsStartIndex-1 : (parentFTIMSpouseIRADeductibleAndPaymentsStartIndex-1)+parentFTIMSpouseIRADeductibleAndPaymentsLength]), // Field # 931

		ParentFTIMSpouseTaxExemptInterest: strings.TrimSpace(s[parentFTIMSpouseTaxExemptInterestStartIndex-1 : (parentFTIMSpouseTaxExemptInterestStartIndex-1)+parentFTIMSpouseTaxExemptInterestLength]), // Field # 932

		ParentFTIMSpouseUntaxedPensionsAmount: strings.TrimSpace(s[parentFTIMSpouseUntaxedPensionsAmountStartIndex-1 : (parentFTIMSpouseUntaxedPensionsAmountStartIndex-1)+parentFTIMSpouseUntaxedPensionsAmountLength]), // Field # 933

		ParentFTIMSpouseScheduleCNetProfitLoss: strings.TrimSpace(s[parentFTIMSpouseScheduleCNetProfitLossStartIndex-1 : (parentFTIMSpouseScheduleCNetProfitLossStartIndex-1)+parentFTIMSpouseScheduleCNetProfitLossLength]), // Field # 934

		ParentFTIMSpouseScheduleAIndicator: strings.TrimSpace(s[parentFTIMSpouseScheduleAIndicatorStartIndex-1 : (parentFTIMSpouseScheduleAIndicatorStartIndex-1)+parentFTIMSpouseScheduleAIndicatorLength]), // Field # 935

		ParentFTIMSpouseScheduleBIndicator: strings.TrimSpace(s[parentFTIMSpouseScheduleBIndicatorStartIndex-1 : (parentFTIMSpouseScheduleBIndicatorStartIndex-1)+parentFTIMSpouseScheduleBIndicatorLength]), // Field # 936

		ParentFTIMSpouseScheduleDIndicator: strings.TrimSpace(s[parentFTIMSpouseScheduleDIndicatorStartIndex-1 : (parentFTIMSpouseScheduleDIndicatorStartIndex-1)+parentFTIMSpouseScheduleDIndicatorLength]), // Field # 937

		ParentFTIMSpouseScheduleEIndicator: strings.TrimSpace(s[parentFTIMSpouseScheduleEIndicatorStartIndex-1 : (parentFTIMSpouseScheduleEIndicatorStartIndex-1)+parentFTIMSpouseScheduleEIndicatorLength]), // Field # 938

		ParentFTIMSpouseScheduleFIndicator: strings.TrimSpace(s[parentFTIMSpouseScheduleFIndicatorStartIndex-1 : (parentFTIMSpouseScheduleFIndicatorStartIndex-1)+parentFTIMSpouseScheduleFIndicatorLength]), // Field # 939

		ParentFTIMSpouseScheduleHIndicator: strings.TrimSpace(s[parentFTIMSpouseScheduleHIndicatorStartIndex-1 : (parentFTIMSpouseScheduleHIndicatorStartIndex-1)+parentFTIMSpouseScheduleHIndicatorLength]), // Field # 940

		ParentFTIMSpouseIRSResponseCode: strings.TrimSpace(s[parentFTIMSpouseIRSResponseCodeStartIndex-1 : (parentFTIMSpouseIRSResponseCodeStartIndex-1)+parentFTIMSpouseIRSResponseCodeLength]), // Field # 941

		FTILabelEnd: strings.TrimSpace(s[ftiLabelEndStartIndex-1 : (ftiLabelEndStartIndex-1)+ftiLabelEndLength]), // Field # 942

		StudentTotalIncome: strings.TrimSpace(s[studentTotalIncomeStartIndex-1 : (studentTotalIncomeStartIndex-1)+studentTotalIncomeLength]), // Field # 944

		ParentTotalIncome: strings.TrimSpace(s[parentTotalIncomeStartIndex-1 : (parentTotalIncomeStartIndex-1)+parentTotalIncomeLength]), // Field # 945

		FISAPTotalIncome: strings.TrimSpace(s[fisapTotalIncomeStartIndex-1 : (fisapTotalIncomeStartIndex-1)+fisapTotalIncomeLength]), // Field # 946

	}

	return r, nil
}

func parseISIRDate(s string) time.Time {
	parsedDate, err := time.Parse(isirDateLayout, s)

	if err != nil {
		return time.Time{}
	}

	return parsedDate
}
