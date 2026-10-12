# Security Policy

The KubeVirtBMC maintainers take the security of the project seriously. This document describes which versions receive security fixes and how to report a vulnerability.

## Supported Versions

Security fixes are made on `main` and backported to the release branch of the latest minor version (see [Branching Strategy](CONTRIBUTING.md#branching-strategy)). Older minor versions do not receive security fixes; please upgrade to the latest release.

| Version                   | Supported          |
| ------------------------- | ------------------ |
| Latest minor (e.g. 0.9.x) | :white_check_mark: |
| Older minor versions      | :x:                |

## Reporting a Vulnerability

**Please do not report security vulnerabilities through public GitHub issues, pull requests, or Discord.**

Report vulnerabilities privately using GitHub's [private vulnerability reporting](https://github.com/kubevirtbmc/kubevirtbmc/security/advisories/new). This sends the report only to the project [maintainers](MAINTAINERS.md).

Please include as much of the following as you can:

- The type of issue (e.g. authentication bypass, privilege escalation, credential exposure, denial of service)
- The affected component (controller, virtBMC IPMI server, Redfish API, etc.) and version or commit
- Full steps to reproduce, including any configuration or manifests needed
- Proof-of-concept or exploit code, if available
- The impact of the issue and how an attacker might use it

## Disclosure Process

1. **Acknowledgement**: A maintainer will acknowledge your report within **3 business days**.
2. **Triage**: The maintainers will confirm the issue, determine affected versions, and assess severity (using [CVSS](https://www.first.org/cvss/)) within **10 business days**, keeping you informed of progress.
3. **Fix**: A fix is developed in a private fork via a GitHub Security Advisory. The reporter is invited to collaborate on and validate the fix.
4. **Release**: Patched releases are published for all supported versions. A CVE is requested through GitHub where appropriate.
5. **Disclosure**: The security advisory is published once fixes are available. We aim to disclose within **90 days** of the initial report, and will coordinate the timing with the reporter.

Reporters will be credited in the advisory unless they prefer to remain anonymous.

## Security Updates

Security advisories are published on the project's [GitHub Security Advisories](https://github.com/kubevirtbmc/kubevirtbmc/security/advisories) page and noted in the release notes of the patched versions.
