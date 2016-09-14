package validator

import "github.com/go-playground/universal-translator"

// TranslationFunc is the function type used to register or override
// custom translations
type TranslationFunc func(ut ut.Translator, fe FieldError) string

// RegisterTranslationsFunc allows for registering of translations
// for a 'ut.Translator' for use withing the 'TranslationFunc'
type RegisterTranslationsFunc func(ut ut.Translator) error
