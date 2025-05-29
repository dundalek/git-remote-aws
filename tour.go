package main

func tour() {
	// git-remote-aws stores encrypted git bundles in S3
	// bundle list content is a file of newline separated hashes
	// uses dynamodb for locking

	// entrypoint to dispatch remote helper commands invoked by git
	gitHelper()
	{
		// just prints "push" and "fetch"
		capabilities()

		// lists refs, locks for reading
		list()
		{
			getBundles()
			// then prints last bundle hash
		}

		// checks if remote bundles include latest local hash
		// shells to git to create bundle
		// encrypts and putObject to s3
		// apends hash to bundle list and unlocks
		// also removes old bundle metadata
		push()

		// goes through list of bundles, fetches the missing ones
		// decrypts into tmp, shells to git to unbundle
		fetch()
	}
}
