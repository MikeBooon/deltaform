package enums

type InputType string

const (
	InputTypeText     InputType = "TEXT"
	InputTypeTextArea InputType = "TEXTAREA"
	InputTypeNumber   InputType = "NUMBER"
	InputTypeDate     InputType = "DATE"
	InputTypeBoolean  InputType = "BOOLEAN"
	InputTypeSelect   InputType = "SELECT"
	InputTypeRadio    InputType = "RADIO"
	InputTypeCheckbox InputType = "CHECKBOX"
)

var InputTypeOptions = []InputType{
	InputTypeText,
	InputTypeTextArea,
	InputTypeNumber,
	InputTypeDate,
	InputTypeBoolean,
	InputTypeSelect,
	InputTypeRadio,
	InputTypeCheckbox,
}
