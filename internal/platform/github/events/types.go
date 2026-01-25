package events

import (
	"time"

	"github.com/IteratorInnovator/git-gram/internal/platform/telegram"
)

// InlineKeyboardButton is an alias for the telegram package type.
type InlineKeyboardButton = telegram.InlineKeyboardButton

// BranchProtectionConfiguration represents a branch protection configuration event.
type BranchProtectionConfiguration struct {
	Action string `json:"action"`

	Repository struct {
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		HTMLURL  string `json:"html_url"`
	} `json:"repository"`

	Sender struct {
		Login   string `json:"login"`
		HTMLURL string `json:"html_url"`
	} `json:"sender"`
}

// CreateEvent represents a branch or tag creation event.
type CreateEvent struct {
	Ref     string `json:"ref"`
	RefType string `json:"ref_type"`

	Repository struct {
		Name     string    `json:"name"`
		FullName string    `json:"full_name"`
		HTMLURL  string    `json:"html_url"`
		PushedAt time.Time `json:"pushed_at"`
	} `json:"repository"`

	Sender struct {
		Login   string `json:"login"`
		HTMLURL string `json:"html_url"`
	} `json:"sender"`
}

// BranchProtectionRuleEvent represents a branch protection rule event.
type BranchProtectionRuleEvent struct {
	Action string `json:"action"`
	Rule   struct {
		ID                                       int           `json:"id"`
		Name                                     string        `json:"name"`
		CreatedAt                                time.Time     `json:"created_at"`
		UpdatedAt                                time.Time     `json:"updated_at"`
		PullRequestReviewsEnforcementLevel       string        `json:"pull_request_reviews_enforcement_level"`
		RequiredApprovingReviewCount             int           `json:"required_approving_review_count"`
		DismissStaleReviewsOnPush                bool          `json:"dismiss_stale_reviews_on_push"`
		RequireCodeOwnerReview                   bool          `json:"require_code_owner_review"`
		AuthorizedDismissalActorsOnly            bool          `json:"authorized_dismissal_actors_only"`
		IgnoreApprovalsFromContributors          bool          `json:"ignore_approvals_from_contributors"`
		RequiredStatusChecks                     []interface{} `json:"required_status_checks"`
		RequiredStatusChecksEnforcementLevel     string        `json:"required_status_checks_enforcement_level"`
		StrictRequiredStatusChecksPolicy         bool          `json:"strict_required_status_checks_policy"`
		SignatureRequirementEnforcementLevel     string        `json:"signature_requirement_enforcement_level"`
		LinearHistoryRequirementEnforcementLevel string        `json:"linear_history_requirement_enforcement_level"`
		AdminEnforced                            bool          `json:"admin_enforced"`
		CreateProtected                          bool          `json:"create_protected"`
		AllowForcePushesEnforcementLevel         string        `json:"allow_force_pushes_enforcement_level"`
		AllowDeletionsEnforcementLevel           string        `json:"allow_deletions_enforcement_level"`
		MergeQueueEnforcementLevel               string        `json:"merge_queue_enforcement_level"`
		RequiredDeploymentsEnforcementLevel      string        `json:"required_deployments_enforcement_level"`
		RequiredConversationResolutionLevel      string        `json:"required_conversation_resolution_level"`
		AuthorizedActorsOnly                     bool          `json:"authorized_actors_only"`
		AuthorizedActorNames                     []interface{} `json:"authorized_actor_names"`
		RequireLastPushApproval                  bool          `json:"require_last_push_approval"`
		LockBranchEnforcementLevel               string        `json:"lock_branch_enforcement_level"`
	} `json:"rule"`
	Changes struct {
		AdminEnforced struct {
			From interface{} `json:"from"`
		} `json:"admin_enforced"`
		AuthorizedActorNames struct {
			From interface{} `json:"from"`
		} `json:"authorized_actor_names"`
		AuthorizedActorsOnly struct {
			From interface{} `json:"from"`
		} `json:"authorized_actors_only"`
		AuthorizedDismissalActorsOnly struct {
			From interface{} `json:"from"`
		} `json:"authorized_dismissal_actors_only"`
		LinearHistoryRequirementEnforcementLevel struct {
			From interface{} `json:"from"`
		} `json:"linear_history_requirement_enforcement_level"`
		LockBranchEnforcementLevel struct {
			From interface{} `json:"from"`
		} `json:"lock_branch_enforcement_level"`
		LockAllowsForkSync struct {
			From interface{} `json:"from"`
		} `json:"lock_allows_fork_sync"`
		PullRequestReviewsEnforcementLevel struct {
			From interface{} `json:"from"`
		} `json:"pull_request_reviews_enforcement_level"`
		RequireLastPushApproval struct {
			From interface{} `json:"from"`
		} `json:"require_last_push_approval"`
		RequiredStatusChecks struct {
			From interface{} `json:"from"`
		} `json:"required_status_checks"`
		RequiredStatusChecksEnforcementLevel struct {
			From interface{} `json:"from"`
		} `json:"required_status_checks_enforcement_level"`
		SignatureRequirementEnforcementLevel struct {
			From string `json:"from"`
		} `json:"signature_requirement_enforcement_level"`
	} `json:"changes"`

	Repository struct {
		Name        string        `json:"name"`
		FullName    string        `json:"full_name"`
		HTMLURL     string        `json:"html_url"`
		Description interface{}   `json:"description"`
		Homepage    interface{}   `json:"homepage"`
		Topics      []interface{} `json:"topics"`
	} `json:"repository"`

	Sender struct {
		Login   string `json:"login"`
		HTMLURL string `json:"html_url"`
	} `json:"sender"`
}

