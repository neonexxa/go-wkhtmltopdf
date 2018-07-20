package pdf

/*
Implements wkhtmltopdf Go bindings. It can be used to convert HTML documents to PDFs.
The package does not use the wkhtmltopdf binary. Instead, it uses the wkhtmltox library directly.

Example
	package main

	import (
		"fmt"
		"log"
		"os"

		pdf "github.com/adrg/go-wkhtmltopdf"
	)

	func main() {
		pdf.Init()
		defer pdf.Destroy()

		// Create object from file
		object, err := pdf.NewObject("sample1.html")
		if err != nil {
			log.Fatal(err)
		}
		object.SetOption("header.center", "This is the header of the first page")
		object.SetOption("footer.right", "[page]")

		// Create object from url
		object2, err := pdf.NewObject("https://google.com")
		if err != nil {
			log.Fatal(err)
		}
		object2.SetOption("footer.right", "[page]")

		// Create object from reader
		file, err := os.Open("sample2.html")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		object3, err := pdf.NewObjectFromReader(file)
		if err != nil {
			log.Fatal(err)
		}
		object3.SetOption("footer.right", "[page]")

		// Create converter
		converter := pdf.NewConverter()
		defer converter.Destroy()

		// Add created objects to the converter
		converter.AddObject(object)
		converter.AddObject(object2)
		converter.AddObject(object3)

		// Add converter options
		converter.SetOption("documentTitle", "Sample document")
		converter.SetOption("margin.left", "10mm")
		converter.SetOption("margin.right", "10mm")
		converter.SetOption("margin.top", "10mm")
		converter.SetOption("margin.bottom", "10mm")

		// Convert the objects and get the output PDF document
		output, err := converter.Convert()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(output))
	}


Converter options
    - size.paperSize                    Paper size of the output document (e.g. A4)
    - size.width                        Width of the output document (e.g.  4cm)
    - size.height                       Height of the output document (e.g. 12in)
    - orientation                       Orientation of the output document (values: Landscape, Portrait)
    - colorMode                         Color mode to use (values: Color, Grayscale)
    - resolution                        Most likely has no effect
    - dpi                               DPI to use for printing (e.g. 80)
    - pageOffset                        Offset added to page numbers when printing headers, footers and tables of contents
    - copies                            Copies of object to print (e.g. 2)
    - collate                           Specifies if the copies should be collated (values: true, false)
    - outline                           Specifies if an outline should be generated (values: true, false)
    - outlineDepth                      The maximum depth level of the outline (e.g. 4)
    - dumpOutline                       Dump an XML representation of the outline to the specified file
    - documentTitle                     Title of the output document
    - useCompression                    Use lossless compression for the output document (values: true, false)
    - margin.top                        Size of the top margin (e.g. 2cm)
    - margin.bottom                     Size of the bottom margin (e.g. 3in)
    - margin.left                       Size of the left margin (e.g. 4mm)
    - margin.right                      Size of the right margin (e.g. 2cm)
    - imageDPI                          Maximum DPI value to use for images in the output document
    - imageQuality                      JPEG compression factor to use when producing the output document (e.g. 92)
    - load.cookieJar                    Path of file used to load and store cookies

Object options
    Load options
    - load.username                     Username to use for logging into a website (e.g. bart)
    - load.password                     Password to use for logging into a website (e.g. elbarto)
    - load.jsdelay                      Amount of time in milliseconds to wait after page load before print start (e.g. 1200)
    - load.windowStatus                 Wait until window.status is equal to this string before rendering page
    - load.zoomFactor                   Zoom of the content (e.g. 2.2)
    - load.blockLocalFileAccess         Disallow local and piped files to access other local files (values: true, false)
    - load.stopSlowScript               Stop slow running javascript (values: true, false)
    - load.debugJavascript              Forward javascript warnings and errors to the warning callback (values: true, false)
    - load.loadErrorHandling            Action to take in case of object conversion failure (values: abort, skip, ignore)
    - load.proxy                        Proxy to use when loading the object

    Header options
    - header.fontSize                   Font size to use for the header (e.g. 13)
    - header.fontName                   Name of the font to use for the header (e.g. verdana)
    - header.left                       Text to print in the left part of the header
    - header.center                     Text to print in the center part of the header
    - header.right                      Text to print in the right part of the header
    - header.line                       Specifies whether a line is printed under the header (values: true, false)
    - header.spacing                    Amount of space to put between the header and the content (e.g. 1.8)
    - header.htmlUrl                    URL for a HTML document to use for the header

    Footer options
    - footer.fontSize                   Font size to use for the footer (e.g. 13)
    - footer.fontName                   Name of the font to use for the footer (e.g. verdana)
    - footer.left                       Text to print in the left part of the footer
    - footer.center                     Text to print in the center part of the footer
    - footer.right                      Text to print in the right part of the footer
    - footer.line                       Specifies whether a line is printed above the footer (values: true, false)
    - footer.spacing                    Amount of space to put between the footer and the content (e.g. 1.8)
    - footer.htmlUrl                    URL for a HTML document to use for the footer

    Web page options
    - web.background                    Specifies if the background is printed (values: true, false)
    - web.loadImages                    Specifies if images are loaded (values: true, false)
    - web.enableJavascript              Specifies if Javascript is enabled (values: true, false)
    - web.enableIntelligentShrinking    Enable intelligent shrinking to fit more content on one page (values: true, false)
    - web.minimumFontSize               Minimum font size allowed (e.g. 9)
    - web.printMediaType                Specifies if the content is printed using the print media type instead of the screen media type (values: true, false)
    - web.defaultEncoding               Specifies the default document encoding (e.g. utf-8)
    - web.userStyleSheet                URL or path to a user specified style sheet
    - web.enablePlugins                 Enable NS plugins (values: true, false)

    Table of contents options
    - tocXsl                            If not empty this object is a table of contents object, "page" is ignored and this xsl style sheet is used to convert the outline XML into a table of contents
    - toc.useDottedLines                Use a dotted line for the table of contents (values: true, false)
    - toc.captionText                   Caption to use when creating a table of contents
    - toc.forwardLinks                  Create links from the table of contents into the actual content (values: true, false)
    - toc.backLinks                     Link back from the content to the table of contents (values: true, false)
    - toc.indentation                   Indentation to use for each table of contents level (e.g. 2em)
    - toc.fontScale                     Incremental scale down of the font for every table of contents level (e.g. 0.8)
    - includeInOutline                  Specifies if HTML sections from the HTML document are included outline and table of contents (values: true, false)
    - pagesCount                        Should we count the pages of this document, in the counter used for TOC, headers and footers?
    - useExternalLinks                  Specifies if external links in the HTML document are converted into external PDF links? (values: true, false)
    - useLocalLinks                     Specifies if internal links in the HTML document are converted into PDF references? (values: true, false)
    - produceForms                      Specifies if HTML forms are converted into PDF forms (values: true, false)

For more information see http://wkhtmltopdf.org/usage/wkhtmltopdf.txt
*/

/*
#cgo LDFLAGS: -L${SRCDIR}/wkhtmltox -lwkhtmltox
#include <stdlib.h>
#include <wkhtmltox/pdf.h>
*/
import "C"

func Init() {
	C.wkhtmltopdf_init(0)
}

func Version() string {
	return C.GoString(C.wkhtmltopdf_version())
}

func Destroy() {
	C.wkhtmltopdf_deinit()
}
