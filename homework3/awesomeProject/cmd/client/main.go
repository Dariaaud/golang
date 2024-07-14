package main

import (
	"awesomeProject/cmd"
	"awesomeProject/proto"
	"time"

	"context"
	"flag"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client_imp proto.AccountManagerClient
}

func (c *Client) Do(cmd cmd.Command) error {
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

type Command struct {
	Port   int
	Host   string
	Cmd    string
	Name   string
	Amount int
}

func (c *Client) create(cmd cmd.Command) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.client_imp.CreateAccount(ctx, &proto.CreateAccountRequest{Name: cmd.Name, Amount: int32(cmd.Amount)})
	return err
}

func (c *Client) delete(cmd cmd.Command) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.client_imp.DeleteAccount(ctx, &proto.DeleteAccountRequest{Name: cmd.Name})
	return err
}

func (c *Client) get(cmd cmd.Command) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.client_imp.GetAccount(ctx, &proto.GetAccountRequest{Name: cmd.Name})
	if err == nil {
		fmt.Printf("account name: %s amount: %d", resp.Name, resp.Amount)
	}
	return err
}

func (c *Client) changeAmount(cmd cmd.Command) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.client_imp.ChangeAmountAccount(ctx, &proto.ChangeAmountAccountRequest{Name: cmd.Name, NewAmount: int32(cmd.Amount)})
	return err
}

func (c *Client) changeName(cmd cmd.Command) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.client_imp.ChangeNameAccount(ctx, &proto.ChangeNameAccountRequest{PrevName: cmd.Name, NewName: cmd.NewName})
	return err
}

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

	conn, err := grpc.NewClient("0.0.0.0:4567", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	cl := Client{proto.NewAccountManagerClient(conn)}

	if err := cl.Do(cmd); err != nil {
		panic(err)
	}
}
