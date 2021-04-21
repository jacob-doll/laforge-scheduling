// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/gen0cide/laforge/ent/command"
	"github.com/gen0cide/laforge/ent/disk"
	"github.com/gen0cide/laforge/ent/ginfilemiddleware"
	"github.com/gen0cide/laforge/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commandFields := schema.Command{}.Fields()
	_ = commandFields
	// commandDescCooldown is the schema descriptor for cooldown field.
	commandDescCooldown := commandFields[7].Descriptor()
	// command.CooldownValidator is a validator for the "cooldown" field. It is called by the builders before save.
	command.CooldownValidator = commandDescCooldown.Validators[0].(func(int) error)
	// commandDescTimeout is the schema descriptor for timeout field.
	commandDescTimeout := commandFields[8].Descriptor()
	// command.TimeoutValidator is a validator for the "timeout" field. It is called by the builders before save.
	command.TimeoutValidator = commandDescTimeout.Validators[0].(func(int) error)
	diskFields := schema.Disk{}.Fields()
	_ = diskFields
	// diskDescSize is the schema descriptor for size field.
	diskDescSize := diskFields[0].Descriptor()
	// disk.SizeValidator is a validator for the "size" field. It is called by the builders before save.
	disk.SizeValidator = diskDescSize.Validators[0].(func(int) error)
	ginfilemiddlewareFields := schema.GinFileMiddleware{}.Fields()
	_ = ginfilemiddlewareFields
	// ginfilemiddlewareDescAccessed is the schema descriptor for accessed field.
	ginfilemiddlewareDescAccessed := ginfilemiddlewareFields[2].Descriptor()
	// ginfilemiddleware.DefaultAccessed holds the default value on creation for the accessed field.
	ginfilemiddleware.DefaultAccessed = ginfilemiddlewareDescAccessed.Default.(bool)
}
