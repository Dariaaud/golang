package main

import (
	"awesomeProject/accounts/dto"
	"awesomeProject/cmd"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {
	portVal := flag.Int("port", 8080, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")

	flag.Parse()

	cmd := Command{
		Port:   *portVal,
		Host:   *hostVal,
		Cmd:    *cmdVal,
		Name:   *nameVal,
		Amount: *amountVal,
	}

	if err := cmd.Do(); err != nil {
		panic(err)
	}
}

type Command struct {
	Port   int
	Host   string
	Cmd    string
	Name   string
	Amount int
}

func (c *Command) Do() error {
	switch cmd.Cmd {
	case "create":
		if err := cmd.create(); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}

		return nil
	case "get":
		if err := cmd.get(); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}
		return nil

	case "delete":
		if err := cmd.delete(); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}
		return nil
	case "change_amount":
		if err := cmd.change_amount(); err != nil {
			return fmt.Errorf("change amount failed: %w", err)
		}
		return nil
	case "change_name":
		if err := cmd.change_name(); err != nil {
			return fmt.Errorf("change name failed: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}

func create(cmd Command) error {
	request := dto.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/create", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func get(cmd cmd.Command) error {
	if len(cmd.Name) == 0 {
		return fmt.Errorf("name is empty")
	}

	url := fmt.Sprintf("http://%s:%d/account/get?name=%s", cmd.Host, cmd.Port, cmd.Name)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http get failed: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}
		fmt.Println(string(body))
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func change_amount(cmd cmd.Command) error {
	if len(cmd.Name) == 0 {
		return fmt.Errorf("name is empty")
	}

	request := dto.ChangeAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	url := fmt.Sprintf("http://%s:%d/account/change_amount", cmd.Host, cmd.Port)
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func change_name(cmd cmd.Command) error {
	if len(cmd.Name) == 0 {
		return fmt.Errorf("name is empty")
	}

	request := dto.PatchAccountRequest{
		Name:    cmd.Name,
		NewName: cmd.Name,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	url := fmt.Sprintf("http://%s:%d/account/change_name", cmd.Host, cmd.Port)
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func delete(cmd cmd.Command) error {
	if len(cmd.Name) == 0 {
		return fmt.Errorf("name is empty")
	}

	url := fmt.Sprintf("http://%s:%d/account/delete?name=%s", cmd.Host, cmd.Port, cmd.Name)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}
