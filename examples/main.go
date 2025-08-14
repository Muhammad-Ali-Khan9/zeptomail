package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jiyamathias/zeptomail"
)

func main() {
	// sendHTMLEmail sends a HTML Email
	sendHTMLEmail()

	// sendHTMLEmail sends a templated email
	sendTemplatedEmail()

	// SendBatchTemplatedEmail sends a batch templated email
	SendBatchTemplatedEmail()

	// AddEmailTemplate can be used to add an email template.
	AddEmailTemplate()

	// UpdateEmailTemplate can be used to update an email template.
	UpdateEmailTemplate()

	// GetEmailTemplate is used to fetch a particular email template.
	GetEmailTemplate()

	// DeleteEmailTemplate is used to delete a template using template key.
	DeleteEmailTemplate()

	// SendBatchHTML sends a batch HTML Email
	SendBatchHTMLEmail()

	// FileCacheUploadAPI is used to upload a file to the cache.
	FileCacheUploadAPI()

	//GetEmailTemplates is used to list the required number of email templates in your ZeptoMail account.
	ListEmailTemplates()

}

// sendHTMLEmail sends a HTML Email
func sendHTMLEmail() {
	zeptomailToken := "your zeptomail authorization token"
	client := zeptomail.New(http.DefaultClient, zeptomailToken)

	data := "<div><b> Kindly click on Verify Account to confirm your account </b></div>"
	req := zeptomail.SendHTMLEmailReq{
		To: []zeptomail.SendEmailTo{
			{
				EmailAddress: zeptomail.EmailAddress{
					Address: "rudra.d@zylker.com",
					Name:    "Rudra",
				},
			},
		},
		From: zeptomail.EmailAddress{
			Address: "accounts@info.zylker.com",
			Name:    "Paula",
		},
		Subject:  "Account Confirmation",
		Htmlbody: data,
	}
	res, err := client.SendHTMLEmail(req)
	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	for _, e := range res.Data {
		fmt.Printf("response message: %v\n", e.Message)
	}
}

// sendHTMLEmail sends a templated email
func sendTemplatedEmail() {
	zeptomailToken := "your zeptomail authorization token"
	tempKey := "your template key"

	client := zeptomail.New(http.DefaultClient, zeptomailToken)
	req := zeptomail.SendTemplatedEmailReq{
		To: []zeptomail.SendEmailTo{
			{
				EmailAddress: zeptomail.EmailAddress{
					Address: "rudra.d@zylker.com",
					Name:    "Rudra",
				},
			},
		},
		From: zeptomail.EmailAddress{
			Address: "rudra.d@zylker.com",
			Name:    "Rudra",
		},
		TemplateKey: tempKey,
		MergeInfo: map[string]interface{}{
			"greeting": "Hello",
		},
	}

	res, err := client.SendTemplatedEmail(req)
	if err != nil {
		fmt.Printf("This is the error: %v", res)
	}

	fmt.Println("Response: ", res)
	for _, e := range res.Data {
		fmt.Printf("response message: %v\n", e.Message)
	}
}

// SendBatchTemplatedEmail sends a batch templated email
func SendBatchTemplatedEmail() {
	zeptomailToken := "your zeptomail authorization token"
	tempKey := "your zeptomail template key"

	client := zeptomail.New(http.DefaultClient, zeptomailToken)
	req := zeptomail.SendBatchTemplatedEmailReq{
		To: []zeptomail.SendBatchTemplateEmailTo{
			{
				EmailAddress: zeptomail.EmailAddress{
					Address: "rudra.d@zylker.com",
					Name:    "Rudra",
				},
				MergeInfo: map[string]interface{}{
					"userID": "my userID",
				},
			},
		},
		From: zeptomail.EmailAddress{
			Address: "accounts@info.zylker.com",
			Name:    "Paula",
		},
		TemplateKey: tempKey,
	}

	res, err := client.SendBatchTemplatedEmail(req)
	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	for _, e := range res.Data {
		fmt.Printf("response message: %v\n", e.Message)
	}
}

