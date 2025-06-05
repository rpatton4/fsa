// Contains constant values in use throughout the module and for use by clients of the module
// Where applicable the names, types and values will match those used by COD and other Ed systems
//
// For a naming convention, the constant will be along the lines of <category><value meaning>, for example
// AwardYear2526 indicating it is a constant referring to the Award Year for 2025-2026
package fsaconstants

type AwardYear int

// Indicates an invalid or unknown or unset Award Year
const AwardYearUnknown AwardYear = -1

// Indicates the Award Year 2024-2025 using the same value as Ed does within the ISIR data stream
const AwardYear2425 AwardYear = 5

// Indicates the Award Year 2025-2026 using the same value as Ed does within the ISIR data stream
const AwardYear2526 AwardYear = 6

// Indicates the Award Year 2026-2027 using the same value as Ed does within the ISIR data stream
const AwardYear2627 AwardYear = 7
