package message_templates

const SingleCommitPush string = `🔔 *New Push to %s*

[%s](%s) pushed 1 commit to ` + "`%s`" + ` at %s

Commit: ` + "`%s`" + `
%s`

const MultipleCommitsPush string = `🔔 *New Push to %s*

[%s](%s) pushed %d commits to ` + "`%s`" + ` at %s

Latest Commit: ` + "`%s`" + `
%s`


