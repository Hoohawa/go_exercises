// Available commands
// pwd - return current working directory
// get <filename> - returns contents of said file or error if not present
// ls - lists files and directories in current directory
// put <filename> - saves file into current directory
// close - closes the connection

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	clientCount := 0
	wd, err := os.Getwd() // Start initial working dir
	checkErr(err)

	listener, err := net.Listen("tcp", "localhost:8000")
	checkErr(err)

	fmt.Println("Started listening on localhost:8000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		client := Client{clientCount, conn, wd}
		go client.HandleConn()
		clientCount++
	}
}

type Client struct {
	Id         int
	Conn       net.Conn
	CurrentDir string
}

func (client *Client) HandleConn() {
	defer client.Conn.Close()

	fmt.Printf("Client %d connected\n", client.Id)

	scanner := bufio.NewReader(client.Conn)
	for {
		cmdLine, err := scanner.ReadString('\n')
		checkErr(err)
		fmt.Printf("[%d]RECEIVED: %s", client.Id, cmdLine)

		cmd, args := parseCommandLine(cmdLine)
		fmt.Printf("[%d]CMD: %s\nARGS: %s\n", client.Id, cmd, strings.Join(args, " "))

		switch {
		case cmd == "pwd":
			client.Write((fmt.Sprintf("FTP> %s\n", client.CurrentDir)))
		case cmd == "cd":
			client.HandleCD(args)
		case cmd == "ls":
			client.HandleLS()
		case cmd == "get":
			client.HandleGet(args)
		case cmd == "put":
			client.HandlePut(args)
		case cmd == "close":
			client.Write((fmt.Sprintf("FTP> Closing connection\n")))
			client.Conn.Close()
			break
		default:
			client.Write(("FTP> ERROR: Unrecognized command"))
		}
	}
}

func (client *Client) HandlePut(args []string) {
	if len(args) == 0 {
		client.Write((fmt.Sprintf("FTP> ERROR: put command requires 1 argument\n")))
	}
	fileName := args[0]
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		client.Write((fmt.Sprintf("FTP> ERROR: Cannot create file: %s %v\n", fileName, err)))
	} else {
		client.Write(("Input file, press Ctrl+D to signal EOF\n"))
		// Copy file contents to a file
		clientReader := bufio.NewReader(client.Conn)
		fileWriter := bufio.NewWriter(file)
		writtenBytes, err := io.Copy(fileWriter, clientReader)
		fileWriter.Flush()
		// Inform client of transfer result
		if err != nil {
			client.Write((fmt.Sprintf("FTP> ERROR: Failed to write file: %s %v\n", fileName, err)))
		} else {
			client.Write((fmt.Sprintf("\nWritten %v bytes to: %s\n", writtenBytes, fileName)))
		}
	}
}

func (client *Client) HandleGet(args []string) {
	if len(args) == 0 {
		client.Write((fmt.Sprintf("FTP> ERROR: get command requires 1 argument\n")))
	}
	fileName := args[0]
	file, err := os.Open(fileName)
	if err != nil {
		client.Write((fmt.Sprintf("FTP> ERROR: Cannot get file: %s %v\n", fileName, err)))
	} else {
		fileReader := bufio.NewReader(file)
		io.Copy(client.Conn, fileReader)
		client.Write(("\n--EOF--"))
	}
}

func (client *Client) HandleLS() {
	fmt.Printf("Handling LS\n")
	dirListing, err := ioutil.ReadDir(client.CurrentDir)
	if err != nil {
		client.Write((fmt.Sprintf("FTP> ERROR: %v\n", client.CurrentDir)))
	} else {
		client.Write((fmt.Sprintf("FTP> Dir contents: %s\n", client.CurrentDir)))
		for i := range dirListing {
			itemType := "file"
			if dirListing[i].IsDir() {
				itemType = "dir"
			}
			client.Write((fmt.Sprintf("%s\t%s\t%v\t%v\n",
				itemType, dirListing[i].Name(), dirListing[i].Size(), dirListing[i].ModTime())))
		}
	}
}

func (client *Client) HandleCD(args []string) {
	if len(args) == 0 {
		client.Write(fmt.Sprintf("FTP> ERROR: cd command requires 1 argument\n"))
	}
	dir := args[0]
	exists, err := exists(dir)
	if err != nil {
		client.Write(fmt.Sprintf("FTP> ERROR: Cannot move to directory dir: %s\n", dir))
	} else if !exists {
		client.Write(fmt.Sprintf("FTP> ERROR: Directory doesn't exist dir: %s\n", dir))
	} else {
		client.CurrentDir = dir
		client.Write(fmt.Sprintf("FTP> Changed directory to: %s\n", dir))
	}
}

func (client *Client) Write(msg string) {
	client.Conn.Write([]byte(msg))
}

func parseCommandLine(cmdLine string) (string, []string) {
	cmdLine = cmdLine[:len(cmdLine)-1] // Strip '\n' from the end
	cmdParts := strings.Split(cmdLine, " ")
	cmd := cmdParts[0]
	args := []string{}
	if len(cmdParts) > 1 {
		args = cmdParts[1:]
	}
	return cmd, args
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
