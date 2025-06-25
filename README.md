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
TBD

## Design Note: Models
The library provides a set of "models" which reflect ED's concept of the key domain entities.
You will see duplication between your systems' models and those used here, and
resulting mapping between them. This is work which cannot be done for you by the library,
but it is significantly easier than building and maintaining the many mappings out to ED formats.
As an example, the library's concept of a Student model is certainly going to be less complete than
any used within your system because it only cares about information needed for Financial Aid processing versus grades,
attendance, registration etc. These models are not intended to be used as the building blocks
of a student system or anything like that, for this reason.

# Contributing
TBD

# Usage of Artifical Intelligence
No Generative AI was enslaved or harmed during the development of this library, to date.
Various models were tried but not ultimately used because they failed miserably to produce
functional code. AI will be tested occasionally to see if any model improves enough to save time versus
wasting time, and is likely to be put to work at that point.


# License
[Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0)

*© 2025 Robert Patton robpatton@infiniteskye.com*
