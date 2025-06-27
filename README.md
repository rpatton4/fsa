# US Federal Student Aid Module for Go
## Table of Contents
- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [AI](#usage-of-artifical-intelligence)
- [License](#license)

# Overview
This library is intended to assist with reading and writing the many different data formats
used by the US Department of Education's (ED) Federal Student Aid systems.  The target systems are:
- Common Origination & Disbursement (COD)
- G5

A library for this purpose is expected to be useful because

- The formats are quite lengthy, for example the ISIR file has over 1,000 fields
- The data is often provided in multiple actual formats, such as both CSV and fixed length
- The data expectations within the formats can be quite flexibile, for example fields which are only populated based on the values in other fields
- The formats change every year, and potentially multiple times per year
- Many formats have been around for decades, and the documentation for them is inconsistent due to many authors over time

And perhaps most importantly **ED doesn't provide libraries for this purpose, so every single system which interacts with ED needs to reinvent this wheel**
and it's not a competitive advantage to do so, it's just grunt work.



# Installation
This is a library for the Go language, and as such it uses the standard module mechanisms.
The module is `github.com/rpatton4/fsa` so you would add the import to your go.mod file
with `go get github.com/rpatton4/fsa`. After that you should see something like the
following in your go.mod file: `github.com/rpatton4/fsa v0.1.6-alpha // indirect`

## Prerequisites
- Go 1.24.x

The library will likely work fine with earlier versions given Go's backwards compatibility, at least to 1.21.x,
but it is only tested on current releases.

# Usage
The entry point for the library is through the fsaservices package, or `github.com/rpatton4/fsa/pkg/fsaservices`
There will be a dedicated set of service functions per high level type of document or data, which can be used to
read from a stream to parse the data into models, or take data from models and produce a resulting stream
in the requested ED format. As different formats are implemented an example of usage for each
will be included here.

## Usage Note: Award Years
TBD

## ISIR Service Functions
The fsaservices.ParseISIRStream function is the entry point for converting a data stream with ISIR data
to the models.  The function and related logic will determine the source format from the data
stream itself, for example which format to use based on award year. Calling into the function is
very straightforward: pass it a stream of ISIR data as received from ED and receive back a slice
of fsamodels.ISIRecord or an error. See [Error Handling](##error-handling) for more information on that
subject. Sample code (with an obviously fake ISIR for brevity) is:
```aiignore
package main

import (
    "github.com/rpatton4/fsa/pkg/fsaservices"
	"github.com/rpatton4/fsa/pkg/fsamodels"
	"strings"
)

func main() {
    s := "replace with valid ISIR line"
    r, err := fsaservices.ParseISIRStream(strings.NewReader(s))
    
    if err != nil {
        fmt.Println(err.Error())
        if len(err.UpstreamErrors()) >0 {
            fmt.Println("Upstream errors: ")
            for i, e := range err.UpstreamErrors {
                fmt.Println(e.Error())
            }
        }
        return
    }
    
    for i, isir := range r {
        fmt.Println(isir[i].TransactionUUID)
    }
}
```

## Design Note: Models
The library provides a set of "models" which reflect ED's concept of the key domain entities.
You will see duplication between your systems' models and those used here, and
resulting mapping between them. This is work which cannot be done for you by the library,
but it is significantly easier than building and maintaining the many mappings out to ED formats.
As an example, the library's concept of a Student model is certainly going to be less complete than
any used within your system because it only cares about information needed for Financial Aid processing versus grades,
attendance, registration etc. These models are not intended to be used as the building blocks
of a student system or anything like that, for this reason.

## Error Handling
The `fsaerrors.Error` type is used to represent errors occurring while transforming data between the models 
and the various formats.  These errors indicate either runtime errors for invalid data such as a field being set to a string
which is longer than what COD will accept, or build time errors for missing or out of date configuration such 
as trying to import a CommonRecord from COD for an award year which is not yet supported by the library.

The custom error type fulfills the Go error interface so it can be used in the standard way, but has been
created to provide some additional functionality:

- An error code, which is numeric identifier for the specific error.  This is intended to be used for providing error message in languages other than English
- Embedded upstream errors, which occur during at least two scenarios:
  - When a validation process is being run on a high level model, so that all validation errors throughout the model can be determined and returned at once
  - When an error occurs several layers deep in the code, to provide more context for the error in something like a stack trace

# Contributing
Anyone is welcome to fork the repository and submit issues for suggestions on improvements, but I'd like to
hold off on PRs for now until the library has implement CommonRecord and at least one or two message classes.

Please feel free to contact me with questions or suggestions which are not marketing or sales related.

# Usage of Artifical Intelligence
No Generative AI was enslaved or harmed during the development of this library, to date.
Various models were tried but not ultimately used because they failed miserably to produce
functional code. AI will be tested occasionally to see if any model gets off acid, and is likely to be put to work at that point.


# License
[Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0)

*© 2025 Robert Patton robpatton@infiniteskye.com*
