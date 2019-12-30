[![Build Status](https://travis-ci.org/kcobos/CC-SIGA.svg?branch=master)](https://travis-ci.org/kcobos/CC-SIGA)
[![CircleCI](https://circleci.com/gh/kcobos/CC-SIGA.svg?style=svg)](https://circleci.com/gh/kcobos/CC-SIGA)
[![codecov](https://codecov.io/gh/kcobos/CC-SIGA/branch/master/graph/badge.svg)](https://codecov.io/gh/kcobos/CC-SIGA)
[![GPL Licence](https://badges.frapsoft.com/os/gpl/gpl.png?v=103)](https://opensource.org/licenses/GPL-3.0/)

# CC-SIGA
SIGA (Sistema Integral de Gestión de Aparcamientos (Integral Parking Management System)) is a system designed for reserved parkings in public area like authorities, ambulances or, its main aim, disabled parkings. It tries to "centralize" any kind of information about these parkings (location, parking type, occupation status (free, occupied, bad occupied...)) and show it as easy as possible to users.

The information about occupation status about parking is provided by sensors disposed in parkings attached to this project. By this way, the system has actual states of all parkings.


* [Infrastructure overview](./docs/index.md#infrastructure-overview)
* [Specifying the infrastructure](./docs/index.md#specifying-the-infrastructure)


## Test
This project has been tested to ensure a function of it self. In this case, this project has ben tested completely in [Travis CI](https://travis-ci.org/kcobos/CC-SIGA) and only a test in [Circle CI](https://circleci.com/gh/kcobos/CC-SIGA) due to its free plans.

Starting with [Circle CI](https://circleci.com/gh/kcobos/CC-SIGA), free plan has only 1 job so only we can test one language and one version of that. We have chosen GoLang language to test due to Circle CI doesn't have the last version of Python language (3.8). This test doesn't upload data to code coverage ([CodeCov](https://codecov.io/gh/kcobos/CC-SIGA)) so, the configuration is very simple:

``` yml
version: 2
jobs:
  build:
    docker: 
      - image: circleci/golang:1.13
    steps: 
      - checkout 
      - run: make goTest
```

Due to this project is a open source project, in [Travis CI](https://travis-ci.org/kcobos/CC-SIGA) we don't have job restrictions and we can test all languages and versions. By default, [Travis](https://docs.travis-ci.com/user/build-matrix/) uses a matrix system to build jobs to test the code. Here, we have two languages for the time being so we couldn't use build matrix system without excluding and comparing each combination because configuration complexity.

We decided to create job per job, [building stages manually](https://docs.travis-ci.com/user/build-stages/) to test all languages (Go an Python) and to test all versions we need. The configuration is:

``` yml
jobs:
  include:
    - stage: test
      name: "Go 1.13 tests"
      language: go
      go:
        - "1.13"
      script:
        - make goTest
      after_success:
        - make goCoverage
        - bash <(curl -s https://codecov.io/bash) -cF go
```
create one stage per language and version. After test success, test coverage will upload to [CodeCov](https://codecov.io/gh/kcobos/CC-SIGA) explaining the language.

In GoLang we have chosen almost all version, from 1.3 to 1.13 or lasted released, and in Python we have chosen only Python3, from 3.4 to 3.8-dev, because Python2 will retire on 2020 (in one day!!).

Due to this project is open source, we don't have to do anymore configuration but add the repository in Travis or CircleCI and add the above configuration file.

## Building tools
To start, we have chosen a common build tool between the both languages (GoLang and Python) because we don't have any dependencies right now and only run tests and coverage. [Make](https://www.gnu.org/software/make/) is the chosen because is a free and open source tool installed in all distributions.

We want to separate every microservice and each one must to have each build tool. GoLang doesn't need an external tool but [Tusk](https://github.com/rliebz/tusk) could ease the job. Python need an external tool and it could be [DoIt](https://pydoit.org/) because it's also free and open source.

buildtool: Makefile

In [Make](Makefile) we have created two rules for each microservice: one for testing and other for coverage except in `Users` that we have other to install dependencies but there isn't now. 

To test all microservices of a language, there is another rule to make it easier. The same for coverage.

To sum all, the buildtool is [make](https://www.gnu.org/software/make/) and the rules are:
 * `<microserviceName>Test`: test that microservice.
 * `<microserviceName>Coverage`: test and coverage that microservice.
 * `<languageName>Test`: test that language (Go or Python).
 * `<languageName>Coverage`: test and coverage that language (Go or Python).
