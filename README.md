FHIR Go and Ember Generator
===========================

The *fhir-golang-generator* repository is a fork of the HL7 FHIR DSTU2 source code, with additions and modifications to support the generation of FHIR code for Go and Ember. This repository is only needed if you intend to make changes to the code generation logic. In that case, the re-generated code should also be committed in the corresponding [fhir](https://github.com/intervention-engine/fhir) or [ember-fhir-adapter](https://github.com/intervention-engine/ember-fhir-adapter) repository.

Install Git
-----------

The *fhir-golang-generator* repository is hosted on GitHub. The Git toolchain is needed to clone the repository locally. At the time this documentation was written, Git 2.7.0 was the latest available release.

To install Git, follow the instructions found in the [Git Book - Installing Git chapter](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).

As an alternative to manual installation, many Mac OS X developers use [Homebrew](http://brew.sh/) to install common development tools. If you prefer to install the latest Git release using Homebrew, execute the following commands:

```
$ brew update
$ brew install git
```

Install Go
----------

This repository generates Go code -- but does not generate it using the conventional Go formatting guidelines. The Go toolchain contains a utility for reformatting code that should be used on the generated Go code. If you're only concerned with the Ember code, you don't need to install Go. Intervention Engine requires Go 1.5 or above. At the time this documentation was written, Go 1.5.3 was the latest available release.

To install Go, follow the instructions found in the [Go Programming Language Getting Started guide](http://golang.org/doc/install).

If you prefer to install the latest Go release using Homebrew, execute the following commands:

```
$ brew update
$ brew install go
```

Be sure to follow the advice in the [Go Programming Language Getting Started guide](http://golang.org/doc/install) regarding setting up environment variables (e.g., $GOROOT, $GOPATH) and your path.

Install Java SDK
----------------

The FHIR build chain is based on the Java programming language, which requires Oracle's Java SDK. At the time this documentation was written, Java SE 8u74 was the latest available release.

To install the Java SE 8 SDK, visit the [Java SE Downloads page](http://www.oracle.com/technetwork/java/javase/downloads/index.html).

If you prefer to install the latest Git release using Homebrew, you'll need to install it from the cask by execute the following commands:

```
$ brew update
$ brew tap caskroom/cask
$ brew cask install java
```

Install Apache Ant
------------------

The FHIR build chain uses the Apache Ant build system. At the time this documentation was written, Ant 1.9.6 was the latest available release.

To install Apache Ant, download the binary from the [Ant Binary Distributions page](http://ant.apache.org/bindownload.cgi) and follow the instructions found in the [Ant Manual: Installing Ant](http://ant.apache.org/manual/install.html#installing).

If you prefer to install the latest Git release using Homebrew, execute the following commands:

```
$ brew update
$ brew install ant
```

Clone the fhir-golang-generator Repository
------------------------------------------

We recommend you create an *intervention-engine* folder within your favorite development location and clone the *fhir-golang-generator* repository there. For this documentation, we'll assume that "your favorite development location" is `~/development`.

```
$ cd ~/development/intervention-engine
$ git clone https://github.com/intervention-engine/fhir-golang-generator.git
```

Run the FHIR Publisher
----------------------

The Go and Ember code generation steps are part of the publisher. The easiest way to run the publisher is to run `publish.bat` (windows) or `publish.sh` (OSX/Linux).

```
$ cd ~/development/intervention-engine/fhir-golang-generator
$ ./publish.sh
```

Note that the build may not finish successfully. This is a known problem (usually with `QuestionnaireAnswers.java`), but does *not* affect the code generation process. As long as you see log output similar to the following, you've successfully generated the code:

```
[java] Produce go Reference Implementation                                        0.396  22sec  568MB
[java] Produce ember Reference Implementation                                     0.866  23sec  582MB
```

Format and Copy the Generated Go Code
-------------------------------------

The Go code is generated to the relative path `implementations/go/base/app`, but it isn't generated using the conventional Go formatting guidelines. Before copying the generated code to another repository (such as [fhir](https://github.com/intervention-engine/fhir)), you should reformat it using `gofmt`:

```
gofmt -w ~/development/intervention-engine/fhir-golang-generator/implementations/go/base/app/
```

After the generated code has been properly formatted, you can copy it to the Intervention Engine [fhir](https://github.com/intervention-engine/fhir) repo to test the changes (and ultimately commit them if they are successful). The following command recursively copies the generated code to the *fhir* repo. Be sure to inspect the changes carefully to ensure that the newly generated code is correct.

```
cp -r ~/development/intervention-engine/fhir-golang-generator/implementations/go/base/app/* $GOPATH/src/github.com/intervention-engine/fhir/
```

Copy the Generated Ember Data Adapter Code
------------------------------------------

The Ember code is generated to the relative path `implementations/ember/base/app`. The following command recursively copies the generated code to the [ember-fhir-adapter](https://github.com/intervention-engine/ember-fhir-adapter) repo. Be sure to inspect the changes carefully to ensure that the newly generated code is correct.

```
cp -r ~/development/intervention-engine/fhir-golang-generator/implementations/ember/base/app/* ~/development/intervention-engine/ember-fhir-adapter/app/
```
