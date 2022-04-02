package main

import (
	"flag"
	"log"
	"os"

	"github.com/realbucksavage/jbicls-conv/conv"
	"github.com/realbucksavage/jbicls-conv/pack"
)

func main() {
	var (
		iclsSource = flag.String("icls", "", "Specify the JetBrains Theme ICLS resource to read")
		packSource = flag.String("pack", "", "Specify the directory containing go templates to be processed")
		target     = flag.String("target", "", "Specify the name of the template to render (relative to the path specified through `-pack`)")
	)
	flag.Parse()

	if *iclsSource == "" {
		log.Fatal("--source is a required flag")
	}

	src, err := os.Open(*iclsSource)
	if err != nil {
		log.Fatalf("cannot open %q: %v", *iclsSource, err)
	}
	defer func() {
		if err := src.Close(); err != nil {
			log.Printf("cannot close resource %q: %v", *iclsSource, err)
		}
	}()

	scheme, err := conv.Read(src)
	if err != nil {
		log.Fatalf("cannot parse source scheme: %v", err)
	}

	log.Printf("parsed scheme name: %s (%s)", scheme.EscapedName(), scheme.Name)

	renderer, err := pack.NewRenderer(*packSource)
	if err != nil {
		log.Fatalf("cannot create theme pack renderer: %v", err)
	}

	log.Fatal(renderer.Run(scheme, *target))
}
