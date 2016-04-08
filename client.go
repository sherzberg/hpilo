package hpilo

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
)

type IloClient struct {
	Hostname string
	Username string
	Password string
}

func NewIloClient(hostname string, username string, password string) *IloClient {

	return &IloClient{
		Hostname: hostname,
		Username: username,
		Password: password,
	}
}

func (c *IloClient) request(xml string) ([]byte, error) {
	conn, err := tls.Dial("tcp", c.Hostname+":443", &tls.Config{
		InsecureSkipVerify: true,
	})
	defer conn.Close()
	if err != nil {
		panic(fmt.Sprintf("client: dial: %s", err))
	}

	HTTP_HEADER := "POST /ribcl HTTP/1.1\r\nHost: localhost\r\nContent-Length: %d\r\nConnection: Close\r\n\r\n"
	XML_HEADER := "<?xml version=\"1.0\"?>\r\n"

	body := []string{fmt.Sprintf(HTTP_HEADER, len(xml+XML_HEADER)), XML_HEADER, xml}

	for _, data := range body {
		_, err := conn.Write([]byte(data))
		if err != nil {
			log.Panic(err)
		}
	}

	response_data, err := ioutil.ReadAll(conn)

	return response_data, err

}

func (c *IloClient) GetFwVersion() (string, error) {
	xml := "<RIBCL\r\n VERSION=\"2.0\"\r\n>\r\n<LOGIN\r\n PASSWORD=\"%s\"\r\n USER_LOGIN=\"%s\"\r\n>\r\n<RIB_INFO\r\n MODE=\"read\"\r\n>\r\n<GET_FW_VERSION\r\n />\r\n</RIB_INFO>\r\n</LOGIN>\r\n</RIBCL>\r\n"

	response, err := c.request(fmt.Sprintf(xml, c.Password, c.Username))
	return string(response), err
}
