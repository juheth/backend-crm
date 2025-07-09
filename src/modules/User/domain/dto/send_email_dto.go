package dto

type SendEmail struct {
	Subject   string
	Recipient string
	PlainText string
	Template  string
}
