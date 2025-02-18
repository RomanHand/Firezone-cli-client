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

func ListUserCmd(apiURL string, apiToken string) *cobra.Command {
	endpoint := "users"
	cmd := &cobra.Command{
		Use:   "listusers",
		Short: "get all users",
		Run: func(cmd *cobra.Command, args []string) {
			data := sender.Send(apiURL, apiToken, endpoint, strings.NewReader(""), "GET")
			var users structures.GetUsersStruct
			json.Unmarshal(data, &users)
			for _, user := range users.Users {
				fmt.Println("---------------",
					"\n📌 ID:", user.ID,
					"\n📧 Email:", user.Email,
					"\n🔹 Role:", user.Role,
					"\n🚫 Disabled at:", user.DisabledAt,
					"\n🔑 Last signed in at:", user.LastSignedInAt,
					"\n🛠️ Last signed in method:", user.LastSignedInMethod,
				)
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
			body.User.PasswordConfirmation = cmd.Flag("password").Value.String()
			jsonBody, _ := json.Marshal(body)
			readerBody := bytes.NewReader(jsonBody)
			data := sender.Send(apiURL, apiToken, endpoint, readerBody, "POST")
			var user structures.SoloUserStruct
			json.Unmarshal(data, &user)
			fmt.Println("---------------",
				"\n📌 ID:", user.User.ID,
				"\n📧 Email:", user.User.Email,
				"\n🔹 Role:", user.User.Role,
				"\n🚫 Disabled at:", user.User.DisabledAt,
				"\n🔑 Last signed in at:", user.User.LastSignedInAt,
				"\n🛠️ Last signed in method:", user.User.LastSignedInMethod,
			)
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
				fmt.Printf("❌ User deletion error: %v\n", msg.Error)
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
			fmt.Println("---------------")
			fmt.Printf("📌 ID: %s\n", user.User.ID)
			fmt.Printf("📧 Email: %s\n", user.User.Email)
			fmt.Printf("🔹 Role: %s\n", user.User.Role)
			fmt.Printf("🚫 Disabled at: %v\n", user.User.DisabledAt)
			fmt.Printf("🔑 Last signed in at: %v\n", user.User.LastSignedInAt)
			fmt.Printf("🛠️ Last signed in method: %s\n", user.User.LastSignedInMethod)
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
				body.User.PasswordConfirmation = cmd.Flag("password").Value.String()
			}
			jsonBody, _ := json.Marshal(body)
			readerBody := bytes.NewReader(jsonBody)
			data := sender.Send(apiURL, apiToken, endpoint, readerBody, "PATCH")
			var user structures.UserStruct
			json.Unmarshal(data, &user)
			fmt.Println("---------------")
			fmt.Printf("📌 ID: %s\n", user.ID)
			fmt.Printf("📧 Email: %s\n", user.Email)
			fmt.Printf("🔹 Role: %s\n", user.Role)
			fmt.Printf("🚫 Disabled at: %v\n", user.DisabledAt)
			fmt.Printf("🔑 Last signed in at: %v\n", user.LastSignedInAt)
			fmt.Printf("🛠️ Last signed in method: %s\n", user.LastSignedInMethod)		
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
