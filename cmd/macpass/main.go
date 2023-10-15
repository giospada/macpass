package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"internal/comunication"

	"github.com/go-ldap/ldap"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

var ldapUri string = ""
var baseDN string = ""
var bindAdmin string = ""
var bindPassword string = ""
var userDNType string = ""
var socketPath string = ""

func main() {
	setConfig()
	user, _ := ldapAuth()
	macAdd, _ := inputMac()
	time, _ := timeRegistered()

	fmt.Print(macAdd + "\t" + user + "\t")
	fmt.Println(time)

	send(comunication.Request{User: user, Mac: macAdd, Duration: time})

}

func ldapAuth() (string, error) {
	l, err := ldap.DialURL(ldapUri)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// First bind with a read only user
	err = l.Bind(bindAdmin, bindPassword)
	if err != nil {
		log.Fatal(err)
	}

	username, password, _ := credentials()

	err = l.Bind(userDNType+"="+username+","+baseDN, password)
	if err != nil {
		log.Fatal(err)
	}

	return username, nil
}

func credentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}
	fmt.Println()

	password := string(bytePassword)
	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}

func inputMac() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a MAC address: ")
	macAdd, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	macAdd = strings.TrimSpace(macAdd)

	if _, err := net.ParseMAC(macAdd); err != nil {
		log.Println(err)
	}

	return macAdd, nil
}

func setConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	viper.AddConfigPath(exPath) // for now the config should be in the same directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found")
			log.Fatal(err)
		} else {
			log.Println("Config file was found but another error was produced")
			log.Fatal(err)
		}
	}

	log.Println("Config file found and successfully parsed")

	ldapUri = viper.GetString("ldapUri")
	baseDN = viper.GetString("baseDN")
	bindAdmin = viper.GetString("bindAdmin")
	bindPassword = viper.GetString("bindPassword")
	userDNType = viper.GetString("userDNType")
	socketPath = viper.GetString("socketPath")
}

func timeRegistered() (time.Duration, error) {
	fmt.Print("Enter the duration for the connection in hours (MAX 4): ")
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		log.Fatal(err)
	}

	if i > 4 {
		i = 4
	} else if i <= 0 {
		i = 1
	}

	return time.Duration(i) * time.Hour, nil
}

func send(r comunication.Request) {
	// Connect to macpassd socket
	conn, err := net.Dial("unix", "/tmp/macpass.sock")
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	conn.Close()
}
