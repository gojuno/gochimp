/**
* Copyright 2012 Matthew Baird
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*  distributed under the License is distributed on an "AS IS" BASIS,
*  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*  See the License for the specific language governing permissions and
*  limitations under the License.
**/
package gochimp

import (
	"fmt"
	"os"
	"testing"
)

var mandrill, err = NewMandrill(os.Getenv("MANDRILL_KEY"))
var user = os.Getenv("MANDRILL_USER")

func TestPing(t *testing.T) {
	response, err := mandrill.Ping()
	if response != "PONG!" {
		t.Error(fmt.Sprintf("failed to return PONG!, returned [%s]", response), err)
	}
}

func TestInfo(t *testing.T) {
	response, err := mandrill.UserInfo()
	if err != nil {
		t.Error("Error:", err)
	}
	if response.Username != user {
		t.Error("wrong user")
	}
}

func TestSenders(t *testing.T) {
	response, err := mandrill.UserSenders()
	if response == nil {
		t.Error("response was nil", err)
	}
	if err != nil {
		t.Error("Error:", err)
	}
}

func TestSending(t *testing.T) {
	var message Message = Message{Html: "<b>hi there</b>", Text: "hello text", Subject: "Test Mail", FromEmail: user,
		FromName: user}
	message.addRecipients(Recipient{Email: user, Name: user})
	response, err := mandrill.MessageSend(message, false)
	if err != nil {
		t.Error("Error:", err)
	}
	if response[0].Email != user {
		t.Errorf("Wrong email recipient, expecting %s, got %s", user, response[0].Email)
	}
}
