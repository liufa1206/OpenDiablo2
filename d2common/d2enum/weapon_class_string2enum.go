// Code generated by "string2enum -samepkg -linecomment -type WeaponClass"; DO NOT EDIT.

package d2enum

import "fmt"

// WeaponClassFromString returns the WeaponClass enum corresponding to s.
func WeaponClassFromString(s string) WeaponClass {
	if len(s) == 0 {
		return 0
	}
	for i := range _WeaponClass_index[:len(_WeaponClass_index)-1] {
		if s == _WeaponClass_name[_WeaponClass_index[i]:_WeaponClass_index[i+1]] {
			return WeaponClass(i)
		}
	}
	panic(fmt.Errorf("unable to locate WeaponClass enum corresponding to %q", s))
}
