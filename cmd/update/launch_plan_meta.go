package update

import (
	"context"
	"fmt"

	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flytectl/clierrors"
	"github.com/flyteorg/flytectl/cmd/config"
	cmdCore "github.com/flyteorg/flytectl/cmd/core"
)

const (
	updateLPMetaShort = "Updates the launch plan metadata"
	updateLPMetaLong  = `
Update the description on the launch plan:
::

 flytectl update launchplan -p flytesnacks -d development  core.advanced.merge_sort.merge_sort --description "Mergesort example"

Archiving launch plan named entity is not supported and would throw an error:
::

 flytectl update launchplan -p flytesnacks -d development  core.advanced.merge_sort.merge_sort --archive

Activating launch plan named entity would be a noop:
::

 flytectl update launchplan -p flytesnacks -d development  core.advanced.merge_sort.merge_sort --activate

Usage
`
)

func getUpdateLPMetaFunc(namedEntityConfig *NamedEntityConfig) func(ctx context.Context, args []string, cmdCtx cmdCore.CommandContext) error {
	return func(ctx context.Context, args []string, cmdCtx cmdCore.CommandContext) error {
		project := config.GetConfig().Project
		domain := config.GetConfig().Domain
		if len(args) != 1 {
			return fmt.Errorf(clierrors.ErrLPNotPassed)
		}
		name := args[0]
		err := namedEntityConfig.UpdateNamedEntity(ctx, name, project, domain, core.ResourceType_LAUNCH_PLAN, cmdCtx)
		if err != nil {
			return fmt.Errorf(clierrors.ErrFailedLPUpdate, name, err)
		}
		fmt.Printf("updated metadata successfully on %v", name)
		return nil
	}
}
