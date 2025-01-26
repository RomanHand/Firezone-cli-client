// commands.go
package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"romanhand.ru/fireZoneCli/internal/sender"
	"romanhand.ru/fireZoneCli/internal/structures"
)

func GetUsersCmd(apiURL string, apiToken string) *cobra.Command {
	endpoint := "users"
	cmd := &cobra.Command{
		Use:   "getusers",
		Short: "get all users",
		Run: func(cmd *cobra.Command, args []string) {
			data := sender.Send(apiURL, apiToken, endpoint, strings.NewReader(""), "GET")
			var users structures.GetUsersStruct
			json.Unmarshal(data, &users)
			for _, user := range users.Users {
				fmt.Println("---------------","\nID:", user.ID,  "\nEmail:", user.Email,  "\nRole:", user.Role,  "\nDisabled at:", user.DisabledAt,  "\nLast signed in at:", user.LastSignedInAt,  "\nLast signed in method:", user.LastSignedInMethod)
			}
		},
	}
	return cmd
}
func CreateUserCmd(apiURL string, apiToken string) *cobra.Command {
	endpoint := "users"
	cmd := &cobra.Command{
		Use:   "createuser",
		Short: "crate user",
		Run: func(cmd *cobra.Command, args []string) {
			var body structures.CreateUserBody
			body.User.Email = cmd.Flag("email").Value.String()
			body.User.Role = cmd.Flag("role").Value.String()
			body.User.Password = cmd.Flag("password").Value.String()
			body.User.Password_confirmation = cmd.Flag("password").Value.String()
			jsonBody, _ := json.Marshal(body)
			readerBody := bytes.NewReader(jsonBody)
			data := sender.Send(apiURL, apiToken, endpoint, readerBody, "POST")
			var user structures.SoloUserStruct
			json.Unmarshal(data, &user)
			fmt.Println("---------------","\nID:", user.User.ID, "\nEmail:", user.User.Email,  "\nRole:", user.User.Role,  "\nDisabled at:", user.User.DisabledAt,  "\nLast signed in at:", user.User.LastSignedInAt,  "\nLast signed in method:", user.User.LastSignedInMethod)
			
		},
	}

	var name, email, role string
	cmd.Flags().StringVarP(&name, "password", "p", "", "Name of the user")
	cmd.Flags().StringVarP(&email, "email", "e", "", "Email of the user")
	cmd.Flags().StringVarP(&role, "role", "r", "", "Role of the user")
	cmd.MarkFlagRequired("pass")
	cmd.MarkFlagRequired("email")
	cmd.MarkFlagRequired("role")
	return cmd
}

func DeleteUserCmd(apiURL string, apiToken string) *cobra.Command {
	
	cmd := &cobra.Command{
		Use:   "deluser",
		Short: "Del user",
		Run: func(cmd *cobra.Command, args []string) {
			endpoint := "users/"+cmd.Flag("id").Value.String()
			data := sender.Send(apiURL, apiToken, endpoint, strings.NewReader(""), "DELETE")
			if data != nil {
				var msg structures.DelUserStruct
				json.Unmarshal(data, &msg)
				fmt.Printf("User deletion error: %v", msg.Error)
			}
		},
	}
	id := ""
	cmd.Flags().StringVarP(&id, "id", "i", "", "id of the user")
	cmd.MarkFlagRequired("id")
	return cmd
}

func GetUserCmd(apiURL string, apiToken string) *cobra.Command {
	
	cmd := &cobra.Command{
		Use:   "getuser",
		Short: "get user",
		Run: func(cmd *cobra.Command, args []string) {
			endpoint := "users/"+cmd.Flag("id").Value.String()
			data := sender.Send(apiURL, apiToken, endpoint, strings.NewReader(""), "GET")
			var user structures.SoloUserStruct
			json.Unmarshal(data, &user)
			fmt.Println("---------------","\nID:", user.User.ID, "\nEmail:", user.User.Email,  "\nRole:", user.User.Role,  "\nDisabled at:", user.User.DisabledAt,  "\nLast signed in at:", user.User.LastSignedInAt,  "\nLast signed in method:", user.User.LastSignedInMethod)
		},
	}
	id := ""
	cmd.Flags().StringVarP(&id, "id", "i", "", "id of the user")
	cmd.MarkFlagRequired("id")
	return cmd
}

func UpdateUserCmd(apiURL string, apiToken string) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "updateuser",
		Short: "update user",
		Run: func(cmd *cobra.Command, args []string) {
			endpoint := "users/"+cmd.Flag("id").Value.String()
			var body structures.CreateUserBody
			if body.User.Email == ""{body.User.Email = cmd.Flag("email").Value.String()}
			if body.User.Role == ""{body.User.Role = cmd.Flag("role").Value.String()}
			if body.User.Password == ""{
				body.User.Password = cmd.Flag("password").Value.String()
				body.User.Password_confirmation = cmd.Flag("password").Value.String()
			}
			jsonBody, _ := json.Marshal(body)
			readerBody := bytes.NewReader(jsonBody)
			data := sender.Send(apiURL, apiToken, endpoint, readerBody, "PATCH")
			var user structures.SoloUserStruct
			json.Unmarshal(data, &user)
			fmt.Println("---------------","\nID:", user.User.ID, "\nEmail:", user.User.Email,  "\nRole:", user.User.Role,  "\nDisabled at:", user.User.DisabledAt,  "\nLast signed in at:", user.User.LastSignedInAt,  "\nLast signed in method:", user.User.LastSignedInMethod)
			
		},
	}

	var name, email, role, id string
	cmd.Flags().StringVarP(&name, "password", "p", "", "Name of the user")
	cmd.Flags().StringVarP(&email, "email", "e", "", "Email of the user")
	cmd.Flags().StringVarP(&role, "role", "r", "", "Role of the user")
	cmd.Flags().StringVarP(&id, "id", "i", "", "id of the user")
	cmd.MarkFlagRequired("id")
	return cmd
}
