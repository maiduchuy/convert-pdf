package ocrworker

const PREPROCESSOR_CONVERTPDF = "convert-pdf"

type Preprocessor interface {
	preprocess(ocrRequest *OcrRequest) error
}
