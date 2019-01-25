package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	verbose  *bool = flag.Bool("verbose", false, "Enabled debugging messages")
	simple   *bool = flag.Bool("simple", true, "download images required by ci-simple")
	dogfood  *bool = flag.Bool("dogfood", false, "download images required by ci-dogfood")
	advanced *bool = flag.Bool("advanced", false, "download images required by ci-advanced")
)

func main() {

	flag.Parse()

	if *verbose {
		log.SetLevel(log.DebugLevel)
	}

	// these are images that get served to clients
	imageDir := "./images"
	// these are images that are used by the server
	// and cannot be mounted by the server
	nonMountImages := "/.images/"

	// use baseURL to get known good images
	baseURL := "https://mirror.deterlab.net/rvn/tests/"
	rvnImageURL := "https://mirror.deterlab.net/rvn/img/"

	log.Debugf("testing for image directory")
	_, err := os.Stat(imageDir)
	if err != nil {
		log.Infof("Image directory not found, downloading..")
		err := os.Mkdir(imageDir, 0755)
		if err != nil {
			log.Fatalf("unable to create image directory")
		}
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Debugf("testing for non-mounted image directory")
	_, err = os.Stat(pwd + nonMountImages)
	if err != nil {
		log.Infof("Image directory not found, downloading..")
		err := os.Mkdir(pwd+nonMountImages, 0755)
		if err != nil {
			log.Fatalf("unable to create image directory")
		}
	}

	err = os.Chdir(imageDir)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if !*dogfood {
		// hardcoded for now, parse unknown json later, or load rvn
		// client.kernel
		// client.initrd
		cKernel := "4.17-generic-kernel"
		cInitrd := "4.17-generic-initramfs"
		_, err = os.Stat(cKernel)
		if err != nil {
			wget(baseURL + cKernel)
		}
		_, err = os.Stat(cInitrd)
		if err != nil {
			wget(baseURL + cInitrd)
		}
	}

	// TODO: parse the json to get the images required
	if *simple || *advanced || *dogfood {
		// client --- image ~~ not in model.js
		//cImage := []string{"ubuntu-1604-disk", "ubuntu-sled", "fedora-27-disk"}
		cImage := []string{"ubuntu-1604-disk"}
		// now we need the kernel we will load with sledc
		//cNewKernel := []string{"ubuntu-1604-kernel", "ubuntu-1804-kernel", "fedora-27-kernel"}
		cNewKernel := []string{"ubuntu-1604-kernel"}
		// now we need the initramfs we will load with sledc
		//cNewInitrd := []string{"ubuntu-1604-initramfs", "ubuntu-1804-initramfs", "fedora-27-initramfs"}
		cNewInitrd := []string{"ubuntu-1604-initramfs"}
		// server.image
		sImage := "debian-buster"
		// custom netboot image (smaller)
		netboot := "netboot"

		for _, clientImage := range cImage {
			_, err = os.Stat(clientImage)
			if err != nil {
				wget(baseURL + clientImage)
			}
		}

		for _, clientKernel := range cNewKernel {
			_, err = os.Stat(clientKernel)
			if err != nil {
				wget(baseURL + clientKernel)
			}
		}
		for _, clientInitrd := range cNewInitrd {
			_, err = os.Stat(clientInitrd)
			if err != nil {
				wget(baseURL + clientInitrd)
			}
		}

		err = os.Chdir(pwd + nonMountImages)
		if err != nil {
			log.Fatalf("%v", err)
		}
		_, err = os.Stat(sImage)
		if err != nil {
			wget(rvnImageURL + sImage)
		}
		_, err = os.Stat(netboot)
		if err != nil {
			cmd := exec.Command("dd", "if=/dev/zero", fmt.Sprintf("of=%s", netboot), "bs=1024", "count=4194304")
			log.Infof("Creating small netboot image.")
			err := cmd.Run()
			if err != nil {
				log.Warnf("Unable to create small netboot image: %v", err)
			}
		}
		err = os.Chdir(pwd)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}
}

func wget(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("unload to get url image: %s")
	}
	defer resp.Body.Close()
	fileName := filepath.Base(url)
	output, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("%s: %v", fileName, err)
	}
	_, err = io.Copy(output, resp.Body)
	if err != nil {
		log.Fatalf("unable to copy file contents")
	}
	log.Infof("%s: Downloaded", fileName)
}
