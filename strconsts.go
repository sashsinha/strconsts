// Provides the same string constants available in the Python string module.
package strconsts

// ASCIILetters is the concatenation of ASCIILowercase and ASCIIUppercase.
// This value is not locale-dependent.
const ASCIILetters = ASCIILowercase + ASCIIUppercase

// ASCIILowercase is the string of lowercase letters 'abcdefghijklmnopqrstuvwxyz'.
// This value is not locale-dependent and will not change.
const ASCIILowercase = "abcdefghijklmnopqrstuvwxyz"

// ASCIIUppercase is the string of uppercase letters 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.
// This value is not locale-dependent and will not change.
const ASCIIUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Digits is the string '0123456789'.
// This value represents all decimal digits in ASCII.
const Digits = "0123456789"

// HexDigits is the string '0123456789abcdefABCDEF'.
// This value represents all hexadecimal digits in ASCII, including both
// lowercase and uppercase forms.
const HexDigits = "0123456789abcdefABCDEF"

// OctDigits is the string '01234567'.
// This value represents all octal digits in ASCII.
const OctDigits = "01234567"

// Punctuation is a string of ASCII characters that are considered punctuation
// characters in the C locale. This includes:
// `!"#$%&'()*+,-./:;<=>?@[\]^_` + "`{|}~`.
const Punctuation = `!"#$%&'()*+,-./:;<=>?@[\]^_` + "`{|}~"

// Printable is a string of ASCII characters that are considered printable.
// This is a combination of Digits, ASCIILetters, Punctuation, and Whitespace.
const Printable = Digits + ASCIILetters + Punctuation + Whitespace

// Whitespace is a string containing all ASCII characters that are considered
// whitespace. This includes:
// - Space (' ')
// - Tab ('\t')
// - Linefeed ('\n')
// - Return ('\r')
// - Formfeed ('\f')
// - Vertical tab ('\v')
const Whitespace = " \t\n\r\f\v"
