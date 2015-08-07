package stages

import (
	"io"

	"github.com/DailyBurn/ratchet/data"
)

// IoReaderWriter performs both the job of a IoReader and IoWriter.
// It will read data from the given io.Reader, write the resulting data to
// the given io.Writer, and (if the write was successful) send the data
// to the next stage of processing.
//
// IoReaderWriter is composed of both a IoReader and IoWriter, so it
// supports all of the same properties and usage options.
type IoReaderWriter struct {
	IoReader
	IoWriter
}

// NewIoReaderWriter returns a new IoReaderWriter wrapping the given io.Reader object
func NewIoReaderWriter(reader io.Reader, writer io.Writer) *IoReaderWriter {
	r := IoReaderWriter{}
	r.IoReader = *NewIoReader(reader)
	r.IoWriter = *NewIoWriter(writer)
	return &r
}

func (r *IoReaderWriter) ProcessData(d data.JSON, outputChan chan data.JSON, killChan chan error) {
	r.ForEachData(killChan, func(d data.JSON) {
		r.IoWriter.ProcessData(d, outputChan, killChan)
		outputChan <- d
	})
}

func (r *IoReaderWriter) Finish(outputChan chan data.JSON, killChan chan error) {
	if outputChan != nil {
		close(outputChan)
	}
}

func (r *IoReaderWriter) String() string {
	return "IoReaderWriter"
}