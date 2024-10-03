package integration

import (
	"io/fs"
	"sync/atomic"
	"time"

	"tobloggan/code/contracts"
	"tobloggan/code/stations"

	"github.com/mdwhatcott/pipelines"
)

type Config struct {
	Clock             contracts.Clock
	Logger            contracts.Logger
	MarkdownConverter stations.Markdown
	FileSystemReader  fs.FS
	FileSystemWriter  contracts.FSWriter
	TargetDirectory   string
	ArticleTemplate   string
	ListingTemplate   string
	BaseURL           string
}

func GenerateBlog(config Config) bool {
	started := config.Clock()
	defer func() { config.Logger.Printf("finished in %s", time.Since(started)) }()

	var (
		failure = new(atomic.Bool)
		input   = make(chan any, 1)

		scanner   = stations.NewSourceScanner(config.FileSystemReader)
		reader    = stations.NewSourceReader(config.FileSystemReader)
		parser    = stations.NewArticleParser()
		validator = contracts.Station(nil)             // stations.NewArticleValidator()
		drafts    = stations.NewDraftRemoval()         // OPTIONAL
		futures   = stations.NewFutureRemoval(started) // OPTIONAL
		markdown  = stations.NewMarkdownConverter(config.MarkdownConverter)
		listing   = contracts.Station(nil) // stations.NewListingRenderer(config.ListingTemplate)
		renderer  = contracts.Station(nil) // stations.NewArticleRenderer(config.ArticleTemplate)
		baseURL   = contracts.Station(nil) // stations.NewBaseURLRewriter(config.BaseURL)
		writer    = contracts.Station(nil) // stations.NewPageWriter(config.TargetDirectory, config.FileSystemWriter)
		reporter  = contracts.Station(nil) // stations.NewReporter(config.Logger, failure)

		pipeline = pipelines.New(input,
			pipelines.Options.Logger(config.Logger),
			pipelines.Options.StationSingleton(scanner),
			pipelines.Options.StationSingleton(reader), pipelines.Options.FanOut(5),
			pipelines.Options.StationSingleton(parser),
			pipelines.Options.StationSingleton(validator),
			pipelines.Options.StationSingleton(drafts),
			pipelines.Options.StationSingleton(futures),
			pipelines.Options.StationSingleton(markdown),
			pipelines.Options.StationSingleton(listing),
			pipelines.Options.StationSingleton(renderer),
			pipelines.Options.StationSingleton(baseURL),
			pipelines.Options.StationSingleton(writer), pipelines.Options.FanOut(5),
			pipelines.Options.StationSingleton(reporter),
		)
	)
	input <- contracts.SourceDirectory(".")
	close(input)
	pipeline.Listen()
	return !failure.Load()
}
