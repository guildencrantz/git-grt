package main

func Fetch(args []string) {
	// curl --verbose -u mhenkel:ZvGPregVQNkD --digest http://gerrit.dev.returnpath.net/a/changes/auth_policy_publisher\~master\~Ib0a47bffa7ddee956d980df3f799ebafb9ae1f2a\?o\=ALL_REVISIONS

	cmd := NewGrtCmd("GET", change_endpoint)
	cmd.rawForm = "o=ALL_REVISIONS"
}

