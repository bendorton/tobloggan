package integration

import (
	"io/fs"
	"sync/atomic"
	"time"

	"github.com/mdwhatcott/pipelines"
	"github.com/mdwhatcott/tobloggan/code/contracts"
	"github.com/mdwhatcott/tobloggan/code/stations"
)

type Config struct {
	Clock            contracts.Clock
	Logger           contracts.Logger
	Markdown         stations.Markdown
	FileSystemReader fs.FS
	FileSystemWriter contracts.FSWriter
	TargetDirectory  string
	ArticleTemplate  string
	ListingTemplate  string
	BaseURL          string
}

func GenerateBlog(config Config) bool {
	started := config.Clock()
	defer func() { config.Logger.Printf("finished in %s", time.Since(started)) }()

	var (
		failure = new(atomic.Bool)
		input   = make(chan any, 1)

		scanner  = stations.NewSourceScanner(config.FileSystemReader)
		reader   = stations.NewSourceReader(config.FileSystemReader)
		parser   = stations.NewArticleParser()
		drafts   = stations.NewDraftRemoval()
		futures  = stations.NewFutureRemoval(started)
		markdown = stations.NewMarkdownConverter(config.Markdown)
		// TODO: topic listings (take unlisted articles into account?)
		listing  = stations.NewListingRenderer(config.ListingTemplate) // TODO: support for unlisted articles
		renderer = stations.NewArticleRenderer(config.ArticleTemplate)
		baseURL  = stations.NewBaseURLRewriter(config.BaseURL)
		writer   = stations.NewPageWriter(config.TargetDirectory, config.FileSystemWriter)
		reporter = stations.NewReporter(config.Logger, failure)

		pipeline = pipelines.New(input,
			pipelines.Options.Logger(config.Logger),
			pipelines.Options.StationSingleton(scanner),
			pipelines.Options.StationSingleton(reader), pipelines.Options.FanOut(5),
			pipelines.Options.StationSingleton(parser),
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
