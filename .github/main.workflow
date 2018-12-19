workflow "New workflow" {
  on = "push"
  resolves = ["Docker Tag"]
}

action "Docker Registry" {
  uses = "actions/docker/login@76ff57a"
  secrets = ["GITHUB_TOKEN"]
}

action "Docker Tag" {
  uses = "actions/docker/tag@76ff57a"
  needs = ["Docker Registry"]
}
