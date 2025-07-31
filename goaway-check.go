package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func usage() {
	fmt.Fprintf(os.Stderr, `Watch Client for Kubernetes (GOAWAY Detection)

Usage:
  GODEBUG=http2debug=2 go run watch-client.go --kubeconfig=/path/to/kubeconfig [--namespace=your-namespace] 2>&1 | tee http2-debug.log

This tool watches pods in the given namespace using client-go and prints connection resets which may be caused by HTTP/2 GOAWAY frames.
Use with GODEBUG=http2debug=2 to observe GOAWAY frames directly.

Flags:
`)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	if len(os.Args) == 1 {
		usage()
	}

	kubeconfig := flag.String("kubeconfig", "", "Path to the kubeconfig file (required)")
	namespace := flag.String("namespace", "", "Namespace to watch (optional)")
	flag.Usage = usage
	flag.Parse()

	if *kubeconfig == "" {
		fmt.Fprintln(os.Stderr, "Error: --kubeconfig is required")
		usage()
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("Failed to build kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create Kubernetes clientset: %v", err)
	}

	for {
		log.Println("🔍 Starting pod watch...")

		watcher, err := clientset.CoreV1().Pods(*namespace).Watch(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Printf("⚠️  Error starting watch: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		ch := watcher.ResultChan()

	WatchLoop:
		for {
			select {
			case event, ok := <-ch:
				if !ok {
					log.Println("\033[31m🚫 Watch channel closed — likely due to GOAWAY or transport close. Reconnecting...\033[0m")
					break WatchLoop
				}
				log.Printf("📦 Event: %s (%T)", event.Type, event.Object)
			case <-time.After(60 * time.Second):
				log.Println("⏱️  Timeout. Closing watch and reconnecting...")
				watcher.Stop()
				break WatchLoop
			}
		}
	}
}
