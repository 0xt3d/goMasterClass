package main

import (
	"fmt"
	"strings"
)

type Cart struct {
	items map[string]map[string]string
}

func NewCart() *Cart {
	return &Cart{items: make(map[string]map[string]string)}
}

func (c *Cart) addItem(itemID string) string {
	c.items[itemID] = make(map[string]string)
	return itemID
}

func (c *Cart) setVariable(itemID, key, value string) {
	c.items[itemID][key] = value
}

func (c *Cart) placeOrder() string {
	return "OrderPlaced"
}

func main() {
	requestedForCountry := "Ireland"
	isCellularServiceRequired := "Yes"
	replacement := "Yes"

	if !strings.Contains(requestedForCountry, "Ireland") && !strings.Contains(requestedForCountry, "Japan") && !strings.Contains(requestedForCountry, "Austria") {
		if isCellularServiceRequired == "Yes" {
			if replacement == "Yes" {
				cart := NewCart()
				providerMap := map[string]string{
					"AT&T Wireless - US":       "AT&T",
					"Verizon - US [Non-Standard]": "Verizon Wireless",
					"Vodafone IE":              "",
					"Vodafone NL":              "Vodafone Netherlands",
					"Vodafone UK":              "Vodafone UK",
					"Vodafone DE":              "Vodafone Germany",
					"Vodafone IT":              "Vodafone Italy",
					"Telus CA":                 "Telus Canada",
					"T-Mobile - US [Non-Standard]": "T-Mobile",
					"Sunrise - Switzerland":    "Sunrise Switzerland",
					"Orange - Belgium":         "Orange Belgium",
					"Orange - France":          "Orange France",
					"Orange - Spain":           "Orange Spain",
					"Docomo - Japan":           "",
					"T-Mobile - Austria":       "",
					"SoftBank - Japan":         "",
				}
				phoneProvider := "AT&T Wireless - US"
				mc := "AT&T Wireless - US"
				carrier := ""

				for key, value := range providerMap {
					if strings.Contains(mc, key) {
						carrier = value
						break
					}
				}

				item := cart.addItem("90520724870f42d0a3fe99f73cbb355f")
				cart.setVariable(item, "x_sakon_requested_for", "requested_for_value")
				cart.setVariable(item, "requested_for", "requested_for_value")
				cart.setVariable(item, "employee_id", "employee_number_value")
				cart.setVariable(item, "email_id", "email_value")
				cart.setVariable(item, "phone_number", "current_bus_phone_value")
				cart.setVariable(item, "device_type", "Tablet")
				cart.setVariable(item, "device_manufacturer", "ipad_device_manufacturer_value")
				cart.setVariable(item, "device_model", "ipad_device_model_value")
				cart.setVariable(item, "device_imei_imei2_number", "ipad_imei2_value")
				cart.setVariable(item, "sim_eid_number", "ipad_eid_value")

				if carrier != "" {
					cart.setVariable(item, "wireless_device_carrier", carrier)
				}

				cart.setVariable(item, "description", "request_item_value")
				sub := cart.placeOrder()
				fmt.Println("replacepad", sub)
			} else {
				cart := NewCart()
				providerMap := map[string]string{
					"AT&T Wireless - US":       "AT&T",
					"Verizon - US [Non-Standard]": "Verizon Wireless",
					"Vodafone IE":              "",
					"Vodafone NL":              "Vodafone Netherlands",
					"Vodafone UK":              "Vodafone UK",
					"Vodafone DE":              "Vodafone Germany",
					"Vodafone IT":              "Vodafone Italy",
					"Telus CA":                 "Telus Canada",
					"T-Mobile - US [Non-Standard]": "T-Mobile",
					"Sunrise - Switzerland":    "Sunrise Switzerland",
					"Orange - Belgium":         "Orange Belgium",
					"Orange - France":          "Orange France",
					"Orange - Spain":           "Orange Spain",
					"Docomo - Japan":           "",
					"T-Mobile - Austria":       "",
					"SoftBank - Japan":         "",
				}
				phoneProvider := "AT&T Wireless - US"
				mc := "AT&T Wireless - US"
				carrier := ""

				for key, value := range providerMap {
					if strings.Contains(mc, key) {
						carrier = value
						break
					}
				}

				item := cart.addItem("4811a2a0874342d0a3fe99f73cbb3572")
				cart.setVariable(item, "x_sakon_requested_for", "requested_for_value")
				cart.setVariable(item, "requested_for", "requested_for_value")
				cart.setVariable(item, "employee_id", "employee_number_value")
				cart.setVariable(item, "email_id", "email_value")
				cart.setVariable(item, "device_type", "Tablet")
				cart.setVariable(item, "device_manufacturer", "ipad_device_manufacturer_value")
				cart.setVariable(item, "device_model", "ipad_device_model_value")
				cart.setVariable(item, "device_imei_imei2_number", "ipad_imei2_value")
				cart.setVariable(item, "sim_eid_number", "ipad_eid_value")

				if carrier != "" {
					cart.setVariable(item, "wireless_device_carrier", carrier)
				}

				cart.setVariable(item, "description", "request_item_value")
				sub := cart.placeOrder()
				fmt.Println("replacepad", sub)
			}
		}
	}
}
liwerhfliuaerfgliaeg