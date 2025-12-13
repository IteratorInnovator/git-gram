package github

type installationResponse struct {
    ID      int64 `json:"id"`
    Account struct {
        Login string `json:"login"`
    } `json:"account"`
}