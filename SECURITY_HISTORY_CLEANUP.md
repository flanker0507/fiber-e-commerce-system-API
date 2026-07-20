# Manual Git History Cleanup Plan

The current working tree no longer tracks `.env` or contains hardcoded application credentials, but deleting or replacing a secret in a new commit does not remove it from older Git objects.

## Sensitive paths found in history

The history audit identified these paths and secret types:

- `.env` — database password and environment credentials
- `auth/service.go` — JWT signing secret
- `config/middtrans.go` — Midtrans server/client key configuration
- `payment/service.go` — Midtrans server/client key configuration

Do not paste any exposed value into an issue, pull request, chat, shell history, or tracked file. Rotate the database password, JWT secret, and both Midtrans keys in their respective systems. History rewriting alone does not revoke a credential.

## Before rewriting

Coordinate a maintenance window and stop all pushes. Install [`git-filter-repo`](https://github.com/newren/git-filter-repo), then work from a fresh clone. In PowerShell, replace the repository URL and destination names as appropriate:

```powershell
$RepositoryUrl = '<repository-url>'
git clone $RepositoryUrl fiber-e-commerce-system-API-clean
git clone --mirror $RepositoryUrl fiber-e-commerce-system-API-backup.git
Set-Location fiber-e-commerce-system-API-clean
git bundle create ..\fiber-e-commerce-system-API-before-filter.bundle --all
```

Confirm that both the mirror clone and bundle exist before continuing. Store them securely because they still contain the exposed credentials.

## Prepare replacements outside the repository

Create `..\secret-replacements.txt` outside the clone. Add one `git-filter-repo` replacement rule per exposed value, replacing the bracketed labels locally with the exact old values:

```text
literal:<OLD_JWT_SECRET>==>REMOVED_JWT_SECRET
literal:<OLD_MIDTRANS_SERVER_KEY>==>REMOVED_MIDTRANS_SERVER_KEY
literal:<OLD_MIDTRANS_CLIENT_KEY>==>REMOVED_MIDTRANS_CLIENT_KEY
```

Do not commit this file. Delete it securely after the rewrite and verification. The database password is removed with every historical `.env` blob; if that password was also copied elsewhere, add it to this replacement file too.

## Rewrite and verify locally

Run the rewrite only after reviewing the backup and replacement file:

```powershell
git filter-repo --path .env --invert-paths --replace-text ..\secret-replacements.txt
git log --all -- .env
git fsck --full
go test ./...
go vet ./...
go build ./...
```

`git log --all -- .env` should produce no commits. Separately search all rewritten refs for each old value using a local secret-scanning tool or patterns file; do not print matches into shared logs.

## Publish the rewritten history

Review the rewritten branches and tags before changing the remote. `git-filter-repo` normally removes the `origin` remote as a safety measure, so add it back and then run the destructive pushes manually:

```powershell
git remote add origin $RepositoryUrl
git push --force --all origin
git push --force --tags origin
```

These commands replace shared history. Protect the backup, notify every collaborator, and expire relevant CI caches and release artifacts that may contain old Git objects.

Afterward, every collaborator must delete and re-clone their local repository, or fetch and hard-reset every local branch to the rewritten remote. Re-cloning is strongly recommended because merging an old clone can reintroduce the removed history.
