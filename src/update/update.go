package update

import (
	as "aliveServices"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/user"
	"runtime"
	"strings"
	"time"
)

//ReqHeader consist of myaddr and mylist ...
type ReqHeader struct {
	MyAddr net.Addr
	MyList []string
}

var OthersHeaders []ReqHeader

//UpdateNUtv starts and collect essential data before init ...
func UpdateNUtv() (time.Time, ReqHeader, error) {
	fmt.Println(" updating  allIP.txt")
	// allIP.AllIP()
	fmt.Println("updated allIP.txt")
	var reqh ReqHeader
	var err error
	reqh.MyAddr = MyAddr()
	reqh.MyList, err = getMovieL()
	if err != nil {
		os.Exit(12)
	}
	// getMovieL()
	fmt.Println("Finding other nodes started")
	err = as.GetRunningServices()
	if err != nil {
		return time.Now(), reqh, fmt.Errorf("error in upating your application, kindly contact Bramhastra")
	}
	fmt.Println("prepared aliveServices.txt")
	fmt.Println("starting listing online vidoes in the network")

	if reqh.MyAddr != nil && reqh.MyList != nil {
		fmt.Println("Your App is Updated!")
		return time.Now(), reqh, nil
	}
	return time.Now(), reqh, fmt.Errorf("sorry")
}

func getMovieL() ([]string, error) {
	var list []string
	// h, err := os.Hostname()
	u, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("not getting hostname")
	}
	h := u.Username

	f, err := os.OpenFile("mylist.txt", os.O_TRUNC|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, fmt.Errorf("something wrong with mylist.txt")
	}
	defer f.Close()
	switch runtime.GOOS {
	case "linux":
		{
			files, err := ioutil.ReadDir(fmt.Sprintf("/home/%s/Desktop", h))
			if err != nil {
				return nil, fmt.Errorf("path for Dektop is not correct%v", err)
			}
			for _, n := range files {
				if strings.HasPrefix(n.Name(), "mov") {
					movies, err := ioutil.ReadDir(fmt.Sprintf("/home/%s/Desktop/%s", h, n.Name()))
					if err != nil {
						return nil, fmt.Errorf("not able to read movies folder")
					}
					for _, movie := range movies {

						// fmt.Println(movie)
						nameMov := ismovie(movie.Name())
						if nameMov != "nil" {
							list = append(list, nameMov)
						}
					}
				}
			}
			for _, listv := range list {
				f.WriteString(listv + "\n")
			}
			return list, nil

		}
	case "windows":
		{
			files, err := ioutil.ReadDir(fmt.Sprintf("C:/Users/%s/Desktop", h))
			if err != nil {
				return nil, fmt.Errorf("path for Dektop is not correct%v", err)
			}
			for _, n := range files {
				if strings.HasPrefix(n.Name(), "mov") {
					movies, err := ioutil.ReadDir(fmt.Sprintf("C:/Users/%s/Desktop/%s", h, n.Name()))
					if err != nil {
						return nil, fmt.Errorf("not able to read movies folder")
					}
					for _, movie := range movies {

						// fmt.Println(movie)
						nameMov := ismovie(movie.Name())
						if nameMov != "nil" {
							list = append(list, nameMov)
						}
					}
				}
			}
			for _, listv := range list {
				f.WriteString(listv + "\n")
			}
			return list, nil

		}

	default:
		{
			return nil, fmt.Errorf("only linux can be accepted as os")
		}
	}

}

func ismovie(mn string) string {
	typeArr := []string{
		".flv",
		".avi",
		".mov",
		".mp4",
		".mpg",
		".asf",
		".mkv",
	}

	for _, r := range typeArr {
		if strings.HasSuffix(mn, r) {
			return mn
		}

	}
	return "nil"
}

func MyAddr() net.Addr {

	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Errorf("cannot find net Interfaces")
	}
	for _, iface := range netInterfaces {
		if strings.HasPrefix(iface.Name, "wlp") {
			address, err := iface.Addrs()
			if err != nil {
				fmt.Errorf("cannot find interfaces address%v", err)
			}
			return address[0]

		}

		// fmt.Println(index, iface.Name, address)
	}
	return nil
}
