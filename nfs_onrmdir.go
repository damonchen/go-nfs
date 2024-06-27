package nfs

import (
	"context"
)

func onRmDir(ctx context.Context, w *Response, userHandle Handler) error {
	return onRemove(ctx, w, userHandle)
}
