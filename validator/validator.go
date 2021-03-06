package validator

import (
	"errors"

	"github.com/JulzDiverse/aviator"
)

//Error Types: Merge-Section
type MergeCombinationError error
type MergeWithCombinationError error
type MergeExceptCombinationError error
type MergeRegexpCombinationError error

//Error Types: ForEach-Section
type ForEachCombinationError error
type ForEachFilesCombinationError error
type ForEachInCombinationError error
type ForEachRegexpCombinationError error
type ForEachWalkCombinationError error

type Validator struct{}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateSpruce(cfg []aviator.Spruce) error {
	for _, spruce := range cfg {
		if !isMergeArrayEmpty(spruce.Merge) {
			err := validateMergeSection(spruce.Merge)
			if err != nil {
				return err
			}
		}

		if !isForEachEmpty(spruce.ForEach) {
			err := validateForEachSection(spruce.ForEach)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validateMergeSection(cfg []aviator.Merge) error {
	for _, merge := range cfg {
		if !isMergeEmpty(merge) {
			err := validateMergeCombinations(merge)
			if err != nil {
				return err
			}
			err = validateMergeWithCombinations(merge.With)
			if err != nil {
				return err
			}

			err = validateMergeExceptCombination(merge)
			if err != nil {
				return err
			}

			err = validateMergeRegexpCombination(merge)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validateForEachSection(forEach aviator.ForEach) error {
	err := validateForEachCombination(forEach)
	if err != nil {
		return err
	}

	err = validateForEachFilesCombinations(forEach)
	if err != nil {
		return err
	}

	err = validateForEachInCombinations(forEach)
	if err != nil {
		return err
	}

	err = validateForEachRegexpCombination(forEach)
	if err != nil {
		return err
	}

	err = validateForEachWalkCombinations(forEach)
	if err != nil {
		return err
	}

	return nil
}

func validateMergeCombinations(merge aviator.Merge) error {
	var err MergeCombinationError
	if (merge.With.Files != nil) && (merge.WithIn != "" || merge.WithAllIn != "") || (merge.WithIn != "" && merge.WithAllIn != "") {
		err = errors.New(
			"INVALID SYNTAX: 'with', 'with_in', and 'with_all_in' are discrete parameters and cannot be defined together",
		)
	}
	return err
}

func validateMergeWithCombinations(with aviator.With) error {
	var err MergeWithCombinationError
	if (len(with.Files) == 0 || with.Files == nil) && (with.InDir != "" || with.Skip == true) {
		err = errors.New(
			"INVALID SYNTAX: 'with.in_dir' or 'with.skip_non_existing' can only be declared in combination with 'with.files'",
		)
	}
	return err
}

func validateMergeExceptCombination(merge aviator.Merge) error {
	var err MergeExceptCombinationError
	if (len(merge.Except) > 0) && (merge.WithIn == "" && merge.WithAllIn == "") {
		err = errors.New(
			"INVALID SYNTAX: 'merge.except' is only allowed in combination with 'merge.with_in' or 'merge.with_all_in'",
		)
	}
	return err
}

func validateMergeRegexpCombination(merge aviator.Merge) error {
	var err MergeRegexpCombinationError
	if (merge.Regexp != "") && (merge.WithIn == "" && merge.WithAllIn == "") {
		err = errors.New(
			"INVALID SYNTAX: 'merge.regexp' is only allowed in combination with 'merge.with_in' or 'merge.with_all_in'",
		)
	}
	return err
}

func validateForEachCombination(forEach aviator.ForEach) error {
	var err ForEachCombinationError
	if forEach.Files != nil && forEach.In != "" {
		err = errors.New(
			"INVALID SYNTAX: Mutually exclusive parameters declared 'for_each.in' and 'for_each.files'",
		)
	}
	return err
}

func validateForEachFilesCombinations(forEach aviator.ForEach) error {
	var err ForEachFilesCombinationError
	if (forEach.InDir != "" || forEach.Skip == true) && forEach.Files == nil {
		err = errors.New(
			"INVALID SYNTAX: 'for_each.in_dir' and 'for_each.skip_non_existing' can only be declared in combination with 'for_each.files'",
		)
	}
	return err
}

func validateForEachInCombinations(forEach aviator.ForEach) error {
	var err ForEachInCombinationError
	if ((forEach.Except != nil || len(forEach.Except) > 0) || forEach.SubDirs == true) && forEach.In == "" {
		err = errors.New(
			"INVALID SYNTAX: 'for_each.except' and 'for_each.include_sub_dirs' can only be declared in combination with 'for_each.in'",
		)
	}
	return err
}

func validateForEachRegexpCombination(forEach aviator.ForEach) error {
	var err ForEachRegexpCombinationError
	if (forEach.Regexp != "") && (forEach.In == "") {
		err = errors.New(
			"INVALID SYNTAX: 'for_each.regexp' is only allowed in combination with 'for_each.in'",
		)
	}
	return err
}

func validateForEachWalkCombinations(forEach aviator.ForEach) error {
	var err ForEachWalkCombinationError
	if (forEach.SubDirs == false) && (forEach.CopyParents == true || forEach.EnableMatching == true || forEach.ForAll != "") {
		err = errors.New(
			"INVALID SYNTAX: 'for_each.copy_parents', 'for_each.enable_matching', 'for_each.for_all' can only be declared in combination with 'for_each.inlcude_sub_dirs'",
		)
	}
	return err
}

func isForEachEmpty(forEach aviator.ForEach) bool {
	if (forEach.Files == nil || len(forEach.Files) == 0) &&
		forEach.InDir == "" &&
		(forEach.Except == nil || len(forEach.Except) == 0) &&
		forEach.In == "" &&
		forEach.Regexp == "" &&
		forEach.Skip == false &&
		forEach.SubDirs == false &&
		forEach.CopyParents == false &&
		forEach.EnableMatching == false &&
		forEach.ForAll == "" {
		return true
	}
	return false
}

func isMergeEmpty(merge aviator.Merge) bool {
	if merge.With.InDir == "" &&
		merge.With.Files == nil &&
		merge.With.Skip == false &&
		merge.WithAllIn == "" &&
		merge.Except == nil &&
		merge.Regexp == "" {
		return true
	}
	return false
}

func isMergeArrayEmpty(merges []aviator.Merge) bool {
	if merges == nil {
		return true
	}
	return false
}
