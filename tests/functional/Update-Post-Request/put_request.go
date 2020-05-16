package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"walk-the-camino/data"
	"walk-the-camino/utils"

	"github.com/DATA-DOG/godog"
)

func theAppIP(ip string) error {
	fmt.Println("ip:", ip)
	return nil
}

func theAppPort(port string) error {
	fmt.Println("port:", port)
	return nil
}

func theTestcaseIs(tc string) error {
	return nil
}

func validateRedisDbThatLinkIs(state string) error {
	return nil
}

func sendPost() error {
	var t *http.Transport
	certFile := os.Getenv(utils.Cert)
	caCert, err := ioutil.ReadFile(certFile)
	if err != nil {
		fmt.Println("Error while fetching cert file", err.Error())
		return err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	config := &tls.Config{
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}
	t = &http.Transport{
		TLSClientConfig:   config,
		DialTLS:           nil,
		DisableKeepAlives: false,
	}
	client := &http.Client{
		Transport: t,
	}
	fmt.Println("About to POST txn")
	url := fmt.Sprintf("https://%s:%s/%s", "walk-the-camino", "4444", "employee")
	emp := data.Employee{"1", "Higgs", "Boson", "Physicist", "Switzerland", "CERN"}
	oData, _ := json.Marshal(emp)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(oData))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while sending txn ", err)
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Did not receive ok response from app")
		return errors.New("did not receive ok response")
	}
	defer resp.Body.Close()
	return nil
}

func sendPut() error {
	var t *http.Transport
	certFile := os.Getenv(utils.Cert)
	caCert, err := ioutil.ReadFile(certFile)
	if err != nil {
		fmt.Println("Error while fetching cert file", err.Error())
		return err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	config := &tls.Config{
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}
	t = &http.Transport{
		TLSClientConfig:   config,
		DialTLS:           nil,
		DisableKeepAlives: false,
	}
	client := &http.Client{
		Transport: t,
	}
	fmt.Println("About to send PUT Transaction")
	url := fmt.Sprintf("https://%s:%s/%s", "walk-the-camino", "4444", "employees/1")
	go func(string) {
		emp := data.Employee{"1", "Sonya", "Sotomayor", "Supreme Court Justice", "USA", "DC"}
		oData, _ := json.Marshal(emp)
		req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(oData))
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error while sending txn ", err)
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Did not receive ok response from app")
		}
		resp.Body.Close()
	}(url)
	time.Sleep(time.Second)
	response, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error while getting updated employee " + err.Error())
		return fmt.Errorf("Error while getting updated employee " + err.Error())
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	employee := data.Employee{}
	if err = json.Unmarshal(body, &employee); err != nil {
		fmt.Printf("Unable to unmarshall while getting updated employee " + err.Error())
		return errors.New("Unable to unmarshall while getting updated employee " + err.Error())
	}
	if employee.Designation != "Supreme Court Justice" {
		log.Printf("expected designation [%s] actual designation [%s]", "Supreme Court Justice \n", employee.Designation)
		return fmt.Errorf("expected designation [%s] actual designation [%s]", "Supreme Court Justice", employee.Designation)
	}
	return nil
}

func PostFeatureContext(s *godog.Suite) {
	s.Step(`^the app ip is "([^"]*)"$`, theAppIP)
	s.Step(`^the app port is "([^"]*)"$`, theAppPort)
	s.Step(`^the testcase is "([^"]*)"$`, theTestcaseIs)
	s.Step(`^Add new employee$`, sendPost)
	s.Step(`^update employee and validate response$`, sendPut)
}
