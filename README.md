# iDeclare

ideclare is an experimental research project for a CI/CD system that:

* uses OCI Images as the composable units of process, with declarative inputs and outputs of artifacts, parameters and metadata (SBOM, Logs, Builds, Packages, Config etc)
* provides tooling for adding, reading, and updating state in the image.
* provides a declarative pipeline language
* a runtime:
  * Local with container execution on Docker or Kubernetes
  * Docker-in-docker for a fully encapsulated runtime
  * A kubernetes controller
* a state specification (yaml/json) which:
  * captures and supports re-entrance
  * supports partial execution/caching (running only that which needs to execute)
