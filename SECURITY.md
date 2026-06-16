# Security policy

costfuse can take actions that stop cloud resources and disable billing, so we
take security seriously.

## Reporting a vulnerability

Please do **not** open a public issue for security problems. Instead, report
them privately using GitHub's "Report a vulnerability" feature under the
repository's Security tab, or email the maintainer directly.

We'll acknowledge your report as quickly as we can and keep you updated on the
fix.

## Scope

Of particular interest: anything that could cause costfuse to stop a protected
resource, take a real action while in dry-run mode, or act on a signal it
should not trust.
