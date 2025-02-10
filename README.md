# iDeclare

ideclare is an experimental research project for a CI/CD system that:

* uses OCI Images as the composable units of process, with declarative inputs and outputs of artifacts, parameters and
  metadata (SBOM, Logs, Builds, Packages, Config etc)
* provides tooling for adding, reading, and updating state in the image.
* provides a declarative pipeline language
* a runtime:
    * Local with container execution on Docker or Kubernetes
    * Docker-in-docker for a fully encapsulated runtime
    * A kubernetes controller
* a state specification (yaml/json) which:
    * captures and supports re-entrance
    * supports partial execution/caching (running only that which needs to execute)

# Jobs

Like any CI/CD system, we isolate work to be done into composable tasks, or for consistancy with typical runtimes (such
as kubernetes)
we call them 'jobs'

`jobs`:

* are largely black boxes, the ideclare runtime tries to integrate with them only on their exposed api.
* expected to declare their inputs
    * in a way that the ideclare runtime can capture identity of the inputs (see [Capturing inputs](#capturing-inputs))
    * with type declarations to make pipelining simpler and clearer.
* expected to declare their outputs (with types)

# Capturing inputs

Inputs into a job, which is an OCI container runtime can be:

* A simple scalar: `memory=<scalar>`
* A value string: `name="my-job"`
* A structured value: `config='{ "key1": "value1" }'` - which could be considered a string
* A string representing a resource: `https://www.example.com/some-resource` or `docker.io/ubuntu`
* A mounted volume

For the purposes of iDeclare, it's important that every input can be forensically traced and identified without
expecting the job to calculate and capture digests/details.

## Current proposal

* Inputs can be declared in a metadata layer (tbd: mime-type) on the job image.
* Inputs are captured by the runtime as part of the job-execution record produced whenever a task is run (regardless of
  completion)
* The runtime will capture (in the job-execution record) verbatim:
    * simple scalars
    * value strings
    * structured values
* Mutable resources with access to immutable references digests: 
  * The runtime will request the current immutable reference for known resource locators, where possible, and rewrite the
    resource locator with the immutable reference. 
  * If it's not possible, the job will not execute
  * Example: `docker.io/library/ubuntu:20.04` -> `docker.io/library/ubuntu:20.04@sha256:8e5c4f0285ecbb4ead070431d29b576a530d3166df73ec44affc1cd27555141b`
  * The input spec can ask for the resource to be mounted. If so, the resolved digest from the host is captured.
* Resources without accessible digests 
  * for now, these are not supported
  * In the future we could add types such as a direct download, and form a digest before mounting the volume.
    * The problem with this is it does not promote repeatable builds.
