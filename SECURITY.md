# Security Policy

## Supported Versions

Security fixes are made on `main` and backported to the latest two minor releases.

| Version | Supported          |
| ------- | ------------------ |
| 0.11.x  | :white_check_mark: |
| 0.10.x  | :white_check_mark: |
| < 0.10  | :x:                |

## Reporting a Vulnerability

**Please do not report security issues through public GitHub issues or pull requests.**

Report privately via:

- **GitHub:** the [Security tab](https://github.com/kubevirtbmc/kubevirtbmc/security) → **"Report a vulnerability"** (preferred)
- **Email:** `<security@kubevirtbmc.io>`

Include a description, affected version(s), component (`virtbmc` or `virtbmc-controller`), and steps to reproduce if possible.

## What to Expect

| Stage                         | Target timeline                       |
| ----------------------------- | ------------------------------------- |
| Acknowledgement of the report | Within 3 business days                |
| Initial assessment            | Within 7 business days                |
| Status updates                | At least every 14 days until resolved |
| Fix for critical/high issues  | Aim for within 30 days                |

- **If accepted:** we'll fix it, release patches for supported versions, publish a GitHub Security Advisory (with a CVE where appropriate), and credit you unless you prefer otherwise
- **If declined:** we'll explain why

We follow coordinated disclosure and will agree on a disclosure date with you.

## Security Advisories

To date, no security vulnerabilities have been reported for KubeVirtBMC. Future advisories will be published on the [Security Advisories](https://github.com/kubevirtbmc/kubevirtbmc/security/advisories) page.
