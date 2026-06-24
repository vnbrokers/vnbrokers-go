// Code generated from Finhay OpenAPI v2; DO NOT EDIT.

package dto

type GlobalNewsDetail struct {
	// Article ID (use for detail lookup)
	ID    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	// Original article URL
	URL string `json:"url,omitempty"`
	// Short summary of the article
	Description *string `json:"description,omitempty"`
	// News source/provider name
	Provider *string `json:"provider,omitempty"`
	// Publication timestamp (ISO 8601)
	PublishedAt string `json:"published_at,omitempty"`
	// Enum: forex, commodities, economic-indicators, stock-market, cryptocurrency
	Category string `json:"category,omitempty"`
	// Full article body/content
	Content *string `json:"content,omitempty"`
}

type GlobalNewsListItem struct {
	// Article ID (use for detail lookup)
	ID    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	// Original article URL
	URL string `json:"url,omitempty"`
	// Short summary of the article
	Description *string `json:"description,omitempty"`
	// News source/provider name
	Provider *string `json:"provider,omitempty"`
	// Publication timestamp (ISO 8601)
	PublishedAt string `json:"published_at,omitempty"`
	// Enum: forex, commodities, economic-indicators, stock-market, cryptocurrency
	Category string `json:"category,omitempty"`
}

type GlobalNewsPage struct {
	Results []GlobalNewsListItem `json:"results,omitempty"`
	// Number of items in the current page
	PageTotal int64 `json:"page_total,omitempty"`
	// Total number of matching articles
	Total        int64  `json:"total,omitempty"`
	CurrentPage  int64  `json:"current_page,omitempty"`
	NextPage     *int64 `json:"next_page,omitempty"`
	PreviousPage *int64 `json:"previous_page,omitempty"`
}

type IndexRecommendationReport struct {
	// Stock symbol
	Stock string `json:"stock,omitempty"`
	// Report title
	Title *string `json:"title,omitempty"`
	// Report source, defaults to `VietStock`
	Source *string `json:"source,omitempty"`
	// Recommended entry price (currently always `null`)
	RecommendationPrice *float64 `json:"recommendation_price,omitempty"`
	// Target price
	TargetPrice *float64 `json:"target_price,omitempty"`
	// Report publish date
	PublishDate *string `json:"publish_date,omitempty"`
	// Report description/summary
	Description *string `json:"description,omitempty"`
	// URL to download the full report PDF
	DownloadURL *string `json:"download_url,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}

type RecommendationReportResponse struct {
	// Latest recommendation description (from the first/most recent report)
	Recommendation        *string                     `json:"recommendation,omitempty"`
	RecommendationReports []IndexRecommendationReport `json:"recommendationReports,omitempty"`
}

type StockEventResponse struct {
	// Event ID
	ID *int64 `json:"id,omitempty"`
	// Internal path/slug of the event
	Path *string `json:"path,omitempty"`
	// Event title
	Title *string `json:"title,omitempty"`
	// Stock symbol
	Stock string `json:"stock,omitempty"`
	// Event body/content
	Body *string `json:"body,omitempty"`
	// Formatted creation date (DD/MM - HH:mm)
	CreatedDate string `json:"createdDate,omitempty"`
	// Date the corporate action takes effect
	ActionDate *string `json:"actionDate,omitempty"`
	// Ex-rights date (ngày GDKHQ)
	GdkhqDate *string `json:"gdkhqDate,omitempty"`
	// Event type code
	EventType *string `json:"eventType,omitempty"`
	// Human-readable event type name
	EventTypeName *string `json:"eventTypeName,omitempty"`
	CreatedAt     *string `json:"createdAt,omitempty"`
	UpdatedAt     *string `json:"updatedAt,omitempty"`
	// Full URL to the event detail page
	URL *string `json:"url,omitempty"`
}
