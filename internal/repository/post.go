package repository

type Post struct {
	ID      int
	Title   string
	Content string
}

type PostRepository struct {
	posts []Post
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		posts: []Post{
			{ID: 1, Title: "First Beginnings", Content: "This is the very first blog post in our series, marking the start of our journey. It introduces the purpose of the blog, what readers can expect from future updates, and sets the tone for everything that follows. We aim to share insights, stories, and experiences that will inform and engage our audience while building a community."},
			{ID: 2, Title: "Second Steps Forward", Content: "The second blog post builds upon the foundation of the first, taking the next step in our narrative. Here we expand on earlier topics, offer more detailed examples, and provide new perspectives that help deepen understanding. It is designed to move the conversation forward and prepare readers for more complex discussions ahead."},
			{ID: 3, Title: "Third Insights Unveiled", Content: "In our third blog post, we begin to explore topics in greater detail, revealing insights that were only briefly mentioned before. This entry focuses on connecting ideas, bridging gaps in knowledge, and adding layers of context. Readers will find more actionable details here, with clear explanations and relevant scenarios for reference."},
			{ID: 4, Title: "Fourth Layer of Depth", Content: "The fourth post in our series continues the journey by going deeper into previously discussed themes. We revisit earlier points with new context, share updated information, and address feedback from our audience. This approach ensures that our content remains relevant, informative, and engaging for both new and returning readers."},
			{ID: 5, Title: "Fifth Perspective Shift", Content: "Our fifth blog post introduces a change in perspective, encouraging readers to view familiar concepts from a fresh angle. This post challenges assumptions, provides alternative viewpoints, and examines potential implications of different approaches. It is designed to provoke thought, encourage discussion, and inspire readers to reconsider what they know."},
			{ID: 6, Title: "Sixth Strategic Overview", Content: "The sixth entry shifts focus toward a broader, strategic overview of the subjects we’ve been exploring. It connects the dots between individual pieces of content, showing how they fit into a bigger picture. This post offers guidance on how readers can apply our shared ideas in real-world scenarios while avoiding common pitfalls."},
			{ID: 7, Title: "Seventh Detailed Analysis", Content: "The seventh blog post offers an in-depth analysis of a specific topic within our broader theme. It uses data, examples, and careful reasoning to explain how certain concepts work in practice. This detailed examination helps readers build stronger understanding, identify key patterns, and apply lessons more effectively in their own contexts."},
			{ID: 8, Title: "Eighth Expansion Phase", Content: "Our eighth blog entry marks an expansion phase, where we broaden the scope of our discussions. We introduce related topics that extend beyond the original boundaries, allowing readers to see connections they might have missed before. This post also serves as a bridge to upcoming content, creating a sense of anticipation."},
			{ID: 9, Title: "Ninth Practical Application", Content: "In the ninth blog post, the focus shifts to practical application. We take theories and ideas discussed in earlier entries and show how they can be implemented in everyday situations. By combining clear instructions with real-world examples, we aim to help readers put knowledge into action and achieve tangible results."},
			{ID: 10, Title: "Tenth Reflection Point", Content: "The tenth entry acts as a reflection point, summarizing key lessons learned so far and assessing their impact. This post looks back at previous entries, identifies recurring themes, and evaluates progress. Readers are encouraged to think critically about how the information applies to their own experiences, helping to solidify learning."},
			{ID: 11, Title: "Eleventh Future Vision", Content: "In our eleventh blog post, we turn our attention toward the future. We explore emerging trends, forecast potential developments, and discuss how readers can prepare for what’s ahead. This forward-looking perspective aims to equip our audience with the knowledge and mindset needed to adapt to change and seize new opportunities."},
			{ID: 12, Title: "Twelfth Grand Conclusion", Content: "The twelfth and final post in this series serves as a grand conclusion, tying together every concept, lesson, and insight shared so far. It reinforces core messages, answers lingering questions, and leaves readers with a clear sense of direction for continued exploration. This post closes one chapter while opening the door to future journeys."},
			{ID: 13, Title: "Thirteenth Reader Engagement", Content: "This post focuses on engaging with our audience more directly. We explore methods to gather feedback, encourage dialogue, and integrate reader contributions into future content. This approach strengthens community bonds and ensures our topics remain aligned with audience interests."},
			{ID: 14, Title: "Fourteenth Behind the Scenes", Content: "The fourteenth entry takes readers behind the scenes, sharing how ideas are developed, researched, and refined before publication. We discuss challenges faced during the process and how they are overcome, providing transparency and deepening reader connection."},
			{ID: 15, Title: "Fifteenth New Horizons", Content: "The fifteenth blog post opens a new chapter by introducing themes and topics that will define the next stage of our journey. It sets the stage for exploration beyond the current series, inviting readers to join us as we venture into fresh territories with renewed curiosity."},
		},
	}
}

func (r *PostRepository) GetAllPosts() []Post {
	return r.posts
}

func (r *PostRepository) GetPostByID(id int) *Post {
	for _, p := range r.posts {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