// DeleteEvent represents a branch or tag deletion event.
type DeleteEvent struct {
	Ref     string `json:"ref"`
	RefType string `json:"ref_type"`

	Repository struct {
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		HTMLURL  string `json:"html_url"`
	} `json:"repository"`

	Sender struct {
		Login   string `json:"login"`
		HTMLURL string `json:"html_url"`
	} `json:"sender"`
}

// PushEvent represents a push event.
type PushEvent struct {
	Ref     string `json:"ref"`
	Compare string `json:"compare"`

	Repository struct {
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		HTMLURL  string `json:"html_url"`
		PushedAt int64  `json:"pushed_at"`
	} `json:"repository"`

	Pusher struct {
		Name string `json:"name"`
	} `json:"pusher"`

	Sender struct {
		Login   string `json:"login"`
		HTMLURL string `json:"html_url"`
	} `json:"sender"`

	Commits []struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	} `json:"commits"`

	HeadCommit struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
	} `json:"head_commit"`
}

// RepositoryEvent represents a repository event.
type RepositoryEvent struct {
	Action string `json:"action"`

	Changes struct {
		Repository struct {
			Name struct {
				From string `json:"from"`
			} `json:"name"`
		} `json:"repository"`
		DefaultBranch struct {
			From string `json:"from"`
		} `json:"default_branch"`
		Description struct {
			From *string `json:"from"`
		} `json:"description"`
		Homepage struct {
			From *string `json:"from"`
		} `json:"homepage"`
		Topics struct {
			From *[]string `json:"from"`
		} `json:"topics"`
	} `json:"changes"`

	Repository struct {
		Name          string     `json:"name"`
		FullName      string     `json:"full_name"`
		HTMLURL       string     `json:"html_url"`
		CreatedAt     time.Time  `json:"created_at"`
		UpdatedAt     time.Time  `json:"updated_at"`
		DefaultBranch string     `json:"default_branch"`
		Description   *string    `json:"description"`
		Homepage      *string    `json:"homepage"`
		Topics        *[]string  `json:"topics"`
	} `json:"repository"`

	Sender struct {
		Login   string `json:"login"`
		HTMLURL string `json:"html_url"`
	} `json:"sender"`
}
