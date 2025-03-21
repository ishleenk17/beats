---
mapped_pages:
  - https://www.elastic.co/guide/en/beats/devguide/current/pr-review.html
---

# Pull request review guidelines [pr-review]

Every change made to Beats must be held to a high standard, and while the responsibility for quality in a pull request ultimately lies with the author, Beats team members have the responsibility as reviewers to verify during their review process. Where this document is unclear or inappropriate let common sense and consensus override it.


## Code Style [_code_style]

Everyone’s got an opinion on style. To avoid spending time on this issue we rely almost exclusively on `go fmt` and [hound](https://houndci.com/) to police style. If neither of these tools complain the code is almost certainly fine. There may be exceptions to this, but they should be extremely rare. Only override the judgement of these tools in the most unusual of situations.


## Flaky Tests [_flaky_tests]

As software projects grow so does the complexity of their test cases and with that the probability of some tests becoming *flaky*. It is everyone’s responsibility to handle flaky tests. If you notice a pull request build failing for a reason that is unrelated to the pushed code follow the procedure below:

1. Create an issue using the "Flaky Test" github issue template with the "Flaky Test" label attached.
2. Create a PR to mute or fix the flaky test.
3. Merge that PR and rebase off of it before continuing with the normal PR process for your original PR.

