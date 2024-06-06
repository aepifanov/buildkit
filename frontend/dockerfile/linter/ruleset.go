package linter

import (
	"fmt"
)

var (
	RuleStageNameCasing = LinterRule[func(string) string]{
		Name:        "StageNameCasing",
		Description: "Stage names should be lowercase",
		URL:         "https://docs.docker.com/go/dockerfile/rule/stage-name-casing/",
		Format: func(stageName string) string {
			return fmt.Sprintf("Stage name '%s' should be lowercase", stageName)
		},
	}
	RuleFromAsCasing = LinterRule[func(string, string) string]{
		Name:        "FromAsCasing",
		Description: "The 'as' keyword should match the case of the 'from' keyword",
		URL:         "https://docs.docker.com/go/dockerfile/rule/from-as-casing/",
		Format: func(from, as string) string {
			return fmt.Sprintf("'%s' and '%s' keywords' casing do not match", as, from)
		},
	}
	RuleNoEmptyContinuations = LinterRule[func() string]{
		Name:        "NoEmptyContinuations",
		Description: "Empty continuation lines will become errors in a future release",
		URL:         "https://docs.docker.com/go/dockerfile/rule/no-empty-continuations/",
		Format: func() string {
			return "Empty continuation line"
		},
	}
	RuleSelfConsistentCommandCasing = LinterRule[func(string) string]{
		Name:        "SelfConsistentCommandCasing",
		Description: "Commands should be in consistent casing (all lower or all upper)",
		URL:         "https://docs.docker.com/go/dockerfile/rule/self-consistent-command-casing/",
		Format: func(command string) string {
			return fmt.Sprintf("Command '%s' should be consistently cased", command)
		},
	}
	RuleFileConsistentCommandCasing = LinterRule[func(string, string) string]{
		Name:        "FileConsistentCommandCasing",
		Description: "All commands within the Dockerfile should use the same casing (either upper or lower)",
		URL:         "https://docs.docker.com/go/dockerfile/rule/file-consistent-command-casing/",
		Format: func(violatingCommand, correctCasing string) string {
			return fmt.Sprintf("Command '%s' should match the case of the command majority (%s)", violatingCommand, correctCasing)
		},
	}
	RuleDuplicateStageName = LinterRule[func(string) string]{
		Name:        "DuplicateStageName",
		Description: "Stage names should be unique",
		URL:         "https://docs.docker.com/go/dockerfile/rule/duplicate-stage-name/",
		Format: func(stageName string) string {
			return fmt.Sprintf("Duplicate stage name %q, stage names should be unique", stageName)
		},
	}
	RuleReservedStageName = LinterRule[func(string) string]{
		Name:        "ReservedStageName",
		Description: "Reserved stage names should not be used to name a stage",
		URL:         "https://docs.docker.com/go/dockerfile/rule/reserved-stage-name/",
		Format: func(reservedStageName string) string {
			return fmt.Sprintf("Stage name should not use the same name as reserved stage %q", reservedStageName)
		},
	}
	RuleJSONArgsRecommended = LinterRule[func(instructionName string) string]{
		Name:        "JSONArgsRecommended",
		Description: "JSON arguments recommended for ENTRYPOINT/CMD to prevent unintended behavior related to OS signals",
		URL:         "https://docs.docker.com/go/dockerfile/rule/json-args-recommended/",
		Format: func(instructionName string) string {
			return fmt.Sprintf("JSON arguments recommended for %s to prevent unintended behavior related to OS signals", instructionName)
		},
	}
	RuleMaintainerDeprecated = LinterRule[func() string]{
		Name:        "MaintainerDeprecated",
		Description: "The maintainer instruction is deprecated, use a label instead to define an image author",
		URL:         "https://docs.docker.com/go/dockerfile/rule/maintainer-deprecated/",
		Format: func() string {
			return "Maintainer instruction is deprecated in favor of using label"
		},
	}
	RuleUndeclaredArgInFrom = LinterRule[func(string, string) string]{
		Name:        "UndeclaredArgInFrom",
		Description: "FROM command must use declared ARGs",
		URL:         "https://docs.docker.com/go/dockerfile/rule/undeclared-arg-in-from/",
		Format: func(baseArg, suggest string) string {
			out := fmt.Sprintf("FROM argument '%s' is not declared", baseArg)
			if suggest != "" {
				out += fmt.Sprintf(" (did you mean %s?)", suggest)
			}
			return out
		},
	}
	RuleWorkdirRelativePath = LinterRule[func(workdir string) string]{
		Name:        "WorkdirRelativePath",
		Description: "Relative workdir without an absolute workdir declared within the build can have unexpected results if the base image changes",
		URL:         "https://docs.docker.com/go/dockerfile/rule/workdir-relative-path/",
		Format: func(workdir string) string {
			return fmt.Sprintf("Relative workdir %q can have unexpected results if the base image changes", workdir)
		},
	}
	RuleUndefinedVar = LinterRule[func(string, string) string]{
		Name:        "UndefinedVar",
		Description: "Variables should be defined before their use",
		URL:         "https://docs.docker.com/go/dockerfile/rule/undefined-var/",
		Format: func(arg, suggest string) string {
			out := fmt.Sprintf("Usage of undefined variable '$%s'", arg)
			if suggest != "" {
				out += fmt.Sprintf(" (did you mean $%s?)", suggest)
			}
			return out
		},
	}
	RuleMultipleInstructionsDisallowed = LinterRule[func(instructionName string) string]{
		Name:        "MultipleInstructionsDisallowed",
		Description: "Multiple instructions of the same type should not be used in the same stage",
		URL:         "https://docs.docker.com/go/dockerfile/rule/multiple-instructions-disallowed/",
		Format: func(instructionName string) string {
			return fmt.Sprintf("Multiple %s instructions should not be used in the same stage because only the last one will be used", instructionName)
		},
	}
	RuleLegacyKeyValueFormat = LinterRule[func(cmdName string) string]{
		Name:        "LegacyKeyValueFormat",
		Description: "Legacy key/value format with whitespace separator should not be used",
		URL:         "https://docs.docker.com/go/dockerfile/rule/legacy-key-value-format/",
		Format: func(cmdName string) string {
			return fmt.Sprintf("\"%s key=value\" should be used instead of legacy \"%s key value\" format", cmdName, cmdName)
		},
	}
	RuleInvalidBaseImagePlatform = LinterRule[func(string, string, string) string]{
		Name:        "InvalidBaseImagePlatform",
		Description: "Base image platform does not match expected target platform",
		Format: func(image, expected, actual string) string {
			return fmt.Sprintf("Base image %s was pulled with platform %q, expected %q for current build", image, actual, expected)
		},
	}
)