// AddEmailTemplate can be used to add an email template.
func AddEmailTemplate() {
	zeptomailToken := "your zeptomail authorization token"
	client := zeptomail.New(http.DefaultClient, zeptomailToken)

	req := zeptomail.AddEmailTemplateReq{
		TemplateName:   "E-invite",
		Subject:        "Invitation to the event",
		HtmlBody:       "<h1> Hi {{name}}</h1>, invitation link {{link}}", // for a html template
		TemplateAlias:  "e-invite emails",
		MailagentAlias: "mail agent alias",
	}

	res, err := client.AddEmailTemplate(req)
	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	fmt.Printf("response message: %v\n", res.Message)
}

// UpdateEmailTemplate can be used to update an email template.
func UpdateEmailTemplate() {
	zeptomailToken := "your zeptomail authorization token"
	client := zeptomail.New(http.DefaultClient, zeptomailToken)

	req := zeptomail.UpdateEmailTemplateReq{
		TemplateName:   "E-invite",
		Subject:        "Invitation to the event",
		HtmlBody:       "<h1> Hi {{name}}</h1>, invitation link {{link}}", // for a html template
		TemplateKey:    "your template key",
		MailagentAlias: "mail agent alias",
	}

	res, err := client.UpdateEmailTemplate(req)
	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	fmt.Printf("response message: %v\n", res.Message)
}

// GetEmailTemplate is used to fetch a particular email template.
func GetEmailTemplate() {
	zeptomailToken := "your zeptomail authorization token"
	var TemplateKey, MailagentAlias string = "your template key", "mail agent alias"

	client := zeptomail.New(http.DefaultClient, zeptomailToken)
	res, err := client.GetEmailTemplate(MailagentAlias, TemplateKey)
	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	fmt.Printf("response message: %v\n", res.Message)
}

// DeleteEmailTemplate is used to delete a template using template key.
func DeleteEmailTemplate() {
	zeptomailToken := "your zeptomail authorization token"
	var TemplateKey, MailagentAlias string = "your template key", "mail agent alias"
	client := zeptomail.New(http.DefaultClient, zeptomailToken)

	res, err := client.DeleteEmailTemplate(MailagentAlias, TemplateKey)
	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	fmt.Printf("response message: %v\n", res)
}

func SendBatchHTMLEmail() {
	zeptomailToken := "your zeptomail authorization token"
	client := zeptomail.New(http.DefaultClient, zeptomailToken)

	data := "<div><b>This is a sample email.{{contact}} {{company}}</b></div"
	req := zeptomail.SendBatchHTMLEmailReq{
		To: []zeptomail.SendBatchTemplateEmailTo{
			{
				EmailAddress: zeptomail.EmailAddress{
					Address: "paul.s@zfashions.com",
					Name:    "Paul",
				},
				MergeInfo: map[string]interface{}{
					"useriD": "useriD",
				},
			},
		},
		From: zeptomail.EmailAddress{
			Address: "invoice@zylker.com",
			Name:    "Invoice",
		},
		Subject:  "Account confirmation",
		Htmlbody: data,
	}

	res, err := client.SendBatchHTMLEmail(req)

	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	for _, e := range res.Data {
		fmt.Printf("response message: %v\n", e.Message)
	}

}

// filecacheupload is used to upload file to file cache
func FileCacheUploadAPI() {
	zeptomailToken := "your zeptomail authorization token"
	client := zeptomail.New(http.DefaultClient, zeptomailToken)

	// Set the file path
	filePath := "path_to_your_file.jpg"

	// Read the file content
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file content:", err)
		return
	}

	// Create the request object
	req := zeptomail.FileCacheUploadAPIReq{
		FileContent: fileContent,
		FileName:    filepath.Base(filePath),
	}

	res, err := client.FileCacheUploadAPI(req)
	if err != nil {
		fmt.Printf("This is the error: %v\n", err.Error())
	}

	for _, e := range res.Data {
		fmt.Printf("response message: %v\n", e.Message)
	}

}

func ListEmailTemplates() {
	zeptomailToken := "your zeptomail authorization token"
	client := zeptomail.New(http.DefaultClient, zeptomailToken)

	var Offset, Limit int
	var MailagentAlias string

	res, err := client.ListEmailTemplates(MailagentAlias, Offset, Limit)
	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	fmt.Printf("response message: %v\n", res.Message)
}
