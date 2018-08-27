package mystringutils

import "strings"

// Upper returns the case of the given argument
func Upper(s string) string {
    return strings.ToUpper(s)
}

//Lower returns the lower of the given argument
func Lower(s string) string {
  return strings.ToLower(s)
}
