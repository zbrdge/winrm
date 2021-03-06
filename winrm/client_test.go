package winrm

import (
	"github.com/masterzen/winrm/soap"
	. "launchpad.net/gocheck"
)

func (s *WinRMSuite) TestNewClient(c *C) {
	client := NewClient("localhost", "Administrator", "v3r1S3cre7")
	c.Assert(client.url, Equals, "http://localhost:5985/wsman")
	c.Assert(client.username, Equals, "Administrator")
	c.Assert(client.password, Equals, "v3r1S3cre7")
}

func (s *WinRMSuite) TestClientCreateShell(c *C) {
	client := NewClient("localhost", "Administrator", "v3r1S3cre7")
	client.http = func(client *Client, message *soap.SoapMessage) (string, error) {
		c.Assert(message.String(), Contains, "http://schemas.xmlsoap.org/ws/2004/09/transfer/Create")
		return createShellResponse, nil
	}

	shell, _ := client.CreateShell()
	c.Assert(shell.ShellId, Equals, "67A74734-DD32-4F10-89DE-49A060483810")
}
