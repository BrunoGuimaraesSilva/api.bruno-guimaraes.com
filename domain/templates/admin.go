package templates

import (
	"bytes"
	"html/template"
)

type AdminEmailData struct {
	Name    string
	Email   string
	Message string
}

const AdminEmailTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>New Message Notification</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
            color: #333;
        }
        .container {
            width: 100%;
            max-width: 600px;
            margin: 0 auto;
            background-color: #ffffff;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        }
        .header {
            text-align: center;
            padding-bottom: 20px;
            border-bottom: 1px solid #eee;
        }
        .header h1 {
            color: #2c3e50;
        }
        .message-content {
            margin-top: 20px;
        }
        .message-content p {
            font-size: 16px;
            line-height: 1.5;
            color: #555;
        }
        .footer {
            text-align: center;
            margin-top: 30px;
            font-size: 12px;
            color: #aaa;
        }
        .footer a {
            color: #3498db;
            text-decoration: none;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>New Message Notification</h1>
        </div>
        <div class="message-content">
            <p><strong>Name:</strong> {{.Name}}</p>
            <p><strong>Email:</strong> {{.Email}}</p>
            <p><strong>Message:</strong></p>
            <p>{{.Message}}</p>
        </div>
        <div class="footer">
            <p>Thanks for using our service!</p>
            <p><a href="mailto:support@bruno-guimaraes.com">Contact Support</a></p>
        </div>
    </div>
</body>
</html>
`

func RenderAdminEmail(data AdminEmailData) (string, error) {
	tmpl, err := template.New("adminEmail").Parse(AdminEmailTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
