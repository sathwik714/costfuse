# Contributing to costfuse

Thanks for your interest. costfuse is in early development, so the best way to
help right now is to try it, file issues, and pick up roadmap items.

## Getting started

```bash
git clone https://github.com/sathwik714/costfuse.git
cd costfuse
make test   # run the tests
make run    # run a dry-run pass
```

## Workflow

1. Open or pick up an issue so work isn't duplicated.
2. Create a branch: `git checkout -b feat/short-description`.
3. Make your change. Keep packages focused (one job each).
4. Run `make lint` and `make test` before pushing.
5. Open a pull request describing what changed and why.

## Commit messages

We use Conventional Commits, e.g. `feat: add gcp collector`,
`fix: correct egress cost weight`, `docs: clarify dry-run behaviour`.

## Code style

Standard Go formatting — run `gofmt` (most editors do this on save).
