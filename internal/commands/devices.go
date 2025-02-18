// commands.go
package commands

import (
	// "bytes"
	"encoding/json"
	"fmt"

	// "go/printer"
	"strings"

	"github.com/spf13/cobra"
	"romanhand.ru/fireZoneCli/internal/sender"
	"romanhand.ru/fireZoneCli/internal/structures"
)

func ListDevicesCmd(apiURL string, apiToken string) *cobra.Command {
	endpoint := "devices"
	cmd := &cobra.Command{
		Use:   "listdevice",
		Short: "get all device",
		Run: func(cmd *cobra.Command, args []string) {
			data := sender.Send(apiURL, apiToken, endpoint, strings.NewReader(""), "GET")
			// fmt.Println(string(data))
			var devices structures.ListDevicesCmd
			json.Unmarshal(data, &devices)
			for _, device := range devices.Devices {
				fmt.Println("---------------")
				fmt.Println("📌 ID:", device.ID)
				fmt.Println("🔹 Name:", device.Name)
				fmt.Println("👤 User ID:", device.UserID)
				fmt.Println("📅 Inserted at:", device.InsertedAt)
				fmt.Println("🔄 Updated at:", device.UpdatedAt)
				fmt.Println("🌐 Endpoint:", device.Endpoint)
				fmt.Println("📶 IPv4:", device.Ipv4)
				fmt.Println("📡 IPv6:", device.Ipv6)
				fmt.Println("🔑 Public Key:", device.PublicKey)
				fmt.Println("🛡️ Preshared Key:", device.PresharedKey)
				fmt.Println("📡 Server Public Key:", device.ServerPublicKey)
				fmt.Println("📌 MTU:", device.Mtu)
				fmt.Println("🔄 Persistent Keepalive:", device.PersistentKeepalive)
				fmt.Println("✅ Use Default Allowed IPs:", device.UseDefaultAllowedIps)
				fmt.Println("✅ Use Default DNS:", device.UseDefaultDNS)
				fmt.Println("✅ Use Default Endpoint:", device.UseDefaultEndpoint)
				fmt.Println("📡 Allowed IPs:", device.AllowedIps)
				fmt.Println("🌎 DNS Servers:", device.DNS)
			}

		},
	}
	return cmd
}
func GetDeviceCmd(apiURL string, apiToken string) *cobra.Command {
	// endpoint := "devices"
	cmd := &cobra.Command{
		Use:   "getdevice",
		Short: "get device",
		Run: func(cmd *cobra.Command, args []string) {
			endpoint := "devices/" + cmd.Flag("id").Value.String()
			data := sender.Send(apiURL, apiToken, endpoint, strings.NewReader(""), "GET")
			var device structures.SoloDeviceStruct
			json.Unmarshal(data, &device)
			fmt.Println("---------------")
			fmt.Println("📌 ID:", device.Device.ID)
			fmt.Println("🔹 Name:", device.Device.Name)
			fmt.Println("👤 User ID:", device.Device.UserID)
			fmt.Println("📅 Inserted at:", device.Device.InsertedAt)
			fmt.Println("🔄 Updated at:", device.Device.UpdatedAt)
			fmt.Println("🌐 Endpoint:", device.Device.Endpoint)
			fmt.Println("📶 IPv4:", device.Device.Ipv4)
			fmt.Println("📡 IPv6:", device.Device.Ipv6)
			fmt.Println("🔑 Public Key:", device.Device.PublicKey)
			fmt.Println("🛡️ Preshared Key:", device.Device.PresharedKey)
			fmt.Println("📡 Server Public Key:", device.Device.ServerPublicKey)
			fmt.Println("📌 MTU:", device.Device.Mtu)
			fmt.Println("🔄 Persistent Keepalive:", device.Device.PersistentKeepalive)
			fmt.Println("✅ Use Default Allowed IPs:", device.Device.UseDefaultAllowedIps)
			fmt.Println("✅ Use Default DNS:", device.Device.UseDefaultDNS)
			fmt.Println("✅ Use Default Endpoint:", device.Device.UseDefaultEndpoint)
			fmt.Println("📡 Allowed IPs:", device.Device.AllowedIps)
			fmt.Println("🌎 DNS Servers:", device.Device.DNS)
		},
	}
	id := ""
	cmd.Flags().StringVarP(&id, "id", "i", "", "id of the device")
	cmd.MarkFlagRequired("id")
	return cmd
}
