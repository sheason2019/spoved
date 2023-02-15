package initial

import "context"

func Initial(ctx context.Context) {
	root, err := initRootUser(ctx)
	if err != nil {
		panic(err)
	}

	err = initSpoved(ctx, root)
	if err != nil {
		panic(err)
	}

	err = initSpovedFe(ctx, root)
	if err != nil {
		panic(err)
	}
}
